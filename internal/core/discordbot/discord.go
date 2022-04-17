package discordbot

import (
	"bot/internal/logger"
	"discordbot/internal/core/discord"
	"discordbot/internal/logger"
	"os"

	"github.com/bwmarrin/discordgo"
)

func (app *App) ConnectDiscord() (err error) {
	s, err := discordgo.New("Bot " + os.Getenv("TOKEN"))
	if err != nil {
		logger.ErrorLog.Fatalln("TOKEN INVALID")
	}

	app.discordClient = discord.New(s)

	app.discordClient.S.AddHandler(app.onReady)

	intent := discordgo.MakeIntent(discordgo.IntentsAllWithoutPrivileged | discordgo.IntentsGuildMembers)
	logger.InfoLog.Println("Added Intents")
	app.discordClient.S.Identify.Intents = intent

	err = app.discordClient.S.Open()
	if err != nil {
		return
	}

	err = <-app.ready
	if err != nil {
		return
	}

	app.discordClient.S.UpdateGameStatus(0, os.Getenv("INITSTATUS"))

	return
}

func (app *App) onReady(s *discordgo.Session, event *discordgo.Ready) {
	app.ready <- func() error {
		return nil
	}()
}
