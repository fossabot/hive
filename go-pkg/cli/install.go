package cli

import (
	"bytes"
	"flag"
	"fmt"
	"github.com/benka-me/hive/go-pkg/library"
	"github.com/benka-me/hive/go-pkg/yaml"
	"log"
	"os/exec"
)

type InstallCommand struct {
	fs *flag.FlagSet

	name string
}
func NewInstallCommand() *InstallCommand {
	gc := &InstallCommand{
		fs: flag.NewFlagSet("i", flag.ContinueOnError),
	}

	gc.fs.StringVar(&gc.name, "name", "World", "name of the person to be greeted")

	return gc
}

func (g *InstallCommand) Name() string {
	return g.fs.Name()
}

func (g *InstallCommand) Init(args []string) error {
	return g.fs.Parse(args)
}

func goGet(repo string) error {
	cmd := exec.Command("go", "get", "-d", repo+"...")
	var out bytes.Buffer
	var errs bytes.Buffer
	cmd.Stdout = &out
	cmd.Stderr= &errs
	err := cmd.Run()
	if err != nil {
		log.Fatal(errs.String(), err)
	}
	fmt.Printf("in all caps: %q\n", out.String())

	return nil
}

func (g *InstallCommand) Run() error {
	config, err := yaml.GetConfig()
	if err != nil {
		return err
	}

	lib, err := library.GetLibray()
	fmt.Println(lib)
	for _, dep := range config.Dependencies {
		err := goGet(lib.Hive[dep].Repo)
		if err != nil {
			return err
		}
	}

	return nil
}
