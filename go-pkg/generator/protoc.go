package generator

import (
	"bytes"
	"fmt"
	_if "github.com/benka-me/hive/go-pkg/cli/if"
	"github.com/benka-me/hive/go-pkg/hive"
	"os/exec"
)

func Protoc (bee hive.Bee) {
	lgs, err := GetLangs(bee.Languages)
	_if.ErrorExit("protoc to pkgs generators", err)

	for _, lg := range *lgs {
		lg.Protoc(&bee)
	}
}

func runProtocCommand(args []string) {
	cmd := exec.Command("protoc", args...)
	var out bytes.Buffer
	var errs bytes.Buffer
	cmd.Stdout = &out
	cmd.Stderr= &errs
	fmt.Println(cmd.Args)
	err := cmd.Run()
	if err != nil {
		fmt.Println(errs.String())
	}
	fmt.Printf("%s\n", out.String())
}
