// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.18.1
// source: com/itcoursee/message/chapter.proto

package message

import (
	chapter "github.com/itcoursee/core/chapter"
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

var File_com_itcoursee_message_chapter_proto protoreflect.FileDescriptor

var file_com_itcoursee_message_chapter_proto_rawDesc = []byte{
	0x0a, 0x23, 0x63, 0x6f, 0x6d, 0x2f, 0x69, 0x74, 0x63, 0x6f, 0x75, 0x72, 0x73, 0x65, 0x65, 0x2f,
	0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x2f, 0x63, 0x68, 0x61, 0x70, 0x74, 0x65, 0x72, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x15, 0x63, 0x6f, 0x6d, 0x2e, 0x69, 0x74, 0x63, 0x6f, 0x75,
	0x72, 0x73, 0x65, 0x65, 0x2e, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x1a, 0x1c, 0x67, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1f, 0x63, 0x6f, 0x6d, 0x2f,
	0x69, 0x74, 0x63, 0x6f, 0x75, 0x72, 0x73, 0x65, 0x65, 0x2f, 0x63, 0x68, 0x61, 0x70, 0x74, 0x65,
	0x72, 0x2f, 0x61, 0x64, 0x64, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x22, 0x63, 0x6f, 0x6d,
	0x2f, 0x69, 0x74, 0x63, 0x6f, 0x75, 0x72, 0x73, 0x65, 0x65, 0x2f, 0x63, 0x68, 0x61, 0x70, 0x74,
	0x65, 0x72, 0x2f, 0x63, 0x6f, 0x75, 0x72, 0x73, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a,
	0x22, 0x63, 0x6f, 0x6d, 0x2f, 0x69, 0x74, 0x63, 0x6f, 0x75, 0x72, 0x73, 0x65, 0x65, 0x2f, 0x63,
	0x68, 0x61, 0x70, 0x74, 0x65, 0x72, 0x2f, 0x64, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x1a, 0x1f, 0x63, 0x6f, 0x6d, 0x2f, 0x69, 0x74, 0x63, 0x6f, 0x75, 0x72, 0x73,
	0x65, 0x65, 0x2f, 0x63, 0x68, 0x61, 0x70, 0x74, 0x65, 0x72, 0x2f, 0x67, 0x65, 0x74, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x22, 0x63, 0x6f, 0x6d, 0x2f, 0x69, 0x74, 0x63, 0x6f, 0x75, 0x72,
	0x73, 0x65, 0x65, 0x2f, 0x63, 0x68, 0x61, 0x70, 0x74, 0x65, 0x72, 0x2f, 0x75, 0x70, 0x64, 0x61,
	0x74, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x32, 0xb0, 0x05, 0x0a, 0x07, 0x43, 0x68, 0x61,
	0x70, 0x74, 0x65, 0x72, 0x12, 0x73, 0x0a, 0x03, 0x41, 0x64, 0x64, 0x12, 0x1d, 0x2e, 0x63, 0x6f,
	0x6d, 0x2e, 0x69, 0x74, 0x63, 0x6f, 0x75, 0x72, 0x73, 0x65, 0x65, 0x2e, 0x63, 0x68, 0x61, 0x70,
	0x74, 0x65, 0x72, 0x2e, 0x41, 0x64, 0x64, 0x52, 0x65, 0x71, 0x1a, 0x1d, 0x2e, 0x63, 0x6f, 0x6d,
	0x2e, 0x69, 0x74, 0x63, 0x6f, 0x75, 0x72, 0x73, 0x65, 0x65, 0x2e, 0x63, 0x68, 0x61, 0x70, 0x74,
	0x65, 0x72, 0x2e, 0x41, 0x64, 0x64, 0x52, 0x73, 0x70, 0x22, 0x2e, 0x82, 0xd3, 0xe4, 0x93, 0x02,
	0x28, 0x22, 0x1a, 0x2f, 0x63, 0x6f, 0x75, 0x72, 0x73, 0x65, 0x73, 0x2f, 0x7b, 0x63, 0x6f, 0x75,
	0x72, 0x73, 0x65, 0x7d, 0x2f, 0x63, 0x68, 0x61, 0x70, 0x74, 0x65, 0x72, 0x73, 0x3a, 0x01, 0x2a,
	0x62, 0x07, 0x63, 0x68, 0x61, 0x70, 0x74, 0x65, 0x72, 0x12, 0x75, 0x0a, 0x03, 0x47, 0x65, 0x74,
	0x12, 0x1d, 0x2e, 0x63, 0x6f, 0x6d, 0x2e, 0x69, 0x74, 0x63, 0x6f, 0x75, 0x72, 0x73, 0x65, 0x65,
	0x2e, 0x63, 0x68, 0x61, 0x70, 0x74, 0x65, 0x72, 0x2e, 0x47, 0x65, 0x74, 0x52, 0x65, 0x71, 0x1a,
	0x1d, 0x2e, 0x63, 0x6f, 0x6d, 0x2e, 0x69, 0x74, 0x63, 0x6f, 0x75, 0x72, 0x73, 0x65, 0x65, 0x2e,
	0x63, 0x68, 0x61, 0x70, 0x74, 0x65, 0x72, 0x2e, 0x47, 0x65, 0x74, 0x52, 0x73, 0x70, 0x22, 0x30,
	0x82, 0xd3, 0xe4, 0x93, 0x02, 0x2a, 0x12, 0x1f, 0x2f, 0x63, 0x6f, 0x75, 0x72, 0x73, 0x65, 0x73,
	0x2f, 0x7b, 0x63, 0x6f, 0x75, 0x72, 0x73, 0x65, 0x7d, 0x2f, 0x63, 0x68, 0x61, 0x70, 0x74, 0x65,
	0x72, 0x73, 0x2f, 0x7b, 0x69, 0x64, 0x7d, 0x62, 0x07, 0x63, 0x68, 0x61, 0x70, 0x74, 0x65, 0x72,
	0x12, 0x8c, 0x01, 0x0a, 0x0c, 0x47, 0x65, 0x74, 0x73, 0x42, 0x79, 0x43, 0x6f, 0x75, 0x72, 0x73,
	0x65, 0x12, 0x26, 0x2e, 0x63, 0x6f, 0x6d, 0x2e, 0x69, 0x74, 0x63, 0x6f, 0x75, 0x72, 0x73, 0x65,
	0x65, 0x2e, 0x63, 0x68, 0x61, 0x70, 0x74, 0x65, 0x72, 0x2e, 0x47, 0x65, 0x74, 0x73, 0x42, 0x79,
	0x43, 0x6f, 0x75, 0x72, 0x73, 0x65, 0x52, 0x65, 0x71, 0x1a, 0x26, 0x2e, 0x63, 0x6f, 0x6d, 0x2e,
	0x69, 0x74, 0x63, 0x6f, 0x75, 0x72, 0x73, 0x65, 0x65, 0x2e, 0x63, 0x68, 0x61, 0x70, 0x74, 0x65,
	0x72, 0x2e, 0x47, 0x65, 0x74, 0x73, 0x42, 0x79, 0x43, 0x6f, 0x75, 0x72, 0x73, 0x65, 0x52, 0x73,
	0x70, 0x22, 0x2c, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x26, 0x12, 0x1a, 0x2f, 0x63, 0x6f, 0x75, 0x72,
	0x73, 0x65, 0x73, 0x2f, 0x7b, 0x63, 0x6f, 0x75, 0x72, 0x73, 0x65, 0x7d, 0x2f, 0x63, 0x68, 0x61,
	0x70, 0x74, 0x65, 0x72, 0x73, 0x62, 0x08, 0x63, 0x68, 0x61, 0x70, 0x74, 0x65, 0x72, 0x73, 0x12,
	0x9f, 0x01, 0x0a, 0x06, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x12, 0x20, 0x2e, 0x63, 0x6f, 0x6d,
	0x2e, 0x69, 0x74, 0x63, 0x6f, 0x75, 0x72, 0x73, 0x65, 0x65, 0x2e, 0x63, 0x68, 0x61, 0x70, 0x74,
	0x65, 0x72, 0x2e, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x52, 0x65, 0x71, 0x1a, 0x20, 0x2e, 0x63,
	0x6f, 0x6d, 0x2e, 0x69, 0x74, 0x63, 0x6f, 0x75, 0x72, 0x73, 0x65, 0x65, 0x2e, 0x63, 0x68, 0x61,
	0x70, 0x74, 0x65, 0x72, 0x2e, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x52, 0x73, 0x70, 0x22, 0x51,
	0x82, 0xd3, 0xe4, 0x93, 0x02, 0x4b, 0x1a, 0x1f, 0x2f, 0x63, 0x6f, 0x75, 0x72, 0x73, 0x65, 0x73,
	0x2f, 0x7b, 0x63, 0x6f, 0x75, 0x72, 0x73, 0x65, 0x7d, 0x2f, 0x63, 0x68, 0x61, 0x70, 0x74, 0x65,
	0x72, 0x73, 0x2f, 0x7b, 0x69, 0x64, 0x7d, 0x3a, 0x01, 0x2a, 0x5a, 0x1c, 0x1a, 0x0e, 0x2f, 0x63,
	0x68, 0x61, 0x70, 0x74, 0x65, 0x72, 0x73, 0x2f, 0x7b, 0x69, 0x64, 0x7d, 0x3a, 0x01, 0x2a, 0x62,
	0x07, 0x63, 0x68, 0x61, 0x70, 0x74, 0x65, 0x72, 0x62, 0x07, 0x63, 0x68, 0x61, 0x70, 0x74, 0x65,
	0x72, 0x12, 0x87, 0x01, 0x0a, 0x06, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x12, 0x20, 0x2e, 0x63,
	0x6f, 0x6d, 0x2e, 0x69, 0x74, 0x63, 0x6f, 0x75, 0x72, 0x73, 0x65, 0x65, 0x2e, 0x63, 0x68, 0x61,
	0x70, 0x74, 0x65, 0x72, 0x2e, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x52, 0x65, 0x71, 0x1a, 0x20,
	0x2e, 0x63, 0x6f, 0x6d, 0x2e, 0x69, 0x74, 0x63, 0x6f, 0x75, 0x72, 0x73, 0x65, 0x65, 0x2e, 0x63,
	0x68, 0x61, 0x70, 0x74, 0x65, 0x72, 0x2e, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x52, 0x73, 0x70,
	0x22, 0x39, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x33, 0x2a, 0x1f, 0x2f, 0x63, 0x6f, 0x75, 0x72, 0x73,
	0x65, 0x73, 0x2f, 0x7b, 0x63, 0x6f, 0x75, 0x72, 0x73, 0x65, 0x7d, 0x2f, 0x63, 0x68, 0x61, 0x70,
	0x74, 0x65, 0x72, 0x73, 0x2f, 0x7b, 0x69, 0x64, 0x7d, 0x5a, 0x10, 0x2a, 0x0e, 0x2f, 0x63, 0x68,
	0x61, 0x70, 0x74, 0x65, 0x72, 0x73, 0x2f, 0x7b, 0x69, 0x64, 0x7d, 0x42, 0x3f, 0x0a, 0x15, 0x63,
	0x6f, 0x6d, 0x2e, 0x69, 0x74, 0x63, 0x6f, 0x75, 0x72, 0x73, 0x65, 0x65, 0x2e, 0x6d, 0x65, 0x73,
	0x73, 0x61, 0x67, 0x65, 0x50, 0x01, 0x5a, 0x24, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63,
	0x6f, 0x6d, 0x2f, 0x69, 0x74, 0x63, 0x6f, 0x75, 0x72, 0x73, 0x65, 0x65, 0x2f, 0x6d, 0x65, 0x73,
	0x73, 0x61, 0x67, 0x65, 0x3b, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x62, 0x06, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x33,
}

var file_com_itcoursee_message_chapter_proto_goTypes = []interface{}{
	(*chapter.AddReq)(nil),          // 0: com.itcoursee.chapter.AddReq
	(*chapter.GetReq)(nil),          // 1: com.itcoursee.chapter.GetReq
	(*chapter.GetsByCourseReq)(nil), // 2: com.itcoursee.chapter.GetsByCourseReq
	(*chapter.UpdateReq)(nil),       // 3: com.itcoursee.chapter.UpdateReq
	(*chapter.DeleteReq)(nil),       // 4: com.itcoursee.chapter.DeleteReq
	(*chapter.AddRsp)(nil),          // 5: com.itcoursee.chapter.AddRsp
	(*chapter.GetRsp)(nil),          // 6: com.itcoursee.chapter.GetRsp
	(*chapter.GetsByCourseRsp)(nil), // 7: com.itcoursee.chapter.GetsByCourseRsp
	(*chapter.UpdateRsp)(nil),       // 8: com.itcoursee.chapter.UpdateRsp
	(*chapter.DeleteRsp)(nil),       // 9: com.itcoursee.chapter.DeleteRsp
}
var file_com_itcoursee_message_chapter_proto_depIdxs = []int32{
	0, // 0: com.itcoursee.message.Chapter.Add:input_type -> com.itcoursee.chapter.AddReq
	1, // 1: com.itcoursee.message.Chapter.Get:input_type -> com.itcoursee.chapter.GetReq
	2, // 2: com.itcoursee.message.Chapter.GetsByCourse:input_type -> com.itcoursee.chapter.GetsByCourseReq
	3, // 3: com.itcoursee.message.Chapter.Update:input_type -> com.itcoursee.chapter.UpdateReq
	4, // 4: com.itcoursee.message.Chapter.Delete:input_type -> com.itcoursee.chapter.DeleteReq
	5, // 5: com.itcoursee.message.Chapter.Add:output_type -> com.itcoursee.chapter.AddRsp
	6, // 6: com.itcoursee.message.Chapter.Get:output_type -> com.itcoursee.chapter.GetRsp
	7, // 7: com.itcoursee.message.Chapter.GetsByCourse:output_type -> com.itcoursee.chapter.GetsByCourseRsp
	8, // 8: com.itcoursee.message.Chapter.Update:output_type -> com.itcoursee.chapter.UpdateRsp
	9, // 9: com.itcoursee.message.Chapter.Delete:output_type -> com.itcoursee.chapter.DeleteRsp
	5, // [5:10] is the sub-list for method output_type
	0, // [0:5] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_com_itcoursee_message_chapter_proto_init() }
func file_com_itcoursee_message_chapter_proto_init() {
	if File_com_itcoursee_message_chapter_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_com_itcoursee_message_chapter_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   0,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_com_itcoursee_message_chapter_proto_goTypes,
		DependencyIndexes: file_com_itcoursee_message_chapter_proto_depIdxs,
	}.Build()
	File_com_itcoursee_message_chapter_proto = out.File
	file_com_itcoursee_message_chapter_proto_rawDesc = nil
	file_com_itcoursee_message_chapter_proto_goTypes = nil
	file_com_itcoursee_message_chapter_proto_depIdxs = nil
}
