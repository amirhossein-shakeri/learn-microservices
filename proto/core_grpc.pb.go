// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v4.25.1
// source: core.proto

package core

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

// HubClusterClient is the client API for HubCluster service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type HubClusterClient interface {
	GetHub(ctx context.Context, in *SelectHub, opts ...grpc.CallOption) (*Hub, error)
	// rpc GetHubs (SelectHubs) return (stream Hubs) {}
	CreateHub(ctx context.Context, in *CreateHubRequest, opts ...grpc.CallOption) (*Hub, error)
	// rpc PauseHub (SelectHub) returns () {} // what about returning nothing?
	// pause / resume / start / run hub?
	TerminateHub(ctx context.Context, in *SelectHub, opts ...grpc.CallOption) (*Empty, error)
}

type hubClusterClient struct {
	cc grpc.ClientConnInterface
}

func NewHubClusterClient(cc grpc.ClientConnInterface) HubClusterClient {
	return &hubClusterClient{cc}
}

func (c *hubClusterClient) GetHub(ctx context.Context, in *SelectHub, opts ...grpc.CallOption) (*Hub, error) {
	out := new(Hub)
	err := c.cc.Invoke(ctx, "/HubCluster/GetHub", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *hubClusterClient) CreateHub(ctx context.Context, in *CreateHubRequest, opts ...grpc.CallOption) (*Hub, error) {
	out := new(Hub)
	err := c.cc.Invoke(ctx, "/HubCluster/CreateHub", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *hubClusterClient) TerminateHub(ctx context.Context, in *SelectHub, opts ...grpc.CallOption) (*Empty, error) {
	out := new(Empty)
	err := c.cc.Invoke(ctx, "/HubCluster/TerminateHub", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// HubClusterServer is the server API for HubCluster service.
// All implementations must embed UnimplementedHubClusterServer
// for forward compatibility
type HubClusterServer interface {
	GetHub(context.Context, *SelectHub) (*Hub, error)
	// rpc GetHubs (SelectHubs) return (stream Hubs) {}
	CreateHub(context.Context, *CreateHubRequest) (*Hub, error)
	// rpc PauseHub (SelectHub) returns () {} // what about returning nothing?
	// pause / resume / start / run hub?
	TerminateHub(context.Context, *SelectHub) (*Empty, error)
	mustEmbedUnimplementedHubClusterServer()
}

// UnimplementedHubClusterServer must be embedded to have forward compatible implementations.
type UnimplementedHubClusterServer struct {
}

func (UnimplementedHubClusterServer) GetHub(context.Context, *SelectHub) (*Hub, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetHub not implemented")
}
func (UnimplementedHubClusterServer) CreateHub(context.Context, *CreateHubRequest) (*Hub, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateHub not implemented")
}
func (UnimplementedHubClusterServer) TerminateHub(context.Context, *SelectHub) (*Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method TerminateHub not implemented")
}
func (UnimplementedHubClusterServer) mustEmbedUnimplementedHubClusterServer() {}

// UnsafeHubClusterServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to HubClusterServer will
// result in compilation errors.
type UnsafeHubClusterServer interface {
	mustEmbedUnimplementedHubClusterServer()
}

func RegisterHubClusterServer(s grpc.ServiceRegistrar, srv HubClusterServer) {
	s.RegisterService(&HubCluster_ServiceDesc, srv)
}

func _HubCluster_GetHub_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SelectHub)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(HubClusterServer).GetHub(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/HubCluster/GetHub",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(HubClusterServer).GetHub(ctx, req.(*SelectHub))
	}
	return interceptor(ctx, in, info, handler)
}

func _HubCluster_CreateHub_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateHubRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(HubClusterServer).CreateHub(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/HubCluster/CreateHub",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(HubClusterServer).CreateHub(ctx, req.(*CreateHubRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _HubCluster_TerminateHub_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SelectHub)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(HubClusterServer).TerminateHub(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/HubCluster/TerminateHub",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(HubClusterServer).TerminateHub(ctx, req.(*SelectHub))
	}
	return interceptor(ctx, in, info, handler)
}

// HubCluster_ServiceDesc is the grpc.ServiceDesc for HubCluster service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var HubCluster_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "HubCluster",
	HandlerType: (*HubClusterServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetHub",
			Handler:    _HubCluster_GetHub_Handler,
		},
		{
			MethodName: "CreateHub",
			Handler:    _HubCluster_CreateHub_Handler,
		},
		{
			MethodName: "TerminateHub",
			Handler:    _HubCluster_TerminateHub_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "core.proto",
}