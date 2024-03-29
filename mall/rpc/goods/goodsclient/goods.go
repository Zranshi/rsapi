// Code generated by goctl. DO NOT EDIT.
// Source: goods.proto

package goodsclient

import (
	"context"

	"rsapi/mall/rpc/goods/goods"

	"github.com/zeromicro/go-zero/zrpc"
	"google.golang.org/grpc"
)

type (
	Request  = goods.Request
	Response = goods.Response

	Goods interface {
		Ping(ctx context.Context, in *Request, opts ...grpc.CallOption) (*Response, error)
	}

	defaultGoods struct {
		cli zrpc.Client
	}
)

func NewGoods(cli zrpc.Client) Goods {
	return &defaultGoods{
		cli: cli,
	}
}

func (m *defaultGoods) Ping(ctx context.Context, in *Request, opts ...grpc.CallOption) (*Response, error) {
	client := goods.NewGoodsClient(m.cli.Conn())
	return client.Ping(ctx, in, opts...)
}
