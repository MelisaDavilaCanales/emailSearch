package emails

import (
	"encoding/csv"
	"log"
	"os"
	"time"

	"github.com/MelisaDavilaCanales/emailSearch/indexer/constant"
)

// logErrorToCSV logs the error and file path to a CSV file.
func LogErrorToCSV(filePath string, err error) {
	f, fileErr := os.OpenFile("error_log.csv", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if fileErr != nil {
		log.Fatalf("error opening CSV file: %v", fileErr)
	}
	defer f.Close() //nolint:errcheck

	writer := csv.NewWriter(f)
	defer writer.Flush()

	fileInfo, errStat := f.Stat()
	if errStat != nil {
		log.Fatalf("error checking file status: %v", err)
	}

	if fileInfo.Size() == 0 {
		err := writer.Write([]string{"Timestamp", "FilePath", "ErrorMessage"})
		if err != nil {
			log.Fatalf("error writing CSV header: %v", err)
		}
	}

	logTime := time.Now().Format(constant.DATE_FORMAT_TO_LOGS)
	record := []string{
		logTime,
		filePath,
		"err.Error()",
	}

	if err := writer.Write(record); err != nil {
		log.Fatalf("error writing to CSV file: %v", err)
	}
}
