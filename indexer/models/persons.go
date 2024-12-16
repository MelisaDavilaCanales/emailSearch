package models

type Persons []Person

type Person struct {
	Name  string `json:"name"`
	Email string `json:"email"`
}
