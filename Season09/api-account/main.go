package main

import (
	"api-account/database"
	"api-account/models"
	"api-account/router"
	"api-account/usecase"
)

func main() {
	db := database.GetDb()
	repository := &models.MysqlRepository{
		DB: db,
	}
	useCase := usecase.NewAccountUseCase(repository)
	//acc := models.User{
	//
	//	FullName: "IKetut Gunawan",
	//	Email:    "iketutg@gmail.com",
	//	//AccountNumber: "",
	//	PhoneNumber: "081213123442",
	//	//Status:        "",
	//}
	//useCase.CreateAccount(&acc)
	router.Server(useCase).Run(":8085")

}
