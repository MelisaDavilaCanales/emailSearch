package constant

//batchesCount int, batchSize int

const (
	// EMAIL_BATCHES_COUNT represents the number of EmailBatches managed by BatchManager. This corresponds to SEND_EMAILS_WORKERS_COUNT, since each worker sending emails uses its own batch.
	EMAIL_BATCHES_COUNT = SEND_EMAILS_WORKERS_COUNT

	// EMAIL_BATCH_SIZE represents the size/capacity of each EmailBatch. That is, the number of emails each batch will contain.
	EMAIL_BATCH_SIZE = 1000
)
