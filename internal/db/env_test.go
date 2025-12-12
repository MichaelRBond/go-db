package db

import (
	"os"
	"testing"
)

func TestGetDbUsername(t *testing.T) {
	setUsername := "test-username"
	os.Clearenv()
	os.Setenv("DB_USERNAME", setUsername)

	receivedUsername := getDbUsername()

	if receivedUsername != setUsername {
		t.Errorf("expected %s, got %s", setUsername, receivedUsername)
	}
}

func TestGetDbPassword(t *testing.T) {
	setPassword := "12345"
	os.Clearenv()
	os.Setenv("DB_PASSWORD", setPassword)

	receivedPassword := getDbPassword()

	if receivedPassword != setPassword {
		t.Errorf("expected %s, got %s", setPassword, receivedPassword)
	}
}

func TestGetDbUrl(t *testing.T) {
	setUrl := "localhost:5432"
	os.Clearenv()
	os.Setenv("DB_URL", setUrl)

	receivedUrl := getDbUrl()

	if receivedUrl != setUrl {
		t.Errorf("expected %s, got %s", setUrl, receivedUrl)
	}
}
