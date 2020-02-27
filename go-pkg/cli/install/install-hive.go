package install

import (
	"context"
	"fmt"
	"github.com/benka-me/cell-user/go-pkg/user"
	rpc_core "github.com/benka-me/hive-server-core/go-pkg/rpc-core"
	"github.com/benka-me/hive/go-pkg/conf"
	"github.com/benka-me/hive/go-pkg/hive"
	"github.com/imdario/mergo"
	"google.golang.org/grpc"
	"log"
	"os"
)

type Hive struct {
	hive.Hive
}

func installHive (h Hive) error {
	//if h.Deps == nil {
	//	h.Deps = make(map[string]hive.Dep)
	//}
	//core, err := rpc_core.ConnectThroughApi("localhost", grpc.WithInsecure())
	//if err != nil {
	//	return err
	//}

	//res, err := core.GetBees(context.TODO(), &hive.BeesReq{
	//	BeeNames: os.Args[2:],
	//	Token: &user.Token{Val: conf.Token()},
	//})
	//if err != nil {
	//	return err
	//}

	deps := mergeRecursif(os.Args[2:])
	fmt.Println("=======> ", deps)


	//err = h.SaveYaml()
	//if err != nil {
	//	return err
	//}
	//
	//if res.StatusMessage != "" {
	//	fmt.Println(res.StatusMessage)
	//}
	return nil
}

func installBee(b Bee, names []string) error {
	deps := mergeRecursif(os.Args[2:])
	fmt.Println("=======> ", deps)
	return nil
}

func getBees(names []string) (*hive.Bees, error) {
	core, err := rpc_core.ConnectThroughApi("localhost", grpc.WithInsecure())
	if err != nil {
		return nil, err
	}

	return core.GetBees(context.TODO(), &hive.BeesReq{
		BeeNames: names,
		Token: &user.Token{Val: conf.Token()},
	})
}

//func (h *Hive) AddDeps (bees *hive.Bees) *Hive {
//	all := make(map[string]hive.Dep)
//
//	for _, b := range bees.Bees {
//		m := oneDep(b)
//		if err := mergo.Merge(&all, m); err != nil {
//			panic(err)
//		}
//	}
//	if err := mergo.Merge(&h.Deps, all); err != nil {
//		panic(err)
//	}
//
//	return h
//}

func mergeRecursif(names []string) map[string]hive.Dep  {
	required, err := getBees(names)
	if err != nil || required == nil {
		log.Fatal("mergeRecursif failed: ", required, err)
	}
	ret := make(map[string]hive.Dep)

	for _, b := range required.Bees {
		m := oneDep(b)
		if err := mergo.Merge(&ret, m); err != nil {
			panic(err)
		}
		got := mergeRecursif(b.Deps)
		if err := mergo.Merge(&ret, got); err != nil {
			panic(err)
		}
	}
	return ret
}

func oneDep(b *hive.Bee) map[string]hive.Dep {
	ret := make(map[string]hive.Dep)
	name := fmt.Sprintf("%s/%s", b.Author, b.Name)

	ret[name] = hive.Dep{
		Port: b.Port,
		Dev:  "localhost",
		Prod: "127.0.0.1",
	}
	return ret
}
