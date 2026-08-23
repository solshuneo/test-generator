package main

import (
	"fmt"

	"test-generator/generator"

	"github.com/gin-gonic/gin"
)

type Form struct {
	Start int `json:"start" binding:"required"`
	End   int `json:"end" binding:"required"`
}

func ValidateMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		var form Form
		if err := c.ShouldBindJSON(&form); err != nil {
			c.JSON(400, gin.H{"error": err.Error()})
			c.Abort()
			return
		}
		c.Set("form", form)
		fmt.Printf("joined")
		c.Next()
		fmt.Printf("joined1")
	}
}

func main() {
	router := gin.New()
	router.GET("/int", ValidateMiddleware(), func(c *gin.Context) {
		form, _ := c.Get("form")
		intNum := generator.GenerateIntNumber(form.(Form).Start, form.(Form).End)
		c.JSON(200, gin.H{"int": intNum})
	})

	router.Run(":3001")
}
