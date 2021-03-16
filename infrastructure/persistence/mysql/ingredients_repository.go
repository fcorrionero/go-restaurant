package mysql

import (
	"fmt"
	"github.com/fcorrionero/go-restaurant/domain"
	"github.com/fcorrionero/go-restaurant/infrastructure/persistence/mysql/models"
	"github.com/google/uuid"
	"gorm.io/gorm"
	"log"
	"strings"
)

func NewIngredientsRepository(db *gorm.DB) IngredientsRepository {
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
	db        *gorm.DB
}

func (r IngredientsRepository) FindByName(name string) *domain.Ingredient {
	var result *domain.Ingredient
	var i models.Ingredient

	name = strings.ToUpper(name)
	r.db.Preload("Allergens").First(&i, "UPPER(ingredient_name) LIKE ?", "%"+name+"%")
	result = r.ingredientAggFromModel(i)
	return result
}

func (r IngredientsRepository) FindById(id uuid.UUID) *domain.Ingredient {
	var result *domain.Ingredient
	var i models.Ingredient
	bId, err := id.MarshalBinary()
	if err != nil {
		log.Println(err.Error())
		return result
	}

	r.db.Preload("Allergens").First(&i, "id = ?", bId)
	result = r.ingredientAggFromModel(i)
	return result
}

func (r IngredientsRepository) ingredientAggFromModel(i models.Ingredient) *domain.Ingredient {
	var ingredient domain.Ingredient
	ingredient.Id, _ = uuid.Parse(i.IdUuid)
	ingredient.Name = i.IngredientName
	for _, a := range i.Allergens {
		aId, _ := uuid.Parse(a.IdUuid)
		allergen := domain.Allergen{
			Id:   aId,
			Name: a.AllergenName,
		}
		ingredient.Allergens = append(ingredient.Allergens, allergen)
	}
	return &ingredient
}

func (r IngredientsRepository) FindAll() []*domain.Ingredient {
	// Without allergens for performance
	var results []*domain.Ingredient

	var ings []models.Ingredient
	res := r.db.Find(&ings)
	if res.Error != nil {
		log.Println(res.Error)
		return results
	}
	for _, i := range ings {
		results = append(results, r.ingredientAggFromModel(i))
	}

	return results
}

func (r IngredientsRepository) FindAllByAllergen(aId uuid.UUID) []*domain.Ingredient {
	var results []*domain.Ingredient
	bId, _ := aId.MarshalBinary()

	var a models.Allergen
	r.db.Preload("Ingredients").First(&a, "id = ?", bId)

	var ids []string
	for _, i := range a.Ingredients {
		ids = append(ids, i.IdUuid)
	}

	var ingredients []models.Ingredient
	r.db.Preload("Allergens").Where("id_uuid in ?", ids).Find(&ingredients)
	for _, i := range ingredients {
		results = append(results, r.ingredientAggFromModel(i))
	}

	return results
}

func (r IngredientsRepository) Save(ingredient *domain.Ingredient) {

	iId, _ := ingredient.Id.MarshalBinary()
	iModel := models.Ingredient{
		Id:             iId,
		IdUuid:         ingredient.Id.String(),
		IngredientName: ingredient.Name,
		Allergens:      nil,
	}
	r.db.Create(&iModel)
	if len(iModel.IdUuid) == 0 {
		log.Println("Error inserting ingredient")
		return
	}

	var aModel models.Allergen
	for _, a := range ingredient.Allergens {
		bId, _ := a.Id.MarshalBinary()
		r.db.First(&aModel, "id = ?", bId)
		if len(aModel.IdUuid) == 0 {
			aModel.IdUuid = a.Id.String()
			aModel.Id, _ = a.Id.MarshalBinary()
			aModel.AllergenName = a.Name
			r.db.Create(&aModel)
			if len(aModel.IdUuid) == 0 {
				log.Println("Error inserting allergen")
				return
			}
		}
		relSql := fmt.Sprintf("INSERT INTO %s VALUES( ?, ? )", r.relTable)
		r.db.Exec(relSql, bId, iId)
	}
}
