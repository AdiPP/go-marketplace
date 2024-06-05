package repository

import "github.com/AdiPP/go-marketplace/pkg/domain/entity"

type ProductRepository interface {
	FindById(id string) (result *entity.ProductEntity, err error)
}
