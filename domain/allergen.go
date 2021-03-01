package domain

import "github.com/google/uuid"

type Allergen struct {
	Id   uuid.UUID
	name string
}
