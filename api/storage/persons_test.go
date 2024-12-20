package storage

import (
	"errors"
	"io"
	"net/http"
	"strings"
	"testing"

	"github.com/MelisaDavilaCanales/emailSearch/api/models"
	"github.com/stretchr/testify/assert"
)

func TestGetPersons(t *testing.T) {
	tc := []struct {
		name         string
		httpResponse *http.Response
		errResponse  error
		params       models.SearchParams
	}{
		{
			name: "Get all persons success",
			httpResponse: &http.Response{
				StatusCode: 200,
				Body: io.NopCloser(strings.NewReader(`{
					"hits": {
						"total": {"value": 1},
						"hits": [
							{
								"_index": "persons",
								"_type": "_doc",
								"_id": "123",
								"_score": 1,
								"_source": {
									"name": "maria",
									"email": "maria@example.com"
								}
							}
						]
					}
				}`)),
			},
			errResponse: nil,
			params: models.SearchParams{
				SearchTerm:  "charles",
				SearchField: "name",
				SortField:   "name",
				SortOrder:   "asc",
				ResultsFrom: 0,
				MaxResults:  10,
			},
		},
		{
			name: "Persons not found",
			httpResponse: &http.Response{
				StatusCode: 404,
				Body: io.NopCloser(strings.NewReader(`{
					"error": "Not Found",
				}`)),
			},
			errResponse: errors.New("Not Found"),
			params:      models.SearchParams{},
		},
	}

	for _, tt := range tc {
		t.Run(tt.name, func(t *testing.T) {
			InitMock(tt.httpResponse, tt.errResponse)
			personsResponse, errResponse := GetPersons(tt.params)
			DisableMock()

			if tt.errResponse != nil {
				assert.Nil(t, personsResponse)
				assert.NotNil(t, errResponse)
			} else {
				assert.NotNil(t, personsResponse)
				assert.Nil(t, errResponse)
			}
		})
	}
}
