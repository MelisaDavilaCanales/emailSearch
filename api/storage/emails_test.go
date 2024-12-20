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

func TestGetEmail(t *testing.T) {

	tc := []struct {
		name         string
		httpResponse *http.Response
		errResponse  error
	}{
		{
			name: "Get email success",
			httpResponse: &http.Response{
				StatusCode: 200,
				Body: io.NopCloser(strings.NewReader(`{
							"_index": "emails",
							"_id": "123",
							"_source": {
								"message_id": "123",
								"date": "2021-06-01",
								"from": "Juan",
								"to": "Melisa",
								"subject": "Test",
								"cc": "",
								"mime_version": "",
								"content_type": "",
								"content_transfer_encoding": "",
								"bcc": "",
								"x_from": "",
								"x_to": "",
								"x_cc": "",
								"x_bcc": "",
								"x_folder": "",
								"x_origin": "",
								"x_file_name": "",
								"content": "Test email content"
							}
						}`)),
			},
			errResponse: nil,
		},
		{
			name: "Email not found",
			httpResponse: &http.Response{
				StatusCode: 404,
				Body: io.NopCloser(strings.NewReader(`{ 
						"error": "Not Found",
					}`)),
			},
			errResponse: errors.New("Not Found"),
		},
	}

	for _, tt := range tc {
		t.Run(tt.name, func(t *testing.T) {
			InitMock(tt.httpResponse, tt.errResponse)
			emailResponse, errResponse := GetMail("123")
			DisableMock()

			if tt.errResponse != nil {
				assert.Nil(t, emailResponse)
				assert.NotNil(t, errResponse)
			} else {
				assert.NotNil(t, emailResponse)
				assert.Nil(t, errResponse)
			}
		})
	}
}

func TestGetEmails(t *testing.T) {

	tc := []struct {
		name         string
		httpResponse *http.Response
		errResponse  error
		params       models.SearchParams
	}{
		{
			name: "Get all emails success",
			httpResponse: &http.Response{
				StatusCode: 200,
				Body: io.NopCloser(strings.NewReader(`{
						"hits": {
						"total": {"value": 1},
							"hits": [
								{
									"_index": "emails",
									"_type": "_doc",
									"_id": "123",
									"_score": 1,
									"_source": {
										"message_id": "123",
										"date": "2021-06-01",
										"from": "Juan",
										"to": "Melisa",
										"subject": "Test",
										"content": "Test"
									}
								}
							]
						}
					}`)),
			},
			errResponse: nil,
			params: models.SearchParams{
				SearchTerm:  "saludo",
				SearchField: "subject",
				SortField:   "date",
				SortOrder:   "asc",
				ResultsFrom: 0,
				MaxResults:  10,
			},
		},
		{
			name: "Emails not found",
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
			emailsResponse, errResponse := GetEmails(tt.params)
			DisableMock()

			if tt.errResponse != nil {
				assert.Nil(t, emailsResponse)
				assert.NotNil(t, errResponse)
			} else {
				assert.NotNil(t, emailsResponse)
				assert.Nil(t, errResponse)
			}
		})
	}
}
