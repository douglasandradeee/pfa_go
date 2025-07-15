package entity

import "errors"

// Order represents an order with its ID, price, tax, and final price.
type Order struct {
	ID         string  `json:"id"`
	Price      float64 `json:"price"`
	Tax        float64 `json:"tax"`
	FinalPrice float64 `json:"final_price"`
}

// NewOrder creates a new order with the given parameters and validates it.
func NewOrder(id string, price float64, tax float64, finalPrice float64) (*Order, error) {
	order := &Order{
		ID:         id,
		Price:      price,
		Tax:        tax,
		FinalPrice: finalPrice,
	}
	// Validate the order fields
	if err := order.IsValid(); err != nil {
		return nil, err
	}
	return order, nil
}

// IsValid checks if the order has valid fields.
func (o *Order) IsValid() error {
	if o.ID == "" {
		return errors.New("order ID cannot be empty")
	}
	if o.Price <= 0 {
		return errors.New("order price cannot be negative or zero")
	}
	if o.Tax <= 0 {
		return errors.New("order tax cannot be negative or zero")
	}
	if o.FinalPrice <= 0 {
		return errors.New("order final price cannot be negative or zero")
	}
	return nil
}
