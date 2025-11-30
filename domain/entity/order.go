package entity

import "order/domain/vo"

type Order struct {
	ID            string
	Name          string
	Quantity      vo.Quantity
	Price         float64
	PaymentMethod vo.Payment
	Location      vo.Location
	ProductID     string
}

func NewOrder(
	id, name, paymentMethodValue, productID string,
	quantityValue int,
	price float64,
	location vo.Location,
) (*Order, error) {
	quantity, err := vo.NewQuantity(quantityValue)
	if err != nil {
		return nil, err
	}
	paymentMethod, err := vo.NewPayment(paymentMethodValue)
	if err != nil {
		return nil, err
	}
	return &Order{
		ID:            id,
		Name:          name,
		Quantity:      *quantity,
		Price:         price,
		PaymentMethod: *paymentMethod,
		Location:      location,
		ProductID:     productID,
	}, nil
}
