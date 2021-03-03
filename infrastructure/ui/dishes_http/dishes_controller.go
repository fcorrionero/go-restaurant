package dishes_http

import (
	"fmt"
	"github.com/fcorrionero/go-restaurant/application/query/find_dish_by_id"
	"github.com/fcorrionero/go-restaurant/application/query/find_dish_by_name"
	"net/http"
)

type DishesHttpController struct {
	findDishByIdQueryHandler   find_dish_by_id.QueryHandler
	findDishByNameQueryHandler find_dish_by_name.QueryHandler
}

func NewDishesHttpController(
	findDishByIdQueryHandler find_dish_by_id.QueryHandler,
	findDishByNameQueryHandler find_dish_by_name.QueryHandler,
) DishesHttpController {
	return DishesHttpController{
		findDishByIdQueryHandler:   findDishByIdQueryHandler,
		findDishByNameQueryHandler: findDishByNameQueryHandler,
	}
}

func (dC DishesHttpController) ById(w http.ResponseWriter, r *http.Request) {
	dish := dC.findDishByIdQueryHandler.Handle(find_dish_by_id.Query{DishId: "ad61989c-4b56-4840-af4b-6e614f5afabf"})
	fmt.Fprintf(w, dish.String())
}

func (dC DishesHttpController) ByName(w http.ResponseWriter, r *http.Request) {
	query := find_dish_by_name.Query{Name: "Paella"}
	dish := dC.findDishByNameQueryHandler.Handle(query)
	fmt.Fprintf(w, dish.String()+"\n")
}
