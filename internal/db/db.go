package db

import (
	"fmt"
)

func InitializeDB() error {
	username := getDbUsername()
	password := getDbPassword()
	dbUrl := getDbUrl()
	fmt.Println("Initializing database for user:", username)
	fmt.Println("Using password:", password)
	fmt.Println("Connecting to DB at URL:", dbUrl)
	return nil
}
