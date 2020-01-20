package cli

import (
	"github.com/benka-me/hive/go-pkg/library"
	"github.com/benka-me/hive/go-pkg/yaml"
	"github.com/urfave/cli"
)


func runInitCell(c *cli.Context) error {
	cell, err := yaml.InitYaml()
	if err != nil {
		return err
	}

	err = library.AddCellToLibrary(cell)
	if err != nil {
		return err
	}


	return nil
}
