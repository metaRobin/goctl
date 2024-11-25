// Code generated by goctl. DO NOT EDIT!
// Source: hello.proto

package client

import (
	"context"

	"github.com/metaRobin/goctl/example/rpc/hello/pb/hello"
	"github.com/zeromicro/go-zero/zrpc"
	"google.golang.org/grpc"
)

type (
	HelloReq  = hello.HelloReq
	HelloResp = hello.HelloResp

	Greet interface {
		SayHello(ctx context.Context, in *HelloReq, opts ...grpc.CallOption) (*HelloResp, error)
	}

	defaultGreet struct {
		cli zrpc.Client
	}
)

func NewGreet(cli zrpc.Client) Greet {
	return &defaultGreet{
		cli: cli,
	}
}

func (m *defaultGreet) SayHello(ctx context.Context, in *HelloReq, opts ...grpc.CallOption) (*HelloResp, error) {
	client := hello.NewGreetClient(m.cli.Conn())
	return client.SayHello(ctx, in, opts...)
}
