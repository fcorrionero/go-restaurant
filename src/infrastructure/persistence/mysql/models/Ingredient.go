package models

type Ingredient struct {
	Id             []byte `gorm:"primaryKey"`
	IdUuid         string
	IngredientName string
	Allergens      []Allergen `gorm:"many2many:ingredients_allergens"`
	Dishes         []Dish     `gorm:"many2many:dishes_ingredients"`
}
