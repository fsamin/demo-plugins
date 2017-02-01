package main

import (
	"encoding/json"
	"fmt"

	"github.com/golang-rennes/demo-plugins/structured-command-line/plugin"
)

func main() {
	s := plugin.Output{"World"}
	b, _ := json.Marshal(s)
	fmt.Print(string(b))
}
