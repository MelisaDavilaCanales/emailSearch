package emails

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strings"
	"time"

	"indexer/models"
)

// ProcessEmailsFiles reads and processes the email file, iterate the file line by line to parse and storing it in an Email structure.
func ProcessEmailsFiles(_ int, path string) (models.Email, error) {
	var (
		email          = models.Email{}
		isHeaderFinish = false
		emailContent   strings.Builder
		currentKey     string
	)

	file, err := os.Open(path)
	if err != nil {
		fmt.Printf("Error opening path: %v\n", err)
		return email, err
	}

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()

		if line == "" {
			isHeaderFinish = true
		}

		if !isHeaderFinish {
			ParseHeaderLine(line, &email, &currentKey)
		} else {
			emailContent.WriteString(line)
			emailContent.WriteString("\n")
		}
	}

	if err := file.Close(); err != nil {
		fmt.Printf("Error closing file: %v\n", err)
		return email, err
	}

	email.Content = emailContent.String()

	return email, nil
}

// ParseHeaderLine processes a header line and assigns its value to the corresponding field in the Email structure.
// If the line is not in key-value format, it's considered part of the content of the previous field and is appended to the current key's value.
func ParseHeaderLine(line string, email *models.Email, currentKey *string) {
	headerLine := strings.SplitN(line, ": ", 2)
	if len(headerLine) == 2 {
		*currentKey = headerLine[0]
		key := headerLine[0]
		value := headerLine[1]

		err := MapHeaderLine(key, value, email)
		if err != nil {
			fmt.Printf("??? Error mapping header line: %v\n", err)
		}
	}

	key := *currentKey
	value := line

	err := MapHeaderLine(key, value, email)
	if err != nil {
		fmt.Printf("??? Error mapping header line: %v\n", err)
	}
}

// MapHeaderLine maps the value of a header line to the appropriate field in the Email structure.
func MapHeaderLine(key string, value string, email *models.Email) error {
	switch key {
	case "Message-ID":
		email.MessageID = value
	case "Date":
		formatedDate, err := convertToDate(value)
		if err != nil {
			return err
		}

		email.Date = formatedDate
	case "From":
		email.From = value
	case "To":
		email.To += value
	case "Subject":
		email.Subject += value
	case "Cc":
		email.Cc += value
	case "X-cc":
		email.XCc += value
	case "Mime-Version":
		email.MimeVersion += value
	case "Content-Type":
		email.ContentType += value
	case "Content-Transfer-Encoding":
		email.ContentTransferEncoding += value
	case "Bcc":
		email.Bcc += value
	case "X-From":
		email.XFrom += value
	case "X-To":
		email.XTo += value
	case "X-bcc":
		email.XBcc += value
	case "X-Folder":
		email.XFolder += value
	case "X-Origin":
		email.XOrigin += value
	case "X-FileName":
		email.XFileName += value
	}

	return nil
}

// Convertir a formato ISO 8601 compatible con ZincSearch
func convertToDate(date string) (string, error) {
	cleanedDate := cleanDateString(date)

	parsedTime, err := time.Parse("Mon, _2 Jan 2006 15:04:05 -0700", cleanedDate)
	if err != nil {
		return "", err
	}

	time := parsedTime.Format("2006-01-02T15:04:05-07:00")

	return time, nil
}

// Función para eliminar la abreviatura de la zona horaria, con expresion regular para eliminanr lo que esté entre paréntesis
func cleanDateString(date string) string {
	reg := regexp.MustCompile(`\s\([^\)]+\)`)
	return reg.ReplaceAllString(date, "")
}
