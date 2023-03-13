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
// 	protoc        v3.21.12
// source: google/ads/googleads/v13/errors/feed_item_set_link_error.proto

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

// Enum describing possible feed item set link errors.
type FeedItemSetLinkErrorEnum_FeedItemSetLinkError int32

const (
	// Enum unspecified.
	FeedItemSetLinkErrorEnum_UNSPECIFIED FeedItemSetLinkErrorEnum_FeedItemSetLinkError = 0
	// The received error code is not known in this version.
	FeedItemSetLinkErrorEnum_UNKNOWN FeedItemSetLinkErrorEnum_FeedItemSetLinkError = 1
	// The feed IDs of the FeedItemSet and FeedItem do not match. Only FeedItems
	// in a given Feed can be linked to a FeedItemSet in that Feed.
	FeedItemSetLinkErrorEnum_FEED_ID_MISMATCH FeedItemSetLinkErrorEnum_FeedItemSetLinkError = 2
	// Cannot add or remove links to a dynamic set.
	FeedItemSetLinkErrorEnum_NO_MUTATE_ALLOWED_FOR_DYNAMIC_SET FeedItemSetLinkErrorEnum_FeedItemSetLinkError = 3
)

// Enum value maps for FeedItemSetLinkErrorEnum_FeedItemSetLinkError.
var (
	FeedItemSetLinkErrorEnum_FeedItemSetLinkError_name = map[int32]string{
		0: "UNSPECIFIED",
		1: "UNKNOWN",
		2: "FEED_ID_MISMATCH",
		3: "NO_MUTATE_ALLOWED_FOR_DYNAMIC_SET",
	}
	FeedItemSetLinkErrorEnum_FeedItemSetLinkError_value = map[string]int32{
		"UNSPECIFIED":                       0,
		"UNKNOWN":                           1,
		"FEED_ID_MISMATCH":                  2,
		"NO_MUTATE_ALLOWED_FOR_DYNAMIC_SET": 3,
	}
)

func (x FeedItemSetLinkErrorEnum_FeedItemSetLinkError) Enum() *FeedItemSetLinkErrorEnum_FeedItemSetLinkError {
	p := new(FeedItemSetLinkErrorEnum_FeedItemSetLinkError)
	*p = x
	return p
}

func (x FeedItemSetLinkErrorEnum_FeedItemSetLinkError) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (FeedItemSetLinkErrorEnum_FeedItemSetLinkError) Descriptor() protoreflect.EnumDescriptor {
	return file_errors_feed_item_set_link_error_proto_enumTypes[0].Descriptor()
}

func (FeedItemSetLinkErrorEnum_FeedItemSetLinkError) Type() protoreflect.EnumType {
	return &file_errors_feed_item_set_link_error_proto_enumTypes[0]
}

func (x FeedItemSetLinkErrorEnum_FeedItemSetLinkError) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use FeedItemSetLinkErrorEnum_FeedItemSetLinkError.Descriptor instead.
func (FeedItemSetLinkErrorEnum_FeedItemSetLinkError) EnumDescriptor() ([]byte, []int) {
	return file_errors_feed_item_set_link_error_proto_rawDescGZIP(), []int{0, 0}
}

// Container for enum describing possible feed item set link errors.
type FeedItemSetLinkErrorEnum struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *FeedItemSetLinkErrorEnum) Reset() {
	*x = FeedItemSetLinkErrorEnum{}
	if protoimpl.UnsafeEnabled {
		mi := &file_errors_feed_item_set_link_error_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FeedItemSetLinkErrorEnum) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FeedItemSetLinkErrorEnum) ProtoMessage() {}

func (x *FeedItemSetLinkErrorEnum) ProtoReflect() protoreflect.Message {
	mi := &file_errors_feed_item_set_link_error_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FeedItemSetLinkErrorEnum.ProtoReflect.Descriptor instead.
func (*FeedItemSetLinkErrorEnum) Descriptor() ([]byte, []int) {
	return file_errors_feed_item_set_link_error_proto_rawDescGZIP(), []int{0}
}

var File_errors_feed_item_set_link_error_proto protoreflect.FileDescriptor

var file_errors_feed_item_set_link_error_proto_rawDesc = []byte{
	0x0a, 0x3e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x64, 0x73, 0x2f, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x61, 0x64, 0x73, 0x2f, 0x76, 0x31, 0x33, 0x2f, 0x65, 0x72, 0x72, 0x6f, 0x72,
	0x73, 0x2f, 0x66, 0x65, 0x65, 0x64, 0x5f, 0x69, 0x74, 0x65, 0x6d, 0x5f, 0x73, 0x65, 0x74, 0x5f,
	0x6c, 0x69, 0x6e, 0x6b, 0x5f, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x12, 0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x61, 0x64, 0x73, 0x2e, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x61, 0x64, 0x73, 0x2e, 0x76, 0x31, 0x33, 0x2e, 0x65, 0x72, 0x72, 0x6f, 0x72,
	0x73, 0x22, 0x8d, 0x01, 0x0a, 0x18, 0x46, 0x65, 0x65, 0x64, 0x49, 0x74, 0x65, 0x6d, 0x53, 0x65,
	0x74, 0x4c, 0x69, 0x6e, 0x6b, 0x45, 0x72, 0x72, 0x6f, 0x72, 0x45, 0x6e, 0x75, 0x6d, 0x22, 0x71,
	0x0a, 0x14, 0x46, 0x65, 0x65, 0x64, 0x49, 0x74, 0x65, 0x6d, 0x53, 0x65, 0x74, 0x4c, 0x69, 0x6e,
	0x6b, 0x45, 0x72, 0x72, 0x6f, 0x72, 0x12, 0x0f, 0x0a, 0x0b, 0x55, 0x4e, 0x53, 0x50, 0x45, 0x43,
	0x49, 0x46, 0x49, 0x45, 0x44, 0x10, 0x00, 0x12, 0x0b, 0x0a, 0x07, 0x55, 0x4e, 0x4b, 0x4e, 0x4f,
	0x57, 0x4e, 0x10, 0x01, 0x12, 0x14, 0x0a, 0x10, 0x46, 0x45, 0x45, 0x44, 0x5f, 0x49, 0x44, 0x5f,
	0x4d, 0x49, 0x53, 0x4d, 0x41, 0x54, 0x43, 0x48, 0x10, 0x02, 0x12, 0x25, 0x0a, 0x21, 0x4e, 0x4f,
	0x5f, 0x4d, 0x55, 0x54, 0x41, 0x54, 0x45, 0x5f, 0x41, 0x4c, 0x4c, 0x4f, 0x57, 0x45, 0x44, 0x5f,
	0x46, 0x4f, 0x52, 0x5f, 0x44, 0x59, 0x4e, 0x41, 0x4d, 0x49, 0x43, 0x5f, 0x53, 0x45, 0x54, 0x10,
	0x03, 0x42, 0xf9, 0x01, 0x0a, 0x23, 0x63, 0x6f, 0x6d, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x2e, 0x61, 0x64, 0x73, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61, 0x64, 0x73, 0x2e, 0x76,
	0x31, 0x33, 0x2e, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x73, 0x42, 0x19, 0x46, 0x65, 0x65, 0x64, 0x49,
	0x74, 0x65, 0x6d, 0x53, 0x65, 0x74, 0x4c, 0x69, 0x6e, 0x6b, 0x45, 0x72, 0x72, 0x6f, 0x72, 0x50,
	0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a, 0x45, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x67,
	0x6f, 0x6c, 0x61, 0x6e, 0x67, 0x2e, 0x6f, 0x72, 0x67, 0x2f, 0x67, 0x65, 0x6e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x2f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61, 0x70, 0x69, 0x73, 0x2f, 0x61, 0x64,
	0x73, 0x2f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61, 0x64, 0x73, 0x2f, 0x76, 0x31, 0x33, 0x2f,
	0x65, 0x72, 0x72, 0x6f, 0x72, 0x73, 0x3b, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x73, 0xa2, 0x02, 0x03,
	0x47, 0x41, 0x41, 0xaa, 0x02, 0x1f, 0x47, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x41, 0x64, 0x73,
	0x2e, 0x47, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x41, 0x64, 0x73, 0x2e, 0x56, 0x31, 0x33, 0x2e, 0x45,
	0x72, 0x72, 0x6f, 0x72, 0x73, 0xca, 0x02, 0x1f, 0x47, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x5c, 0x41,
	0x64, 0x73, 0x5c, 0x47, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x41, 0x64, 0x73, 0x5c, 0x56, 0x31, 0x33,
	0x5c, 0x45, 0x72, 0x72, 0x6f, 0x72, 0x73, 0xea, 0x02, 0x23, 0x47, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x3a, 0x3a, 0x41, 0x64, 0x73, 0x3a, 0x3a, 0x47, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x41, 0x64, 0x73,
	0x3a, 0x3a, 0x56, 0x31, 0x33, 0x3a, 0x3a, 0x45, 0x72, 0x72, 0x6f, 0x72, 0x73, 0x62, 0x06, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_errors_feed_item_set_link_error_proto_rawDescOnce sync.Once
	file_errors_feed_item_set_link_error_proto_rawDescData = file_errors_feed_item_set_link_error_proto_rawDesc
)

func file_errors_feed_item_set_link_error_proto_rawDescGZIP() []byte {
	file_errors_feed_item_set_link_error_proto_rawDescOnce.Do(func() {
		file_errors_feed_item_set_link_error_proto_rawDescData = protoimpl.X.CompressGZIP(file_errors_feed_item_set_link_error_proto_rawDescData)
	})
	return file_errors_feed_item_set_link_error_proto_rawDescData
}

var file_errors_feed_item_set_link_error_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_errors_feed_item_set_link_error_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_errors_feed_item_set_link_error_proto_goTypes = []interface{}{
	(FeedItemSetLinkErrorEnum_FeedItemSetLinkError)(0), // 0: google.ads.googleads.v13.errors.FeedItemSetLinkErrorEnum.FeedItemSetLinkError
	(*FeedItemSetLinkErrorEnum)(nil),                   // 1: google.ads.googleads.v13.errors.FeedItemSetLinkErrorEnum
}
var file_errors_feed_item_set_link_error_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_errors_feed_item_set_link_error_proto_init() }
func file_errors_feed_item_set_link_error_proto_init() {
	if File_errors_feed_item_set_link_error_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_errors_feed_item_set_link_error_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*FeedItemSetLinkErrorEnum); i {
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
			RawDescriptor: file_errors_feed_item_set_link_error_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_errors_feed_item_set_link_error_proto_goTypes,
		DependencyIndexes: file_errors_feed_item_set_link_error_proto_depIdxs,
		EnumInfos:         file_errors_feed_item_set_link_error_proto_enumTypes,
		MessageInfos:      file_errors_feed_item_set_link_error_proto_msgTypes,
	}.Build()
	File_errors_feed_item_set_link_error_proto = out.File
	file_errors_feed_item_set_link_error_proto_rawDesc = nil
	file_errors_feed_item_set_link_error_proto_goTypes = nil
	file_errors_feed_item_set_link_error_proto_depIdxs = nil
}
