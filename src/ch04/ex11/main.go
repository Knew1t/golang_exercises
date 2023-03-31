/*
 Exercise 4.11: Build a tool that lets users create, read, update, and delete GitHub issues from the command line,
 invoking their preferred text editor when substantial text input is required.
*/

package main

import (
	// "fmt"
	// "os"
	// "log"
	// "os"

	"ch04/ex11/github"
)

func main() {
	// result, err := github.SearchIssues(os.Args[1:])
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// fmt.Printf("%d issues:\n", result.TotalCount)
	// fmt.Println("Less Than A Month")
	// for _, item := range result.Items {
	// 		fmt.Printf("#%-5d %s %9.9s %.55s\n", item.Number, item.CreateAt.Format("2006.01.02"), item.User.Login, item.Title)
	// }
	github.CreateIssue()
	// fmt.Println("response")
	// fmt.Println(result.HTMLURL)
}
