package router

import (
	"github.com/gin-gonic/gin"
	"web-api-gin/controller"
)

func StartServer() (router *gin.Engine) {

	router = gin.Default()

	router.POST("/cars", controller.CreateCars)
	router.GET("/cars/:carID", controller.GetCars)
	router.DELETE("/cars/:carID", controller.DeleteCars)

	return router

}
