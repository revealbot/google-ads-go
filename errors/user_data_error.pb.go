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
// source: google/ads/googleads/v19/errors/user_data_error.proto

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

// Enum describing possible request errors.
type UserDataErrorEnum_UserDataError int32

const (
	// Enum unspecified.
	UserDataErrorEnum_UNSPECIFIED UserDataErrorEnum_UserDataError = 0
	// The received error code is not known in this version.
	UserDataErrorEnum_UNKNOWN UserDataErrorEnum_UserDataError = 1
	// Customer is not allowed to perform operations related to Customer Match.
	UserDataErrorEnum_OPERATIONS_FOR_CUSTOMER_MATCH_NOT_ALLOWED UserDataErrorEnum_UserDataError = 2
	// Maximum number of user identifiers allowed for each request is 100 and
	// for each operation is 20.
	UserDataErrorEnum_TOO_MANY_USER_IDENTIFIERS UserDataErrorEnum_UserDataError = 3
	// Current user list is not applicable for the given customer.
	UserDataErrorEnum_USER_LIST_NOT_APPLICABLE UserDataErrorEnum_UserDataError = 4
)

// Enum value maps for UserDataErrorEnum_UserDataError.
var (
	UserDataErrorEnum_UserDataError_name = map[int32]string{
		0: "UNSPECIFIED",
		1: "UNKNOWN",
		2: "OPERATIONS_FOR_CUSTOMER_MATCH_NOT_ALLOWED",
		3: "TOO_MANY_USER_IDENTIFIERS",
		4: "USER_LIST_NOT_APPLICABLE",
	}
	UserDataErrorEnum_UserDataError_value = map[string]int32{
		"UNSPECIFIED": 0,
		"UNKNOWN":     1,
		"OPERATIONS_FOR_CUSTOMER_MATCH_NOT_ALLOWED": 2,
		"TOO_MANY_USER_IDENTIFIERS":                 3,
		"USER_LIST_NOT_APPLICABLE":                  4,
	}
)

func (x UserDataErrorEnum_UserDataError) Enum() *UserDataErrorEnum_UserDataError {
	p := new(UserDataErrorEnum_UserDataError)
	*p = x
	return p
}

func (x UserDataErrorEnum_UserDataError) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (UserDataErrorEnum_UserDataError) Descriptor() protoreflect.EnumDescriptor {
	return file_errors_user_data_error_proto_enumTypes[0].Descriptor()
}

func (UserDataErrorEnum_UserDataError) Type() protoreflect.EnumType {
	return &file_errors_user_data_error_proto_enumTypes[0]
}

func (x UserDataErrorEnum_UserDataError) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use UserDataErrorEnum_UserDataError.Descriptor instead.
func (UserDataErrorEnum_UserDataError) EnumDescriptor() ([]byte, []int) {
	return file_errors_user_data_error_proto_rawDescGZIP(), []int{0, 0}
}

// Container for enum describing possible user data errors.
type UserDataErrorEnum struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *UserDataErrorEnum) Reset() {
	*x = UserDataErrorEnum{}
	mi := &file_errors_user_data_error_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *UserDataErrorEnum) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UserDataErrorEnum) ProtoMessage() {}

func (x *UserDataErrorEnum) ProtoReflect() protoreflect.Message {
	mi := &file_errors_user_data_error_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UserDataErrorEnum.ProtoReflect.Descriptor instead.
func (*UserDataErrorEnum) Descriptor() ([]byte, []int) {
	return file_errors_user_data_error_proto_rawDescGZIP(), []int{0}
}

var File_errors_user_data_error_proto protoreflect.FileDescriptor

var file_errors_user_data_error_proto_rawDesc = string([]byte{
	0x0a, 0x35, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x64, 0x73, 0x2f, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x61, 0x64, 0x73, 0x2f, 0x76, 0x31, 0x39, 0x2f, 0x65, 0x72, 0x72, 0x6f, 0x72,
	0x73, 0x2f, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x64, 0x61, 0x74, 0x61, 0x5f, 0x65, 0x72, 0x72, 0x6f,
	0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e,
	0x61, 0x64, 0x73, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61, 0x64, 0x73, 0x2e, 0x76, 0x31,
	0x39, 0x2e, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x73, 0x22, 0xaf, 0x01, 0x0a, 0x11, 0x55, 0x73, 0x65,
	0x72, 0x44, 0x61, 0x74, 0x61, 0x45, 0x72, 0x72, 0x6f, 0x72, 0x45, 0x6e, 0x75, 0x6d, 0x22, 0x99,
	0x01, 0x0a, 0x0d, 0x55, 0x73, 0x65, 0x72, 0x44, 0x61, 0x74, 0x61, 0x45, 0x72, 0x72, 0x6f, 0x72,
	0x12, 0x0f, 0x0a, 0x0b, 0x55, 0x4e, 0x53, 0x50, 0x45, 0x43, 0x49, 0x46, 0x49, 0x45, 0x44, 0x10,
	0x00, 0x12, 0x0b, 0x0a, 0x07, 0x55, 0x4e, 0x4b, 0x4e, 0x4f, 0x57, 0x4e, 0x10, 0x01, 0x12, 0x2d,
	0x0a, 0x29, 0x4f, 0x50, 0x45, 0x52, 0x41, 0x54, 0x49, 0x4f, 0x4e, 0x53, 0x5f, 0x46, 0x4f, 0x52,
	0x5f, 0x43, 0x55, 0x53, 0x54, 0x4f, 0x4d, 0x45, 0x52, 0x5f, 0x4d, 0x41, 0x54, 0x43, 0x48, 0x5f,
	0x4e, 0x4f, 0x54, 0x5f, 0x41, 0x4c, 0x4c, 0x4f, 0x57, 0x45, 0x44, 0x10, 0x02, 0x12, 0x1d, 0x0a,
	0x19, 0x54, 0x4f, 0x4f, 0x5f, 0x4d, 0x41, 0x4e, 0x59, 0x5f, 0x55, 0x53, 0x45, 0x52, 0x5f, 0x49,
	0x44, 0x45, 0x4e, 0x54, 0x49, 0x46, 0x49, 0x45, 0x52, 0x53, 0x10, 0x03, 0x12, 0x1c, 0x0a, 0x18,
	0x55, 0x53, 0x45, 0x52, 0x5f, 0x4c, 0x49, 0x53, 0x54, 0x5f, 0x4e, 0x4f, 0x54, 0x5f, 0x41, 0x50,
	0x50, 0x4c, 0x49, 0x43, 0x41, 0x42, 0x4c, 0x45, 0x10, 0x04, 0x42, 0xf2, 0x01, 0x0a, 0x23, 0x63,
	0x6f, 0x6d, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x61, 0x64, 0x73, 0x2e, 0x67, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x61, 0x64, 0x73, 0x2e, 0x76, 0x31, 0x39, 0x2e, 0x65, 0x72, 0x72, 0x6f,
	0x72, 0x73, 0x42, 0x12, 0x55, 0x73, 0x65, 0x72, 0x44, 0x61, 0x74, 0x61, 0x45, 0x72, 0x72, 0x6f,
	0x72, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a, 0x45, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x2e, 0x67, 0x6f, 0x6c, 0x61, 0x6e, 0x67, 0x2e, 0x6f, 0x72, 0x67, 0x2f, 0x67, 0x65, 0x6e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61, 0x70, 0x69, 0x73, 0x2f,
	0x61, 0x64, 0x73, 0x2f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61, 0x64, 0x73, 0x2f, 0x76, 0x31,
	0x39, 0x2f, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x73, 0x3b, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x73, 0xa2,
	0x02, 0x03, 0x47, 0x41, 0x41, 0xaa, 0x02, 0x1f, 0x47, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x41,
	0x64, 0x73, 0x2e, 0x47, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x41, 0x64, 0x73, 0x2e, 0x56, 0x31, 0x39,
	0x2e, 0x45, 0x72, 0x72, 0x6f, 0x72, 0x73, 0xca, 0x02, 0x1f, 0x47, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x5c, 0x41, 0x64, 0x73, 0x5c, 0x47, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x41, 0x64, 0x73, 0x5c, 0x56,
	0x31, 0x39, 0x5c, 0x45, 0x72, 0x72, 0x6f, 0x72, 0x73, 0xea, 0x02, 0x23, 0x47, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x3a, 0x3a, 0x41, 0x64, 0x73, 0x3a, 0x3a, 0x47, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x41,
	0x64, 0x73, 0x3a, 0x3a, 0x56, 0x31, 0x39, 0x3a, 0x3a, 0x45, 0x72, 0x72, 0x6f, 0x72, 0x73, 0x62,
	0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
})

var (
	file_errors_user_data_error_proto_rawDescOnce sync.Once
	file_errors_user_data_error_proto_rawDescData []byte
)

func file_errors_user_data_error_proto_rawDescGZIP() []byte {
	file_errors_user_data_error_proto_rawDescOnce.Do(func() {
		file_errors_user_data_error_proto_rawDescData = protoimpl.X.CompressGZIP(unsafe.Slice(unsafe.StringData(file_errors_user_data_error_proto_rawDesc), len(file_errors_user_data_error_proto_rawDesc)))
	})
	return file_errors_user_data_error_proto_rawDescData
}

var file_errors_user_data_error_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_errors_user_data_error_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_errors_user_data_error_proto_goTypes = []any{
	(UserDataErrorEnum_UserDataError)(0), // 0: google.ads.googleads.v19.errors.UserDataErrorEnum.UserDataError
	(*UserDataErrorEnum)(nil),            // 1: google.ads.googleads.v19.errors.UserDataErrorEnum
}
var file_errors_user_data_error_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_errors_user_data_error_proto_init() }
func file_errors_user_data_error_proto_init() {
	if File_errors_user_data_error_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: unsafe.Slice(unsafe.StringData(file_errors_user_data_error_proto_rawDesc), len(file_errors_user_data_error_proto_rawDesc)),
			NumEnums:      1,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_errors_user_data_error_proto_goTypes,
		DependencyIndexes: file_errors_user_data_error_proto_depIdxs,
		EnumInfos:         file_errors_user_data_error_proto_enumTypes,
		MessageInfos:      file_errors_user_data_error_proto_msgTypes,
	}.Build()
	File_errors_user_data_error_proto = out.File
	file_errors_user_data_error_proto_goTypes = nil
	file_errors_user_data_error_proto_depIdxs = nil
}
