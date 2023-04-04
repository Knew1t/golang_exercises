/*
 Exercise 4.11: Build a tool that lets users create, read, update, and delete GitHub issues from the command line,
 invoking their preferred text editor when substantial text input is required.
*/

package main

import (
	"fmt"
	"os"
	"strconv"

	// "fmt"
	"ch04/ex11/github"
)

func main() {
	if len(os.Args) > 1 {
		switch arg := os.Args[1:]; arg[0] {
		case "-c":
      if len(arg) > 2 {
        github.CreateIssue()
      }

		case "-l":
			github.ListIssues()

		case "-r":
			if len(arg) == 2 {
				id, err := strconv.Atoi(arg[1])
        if err != nil {
          panic(err)
        }
				github.ReadIssue(id)
			}
		}
	} else {
    fmt.Println("usage: main [-clru] [issue id]")
    fmt.Println("-c for create")
    fmt.Println("-l for list")
    fmt.Println("-r for read (put id into it)")
    fmt.Println("-uo reopen issue")
    fmt.Println("-uc close issue")
  }
}
