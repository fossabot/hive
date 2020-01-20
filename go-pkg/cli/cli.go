package cli

import (
	"fmt"
	"github.com/urfave/cli"
	"log"
	"os"
	"sort"
)

//type Runner interface {
//	Init([]string) error
//	Run() error
//	Name() string
//}
//
//func Root(args []string) error {
//	if len(args) < 1 {
//		return errors.New("You must pass a sub-command")
//	}
//
//	cmds := []Runner{
//		NewGreetCommand(),
//		NewProtocCommand(),
//		NewInitCommand(),
//		NewInstallCommand(),
//	}
//
//	subcommand := os.Args[1]
//
//	for _, cmd := range cmds {
//		if cmd.Name() == subcommand {
//			cmd.Init(os.Args[2:])
//			return cmd.Run()
//		}
//	}
//
//	return fmt.Errorf("Unknown subcommand: %s", subcommand)
//}

func Run()  {
	app := cli.NewApp()
	app.Name = "Hive"
	app.Usage = "Manage your microservices"
	app.Action = func (c *cli.Context) error {
		fmt.Println("app AAction!")
		return nil
	}

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
					Name:                   "cell",
					Action: runInitCell,
					Usage:                  "init cell microservice",
				},
				{
					Name:                   "hive",
					Action: runInitHive,
					Usage:                  "init hive application",
				},
			},
		},
		{
			Name:                   "install",
			Aliases:                []string{"i"},
			Action: runInstall,
			Usage:                  "install dependencies",
		},
		{
			Name:                   "list",
			Aliases:                []string{"ls"},
			Action: runList,
			Usage:                  "list dependencies",
		},
		{
			Name:                   "remove",
			Aliases:                []string{"rm", "delete", "del"},
			Action: runRemove,
			Usage:                  "remove dependencies",
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
