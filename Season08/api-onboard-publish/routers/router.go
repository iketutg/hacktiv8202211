package routers

import (
	"api-onboard-publish/auth"
	"api-onboard-publish/controller"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"strings"
)

func Server() *gin.Engine {
	engine := gin.New()
	engine.Use(JWTAuth())
	engine.POST("/register", controller.Register)

	return engine
}

func JWTAuth() gin.HandlerFunc {
	return func(context *gin.Context) {
		datatoken := context.GetHeader("Authorization")
		fmt.Println("Data token : ", datatoken)
		if !strings.HasPrefix(datatoken, "Bearer") {
			context.JSON(http.StatusMethodNotAllowed, gin.H{
				"Error": "Request token salah",
			})
			context.Abort()
			return
		}

		token := strings.Split(datatoken, " ")
		if len(token) != 2 {
			context.JSON(http.StatusMethodNotAllowed, gin.H{
				"Error": "Request token salah",
			})
			context.Abort()
			return
		}

		tokenStr := token[1]
		fmt.Println("Token : ", tokenStr)
		if tokenStr == "" {
			context.JSON(http.StatusMethodNotAllowed, gin.H{
				"Error": "Request token salah",
			})
			context.Abort()
			return
		}

		if !auth.ValidateJWT(tokenStr) {
			context.JSON(http.StatusMethodNotAllowed, gin.H{
				"Error": "Validate token Gagal",
			})
			context.Abort()
			return
		}

		context.Next()
	}
}
