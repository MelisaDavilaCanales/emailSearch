package storage

import (
	"bytes"
	"fmt"
	"io"
	"net/http"

	"github.com/MelisaDavilaCanales/emailSearch/api/config"
	"github.com/MelisaDavilaCanales/emailSearch/api/models"
)

// DoRequest make a request to the database, make some validations and return a response.
func DoRequest(method string, url string, data io.Reader) (*http.Response, error) {
	req, err := http.NewRequest(method, url, data)
	if err != nil {
		return nil, err
	}

	req.SetBasicAuth(config.DB_USER, config.DB_PASSWORD)
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("User-Agent", "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_4) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/81.0.4044.138 Safari/537.36")

	client := &http.Client{}

	resp, err := client.Do(req)
	if err != nil {
		fmt.Println("request error:", err)

		return nil, &models.InternalServerError{
			Message: "Internal Server Error",
		}
	}

	var bodyBuffer bytes.Buffer
	tee := io.TeeReader(resp.Body, &bodyBuffer)

	bodyContent, readErr := io.ReadAll(tee)
	if readErr != nil {
		return resp, fmt.Errorf("reading response body: %s", readErr)
	}

	resp.Body = io.NopCloser(&bodyBuffer)

	// Betwen 400-500
	if resp.StatusCode >= 400 && resp.StatusCode < 500 {
		fmt.Printf("Status code: %d. /\n Response content: %v", resp.StatusCode, bodyContent)

		return nil, &models.NotFoundError{
			Message: "Not Found",
		}
	}

	// Whatever, at less between 200-299
	if resp.StatusCode < 200 || resp.StatusCode >= 300 {
		fmt.Printf("Status code: %d. /\n Response content: %v", resp.StatusCode, bodyContent)

		return nil, &models.InternalServerError{
			Message: "Internal Server Error",
		}
	}

	return resp, nil
}
