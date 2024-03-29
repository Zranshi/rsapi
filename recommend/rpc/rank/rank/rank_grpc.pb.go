// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v4.25.1
// source: rank.proto

package rank

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
	Rank_Ping_FullMethodName = "/rank.Rank/Ping"
)

// RankClient is the client API for Rank service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type RankClient interface {
	Ping(ctx context.Context, in *Request, opts ...grpc.CallOption) (*Response, error)
}

type rankClient struct {
	cc grpc.ClientConnInterface
}

func NewRankClient(cc grpc.ClientConnInterface) RankClient {
	return &rankClient{cc}
}

func (c *rankClient) Ping(ctx context.Context, in *Request, opts ...grpc.CallOption) (*Response, error) {
	out := new(Response)
	err := c.cc.Invoke(ctx, Rank_Ping_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// RankServer is the server API for Rank service.
// All implementations must embed UnimplementedRankServer
// for forward compatibility
type RankServer interface {
	Ping(context.Context, *Request) (*Response, error)
	mustEmbedUnimplementedRankServer()
}

// UnimplementedRankServer must be embedded to have forward compatible implementations.
type UnimplementedRankServer struct {
}

func (UnimplementedRankServer) Ping(context.Context, *Request) (*Response, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Ping not implemented")
}
func (UnimplementedRankServer) mustEmbedUnimplementedRankServer() {}

// UnsafeRankServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to RankServer will
// result in compilation errors.
type UnsafeRankServer interface {
	mustEmbedUnimplementedRankServer()
}

func RegisterRankServer(s grpc.ServiceRegistrar, srv RankServer) {
	s.RegisterService(&Rank_ServiceDesc, srv)
}

func _Rank_Ping_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Request)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RankServer).Ping(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Rank_Ping_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RankServer).Ping(ctx, req.(*Request))
	}
	return interceptor(ctx, in, info, handler)
}

// Rank_ServiceDesc is the grpc.ServiceDesc for Rank service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Rank_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "rank.Rank",
	HandlerType: (*RankServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Ping",
			Handler:    _Rank_Ping_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "rank.proto",
}
