package main

import (
	"github.com/sharmajsr/oms/api_server"
	"github.com/sharmajsr/oms/db"
)

func main() {

	db.Init()

	api_server.RunServer()
}
