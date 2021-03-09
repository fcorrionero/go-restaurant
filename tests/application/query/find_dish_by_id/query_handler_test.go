package find_dish_by_id

import (
	"github.com/fcorrionero/go-restaurant/application/query/find_dish_by_id"
	"github.com/fcorrionero/go-restaurant/tests/mocks"
	"github.com/golang/mock/gomock"
	"github.com/google/uuid"
	"testing"
)

func TestDishMushBeFound(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	m := mocks.NewMockDishesRepository(ctrl)

	id := uuid.New()
	m.EXPECT().FindDishById(id).Times(1)

	queryHandler := find_dish_by_id.QueryHandler{DishesRepository: m}
	query := find_dish_by_id.Query{DishId: id.String()}

	queryHandler.Handle(query)

}
