package main

import (
	"C"
	"strings"
)

type worldGreeter struct{}

func (g worldGreeter) Greetings(args ...string) string {
	return "World " + strings.Join(args, " ")
}

var Greeter = worldGreeter{}
