package main

import (
	"api-onboard-publish/routers"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	gin.SetMode(gin.DebugMode)

	server := http.Server{
		Addr:    ":8086",
		Handler: routers.Server(),
	}

	go func() {
		if err := server.ListenAndServe(); err != nil {
			log.Printf("Server Listen Error : %s ", err.Error())
		}
	}()

	signals := make(chan os.Signal)

	signal.Notify(signals, syscall.SIGINT, syscall.SIGTERM)
	<-signals
	log.Println("Server is Shutdown")

}
