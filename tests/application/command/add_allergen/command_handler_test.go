package add_allergen

import (
	"github.com/fcorrionero/go-restaurant/src/application/command/add_allergen"
	"github.com/fcorrionero/go-restaurant/src/domain"
	"github.com/fcorrionero/go-restaurant/tests/mocks"
	"github.com/golang/mock/gomock"
	"github.com/google/uuid"
	"testing"
)

func TestAllergenMustBeSaved(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	m := mocks.NewMockAllergensRepository(ctrl)
	a := domain.Allergen{
		Id:   uuid.New(),
		Name: "Allergen Name",
	}
	m.EXPECT().Save(&a).Times(1)

	command := add_allergen.Command{Name: a.Name, Id: a.Id.String()}
	commandHandler := add_allergen.New(m)

	commandHandler.Handle(command)
}
