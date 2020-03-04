package initier

import (
	"fmt"
	"github.com/benka-me/hive/go-pkg/cli/scan"
	"github.com/benka-me/hive/go-pkg/conf"
	"github.com/benka-me/hive/go-pkg/hive"
	"github.com/urfave/cli"
	"os"
)

var HivePath = fmt.Sprintf("%s/hive", os.Getenv("HOME"))
func setAuthor(hive *hive.Hive)  {
	author := conf.GetUsername()
	if author == "" {
		hive.Name = scan.Step(
			"No username found, please login or type a temporary username (not recommended):",
			"required,lte=20,gte=2",
			func (s string) error { return nil } )
	} else {
		hive.Author = author
	}
}

func Hive(c *cli.Context) error {
	hive := hive.InitHiveAskUser()
	setAuthor(hive)


	err := hive.SaveLocal()
	if err != nil {
		return err
	}
	return nil
}
