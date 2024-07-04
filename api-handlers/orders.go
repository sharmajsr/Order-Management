package api_handlers

import (
	"encoding/json"
	"fmt"
	"io"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/sharmajsr/oms/db"
)

func GetOrder(ctx *gin.Context) {
	orderId := ctx.Param("id")
	ord, err := strconv.Atoi(orderId)
	if err != nil {
		fmt.Println("Error converting id from strint to int")
	}
	item, err := db.GetQueries().OrderById(ctx, int64(ord))
	fmt.Println(item)
	if err != nil {
		fmt.Println(err)
		ctx.IndentedJSON(404, gin.H{"code": "PAGE_NOT_FOUND", "message": "Page not found"})
		return
	}
	ctx.JSON(200, item)
}

func CreateOrder(ctx *gin.Context) {
	bodyBytes, err := io.ReadAll(ctx.Request.Body)
	if err != nil {
		fmt.Println("Error reading request body : ", err)
	}
	var createOrderParams db.CreateOrdersParams
	err = json.Unmarshal(bodyBytes, &createOrderParams)
	if err != nil {
		fmt.Println("JSON Unmarshall error ")
	}
	err = db.GetQueries().CreateOrders(ctx, createOrderParams)
	if err != nil {
		fmt.Println("Error inserting Product in dB : ", err)
	}
}
