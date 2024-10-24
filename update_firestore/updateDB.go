package main

import (
	"log"
)

func updateDB(user_id string, user_data FirestoreUserData) {
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