package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"strings"

	"github.com/docker/docker/pkg/namesgenerator"
	"github.com/golang-rennes/demo-plugins/structured-command-line/plugin"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	b, err := ioutil.ReadAll(reader)
	if err != nil {
		fmt.Printf("Error reading stdin : %s", err)
		os.Exit(1)
	}

	var in = &plugin.Input{}
	if err := json.Unmarshal(b, in); err != nil {
		fmt.Printf("Error unmarshalling stdin : %s", err)
		os.Exit(1)
	}

	name := namesgenerator.GetRandomName(0)
	if len(*in) > 1 {
		s := strings.Join([]string(*in), "_")
		out := plugin.Output{name + "_" + s}
		bOut, _ := json.Marshal(out)
		fmt.Print(string(bOut))
		return
	}

	out := plugin.Output{name}
	bOut, _ := json.Marshal(out)
	fmt.Print(string(bOut))
}
