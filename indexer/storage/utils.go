package storage

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"
)

func DoRequest(method string, url string, data io.Reader) (*http.Response, error) {
	req, err := http.NewRequest(method, url, data)
	if err != nil {
		return nil, err
	}

	req.SetBasicAuth(os.Getenv("DB_USER"), os.Getenv("DB_PASSWORD"))
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("User-Agent", "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_4) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/81.0.4044.138 Safari/537.36")

	client := &http.Client{}

	res, err := client.Do(req)
	if err != nil {
		return res, err
	}

	return res, nil
}

func CreateIndex(indexName, indexDataStr string) error {
	isExist := checkIndexExists(indexName)
	if isExist {
		return nil
	}

	url := os.Getenv("ZINC_SEARCH_API_URL") + "index"

	indexData := strings.NewReader(indexDataStr)

	res, err := DoRequest("POST", url, indexData)
	if err != nil {
		return err
	}

	if res.StatusCode != http.StatusCreated {
		bodyBytes, err := io.ReadAll(res.Body)

		if err != nil {
			fmt.Println("Error reading response body")
		}

		bodyString := string(bodyBytes)

		return fmt.Errorf("status code: %d, response body: %s", res.StatusCode, bodyString)
	}

	closeErr := res.Body.Close()
	if closeErr != nil {
		fmt.Println("Error closing response body:", closeErr)
	}

	fmt.Println("Index created successfully")

	return nil
}

func checkIndexExists(indexName string) bool {
	url := os.Getenv("ZINC_SEARCH_API_URL") + "index/" + indexName
	_, err := DoRequest("GET", url, nil)

	return err == nil
}
