package shared

import (
	"encoding/json"
	"errors"
	Errors "go-template-api/internal/errors"
	"net/http"
	"strings"

	"github.com/go-playground/validator/v10"
)

func GetHttpError(err error) *Errors.HttpErr {

	if HttpErr := parseInvalidJSONErr(err); HttpErr != nil {
		return HttpErr
	}

	if HttpErr := parseInvalidPayloadErr(err); HttpErr != nil {
		return HttpErr
	}

	if errors.Is(err, Errors.ResourceNotFoundErr) {
		return &Errors.HttpErr{Status: http.StatusNotFound, Description: err.Error()}
	}

	return &Errors.HttpErr{Status: http.StatusInternalServerError, Description: "Internal error"}
}

func parseInvalidJSONErr(err error) *Errors.HttpErr {
	var jsonSyntaxError *json.SyntaxError
	if errors.As(err, &jsonSyntaxError) {
		return &Errors.HttpErr{Status: http.StatusBadRequest, Description: err.Error()}
	}
	if err.Error() == "EOF" {
		description := "Unexpected JSON payload"
		return &Errors.HttpErr{Status: http.StatusBadRequest, Description: description}
	}
	return nil
}

func parseInvalidPayloadErr(err error) *Errors.HttpErr {
	var unmarshalTypeError *json.UnmarshalTypeError
	if errors.As(err, &unmarshalTypeError) {
		var description = "Payload field " + unmarshalTypeError.Field + "should be of type" + unmarshalTypeError.Type.Name()
		return &Errors.HttpErr{Status: http.StatusBadRequest, Description: description}
	}

	var validationErrors validator.ValidationErrors
	if errors.As(err, &validationErrors) {
		var description = "Payload field '" + strings.ToLower(validationErrors[0].Field()) + "' failed for tag '" + validationErrors[0].Tag() + "'"
		return &Errors.HttpErr{Status: http.StatusBadRequest, Description: description}
	}
	return nil
}
