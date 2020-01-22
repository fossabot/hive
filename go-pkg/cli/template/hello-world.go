package rpc

import (
	"context"
	"fmt"
	"github.com/benka-me/hive/go-pkg/hive"
)

func (c *App) HelloWorld(ctx context.Context, user *hive.User) (*hive.Welcome, error) {

	return &hive.Welcome{Message: fmt.Sprintf("Hello %s", user.Username)}, nil
}
