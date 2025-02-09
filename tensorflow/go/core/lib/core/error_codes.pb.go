// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0-devel
// 	protoc        v3.15.8
// source: tensorflow/core/lib/core/error_codes.proto

package core

import (
	for_core_protos_go_proto "github.com/tensorflow/tensorflow/tensorflow/go/core/protobuf/for_core_protos_go_proto"
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

// Symbols defined in public import of tensorflow/core/protobuf/error_codes.proto.

type Code = for_core_protos_go_proto.Code

const Code_OK = for_core_protos_go_proto.Code_OK
const Code_CANCELLED = for_core_protos_go_proto.Code_CANCELLED
const Code_UNKNOWN = for_core_protos_go_proto.Code_UNKNOWN
const Code_INVALID_ARGUMENT = for_core_protos_go_proto.Code_INVALID_ARGUMENT
const Code_DEADLINE_EXCEEDED = for_core_protos_go_proto.Code_DEADLINE_EXCEEDED
const Code_NOT_FOUND = for_core_protos_go_proto.Code_NOT_FOUND
const Code_ALREADY_EXISTS = for_core_protos_go_proto.Code_ALREADY_EXISTS
const Code_PERMISSION_DENIED = for_core_protos_go_proto.Code_PERMISSION_DENIED
const Code_UNAUTHENTICATED = for_core_protos_go_proto.Code_UNAUTHENTICATED
const Code_RESOURCE_EXHAUSTED = for_core_protos_go_proto.Code_RESOURCE_EXHAUSTED
const Code_FAILED_PRECONDITION = for_core_protos_go_proto.Code_FAILED_PRECONDITION
const Code_ABORTED = for_core_protos_go_proto.Code_ABORTED
const Code_OUT_OF_RANGE = for_core_protos_go_proto.Code_OUT_OF_RANGE
const Code_UNIMPLEMENTED = for_core_protos_go_proto.Code_UNIMPLEMENTED
const Code_INTERNAL = for_core_protos_go_proto.Code_INTERNAL
const Code_UNAVAILABLE = for_core_protos_go_proto.Code_UNAVAILABLE
const Code_DATA_LOSS = for_core_protos_go_proto.Code_DATA_LOSS
const Code_DO_NOT_USE_RESERVED_FOR_FUTURE_EXPANSION_USE_DEFAULT_IN_SWITCH_INSTEAD_ = for_core_protos_go_proto.Code_DO_NOT_USE_RESERVED_FOR_FUTURE_EXPANSION_USE_DEFAULT_IN_SWITCH_INSTEAD_

var Code_name = for_core_protos_go_proto.Code_name
var Code_value = for_core_protos_go_proto.Code_value

var File_tensorflow_core_lib_core_error_codes_proto protoreflect.FileDescriptor

var file_tensorflow_core_lib_core_error_codes_proto_rawDesc = []byte{
	0x0a, 0x2a, 0x74, 0x65, 0x6e, 0x73, 0x6f, 0x72, 0x66, 0x6c, 0x6f, 0x77, 0x2f, 0x63, 0x6f, 0x72,
	0x65, 0x2f, 0x6c, 0x69, 0x62, 0x2f, 0x63, 0x6f, 0x72, 0x65, 0x2f, 0x65, 0x72, 0x72, 0x6f, 0x72,
	0x5f, 0x63, 0x6f, 0x64, 0x65, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x2a, 0x74, 0x65,
	0x6e, 0x73, 0x6f, 0x72, 0x66, 0x6c, 0x6f, 0x77, 0x2f, 0x63, 0x6f, 0x72, 0x65, 0x2f, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x5f, 0x63, 0x6f, 0x64,
	0x65, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x00, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x33,
}

var file_tensorflow_core_lib_core_error_codes_proto_goTypes = []interface{}{}
var file_tensorflow_core_lib_core_error_codes_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_tensorflow_core_lib_core_error_codes_proto_init() }
func file_tensorflow_core_lib_core_error_codes_proto_init() {
	if File_tensorflow_core_lib_core_error_codes_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_tensorflow_core_lib_core_error_codes_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   0,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_tensorflow_core_lib_core_error_codes_proto_goTypes,
		DependencyIndexes: file_tensorflow_core_lib_core_error_codes_proto_depIdxs,
	}.Build()
	File_tensorflow_core_lib_core_error_codes_proto = out.File
	file_tensorflow_core_lib_core_error_codes_proto_rawDesc = nil
	file_tensorflow_core_lib_core_error_codes_proto_goTypes = nil
	file_tensorflow_core_lib_core_error_codes_proto_depIdxs = nil
}
