package models

import (
	"errors"
	"fmt"
	"gorm.io/gorm"
)

type Product struct {
	gorm.Model
	Name   string `gorm: "not null;type: varchar(50)"`
	Brand  string `gorm: "not null;type: varchar(100)"`
	UserID uint
}

func (p *Product) BeforeCreate(tx *gorm.DB) error {
	fmt.Println("Product Before Create")

	if len(p.Name) < 3 {
		return errors.New("Product Name to short")
	}
	return nil
}
