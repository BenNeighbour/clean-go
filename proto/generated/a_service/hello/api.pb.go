// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.34.2
// 	protoc        (unknown)
// source: a_service/hello/api.proto

package hello

import (
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

var File_a_service_hello_api_proto protoreflect.FileDescriptor

var file_a_service_hello_api_proto_rawDesc = []byte{
	0x0a, 0x19, 0x61, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2f, 0x68, 0x65, 0x6c, 0x6c,
	0x6f, 0x2f, 0x61, 0x70, 0x69, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x05, 0x68, 0x65, 0x6c,
	0x6c, 0x6f, 0x1a, 0x19, 0x61, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2f, 0x68, 0x65,
	0x6c, 0x6c, 0x6f, 0x2f, 0x64, 0x74, 0x6f, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x32, 0x3d, 0x0a,
	0x05, 0x48, 0x65, 0x6c, 0x6c, 0x6f, 0x12, 0x34, 0x0a, 0x08, 0x53, 0x61, 0x79, 0x48, 0x65, 0x6c,
	0x6c, 0x6f, 0x12, 0x13, 0x2e, 0x68, 0x65, 0x6c, 0x6c, 0x6f, 0x2e, 0x48, 0x65, 0x6c, 0x6c, 0x6f,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x11, 0x2e, 0x68, 0x65, 0x6c, 0x6c, 0x6f, 0x2e,
	0x48, 0x65, 0x6c, 0x6c, 0x6f, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x22, 0x00, 0x42, 0x84, 0x01, 0x0a,
	0x09, 0x63, 0x6f, 0x6d, 0x2e, 0x68, 0x65, 0x6c, 0x6c, 0x6f, 0x42, 0x08, 0x41, 0x70, 0x69, 0x50,
	0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a, 0x39, 0x62, 0x65, 0x6e, 0x6e, 0x65, 0x69, 0x67, 0x68,
	0x62, 0x6f, 0x75, 0x72, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65,
	0x73, 0x2f, 0x62, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x2f, 0x61, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2f, 0x68, 0x65, 0x6c, 0x6c,
	0x6f, 0xa2, 0x02, 0x03, 0x48, 0x58, 0x58, 0xaa, 0x02, 0x05, 0x48, 0x65, 0x6c, 0x6c, 0x6f, 0xca,
	0x02, 0x05, 0x48, 0x65, 0x6c, 0x6c, 0x6f, 0xe2, 0x02, 0x11, 0x48, 0x65, 0x6c, 0x6c, 0x6f, 0x5c,
	0x47, 0x50, 0x42, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0xea, 0x02, 0x05, 0x48, 0x65,
	0x6c, 0x6c, 0x6f, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var file_a_service_hello_api_proto_goTypes = []any{
	(*HelloRequest)(nil), // 0: hello.HelloRequest
	(*HelloReply)(nil),   // 1: hello.HelloReply
}
var file_a_service_hello_api_proto_depIdxs = []int32{
	0, // 0: hello.Hello.SayHello:input_type -> hello.HelloRequest
	1, // 1: hello.Hello.SayHello:output_type -> hello.HelloReply
	1, // [1:2] is the sub-list for method output_type
	0, // [0:1] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_a_service_hello_api_proto_init() }
func file_a_service_hello_api_proto_init() {
	if File_a_service_hello_api_proto != nil {
		return
	}
	file_a_service_hello_dto_proto_init()
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_a_service_hello_api_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   0,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_a_service_hello_api_proto_goTypes,
		DependencyIndexes: file_a_service_hello_api_proto_depIdxs,
	}.Build()
	File_a_service_hello_api_proto = out.File
	file_a_service_hello_api_proto_rawDesc = nil
	file_a_service_hello_api_proto_goTypes = nil
	file_a_service_hello_api_proto_depIdxs = nil
}