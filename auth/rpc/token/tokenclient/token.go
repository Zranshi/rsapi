// Code generated by goctl. DO NOT EDIT.
// Source: token.proto

package tokenclient

import (
	"context"

	"rsapi/auth/rpc/token/token"

	"github.com/zeromicro/go-zero/zrpc"
	"google.golang.org/grpc"
)

type (
	Request  = token.Request
	Response = token.Response

	Token interface {
		Ping(ctx context.Context, in *Request, opts ...grpc.CallOption) (*Response, error)
	}

	defaultToken struct {
		cli zrpc.Client
	}
)

func NewToken(cli zrpc.Client) Token {
	return &defaultToken{
		cli: cli,
	}
}

func (m *defaultToken) Ping(ctx context.Context, in *Request, opts ...grpc.CallOption) (*Response, error) {
	client := token.NewTokenClient(m.cli.Conn())
	return client.Ping(ctx, in, opts...)
}
