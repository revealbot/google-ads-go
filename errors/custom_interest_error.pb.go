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
// source: google/ads/googleads/v19/errors/custom_interest_error.proto

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

// Enum describing possible custom interest errors.
type CustomInterestErrorEnum_CustomInterestError int32

const (
	// Enum unspecified.
	CustomInterestErrorEnum_UNSPECIFIED CustomInterestErrorEnum_CustomInterestError = 0
	// The received error code is not known in this version.
	CustomInterestErrorEnum_UNKNOWN CustomInterestErrorEnum_CustomInterestError = 1
	// Duplicate custom interest name ignoring case.
	CustomInterestErrorEnum_NAME_ALREADY_USED CustomInterestErrorEnum_CustomInterestError = 2
	// In the remove custom interest member operation, both member ID and
	// pair [type, parameter] are not present.
	CustomInterestErrorEnum_CUSTOM_INTEREST_MEMBER_ID_AND_TYPE_PARAMETER_NOT_PRESENT_IN_REMOVE CustomInterestErrorEnum_CustomInterestError = 3
	// The pair of [type, parameter] does not exist.
	CustomInterestErrorEnum_TYPE_AND_PARAMETER_NOT_FOUND CustomInterestErrorEnum_CustomInterestError = 4
	// The pair of [type, parameter] already exists.
	CustomInterestErrorEnum_TYPE_AND_PARAMETER_ALREADY_EXISTED CustomInterestErrorEnum_CustomInterestError = 5
	// Unsupported custom interest member type.
	CustomInterestErrorEnum_INVALID_CUSTOM_INTEREST_MEMBER_TYPE CustomInterestErrorEnum_CustomInterestError = 6
	// Cannot remove a custom interest while it's still being targeted.
	CustomInterestErrorEnum_CANNOT_REMOVE_WHILE_IN_USE CustomInterestErrorEnum_CustomInterestError = 7
	// Cannot mutate custom interest type.
	CustomInterestErrorEnum_CANNOT_CHANGE_TYPE CustomInterestErrorEnum_CustomInterestError = 8
)

// Enum value maps for CustomInterestErrorEnum_CustomInterestError.
var (
	CustomInterestErrorEnum_CustomInterestError_name = map[int32]string{
		0: "UNSPECIFIED",
		1: "UNKNOWN",
		2: "NAME_ALREADY_USED",
		3: "CUSTOM_INTEREST_MEMBER_ID_AND_TYPE_PARAMETER_NOT_PRESENT_IN_REMOVE",
		4: "TYPE_AND_PARAMETER_NOT_FOUND",
		5: "TYPE_AND_PARAMETER_ALREADY_EXISTED",
		6: "INVALID_CUSTOM_INTEREST_MEMBER_TYPE",
		7: "CANNOT_REMOVE_WHILE_IN_USE",
		8: "CANNOT_CHANGE_TYPE",
	}
	CustomInterestErrorEnum_CustomInterestError_value = map[string]int32{
		"UNSPECIFIED":       0,
		"UNKNOWN":           1,
		"NAME_ALREADY_USED": 2,
		"CUSTOM_INTEREST_MEMBER_ID_AND_TYPE_PARAMETER_NOT_PRESENT_IN_REMOVE": 3,
		"TYPE_AND_PARAMETER_NOT_FOUND":                                       4,
		"TYPE_AND_PARAMETER_ALREADY_EXISTED":                                 5,
		"INVALID_CUSTOM_INTEREST_MEMBER_TYPE":                                6,
		"CANNOT_REMOVE_WHILE_IN_USE":                                         7,
		"CANNOT_CHANGE_TYPE":                                                 8,
	}
)

func (x CustomInterestErrorEnum_CustomInterestError) Enum() *CustomInterestErrorEnum_CustomInterestError {
	p := new(CustomInterestErrorEnum_CustomInterestError)
	*p = x
	return p
}

func (x CustomInterestErrorEnum_CustomInterestError) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (CustomInterestErrorEnum_CustomInterestError) Descriptor() protoreflect.EnumDescriptor {
	return file_errors_custom_interest_error_proto_enumTypes[0].Descriptor()
}

func (CustomInterestErrorEnum_CustomInterestError) Type() protoreflect.EnumType {
	return &file_errors_custom_interest_error_proto_enumTypes[0]
}

func (x CustomInterestErrorEnum_CustomInterestError) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use CustomInterestErrorEnum_CustomInterestError.Descriptor instead.
func (CustomInterestErrorEnum_CustomInterestError) EnumDescriptor() ([]byte, []int) {
	return file_errors_custom_interest_error_proto_rawDescGZIP(), []int{0, 0}
}

// Container for enum describing possible custom interest errors.
type CustomInterestErrorEnum struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *CustomInterestErrorEnum) Reset() {
	*x = CustomInterestErrorEnum{}
	mi := &file_errors_custom_interest_error_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *CustomInterestErrorEnum) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CustomInterestErrorEnum) ProtoMessage() {}

func (x *CustomInterestErrorEnum) ProtoReflect() protoreflect.Message {
	mi := &file_errors_custom_interest_error_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CustomInterestErrorEnum.ProtoReflect.Descriptor instead.
func (*CustomInterestErrorEnum) Descriptor() ([]byte, []int) {
	return file_errors_custom_interest_error_proto_rawDescGZIP(), []int{0}
}

var File_errors_custom_interest_error_proto protoreflect.FileDescriptor

var file_errors_custom_interest_error_proto_rawDesc = string([]byte{
	0x0a, 0x3b, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x64, 0x73, 0x2f, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x61, 0x64, 0x73, 0x2f, 0x76, 0x31, 0x39, 0x2f, 0x65, 0x72, 0x72, 0x6f, 0x72,
	0x73, 0x2f, 0x63, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x5f, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x65, 0x73,
	0x74, 0x5f, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x1f, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x61, 0x64, 0x73, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x61, 0x64, 0x73, 0x2e, 0x76, 0x31, 0x39, 0x2e, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x73, 0x22, 0xd9,
	0x02, 0x0a, 0x17, 0x43, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x49, 0x6e, 0x74, 0x65, 0x72, 0x65, 0x73,
	0x74, 0x45, 0x72, 0x72, 0x6f, 0x72, 0x45, 0x6e, 0x75, 0x6d, 0x22, 0xbd, 0x02, 0x0a, 0x13, 0x43,
	0x75, 0x73, 0x74, 0x6f, 0x6d, 0x49, 0x6e, 0x74, 0x65, 0x72, 0x65, 0x73, 0x74, 0x45, 0x72, 0x72,
	0x6f, 0x72, 0x12, 0x0f, 0x0a, 0x0b, 0x55, 0x4e, 0x53, 0x50, 0x45, 0x43, 0x49, 0x46, 0x49, 0x45,
	0x44, 0x10, 0x00, 0x12, 0x0b, 0x0a, 0x07, 0x55, 0x4e, 0x4b, 0x4e, 0x4f, 0x57, 0x4e, 0x10, 0x01,
	0x12, 0x15, 0x0a, 0x11, 0x4e, 0x41, 0x4d, 0x45, 0x5f, 0x41, 0x4c, 0x52, 0x45, 0x41, 0x44, 0x59,
	0x5f, 0x55, 0x53, 0x45, 0x44, 0x10, 0x02, 0x12, 0x46, 0x0a, 0x42, 0x43, 0x55, 0x53, 0x54, 0x4f,
	0x4d, 0x5f, 0x49, 0x4e, 0x54, 0x45, 0x52, 0x45, 0x53, 0x54, 0x5f, 0x4d, 0x45, 0x4d, 0x42, 0x45,
	0x52, 0x5f, 0x49, 0x44, 0x5f, 0x41, 0x4e, 0x44, 0x5f, 0x54, 0x59, 0x50, 0x45, 0x5f, 0x50, 0x41,
	0x52, 0x41, 0x4d, 0x45, 0x54, 0x45, 0x52, 0x5f, 0x4e, 0x4f, 0x54, 0x5f, 0x50, 0x52, 0x45, 0x53,
	0x45, 0x4e, 0x54, 0x5f, 0x49, 0x4e, 0x5f, 0x52, 0x45, 0x4d, 0x4f, 0x56, 0x45, 0x10, 0x03, 0x12,
	0x20, 0x0a, 0x1c, 0x54, 0x59, 0x50, 0x45, 0x5f, 0x41, 0x4e, 0x44, 0x5f, 0x50, 0x41, 0x52, 0x41,
	0x4d, 0x45, 0x54, 0x45, 0x52, 0x5f, 0x4e, 0x4f, 0x54, 0x5f, 0x46, 0x4f, 0x55, 0x4e, 0x44, 0x10,
	0x04, 0x12, 0x26, 0x0a, 0x22, 0x54, 0x59, 0x50, 0x45, 0x5f, 0x41, 0x4e, 0x44, 0x5f, 0x50, 0x41,
	0x52, 0x41, 0x4d, 0x45, 0x54, 0x45, 0x52, 0x5f, 0x41, 0x4c, 0x52, 0x45, 0x41, 0x44, 0x59, 0x5f,
	0x45, 0x58, 0x49, 0x53, 0x54, 0x45, 0x44, 0x10, 0x05, 0x12, 0x27, 0x0a, 0x23, 0x49, 0x4e, 0x56,
	0x41, 0x4c, 0x49, 0x44, 0x5f, 0x43, 0x55, 0x53, 0x54, 0x4f, 0x4d, 0x5f, 0x49, 0x4e, 0x54, 0x45,
	0x52, 0x45, 0x53, 0x54, 0x5f, 0x4d, 0x45, 0x4d, 0x42, 0x45, 0x52, 0x5f, 0x54, 0x59, 0x50, 0x45,
	0x10, 0x06, 0x12, 0x1e, 0x0a, 0x1a, 0x43, 0x41, 0x4e, 0x4e, 0x4f, 0x54, 0x5f, 0x52, 0x45, 0x4d,
	0x4f, 0x56, 0x45, 0x5f, 0x57, 0x48, 0x49, 0x4c, 0x45, 0x5f, 0x49, 0x4e, 0x5f, 0x55, 0x53, 0x45,
	0x10, 0x07, 0x12, 0x16, 0x0a, 0x12, 0x43, 0x41, 0x4e, 0x4e, 0x4f, 0x54, 0x5f, 0x43, 0x48, 0x41,
	0x4e, 0x47, 0x45, 0x5f, 0x54, 0x59, 0x50, 0x45, 0x10, 0x08, 0x42, 0xf8, 0x01, 0x0a, 0x23, 0x63,
	0x6f, 0x6d, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x61, 0x64, 0x73, 0x2e, 0x67, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x61, 0x64, 0x73, 0x2e, 0x76, 0x31, 0x39, 0x2e, 0x65, 0x72, 0x72, 0x6f,
	0x72, 0x73, 0x42, 0x18, 0x43, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x49, 0x6e, 0x74, 0x65, 0x72, 0x65,
	0x73, 0x74, 0x45, 0x72, 0x72, 0x6f, 0x72, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a, 0x45,
	0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x67, 0x6f, 0x6c, 0x61, 0x6e, 0x67, 0x2e, 0x6f, 0x72,
	0x67, 0x2f, 0x67, 0x65, 0x6e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x67, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x61, 0x70, 0x69, 0x73, 0x2f, 0x61, 0x64, 0x73, 0x2f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x61, 0x64, 0x73, 0x2f, 0x76, 0x31, 0x39, 0x2f, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x73, 0x3b, 0x65,
	0x72, 0x72, 0x6f, 0x72, 0x73, 0xa2, 0x02, 0x03, 0x47, 0x41, 0x41, 0xaa, 0x02, 0x1f, 0x47, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x41, 0x64, 0x73, 0x2e, 0x47, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x41,
	0x64, 0x73, 0x2e, 0x56, 0x31, 0x39, 0x2e, 0x45, 0x72, 0x72, 0x6f, 0x72, 0x73, 0xca, 0x02, 0x1f,
	0x47, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x5c, 0x41, 0x64, 0x73, 0x5c, 0x47, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x41, 0x64, 0x73, 0x5c, 0x56, 0x31, 0x39, 0x5c, 0x45, 0x72, 0x72, 0x6f, 0x72, 0x73, 0xea,
	0x02, 0x23, 0x47, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x3a, 0x3a, 0x41, 0x64, 0x73, 0x3a, 0x3a, 0x47,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x41, 0x64, 0x73, 0x3a, 0x3a, 0x56, 0x31, 0x39, 0x3a, 0x3a, 0x45,
	0x72, 0x72, 0x6f, 0x72, 0x73, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
})

var (
	file_errors_custom_interest_error_proto_rawDescOnce sync.Once
	file_errors_custom_interest_error_proto_rawDescData []byte
)

func file_errors_custom_interest_error_proto_rawDescGZIP() []byte {
	file_errors_custom_interest_error_proto_rawDescOnce.Do(func() {
		file_errors_custom_interest_error_proto_rawDescData = protoimpl.X.CompressGZIP(unsafe.Slice(unsafe.StringData(file_errors_custom_interest_error_proto_rawDesc), len(file_errors_custom_interest_error_proto_rawDesc)))
	})
	return file_errors_custom_interest_error_proto_rawDescData
}

var file_errors_custom_interest_error_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_errors_custom_interest_error_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_errors_custom_interest_error_proto_goTypes = []any{
	(CustomInterestErrorEnum_CustomInterestError)(0), // 0: google.ads.googleads.v19.errors.CustomInterestErrorEnum.CustomInterestError
	(*CustomInterestErrorEnum)(nil),                  // 1: google.ads.googleads.v19.errors.CustomInterestErrorEnum
}
var file_errors_custom_interest_error_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_errors_custom_interest_error_proto_init() }
func file_errors_custom_interest_error_proto_init() {
	if File_errors_custom_interest_error_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: unsafe.Slice(unsafe.StringData(file_errors_custom_interest_error_proto_rawDesc), len(file_errors_custom_interest_error_proto_rawDesc)),
			NumEnums:      1,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_errors_custom_interest_error_proto_goTypes,
		DependencyIndexes: file_errors_custom_interest_error_proto_depIdxs,
		EnumInfos:         file_errors_custom_interest_error_proto_enumTypes,
		MessageInfos:      file_errors_custom_interest_error_proto_msgTypes,
	}.Build()
	File_errors_custom_interest_error_proto = out.File
	file_errors_custom_interest_error_proto_goTypes = nil
	file_errors_custom_interest_error_proto_depIdxs = nil
}
