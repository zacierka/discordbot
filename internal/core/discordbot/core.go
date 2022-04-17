package discordbot

import (
	"discordbot/internal/core/discord"
	"discordbot/internal/core/storage"
	"discordbot/internal/logger"
	"os"
	"os/signal"
	"syscall"
)

type App struct {
	discordClient *discord.Session
	BotID         string
	storage       storage.Storer
	ready         chan error
}

func Start() {
	var app = &App{}
	var err error

	app.storage, err = storage.New()
	if err != nil {
		logger.ErrorLog.Println("Could not initialize storage on app")
	}

	err = app.ConnectDiscord()
	if err != nil {
		logger.ErrorLog.Fatalln("Could not initialize bot")
	}

	signals := make(chan os.Signal, 1)
	signal.Notify(signals, syscall.SIGTERM, syscall.SIGKILL)
	<-signals
}
