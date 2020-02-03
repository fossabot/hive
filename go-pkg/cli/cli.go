package cli

import (
	"errors"
	"github.com/benka-me/hive-server-core/go-pkg/core"
	"github.com/benka-me/hive/go-pkg/cli/initier"
	"github.com/benka-me/hive/go-pkg/cli/install"
	"github.com/benka-me/hive/go-pkg/cli/list"
	"github.com/benka-me/hive/go-pkg/cli/remove"
	"github.com/benka-me/hive/go-pkg/user"
	"github.com/urfave/cli"
	"log"
	"os"
	"sort"
)

var gopath = os.Getenv("GOPATH")

func Run()  {
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
			Usage:                  "init cell or hive",
			Subcommands: cli.Commands{
				{
					Name:   "cell",
					Action: initier.Cell,
					Usage:  "init cell (microservice)",
				},
				{
					Name:   "hive",
					Action: initier.Hive,
					Usage:  "init hive application",
				},
			},
		},
		{
			Name:                   "register",
			Action: user.RunRegister,
			Usage:                  "register to hivemicrocell.com",
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
				cell, err := core.GetCell()
				if err != nil {
					return errors.New("no cell.yaml file founded")
				}
				cell.Protoc()
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
