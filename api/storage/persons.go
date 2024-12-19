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
		query = buildAllPersonsQuery(params.SortField, params.SortOrder, params.ResultsFrom, params.MaxResults)
	} else {
		query = buildFilteredPersonsQuery(params.SearchTerm, params.SearchField, params.SortField, params.SortOrder, params.ResultsFrom, params.MaxResults)
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

func buildAllPersonsQuery(sortField, sortOrder string, from, max int) string {

	sort := buildSort(sortField, sortOrder)

	query := fmt.Sprintf(`
		{
			"search_type": "matchall",
			"sort_fields": ["%s"],
			"from": %d,
			"max_results": %d,
			"_source": []
		}`, sort, from, max)

	fmt.Println(query)
	return query

}

func buildFilteredPersonsQuery(searchTerm, searchField, sortField, sortOrder string, from, max int) string {
	if searchField == "" {
		searchField = "_all"
	}

	sort := buildSort(sortField, sortOrder)

	query := fmt.Sprintf(`
		{
		"search_type": "match",
		"query": {
			"term": "%v",
			"field":"%v"
		},
		"sort_fields": ["%s"],
		"from": %d,
		"max_results": %d,
		"_source": []
	}`, searchTerm, searchField, sort, from, max)

	fmt.Println(query)
	return query
}
