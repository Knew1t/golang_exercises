package main

import (
	"os"
	"reflect"
	"testing"
)

func Test(t *testing.T) {
	os.Args = []string{"hello", "dolly"}
	x := make(map[int]int)
	x = map[int]int{0: 37, 1: 5, 2: 5}
	y := echoIdValue()
	if reflect.DeepEqual(x, y) {
		t.Error("wrong")
	}
}
