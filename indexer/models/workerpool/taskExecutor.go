package models

// TaskExecutor is a function type that executes a task. Is used by workers to perform different tasks.
//
// Parameters:
//	- workerIndex: the index of the worker executing the task
//	- input: the data required to perform the task
//
// Returns:
//	- result: the result of the task when it completes successfully
//	- err: an error that occurs during execution, or nil if no error occurs
type TaskExecutor[Input any, Output any] func(workerIndex int, input Input) (result Output, err error)
