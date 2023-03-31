package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"strings"
)

type Issue struct {
	Title string `json:"title"`
	Body  string `json:"body"`
}

func main() {
	url := "https://api.github.com/repos/:owner/:repo/issues"

	// Replace :owner and :repo with your Github username and repository name
	repoUrl := strings.Replace(url, ":owner/:repo", "Knew1t/configs", -1)

	// Create a new issue
	issue := &Issue{
		Title: "New issue title",
		Body:  "New issue description...",
	}

	// Convert the issue to json
	issueData, err := json.Marshal(issue)
	if err != nil {
		panic(err)
	}

	// Create a new http request
	req, err := http.NewRequest("POST", repoUrl, bytes.NewBuffer(issueData))
	if err != nil {
		panic(err)
	}

	// Set the Content-Type header
	req.Header.Set("Content-Type", "application/json")

	// Add your Github access token to the Authorization header
	req.Header.Set("Authorization", "Bearer TOKEN")

	// Create a new http client and perform the request
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	// Print out the response
	fmt.Println(resp.Status)
}
