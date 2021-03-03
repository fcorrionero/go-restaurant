package mongo

import (
	"context"
	"fmt"
	"github.com/fcorrionero/go-restaurant/domain"
	"github.com/google/uuid"
	"go.mongodb.org/mongo-driver/bson"
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

func (r DishesRepository) FindDishesByAllergen(allergenId uuid.UUID) []domain.DishAggregate {
	var dishes []domain.DishAggregate

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

	filter := bson.M{"name": name}
	err := r.collection.FindOne(r.context, filter).Decode(&result)
	if err != nil {
		log.Fatal(err)
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
