package router

import (
	"api-auth/controllers"
	"api-auth/usecase"
	"fmt"
	ginzap "github.com/gin-contrib/zap"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"net/http"
	"time"
)

func Server(useCase usecase.UserUseCase) *gin.Engine {
	engine := gin.New()

	logger, _ := zap.NewDevelopment()

	engine.Use(ginzap.Ginzap(logger, time.RFC3339, true))
	engine.Use(ginzap.RecoveryWithZap(logger, true))
	engine.GET("/healthcheck", func(context *gin.Context) {
		context.String(http.StatusOK, "Server UP "+fmt.Sprintf("%v", time.Now().Unix()))
	})
	controller := controllers.NewUserController(useCase)
	//create Account
	engine.POST("/user", controller.CreateUser)
	engine.POST("/login", controller.LoginUser)

	group := engine.Group("/v1/users").Use(BasicAuth(useCase))
	{
		group.GET("/info", controller.GetInfo)
	}

	return engine
}

func BasicAuth(useCase usecase.UserUseCase) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		fmt.Println("Before call controller ")
		username, password, ok := ctx.Request.BasicAuth()
		if ok {
			user, err := useCase.GetUserByUserName(username)
			if err == nil {
				if password == user.Password {
					fmt.Println("Password sama")
					ctx.Next()
					return
				}
			}
		}

		ctx.Abort()
		ctx.JSON(http.StatusBadRequest, "")
	}

}
