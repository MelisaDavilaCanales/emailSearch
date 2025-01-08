package constant

const (

	// BUFFER_CAPACITY is the capacity of the buffer for the channels.
	BUFFER_CAPACITY = 50

	// PROCESS_EMAILS_WORKERS_COUNT Number of workers (goroutines) for the email-processing WorkerPool
	PROCESS_EMAILS_WORKERS_COUNT = 2

	// SEND_EMAILS_WORKERS_COUNT Number of workers (goroutines) for the email-sending WorkerPool
	SEND_EMAILS_WORKERS_COUNT = 2

	// STRUCTURE_PERSONS_WORKERS_COUNT Number of workers (goroutines) for the person-structuring WorkerPool
	STRUCTURE_PERSONS_WORKERS_COUNT = 2

	// BUILD_PERSON_BATCHES_WORKERS_COUNT Number of workers (goroutines) for the personBatches-building WorkerPool
	BUILD_PERSON_BATCHES_WORKERS_COUNT = 2
)
