package constant

const (

	// BUFFER_CAPACITY is the capacity of the buffer for the channels.
	BUFFER_CAPACITY = 1000

	// PROCESS_EMAILS_WORKERS_COUNT Number of workers (goroutines) for the email-processing WorkerPool
	PROCESS_EMAILS_WORKERS_COUNT = 10

	// SEND_EMAILS_WORKERS_COUNT Number of workers (goroutines) for the email-sending WorkerPool
	SEND_EMAILS_WORKERS_COUNT = 10

	// STRUCTURE_PERSONS_WORKERS_COUNT Number of workers (goroutines) for the person-structuring WorkerPool
	STRUCTURE_PERSONS_WORKERS_COUNT = 10
)
