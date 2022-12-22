package controllers

import (
	"api-account/models"
	"api-account/usecase"
	"github.com/gin-gonic/gin"
	"net/http"
)

type AccountController struct {
	accUseCase usecase.AccountUseCase
}

func NewAccountController(uc usecase.AccountUseCase) *AccountController {
	return &AccountController{
		accUseCase: uc,
	}
}

func (ctrl *AccountController) CreateAccount(ctx *gin.Context) {
	var account models.Acccount

	if err := ctx.ShouldBindJSON(&account); err != nil {
		ctx.JSON(http.StatusBadRequest, err.Error())
		ctx.Abort()
	}

	if err := ctrl.accUseCase.CreateAccount(&account); err != nil {
		ctx.JSON(http.StatusInternalServerError, err.Error())
		ctx.Abort()
	}

	ctx.JSON(http.StatusCreated, account)
}
