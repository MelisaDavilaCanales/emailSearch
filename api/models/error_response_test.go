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
