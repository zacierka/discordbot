package discordbot

import (
	"discordbot/internal/core/storage"
	"discordbot/internal/logger"
	"os"

	"github.com/bwmarrin/discordgo"
)

type DiscordBot struct {
	dg *discordgo.Session
}

func SetupBot() (*DiscordBot, error) {

	TOKEN := os.Getenv("TOKEN")
	if TOKEN == "" {

		logger.ErrorLog.Println("FATAL Could not retrieve TOKEN")
		return &DiscordBot{}, nil
	}

	dg, err := discordgo.New("Bot " + TOKEN)
	if err != nil {
		logger.ErrorLog.Println("Error Creating Bot Instance")
		return &DiscordBot{}, discordgo.ErrNilState
	}

	dg.Identify.Intents = discordgo.MakeIntent(discordgo.IntentsGuildMessages)
	logger.InfoLog.Println("Set Intents: IntentsGuildMessages")

	dg.AddHandler(messageCreate)
	dg.AddHandler(ready)

	logger.InfoLog.Println("Created Discord Bot Instance")

	return &DiscordBot{dg: dg}, nil
}

func (bot *DiscordBot) Open() error {
	if err := bot.dg.Open(); err != nil {
		return discordgo.ErrWSNotFound
	}
	return nil
}

func (bot *DiscordBot) Close() (bool, error) {
	err := bot.dg.Close()
	if err != nil {
		return false, discordgo.ErrStatusOffline
	}

	if ret := storage.Close(); !ret {
		logger.ErrorLog.Println("Could not close database connection on Close")
	}
	return true, nil
}

func ready(s *discordgo.Session, m *discordgo.Ready) {
	logger.InfoLog.Println("Bot Logged in as ", s.State.User.Username)
}
