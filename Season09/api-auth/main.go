package main

import (
	"api-auth/database"
	"api-auth/models"
	"api-auth/router"
	"api-auth/usecase"
)

func main() {
	db := database.GetDb()
	repository := &models.MysqlRepository{
		DB: db,
	}
	useCase := usecase.NewUserUseCase(repository)
	//acc := models.User{
	//
	//	FullName: "IKetut Gunawan",
	//	Email:    "iketutg@gmail.com",
	//	//AccountNumber: "",
	//	PhoneNumber: "081213123442",
	//	//Status:        "",
	//}
	//useCase.CreateAccount(&acc)
	router.Server(useCase).Run(":8087")

}
