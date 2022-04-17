package discordbot

import (
	"discordbot/internal/core/storage"

	"github.com/bwmarrin/discordgo"
)

func HandlePingDBCommand(s *discordgo.Session, m *discordgo.Message, d *storage.Storer) {
	s.ChannelMessageSend(m.ChannelID, d.GetVersion())
}
