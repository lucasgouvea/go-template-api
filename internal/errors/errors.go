package errors

import "errors"

type HttpErr struct {
	Status      int    `json:"status"`
	Description string `json:"description"`
}

var ResourceNotFoundErr = errors.New("Resource not found in database.")
