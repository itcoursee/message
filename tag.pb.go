// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.30.0
// 	protoc        v3.21.9
// source: com/itcoursee/message/tag.proto

package message

import (
	tag "github.com/itcoursee/core/tag"
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

var File_com_itcoursee_message_tag_proto protoreflect.FileDescriptor

var file_com_itcoursee_message_tag_proto_rawDesc = []byte{
	0x0a, 0x1f, 0x63, 0x6f, 0x6d, 0x2f, 0x69, 0x74, 0x63, 0x6f, 0x75, 0x72, 0x73, 0x65, 0x65, 0x2f,
	0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x2f, 0x74, 0x61, 0x67, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x12, 0x15, 0x63, 0x6f, 0x6d, 0x2e, 0x69, 0x74, 0x63, 0x6f, 0x75, 0x72, 0x73, 0x65, 0x65,
	0x2e, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x1a, 0x1c, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x2f, 0x61, 0x70, 0x69, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x20, 0x63, 0x6f, 0x6d, 0x2f, 0x69, 0x74, 0x63, 0x6f,
	0x75, 0x72, 0x73, 0x65, 0x65, 0x2f, 0x63, 0x6f, 0x72, 0x65, 0x2f, 0x74, 0x61, 0x67, 0x2f, 0x61,
	0x64, 0x64, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x23, 0x63, 0x6f, 0x6d, 0x2f, 0x69, 0x74,
	0x63, 0x6f, 0x75, 0x72, 0x73, 0x65, 0x65, 0x2f, 0x63, 0x6f, 0x72, 0x65, 0x2f, 0x74, 0x61, 0x67,
	0x2f, 0x64, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x20, 0x63,
	0x6f, 0x6d, 0x2f, 0x69, 0x74, 0x63, 0x6f, 0x75, 0x72, 0x73, 0x65, 0x65, 0x2f, 0x63, 0x6f, 0x72,
	0x65, 0x2f, 0x74, 0x61, 0x67, 0x2f, 0x67, 0x65, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a,
	0x23, 0x63, 0x6f, 0x6d, 0x2f, 0x69, 0x74, 0x63, 0x6f, 0x75, 0x72, 0x73, 0x65, 0x65, 0x2f, 0x63,
	0x6f, 0x72, 0x65, 0x2f, 0x74, 0x61, 0x67, 0x2f, 0x70, 0x61, 0x67, 0x69, 0x6e, 0x67, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x23, 0x63, 0x6f, 0x6d, 0x2f, 0x69, 0x74, 0x63, 0x6f, 0x75, 0x72,
	0x73, 0x65, 0x65, 0x2f, 0x63, 0x6f, 0x72, 0x65, 0x2f, 0x74, 0x61, 0x67, 0x2f, 0x75, 0x70, 0x64,
	0x61, 0x74, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x32, 0xb1, 0x03, 0x0a, 0x03, 0x54, 0x61,
	0x67, 0x12, 0x47, 0x0a, 0x03, 0x41, 0x64, 0x64, 0x12, 0x1e, 0x2e, 0x63, 0x6f, 0x6d, 0x2e, 0x69,
	0x74, 0x63, 0x6f, 0x75, 0x72, 0x73, 0x65, 0x65, 0x2e, 0x63, 0x6f, 0x72, 0x65, 0x2e, 0x74, 0x61,
	0x67, 0x2e, 0x41, 0x64, 0x64, 0x52, 0x65, 0x71, 0x1a, 0x1e, 0x2e, 0x63, 0x6f, 0x6d, 0x2e, 0x69,
	0x74, 0x63, 0x6f, 0x75, 0x72, 0x73, 0x65, 0x65, 0x2e, 0x63, 0x6f, 0x72, 0x65, 0x2e, 0x74, 0x61,
	0x67, 0x2e, 0x41, 0x64, 0x64, 0x52, 0x73, 0x70, 0x22, 0x00, 0x12, 0x5e, 0x0a, 0x03, 0x47, 0x65,
	0x74, 0x12, 0x1e, 0x2e, 0x63, 0x6f, 0x6d, 0x2e, 0x69, 0x74, 0x63, 0x6f, 0x75, 0x72, 0x73, 0x65,
	0x65, 0x2e, 0x63, 0x6f, 0x72, 0x65, 0x2e, 0x74, 0x61, 0x67, 0x2e, 0x47, 0x65, 0x74, 0x52, 0x65,
	0x71, 0x1a, 0x1e, 0x2e, 0x63, 0x6f, 0x6d, 0x2e, 0x69, 0x74, 0x63, 0x6f, 0x75, 0x72, 0x73, 0x65,
	0x65, 0x2e, 0x63, 0x6f, 0x72, 0x65, 0x2e, 0x74, 0x61, 0x67, 0x2e, 0x47, 0x65, 0x74, 0x52, 0x73,
	0x70, 0x22, 0x17, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x11, 0x62, 0x03, 0x74, 0x61, 0x67, 0x12, 0x0a,
	0x2f, 0x74, 0x61, 0x67, 0x73, 0x2f, 0x7b, 0x69, 0x64, 0x7d, 0x12, 0x5d, 0x0a, 0x06, 0x50, 0x61,
	0x67, 0x69, 0x6e, 0x67, 0x12, 0x21, 0x2e, 0x63, 0x6f, 0x6d, 0x2e, 0x69, 0x74, 0x63, 0x6f, 0x75,
	0x72, 0x73, 0x65, 0x65, 0x2e, 0x63, 0x6f, 0x72, 0x65, 0x2e, 0x74, 0x61, 0x67, 0x2e, 0x50, 0x61,
	0x67, 0x69, 0x6e, 0x67, 0x52, 0x65, 0x71, 0x1a, 0x21, 0x2e, 0x63, 0x6f, 0x6d, 0x2e, 0x69, 0x74,
	0x63, 0x6f, 0x75, 0x72, 0x73, 0x65, 0x65, 0x2e, 0x63, 0x6f, 0x72, 0x65, 0x2e, 0x74, 0x61, 0x67,
	0x2e, 0x50, 0x61, 0x67, 0x69, 0x6e, 0x67, 0x52, 0x73, 0x70, 0x22, 0x0d, 0x82, 0xd3, 0xe4, 0x93,
	0x02, 0x07, 0x12, 0x05, 0x2f, 0x74, 0x61, 0x67, 0x73, 0x12, 0x50, 0x0a, 0x06, 0x55, 0x70, 0x64,
	0x61, 0x74, 0x65, 0x12, 0x21, 0x2e, 0x63, 0x6f, 0x6d, 0x2e, 0x69, 0x74, 0x63, 0x6f, 0x75, 0x72,
	0x73, 0x65, 0x65, 0x2e, 0x63, 0x6f, 0x72, 0x65, 0x2e, 0x74, 0x61, 0x67, 0x2e, 0x55, 0x70, 0x64,
	0x61, 0x74, 0x65, 0x52, 0x65, 0x71, 0x1a, 0x21, 0x2e, 0x63, 0x6f, 0x6d, 0x2e, 0x69, 0x74, 0x63,
	0x6f, 0x75, 0x72, 0x73, 0x65, 0x65, 0x2e, 0x63, 0x6f, 0x72, 0x65, 0x2e, 0x74, 0x61, 0x67, 0x2e,
	0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x52, 0x73, 0x70, 0x22, 0x00, 0x12, 0x50, 0x0a, 0x06, 0x44,
	0x65, 0x6c, 0x65, 0x74, 0x65, 0x12, 0x21, 0x2e, 0x63, 0x6f, 0x6d, 0x2e, 0x69, 0x74, 0x63, 0x6f,
	0x75, 0x72, 0x73, 0x65, 0x65, 0x2e, 0x63, 0x6f, 0x72, 0x65, 0x2e, 0x74, 0x61, 0x67, 0x2e, 0x44,
	0x65, 0x6c, 0x65, 0x74, 0x65, 0x52, 0x65, 0x71, 0x1a, 0x21, 0x2e, 0x63, 0x6f, 0x6d, 0x2e, 0x69,
	0x74, 0x63, 0x6f, 0x75, 0x72, 0x73, 0x65, 0x65, 0x2e, 0x63, 0x6f, 0x72, 0x65, 0x2e, 0x74, 0x61,
	0x67, 0x2e, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x52, 0x73, 0x70, 0x22, 0x00, 0x42, 0x3f, 0x0a,
	0x15, 0x63, 0x6f, 0x6d, 0x2e, 0x69, 0x74, 0x63, 0x6f, 0x75, 0x72, 0x73, 0x65, 0x65, 0x2e, 0x6d,
	0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x50, 0x01, 0x5a, 0x24, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62,
	0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x69, 0x74, 0x63, 0x6f, 0x75, 0x72, 0x73, 0x65, 0x65, 0x2f, 0x6d,
	0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x3b, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x62, 0x06,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var file_com_itcoursee_message_tag_proto_goTypes = []interface{}{
	(*tag.AddReq)(nil),    // 0: com.itcoursee.core.tag.AddReq
	(*tag.GetReq)(nil),    // 1: com.itcoursee.core.tag.GetReq
	(*tag.PagingReq)(nil), // 2: com.itcoursee.core.tag.PagingReq
	(*tag.UpdateReq)(nil), // 3: com.itcoursee.core.tag.UpdateReq
	(*tag.DeleteReq)(nil), // 4: com.itcoursee.core.tag.DeleteReq
	(*tag.AddRsp)(nil),    // 5: com.itcoursee.core.tag.AddRsp
	(*tag.GetRsp)(nil),    // 6: com.itcoursee.core.tag.GetRsp
	(*tag.PagingRsp)(nil), // 7: com.itcoursee.core.tag.PagingRsp
	(*tag.UpdateRsp)(nil), // 8: com.itcoursee.core.tag.UpdateRsp
	(*tag.DeleteRsp)(nil), // 9: com.itcoursee.core.tag.DeleteRsp
}
var file_com_itcoursee_message_tag_proto_depIdxs = []int32{
	0, // 0: com.itcoursee.message.Tag.Add:input_type -> com.itcoursee.core.tag.AddReq
	1, // 1: com.itcoursee.message.Tag.Get:input_type -> com.itcoursee.core.tag.GetReq
	2, // 2: com.itcoursee.message.Tag.Paging:input_type -> com.itcoursee.core.tag.PagingReq
	3, // 3: com.itcoursee.message.Tag.Update:input_type -> com.itcoursee.core.tag.UpdateReq
	4, // 4: com.itcoursee.message.Tag.Delete:input_type -> com.itcoursee.core.tag.DeleteReq
	5, // 5: com.itcoursee.message.Tag.Add:output_type -> com.itcoursee.core.tag.AddRsp
	6, // 6: com.itcoursee.message.Tag.Get:output_type -> com.itcoursee.core.tag.GetRsp
	7, // 7: com.itcoursee.message.Tag.Paging:output_type -> com.itcoursee.core.tag.PagingRsp
	8, // 8: com.itcoursee.message.Tag.Update:output_type -> com.itcoursee.core.tag.UpdateRsp
	9, // 9: com.itcoursee.message.Tag.Delete:output_type -> com.itcoursee.core.tag.DeleteRsp
	5, // [5:10] is the sub-list for method output_type
	0, // [0:5] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_com_itcoursee_message_tag_proto_init() }
func file_com_itcoursee_message_tag_proto_init() {
	if File_com_itcoursee_message_tag_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_com_itcoursee_message_tag_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   0,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_com_itcoursee_message_tag_proto_goTypes,
		DependencyIndexes: file_com_itcoursee_message_tag_proto_depIdxs,
	}.Build()
	File_com_itcoursee_message_tag_proto = out.File
	file_com_itcoursee_message_tag_proto_rawDesc = nil
	file_com_itcoursee_message_tag_proto_goTypes = nil
	file_com_itcoursee_message_tag_proto_depIdxs = nil
}
