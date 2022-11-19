package shared

import (
	"encoding/json"
	"errors"
	"net/http"
	"strconv"

	Response "go-api/internal/shared/response"

	"github.com/gin-gonic/gin"
	Validator "github.com/go-playground/validator/v10"
)

func handleQueryParamError(err error, context *gin.Context) bool {
	var numError *strconv.NumError
	if errors.As(err, &numError) {
		var expectedType string
		if numError.Func == "ParseInt" {
			expectedType = "integer"
		}
		if numError.Func == "ParseBool" {
			expectedType = "bool"
		}

		description := "Invalid query value '" + numError.Num + "', expected type to be " + expectedType
		Response.SendError(context, http.StatusBadRequest, Response.NewError([]string{description}))
		return true
	}

	var validationErrors Validator.ValidationErrors
	if errors.As(err, &validationErrors) {
		validationError := validationErrors[0]
		if validationError.Tag() == "min" {
			description := "Invalid query value '" + validationError.Namespace() + "', expected value to be higher than or equal to " + validationError.Param()
			Response.SendError(context, http.StatusBadRequest, Response.NewError([]string{description}))
		}

		return true
	}
	return false
}

func handleInvalidJSON(err error, context *gin.Context) bool {
	if err.Error() == "EOF" {
		description := "Unexpected JSON payload"
		Response.SendError(context, http.StatusBadRequest, Response.NewError([]string{description}))
		return true
	}
	return false
}

func handleInvalidJSONPayload(err error, context *gin.Context) bool {
	var unmarshalTypeError *json.UnmarshalTypeError
	if errors.As(err, &unmarshalTypeError) {
		var description = "Field " + unmarshalTypeError.Field + "should be of type" + unmarshalTypeError.Type.Name()
		Response.SendError(context, http.StatusBadRequest, Response.NewError([]string{description}))
		return true
	}
	return false
}

func HandleRequestError(err error, context *gin.Context) {

	if handleQueryParamError(err, context) {
		return
	}

	if handleInvalidJSON(err, context) {
		return
	}

	if handleInvalidJSONPayload(err, context) {
		return
	}

	Response.SendError(context, http.StatusInternalServerError, Response.NewError([]string{err.Error()}))
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
