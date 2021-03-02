package domain

import (
	"github.com/fcorrionero/go-restaurant/domain"
	"github.com/google/uuid"
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
