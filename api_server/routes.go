package api_server

import (
	"net/http"

	"github.com/gin-gonic/gin"
	api_handlers "github.com/sharmajsr/oms/api-handlers"
)

type routeEntry struct {
	path   string
	hFunc  gin.HandlerFunc
	method string
}

type RouteTable map[string][]routeEntry

var ServiceRoutes RouteTable = RouteTable{
	"/orders": []routeEntry{
		{"orders/:id", api_handlers.GetOrder, http.MethodGet},
		{"orders", api_handlers.CreateOrder, http.MethodPost},
	},
	"/products": []routeEntry{
		{"products/:id", api_handlers.GetProducts, http.MethodGet},
		{"products", api_handlers.CreateProducts, http.MethodPost},
	},
	"/user": []routeEntry{
		{"user/:id", api_handlers.GetUser, http.MethodGet},
		{"user", api_handlers.CreateUser, http.MethodPost},
	},
}

func SetAllRoutes(router *gin.Engine) {

	// router.Use()

	for _, routeEntries := range ServiceRoutes {
		for _, routeEntry := range routeEntries {
			router.Handle(routeEntry.method, routeEntry.path, routeEntry.hFunc)
		}
	}
}
