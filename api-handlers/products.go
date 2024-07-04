package api_handlers

import (
	"encoding/json"
	"fmt"
	"io"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/sharmajsr/oms/db"
)

func GetProducts(ctx *gin.Context) {
	id1, isIdPresent := ctx.Params.Get("id")

	if !isIdPresent {
		fmt.Println("ID Not present in Get Query ")
	}
	id, err := strconv.Atoi(id1)
	if err != nil {
		fmt.Println("Error convert id to int")
	}
	item, err := db.GetQueries().ProductById(ctx, int64(id))
	if err != nil {
		fmt.Println("Error fetching item from DB ", err)
	}
	fmt.Println(item)
	ctx.JSON(200, item)

}

func CreateProducts(ctx *gin.Context) {
	bodyBytes, err := io.ReadAll(ctx.Request.Body)
	if err != nil {
		fmt.Println("Error reading request body : ", err)
	}
	var createProductParams db.CreateProductParams
	err = json.Unmarshal(bodyBytes, &createProductParams)
	if err != nil {
		fmt.Println("JSON Unmarshall error ")
	}
	err = db.GetQueries().CreateProduct(ctx, createProductParams)
	if err != nil {
		fmt.Println("Error inserting Product in dB : ", err)
	}
}
