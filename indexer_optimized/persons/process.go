package persons

import (
	"regexp"
	"strings"
	"sync"

	"github.com/MelisaDavilaCanales/emailSearch/indexer/constant"
	"github.com/MelisaDavilaCanales/emailSearch/indexer/models"
	models_wp "github.com/MelisaDavilaCanales/emailSearch/indexer/models/workerpool"
)

var (
	// Persons is the Map to add the people, it is used to avoid duplicates and then all the data will be added to a batch as structured Persons
	Persons map[string]string
	mu      sync.Mutex
)

func init() {
	// The Persons map is initialized to allow access by any Goroutine.
	Persons = make(map[string]string)
}

// StructurePersons is a function that takes an email and processes the data to extract the individuals who send and receive the email.
// It extracts the email addresses from the From, To, and Cc fields and their corresponding names from the X-From, X-To, and X-Cc fields.
// The data is cleaned and each person, along with their email, is added to the Persons map.
func StructurePersons(_ int, data models_wp.Result[*models.EmailData]) (models.Person, error) {
	if data.Err != nil {
		return models.Person{}, data.Err
	}

	email := data.Value.EmailStruct

	emailsOfFromField := convertToArray(email.From)
	namesOfXFromField := convertToArray(email.XFrom)

	emailsOfToField := convertToArray(email.To)
	namesOfXToField := convertToArray(email.XTo)

	emailsOfCcField := convertToArray(email.Cc)
	namesOfXCcField := convertToArray(email.XCc)

	cleanPersonName(&namesOfXFromField)
	cleanPersonName(&namesOfXToField)
	cleanPersonName(&namesOfXCcField)

	cleanPersonEmail(&emailsOfFromField)
	cleanPersonEmail(&emailsOfToField)
	cleanPersonEmail(&emailsOfCcField)

	processAndAppend(emailsOfFromField, namesOfXFromField)
	processAndAppend(emailsOfToField, namesOfXToField)
	processAndAppend(emailsOfCcField, namesOfXCcField)

	return models.Person{}, nil
}

// cleanPersonName is a function that cleans up the names of individuals by removing any unwanted characters or tags e.g. < >
func cleanPersonName(names *[]string) {
	for i, name := range *names {
		re := regexp.MustCompile(constant.TAG_CONTENT_REGEX)
		matches := re.FindStringSubmatch(name)

		if len(matches) > 0 {
			nameClean := re.ReplaceAllString(name, "")
			(*names)[i] = strings.TrimSpace(nameClean)
		}

		(*names)[i] = strings.TrimSpace(name)
	}
}

// cleanPersonEmail remove unwanted Chars e.g. e-mail, <email.>, <., <'.' and >
func cleanPersonEmail(emails *[]string) {
	regexp := regexp.MustCompile(`(?i)e-mail|<email.>|<\.\s*|<'?'\s*|\s*>`)

	for i, email := range *emails {
		emailClean := regexp.ReplaceAllString(email, "")

		emailClean = strings.ReplaceAll(emailClean, "<.", "")
		emailClean = strings.ReplaceAll(emailClean, "<'.", "")
		emailClean = strings.ReplaceAll(emailClean, ".'", "")

		(*emails)[i] = strings.TrimSpace(emailClean)
	}
}

// convertToArray is a function that splits a string containing multiple email addresses and returns a slice of email addresses,
func convertToArray(input string) []string {
	if input == "" {
		return []string{}
	}

	values := strings.Split(input, ", ")
	return values
}

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
func appendPerson(correo, nombre string) {
	if correo == "" || nombre == "" {
		return
	}

	mu.Lock()
	defer mu.Unlock()

	if _, exists := Persons[correo]; exists {
		return
	}

	Persons[correo] = nombre
}
