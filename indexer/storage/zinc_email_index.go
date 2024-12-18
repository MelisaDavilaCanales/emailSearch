package storage

import (
	"fmt"

	"github.com/MelisaDavilaCanales/emailSearch/indexer/config"
	"github.com/MelisaDavilaCanales/emailSearch/indexer/constant"
)

func CreateEmailIndex() error {
	indexDataStr := buildEmailIndex()
	return CreateIndex(constant.EMAIL_INDEX_NAME, indexDataStr)
}

func buildEmailIndex() string {
	indexData := fmt.Sprintf(`{
		"name": "%v",
		"storage_type": "disk",
		"shard_num": %d,
		"mappings": {
			"properties": {
				"message_id": {
					"type": "text",
					"index": true,
					"sortable": true,
					"highlightable": true
				},
				"Date": {
					"type": "date",
					"format":"2006-01-02T15:04:05Z",
					"index": true,
					"sortable": true,
					"highlightable": true
				},
				"from": {
					"type": "text",
					"index": true,
					"sortable": true,
					"highlightable": true
				},
				"to": {
					"type": "text",
					"index": true,
					"sortable": true,
					"highlightable": true
				},
				"subject": {
					"type": "text",
					"index": true,
					"sortable": true,
					"highlightable": true
				},
				"cc": {
					"type": "text",
					"index": true,
					"sortable": true,
					"highlightable": true
				},
				"mime_version": {
					"type": "text",
					"index": true,
					"sortable": true,
					"highlightable": true
				},
				"content_type": {
					"type": "text",
					"index": true,
					"sortable": true,
					"highlightable": true
				},
				"content_transfer_encoding": {
					"type": "text",
					"index": true,
					"sortable": true,
					"highlightable": true
				},
				"bcc": {
					"type": "text",
					"index": true,
					"sortable": true,
					"highlightable": true
				},
				"x_from": {
					"type": "text",
					"index": true,
					"sortable": true,
					"highlightable": true
				},
				"x_to": {
					"type": "text",
					"index": true,
					"sortable": true,
					"highlightable": true
				},
				"x_cc": {
					"type": "text",
					"index": true,
					"sortable": true,
					"highlightable": true
				},
				"x_bcc": {
					"type": "text",
					"index": true,
					"sortable": true,
					"highlightable": true
				},
				"x_folder": {
					"type": "text",
					"index": true,
					"sortable": true,
					"highlightable": true
				},
				"x_origin": {
					"type": "text",
					"index": true,
					"sortable": true,
					"highlightable": true
				},
				"x_file_name": {
					"type": "text",
					"index": true,
					"sortable": true,
					"highlightable": true
				},
				"content": {
					"type": "text",
					"index": true,
					"sortable": true,
					"highlightable": true
				}
			}
		}
	}`, constant.EMAIL_INDEX_NAME, config.NUM_CPUS)

	return indexData
}
