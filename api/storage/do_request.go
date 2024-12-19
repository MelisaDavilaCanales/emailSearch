package storage

import (
	"bytes"
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

	// PrintLogs(resp, bodyContent)

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

	return resp, nil
}

func setHeaders(req *http.Request) {
	req.SetBasicAuth(config.DB_USER, config.DB_PASSWORD)
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("User-Agent", "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_4) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/81.0.4044.138 Safari/537.36")
}
