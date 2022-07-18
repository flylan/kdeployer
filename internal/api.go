package internal

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func Start() error {
	router := initRouter()
	return router.Run(":80")
}

func initRouter() *gin.Engine {
	router := gin.Default()
	router.GET("/", func(context *gin.Context) {
		context.String(http.StatusOK, "hello world")
	})
	return router
}
