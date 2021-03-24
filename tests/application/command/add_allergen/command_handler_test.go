package add_allergen

import (
	"errors"
	"github.com/fcorrionero/go-restaurant/src/application/command/add_allergen"
	"github.com/fcorrionero/go-restaurant/src/domain"
	"github.com/fcorrionero/go-restaurant/tests/mocks"
	"github.com/golang/mock/gomock"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
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

func TestErrorShouldAppear(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	m := mocks.NewMockAllergensRepository(ctrl)
	a := domain.Allergen{
		Id:   uuid.New(),
		Name: "Allergen Name",
	}
	err := errors.New("test saving error")
	m.EXPECT().Save(&a).Times(1).Return(err)
	command := add_allergen.Command{Name: a.Name, Id: a.Id.String()}
	commandHandler := add_allergen.New(m)

	errH := commandHandler.Handle(command)

	assert.True(t, errH == err)
}

func TestOnlyValidUuidAllowed(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	id := "invalid-uuid"
	name := "Allergen Name"

	command := add_allergen.Command{Name: name, Id: id}
	m := mocks.NewMockAllergensRepository(ctrl)
	commandHandler := add_allergen.New(m)

	err := commandHandler.Handle(command)
	assert.True(t, err != nil)
}
