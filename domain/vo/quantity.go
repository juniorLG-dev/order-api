package vo

import "errors"

type Quantity struct {
	Value int
}

func NewQuantity(value int) (*Quantity, error) {
	if value < 1 || value > 10000 {
		return nil, errors.New("invalid quantity of products")
	}
	return &Quantity{
		Value: value,
	}, nil
}
