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

var HivePath = fmt.Sprintf("%s/hive", os.Getenv("HOME"))

//func (h *Hive) Dep

type Hives []*Hive

func GetLocalHiveFromString(namespace string) (Hive, error) {
	author, name := Explode(namespace)
	hive := Hive{}
	dat, err := ioutil.ReadFile(fmt.Sprintf("%s/%s/%s/hive.yaml", HivePath, author, name))
	if err != nil {
		return hive, err
	}

	err = yaml.Unmarshal(dat, &hive)
	if err != nil {
		return hive, err
	}

	return hive, nil
}

func GetLocalHiveCurrentDir() (Hive, error) {
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

func (hive *Hive) SaveLocal() error {
	data, err := yaml.Marshal(hive)
	if err != nil {
		return err
	}

	err = os.MkdirAll(fmt.Sprintf("%s/%s/%s", HivePath, hive.Author, hive.Name), 0777) //TODO perm
	if err != nil {
		return err
	}

	path := fmt.Sprintf("%s/%s/%s/hive.yaml", HivePath, hive.Author, hive.Name)
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

	hive.FillMeta()

	return hive
}

func (hive *Hive) FillMeta() {
	hive.PkgNameCamel = strings.Title(scan.KebabToCamelCase(hive.PkgName))
}
