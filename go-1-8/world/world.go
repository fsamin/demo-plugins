package main

import (
	"C"
	"strings"

	"github.com/golang-rennes/demo-plugins/go-1-8/types"
)

type worldGreeter struct{}

func (g worldGreeter) Greetings(args ...string) string {
	return "World " + strings.Join(args, " ")
}

var Greeter = types.Greeter(worldGreeter{})
