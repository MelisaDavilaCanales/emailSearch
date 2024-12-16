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

func GetMail(id string) (*models.Email, *models.ResponseError) {

	var ResponseData *models.EmailDocResponse

	url := os.Getenv("ZINC_SEARCH_API_URL") + constant.EMAIL_INDEX_NAME + "/_doc/" + id

	res, err := DoRequest("GET", url, nil)
	if err != nil {
		return nil, models.NewResponseError(http.StatusInternalServerError, "Error making request", err)
	}
	defer res.Body.Close()

	err = json.NewDecoder(res.Body).Decode(&ResponseData)
	if err != nil {
		return nil, models.NewResponseError(http.StatusInternalServerError, "Error decoding response body", err)
	}

	return &ResponseData.Email, nil
}

func GetAllEmails(from, max int) (*models.EmailHitsData, *models.ResponseError) {

	var ResponseData models.EmailSearchResponse

	query := fmt.Sprintf(`
		{
			"search_type": "matchall",
			"sort_fields": ["-date"],
			"from": %v,
			"max_results": %v,
			"_source": [
			"to", "from","date", "subject", "message_id"
			]
		}`, from, max)

	url := os.Getenv("ZINC_SEARCH_API_URL") + constant.EMAIL_INDEX_NAME + "/_search"

	res, err := DoRequest("POST", url, bytes.NewBuffer([]byte(query)))
	if err != nil {
		return nil, models.NewResponseError(http.StatusInternalServerError, "Error making request", err)
	}
	defer res.Body.Close()

	err = json.NewDecoder(res.Body).Decode(&ResponseData)
	if err != nil {
		return nil, models.NewResponseError(http.StatusInternalServerError, "Error decoding response body", err)
	}

	return &ResponseData.EmailHitsData, nil
}

func SearchEmails(term, field string, from, max int) (*models.EmailHitsData, *models.ResponseError) {

	var ResponseData models.EmailSearchResponse
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
		"_source": [
			 "from", "to", "date","subject", "message_id"
		]
	}`, term, field, from, max)

	url := os.Getenv("ZINC_SEARCH_API_URL") + constant.EMAIL_INDEX_NAME + "/_search"

	res, err := DoRequest("POST", url, strings.NewReader(query))
	if err != nil {
		return nil, models.NewResponseError(http.StatusInternalServerError, "Error making request", err)
	}

	defer res.Body.Close()

	err = json.NewDecoder(res.Body).Decode(&ResponseData)
	if err != nil {
		return nil, models.NewResponseError(http.StatusInternalServerError, "Error decoding response body", err)
	}

	return &ResponseData.EmailHitsData, nil
}
