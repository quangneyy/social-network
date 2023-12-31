// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.9.2
// source: newsfeed.proto

package newsfeed

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

// NewsfeedClient is the client API for Newsfeed service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type NewsfeedClient interface {
	Newsfeed(ctx context.Context, in *NewsfeedRequest, opts ...grpc.CallOption) (*NewsfeedResponse, error)
}

type newsfeedClient struct {
	cc grpc.ClientConnInterface
}

func NewNewsfeedClient(cc grpc.ClientConnInterface) NewsfeedClient {
	return &newsfeedClient{cc}
}

func (c *newsfeedClient) Newsfeed(ctx context.Context, in *NewsfeedRequest, opts ...grpc.CallOption) (*NewsfeedResponse, error) {
	out := new(NewsfeedResponse)
	err := c.cc.Invoke(ctx, "/newsfeed.Newsfeed/Newsfeed", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// NewsfeedServer is the server API for Newsfeed service.
// All implementations must embed UnimplementedNewsfeedServer
// for forward compatibility
type NewsfeedServer interface {
	Newsfeed(context.Context, *NewsfeedRequest) (*NewsfeedResponse, error)
	mustEmbedUnimplementedNewsfeedServer()
}

// UnimplementedNewsfeedServer must be embedded to have forward compatible implementations.
type UnimplementedNewsfeedServer struct {
}

func (UnimplementedNewsfeedServer) Newsfeed(context.Context, *NewsfeedRequest) (*NewsfeedResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Newsfeed not implemented")
}
func (UnimplementedNewsfeedServer) mustEmbedUnimplementedNewsfeedServer() {}

// UnsafeNewsfeedServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to NewsfeedServer will
// result in compilation errors.
type UnsafeNewsfeedServer interface {
	mustEmbedUnimplementedNewsfeedServer()
}

func RegisterNewsfeedServer(s grpc.ServiceRegistrar, srv NewsfeedServer) {
	s.RegisterService(&Newsfeed_ServiceDesc, srv)
}

func _Newsfeed_Newsfeed_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(NewsfeedRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(NewsfeedServer).Newsfeed(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/newsfeed.Newsfeed/Newsfeed",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(NewsfeedServer).Newsfeed(ctx, req.(*NewsfeedRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// Newsfeed_ServiceDesc is the grpc.ServiceDesc for Newsfeed service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Newsfeed_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "newsfeed.Newsfeed",
	HandlerType: (*NewsfeedServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Newsfeed",
			Handler:    _Newsfeed_Newsfeed_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "newsfeed.proto",
}
