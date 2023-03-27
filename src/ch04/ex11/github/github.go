package github

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"net/url"
	"strings"
	"time"
)

const (
	IssuesURL       = "https://api.github.com/search/issues"
	RepoForIssueURL = "https://api.github.com/repos/Knew1t/configs/issues/"
	token           = " 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"
)

type IssuesSearchResult struct {
	TotalCount int `json:"total_count"`
	Items      []*Issue
}

type IssueCreateResult struct {
	Id      int
	HTMLURL string `json:"html_url"`
}

type Issue struct {
	Number   int
	HTMLURL  string `json:"html_url"`
	Title    string
	State    string
	User     *User
	CreateAt time.Time `json:"created_at"`
	Body     string    // in markdown format
}

type User struct {
	Login   string
	HTMLURL string `json:"html_url"`
}

// SearchIssues queries the GitHub issue tracker.
func SearchIssues(terms []string) (*IssuesSearchResult, error) {
	q := url.QueryEscape(strings.Join(terms, " "))
	fmt.Println(strings.Join(terms, " "))
	fmt.Println(IssuesURL + "?q=" + q)
	resp, err := http.Get(IssuesURL + "?q=" + q)
	if err != nil {
		return nil, err
	}
	// We must close resp.Body on all execution paths.
	// (Chapter 5 presents 'defer', which makes this simpler.)
	if resp.StatusCode != http.StatusOK {
		resp.Body.Close()
		return nil, fmt.Errorf("search query failed: %s", resp.Status)
	}
	var result IssuesSearchResult
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		resp.Body.Close()
		return nil, err
	}
	resp.Body.Close()
	return &result, nil
}

func CreateIssue( /* content []string */ ) (*IssueCreateResult, error) {
	var result IssueCreateResult
	client := &http.Client{}
	issue := Issue{Number: 1, Title: "test", CreateAt: time.Now(), Body: "I am creating an Issue"}
	json_issue, err := json.Marshal(issue)
	if err != nil {
		log.Fatalf("Json marshaling failed %s", err)
	}
	issueRequest, _ := http.NewRequest("POST", RepoForIssueURL, bytes.NewBuffer(json_issue))
	issueRequest.Header.Set("Authorisation", "token "+token)
	issueRequest.Header.Set("Accept", "application/vnd.github+json")
	issueRequest.Header.Set("Content-type", "application/json")

	fmt.Println("ISSUEREQUEST")
	fmt.Println(issueRequest)
	resp, err := client.Do(issueRequest)

	fmt.Println("RESP")
	fmt.Println(resp)
	if err != nil {
		log.Fatal(err)
	}
	// resp, err := http.Post(RepoForIssueURL, "application/json", bytes.NewBuffer(json_issue))
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		resp.Body.Close()
		return nil, err
	}
	return &result, nil
}
