package find_dish_by_name

import (
	"github.com/fcorrionero/go-restaurant/application/query/find_dish_by_name"
	"github.com/fcorrionero/go-restaurant/domain"
	"github.com/fcorrionero/go-restaurant/tests/mocks"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestDishMustBeFound(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	m := mocks.NewMockDishesRepository(ctrl)
	name := "Paella"
	dish := domain.DishAggregate{Name: name}
	m.EXPECT().FindDishByName(name).Times(1).Return(dish)

	queryHandler := find_dish_by_name.New(m)
	query := find_dish_by_name.Query{Name: name}
	result := queryHandler.Handle(query)
	assert.Equal(t, dish, result, "Dishes must be equals")
}
