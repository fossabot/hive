package install

import (
	"context"
	"errors"
	"fmt"
	"github.com/benka-me/cell-user/go-pkg/user"
	rpc_core "github.com/benka-me/hive-server-core/go-pkg/rpc-core"
	"github.com/benka-me/hive/go-pkg/conf"
	"github.com/benka-me/hive/go-pkg/hive"
	"github.com/urfave/cli"
	"google.golang.org/grpc"
	"os"
)

type Hive struct {
	hive.Hive
}

func RunInstall(c *cli.Context) error {
	if len(os.Args) < 3 {
		return errors.New("bad argument")
	}

	name := os.Args[2]
	h := Hive{
		hive.GetHiveMust(),
	}

	core, err := rpc_core.ConnectThroughApi("localhost", grpc.WithInsecure())
	if err != nil {
		return err
	}
	fmt.Println(name, h)

	b, err := core.GetBee(context.TODO(), &hive.BeeReq{BeeName:name, Token:&user.Token{Val: conf.Token()}})
	fmt.Println(b, err)

	return nil
}
