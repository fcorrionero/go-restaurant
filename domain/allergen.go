package domain

import "github.com/google/uuid"

type Allergen struct {
	Id   uuid.UUID
	Name string
}

func (a Allergen) String() string {
	return a.Name
}
