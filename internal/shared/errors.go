package shared

import (
	"encoding/json"
	"errors"
	"net/http"
	"strings"

	"github.com/go-playground/validator/v10"
)

type HttpError struct {
	Status      int    `json:"status"`
	Description string `json:"description"`
}

func GetHttpError(err error) *HttpError {

	if httpError := parseInvalidJSONErr(err); httpError != nil {
		return httpError
	}

	if httpError := parseInvalidPayloadErr(err); httpError != nil {
		return httpError
	}

	return &HttpError{Status: http.StatusInternalServerError, Description: "Internal error"}
}

func parseInvalidJSONErr(err error) *HttpError {
	var jsonSyntaxError *json.SyntaxError
	if errors.As(err, &jsonSyntaxError) {
		return &HttpError{Status: http.StatusBadRequest, Description: err.Error()}
	}
	if err.Error() == "EOF" {
		description := "Unexpected JSON payload"
		return &HttpError{Status: http.StatusBadRequest, Description: description}
	}
	return nil
}

func parseInvalidPayloadErr(err error) *HttpError {
	var unmarshalTypeError *json.UnmarshalTypeError
	if errors.As(err, &unmarshalTypeError) {
		var description = "Payload field " + unmarshalTypeError.Field + "should be of type" + unmarshalTypeError.Type.Name()
		return &HttpError{Status: http.StatusBadRequest, Description: description}
	}

	var validationErrors validator.ValidationErrors
	if errors.As(err, &validationErrors) {
		var description = "Payload field '" + strings.ToLower(validationErrors[0].Field()) + "' failed for tag '" + validationErrors[0].Tag() + "'"
		return &HttpError{Status: http.StatusBadRequest, Description: description}
	}
	return nil
}
