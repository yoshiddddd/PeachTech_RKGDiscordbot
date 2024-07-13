package main

import (
	"log"
	"os"

	"github.com/aws/aws-lambda-go/lambda"
	"github.com/bwmarrin/discordgo"
	"github.com/joho/godotenv"
)

func main() {
	lambda.Start(handler)
	// handler()
}

func handler() {
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Error loading .env file: %v", err)
	}

	DISCORD_TOKEN := os.Getenv("DISCORD_BOT_TOKEN")
	PROJECT_ID := os.Getenv("FIRESTORE_PROJECT_ID")
	COLLECTION_NAME := os.Getenv("FIRESTORE_COLLECTION_NAME")
	DISCORD_CHANNEL_ID := os.Getenv("DISCORD_CHANNEL_ID")
	CREDENTIALS_FILE := os.Getenv("GOOGLE_APPLICATION_CREDENTIALS")
	GITHUB_TOKEN := os.Getenv("GITHUB_TOKEN")

	if DISCORD_TOKEN == "" || PROJECT_ID == "" || COLLECTION_NAME == "" || DISCORD_CHANNEL_ID == "" || CREDENTIALS_FILE == "" || GITHUB_TOKEN == "" {
		log.Fatalf("One or more required environment variables are missing")
	}

	dg, err := discordgo.New("Bot " + DISCORD_TOKEN)
	if err != nil {
		log.Fatalf("Error creating Discord session: %v", err)
	}
	err = dg.Open()
	if err != nil {
		log.Fatalf("Error opening Discord session: %v", err)
	}
	defer dg.Close()

	userData, err := fetchUserData(PROJECT_ID, COLLECTION_NAME, CREDENTIALS_FILE, GITHUB_TOKEN)
	if err != nil {
		log.Fatalf("Error fetching user data: %v", err)
		return
	}

	err = sendDiscordMessage(dg, DISCORD_CHANNEL_ID, userData)
	if err != nil {
		log.Printf("Error sending embed message: %v", err)
	} else {
		log.Printf("Embed message sent successfully")
	}
}

