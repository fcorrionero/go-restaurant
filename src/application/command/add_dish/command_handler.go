package add_dish

import (
	"github.com/fcorrionero/go-restaurant/src/domain"
	"github.com/google/uuid"
	"log"
	"time"
)

type CommandHandler struct {
	DishesRepository      domain.DishesRepository
	ReadDishesRepository  domain.DishesRepository
	IngredientsRepository domain.IngredientsRepository
}

func New(dRepo domain.DishesRepository, rDRepo domain.DishesRepository, iRepo domain.IngredientsRepository) CommandHandler {
	return CommandHandler{
		DishesRepository:      dRepo,
		ReadDishesRepository:  rDRepo,
		IngredientsRepository: iRepo,
	}
}

func (c CommandHandler) Handle(command Command) {
	id, err := uuid.Parse(command.Id)
	if err != nil {
		log.Println(err.Error())
		return
	}

	d := domain.DishAggregate{
		Id:          id,
		Ingredients: nil,
		DateTime:    time.Now(),
		Name:        command.Name,
	}

	for _, iSId := range command.IngredientIds {
		iId, _ := uuid.Parse(iSId)
		i := c.IngredientsRepository.FindById(iId)
		d.Ingredients = append(d.Ingredients, *i)
	}

	c.DishesRepository.SaveDish(&d)

	go func() {
		dIA := c.DishesRepository.FindDishById(id)
		c.ReadDishesRepository.SaveDish(dIA)
	}()
}
