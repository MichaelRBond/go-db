package db

import (
	"os"
	"strings"
	"testing"
)

var originalEnv map[string]string

func TestMain(m *testing.M) {
	backupEnv()
	code := m.Run()
	restoreEnv()
	os.Exit(code)
}

func backupEnv() {
	originalEnv = make(map[string]string)
	for _, env := range os.Environ() {
		parts := strings.SplitN(env, "=", 2)
		key := parts[0]
		val := parts[1]
		originalEnv[key] = val
	}
}

func restoreEnv() {
	os.Clearenv()
	for k, v := range originalEnv {
		os.Setenv(k, v)
	}
}
