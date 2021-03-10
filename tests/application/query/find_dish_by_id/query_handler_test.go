package find_dish_by_id

import (
	"github.com/fcorrionero/go-restaurant/application/query/find_dish_by_id"
	"github.com/fcorrionero/go-restaurant/domain"
	"github.com/fcorrionero/go-restaurant/tests/mocks"
	"github.com/golang/mock/gomock"
	"github.com/google/uuid"
	"reflect"
	"testing"
)

func TestDishMushBeFound(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	m := mocks.NewMockDishesRepository(ctrl)

	id := uuid.New()
	dish := domain.DishAggregate{Id: id}
	m.EXPECT().FindDishById(id).Times(1).Return(dish)

	queryHandler := find_dish_by_id.QueryHandler{DishesRepository: m}
	query := find_dish_by_id.Query{DishId: id.String()}

	result := queryHandler.Handle(query)
	if reflect.TypeOf(result) != reflect.TypeOf(dish) {
		t.Error("QueryHandler must return a DishAggregate")
	}

}
