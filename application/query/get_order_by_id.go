package query

import (
	"database/sql"
	"encoding/json"
	"order/domain/vo"
)

type GetOrderByID struct {
	db *sql.DB
}

func NewGetOrderByID(db *sql.DB) *GetOrderByID {
	return &GetOrderByID{
		db: db,
	}
}

func (g *GetOrderByID) Run(id string) (*GetOrderByIDOutput, error) {
	var order GetOrderByIDOutput
	var locationPayload []byte
	err := g.db.QueryRow("SELECT * FROM orders WHERE id = ?", id).Scan(
		&order.ID,
		&order.Name,
		&order.Quantity,
		&order.Price,
		&order.PaymentMethod,
		&locationPayload,
		&order.ProductID,
		&order.CreatedAt,
		&order.UpdatedAt,
	)
	if err != nil {
		return nil, err
	}
	if err := json.Unmarshal(locationPayload, &order.Location); err != nil {
		return nil, err
	}
	return &order, nil
}

type GetOrderByIDOutput struct {
	ID            string
	Name          string      `json:"name"`
	Quantity      int         `json:"quantity"`
	Price         float64     `json:"price"`
	PaymentMethod string      `json:"payment_method"`
	Location      vo.Location `json:"location"`
	ProductID     string      `json:"product_id"`
	CreatedAt     string      `json:"created_at"`
	UpdatedAt     string      `json:"updated_at"`
}
