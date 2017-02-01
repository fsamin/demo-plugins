package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"strings"

	"github.com/golang-rennes/demo-plugins/structured-command-line/plugin"
)

func reverse(s string) (result string) {
	for _, v := range s {
		result = string(v) + result
	}
	return
}

func main() {
	reader := bufio.NewReader(os.Stdin)
	b, err := ioutil.ReadAll(reader)
	if err != nil {
		fmt.Printf("Error reading stdin : %s", err)
		os.Exit(1)
	}

	var in = plugin.Input{}
	if err := json.Unmarshal(b, &in); err != nil {
		fmt.Printf("Error unmarshalling stdin : %s", err)
		os.Exit(1)
	}

	s := strings.Join([]string(in), "_")
	name := reverse(s)

	out := plugin.Output{name}
	bOut, _ := json.Marshal(out)
	fmt.Print(string(bOut))
}
