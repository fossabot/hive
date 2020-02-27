package request

import (
	"context"
	rpc_core "github.com/benka-me/hive-server-core/go-pkg/rpc-core"
	"github.com/benka-me/hive/go-pkg/conf"
	"github.com/benka-me/hive/go-pkg/hive"
	"google.golang.org/grpc"
)

func GetBees(urls hive.URLs) ([]*hive.Bee, error) {
	core, err := rpc_core.ConnectThroughApi("localhost", grpc.WithInsecure())
	if err != nil {
		return []*hive.Bee{}, err
	}

	res, err := core.GetBees(context.Background(), &hive.BeesReq{
		Token:    conf.UserToken(),
		BeeNames: urls.Array(),
	})
	return res.Bees, err
}

func GetBee(url hive.URL) (*hive.Bee, error) {
	core, err := rpc_core.ConnectThroughApi("localhost", grpc.WithInsecure())
	if err != nil {
		return &hive.Bee{}, err
	}

	res, err := core.GetBee(context.Background(), &hive.BeeReq{
		Token:    conf.UserToken(),
		BeeName: url.Str(),
	})
	return res, err
}
