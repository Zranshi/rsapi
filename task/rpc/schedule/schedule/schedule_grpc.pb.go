// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v4.25.1
// source: schedule.proto

package schedule

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
	Schedule_Ping_FullMethodName = "/schedule.Schedule/Ping"
)

// ScheduleClient is the client API for Schedule service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type ScheduleClient interface {
	Ping(ctx context.Context, in *Request, opts ...grpc.CallOption) (*Response, error)
}

type scheduleClient struct {
	cc grpc.ClientConnInterface
}

func NewScheduleClient(cc grpc.ClientConnInterface) ScheduleClient {
	return &scheduleClient{cc}
}

func (c *scheduleClient) Ping(ctx context.Context, in *Request, opts ...grpc.CallOption) (*Response, error) {
	out := new(Response)
	err := c.cc.Invoke(ctx, Schedule_Ping_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ScheduleServer is the server API for Schedule service.
// All implementations must embed UnimplementedScheduleServer
// for forward compatibility
type ScheduleServer interface {
	Ping(context.Context, *Request) (*Response, error)
	mustEmbedUnimplementedScheduleServer()
}

// UnimplementedScheduleServer must be embedded to have forward compatible implementations.
type UnimplementedScheduleServer struct {
}

func (UnimplementedScheduleServer) Ping(context.Context, *Request) (*Response, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Ping not implemented")
}
func (UnimplementedScheduleServer) mustEmbedUnimplementedScheduleServer() {}

// UnsafeScheduleServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to ScheduleServer will
// result in compilation errors.
type UnsafeScheduleServer interface {
	mustEmbedUnimplementedScheduleServer()
}

func RegisterScheduleServer(s grpc.ServiceRegistrar, srv ScheduleServer) {
	s.RegisterService(&Schedule_ServiceDesc, srv)
}

func _Schedule_Ping_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Request)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ScheduleServer).Ping(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Schedule_Ping_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ScheduleServer).Ping(ctx, req.(*Request))
	}
	return interceptor(ctx, in, info, handler)
}

// Schedule_ServiceDesc is the grpc.ServiceDesc for Schedule service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Schedule_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "schedule.Schedule",
	HandlerType: (*ScheduleServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Ping",
			Handler:    _Schedule_Ping_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "schedule.proto",
}
