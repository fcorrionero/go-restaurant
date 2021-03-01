package domain

import (
	"github.com/google/uuid"
	"time"
)

type DishAggregate struct {
	Id          uuid.UUID
	Ingredients []Ingredient
	DateTime    time.Time
}
