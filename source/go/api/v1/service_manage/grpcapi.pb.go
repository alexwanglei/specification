// Code generated by protoc-gen-go. DO NOT EDIT.
// source: grpcapi.proto

package service_manage // import "github.com/polarismesh/specification/source/go/api/v1/service_manage"

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

import (
	context "golang.org/x/net/context"
	grpc "google.golang.org/grpc"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// PolarisGRPCClient is the client API for PolarisGRPC service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type PolarisGRPCClient interface {
	// 客户端上报
	ReportClient(ctx context.Context, in *Client, opts ...grpc.CallOption) (*Response, error)
	// 被调方注册服务实例
	RegisterInstance(ctx context.Context, in *Instance, opts ...grpc.CallOption) (*Response, error)
	// 被调方反注册服务实例
	DeregisterInstance(ctx context.Context, in *Instance, opts ...grpc.CallOption) (*Response, error)
	// 统一发现接口
	Discover(ctx context.Context, opts ...grpc.CallOption) (PolarisGRPC_DiscoverClient, error)
	// 被调方上报心跳
	Heartbeat(ctx context.Context, in *Instance, opts ...grpc.CallOption) (*Response, error)
}

type polarisGRPCClient struct {
	cc *grpc.ClientConn
}

func NewPolarisGRPCClient(cc *grpc.ClientConn) PolarisGRPCClient {
	return &polarisGRPCClient{cc}
}

func (c *polarisGRPCClient) ReportClient(ctx context.Context, in *Client, opts ...grpc.CallOption) (*Response, error) {
	out := new(Response)
	err := c.cc.Invoke(ctx, "/v1.PolarisGRPC/ReportClient", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *polarisGRPCClient) RegisterInstance(ctx context.Context, in *Instance, opts ...grpc.CallOption) (*Response, error) {
	out := new(Response)
	err := c.cc.Invoke(ctx, "/v1.PolarisGRPC/RegisterInstance", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *polarisGRPCClient) DeregisterInstance(ctx context.Context, in *Instance, opts ...grpc.CallOption) (*Response, error) {
	out := new(Response)
	err := c.cc.Invoke(ctx, "/v1.PolarisGRPC/DeregisterInstance", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *polarisGRPCClient) Discover(ctx context.Context, opts ...grpc.CallOption) (PolarisGRPC_DiscoverClient, error) {
	stream, err := c.cc.NewStream(ctx, &_PolarisGRPC_serviceDesc.Streams[0], "/v1.PolarisGRPC/Discover", opts...)
	if err != nil {
		return nil, err
	}
	x := &polarisGRPCDiscoverClient{stream}
	return x, nil
}

type PolarisGRPC_DiscoverClient interface {
	Send(*DiscoverRequest) error
	Recv() (*DiscoverResponse, error)
	grpc.ClientStream
}

type polarisGRPCDiscoverClient struct {
	grpc.ClientStream
}

func (x *polarisGRPCDiscoverClient) Send(m *DiscoverRequest) error {
	return x.ClientStream.SendMsg(m)
}

func (x *polarisGRPCDiscoverClient) Recv() (*DiscoverResponse, error) {
	m := new(DiscoverResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *polarisGRPCClient) Heartbeat(ctx context.Context, in *Instance, opts ...grpc.CallOption) (*Response, error) {
	out := new(Response)
	err := c.cc.Invoke(ctx, "/v1.PolarisGRPC/Heartbeat", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// PolarisGRPCServer is the server API for PolarisGRPC service.
type PolarisGRPCServer interface {
	// 客户端上报
	ReportClient(context.Context, *Client) (*Response, error)
	// 被调方注册服务实例
	RegisterInstance(context.Context, *Instance) (*Response, error)
	// 被调方反注册服务实例
	DeregisterInstance(context.Context, *Instance) (*Response, error)
	// 统一发现接口
	Discover(PolarisGRPC_DiscoverServer) error
	// 被调方上报心跳
	Heartbeat(context.Context, *Instance) (*Response, error)
}

func RegisterPolarisGRPCServer(s *grpc.Server, srv PolarisGRPCServer) {
	s.RegisterService(&_PolarisGRPC_serviceDesc, srv)
}

func _PolarisGRPC_ReportClient_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Client)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PolarisGRPCServer).ReportClient(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/v1.PolarisGRPC/ReportClient",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PolarisGRPCServer).ReportClient(ctx, req.(*Client))
	}
	return interceptor(ctx, in, info, handler)
}

func _PolarisGRPC_RegisterInstance_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Instance)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PolarisGRPCServer).RegisterInstance(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/v1.PolarisGRPC/RegisterInstance",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PolarisGRPCServer).RegisterInstance(ctx, req.(*Instance))
	}
	return interceptor(ctx, in, info, handler)
}

func _PolarisGRPC_DeregisterInstance_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Instance)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PolarisGRPCServer).DeregisterInstance(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/v1.PolarisGRPC/DeregisterInstance",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PolarisGRPCServer).DeregisterInstance(ctx, req.(*Instance))
	}
	return interceptor(ctx, in, info, handler)
}

func _PolarisGRPC_Discover_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(PolarisGRPCServer).Discover(&polarisGRPCDiscoverServer{stream})
}

type PolarisGRPC_DiscoverServer interface {
	Send(*DiscoverResponse) error
	Recv() (*DiscoverRequest, error)
	grpc.ServerStream
}

type polarisGRPCDiscoverServer struct {
	grpc.ServerStream
}

func (x *polarisGRPCDiscoverServer) Send(m *DiscoverResponse) error {
	return x.ServerStream.SendMsg(m)
}

func (x *polarisGRPCDiscoverServer) Recv() (*DiscoverRequest, error) {
	m := new(DiscoverRequest)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func _PolarisGRPC_Heartbeat_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Instance)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PolarisGRPCServer).Heartbeat(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/v1.PolarisGRPC/Heartbeat",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PolarisGRPCServer).Heartbeat(ctx, req.(*Instance))
	}
	return interceptor(ctx, in, info, handler)
}

var _PolarisGRPC_serviceDesc = grpc.ServiceDesc{
	ServiceName: "v1.PolarisGRPC",
	HandlerType: (*PolarisGRPCServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "ReportClient",
			Handler:    _PolarisGRPC_ReportClient_Handler,
		},
		{
			MethodName: "RegisterInstance",
			Handler:    _PolarisGRPC_RegisterInstance_Handler,
		},
		{
			MethodName: "DeregisterInstance",
			Handler:    _PolarisGRPC_DeregisterInstance_Handler,
		},
		{
			MethodName: "Heartbeat",
			Handler:    _PolarisGRPC_Heartbeat_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "Discover",
			Handler:       _PolarisGRPC_Discover_Handler,
			ServerStreams: true,
			ClientStreams: true,
		},
	},
	Metadata: "grpcapi.proto",
}

func init() { proto.RegisterFile("grpcapi.proto", fileDescriptor_grpcapi_76e09f6be072a541) }

var fileDescriptor_grpcapi_76e09f6be072a541 = []byte{
	// 289 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x94, 0x91, 0x31, 0x4f, 0xc3, 0x30,
	0x10, 0x85, 0x49, 0x07, 0x04, 0x26, 0x45, 0xc8, 0x30, 0x65, 0x64, 0x02, 0x06, 0xbb, 0x29, 0x03,
	0x03, 0x5b, 0x1b, 0x09, 0xd8, 0xaa, 0xb0, 0xb1, 0x20, 0xc7, 0x3a, 0x52, 0x4b, 0x8d, 0x6d, 0xce,
	0x8e, 0xff, 0x02, 0x3f, 0x80, 0x3f, 0x8c, 0x12, 0xc7, 0x12, 0x15, 0x12, 0x52, 0xb7, 0x7b, 0xf7,
	0xfc, 0x9e, 0xe5, 0xcf, 0x64, 0xde, 0xa2, 0x95, 0xc2, 0x2a, 0x66, 0xd1, 0x78, 0x43, 0x67, 0xa1,
	0x2c, 0x72, 0xb9, 0x53, 0xa0, 0x7d, 0xdc, 0x14, 0x73, 0x07, 0x18, 0x94, 0x84, 0x24, 0x11, 0x3e,
	0x7b, 0x70, 0xc9, 0x3d, 0x47, 0x70, 0xd6, 0x68, 0x37, 0xd9, 0xcb, 0xaf, 0x19, 0x39, 0xdb, 0x98,
	0x9d, 0x40, 0xe5, 0x9e, 0xea, 0xcd, 0x9a, 0xde, 0x91, 0xbc, 0x06, 0x6b, 0xd0, 0xaf, 0xc7, 0x4e,
	0x4a, 0x58, 0x28, 0x59, 0x9c, 0x8b, 0x7c, 0x98, 0xeb, 0x29, 0x7f, 0x7d, 0x44, 0x17, 0xe4, 0xa2,
	0x86, 0x56, 0x39, 0x0f, 0xf8, 0xa2, 0x9d, 0x17, 0x5a, 0x02, 0x1d, 0xcf, 0x24, 0xf5, 0x27, 0xb1,
	0x24, 0xb4, 0x02, 0x3c, 0x2c, 0xf3, 0x48, 0x4e, 0x2a, 0xe5, 0xa4, 0x09, 0x80, 0xf4, 0x72, 0xf0,
	0x92, 0xaa, 0xe3, 0xc3, 0x8a, 0xab, 0xfd, 0x65, 0x0a, 0xde, 0x64, 0x8b, 0x8c, 0xde, 0x92, 0xd3,
	0x67, 0x10, 0xe8, 0x1b, 0x10, 0xfe, 0xff, 0x7b, 0x56, 0xdf, 0x19, 0x79, 0x90, 0xa6, 0x63, 0x1e,
	0xb4, 0x1c, 0x69, 0x46, 0x2a, 0xcc, 0x59, 0x90, 0xea, 0x43, 0x49, 0xe1, 0x95, 0xd1, 0x6c, 0x20,
	0x1f, 0x4a, 0x96, 0x18, 0x77, 0x42, 0x8b, 0x16, 0x56, 0xf4, 0x17, 0xc2, 0xd7, 0x68, 0xbd, 0x55,
	0xad, 0xf2, 0xdb, 0xbe, 0x61, 0xd2, 0x74, 0x7c, 0xea, 0xea, 0xc0, 0x6d, 0xf9, 0x5e, 0x1f, 0x77,
	0xa6, 0x47, 0x09, 0xbc, 0x35, 0x5c, 0x58, 0xc5, 0x43, 0xc9, 0xa7, 0xe6, 0xf7, 0xd8, 0xdc, 0x1c,
	0x8f, 0xdf, 0x74, 0xff, 0x13, 0x00, 0x00, 0xff, 0xff, 0x17, 0x3f, 0xdb, 0xbd, 0xf7, 0x01, 0x00,
	0x00,
}
