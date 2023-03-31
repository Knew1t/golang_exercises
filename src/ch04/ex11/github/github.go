package github

import (
	// "bytes"
	"bytes"
	"encoding/json"
	"fmt"
	// "log"
	"net/http"
	// "net/url"
	// "strings"
	"time"
)

const (
	IssuesURL       = "https://api.github.com/search/issues"
	RepoForIssueURL = "https://api.github.com/repos/Knew1t/configs/issues"
	token           = " tkn"
)

// type IssuesSearchResult struct {
// 	TotalCount int `json:"total_count"`
// 	Items      []*Issue
// }
//
// type IssueCreateResult struct {
// 	Id      int
// 	HTMLURL string `json:"html_url"`
// }

type Issue struct {
	Number   int
	HTMLURL  string `json:"html_url,omitempty"`
	Title    string `json:"title"`
	State    string `json:"State,omitempty"`
	User     *User
	CreateAt time.Time `json:"created_at,omitempty"`
	Body     string    `json:"body"` // in markdown format
}

type User struct {
	Login   string
	HTMLURL string `json:"html_url"`
}

//
// SearchIssues queries the GitHub issue tracker.
// func SearchIssues(terms []string) (*IssuesSearchResult, error) {
// 	q := url.QueryEscape(strings.Join(terms, " "))
// 	fmt.Println(strings.Join(terms, " "))
// 	fmt.Println(IssuesURL + "?q=" + q)
// 	resp, err := http.Get(IssuesURL + "?q=" + q)
// 	if err != nil {
// 		return nil, err
// 	}
// 	// We must close resp.Body on all execution paths.
// 	// (Chapter 5 presents 'defer', which makes this simpler.)
// 	if resp.StatusCode != http.StatusOK {
// 		resp.Body.Close()
// 		return nil, fmt.Errorf("search query failed: %s", resp.Status)
// 	}
// 	var result IssuesSearchResult
// 	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
// 		resp.Body.Close()
// 		return nil, err
// 	}
// 	resp.Body.Close()
// 	return &result, nil
// }

func CreateIssue( /* content []string */ ) /* *IssueCreateResult, error */ {
	newIssue := Issue{
		Title: "New Issue",
		Body:  "Test issue body",
	}
	data, err := json.Marshal(newIssue)
	if err != nil {
		panic(err)
	}

	fmt.Printf("%s\n", data)
	fmt.Println(RepoForIssueURL)
	req, err := http.NewRequest("POST", RepoForIssueURL, bytes.NewBuffer(data))
	if err != nil {
		panic(err)
	}
	// req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "Bearer"+token)
	req.Header.Add("Accept", "application/vnd.github+json")
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()
	fmt.Println(resp.Status)
}
