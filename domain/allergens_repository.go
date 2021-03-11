package domain

type AllergensRepository interface {
	FindByName(name string) *Allergen
	FindAllById(ids []int) []*Allergen
	Save(allergen Allergen)
}
