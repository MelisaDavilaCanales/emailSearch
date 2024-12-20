package storage

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestBuildSort(t *testing.T) {
	tc := []struct {
		name             string
		sortField        string
		sortOrder        string
		sortFieldDefault string
		sortOrderDefault string
		expectedSort     string
	}{
		{
			name:             "Success case",
			sortField:        "name",
			sortOrder:        "asc",
			sortFieldDefault: "name",
			sortOrderDefault: "-",
			expectedSort:     "name",
		},
	}

	for _, tt := range tc {
		t.Run(tt.name, func(t *testing.T) {
			result := buildSort(tt.sortField, tt.sortOrder, tt.sortFieldDefault, tt.sortOrderDefault)

			fmt.Println("expectedSort: ", tt.expectedSort)
			fmt.Println("result: ", result)

			assert.Equal(t, tt.expectedSort, result)
		})
	}
}
