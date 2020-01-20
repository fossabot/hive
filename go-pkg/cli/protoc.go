package cli

import (
	"bytes"
	"fmt"
	"github.com/benka-me/hive/go-pkg/library"
	"github.com/benka-me/hive/go-pkg/yaml"
	"github.com/urfave/cli"
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

func runProtoc(c *cli.Context) error {
	config, err := yaml.GetConfig()
	if err != nil {
		return err
	}

	lib, err := library.GetLibrary()

	for _, dep := range config.Dependencies {
		_ = protoc(lib, dep)
	}

	return nil
}
