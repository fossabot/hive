package cli

import (
	"errors"
	"fmt"
	"github.com/benka-me/hive/go-pkg/library"
	"github.com/urfave/cli"
)

func runInitCell(c *cli.Context) error {
	cell := library.Cell{
		Name:    "",
		Repo:    "",
		PkgName: "",
	}
	var tmp string

	err := errors.New("Todo")
	for err != nil {
		fmt.Println("Name of the new cell microservice: ")
		if _, err = fmt.Scanf("%s", &tmp); err != nil {
			fmt.Printf("%s\n", err)
		} else if len(tmp) < 2 {
			fmt.Println("Cell name must be at least 2 characters long")
		} else if len(tmp) > 20 {
			fmt.Println("Cell name must be less than 20 characters long")
		} else {
			cell.Name = tmp
		}
	}

	err = errors.New("Todo")
	for err != nil {
		fmt.Println("Package name (2 - 4 chars long) for packages and types building")
		if _, err = fmt.Scanf("%s", &tmp); err != nil {
			fmt.Printf("%s\n", err)
		} else if len(tmp) < 2 {
			fmt.Println("Pkg name too short: ", len(tmp))
		} else if len(tmp) > 4 {
			fmt.Println("Pkg name too long: ", len(tmp))
		} else {
			cell.PkgName = tmp
		}
	}

	err = errors.New("Todo")
	for err != nil {
		fmt.Println("Repository: ")
		if _, err = fmt.Scanf("%s", &tmp); err != nil {
			fmt.Printf("%s\n", err)
		} else if len(tmp) < 5 {
			fmt.Println("adresse seems to long: ", len(tmp))
		} else if len(tmp) > 80 {
			fmt.Println("adresse seems to long: ", len(tmp))
		} else {
			cell.PkgName = tmp
		}
	}

	fmt.Println(cell)
	return nil
}
