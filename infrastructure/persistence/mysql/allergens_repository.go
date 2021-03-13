package mysql

import (
	"database/sql"
	"fmt"
	"github.com/fcorrionero/go-restaurant/domain"
	"github.com/google/uuid"
	"log"
)

type AllergensRepository struct {
	table string
	db    *sql.DB
}

func NewAllergensRepository(db *sql.DB) AllergensRepository {
	return AllergensRepository{
		table: "allergens",
		db:    db,
	}
}

func (r AllergensRepository) FindByName(name string) *domain.Allergen {
	result := domain.Allergen{}
	sqlStmt := fmt.Sprintf("SELECT * FROM %s WHERE allergen_name = ?", r.table)
	stmtOut, err := r.db.Prepare(sqlStmt)
	if err != nil {
		log.Println(err.Error())
		return &result
	}
	defer func() {
		err := stmtOut.Close()
		if nil != err {
			log.Println(err.Error())
		}
	}()

	var sId, aName string
	var bId []byte
	err = stmtOut.QueryRow(name).Scan(&bId, &sId, &aName)
	if err != nil {
		log.Println(err.Error())
		return &result
	}
	result.Id, err = uuid.Parse(sId)
	result.Name = aName
	if nil != err {
		log.Println(err.Error())
	}
	return &result
}

func (r AllergensRepository) FindAllById(ids []int) []*domain.Allergen {
	panic("implement me")
}

func (r AllergensRepository) Save(allergen domain.Allergen) {
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

	bId, _ := allergen.Id.MarshalBinary()
	sId := allergen.Id.String()
	_, err = stmtIns.Exec(bId, sId, allergen.Name)
	if err != nil {
		log.Println(err.Error())
		return
	}
}
