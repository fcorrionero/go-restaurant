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
	exists := false
	for _, a := range i.Allergens {
		if a.Id == allergen.Id {
			exists = true
		}
	}
	if !exists {
		i.Allergens = append(i.Allergens, allergen)
	}
}

func (i *Ingredient) RemoveAllergen(allergen Allergen) {
	var allergens []Allergen
	for _, a := range i.Allergens {
		if a.Id != allergen.Id {
			allergens = append(allergens, a)
		}
	}
	i.Allergens = allergens
}
