package storage

import (
	"encoding/json"
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

	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}

	if resp.StatusCode != http.StatusOK {
		bodyContent, err := io.ReadAll(resp.Body)
		if err != nil {
			return nil, fmt.Errorf("request error, status code: %d", resp.StatusCode)
		}

		return nil, fmt.Errorf("staus code %d: %s", resp.StatusCode, string(bodyContent))
	}

	return resp, nil
}

// PrintLogs prints the response status code and body content, can be used for debugging.
func PrintLogs(bodyContent []byte, status int) {
	fmt.Println("=========================================")
	fmt.Println("Response StatusCode:", status)

	var jsonBody interface{}
	if jsonErr := json.Unmarshal(bodyContent, &jsonBody); jsonErr == nil {
		prettyJSON, err := json.MarshalIndent(jsonBody, "", "  ")
		if err != nil {
			fmt.Println("Response Body (as string):", string(bodyContent))
		}

		fmt.Println("Response Body:", string(prettyJSON))
	}

	fmt.Println("=========================================")
}
