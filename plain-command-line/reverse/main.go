package main

import (
	"fmt"
	"os"
	"strings"
)

func reverse(s string) (result string) {
	for _, v := range s {
		result = string(v) + result
	}
	return
}

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Invalid usage. need at least 1 argument")
		os.Exit(1)
	}
	s := strings.Join(os.Args[1:], "_")
	name := reverse(s)
	fmt.Print(name)
}
