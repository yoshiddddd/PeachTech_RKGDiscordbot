package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"time"
)

const githubAPI = "https://api.github.com/graphql"
const query = `
query($username: String!, $from: DateTime!, $to: DateTime!) {
  user(login: $username) {
    contributionsCollection(from: $from, to: $to) {
      contributionCalendar {
        totalContributions
      }
    }
  }
}
`

type Variables struct {
	Username string `json:"username"`
	From     string `json:"from"`
	To       string `json:"to"`
}

type GraphQLRequest struct {
	Query     string   `json:"query"`
	Variables Variables `json:"variables"`
}

type ContributionCalendar struct {
	TotalContributions int `json:"totalContributions"`
}

type ContributionsCollection struct {
	ContributionCalendar ContributionCalendar `json:"contributionCalendar"`
}

type User struct {
	ContributionsCollection ContributionsCollection `json:"contributionsCollection"`
}

type Data struct {
	User User `json:"user"`
}

type GraphQLResponse struct {
	Data Data `json:"data"`
}

func getWeeklyContributions(username, token string) (int, error) {
	jst := time.FixedZone("JST", 9*60*60)
	now := time.Now().In(jst)
	oneDayAgo := now.AddDate(0, 0, -1)
	oneWeekAgo := now.AddDate(0, 0, -7)

	variables := Variables{
		Username: username,
		From:     oneWeekAgo.Format(time.RFC3339),
		To:       oneDayAgo.Format(time.RFC3339),
	}

	requestBody := GraphQLRequest{
		Query:     query,
		Variables: variables,
	}

	body, err := json.Marshal(requestBody)
	if err != nil {
		return 0, fmt.Errorf("failed to marshal request body: %v", err)
	}

	req, err := http.NewRequest("POST", githubAPI, bytes.NewBuffer(body))
	if err != nil {
		return 0, fmt.Errorf("failed to create request: %v", err)
	}

	req.Header.Set("Authorization", "Bearer "+token)
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return 0, fmt.Errorf("failed to execute request: %v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return 0, fmt.Errorf("unexpected status code: %d", resp.StatusCode)
	}

	responseBody, err := io.ReadAll(resp.Body)
	if err != nil {
		return 0, fmt.Errorf("failed to read response body: %v", err)
	}

	var response GraphQLResponse
	err = json.Unmarshal(responseBody, &response)
	if err != nil {
		return 0, fmt.Errorf("failed to unmarshal response body: %v", err)
	}

	return response.Data.User.ContributionsCollection.ContributionCalendar.TotalContributions, nil
}
