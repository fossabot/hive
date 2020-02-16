package initier

import (
	"github.com/benka-me/hive/go-pkg/hive"
	"github.com/urfave/cli"
	"os"
)

var gopath = os.Getenv("GOPATH")

func Bee(c *cli.Context) error {
	bee := hive.InitBeeAskUser()

	err := bee.GenerateFiles()
	if err != nil {
		return err
	}

	err = bee.SaveYaml()
	if err != nil {
		return err
	}
	return err
}
