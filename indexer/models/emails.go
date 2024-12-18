package models

import (
	"fmt"
	"time"

	"github.com/MelisaDavilaCanales/emailSearch/indexer/constant"
)

// Email represents the structure of an email.
type Email struct {
	MessageID               string    `json:"message_id"`
	Date                    time.Time `json:"date"`
	From                    string    `json:"from"`
	To                      string    `json:"to"`
	Subject                 string    `json:"subject"`
	Cc                      string    `json:"cc"`
	MimeVersion             string    `json:"mime_version"`
	ContentType             string    `json:"content_type"`
	ContentTransferEncoding string    `json:"content_transfer_encoding"`
	Bcc                     string    `json:"bcc"`
	XFrom                   string    `json:"x_from"`
	XTo                     string    `json:"x_to"`
	XCc                     string    `json:"x_cc"`
	XBcc                    string    `json:"x_bcc"`
	XFolder                 string    `json:"x_folder"`
	XOrigin                 string    `json:"x_origin"`
	XFileName               string    `json:"x_file_name"`
	Content                 string    `json:"content"`
}

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
	Emails    [constant.EMAIL_BATCH_SIZE]Email
	NextIndex int
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
		Emails:    [constant.EMAIL_BATCH_SIZE]Email{},
		NextIndex: 0,
	}
}

// GetBatchID returns the Id of the EmailBatch
func (eb *EmailBatch) GetBatchID() int {
	return eb.Id
}

// Reset resets all Emails in the EmailBatch, resetting it to an empty state.
func (eb *EmailBatch) Reset() {
	eb.Emails = [constant.EMAIL_BATCH_SIZE]Email{}
	eb.NextIndex = 0
}

// IsFull returns true if EmailBatch is at full capacity, false otherwise.
func (eb *EmailBatch) IsFull() bool {
	return eb.NextIndex >= eb.Size
}

// AddItem adds an Email to the EmailBatch.
func (eb *EmailBatch) AddItem(item interface{}) error {
	email, ok := item.(Email)
	if !ok {
		return fmt.Errorf("item is not of type Email")
	}

	if eb.IsFull() {
		return fmt.Errorf("batch is full")
	}

	eb.Emails[eb.NextIndex] = email
	eb.NextIndex++

	return nil
}
