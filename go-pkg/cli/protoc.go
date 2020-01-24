package cli

import (
	"bytes"
	"errors"
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

func protoc(lib library.Library, dep, pbFile string) error {
	args := []string{
		"-I=.",
		fmt.Sprintf("-I=%s/src/", gopath),
		fmt.Sprintf("-I=%s/src/github.com/gogo/protobuf/protobuf", gopath),
		fmt.Sprintf("-I=%s/src/%s/protobuf", gopath, lib.Hive[dep].Repo),
		fmt.Sprintf("--gogoslick_out=plugins=grpc:%s/src", gopath),
		pbFile,
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
	lib, err := library.GetLibrary()
	if err != nil {
		return err
	}

	config, err := yaml.GetHiveConfig()
	if err == nil {
		for _, dep := range config.Dependencies {
			_ = protoc(lib, dep, fmt.Sprintf("%s/src/%s/protobuf/%s.proto",gopath, lib.Hive[dep].Repo, lib.Hive[dep].PkgName))
		}
	} else {
		cell, err := yaml.GetCellConfig()
		if err != nil {
			return errors.New("no config file founded")
		}
		_ = protoc(lib, cell.Name, fmt.Sprintf("%s/src/%s/protobuf/%s.proto",gopath, cell.Repo, cell.PkgName))
		_ = protoc(lib, cell.Name, fmt.Sprintf("%s/src/%s/protobuf/rpc-%s.proto",gopath, cell.Repo, cell.PkgName))
	}


	return nil
}
