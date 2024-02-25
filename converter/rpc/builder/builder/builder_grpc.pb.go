// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v4.25.1
// source: builder.proto

package builder

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
	Builder_Ping_FullMethodName = "/builder.Builder/Ping"
)

// BuilderClient is the client API for Builder service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type BuilderClient interface {
	Ping(ctx context.Context, in *Request, opts ...grpc.CallOption) (*Response, error)
}

type builderClient struct {
	cc grpc.ClientConnInterface
}

func NewBuilderClient(cc grpc.ClientConnInterface) BuilderClient {
	return &builderClient{cc}
}

func (c *builderClient) Ping(ctx context.Context, in *Request, opts ...grpc.CallOption) (*Response, error) {
	out := new(Response)
	err := c.cc.Invoke(ctx, Builder_Ping_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// BuilderServer is the server API for Builder service.
// All implementations must embed UnimplementedBuilderServer
// for forward compatibility
type BuilderServer interface {
	Ping(context.Context, *Request) (*Response, error)
	mustEmbedUnimplementedBuilderServer()
}

// UnimplementedBuilderServer must be embedded to have forward compatible implementations.
type UnimplementedBuilderServer struct {
}

func (UnimplementedBuilderServer) Ping(context.Context, *Request) (*Response, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Ping not implemented")
}
func (UnimplementedBuilderServer) mustEmbedUnimplementedBuilderServer() {}

// UnsafeBuilderServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to BuilderServer will
// result in compilation errors.
type UnsafeBuilderServer interface {
	mustEmbedUnimplementedBuilderServer()
}

func RegisterBuilderServer(s grpc.ServiceRegistrar, srv BuilderServer) {
	s.RegisterService(&Builder_ServiceDesc, srv)
}

func _Builder_Ping_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Request)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BuilderServer).Ping(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Builder_Ping_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BuilderServer).Ping(ctx, req.(*Request))
	}
	return interceptor(ctx, in, info, handler)
}

// Builder_ServiceDesc is the grpc.ServiceDesc for Builder service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Builder_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "builder.Builder",
	HandlerType: (*BuilderServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Ping",
			Handler:    _Builder_Ping_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "builder.proto",
}