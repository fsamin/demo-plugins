package main

import (
	"C"
	"strings"

	"github.com/golang-rennes/demo-plugins/go-1-8/types"
)

func reverse(s string) (result string) {
	for _, v := range s {
		result = string(v) + result
	}
	return
}

var Greetings types.MyFunc = func(args ...string) string {
	name := reverse(strings.Join(args, "_"))
	return name
}
