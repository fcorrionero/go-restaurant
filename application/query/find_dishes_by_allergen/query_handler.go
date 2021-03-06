package find_dishes_by_allergen

import "github.com/fcorrionero/go-restaurant/domain"

type QueryHandler struct {
	DishesRepository domain.DishesRepository
}

func New(repo domain.DishesRepository) QueryHandler {
	return QueryHandler{
		DishesRepository: repo,
	}
}

func (q QueryHandler) Handle(query Query) []*domain.DishAggregate {
	return q.DishesRepository.FindDishesByAllergen(query.AllergenName)
}
