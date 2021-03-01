package domain

import "github.com/google/uuid"

type Ingredient struct {
	Id        uuid.UUID
	Allergens []Allergen
	name      string
}
