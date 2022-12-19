// Copyright 2022 Google LLC
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
// 	protoc-gen-go v1.26.0
// 	protoc        v3.21.7
// source: google/ads/googleads/v12/errors/feed_item_error.proto

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

// Enum describing possible feed item errors.
type FeedItemErrorEnum_FeedItemError int32

const (
	// Enum unspecified.
	FeedItemErrorEnum_UNSPECIFIED FeedItemErrorEnum_FeedItemError = 0
	// The received error code is not known in this version.
	FeedItemErrorEnum_UNKNOWN FeedItemErrorEnum_FeedItemError = 1
	// Cannot convert the feed attribute value from string to its real type.
	FeedItemErrorEnum_CANNOT_CONVERT_ATTRIBUTE_VALUE_FROM_STRING FeedItemErrorEnum_FeedItemError = 2
	// Cannot operate on removed feed item.
	FeedItemErrorEnum_CANNOT_OPERATE_ON_REMOVED_FEED_ITEM FeedItemErrorEnum_FeedItemError = 3
	// Date time zone does not match the account's time zone.
	FeedItemErrorEnum_DATE_TIME_MUST_BE_IN_ACCOUNT_TIME_ZONE FeedItemErrorEnum_FeedItemError = 4
	// Feed item with the key attributes could not be found.
	FeedItemErrorEnum_KEY_ATTRIBUTES_NOT_FOUND FeedItemErrorEnum_FeedItemError = 5
	// Url feed attribute value is not valid.
	FeedItemErrorEnum_INVALID_URL FeedItemErrorEnum_FeedItemError = 6
	// Some key attributes are missing.
	FeedItemErrorEnum_MISSING_KEY_ATTRIBUTES FeedItemErrorEnum_FeedItemError = 7
	// Feed item has same key attributes as another feed item.
	FeedItemErrorEnum_KEY_ATTRIBUTES_NOT_UNIQUE FeedItemErrorEnum_FeedItemError = 8
	// Cannot modify key attributes on an existing feed item.
	FeedItemErrorEnum_CANNOT_MODIFY_KEY_ATTRIBUTE_VALUE FeedItemErrorEnum_FeedItemError = 9
	// The feed attribute value is too large.
	FeedItemErrorEnum_SIZE_TOO_LARGE_FOR_MULTI_VALUE_ATTRIBUTE FeedItemErrorEnum_FeedItemError = 10
	// Feed is read only.
	FeedItemErrorEnum_LEGACY_FEED_TYPE_READ_ONLY FeedItemErrorEnum_FeedItemError = 11
)

// Enum value maps for FeedItemErrorEnum_FeedItemError.
var (
	FeedItemErrorEnum_FeedItemError_name = map[int32]string{
		0:  "UNSPECIFIED",
		1:  "UNKNOWN",
		2:  "CANNOT_CONVERT_ATTRIBUTE_VALUE_FROM_STRING",
		3:  "CANNOT_OPERATE_ON_REMOVED_FEED_ITEM",
		4:  "DATE_TIME_MUST_BE_IN_ACCOUNT_TIME_ZONE",
		5:  "KEY_ATTRIBUTES_NOT_FOUND",
		6:  "INVALID_URL",
		7:  "MISSING_KEY_ATTRIBUTES",
		8:  "KEY_ATTRIBUTES_NOT_UNIQUE",
		9:  "CANNOT_MODIFY_KEY_ATTRIBUTE_VALUE",
		10: "SIZE_TOO_LARGE_FOR_MULTI_VALUE_ATTRIBUTE",
		11: "LEGACY_FEED_TYPE_READ_ONLY",
	}
	FeedItemErrorEnum_FeedItemError_value = map[string]int32{
		"UNSPECIFIED": 0,
		"UNKNOWN":     1,
		"CANNOT_CONVERT_ATTRIBUTE_VALUE_FROM_STRING": 2,
		"CANNOT_OPERATE_ON_REMOVED_FEED_ITEM":        3,
		"DATE_TIME_MUST_BE_IN_ACCOUNT_TIME_ZONE":     4,
		"KEY_ATTRIBUTES_NOT_FOUND":                   5,
		"INVALID_URL":                                6,
		"MISSING_KEY_ATTRIBUTES":                     7,
		"KEY_ATTRIBUTES_NOT_UNIQUE":                  8,
		"CANNOT_MODIFY_KEY_ATTRIBUTE_VALUE":          9,
		"SIZE_TOO_LARGE_FOR_MULTI_VALUE_ATTRIBUTE":   10,
		"LEGACY_FEED_TYPE_READ_ONLY":                 11,
	}
)

func (x FeedItemErrorEnum_FeedItemError) Enum() *FeedItemErrorEnum_FeedItemError {
	p := new(FeedItemErrorEnum_FeedItemError)
	*p = x
	return p
}

func (x FeedItemErrorEnum_FeedItemError) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (FeedItemErrorEnum_FeedItemError) Descriptor() protoreflect.EnumDescriptor {
	return file_errors_feed_item_error_proto_enumTypes[0].Descriptor()
}

func (FeedItemErrorEnum_FeedItemError) Type() protoreflect.EnumType {
	return &file_errors_feed_item_error_proto_enumTypes[0]
}

func (x FeedItemErrorEnum_FeedItemError) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use FeedItemErrorEnum_FeedItemError.Descriptor instead.
func (FeedItemErrorEnum_FeedItemError) EnumDescriptor() ([]byte, []int) {
	return file_errors_feed_item_error_proto_rawDescGZIP(), []int{0, 0}
}

// Container for enum describing possible feed item errors.
type FeedItemErrorEnum struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *FeedItemErrorEnum) Reset() {
	*x = FeedItemErrorEnum{}
	if protoimpl.UnsafeEnabled {
		mi := &file_errors_feed_item_error_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FeedItemErrorEnum) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FeedItemErrorEnum) ProtoMessage() {}

func (x *FeedItemErrorEnum) ProtoReflect() protoreflect.Message {
	mi := &file_errors_feed_item_error_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FeedItemErrorEnum.ProtoReflect.Descriptor instead.
func (*FeedItemErrorEnum) Descriptor() ([]byte, []int) {
	return file_errors_feed_item_error_proto_rawDescGZIP(), []int{0}
}

var File_errors_feed_item_error_proto protoreflect.FileDescriptor

var file_errors_feed_item_error_proto_rawDesc = []byte{
	0x0a, 0x35, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x64, 0x73, 0x2f, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x61, 0x64, 0x73, 0x2f, 0x76, 0x31, 0x32, 0x2f, 0x65, 0x72, 0x72, 0x6f, 0x72,
	0x73, 0x2f, 0x66, 0x65, 0x65, 0x64, 0x5f, 0x69, 0x74, 0x65, 0x6d, 0x5f, 0x65, 0x72, 0x72, 0x6f,
	0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e,
	0x61, 0x64, 0x73, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61, 0x64, 0x73, 0x2e, 0x76, 0x31,
	0x32, 0x2e, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x73, 0x22, 0xa7, 0x03, 0x0a, 0x11, 0x46, 0x65, 0x65,
	0x64, 0x49, 0x74, 0x65, 0x6d, 0x45, 0x72, 0x72, 0x6f, 0x72, 0x45, 0x6e, 0x75, 0x6d, 0x22, 0x91,
	0x03, 0x0a, 0x0d, 0x46, 0x65, 0x65, 0x64, 0x49, 0x74, 0x65, 0x6d, 0x45, 0x72, 0x72, 0x6f, 0x72,
	0x12, 0x0f, 0x0a, 0x0b, 0x55, 0x4e, 0x53, 0x50, 0x45, 0x43, 0x49, 0x46, 0x49, 0x45, 0x44, 0x10,
	0x00, 0x12, 0x0b, 0x0a, 0x07, 0x55, 0x4e, 0x4b, 0x4e, 0x4f, 0x57, 0x4e, 0x10, 0x01, 0x12, 0x2e,
	0x0a, 0x2a, 0x43, 0x41, 0x4e, 0x4e, 0x4f, 0x54, 0x5f, 0x43, 0x4f, 0x4e, 0x56, 0x45, 0x52, 0x54,
	0x5f, 0x41, 0x54, 0x54, 0x52, 0x49, 0x42, 0x55, 0x54, 0x45, 0x5f, 0x56, 0x41, 0x4c, 0x55, 0x45,
	0x5f, 0x46, 0x52, 0x4f, 0x4d, 0x5f, 0x53, 0x54, 0x52, 0x49, 0x4e, 0x47, 0x10, 0x02, 0x12, 0x27,
	0x0a, 0x23, 0x43, 0x41, 0x4e, 0x4e, 0x4f, 0x54, 0x5f, 0x4f, 0x50, 0x45, 0x52, 0x41, 0x54, 0x45,
	0x5f, 0x4f, 0x4e, 0x5f, 0x52, 0x45, 0x4d, 0x4f, 0x56, 0x45, 0x44, 0x5f, 0x46, 0x45, 0x45, 0x44,
	0x5f, 0x49, 0x54, 0x45, 0x4d, 0x10, 0x03, 0x12, 0x2a, 0x0a, 0x26, 0x44, 0x41, 0x54, 0x45, 0x5f,
	0x54, 0x49, 0x4d, 0x45, 0x5f, 0x4d, 0x55, 0x53, 0x54, 0x5f, 0x42, 0x45, 0x5f, 0x49, 0x4e, 0x5f,
	0x41, 0x43, 0x43, 0x4f, 0x55, 0x4e, 0x54, 0x5f, 0x54, 0x49, 0x4d, 0x45, 0x5f, 0x5a, 0x4f, 0x4e,
	0x45, 0x10, 0x04, 0x12, 0x1c, 0x0a, 0x18, 0x4b, 0x45, 0x59, 0x5f, 0x41, 0x54, 0x54, 0x52, 0x49,
	0x42, 0x55, 0x54, 0x45, 0x53, 0x5f, 0x4e, 0x4f, 0x54, 0x5f, 0x46, 0x4f, 0x55, 0x4e, 0x44, 0x10,
	0x05, 0x12, 0x0f, 0x0a, 0x0b, 0x49, 0x4e, 0x56, 0x41, 0x4c, 0x49, 0x44, 0x5f, 0x55, 0x52, 0x4c,
	0x10, 0x06, 0x12, 0x1a, 0x0a, 0x16, 0x4d, 0x49, 0x53, 0x53, 0x49, 0x4e, 0x47, 0x5f, 0x4b, 0x45,
	0x59, 0x5f, 0x41, 0x54, 0x54, 0x52, 0x49, 0x42, 0x55, 0x54, 0x45, 0x53, 0x10, 0x07, 0x12, 0x1d,
	0x0a, 0x19, 0x4b, 0x45, 0x59, 0x5f, 0x41, 0x54, 0x54, 0x52, 0x49, 0x42, 0x55, 0x54, 0x45, 0x53,
	0x5f, 0x4e, 0x4f, 0x54, 0x5f, 0x55, 0x4e, 0x49, 0x51, 0x55, 0x45, 0x10, 0x08, 0x12, 0x25, 0x0a,
	0x21, 0x43, 0x41, 0x4e, 0x4e, 0x4f, 0x54, 0x5f, 0x4d, 0x4f, 0x44, 0x49, 0x46, 0x59, 0x5f, 0x4b,
	0x45, 0x59, 0x5f, 0x41, 0x54, 0x54, 0x52, 0x49, 0x42, 0x55, 0x54, 0x45, 0x5f, 0x56, 0x41, 0x4c,
	0x55, 0x45, 0x10, 0x09, 0x12, 0x2c, 0x0a, 0x28, 0x53, 0x49, 0x5a, 0x45, 0x5f, 0x54, 0x4f, 0x4f,
	0x5f, 0x4c, 0x41, 0x52, 0x47, 0x45, 0x5f, 0x46, 0x4f, 0x52, 0x5f, 0x4d, 0x55, 0x4c, 0x54, 0x49,
	0x5f, 0x56, 0x41, 0x4c, 0x55, 0x45, 0x5f, 0x41, 0x54, 0x54, 0x52, 0x49, 0x42, 0x55, 0x54, 0x45,
	0x10, 0x0a, 0x12, 0x1e, 0x0a, 0x1a, 0x4c, 0x45, 0x47, 0x41, 0x43, 0x59, 0x5f, 0x46, 0x45, 0x45,
	0x44, 0x5f, 0x54, 0x59, 0x50, 0x45, 0x5f, 0x52, 0x45, 0x41, 0x44, 0x5f, 0x4f, 0x4e, 0x4c, 0x59,
	0x10, 0x0b, 0x42, 0xf2, 0x01, 0x0a, 0x23, 0x63, 0x6f, 0x6d, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x2e, 0x61, 0x64, 0x73, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61, 0x64, 0x73, 0x2e,
	0x76, 0x31, 0x32, 0x2e, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x73, 0x42, 0x12, 0x46, 0x65, 0x65, 0x64,
	0x49, 0x74, 0x65, 0x6d, 0x45, 0x72, 0x72, 0x6f, 0x72, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01,
	0x5a, 0x45, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x67, 0x6f, 0x6c, 0x61, 0x6e, 0x67, 0x2e,
	0x6f, 0x72, 0x67, 0x2f, 0x67, 0x65, 0x6e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x61, 0x70, 0x69, 0x73, 0x2f, 0x61, 0x64, 0x73, 0x2f, 0x67, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x61, 0x64, 0x73, 0x2f, 0x76, 0x31, 0x32, 0x2f, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x73,
	0x3b, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x73, 0xa2, 0x02, 0x03, 0x47, 0x41, 0x41, 0xaa, 0x02, 0x1f,
	0x47, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x41, 0x64, 0x73, 0x2e, 0x47, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x41, 0x64, 0x73, 0x2e, 0x56, 0x31, 0x32, 0x2e, 0x45, 0x72, 0x72, 0x6f, 0x72, 0x73, 0xca,
	0x02, 0x1f, 0x47, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x5c, 0x41, 0x64, 0x73, 0x5c, 0x47, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x41, 0x64, 0x73, 0x5c, 0x56, 0x31, 0x32, 0x5c, 0x45, 0x72, 0x72, 0x6f, 0x72,
	0x73, 0xea, 0x02, 0x23, 0x47, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x3a, 0x3a, 0x41, 0x64, 0x73, 0x3a,
	0x3a, 0x47, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x41, 0x64, 0x73, 0x3a, 0x3a, 0x56, 0x31, 0x32, 0x3a,
	0x3a, 0x45, 0x72, 0x72, 0x6f, 0x72, 0x73, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_errors_feed_item_error_proto_rawDescOnce sync.Once
	file_errors_feed_item_error_proto_rawDescData = file_errors_feed_item_error_proto_rawDesc
)

func file_errors_feed_item_error_proto_rawDescGZIP() []byte {
	file_errors_feed_item_error_proto_rawDescOnce.Do(func() {
		file_errors_feed_item_error_proto_rawDescData = protoimpl.X.CompressGZIP(file_errors_feed_item_error_proto_rawDescData)
	})
	return file_errors_feed_item_error_proto_rawDescData
}

var file_errors_feed_item_error_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_errors_feed_item_error_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_errors_feed_item_error_proto_goTypes = []interface{}{
	(FeedItemErrorEnum_FeedItemError)(0), // 0: google.ads.googleads.v12.errors.FeedItemErrorEnum.FeedItemError
	(*FeedItemErrorEnum)(nil),            // 1: google.ads.googleads.v12.errors.FeedItemErrorEnum
}
var file_errors_feed_item_error_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_errors_feed_item_error_proto_init() }
func file_errors_feed_item_error_proto_init() {
	if File_errors_feed_item_error_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_errors_feed_item_error_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*FeedItemErrorEnum); i {
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
			RawDescriptor: file_errors_feed_item_error_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_errors_feed_item_error_proto_goTypes,
		DependencyIndexes: file_errors_feed_item_error_proto_depIdxs,
		EnumInfos:         file_errors_feed_item_error_proto_enumTypes,
		MessageInfos:      file_errors_feed_item_error_proto_msgTypes,
	}.Build()
	File_errors_feed_item_error_proto = out.File
	file_errors_feed_item_error_proto_rawDesc = nil
	file_errors_feed_item_error_proto_goTypes = nil
	file_errors_feed_item_error_proto_depIdxs = nil
}
