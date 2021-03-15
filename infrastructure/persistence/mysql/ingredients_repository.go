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
	var results []*domain.Ingredient
	relSql := fmt.Sprintf(
		"SELECT " +
			"i.id_uuid, i.ingredient_name, a.id_uuid, a.allergen_name " +
			"FROM ingredients i " +
			"INNER JOIN ingredients_allergens ia on i.id = ia.ingredient_id " +
			"INNER JOIN allergens a on ia.allergen_id = a.id " +
			"WHERE ia.allergen_id = ? ORDER BY i.id_uuid")
	sqlStmt, err := r.db.Prepare(relSql)
	if nil != err {
		log.Println(err.Error())
		return results
	}
	defer func() {
		err := sqlStmt.Close()
		if nil != err {
			log.Println(err.Error())
		}
	}()
	bId, err := aId.MarshalBinary()
	if err != nil {
		log.Println(err.Error())
		return results
	}
	rows, err := sqlStmt.Query(bId)
	if err != nil {
		log.Println(err.Error())
		return results
	}
	var alId string
	var iId string
	var iName string
	var aName string
	var ingredient domain.Ingredient
	for rows.Next() {
		err := rows.Scan(&iId, &iName, &alId, &aName)
		if err != nil {
			log.Println(err.Error())
			return results
		}
		fiId, _ := uuid.Parse(iId)
		if fiId != ingredient.Id {
			if len(ingredient.Name) > 0 {
				results = append(results, &ingredient)
			}
			ingredient = domain.Ingredient{
				Id:        fiId,
				Allergens: nil,
				Name:      iName,
			}
		}
		faId, _ := uuid.Parse(alId)
		allergen := domain.Allergen{
			Id:   faId,
			Name: aName,
		}
		ingredient.AddAllergen(allergen)
	}
	results = append(results, &ingredient)

	return results
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
