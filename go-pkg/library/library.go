package library

import (
	"errors"
	"fmt"
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"os"
)

type Cell struct {
	Name string        `bson:"name" json:"name" yaml:"name"`
	PkgName string     `bson:"pkgName" json:"pkgName" yaml:"pkgName"`
	Repo string        `bson:"repo" json:"repo" yaml:"repo"`
	Author string      `bson:"author" json:"author" yaml:"author"`
	AuthorEmail string `bson:"authorEmail" json:"authorEmail" yaml:"authorEmail"`
	Port int           `bson:"port" json:"port" yaml:"port"`
}

type Library struct {
	Hive map[string]Cell
}

var (
	libraryFile string
	library Library
	gopath string
)

func init() {
	gopath = os.Getenv("GOPATH")
	if len(gopath) < 3 {
		panic(errors.New("gopath not specified: " + gopath))
	}
	libraryFile =  fmt.Sprintf("%s/src/github.com/benka-me/hive/data/library.yaml", gopath)

	_, err := os.Stat(libraryFile)
	if os.IsNotExist(err) {
		panic(err)
	}
}

func GetLibrary() (Library, error) {
	dat, err := ioutil.ReadFile(libraryFile)
	if err != nil {
		return library, err
	}

	err = yaml.Unmarshal(dat, &library)
	if err != nil {
		return library, err
	}

	return library, nil
}

func AddCellToLibrary (cell Cell) error {
	lib, err := GetLibrary()
	if err != nil {
		return err
	}

	lib.Hive[cell.Name] = cell

	err = SaveLibrary(lib)
	if err != nil {
		return err
	}
	return nil
}

func SaveLibrary(library Library) error {
	data, err := yaml.Marshal(library)
	if err != nil {
		return err
	}

	err = ioutil.WriteFile(libraryFile, data, 0644)
	if err != nil {
		return err
	}

	return nil
}
