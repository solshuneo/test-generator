package main

import (
	"fmt"
	"net/http"

	"test-generator/generator"

	"github.com/gin-gonic/gin"
)

type Variable struct {
	Min int `yaml:"min"`
	Max int `yaml:"max"`
}

func ValidateMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		var number Variable
		fmt.Printf("Here")
		if err := c.ShouldBindYAML(&number); err != nil {
			c.JSON(http.StatusBadRequest, Response{
				Message: InvalidInput,
			})
			c.Abort()
			return
		}
		c.Set("number", number)
		fmt.Printf("Here1")
		c.Next()
	}
}

func main() {
	router := gin.New()
	router.POST("/number", ValidateMiddleware(), func(c *gin.Context) {
		numberAny, okNumber := c.Get("number")
		if !okNumber {
			c.JSON(500, Response{
				Message: InternalServerError,
			})
			c.Abort()
			return
		}
		number, ok := numberAny.(Variable)
		if !ok {
			c.JSON(500, Response{
				Message: InternalServerError,
			})
			c.Abort()
			return
		}
		fmt.Printf("start: %d, end: %d\n", number.Min, number.Max)
		start := number.Min
		end := number.Max

		genNumber, error := generator.GenerateIntNumber(start, end)
		if error != nil {
			c.JSON(500, Response{
				Message: InternalServerError,
			})
			c.Abort()
			return
		}
		c.JSON(200, Response{
			Data: genNumber,
		})
	})

	router.Run(":3001")
}
