package find_dish_by_id

import (
	"github.com/fcorrionero/go-restaurant/application/query/find_dish_by_id"
	"github.com/fcorrionero/go-restaurant/domain"
	"github.com/fcorrionero/go-restaurant/tests/mocks"
	"github.com/golang/mock/gomock"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestDishMushBeFound(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	m := mocks.NewMockDishesRepository(ctrl)

	id := uuid.New()
	dish := domain.DishAggregate{Id: id}
	m.EXPECT().FindDishById(id).Times(1).Return(&dish)

	queryHandler := find_dish_by_id.New(m)
	query := find_dish_by_id.Query{DishId: id.String()}

	result := queryHandler.Handle(query)
	assert.Equal(t, result, &dish, "Dish found does not match expected result")
}

func TestOnlyValidUuidAreAllowed(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	m := mocks.NewMockDishesRepository(ctrl)
	id := "INVALID_ID"
	query := find_dish_by_id.Query{DishId: id}
	queryHandler := find_dish_by_id.New(m)
	result := queryHandler.Handle(query)
	assert.True(t, len(result.Name) == 0, "Dish is not empty")
}
