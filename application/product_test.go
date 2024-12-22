package application_test

import (
	"testing"

	"github.com/bruno3du/hexagonal/application"
	uuid "github.com/satori/go.uuid"
	"github.com/stretchr/testify/require"
)

func TestApplicationProduct_Enable(t *testing.T) {
	product := application.Product{
		Name:   "HELLO",
		Price:  10,
		Status: application.DISABLED,
	}

	err := product.Enable()

	require.Nil(t, err)

	product.Price = 0

	err = product.Enable()

	require.NotNil(t, err)
	require.Equal(t, "the price must be greater than zero to enable the product", err.Error())
}

func TestApplicationProduct_Disable(t *testing.T) {
	product := application.Product{
		Name:   "HELLO",
		Price:  0,
		Status: application.ENABLED,
	}

	err := product.Disable()

	require.Nil(t, err)

	product.Price = 10

	err = product.Disable()

	require.NotNil(t, err)
	require.Equal(t, "the price must be zero in order to have the product disabled", err.Error())
}

func TestApplicationProduct_IsValid(t *testing.T) {
	product := application.Product{
		ID:     uuid.NewV4().String(),
		Name:   "HELLO",
		Price:  10,
		Status: application.DISABLED,
	}

	isValid, err := product.IsValid()

	// expect to be valid
	require.Nil(t, err)
	require.True(t, isValid)

	product.Status = "INVALID"

	isValid, err = product.IsValid()

	// expect to be invalid
	require.NotNil(t, err)
	require.False(t, isValid)
	require.Equal(t, "the status must be enabled or disabled", err.Error())

	// expect to be valid
	product.Status = application.ENABLED
	_, err = product.IsValid()
	require.Nil(t, err)

	// expect to be invalid
	product.Price = -10
	_, err = product.IsValid()
	require.NotNil(t, err)
	require.Equal(t, "the price must be greater or equal to zero", err.Error())

	// expect to be valid
	product.Price = 10
	_, err = product.IsValid()
	require.Nil(t, err)
}

func TestApplicationProduct_GetPrice(t *testing.T) {
	product := application.Product{
		ID:     uuid.NewV4().String(),
		Name:   "HELLO",
		Price:  10,
		Status: application.DISABLED,
	}

	require.Equal(t, 10.0, product.GetPrice())
}

func TestApplicationProduct_GetStatus(t *testing.T) {
	product := application.Product{
		ID:     uuid.NewV4().String(),
		Name:   "HELLO",
		Price:  10,
		Status: application.DISABLED,
	}

	require.Equal(t, "disabled", product.GetStatus())
}

