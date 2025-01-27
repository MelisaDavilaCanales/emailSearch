package emails

import (
	"errors"
	"fmt"

	"github.com/MelisaDavilaCanales/emailSearch/indexer/constant"
	"github.com/MelisaDavilaCanales/emailSearch/indexer/models"
	models_wp "github.com/MelisaDavilaCanales/emailSearch/indexer/models/workerpool"
	"github.com/MelisaDavilaCanales/emailSearch/indexer/storage"
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
func ProcessAndSendEmails(idWorker int, data models_wp.Result[*models.EmailData]) (bool, error) {
	if data.Err != nil {
		return false, data.Err
	}

	idBatch := idWorker
	emailData := data.Value
	path := data.Value.EmailPath

	batch, err := addToEmailBatch(idBatch, emailData)
	if err != nil {
		LogErrorToCSV(path, err)
		return false, err
	}

	if batch.IsFull() {
		bulk := createBulk(batch)
		if err := storage.SendBulk(bulk); err != nil {
			for _, emailData := range batch.EmailData {
				LogErrorToCSV(emailData.EmailPath, fmt.Errorf("send data bulk: %w", err))
			}

			return false, fmt.Errorf("send data bulk: %w", err)
		}

		batch.Reset()
	}

	return true, nil
}

// addToEmailBatch retrieves the batch based on the ID, adds the email to the batch, and returns the batch.
func addToEmailBatch(idBatch int, emailData *models.EmailData) (*models.EmailBatch, error) {
	myBatch, err := EmailBatchManager.GetBatchById(idBatch)
	if err != nil {
		return &models.EmailBatch{}, fmt.Errorf("get batch: %w", err)
	}

	err = myBatch.AddItem(*emailData)
	if err != nil {
		return nil, fmt.Errorf("add item: %w", err)
	}

	batch, ok := myBatch.(*models.EmailBatch)
	if !ok {
		return nil, errors.New("cast emailBatch")
	}

	return batch, nil
}

func createBulk(batch *models.EmailBatch) *models.EmailBulkData {
	var dataToBulk [constant.EMAIL_BATCH_SIZE]models.Email

	for i, emailData := range batch.EmailData {
		if emailData.EmailStruct == nil {
			LogErrorToCSV(batch.EmailData[i].EmailPath, errors.New("email struct is nil"))

			continue
		}

		dataToBulk[i] = *emailData.EmailStruct
	}

	bulk := models.NewEmailBulkData(constant.EMAIL_INDEX_NAME, dataToBulk)

	return bulk
}
