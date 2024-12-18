package storage

import (
	"fmt"

	"github.com/MelisaDavilaCanales/emailSearch/indexer/config"
	"github.com/MelisaDavilaCanales/emailSearch/indexer/constant"
)

func CreatePersonIndex() error {
	indexDataStr := buildPersonIndex()
	return CreateIndex(constant.PERSON_INDEX_NAME, indexDataStr)
}

func buildPersonIndex() string {
	indexData := fmt.Sprintf(`{
		"name": "%v",
		"storage_type": "disk",
		"shard_num": %d,
		"mappings": {
			"properties": {
				"email": {
					"type": "text",
					"index": true,
					"sortable": true,
					"highlightable": true
				},
				"name": {
					"type": "text",
					"index": true,
					"sortable": true,
					"highlightable": true
				}
			}
		}
	}`, constant.PERSON_INDEX_NAME, config.NUM_CPUS)

	return indexData
}
