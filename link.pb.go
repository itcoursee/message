// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.18.1
// source: com/itcoursee/message/link.proto

package protocol

import (
	link "github.com/itcoursee/core/link"
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

var File_com_itcoursee_message_link_proto protoreflect.FileDescriptor

var file_com_itcoursee_message_link_proto_rawDesc = []byte{
	0x0a, 0x20, 0x63, 0x6f, 0x6d, 0x2f, 0x69, 0x74, 0x63, 0x6f, 0x75, 0x72, 0x73, 0x65, 0x65, 0x2f,
	0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x2f, 0x6c, 0x69, 0x6e, 0x6b, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x12, 0x15, 0x63, 0x6f, 0x6d, 0x2e, 0x69, 0x74, 0x63, 0x6f, 0x75, 0x72, 0x73, 0x65,
	0x65, 0x2e, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x1a, 0x1c, 0x67, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1c, 0x63, 0x6f, 0x6d, 0x2f, 0x69, 0x74, 0x63,
	0x6f, 0x75, 0x72, 0x73, 0x65, 0x65, 0x2f, 0x6c, 0x69, 0x6e, 0x6b, 0x2f, 0x61, 0x64, 0x64, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1f, 0x63, 0x6f, 0x6d, 0x2f, 0x69, 0x74, 0x63, 0x6f, 0x75,
	0x72, 0x73, 0x65, 0x65, 0x2f, 0x6c, 0x69, 0x6e, 0x6b, 0x2f, 0x64, 0x65, 0x6c, 0x65, 0x74, 0x65,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1c, 0x63, 0x6f, 0x6d, 0x2f, 0x69, 0x74, 0x63, 0x6f,
	0x75, 0x72, 0x73, 0x65, 0x65, 0x2f, 0x6c, 0x69, 0x6e, 0x6b, 0x2f, 0x67, 0x65, 0x74, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1f, 0x63, 0x6f, 0x6d, 0x2f, 0x69, 0x74, 0x63, 0x6f, 0x75, 0x72,
	0x73, 0x65, 0x65, 0x2f, 0x6c, 0x69, 0x6e, 0x6b, 0x2f, 0x70, 0x61, 0x67, 0x69, 0x6e, 0x67, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1f, 0x63, 0x6f, 0x6d, 0x2f, 0x69, 0x74, 0x63, 0x6f, 0x75,
	0x72, 0x73, 0x65, 0x65, 0x2f, 0x6c, 0x69, 0x6e, 0x6b, 0x2f, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x32, 0xd3, 0x03, 0x0a, 0x04, 0x4c, 0x69, 0x6e, 0x6b, 0x12,
	0x56, 0x0a, 0x03, 0x41, 0x64, 0x64, 0x12, 0x1a, 0x2e, 0x63, 0x6f, 0x6d, 0x2e, 0x69, 0x74, 0x63,
	0x6f, 0x75, 0x72, 0x73, 0x65, 0x65, 0x2e, 0x6c, 0x69, 0x6e, 0x6b, 0x2e, 0x41, 0x64, 0x64, 0x52,
	0x65, 0x71, 0x1a, 0x1a, 0x2e, 0x63, 0x6f, 0x6d, 0x2e, 0x69, 0x74, 0x63, 0x6f, 0x75, 0x72, 0x73,
	0x65, 0x65, 0x2e, 0x6c, 0x69, 0x6e, 0x6b, 0x2e, 0x41, 0x64, 0x64, 0x52, 0x73, 0x70, 0x22, 0x17,
	0x82, 0xd3, 0xe4, 0x93, 0x02, 0x11, 0x22, 0x06, 0x2f, 0x6c, 0x69, 0x6e, 0x6b, 0x73, 0x3a, 0x01,
	0x2a, 0x62, 0x04, 0x6c, 0x69, 0x6e, 0x6b, 0x12, 0x58, 0x0a, 0x03, 0x47, 0x65, 0x74, 0x12, 0x1a,
	0x2e, 0x63, 0x6f, 0x6d, 0x2e, 0x69, 0x74, 0x63, 0x6f, 0x75, 0x72, 0x73, 0x65, 0x65, 0x2e, 0x6c,
	0x69, 0x6e, 0x6b, 0x2e, 0x47, 0x65, 0x74, 0x52, 0x65, 0x71, 0x1a, 0x1a, 0x2e, 0x63, 0x6f, 0x6d,
	0x2e, 0x69, 0x74, 0x63, 0x6f, 0x75, 0x72, 0x73, 0x65, 0x65, 0x2e, 0x6c, 0x69, 0x6e, 0x6b, 0x2e,
	0x47, 0x65, 0x74, 0x52, 0x73, 0x70, 0x22, 0x19, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x13, 0x12, 0x0b,
	0x2f, 0x6c, 0x69, 0x6e, 0x6b, 0x73, 0x2f, 0x7b, 0x69, 0x64, 0x7d, 0x62, 0x04, 0x6c, 0x69, 0x6e,
	0x6b, 0x12, 0x56, 0x0a, 0x06, 0x50, 0x61, 0x67, 0x69, 0x6e, 0x67, 0x12, 0x1d, 0x2e, 0x63, 0x6f,
	0x6d, 0x2e, 0x69, 0x74, 0x63, 0x6f, 0x75, 0x72, 0x73, 0x65, 0x65, 0x2e, 0x6c, 0x69, 0x6e, 0x6b,
	0x2e, 0x50, 0x61, 0x67, 0x69, 0x6e, 0x67, 0x52, 0x65, 0x71, 0x1a, 0x1d, 0x2e, 0x63, 0x6f, 0x6d,
	0x2e, 0x69, 0x74, 0x63, 0x6f, 0x75, 0x72, 0x73, 0x65, 0x65, 0x2e, 0x6c, 0x69, 0x6e, 0x6b, 0x2e,
	0x50, 0x61, 0x67, 0x69, 0x6e, 0x67, 0x52, 0x73, 0x70, 0x22, 0x0e, 0x82, 0xd3, 0xe4, 0x93, 0x02,
	0x08, 0x12, 0x06, 0x2f, 0x6c, 0x69, 0x6e, 0x6b, 0x73, 0x12, 0x64, 0x0a, 0x06, 0x55, 0x70, 0x64,
	0x61, 0x74, 0x65, 0x12, 0x1d, 0x2e, 0x63, 0x6f, 0x6d, 0x2e, 0x69, 0x74, 0x63, 0x6f, 0x75, 0x72,
	0x73, 0x65, 0x65, 0x2e, 0x6c, 0x69, 0x6e, 0x6b, 0x2e, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x52,
	0x65, 0x71, 0x1a, 0x1d, 0x2e, 0x63, 0x6f, 0x6d, 0x2e, 0x69, 0x74, 0x63, 0x6f, 0x75, 0x72, 0x73,
	0x65, 0x65, 0x2e, 0x6c, 0x69, 0x6e, 0x6b, 0x2e, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x52, 0x73,
	0x70, 0x22, 0x1c, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x16, 0x1a, 0x0b, 0x2f, 0x6c, 0x69, 0x6e, 0x6b,
	0x73, 0x2f, 0x7b, 0x69, 0x64, 0x7d, 0x3a, 0x01, 0x2a, 0x62, 0x04, 0x6c, 0x69, 0x6e, 0x6b, 0x12,
	0x5b, 0x0a, 0x06, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x12, 0x1d, 0x2e, 0x63, 0x6f, 0x6d, 0x2e,
	0x69, 0x74, 0x63, 0x6f, 0x75, 0x72, 0x73, 0x65, 0x65, 0x2e, 0x6c, 0x69, 0x6e, 0x6b, 0x2e, 0x44,
	0x65, 0x6c, 0x65, 0x74, 0x65, 0x52, 0x65, 0x71, 0x1a, 0x1d, 0x2e, 0x63, 0x6f, 0x6d, 0x2e, 0x69,
	0x74, 0x63, 0x6f, 0x75, 0x72, 0x73, 0x65, 0x65, 0x2e, 0x6c, 0x69, 0x6e, 0x6b, 0x2e, 0x44, 0x65,
	0x6c, 0x65, 0x74, 0x65, 0x52, 0x73, 0x70, 0x22, 0x13, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x0d, 0x2a,
	0x0b, 0x2f, 0x6c, 0x69, 0x6e, 0x6b, 0x73, 0x2f, 0x7b, 0x69, 0x64, 0x7d, 0x42, 0x40, 0x0a, 0x15,
	0x63, 0x6f, 0x6d, 0x2e, 0x69, 0x74, 0x63, 0x6f, 0x75, 0x72, 0x73, 0x65, 0x65, 0x2e, 0x6d, 0x65,
	0x73, 0x73, 0x61, 0x67, 0x65, 0x50, 0x01, 0x5a, 0x25, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e,
	0x63, 0x6f, 0x6d, 0x2f, 0x69, 0x74, 0x63, 0x6f, 0x75, 0x72, 0x73, 0x65, 0x65, 0x2f, 0x6d, 0x65,
	0x73, 0x73, 0x61, 0x67, 0x65, 0x3b, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x6f, 0x6c, 0x62, 0x06,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var file_com_itcoursee_message_link_proto_goTypes = []interface{}{
	(*link.AddReq)(nil),    // 0: com.itcoursee.link.AddReq
	(*link.GetReq)(nil),    // 1: com.itcoursee.link.GetReq
	(*link.PagingReq)(nil), // 2: com.itcoursee.link.PagingReq
	(*link.UpdateReq)(nil), // 3: com.itcoursee.link.UpdateReq
	(*link.DeleteReq)(nil), // 4: com.itcoursee.link.DeleteReq
	(*link.AddRsp)(nil),    // 5: com.itcoursee.link.AddRsp
	(*link.GetRsp)(nil),    // 6: com.itcoursee.link.GetRsp
	(*link.PagingRsp)(nil), // 7: com.itcoursee.link.PagingRsp
	(*link.UpdateRsp)(nil), // 8: com.itcoursee.link.UpdateRsp
	(*link.DeleteRsp)(nil), // 9: com.itcoursee.link.DeleteRsp
}
var file_com_itcoursee_message_link_proto_depIdxs = []int32{
	0, // 0: com.itcoursee.message.Link.Add:input_type -> com.itcoursee.link.AddReq
	1, // 1: com.itcoursee.message.Link.Get:input_type -> com.itcoursee.link.GetReq
	2, // 2: com.itcoursee.message.Link.Paging:input_type -> com.itcoursee.link.PagingReq
	3, // 3: com.itcoursee.message.Link.Update:input_type -> com.itcoursee.link.UpdateReq
	4, // 4: com.itcoursee.message.Link.Delete:input_type -> com.itcoursee.link.DeleteReq
	5, // 5: com.itcoursee.message.Link.Add:output_type -> com.itcoursee.link.AddRsp
	6, // 6: com.itcoursee.message.Link.Get:output_type -> com.itcoursee.link.GetRsp
	7, // 7: com.itcoursee.message.Link.Paging:output_type -> com.itcoursee.link.PagingRsp
	8, // 8: com.itcoursee.message.Link.Update:output_type -> com.itcoursee.link.UpdateRsp
	9, // 9: com.itcoursee.message.Link.Delete:output_type -> com.itcoursee.link.DeleteRsp
	5, // [5:10] is the sub-list for method output_type
	0, // [0:5] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_com_itcoursee_message_link_proto_init() }
func file_com_itcoursee_message_link_proto_init() {
	if File_com_itcoursee_message_link_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_com_itcoursee_message_link_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   0,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_com_itcoursee_message_link_proto_goTypes,
		DependencyIndexes: file_com_itcoursee_message_link_proto_depIdxs,
	}.Build()
	File_com_itcoursee_message_link_proto = out.File
	file_com_itcoursee_message_link_proto_rawDesc = nil
	file_com_itcoursee_message_link_proto_goTypes = nil
	file_com_itcoursee_message_link_proto_depIdxs = nil
}
