package main

import (
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/bwmarrin/discordgo"
	_ "github.com/joho/godotenv/autoload"
)

var session *discordgo.Session

var (
	botToken             = lookupRequired("BOT_TOKEN")
	guildID              = lookup("GUILD_ID", "")
	removeCommandsOnExit = lookupBool("REMOVE_COMMANDS_ON_EXIT", false)

	// commands = []*discordgo.ApplicationCommand{}
)

const (
	botIntents = (discordgo.IntentGuildMessages |
		discordgo.IntentGuildMessageReactions |
		discordgo.IntentMessageContent |
		discordgo.IntentGuildScheduledEvents |
		discordgo.IntentGuildMessagePolls)

	prefix string = "!metro"
)

func init() {
	var err error
	session, err = discordgo.New("Bot " + botToken)
	if err != nil {
		log.Panic("Error creating bot instance, please make sure BOT_TOKEN is a valid Discord bot token")
	}

	session.Identify.Intents = botIntents
}

func main() {
	session.AddHandler(sessionHandler)

	err := session.Open()
	if err != nil {
		log.Fatalf("Cannot open the session: %v", err)
	}

	// log.Println("Adding commands...")
	// registeredCommands := make([]*discordgo.ApplicationCommand, len(commands))
	// for i, v := range commands {
	// 	cmd, err := session.ApplicationCommandCreate(session.State.User.ID, guildID, v)
	// 	if err != nil {
	// 		log.Panicf("Cannot create '%v' command: %v", v.Name, err)
	// 	}
	// 	registeredCommands[i] = cmd
	// }

	defer session.Close()

	log.Println("Bot is now online!")

	// Ensure the bot closes when interrupt signal is sent
	sc := make(chan os.Signal, 1)
	signal.Notify(sc, syscall.SIGINT, syscall.SIGTERM, os.Interrupt)
	<-sc
}
