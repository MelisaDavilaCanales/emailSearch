package persons

import (
	"fmt"
	"indexer/constant"
	"indexer/models"
	"indexer/storage"
)

// BuildAndSendPersonBulk constructs a bulk of persons using the "Persons" map defined and populated in the process.go file.
// It adds all Person entries from the map to the bulk and then sends the bulk to the ZincSearch API.
func BuildAndSendPersonBulk() {

	personBulk := models.NewPersonBulkData(constant.PERSON_INDEX_NAME, models.Persons{})

	for email, name := range Persons {
		personBulk.Persons = append(personBulk.Persons, models.Person{
			Name:  name,
			Email: email,
		})
	}

	if err := storage.SendBulk(personBulk); err != nil {
		fmt.Println("failed to send bulk: %w", err)
	}
}
