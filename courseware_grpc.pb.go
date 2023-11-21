// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v3.21.9
// source: com/itcoursee/message/courseware.proto

package message

import (
	context "context"
	courseware "github.com/itcoursee/core/courseware"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

const (
	Courseware_AddByCourse_FullMethodName   = "/com.itcoursee.message.Courseware/AddByCourse"
	Courseware_AddByLecture_FullMethodName  = "/com.itcoursee.message.Courseware/AddByLecture"
	Courseware_Get_FullMethodName           = "/com.itcoursee.message.Courseware/Get"
	Courseware_GetsByCourse_FullMethodName  = "/com.itcoursee.message.Courseware/GetsByCourse"
	Courseware_GetsByLecture_FullMethodName = "/com.itcoursee.message.Courseware/GetsByLecture"
	Courseware_Update_FullMethodName        = "/com.itcoursee.message.Courseware/Update"
	Courseware_Delete_FullMethodName        = "/com.itcoursee.message.Courseware/Delete"
)

// CoursewareClient is the client API for Courseware service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type CoursewareClient interface {
	// 创建课程课件
	// protolint:disable:next MAX_LINE_LENGTH
	AddByCourse(ctx context.Context, in *courseware.AddByCourseReq, opts ...grpc.CallOption) (*courseware.AddRsp, error)
	// 创建讲次课件
	// protolint:disable:next MAX_LINE_LENGTH
	AddByLecture(ctx context.Context, in *courseware.AddByLectureReq, opts ...grpc.CallOption) (*courseware.AddRsp, error)
	// 获取
	Get(ctx context.Context, in *courseware.GetReq, opts ...grpc.CallOption) (*courseware.GetRsp, error)
	// 获得课程所有的课件
	// protolint:disable:next MAX_LINE_LENGTH
	GetsByCourse(ctx context.Context, in *courseware.GetsByCourseReq, opts ...grpc.CallOption) (*courseware.GetsByOwnerRsp, error)
	// 获得讲次所有的课件
	// protolint:disable:next MAX_LINE_LENGTH
	GetsByLecture(ctx context.Context, in *courseware.GetsByLectureReq, opts ...grpc.CallOption) (*courseware.GetsByOwnerRsp, error)
	// 修改
	Update(ctx context.Context, in *courseware.UpdateReq, opts ...grpc.CallOption) (*courseware.UpdateRsp, error)
	// 删除
	Delete(ctx context.Context, in *courseware.DeleteReq, opts ...grpc.CallOption) (*courseware.DeleteRsp, error)
}

type coursewareClient struct {
	cc grpc.ClientConnInterface
}

func NewCoursewareClient(cc grpc.ClientConnInterface) CoursewareClient {
	return &coursewareClient{cc}
}

func (c *coursewareClient) AddByCourse(ctx context.Context, in *courseware.AddByCourseReq, opts ...grpc.CallOption) (*courseware.AddRsp, error) {
	out := new(courseware.AddRsp)
	err := c.cc.Invoke(ctx, Courseware_AddByCourse_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *coursewareClient) AddByLecture(ctx context.Context, in *courseware.AddByLectureReq, opts ...grpc.CallOption) (*courseware.AddRsp, error) {
	out := new(courseware.AddRsp)
	err := c.cc.Invoke(ctx, Courseware_AddByLecture_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *coursewareClient) Get(ctx context.Context, in *courseware.GetReq, opts ...grpc.CallOption) (*courseware.GetRsp, error) {
	out := new(courseware.GetRsp)
	err := c.cc.Invoke(ctx, Courseware_Get_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *coursewareClient) GetsByCourse(ctx context.Context, in *courseware.GetsByCourseReq, opts ...grpc.CallOption) (*courseware.GetsByOwnerRsp, error) {
	out := new(courseware.GetsByOwnerRsp)
	err := c.cc.Invoke(ctx, Courseware_GetsByCourse_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *coursewareClient) GetsByLecture(ctx context.Context, in *courseware.GetsByLectureReq, opts ...grpc.CallOption) (*courseware.GetsByOwnerRsp, error) {
	out := new(courseware.GetsByOwnerRsp)
	err := c.cc.Invoke(ctx, Courseware_GetsByLecture_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *coursewareClient) Update(ctx context.Context, in *courseware.UpdateReq, opts ...grpc.CallOption) (*courseware.UpdateRsp, error) {
	out := new(courseware.UpdateRsp)
	err := c.cc.Invoke(ctx, Courseware_Update_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *coursewareClient) Delete(ctx context.Context, in *courseware.DeleteReq, opts ...grpc.CallOption) (*courseware.DeleteRsp, error) {
	out := new(courseware.DeleteRsp)
	err := c.cc.Invoke(ctx, Courseware_Delete_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// CoursewareServer is the server API for Courseware service.
// All implementations must embed UnimplementedCoursewareServer
// for forward compatibility
type CoursewareServer interface {
	// 创建课程课件
	// protolint:disable:next MAX_LINE_LENGTH
	AddByCourse(context.Context, *courseware.AddByCourseReq) (*courseware.AddRsp, error)
	// 创建讲次课件
	// protolint:disable:next MAX_LINE_LENGTH
	AddByLecture(context.Context, *courseware.AddByLectureReq) (*courseware.AddRsp, error)
	// 获取
	Get(context.Context, *courseware.GetReq) (*courseware.GetRsp, error)
	// 获得课程所有的课件
	// protolint:disable:next MAX_LINE_LENGTH
	GetsByCourse(context.Context, *courseware.GetsByCourseReq) (*courseware.GetsByOwnerRsp, error)
	// 获得讲次所有的课件
	// protolint:disable:next MAX_LINE_LENGTH
	GetsByLecture(context.Context, *courseware.GetsByLectureReq) (*courseware.GetsByOwnerRsp, error)
	// 修改
	Update(context.Context, *courseware.UpdateReq) (*courseware.UpdateRsp, error)
	// 删除
	Delete(context.Context, *courseware.DeleteReq) (*courseware.DeleteRsp, error)
	mustEmbedUnimplementedCoursewareServer()
}

// UnimplementedCoursewareServer must be embedded to have forward compatible implementations.
type UnimplementedCoursewareServer struct {
}

func (UnimplementedCoursewareServer) AddByCourse(context.Context, *courseware.AddByCourseReq) (*courseware.AddRsp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AddByCourse not implemented")
}
func (UnimplementedCoursewareServer) AddByLecture(context.Context, *courseware.AddByLectureReq) (*courseware.AddRsp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AddByLecture not implemented")
}
func (UnimplementedCoursewareServer) Get(context.Context, *courseware.GetReq) (*courseware.GetRsp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Get not implemented")
}
func (UnimplementedCoursewareServer) GetsByCourse(context.Context, *courseware.GetsByCourseReq) (*courseware.GetsByOwnerRsp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetsByCourse not implemented")
}
func (UnimplementedCoursewareServer) GetsByLecture(context.Context, *courseware.GetsByLectureReq) (*courseware.GetsByOwnerRsp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetsByLecture not implemented")
}
func (UnimplementedCoursewareServer) Update(context.Context, *courseware.UpdateReq) (*courseware.UpdateRsp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Update not implemented")
}
func (UnimplementedCoursewareServer) Delete(context.Context, *courseware.DeleteReq) (*courseware.DeleteRsp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Delete not implemented")
}
func (UnimplementedCoursewareServer) mustEmbedUnimplementedCoursewareServer() {}

// UnsafeCoursewareServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to CoursewareServer will
// result in compilation errors.
type UnsafeCoursewareServer interface {
	mustEmbedUnimplementedCoursewareServer()
}

func RegisterCoursewareServer(s grpc.ServiceRegistrar, srv CoursewareServer) {
	s.RegisterService(&Courseware_ServiceDesc, srv)
}

func _Courseware_AddByCourse_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(courseware.AddByCourseReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CoursewareServer).AddByCourse(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Courseware_AddByCourse_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CoursewareServer).AddByCourse(ctx, req.(*courseware.AddByCourseReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Courseware_AddByLecture_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(courseware.AddByLectureReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CoursewareServer).AddByLecture(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Courseware_AddByLecture_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CoursewareServer).AddByLecture(ctx, req.(*courseware.AddByLectureReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Courseware_Get_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(courseware.GetReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CoursewareServer).Get(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Courseware_Get_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CoursewareServer).Get(ctx, req.(*courseware.GetReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Courseware_GetsByCourse_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(courseware.GetsByCourseReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CoursewareServer).GetsByCourse(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Courseware_GetsByCourse_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CoursewareServer).GetsByCourse(ctx, req.(*courseware.GetsByCourseReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Courseware_GetsByLecture_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(courseware.GetsByLectureReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CoursewareServer).GetsByLecture(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Courseware_GetsByLecture_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CoursewareServer).GetsByLecture(ctx, req.(*courseware.GetsByLectureReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Courseware_Update_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(courseware.UpdateReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CoursewareServer).Update(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Courseware_Update_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CoursewareServer).Update(ctx, req.(*courseware.UpdateReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Courseware_Delete_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(courseware.DeleteReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CoursewareServer).Delete(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Courseware_Delete_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CoursewareServer).Delete(ctx, req.(*courseware.DeleteReq))
	}
	return interceptor(ctx, in, info, handler)
}

// Courseware_ServiceDesc is the grpc.ServiceDesc for Courseware service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Courseware_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "com.itcoursee.message.Courseware",
	HandlerType: (*CoursewareServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "AddByCourse",
			Handler:    _Courseware_AddByCourse_Handler,
		},
		{
			MethodName: "AddByLecture",
			Handler:    _Courseware_AddByLecture_Handler,
		},
		{
			MethodName: "Get",
			Handler:    _Courseware_Get_Handler,
		},
		{
			MethodName: "GetsByCourse",
			Handler:    _Courseware_GetsByCourse_Handler,
		},
		{
			MethodName: "GetsByLecture",
			Handler:    _Courseware_GetsByLecture_Handler,
		},
		{
			MethodName: "Update",
			Handler:    _Courseware_Update_Handler,
		},
		{
			MethodName: "Delete",
			Handler:    _Courseware_Delete_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "com/itcoursee/message/courseware.proto",
}
