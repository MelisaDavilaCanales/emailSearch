package storage

import (
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

	// Betwen 400-500
	if resp.StatusCode >= 400 && resp.StatusCode < 500 {
		bodyContent, _ := io.ReadAll(resp.Body) // nolint:errcheck
		PrintLogs(bodyContent, resp.StatusCode)

		return nil, &models.NotFoundError{
			Message: "Not Found",
		}
	}

	// Whatever, at less between 200-299
	if resp.StatusCode < 200 || resp.StatusCode >= 300 {
		bodyContent, _ := io.ReadAll(resp.Body) // nolint:errcheck
		PrintLogs(bodyContent, resp.StatusCode)

		return nil, &models.InternalServerError{
			Message: "Internal Server Error",
		}
	}

	return resp, nil
}
