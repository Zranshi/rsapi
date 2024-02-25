// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v4.25.1
// source: converter.proto

package converter

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
	Converter_Ping_FullMethodName = "/converter.Converter/Ping"
)

// ConverterClient is the client API for Converter service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type ConverterClient interface {
	Ping(ctx context.Context, in *Request, opts ...grpc.CallOption) (*Response, error)
}

type converterClient struct {
	cc grpc.ClientConnInterface
}

func NewConverterClient(cc grpc.ClientConnInterface) ConverterClient {
	return &converterClient{cc}
}

func (c *converterClient) Ping(ctx context.Context, in *Request, opts ...grpc.CallOption) (*Response, error) {
	out := new(Response)
	err := c.cc.Invoke(ctx, Converter_Ping_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ConverterServer is the server API for Converter service.
// All implementations must embed UnimplementedConverterServer
// for forward compatibility
type ConverterServer interface {
	Ping(context.Context, *Request) (*Response, error)
	mustEmbedUnimplementedConverterServer()
}

// UnimplementedConverterServer must be embedded to have forward compatible implementations.
type UnimplementedConverterServer struct {
}

func (UnimplementedConverterServer) Ping(context.Context, *Request) (*Response, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Ping not implemented")
}
func (UnimplementedConverterServer) mustEmbedUnimplementedConverterServer() {}

// UnsafeConverterServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to ConverterServer will
// result in compilation errors.
type UnsafeConverterServer interface {
	mustEmbedUnimplementedConverterServer()
}

func RegisterConverterServer(s grpc.ServiceRegistrar, srv ConverterServer) {
	s.RegisterService(&Converter_ServiceDesc, srv)
}

func _Converter_Ping_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Request)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ConverterServer).Ping(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Converter_Ping_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ConverterServer).Ping(ctx, req.(*Request))
	}
	return interceptor(ctx, in, info, handler)
}

// Converter_ServiceDesc is the grpc.ServiceDesc for Converter service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Converter_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "converter.Converter",
	HandlerType: (*ConverterServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Ping",
			Handler:    _Converter_Ping_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "converter.proto",
}
