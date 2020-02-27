package hive 

import (
	"fmt"
	"github.com/benka-me/hive/go-pkg/cli/scan"
	"github.com/go-playground/validator"
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"strconv"
	"strings"
)

func (bee *Bee) Protoc() {
	bee.FillMeta()
	bee.SaveYaml()
	langs := bee.Languages.ToGenerators()

	for _, lang := range *langs {
		lang.Protoc()
	}
}

func (bee *Bee) GenerateFiles() error {
	bee.FillMeta()
	langs := bee.Languages.ToGenerators()

	for _, l := range *langs {
		err := l.PublicFiles(bee)
		if err != nil {
			return err
		}
	}

	for _, lang := range *langs {
		lang.Protoc()
	}
	return nil
}

func InitBeeAskUser() *Bee {
	bee := &Bee{}
	scan.V = validator.New()

	bee.Name = strings.ToLower(scan.Step(
		"Name of the new bee microservice ",
		"required,lte=20,gte=3",
		scan.IsAlphanumDash))

	bee.PkgName = strings.ToLower(scan.Step(
		"Package name (2 - 7 chars long, shorter is better) for packages and types building",
		"required,lte=7,gte=2",
		scan.IsAlphaNum))

	bee.Repo = strings.Replace(scan.Step(
		"Git repository",
		"required,gte=5",
		func(s string) error {return nil}), " ", "", -1)

	//BoolStr := Step(
	//	"do you want it public? (you can change later) y/n",
	//	"required,lte=1,gte=1",
	//	func(s string) error {
	//		if !strings.ContainsAny(s, "yn") {
	//			return errors.New("y or n only")
	//		}
	//		return nil
	//	})
	//bee.Public = strings.Compare(BoolStr, "y") == 0

	//portValidator := portValidator(bee.Public)
	pInt, _ := strconv.Atoi(scan.Step(
		"Port",
		"required,number",
		func(s string) error {return nil}))
	bee.Port= int32(pInt)

	bee.FillMeta()

	return bee
}

func GetYamlBeeLocal() (Bee, error) {
	bee := Bee{}
	dat, err := ioutil.ReadFile("./bee.yaml")

	if err != nil {
		return bee, err
	}

	err = yaml.Unmarshal(dat, &bee)
	if err != nil {
		return bee, err
	}

	return bee, nil
}

func (bee *Bee) SaveYaml() error {
	data, err := yaml.Marshal(bee)
	if err != nil {
		return err
	}

	path := fmt.Sprintf("%s/src/%s/bee.yaml", gopath, bee.Repo)
	err = ioutil.WriteFile(path, data, 0644)
	if err != nil {
		return err
	}

	return nil
}

func (bee *Bee) FillMeta() {
	bee.PkgNameCamel = strings.Title(scan.KebabToCamelCase(bee.PkgName))
	bee.Languages =  &Languages{
		Go: &Go{Setup: &LanguageSetup{
			Active: true,
			Repo:  bee.Repo,
			Files: []string{
				fmt.Sprintf("%s.proto", bee.PkgName),
				fmt.Sprintf("rpc-%s.proto", bee.PkgName),
			},
			PkgName: bee.PkgName,
		}},
		Javascript: &Javascript{Setup: &LanguageSetup{
			Active: true,
			Repo:   bee.Repo,
			Files: []string{
				fmt.Sprintf("%s.proto", bee.PkgName),
				fmt.Sprintf("rpc-%s.proto", bee.PkgName),
			},
			PkgName: bee.PkgName,
		}},
	}
}
