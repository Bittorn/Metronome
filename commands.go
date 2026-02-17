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

			var total_messages int
			var total_edits int
			var author string

			var last_message_id string = "0"

			var should_continue bool = true

			for should_continue {
				st, err := s.ChannelMessages(i.ChannelID, 100, last_message_id, "0", "0")
				if err != nil {
					handleError(err, s, i)
					return
				}
				// gotta fix this, does not actually return number of items
				// log.Printf("Number of items: %d\n", len(st))

				if last_message_id == st[0].ID {
					should_continue = false
					break
				}

				for index, message := range st {
					if index == 0 {
						last_message_id = message.ID
					}
					// log.Printf("Message author: %s\n", message.EditedTimestamp)
					if message.Author.ID == options[0].Value {
						author = message.Author.Username
						total_messages += 1
						if message.EditedTimestamp != nil {
							total_edits += 1
						}
					}
				}
			}
			s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
				Type: discordgo.InteractionResponseChannelMessageWithSource,
				Data: &discordgo.InteractionResponseData{
					Content: "Of the past 100 messages, " + author + " has sent **" + strconv.Itoa(total_messages) + "**, with **" + strconv.Itoa(total_edits) + "** of them being edited.",
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
