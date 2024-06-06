package repository

import "github.com/AdiPP/go-marketplace/pkg/domain/entity"

type OrderRepository interface {
	Save(order *entity.OrderEntity) (*entity.OrderEntity, error)
}
