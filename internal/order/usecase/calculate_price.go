package usecase

import "github.com/douglasandradeee/pfa-go/internal/order/entity"

type OrderInputDTO struct {
	ID    string
	Price float64
	Tax   float64
}

type OrderOutputDTO struct {
	ID         string
	Price      float64
	Tax        float64
	FinalPrice float64
}

// CalculatePriceUseCase is a use case for calculating the final price of an order
type CalculatePriceUseCase struct {
	OrderRepository entity.OrderRepositoryInterface
}

// NewCalculateFinalPriceUseCase creates a new instance of CalculateFinalPriceUseCase
func NewCalculateFinalPriceUseCase(orderRepository entity.OrderRepositoryInterface) *CalculatePriceUseCase {
	return &CalculatePriceUseCase{OrderRepository: orderRepository}
}

// Execute calculates the final price of an order and saves it to the repository
func (c *CalculatePriceUseCase) Execute(input OrderInputDTO) (*OrderOutputDTO, error) {
	order, err := entity.NewOrder(input.ID, input.Price, input.Tax)
	if err != nil {
		return nil, err
	}

	err = order.CalculateFinalPrice()
	if err != nil {
		return nil, err
	}
	err = c.OrderRepository.Save(order)
	if err != nil {
		return nil, err
	}
	return &OrderOutputDTO{
		ID:         order.ID,
		Price:      order.Price,
		Tax:        order.Tax,
		FinalPrice: order.FinalPrice,
	}, nil
}
