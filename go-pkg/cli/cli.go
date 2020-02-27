package cli

import (
	"errors"
	"fmt"
	"github.com/benka-me/hive/go-pkg/cli/initier"
	"github.com/benka-me/hive/go-pkg/cli/install"
	"github.com/benka-me/hive/go-pkg/cli/list"
	"github.com/benka-me/hive/go-pkg/cli/privatish"
	"github.com/benka-me/hive/go-pkg/cli/publish"
	"github.com/benka-me/hive/go-pkg/cli/remove"
	"github.com/benka-me/hive/go-pkg/hive"
	"github.com/benka-me/hive/go-pkg/user"
	"github.com/urfave/cli"
	"log"
	"os"
	"sort"
)

var gopath = os.Getenv("GOPATH")
var HivePath = fmt.Sprintf("%s/hive", os.Getenv("HOME"))

func Run()  {
	if _, err := os.Stat(HivePath); os.IsNotExist(err) {
		_ = os.Mkdir(HivePath, 0777) //TODO permission
	}
	app := cli.NewApp()
	app.Name = "Hive"
	app.Usage = "Manage your microservices based application"
	//app.Action = func (c *cli.Context) error {
	//	fmt.Println("app AAction!")
	//	return nil
	//}

	app.Flags = []cli.Flag{
		&cli.StringFlag{
			Name:  "lang, l",
			Value: "english",
			Usage: "Language for the greeting",
		},
		&cli.StringFlag{
			Name:  "config, c",
			Usage: "Load configuration from `FILE`",
		},
	}
	app.Commands = cli.Commands{
		{
			Name:                   "init",
			Usage:                  "init bee or hive",
			Subcommands: cli.Commands{
				{
					Name:   "bee",
					Action: initier.Bee,
					Usage:  "init bee (micro-service)",
				},
				{
					Name:   "hive",
					Action: initier.Hive,
					Usage:  "init hive application",
				},
			},
		},
		{
			Name:                   "push",
			Action: publish.Push,
			Usage:                  "push on hive-bees.com",
		},
		{
			Name:                   "publish",
			Action: publish.Run,
			Usage:                  "publish on hive-and-bees.com",
		},
		{
			Name:                   "privatish",
			Action: privatish.Run,
			Usage:                  "privatish on hive-and-bees.com",
		},
		{
			Name:                   "register",
			Action: user.RunRegister,
			Usage:                  "register to hive-micro-bee.com",
		},
		{
			Name:                   "login",
			Action: user.RunLogin,
			Usage:                  "login",
		},
		{
			Name:                   "whoami",
			Action: user.RunWhoAmI,
			Usage:                  "Don't forget who you are",
		},
		{
			Name:    "protoc",
			Aliases: []string{"proto", "gnr"},
			Action: func(c *cli.Context) error {
				bee, err := hive.GetYamlBeeLocal()
				if err != nil {
					return errors.New("no bee.yaml file founded")
				}
				bee.Protoc()
				return nil
			},
			Usage:   "generate protobuf files",
		},
		{
			Name:    "install",
			Aliases: []string{"i"},
			Action:  install.RunInstall,
			Usage:   "install dependencies",
		},
		{
			Name:    "list",
			Aliases: []string{"ls"},
			Action:  list.RunList,
			Usage:   "list dependencies",
		},
		{
			Name:    "remove",
			Aliases: []string{"rm", "delete", "del"},
			Action:  remove.RunRemove,
			Usage:   "remove dependencies",
		},
	}

	sort.Sort(cli.FlagsByName(app.Flags))
	sort.Sort(cli.CommandsByName(app.Commands))

	err := app.Run(os.Args)
	if err != nil {
		log.Println(err)
		os.Exit(0)
	}
}
