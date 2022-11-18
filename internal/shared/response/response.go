package shared_response

import (
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

func New[T any](schemas []T) IResponse[T] {
	response := new(Response[T])
	response.Data = schemas
	return *response
}

func Send[T any](context *gin.Context, status int, response IResponse[T]) {
	sendResponse(context, status, response)
}
