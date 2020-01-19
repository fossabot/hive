package yaml

import (
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"os"
)

type Config struct {
	Name string
	Repo string
	Dependencies []string
}

func GetConfig() (Config, error) {
	config := Config{}
	file :=  "./hive.yaml"

	_, err := os.Stat(file)
	if os.IsNotExist(err) {
		return config, err
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