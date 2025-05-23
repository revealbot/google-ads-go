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
// source: google/ads/googleads/v19/errors/asset_group_asset_error.proto

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

// Enum describing possible asset group asset errors.
type AssetGroupAssetErrorEnum_AssetGroupAssetError int32

const (
	// Enum unspecified.
	AssetGroupAssetErrorEnum_UNSPECIFIED AssetGroupAssetErrorEnum_AssetGroupAssetError = 0
	// The received error code is not known in this version.
	AssetGroupAssetErrorEnum_UNKNOWN AssetGroupAssetErrorEnum_AssetGroupAssetError = 1
	// Cannot add duplicated asset group asset.
	AssetGroupAssetErrorEnum_DUPLICATE_RESOURCE AssetGroupAssetErrorEnum_AssetGroupAssetError = 2
	// Expandable tags are not allowed in description assets.
	AssetGroupAssetErrorEnum_EXPANDABLE_TAGS_NOT_ALLOWED_IN_DESCRIPTION AssetGroupAssetErrorEnum_AssetGroupAssetError = 3
	// Ad customizers are not supported in assetgroup's text assets.
	AssetGroupAssetErrorEnum_AD_CUSTOMIZER_NOT_SUPPORTED AssetGroupAssetErrorEnum_AssetGroupAssetError = 4
	// Cannot add a HotelPropertyAsset to an AssetGroup that isn't linked
	// to the parent campaign's hotel_property_asset_set field.
	AssetGroupAssetErrorEnum_HOTEL_PROPERTY_ASSET_NOT_LINKED_TO_CAMPAIGN AssetGroupAssetErrorEnum_AssetGroupAssetError = 5
)

// Enum value maps for AssetGroupAssetErrorEnum_AssetGroupAssetError.
var (
	AssetGroupAssetErrorEnum_AssetGroupAssetError_name = map[int32]string{
		0: "UNSPECIFIED",
		1: "UNKNOWN",
		2: "DUPLICATE_RESOURCE",
		3: "EXPANDABLE_TAGS_NOT_ALLOWED_IN_DESCRIPTION",
		4: "AD_CUSTOMIZER_NOT_SUPPORTED",
		5: "HOTEL_PROPERTY_ASSET_NOT_LINKED_TO_CAMPAIGN",
	}
	AssetGroupAssetErrorEnum_AssetGroupAssetError_value = map[string]int32{
		"UNSPECIFIED":        0,
		"UNKNOWN":            1,
		"DUPLICATE_RESOURCE": 2,
		"EXPANDABLE_TAGS_NOT_ALLOWED_IN_DESCRIPTION":  3,
		"AD_CUSTOMIZER_NOT_SUPPORTED":                 4,
		"HOTEL_PROPERTY_ASSET_NOT_LINKED_TO_CAMPAIGN": 5,
	}
)

func (x AssetGroupAssetErrorEnum_AssetGroupAssetError) Enum() *AssetGroupAssetErrorEnum_AssetGroupAssetError {
	p := new(AssetGroupAssetErrorEnum_AssetGroupAssetError)
	*p = x
	return p
}

func (x AssetGroupAssetErrorEnum_AssetGroupAssetError) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (AssetGroupAssetErrorEnum_AssetGroupAssetError) Descriptor() protoreflect.EnumDescriptor {
	return file_errors_asset_group_asset_error_proto_enumTypes[0].Descriptor()
}

func (AssetGroupAssetErrorEnum_AssetGroupAssetError) Type() protoreflect.EnumType {
	return &file_errors_asset_group_asset_error_proto_enumTypes[0]
}

func (x AssetGroupAssetErrorEnum_AssetGroupAssetError) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use AssetGroupAssetErrorEnum_AssetGroupAssetError.Descriptor instead.
func (AssetGroupAssetErrorEnum_AssetGroupAssetError) EnumDescriptor() ([]byte, []int) {
	return file_errors_asset_group_asset_error_proto_rawDescGZIP(), []int{0, 0}
}

// Container for enum describing possible asset group asset errors.
type AssetGroupAssetErrorEnum struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *AssetGroupAssetErrorEnum) Reset() {
	*x = AssetGroupAssetErrorEnum{}
	mi := &file_errors_asset_group_asset_error_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *AssetGroupAssetErrorEnum) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AssetGroupAssetErrorEnum) ProtoMessage() {}

func (x *AssetGroupAssetErrorEnum) ProtoReflect() protoreflect.Message {
	mi := &file_errors_asset_group_asset_error_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AssetGroupAssetErrorEnum.ProtoReflect.Descriptor instead.
func (*AssetGroupAssetErrorEnum) Descriptor() ([]byte, []int) {
	return file_errors_asset_group_asset_error_proto_rawDescGZIP(), []int{0}
}

var File_errors_asset_group_asset_error_proto protoreflect.FileDescriptor

var file_errors_asset_group_asset_error_proto_rawDesc = string([]byte{
	0x0a, 0x3d, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x64, 0x73, 0x2f, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x61, 0x64, 0x73, 0x2f, 0x76, 0x31, 0x39, 0x2f, 0x65, 0x72, 0x72, 0x6f, 0x72,
	0x73, 0x2f, 0x61, 0x73, 0x73, 0x65, 0x74, 0x5f, 0x67, 0x72, 0x6f, 0x75, 0x70, 0x5f, 0x61, 0x73,
	0x73, 0x65, 0x74, 0x5f, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12,
	0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x61, 0x64, 0x73, 0x2e, 0x67, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x61, 0x64, 0x73, 0x2e, 0x76, 0x31, 0x39, 0x2e, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x73,
	0x22, 0xeb, 0x01, 0x0a, 0x18, 0x41, 0x73, 0x73, 0x65, 0x74, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x41,
	0x73, 0x73, 0x65, 0x74, 0x45, 0x72, 0x72, 0x6f, 0x72, 0x45, 0x6e, 0x75, 0x6d, 0x22, 0xce, 0x01,
	0x0a, 0x14, 0x41, 0x73, 0x73, 0x65, 0x74, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x41, 0x73, 0x73, 0x65,
	0x74, 0x45, 0x72, 0x72, 0x6f, 0x72, 0x12, 0x0f, 0x0a, 0x0b, 0x55, 0x4e, 0x53, 0x50, 0x45, 0x43,
	0x49, 0x46, 0x49, 0x45, 0x44, 0x10, 0x00, 0x12, 0x0b, 0x0a, 0x07, 0x55, 0x4e, 0x4b, 0x4e, 0x4f,
	0x57, 0x4e, 0x10, 0x01, 0x12, 0x16, 0x0a, 0x12, 0x44, 0x55, 0x50, 0x4c, 0x49, 0x43, 0x41, 0x54,
	0x45, 0x5f, 0x52, 0x45, 0x53, 0x4f, 0x55, 0x52, 0x43, 0x45, 0x10, 0x02, 0x12, 0x2e, 0x0a, 0x2a,
	0x45, 0x58, 0x50, 0x41, 0x4e, 0x44, 0x41, 0x42, 0x4c, 0x45, 0x5f, 0x54, 0x41, 0x47, 0x53, 0x5f,
	0x4e, 0x4f, 0x54, 0x5f, 0x41, 0x4c, 0x4c, 0x4f, 0x57, 0x45, 0x44, 0x5f, 0x49, 0x4e, 0x5f, 0x44,
	0x45, 0x53, 0x43, 0x52, 0x49, 0x50, 0x54, 0x49, 0x4f, 0x4e, 0x10, 0x03, 0x12, 0x1f, 0x0a, 0x1b,
	0x41, 0x44, 0x5f, 0x43, 0x55, 0x53, 0x54, 0x4f, 0x4d, 0x49, 0x5a, 0x45, 0x52, 0x5f, 0x4e, 0x4f,
	0x54, 0x5f, 0x53, 0x55, 0x50, 0x50, 0x4f, 0x52, 0x54, 0x45, 0x44, 0x10, 0x04, 0x12, 0x2f, 0x0a,
	0x2b, 0x48, 0x4f, 0x54, 0x45, 0x4c, 0x5f, 0x50, 0x52, 0x4f, 0x50, 0x45, 0x52, 0x54, 0x59, 0x5f,
	0x41, 0x53, 0x53, 0x45, 0x54, 0x5f, 0x4e, 0x4f, 0x54, 0x5f, 0x4c, 0x49, 0x4e, 0x4b, 0x45, 0x44,
	0x5f, 0x54, 0x4f, 0x5f, 0x43, 0x41, 0x4d, 0x50, 0x41, 0x49, 0x47, 0x4e, 0x10, 0x05, 0x42, 0xf9,
	0x01, 0x0a, 0x23, 0x63, 0x6f, 0x6d, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x61, 0x64,
	0x73, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61, 0x64, 0x73, 0x2e, 0x76, 0x31, 0x39, 0x2e,
	0x65, 0x72, 0x72, 0x6f, 0x72, 0x73, 0x42, 0x19, 0x41, 0x73, 0x73, 0x65, 0x74, 0x47, 0x72, 0x6f,
	0x75, 0x70, 0x41, 0x73, 0x73, 0x65, 0x74, 0x45, 0x72, 0x72, 0x6f, 0x72, 0x50, 0x72, 0x6f, 0x74,
	0x6f, 0x50, 0x01, 0x5a, 0x45, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x67, 0x6f, 0x6c, 0x61,
	0x6e, 0x67, 0x2e, 0x6f, 0x72, 0x67, 0x2f, 0x67, 0x65, 0x6e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f,
	0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61, 0x70, 0x69, 0x73, 0x2f, 0x61, 0x64, 0x73, 0x2f, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61, 0x64, 0x73, 0x2f, 0x76, 0x31, 0x39, 0x2f, 0x65, 0x72, 0x72,
	0x6f, 0x72, 0x73, 0x3b, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x73, 0xa2, 0x02, 0x03, 0x47, 0x41, 0x41,
	0xaa, 0x02, 0x1f, 0x47, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x41, 0x64, 0x73, 0x2e, 0x47, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x41, 0x64, 0x73, 0x2e, 0x56, 0x31, 0x39, 0x2e, 0x45, 0x72, 0x72, 0x6f,
	0x72, 0x73, 0xca, 0x02, 0x1f, 0x47, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x5c, 0x41, 0x64, 0x73, 0x5c,
	0x47, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x41, 0x64, 0x73, 0x5c, 0x56, 0x31, 0x39, 0x5c, 0x45, 0x72,
	0x72, 0x6f, 0x72, 0x73, 0xea, 0x02, 0x23, 0x47, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x3a, 0x3a, 0x41,
	0x64, 0x73, 0x3a, 0x3a, 0x47, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x41, 0x64, 0x73, 0x3a, 0x3a, 0x56,
	0x31, 0x39, 0x3a, 0x3a, 0x45, 0x72, 0x72, 0x6f, 0x72, 0x73, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x33,
})

var (
	file_errors_asset_group_asset_error_proto_rawDescOnce sync.Once
	file_errors_asset_group_asset_error_proto_rawDescData []byte
)

func file_errors_asset_group_asset_error_proto_rawDescGZIP() []byte {
	file_errors_asset_group_asset_error_proto_rawDescOnce.Do(func() {
		file_errors_asset_group_asset_error_proto_rawDescData = protoimpl.X.CompressGZIP(unsafe.Slice(unsafe.StringData(file_errors_asset_group_asset_error_proto_rawDesc), len(file_errors_asset_group_asset_error_proto_rawDesc)))
	})
	return file_errors_asset_group_asset_error_proto_rawDescData
}

var file_errors_asset_group_asset_error_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_errors_asset_group_asset_error_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_errors_asset_group_asset_error_proto_goTypes = []any{
	(AssetGroupAssetErrorEnum_AssetGroupAssetError)(0), // 0: google.ads.googleads.v19.errors.AssetGroupAssetErrorEnum.AssetGroupAssetError
	(*AssetGroupAssetErrorEnum)(nil),                   // 1: google.ads.googleads.v19.errors.AssetGroupAssetErrorEnum
}
var file_errors_asset_group_asset_error_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_errors_asset_group_asset_error_proto_init() }
func file_errors_asset_group_asset_error_proto_init() {
	if File_errors_asset_group_asset_error_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: unsafe.Slice(unsafe.StringData(file_errors_asset_group_asset_error_proto_rawDesc), len(file_errors_asset_group_asset_error_proto_rawDesc)),
			NumEnums:      1,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_errors_asset_group_asset_error_proto_goTypes,
		DependencyIndexes: file_errors_asset_group_asset_error_proto_depIdxs,
		EnumInfos:         file_errors_asset_group_asset_error_proto_enumTypes,
		MessageInfos:      file_errors_asset_group_asset_error_proto_msgTypes,
	}.Build()
	File_errors_asset_group_asset_error_proto = out.File
	file_errors_asset_group_asset_error_proto_goTypes = nil
	file_errors_asset_group_asset_error_proto_depIdxs = nil
}
