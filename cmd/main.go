package main

import (
	"database/sql"
	"fmt"

	"github.com/douglasandradeee/pfa-go/internal/order/infra/database"
	"github.com/douglasandradeee/pfa-go/internal/order/usecase"
	"github.com/google/uuid"
)

func main() {
	db, err := sql.Open("mysql", "root:root@tcp(mysql:3306)/orders")
	if err != nil {
		panic(err)
	}
	defer db.Close()
	repository := database.NewOrderRepository(db)
	uc := usecase.NewCalculateFinalPriceUseCase(repository)
	input := usecase.OrderInputDTO{
		ID:    uuid.NewString(),
		Price: 100.00,
		Tax:   15.00,
	}
	output, err := uc.Execute(input)
	if err != nil {
		panic(err)
	}
	fmt.Println("Order ID:", output.ID)
	fmt.Println("Final Price:", output.FinalPrice)
	// Uncomment the following lines to test the repository methods
}
