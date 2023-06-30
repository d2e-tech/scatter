// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v4.23.3
// source: pb/scatter.proto

package scatter

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

// ScatterClient is the client API for Scatter service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type ScatterClient interface {
	// Ingests observations
	Ingest(ctx context.Context, in *IngestRequest, opts ...grpc.CallOption) (*IngestReply, error)
}

type scatterClient struct {
	cc grpc.ClientConnInterface
}

func NewScatterClient(cc grpc.ClientConnInterface) ScatterClient {
	return &scatterClient{cc}
}

func (c *scatterClient) Ingest(ctx context.Context, in *IngestRequest, opts ...grpc.CallOption) (*IngestReply, error) {
	out := new(IngestReply)
	err := c.cc.Invoke(ctx, "/pb.Scatter/Ingest", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ScatterServer is the server API for Scatter service.
// All implementations must embed UnimplementedScatterServer
// for forward compatibility
type ScatterServer interface {
	// Ingests observations
	Ingest(context.Context, *IngestRequest) (*IngestReply, error)
	mustEmbedUnimplementedScatterServer()
}

// UnimplementedScatterServer must be embedded to have forward compatible implementations.
type UnimplementedScatterServer struct {
}

func (UnimplementedScatterServer) Ingest(context.Context, *IngestRequest) (*IngestReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Ingest not implemented")
}
func (UnimplementedScatterServer) mustEmbedUnimplementedScatterServer() {}

// UnsafeScatterServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to ScatterServer will
// result in compilation errors.
type UnsafeScatterServer interface {
	mustEmbedUnimplementedScatterServer()
}

func RegisterScatterServer(s grpc.ServiceRegistrar, srv ScatterServer) {
	s.RegisterService(&Scatter_ServiceDesc, srv)
}

func _Scatter_Ingest_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(IngestRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ScatterServer).Ingest(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.Scatter/Ingest",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ScatterServer).Ingest(ctx, req.(*IngestRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// Scatter_ServiceDesc is the grpc.ServiceDesc for Scatter service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Scatter_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "pb.Scatter",
	HandlerType: (*ScatterServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Ingest",
			Handler:    _Scatter_Ingest_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "pb/scatter.proto",
}