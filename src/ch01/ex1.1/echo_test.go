package echo

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
