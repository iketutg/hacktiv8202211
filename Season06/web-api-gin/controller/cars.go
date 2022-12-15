package controller

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

type Car struct {
	ID    string `json:"id"`
	Brand string `json:"brand"`
	Model string `json:"model"`
	Harga string `json:"harga"`
}

var dataCars = []Car{}

func CreateCars(ctx *gin.Context) {
	var newCar Car
	if err := ctx.ShouldBindJSON(&newCar); err != nil {
		ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}

	newCar.ID = fmt.Sprintf("CR-%02d", len(dataCars)+1)

	dataCars = append(dataCars, newCar)
	//Send Kafka
	//carByte, _ := json.Marshal(newCar)
	//

	ctx.JSON(http.StatusCreated, gin.H{
		"car": newCar,
	})
}

func GetCars(ctx *gin.Context) {
	carID := ctx.Param("carID")

	var responseCar Car
	ok := false
	for _, v := range dataCars {
		if carID == v.ID {
			responseCar = v
			ok = true
			break
		}
	}

	if ok {
		ctx.JSON(http.StatusOK, gin.H{
			"kode":    http.StatusOK,
			"data":    responseCar,
			"Message": "Success",
		})
		return
	}

	ctx.JSON(http.StatusNotFound, gin.H{
		"kode":    -1,
		"data":    nil,
		"Message": "Not Found",
	})

}

func DeleteCars(ctx *gin.Context) {
	carID := ctx.Param("carID")

	var responseCar Car
	var idxCar int
	ok := false
	for idx, v := range dataCars {
		if carID == v.ID {
			responseCar = v
			ok = true
			idxCar = idx
			break
		}
	}

	if ok {
		copy(dataCars[idxCar:], dataCars[idxCar+1:])
		dataCars[len(dataCars)-1] = Car{}
		dataCars = dataCars[:len(dataCars)-1]

		ctx.JSON(http.StatusOK, gin.H{
			"kode":    http.StatusOK,
			"data":    responseCar,
			"Message": "Delete Success",
		})
		return
	}

	ctx.JSON(http.StatusNotFound, gin.H{
		"kode":    -1,
		"data":    nil,
		"Message": "Not Found",
	})

}
