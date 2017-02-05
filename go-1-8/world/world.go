package main

import (
	"C"
	"strings"
)

func Greetings(args ...string) string {
	return "World " + strings.Join(args, " ")
}
