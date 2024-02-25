// Code generated by goctl. DO NOT EDIT.
// Source: analyst.proto

package analystclient

import (
	"context"

	"rsapi/stock/rpc/analyst/analyst"

	"github.com/zeromicro/go-zero/zrpc"
	"google.golang.org/grpc"
)

type (
	Request  = analyst.Request
	Response = analyst.Response

	Analyst interface {
		Ping(ctx context.Context, in *Request, opts ...grpc.CallOption) (*Response, error)
	}

	defaultAnalyst struct {
		cli zrpc.Client
	}
)

func NewAnalyst(cli zrpc.Client) Analyst {
	return &defaultAnalyst{
		cli: cli,
	}
}

func (m *defaultAnalyst) Ping(ctx context.Context, in *Request, opts ...grpc.CallOption) (*Response, error) {
	client := analyst.NewAnalystClient(m.cli.Conn())
	return client.Ping(ctx, in, opts...)
}
