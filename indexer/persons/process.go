package persons

import (
	"fmt"
	"indexer/models"
	models_wp "indexer/models/workerpool"
	"regexp"
	"strings"
	"sync"
)

var (
	// Persons is the Map to add the people, it is used to avoid duplicates and then all the data will be added to a batch as structured Persons
	Persons map[string]string
	// mu is the mutex instance used to protect access to the Persons map and prevent race conditions.
	mu              sync.Mutex
	ExistingPersons int
	UniquePersons   int
	TotalPersons    int
)

func init() {
	// The Persons map is initialized to allow access by any Goroutine.
	Persons = make(map[string]string)
}

// ################ Retornar puntero
// StructurePersons is a function that takes an email and processes the data to extract the individuals who send and receive the email.
// It extracts the email addresses from the From, To, and Cc fields and their corresponding names from the X-From, X-To, and X-Cc fields.
// The data is cleaned and each person, along with their email, is added to the Persons map.
func StructurePersons(idWorker int, data models_wp.Result[models.Email]) (models.Person, error) {
	if data.Err != nil {
		return models.Person{}, data.Err
	}

	email := data.Value

	emailsOfFromField := splitAndClean(email.From)
	namesOfXFromField := splitAndClean(email.XFrom)

	emailsOfToField := splitAndClean(email.To)
	namesOfXToField := splitAndClean(email.XTo)

	emailsOfCcField := splitAndClean(email.Cc)
	namesOfXCcField := splitAndClean(email.XCc)

	cleanNamePerson(&namesOfXFromField)
	cleanNamePerson(&namesOfXToField)
	cleanNamePerson(&namesOfXCcField)

	processAndAppend(emailsOfFromField, namesOfXFromField)
	processAndAppend(emailsOfToField, namesOfXToField)
	processAndAppend(emailsOfCcField, namesOfXCcField)

	return models.Person{}, nil
}

// cleanNamePerson is a function that cleans up the names of individuals by removing any unwanted characters or tags.
func cleanNamePerson(names *[]string) {

	for i, name := range *names {
		re := regexp.MustCompile(`<[^>]*>`)
		matches := re.FindStringSubmatch(name)

		if len(matches) > 0 {
			nameClean := re.ReplaceAllString(name, "")
			(*names)[i] = strings.TrimSpace(nameClean)
		} else {
			(*names)[i] = strings.TrimSpace(name)
		}
	}
}

// ################ REFACRORIZAR CODIGO SI ES NECESARIO
// splitAndClean is a function that splits a string containing multiple email addresses and returns a slice of email addresses,
// each cleaned of surrounding spaces.
func splitAndClean(input string) []string {
	if input == "" {
		return []string{}
	}

	values := strings.Split(input, ", ")
	for i := range values {
		values[i] = strings.TrimSpace(values[i])
	}

	return values
}

// ################ REFACRORIZAR CODIGO SI ES NECESARIO
// processAndAppend is a function that takes two slices and adds them to the Persons map, ensuring via the index that each email is matched with the correct name.
//
// Parameters:
//   - emails: A slice of email addresses involved.
//   - names: A slice of names corresponding to the emails.
func processAndAppend(emails, names []string) {
	for i := range emails {
		if i < len(names) {
			if emails[i] == "" || names[i] == "" {
				continue
			}
			appendPerson(emails[i], names[i])
		}
	}
}

// appendPerson is a function that adds a record to the Persons map, using a mutex to prevent data loss or duplication.
func appendPerson(correo, nombre string) error {
	if correo == "" || nombre == "" {
		return nil
	}

	TotalPersons++

	mu.Lock()
	defer mu.Unlock()

	if _, exists := Persons[correo]; exists {
		ExistingPersons++
		return nil
	} else {
		fmt.Println(correo)
		UniquePersons++
	}

	Persons[correo] = nombre
	return nil
}
