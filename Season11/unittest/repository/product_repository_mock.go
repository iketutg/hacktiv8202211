package repository

import (
	"github.com/stretchr/testify/mock"
	"unittest/entity"
)

type ProductRepositoryMock struct {
	Mock mock.Mock
}

func (repo *ProductRepositoryMock) FindById(id string) *entity.Product {
	mockProduct := repo.Mock.Called(id)

	if mockProduct.Get(0) == nil {
		return nil
	}
	product := mockProduct.Get(0).(entity.Product)
	return &product
}
