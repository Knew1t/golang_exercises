package main

import (
  "fmt"
  "os"
)

func echoIdValue() {
  var s string
  for id, arg := range os.Args{
    fmt.Println(id, len(arg))
    s += string(id) + " " +string(len(arg)) + "\n"
  }
}

func main() {
  echoIdValue()
}


