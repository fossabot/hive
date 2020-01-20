package cli

import (
	"fmt"
	"github.com/benka-me/hive/go-pkg/library"
	"github.com/benka-me/hive/go-pkg/yaml"
	"github.com/urfave/cli"
	"io/ioutil"
	"os"
	"text/template"
)


func runInitCell(c *cli.Context) error {
	cell, err := yaml.InitYaml()
	if err != nil {
		return err
	}

	err = library.AddCellToLibrary(cell)
	if err != nil {
		return err
	}

	createFiles(cell)
	return nil
}

func createFiles(cell library.Cell) error {
	dat, err := ioutil.ReadFile(fmt.Sprintf("%s/src/github.com/benka-me/hive/go-pkg/cli/template/template.proto", gopath))
	if err != nil {
		return err
	}

	repoPath := fmt.Sprintf("%s/src/%s", gopath, cell.Repo)

	tmpl, err := template.New("test").Parse(string(dat))
	if err != nil {
		return err
	}

	fmt.Println("========", repoPath)

	var perm os.FileMode
	perm = 0644
	err = os.MkdirAll(fmt.Sprintf("%s/go-pkg/http/rpc", repoPath), perm)
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

	f, err := os.Create(fmt.Sprintf("%s/src/%s/protobuf/%s.proto", gopath, cell.Repo, cell.PkgName))
	if err != nil {
		return err
	}

	err = tmpl.Execute(f, cell)
	if err != nil {
		return err
	}

	return nil
}
