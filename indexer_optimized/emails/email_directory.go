package emails

import (
	"fmt"
	"os"
	"path/filepath"
)

var CantGoRutine int // Variable para contar la cantidad de goroutines lanzadas.

// GetDirectory retrieves the directory path from the command-line arguments and verifies its existence.
func GetDirectory() (string, error) {
	if len(os.Args) < 2 {
		return "", fmt.Errorf("no directory provided")
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

// ProcessSubDirectory processes a single subdirectory (recursively scans its contents).
func ProcessSubDirectory(pathDirectory string, emailCh chan string) {
	files, err := os.ReadDir(pathDirectory)
	if err != nil {
		fmt.Printf("error reading directory %s: %v", pathDirectory, err)
	}

	for _, file := range files {
		filePath := filepath.Join(pathDirectory, file.Name())
		if file.IsDir() {
			ProcessSubDirectory(filePath, emailCh)
		} else {
			emailCh <- filePath
		}
	}
}

// ProcessEmailDirectory retrieves the directory path from the command-line arguments, processes the specified directory,
// and scans it for email files, and sends their paths to the provided channel.
func ProcessEmailDirectory(emailPathCh chan string) error {
	dir, err := GetDirectory()
	if err != nil {
		return err
	}

	path := dir
	if err = scanDirectory(path, emailPathCh); err != nil {
		return err
	}

	return nil
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
