package storage

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/MelisaDavilaCanales/emailSearch/api/config"
	"github.com/MelisaDavilaCanales/emailSearch/api/models"
)

// DoRequest sends an HTTP request to the specified URL with the specified method and data.
func DoRequest(method string, url string, data io.Reader) (*http.Response, error) {
	req, err := http.NewRequest(method, url, data)
	if err != nil {
		return nil, err
	}

	setHeaders(req)

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return resp, err
	}

	var bodyBuffer bytes.Buffer
	tee := io.TeeReader(resp.Body, &bodyBuffer)

	bodyContent, readErr := io.ReadAll(tee)
	if readErr != nil {
		return resp, fmt.Errorf("reading response body: %s", readErr)
	}

	// printLogs(resp, bodyContent)

	resp.Body = io.NopCloser(&bodyBuffer)

	if resp.StatusCode >= 400 && resp.StatusCode < 500 {
		return resp, &models.NotFoundError{
			Message: string(bodyContent),
		}
	}

	if resp.StatusCode < 200 || resp.StatusCode >= 300 {
		return resp, &models.InternalServerError{
			Message: "internalServerError",
		}
	}

	// Si el código es exitoso (200-299), no hacemos nada y simplemente retornamos la respuesta.
	return resp, nil
}

func setHeaders(req *http.Request) {
	req.SetBasicAuth(config.DB_USER, config.DB_PASSWORD)
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("User-Agent", "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_4) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/81.0.4044.138 Safari/537.36")
}

func printLogs(resp *http.Response, bodyContent []byte) {
	fmt.Println("=========================================")
	fmt.Println("Response StatusCode:", resp.StatusCode)

	var jsonBody interface{}
	if jsonErr := json.Unmarshal(bodyContent, &jsonBody); jsonErr == nil {
		prettyJSON, err := json.MarshalIndent(jsonBody, "", "  ")
		if err != nil {
			fmt.Println("Response Body (as string):", string(bodyContent))
		}

		fmt.Println("Response Body (as JSON):", string(prettyJSON))
	}

	fmt.Println("=========================================")
}

func buildSort(sortField, sortOrder string) string {

	if sortField == "" {
		sortField = "@timestamp"
	}

	if sortOrder == "desc" || sortOrder == "" {
		sortOrder = "-"
	} else {
		sortOrder = ""
	}

	sort := fmt.Sprintf(`%s%s`, sortOrder, sortField)

	return sort
}
