package api_handlers

import (
	"encoding/json"
	"fmt"
	"io"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/sharmajsr/oms/db"
)

func GetUser(ctx *gin.Context) {
	id1, isIdPresent := ctx.Params.Get("id")

	if !isIdPresent {
		fmt.Println("ID Not present in Get Query ")
	}
	id, err := strconv.Atoi(id1)
	if err != nil {
		fmt.Println("Error convert id to int")
	}
	item, err := db.GetQueries().UserById(ctx, int32(id))
	if err != nil {
		fmt.Println("Error fetching item from DB ", err)
	}
	fmt.Println(item)
	ctx.JSON(200, item)
}

func CreateUser(ctx *gin.Context) {
	bodyBytes, err := io.ReadAll(ctx.Request.Body)
	if err != nil {
		fmt.Println("Error reading request body : ", err)
	}
	var createUserParams db.CreateUserParams
	err = json.Unmarshal(bodyBytes, &createUserParams)
	if err != nil {
		fmt.Println("JSON Unmarshall error ")
	}
	err = db.GetQueries().CreateUser(ctx, createUserParams)
	if err != nil {
		fmt.Println("Error inserting Product in dB : ", err)
	}
}
