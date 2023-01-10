package models

import "gorm.io/gorm"

type Acccount struct {
	gorm.Model
	FullName      string
	Email         string
	AccountNumber string
	PhoneNumber   string
	Status        string
}