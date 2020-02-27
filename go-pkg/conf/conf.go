package conf

import (
	"fmt"
	"github.com/benka-me/cell-user/go-pkg/user"
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

func GetUsername() string {
	c := ParseYaml()
	return c.Username
}

func UserToken() *user.Token {
	c := ParseYaml()
	return &user.Token{Val:c.Token}
}

func Token() string {
	c := ParseYaml()
	return c.Token
}

func (c *Conf) SaveYaml() error {
	b, err := yaml.Marshal(c)
	if err != nil {
		return err
	}

	return ioutil.WriteFile(file, b, 0644)
}

func ParseYaml() Conf {
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

func (c *Conf) ParseYaml() error {
	ParseYaml()
	return nil
}

