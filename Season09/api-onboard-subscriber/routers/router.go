package routers

import (
	"fmt"
	ginzap "github.com/gin-contrib/zap"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"net/http"
	"time"
)

func Server() *gin.Engine {
	engine := gin.New()

	logger, _ := zap.NewDevelopment()

	engine.Use(ginzap.Ginzap(logger, time.RFC3339, true))
	engine.Use(ginzap.RecoveryWithZap(logger, true))
	engine.GET("/healthcheck", func(context *gin.Context) {
		context.String(http.StatusOK, "Server UP "+fmt.Sprintf("%v", time.Now().Unix()))
	})

	return engine

}
