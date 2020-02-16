package rpc

import (
	"context"
	"fmt"
	"github.com/benka-me/hive/go-pkg/hive"
)

func (c *App) HelloWorld(ctx context.Context, req *gtgg.Request) (*gtgg.Greeting, error) {

	return &gtgg.Greeting{Msg: fmt.Sprintf("Hello %s", req.Msg)}, nil
}
