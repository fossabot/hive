package conf

import (
	"errors"
	"fmt"
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"os"
)

type Conf struct {
	Username string
	Token string
	Hives map[string]string
}


var file = fmt.Sprintf("%s/.hive_conf.yaml", os.Getenv("HOME"))

func (c *Conf) Write() error {
	b, err := yaml.Marshal(c)
	if err != nil {
		return err
	}

	return ioutil.WriteFile(file, b, 0644)
}

func Parse() Conf {
	c := Conf{}
	yamlFile, _ := ioutil.ReadFile(file)
	_ = yaml.Unmarshal(yamlFile, &c)

	return c
}

func (c *Conf) Parse() error {
	if c == nil {
		c = &Conf{}
	}
	yamlFile, err := ioutil.ReadFile(file)
	if err != nil {
		return errors.New("cannot parse or write file on " + file)
	}

	err = yaml.Unmarshal(yamlFile, c)
	if err != nil {
		return err
	}

	return nil
}

