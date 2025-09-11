package main

import (
	"log"

	"github.com/bwmarrin/discordgo"
)

func init() {
	botToken := lookupRequired("BOT_TOKEN")
	guildId := lookup("GUILD_ID", "")
	removeCommandsOnExit := lookupBool("REMOVE_COMMANDS_ON_EXIT", false)

	s, err := discordgo.New("Bot " + botToken)
	if err != nil {
		log.Panic("Error creating bot instance, please make sure BOT_TOKEN is a valid Discord bot token")
	}
}

func main() {
	// do stuff
}
