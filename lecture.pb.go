// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.18.1
// source: com/itcoursee/message/lecture.proto

package message

import (
	lecture "github.com/itcoursee/core/lecture"
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

var File_com_itcoursee_message_lecture_proto protoreflect.FileDescriptor

var file_com_itcoursee_message_lecture_proto_rawDesc = []byte{
	0x0a, 0x23, 0x63, 0x6f, 0x6d, 0x2f, 0x69, 0x74, 0x63, 0x6f, 0x75, 0x72, 0x73, 0x65, 0x65, 0x2f,
	0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x2f, 0x6c, 0x65, 0x63, 0x74, 0x75, 0x72, 0x65, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x15, 0x63, 0x6f, 0x6d, 0x2e, 0x69, 0x74, 0x63, 0x6f, 0x75,
	0x72, 0x73, 0x65, 0x65, 0x2e, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x1a, 0x1c, 0x67, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x24, 0x63, 0x6f, 0x6d, 0x2f,
	0x69, 0x74, 0x63, 0x6f, 0x75, 0x72, 0x73, 0x65, 0x65, 0x2f, 0x63, 0x6f, 0x72, 0x65, 0x2f, 0x6c,
	0x65, 0x63, 0x74, 0x75, 0x72, 0x65, 0x2f, 0x61, 0x64, 0x64, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x1a, 0x27, 0x63, 0x6f, 0x6d, 0x2f, 0x69, 0x74, 0x63, 0x6f, 0x75, 0x72, 0x73, 0x65, 0x65, 0x2f,
	0x63, 0x6f, 0x72, 0x65, 0x2f, 0x6c, 0x65, 0x63, 0x74, 0x75, 0x72, 0x65, 0x2f, 0x63, 0x6f, 0x75,
	0x72, 0x73, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x27, 0x63, 0x6f, 0x6d, 0x2f, 0x69,
	0x74, 0x63, 0x6f, 0x75, 0x72, 0x73, 0x65, 0x65, 0x2f, 0x63, 0x6f, 0x72, 0x65, 0x2f, 0x6c, 0x65,
	0x63, 0x74, 0x75, 0x72, 0x65, 0x2f, 0x64, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x1a, 0x24, 0x63, 0x6f, 0x6d, 0x2f, 0x69, 0x74, 0x63, 0x6f, 0x75, 0x72, 0x73, 0x65,
	0x65, 0x2f, 0x63, 0x6f, 0x72, 0x65, 0x2f, 0x6c, 0x65, 0x63, 0x74, 0x75, 0x72, 0x65, 0x2f, 0x67,
	0x65, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x29, 0x63, 0x6f, 0x6d, 0x2f, 0x69, 0x74,
	0x63, 0x6f, 0x75, 0x72, 0x73, 0x65, 0x65, 0x2f, 0x63, 0x6f, 0x72, 0x65, 0x2f, 0x6c, 0x65, 0x63,
	0x74, 0x75, 0x72, 0x65, 0x2f, 0x70, 0x6c, 0x61, 0x79, 0x62, 0x61, 0x63, 0x6b, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x1a, 0x29, 0x63, 0x6f, 0x6d, 0x2f, 0x69, 0x74, 0x63, 0x6f, 0x75, 0x72, 0x73,
	0x65, 0x65, 0x2f, 0x63, 0x6f, 0x72, 0x65, 0x2f, 0x6c, 0x65, 0x63, 0x74, 0x75, 0x72, 0x65, 0x2f,
	0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x27,
	0x63, 0x6f, 0x6d, 0x2f, 0x69, 0x74, 0x63, 0x6f, 0x75, 0x72, 0x73, 0x65, 0x65, 0x2f, 0x63, 0x6f,
	0x72, 0x65, 0x2f, 0x6c, 0x65, 0x63, 0x74, 0x75, 0x72, 0x65, 0x2f, 0x75, 0x70, 0x64, 0x61, 0x74,
	0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x32, 0xba, 0x08, 0x0a, 0x07, 0x4c, 0x65, 0x63, 0x74,
	0x75, 0x72, 0x65, 0x12, 0x7d, 0x0a, 0x03, 0x41, 0x64, 0x64, 0x12, 0x22, 0x2e, 0x63, 0x6f, 0x6d,
	0x2e, 0x69, 0x74, 0x63, 0x6f, 0x75, 0x72, 0x73, 0x65, 0x65, 0x2e, 0x63, 0x6f, 0x72, 0x65, 0x2e,
	0x6c, 0x65, 0x63, 0x74, 0x75, 0x72, 0x65, 0x2e, 0x41, 0x64, 0x64, 0x52, 0x65, 0x71, 0x1a, 0x22,
	0x2e, 0x63, 0x6f, 0x6d, 0x2e, 0x69, 0x74, 0x63, 0x6f, 0x75, 0x72, 0x73, 0x65, 0x65, 0x2e, 0x63,
	0x6f, 0x72, 0x65, 0x2e, 0x6c, 0x65, 0x63, 0x74, 0x75, 0x72, 0x65, 0x2e, 0x41, 0x64, 0x64, 0x52,
	0x73, 0x70, 0x22, 0x2e, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x28, 0x22, 0x1a, 0x2f, 0x63, 0x6f, 0x75,
	0x72, 0x73, 0x65, 0x73, 0x2f, 0x7b, 0x63, 0x6f, 0x75, 0x72, 0x73, 0x65, 0x7d, 0x2f, 0x6c, 0x65,
	0x63, 0x74, 0x75, 0x72, 0x65, 0x73, 0x3a, 0x01, 0x2a, 0x62, 0x07, 0x6c, 0x65, 0x63, 0x74, 0x75,
	0x72, 0x65, 0x12, 0x7f, 0x0a, 0x03, 0x47, 0x65, 0x74, 0x12, 0x22, 0x2e, 0x63, 0x6f, 0x6d, 0x2e,
	0x69, 0x74, 0x63, 0x6f, 0x75, 0x72, 0x73, 0x65, 0x65, 0x2e, 0x63, 0x6f, 0x72, 0x65, 0x2e, 0x6c,
	0x65, 0x63, 0x74, 0x75, 0x72, 0x65, 0x2e, 0x47, 0x65, 0x74, 0x52, 0x65, 0x71, 0x1a, 0x22, 0x2e,
	0x63, 0x6f, 0x6d, 0x2e, 0x69, 0x74, 0x63, 0x6f, 0x75, 0x72, 0x73, 0x65, 0x65, 0x2e, 0x63, 0x6f,
	0x72, 0x65, 0x2e, 0x6c, 0x65, 0x63, 0x74, 0x75, 0x72, 0x65, 0x2e, 0x47, 0x65, 0x74, 0x52, 0x73,
	0x70, 0x22, 0x30, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x2a, 0x12, 0x1f, 0x2f, 0x63, 0x6f, 0x75, 0x72,
	0x73, 0x65, 0x73, 0x2f, 0x7b, 0x63, 0x6f, 0x75, 0x72, 0x73, 0x65, 0x7d, 0x2f, 0x6c, 0x65, 0x63,
	0x74, 0x75, 0x72, 0x65, 0x73, 0x2f, 0x7b, 0x69, 0x64, 0x7d, 0x62, 0x07, 0x6c, 0x65, 0x63, 0x74,
	0x75, 0x72, 0x65, 0x12, 0x96, 0x01, 0x0a, 0x0c, 0x47, 0x65, 0x74, 0x73, 0x42, 0x79, 0x43, 0x6f,
	0x75, 0x72, 0x73, 0x65, 0x12, 0x2b, 0x2e, 0x63, 0x6f, 0x6d, 0x2e, 0x69, 0x74, 0x63, 0x6f, 0x75,
	0x72, 0x73, 0x65, 0x65, 0x2e, 0x63, 0x6f, 0x72, 0x65, 0x2e, 0x6c, 0x65, 0x63, 0x74, 0x75, 0x72,
	0x65, 0x2e, 0x47, 0x65, 0x74, 0x73, 0x42, 0x79, 0x43, 0x6f, 0x75, 0x72, 0x73, 0x65, 0x52, 0x65,
	0x71, 0x1a, 0x2b, 0x2e, 0x63, 0x6f, 0x6d, 0x2e, 0x69, 0x74, 0x63, 0x6f, 0x75, 0x72, 0x73, 0x65,
	0x65, 0x2e, 0x63, 0x6f, 0x72, 0x65, 0x2e, 0x6c, 0x65, 0x63, 0x74, 0x75, 0x72, 0x65, 0x2e, 0x47,
	0x65, 0x74, 0x73, 0x42, 0x79, 0x43, 0x6f, 0x75, 0x72, 0x73, 0x65, 0x52, 0x73, 0x70, 0x22, 0x2c,
	0x82, 0xd3, 0xe4, 0x93, 0x02, 0x26, 0x12, 0x1a, 0x2f, 0x63, 0x6f, 0x75, 0x72, 0x73, 0x65, 0x73,
	0x2f, 0x7b, 0x63, 0x6f, 0x75, 0x72, 0x73, 0x65, 0x7d, 0x2f, 0x6c, 0x65, 0x63, 0x74, 0x75, 0x72,
	0x65, 0x73, 0x62, 0x08, 0x6c, 0x65, 0x63, 0x74, 0x75, 0x72, 0x65, 0x73, 0x12, 0xa2, 0x01, 0x0a,
	0x0b, 0x47, 0x65, 0x74, 0x50, 0x6c, 0x61, 0x79, 0x62, 0x61, 0x63, 0x6b, 0x12, 0x2a, 0x2e, 0x63,
	0x6f, 0x6d, 0x2e, 0x69, 0x74, 0x63, 0x6f, 0x75, 0x72, 0x73, 0x65, 0x65, 0x2e, 0x63, 0x6f, 0x72,
	0x65, 0x2e, 0x6c, 0x65, 0x63, 0x74, 0x75, 0x72, 0x65, 0x2e, 0x47, 0x65, 0x74, 0x50, 0x6c, 0x61,
	0x79, 0x62, 0x61, 0x63, 0x6b, 0x52, 0x65, 0x71, 0x1a, 0x2a, 0x2e, 0x63, 0x6f, 0x6d, 0x2e, 0x69,
	0x74, 0x63, 0x6f, 0x75, 0x72, 0x73, 0x65, 0x65, 0x2e, 0x63, 0x6f, 0x72, 0x65, 0x2e, 0x6c, 0x65,
	0x63, 0x74, 0x75, 0x72, 0x65, 0x2e, 0x47, 0x65, 0x74, 0x50, 0x6c, 0x61, 0x79, 0x62, 0x61, 0x63,
	0x6b, 0x52, 0x73, 0x70, 0x22, 0x3b, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x35, 0x12, 0x29, 0x2f, 0x63,
	0x6f, 0x75, 0x72, 0x73, 0x65, 0x73, 0x2f, 0x7b, 0x63, 0x6f, 0x75, 0x72, 0x73, 0x65, 0x7d, 0x2f,
	0x6c, 0x65, 0x63, 0x74, 0x75, 0x72, 0x65, 0x73, 0x2f, 0x7b, 0x69, 0x64, 0x7d, 0x2f, 0x70, 0x6c,
	0x61, 0x79, 0x62, 0x61, 0x63, 0x6b, 0x73, 0x62, 0x08, 0x70, 0x6c, 0x61, 0x79, 0x62, 0x61, 0x63,
	0x6b, 0x12, 0xa9, 0x01, 0x0a, 0x06, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x12, 0x25, 0x2e, 0x63,
	0x6f, 0x6d, 0x2e, 0x69, 0x74, 0x63, 0x6f, 0x75, 0x72, 0x73, 0x65, 0x65, 0x2e, 0x63, 0x6f, 0x72,
	0x65, 0x2e, 0x6c, 0x65, 0x63, 0x74, 0x75, 0x72, 0x65, 0x2e, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65,
	0x52, 0x65, 0x71, 0x1a, 0x25, 0x2e, 0x63, 0x6f, 0x6d, 0x2e, 0x69, 0x74, 0x63, 0x6f, 0x75, 0x72,
	0x73, 0x65, 0x65, 0x2e, 0x63, 0x6f, 0x72, 0x65, 0x2e, 0x6c, 0x65, 0x63, 0x74, 0x75, 0x72, 0x65,
	0x2e, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x52, 0x73, 0x70, 0x22, 0x51, 0x82, 0xd3, 0xe4, 0x93,
	0x02, 0x4b, 0x1a, 0x1f, 0x2f, 0x63, 0x6f, 0x75, 0x72, 0x73, 0x65, 0x73, 0x2f, 0x7b, 0x63, 0x6f,
	0x75, 0x72, 0x73, 0x65, 0x7d, 0x2f, 0x6c, 0x65, 0x63, 0x74, 0x75, 0x72, 0x65, 0x73, 0x2f, 0x7b,
	0x69, 0x64, 0x7d, 0x3a, 0x01, 0x2a, 0x5a, 0x1c, 0x1a, 0x0e, 0x2f, 0x6c, 0x65, 0x63, 0x74, 0x75,
	0x72, 0x65, 0x73, 0x2f, 0x7b, 0x69, 0x64, 0x7d, 0x3a, 0x01, 0x2a, 0x62, 0x07, 0x6c, 0x65, 0x63,
	0x74, 0x75, 0x72, 0x65, 0x62, 0x07, 0x6c, 0x65, 0x63, 0x74, 0x75, 0x72, 0x65, 0x12, 0xc3, 0x01,
	0x0a, 0x08, 0x52, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x12, 0x27, 0x2e, 0x63, 0x6f, 0x6d,
	0x2e, 0x69, 0x74, 0x63, 0x6f, 0x75, 0x72, 0x73, 0x65, 0x65, 0x2e, 0x63, 0x6f, 0x72, 0x65, 0x2e,
	0x6c, 0x65, 0x63, 0x74, 0x75, 0x72, 0x65, 0x2e, 0x52, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65,
	0x52, 0x65, 0x71, 0x1a, 0x27, 0x2e, 0x63, 0x6f, 0x6d, 0x2e, 0x69, 0x74, 0x63, 0x6f, 0x75, 0x72,
	0x73, 0x65, 0x65, 0x2e, 0x63, 0x6f, 0x72, 0x65, 0x2e, 0x6c, 0x65, 0x63, 0x74, 0x75, 0x72, 0x65,
	0x2e, 0x52, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x52, 0x73, 0x70, 0x22, 0x65, 0x82, 0xd3,
	0xe4, 0x93, 0x02, 0x5f, 0x1a, 0x29, 0x2f, 0x63, 0x6f, 0x75, 0x72, 0x73, 0x65, 0x73, 0x2f, 0x7b,
	0x63, 0x6f, 0x75, 0x72, 0x73, 0x65, 0x7d, 0x2f, 0x6c, 0x65, 0x63, 0x74, 0x75, 0x72, 0x65, 0x73,
	0x2f, 0x7b, 0x69, 0x64, 0x7d, 0x2f, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x73, 0x3a,
	0x01, 0x2a, 0x5a, 0x26, 0x1a, 0x18, 0x2f, 0x6c, 0x65, 0x63, 0x74, 0x75, 0x72, 0x65, 0x73, 0x2f,
	0x7b, 0x69, 0x64, 0x7d, 0x2f, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x73, 0x3a, 0x01,
	0x2a, 0x62, 0x07, 0x6c, 0x65, 0x63, 0x74, 0x75, 0x72, 0x65, 0x62, 0x07, 0x6c, 0x65, 0x63, 0x74,
	0x75, 0x72, 0x65, 0x12, 0x7f, 0x0a, 0x06, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x12, 0x25, 0x2e,
	0x63, 0x6f, 0x6d, 0x2e, 0x69, 0x74, 0x63, 0x6f, 0x75, 0x72, 0x73, 0x65, 0x65, 0x2e, 0x63, 0x6f,
	0x72, 0x65, 0x2e, 0x6c, 0x65, 0x63, 0x74, 0x75, 0x72, 0x65, 0x2e, 0x44, 0x65, 0x6c, 0x65, 0x74,
	0x65, 0x52, 0x65, 0x71, 0x1a, 0x25, 0x2e, 0x63, 0x6f, 0x6d, 0x2e, 0x69, 0x74, 0x63, 0x6f, 0x75,
	0x72, 0x73, 0x65, 0x65, 0x2e, 0x63, 0x6f, 0x72, 0x65, 0x2e, 0x6c, 0x65, 0x63, 0x74, 0x75, 0x72,
	0x65, 0x2e, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x52, 0x73, 0x70, 0x22, 0x27, 0x82, 0xd3, 0xe4,
	0x93, 0x02, 0x21, 0x2a, 0x1f, 0x2f, 0x63, 0x6f, 0x75, 0x72, 0x73, 0x65, 0x73, 0x2f, 0x7b, 0x63,
	0x6f, 0x75, 0x72, 0x73, 0x65, 0x7d, 0x2f, 0x6c, 0x65, 0x63, 0x74, 0x75, 0x72, 0x65, 0x73, 0x2f,
	0x7b, 0x69, 0x64, 0x7d, 0x42, 0x3f, 0x0a, 0x15, 0x63, 0x6f, 0x6d, 0x2e, 0x69, 0x74, 0x63, 0x6f,
	0x75, 0x72, 0x73, 0x65, 0x65, 0x2e, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x50, 0x01, 0x5a,
	0x24, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x69, 0x74, 0x63, 0x6f,
	0x75, 0x72, 0x73, 0x65, 0x65, 0x2f, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x3b, 0x6d, 0x65,
	0x73, 0x73, 0x61, 0x67, 0x65, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var file_com_itcoursee_message_lecture_proto_goTypes = []interface{}{
	(*lecture.AddReq)(nil),          // 0: com.itcoursee.core.lecture.AddReq
	(*lecture.GetReq)(nil),          // 1: com.itcoursee.core.lecture.GetReq
	(*lecture.GetsByCourseReq)(nil), // 2: com.itcoursee.core.lecture.GetsByCourseReq
	(*lecture.GetPlaybackReq)(nil),  // 3: com.itcoursee.core.lecture.GetPlaybackReq
	(*lecture.UpdateReq)(nil),       // 4: com.itcoursee.core.lecture.UpdateReq
	(*lecture.ResourceReq)(nil),     // 5: com.itcoursee.core.lecture.ResourceReq
	(*lecture.DeleteReq)(nil),       // 6: com.itcoursee.core.lecture.DeleteReq
	(*lecture.AddRsp)(nil),          // 7: com.itcoursee.core.lecture.AddRsp
	(*lecture.GetRsp)(nil),          // 8: com.itcoursee.core.lecture.GetRsp
	(*lecture.GetsByCourseRsp)(nil), // 9: com.itcoursee.core.lecture.GetsByCourseRsp
	(*lecture.GetPlaybackRsp)(nil),  // 10: com.itcoursee.core.lecture.GetPlaybackRsp
	(*lecture.UpdateRsp)(nil),       // 11: com.itcoursee.core.lecture.UpdateRsp
	(*lecture.ResourceRsp)(nil),     // 12: com.itcoursee.core.lecture.ResourceRsp
	(*lecture.DeleteRsp)(nil),       // 13: com.itcoursee.core.lecture.DeleteRsp
}
var file_com_itcoursee_message_lecture_proto_depIdxs = []int32{
	0,  // 0: com.itcoursee.message.Lecture.Add:input_type -> com.itcoursee.core.lecture.AddReq
	1,  // 1: com.itcoursee.message.Lecture.Get:input_type -> com.itcoursee.core.lecture.GetReq
	2,  // 2: com.itcoursee.message.Lecture.GetsByCourse:input_type -> com.itcoursee.core.lecture.GetsByCourseReq
	3,  // 3: com.itcoursee.message.Lecture.GetPlayback:input_type -> com.itcoursee.core.lecture.GetPlaybackReq
	4,  // 4: com.itcoursee.message.Lecture.Update:input_type -> com.itcoursee.core.lecture.UpdateReq
	5,  // 5: com.itcoursee.message.Lecture.Resource:input_type -> com.itcoursee.core.lecture.ResourceReq
	6,  // 6: com.itcoursee.message.Lecture.Delete:input_type -> com.itcoursee.core.lecture.DeleteReq
	7,  // 7: com.itcoursee.message.Lecture.Add:output_type -> com.itcoursee.core.lecture.AddRsp
	8,  // 8: com.itcoursee.message.Lecture.Get:output_type -> com.itcoursee.core.lecture.GetRsp
	9,  // 9: com.itcoursee.message.Lecture.GetsByCourse:output_type -> com.itcoursee.core.lecture.GetsByCourseRsp
	10, // 10: com.itcoursee.message.Lecture.GetPlayback:output_type -> com.itcoursee.core.lecture.GetPlaybackRsp
	11, // 11: com.itcoursee.message.Lecture.Update:output_type -> com.itcoursee.core.lecture.UpdateRsp
	12, // 12: com.itcoursee.message.Lecture.Resource:output_type -> com.itcoursee.core.lecture.ResourceRsp
	13, // 13: com.itcoursee.message.Lecture.Delete:output_type -> com.itcoursee.core.lecture.DeleteRsp
	7,  // [7:14] is the sub-list for method output_type
	0,  // [0:7] is the sub-list for method input_type
	0,  // [0:0] is the sub-list for extension type_name
	0,  // [0:0] is the sub-list for extension extendee
	0,  // [0:0] is the sub-list for field type_name
}

func init() { file_com_itcoursee_message_lecture_proto_init() }
func file_com_itcoursee_message_lecture_proto_init() {
	if File_com_itcoursee_message_lecture_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_com_itcoursee_message_lecture_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   0,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_com_itcoursee_message_lecture_proto_goTypes,
		DependencyIndexes: file_com_itcoursee_message_lecture_proto_depIdxs,
	}.Build()
	File_com_itcoursee_message_lecture_proto = out.File
	file_com_itcoursee_message_lecture_proto_rawDesc = nil
	file_com_itcoursee_message_lecture_proto_goTypes = nil
	file_com_itcoursee_message_lecture_proto_depIdxs = nil
}
