package add_allergen

import (
	"github.com/fcorrionero/go-restaurant/src/domain"
	"github.com/google/uuid"
	"log"
)

type CommandHandler struct {
	AllergenRepository domain.AllergensRepository
}

func New(repo domain.AllergensRepository) CommandHandler {
	return CommandHandler{AllergenRepository: repo}
}

func (c CommandHandler) Handle(command Command) {
	id, err := uuid.Parse(command.Id)
	if err != nil {
		log.Println("Invalid allergen id")
		return
	}
	a := domain.Allergen{
		Id:   id,
		Name: command.Name,
	}
	c.AllergenRepository.Save(&a)
}
