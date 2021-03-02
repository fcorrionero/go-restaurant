package domain

import "github.com/google/uuid"

type DishesRepository interface {
	FindDishesByAllergen(allergenId uuid.UUID) []DishAggregate
	FindDishById(dishId uuid.UUID) DishAggregate
	FindDishByName(name string) DishAggregate
	SaveDish(aggregate DishAggregate)
}
