// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.18.1
// source: com/itcoursee/message/courseware.proto

package protocol

import (
	courseware "github.com/itcoursee/core/courseware"
	_ "google.golang.org/genproto/googleapis/api/annotations"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

var File_com_itcoursee_message_courseware_proto protoreflect.FileDescriptor

var file_com_itcoursee_message_courseware_proto_rawDesc = []byte{
	0x0a, 0x26, 0x63, 0x6f, 0x6d, 0x2f, 0x69, 0x74, 0x63, 0x6f, 0x75, 0x72, 0x73, 0x65, 0x65, 0x2f,
	0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x2f, 0x63, 0x6f, 0x75, 0x72, 0x73, 0x65, 0x77, 0x61,
	0x72, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x15, 0x63, 0x6f, 0x6d, 0x2e, 0x69, 0x74,
	0x63, 0x6f, 0x75, 0x72, 0x73, 0x65, 0x65, 0x2e, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x1a,
	0x1c, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x61, 0x6e, 0x6e, 0x6f,
	0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x22, 0x63,
	0x6f, 0x6d, 0x2f, 0x69, 0x74, 0x63, 0x6f, 0x75, 0x72, 0x73, 0x65, 0x65, 0x2f, 0x63, 0x6f, 0x75,
	0x72, 0x73, 0x65, 0x77, 0x61, 0x72, 0x65, 0x2f, 0x61, 0x64, 0x64, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x1a, 0x25, 0x63, 0x6f, 0x6d, 0x2f, 0x69, 0x74, 0x63, 0x6f, 0x75, 0x72, 0x73, 0x65, 0x65,
	0x2f, 0x63, 0x6f, 0x75, 0x72, 0x73, 0x65, 0x77, 0x61, 0x72, 0x65, 0x2f, 0x63, 0x6f, 0x75, 0x72,
	0x73, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x25, 0x63, 0x6f, 0x6d, 0x2f, 0x69, 0x74,
	0x63, 0x6f, 0x75, 0x72, 0x73, 0x65, 0x65, 0x2f, 0x63, 0x6f, 0x75, 0x72, 0x73, 0x65, 0x77, 0x61,
	0x72, 0x65, 0x2f, 0x64, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a,
	0x22, 0x63, 0x6f, 0x6d, 0x2f, 0x69, 0x74, 0x63, 0x6f, 0x75, 0x72, 0x73, 0x65, 0x65, 0x2f, 0x63,
	0x6f, 0x75, 0x72, 0x73, 0x65, 0x77, 0x61, 0x72, 0x65, 0x2f, 0x67, 0x65, 0x74, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x1a, 0x26, 0x63, 0x6f, 0x6d, 0x2f, 0x69, 0x74, 0x63, 0x6f, 0x75, 0x72, 0x73,
	0x65, 0x65, 0x2f, 0x63, 0x6f, 0x75, 0x72, 0x73, 0x65, 0x77, 0x61, 0x72, 0x65, 0x2f, 0x6c, 0x65,
	0x63, 0x74, 0x75, 0x72, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x24, 0x63, 0x6f, 0x6d,
	0x2f, 0x69, 0x74, 0x63, 0x6f, 0x75, 0x72, 0x73, 0x65, 0x65, 0x2f, 0x63, 0x6f, 0x75, 0x72, 0x73,
	0x65, 0x77, 0x61, 0x72, 0x65, 0x2f, 0x6f, 0x77, 0x6e, 0x65, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x1a, 0x25, 0x63, 0x6f, 0x6d, 0x2f, 0x69, 0x74, 0x63, 0x6f, 0x75, 0x72, 0x73, 0x65, 0x65,
	0x2f, 0x63, 0x6f, 0x75, 0x72, 0x73, 0x65, 0x77, 0x61, 0x72, 0x65, 0x2f, 0x75, 0x70, 0x64, 0x61,
	0x74, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x32, 0xed, 0x07, 0x0a, 0x0a, 0x43, 0x6f, 0x75,
	0x72, 0x73, 0x65, 0x77, 0x61, 0x72, 0x65, 0x12, 0x8f, 0x01, 0x0a, 0x0b, 0x41, 0x64, 0x64, 0x42,
	0x79, 0x43, 0x6f, 0x75, 0x72, 0x73, 0x65, 0x12, 0x28, 0x2e, 0x63, 0x6f, 0x6d, 0x2e, 0x69, 0x74,
	0x63, 0x6f, 0x75, 0x72, 0x73, 0x65, 0x65, 0x2e, 0x63, 0x6f, 0x75, 0x72, 0x73, 0x65, 0x77, 0x61,
	0x72, 0x65, 0x2e, 0x41, 0x64, 0x64, 0x42, 0x79, 0x43, 0x6f, 0x75, 0x72, 0x73, 0x65, 0x52, 0x65,
	0x71, 0x1a, 0x20, 0x2e, 0x63, 0x6f, 0x6d, 0x2e, 0x69, 0x74, 0x63, 0x6f, 0x75, 0x72, 0x73, 0x65,
	0x65, 0x2e, 0x63, 0x6f, 0x75, 0x72, 0x73, 0x65, 0x77, 0x61, 0x72, 0x65, 0x2e, 0x41, 0x64, 0x64,
	0x52, 0x73, 0x70, 0x22, 0x34, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x2e, 0x22, 0x1d, 0x2f, 0x63, 0x6f,
	0x75, 0x72, 0x73, 0x65, 0x73, 0x2f, 0x7b, 0x63, 0x6f, 0x75, 0x72, 0x73, 0x65, 0x7d, 0x2f, 0x63,
	0x6f, 0x75, 0x72, 0x73, 0x65, 0x77, 0x61, 0x72, 0x65, 0x73, 0x3a, 0x01, 0x2a, 0x62, 0x0a, 0x63,
	0x6f, 0x75, 0x72, 0x73, 0x65, 0x77, 0x61, 0x72, 0x65, 0x12, 0xa4, 0x01, 0x0a, 0x0c, 0x41, 0x64,
	0x64, 0x42, 0x79, 0x4c, 0x65, 0x63, 0x74, 0x75, 0x72, 0x65, 0x12, 0x29, 0x2e, 0x63, 0x6f, 0x6d,
	0x2e, 0x69, 0x74, 0x63, 0x6f, 0x75, 0x72, 0x73, 0x65, 0x65, 0x2e, 0x63, 0x6f, 0x75, 0x72, 0x73,
	0x65, 0x77, 0x61, 0x72, 0x65, 0x2e, 0x41, 0x64, 0x64, 0x42, 0x79, 0x4c, 0x65, 0x63, 0x74, 0x75,
	0x72, 0x65, 0x52, 0x65, 0x71, 0x1a, 0x20, 0x2e, 0x63, 0x6f, 0x6d, 0x2e, 0x69, 0x74, 0x63, 0x6f,
	0x75, 0x72, 0x73, 0x65, 0x65, 0x2e, 0x63, 0x6f, 0x75, 0x72, 0x73, 0x65, 0x77, 0x61, 0x72, 0x65,
	0x2e, 0x41, 0x64, 0x64, 0x52, 0x73, 0x70, 0x22, 0x47, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x41, 0x22,
	0x30, 0x2f, 0x63, 0x6f, 0x75, 0x72, 0x73, 0x65, 0x73, 0x2f, 0x7b, 0x63, 0x6f, 0x75, 0x72, 0x73,
	0x65, 0x7d, 0x2f, 0x6c, 0x65, 0x63, 0x74, 0x75, 0x72, 0x65, 0x73, 0x2f, 0x7b, 0x6c, 0x65, 0x63,
	0x74, 0x75, 0x72, 0x65, 0x7d, 0x2f, 0x63, 0x6f, 0x75, 0x72, 0x73, 0x65, 0x77, 0x61, 0x72, 0x65,
	0x73, 0x3a, 0x01, 0x2a, 0x62, 0x0a, 0x63, 0x6f, 0x75, 0x72, 0x73, 0x65, 0x77, 0x61, 0x72, 0x65,
	0x12, 0x70, 0x0a, 0x03, 0x47, 0x65, 0x74, 0x12, 0x20, 0x2e, 0x63, 0x6f, 0x6d, 0x2e, 0x69, 0x74,
	0x63, 0x6f, 0x75, 0x72, 0x73, 0x65, 0x65, 0x2e, 0x63, 0x6f, 0x75, 0x72, 0x73, 0x65, 0x77, 0x61,
	0x72, 0x65, 0x2e, 0x47, 0x65, 0x74, 0x52, 0x65, 0x71, 0x1a, 0x20, 0x2e, 0x63, 0x6f, 0x6d, 0x2e,
	0x69, 0x74, 0x63, 0x6f, 0x75, 0x72, 0x73, 0x65, 0x65, 0x2e, 0x63, 0x6f, 0x75, 0x72, 0x73, 0x65,
	0x77, 0x61, 0x72, 0x65, 0x2e, 0x47, 0x65, 0x74, 0x52, 0x73, 0x70, 0x22, 0x25, 0x82, 0xd3, 0xe4,
	0x93, 0x02, 0x1f, 0x12, 0x11, 0x2f, 0x63, 0x6f, 0x75, 0x72, 0x73, 0x65, 0x77, 0x61, 0x72, 0x65,
	0x73, 0x2f, 0x7b, 0x69, 0x64, 0x7d, 0x62, 0x0a, 0x63, 0x6f, 0x75, 0x72, 0x73, 0x65, 0x77, 0x61,
	0x72, 0x65, 0x12, 0x97, 0x01, 0x0a, 0x0c, 0x47, 0x65, 0x74, 0x73, 0x42, 0x79, 0x43, 0x6f, 0x75,
	0x72, 0x73, 0x65, 0x12, 0x29, 0x2e, 0x63, 0x6f, 0x6d, 0x2e, 0x69, 0x74, 0x63, 0x6f, 0x75, 0x72,
	0x73, 0x65, 0x65, 0x2e, 0x63, 0x6f, 0x75, 0x72, 0x73, 0x65, 0x77, 0x61, 0x72, 0x65, 0x2e, 0x47,
	0x65, 0x74, 0x73, 0x42, 0x79, 0x43, 0x6f, 0x75, 0x72, 0x73, 0x65, 0x52, 0x65, 0x71, 0x1a, 0x28,
	0x2e, 0x63, 0x6f, 0x6d, 0x2e, 0x69, 0x74, 0x63, 0x6f, 0x75, 0x72, 0x73, 0x65, 0x65, 0x2e, 0x63,
	0x6f, 0x75, 0x72, 0x73, 0x65, 0x77, 0x61, 0x72, 0x65, 0x2e, 0x47, 0x65, 0x74, 0x73, 0x42, 0x79,
	0x4f, 0x77, 0x6e, 0x65, 0x72, 0x52, 0x73, 0x70, 0x22, 0x32, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x2c,
	0x12, 0x1d, 0x2f, 0x63, 0x6f, 0x75, 0x72, 0x73, 0x65, 0x73, 0x2f, 0x7b, 0x63, 0x6f, 0x75, 0x72,
	0x73, 0x65, 0x7d, 0x2f, 0x63, 0x6f, 0x75, 0x72, 0x73, 0x65, 0x77, 0x61, 0x72, 0x65, 0x73, 0x62,
	0x0b, 0x63, 0x6f, 0x75, 0x72, 0x73, 0x65, 0x77, 0x61, 0x72, 0x65, 0x73, 0x12, 0xac, 0x01, 0x0a,
	0x0d, 0x47, 0x65, 0x74, 0x73, 0x42, 0x79, 0x4c, 0x65, 0x63, 0x74, 0x75, 0x72, 0x65, 0x12, 0x2a,
	0x2e, 0x63, 0x6f, 0x6d, 0x2e, 0x69, 0x74, 0x63, 0x6f, 0x75, 0x72, 0x73, 0x65, 0x65, 0x2e, 0x63,
	0x6f, 0x75, 0x72, 0x73, 0x65, 0x77, 0x61, 0x72, 0x65, 0x2e, 0x47, 0x65, 0x74, 0x73, 0x42, 0x79,
	0x4c, 0x65, 0x63, 0x74, 0x75, 0x72, 0x65, 0x52, 0x65, 0x71, 0x1a, 0x28, 0x2e, 0x63, 0x6f, 0x6d,
	0x2e, 0x69, 0x74, 0x63, 0x6f, 0x75, 0x72, 0x73, 0x65, 0x65, 0x2e, 0x63, 0x6f, 0x75, 0x72, 0x73,
	0x65, 0x77, 0x61, 0x72, 0x65, 0x2e, 0x47, 0x65, 0x74, 0x73, 0x42, 0x79, 0x4f, 0x77, 0x6e, 0x65,
	0x72, 0x52, 0x73, 0x70, 0x22, 0x45, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x3f, 0x12, 0x30, 0x2f, 0x63,
	0x6f, 0x75, 0x72, 0x73, 0x65, 0x73, 0x2f, 0x7b, 0x63, 0x6f, 0x75, 0x72, 0x73, 0x65, 0x7d, 0x2f,
	0x6c, 0x65, 0x63, 0x74, 0x75, 0x72, 0x65, 0x73, 0x2f, 0x7b, 0x6c, 0x65, 0x63, 0x74, 0x75, 0x72,
	0x65, 0x7d, 0x2f, 0x63, 0x6f, 0x75, 0x72, 0x73, 0x65, 0x77, 0x61, 0x72, 0x65, 0x73, 0x62, 0x0b,
	0x63, 0x6f, 0x75, 0x72, 0x73, 0x65, 0x77, 0x61, 0x72, 0x65, 0x73, 0x12, 0x7c, 0x0a, 0x06, 0x55,
	0x70, 0x64, 0x61, 0x74, 0x65, 0x12, 0x23, 0x2e, 0x63, 0x6f, 0x6d, 0x2e, 0x69, 0x74, 0x63, 0x6f,
	0x75, 0x72, 0x73, 0x65, 0x65, 0x2e, 0x63, 0x6f, 0x75, 0x72, 0x73, 0x65, 0x77, 0x61, 0x72, 0x65,
	0x2e, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x52, 0x65, 0x71, 0x1a, 0x23, 0x2e, 0x63, 0x6f, 0x6d,
	0x2e, 0x69, 0x74, 0x63, 0x6f, 0x75, 0x72, 0x73, 0x65, 0x65, 0x2e, 0x63, 0x6f, 0x75, 0x72, 0x73,
	0x65, 0x77, 0x61, 0x72, 0x65, 0x2e, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x52, 0x73, 0x70, 0x22,
	0x28, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x22, 0x1a, 0x11, 0x2f, 0x63, 0x6f, 0x75, 0x72, 0x73, 0x65,
	0x77, 0x61, 0x72, 0x65, 0x73, 0x2f, 0x7b, 0x69, 0x64, 0x7d, 0x3a, 0x01, 0x2a, 0x62, 0x0a, 0x63,
	0x6f, 0x75, 0x72, 0x73, 0x65, 0x77, 0x61, 0x72, 0x65, 0x12, 0x6d, 0x0a, 0x06, 0x44, 0x65, 0x6c,
	0x65, 0x74, 0x65, 0x12, 0x23, 0x2e, 0x63, 0x6f, 0x6d, 0x2e, 0x69, 0x74, 0x63, 0x6f, 0x75, 0x72,
	0x73, 0x65, 0x65, 0x2e, 0x63, 0x6f, 0x75, 0x72, 0x73, 0x65, 0x77, 0x61, 0x72, 0x65, 0x2e, 0x44,
	0x65, 0x6c, 0x65, 0x74, 0x65, 0x52, 0x65, 0x71, 0x1a, 0x23, 0x2e, 0x63, 0x6f, 0x6d, 0x2e, 0x69,
	0x74, 0x63, 0x6f, 0x75, 0x72, 0x73, 0x65, 0x65, 0x2e, 0x63, 0x6f, 0x75, 0x72, 0x73, 0x65, 0x77,
	0x61, 0x72, 0x65, 0x2e, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x52, 0x73, 0x70, 0x22, 0x19, 0x82,
	0xd3, 0xe4, 0x93, 0x02, 0x13, 0x2a, 0x11, 0x2f, 0x63, 0x6f, 0x75, 0x72, 0x73, 0x65, 0x77, 0x61,
	0x72, 0x65, 0x73, 0x2f, 0x7b, 0x69, 0x64, 0x7d, 0x42, 0x40, 0x0a, 0x15, 0x63, 0x6f, 0x6d, 0x2e,
	0x69, 0x74, 0x63, 0x6f, 0x75, 0x72, 0x73, 0x65, 0x65, 0x2e, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67,
	0x65, 0x50, 0x01, 0x5a, 0x25, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f,
	0x69, 0x74, 0x63, 0x6f, 0x75, 0x72, 0x73, 0x65, 0x65, 0x2f, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67,
	0x65, 0x3b, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x6f, 0x6c, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x33,
}

var file_com_itcoursee_message_courseware_proto_goTypes = []interface{}{
	(*courseware.AddByCourseReq)(nil),   // 0: com.itcoursee.courseware.AddByCourseReq
	(*courseware.AddByLectureReq)(nil),  // 1: com.itcoursee.courseware.AddByLectureReq
	(*courseware.GetReq)(nil),           // 2: com.itcoursee.courseware.GetReq
	(*courseware.GetsByCourseReq)(nil),  // 3: com.itcoursee.courseware.GetsByCourseReq
	(*courseware.GetsByLectureReq)(nil), // 4: com.itcoursee.courseware.GetsByLectureReq
	(*courseware.UpdateReq)(nil),        // 5: com.itcoursee.courseware.UpdateReq
	(*courseware.DeleteReq)(nil),        // 6: com.itcoursee.courseware.DeleteReq
	(*courseware.AddRsp)(nil),           // 7: com.itcoursee.courseware.AddRsp
	(*courseware.GetRsp)(nil),           // 8: com.itcoursee.courseware.GetRsp
	(*courseware.GetsByOwnerRsp)(nil),   // 9: com.itcoursee.courseware.GetsByOwnerRsp
	(*courseware.UpdateRsp)(nil),        // 10: com.itcoursee.courseware.UpdateRsp
	(*courseware.DeleteRsp)(nil),        // 11: com.itcoursee.courseware.DeleteRsp
}
var file_com_itcoursee_message_courseware_proto_depIdxs = []int32{
	0,  // 0: com.itcoursee.message.Courseware.AddByCourse:input_type -> com.itcoursee.courseware.AddByCourseReq
	1,  // 1: com.itcoursee.message.Courseware.AddByLecture:input_type -> com.itcoursee.courseware.AddByLectureReq
	2,  // 2: com.itcoursee.message.Courseware.Get:input_type -> com.itcoursee.courseware.GetReq
	3,  // 3: com.itcoursee.message.Courseware.GetsByCourse:input_type -> com.itcoursee.courseware.GetsByCourseReq
	4,  // 4: com.itcoursee.message.Courseware.GetsByLecture:input_type -> com.itcoursee.courseware.GetsByLectureReq
	5,  // 5: com.itcoursee.message.Courseware.Update:input_type -> com.itcoursee.courseware.UpdateReq
	6,  // 6: com.itcoursee.message.Courseware.Delete:input_type -> com.itcoursee.courseware.DeleteReq
	7,  // 7: com.itcoursee.message.Courseware.AddByCourse:output_type -> com.itcoursee.courseware.AddRsp
	7,  // 8: com.itcoursee.message.Courseware.AddByLecture:output_type -> com.itcoursee.courseware.AddRsp
	8,  // 9: com.itcoursee.message.Courseware.Get:output_type -> com.itcoursee.courseware.GetRsp
	9,  // 10: com.itcoursee.message.Courseware.GetsByCourse:output_type -> com.itcoursee.courseware.GetsByOwnerRsp
	9,  // 11: com.itcoursee.message.Courseware.GetsByLecture:output_type -> com.itcoursee.courseware.GetsByOwnerRsp
	10, // 12: com.itcoursee.message.Courseware.Update:output_type -> com.itcoursee.courseware.UpdateRsp
	11, // 13: com.itcoursee.message.Courseware.Delete:output_type -> com.itcoursee.courseware.DeleteRsp
	7,  // [7:14] is the sub-list for method output_type
	0,  // [0:7] is the sub-list for method input_type
	0,  // [0:0] is the sub-list for extension type_name
	0,  // [0:0] is the sub-list for extension extendee
	0,  // [0:0] is the sub-list for field type_name
}

func init() { file_com_itcoursee_message_courseware_proto_init() }
func file_com_itcoursee_message_courseware_proto_init() {
	if File_com_itcoursee_message_courseware_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_com_itcoursee_message_courseware_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   0,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_com_itcoursee_message_courseware_proto_goTypes,
		DependencyIndexes: file_com_itcoursee_message_courseware_proto_depIdxs,
	}.Build()
	File_com_itcoursee_message_courseware_proto = out.File
	file_com_itcoursee_message_courseware_proto_rawDesc = nil
	file_com_itcoursee_message_courseware_proto_goTypes = nil
	file_com_itcoursee_message_courseware_proto_depIdxs = nil
}
