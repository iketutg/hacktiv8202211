package main

import (
	"errors"
	"fmt"
	"gorm.io/gorm"
	"latihancrudgorm/database"
	"latihancrudgorm/models"
)

func main() {
	database.DbConnection()
	database.Migrate()
	createUser("Iketutg@gmail.com")
	createProduct(1, "ABC", "Kecap")
	getUserById(1)
	updateUserById(1, "iketutg@yahoo.com")

}

func updateUserById(userId uint, email string) {
	db := database.GetDb()
	user := models.User{}
	err := db.Model(&user).Where("id", userId).Updates(
		models.User{
			Email: email,
		},
	).Error

	if err != nil {
		fmt.Println("Update Failed : ", err.Error())
		return
	}
	fmt.Println("Update Success : ", user)
}

func getUserById(userId uint) {
	db := database.GetDb()
	user := models.User{}
	err := db.First(&user, "id= ?", userId).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			fmt.Println("Data tidak di temukan ")
			return
		}

		fmt.Println("Find User Error : ", err.Error())
		return
	}

	fmt.Println("Data ditemukan : ", user)

}

func createProduct(userId uint, brand string, name string) {
	db := database.GetDb()
	product := models.Product{
		UserID: userId,
		Brand:  brand,
		Name:   name,
	}

	err := db.Create(&product).Error
	if err != nil {
		fmt.Println("Crete Product Error : ", err.Error())
		return
	}
	fmt.Println("Create Product Success :", product)
}

func createUser(email string) {
	db := database.GetDb()
	user := models.User{
		Email: email,
	}
	err := db.Create(&user).Error
	if err != nil {
		fmt.Println("Crete User Error : ", err.Error())
		return
	}
	fmt.Println("Create User Success :", user)
}
