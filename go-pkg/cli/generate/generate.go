package generate

import (
	_if "github.com/benka-me/hive/go-pkg/cli/if"
	"github.com/benka-me/hive/go-pkg/cli/scan"
	"github.com/benka-me/hive/go-pkg/generator"
	"github.com/benka-me/hive/go-pkg/hive"
	"github.com/urfave/cli"
)

func GenerateEntryPointFiles(c *cli.Context) error {
	bee, err := hive.GetLocalBeeCurrentDir()
	_if.ErrorExit("generate entry point files", err)


	generator.GenerateEntryPointFiles(&bee)
	return nil
}

func GenerateClientsFiles(c *cli.Context) error {
	bee, err := hive.GetLocalBeeCurrentDir()
	_if.ErrorExit("generate clients files", err)

	generator.GenerateClientsFilesFor(&bee)
	return nil
}

func All(c *cli.Context) error {
	if !scan.StepBool("this will rewrite all generated files are you sure? ") {return nil}

	bee, err := hive.GetLocalBeeCurrentDir()
	_if.ErrorExit("generate all", err)

	generator.GenerateEntryPointFiles(&bee)
	generator.GenerateClientsFilesFor(&bee)
	return nil
}
