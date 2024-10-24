package main


import (
	"fmt"
	"os"
	"context"
	"log"
	"cloud.google.com/go/firestore"
	"google.golang.org/api/option"
)

func main(){
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Error loading .env file: %v", err)
	}
	PROJECT_ID := os.Getenv("FIRESTORE_PROJECT_ID")
	USER_COLLECTION := os.Getenv("FIRESTORE_USER_COLLECTION_NAME")
	RKG_COLLECTION := os.Getenv("FIRESTORE_RKG_COLLECTION_NAME")
	CREDENTIALS_FILE := os.Getenv("GOOGLE_APPLICATION_CREDENTIALS")
	GITHUB_TOKEN := os.Getenv("GITHUB_TOKEN")
	if PROJECT_ID == "" || USER_COLLECTION == "" || RKG_COLLECTION == "" || CREDENTIALS_FILE == "" || GITHUB_TOKEN == "" {
		log.Fatalf("One or more required environment variables are missing")
	}
	
}