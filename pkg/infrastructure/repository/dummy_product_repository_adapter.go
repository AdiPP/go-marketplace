package repository

import (
	"fmt"
	"github.com/AdiPP/go-marketplace/pkg/domain/entity"
	"github.com/google/uuid"
)

type DummyProductRepositoryAdapter struct {
}

func NewDummyProductRepositoryAdapter() *DummyProductRepositoryAdapter {
	return &DummyProductRepositoryAdapter{}
}

func (d DummyProductRepositoryAdapter) FindById(id string) (result *entity.ProductEntity, err error) {
	return &entity.ProductEntity{
		ProductId:    uuid.New().String(),
		ProductName:  fmt.Sprintf("Product: %s", id),
		ProductPrice: 10.50,
	}, nil
}
