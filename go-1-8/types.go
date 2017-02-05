package main

type Plugin struct {
	Path      string
	Greetings func(...string) string
}
