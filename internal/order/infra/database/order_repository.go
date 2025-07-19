package database

import (
	"database/sql"
	"errors"

	"github.com/douglasandradeee/pfa-go/internal/order/entity"
	_ "github.com/go-sql-driver/mysql"
)

type OrderRepository struct {
	Db *sql.DB
}

func NewOrderRepository(db *sql.DB) *OrderRepository {
	return &OrderRepository{Db: db}
}

// Save persists the order to the database.
func (r *OrderRepository) Save(order *entity.Order) error {
	// Ensure the order is not nil and is valid before saving
	if order == nil {
		return errors.New("order cannot be nil")
	}
	// Validate the order before saving
	if err := order.IsValid(); err != nil {
		return err
	}

	tx, err := r.Db.Begin()
	if err != nil {
		return err
	}
	defer tx.Rollback()

	stmt, err := tx.Prepare("INSERT INTO orders (id, price, tax, final_price) VALUES (?, ?, ?, ?)")
	if err != nil {
		return err
	}
	defer stmt.Close()

	_, err = stmt.Exec(order.ID, order.Price, order.Tax, order.FinalPrice)
	if err != nil {
		return err
	}

	return tx.Commit()
}
