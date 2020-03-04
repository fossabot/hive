package dive


import (
	_if "github.com/benka-me/hive/go-pkg/cli/if"
	"github.com/benka-me/hive/go-pkg/hive"
	"github.com/benka-me/hive/go-pkg/request"
)

func GetSubDependenciesNamespaces(deps hive.Namespaces) hive.Namespaces {
	allDeps := make(hive.Namespaces, 0)
	for _, d := range deps {
		subDeps := 	Dive(d)
		allDeps = hive.AppendNamespace(allDeps, subDeps...)
	}

	return allDeps
}

func Dive(current hive.Namespace) hive.Namespaces {
	ret := make(hive.Namespaces, 0)

	bee, err := request.GetRemoteBee(current)
	_if.ErrorExit("dive get remote bee", err)
	bee.GetLocal()

	for _, sub := range bee.GetSubDependencies() {
		subNamespaces := Dive(sub)
		//TODO check cycle(if sub containsString current)
		ret = hive.AppendNamespace(ret, subNamespaces...)
	}
	ret = hive.AppendNamespace(ret, current)
	return ret
}

