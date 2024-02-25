// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v4.25.1
// source: p2p.proto

package p2p

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
	P2P_Ping_FullMethodName = "/p2p.P2p/Ping"
)

// P2PClient is the client API for P2P service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type P2PClient interface {
	Ping(ctx context.Context, in *Request, opts ...grpc.CallOption) (*Response, error)
}

type p2PClient struct {
	cc grpc.ClientConnInterface
}

func NewP2PClient(cc grpc.ClientConnInterface) P2PClient {
	return &p2PClient{cc}
}

func (c *p2PClient) Ping(ctx context.Context, in *Request, opts ...grpc.CallOption) (*Response, error) {
	out := new(Response)
	err := c.cc.Invoke(ctx, P2P_Ping_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// P2PServer is the server API for P2P service.
// All implementations must embed UnimplementedP2PServer
// for forward compatibility
type P2PServer interface {
	Ping(context.Context, *Request) (*Response, error)
	mustEmbedUnimplementedP2PServer()
}

// UnimplementedP2PServer must be embedded to have forward compatible implementations.
type UnimplementedP2PServer struct {
}

func (UnimplementedP2PServer) Ping(context.Context, *Request) (*Response, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Ping not implemented")
}
func (UnimplementedP2PServer) mustEmbedUnimplementedP2PServer() {}

// UnsafeP2PServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to P2PServer will
// result in compilation errors.
type UnsafeP2PServer interface {
	mustEmbedUnimplementedP2PServer()
}

func RegisterP2PServer(s grpc.ServiceRegistrar, srv P2PServer) {
	s.RegisterService(&P2P_ServiceDesc, srv)
}

func _P2P_Ping_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Request)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(P2PServer).Ping(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: P2P_Ping_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(P2PServer).Ping(ctx, req.(*Request))
	}
	return interceptor(ctx, in, info, handler)
}

// P2P_ServiceDesc is the grpc.ServiceDesc for P2P service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var P2P_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "p2p.P2p",
	HandlerType: (*P2PServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Ping",
			Handler:    _P2P_Ping_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "p2p.proto",
}