// Copyright 2025 Google LLC
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.5
// 	protoc        v3.21.12
// source: google/ads/googleads/v19/errors/field_error.proto

package errors

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
	unsafe "unsafe"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

// Enum describing possible field errors.
type FieldErrorEnum_FieldError int32

const (
	// Enum unspecified.
	FieldErrorEnum_UNSPECIFIED FieldErrorEnum_FieldError = 0
	// The received error code is not known in this version.
	FieldErrorEnum_UNKNOWN FieldErrorEnum_FieldError = 1
	// The required field was not present.
	FieldErrorEnum_REQUIRED FieldErrorEnum_FieldError = 2
	// The field attempted to be mutated is immutable.
	FieldErrorEnum_IMMUTABLE_FIELD FieldErrorEnum_FieldError = 3
	// The field's value is invalid.
	FieldErrorEnum_INVALID_VALUE FieldErrorEnum_FieldError = 4
	// The field cannot be set.
	FieldErrorEnum_VALUE_MUST_BE_UNSET FieldErrorEnum_FieldError = 5
	// The required repeated field was empty.
	FieldErrorEnum_REQUIRED_NONEMPTY_LIST FieldErrorEnum_FieldError = 6
	// The field cannot be cleared.
	FieldErrorEnum_FIELD_CANNOT_BE_CLEARED FieldErrorEnum_FieldError = 7
	// The field's value is on a deny-list for this field.
	FieldErrorEnum_BLOCKED_VALUE FieldErrorEnum_FieldError = 9
	// The field's value cannot be modified, except for clearing.
	FieldErrorEnum_FIELD_CAN_ONLY_BE_CLEARED FieldErrorEnum_FieldError = 10
)

// Enum value maps for FieldErrorEnum_FieldError.
var (
	FieldErrorEnum_FieldError_name = map[int32]string{
		0:  "UNSPECIFIED",
		1:  "UNKNOWN",
		2:  "REQUIRED",
		3:  "IMMUTABLE_FIELD",
		4:  "INVALID_VALUE",
		5:  "VALUE_MUST_BE_UNSET",
		6:  "REQUIRED_NONEMPTY_LIST",
		7:  "FIELD_CANNOT_BE_CLEARED",
		9:  "BLOCKED_VALUE",
		10: "FIELD_CAN_ONLY_BE_CLEARED",
	}
	FieldErrorEnum_FieldError_value = map[string]int32{
		"UNSPECIFIED":               0,
		"UNKNOWN":                   1,
		"REQUIRED":                  2,
		"IMMUTABLE_FIELD":           3,
		"INVALID_VALUE":             4,
		"VALUE_MUST_BE_UNSET":       5,
		"REQUIRED_NONEMPTY_LIST":    6,
		"FIELD_CANNOT_BE_CLEARED":   7,
		"BLOCKED_VALUE":             9,
		"FIELD_CAN_ONLY_BE_CLEARED": 10,
	}
)

func (x FieldErrorEnum_FieldError) Enum() *FieldErrorEnum_FieldError {
	p := new(FieldErrorEnum_FieldError)
	*p = x
	return p
}

func (x FieldErrorEnum_FieldError) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (FieldErrorEnum_FieldError) Descriptor() protoreflect.EnumDescriptor {
	return file_errors_field_error_proto_enumTypes[0].Descriptor()
}

func (FieldErrorEnum_FieldError) Type() protoreflect.EnumType {
	return &file_errors_field_error_proto_enumTypes[0]
}

func (x FieldErrorEnum_FieldError) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use FieldErrorEnum_FieldError.Descriptor instead.
func (FieldErrorEnum_FieldError) EnumDescriptor() ([]byte, []int) {
	return file_errors_field_error_proto_rawDescGZIP(), []int{0, 0}
}

// Container for enum describing possible field errors.
type FieldErrorEnum struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *FieldErrorEnum) Reset() {
	*x = FieldErrorEnum{}
	mi := &file_errors_field_error_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *FieldErrorEnum) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FieldErrorEnum) ProtoMessage() {}

func (x *FieldErrorEnum) ProtoReflect() protoreflect.Message {
	mi := &file_errors_field_error_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FieldErrorEnum.ProtoReflect.Descriptor instead.
func (*FieldErrorEnum) Descriptor() ([]byte, []int) {
	return file_errors_field_error_proto_rawDescGZIP(), []int{0}
}

var File_errors_field_error_proto protoreflect.FileDescriptor

var file_errors_field_error_proto_rawDesc = string([]byte{
	0x0a, 0x31, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x64, 0x73, 0x2f, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x61, 0x64, 0x73, 0x2f, 0x76, 0x31, 0x39, 0x2f, 0x65, 0x72, 0x72, 0x6f, 0x72,
	0x73, 0x2f, 0x66, 0x69, 0x65, 0x6c, 0x64, 0x5f, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x12, 0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x61, 0x64, 0x73, 0x2e,
	0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61, 0x64, 0x73, 0x2e, 0x76, 0x31, 0x39, 0x2e, 0x65, 0x72,
	0x72, 0x6f, 0x72, 0x73, 0x22, 0xf7, 0x01, 0x0a, 0x0e, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x45, 0x72,
	0x72, 0x6f, 0x72, 0x45, 0x6e, 0x75, 0x6d, 0x22, 0xe4, 0x01, 0x0a, 0x0a, 0x46, 0x69, 0x65, 0x6c,
	0x64, 0x45, 0x72, 0x72, 0x6f, 0x72, 0x12, 0x0f, 0x0a, 0x0b, 0x55, 0x4e, 0x53, 0x50, 0x45, 0x43,
	0x49, 0x46, 0x49, 0x45, 0x44, 0x10, 0x00, 0x12, 0x0b, 0x0a, 0x07, 0x55, 0x4e, 0x4b, 0x4e, 0x4f,
	0x57, 0x4e, 0x10, 0x01, 0x12, 0x0c, 0x0a, 0x08, 0x52, 0x45, 0x51, 0x55, 0x49, 0x52, 0x45, 0x44,
	0x10, 0x02, 0x12, 0x13, 0x0a, 0x0f, 0x49, 0x4d, 0x4d, 0x55, 0x54, 0x41, 0x42, 0x4c, 0x45, 0x5f,
	0x46, 0x49, 0x45, 0x4c, 0x44, 0x10, 0x03, 0x12, 0x11, 0x0a, 0x0d, 0x49, 0x4e, 0x56, 0x41, 0x4c,
	0x49, 0x44, 0x5f, 0x56, 0x41, 0x4c, 0x55, 0x45, 0x10, 0x04, 0x12, 0x17, 0x0a, 0x13, 0x56, 0x41,
	0x4c, 0x55, 0x45, 0x5f, 0x4d, 0x55, 0x53, 0x54, 0x5f, 0x42, 0x45, 0x5f, 0x55, 0x4e, 0x53, 0x45,
	0x54, 0x10, 0x05, 0x12, 0x1a, 0x0a, 0x16, 0x52, 0x45, 0x51, 0x55, 0x49, 0x52, 0x45, 0x44, 0x5f,
	0x4e, 0x4f, 0x4e, 0x45, 0x4d, 0x50, 0x54, 0x59, 0x5f, 0x4c, 0x49, 0x53, 0x54, 0x10, 0x06, 0x12,
	0x1b, 0x0a, 0x17, 0x46, 0x49, 0x45, 0x4c, 0x44, 0x5f, 0x43, 0x41, 0x4e, 0x4e, 0x4f, 0x54, 0x5f,
	0x42, 0x45, 0x5f, 0x43, 0x4c, 0x45, 0x41, 0x52, 0x45, 0x44, 0x10, 0x07, 0x12, 0x11, 0x0a, 0x0d,
	0x42, 0x4c, 0x4f, 0x43, 0x4b, 0x45, 0x44, 0x5f, 0x56, 0x41, 0x4c, 0x55, 0x45, 0x10, 0x09, 0x12,
	0x1d, 0x0a, 0x19, 0x46, 0x49, 0x45, 0x4c, 0x44, 0x5f, 0x43, 0x41, 0x4e, 0x5f, 0x4f, 0x4e, 0x4c,
	0x59, 0x5f, 0x42, 0x45, 0x5f, 0x43, 0x4c, 0x45, 0x41, 0x52, 0x45, 0x44, 0x10, 0x0a, 0x42, 0xef,
	0x01, 0x0a, 0x23, 0x63, 0x6f, 0x6d, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x61, 0x64,
	0x73, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61, 0x64, 0x73, 0x2e, 0x76, 0x31, 0x39, 0x2e,
	0x65, 0x72, 0x72, 0x6f, 0x72, 0x73, 0x42, 0x0f, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x45, 0x72, 0x72,
	0x6f, 0x72, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a, 0x45, 0x67, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x2e, 0x67, 0x6f, 0x6c, 0x61, 0x6e, 0x67, 0x2e, 0x6f, 0x72, 0x67, 0x2f, 0x67, 0x65, 0x6e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61, 0x70, 0x69, 0x73,
	0x2f, 0x61, 0x64, 0x73, 0x2f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61, 0x64, 0x73, 0x2f, 0x76,
	0x31, 0x39, 0x2f, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x73, 0x3b, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x73,
	0xa2, 0x02, 0x03, 0x47, 0x41, 0x41, 0xaa, 0x02, 0x1f, 0x47, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e,
	0x41, 0x64, 0x73, 0x2e, 0x47, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x41, 0x64, 0x73, 0x2e, 0x56, 0x31,
	0x39, 0x2e, 0x45, 0x72, 0x72, 0x6f, 0x72, 0x73, 0xca, 0x02, 0x1f, 0x47, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x5c, 0x41, 0x64, 0x73, 0x5c, 0x47, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x41, 0x64, 0x73, 0x5c,
	0x56, 0x31, 0x39, 0x5c, 0x45, 0x72, 0x72, 0x6f, 0x72, 0x73, 0xea, 0x02, 0x23, 0x47, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x3a, 0x3a, 0x41, 0x64, 0x73, 0x3a, 0x3a, 0x47, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x41, 0x64, 0x73, 0x3a, 0x3a, 0x56, 0x31, 0x39, 0x3a, 0x3a, 0x45, 0x72, 0x72, 0x6f, 0x72, 0x73,
	0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
})

var (
	file_errors_field_error_proto_rawDescOnce sync.Once
	file_errors_field_error_proto_rawDescData []byte
)

func file_errors_field_error_proto_rawDescGZIP() []byte {
	file_errors_field_error_proto_rawDescOnce.Do(func() {
		file_errors_field_error_proto_rawDescData = protoimpl.X.CompressGZIP(unsafe.Slice(unsafe.StringData(file_errors_field_error_proto_rawDesc), len(file_errors_field_error_proto_rawDesc)))
	})
	return file_errors_field_error_proto_rawDescData
}

var file_errors_field_error_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_errors_field_error_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_errors_field_error_proto_goTypes = []any{
	(FieldErrorEnum_FieldError)(0), // 0: google.ads.googleads.v19.errors.FieldErrorEnum.FieldError
	(*FieldErrorEnum)(nil),         // 1: google.ads.googleads.v19.errors.FieldErrorEnum
}
var file_errors_field_error_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_errors_field_error_proto_init() }
func file_errors_field_error_proto_init() {
	if File_errors_field_error_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: unsafe.Slice(unsafe.StringData(file_errors_field_error_proto_rawDesc), len(file_errors_field_error_proto_rawDesc)),
			NumEnums:      1,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_errors_field_error_proto_goTypes,
		DependencyIndexes: file_errors_field_error_proto_depIdxs,
		EnumInfos:         file_errors_field_error_proto_enumTypes,
		MessageInfos:      file_errors_field_error_proto_msgTypes,
	}.Build()
	File_errors_field_error_proto = out.File
	file_errors_field_error_proto_goTypes = nil
	file_errors_field_error_proto_depIdxs = nil
}
