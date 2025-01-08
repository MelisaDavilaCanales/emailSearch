package config

import (
	"errors"
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

	// TRY_CONNECTION_API_URL represents the URL to try to connect to the database.
	TRY_CONNECTION_API_URL string

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
	TRY_CONNECTION_API_URL = os.Getenv("ZINC_SEARCH_API_URL") + "index/"
	EMAIL_DIR_SUBPATH = ""

}

func SetEnvVars() error {

	if os.Getenv("APP_MODE") != "release" {
		err := godotenv.Load()
		if err != nil {
			return err
		}
	}

	if os.Getenv("DB_USER") == "" {
		return errors.New("DB_USER environment variable not set")
	}

	if os.Getenv("DB_PASSWORD") == "" {
		return errors.New("DB_PASSWORD environment variable not set")
	}

	if os.Getenv("ZINC_SEARCH_API_URL") == "" {
		return errors.New("ZINC_SEARCH_API_URL environment variable not set")
	}

	setGlobalEnvVars()

	return nil
}
