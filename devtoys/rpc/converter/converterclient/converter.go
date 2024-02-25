// Code generated by goctl. DO NOT EDIT.
// Source: converter.proto

package converterclient

import (
	"context"

	"rsapi/devtoys/rpc/converter/converter"

	"github.com/zeromicro/go-zero/zrpc"
	"google.golang.org/grpc"
)

type (
	Request  = converter.Request
	Response = converter.Response

	Converter interface {
		Ping(ctx context.Context, in *Request, opts ...grpc.CallOption) (*Response, error)
	}

	defaultConverter struct {
		cli zrpc.Client
	}
)

func NewConverter(cli zrpc.Client) Converter {
	return &defaultConverter{
		cli: cli,
	}
}

func (m *defaultConverter) Ping(ctx context.Context, in *Request, opts ...grpc.CallOption) (*Response, error) {
	client := converter.NewConverterClient(m.cli.Conn())
	return client.Ping(ctx, in, opts...)
}
