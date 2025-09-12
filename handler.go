package main

import (
	"strings"

	"github.com/bwmarrin/discordgo"
)

func sessionHandler(session *discordgo.Session, message *discordgo.MessageCreate) {
	// Guard if message was sent by us
	if message.Author.ID == session.State.User.ID {
		return
	}

	args := strings.Split(message.Content, " ")

	if args[0] != prefix {
		return
	}

	switch args[1] {
	case "ping":
		session.ChannelMessageSend(message.ChannelID, "Pong!")
	}
}
