package models

import (
	"fmt"
	"strings"
)

type ResponseError struct {
	StatusCode int    `json:"status_code"`
	Message    string `json:"message"`
	Err        error  `json:"error"` //nolint:tagliatelle
}

func NewResponseError(statusCode int, message string, err error) *ResponseError {
	return &ResponseError{
		StatusCode: statusCode,
		Message:    message,
		Err:        err,
	}
}

func (re *ResponseError) Error() string {
	return fmt.Sprintf(`
	{
		"status":%v,
		"message":"%v",
		"error":"%v"
	}`, re.StatusCode, re.Message, strings.ReplaceAll(re.Err.Error(), "\"", "'"))
}
