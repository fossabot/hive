package yaml

import (
	"errors"
	"fmt"
	"github.com/benka-me/hive/go-pkg/library"
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"os"
)

type Config struct {
	Name string
	Repo string
	Dependencies []string
}
var gopath = os.Getenv("GOPATH")

func GetCellConfig() (library.Cell, error) {
	const file = "./cell.yaml"
	config := library.Cell{}

	_, err := os.Stat(file)
	if os.IsNotExist(err) {
		return config, errors.New("no hive application detected, please create a new one or move to root directory of a hive application")
	}

	dat, err := ioutil.ReadFile(file)
	if err != nil {
		return config, err
	}


	err = yaml.Unmarshal(dat, &config)
	if err != nil {
		return config, err
	}


	return config, nil
}

func GetHiveConfig() (Config, error) {
	const file = "./hive.yaml"
	config := Config{}

	_, err := os.Stat(file)
	if os.IsNotExist(err) {
		return config, errors.New("no hive application detected, please create a new one or move to root directory of a hive application")
	}

	dat, err := ioutil.ReadFile(file)
	if err != nil {
		return config, err
	}


	err = yaml.Unmarshal(dat, &config)
	if err != nil {
		return config, err
	}


	return config, nil
}

func SaveConfig(config Config, file string) error {
	data, err := yaml.Marshal(config)
	if err != nil {
		return err
	}

	err = ioutil.WriteFile(file, data, 0644)
	if err != nil {
		return err
	}
	return nil
}

func GetCell() (library.Cell, error) {
	cell := library.Cell{}
	dat, err := ioutil.ReadFile("./cell.yaml")
	if err != nil {
		return cell, err
	}

	err = yaml.Unmarshal(dat, &cell)
	if err != nil {
		return cell, err
	}

	return cell, nil
}

func SaveCell(cell library.Cell) error {
	data, err := yaml.Marshal(cell)
	if err != nil {
		return err
	}

	path := fmt.Sprintf("%s/src/%s/cell.yaml", gopath, cell.Repo)
	fmt.Println(path)
	err = ioutil.WriteFile(path, data, 0644)
	if err != nil {
		return err
	}

	err = library.AddCellToLibrary(cell)
	if err != nil {
		return err
	}

	return nil
}
