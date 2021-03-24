package dishes_http

import (
	"fmt"
	"github.com/fcorrionero/go-restaurant/src/application/command/add_allergen"
	"github.com/fcorrionero/go-restaurant/src/application/command/add_dish"
	"github.com/fcorrionero/go-restaurant/src/application/command/add_ingredient"
	"github.com/google/uuid"
	"log"
	"net/http"
)

type CrudHttpController struct {
	addAllergen   add_allergen.CommandHandler
	addIngredient add_ingredient.CommandHandler
	addDish       add_dish.CommandHandler
}

func NewCrudController(
	addAllergen add_allergen.CommandHandler,
	addIngredient add_ingredient.CommandHandler,
	addDish add_dish.CommandHandler) CrudHttpController {

	return CrudHttpController{
		addAllergen:   addAllergen,
		addIngredient: addIngredient,
		addDish:       addDish,
	}

}

func (c CrudHttpController) AddAllergen(w http.ResponseWriter, r *http.Request) {
	if err := r.ParseForm(); err != nil {
		log.Println(err.Error())
		return
	}
	aName := r.FormValue("name")
	command := add_allergen.Command{
		Id:   uuid.New().String(),
		Name: aName,
	}

	err := c.addAllergen.Handle(command)
	if err != nil {
		log.Println(err)
		fmt.Println(err.Error())
		return
	}
}

func (c CrudHttpController) AddIngredient(w http.ResponseWriter, r *http.Request) {
	if err := r.ParseForm(); err != nil {
		log.Println(err.Error())
		fmt.Println(err.Error())
		return
	}
	iName := r.FormValue("name")
	id := uuid.New().String()

	command := add_ingredient.Command{
		Id:           id,
		Name:         iName,
		AllergensIds: nil,
	}

	err := c.addIngredient.Handle(command)
	if err != nil {
		log.Println(err.Error())
	}
}
