package main

import (
	"net/http"
)

func main() {
	repo := InitializeDishesRepository()
	dishesController := InitializeDishesHttpController(repo)
	http.HandleFunc("/name", dishesController.ByName)
	http.HandleFunc("/id", dishesController.ById)
	http.HandleFunc("/allergen", dishesController.ByAllergen)
	err := http.ListenAndServe(":8090", nil)
	if err != nil {
		panic(err)
	}
}
