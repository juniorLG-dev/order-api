package repository

import (
	"database/sql"
	"encoding/json"
	"order/domain/entity"
)

type SQLRepository struct {
	db *sql.DB
}

func NewSQLRepository(db *sql.DB) *SQLRepository {
	return &SQLRepository{
		db: db,
	}
}

func (s *SQLRepository) Save(order entity.Order) error {
	locationPayload, err := json.Marshal(order.Location)
	if err != nil {
		return err
	}
	_, err = s.db.Exec(`
		INSERT INTO orders (id, name, quantity, price, payment_method, location, product_id, createdAt, updatedAt)
		VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?)
	`, order.ID,
		order.Name,
		order.Quantity.Value,
		order.Price,
		order.PaymentMethod.Value,
		locationPayload,
		order.ProductID,
		order.DateInformation.CreatedAt,
		order.DateInformation.UpdatedAt,
	)
	return err
}
