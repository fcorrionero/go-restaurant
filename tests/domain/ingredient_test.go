package domain

import (
	"github.com/fcorrionero/go-restaurant/domain"
	"github.com/google/uuid"
	"testing"
)

func TestAllergenShouldBeAddedToIngredient(t *testing.T) {
	i := domain.Ingredient{
		Id:        uuid.New(),
		Allergens: nil,
		Name:      "Pasta",
	}

	if len(i.Allergens) > 0 {
		t.Error("Allergens should be empty")
	}

	allergen := domain.Allergen{
		Id:   uuid.New(),
		Name: "Gluten",
	}

	i.AddAllergen(allergen)
	if len(i.Allergens) != 1 {
		t.Error("Should be only 1 allergen")
	}

	if i.Allergens[0].Name != allergen.Name {
		t.Error("Names should be equals")
	}
}
