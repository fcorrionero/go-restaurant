package add_ingredient

import (
	"github.com/fcorrionero/go-restaurant/src/application/command/add_ingredient"
	"github.com/fcorrionero/go-restaurant/src/domain"
	"github.com/fcorrionero/go-restaurant/tests/mocks"
	"github.com/golang/mock/gomock"
	"github.com/google/uuid"
	"testing"
)

func TestIngredientMustBeSaved(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	iM := mocks.NewMockIngredientsRepository(ctrl)
	aM := mocks.NewMockAllergensRepository(ctrl)

	iId := uuid.New()
	aId1 := uuid.New()
	aId2 := uuid.New()

	iName := "Ingredient Name"

	command := add_ingredient.Command{
		Id:           iId.String(),
		Name:         iName,
		AllergensIds: []string{aId1.String(), aId2.String()},
	}

	a1 := domain.Allergen{
		Id:   aId1,
		Name: "Alergen1",
	}
	a2 := domain.Allergen{
		Id:   aId1,
		Name: "Alergen2",
	}
	aM.EXPECT().FindById(aId1).Times(1).Return(&a1)
	aM.EXPECT().FindById(aId2).Times(1).Return(&a2)

	i := domain.Ingredient{
		Id:        iId,
		Allergens: []domain.Allergen{a1, a2},
		Name:      iName,
	}

	iM.EXPECT().Save(&i).Times(1)

	commandHandler := add_ingredient.New(iM, aM)
	commandHandler.Handle(command)
}
