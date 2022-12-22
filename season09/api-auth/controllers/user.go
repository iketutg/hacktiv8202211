package controllers

import (
	"api-auth/models"
	"api-auth/usecase"
	"github.com/gin-gonic/gin"
	"net/http"
)

type UserController struct {
	userUseCase usecase.UserUseCase
}

func NewUserController(uc usecase.UserUseCase) *UserController {
	return &UserController{
		userUseCase: uc,
	}
}

func (ctrl *UserController) CreateUser(ctx *gin.Context) {
	var user models.User

	if err := ctx.ShouldBindJSON(&user); err != nil {
		ctx.JSON(http.StatusBadRequest, err.Error())
		ctx.Abort()
	}

	if err := ctrl.userUseCase.CreateUser(&user); err != nil {
		ctx.JSON(http.StatusInternalServerError, err.Error())
		ctx.Abort()
	}

	ctx.JSON(http.StatusCreated, user)
}

func (ctrl *UserController) GetInfo(ctx *gin.Context) {
	ctx.JSON(http.StatusCreated, "Hello")
}
