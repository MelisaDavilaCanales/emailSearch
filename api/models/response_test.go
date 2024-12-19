package models

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewResponse(t *testing.T) {
	tc := []struct {
		name     string
		message  string
		data     interface{}
		expected *Response
	}{
		{
			name:     "Response with string data",
			message:  "Success",
			data:     "Example data",
			expected: &Response{Message: "Success", Data: "Example data"},
		},
		{
			name:     "Response with nil data",
			message:  "Data not found",
			data:     nil,
			expected: &Response{Message: "Data not found", Data: nil},
		},
	}

	for _, tt := range tc {
		t.Run(tt.name, func(t *testing.T) {
			response := NewResponse(tt.message, tt.data)
			assert.Equal(t, tt.expected, response)
		})
	}
}

func TestNewEmailsResponseData(t *testing.T) {
	tc := []struct {
		name       string
		totalPages int
		page       int
		pageSize   int
		emails     []EmailSummary
		expected   *EmailsResponseData
	}{
		{
			name:       "Emails response with an email summary",
			totalPages: 5,
			page:       1,
			pageSize:   10,
			emails: []EmailSummary{
				{
					Id:      "2e7hSmG5bO3",
					Date:    "2000-12-01T06:09:00-08:00",
					From:    "person1@example.com",
					To:      "person2@example.com",
					Subject: "Test Subject",
				},
			},
			expected: &EmailsResponseData{
				TotalPages: 5,
				Page:       1,
				PageSize:   10,
				Emails: []EmailSummary{
					{
						Id:      "2e7hSmG5bO3",
						Date:    "2000-12-01T06:09:00-08:00",
						From:    "person1@example.com",
						To:      "person2@example.com",
						Subject: "Test Subject",
					},
				},
			},
		},
		{
			name:       "Emails response with no data",
			totalPages: 0,
			page:       0,
			pageSize:   0,
			emails:     nil,
			expected:   &EmailsResponseData{TotalPages: 0, Page: 0, PageSize: 0, Emails: nil},
		},
	}

	for _, tt := range tc {
		t.Run(tt.name, func(t *testing.T) {
			emailsResponse := NewEmailsResponseData(tt.totalPages, tt.page, tt.pageSize, tt.emails)
			assert.Equal(t, tt.expected, emailsResponse)
		})
	}
}

func TestNewPersonResponseData(t *testing.T) {
	tc := []struct {
		name       string
		totalPages int
		page       int
		pageSize   int
		persons    []Person
		expected   *PersonResponseData
	}{
		{
			name:       "Persons response with data",
			totalPages: 3,
			page:       1,
			pageSize:   20,
			persons:    []Person{{Name: "Martin Zepernick", Email: "martin@example.com"}},
			expected:   &PersonResponseData{TotalPages: 3, Page: 1, PageSize: 20, Persons: []Person{{Name: "Martin Zepernick", Email: "martin@example.com"}}},
		},
		{
			name:       "Persons response with no data",
			totalPages: 0,
			page:       0,
			pageSize:   0,
			persons:    nil,
			expected:   &PersonResponseData{TotalPages: 0, Page: 0, PageSize: 0, Persons: nil},
		},
	}

	for _, tt := range tc {
		t.Run(tt.name, func(t *testing.T) {
			personsResponse := NewPersonResponseData(tt.totalPages, tt.page, tt.pageSize, tt.persons)
			assert.Equal(t, tt.expected, personsResponse)
		})
	}
}
