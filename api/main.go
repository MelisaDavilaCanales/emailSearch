package main

import (
	"log"

	"backend/config"
	"backend/server"
)

func main() {
	if err := config.LoadEnvVars(); err != nil {
		log.Fatal("Error loading .env file: ", err)
	}

	server := server.NewServer()

	err := server.Run()
	if err != nil {
		log.Fatal("Error starting server:", err)
	}
}
