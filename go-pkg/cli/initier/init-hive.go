package initier

import (
	"github.com/benka-me/hive/go-pkg/hive"
	"github.com/urfave/cli"
)

func Hive(c *cli.Context) error {
	hive := hive.InitHiveAskUser()

	err := hive.SaveYaml()
	if err != nil {
		return err
	}
	return nil
}
