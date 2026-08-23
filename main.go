package main

import (
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
		if form.Start > form.End {
			c.JSON(400, gin.H{"error": "start must be less than end"})
			c.Abort()
			return
		}
		c.Set("form", form)
		c.Next()
	}
}

func main() {
	router := gin.New()
	router.GET("/int", ValidateMiddleware(), func(c *gin.Context) {
		formAny, okForm := c.Get("form")
		if !okForm {
			c.JSON(500, gin.H{"error": "form not found"})
			c.Abort()
			return
		}
		form, ok := formAny.(Form)
		if !ok {
			c.JSON(500, gin.H{"error": "invalid form"})
			c.Abort()
			return
		}
		start := form.Start
		end := form.End
		intNum, error := generator.GenerateIntNumber(start, end)
		if error != nil {
			c.JSON(500, gin.H{"error": error.Error()})
			c.Abort()
			return
		}
		c.JSON(200, gin.H{"int": intNum})
	})

	router.Run(":3001")
}
