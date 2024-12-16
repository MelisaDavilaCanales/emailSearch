package models

// Batch represents an interface for managing a collection of items in a batch.
//
// Methods:
//   - Reset(): Resets the batch, clearing all items and resetting its state.
//   - GetBatchID(): Returns the unique identifier of the batch.
//   - AddItem(item interface{}) error: Adds an item to the batch. Returns an error if the item cannot be added.
type Batch interface {
	Reset()
	GetBatchID() int
	AddItem(item interface{}) error
}
