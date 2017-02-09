package main

import (
	"C"
	"strings"

	"github.com/golang-rennes/demo-plugins/go-1-8/types"
)

type reverseGreeter struct{}

func reverse(s string) (result string) {
	for _, v := range s {
		result = string(v) + result
	}
	return
}

func (g reverseGreeter) Greetings(args ...string) string {
	name := reverse(strings.Join(args, "_"))
	return name
}

var Greeter = types.Greeter(reverseGreeter{})
