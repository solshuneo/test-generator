package main

const (
	InternalServerError = "internal server error"
	InvalidInput        = "invalid input"
	LeftLessThanRight   = "left must be less than or equal to right"
)

type Response struct {
	Data    any    `json:"data,omitempty"`
	Message string `json:"message,omitempty"`
}
