package env_test

import (
	"discordbot/internal/env"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestEnv(t *testing.T) {

	assert := assert.New(t)

	env.LoadEnv()

	var a string = "LOADED"
	var b string = os.Getenv("LOADED")

	assert.Equal(a, b, "Env File loaded if results are the same")
}
