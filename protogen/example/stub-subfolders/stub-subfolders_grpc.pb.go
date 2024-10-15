// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.5.1
// - protoc             v5.28.2
// source: stub-subfolders.proto

package stub_subfolders

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
	Gripmock_SayHello_FullMethodName = "/stub_subfolders.Gripmock/SayHello"
)

// GripmockClient is the client API for Gripmock service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
//
// The Gripmock service definition.
type GripmockClient interface {
	// simple unary method
	SayHello(ctx context.Context, in *Request, opts ...grpc.CallOption) (*Reply, error)
}

type gripmockClient struct {
	cc grpc.ClientConnInterface
}

func NewGripmockClient(cc grpc.ClientConnInterface) GripmockClient {
	return &gripmockClient{cc}
}

func (c *gripmockClient) SayHello(ctx context.Context, in *Request, opts ...grpc.CallOption) (*Reply, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(Reply)
	err := c.cc.Invoke(ctx, Gripmock_SayHello_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// GripmockServer is the server API for Gripmock service.
// All implementations must embed UnimplementedGripmockServer
// for forward compatibility.
//
// The Gripmock service definition.
type GripmockServer interface {
	// simple unary method
	SayHello(context.Context, *Request) (*Reply, error)
	mustEmbedUnimplementedGripmockServer()
}

// UnimplementedGripmockServer must be embedded to have
// forward compatible implementations.
//
// NOTE: this should be embedded by value instead of pointer to avoid a nil
// pointer dereference when methods are called.
type UnimplementedGripmockServer struct{}

func (UnimplementedGripmockServer) SayHello(context.Context, *Request) (*Reply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SayHello not implemented")
}
func (UnimplementedGripmockServer) mustEmbedUnimplementedGripmockServer() {}
func (UnimplementedGripmockServer) testEmbeddedByValue()                  {}

// UnsafeGripmockServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to GripmockServer will
// result in compilation errors.
type UnsafeGripmockServer interface {
	mustEmbedUnimplementedGripmockServer()
}

func RegisterGripmockServer(s grpc.ServiceRegistrar, srv GripmockServer) {
	// If the following call pancis, it indicates UnimplementedGripmockServer was
	// embedded by pointer and is nil.  This will cause panics if an
	// unimplemented method is ever invoked, so we test this at initialization
	// time to prevent it from happening at runtime later due to I/O.
	if t, ok := srv.(interface{ testEmbeddedByValue() }); ok {
		t.testEmbeddedByValue()
	}
	s.RegisterService(&Gripmock_ServiceDesc, srv)
}

func _Gripmock_SayHello_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Request)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GripmockServer).SayHello(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Gripmock_SayHello_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GripmockServer).SayHello(ctx, req.(*Request))
	}
	return interceptor(ctx, in, info, handler)
}

// Gripmock_ServiceDesc is the grpc.ServiceDesc for Gripmock service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Gripmock_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "stub_subfolders.Gripmock",
	HandlerType: (*GripmockServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "SayHello",
			Handler:    _Gripmock_SayHello_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "stub-subfolders.proto",
}
