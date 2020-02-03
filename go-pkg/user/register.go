package user

import (
	"context"
	"fmt"
	rpc_user "github.com/benka-me/cell-user/go-pkg/rpc-user"
	"github.com/benka-me/cell-user/go-pkg/user"
	"github.com/benka-me/hive/go-pkg/cli/scan"
	"github.com/go-playground/validator"
	"github.com/urfave/cli"
	"google.golang.org/grpc"
)

func RunRegister(c *cli.Context) error {
	scan.V = validator.New()
	userCell, err := rpc_user.ConnectThroughApi("localhost", grpc.WithInsecure())
	if err != nil {
		return err
	}
	req := &user.RegisterReq{}

	req.Username = scan.Step(
		"Username",
		"required,lte=20,gte=3",
		func(s string) error {return nil})

	req.Email = scan.Step("Email", "email", func(s string) error {return nil})
	req.Password = scan.GetPassword("Password")

	res, err := userCell.Register(context.Background(), req)

	fmt.Println("res: ", res, err)

	return nil
}
