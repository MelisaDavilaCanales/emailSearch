package emails

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
)

// ProcessEmailDirectory retrieves the directory path from the command-line arguments, processes the specified directory,
// and scans it for email files, and sends their paths to the provided channel.
func ProcessEmailDirectory(emailPathCh chan string) {

	dir, err := getDirectory()
	if err != nil {
		log.Fatal(err, " please provide a valid directory, example: go run main.go <directory>")
	}

	path := dir + "/maildir"
	if err = scanDirectory(path, emailPathCh); err != nil {
		log.Fatal(err)
	}

}

// getDirectory retrieves the directory path from the command-line arguments and verifies it's existence.
func getDirectory() (string, error) {

	if len(os.Args) < 2 {
		return "", fmt.Errorf("no directory provided ")
	}

	dirPath := os.Args[1]
	_, err := os.Stat(dirPath)
	if err != nil {
		if os.IsNotExist(err) {
			return "", fmt.Errorf("directory does not exist: %s", dirPath)
		}
		return "", fmt.Errorf("error checking directory: %v", err)
	}

	return dirPath, nil
}

// scanDirectory recursively scans the specified directory for files and sends their paths to the provided channel.
func scanDirectory(path string, emailPathCh chan string) error {

	files, err := os.ReadDir(path)
	if err != nil {
		return fmt.Errorf("error reading directory: %v", err)
	}

	for _, file := range files {
		filePath := filepath.Join(path, file.Name())
		if file.IsDir() {
			scanDirectory(filePath, emailPathCh)
		} else {
			emailPathCh <- filePath
		}
	}

	return nil
}
