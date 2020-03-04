package hive

import "github.com/benka-me/hive/go-pkg/git"

func (dep Namespace) ConcatConsumer(cons Namespace) {
	dep.GetLocalHive()
}
func (dep Namespace) ConcatDependency(cons Namespace) {

}

func (parent *Bee) ConcatSubDependenciesFrom(childrens []*Bee) *Bee {
	for _, child := range childrens {
		parent.ConsumersConcatTo(child)
	}

	childrensNamespaces := BeesToNamespacesFrom(childrens)
	for _, childs := range childrensNamespaces {
		if !containsString(parent.Deps, childs.NamespaceStr()) {
			parent.Deps = append(parent.Deps, childs.NamespaceStr())
		}
	}

	return parent
}


func containsNamespace(namespaces Namespaces, namespace Namespace) bool {
	for _, a := range namespaces {
		if a == namespace {
			return true
		}
	}
	return false
}

func containsString(arr []string, str string) bool {
	for _, a := range arr {
		if a == str {
			return true
		}
	}
	return false
}

func (parent *Bee) ConsumersConcatTo(child *Bee) {
	child.GetLocal()
	for _, c := range parent.Cons {
		if !containsString(child.Cons, c) {
			child.Cons = append(child.Deps, c)
		}
	}
	child.SaveLocal()
}

func (h *Hive) ConcatConsumerToLocalBee(b *Bee) {
	b.GetLocal()
	if !containsString(b.Cons, h.NamespaceStr()) {
		b.Cons = append(b.Deps, h.NamespaceStr())
	}
	b.SaveLocal()
}

func (cons *Hive) ConcatDependencyFromNamespace(namespace Namespace) {
	bee := GetLocalBeeFromStringNamespace(namespace)
	_, _ = git.Clone(bee.Repo)
	cons.ConcatDependencyFromBee(bee)
}

func (cons *Hive) ConcatDependencyFromBee(dep *Bee) {
	cons.ConcatConsumerToLocalBee(dep)
	if _, ok := cons.Deps[dep.GetNamespaceStr()]; !ok {
		cons.Deps[dep.GetNamespaceStr()] = &Dep{
			Port:         dep.Port,
			Dev:          "localhost",
			Prod:         "127.0.0.1",
			PkgName:      dep.PkgName,
			PkgNameCamel: dep.PkgNameCamel,
		}
	}
}
