package services

import (
	"errors"
	"unittest/entity"
	"unittest/repository"
)

type ProductService struct {
	Repository repository.ProductRepository
}

func (svc ProductService) GetOneProductBy(id string) (*entity.Product, error) {
	product := svc.Repository.FindById(id)
	if product == nil {
		return nil, errors.New("Product Not found")
	}
	return product, nil
}
