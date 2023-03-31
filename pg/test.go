package main

import (
	"encoding/json"
	"fmt"
	"log"
)

type candy struct {
	Name   string
	Amout  int
	Rating float32
}

func main() {
	var c = []candy{{Name: "mars", Amout: 2, Rating: 3.0}, {"snickers", 2, 3.2}}
	data, err := json.MarshalIndent(c, "", "    ")
	if err != nil {
		log.Fatalf("json marsh f:%s", err)
	}
	fmt.Printf("%s\n", data)
	var titles []struct{ Name string }
	if err := json.Unmarshal(data, &titles); err != nil {
		log.Fatalf("asdf: %s", err)
	}
	fmt.Println(titles)

}
