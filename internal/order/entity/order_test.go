package entity_test

import (
	"testing"

	"github.com/douglasandradeee/pfa-go/internal/order/entity"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
)

func TestGivenEmptyId_WhenCreateAnewOrder_ThenShouldReceiveAndError(t *testing.T) {
	order := entity.Order{}
	assert.Error(t, order.IsValid(), "Expected error for empty ID")
}

func TestGivenEmptyPrice_WhenCreateAnewOrder_ThenShouldReceiveAndError(t *testing.T) {
	order := entity.Order{
		ID:         uuid.NewString(),
		Tax:        1,
		FinalPrice: 1,
	}
	assert.Error(t, order.IsValid(), "Expected error for empty price")
}

func TestGivenEmptyTax_WhenCreateAnewOrder_ThenShouldReceiveAndError(t *testing.T) {
	order := entity.Order{
		ID:         uuid.NewString(),
		Price:      1,
		FinalPrice: 0,
	}
	assert.Error(t, order.IsValid(), "Expected error for empty tax")
}

func TestGivenNegativePrice_WhenCreateAnewOrder_ThenShouldReceiveAndError(t *testing.T) {
	order := entity.Order{
		ID:         uuid.NewString(),
		Price:      -1.5,
		Tax:        1,
		FinalPrice: 1,
	}
	assert.Error(t, order.IsValid(), "Expected error for negative price")
}

func TestGivenNegativeTax_WhenCreateAnewOrder_ThenShouldReceiveAndError(t *testing.T) {
	order := entity.Order{
		ID:         uuid.NewString(),
		Price:      1,
		Tax:        -1.5,
		FinalPrice: 1,
	}
	assert.Error(t, order.IsValid(), "Expected error for negative tax")
}

func TestGivenAValidParams_WhenCallNewOrder_ThenShould_ReceiveCreateOrderWithAllParams(t *testing.T) {
	order, err := entity.NewOrder(
		"123",
		100.0,
		10.0,
	)
	assert.NoError(t, err)
	assert.Equal(t, "123", order.ID)
	assert.Equal(t, 100.0, order.Price)
	assert.Equal(t, 10.0, order.Tax)
}

func TestGivenAValidParams_WhenCallCalculateFinalPrice_ThenShouldCalculeteFinalPriceAndSetItOnFinalPriceProperty(t *testing.T) {
	order, err := entity.NewOrder("123", 10, 2)
	assert.NoError(t, err)
	err = order.CalculateFinalPrice()
	assert.NoError(t, err)
	assert.Equal(t, 12.0, order.FinalPrice)
}

func TestGivenInvalidOrder_WhenCallCalculateFinalPrice_ThenShouldReturnError(t *testing.T) {
	// Testando com ordem inválida (preço negativo)
	order := &entity.Order{
		ID:    "123",
		Price: -10,
		Tax:   2,
	}
	err := order.CalculateFinalPrice()
	assert.Error(t, err)
	assert.Equal(t, 0.0, order.FinalPrice) // FinalPrice não deve ser alterado
}

func TestGivenInvalidParams_WhenCallNewOrder_ThenShouldReturnError(t *testing.T) {
	// Testando com parâmetros inválidos (taxa negativa)
	order, err := entity.NewOrder("123", 10, -2)
	assert.Error(t, err)
	assert.Nil(t, order)

	// Testando com parâmetros inválidos (ID vazio)
	order, err = entity.NewOrder("", 10, 2)
	assert.Error(t, err)
	assert.Nil(t, order)

	// Testando com parâmetros inválidos (preço zero)
	order, err = entity.NewOrder("123", 0, 2)
	assert.Error(t, err)
	assert.Nil(t, order)
}
