package storage

import (
	"fmt"
	"io"
	"net/http"

	"github.com/MelisaDavilaCanales/emailSearch/indexer/config"
)

func TryConnectionAPI() error {
	_, err := DoRequest(http.MethodGet, config.TRY_CONNECTION_API_URL, nil)
	if err != nil {
		return err
	}

	fmt.Println("Connection to database successful.")

	return nil
}

// DoRequest sends an HTTP request to the specified URL with the specified method and data.
func DoRequest(method string, url string, data io.Reader) (*http.Response, error) {
	req, err := http.NewRequest(method, url, data)
	if err != nil {
		return nil, err
	}

	req.SetBasicAuth(config.DB_USER, config.DB_PASSWORD)
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("User-Agent", "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_4) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/81.0.4044.138 Safari/537.36")

	client := &http.Client{}

	res, err := client.Do(req)
	if err != nil {
		return nil, err
	}

	if res.StatusCode != http.StatusOK {
		responseBody, err := io.ReadAll(res.Body)
		if err != nil {
			return nil, fmt.Errorf("request error, status code: %d", res.StatusCode)
		}

		return nil, fmt.Errorf("staus code %d: %s", res.StatusCode, string(responseBody))
	}

	return res, nil
}
