/*
Exercise 4.10: Modify issues to report the results in age categories, say less than a month old, less than a year old, and more than a year old.
*/

package main

import (
	"fmt"
	"log"
	"os"
	"time"

	"ch04/ex10/github"
)

func main() {
	result, err := github.SearchIssues(os.Args[1:])
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%d issues:\n", result.TotalCount)
	oneMonthEarlier := time.Now().AddDate(0, -1, 0)
	oneYearEarlier := time.Now().AddDate(-1, 0, 0)
	fmt.Println("Less Than A Month")
	for _, item := range result.Items {
		if item.CreateAt.After(oneMonthEarlier) {
			fmt.Printf("#%-5d %s %9.9s %.55s\n", item.Number, item.CreateAt.Format("2006.01.02"), item.User.Login, item.Title)
		}
	}
	fmt.Println("Less Than A Year")
	for _, item := range result.Items {
		if item.CreateAt.After(oneYearEarlier) && item.CreateAt.Before(oneMonthEarlier) {
			fmt.Printf("#%-5d %s %9.9s %.55s\n", item.Number, item.CreateAt.Format("2006.01.02"), item.User.Login, item.Title)
		}
	}
	fmt.Println("More Than A Year old")
	for _, item := range result.Items {
		if item.CreateAt.Before(oneYearEarlier) {
			fmt.Printf("#%-5d %s %9.9s %.55s\n", item.Number, item.CreateAt.Format("2006.01.02"), item.User.Login, item.Title)
		}
	}
}
