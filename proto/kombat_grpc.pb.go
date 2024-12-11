// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.5.1
// - protoc             v5.28.3
// source: proto/kombat.proto

package Kombat

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
	KombatService_GetKombatInfo_FullMethodName     = "/Kombat.KombatService/GetKombatInfo"
	KombatService_GetKombatList_FullMethodName     = "/Kombat.KombatService/GetKombatList"
	KombatService_AddKombats_FullMethodName        = "/Kombat.KombatService/addKombats"
	KombatService_GetKombatByType_FullMethodName   = "/Kombat.KombatService/GetKombatByType"
	KombatService_GetKombatBySkills_FullMethodName = "/Kombat.KombatService/GetKombatBySkills"
)

// KombatServiceClient is the client API for KombatService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type KombatServiceClient interface {
	GetKombatInfo(ctx context.Context, in *KombatRequest, opts ...grpc.CallOption) (*KombatResponse, error)
	GetKombatList(ctx context.Context, in *Empty, opts ...grpc.CallOption) (grpc.ServerStreamingClient[KombatResponse], error)
	AddKombats(ctx context.Context, opts ...grpc.CallOption) (grpc.ClientStreamingClient[NewKombatRequest, AddKombatResponse], error)
	GetKombatByType(ctx context.Context, opts ...grpc.CallOption) (grpc.BidiStreamingClient[KombatTypeRequest, KombatResponse], error)
	GetKombatBySkills(ctx context.Context, opts ...grpc.CallOption) (grpc.BidiStreamingClient[KombatSkillsRequest, KombatResponse], error)
}

type kombatServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewKombatServiceClient(cc grpc.ClientConnInterface) KombatServiceClient {
	return &kombatServiceClient{cc}
}

func (c *kombatServiceClient) GetKombatInfo(ctx context.Context, in *KombatRequest, opts ...grpc.CallOption) (*KombatResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(KombatResponse)
	err := c.cc.Invoke(ctx, KombatService_GetKombatInfo_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *kombatServiceClient) GetKombatList(ctx context.Context, in *Empty, opts ...grpc.CallOption) (grpc.ServerStreamingClient[KombatResponse], error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	stream, err := c.cc.NewStream(ctx, &KombatService_ServiceDesc.Streams[0], KombatService_GetKombatList_FullMethodName, cOpts...)
	if err != nil {
		return nil, err
	}
	x := &grpc.GenericClientStream[Empty, KombatResponse]{ClientStream: stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

// This type alias is provided for backwards compatibility with existing code that references the prior non-generic stream type by name.
type KombatService_GetKombatListClient = grpc.ServerStreamingClient[KombatResponse]

func (c *kombatServiceClient) AddKombats(ctx context.Context, opts ...grpc.CallOption) (grpc.ClientStreamingClient[NewKombatRequest, AddKombatResponse], error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	stream, err := c.cc.NewStream(ctx, &KombatService_ServiceDesc.Streams[1], KombatService_AddKombats_FullMethodName, cOpts...)
	if err != nil {
		return nil, err
	}
	x := &grpc.GenericClientStream[NewKombatRequest, AddKombatResponse]{ClientStream: stream}
	return x, nil
}

// This type alias is provided for backwards compatibility with existing code that references the prior non-generic stream type by name.
type KombatService_AddKombatsClient = grpc.ClientStreamingClient[NewKombatRequest, AddKombatResponse]

func (c *kombatServiceClient) GetKombatByType(ctx context.Context, opts ...grpc.CallOption) (grpc.BidiStreamingClient[KombatTypeRequest, KombatResponse], error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	stream, err := c.cc.NewStream(ctx, &KombatService_ServiceDesc.Streams[2], KombatService_GetKombatByType_FullMethodName, cOpts...)
	if err != nil {
		return nil, err
	}
	x := &grpc.GenericClientStream[KombatTypeRequest, KombatResponse]{ClientStream: stream}
	return x, nil
}

// This type alias is provided for backwards compatibility with existing code that references the prior non-generic stream type by name.
type KombatService_GetKombatByTypeClient = grpc.BidiStreamingClient[KombatTypeRequest, KombatResponse]

func (c *kombatServiceClient) GetKombatBySkills(ctx context.Context, opts ...grpc.CallOption) (grpc.BidiStreamingClient[KombatSkillsRequest, KombatResponse], error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	stream, err := c.cc.NewStream(ctx, &KombatService_ServiceDesc.Streams[3], KombatService_GetKombatBySkills_FullMethodName, cOpts...)
	if err != nil {
		return nil, err
	}
	x := &grpc.GenericClientStream[KombatSkillsRequest, KombatResponse]{ClientStream: stream}
	return x, nil
}

// This type alias is provided for backwards compatibility with existing code that references the prior non-generic stream type by name.
type KombatService_GetKombatBySkillsClient = grpc.BidiStreamingClient[KombatSkillsRequest, KombatResponse]

// KombatServiceServer is the server API for KombatService service.
// All implementations must embed UnimplementedKombatServiceServer
// for forward compatibility.
type KombatServiceServer interface {
	GetKombatInfo(context.Context, *KombatRequest) (*KombatResponse, error)
	GetKombatList(*Empty, grpc.ServerStreamingServer[KombatResponse]) error
	AddKombats(grpc.ClientStreamingServer[NewKombatRequest, AddKombatResponse]) error
	GetKombatByType(grpc.BidiStreamingServer[KombatTypeRequest, KombatResponse]) error
	GetKombatBySkills(grpc.BidiStreamingServer[KombatSkillsRequest, KombatResponse]) error
	mustEmbedUnimplementedKombatServiceServer()
}

// UnimplementedKombatServiceServer must be embedded to have
// forward compatible implementations.
//
// NOTE: this should be embedded by value instead of pointer to avoid a nil
// pointer dereference when methods are called.
type UnimplementedKombatServiceServer struct{}

func (UnimplementedKombatServiceServer) GetKombatInfo(context.Context, *KombatRequest) (*KombatResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetKombatInfo not implemented")
}
func (UnimplementedKombatServiceServer) GetKombatList(*Empty, grpc.ServerStreamingServer[KombatResponse]) error {
	return status.Errorf(codes.Unimplemented, "method GetKombatList not implemented")
}
func (UnimplementedKombatServiceServer) AddKombats(grpc.ClientStreamingServer[NewKombatRequest, AddKombatResponse]) error {
	return status.Errorf(codes.Unimplemented, "method AddKombats not implemented")
}
func (UnimplementedKombatServiceServer) GetKombatByType(grpc.BidiStreamingServer[KombatTypeRequest, KombatResponse]) error {
	return status.Errorf(codes.Unimplemented, "method GetKombatByType not implemented")
}
func (UnimplementedKombatServiceServer) GetKombatBySkills(grpc.BidiStreamingServer[KombatSkillsRequest, KombatResponse]) error {
	return status.Errorf(codes.Unimplemented, "method GetKombatBySkills not implemented")
}
func (UnimplementedKombatServiceServer) mustEmbedUnimplementedKombatServiceServer() {}
func (UnimplementedKombatServiceServer) testEmbeddedByValue()                       {}

// UnsafeKombatServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to KombatServiceServer will
// result in compilation errors.
type UnsafeKombatServiceServer interface {
	mustEmbedUnimplementedKombatServiceServer()
}

func RegisterKombatServiceServer(s grpc.ServiceRegistrar, srv KombatServiceServer) {
	// If the following call pancis, it indicates UnimplementedKombatServiceServer was
	// embedded by pointer and is nil.  This will cause panics if an
	// unimplemented method is ever invoked, so we test this at initialization
	// time to prevent it from happening at runtime later due to I/O.
	if t, ok := srv.(interface{ testEmbeddedByValue() }); ok {
		t.testEmbeddedByValue()
	}
	s.RegisterService(&KombatService_ServiceDesc, srv)
}

func _KombatService_GetKombatInfo_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(KombatRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(KombatServiceServer).GetKombatInfo(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: KombatService_GetKombatInfo_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(KombatServiceServer).GetKombatInfo(ctx, req.(*KombatRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _KombatService_GetKombatList_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(Empty)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(KombatServiceServer).GetKombatList(m, &grpc.GenericServerStream[Empty, KombatResponse]{ServerStream: stream})
}

// This type alias is provided for backwards compatibility with existing code that references the prior non-generic stream type by name.
type KombatService_GetKombatListServer = grpc.ServerStreamingServer[KombatResponse]

func _KombatService_AddKombats_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(KombatServiceServer).AddKombats(&grpc.GenericServerStream[NewKombatRequest, AddKombatResponse]{ServerStream: stream})
}

// This type alias is provided for backwards compatibility with existing code that references the prior non-generic stream type by name.
type KombatService_AddKombatsServer = grpc.ClientStreamingServer[NewKombatRequest, AddKombatResponse]

func _KombatService_GetKombatByType_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(KombatServiceServer).GetKombatByType(&grpc.GenericServerStream[KombatTypeRequest, KombatResponse]{ServerStream: stream})
}

// This type alias is provided for backwards compatibility with existing code that references the prior non-generic stream type by name.
type KombatService_GetKombatByTypeServer = grpc.BidiStreamingServer[KombatTypeRequest, KombatResponse]

func _KombatService_GetKombatBySkills_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(KombatServiceServer).GetKombatBySkills(&grpc.GenericServerStream[KombatSkillsRequest, KombatResponse]{ServerStream: stream})
}

// This type alias is provided for backwards compatibility with existing code that references the prior non-generic stream type by name.
type KombatService_GetKombatBySkillsServer = grpc.BidiStreamingServer[KombatSkillsRequest, KombatResponse]

// KombatService_ServiceDesc is the grpc.ServiceDesc for KombatService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var KombatService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "Kombat.KombatService",
	HandlerType: (*KombatServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetKombatInfo",
			Handler:    _KombatService_GetKombatInfo_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "GetKombatList",
			Handler:       _KombatService_GetKombatList_Handler,
			ServerStreams: true,
		},
		{
			StreamName:    "addKombats",
			Handler:       _KombatService_AddKombats_Handler,
			ClientStreams: true,
		},
		{
			StreamName:    "GetKombatByType",
			Handler:       _KombatService_GetKombatByType_Handler,
			ServerStreams: true,
			ClientStreams: true,
		},
		{
			StreamName:    "GetKombatBySkills",
			Handler:       _KombatService_GetKombatBySkills_Handler,
			ServerStreams: true,
			ClientStreams: true,
		},
	},
	Metadata: "proto/kombat.proto",
}
