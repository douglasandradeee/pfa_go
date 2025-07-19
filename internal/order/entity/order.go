package entity

import "errors"

type OrderRepositoryInterface interface {
	Save(order *Order) error
}

// Order represents an order with its ID, price, tax, and final price.
type Order struct {
	ID         string  `json:"id"`
	Price      float64 `json:"price"`
	Tax        float64 `json:"tax"`
	FinalPrice float64 `json:"final_price"`
}

// CREATE TABLE orders (
//     id VARCHAR(36) PRIMARY KEY,
//     price DECIMAL(10,2) NOT NULL,
//     tax DECIMAL(10,2) NOT NULL,
//     final_price DECIMAL(10,2) NOT NULL
// );

// NewOrder creates a new order with the given parameters and validates it.
func NewOrder(id string, price float64, tax float64) (*Order, error) {
	order := &Order{
		ID:    id,
		Price: price,
		Tax:   tax,
	}
	// Validate the order fields
	if err := order.IsValid(); err != nil {
		return nil, err
	}
	return order, nil
}

// CalculateFinalPrice calculates the final price of the order based on its price and tax.
func (o *Order) CalculateFinalPrice() error {
	if err := o.IsValid(); err != nil {
		return err
	}
	o.FinalPrice = o.Price + o.Tax
	return nil
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
	return nil
}
