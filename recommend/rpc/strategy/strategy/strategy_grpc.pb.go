// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v4.25.1
// source: strategy.proto

package strategy

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

const (
	Strategy_Ping_FullMethodName = "/strategy.Strategy/Ping"
)

// StrategyClient is the client API for Strategy service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type StrategyClient interface {
	Ping(ctx context.Context, in *Request, opts ...grpc.CallOption) (*Response, error)
}

type strategyClient struct {
	cc grpc.ClientConnInterface
}

func NewStrategyClient(cc grpc.ClientConnInterface) StrategyClient {
	return &strategyClient{cc}
}

func (c *strategyClient) Ping(ctx context.Context, in *Request, opts ...grpc.CallOption) (*Response, error) {
	out := new(Response)
	err := c.cc.Invoke(ctx, Strategy_Ping_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// StrategyServer is the server API for Strategy service.
// All implementations must embed UnimplementedStrategyServer
// for forward compatibility
type StrategyServer interface {
	Ping(context.Context, *Request) (*Response, error)
	mustEmbedUnimplementedStrategyServer()
}

// UnimplementedStrategyServer must be embedded to have forward compatible implementations.
type UnimplementedStrategyServer struct {
}

func (UnimplementedStrategyServer) Ping(context.Context, *Request) (*Response, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Ping not implemented")
}
func (UnimplementedStrategyServer) mustEmbedUnimplementedStrategyServer() {}

// UnsafeStrategyServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to StrategyServer will
// result in compilation errors.
type UnsafeStrategyServer interface {
	mustEmbedUnimplementedStrategyServer()
}

func RegisterStrategyServer(s grpc.ServiceRegistrar, srv StrategyServer) {
	s.RegisterService(&Strategy_ServiceDesc, srv)
}

func _Strategy_Ping_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Request)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(StrategyServer).Ping(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Strategy_Ping_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(StrategyServer).Ping(ctx, req.(*Request))
	}
	return interceptor(ctx, in, info, handler)
}

// Strategy_ServiceDesc is the grpc.ServiceDesc for Strategy service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Strategy_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "strategy.Strategy",
	HandlerType: (*StrategyServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Ping",
			Handler:    _Strategy_Ping_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "strategy.proto",
}
