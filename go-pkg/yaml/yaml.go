package yaml

import (
	"errors"
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
const file = "./hive.yaml"

func GetConfig() (Config, error) {
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

func SaveConfig(config Config) error {
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

	err = ioutil.WriteFile("./cell.yaml", data, 0644)
	if err != nil {
		return err
	}

	return nil
}
