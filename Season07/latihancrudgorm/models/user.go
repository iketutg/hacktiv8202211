package models

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Email   string `gorm:"not null;unique;type:varchar(250)"`
	Product []Product
}
