package repository

import (
	"fmt"
	"github.com/AdiPP/go-marketplace/pkg/domain/entity"
	"github.com/google/uuid"
)

type DummyRepositoryAdapter struct {
}

func (d *DummyRepositoryAdapter) Save(order *entity.OrderEntity) (*entity.OrderEntity, error) {
	return order, nil
}

func NewDummyRepositoryAdapter() *DummyRepositoryAdapter {
	return &DummyRepositoryAdapter{}
}

func (d *DummyRepositoryAdapter) FindById(id string) (result *entity.ProductEntity, err error) {
	return &entity.ProductEntity{
		ProductId:    uuid.New().String(),
		ProductName:  fmt.Sprintf("Product: %s", id),
		ProductPrice: 10.50,
	}, nil
}
