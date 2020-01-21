package rpc

import (
	"context"
	"github.com/benka-me/hive/go-pkg/hive"
)

func (c *Cell) HelloWorld(ctx context.Context, user *hive.User) (*hive.Welcome, error) {

	return &hive.Welcome{}, nil
}
