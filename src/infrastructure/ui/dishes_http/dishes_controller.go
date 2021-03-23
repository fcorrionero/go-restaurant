package dishes_http

import (
	"fmt"
	"github.com/fcorrionero/go-restaurant/src/application/query/find_dish_by_id"
	"github.com/fcorrionero/go-restaurant/src/application/query/find_dish_by_name"
	"github.com/fcorrionero/go-restaurant/src/application/query/find_dishes_by_allergen"
	"log"
	"net/http"
	"strings"
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

func (dC DishesHttpController) Search(w http.ResponseWriter, r *http.Request) {
	query := r.URL.Query()
	id, idExists := query["id"]
	if idExists && len(id) > 0 {
		dC.byId(w, strings.Join(id, ""))
		return
	}

	name, nameExists := query["name"]
	if nameExists && len(name) > 0 {
		dC.byName(w, strings.Join(name, ""))
		return
	}

	allergen, aExists := query["allergen"]
	if aExists && len(allergen) > 0 {
		dC.byAllergen(w, strings.Join(allergen, ""))
		return
	}

	fmt.Fprintf(w, "Dish(es) not found")
}

func (dC DishesHttpController) byId(w http.ResponseWriter, id string) {
	// id : "ad61989c-4b56-4840-af4b-6e614f5afabf"
	dish := dC.findDishByIdQueryHandler.Handle(find_dish_by_id.Query{DishId: id})
	fmt.Fprintf(w, dish.String())
}

func (dC DishesHttpController) byName(w http.ResponseWriter, name string) {
	// name : paella
	query := find_dish_by_name.Query{Name: name}
	dish := dC.findDishByNameQueryHandler.Handle(query)
	fmt.Fprintf(w, dish.String()+"\n")
}

func (dC DishesHttpController) byAllergen(w http.ResponseWriter, allergen string) {
	// allergen : Gluten
	query := find_dishes_by_allergen.Query{AllergenName: allergen}

	dishes := dC.findDishByAllergenQueryHandler.Handle(query)

	log.Println(dishes)
	for _, dish := range dishes {
		fmt.Fprintf(w, dish.String()+"\n")
	}

}
