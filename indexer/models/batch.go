package models

// Batch represents an interface for managing a collection of items in a batch.
//
// Methods:
//   - Reset(): Resets the batch, clearing all items and resetting its state.
//   - GetBatchID(): Returns the unique identifier of the batch.
//   - AddItem(): Adds an item to the batch and returns an error if the item cannot be added.
type Batch interface {
	Reset()
	GetBatchID() int
	AddItem(interface{}) error
}
