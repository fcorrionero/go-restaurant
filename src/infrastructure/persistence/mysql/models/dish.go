package models

import "time"

type Dish struct {
	Id          []byte `gorm:"primaryKey"`
	IdUuid      string
	DishName    string
	CreatedAt   time.Time
	Ingredients []Ingredient `gorm:"many2many:dishes_ingredients"`
}
