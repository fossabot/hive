package initier

import (
	"fmt"
	"github.com/benka-me/hive/go-pkg/hive"
	"github.com/urfave/cli"
	"os"
)

var HivePath = fmt.Sprintf("%s/hive", os.Getenv("HOME"))
func Hive(c *cli.Context) error {
	hive := hive.InitHiveAskUser()
	if _, err := os.Stat(HivePath); os.IsNotExist(err) {
		_ = os.Mkdir(HivePath + "/" + hive.Name, 0777) //TODO permission
	}

	err := hive.SaveYaml()
	if err != nil {
		return err
	}
	return nil
}
