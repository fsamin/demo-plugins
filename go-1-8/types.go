package main

import "github.com/golang-rennes/demo-plugins/go-1-8/types"

type Plugin struct {
	Path string
	types.Greeter
}
