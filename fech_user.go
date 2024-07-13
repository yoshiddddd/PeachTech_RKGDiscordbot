package main

import (
	"context"
	"fmt"
	"log"
	"sort"
	// "time"

	"cloud.google.com/go/firestore"
	"google.golang.org/api/option"
)

type FirestoreData struct {
	Githubid string `firestore:"githubID"`
	Name     string `firestore:"name"`
}

type Userdata struct {
	Githubid      string
	Name          string
	Contributions int
}

func fetchUserData(projectID, collectionName, credentialsFile, githubToken string) ([]Userdata, error) {
	ctx := context.Background()
	sa := option.WithCredentialsFile(credentialsFile)
	client, err := firestore.NewClient(ctx, projectID, sa)
	if err != nil {
		return nil, fmt.Errorf("error creating Firestore client: %v", err)
	}
	defer client.Close()

	docs, err := client.Collection(collectionName).Documents(ctx).GetAll()
	if err != nil {
		return nil, fmt.Errorf("error getting documents: %v", err)
	}

	var userData []Userdata
	// _はブランク識別子→インデックス変数がいらないから
	for _, doc := range docs {
		var data FirestoreData
		if err := doc.DataTo(&data); err != nil {
			log.Printf("Error converting document data: %v", err)
			continue
		}
		contributions, err := getWeeklyContributions(data.Githubid, githubToken)
		if err != nil {
			log.Fatalf("Error getting weekly contributions: %v", err)
			return nil, err
		}
		userData = append(userData, Userdata{Githubid: data.Githubid, Name: data.Name, Contributions: contributions})
	}

	sort.Slice(userData, func(i, j int) bool {
		return userData[i].Contributions > userData[j].Contributions
	})

	return userData, nil
}
