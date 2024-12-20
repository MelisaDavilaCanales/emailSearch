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
	var ResponseData models.PersonSearchResponse

	query := buildPersonQuery(params)
	fmt.Println(query)

	url := config.GET_PERSONS_API_URL

	res, err := DoRequestFunc(http.MethodPost, url, strings.NewReader(query))
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

func buildPersonQuery(params models.SearchParams) string {
	sort := buildSort(params.SortField, params.SortOrder, "name", "")

	if params.SearchTerm == "" {
		return fmt.Sprintf(`
		{
			"search_type": "matchall",
			"sort_fields": ["%s"],
			"from": %d,
			"max_results": %d,
			"_source": []
		}`, sort, params.ResultsFrom, params.MaxResults)
	}

	if params.SearchField == "" {
		params.SearchField = "_all"
	}

	return fmt.Sprintf(`
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
		}`, params.SearchTerm, params.SearchField, sort, params.ResultsFrom, params.MaxResults)
}
