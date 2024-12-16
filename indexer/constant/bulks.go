package constant

const (
	// EMAIL_BULK_SIZE defines the maximum size/capacity of a PersonBulk, that is the number of persons an EmailBulk contains.
	// It corresponds to EMAIL_BATCH_SIZE because for each emailBatch a PersonBulk is sent to the ZincSearch API with the same number of elements (Emails)
	EMAIL_BULK_SIZE = EMAIL_BATCH_SIZE
)
