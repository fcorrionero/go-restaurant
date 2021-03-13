package mysql

import (
	"database/sql"
	"github.com/fcorrionero/go-restaurant/domain"
)

func NewIngredientsRepository(db *sql.DB) IngredientsRepository {
	return IngredientsRepository{
		table: "ingredients",
		db:    db,
	}
}

type IngredientsRepository struct {
	table string
	db    *sql.DB
}

func (r IngredientsRepository) FindByName(name string) *domain.Ingredient {
	panic("implement me")
}

func (r IngredientsRepository) FindAllByIds(ids []int) []*domain.Ingredient {
	panic("implement me")
}

func (r IngredientsRepository) Save(ingredient domain.Ingredient) {
	panic("implement me")
}
