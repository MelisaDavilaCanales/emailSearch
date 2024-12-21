package storage

import (
	"fmt"

	"github.com/MelisaDavilaCanales/emailSearch/indexer/config"
	"github.com/MelisaDavilaCanales/emailSearch/indexer/constant"
)

// CreatePersonIndex creates an index for storing persons.
func CreatePersonIndex() error {
	indexDataStr := buildPersonIndex()

	mssg, err := CreateIndex(constant.PERSON_INDEX_NAME, indexDataStr)
	if err != nil {
		return err
	}

	fmt.Println("Person index: ", mssg)

	return nil
}

func buildPersonIndex() string {
	indexData := fmt.Sprintf(`{
		"name": "%v",
		"storage_type": "disk",
		"shard_num": %d,
		"mappings": {
			"properties": {
				"email": {
					"type": "keyword",
					"index": true,
					"sortable": true,
					"highlightable": true
				},
				"name": {
					"type": "keyword",
					"index": true,
					"sortable": true,
					"highlightable": true
				}
			}
		}
	}`, constant.PERSON_INDEX_NAME, config.NUM_CPUS)

	return indexData
}
