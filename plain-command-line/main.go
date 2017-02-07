package main

import (
	"fmt"
	"log"
	"os"
	"os/exec"
)

func main() {
	p, err := registerPlugin("github.com/golang-rennes/demo-plugins/plain-command-line/world", "world")
	if err != nil {
		panic(err)
	}
	if err := execPlugin(p); err != nil {
		panic(err)
	}
}

func registerPlugin(p, c string) (*Plugin, error) {
	b, err := exec.Command("go", "install", p).CombinedOutput()
	if err != nil {
		log.Println(err)
		return nil, err
	}

	log.Printf("Plugin successfully installed %s\n", b)

	plugin := &Plugin{
		Cmd:     c,
		Package: p,
	}

	return plugin, nil
}

func execPlugin(p *Plugin) error {
	var args []string
	if len(os.Args) > 1 {
		args = os.Args[1:]
	}
	if err := runPlugin(p, args...); err != nil {
		return err
	}
	return nil
}

func runPlugin(p *Plugin, args ...string) error {
	b, err := exec.Command(p.Cmd, args...).CombinedOutput()
	if err != nil {
		log.Printf("Plugin %s error: %s\n", p.Cmd, err)
		return err
	}

	log.Printf("Plugin %s successfully executed\n", p.Cmd)

	fmt.Println("Hello " + string(b))

	return nil
}

/*
func main() {
	var pluginExecs = []struct {
		Pkg string
		Cmd string
	}{
		{"github.com/golang-rennes/demo-plugins/plain-command-line/world", "world"},
		{"github.com/golang-rennes/demo-plugins/plain-command-line/name-generator", "name-generator"},
		{"github.com/golang-rennes/demo-plugins/plain-command-line/reverse", "reverse"},
	}

	for _, e := range pluginExecs {
		p, err := registerPlugin(e.Pkg, e.Cmd)
		if err != nil {
			panic(err)
		}
		if err := execPlugin(p); err != nil {
			panic(err)
		}
	}
}
*/
