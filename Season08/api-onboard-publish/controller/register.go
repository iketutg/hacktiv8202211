package controller

import (
	"api-onboard-publish/kafka"
	"api-onboard-publish/models"
	"encoding/json"
	"errors"
	"github.com/gin-gonic/gin"
	"net/http"
)

func Register(ctx *gin.Context) {
	var person models.Person

	if err := ctx.ShouldBindJSON(&person); err != nil {
		ctx.AbortWithError(http.StatusBadRequest, err)
	}
	//connect kafka
	producer, err := kafka.ConnProducer()
	if err != nil {
		ctx.AbortWithError(http.StatusInternalServerError, err)
	}
	personByte, _ := json.Marshal(person)

	if kafka.Publish(personByte, producer) != nil {
		ctx.AbortWithError(http.StatusInternalServerError, errors.New("Publish Kafka Error"))
	}

	ctx.JSON(http.StatusOK, gin.H{
		"message": "Success",
		"code":    200,
		"data":    person,
	})
}
