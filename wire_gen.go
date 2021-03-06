// Code generated by Wire. DO NOT EDIT.

//go:generate go run github.com/google/wire/cmd/wire
//+build !wireinject

package main

import (
	"github.com/fcorrionero/go-restaurant/application/query/find_dish_by_id"
	"github.com/fcorrionero/go-restaurant/application/query/find_dish_by_name"
	"github.com/fcorrionero/go-restaurant/application/query/find_dishes_by_allergen"
	"github.com/fcorrionero/go-restaurant/domain"
	"github.com/fcorrionero/go-restaurant/infrastructure/persistence/mongo"
	"github.com/fcorrionero/go-restaurant/infrastructure/ui/dishes_http"
)

// Injectors from wire.go:

func InitializeDishesRepository() domain.DishesRepository {
	dishesRepository := NewDishesRepository()
	return dishesRepository
}

func InitializeDishesHttpController(dishesRepository domain.DishesRepository) dishes_http.DishesHttpController {
	queryHandler := find_dish_by_id.New(dishesRepository)
	find_dish_by_nameQueryHandler := find_dish_by_name.New(dishesRepository)
	find_dishes_by_allergenQueryHandler := find_dishes_by_allergen.New(dishesRepository)
	dishesHttpController := dishes_http.NewDishesHttpController(queryHandler, find_dish_by_nameQueryHandler, find_dishes_by_allergenQueryHandler)
	return dishesHttpController
}

// wire.go:

func NewDishesRepository() domain.DishesRepository {
	return mongo.New("root", "example", "0.0.0.0", "27017", "go-restaurant")
}
