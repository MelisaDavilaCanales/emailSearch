package emails

import (
	"encoding/csv"
	"fmt"
	"os"
	"time"

	"github.com/MelisaDavilaCanales/emailSearch/indexer/constant"
)

// logErrorToCSV logs the error and file path to a CSV file.
func LogErrorToCSV(filePath string, err error) {
	f, fileErr := os.OpenFile("error_log.csv", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if fileErr != nil {
		fmt.Printf("Error opening error log CSV file: %v\n", fileErr)
		return
	}
	defer f.Close() //nolint:errcheck

	writer := csv.NewWriter(f)
	defer writer.Flush()

	fileInfo, errStat := f.Stat()
	if errStat != nil {
		fmt.Printf("Error checking status of error log file: %v\n", err)
		return
	}

	if fileInfo.Size() == 0 {
		err := writer.Write([]string{"timestamp", "filePath", "errorMessage"})
		if err != nil {
			fmt.Printf("Error writing header to error log CSV file: %v\n", err)
			return
		}
	}

	logTime := time.Now().Format(constant.DATE_FORMAT_TO_LOGS)
	record := []string{
		logTime,
		filePath,
		err.Error(),
	}

	if err := writer.Write(record); err != nil {
		fmt.Printf("Error writing error record to CSV log file: %v\n", err)
	}
}
