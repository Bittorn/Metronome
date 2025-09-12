package main

import (
	"log"

	"github.com/bwmarrin/discordgo"
)

var (
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
			e, err := s.ApplicationEmoji(i.AppID, "1416080985417584741")
			if err != nil {
				handleError(err, s, i)
				return
			}
			s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
				Type: discordgo.InteractionResponseChannelMessageWithSource,
				Data: &discordgo.InteractionResponseData{
					Flags:   discordgo.MessageFlagsEphemeral,
					Content: e.MessageFormat() + " Pong!",
				},
			})
		},
	}
)

func handleError(err error, s *discordgo.Session, i *discordgo.InteractionCreate) {
	s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
		Type: discordgo.InteractionResponseChannelMessageWithSource,
		Data: &discordgo.InteractionResponseData{
			Flags:   discordgo.MessageFlagsEphemeral,
			Content: "Oops, something went wrong!",
		},
	})
	log.Printf("Error when processing application emoji: %s\n", err)
}
