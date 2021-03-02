package domain

import "github.com/google/uuid"

type Ingredient struct {
	Id        uuid.UUID
	Allergens []Allergen
	Name      string
}

func (i Ingredient) String() string {
	var allergens string
	for _, a := range i.Allergens {
		allergens += a.String() + " "
	}
	return i.Name + "(" + allergens + ")"
}

func (i *Ingredient) AddAllergen(allergen Allergen) {
	i.Allergens = append(i.Allergens, allergen)
}
