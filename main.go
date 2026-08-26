package main

import (
	"net/http"

	"test-generator/generator"

	"github.com/gin-gonic/gin"
)

func HandlerNumber(c *gin.Context) {
	var number generator.Variable
	if err := c.ShouldBindYAML(&number); err != nil {
		ErrorResponse(c, http.StatusFailedDependency, InvalidInput)
		c.Abort()
		return
	}
	start := number.Min
	end := number.Max

	genNumber, error := generator.GenerateIntNumber(start, end)
	if error != nil {
		InternalServerErrorResponse(c)
		c.Abort()
		return
	}
	SuccessResponse(c, http.StatusOK, genNumber)
}

func main() {
	router := gin.New()

	router.POST("/number", HandlerNumber)

	router.Run(":3001")
}
