package main

import (
	"api-onboard-subscriber/kafka"
	"api-onboard-subscriber/routers"
	"fmt"
	"github.com/Shopify/sarama"
	"github.com/gin-gonic/gin"

	"log"
	"os"
)

func main() {

	sarama.Logger = log.New(os.Stdout, "", log.Ltime)

	consumer, err := kafka.ConnConsumer()
	if err != nil {
		fmt.Println("Error Consumer connection ", err.Error())
		os.Exit(1)
	}
	go kafka.Subscriber(consumer)
	gin.SetMode(gin.DebugMode)
	routers.Server().Run(":8083")
}
