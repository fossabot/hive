package rpc

import (
	"fmt"
	rpcHive "github.com/benka-me/hive/go-pkg/rpc-hive"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"log"
	"net"
)

type App struct {}

const port = ":8012"

var grpcServer *grpc.Server

func Server_2_0() {
	var err error
	var app *App

	grpcServer = grpc.NewServer()
	lis, err := net.Listen("tcp", port)

	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	{
		rpcHive.RegisterHiveServer(grpcServer, app)
		reflection.Register(grpcServer)
	}
	fmt.Println("HiveServer Grpc 2.0 on port ", port)
	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
