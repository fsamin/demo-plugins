package main

import (
	"C"
	"strings"
)

func reverse(s string) (result string) {
	for _, v := range s {
		result = string(v) + result
	}
	return
}

func Greetings(args ...string) string {
	name := reverse(strings.Join(args, "_"))
	return name
}
