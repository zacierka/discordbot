package discordbot

import (
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
	app.discordClient.S.AddHandler(app.onMessage)

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

	app.BotID = app.discordClient.S.State.User.ID

	app.discordClient.S.UpdateGameStatus(0, os.Getenv("INITSTATUS"))

	return
}

func (app *App) onReady(s *discordgo.Session, event *discordgo.Ready) {
	app.ready <- func() error {
		return nil
	}()
}

func (app *App) onMessage(s *discordgo.Session, event *discordgo.MessageCreate) {
	if len(app.ready) > 0 {
		<-app.ready
	}

	ch, err := app.discordClient.S.Channel(event.ChannelID)
	if err != nil {
		return
	}
	if ch.Type != discordgo.ChannelTypeGuildText {
		return
	}

	if event.Message.Author.ID == app.BotID {
		return
	}

	logger.InfoLog.Println("Message Recieved: ", event.Content)
	s.ChannelMessageSend(ch.ID, app.storage.GetVersion())
}
