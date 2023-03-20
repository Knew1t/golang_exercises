// Exercise 7.2:
// Write a function CountingWriter with the signature below
//
// func CountingWriter(w io.Writer) (io.Writer, *int64)
//
// that, given an io.Writer, returns a new Writer that wraps the original, and a pointer
// to an int64variable that at any moment contains the number of bytes written to the new Writer.
package main

import (
	"fmt"
	"io"
)

type WriteWrap struct {
	n int64
	w io.Writer
}

func (c *WriteWrap) Write(p []byte) (int, error) {
	*&c.n += int64(len(p)) 
	return len(p), nil
}

func CountingWriter(w io.Writer) (io.Writer, *int64) {
	var newWriter WriteWrap
	newWriter.w = w
	return w, &newWriter.n
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
	var hey WriteWrap
	hey.Write([]byte("hello"))
	fmt.Println(hey)
	fmt.Println(hey.n)
}
