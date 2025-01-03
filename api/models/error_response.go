package models

import (
	"fmt"
	"strings"
)

// ResponseError is a struct that represents an error response
type ResponseError struct {
	StatusCode int    `json:"status_code"`
	Message    string `json:"message"`
	Err        error  `json:"error"` //nolint:tagliatelle
}

// NewResponseError is a function that creates a new ResponseError instance.
// Returns a pointer to a ResponseError instance.
func NewResponseError(statusCode int, message string, err error) *ResponseError {
	return &ResponseError{
		StatusCode: statusCode,
		Message:    message,
		Err:        err,
	}
}

// Error is a method that returns a string representation of the ResponseError instance.
func (re *ResponseError) Error() string {
	return fmt.Sprintf(`
	{
		"status_code":%v,
		"message":"%v",
		"error":"%v"
	}`, re.StatusCode, re.Message, strings.ReplaceAll(re.Err.Error(), "\"", "'"))
}
