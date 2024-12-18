package config

import (
	"backend/constant"
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

	// EMAIL_DIR_SUBPATH is the subpath to the email directory.
	EMAIL_DIR_SUBPATH = "/maildir"

	// GET_EMAILS_API_URL represents the URL to get emails.
	GET_EMAILS_API_URL string

	// GET_EMAIL_API_URL represents the URL to get an email.
	GET_EMAIL_API_URL string

	// GET_PERSONS_API_URL represents the URL to get persons.
	GET_PERSONS_API_URL string
)

func setGlobalEnvVars() {
	NUM_CPUS = runtime.NumCPU()
	DB_USER = os.Getenv("DB_USER")
	DB_PASSWORD = os.Getenv("DB_PASSWORD")

	GET_EMAILS_API_URL = os.Getenv("ZINC_SEARCH_API_URL") + constant.EMAIL_INDEX_NAME + "/_search"
	GET_EMAIL_API_URL = os.Getenv("ZINC_SEARCH_API_URL") + constant.EMAIL_INDEX_NAME + "/_doc/"

	GET_PERSONS_API_URL = os.Getenv("ZINC_SEARCH_API_URL") + constant.PERSON_INDEX_NAME + "/_search"
}

func LoadEnvVars() error {
	err := godotenv.Load()
	if err != nil {
		return err
	}

	setGlobalEnvVars()

	return nil
}
