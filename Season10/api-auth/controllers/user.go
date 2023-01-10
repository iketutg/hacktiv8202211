package controllers

import (
	"api-auth/auth"
	"api-auth/models"
	"api-auth/usecase"
	"fmt"
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
func (ctrl *UserController) LoginUser(ctx *gin.Context) {
	var login models.LoginRequest
	if err := ctx.ShouldBindJSON(&login); err != nil {
		ctx.JSON(http.StatusBadRequest, err.Error())
		//ctx.Abort()
		return
	}

	loginresonpse := models.LoginResponse{
		Status:  -1,
		Message: "Failed",
		Token:   "",
	}
	user, err := ctrl.userUseCase.GetUserByUserName(login.UserName)
	if err != nil {
		fmt.Println("Get user Error : ", err.Error())
		ctx.JSON(http.StatusNotFound, loginresonpse)
		ctx.Abort()
		return
	}

	if user.Password == login.Password {
		token := auth.GetToken(user.UserName, user.Email)
		loginresonpse.Status = 200
		loginresonpse.Message = "Success"
		loginresonpse.Token = token
		ctx.JSON(http.StatusOK, loginresonpse)
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
	fmt.Println("Controller Get Info Call")
	ctx.JSON(http.StatusCreated, "Hello")
}
