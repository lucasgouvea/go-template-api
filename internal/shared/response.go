package shared

import (
	"os"

	"github.com/gin-gonic/gin"
)

type IResponse[T any] interface {
	GetData() []T
}

type Response struct {
	Data []any `json:"data"`
}

type ErrorResponse struct {
	Errors []any `json:"error"`
}

func sendResponse(context *gin.Context, status int, response any) {
	if os.Getenv("ENVIRONMENT") != "PRODUCTION" {
		context.IndentedJSON(status, response)
	} else {
		context.JSON(status, response)
	}
}

func HandleResponse(context *gin.Context, status int, data []any) {
	var response = Response{Data: data}
	sendResponse(context, status, response)
}

func HandleErrorResponse(context *gin.Context, status int, data []any) {
	var errorResponse = ErrorResponse{Errors: data}
	sendResponse(context, status, errorResponse)
}
