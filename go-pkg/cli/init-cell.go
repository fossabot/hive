package cli

import (
	"fmt"
	"github.com/benka-me/hive/go-pkg/library"
	"github.com/urfave/cli"
	"io"
	"io/ioutil"
	"os"
	"strings"
	"text/template"
)


func runInitCell(c *cli.Context) error {
	cell, err := askUser()
	if err != nil {
		return err
	}

	cell.PkgNameCamel = strings.Title(kebabToCamelCase(cell.PkgName))
	err = library.AddCellToLibrary(cell)
	if err != nil {
		return err
	}

	return createFiles(cell)
}

func createFiles(cell library.Cell) error {
	repoPath := fmt.Sprintf("%s/src/%s", gopath, cell.Repo)
	var perm os.FileMode = 0777
	cell.PkgNameCamel = strings.Title(kebabToCamelCase(cell.PkgName))

	//create directories
	err := os.MkdirAll(fmt.Sprintf("%s/go-pkg/provider", repoPath), perm)
	if err != nil {
		return err
	}
	err = os.MkdirAll(fmt.Sprintf("%s/go-pkg/http/rpc", repoPath), perm)
	if err != nil {
		return err
	}
	err = os.MkdirAll(fmt.Sprintf("%s/go-pkg/%s", repoPath, cell.PkgName), perm)
	if err != nil {
		return err
	}
	err = os.MkdirAll(fmt.Sprintf("%s/protobuf", repoPath), perm)
	if err != nil {
		return err
	}
	err = os.MkdirAll(fmt.Sprintf("%s/cmd", repoPath), perm)
	if err != nil {
		return err
	}

	//generate .proto file
	proto := Code{
		Interface: cell,
		Template:  fmt.Sprintf("%s/src/github.com/benka-me/hive/go-pkg/cli/template/template.proto", gopath),
		Target:    fmt.Sprintf("%s/src/%s/protobuf/%s.proto", gopath, cell.Repo, cell.PkgName),
		Name:      "proto",
	}
	err = proto.generator()
	if err != nil {
		return err
	}

	//generate main.go
	main := Code{
		Interface: cell,
		Template:  fmt.Sprintf("%s/src/github.com/benka-me/hive/go-pkg/cli/template/maingo", gopath),
		Target:    fmt.Sprintf("%s/src/%s/cmd/main.go", gopath, cell.Repo),
		Name:      "main",
	}
	err = main.generator()
	if err != nil {
		return err
	}

	//generate ServerGrpc_2.0.go
	server := Code{
		Interface: cell,
		Template:  fmt.Sprintf("%s/src/github.com/benka-me/hive/go-pkg/cli/template/server-grpc-2.0go", gopath),
		Target:    fmt.Sprintf("%s/src/%s/go-pkg/http/rpc/server-grpc-2.0.go", gopath, cell.Repo),
		Name:      "server",
	}
	err = server.generator()
	if err != nil {
		return err
	}

	//generate provider
	provider := Code{
		Interface: cell,
		Template:  fmt.Sprintf("%s/src/github.com/benka-me/hive/go-pkg/cli/template/providergo", gopath),
		Target:    fmt.Sprintf("%s/go-pkg/provider/provider.go", repoPath),
		Name: "provider",
	}
	err = provider.generator()
	if err != nil {
		return err
	}

	lib, err := library.GetLibrary()
	_ = protoc(lib, cell.Name)

	//copy hello-world.go
	_, err = copy(fmt.Sprintf("%s/src/github.com/benka-me/hive/go-pkg/cli/template/hello-world.go", gopath), fmt.Sprintf("%s/src/%s/go-pkg/http/rpc/hello-world.go", gopath, cell.Repo))

	return nil
}

type Code struct {
	Interface interface{}
	Template string
	Target string
	Name string
}

func (code Code) generator() error {
	dat, err := ioutil.ReadFile(code.Template)
	if err != nil {
		return err
	}

	tmpl, err := template.New(code.Name).Parse(string(dat))
	if err != nil {
		return err
	}

	f, err := os.Create(code.Target)
	if err != nil {
		return err
	}

	err = tmpl.Execute(f, code.Interface)
	if err != nil {
		return err
	}
	return nil
}

func kebabToCamelCase(kebab string) (camelCase string) {
	isToUpper := false
	for _, runeValue := range kebab {
		if isToUpper {
			camelCase += strings.ToUpper(string(runeValue))
			isToUpper = false
		} else {
			if runeValue == '-' {
				isToUpper = true
			} else {
				camelCase += string(runeValue)
			}
		}
	}
	return
}

func copy(src, dst string) (int64, error) {
	sourceFileStat, err := os.Stat(src)
	if err != nil {
		return 0, err
	}

	if !sourceFileStat.Mode().IsRegular() {
		return 0, fmt.Errorf("%s is not a regular file", src)
	}

	source, err := os.Open(src)
	if err != nil {
		return 0, err
	}
	defer source.Close()

	destination, err := os.Create(dst)
	if err != nil {
		return 0, err
	}
	defer destination.Close()
	nBytes, err := io.Copy(destination, source)
	return nBytes, err
}
