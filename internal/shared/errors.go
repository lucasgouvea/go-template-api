package shared

import (
	"encoding/json"
	"errors"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	Validator "github.com/go-playground/validator/v10"
)

type ErrorMessage struct {
	Field   string `json:"field"`
	Message string `json:"message"`
}

func HandleRequestError(error error, context *gin.Context) {

	var data []any
	if error.Error() == "EOF" {
		data = append(data, gin.H{"errors": "Unexpected JSON payload"})
		HandleErrorResponse(context, http.StatusBadRequest, data)
		return
	}

	var unmarshalTypeError *json.UnmarshalTypeError
	if errors.As(error, &unmarshalTypeError) {
		var sb strings.Builder
		sb.WriteString("Should be of type ")
		sb.WriteString(unmarshalTypeError.Type.Name())
		var errorMessage = ErrorMessage{Field: unmarshalTypeError.Field, Message: sb.String()}
		data = append(data, errorMessage)
		HandleErrorResponse(context, http.StatusBadRequest, data)
		return
	}

	var validationErrors Validator.ValidationErrors
	if errors.As(error, &validationErrors) {
		errorMessages := make([]ErrorMessage, len(validationErrors))
		for i, fe := range validationErrors {
			var field = fe.Field()
			field = JoinCamelCaseWith_(field)
			errorMessages[i] = ErrorMessage{Field: strings.ToLower(field), Message: GetErrorMessage(fe)}
			data = append(data, errorMessages[i])
		}
		HandleErrorResponse(context, http.StatusBadRequest, data)
		return
	}

	data = append(data, error)
	HandleErrorResponse(context, http.StatusInternalServerError, data)
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
