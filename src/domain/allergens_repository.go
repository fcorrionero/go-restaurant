package domain

import "github.com/google/uuid"

type AllergensRepository interface {
	FindByName(name string) *Allergen
	FindById(id uuid.UUID) *Allergen
	FindAll() []*Allergen
	Save(allergen *Allergen)
}
