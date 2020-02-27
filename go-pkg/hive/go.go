package hive

import (
	"bytes"
	"fmt"
	"github.com/benka-me/hive/go-pkg/generator"
	"os"
	"os/exec"
)

func (g Go) Protoc() {
	for _, f := range g.Setup.Files {
		g.oneProtoc(f)
	}
}

func (g Go) oneProtoc (file string) {
	args := []string{
		"-I=.",
		fmt.Sprintf("-I=%s/src/", gopath),
		fmt.Sprintf("-I=%s/src/github.com/gogo/protobuf/protobuf", gopath),
		fmt.Sprintf("-I=%s/src/%s/protobuf", gopath, g.Setup.Repo),
		fmt.Sprintf("--gogoslick_out=plugins=grpc:%s/src", gopath),
		file,
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
}

func (g Go) PublicFiles (bee *Bee) error {
	repo := g.Setup.Repo
	pkgName := g.Setup.PkgName
	repoPath := fmt.Sprintf("%s/src/%s", gopath, repo)
	var perm os.FileMode = 0777 //TODO change file mode

	//create directories
	err := os.MkdirAll(fmt.Sprintf("%s/go-pkg/http/rpc", repoPath), perm)
	if err != nil {
		return err
	}
	err = os.MkdirAll(fmt.Sprintf("%s/go-pkg/rpc-%s", repoPath, pkgName), perm)
	if err != nil {
		return err
	}
	err = os.MkdirAll(fmt.Sprintf("%s/go-pkg/%s", repoPath, pkgName), perm)
	if err != nil {
		return err
	}
	err = os.MkdirAll(fmt.Sprintf("%s/protobuf", repoPath), perm)
	if err != nil {
		return err
	}

	//generate defs.proto file
	err = generator.Code{
		Interface: bee,
		Template:  fmt.Sprintf("%s/src/github.com/benka-me/hive/go-pkg/generator/template/defs.proto", gopath),
		Target:    fmt.Sprintf("%s/src/%s/protobuf/%s.proto", gopath, repo, pkgName),
		Name:      "defs",
	}.Generate()
	if err != nil {
		return err
	}

	//generate rpc.proto file
	rpc := generator.Code{
		Interface: bee,
		Template:  fmt.Sprintf("%s/src/github.com/benka-me/hive/go-pkg/generator/template/rpc.proto", gopath),
		Target:    fmt.Sprintf("%s/src/%s/protobuf/rpc-%s.proto", gopath,repo, pkgName),
		Name:      "rpc",
	}
	err = rpc.Generate()
	if err != nil {
		return err
	}

	//generate main.go
	main := generator.Code{
		Interface: bee,
		Template:  fmt.Sprintf("%s/src/github.com/benka-me/hive/go-pkg/generator/template/go/maingo", gopath),
		Target:    fmt.Sprintf("%s/src/%s/main.go", gopath, repo),
		Name:      "main",
	}
	err = main.Generate()
	if err != nil {
		return err
	}

	//generate server-grpc_2.0.go
	server := generator.Code{
		Interface: bee,
		Template:  fmt.Sprintf("%s/src/github.com/benka-me/hive/go-pkg/generator/template/go/server-grpc-2.0go", gopath),
		Target:    fmt.Sprintf("%s/src/%s/go-pkg/http/rpc/server-grpc-2.0.go", gopath, repo),
		Name:      "server",
	}
	err = server.Generate()
	if err != nil {
		return err
	}

	//generate provider
	provider := generator.Code{
		Interface: bee,
		Template:  fmt.Sprintf("%s/src/github.com/benka-me/hive/go-pkg/generator/template/go/providergo", gopath),
		Target:    fmt.Sprintf("%s/go-pkg/rpc-%s/provider.go", repoPath, pkgName),
		Name:      "provider",
	}
	err = provider.Generate()
	if err != nil {
		return err
	}

	//generate hello-world.go
	hello := generator.Code{
		Interface: bee,
		Template:  fmt.Sprintf("%s/src/github.com/benka-me/hive/go-pkg/generator/template/go/hello-worldgo", gopath),
		Target:    fmt.Sprintf("%s/go-pkg/http/rpc/hello-world.go", repoPath),
		Name:      "hello",
	}
	err = hello.Generate()
	if err != nil {
		return err
	}

	return nil
}
