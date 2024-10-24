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
type FirestoreUserData struct {
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
	log.Printf("月曜日 : %v", getThisWeekMonday())
	//firesotreへの接続
	ctx := context.Background()
	sa := option.WithCredentialsFile(CREDENTIALS_FILE)
	client, err := firestore.NewClient(ctx, PROJECT_ID, sa)
	if err != nil {
		log.Fatalf("Error creating Firestore client: %v", err)
	}
	defer client.Close()
	user_docs, err := client.Collection(USER_COLLECTION).Documents(ctx).GetAll()
	//ユーザーごとに情報追加する
		for _, doc := range user_docs {
			var user_data FirestoreUserData
			doc.DataTo(&user_data);
			// log.Printf("doc: %v", user_data)
			user_id := user_data.Githubid
			contributions, err := getWeeklyContributions(user_id, GITHUB_TOKEN)
			if err != nil {
				log.Fatalf("Error getting weekly contributions: %v", err)
				return
			}
			log.Printf("user_id: %s, contributions: %d", user_id, contributions)
		}
}