package services

import (
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"testing"
	"unittest/entity"
	"unittest/repository"
)

var productRepository = &repository.ProductRepositoryMock{
	Mock: mock.Mock{},
}
var productService = ProductService{
	Repository: productRepository,
}

func TestProductService_GetOneProductBy(t *testing.T) {
	id := "123456"
	product := entity.Product{
		Id:   id,
		Name: "Buku",
	}

	productRepository.Mock.On("FindById", id).Return(product)

	productResult, err := productService.GetOneProductBy(id)

	assert.Nil(t, err)
	assert.NotNil(t, productResult)
	assert.Equal(t, &product, productResult, "result %v has to be %v ", &product, productResult)

}
