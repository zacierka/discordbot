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
		logger.LOGFATAL("TOKEN INVALID")
	}

	app.discordClient = discord.New(s)

	app.discordClient.S.AddHandler(app.onReady)
	app.discordClient.S.AddHandler(app.onMessage)

	intent := discordgo.MakeIntent(discordgo.IntentGuildMessages | discordgo.IntentGuildMembers)
	logger.LOGMSG("Added Intents")
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

	if bot := event.Author.Bot; bot {
		return
	}

	ch, err := app.discordClient.S.Channel(event.ChannelID)
	if err != nil {
		return
	}
	if ch.Type != discordgo.ChannelTypeGuildText {
		return
	}

	logger.LOGMSG("Message Recieved: ", event.Content)
	// Add some message handler here.
	// CommandManager of some sort
	//  - feature to include:
	//    + enable/disable feature for group of commands requiring database connection
	//    + enable/disable feature for group of commands requiring minecraft server status to be active
	//  - maybe something along the lines of a module system where command categories can be disabled due to limitations or services
}
