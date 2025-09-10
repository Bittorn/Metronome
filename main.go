package main

import (
	"fmt"
	"log"
	"os"
)

func main() {
	botToken, ok := os.LookupEnv("BOT_TOKEN")
	if !ok {
		log.Printf("Value not found : %v \n", ok)
		log.Panic("Your bot token is probably not set, please make sure BOT_TOKEN is present in your .env file")
	}

	fmt.Printf("Value found : %s \n", botToken)
}
