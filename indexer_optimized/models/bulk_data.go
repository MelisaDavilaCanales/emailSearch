package models

import "github.com/MelisaDavilaCanales/emailSearch/indexer/constant"

// EmailBulkData represents the structure required to send a bulkV2 of emails to the ZincSearch API.
type EmailBulkData struct {
	Index  string                          `json:"index"`
	Emails [constant.EMAIL_BULK_SIZE]Email `json:"records"` //nolint:tagliatelle // The field name sent in the BulkV2 request is "records"
}

// NewEmailBulkData Creates a new EmailBulkData instance.
//
// Parameters:
//   - indexName: The name of the ZincSearch index for Emails.
//   - emails: The array of Email to be sent in the bulkV2 request.
//
// Returns a pointer to the new EmailBulkData instance.
func NewEmailBulkData(indexName string, emails [constant.EMAIL_BULK_SIZE]Email) *EmailBulkData {
	return &EmailBulkData{
		Index:  indexName,
		Emails: emails,
	}
}

// PersonBulkData represents the structure required to send a bulkV2 of persons to the ZincSearch API.
type PersonBulkData struct {
	Index   string  `json:"index"`
	Persons Persons `json:"records"` //nolint:tagliatelle // The field name sent in the BulkV2 request is "records"
}

// NewPersonBulkData Creates a new PersonBulkData instance.
//
// Parameters:
//   - indexName: The name of the ZincSearch index for Persons.
//   - persons: it's a slice of Person to be sent in the bulkV2 request.
//
// Returns a pointer to the new PersonBulkData instance.
func NewPersonBulkData(indexName string, persons Persons) *PersonBulkData {
	return &PersonBulkData{
		Index:   indexName,
		Persons: persons,
	}
}
