package main

import (
	"bytes"
	"fmt"
	"log"

	docker "github.com/fsouza/go-dockerclient"
)

var (
	client *docker.Client
)

func init() {
	var err error
	client, err = docker.NewClientFromEnv()
	if err != nil {
		panic(err)
	}
}

func main() {
	p, err := registerPlugin("debian", "jessie", `echo "world"`)
	if err != nil {
		panic(err)
	}
	if err := execPlugin(p); err != nil {
		panic(err)
	}
}

func registerPlugin(r, t, c string) (*Plugin, error) {
	if err := client.PullImage(docker.PullImageOptions{
		Repository: r,
		Tag:        t,
	}, docker.AuthConfiguration{}); err != nil {
		return nil, err
	}

	log.Printf("Plugin successfully installed\n")

	plugin := &Plugin{
		Repository: r,
		Tag:        t,
		Cmd:        c,
	}

	return plugin, nil
}

func execPlugin(p *Plugin) error {
	config := docker.Config{
		AttachStdout: true,
		AttachStdin:  true,
		Image:        p.Repository + ":" + p.Tag,
		Entrypoint:   []string{"bash", "-c"},
		Cmd:          []string{p.Cmd},
	}
	opts := docker.CreateContainerOptions{Name: "Plugin", Config: &config}
	container, err := client.CreateContainer(opts)
	if err != nil {
		return err
	}

	if err := client.StartContainer(container.ID, &docker.HostConfig{}); err != nil {
		return err
	}

	log.Printf("Plugin %s successfully executed\n", p.Cmd)

	var buf bytes.Buffer
	if err := client.Logs(docker.LogsOptions{
		Container:    container.ID,
		OutputStream: &buf,
		Stdout:       true,
		Stderr:       true,
	}); err != nil {
		log.Println("Error: cannot get logs")
		return err
	}

	fmt.Println("Hello " + buf.String())

	return nil
}
