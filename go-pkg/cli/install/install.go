package install

import (
	"errors"
	"github.com/benka-me/hive/go-pkg/hive"
	"github.com/urfave/cli"
	"os"
)


func RunInstall(c *cli.Context) error {
	var h Hive
	var b Bee
	var errHive error
	var errBee error

	if len(os.Args) < 3 {
		return errors.New("bad argument")
	}

	b.Bee, errBee = hive.GetYamlBeeLocal()
	h.Hive, errHive = hive.GetYamlHiveLocal()

	if errBee == nil && errHive != nil {
		err := installBee(b, os.Args[2:])
		if err != nil {
			return err
		}
	} else if errBee != nil && errHive == nil {
		err := installHive(h)
		if err != nil {
			return err
		}
	} else if errBee != nil && errHive != nil {
		return errors.New("no hive.yaml or bee.yaml detected")
	} else if errBee == nil && errHive == nil {
		return errors.New("both hive.yaml and bee.yaml detected, please remove one")
	}
	return nil
}
