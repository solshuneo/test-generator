package main

import (
	"net/http"

	"test-generator/generator"

	"github.com/gin-gonic/gin"
)

func ValidateMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		var number generator.Variable
		if err := c.ShouldBindYAML(&number); err != nil {
			ErrorResponse(c, http.StatusFailedDependency, InvalidInput)
			c.Abort()
			return
		}
		c.Set("number", number)
		c.Next()
	}
}

func main() {
	router := gin.New()
	router.POST("/number", ValidateMiddleware(), func(c *gin.Context) {
		numberAny, okNumber := c.Get("number")
		if !okNumber {
			InternalServerErrorResponse(c)
			c.Abort()
			return
		}
		number, ok := numberAny.(generator.Variable)
		if !ok {
			InternalServerErrorResponse(c)
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
	})

	router.Run(":3001")
}
