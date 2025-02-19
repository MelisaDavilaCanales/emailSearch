package models

import (
	"fmt"
)

// BatchManager manages a collection of Batch instances, allowing the efficient handling of batches across different types.
type BatchManager struct {
	BatchesCount int
	Batches      map[int]Batch
	BatchSize    int
}

// NewBatchManager creates a new BatchManager instance.
//
// Parameters:
//   - batchesCount: The number of batches to be managed by the BatchManager.
//   - batchSize: The size of each batch.
//
// Returns a pointer to the BatchManager instance.
func NewBatchManager(batchesCount int, batchSize int) *BatchManager {
	return &BatchManager{
		Batches:      make(map[int]Batch),
		BatchesCount: batchesCount,
		BatchSize:    batchSize,
	}
}

// InitBatches creates new Batch instances using the specified function.
//
// Parameters:
//   - newBatchFunc: The function must return a new Batch instance. It receives the batch id and the batch size as parameters.
func (bm *BatchManager) InitBatches(newBatchFunc func(id int) Batch) {
	for i := 0; i < bm.BatchesCount; i++ {
		newBatch := newBatchFunc(i + 1)
		bm.Batches[i+1] = newBatch
	}
}

// GetBatchById returns the Batch instance with the specified ID, if it exists.
// If the batch is not found, returns an error.
func (bm *BatchManager) GetBatchById(id int) (Batch, error) {
	if batch, ok := bm.Batches[id]; ok {
		return batch, nil
	}

	return nil, fmt.Errorf("batch with id %d not found", id)
}
