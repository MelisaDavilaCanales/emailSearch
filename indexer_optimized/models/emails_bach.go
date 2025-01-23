package models

import (
	"errors"

	"github.com/MelisaDavilaCanales/emailSearch/indexer/constant"
)

// EmailBatch implements the Batch interface.
//
// Fields:
//   - Id: The Id of the EmailBatch
//   - Size: The size of the EmailBatch (The maximum number of Emails that can be stored in the EmailBatch).
//   - Emails: It is an array of Emails of size EMAIL_BATCH_SIZE
//   - NextIndex: Indicates the next available index to add an element
type EmailBatch struct {
	Id        int
	Size      int
	EmailData [constant.EMAIL_BATCH_SIZE]EmailData
	NextIndex int
}

// EmailData represents the structure of an email and its path.
type EmailData struct {
	EmailPath   string
	EmailStruct *Email
}

// NewEmailBatch initializes and returns a new instance of EmailBatch.
//
// Parameters:
//   - idBatch: The Id for the EmailBatch.
//
// Returns a pointer to the new EmailBatch instance.
func NewEmailBatch(idBatch int) *EmailBatch {
	return &EmailBatch{
		Id:        idBatch,
		Size:      constant.EMAIL_BATCH_SIZE,
		EmailData: [constant.EMAIL_BATCH_SIZE]EmailData{},
		NextIndex: 0,
	}
}

// GetBatchID returns the Id of the EmailBatch
func (eb *EmailBatch) GetBatchID() int {
	return eb.Id
}

// Reset resets all Emails in the EmailBatch, resetting it to an empty state.
func (eb *EmailBatch) Reset() {
	eb.EmailData = [constant.EMAIL_BATCH_SIZE]EmailData{}
	eb.NextIndex = 0
}

// IsFull returns true if EmailBatch is at full capacity, false otherwise.
func (eb *EmailBatch) IsFull() bool {
	return eb.NextIndex >= eb.Size
}

// AddItem adds an Email to the EmailBatch.
func (eb *EmailBatch) AddItem(item interface{}) error {
	emailData, ok := item.(EmailData)
	if !ok {
		return errors.New("item is not of type emailData")
	}

	if eb.IsFull() {
		return errors.New("batch is full")
	}

	eb.EmailData[eb.NextIndex] = emailData
	eb.NextIndex++

	return nil
}
