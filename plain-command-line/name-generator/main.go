package main

import (
	"fmt"

	"os"
	"strings"

	"github.com/docker/docker/pkg/namesgenerator"
)

func main() {
	name := namesgenerator.GetRandomName(0)
	if len(os.Args) > 1 {
		s := strings.Join(os.Args[1:], "_")
		fmt.Print(name + "_" + s)
		return
	}
	fmt.Print(name)
}
