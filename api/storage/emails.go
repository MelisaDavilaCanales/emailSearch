package storage

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strings"

	"github.com/MelisaDavilaCanales/emailSearch/api/config"
	"github.com/MelisaDavilaCanales/emailSearch/api/models"
)

var (
	DoRequestFunc = DoRequest
)

func InitMock(httpResponse *http.Response, err error) {
	DoRequestFunc = func(_ string, _ string, _ io.Reader) (*http.Response, error) {
		return httpResponse, err
	}
}

func DisableMock() {
	DoRequestFunc = DoRequest
}

func GetMail(id string) (*models.Email, error) {
	var (
		ResponseData *models.EmailDocResponse
		url          string
	)

	url = config.GET_EMAIL_API_URL + id

	res, err := DoRequestFunc(http.MethodGet, url, nil)
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
	var ResponseData models.EmailSearchResponse

	url := config.GET_EMAILS_API_URL

	query := buildEmailQuery(params)
	fmt.Println(query)

	res, err := DoRequestFunc(http.MethodPost, url, strings.NewReader(query))
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

func buildEmailQuery(params models.SearchParams) string {
	sort := buildSort(params.SortField, params.SortOrder, "date", "-")

	if params.SearchTerm == "" {
		return fmt.Sprintf(`
		{
			"search_type": "matchall",
			"sort_fields": ["%s"],
			"from": %d,
			"max_results": %d,
			"_source": [ "to", "from","date", "subject"]
		}`, sort, params.ResultsFrom, params.MaxResults)
	}

	if params.SearchField == "" {
		params.SearchField = "_all"
	}

	return fmt.Sprintf(`
		{
			"search_type": "match",
			"query": {
				"term": "%s",
				"field": "%s"
			},
			"sort_fields": ["%s"],
			"from": %d,
			"max_results": %d,
			"_source": [ "to", "from","date", "subject"]
		}`, params.SearchTerm, params.SearchField, sort, params.ResultsFrom, params.MaxResults)
}
