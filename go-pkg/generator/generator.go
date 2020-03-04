package generator

import (
	"fmt"
	_if "github.com/benka-me/hive/go-pkg/cli/if"
	"github.com/benka-me/hive/go-pkg/hive"
	"io/ioutil"
	"os"
	"text/template"
)


type Protobuf struct {
	Repo string
	Files []string
}

type Code struct {
	Interface interface{}
	Template string
	Target string
	Name string
}

var gopath = os.Getenv("GOPATH")
var Templates = fmt.Sprintf("%s/src/github.com/benka-me/hive/go-pkg/generator/templates", gopath)
var GoTemplates = fmt.Sprintf("%s/go", Templates)
var ProtobufTemplates = fmt.Sprintf("%s/protobuf", Templates)

func GenerateEntryPointFiles(bee *hive.Bee) {
	bee.FillDefaultMeta()

	generators, err := ToPkgsGenerators(bee.Languages)
	_if.ErrorExit("generate entry points files", err)

	for _, gen := range *generators {
		err := gen.EntryPointFiles(bee)
		_if.ErrorExit("generate entry points files", err)
	}

	for _, lang := range *generators {
		lang.Protoc(bee)
	}
}

func GenerateClientsFilesFor(bee *hive.Bee) {
	bee.FillDefaultMeta()

	err := ServerLang(bee).ClientsFiles(bee)
	_if.ErrorExit("generate clients files", err)
}

func (code Code) Generate() error {
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
