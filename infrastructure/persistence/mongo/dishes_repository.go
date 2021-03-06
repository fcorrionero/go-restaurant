package mongo

import (
	"context"
	"fmt"
	"github.com/fcorrionero/go-restaurant/domain"
	"github.com/google/uuid"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
)

type DishesRepository struct {
	client     *mongo.Client
	context    context.Context
	collection *mongo.Collection
}

func New(userName string, password string, host string, port string, database string) DishesRepository {
	dishesRepository := DishesRepository{}
	dishesRepository.connect(userName, password, host, port)
	dishesRepository.collection = dishesRepository.client.Database(database).Collection("dishes")
	return dishesRepository
}

func (r *DishesRepository) connect(userName string, password string, host string, port string) {
	credential := options.Credential{
		Username: userName,
		Password: password,
	}
	client, err := mongo.NewClient(options.Client().ApplyURI(fmt.Sprintf("mongodb://%s:%s", host, port)).SetAuth(credential))
	if err != nil {
		log.Fatal(err)
		return
	}
	r.client = client

	ctx := context.Background()
	err = client.Connect(ctx)
	if err != nil {
		log.Fatal(err)
		return
	}

	r.context = ctx

	return
}

func (r DishesRepository) FindDishesByAllergenId(allergenId uuid.UUID) []*domain.DishAggregate {
	var dishes []*domain.DishAggregate

	return dishes
}

func (r DishesRepository) FindDishById(dishId uuid.UUID) domain.DishAggregate {
	result := domain.DishAggregate{}

	filter := bson.M{"id": dishId}
	err := r.collection.FindOne(r.context, filter).Decode(&result)
	if err != nil {
		log.Fatal(err)
	}
	return result
}

func (r DishesRepository) FindDishByName(name string) domain.DishAggregate {
	result := domain.DishAggregate{}

	// Search by name case insensitive with a regex expression
	key := "name"
	regex := bson.M{"$regex": primitive.Regex{Pattern: ".*" + name + ".*", Options: "i"}}
	filter2 := bson.M{key: regex}

	err := r.collection.FindOne(r.context, filter2).Decode(&result)
	if err != nil {
		log.Println(err.Error() + " " + name)
		log.Println(filter2)
	}
	return result
}

func (r DishesRepository) SaveDish(aggregate domain.DishAggregate) {
	if r.client == nil {
		return
	}
	_, err := r.collection.InsertOne(r.context, aggregate)
	if err != nil {
		panic("This dish exists ")
	}
}

func (r DishesRepository) ConfigureDB() {
	_, _ = r.collection.Indexes().CreateOne(
		context.Background(),
		mongo.IndexModel{
			Keys:    bson.D{{Key: "id", Value: 1}},
			Options: options.Index().SetUnique(true),
		})
}

func (r DishesRepository) FindDishesByAllergen(allergen string) []*domain.DishAggregate {
	var dishes []*domain.DishAggregate

	regex := bson.M{"$regex": primitive.Regex{Pattern: ".*" + allergen + ".*", Options: "i"}}
	filter := bson.M{"ingredients.allergens.name": regex}
	cursor, err := r.collection.Find(r.context, filter)
	if err != nil {
		log.Println(err)
		return dishes
	}
	for cursor.Next(context.TODO()) {
		// create a value into which the single document can be decoded
		var d domain.DishAggregate
		err := cursor.Decode(&d)
		if err != nil {
			log.Println(err)
			return dishes
		}
		dishes = append(dishes, &d)
	}
	log.Println(filter)
	return dishes
}
