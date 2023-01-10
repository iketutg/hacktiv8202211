package router

import (
	"api-account/controllers"
	"api-account/usecase"
	"fmt"
	ginzap "github.com/gin-contrib/zap"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"net/http"
	"time"
)

func Server(useCase usecase.AccountUseCase) *gin.Engine {
	engine := gin.New()

	logger, _ := zap.NewDevelopment()

	engine.Use(ginzap.Ginzap(logger, time.RFC3339, true))
	engine.Use(ginzap.RecoveryWithZap(logger, true))
	engine.GET("/healthcheck", func(context *gin.Context) {
		context.String(http.StatusOK, "Server UP "+fmt.Sprintf("%v", time.Now().Unix()))
	})
	controller := controllers.NewAccountController(useCase)
	//create Account
	engine.POST("/account", controller.CreateAccount)

	return engine
}
