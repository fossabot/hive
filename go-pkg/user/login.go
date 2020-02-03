package user

import (
	"context"
	"fmt"
	"github.com/benka-me/cell-user/go-pkg/jwt"
	rpc_user "github.com/benka-me/cell-user/go-pkg/rpc-user"
	"github.com/benka-me/cell-user/go-pkg/user"
	"github.com/benka-me/hive/go-pkg/cli/scan"
	"github.com/benka-me/hive/go-pkg/conf"
	"github.com/go-playground/validator"
	"github.com/urfave/cli"
	"google.golang.org/grpc"
	"google.golang.org/grpc/status"
	"log"
)

func RunLogin(c *cli.Context) error {

	scan.V = validator.New()
	userCell, err := rpc_user.ConnectThroughApi("localhost", grpc.WithInsecure())
	if err != nil {
		return err
	}
	req := &user.LoginReq{}

	username := scan.Step(
		"Username",
		"required,lte=20,gte=3",
		func(s string) error {return nil})

	req.Identifier = user.GetReq{OneOf:&user.GetReq_Name{Name:username}}
	req.Password = scan.GetPassword("Password")

	res, err := userCell.Login(context.Background(), req)
	if err != nil {
		return err
	}

	cnf := conf.Conf{
		Username: res.Data.Username,
		Token: res.Auth,
	}
	err = cnf.Write()
	if err != nil {
		return err
	}

	return nil
}

func RunWhoAmI(c *cli.Context) error {
	cnf := conf.Conf{}
	err := cnf.Parse()
	if err != nil {
		return err
	}

	if cnf.Username == "" {
		fmt.Println("config username empty, please login")
		return nil
	} else if cnf.Token == "" {
		fmt.Println("username is "+ cnf.Username + " but not authenticated, please login")
		return nil
	}

	userCell, err := rpc_user.ConnectThroughApi("localhost", grpc.WithInsecure())
	if err != nil {
		return err
	}
	req := &user.Token{
		Val: cnf.Token,
	}

	_, err = userCell.Auth(context.Background(), req)
	if err != nil {
		st := status.Convert(err)
		if st.Code() == jwt.TokenExpired {
			cnf.Token = ""
			err := cnf.Write()
			if err != nil {
				log.Println(err)
			}
			fmt.Println("Your username is: ", cnf.Username, " but your token has expired")
			return err
		}
		fmt.Println("Your username is: ", cnf.Username, " but not authenticated.")
		return err
	}
	fmt.Println("You are authenticated as: ", cnf.Username)

	return nil
}
