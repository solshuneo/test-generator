package main

import (
	"test-generator/generator"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.New()

	router.GET("/int", func(c *gin.Context) {
		intNum := generator.GenerateIntNumber(0, 100)
		c.JSON(200, gin.H{"int": intNum})
	})

	router.GET("/float", func(c *gin.Context) {
		floatNum := generator.FloatWithPrecision(generator.GenerateFloatNumber(0.0, 100.0), 3)
		c.JSON(200, gin.H{"float": floatNum})
	})

	router.Run(":3001")
}
