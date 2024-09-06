// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package pb

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

// ExecutorClient is the client API for Executor service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type ExecutorClient interface {
	Start(ctx context.Context, in *StartRequest, opts ...grpc.CallOption) (*StartResult, error)
	Terminate(ctx context.Context, in *TermRequest, opts ...grpc.CallOption) (*TermResult, error)
}

type executorClient struct {
	cc grpc.ClientConnInterface
}

func NewExecutorClient(cc grpc.ClientConnInterface) ExecutorClient {
	return &executorClient{cc}
}

func (c *executorClient) Start(ctx context.Context, in *StartRequest, opts ...grpc.CallOption) (*StartResult, error) {
	out := new(StartResult)
	err := c.cc.Invoke(ctx, "/pb.Executor/Start", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *executorClient) Terminate(ctx context.Context, in *TermRequest, opts ...grpc.CallOption) (*TermResult, error) {
	out := new(TermResult)
	err := c.cc.Invoke(ctx, "/pb.Executor/Terminate", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ExecutorServer is the server API for Executor service.
// All implementations should embed UnimplementedExecutorServer
// for forward compatibility
type ExecutorServer interface {
	Start(context.Context, *StartRequest) (*StartResult, error)
	Terminate(context.Context, *TermRequest) (*TermResult, error)
}

// UnimplementedExecutorServer should be embedded to have forward compatible implementations.
type UnimplementedExecutorServer struct {
}

func (UnimplementedExecutorServer) Start(context.Context, *StartRequest) (*StartResult, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Start not implemented")
}
func (UnimplementedExecutorServer) Terminate(context.Context, *TermRequest) (*TermResult, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Terminate not implemented")
}

// UnsafeExecutorServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to ExecutorServer will
// result in compilation errors.
type UnsafeExecutorServer interface {
	mustEmbedUnimplementedExecutorServer()
}

func RegisterExecutorServer(s grpc.ServiceRegistrar, srv ExecutorServer) {
	s.RegisterService(&Executor_ServiceDesc, srv)
}

func _Executor_Start_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(StartRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ExecutorServer).Start(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.Executor/Start",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ExecutorServer).Start(ctx, req.(*StartRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Executor_Terminate_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(TermRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ExecutorServer).Terminate(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.Executor/Terminate",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ExecutorServer).Terminate(ctx, req.(*TermRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// Executor_ServiceDesc is the grpc.ServiceDesc for Executor service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Executor_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "pb.Executor",
	HandlerType: (*ExecutorServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Start",
			Handler:    _Executor_Start_Handler,
		},
		{
			MethodName: "Terminate",
			Handler:    _Executor_Terminate_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "pb/executor.proto",
}
