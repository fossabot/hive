package engine

import (
	"fmt"
	"github.com/benka-me/hive/go-pkg/hive"
	"google.golang.org/grpc"
	"os"
	"strconv"
)

func (r *RunConfig) ThisPort(namespace string) string {
	if _, ok := r.Deps[namespace] ; !ok {
		fmt.Println(namespace + " is not subscribed")
		os.Exit(1)
	}
	i := r.Deps[namespace].Port
	return ":" + strconv.FormatInt(int64(i), 10)
}

func (r *RunConfig) GrpcConn(namespace string, gateway bool, options ...grpc.DialOption) (*grpc.ClientConn, error){
	var host string
	if r.Dev {
		host = r.Deps[namespace].Dev
	} else {
		host =  r.Deps[namespace].Prod
	}

	var port string
	var isGateway string
	if gateway {
		port = strconv.FormatInt(int64(r.GatewayPort,), 10)
		isGateway = "(Through gateway)"
	} else {
		port = strconv.FormatInt(int64(r.Deps[namespace].Port,), 10)
	}
	address := fmt.Sprintf("%s:%s", host, port)
	fmt.Printf("New client to %s service on %s %s\n", namespace, address, isGateway)
	return grpc.Dial(address, options...)
}

func Parse(namespace string, dev bool) *RunConfig {
	hive, err := hive.GetLocalHiveFromString(namespace)
	if err != nil {panic(err)}

	return &RunConfig{
		Dev:  dev,
		Deps: hive.Deps,
	}
}

type RunConfig struct {
	Dev bool
	Deps map[string]*hive.Dep
	GatewayPort int
}
