package publish

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
)

func Run(c *cli.Context) error {
	core, err := rpc_core.ConnectThroughApi("localhost", grpc.WithInsecure())
	if err != nil {
		return err
	}

	b, errBee := hive.Getbee()
	h, errHive := hive.GetHive()

	if errBee == nil && errHive != nil {
		_, err = core.PublishBee(context.Background(), &hive.PushBee{
			Bee:b,
			Token: &user.Token{Val:conf.Parse().Token},
		})
	} else if errBee != nil && errHive == nil {
		_, err = core.PublishHive(context.Background(), &hive.PushHive{
			Hive:h,
			Token: &user.Token{Val:conf.Parse().Token},
		})
	} else if errBee != nil && errHive != nil {
		return errors.New("no hive.yaml or bee.yaml detected")
	} else if errBee == nil && errHive == nil {
		return errors.New("both hive.yaml and bee.yaml detected, please remove one")
	}

	if err != nil {
		return err
	}
	fmt.Println("Successfully published")
	return nil
}