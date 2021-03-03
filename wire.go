//+build wireinject

package main

import (
	"github.com/fcorrionero/go-restaurant/application/query/find_dish_by_id"
	"github.com/fcorrionero/go-restaurant/application/query/find_dish_by_name"
	"github.com/fcorrionero/go-restaurant/domain"
	"github.com/fcorrionero/go-restaurant/infrastructure/persistence/mongo"
	"github.com/fcorrionero/go-restaurant/infrastructure/ui/dishes_http"
	"github.com/google/wire"
)

func InitializeDishesRepository() domain.DishesRepository {
	wire.Build(NewDishesRepository)
	return mongo.DishesRepository{}
}

func InitializeDishesHttpController(dishesRepository domain.DishesRepository) dishes_http.DishesHttpController {
	wire.Build(dishes_http.NewDishesHttpController, find_dish_by_id.New, find_dish_by_name.New)
	return dishes_http.DishesHttpController{}
}

func NewDishesRepository() domain.DishesRepository {
	return mongo.New("root", "example", "0.0.0.0", "27017", "go-restaurant")
}
