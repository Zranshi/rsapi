// Code generated by goctl. DO NOT EDIT.
// Source: cart.proto

package cartclient

import (
	"context"

	"rsapi/mall/rpc/cart/cart"

	"github.com/zeromicro/go-zero/zrpc"
	"google.golang.org/grpc"
)

type (
	Request  = cart.Request
	Response = cart.Response

	Cart interface {
		Ping(ctx context.Context, in *Request, opts ...grpc.CallOption) (*Response, error)
	}

	defaultCart struct {
		cli zrpc.Client
	}
)

func NewCart(cli zrpc.Client) Cart {
	return &defaultCart{
		cli: cli,
	}
}

func (m *defaultCart) Ping(ctx context.Context, in *Request, opts ...grpc.CallOption) (*Response, error) {
	client := cart.NewCartClient(m.cli.Conn())
	return client.Ping(ctx, in, opts...)
}
