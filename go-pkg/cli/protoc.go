package cli

import (
	"flag"
	"fmt"
)

type ProtocCommand struct {
	fs *flag.FlagSet

	name string
}
func NewProtocCommand() *ProtocCommand {
	gc := &ProtocCommand{
		fs: flag.NewFlagSet("protoc", flag.ContinueOnError),
	}

	gc.fs.StringVar(&gc.name, "name", "World", "name of the person to be greeted")

	return gc
}


func (g *ProtocCommand) Name() string {
	return g.fs.Name()
}

func (g *ProtocCommand) Init(args []string) error {
	return g.fs.Parse(args)
}

func (g *ProtocCommand) Run() error {
	fmt.Println("Hello", g.name, "!")
	return nil
}
