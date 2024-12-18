package models

// Persons represents a slice of persons involved in email communication.
type Persons []Person

// Person represents the data structure for an individual involved in email communication.
type Person struct {
	Name  string `json:"name"`
	Email string `json:"email"`
}
