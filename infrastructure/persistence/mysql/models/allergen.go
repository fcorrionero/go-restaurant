package models

type Allergen struct {
	Id           []byte `gorm:"primaryKey"`
	IdUuid       string
	AllergenName string
	Ingredients  []Ingredient `gorm:"many2many:ingredients_allergens"`
}
