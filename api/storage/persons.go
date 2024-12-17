package storage

import (
	"backend/constant"
	"backend/models"
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"strings"
)

func GetAllPersons(from, max int) (*models.PersonHitsData, *models.ResponseError) {

	var ResponseData models.PersonSearchResponse
	//sort=name& =asc
	query := fmt.Sprintf(`
		{
			"search_type": "matchall",
			"sort_fields": ["-date"],
			"from": %v,
			"max_results": %v,
			"_source": ["name", "email"]
		}`, from, max)

	url := os.Getenv("ZINC_SEARCH_API_URL") + constant.PERSON_INDEX_NAME + "/_search"

	res, err := DoRequest("POST", url, bytes.NewBuffer([]byte(query)))
	if err != nil {
		return nil, models.NewResponseError(http.StatusInternalServerError, "Error making request", err)
	}
	defer res.Body.Close()

	err = json.NewDecoder(res.Body).Decode(&ResponseData)
	if err != nil {
		return nil, models.NewResponseError(http.StatusInternalServerError, "Error decoding response body", err)
	}

	return &ResponseData.PersonHitsData, nil
}

func SearchPersons(term, field string, from, max int) (*models.PersonHitsData, *models.ResponseError) {

	var ResponseData models.PersonSearchResponse

	query := fmt.Sprintf(`
		{
		"search_type": "match",
		"query": {
			"term": "%v",
			"field":"%v"
		},
		"sort_fields": ["-date"],
		"from": %v,
		"max_results": %v,
		"_source": ["name", "email"]
	}`, term, field, from, max)

	url := os.Getenv("ZINC_SEARCH_API_URL") + constant.PERSON_INDEX_NAME + "/_search"

	res, err := DoRequest("POST", url, strings.NewReader(query))
	if err != nil {
		return nil, models.NewResponseError(http.StatusInternalServerError, "Error making request", err)
	}

	defer res.Body.Close()

	err = json.NewDecoder(res.Body).Decode(&ResponseData)
	if err != nil {
		return nil, models.NewResponseError(http.StatusInternalServerError, "Error decoding response body", err)
	}

	return &ResponseData.PersonHitsData, nil
}
