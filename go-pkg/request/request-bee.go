package request

import (
	"context"
	rpc_core "github.com/benka-me/hive-server-core/go-pkg/rpc-core"
	"github.com/benka-me/hive/go-pkg/conf"
	"github.com/benka-me/hive/go-pkg/hive"
	"google.golang.org/grpc"
)

func GetRemoteBees(urls hive.Namespaces) ([]*hive.Bee, error) {
	core, err := rpc_core.ConnectThroughApi("localhost", grpc.WithInsecure())
	if err != nil {
		return []*hive.Bee{}, err
	}

	res, err := core.GetBees(context.Background(), &hive.BeesReq{
		Token:    conf.UserToken(),
		BeeNames: urls.Array(),
	})
	if res == nil {
		return []*hive.Bee{}, err
	}
	return res.Bees, err
}

func GetRemoteBee(url hive.Namespace) (*hive.Bee, error) {
	core, err := rpc_core.ConnectThroughApi("localhost", grpc.WithInsecure())
	if err != nil {
		return &hive.Bee{}, err
	}

	res, err := core.GetBee(context.Background(), &hive.BeeReq{
		Token:    conf.UserToken(),
		BeeName: url.NamespaceStr(),
	})
	return res, err
}
