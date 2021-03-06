package dishes_http

import (
	"fmt"
	"github.com/fcorrionero/go-restaurant/application/query/find_dish_by_id"
	"github.com/fcorrionero/go-restaurant/application/query/find_dish_by_name"
	"github.com/fcorrionero/go-restaurant/application/query/find_dishes_by_allergen"
	"log"
	"net/http"
)

type DishesHttpController struct {
	findDishByIdQueryHandler       find_dish_by_id.QueryHandler
	findDishByNameQueryHandler     find_dish_by_name.QueryHandler
	findDishByAllergenQueryHandler find_dishes_by_allergen.QueryHandler
}

func NewDishesHttpController(
	findDishByIdQueryHandler find_dish_by_id.QueryHandler,
	findDishByNameQueryHandler find_dish_by_name.QueryHandler,
	findDishByAllergenQueryHandler find_dishes_by_allergen.QueryHandler,
) DishesHttpController {
	return DishesHttpController{
		findDishByIdQueryHandler:       findDishByIdQueryHandler,
		findDishByNameQueryHandler:     findDishByNameQueryHandler,
		findDishByAllergenQueryHandler: findDishByAllergenQueryHandler,
	}
}

func (dC DishesHttpController) ById(w http.ResponseWriter, r *http.Request) {
	dish := dC.findDishByIdQueryHandler.Handle(find_dish_by_id.Query{DishId: "ad61989c-4b56-4840-af4b-6e614f5afabf"})
	fmt.Fprintf(w, dish.String())
}

func (dC DishesHttpController) ByName(w http.ResponseWriter, r *http.Request) {
	query := find_dish_by_name.Query{Name: "PaElL"}
	dish := dC.findDishByNameQueryHandler.Handle(query)
	fmt.Fprintf(w, dish.String()+"\n")
}

func (dC DishesHttpController) ByAllergen(w http.ResponseWriter, r *http.Request) {
	query := find_dishes_by_allergen.Query{AllergenName: "Gluten"}

	dishes := dC.findDishByAllergenQueryHandler.Handle(query)

	log.Println(dishes)
	for _, dish := range dishes {
		fmt.Fprintf(w, dish.String()+"\n")
	}

}
