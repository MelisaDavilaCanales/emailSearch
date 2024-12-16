package models

import "sync"

type Worker[Input any, Output any] struct {
	Id           int
	taskExecutor TaskExecutor[Input, Output]
	taskQueueCh  chan Input
	wg           *sync.WaitGroup
	resultChList []chan Result[Output]
}

// NewWorker creates a new instance of Worker.
//
// Parameters:
//
//	- id: The identifier for the worker.
//	- taskExecutor: The function that will execute the tasks.
//	- taskQueueCh: The channel that the worker will listen to and take tasks.
//	- resultChList: A list of channels through which the result returned by the taskExecutor will be sent.
//	- wg: A *sync.WaitGroup that the worker belongs to, used to decrement the counter when the worker finish.
//
// Returns:
// A pointer to a Worker instance configured with the provided parameters.
func NewWorker[Input any, Output any](
	id int,
	taskExecutor TaskExecutor[Input, Output],
	tasksCh chan Input,
	wg *sync.WaitGroup,
	resultChList []chan Result[Output],
) *Worker[Input, Output] {
	return &Worker[Input, Output]{
		Id:           id,
		taskExecutor: taskExecutor,
		taskQueueCh:  tasksCh,
		resultChList: resultChList,
		wg:           wg,
	}
}

// Start begins the worker's task processing.
// The worker listens to the task channel, it is executed by the taskExecutor, and the result is sent by resultChanels
// The worker continues processing until the task channel is closed.
// Note: The worker does not handle errors internally; errors are passed through the result.
func (w *Worker[Input, Output]) Start() {
	go func() {
		defer w.wg.Done()
		for task := range w.taskQueueCh {
			result, err := w.taskExecutor(w.Id, task)
			if len(w.resultChList) > 0 {
				for _, resultCh := range w.resultChList {
					resultCh <- Result[Output]{Value: result, Err: err}
				}
			}
		}
	}()
}
