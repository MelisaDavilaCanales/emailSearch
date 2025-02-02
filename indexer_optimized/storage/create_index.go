package storage

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/MelisaDavilaCanales/emailSearch/indexer/config"
)

// CreateIndex creates an index with the specified name and data.
func CreateIndex(indexName, indexDataStr string) (string, error) {
	isExist := checkIndexExists(indexName)
	if isExist {
		return fmt.Sprintf("Index %s already exists", indexName), nil
	}

	url := config.CREATE_INDEX_API_URL

	indexData := strings.NewReader(indexDataStr)

	_, err := DoRequest(http.MethodPost, url, indexData) // nolint: errcheck
	if err != nil {
		return "", fmt.Errorf("create index request %w", err)
	}

	return "Index created successfully", nil
}

func checkIndexExists(indexName string) bool {
	url := config.CHECK_INDEX_EXISTS_API_URL + indexName
	_, err := DoRequest(http.MethodGet, url, nil) // nolint: errcheck

	return err == nil
}
