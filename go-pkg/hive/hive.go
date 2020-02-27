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

//func GetHiveMust() Hive {
//	h, err := GetYamlHiveLocal()
//	if err != nil {
//		fmt.Println("===> hive.yaml not found ", err)
//		os.Exit(1)
//	}
//	return h
//}
//
func GetHivePath(name string) (Hive, error) {
	hive := Hive{}
	dat, err := ioutil.ReadFile(fmt.Sprintf("%s/%s/hive.yaml", HivePath, name))

	if err != nil {
		return hive, err
	}

	err = yaml.Unmarshal(dat, &hive)
	if err != nil {
		return hive, err
	}

	return hive, nil
}
func GetYamlHiveLocal() (Hive, error) {
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
	if _, err := os.Stat(HivePath); os.IsNotExist(err) {
		_ = os.Mkdir(HivePath + "/" + hive.Name, 0777) //TODO permission
	}

	path := fmt.Sprintf("%s/%s/hive.yaml", HivePath, hive.Name)
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
