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
// source: google/ads/googleads/v14/common/asset_usage.proto

package common

import (
	enums "github.com/revealbot/google-ads-go/enums"
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

// Contains the usage information of the asset.
type AssetUsage struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Resource name of the asset.
	Asset string `protobuf:"bytes,1,opt,name=asset,proto3" json:"asset,omitempty"`
	// The served field type of the asset.
	ServedAssetFieldType enums.ServedAssetFieldTypeEnum_ServedAssetFieldType `protobuf:"varint,2,opt,name=served_asset_field_type,json=servedAssetFieldType,proto3,enum=google.ads.googleads.v14.enums.ServedAssetFieldTypeEnum_ServedAssetFieldType" json:"served_asset_field_type,omitempty"`
}

func (x *AssetUsage) Reset() {
	*x = AssetUsage{}
	if protoimpl.UnsafeEnabled {
		mi := &file_common_asset_usage_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AssetUsage) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AssetUsage) ProtoMessage() {}

func (x *AssetUsage) ProtoReflect() protoreflect.Message {
	mi := &file_common_asset_usage_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AssetUsage.ProtoReflect.Descriptor instead.
func (*AssetUsage) Descriptor() ([]byte, []int) {
	return file_common_asset_usage_proto_rawDescGZIP(), []int{0}
}

func (x *AssetUsage) GetAsset() string {
	if x != nil {
		return x.Asset
	}
	return ""
}

func (x *AssetUsage) GetServedAssetFieldType() enums.ServedAssetFieldTypeEnum_ServedAssetFieldType {
	if x != nil {
		return x.ServedAssetFieldType
	}
	return enums.ServedAssetFieldTypeEnum_ServedAssetFieldType(0)
}

var File_common_asset_usage_proto protoreflect.FileDescriptor

var file_common_asset_usage_proto_rawDesc = []byte{
	0x0a, 0x31, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x64, 0x73, 0x2f, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x61, 0x64, 0x73, 0x2f, 0x76, 0x31, 0x34, 0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x6f,
	0x6e, 0x2f, 0x61, 0x73, 0x73, 0x65, 0x74, 0x5f, 0x75, 0x73, 0x61, 0x67, 0x65, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x12, 0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x61, 0x64, 0x73, 0x2e,
	0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61, 0x64, 0x73, 0x2e, 0x76, 0x31, 0x34, 0x2e, 0x63, 0x6f,
	0x6d, 0x6d, 0x6f, 0x6e, 0x1a, 0x3c, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x64, 0x73,
	0x2f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61, 0x64, 0x73, 0x2f, 0x76, 0x31, 0x34, 0x2f, 0x65,
	0x6e, 0x75, 0x6d, 0x73, 0x2f, 0x73, 0x65, 0x72, 0x76, 0x65, 0x64, 0x5f, 0x61, 0x73, 0x73, 0x65,
	0x74, 0x5f, 0x66, 0x69, 0x65, 0x6c, 0x64, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x22, 0xa9, 0x01, 0x0a, 0x0a, 0x41, 0x73, 0x73, 0x65, 0x74, 0x55, 0x73, 0x61, 0x67,
	0x65, 0x12, 0x14, 0x0a, 0x05, 0x61, 0x73, 0x73, 0x65, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x05, 0x61, 0x73, 0x73, 0x65, 0x74, 0x12, 0x84, 0x01, 0x0a, 0x17, 0x73, 0x65, 0x72, 0x76,
	0x65, 0x64, 0x5f, 0x61, 0x73, 0x73, 0x65, 0x74, 0x5f, 0x66, 0x69, 0x65, 0x6c, 0x64, 0x5f, 0x74,
	0x79, 0x70, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x4d, 0x2e, 0x67, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x2e, 0x61, 0x64, 0x73, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61, 0x64, 0x73,
	0x2e, 0x76, 0x31, 0x34, 0x2e, 0x65, 0x6e, 0x75, 0x6d, 0x73, 0x2e, 0x53, 0x65, 0x72, 0x76, 0x65,
	0x64, 0x41, 0x73, 0x73, 0x65, 0x74, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x54, 0x79, 0x70, 0x65, 0x45,
	0x6e, 0x75, 0x6d, 0x2e, 0x53, 0x65, 0x72, 0x76, 0x65, 0x64, 0x41, 0x73, 0x73, 0x65, 0x74, 0x46,
	0x69, 0x65, 0x6c, 0x64, 0x54, 0x79, 0x70, 0x65, 0x52, 0x14, 0x73, 0x65, 0x72, 0x76, 0x65, 0x64,
	0x41, 0x73, 0x73, 0x65, 0x74, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x54, 0x79, 0x70, 0x65, 0x42, 0xef,
	0x01, 0x0a, 0x23, 0x63, 0x6f, 0x6d, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x61, 0x64,
	0x73, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61, 0x64, 0x73, 0x2e, 0x76, 0x31, 0x34, 0x2e,
	0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x42, 0x0f, 0x41, 0x73, 0x73, 0x65, 0x74, 0x55, 0x73, 0x61,
	0x67, 0x65, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a, 0x45, 0x67, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x2e, 0x67, 0x6f, 0x6c, 0x61, 0x6e, 0x67, 0x2e, 0x6f, 0x72, 0x67, 0x2f, 0x67, 0x65, 0x6e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61, 0x70, 0x69, 0x73,
	0x2f, 0x61, 0x64, 0x73, 0x2f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61, 0x64, 0x73, 0x2f, 0x76,
	0x31, 0x34, 0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x3b, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e,
	0xa2, 0x02, 0x03, 0x47, 0x41, 0x41, 0xaa, 0x02, 0x1f, 0x47, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e,
	0x41, 0x64, 0x73, 0x2e, 0x47, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x41, 0x64, 0x73, 0x2e, 0x56, 0x31,
	0x34, 0x2e, 0x43, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0xca, 0x02, 0x1f, 0x47, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x5c, 0x41, 0x64, 0x73, 0x5c, 0x47, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x41, 0x64, 0x73, 0x5c,
	0x56, 0x31, 0x34, 0x5c, 0x43, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0xea, 0x02, 0x23, 0x47, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x3a, 0x3a, 0x41, 0x64, 0x73, 0x3a, 0x3a, 0x47, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x41, 0x64, 0x73, 0x3a, 0x3a, 0x56, 0x31, 0x34, 0x3a, 0x3a, 0x43, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e,
	0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_common_asset_usage_proto_rawDescOnce sync.Once
	file_common_asset_usage_proto_rawDescData = file_common_asset_usage_proto_rawDesc
)

func file_common_asset_usage_proto_rawDescGZIP() []byte {
	file_common_asset_usage_proto_rawDescOnce.Do(func() {
		file_common_asset_usage_proto_rawDescData = protoimpl.X.CompressGZIP(file_common_asset_usage_proto_rawDescData)
	})
	return file_common_asset_usage_proto_rawDescData
}

var file_common_asset_usage_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_common_asset_usage_proto_goTypes = []interface{}{
	(*AssetUsage)(nil), // 0: google.ads.googleads.v14.common.AssetUsage
	(enums.ServedAssetFieldTypeEnum_ServedAssetFieldType)(0), // 1: google.ads.googleads.v14.enums.ServedAssetFieldTypeEnum.ServedAssetFieldType
}
var file_common_asset_usage_proto_depIdxs = []int32{
	1, // 0: google.ads.googleads.v14.common.AssetUsage.served_asset_field_type:type_name -> google.ads.googleads.v14.enums.ServedAssetFieldTypeEnum.ServedAssetFieldType
	1, // [1:1] is the sub-list for method output_type
	1, // [1:1] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_common_asset_usage_proto_init() }
func file_common_asset_usage_proto_init() {
	if File_common_asset_usage_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_common_asset_usage_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AssetUsage); i {
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
			RawDescriptor: file_common_asset_usage_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_common_asset_usage_proto_goTypes,
		DependencyIndexes: file_common_asset_usage_proto_depIdxs,
		MessageInfos:      file_common_asset_usage_proto_msgTypes,
	}.Build()
	File_common_asset_usage_proto = out.File
	file_common_asset_usage_proto_rawDesc = nil
	file_common_asset_usage_proto_goTypes = nil
	file_common_asset_usage_proto_depIdxs = nil
}
