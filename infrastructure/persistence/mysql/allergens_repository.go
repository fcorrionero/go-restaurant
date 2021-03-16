package mysql

import (
	"github.com/fcorrionero/go-restaurant/domain"
	"github.com/fcorrionero/go-restaurant/infrastructure/persistence/mysql/models"
	"github.com/google/uuid"
	"gorm.io/gorm"
	"log"
	"strings"
)

type AllergensRepository struct {
	table string
	db    *gorm.DB
}

func NewAllergensRepository(db *gorm.DB) AllergensRepository {
	return AllergensRepository{
		table: "allergens",
		db:    db,
	}
}

func (r AllergensRepository) FindByName(name string) *domain.Allergen {
	var result *domain.Allergen
	var m models.Allergen

	name = strings.ToUpper(name)
	res := r.db.First(&m, "UPPER(allergen_name) LIKE ?", "%"+name+"%")
	if res.Error != nil {
		log.Println(res.Error)
		return result
	}

	result = r.allergenAggFromModel(m)
	return result
}

func (r AllergensRepository) FindAll() []*domain.Allergen {
	var results []*domain.Allergen
	var algs []models.Allergen

	res := r.db.Find(&algs)
	if res.Error != nil {
		log.Println(res.Error)
		return results
	}
	for _, a := range algs {
		results = append(results, r.allergenAggFromModel(a))
	}

	return results
}

func (r AllergensRepository) FindById(id uuid.UUID) *domain.Allergen {
	var result *domain.Allergen
	var m models.Allergen
	bId, err := id.MarshalBinary()
	if err != nil {
		log.Println(err.Error())
		return result
	}
	r.db.First(&m, "id = ?", bId)
	result = r.allergenAggFromModel(m)
	return result
}

func (r AllergensRepository) Save(allergen *domain.Allergen) {

	aId, _ := allergen.Id.MarshalBinary()
	aModel := models.Allergen{
		Id:           aId,
		IdUuid:       allergen.Id.String(),
		AllergenName: allergen.Name,
		Ingredients:  nil,
	}

	r.db.Create(&aModel)
	if len(aModel.IdUuid) == 0 {
		log.Println("Error inserting allergen")
		return
	}
}

func (r AllergensRepository) allergenAggFromModel(a models.Allergen) *domain.Allergen {
	var allergen domain.Allergen
	allergen.Id, _ = uuid.Parse(a.IdUuid)
	allergen.Name = a.AllergenName

	return &allergen
}
