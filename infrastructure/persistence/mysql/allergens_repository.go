package mysql

import (
	"database/sql"
	"github.com/fcorrionero/go-restaurant/domain"
)

type AllergensRepository struct {
	table string
	db    *sql.DB
}

func NewAllergensRepository(table string, db *sql.DB) AllergensRepository {
	return AllergensRepository{
		table: table,
		db:    db,
	}
}

func (r AllergensRepository) FindByName(name string) *domain.Allergen {
	panic("implement me")
}

func (r AllergensRepository) FindAllById(ids []int) []*domain.Allergen {
	panic("implement me")
}

func (r AllergensRepository) Save(allergen domain.Allergen) {
	panic("implement me")
}
