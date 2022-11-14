package shared

import (
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
	Validator "github.com/go-playground/validator/v10"
)

type Model[T any] struct {
	Data T
	Hash string
}

type ErrorMessage struct {
	Field   string `json:"field"`
	Message string `json:"message"`
}

func HandleRequestError(error error, context *gin.Context) {
	if error.Error() == "EOF" {
		context.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"errors": "Unexpected JSON payload"})
	}
	var validationErrors Validator.ValidationErrors
	if errors.As(error, &validationErrors) {
		errorMessages := make([]ErrorMessage, len(validationErrors))
		for i, fe := range validationErrors {
			errorMessages[i] = ErrorMessage{Field: fe.Field(), Message: GetErrorMessage(fe)}
		}
		context.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"errors": errorMessages})
	}
}

func GetErrorMessage(fe Validator.FieldError) string {
	switch fe.Tag() {
	case "required":
		return "This field is required"
	case "lte":
		return "Should be less than " + fe.Param()
	case "gte":
		return "Should be greater than " + fe.Param()
	}
	return "Unknown error"
}
