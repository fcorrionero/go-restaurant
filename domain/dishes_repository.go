package domain

import "github.com/google/uuid"

type DishesRepository interface {
	FindDishesByAllergenId(allergenId uuid.UUID) []*DishAggregate
	FindDishById(dishId uuid.UUID) DishAggregate
	FindDishByName(name string) DishAggregate
	FindDishesByAllergen(allergen string) []*DishAggregate
	SaveDish(aggregate DishAggregate)
}
