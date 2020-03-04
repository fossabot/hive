package install

import (
	"errors"
	_if "github.com/benka-me/hive/go-pkg/cli/if"
	"github.com/benka-me/hive/go-pkg/generator"
	"github.com/benka-me/hive/go-pkg/git"
	"github.com/benka-me/hive/go-pkg/hive"
	"github.com/benka-me/hive/go-pkg/hive/dive"
	"github.com/benka-me/hive/go-pkg/request"
	"github.com/urfave/cli"
	"os"
)

func RunInstall(c *cli.Context) error {
	var depMode = false
	if len(os.Args) < 3 {
		depMode = true
	}
	urls := hive.ArrayToNamespaces(os.Args[2:])

	b, errBee := hive.GetLocalBeeCurrentDir()
	h, errHive := hive.GetLocalHiveCurrentDir()

	if errBee == nil && errHive != nil {
		if depMode {
			urls = b.GetSubDependencies()
		}
		err := installFromBee(b, urls)
		if err != nil {return err}
	} else if errBee != nil && errHive == nil {
		if depMode {
			urls = h.GetDependencies()
		}
		err := installFromHive(h, urls)
		if err != nil {return err}
	} else if errBee != nil && errHive != nil {
		return errors.New("no hive.yaml or bee.yaml detected")
	} else if errBee == nil && errHive == nil {
		return errors.New("both hive.yaml and bee.yaml detected, please remove one")
	}
	return nil
}

func installFromHive(target hive.Hive, requiredNamespaces hive.Namespaces) error {
	if target.Deps == nil {
		target.Deps = make(map[string]*hive.Dep)
	}

	allDependenciesNamespace := dive.GetSubDependenciesNamespaces(requiredNamespaces)

	allDependencies, err := request.GetRemoteBees(allDependenciesNamespace)
	_if.ErrorExit("get remote bees", err)

	for _, dep := range allDependencies {
		target.ConcatDependencyFromBee(dep)
		_, _ = git.Clone(dep.Repo)
	}

	return target.SaveLocal()
}

func installFromBee(target hive.Bee, requiredNamespaces hive.Namespaces) error {
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
