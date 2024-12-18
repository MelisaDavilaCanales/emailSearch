package config

import (
	"fmt"
	"os"
	"runtime"

	"github.com/joho/godotenv"
)

var (
	// NUM_CPUS is the number of CPUs available on the machine.
	NUM_CPUS int

	// DB_USER is the username to connect to the database.
	DB_USER string

	// DB_PASSWORD is the password to connect to the database.
	DB_PASSWORD string

	// SEND_BULK_API_URL represents the URL to send a bulk of data.
	SEND_BULK_API_URL string

	// CREATE_INDEX_API_URL represents the URL to create an index.
	CREATE_INDEX_API_URL string

	// CHECK_INDEX_EXISTS_API_URL represents the URL to check if an index exists.
	CHECK_INDEX_EXISTS_API_URL string

	// EMAIL_DIR_SUBPATH is the subpath to the email directory.
	EMAIL_DIR_SUBPATH string
)

func setGlobalEnvVars() {
	NUM_CPUS = runtime.NumCPU()

	DB_USER = os.Getenv("DB_USER")
	DB_PASSWORD = os.Getenv("DB_PASSWORD")

	SEND_BULK_API_URL = os.Getenv("ZINC_SEARCH_API_URL") + "_bulkv2"
	CREATE_INDEX_API_URL = os.Getenv("ZINC_SEARCH_API_URL") + "index"
	CHECK_INDEX_EXISTS_API_URL = os.Getenv("ZINC_SEARCH_API_URL") + "index/"

	EMAIL_DIR_SUBPATH = "/maildir"
}

func SetEnvVars() error {
	err := godotenv.Load()
	if err != nil {
		return err
	}

	if os.Getenv("DB_USER") == "" {
		return fmt.Errorf("DB_USER environment variable not set")
	}

	if os.Getenv("DB_PASSWORD") == "" {
		return fmt.Errorf("DB_PASSWORD environment variable not set")
	}

	if os.Getenv("ZINC_SEARCH_API_URL") == "" {
		return fmt.Errorf("ZINC_SEARCH_API_URL environment variable not set")
	}

	setGlobalEnvVars()

	return nil
}
