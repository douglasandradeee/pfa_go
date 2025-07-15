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

func TestGivenEmptyFinalPrice_WhenCreateAnewOrder_ThenShouldReceiveAndError(t *testing.T) {
	order := entity.Order{
		ID:    uuid.NewString(),
		Price: 1,
		Tax:   1,
	}
	assert.Error(t, order.IsValid(), "Expected error for empty final price")
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

func TestGivenNegativeFinalPrice_WhenCreateAnewOrder_ThenShouldReceiveAndError(t *testing.T) {
	order := entity.Order{
		ID:         uuid.NewString(),
		Price:      1,
		Tax:        1,
		FinalPrice: -1.5,
	}
	assert.Error(t, order.IsValid(), "Expected error for negative final price")
}

func TestGivenAValidParams_WhenCallNewOrder_ThenShould_ReceiveCreateOrderWithAllParams(t *testing.T) {
	order, err := entity.NewOrder(
		"123",
		100.0,
		10.0,
		110.0,
	)
	assert.NoError(t, err)
	assert.Equal(t, "123", order.ID)
	assert.Equal(t, 100.0, order.Price)
	assert.Equal(t, 10.0, order.Tax)
	assert.Equal(t, 110.0, order.FinalPrice)
}
