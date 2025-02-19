package storage

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/MelisaDavilaCanales/emailSearch/indexer/config"
	"github.com/MelisaDavilaCanales/emailSearch/indexer/models"
)

// validBulkData defines the valid data types for sending a bulkV2 request to the ZincSearch API.
type validBulkData interface {
	*models.PersonBulkData | *models.EmailBulkData
}

// SendBulk sends data in bulkV2 format via a POST request to the ZincSearch API.
func SendBulk[T validBulkData](bulk T) error {
	url := config.SEND_BULK_API_URL

	switch any(bulk).(type) {
	case *models.PersonBulkData:
		fmt.Println("Sending the bulk of Persons ... Please wait")
	}

	bulkData, err := json.Marshal(bulk)
	if err != nil {
		return fmt.Errorf("marshal bulk: %w", err)
	}

	resp, err := DoRequest(http.MethodPost, url, bytes.NewBuffer(bulkData))
	if err != nil {
		return fmt.Errorf("send bulk: %w", err)
	}
	defer resp.Body.Close() // nolint: errcheck

	switch any(bulk).(type) {
	case *models.PersonBulkData:
		fmt.Println("Send bulk of Persons completed")
	case *models.EmailBulkData:
		fmt.Println("Send bulk of Emails completed")
	}

	return nil
}
