// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package bb

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

// DdddClient is the client API for Dddd service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type DdddClient interface {
	SayHello22(ctx context.Context, in *HelloRequest, opts ...grpc.CallOption) (*HelloReply, error)
}

type ddddClient struct {
	cc grpc.ClientConnInterface
}

func NewDdddClient(cc grpc.ClientConnInterface) DdddClient {
	return &ddddClient{cc}
}

func (c *ddddClient) SayHello22(ctx context.Context, in *HelloRequest, opts ...grpc.CallOption) (*HelloReply, error) {
	out := new(HelloReply)
	err := c.cc.Invoke(ctx, "/aa.bb.Dddd/SayHello22", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// DdddServer is the server API for Dddd service.
// All implementations must embed UnimplementedDdddServer
// for forward compatibility
type DdddServer interface {
	SayHello22(context.Context, *HelloRequest) (*HelloReply, error)
	mustEmbedUnimplementedDdddServer()
}

// UnimplementedDdddServer must be embedded to have forward compatible implementations.
type UnimplementedDdddServer struct {
}

func (UnimplementedDdddServer) SayHello22(context.Context, *HelloRequest) (*HelloReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SayHello22 not implemented")
}
func (UnimplementedDdddServer) mustEmbedUnimplementedDdddServer() {}

// UnsafeDdddServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to DdddServer will
// result in compilation errors.
type UnsafeDdddServer interface {
	mustEmbedUnimplementedDdddServer()
}

func RegisterDdddServer(s grpc.ServiceRegistrar, srv DdddServer) {
	s.RegisterService(&Dddd_ServiceDesc, srv)
}

func _Dddd_SayHello22_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(HelloRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DdddServer).SayHello22(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/aa.bb.Dddd/SayHello22",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DdddServer).SayHello22(ctx, req.(*HelloRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// Dddd_ServiceDesc is the grpc.ServiceDesc for Dddd service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Dddd_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "aa.bb.Dddd",
	HandlerType: (*DdddServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "SayHello22",
			Handler:    _Dddd_SayHello22_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "Cc.proto",
}
