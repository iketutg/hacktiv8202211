package main

import (
	"fmt"
	"github.com/gin-contrib/cache"
	"github.com/gin-contrib/cache/persistence"
	"github.com/gin-gonic/gin"
	"time"
)

func main() {
	engine := gin.Default()
	duration, _ := time.ParseDuration("2m")
	rStore := persistence.NewRedisCache("localhost:6379", "", duration)

	engine.GET("/ping", func(context *gin.Context) {
		context.String(200, "pong "+fmt.Sprintf(" Time : %v", time.Now().Unix()))
	})

	engine.GET("/ping_cache", cache.CachePage(rStore, time.Minute, func(context *gin.Context) {
		context.String(200, "pong cache "+fmt.Sprintf(" Time : %v", time.Now().Unix()))
	}))

	engine.Run(":8000")

}
