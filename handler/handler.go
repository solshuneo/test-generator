package handler

import (
	"net/http"

	"test-generator/generator"
	"test-generator/response"

	"github.com/gin-gonic/gin"
)

const (
	MinGreaterThanMax = "Min > Max"
)

func HandlerNumber(c *gin.Context) {
	var number generator.Variable
	if err := c.ShouldBindYAML(&number); err != nil {
		response.ErrorResponse(c, http.StatusFailedDependency, response.InvalidInput)
		c.Abort()
		return
	}
	start := number.Min
	end := number.Max

	genNumber, error := generator.GenerateIntNumber(start, end)
	if error != nil {
		response.ErrorResponse(c, http.StatusBadRequest, MinGreaterThanMax)
		c.Abort()
		return
	}
	response.SuccessResponse(c, http.StatusOK, genNumber)
}
