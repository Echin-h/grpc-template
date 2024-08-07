// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.5.1
// - protoc             (unknown)
// source: proto/hello/v1/hello.proto

package helloV1

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.64.0 or later.
const _ = grpc.SupportPackageIsVersion9

const (
	GreeterService_SayHello_FullMethodName        = "/proto.hello.v1.GreeterService/SayHello"
	GreeterService_SayHelloAgain_FullMethodName   = "/proto.hello.v1.GreeterService/SayHelloAgain"
	GreeterService_SayHelloThirdly_FullMethodName = "/proto.hello.v1.GreeterService/SayHelloThirdly"
)

// GreeterServiceClient is the client API for GreeterService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type GreeterServiceClient interface {
	SayHello(ctx context.Context, in *SayHelloRequest, opts ...grpc.CallOption) (*SayHelloResponse, error)
	SayHelloAgain(ctx context.Context, in *SayHelloAgainRequest, opts ...grpc.CallOption) (*SayHelloAgainResponse, error)
	SayHelloThirdly(ctx context.Context, in *SayHelloThirdlyRequest, opts ...grpc.CallOption) (*SayHelloThirdlyResponse, error)
}

type greeterServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewGreeterServiceClient(cc grpc.ClientConnInterface) GreeterServiceClient {
	return &greeterServiceClient{cc}
}

func (c *greeterServiceClient) SayHello(ctx context.Context, in *SayHelloRequest, opts ...grpc.CallOption) (*SayHelloResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(SayHelloResponse)
	err := c.cc.Invoke(ctx, GreeterService_SayHello_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *greeterServiceClient) SayHelloAgain(ctx context.Context, in *SayHelloAgainRequest, opts ...grpc.CallOption) (*SayHelloAgainResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(SayHelloAgainResponse)
	err := c.cc.Invoke(ctx, GreeterService_SayHelloAgain_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *greeterServiceClient) SayHelloThirdly(ctx context.Context, in *SayHelloThirdlyRequest, opts ...grpc.CallOption) (*SayHelloThirdlyResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(SayHelloThirdlyResponse)
	err := c.cc.Invoke(ctx, GreeterService_SayHelloThirdly_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// GreeterServiceServer is the server API for GreeterService service.
// All implementations must embed UnimplementedGreeterServiceServer
// for forward compatibility.
type GreeterServiceServer interface {
	SayHello(context.Context, *SayHelloRequest) (*SayHelloResponse, error)
	SayHelloAgain(context.Context, *SayHelloAgainRequest) (*SayHelloAgainResponse, error)
	SayHelloThirdly(context.Context, *SayHelloThirdlyRequest) (*SayHelloThirdlyResponse, error)
	mustEmbedUnimplementedGreeterServiceServer()
}

// UnimplementedGreeterServiceServer must be embedded to have
// forward compatible implementations.
//
// NOTE: this should be embedded by value instead of pointer to avoid a nil
// pointer dereference when methods are called.
type UnimplementedGreeterServiceServer struct{}

func (UnimplementedGreeterServiceServer) SayHello(context.Context, *SayHelloRequest) (*SayHelloResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SayHello not implemented")
}
func (UnimplementedGreeterServiceServer) SayHelloAgain(context.Context, *SayHelloAgainRequest) (*SayHelloAgainResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SayHelloAgain not implemented")
}
func (UnimplementedGreeterServiceServer) SayHelloThirdly(context.Context, *SayHelloThirdlyRequest) (*SayHelloThirdlyResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SayHelloThirdly not implemented")
}
func (UnimplementedGreeterServiceServer) mustEmbedUnimplementedGreeterServiceServer() {}
func (UnimplementedGreeterServiceServer) testEmbeddedByValue()                        {}

// UnsafeGreeterServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to GreeterServiceServer will
// result in compilation errors.
type UnsafeGreeterServiceServer interface {
	mustEmbedUnimplementedGreeterServiceServer()
}

func RegisterGreeterServiceServer(s grpc.ServiceRegistrar, srv GreeterServiceServer) {
	// If the following call pancis, it indicates UnimplementedGreeterServiceServer was
	// embedded by pointer and is nil.  This will cause panics if an
	// unimplemented method is ever invoked, so we test this at initialization
	// time to prevent it from happening at runtime later due to I/O.
	if t, ok := srv.(interface{ testEmbeddedByValue() }); ok {
		t.testEmbeddedByValue()
	}
	s.RegisterService(&GreeterService_ServiceDesc, srv)
}

func _GreeterService_SayHello_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SayHelloRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GreeterServiceServer).SayHello(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: GreeterService_SayHello_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GreeterServiceServer).SayHello(ctx, req.(*SayHelloRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _GreeterService_SayHelloAgain_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SayHelloAgainRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GreeterServiceServer).SayHelloAgain(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: GreeterService_SayHelloAgain_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GreeterServiceServer).SayHelloAgain(ctx, req.(*SayHelloAgainRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _GreeterService_SayHelloThirdly_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SayHelloThirdlyRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GreeterServiceServer).SayHelloThirdly(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: GreeterService_SayHelloThirdly_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GreeterServiceServer).SayHelloThirdly(ctx, req.(*SayHelloThirdlyRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// GreeterService_ServiceDesc is the grpc.ServiceDesc for GreeterService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var GreeterService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "proto.hello.v1.GreeterService",
	HandlerType: (*GreeterServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "SayHello",
			Handler:    _GreeterService_SayHello_Handler,
		},
		{
			MethodName: "SayHelloAgain",
			Handler:    _GreeterService_SayHelloAgain_Handler,
		},
		{
			MethodName: "SayHelloThirdly",
			Handler:    _GreeterService_SayHelloThirdly_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "proto/hello/v1/hello.proto",
}
