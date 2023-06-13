package main

import ("fmt" 
"os"
"strings")

func echo() string {
  str := strings.Join(os.Args, " ")
  fmt.Println(str)
  return str
}

func main() {
  echo()
}

