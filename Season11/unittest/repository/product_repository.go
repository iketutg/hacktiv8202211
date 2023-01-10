package repository

import "unittest/entity"

type ProductRepository interface {
	FindById(id string) *entity.Product
}
