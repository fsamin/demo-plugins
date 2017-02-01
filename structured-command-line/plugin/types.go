package plugin

type Plugin struct {
	Package string
	Cmd     string
}

type Input []string

type Output struct {
	S string
}

func (o *Output) String() string {
	return o.S
}
