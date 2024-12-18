package emails

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"strings"
	"time"

	"github.com/MelisaDavilaCanales/emailSearch/indexer/constant"
	"github.com/MelisaDavilaCanales/emailSearch/indexer/models"
)

type HeaderKey string

const (
	MESSAGE_ID                HeaderKey = "Message-ID"
	DATE                      HeaderKey = "Date"
	FROM                      HeaderKey = "From"
	TO                        HeaderKey = "To"
	SUBJECT                   HeaderKey = "Subject"
	CC                        HeaderKey = "Cc"
	X_CC                      HeaderKey = "X-cc"
	MIME_VERSION              HeaderKey = "Mime-Version"
	CONTENT_TYPE              HeaderKey = "Content-Type"
	CONTENT_TRANSFER_ENCODING HeaderKey = "Content-Transfer-Encoding"
	BCC                       HeaderKey = "Bcc"
	X_FROM                    HeaderKey = "X-From"
	X_TO                      HeaderKey = "X-To"
	X_BCC                     HeaderKey = "X-bcc"
	X_FOLDER                  HeaderKey = "X-Folder"
	X_ORIGIN                  HeaderKey = "X-Origin"
	X_FILE_NAME               HeaderKey = "X-FileName"
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

		err := MapHeaderLine(HeaderKey(key), value, email)
		if err != nil {
			fmt.Printf("Error mapping header line: %v\n", err)
		}

		return
	}

	key := *currentKey
	value := line

	err := MapHeaderLine(HeaderKey(key), value, email)
	if err != nil {
		fmt.Printf("Error mapping header line: %v\n", err)
	}
}

// MapHeaderLine maps the value of a header line to the appropriate field in the Email structure.
func MapHeaderLine(key HeaderKey, value string, email *models.Email) error {
	switch key {
	case MESSAGE_ID:
		email.MessageID = value
	case DATE:
		formatedDate, err := ConvertDateFormat(value)
		if err != nil {
			return err
		}
		email.Date = formatedDate
	case FROM:
		email.From = value
	case TO:
		email.To += value
	case SUBJECT:
		email.Subject += value
	case CC:
		email.Cc += value
	case X_CC:
		email.XCc += value
	case MIME_VERSION:
		email.MimeVersion += value
	case CONTENT_TYPE:
		email.ContentType += value
	case CONTENT_TRANSFER_ENCODING:
		email.ContentTransferEncoding += value
	case BCC:
		email.Bcc += value
	case X_FROM:
		email.XFrom += value
	case X_TO:
		email.XTo += value
	case X_BCC:
		email.XBcc += value
	case X_FOLDER:
		email.XFolder += value
	case X_ORIGIN:
		email.XOrigin += value
	case X_FILE_NAME:
		email.XFileName += value
	default:
		return errors.New("unknown header key")
	}

	return nil
}

func ConvertDateFormat(date string) (time.Time, error) {
	time, err := time.Parse(constant.DATE_FORMAT, date)

	return time, err
}
