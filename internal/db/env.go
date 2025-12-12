package db

import (
	"fmt"
	"os"
)

func getDbUsername() string {
	return getEnvVariable("DB_USERNAME")
}

func getDbPassword() string {
	return getEnvVariable("DB_PASSWORD")
}

func getDbUrl() string {
	return getEnvVariable("DB_URL")
}

func getEnvVariable(envName string) string {
	value := os.Getenv(envName)
	if value == "" {
		fmt.Printf("Required ENV variables '%s' not set.\n", envName)
		os.Exit(1)
	}
	return value
}
