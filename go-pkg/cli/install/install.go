package install

import (
	"errors"
	_if "github.com/benka-me/hive/go-pkg/cli/if"
	"github.com/benka-me/hive/go-pkg/generator"
	"github.com/benka-me/hive/go-pkg/hive"
	"github.com/benka-me/hive/go-pkg/hive/dive"
	"github.com/benka-me/hive/go-pkg/request"
	"github.com/urfave/cli"
	"os"
)

func RunInstall(c *cli.Context) error {
	if len(os.Args) < 3 {
		return errors.New("bad argument")
	}
	urls := hive.ArrayToNamespaces(os.Args[2:])

	b, errBee := hive.GetLocalBeeCurrentDir()
	h, errHive := hive.GetLocalHiveCurrentDir()

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

func installHive (target hive.Hive, requiredNamespaces hive.Namespaces) error {
	if target.Deps == nil {
		target.Deps = make(map[string]*hive.Dep)
	}

	allDependenciesNamespace := dive.GetSubDependenciesNamespaces(requiredNamespaces)

	allDependencies, err := request.GetRemoteBees(allDependenciesNamespace)
	_if.ErrorExit("get remote bees", err)

	for _, dep := range allDependencies {
		target.ConcatDependencyFromBee(dep)
	}

	return target.SaveLocal()
}

func installBee(target hive.Bee, requiredNamespaces hive.Namespaces) error {
	required, err := request.GetRemoteBees(requiredNamespaces)
	_if.ErrorExit("get remote bees", err)

	_ = target.ConcatSubDependenciesFrom(required).SaveLocal()
	generator.GenerateClientsFilesFor(&target)

	allDependenciesNamespace := dive.GetSubDependenciesNamespaces(requiredNamespaces)

	for _, consumerNamespace := range target.Cons {
		consumer, err := hive.GetLocalHiveFromString(consumerNamespace)
		if err != nil {
			return err
		}
		for _, depNamespace := range allDependenciesNamespace {
			consumer.ConcatDependencyFromNamespace(depNamespace)
		}
		consumer.SaveLocal()
	}
	return nil
}
