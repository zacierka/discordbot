package main

import (
	"discordbot/internal/core/discordbot"
	"discordbot/internal/env"
	"discordbot/internal/logger"
	"os"
	"os/signal"
	"syscall"
)

func main() {

	env.LoadEnv()

	bot, err := discordbot.SetupBot()
	if err != nil {
		logger.ErrorLog.Println("Could not create Bot")
		os.Exit(1)
	}

	err = bot.Open()
	if err != nil {
		logger.ErrorLog.Println("A problem occurred while opening a connection.", err)
		return
	}

	sc := make(chan os.Signal, 1)
	signal.Notify(sc, syscall.SIGINT, syscall.SIGTERM, os.Interrupt, os.Kill)
	<-sc

	// Cleanly close down the Discord session.
	bot.Close()

}
