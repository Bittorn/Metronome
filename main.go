package main

import (
	"fmt"
	"log"
	"os"
)

func main() {
	botToken, ok := os.LookupEnv("BOT_TOKEN")
	if !ok {
		log.Panic("Bot token was not found, please make sure BOT_TOKEN is present in your .env file")
	}

	fmt.Printf("Value found : %s \n", botToken)
}
