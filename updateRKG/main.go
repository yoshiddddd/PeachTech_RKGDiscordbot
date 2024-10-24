package main


import (
	// "fmt"
	"os"
	"context"
	"log"
	"cloud.google.com/go/firestore"
	"google.golang.org/api/option"
	"github.com/joho/godotenv"
)
type FirestoreData struct {
	Githubid string `firestore:"githubID"`
	Name     string `firestore:"name"`
}

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
	//firesotreへの接続
	ctx := context.Background()
	sa := option.WithCredentialsFile(CREDENTIALS_FILE)
	client, err := firestore.NewClient(ctx, PROJECT_ID, sa)
	if err != nil {
		log.Fatalf("Error creating Firestore client: %v", err)
	}
	defer client.Close()
	docs, err := client.Collection(USER_COLLECTION).Documents(ctx).GetAll()
	//ユーザーごとに情報追加する
		for _, doc := range docs {
			var data FirestoreData
			doc.DataTo(&data);
			log.Printf("doc: %v", data)
		}
}