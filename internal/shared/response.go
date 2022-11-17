package shared

import (
	"os"

	"github.com/gin-gonic/gin"
)

type IResponse[T any] interface {
	GetData() []T
}

type Response[T any] struct {
	Data []T `json:"data"`
}

func (response Response[T]) GetData() []T {
	return response.Data
}

func NewResponse[T any](schemas []T) IResponse[T] {
	response := new(Response[T])
	response.Data = schemas
	return *response
}

type ErrorResponse struct {
	Errors []any `json:"errors"`
}

func sendResponse(context *gin.Context, status int, data any) {
	if os.Getenv("ENVIRONMENT") != "PRODUCTION" {
		context.IndentedJSON(status, data)
	} else {
		context.JSON(status, data)
	}
}

func HandleResponse[T any](context *gin.Context, status int, response IResponse[T]) {
	sendResponse(context, status, response)
}

func HandleErrorResponse(context *gin.Context, status int, data []any) {
	var errorResponse = ErrorResponse{Errors: data}
	sendResponse(context, status, errorResponse)
}
