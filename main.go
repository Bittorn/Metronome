package main

import (
	"log"
	"os"

	"github.com/bwmarrin/discordgo"
)

func main() {
	botToken, ok := os.LookupEnv("BOT_TOKEN")
	if !ok {
		log.Panic("Bot token was not found, please make sure BOT_TOKEN is present in your .env file")
	}

	bot, err := discordgo.New("Bot " + botToken)
	if err != nil {
		log.Panic("Error creating bot instance, please make sure BOT_TOKEN is a valid Discord bot token")
	}
}
