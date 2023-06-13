package main

import (
  "testing"
  "os"
)


func Test(t *testing.T) {
  os.Args = []string{"one", "two"}
  expected := "one two"
  output := echo()
  if expected != output {
    t.Error("wrong!")
  } else {
  }
}
func Test2(t *testing.T) {
  os.Args = []string{"hey", "hey"}
  expected := "hey hey"
  output := echo()
  if expected != output {
    t.Error("wrong!")
  } else {
  }
}
