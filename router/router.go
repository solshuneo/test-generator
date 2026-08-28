package router

import (
	"test-generator/handler"

	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {
	router := gin.Default()

	router.POST("/number", handler.HandlerNumber)

	return router
}
