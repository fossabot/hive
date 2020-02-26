package install

import (
	"fmt"
	"github.com/benka-me/hive/go-pkg/hive"
	"github.com/urfave/cli"
)

type Hive struct {
	hive.Hive
}

func RunInstall(c *cli.Context) error {
	hive := Hive{
		hive.GetHiveMust(),
	}
	fmt.Println(hive)

	return nil
}
