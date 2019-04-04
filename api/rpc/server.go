package rpc

import (
	"context"
	"github.com/codersgarage/golang-restful-boilerplate/proto/defs"
)

type RPCServer struct {
}

func (s *RPCServer) CreateMonkey(ctx context.Context, req *defs.ReqCreateMonkey) (*defs.ResCreateMonkey, error) {
	return &defs.ResCreateMonkey{}, nil
}
