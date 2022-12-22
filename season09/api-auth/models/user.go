package models

import "gorm.io/gorm"

type User struct {
	gorm.Model
	UserName string
	Password string
	Email    string
	Status   string
}
