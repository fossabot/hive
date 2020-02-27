package conf

import (
	"errors"
	"fmt"
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"log"
	"os"
)

type Conf struct {
	Username string
	Token string
	Hives map[string]string
}


var file = fmt.Sprintf("%s/.hive_conf.yaml", os.Getenv("HOME"))
var HivePath = fmt.Sprintf("%s/hive", os.Getenv("HOME"))

func Hives(fn func (s string)) {
	dirs, err := ioutil.ReadDir(HivePath)
	if err != nil {
		log.Fatal(err)
	}

	for _, f := range dirs {
		fn(f.Name())
	}
}

func Token() string {
	c := Parse()
	return c.Token
}

func (c *Conf) Write() error {
	b, err := yaml.Marshal(c)
	if err != nil {
		return err
	}

	return ioutil.WriteFile(file, b, 0644)
}

func Parse() Conf {
	c := Conf{}
	yamlFile, err := ioutil.ReadFile(file)
	if err != nil {
		fmt.Println("parse " + file, err)
		os.Exit(1)
	}
	err = yaml.Unmarshal(yamlFile, &c)
	if err != nil {
		fmt.Println("unmarshall " + file, err)
		os.Exit(1)
	}

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

