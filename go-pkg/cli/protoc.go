package cli

import (
	"bytes"
	"errors"
	"flag"
	"fmt"
	"github.com/benka-me/hive/go-pkg/library"
	"github.com/benka-me/hive/go-pkg/yaml"
	"os"
	"os/exec"
)

var (
	gopath string = os.Getenv("GOPATH")
)

func protoc(lib library.Library, dep string) error {
	args := []string{
		"-I=.",
		fmt.Sprintf("-I=%s/src/", gopath),
		fmt.Sprintf("-I=%s/src/github.com/gogo/protobuf/protobuf", gopath),
		fmt.Sprintf("-I=%s/src/%s/protobuf", gopath, lib.Hive[dep].Repo),
		fmt.Sprintf("--gogofaster_out=plugins=grpc:%s/src", gopath),
		fmt.Sprintf("%s/src/%s/protobuf/%s.proto",gopath, lib.Hive[dep].Repo, lib.Hive[dep].PkgName),
	}
	cmd := exec.Command("protoc", args...)
	var out bytes.Buffer
	var errs bytes.Buffer
	cmd.Stdout = &out
	cmd.Stderr= &errs
	fmt.Println(cmd.Args)
	err := cmd.Run()
	if err != nil {
		fmt.Println(errs.String())
	}
	fmt.Printf("%s\n", out.String())
	return nil
}

func (g *ProtocCommand) Run() error {
	config, err := yaml.GetConfig()
	if err != nil {
		return err
	}

	lib, err := library.GetLibray()

	for _, dep := range config.Dependencies {
		protoc(lib, dep)
	}

	return nil
}

type ProtocCommand struct {
	fs *flag.FlagSet
	name string
}

func NewProtocCommand() *ProtocCommand {
	return &ProtocCommand{
		fs: flag.NewFlagSet("protoc", flag.ContinueOnError),
	}
}

func (g *ProtocCommand) Name() string {
	return g.fs.Name()
}

func (g *ProtocCommand) Init(args []string) error {
	if len(gopath) < 3 {
		return errors.New("bad gopath: " + gopath)
	}
	return g.fs.Parse(args)
}

