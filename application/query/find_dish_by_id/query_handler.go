package find_dish_by_id

import (
	"github.com/fcorrionero/go-restaurant/domain"
	"github.com/google/uuid"
)

type QueryHandler struct {
	DishesRepository domain.DishesRepository
}

func New(repo domain.DishesRepository) QueryHandler {
	return QueryHandler{
		DishesRepository: repo,
	}
}

func (q QueryHandler) Handle(query Query) *domain.DishAggregate {
	id, err := uuid.Parse(query.DishId)
	if err != nil {
		return &domain.DishAggregate{}
	}

	return q.DishesRepository.FindDishById(id)
}
