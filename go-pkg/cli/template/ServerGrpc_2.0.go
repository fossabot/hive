package template

import (
	"fmt"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"log"
	"net"
)

type Server struct {}
const port = ":8037"
var grpcServer *grpc.Server
var guide *Server

func Server_2_0() {
	var err error
	grpcServer = grpc.NewServer()
	lis, err := net.Listen("tcp", port)

	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	{
		//api.RegisterApiServer(grpcServer, guide)
		reflection.Register(grpcServer)
	}
	fmt.Println("Server Grpc 2.0 on port ", port)
	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
