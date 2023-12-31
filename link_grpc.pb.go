// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v3.21.9
// source: com/itcoursee/message/link.proto

package message

import (
	context "context"
	link "github.com/yangzone/core/link"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

const (
	Link_Paging_FullMethodName = "/com.itcoursee.message.Link/Paging"
)

// LinkClient is the client API for Link service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type LinkClient interface {
	// 分页
	// protolint:disable:next MAX_LINE_LENGTH
	Paging(ctx context.Context, in *link.PagingReq, opts ...grpc.CallOption) (*link.PagingRsp, error)
}

type linkClient struct {
	cc grpc.ClientConnInterface
}

func NewLinkClient(cc grpc.ClientConnInterface) LinkClient {
	return &linkClient{cc}
}

func (c *linkClient) Paging(ctx context.Context, in *link.PagingReq, opts ...grpc.CallOption) (*link.PagingRsp, error) {
	out := new(link.PagingRsp)
	err := c.cc.Invoke(ctx, Link_Paging_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// LinkServer is the server API for Link service.
// All implementations must embed UnimplementedLinkServer
// for forward compatibility
type LinkServer interface {
	// 分页
	// protolint:disable:next MAX_LINE_LENGTH
	Paging(context.Context, *link.PagingReq) (*link.PagingRsp, error)
	mustEmbedUnimplementedLinkServer()
}

// UnimplementedLinkServer must be embedded to have forward compatible implementations.
type UnimplementedLinkServer struct {
}

func (UnimplementedLinkServer) Paging(context.Context, *link.PagingReq) (*link.PagingRsp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Paging not implemented")
}
func (UnimplementedLinkServer) mustEmbedUnimplementedLinkServer() {}

// UnsafeLinkServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to LinkServer will
// result in compilation errors.
type UnsafeLinkServer interface {
	mustEmbedUnimplementedLinkServer()
}

func RegisterLinkServer(s grpc.ServiceRegistrar, srv LinkServer) {
	s.RegisterService(&Link_ServiceDesc, srv)
}

func _Link_Paging_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(link.PagingReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(LinkServer).Paging(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Link_Paging_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(LinkServer).Paging(ctx, req.(*link.PagingReq))
	}
	return interceptor(ctx, in, info, handler)
}

// Link_ServiceDesc is the grpc.ServiceDesc for Link service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Link_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "com.itcoursee.message.Link",
	HandlerType: (*LinkServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Paging",
			Handler:    _Link_Paging_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "com/itcoursee/message/link.proto",
}
