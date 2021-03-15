package domain

import (
	"github.com/fcorrionero/go-restaurant/domain"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
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

func TestDuplicatedAllergensAreNotAllowed(t *testing.T) {
	i := domain.Ingredient{
		Id:        uuid.New(),
		Allergens: nil,
		Name:      "Pasta",
	}

	assert.True(t, len(i.Allergens) == 0, "Allergens should be empty")

	allergen := domain.Allergen{
		Id:   uuid.New(),
		Name: "Gluten",
	}

	allergen2 := domain.Allergen{
		Id:   allergen.Id,
		Name: "Gluten",
	}
	i.AddAllergen(allergen)
	assert.True(t, len(i.Allergens) == 1, "One allergen should be added")

	i.AddAllergen(allergen2)
	assert.True(t, len(i.Allergens) == 1, "Duplicated allergens are not allowed")
}

func TestRemoveAllergenFromIngredient(t *testing.T) {
	i := domain.Ingredient{
		Id:        uuid.New(),
		Allergens: nil,
		Name:      "Pasta",
	}
	allergen := domain.Allergen{
		Id:   uuid.New(),
		Name: "Gluten",
	}

	allergen2 := domain.Allergen{
		Id:   uuid.New(),
		Name: "Crustáceos",
	}
	i.AddAllergen(allergen)
	i.AddAllergen(allergen2)
	assert.True(t, len(i.Allergens) == 2, "Ingredient must have two allergens")

	i.RemoveAllergen(allergen2)
	assert.True(t, len(i.Allergens) == 1, "Ingredient must have one allergens")
}
