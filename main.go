package main

import (
	"net/http"
)

func main() {
	repo := InitializeDishesRepository()
	dishesController := InitializeDishesHttpController(repo)
	http.HandleFunc("/search", dishesController.Search)
	err := http.ListenAndServe(":8090", nil)
	if err != nil {
		panic(err)
	}
}
