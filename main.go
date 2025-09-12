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
	botToken             = lookupRequired("BOT_TOKEN")                  //  Bot access token
	guildID              = lookup("GUILD_ID", "")                       //  If not passed - bot registers commands globally
	removeCommandsOnExit = lookupBool("REMOVE_COMMANDS_ON_EXIT", false) // Remove all commands after shutdowning or not, defaults to false

	commands = []*discordgo.ApplicationCommand{
		{
			Name: "ping",
			// All commands and options must have a description
			// Commands/options without description will fail the registration
			// of the command.
			Description: "Pings the bot",
		},
	}

	commandHandlers = map[string]func(s *discordgo.Session, i *discordgo.InteractionCreate){
		"ping": func(s *discordgo.Session, i *discordgo.InteractionCreate) {
			s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
				Type: discordgo.InteractionResponseChannelMessageWithSource,
				Data: &discordgo.InteractionResponseData{
					Content: "Pong!",
				},
			})
		},
	}
)

const (
	botIntents = (discordgo.IntentGuildMessages |
		discordgo.IntentGuildMessageReactions |
		discordgo.IntentMessageContent |
		discordgo.IntentGuildScheduledEvents |
		discordgo.IntentGuildMessagePolls)
)

func init() {
	var err error
	session, err = discordgo.New("Bot " + botToken)
	if err != nil {
		log.Panic("Error creating bot instance, please make sure BOT_TOKEN is a valid Discord bot token")
	}

	// Register command handlers
	session.AddHandler(func(s *discordgo.Session, i *discordgo.InteractionCreate) {
		if h, ok := commandHandlers[i.ApplicationCommandData().Name]; ok {
			h(s, i)
		}
	})

	session.Identify.Intents = botIntents
}

func main() {
	session.AddHandler(func(s *discordgo.Session, r *discordgo.Ready) {
		log.Printf("Logged in as: %v#%v", s.State.User.Username, s.State.User.Discriminator)
	})

	err := session.Open()
	if err != nil {
		log.Fatalf("Cannot open the session: %v", err)
	}

	log.Println("Registering commands...")
	registeredCommands := make([]*discordgo.ApplicationCommand, len(commands))
	for i, v := range commands {
		cmd, err := session.ApplicationCommandCreate(session.State.User.ID, guildID, v)
		if err != nil {
			log.Panicf("Cannot create '%v' command: %v", v.Name, err)
		}
		registeredCommands[i] = cmd
	}

	defer session.Close()

	log.Println("Bot is now online!")

	// Ensure the bot closes when interrupt signal is sent
	stop := make(chan os.Signal, 1)
	signal.Notify(stop, syscall.SIGINT, syscall.SIGTERM, os.Interrupt)
	<-stop

	if removeCommandsOnExit {
		for _, v := range registeredCommands {
			err := session.ApplicationCommandDelete(session.State.User.ID, guildID, v.ID)
			if err != nil {
				log.Panicf("Cannot delete '%v' command: %v", v.Name, err)
			}
		}
	}

	log.Println("Gracefully shutting down.")
}
