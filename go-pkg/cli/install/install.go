package install

import (
	"context"
	"fmt"
	"github.com/benka-me/cell-user/go-pkg/user"
	rpc_core "github.com/benka-me/hive-server-core/go-pkg/rpc-core"
	"github.com/benka-me/hive/go-pkg/hive"
	"github.com/urfave/cli"
	"google.golang.org/grpc"
)

type Hive struct {
	hive.Hive
}

func RunInstall(c *cli.Context) error {
	h := Hive{
		hive.GetHiveMust(),
	}

	core, err := rpc_core.ConnectThroughApi("localhost", grpc.WithInsecure())
	if err != nil {
		return err
	}
	fmt.Println(core, h)
	b, err := core.GetBee(context.TODO(), &hive.BeeReq{BeeName:"", Token:&user.Token{Val:""}})
	fmt.Println(b)

	return nil
}
