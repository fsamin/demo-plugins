package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"os"
	"os/exec"

	"github.com/golang-rennes/demo-plugins/structured-command-line/plugin"
)

/*
func main() {
	p, err := registerPlugin("github.com/golang-rennes/demo-plugins/structured-command-line/world", "world")
	if err != nil {
		panic(err)
	}
	if err := execPlugin(p); err != nil {
		panic(err)
	}
}*/

func registerPlugin(p, c string) (*plugin.Plugin, error) {
	b, err := exec.Command("go", "install", p).CombinedOutput()
	if err != nil {
		log.Printf("Plugin %s %s install error %s %s\n", p, c, b, err)
		return nil, err
	}

	log.Printf("Plugin successfully installed %s\n", b)

	plugin := &plugin.Plugin{
		Cmd:     c,
		Package: p,
	}

	return plugin, nil
}

func execPlugin(p *plugin.Plugin) error {
	var args []string
	if len(os.Args) > 1 {
		args = os.Args[1:]
	}
	i := plugin.Input(args)

	if err := runPlugin(p, i); err != nil {
		return err
	}
	return nil
}

func runPlugin(p *plugin.Plugin, i plugin.Input) error {
	b, err := json.Marshal(i)
	if err != nil {
		return err
	}

	cmd := exec.Command(p.Cmd)

	var stdout bytes.Buffer
	cmd.Stdout = &stdout
	cmd.Stderr = &stdout

	stdin, err := cmd.StdinPipe()
	if err != nil {
		log.Printf("Plugin %s error: %s\n", p.Cmd, err)
		return err
	}

	if err := cmd.Start(); err != nil {
		log.Printf("Plugin %s error: %s\n", p.Cmd, err)
		return err
	}

	if _, err := stdin.Write(b); err != nil {
		log.Printf("Plugin %s error: %s\n", p.Cmd, err)
		return err
	}

	if err := stdin.Close(); err != nil {
		log.Printf("Plugin %s error: %s\n", p.Cmd, err)
		return err
	}

	if err := cmd.Wait(); err != nil {
		log.Printf("Plugin %s error: %s\n", p.Cmd, err)
		return err
	}

	log.Printf("Plugin %s successfully executed\n", p.Cmd)

	var out = &plugin.Output{}
	if err := json.Unmarshal(stdout.Bytes(), out); err != nil {
		log.Printf("Plugin %s error: %s\n", p.Cmd, err)
		return err
	}

	fmt.Printf("Hello %v\n", out)

	return nil
}

func main() {
	var pluginExecs = []struct {
		Pkg string
		Cmd string
	}{
		{"github.com/golang-rennes/demo-plugins/structured-command-line/world", "world"},
		{"github.com/golang-rennes/demo-plugins/structured-command-line/name-generator", "name-generator"},
		{"github.com/golang-rennes/demo-plugins/structured-command-line/reverse", "reverse"},
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
