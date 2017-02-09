package main

import (
	"C"
	"strings"

	"github.com/golang-rennes/demo-plugins/go-1-8/types"
)

var Greetings types.MyFunc = func(args ...string) string {
	return "World " + strings.Join(args, " ")
}
