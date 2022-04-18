package env

import (
	"discordbot/internal/logger"
	"os"

	"github.com/joho/godotenv"
)

func LoadEnv(path string) {
	err := godotenv.Load(path)
	if err != nil {
		logger.LOGFATAL("Error loading .env file")
	}
	loaded := os.Getenv("LOADED")
	if loaded == "" {
		loaded = "FAILED"
	}

	logger.LOGMSG("Env File Status: ", loaded)
}
