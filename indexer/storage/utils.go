package storage

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strings"

	"github.com/MelisaDavilaCanales/emailSearch/indexer/config"
)

func TryConnectionAPI() error {
	_, err := DoRequest(http.MethodGet, config.TRY_CONNECTION_API_URL, nil)
	if err != nil {
		return err
	}

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

	return res, nil
}

// CreateIndex creates an index with the specified name and data.
func CreateIndex(indexName, indexDataStr string) error {
	isExist := checkIndexExists(indexName)
	if isExist {
		return nil
	}

	url := config.CREATE_INDEX_API_URL

	indexData := strings.NewReader(indexDataStr)

	res, err := DoRequest(http.MethodPost, url, indexData)
	if err != nil {
		return fmt.Errorf("create index request %w", err)
	}

	if res.StatusCode != http.StatusCreated {
		return fmt.Errorf("create index status: %s", res.Status)
	}
	defer res.Body.Close() // nolint: errcheck

	return nil
}

func checkIndexExists(indexName string) bool {
	url := config.CHECK_INDEX_EXISTS_API_URL + indexName
	_, err := DoRequest(http.MethodGet, url, nil)

	return err == nil
}

// PrintLogs prints the response logs, can be used for debugging.
func PrintLogs(resp *http.Response) {
	var bodyBuffer bytes.Buffer
	tee := io.TeeReader(resp.Body, &bodyBuffer)

	bodyContent, readErr := io.ReadAll(tee)
	if readErr != nil {
		fmt.Println("read response body %w", readErr)
	}

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
