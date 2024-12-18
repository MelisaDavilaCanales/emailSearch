package storage

import (
	"backend/constant"
	"backend/models"
	"encoding/json"
	"fmt"
	"os"
	"strings"
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

	url := os.Getenv("ZINC_SEARCH_API_URL") + constant.PERSON_INDEX_NAME + "/_search"

	res, err := DoRequest("POST", url, strings.NewReader(query))
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
