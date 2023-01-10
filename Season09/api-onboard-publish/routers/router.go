package routers

import (
	"api-onboard-publish/controller"
	"github.com/gin-gonic/gin"
)

func Server() *gin.Engine {
	engine := gin.New()

	engine.POST("/register", controller.Register)

	return engine
}
