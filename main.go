package main

import (
	"log"
	"os"
	"strconv"

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

func lookup(envString string, defaultValue string) string {
	value, ok := os.LookupEnv(envString)
	if !ok {
		log.Printf("Optional env variable %s was not found, using default value %s \n", envString, defaultValue)
		return defaultValue
	}
	return value
}

func lookupRequired(envString string) string {
	value, ok := os.LookupEnv(envString)
	if !ok {
		log.Panicf("Required env variable %s was not found, please make sure it is present in your .env file \n", envString)
	}
	return value
}

func lookupBool(envString string, defaultValue bool) bool {
	l := lookup("REMOVE_COMMANDS_ON_EXIT", "false")
	v, err := strconv.ParseBool(l)
	if err != nil {
		log.Printf("Error parsing env variable %s: expected bool, received %s \n", envString, l)
		return defaultValue
	}
	return v
}

func main() {
	// do stuff
}
