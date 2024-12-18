package storage

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"

	"backend/config"
	"backend/models"
)

func GetMail(id string) (*models.Email, *models.ResponseError) {
	var ResponseData *models.EmailDocResponse

	url := config.GET_EMAIL_API_URL + id

	res, err := DoRequest(http.MethodGet, url, nil)
	if err != nil {
		errResponse := models.NewResponseError(http.StatusNotFound, "Error making request", err)
		return nil, errResponse
	}

	if res.StatusCode == 400 {
		errResponse := models.NewResponseError(http.StatusNotFound, "Error getting email", fmt.Errorf("id no found %s", id))
		return nil, errResponse
	}
	defer res.Body.Close() //nolint:errcheck

	err = json.NewDecoder(res.Body).Decode(&ResponseData)
	if err != nil {
		errResponse := models.NewResponseError(http.StatusInternalServerError, "Error decoding response body", err)
		return nil, errResponse
	}

	return &ResponseData.Email, nil
}

func GetEmails(term, field string, from, max int) (*models.EmailHitsData, error) {
	var ResponseData models.EmailSearchResponse

	var query string

	if term == "" || field == "" {
		query = buildAllEmailsQuery(from, max)
	} else {
		query = buildFilteredEmailsQuery(term, field, from, max)
	}

	url := config.GET_EMAILS_API_URL

	res, err := DoRequest(http.MethodPost, url, strings.NewReader(query))
	if err != nil {
		return nil, err
	}

	defer res.Body.Close() //nolint:errcheck

	err = json.NewDecoder(res.Body).Decode(&ResponseData)
	if err != nil {
		return nil, fmt.Errorf("error decoding response body: %s", err)
	}

	return &ResponseData.EmailHitsData, nil
}

func buildAllEmailsQuery(from, max int) string {
	return fmt.Sprintf(`
		{
			"search_type": "matchall",
			"sort_fields": ["-date"],
			"from": %v,
			"max_results": %v,
			"_source": [
			"to", "from","date", "subject", "message_id"
			]
		}`, from, max)
}

func buildFilteredEmailsQuery(term, field string, from, max int) string {
	return fmt.Sprintf(`
		{
		"search_type": "match",
		"query": {
			"term": "%v",
			"field":"%v"
		},
		"sort_fields": ["-date"],
		"from": %v,
		"max_results": %v,
		"_source": [
			 "from", "to", "date","subject", "message_id"
		]
	}`, term, field, from, max)
}
