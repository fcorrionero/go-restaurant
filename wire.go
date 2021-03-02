//+build wireinject

package main

import (
	"github.com/fcorrionero/go-restaurant/domain"
	"github.com/fcorrionero/go-restaurant/infrastructure/persistence/mongo"
	"github.com/google/wire"
)

func InitializeDishesRepository() domain.DishesRepository {
	wire.Build(NewDishesRepository)
	return mongo.DishesRepository{}
}

func NewDishesRepository() domain.DishesRepository {
	return mongo.New("root", "example", "0.0.0.0", "27017", "go-restaurant")
}
