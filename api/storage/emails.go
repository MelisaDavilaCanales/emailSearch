package storage

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"

	"github.com/MelisaDavilaCanales/emailSearch/api/config"
	"github.com/MelisaDavilaCanales/emailSearch/api/models"
)

func GetMail(id string) (*models.Email, error) {
	var (
		ResponseData *models.EmailDocResponse
		url          string
	)

	url = config.GET_EMAIL_API_URL + id

	res, err := DoRequest(http.MethodGet, url, nil)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close() //nolint:errcheck

	err = json.NewDecoder(res.Body).Decode(&ResponseData)
	if err != nil {
		return nil, fmt.Errorf("%s: decoding response body", err)
	}

	return &ResponseData.Email, nil
}

func GetEmails(params models.SearchParams) (*models.EmailHitsData, error) {
	var (
		ResponseData models.EmailSearchResponse
		query        string
		url          string
	)

	if params.SearchTerm == "" {
		query = buildAllEmailsQuery(params.ResultsFrom, params.MaxResults)
	} else {
		query = buildFilteredEmailsQuery(params.SearchTerm, params.SearchField, params.ResultsFrom, params.MaxResults)
	}

	url = config.GET_EMAILS_API_URL

	res, err := DoRequest(http.MethodPost, url, strings.NewReader(query))
	if err != nil {
		return nil, err
	}
	defer res.Body.Close() //nolint:errcheck

	err = json.NewDecoder(res.Body).Decode(&ResponseData)
	if err != nil {
		return nil, fmt.Errorf("%s: decoding response body", err)
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
	if field == "" {
		field = "_all"
	}

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
