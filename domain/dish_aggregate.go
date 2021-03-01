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

func (d DishAggregate) String() string {
	var ingredients string
	for _, ing := range d.Ingredients {
		ingredients += ing.String() + " "
	}

	return d.Id.String() + " [" + ingredients + "] " + d.DateTime.Format("2006-01-02 15:04:05")
}
