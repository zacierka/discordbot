package main

import (
	"discordbot/internal/core/discordbot"
	"discordbot/internal/env"
)

func main() {

	env.LoadEnv(".env")

	discordbot.Start()
}
