package controllers

import (
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/MelisaDavilaCanales/emailSearch/api/models"
	"github.com/stretchr/testify/assert"
)

func TestGetEmails(t *testing.T) {
	tc := []struct {
		name          string
		emailHistData *models.EmailHitsData
		errResponse   error
		statusCode    int
	}{
		{
			name: "Get email success",
			emailHistData: &models.EmailHitsData{
				Total: models.Total{
					Value: 1,
				},
				Hits: []models.EmailHit{
					{
						Index: "emails",
						Type:  "_doc",
						ID:    "123",
						Score: 1,
						Email: models.EmailOfHit{
							Date:    "2021-06-01",
							From:    "juan@example.com",
							To:      "melisa@example.com",
							Subject: "Test Email",
						},
					},
				},
			},
			errResponse: nil,
			statusCode:  http.StatusOK,
		},
		{
			name: "Email not found",
			emailHistData: &models.EmailHitsData{
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
			InitMockGetEmails(tt.emailHistData, tt.errResponse)

			rr := httptest.NewRecorder()

			req, err := http.NewRequest(http.MethodGet, "/emails", nil)
			if err != nil {
				t.Fatal(err)
			}

			GetEmails(rr, req)

			DisableMockGetEmails()

			if tt.errResponse != nil {
				assert.Equal(t, http.StatusInternalServerError, tt.statusCode)
			} else {
				assert.Equal(t, http.StatusOK, tt.statusCode)
			}
		})
	}
}

func TestGetEmail(t *testing.T) {
	tc := []struct {
		name        string
		email       *models.Email
		errResponse error
		statusCode  int
	}{
		{
			name: "Get email success",
			email: &models.Email{
				MessageID: "123",
				Date:      "2021-06-01",
				From:      "maria vega",
				To:        "juan perez",
				Subject:   "Test Email",
				Content:   "Test email content",
			},
			errResponse: nil,
			statusCode:  http.StatusOK,
		},
		{
			name:        "Email not found",
			email:       nil,
			errResponse: errors.New("Not Found"),
			statusCode:  http.StatusNotFound,
		},
	}

	for _, tt := range tc {
		t.Run(tt.name, func(t *testing.T) {
			InitMockGetEmail(tt.email, tt.errResponse)

			// Create a new recorder to record the response
			rr := httptest.NewRecorder()

			req, err := http.NewRequest(http.MethodGet, "/emails/123", nil)
			if err != nil {
				t.Fatal(err)
			}

			GetEmail(rr, req)

			DisableMockGetEmail()

			if tt.errResponse != nil {
				assert.Equal(t, http.StatusNotFound, tt.statusCode)
			} else {
				assert.Equal(t, http.StatusOK, tt.statusCode)
			}
		})
	}
}
