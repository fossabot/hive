package rpc_hive

import (
    "google.golang.org/grpc"
)

func ClientProvider (server string, options ...grpc.DialOption) (HiveClient, error) {
	conn, err := grpc.Dial(server + ":8012", options...)
	if err != nil {
		panic("Cannot dial")
	}

	return NewHiveClient(conn), nil
}

func ConnectThroughApi(server string, options ...grpc.DialOption) (HiveClient, error) {
	conn, err := grpc.Dial(server + ":8080", options...)
	if err != nil {
		panic("Cannot dial")
	}

	return NewHiveClient(conn), nil
}
