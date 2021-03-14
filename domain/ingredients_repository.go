package domain

import "github.com/google/uuid"

type IngredientsRepository interface {
	FindByName(name string) *Ingredient
	FindById(id uuid.UUID) *Ingredient
	FindAll() []*Ingredient
	FindAllByAllergen(aId uuid.UUID) []*Ingredient
	Save(ingredient *Ingredient)
}
