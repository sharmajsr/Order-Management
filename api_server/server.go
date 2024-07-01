package api_server

import (
	"fmt"
	"net/http"
	"sync"

	"github.com/gin-gonic/gin"
)

var (
	OmsServer *OMSServer
	once      sync.Once
)

type OMSServer struct {
	srv    *http.Server
	router *gin.Engine
}

func init() {
	once.Do(func() {
		OmsServer = &OMSServer{}
	})
}
func RunServer() {

	port := 8080
	OmsServer.router = gin.Default()
	SetAllRoutes(OmsServer.router)
	OmsServer.srv = &http.Server{Addr: fmt.Sprintf(":%d", port), Handler: OmsServer.router}
	fmt.Printf("Order Management Server started at port %d", port)

	if err := OmsServer.srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
		fmt.Println("Server Closed")
	}
}
