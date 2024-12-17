package models

import "sync"

type WorkerPool[Input any, Output any] struct {
	taskExecutor TaskExecutor[Input, Output]
	taskQueueCh  chan Input
	resultChList []chan Result[Output]
	workerCount  int
	wg           *sync.WaitGroup
}

// NewWorkerPool creates a new instance of WorkerPool.
//
// Parameters:
//   - taskExecutor: A function that each worker will execute. This function processes the task with the provided data and returns a result.
//   - wg: A *sync.WaitGroup used to add and wait for all workers to complete their processing.
//   - workerCount: The number of workers that will be created in the pool.
//   - taskQueueCh: An input channel that the workers will use to receive tasks for processing.
//   - resultChList: A variadic parameter that can accept none, one, or multiple result channels, managed as a slice of channels.
//     The same result generated by the taskExecutor will be sent to all result channels provided.
//     If you do not need to send results through channels, simply do not pass any channels.
//
// Returns:
// A pointer to a WorkerPool instance configured with the provided parameters.
func NewWorkerPool[Input any, Output any](taskExecutor TaskExecutor[Input, Output], wg *sync.WaitGroup, workerCount int, taskQueueCh chan Input, resultChList ...chan Result[Output]) *WorkerPool[Input, Output] {
	return &WorkerPool[Input, Output]{
		taskExecutor: taskExecutor,
		wg:           wg,
		workerCount:  workerCount,
		taskQueueCh:  taskQueueCh,
		resultChList: resultChList,
	}
}

// Start initializes the workers by adding them to the WaitGroup and assigning each one an Id.
// It then starts each worker and waits for all of them to finish before close the result channels.
func (wp *WorkerPool[Input, Output]) Start() {
	for i := 0; i < wp.workerCount; i++ {
		wp.wg.Add(1)
		worker := NewWorker(i+1, wp.taskExecutor, wp.taskQueueCh, wp.wg, wp.resultChList)
		worker.Start()
	}

	go func() {
		wp.wg.Wait()

		if len(wp.resultChList) > 0 {
			for _, ch := range wp.resultChList {
				close(ch)
			}
		}
	}()
}
