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
var PROJECT_ID string
var USER_COLLECTION string
var RKG_COLLECTION string
var CREDENTIALS_FILE string
var GITHUB_TOKEN string
var client *firestore.Client
var ctx context.Context
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
	user_docs, err := client.Collection(USER_COLLECTION).Documents(ctx).GetAll()
	//ユーザーごとに情報追加する
		for _, doc := range user_docs {
			var user_data FirestoreUserData
			doc.DataTo(&user_data);
			// log.Printf("doc: %v", user_data)
			user_id := user_data.Githubid
			// updateDB(user_id, user_data )
			contributions, err := getContributions(user_id, GITHUB_TOKEN,true)
			if err != nil {
				log.Fatalf("Error getting weekly contributions: %v", err)
				return
			}
			year_contributions, err := getContributions(user_id, GITHUB_TOKEN,false)
			if err != nil {
				log.Fatalf("Error getting weekly contributions: %v", err)
				return
			}
			_, err = client.Collection(RKG_COLLECTION).Doc(user_id).Set(ctx, map[string]interface{}{
				"name"	: user_data.Name,
				"user_id": user_id,
				"contributions": contributions,
				"year_contributions": year_contributions,
			})
			if err != nil {
				// Handle any errors in an appropriate way, such as returning them.
				log.Printf("An error has occurred: %s", err)
			}
		}
}