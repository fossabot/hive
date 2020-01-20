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
			Name:                   "install",
			Aliases:                []string{"i"},
			Action: runInstall,
			Usage:                  "install dependencies",
			BashComplete: func(c *cli.Context) {},
			UsageText:              "",
			Description:            "",
			ArgsUsage:              "Argss Usaaaaage",
			Category:               "",
			Before:                 nil,
			After:                  nil,
			OnUsageError:           nil,
			Flags:                  []cli.Flag{},
			SkipFlagParsing:        true,
			HideHelp:               false,
			Hidden:                 false,
			UseShortOptionHandling: false,
			HelpName:               "",
			CustomHelpTemplate:     "",
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
