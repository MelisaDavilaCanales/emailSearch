package config

import (
	"fmt"

	"github.com/joho/godotenv"
)

func GetEnvVars() error {
	err := godotenv.Load()
	if err != nil {
		return err
	}
	fmt.Println("Environment variables loaded successfully")
	return nil
}
