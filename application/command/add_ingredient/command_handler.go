package add_ingredient

import (
	"github.com/fcorrionero/go-restaurant/domain"
	"github.com/google/uuid"
	"log"
)

type CommandHandler struct {
	IngredientsRepository domain.IngredientsRepository
	AllergensRepository   domain.AllergensRepository
}

func New(iRepo domain.IngredientsRepository, aRepo domain.AllergensRepository) CommandHandler {
	return CommandHandler{IngredientsRepository: iRepo, AllergensRepository: aRepo}
}

func (c CommandHandler) Handle(command Command) {
	id, err := uuid.Parse(command.Id)
	if err != nil {
		log.Println("Invalid allergen id")
		return
	}
	i := domain.Ingredient{
		Id:        id,
		Allergens: nil,
		Name:      command.Name,
	}

	for _, aId := range command.AllergensIds {
		aId, _ := uuid.Parse(aId)
		a := c.AllergensRepository.FindById(aId)
		i.Allergens = append(i.Allergens, *a)
	}

	c.IngredientsRepository.Save(&i)
}
