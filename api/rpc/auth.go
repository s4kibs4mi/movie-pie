package rpc

import (
	"context"
	"fmt"
	"github.com/twitchtv/twirp"
)

func AuthHook() *twirp.ServerHooks {
	hook := &twirp.ServerHooks{}
	hook.RequestReceived = func(ctx context.Context) (context.Context, error) {
		fmt.Println("Auth Hook Request Received...")

		appKeyValues := ctx.Value("Appkey")
		if appKeyValues == nil {
			return ctx, twirp.NewError(twirp.Aborted, "Unauthorized request")
		}
		appKeys := appKeyValues.([]string)
		if len(appKeys) > 0 && appKeys[0] == "1" {
			return ctx, nil
		}
		return ctx, twirp.NewError(twirp.Aborted, "Unauthorized request")
	}
	return hook
}
