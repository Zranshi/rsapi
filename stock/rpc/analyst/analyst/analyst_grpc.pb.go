// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v4.25.1
// source: analyst.proto

package analyst

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
	Analyst_Ping_FullMethodName = "/analyst.Analyst/Ping"
)

// AnalystClient is the client API for Analyst service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type AnalystClient interface {
	Ping(ctx context.Context, in *Request, opts ...grpc.CallOption) (*Response, error)
}

type analystClient struct {
	cc grpc.ClientConnInterface
}

func NewAnalystClient(cc grpc.ClientConnInterface) AnalystClient {
	return &analystClient{cc}
}

func (c *analystClient) Ping(ctx context.Context, in *Request, opts ...grpc.CallOption) (*Response, error) {
	out := new(Response)
	err := c.cc.Invoke(ctx, Analyst_Ping_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// AnalystServer is the server API for Analyst service.
// All implementations must embed UnimplementedAnalystServer
// for forward compatibility
type AnalystServer interface {
	Ping(context.Context, *Request) (*Response, error)
	mustEmbedUnimplementedAnalystServer()
}

// UnimplementedAnalystServer must be embedded to have forward compatible implementations.
type UnimplementedAnalystServer struct {
}

func (UnimplementedAnalystServer) Ping(context.Context, *Request) (*Response, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Ping not implemented")
}
func (UnimplementedAnalystServer) mustEmbedUnimplementedAnalystServer() {}

// UnsafeAnalystServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to AnalystServer will
// result in compilation errors.
type UnsafeAnalystServer interface {
	mustEmbedUnimplementedAnalystServer()
}

func RegisterAnalystServer(s grpc.ServiceRegistrar, srv AnalystServer) {
	s.RegisterService(&Analyst_ServiceDesc, srv)
}

func _Analyst_Ping_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Request)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AnalystServer).Ping(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Analyst_Ping_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AnalystServer).Ping(ctx, req.(*Request))
	}
	return interceptor(ctx, in, info, handler)
}

// Analyst_ServiceDesc is the grpc.ServiceDesc for Analyst service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Analyst_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "analyst.Analyst",
	HandlerType: (*AnalystServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Ping",
			Handler:    _Analyst_Ping_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "analyst.proto",
}
