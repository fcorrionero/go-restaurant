package domain

type IngredientsRepository interface {
	FindByName(name string) *Ingredient
	FindAllByIds(ids []int) []*Ingredient
	Save(ingredient Ingredient)
}
