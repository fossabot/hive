package initier

import (
	"github.com/benka-me/hive-server-core/go-pkg/core"
	"github.com/urfave/cli"
	"os"
)

var gopath = os.Getenv("GOPATH")

func Cell(c *cli.Context) error {
	cell := core.InitAskUser()

	err := cell.GenerateFiles()
	if err != nil {
		return err
	}

	err = cell.SaveYaml()
	if err != nil {
		return err
	}
	return err
}
