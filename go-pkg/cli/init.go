package cli

import (
	"flag"
	"fmt"
)

type InitCommand struct {
	fs *flag.FlagSet

	name string
}
func NewInitCommand() *InitCommand {
	gc := &InitCommand{
		fs: flag.NewFlagSet("init", flag.ContinueOnError),
	}

	gc.fs.StringVar(&gc.name, "name", "World", "name of the person to be greeted")

	return gc
}


func (g *InitCommand) Name() string {
	return g.fs.Name()
}

func (g *InitCommand) Init(args []string) error {
	return g.fs.Parse(args)
}

func (g *InitCommand) Run() error {
	fmt.Println("Hello", g.name, "!")
	return nil
}
