/*
 Exercise 4.11: Build a tool that lets users create, read, update, and delete GitHub issues from the command line,
 invoking their preferred text editor when substantial text input is required.
*/

package main

import (
	"ch04/ex11/github"
)

func main() {
	// github.CreateIssue()
	github.ListIssues()
	github.ReadIssue(17)
}
