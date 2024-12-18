package config

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"

	"backend/constant"
)

var (
	// DB_USER is the username to connect to the database.
	DB_USER string

	// DB_PASSWORD is the password to connect to the database.
	DB_PASSWORD string

	// GET_EMAILS_API_URL represents the URL to get emails.
	GET_EMAILS_API_URL string

	// GET_EMAIL_API_URL represents the URL to get an email.
	GET_EMAIL_API_URL string

	// GET_PERSONS_API_URL represents the URL to get persons.
	GET_PERSONS_API_URL string

	// API_PORT is the port where the API will run.
	API_PORT string
)

func setGlobalEnvVars() {
	DB_USER = os.Getenv("DB_USER")
	DB_PASSWORD = os.Getenv("DB_PASSWORD")

	GET_EMAILS_API_URL = os.Getenv("ZINC_SEARCH_API_URL") + constant.EMAIL_INDEX_NAME + "/_search"
	GET_EMAIL_API_URL = os.Getenv("ZINC_SEARCH_API_URL") + constant.EMAIL_INDEX_NAME + "/_doc/"

	GET_PERSONS_API_URL = os.Getenv("ZINC_SEARCH_API_URL") + constant.PERSON_INDEX_NAME + "/_search"

	API_PORT = os.Getenv("API_PORT")
}

func LoadEnvVars() error {
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

	if os.Getenv("API_PORT") == "" {
		return fmt.Errorf("API_PORT environment variable not set")
	}

	setGlobalEnvVars()

	return nil
}
