package main

import (
	"fmt"
	"github.com/benka-me/hive/go-pkg/cli"
	"os"
)

func main() {
	if err := cli.Root(os.Args[1:]); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
