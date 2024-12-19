package storage

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"

	"github.com/MelisaDavilaCanales/emailSearch/api/config"
	"github.com/MelisaDavilaCanales/emailSearch/api/models"
)

func GetPersons(params models.SearchParams) (*models.PersonHitsData, error) {
	var (
		ResponseData models.PersonSearchResponse
		query        string
		url          string
	)

	if params.SearchTerm == "" {
		query = buildAllPersonsQuery(params.ResultsFrom, params.MaxResults)
	} else {
		query = buildFilteredPersonsQuery(params.SearchTerm, params.SearchField, params.ResultsFrom, params.MaxResults)
	}

	url = config.GET_PERSONS_API_URL

	res, err := DoRequest(http.MethodPost, url, strings.NewReader(query))
	if err != nil {
		return nil, err
	}
	defer res.Body.Close() //nolint:errcheck

	err = json.NewDecoder(res.Body).Decode(&ResponseData)
	if err != nil {
		return nil, fmt.Errorf("decoding response body: %s", err)
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
			 "name", "email"
		]
	}`, term, field, from, max)
}
