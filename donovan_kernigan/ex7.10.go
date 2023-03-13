package main

import (
	// "bufio"
	// "bytes"
	"bufio"
	"fmt"
	// "strings"
)

type LinesCounter int
type WordsCounter int

func (l *LinesCounter) Write(p []byte) (int, error) {
	(*l) += LinesCounter(Scan(p, bufio.ScanLines) - 1)
	return int(*l), nil
}

func (w *WordsCounter) Write(p []byte) (int, error) {
	(*w) += WordsCounter(Scan(p, bufio.ScanWords) - 1)
	return int(*w), nil
}

func Scan(p []byte, fn func(p []byte, atEof bool) (int, []byte, error)) int {
	var word = []byte("")
	counter := 0
	var add int
	for start := 0; word != nil; {
		add, word, _ = fn(p[start:], true)
		start += add
		counter++
	}
	return counter
}

func main() {
	var b LinesCounter
	b.Write([]byte("hey\n hey\ney ,w"))

	fmt.Println(b)
}
