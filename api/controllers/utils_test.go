package controllers

import (
	"net/http"
	"testing"

	"github.com/MelisaDavilaCanales/emailSearch/api/models"
	"github.com/stretchr/testify/assert"
)

func TestGetQueryParams(t *testing.T) {
	// Crear el escenario de prueba para los diferentes parámetros de la URL
	tests := []struct {
		name           string
		url            string
		expectedParams models.QueryParams
	}{
		{
			name: "All params provided",
			url:  "/test?page=1&page_size=10&term=test&field=from&sort=date&order=desc",
			expectedParams: models.QueryParams{
				PageNumber:  "1",
				PageSize:    "10",
				SearchTerm:  "test",
				SearchField: "from",
				SortField:   "date",
				SortOrder:   "desc",
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			req, err := http.NewRequest(http.MethodGet, tt.url, nil)
			if err != nil {
				t.Fatal(err)
			}

			params := getQueryParams(req)

			assert.Equal(t, tt.expectedParams, params)
		})
	}
}
