package config

import (
	"runtime"

	"github.com/joho/godotenv"
)

var (
	// NUM_CPUS is the number of CPUs available on the machine.
	NUM_CPUS int
)

func init() {
	NUM_CPUS = runtime.NumCPU()
}

func LoadEnvVars() error {
	err := godotenv.Load()
	if err != nil {
		return err
	}
	return nil
}
