package main

import (
	"github.com/fcorrionero/go-restaurant/domain"
	"github.com/google/uuid"
	"net/http"
)

func main() {

	//id := uuid.New()
	//
	//fmt.Println(id)
	//
	//fmt.Println(id.MarshalBinary())
	//
	//binary, _ := id.MarshalBinary()
	//
	//fmt.Println(uuid.FromBytes(binary))

	db := StartMysqlDB()
	defer db.Close()

	iRepo := NewMysqlIngredientsRepository(db)
	aId, _ := uuid.Parse("64d786cf-9a36-4185-8569-e1336cc40fd1")
	allergen := domain.Allergen{
		Id:   aId,
		Name: "Crustáceos",
	}
	ingredient := domain.Ingredient{
		Id:        uuid.New(),
		Allergens: []domain.Allergen{allergen},
		Name:      "Gambas",
	}

	iRepo.Save(&ingredient)

	//aRepo := NewMysqlAllergensRepository(db)
	//allergen := aRepo.FindByName("Gluten")
	//id, _ := uuid.Parse("f329599a-5f9e-4930-86df-0468dd645372")
	//allergen := aRepo.FindById(id)
	//log.Println(allergen.String())

	repo := InitializeDishesRepository()
	dishesController := InitializeDishesHttpController(repo)
	http.HandleFunc("/search", dishesController.Search)
	err := http.ListenAndServe(":8090", nil)
	if err != nil {
		panic(err)
	}
}
