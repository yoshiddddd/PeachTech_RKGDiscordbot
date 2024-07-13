package main
import
(
	"fmt"
	"os"
	"context"
	// "log"
	"cloud.google.com/go/firestore"
	// "google.golang.org/api/option"
)
func pushRankingData(userData []Userdata) error {
	PROJECT_ID := os.Getenv("FIRESTORE_PROJECT_ID")
	ctx := context.Background()
	client, err := firestore.NewClient(ctx, PROJECT_ID)
	if err != nil {
		return fmt.Errorf("error creating Firestore client: %v", err)
	}
	defer client.Close()

	userCollection := client.Collection("users")
	firststiter := userCollection.Where("githubID", "==", userData[0].Githubid).Documents(ctx)
	seconditer := userCollection.Where("githubID", "==", userData[1].Githubid).Documents(ctx)
	thrditer := userCollection.Where("githubID", "==", userData[2].Githubid).Documents(ctx)
	firstdoc , err := firststiter.Next() 
	seconddoc , err := seconditer.Next() 
	threedoc , err := thrditer.Next() 
	// if err == firestore.Done{
	// 	return fmt.Errorf("error getting document: %v", err)
	// }

	_, err = firstdoc.Ref.Update(ctx, []firestore.Update{
		{Path: "1st_num", Value: firestore.Increment(1)},
	})
	_, err = seconddoc.Ref.Update(ctx, []firestore.Update{
		{Path: "2nd_num", Value: firestore.Increment(1)},
	})
	_, err = threedoc.Ref.Update(ctx, []firestore.Update{
		{Path: "3rd_num", Value: firestore.Increment(1)},
	})
	if err != nil {
		return fmt.Errorf("error updating document: %v", err)
	}

	fmt.Printf("userData1ster: %s\n", userData[0].Githubid)
	return err
}
