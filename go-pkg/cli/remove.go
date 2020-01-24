package cli

import (
	"errors"
	"github.com/benka-me/hive/go-pkg/yaml"
	"github.com/urfave/cli"
	"os"
	"strings"
)

func runRemove(c *cli.Context) error {
	config, err := yaml.GetHiveConfig()
	if err != nil {
		return err
	}

	if len(os.Args) == 2 {
		return errors.New("please indicate cell to remove")
	} else {
		for _, del := range os.Args[2:] {
			for i, dep := range config.Dependencies {
				if 0 == strings.Compare(del, dep) {
					config.Dependencies = append(config.Dependencies[:i], config.Dependencies[i+1:]...)
				}
			}
		}

		err = yaml.SaveConfig(config, "./hive.yaml")
		if err != nil {
			return err
		}
	}



	return nil
}
