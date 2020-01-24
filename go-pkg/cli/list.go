package cli

import (
	"fmt"
	"github.com/benka-me/hive/go-pkg/yaml"
	"github.com/urfave/cli"
)

func runList(c *cli.Context) error {
	config, err := yaml.GetHiveConfig()
	if err != nil {
		return err
	}
	for _, dep := range config.Dependencies {
		fmt.Println(dep)
	}

	return nil
}
