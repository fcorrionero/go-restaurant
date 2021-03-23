package domain

import (
	"github.com/fcorrionero/go-restaurant/src/domain"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

func TestIngredientShouldBeAddedToDish(t *testing.T) {
	dish := domain.DishAggregate{
		Id:          uuid.UUID{},
		Ingredients: nil,
		DateTime:    time.Time{},
	}

	if len(dish.Ingredients) > 0 {
		t.Error("Ingredients should be empty")
	}

	ingredient := domain.Ingredient{
		Id:        uuid.New(),
		Allergens: nil,
		Name:      "Curcuma",
	}

	dish.AddIngredient(ingredient)
	if len(dish.Ingredients) != 1 {
		t.Error("Should be only 1 ingredient")
	}

	if dish.Ingredients[0].Name != ingredient.Name {
		t.Error("Names should be equals")
	}

}

func TestDuplicatedIngredientsAreNotAllowed(t *testing.T) {
	dish := domain.DishAggregate{
		Id:          uuid.UUID{},
		Ingredients: nil,
		DateTime:    time.Time{},
	}

	assert.True(t, len(dish.Ingredients) == 0, "Dish should not have ingredients")
	ingredient := domain.Ingredient{
		Id:        uuid.New(),
		Allergens: nil,
		Name:      "Curcuma",
	}
	ingredient2 := domain.Ingredient{
		Id:        ingredient.Id,
		Allergens: nil,
		Name:      "Curcuma",
	}

	dish.AddIngredient(ingredient)
	assert.True(t, len(dish.Ingredients) == 1, "Dish should have one ingredient")
	dish.AddIngredient(ingredient2)
	assert.True(t, len(dish.Ingredients) == 1, "Dish should not have duplicated ingredients")
}

func TestRemoveIngredientFromDish(t *testing.T) {
	dish := domain.DishAggregate{
		Id:          uuid.UUID{},
		Ingredients: nil,
		DateTime:    time.Time{},
	}
	ingredient := domain.Ingredient{
		Id:        uuid.New(),
		Allergens: nil,
		Name:      "Curcuma",
	}
	ingredient2 := domain.Ingredient{
		Id:        uuid.New(),
		Allergens: nil,
		Name:      "Arroz",
	}
	dish.AddIngredient(ingredient)
	dish.AddIngredient(ingredient2)
	assert.True(t, len(dish.Ingredients) == 2, "Dish should have two ingredients")

	dish.RemoveIngredient(ingredient2)
	assert.True(t, len(dish.Ingredients) == 1, "Dish should have one ingredients")
}
