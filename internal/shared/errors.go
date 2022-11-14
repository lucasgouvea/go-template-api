package shared

import (
	"encoding/json"
	"errors"
	"net/http"
	"os"
	"strings"

	"github.com/gin-gonic/gin"
	Validator "github.com/go-playground/validator/v10"
)

type ErrorMessage struct {
	Field   string `json:"field"`
	Message string `json:"message"`
}

func HandleRequestError(error error, context *gin.Context) {
	if error.Error() == "EOF" {
		HandleResponse(context, http.StatusBadRequest, gin.H{"errors": "Unexpected JSON payload"})
		return
	}

	var unmarshalTypeError *json.UnmarshalTypeError
	if errors.As(error, &unmarshalTypeError) {
		var sb strings.Builder
		sb.WriteString("Should be of type ")
		sb.WriteString(unmarshalTypeError.Type.Name())
		var errorMessage = ErrorMessage{Field: unmarshalTypeError.Field, Message: sb.String()}
		HandleResponse(context, http.StatusBadRequest, errorMessage)
		return
	}

	var validationErrors Validator.ValidationErrors
	if errors.As(error, &validationErrors) {
		errorMessages := make([]ErrorMessage, len(validationErrors))
		for i, fe := range validationErrors {
			errorMessages[i] = ErrorMessage{Field: strings.ToLower(fe.Field()), Message: GetErrorMessage(fe)}
		}
		HandleResponse(context, http.StatusBadRequest, errorMessages)
		return
	}

	HandleResponse(context, http.StatusInternalServerError, error)
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

func HandleResponse(context *gin.Context, status int, data any) {
	if os.Getenv("ENVIRONMENT") != "PRODUCTION" {
		context.IndentedJSON(status, data)
	} else {
		context.JSON(status, data)
	}

}
