package engine

import (
	"fmt"
	"github.com/benka-me/hive/go-pkg/hive"
	"google.golang.org/grpc"
	"strconv"
)

func (r *RunConfig) GrpcConn(url string, gateway bool, options ...grpc.DialOption) (*grpc.ClientConn, error){
	var host string
	if r.Dev {
		host = r.Deps[url].Dev
	} else {
		host =  r.Deps[url].Prod
	}

	var port string
	var isGateway string
	if gateway {
		port = "8088"
		isGateway = "(Through gateway)"
	} else {
		port = strconv.FormatInt(int64(r.Deps[url].Port,), 10)
	}
	address := fmt.Sprintf("%s:%s", host, port)
	fmt.Printf("Connected to %s service on %s %s\n", url, address, isGateway)
	return grpc.Dial(address, options...)
}
func Parse(url string, dev bool) *RunConfig {
	hive, err := hive.GetLocalHiveFromString(url)
	if err != nil {panic(err)}

	return &RunConfig{
		Dev:  dev,
		Deps: hive.Deps,
	}
}

type RunConfig struct {
	Dev bool
	Deps map[string]*hive.Dep
}
