package shared_response

import (
	"github.com/gin-gonic/gin"
)

type IResponseError interface {
	GetErrors() []BaseError
}

type BaseError struct {
	Description string `json:"description"`
}

type ResponseError struct {
	Errors []BaseError `json:"errors"`
}

func (response ResponseError) GetErrors() []BaseError {
	return response.Errors
}

func NewError(descriptions []string) IResponseError {
	responseError := new(ResponseError)
	responseError.Errors = []BaseError{}
	for _, description := range descriptions {
		responseError.Errors = append(responseError.Errors, BaseError{Description: description})
	}
	return *responseError
}

func SendError(context *gin.Context, status int, responseError IResponseError) {
	sendResponse(context, status, responseError)
}
