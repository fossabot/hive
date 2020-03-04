package initier

import (
	"github.com/benka-me/hive/go-pkg/conf"
	"github.com/benka-me/hive/go-pkg/generator"
	"github.com/benka-me/hive/go-pkg/hive"
	"github.com/urfave/cli"
	"os"
)

var gopath = os.Getenv("GOPATH")

func Bee(c *cli.Context) error {
	bee := hive.InitBeeAskUser()

	generator.GenerateEntryPointFiles(bee)

	conf.AddBee(bee.GetNamespaceStr(), bee.Repo)
	err := bee.SaveLocal()
	if err != nil {
		return err
	}
	return err
}
