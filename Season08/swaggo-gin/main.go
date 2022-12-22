package main

import (
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"log"
	"net/http"
	_ "swaggo-gin/docs"
)

// @title HealthCheck API
// @version 0.1

// @description this is description
// @termOfService http://iketutg.com/termp

// @Company IKG
// @contact email : iketutg@gmail.com

// @license name Apache 2.0

// @host 127.0.0.1:8000
// @BasePath /
// @scheme http
func main() {
	engine := gin.New()

	engine.GET("/health", Healthcheck)
	url := ginSwagger.URL("http://localhost:8001/swagger/doc.json")
	engine.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler, url))

	if err := engine.Run(":8001"); err != nil {
		log.Fatalln(err)
	}
}

// HealthCheck godoc
// @Summary Create order
// @Description create Order include detail item
// @Tags root
// @Accept */*
// @Produce json
// @Success 200 {object} map[string]interface{}
// @Router /health [get]
func Healthcheck(ctx *gin.Context) {
	res := map[string]interface{}{
		"server": "up",
	}
	ctx.JSON(http.StatusOK, res)
}
