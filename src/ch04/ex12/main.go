package main

import (
	// "encoding/json"
	// "bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"strconv"
	// "net/url"
	// "strings"
)

type ComicInfo struct {
	Num        int
	Safe_title string `json:"safe_title"`
	Transcript string
}

type Comics struct {
	mp map[int]ComicInfo
}

func parseComic(latest_id int) (*map[int]ComicInfo, error) {
	var ar []ComicInfo
	m := make(map[int]ComicInfo)
	for i := 1; i <= latest_id; i++ {
		resp, err := http.Get("https://xkcd.com/" + strconv.Itoa(i) + "/info.0.json")
		if err != nil {
			panic(err)
		}
		if resp.StatusCode != http.StatusOK {
			resp.Body.Close()
			return nil, fmt.Errorf("search query failed: %s", resp.Status)
		}

		var result ComicInfo
		if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
			resp.Body.Close()
			return nil, err
		}
		ar = append(ar, result)
		m[i] = ar[i-1]
		// fmt.Println(m[i])
	}
	return &m, nil
}

func parseIndex(id int) {
	jsonFile, err := os.Open("test.json")
	if err != nil {
		fmt.Println(err)
	}
	byteValue, _ := ioutil.ReadAll(jsonFile)
	var m map[int]ComicInfo
	json.Unmarshal(byteValue, &m)
	// json.NewEncoder(byteValue).Encode(m)
  fmt.Println("Num:", id)
  fmt.Println("Title:",m[id].Safe_title)
  fmt.Println("Transcript:\n",m[id].Transcript)
}

func CreateIndex(latest_id int) {
	ar, _ := parseComic(latest_id)
	file, _ := json.MarshalIndent(ar, "", " ")
	_ = ioutil.WriteFile("test.json", file, 0o644)
}

func main() {
  CreateIndex(2758)
	// id, _ := strconv.Atoi(os.Args[1])
	// parseIndex(id)
}
