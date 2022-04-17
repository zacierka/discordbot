package env

import (
	"discordbot/internal/logger"
	"os"

	"github.com/joho/godotenv"
)

func LoadEnv() {
	err := godotenv.Load("../../.env")
	if err != nil {
		logger.ErrorLog.Fatal("Error loading .env file")
	}
	loaded := os.Getenv("LOADED")

	logger.InfoLog.Println("Loaded .env file ", loaded)
}
