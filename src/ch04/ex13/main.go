package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"
)

type Movie struct {
	Title  string
	Poster string
}

const OmdbULR = "http://www.omdbapi.com/?apikey=[yourkey]&"

func main() {
	url := strings.Replace(OmdbULR, "[yourkey]", os.Getenv("APIKEY"), -1)
	resp, err := http.Get(url + "t=" + os.Args[1])
	if err != nil {
		panic(err)
	}
	if resp.StatusCode != http.StatusOK {
		resp.Body.Close()
	}
	fmt.Println(resp)
	var data Movie
	if err = json.NewDecoder(resp.Body).Decode(&data); err != nil {
		resp.Body.Close()
	}
	resp, err = http.Get(data.Poster + ".jpg") // getting poster data
	file, err := os.Create(os.Args[1])
	if err != nil {
		panic(err)
	}
	defer file.Close()
	_, err = io.Copy(file, resp.Body)
}
