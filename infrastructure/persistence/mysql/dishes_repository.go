package mysql

import (
	"github.com/fcorrionero/go-restaurant/domain"
	"github.com/fcorrionero/go-restaurant/infrastructure/persistence/mysql/models"
	"github.com/google/uuid"
	"gorm.io/gorm"
	"log"
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
	return []*domain.DishAggregate{}
}

func (r DishesRepository) FindDishById(dishId uuid.UUID) *domain.DishAggregate {
	var result *domain.DishAggregate
	var d models.Dish
	bId, err := dishId.MarshalBinary()
	if err != nil {
		log.Println(err.Error())
		return result
	}

	r.db.Preload("Ingredients").Preload("Allergens").First(&d, "id = ?", bId)

	panic("implement me")
}

func (r DishesRepository) FindDishByName(name string) *domain.DishAggregate {
	panic("Implement me")
}

func (r DishesRepository) FindDishesByAllergen(allergen string) []*domain.DishAggregate {
	return []*domain.DishAggregate{}
}

func (r DishesRepository) SaveDish(aggregate *domain.DishAggregate) {
	panic("implement me")
}
