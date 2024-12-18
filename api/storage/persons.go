package storage

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"

	"backend/config"
	"backend/models"
)

// sort=name& =asc
func GetPersons(term, field string, from, max int) (*models.PersonHitsData, error) {

	var ResponseData models.PersonSearchResponse
	var query string

	if term == "" || field == "" {
		query = buildAllPersonsQuery(from, max)
	} else {
		query = buildFilteredPersonsQuery(term, field, from, max)
	}

	url := config.GET_PERSONS_API_URL
	res, err := DoRequest(http.MethodPost, url, strings.NewReader(query))
	if err != nil {
		return nil, fmt.Errorf("error making request: %s", err)
	}

	defer res.Body.Close()

	err = json.NewDecoder(res.Body).Decode(&ResponseData)
	if err != nil {
		return nil, fmt.Errorf("error decoding response body: %s", err)
	}

	return &ResponseData.PersonHitsData, nil
}

func buildAllPersonsQuery(from, max int) string {
	return fmt.Sprintf(`
		{
			"search_type": "matchall",
			"sort_fields": ["-date"],
			"from": %v,
			"max_results": %v,
			"_source": [
			"name", "email"
			]
		}`, from, max)
}

func buildFilteredPersonsQuery(term, field string, from, max int) string {
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
			 "name", "email"
		]
	}`, term, field, from, max)
}
