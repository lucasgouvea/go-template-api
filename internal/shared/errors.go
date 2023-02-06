package shared

import (
	"encoding/json"
	"errors"
	"net/http"
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
	if err.Error() == "EOF" {
		description := "Unexpected JSON payload"
		return &HttpError{Status: http.StatusBadRequest, Description: description}
	}
	return nil
}

func parseInvalidPayloadErr(err error) *HttpError {
	var unmarshalTypeError *json.UnmarshalTypeError
	if errors.As(err, &unmarshalTypeError) {
		var description = "Field " + unmarshalTypeError.Field + "should be of type" + unmarshalTypeError.Type.Name()
		return &HttpError{Status: http.StatusBadRequest, Description: description}
	}
	return nil
}
