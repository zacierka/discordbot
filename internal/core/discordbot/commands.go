package discordbot

import (
	"discordbot/internal/core/storage"
	"discordbot/internal/logger"
	"strings"
	"time"

	"github.com/bwmarrin/discordgo"
)

var Db *storage.SqlStorer

func init() {
	Db, _ = storage.New()
}

func executeCommand(s *discordgo.Session, m *discordgo.Message) {

	msg := strings.Split(strings.TrimSpace(m.Content), ".")[1]

	if len(msg) > 2 {
		msg = strings.Split(strings.Split(m.Content, " ")[0], ".")[1]
	}

	logger.InfoLog.Println("Command Called ", msg)
	switch msg {
	case "info":
		HandlePingDBCommand(s, m, Db)
	default:
		HandleUnknownCommand(s, m, msg)
	}
}

func HandleUnknownCommand(s *discordgo.Session, m *discordgo.Message, msg string) {

	c, err := s.UserChannelCreate(m.Author.ID)
	if err != nil {
		println("Unable to open User Channel: ", err)
		return
	}
	s.ChannelMessageSend(c.ID, "The command ` "+msg+" ` is not recognized.")
}

func HandleWrongPermissions(s *discordgo.Session, m *discordgo.Message, msg string) {

	c, err := s.UserChannelCreate(m.Author.ID)
	if err != nil {
		println("Unable to open User Channel: ", err)
		return
	}
	s.ChannelMessageSend(c.ID, "The command ` "+msg+" ` is not available to you.")
}

func SendTimedMessage(s *discordgo.Session, m *discordgo.Message, msg string, length time.Duration) {
	res, _ := s.ChannelMessageSend(m.ChannelID, msg)
	time.AfterFunc(length*time.Second, func() {
		s.ChannelMessageDelete(m.ChannelID, res.ID)
	})
}
