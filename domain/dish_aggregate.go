package domain

import (
	"github.com/google/uuid"
	"time"
)

type DishAggregate struct {
	Id          uuid.UUID
	Ingredients []Ingredient
	DateTime    time.Time
	Name        string
}

func (d DishAggregate) String() string {
	var ingredients string
	for _, ing := range d.Ingredients {
		ingredients += ing.String() + " "
	}

	return d.Id.String() + " " + d.Name + " [" + ingredients + "] " + d.DateTime.Format("2006-01-02 15:04:05")
}

func (d *DishAggregate) AddIngredient(ingredient Ingredient) {
	exists := false
	for _, i := range d.Ingredients {
		if i.Id == ingredient.Id {
			exists = true
		}
	}
	if !exists {
		d.Ingredients = append(d.Ingredients, ingredient)
	}
}

func (d *DishAggregate) RemoveIngredient(ingredient Ingredient) {
	var ingredients []Ingredient
	for _, i := range d.Ingredients {
		if i.Id != ingredient.Id {
			ingredients = append(ingredients, i)
		}
	}
	d.Ingredients = ingredients
}
