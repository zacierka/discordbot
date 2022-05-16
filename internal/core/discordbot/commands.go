package discordbot

import (
	"discordbot/internal/logger"
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"strings"
	"time"

	"github.com/bwmarrin/discordgo"
)

type Players struct {
	Endpoint    string   `json:"endpoint"`
	ID          int      `json:"id"`
	Identifiers []string `json:"identifiers"`
	Name        string   `json:"name"`
	Ping        int      `json:"ping"`
}

func handleCommand(s *discordgo.Session, msg *discordgo.MessageCreate) {

	cmd, args, err := parseCommand(msg.Content)
	if err != nil {
		logger.LOGMSG("Couldnt Parse Command ", cmd)
	}
	switch cmd {
	case "derek":
		handleCommandDerek(s, msg.ChannelID)
	case "socials":
		logger.LOGMSG("Socials Command Not implemented yet")
		//handleCommandSocials(s, msg.ChannelID)
	default:
		logger.LOGMSG("Couldnt find Command ", cmd, " "+strings.Join(args, " "))
	}
}

func parseCommand(str string) (cmd string, args []string, err error) {

	var rcmd string
	var rargs []string = make([]string, 0)

	if !strings.Contains(str, " ") {
		return strings.Trim(str, " ")[1:], rargs, nil
	}

	splittext := strings.Split(str, " ")
	rcmd = splittext[0][1:]
	if len(splittext) > 1 {
		rargs = splittext[1:]
	}

	logger.LOGMSG(rcmd)

	return rcmd, rargs, nil
}

func handleCommandDerek(s *discordgo.Session, ch string) {
	/*
		HTTP GET REQUEST TO FIVEM ESCAPERP3.0 SERVER /players.json
		parse list of players for dereks name
	*/
	url := "http://147.135.31.106:30130/players.json"

	httpClient := http.Client{
		Timeout: time.Second * 2, // Timeout after 2 seconds
	}

	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		log.Fatal(err)
	}

	req.Header.Set("User-Agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/101.0.4951.67 Safari/537.36")

	res, getErr := httpClient.Do(req)
	if getErr != nil {
		log.Fatal(getErr)
	}

	if res.Body != nil {
		defer res.Body.Close()
	}

	body, readErr := ioutil.ReadAll(res.Body)
	if readErr != nil {
		log.Fatal(readErr)
	}

	var p []Players
	if err := json.Unmarshal([]byte(body), &p); err != nil {
		panic(err)
	}

	embed := &discordgo.MessageEmbed{
		Color: 0x00ff00, // Green
		Thumbnail: &discordgo.MessageEmbedThumbnail{
			URL: "https://servers-live.fivem.net/servers/icon/m3x69d/-541801529.png",
		},
		Timestamp: time.Now().Format(time.RFC3339), // Discord wants ISO8601; RFC3339 is an extension of ISO8601 and should be completely compatible.
		Title:     "Yes",
	}
	var isDerek bool = false
	for _, val := range p {
		if val.Name == "Revik" {
			isDerek = true
		}
	}

	if isDerek {
		s.ChannelMessageSendEmbed(ch, embed)
	} else {
		embed.Title = "No"
		embed.Thumbnail = nil
		s.ChannelMessageSendEmbed(ch, embed)
	}

}

func handleCommandSocials(s *discordgo.Session, ch string) {
	s.ChannelMessageSend(ch, "SOCIALS HERE")
}
