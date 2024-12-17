package emails

import (
	"fmt"

	"indexer/constant"
	"indexer/models"
	models_wp "indexer/models/workerpool"
	"indexer/storage"
)

var EmailBatchManager *models.BatchManager

// init initializes the EmailBatchManager to manage email batches. It ensures that workers can access their respective batches at any time.
// The manager is set up with the number of batches and batch size defined in the constants.
func init() {
	EmailBatchManager = models.NewBatchManager(constant.EMAIL_BATCHES_COUNT, constant.EMAIL_BATCH_SIZE)
	EmailBatchManager.InitBatches(func(id int) models.Batch {
		return models.NewEmailBatch(id)
	})
}

// ProcessAndSendEmails retrieves the appropriate batch based on the worker ID, adds the email to the batch, and if the batch is full,
// creates a bulk of emails whith the elements of that batch and sends it to the API.
func ProcessAndSendEmails(idWorker int, data models_wp.Result[models.Email]) (bool, error) {
	if data.Err != nil {
		return false, data.Err
	}

	idBatch := idWorker
	email := data.Value

	batch, err := buildEmailBatch(idBatch, email)
	if err != nil {
		return false, err
	}

	if batch.IsFull() {
		bulk := createBulk(batch)
		if err := storage.SendBulk(bulk); err != nil {
			return false, fmt.Errorf("failed to send batch: %w", err)
		}

		batch.Reset()
	}

	return true, nil
}

// buildEmailBatch retrieves the batch based on the ID, adds the email to the batch, and returns the batch.
func buildEmailBatch(idBatch int, email models.Email) (*models.EmailBatch, error) {
	myBatch, err := EmailBatchManager.GetBatchById(idBatch)
	if err != nil {
		return &models.EmailBatch{}, fmt.Errorf("failed to get batch: %w", err)
	}

	if err = myBatch.AddItem(email); err != nil {
		return &models.EmailBatch{}, fmt.Errorf("failed to add email to batch: %w", err)
	}

	//################## Que retornar nil o &models.EmailBatch{}
	batch, ok := myBatch.(*models.EmailBatch)
	if !ok {
		return nil, fmt.Errorf("failed to cast batch")
	}

	return batch, nil
}

func createBulk(batch *models.EmailBatch) *models.EmailBulkData {
	bulk := models.NewEmailBulkData(constant.EMAIL_INDEX_NAME, batch.Emails)
	return bulk
}
