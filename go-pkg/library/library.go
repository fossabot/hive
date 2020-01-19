package library

import (
	"errors"
	"fmt"
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"os"
)

type Cell struct {
	Name string
	Repo string
	PkgName string `yaml:"pkgName"`
}

type Library struct {
	Hive map[string]Cell
}

func GetLibray() (Library, error) {
	library := Library{}
	gopath := os.Getenv("GOPATH")
	if len(gopath) < 3 {
		return library, errors.New("gopath not specified: " + gopath)
	}
	file :=  fmt.Sprintf("%s/src/github.com/benka-me/hive/data/library.yaml", gopath)

	_, err := os.Stat(file)
	if os.IsNotExist(err) {
		return library, err
	}

	dat, err := ioutil.ReadFile(file)
	if err != nil {
		return library, err
	}

	err = yaml.Unmarshal(dat, &library)
	if err != nil {
		return library, err
	}


	return library, nil
}
