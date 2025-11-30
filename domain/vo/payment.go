package vo

import "errors"

type Payment struct {
	Value string
}

func NewPayment(method string) (*Payment, error) {
	allowedMethods := []string{"PIX", "Crédito", "Débito"}
	for _, alloallowedMethod := range allowedMethods {
		if method == alloallowedMethod {
			return &Payment{
				Value: method,
			}, nil
		}
	}
	return nil, errors.New("this payment method is not allowed")
}
