package main

import (
	"net/http"
)

func main() {
	repo := InitializeDishesRepository()
	dishesController := InitializeDishesHttpController(repo)
	crudController := InitializeCrudHttpController()
	http.HandleFunc("/search", dishesController.Search)
	http.HandleFunc("/add-allergen", crudController.AddAllergen)
	err := http.ListenAndServe(":8090", nil)
	if err != nil {
		panic(err)
	}
}
