package dishes_http

import (
	"github.com/fcorrionero/go-restaurant/src/application/command/add_allergen"
	"github.com/fcorrionero/go-restaurant/src/application/command/add_dish"
	"github.com/fcorrionero/go-restaurant/src/application/command/add_ingredient"
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
	id := r.FormValue("id")

	command := add_allergen.Command{
		Id:   id,
		Name: aName,
	}

	c.addAllergen.Handle(command)
}
