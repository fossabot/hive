package generator

import (
	"bytes"
	"fmt"
	"github.com/benka-me/hive/go-pkg/hive"
	"os"
	"os/exec"
)

type Go hive.Go

func (g Go) ClientsFiles(bee *hive.Bee) error {
	repo := bee.Repo
	repoPath := fmt.Sprintf("%s/src/%s", gopath, repo)

	//generate clients
	err := Code{
		Interface: bee,
		Target:    fmt.Sprintf("%s/go-pkg/http/rpc/clients.go", repoPath),
		Name:      "clients",
	}.GenerateClientsGo()
	if err != nil {
		return err
	}

	return nil
}

func (g Go) EntryPointFiles(bee *hive.Bee) error {
	repo := bee.Repo
	pkgName := bee.PkgName
	repoPath := fmt.Sprintf("%s/src/%s", gopath, repo)
	var perm os.FileMode = 0777 //TODO change file mode

	//create directories
	err := os.MkdirAll(fmt.Sprintf("%s/go-pkg/http/rpc", repoPath), perm)
	if err != nil {return err}

	err = os.MkdirAll(fmt.Sprintf("%s/go-pkg/rpc-%s", repoPath, pkgName), perm)
	if err != nil {return err}

	err = os.MkdirAll(fmt.Sprintf("%s/go-pkg/%s", repoPath, pkgName), perm)
	if err != nil {return err}

	err = os.MkdirAll(fmt.Sprintf("%s/protobuf", repoPath), perm)
	if err != nil {return err}

	//generate defs.proto file
	err = Code{
		Interface: bee,
		Template:  fmt.Sprintf("%s/defs.proto", ProtobufTemplates),
		Target:    fmt.Sprintf("%s/src/%s/protobuf/%s.proto", gopath, repo, pkgName),
		Name:      "defs",
	}.Generate()
	if err != nil {
		return err
	}

	//generate rpc.proto file
	rpc := Code{
		Interface: bee,
		Template:  fmt.Sprintf("%s/rpc.proto", ProtobufTemplates),
		Target:    fmt.Sprintf("%s/src/%s/protobuf/rpc-%s.proto", gopath,repo, pkgName),
		Name:      "rpc",
	}
	err = rpc.Generate()
	if err != nil {
		return err
	}

	//generate main.go
	main := Code{
		Interface: bee,
		Template:  fmt.Sprintf("%s/main-go", GoTemplates),
		Target:    fmt.Sprintf("%s/src/%s/main.go", gopath, repo),
		Name:      "main",
	}
	err = main.Generate()
	if err != nil {
		return err
	}

	//generate server-grpc_2.0.go
	server := Code{
		Interface: bee,
		Template:  fmt.Sprintf("%s/server-grpc-2.0-go", GoTemplates),
		Target:    fmt.Sprintf("%s/src/%s/go-pkg/http/rpc/server-grpc-2.0.go", gopath, repo),
		Name:      "server",
	}
	err = server.Generate()
	if err != nil {
		return err
	}

	//generate hello-world.go
	hello := Code{
		Interface: bee,
		Template:  fmt.Sprintf("%s/hello-world-go", GoTemplates),
		Target:    fmt.Sprintf("%s/go-pkg/http/rpc/hello-world.go", repoPath),
		Name:      "hello",
	}
	err = hello.Generate()
	if err != nil {
		return err
	}
	return nil
}

func (g Go) Protoc(bee *hive.Bee) {
	//repoPath := fmt.Sprintf("%s/src/%s", gopath, bee.Repo)
	//args := make([]string, 5)
	//args = []string{
	//	"-I=.",
	//	fmt.Sprintf("-I=%s/src/", gopath),
	//	fmt.Sprintf("-I=%s/src/github.com/gogo/protobuf/protobuf", gopath),
	//	fmt.Sprintf("-I=%s/src/%s/protobuf", gopath, bee.Repo),
	//	fmt.Sprintf("--%s_out=plugins=grpc:%s/src", "gogoslick",  gopath),
	//}
	//
	for _, f := range bee.ProtoSetup.Files {
		//args = append(args, fmt.Sprintf("%s/protobuf/%s", repoPath, f))
		g.oneProtoc(f, bee.Repo)
	}
	//runProtocCommand(args)
}

func (g Go) oneProtoc (file, repo string) { //TODO delete this
	args := []string{
		"-I=.",
		fmt.Sprintf("-I=%s/src/", gopath),
		fmt.Sprintf("-I=%s/src/github.com/gogo/protobuf/protobuf", gopath),
		fmt.Sprintf("-I=%s/src/%s/protobuf", gopath, repo),
		fmt.Sprintf("--%s_out=plugins=grpc:%s/src", "gogoslick",  gopath),
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

