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

	url := config.GET_PERSONS_API_URL

	res, err := DoRequestFunc(http.MethodPost, url, strings.NewReader(query))
	if err != nil {
		return nil, err
	}
	defer res.Body.Close() //nolint:errcheck

	err = json.NewDecoder(res.Body).Decode(&ResponseData)
	if err != nil {
		fmt.Printf("decoding response body: %s", err)
		return nil, fmt.Errorf("decoding response body")
	}

	return &ResponseData.PersonHitsData, nil
}

func buildPersonQuery(params models.SearchParams) string {
	sort := buildSort(params.SortField, params.SortOrder, "name", "")

	if params.SearchTerm == "" {
		return fmt.Sprintf(`
		{
			"search_type": "query",
			"query": {
				"match_all": {}
			},
			"sort": ["%s"],
			"from": %d,
			"size": %d,
			"_source": []
		}`, sort, params.ResultsFrom, params.MaxResults)
	}

	if params.SearchField == "" {
		params.SearchField = "_all"
	}

	return fmt.Sprintf(`
		{
			"search_type": "query",
			"query": {
				"match": {
					"%s": {
						"query": "\"%s\"",
						"operator": "AND"
					}
				}
			},
			"sort": ["%s"],
			"from": %d,
			"size": %d,
			"_source": []
		}`, params.SearchField, params.SearchTerm, sort, params.ResultsFrom, params.MaxResults)
}
