package main

import (
	"fmt"
	"github.com/fcorrionero/go-restaurant/domain"
	"github.com/google/uuid"
	"time"
)

func main() {
	var ingredients []domain.Ingredient

	ingredient := domain.Ingredient{
		Id:        uuid.UUID{},
		Allergens: nil,
		Name:      "Arroz",
	}

	ingredients = append(ingredients, ingredient)

	dish := domain.DishAggregate{
		Id:          uuid.New(),
		Ingredients: ingredients,
		DateTime:    time.Now(),
	}
	fmt.Println("RESTAURANT APP : " + dish.String())
}
