package repository

import "order/domain/entity"

type OrderRepository interface {
	Save(order entity.Order) error
}
