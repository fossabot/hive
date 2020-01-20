package cli

import (
	"bytes"
	"fmt"
	"github.com/benka-me/hive/go-pkg/library"
	"github.com/benka-me/hive/go-pkg/yaml"
	"github.com/urfave/cli"
	"os"
	"os/exec"
)

func goGet(repo string) error {
	cmd := exec.Command("go", "get", "-d", "-v", fmt.Sprintf("%s/cmd", repo))
	var out bytes.Buffer
	var errs bytes.Buffer
	cmd.Stdout = &out
	cmd.Stderr= &errs
	fmt.Println(cmd.Args)
	err := cmd.Run()
	if err != nil {
		fmt.Println(errs.String())
	}
	fmt.Println(out.String())

	return err
}

func installAll(lib library.Library, config yaml.Config) error {
	fmt.Println("install all")
	for _, dep := range config.Dependencies {
		_ = goGet(lib.Hive[dep].Repo)
	}
	return nil
}

func installDep (lib library.Library, config *yaml.Config, dep string) error {
	err := goGet(lib.Hive[dep].Repo)
	if err != nil {
		return err
	}
	config.Dependencies = append(config.Dependencies, dep)

	return nil
}

func installUnits (lib library.Library, config *yaml.Config) error {
	for _, dep := range os.Args[2:] {
		installDep(lib, config, dep)
	}
	return nil
}

func runInstall(c *cli.Context) error {
	config, err := yaml.GetConfig()
	if err != nil {
		return err
	}

	lib, err := library.GetLibrary()
	if len(os.Args) == 2 {
		err := installAll(lib, config)
		if err != nil {
			return err
		}
	} else {
		err := installUnits(lib, &config)
		if err != nil {
			return err
		}

		err = yaml.SaveConfig(config)
		if err != nil {
			return err
		}
	}



	return nil
}
