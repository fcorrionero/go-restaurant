package mysql

import (
	"database/sql"
	"fmt"
	"github.com/fcorrionero/go-restaurant/domain"
	"github.com/google/uuid"
	"log"
)

func NewIngredientsRepository(db *sql.DB) IngredientsRepository {
	return IngredientsRepository{
		table:     "ingredients",
		relations: "allergens",
		relTable:  "ingredients_allergens",
		db:        db,
	}
}

type IngredientsRepository struct {
	table     string
	relations string
	relTable  string
	db        *sql.DB
}

func (r IngredientsRepository) FindByName(name string) *domain.Ingredient {
	panic("implement me")
}

func (r IngredientsRepository) FindById(id uuid.UUID) *domain.Ingredient {
	panic("implement me")
}

func (r IngredientsRepository) FindAll() []*domain.Ingredient {
	panic("implement me")
}

func (r IngredientsRepository) FindAllByAllergen(aId uuid.UUID) []*domain.Ingredient {
	panic("implement me")
}

func (r IngredientsRepository) Save(ingredient *domain.Ingredient) {
	sqlStmt := fmt.Sprintf("INSERT INTO %s VALUES( ?, ?, ? )", r.table)
	stmtIns, err := r.db.Prepare(sqlStmt) // ? = placeholder
	if err != nil {
		log.Println(err.Error())
		return
	}
	defer func() {
		err := stmtIns.Close()
		if nil != err {
			log.Println(err.Error())
		}
	}()

	bId, _ := ingredient.Id.MarshalBinary()
	sId := ingredient.Id.String()
	_, err = stmtIns.Exec(bId, sId, ingredient.Name)
	if err != nil {
		log.Println(err.Error())
		return
	}
	aRepo := AllergensRepository{
		table: r.relations,
		db:    r.db,
	}
	for _, a := range ingredient.Allergens {
		r.saveRelations(ingredient, aRepo, a)
	}
}

func (r IngredientsRepository) saveRelations(ingredient *domain.Ingredient, aRepo AllergensRepository, a domain.Allergen) {
	allergen := aRepo.FindById(a.Id)
	if len(allergen.Name) == 0 {
		aRepo.Save(&a)
	}
	relSql := fmt.Sprintf("INSERT INTO %s VALUES( ?, ? )", r.relTable)
	stmtIns, _ := r.db.Prepare(relSql)
	bIId, _ := ingredient.Id.MarshalBinary()
	bAId, _ := a.Id.MarshalBinary()
	_, err := stmtIns.Exec(bAId, bIId)
	if err != nil {
		log.Println(err.Error())
	}
	stmtIns.Close()
}
