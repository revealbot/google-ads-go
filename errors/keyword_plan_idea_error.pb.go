// Copyright 2023 Google LLC
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
// 	protoc-gen-go v1.28.1
// 	protoc        v3.21.12
// source: google/ads/googleads/v16/errors/keyword_plan_idea_error.proto

package errors

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

// Enum describing possible errors from KeywordPlanIdeaService.
type KeywordPlanIdeaErrorEnum_KeywordPlanIdeaError int32

const (
	// Enum unspecified.
	KeywordPlanIdeaErrorEnum_UNSPECIFIED KeywordPlanIdeaErrorEnum_KeywordPlanIdeaError = 0
	// The received error code is not known in this version.
	KeywordPlanIdeaErrorEnum_UNKNOWN KeywordPlanIdeaErrorEnum_KeywordPlanIdeaError = 1
	// Error when crawling the input URL.
	KeywordPlanIdeaErrorEnum_URL_CRAWL_ERROR KeywordPlanIdeaErrorEnum_KeywordPlanIdeaError = 2
	// The input has an invalid value.
	KeywordPlanIdeaErrorEnum_INVALID_VALUE KeywordPlanIdeaErrorEnum_KeywordPlanIdeaError = 3
)

// Enum value maps for KeywordPlanIdeaErrorEnum_KeywordPlanIdeaError.
var (
	KeywordPlanIdeaErrorEnum_KeywordPlanIdeaError_name = map[int32]string{
		0: "UNSPECIFIED",
		1: "UNKNOWN",
		2: "URL_CRAWL_ERROR",
		3: "INVALID_VALUE",
	}
	KeywordPlanIdeaErrorEnum_KeywordPlanIdeaError_value = map[string]int32{
		"UNSPECIFIED":     0,
		"UNKNOWN":         1,
		"URL_CRAWL_ERROR": 2,
		"INVALID_VALUE":   3,
	}
)

func (x KeywordPlanIdeaErrorEnum_KeywordPlanIdeaError) Enum() *KeywordPlanIdeaErrorEnum_KeywordPlanIdeaError {
	p := new(KeywordPlanIdeaErrorEnum_KeywordPlanIdeaError)
	*p = x
	return p
}

func (x KeywordPlanIdeaErrorEnum_KeywordPlanIdeaError) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (KeywordPlanIdeaErrorEnum_KeywordPlanIdeaError) Descriptor() protoreflect.EnumDescriptor {
	return file_errors_keyword_plan_idea_error_proto_enumTypes[0].Descriptor()
}

func (KeywordPlanIdeaErrorEnum_KeywordPlanIdeaError) Type() protoreflect.EnumType {
	return &file_errors_keyword_plan_idea_error_proto_enumTypes[0]
}

func (x KeywordPlanIdeaErrorEnum_KeywordPlanIdeaError) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use KeywordPlanIdeaErrorEnum_KeywordPlanIdeaError.Descriptor instead.
func (KeywordPlanIdeaErrorEnum_KeywordPlanIdeaError) EnumDescriptor() ([]byte, []int) {
	return file_errors_keyword_plan_idea_error_proto_rawDescGZIP(), []int{0, 0}
}

// Container for enum describing possible errors from KeywordPlanIdeaService.
type KeywordPlanIdeaErrorEnum struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *KeywordPlanIdeaErrorEnum) Reset() {
	*x = KeywordPlanIdeaErrorEnum{}
	if protoimpl.UnsafeEnabled {
		mi := &file_errors_keyword_plan_idea_error_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *KeywordPlanIdeaErrorEnum) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*KeywordPlanIdeaErrorEnum) ProtoMessage() {}

func (x *KeywordPlanIdeaErrorEnum) ProtoReflect() protoreflect.Message {
	mi := &file_errors_keyword_plan_idea_error_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use KeywordPlanIdeaErrorEnum.ProtoReflect.Descriptor instead.
func (*KeywordPlanIdeaErrorEnum) Descriptor() ([]byte, []int) {
	return file_errors_keyword_plan_idea_error_proto_rawDescGZIP(), []int{0}
}

var File_errors_keyword_plan_idea_error_proto protoreflect.FileDescriptor

var file_errors_keyword_plan_idea_error_proto_rawDesc = []byte{
	0x0a, 0x3d, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x64, 0x73, 0x2f, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x61, 0x64, 0x73, 0x2f, 0x76, 0x31, 0x36, 0x2f, 0x65, 0x72, 0x72, 0x6f, 0x72,
	0x73, 0x2f, 0x6b, 0x65, 0x79, 0x77, 0x6f, 0x72, 0x64, 0x5f, 0x70, 0x6c, 0x61, 0x6e, 0x5f, 0x69,
	0x64, 0x65, 0x61, 0x5f, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12,
	0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x61, 0x64, 0x73, 0x2e, 0x67, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x61, 0x64, 0x73, 0x2e, 0x76, 0x31, 0x36, 0x2e, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x73,
	0x22, 0x78, 0x0a, 0x18, 0x4b, 0x65, 0x79, 0x77, 0x6f, 0x72, 0x64, 0x50, 0x6c, 0x61, 0x6e, 0x49,
	0x64, 0x65, 0x61, 0x45, 0x72, 0x72, 0x6f, 0x72, 0x45, 0x6e, 0x75, 0x6d, 0x22, 0x5c, 0x0a, 0x14,
	0x4b, 0x65, 0x79, 0x77, 0x6f, 0x72, 0x64, 0x50, 0x6c, 0x61, 0x6e, 0x49, 0x64, 0x65, 0x61, 0x45,
	0x72, 0x72, 0x6f, 0x72, 0x12, 0x0f, 0x0a, 0x0b, 0x55, 0x4e, 0x53, 0x50, 0x45, 0x43, 0x49, 0x46,
	0x49, 0x45, 0x44, 0x10, 0x00, 0x12, 0x0b, 0x0a, 0x07, 0x55, 0x4e, 0x4b, 0x4e, 0x4f, 0x57, 0x4e,
	0x10, 0x01, 0x12, 0x13, 0x0a, 0x0f, 0x55, 0x52, 0x4c, 0x5f, 0x43, 0x52, 0x41, 0x57, 0x4c, 0x5f,
	0x45, 0x52, 0x52, 0x4f, 0x52, 0x10, 0x02, 0x12, 0x11, 0x0a, 0x0d, 0x49, 0x4e, 0x56, 0x41, 0x4c,
	0x49, 0x44, 0x5f, 0x56, 0x41, 0x4c, 0x55, 0x45, 0x10, 0x03, 0x42, 0xf9, 0x01, 0x0a, 0x23, 0x63,
	0x6f, 0x6d, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x61, 0x64, 0x73, 0x2e, 0x67, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x61, 0x64, 0x73, 0x2e, 0x76, 0x31, 0x36, 0x2e, 0x65, 0x72, 0x72, 0x6f,
	0x72, 0x73, 0x42, 0x19, 0x4b, 0x65, 0x79, 0x77, 0x6f, 0x72, 0x64, 0x50, 0x6c, 0x61, 0x6e, 0x49,
	0x64, 0x65, 0x61, 0x45, 0x72, 0x72, 0x6f, 0x72, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a,
	0x45, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x67, 0x6f, 0x6c, 0x61, 0x6e, 0x67, 0x2e, 0x6f,
	0x72, 0x67, 0x2f, 0x67, 0x65, 0x6e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x67, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x61, 0x70, 0x69, 0x73, 0x2f, 0x61, 0x64, 0x73, 0x2f, 0x67, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x61, 0x64, 0x73, 0x2f, 0x76, 0x31, 0x36, 0x2f, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x73, 0x3b,
	0x65, 0x72, 0x72, 0x6f, 0x72, 0x73, 0xa2, 0x02, 0x03, 0x47, 0x41, 0x41, 0xaa, 0x02, 0x1f, 0x47,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x41, 0x64, 0x73, 0x2e, 0x47, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x41, 0x64, 0x73, 0x2e, 0x56, 0x31, 0x36, 0x2e, 0x45, 0x72, 0x72, 0x6f, 0x72, 0x73, 0xca, 0x02,
	0x1f, 0x47, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x5c, 0x41, 0x64, 0x73, 0x5c, 0x47, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x41, 0x64, 0x73, 0x5c, 0x56, 0x31, 0x36, 0x5c, 0x45, 0x72, 0x72, 0x6f, 0x72, 0x73,
	0xea, 0x02, 0x23, 0x47, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x3a, 0x3a, 0x41, 0x64, 0x73, 0x3a, 0x3a,
	0x47, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x41, 0x64, 0x73, 0x3a, 0x3a, 0x56, 0x31, 0x36, 0x3a, 0x3a,
	0x45, 0x72, 0x72, 0x6f, 0x72, 0x73, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_errors_keyword_plan_idea_error_proto_rawDescOnce sync.Once
	file_errors_keyword_plan_idea_error_proto_rawDescData = file_errors_keyword_plan_idea_error_proto_rawDesc
)

func file_errors_keyword_plan_idea_error_proto_rawDescGZIP() []byte {
	file_errors_keyword_plan_idea_error_proto_rawDescOnce.Do(func() {
		file_errors_keyword_plan_idea_error_proto_rawDescData = protoimpl.X.CompressGZIP(file_errors_keyword_plan_idea_error_proto_rawDescData)
	})
	return file_errors_keyword_plan_idea_error_proto_rawDescData
}

var file_errors_keyword_plan_idea_error_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_errors_keyword_plan_idea_error_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_errors_keyword_plan_idea_error_proto_goTypes = []interface{}{
	(KeywordPlanIdeaErrorEnum_KeywordPlanIdeaError)(0), // 0: google.ads.googleads.v16.errors.KeywordPlanIdeaErrorEnum.KeywordPlanIdeaError
	(*KeywordPlanIdeaErrorEnum)(nil),                   // 1: google.ads.googleads.v16.errors.KeywordPlanIdeaErrorEnum
}
var file_errors_keyword_plan_idea_error_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_errors_keyword_plan_idea_error_proto_init() }
func file_errors_keyword_plan_idea_error_proto_init() {
	if File_errors_keyword_plan_idea_error_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_errors_keyword_plan_idea_error_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*KeywordPlanIdeaErrorEnum); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_errors_keyword_plan_idea_error_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_errors_keyword_plan_idea_error_proto_goTypes,
		DependencyIndexes: file_errors_keyword_plan_idea_error_proto_depIdxs,
		EnumInfos:         file_errors_keyword_plan_idea_error_proto_enumTypes,
		MessageInfos:      file_errors_keyword_plan_idea_error_proto_msgTypes,
	}.Build()
	File_errors_keyword_plan_idea_error_proto = out.File
	file_errors_keyword_plan_idea_error_proto_rawDesc = nil
	file_errors_keyword_plan_idea_error_proto_goTypes = nil
	file_errors_keyword_plan_idea_error_proto_depIdxs = nil
}
