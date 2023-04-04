package github

import (
	// "bytes"
	"bytes"
	"encoding/json"
	"fmt"
	"os"
	"strconv"

	// "log"
	"net/http"
	// "net/url"
	// "strings"
	"time"
)

const (
	IssuesURL       = "https://api.github.com/search/issues"
	RepoForIssueURL = "https://api.github.com/repos/Knew1t/configs/issues"
)

type Issue struct {
	Number   int
	HTMLURL  string `json:"html_url,omitempty"`
	Title    string `json:"title"`
	State    string `json:"state"`
	User     *User
	CreateAt time.Time `json:"created_at,omitempty"`
	Body     string    `json:"body"` // in markdown format
}

type User struct {
	Login   string
	HTMLURL string `json:"html_url"`
}

func CreateIssue( /* content []string */ ) /* *IssueCreateResult, error */ {
	token := os.Getenv("GITHUB_TKN")
	newIssue := Issue{
		Title: "New Issue",
		Body:  "Test issue body",
	}
	data, err := json.Marshal(newIssue)
	if err != nil {
		panic(err)
	}
	req, err := http.NewRequest("POST", RepoForIssueURL, bytes.NewBuffer(data))
	if err != nil {
		panic(err)
	}
	req.Header.Set("Content-Type", "application/json")
	// req.Header.Add("Accept", "application/vnd.github+json")
	req.Header.Set("Authorization", "Bearer "+token)
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()
	fmt.Println(resp)
}

func ListIssues() *[]Issue {
	token := os.Getenv("GITHUB_TKN")
	req, err := http.NewRequest("GET", RepoForIssueURL+"?state=all", nil)
	if err != nil {
		panic(err)
	}
	req.Header.Set("Accept", "application/vnd.github+json")
	req.Header.Set("Authorization", "Bearer "+token)
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()
	var result []Issue
	json.NewDecoder(resp.Body).Decode(&result)
	for _, elem := range result {
		fmt.Println(elem.Number)
		fmt.Println(elem.State)
		fmt.Println(elem.Title)
		fmt.Println(elem.Body)
	}
	return &result
}


func ReadIssue(id int) {
	token := os.Getenv("GITHUB_TKN")
	req, err := http.NewRequest("GET", RepoForIssueURL+"/"+strconv.Itoa(id), nil)
	if err != nil {
		panic(err)
	}
	req.Header.Set("Accept", "application/vnd.github+json")
	req.Header.Set("Authorization", "Bearer "+token)
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()
	var result Issue
	json.NewDecoder(resp.Body).Decode(&result)
	fmt.Println(result.Number)
	fmt.Println(result.State)
	fmt.Println(result.Title)
	fmt.Println(result.Body)
}

func UpdateIssue(id int) {
	token := os.Getenv("GITHUB_TKN")
	data, err := json.Marshal(Issue{Title: "Upd", Body: "Upd", State: "closed"})
	if err != nil {
		panic(err)
	}
	req, err := http.NewRequest("PATCH", RepoForIssueURL+"/"+strconv.Itoa(id), bytes.NewBuffer(data))
	if err != nil {
		panic(err)
	}
	req.Header.Set("Accept", "application/vnd.github+json")
	req.Header.Set("Authorization", "Bearer "+token)
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()
}

func DeleteIssue(id int) {
	token := os.Getenv("GITHUB_TKN")
	req, err := http.NewRequest("DELETE", RepoForIssueURL+"/"+strconv.Itoa(id), nil)
	if err != nil {
		panic(err)
	}
	req.Header.Set("Accept", "application/vnd.github+json")
	req.Header.Set("Authorization", "Bearer "+token)
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()
	fmt.Println(resp.Status)
}
