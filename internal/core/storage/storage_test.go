package storage_test

import (
	"discordbot/internal/core/storage"
	"discordbot/internal/env"
	"discordbot/internal/logger"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestStorageConnection(t *testing.T) {

	assert := assert.New(t)

	env.LoadEnv("../../../.development.env")

	logger.TestLog.Println("Database PATH: ", os.Getenv("DB_ADDR"))

	storer, err := storage.New()
	if err != nil {
		t.Fail()
	}

	var a string = "8.0.28"
	var b = storer.GetVersion()

	assert.Equal(a, b, "MySql version should be the same")
}
