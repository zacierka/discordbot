package discordbot

import (
	"strings"

	"github.com/bwmarrin/discordgo"
)

func messageCreate(s *discordgo.Session, m *discordgo.MessageCreate) {

	if s == nil || m == nil {
		return
	}

	if m.Author.ID == s.State.User.ID {
		return
	}

	if m.Content == "" {
		return
	}

	if m.Content[0] == '.' && strings.Count(m.Content, ".") < 2 {
		executeCommand(s, m.Message)
		return
	}
}
