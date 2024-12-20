package controllers

import (
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/MelisaDavilaCanales/emailSearch/api/models"
	"github.com/stretchr/testify/assert"
)

func TestGetPersons(t *testing.T) {
	tc := []struct {
		name           string
		personHistData *models.PersonHitsData
		errResponse    error
		statusCode     int
	}{
		{
			name: "Get persons success",
			personHistData: &models.PersonHitsData{
				Total: models.Total{
					Value: 1,
				},
				Hits: []models.PersonHit{
					{
						Index: "persons",
						Type:  "_doc",
						ID:    "123",
						Score: 1,
						Person: models.PersonOfHit{
							Name:  "steve rivera",
							Email: "steve@example.com",
						},
					},
				},
			},
			errResponse: nil,
			statusCode:  http.StatusOK,
		},
		{
			name: "Persons not found",
			personHistData: &models.PersonHitsData{
				Total: models.Total{
					Value: 0,
				},
				Hits: nil,
			},
			errResponse: errors.New("Not Found"),
			statusCode:  http.StatusInternalServerError,
		},
	}

	for _, tt := range tc {
		t.Run(tt.name, func(t *testing.T) {
			InitMockGetPersons(tt.personHistData, tt.errResponse)
			defer DisableMockGetPersons()

			req := httptest.NewRequest(http.MethodGet, "/persons", nil)
			w := httptest.NewRecorder()

			GetPersons(w, req)

			if tt.errResponse != nil {
				assert.Equal(t, http.StatusInternalServerError, tt.statusCode)
			} else {
				assert.Equal(t, http.StatusOK, tt.statusCode)
			}
		})
	}
}
