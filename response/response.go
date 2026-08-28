package response

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

const (
	InvalidInput      = "invalid input"
	LeftLessThanRight = "left must be less than or equal to right"
)

type Response struct {
	Data    any    `json:"data,omitempty"`
	Message string `json:"message,omitempty"`
}

func InternalServerErrorResponse(c *gin.Context) Response {
	c.JSON(http.StatusInternalServerError, Response{
		Message: "Internal Server Error",
	})
	return Response{}
}

func ErrorResponse(c *gin.Context, code int, message string) Response {
	c.JSON(code, Response{
		Message: message,
	})
	return Response{}
}

func SuccessResponse(c *gin.Context, code int, data any) Response {
	c.JSON(code, Response{
		Data: data,
	})
	return Response{}
}
