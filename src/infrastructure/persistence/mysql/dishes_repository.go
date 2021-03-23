package mysql

import (
	"github.com/fcorrionero/go-restaurant/src/domain"
	"github.com/fcorrionero/go-restaurant/src/infrastructure/persistence/mysql/models"
	"github.com/google/uuid"
	"gorm.io/gorm"
	"log"
	"strings"
)

func NewDishesRepository(db *gorm.DB) DishesRepository {
	return DishesRepository{
		table: "dishes",
		db:    db,
	}
}

type DishesRepository struct {
	table string
	db    *gorm.DB
}

func (r DishesRepository) FindDishesByAllergenId(allergenId uuid.UUID) []*domain.DishAggregate {
	var a models.Allergen
	baId, err := allergenId.MarshalBinary()
	if err != nil {
		log.Println(err.Error())
		return []*domain.DishAggregate{}
	}
	r.db.Preload("Ingredients.Dishes").First(&a, "id LIKE ?", baId)
	var dIds []string
	for _, i := range a.Ingredients {
		for _, d := range i.Dishes {
			dIds = append(dIds, d.IdUuid)
		}
	}

	results := r.dishesAggFromDishStringIds(dIds)

	return results
}

func (r DishesRepository) dishesAggFromDishStringIds(dIds []string) []*domain.DishAggregate {
	var dishes []models.Dish
	r.db.Preload("Ingredients.Allergens").Find(&dishes, "id_uuid IN ?", dIds)
	var results []*domain.DishAggregate
	for _, d := range dishes {
		dAgg := r.dishAggFromModel(d)
		results = append(results, &dAgg)
	}
	return results
}

func (r DishesRepository) FindDishById(dishId uuid.UUID) *domain.DishAggregate {
	var result domain.DishAggregate
	var d models.Dish
	bId, err := dishId.MarshalBinary()
	if err != nil {
		log.Println(err.Error())
		return &result
	}

	r.db.Preload("Ingredients.Allergens").First(&d, "id = ?", bId)

	result = r.dishAggFromModel(d)
	return &result
}

func (r DishesRepository) FindDishByName(name string) *domain.DishAggregate {
	var result domain.DishAggregate
	var d models.Dish

	name = strings.ToUpper(name)
	r.db.Preload("Ingredients.Allergens").First(&d, "UPPER(dish_name) LIKE ?", "%"+name+"%")

	result = r.dishAggFromModel(d)
	return &result
}

func (r DishesRepository) FindDishesByAllergen(allergen string) []*domain.DishAggregate {
	var a models.Allergen
	allergen = strings.ToUpper(allergen)
	r.db.Preload("Ingredients.Dishes").First(&a, "UPPER(allergen_name) LIKE ?", "%"+allergen+"%")
	var dIds []string

	for _, i := range a.Ingredients {
		for _, d := range i.Dishes {
			dIds = append(dIds, d.IdUuid)
		}
	}

	results := r.dishesAggFromDishStringIds(dIds)

	return results
}

func (r DishesRepository) SaveDish(aggregate *domain.DishAggregate) {
	iId, _ := aggregate.Id.MarshalBinary()
	aModel := models.Dish{
		Id:          iId,
		IdUuid:      aggregate.Id.String(),
		DishName:    aggregate.Name,
		CreatedAt:   aggregate.DateTime,
		Ingredients: nil,
	}

	for _, i := range aggregate.Ingredients {
		biId, _ := i.Id.MarshalBinary()
		iModel := models.Ingredient{
			Id:             biId,
			IdUuid:         i.Id.String(),
			IngredientName: i.Name,
			Allergens:      nil,
		}
		aModel.Ingredients = append(aModel.Ingredients, iModel)
	}

	r.db.Create(&aModel)
}

func (r DishesRepository) dishAggFromModel(d models.Dish) domain.DishAggregate {
	var result domain.DishAggregate

	dId, err := uuid.Parse(d.IdUuid)
	if err != nil {
		return result
	}

	result.Name = d.DishName
	result.Id = dId
	result.DateTime = d.CreatedAt
	for _, i := range d.Ingredients {
		id, _ := uuid.Parse(i.IdUuid)
		ing := domain.Ingredient{
			Id:        id,
			Allergens: nil,
			Name:      i.IngredientName,
		}
		for _, a := range i.Allergens {
			aId, _ := uuid.Parse(a.IdUuid)
			alg := domain.Allergen{
				Id:   aId,
				Name: a.AllergenName,
			}
			ing.Allergens = append(ing.Allergens, alg)
		}
		result.Ingredients = append(result.Ingredients, ing)
	}

	return result
}
