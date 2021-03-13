package mysql

import (
	"database/sql"
	"github.com/fcorrionero/go-restaurant/domain"
	"github.com/google/uuid"
)

func NewDishesRepository(db *sql.DB) DishesRepository {
	return DishesRepository{
		table: "dishes",
		db:    db,
	}
}

type DishesRepository struct {
	table string
	db    *sql.DB
}

func (r DishesRepository) FindDishesByAllergenId(allergenId uuid.UUID) []*domain.DishAggregate {
	return []*domain.DishAggregate{}
}

func (r DishesRepository) FindDishById(dishId uuid.UUID) domain.DishAggregate {
	return domain.DishAggregate{}
}

func (r DishesRepository) FindDishByName(name string) domain.DishAggregate {
	return domain.DishAggregate{}
}

func (r DishesRepository) FindDishesByAllergen(allergen string) []*domain.DishAggregate {
	return []*domain.DishAggregate{}
}

func (r DishesRepository) SaveDish(aggregate domain.DishAggregate) {
	panic("implement me")
}
