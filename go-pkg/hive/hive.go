package hive

import (
	"fmt"
	"github.com/benka-me/hive/go-pkg/cli/scan"
	"github.com/go-playground/validator"
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"os"
	"strings"
)

var gopath = os.Getenv("GOPATH")

func GetHive() (Hive, error) {
	hive := Hive{}
	dat, err := ioutil.ReadFile("./hive.yaml")

	if err != nil {
		return hive, err
	}

	err = yaml.Unmarshal(dat, &hive)
	if err != nil {
		return hive, err
	}

	return hive, nil
}

func (hive *Hive) SaveYaml() error {
	data, err := yaml.Marshal(hive)
	if err != nil {
		return err
	}

	path := fmt.Sprintf("%s/src/%s/hive.yaml", gopath, hive.Repo)
	err = ioutil.WriteFile(path, data, 0644)
	if err != nil {
		return err
	}

	return nil
}

func InitHiveAskUser() *Hive {
	hive := &Hive{}
	scan.V = validator.New()

	hive.Name = strings.ToLower(scan.Step(
		"Name of the new hive micro-service based application",
		"required,lte=20,gte=3",
		scan.IsAlphanumDash))

	hive.PkgName = strings.ToLower(scan.Step(
		"Package name (2 - 7 chars long, shorter is better) for packages and types building",
		"required,lte=7,gte=2",
		scan.IsAlphaNum))

	hive.Repo = strings.Replace(scan.Step(
		"Git repository",
		"required,gte=5",
		func(s string) error {return nil}), " ", "", -1)

	//portValidator := portValidator(hive.Public)
	hive.FillMeta()

	return hive
}

func (hive *Hive) FillMeta() {
	hive.PkgNameCamel = strings.Title(scan.KebabToCamelCase(hive.PkgName))
}
