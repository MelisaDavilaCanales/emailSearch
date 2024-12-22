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

	email := &models.Email{
		ID:                      ResponseData.ID,
		MessageID:               ResponseData.Email.MessageID,
		Date:                    ResponseData.Email.Date,
		From:                    ResponseData.Email.From,
		To:                      ResponseData.Email.To,
		Subject:                 ResponseData.Email.Subject,
		Cc:                      ResponseData.Email.Cc,
		MimeVersion:             ResponseData.Email.MimeVersion,
		ContentType:             ResponseData.Email.ContentType,
		ContentTransferEncoding: ResponseData.Email.ContentTransferEncoding,
		Bcc:                     ResponseData.Email.Bcc,
		XFrom:                   ResponseData.Email.XFrom,
		XTo:                     ResponseData.Email.XTo,
		XCc:                     ResponseData.Email.XCc,
		XBcc:                    ResponseData.Email.XBcc,
		XFolder:                 ResponseData.Email.XFolder,
		XOrigin:                 ResponseData.Email.XOrigin,
		XFileName:               ResponseData.Email.XFileName,
		Content:                 ResponseData.Email.Content,
	}

	return email, nil
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
			"search_type": "querystring",
			"query": {
				"query_string": {
					"query": "%s:%s"
				}
			},
			"sort_fields": ["%s"],
			"from": %d,
			"max_results": %d,
			"_source": [ "to", "from","date", "subject"]
		}`, params.SearchField, params.SearchTerm, sort, params.ResultsFrom, params.MaxResults)
}
