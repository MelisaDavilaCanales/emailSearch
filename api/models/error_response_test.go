package models

import (
	"errors"
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewResponseError(t *testing.T) {
	tc := []struct {
		name     string
		status   int
		message  string
		err      error
		expected *ResponseError
	}{
		{
			name:     "Not Found Response Error",
			status:   http.StatusNotFound,
			message:  "Not found",
			err:      errors.New("not found"),
			expected: &ResponseError{StatusCode: http.StatusNotFound, Message: "Not found", Err: errors.New("not found")},
		},
	}

	for _, tt := range tc {
		t.Run(tt.name, func(t *testing.T) {
			responseErr := NewResponseError(tt.status, tt.message, tt.err)
			assert.Equal(t, tt.expected, responseErr)
		})
	}
}

func TestNewResponseError_Error(t *testing.T) {
	tc := []struct {
		name     string
		status   int
		message  string
		err      error
		expected string
	}{
		{
			name:    "Not Found Response Error",
			status:  http.StatusNotFound,
			message: "Not found",
			err:     errors.New("not found"),
			expected: `{
		"status_code":404,
		"message":"Not found",
		"error":"not found"
	}`,
		},
		{
			name:    "Internal Server Error",
			status:  http.StatusInternalServerError,
			message: "Internal server error",
			err:     errors.New("server failure"),
			expected: `{
		"status_code":500,
		"message":"Internal server error",
		"error":"server failure"
	}`,
		},
	}

	for _, tt := range tc {
		t.Run(tt.name, func(t *testing.T) {
			// Create the ResponseError instance
			responseErr := NewResponseError(tt.status, tt.message, tt.err)

			// Call the Error() method and check if the output matches the expected string
			assert.Equal(t, tt.expected, responseErr.Error())
		})
	}
}
