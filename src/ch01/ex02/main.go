// Exercise 1.2: Modify the echo program to print the index and value of each of its arguments, one per line.
package main

import (
  "fmt"
  "os"
)

func echoIdValue() map [int]int {
  resValue := make(map[int]int)
  for i, elem := range os.Args {
    fmt.Println(i, len(elem))
    resValue[i] = len(elem)
  }
  return resValue
}

func main() {
  fmt.Println(echoIdValue())
}


