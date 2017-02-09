package main

import (
	"fmt"
	"log"
	"os"
	"plugin"

	"github.com/golang-rennes/demo-plugins/go-1-8/types"
)

func main() {
	var pluginExecs = []struct {
		Path string
	}{
		{"/opt/plugins/world.so"},
		{"/opt/plugins/reverse.so"},
	}

	for i := range pluginExecs {
		p, err := registerPlugin(pluginExecs[i].Path)
		if err != nil {
			panic(err)
		}
		if err := execPlugin(p); err != nil {
			panic(err)
		}
	}
}

func registerPlugin(path string) (*Plugin, error) {
	p, err := plugin.Open(path)
	if err != nil {
		return nil, err
	}

	funcSymbol, err := p.Lookup("Greetings")
	if err != nil {
		return nil, err
	}

	greet := *funcSymbol.(*types.MyFunc)

	log.Printf("Plugin successfully installed\n")

	plugin := &Plugin{
		Path:      path,
		Greetings: greet,
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
	s := p.Greetings(args...)
	log.Printf("Plugin %s successfully executed\n", p.Path)

	fmt.Println("Hello " + s)

	return nil
}
