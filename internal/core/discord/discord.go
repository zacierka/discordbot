package discord

import (
	"discordbot/internal/logger"

	"github.com/bwmarrin/discordgo"
)

type Session struct {
	S *discordgo.Session
}

func New(s *discordgo.Session) (d *Session) {
	d = &Session{S: s}
	s.AddHandler(d.ready)
	return
}

func (s *Session) ready(session *discordgo.Session, event *discordgo.Ready) {
	logger.InfoLog.Println("Bot established connection to discord API")
}
