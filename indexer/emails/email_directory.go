package emails

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/MelisaDavilaCanales/emailSearch/indexer/config"
)

// ProcessEmailDirectory retrieves the directory path from the command-line arguments, processes the specified directory,
// and scans it for email files, and sends their paths to the provided channel.
func ProcessEmailDirectory(emailPathCh chan string) error {
	dir, err := getDirectory()
	if err != nil {
		return err
	}

	path := dir + config.EMAIL_DIR_SUBPATH
	if err = scanDirectory(path, emailPathCh); err != nil {
		return err
	}

	return nil
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
			return "", fmt.Errorf("directory %s does not exist", dirPath)
		}

		return "", fmt.Errorf("checking directory: %v", err)
	}

	return dirPath, nil
}

// scanDirectory recursively scans the specified directory for files and sends their paths to the provided channel.
func scanDirectory(path string, emailPathCh chan string) error {
	files, err := os.ReadDir(path)
	if err != nil {
		return fmt.Errorf("reading directory: %v", err)
	}

	for _, file := range files {
		filePath := filepath.Join(path, file.Name())
		if file.IsDir() {
			err := scanDirectory(filePath, emailPathCh)
			if err != nil {
				return err
			}
		}

		emailPathCh <- filePath
	}

	return nil
}
