package main

import (
	"log"
	"strconv"

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
		{
			Name:        "count-messages",
			Description: "Counts the messages sent by the given user",
			Options: []*discordgo.ApplicationCommandOption{
				{
					Type:        discordgo.ApplicationCommandOptionUser,
					Name:        "user",
					Description: "The user to count messages from",
					Required:    true,
				},
			},
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
		"count-messages": func(s *discordgo.Session, i *discordgo.InteractionCreate) {
			options := i.ApplicationCommandData().Options

			st, err := s.ChannelMessages(i.ChannelID, 100, "0", "0", "0")
			if err != nil {
				handleError(err, s, i)
				return
			}

			log.Printf("Counting user messages: %s\n", options[0].Value)
			var total_messages int
			var total_edits int
			var author string

			var should_continue bool = true
			for _, message := range st {
				log.Printf("Message author: %s\n", message.EditedTimestamp)
				if message.Author.ID == options[0].Value {
					author = message.Author.Username
					total_messages += 1
					if message.EditedTimestamp != nil {
						total_edits += 1
					}
				}
				if len(st) != 100 {
					should_continue = false
				}
			}
			s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
				Type: discordgo.InteractionResponseChannelMessageWithSource,
				Data: &discordgo.InteractionResponseData{
					Content: author + " has sent **" + strconv.Itoa(total_messages) + "** messages, with **" + strconv.Itoa(total_edits) + "** of them being edited.",
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
