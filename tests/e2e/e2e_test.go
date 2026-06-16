//go:build e2e
// +build e2e

package e2e

import (
	"log"
	"os"
	"testing"

	"github.com/joho/godotenv"
)

func TestMain(m *testing.M) {
	err := godotenv.Load("../../.env")
	if err != nil {
		log.Println("No .env file found, relying on environment variables")
	}

	if os.Getenv("SN_INSTANCE") == "" {
		log.Fatal("SN_INSTANCE environment variable is not set")
	}

	os.Exit(m.Run())
}
