package install

import (
	"errors"
	"fmt"
	"github.com/benka-me/hive/go-pkg/hive"
	"github.com/benka-me/hive/go-pkg/request"
	"github.com/urfave/cli"
	"os"
)

func ifErrorExit (err error) {
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func RunInstall(c *cli.Context) error {
	if len(os.Args) < 3 {
		return errors.New("bad argument")
	}
	urls := hive.ArrToURLs(os.Args[2:])

	b, errBee := hive.GetYamlBeeLocal()
	h, errHive := hive.GetYamlHiveLocal()

	if errBee == nil && errHive != nil {
		err := installBee(b, urls)
		if err != nil {
			return err
		}
	} else if errBee != nil && errHive == nil {
		err := installHive(h, urls)
		if err != nil {
			return err
		}
	} else if errBee != nil && errHive != nil {
		return errors.New("no hive.yaml or bee.yaml detected")
	} else if errBee == nil && errHive == nil {
		return errors.New("both hive.yaml and bee.yaml detected, please remove one")
	}
	return nil
}

func installHive (target hive.Hive, urls hive.URLs) error {
	return nil
}


func dive(self hive.URL) hive.URLs {
	ret := make(hive.URLs, 0)
	bee, err := request.GetBee(self)
	ifErrorExit(err)
	for _, urlSub := range bee.DepURLs() {
		urlsSub := dive(urlSub)
		//TODO check cycle(if urlSub contains himself -> self)
		ret = hive.AppendURL(ret, urlsSub...)
	}
	ret = hive.AppendURL(ret, self)
	return ret
}

func installBee(target hive.Bee, urls hive.URLs) error {
	bees, err := request.GetBees(urls)
	ifErrorExit(err)

	_ = target.ConcatDeps(bees).SaveYaml()

	deps := make(hive.URLs, 0)
	deps = hive.AppendURL(deps, urls...)

	fmt.Println(deps)
	allDeps := make(hive.URLs, 0)
	for _, d := range deps {
		subs := dive(d)
		allDeps = hive.AppendURL(allDeps, subs...)
	}

	fmt.Println(allDeps)

	return nil
}

