// Copyright 2024 Google LLC
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
// source: google/ads/googleads/v17/enums/asset_link_primary_status.proto

package enums

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

// Enum Provides insight into why an asset is not serving or not serving
// at full capacity for a particular link level. There could be one status
// with multiple reasons. For example, a sitelink might be paused by the user,
// but also limited in serving due to violation of an alcohol policy. In
// this case, the PrimaryStatus will be returned as PAUSED, since the asset's
// effective status is determined by its paused state.
type AssetLinkPrimaryStatusEnum_AssetLinkPrimaryStatus int32

const (
	// Not specified.
	AssetLinkPrimaryStatusEnum_UNSPECIFIED AssetLinkPrimaryStatusEnum_AssetLinkPrimaryStatus = 0
	// Used for return value only. Represents value unknown in this version.
	AssetLinkPrimaryStatusEnum_UNKNOWN AssetLinkPrimaryStatusEnum_AssetLinkPrimaryStatus = 1
	// The asset is eligible to serve.
	AssetLinkPrimaryStatusEnum_ELIGIBLE AssetLinkPrimaryStatusEnum_AssetLinkPrimaryStatus = 2
	// The user-specified asset link status is paused.
	AssetLinkPrimaryStatusEnum_PAUSED AssetLinkPrimaryStatusEnum_AssetLinkPrimaryStatus = 3
	// The user-specified asset link status is removed.
	AssetLinkPrimaryStatusEnum_REMOVED AssetLinkPrimaryStatusEnum_AssetLinkPrimaryStatus = 4
	// The asset may serve in the future.
	AssetLinkPrimaryStatusEnum_PENDING AssetLinkPrimaryStatusEnum_AssetLinkPrimaryStatus = 5
	// The asset is serving in a partial capacity.
	AssetLinkPrimaryStatusEnum_LIMITED AssetLinkPrimaryStatusEnum_AssetLinkPrimaryStatus = 6
	// The asset is not eligible to serve.
	AssetLinkPrimaryStatusEnum_NOT_ELIGIBLE AssetLinkPrimaryStatusEnum_AssetLinkPrimaryStatus = 7
)

// Enum value maps for AssetLinkPrimaryStatusEnum_AssetLinkPrimaryStatus.
var (
	AssetLinkPrimaryStatusEnum_AssetLinkPrimaryStatus_name = map[int32]string{
		0: "UNSPECIFIED",
		1: "UNKNOWN",
		2: "ELIGIBLE",
		3: "PAUSED",
		4: "REMOVED",
		5: "PENDING",
		6: "LIMITED",
		7: "NOT_ELIGIBLE",
	}
	AssetLinkPrimaryStatusEnum_AssetLinkPrimaryStatus_value = map[string]int32{
		"UNSPECIFIED":  0,
		"UNKNOWN":      1,
		"ELIGIBLE":     2,
		"PAUSED":       3,
		"REMOVED":      4,
		"PENDING":      5,
		"LIMITED":      6,
		"NOT_ELIGIBLE": 7,
	}
)

func (x AssetLinkPrimaryStatusEnum_AssetLinkPrimaryStatus) Enum() *AssetLinkPrimaryStatusEnum_AssetLinkPrimaryStatus {
	p := new(AssetLinkPrimaryStatusEnum_AssetLinkPrimaryStatus)
	*p = x
	return p
}

func (x AssetLinkPrimaryStatusEnum_AssetLinkPrimaryStatus) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (AssetLinkPrimaryStatusEnum_AssetLinkPrimaryStatus) Descriptor() protoreflect.EnumDescriptor {
	return file_enums_asset_link_primary_status_proto_enumTypes[0].Descriptor()
}

func (AssetLinkPrimaryStatusEnum_AssetLinkPrimaryStatus) Type() protoreflect.EnumType {
	return &file_enums_asset_link_primary_status_proto_enumTypes[0]
}

func (x AssetLinkPrimaryStatusEnum_AssetLinkPrimaryStatus) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use AssetLinkPrimaryStatusEnum_AssetLinkPrimaryStatus.Descriptor instead.
func (AssetLinkPrimaryStatusEnum_AssetLinkPrimaryStatus) EnumDescriptor() ([]byte, []int) {
	return file_enums_asset_link_primary_status_proto_rawDescGZIP(), []int{0, 0}
}

// Provides the primary status of an asset link.
// For example: a sitelink may be paused for a particular campaign.
type AssetLinkPrimaryStatusEnum struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *AssetLinkPrimaryStatusEnum) Reset() {
	*x = AssetLinkPrimaryStatusEnum{}
	if protoimpl.UnsafeEnabled {
		mi := &file_enums_asset_link_primary_status_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AssetLinkPrimaryStatusEnum) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AssetLinkPrimaryStatusEnum) ProtoMessage() {}

func (x *AssetLinkPrimaryStatusEnum) ProtoReflect() protoreflect.Message {
	mi := &file_enums_asset_link_primary_status_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AssetLinkPrimaryStatusEnum.ProtoReflect.Descriptor instead.
func (*AssetLinkPrimaryStatusEnum) Descriptor() ([]byte, []int) {
	return file_enums_asset_link_primary_status_proto_rawDescGZIP(), []int{0}
}

var File_enums_asset_link_primary_status_proto protoreflect.FileDescriptor

var file_enums_asset_link_primary_status_proto_rawDesc = []byte{
	0x0a, 0x3e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x64, 0x73, 0x2f, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x61, 0x64, 0x73, 0x2f, 0x76, 0x31, 0x37, 0x2f, 0x65, 0x6e, 0x75, 0x6d, 0x73,
	0x2f, 0x61, 0x73, 0x73, 0x65, 0x74, 0x5f, 0x6c, 0x69, 0x6e, 0x6b, 0x5f, 0x70, 0x72, 0x69, 0x6d,
	0x61, 0x72, 0x79, 0x5f, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x12, 0x1e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x61, 0x64, 0x73, 0x2e, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x61, 0x64, 0x73, 0x2e, 0x76, 0x31, 0x37, 0x2e, 0x65, 0x6e, 0x75, 0x6d, 0x73,
	0x22, 0xa8, 0x01, 0x0a, 0x1a, 0x41, 0x73, 0x73, 0x65, 0x74, 0x4c, 0x69, 0x6e, 0x6b, 0x50, 0x72,
	0x69, 0x6d, 0x61, 0x72, 0x79, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x45, 0x6e, 0x75, 0x6d, 0x22,
	0x89, 0x01, 0x0a, 0x16, 0x41, 0x73, 0x73, 0x65, 0x74, 0x4c, 0x69, 0x6e, 0x6b, 0x50, 0x72, 0x69,
	0x6d, 0x61, 0x72, 0x79, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x0f, 0x0a, 0x0b, 0x55, 0x4e,
	0x53, 0x50, 0x45, 0x43, 0x49, 0x46, 0x49, 0x45, 0x44, 0x10, 0x00, 0x12, 0x0b, 0x0a, 0x07, 0x55,
	0x4e, 0x4b, 0x4e, 0x4f, 0x57, 0x4e, 0x10, 0x01, 0x12, 0x0c, 0x0a, 0x08, 0x45, 0x4c, 0x49, 0x47,
	0x49, 0x42, 0x4c, 0x45, 0x10, 0x02, 0x12, 0x0a, 0x0a, 0x06, 0x50, 0x41, 0x55, 0x53, 0x45, 0x44,
	0x10, 0x03, 0x12, 0x0b, 0x0a, 0x07, 0x52, 0x45, 0x4d, 0x4f, 0x56, 0x45, 0x44, 0x10, 0x04, 0x12,
	0x0b, 0x0a, 0x07, 0x50, 0x45, 0x4e, 0x44, 0x49, 0x4e, 0x47, 0x10, 0x05, 0x12, 0x0b, 0x0a, 0x07,
	0x4c, 0x49, 0x4d, 0x49, 0x54, 0x45, 0x44, 0x10, 0x06, 0x12, 0x10, 0x0a, 0x0c, 0x4e, 0x4f, 0x54,
	0x5f, 0x45, 0x4c, 0x49, 0x47, 0x49, 0x42, 0x4c, 0x45, 0x10, 0x07, 0x42, 0xf5, 0x01, 0x0a, 0x22,
	0x63, 0x6f, 0x6d, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x61, 0x64, 0x73, 0x2e, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61, 0x64, 0x73, 0x2e, 0x76, 0x31, 0x37, 0x2e, 0x65, 0x6e, 0x75,
	0x6d, 0x73, 0x42, 0x1b, 0x41, 0x73, 0x73, 0x65, 0x74, 0x4c, 0x69, 0x6e, 0x6b, 0x50, 0x72, 0x69,
	0x6d, 0x61, 0x72, 0x79, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50,
	0x01, 0x5a, 0x43, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x67, 0x6f, 0x6c, 0x61, 0x6e, 0x67,
	0x2e, 0x6f, 0x72, 0x67, 0x2f, 0x67, 0x65, 0x6e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x67, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x61, 0x70, 0x69, 0x73, 0x2f, 0x61, 0x64, 0x73, 0x2f, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x61, 0x64, 0x73, 0x2f, 0x76, 0x31, 0x37, 0x2f, 0x65, 0x6e, 0x75, 0x6d, 0x73,
	0x3b, 0x65, 0x6e, 0x75, 0x6d, 0x73, 0xa2, 0x02, 0x03, 0x47, 0x41, 0x41, 0xaa, 0x02, 0x1e, 0x47,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x41, 0x64, 0x73, 0x2e, 0x47, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x41, 0x64, 0x73, 0x2e, 0x56, 0x31, 0x37, 0x2e, 0x45, 0x6e, 0x75, 0x6d, 0x73, 0xca, 0x02, 0x1e,
	0x47, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x5c, 0x41, 0x64, 0x73, 0x5c, 0x47, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x41, 0x64, 0x73, 0x5c, 0x56, 0x31, 0x37, 0x5c, 0x45, 0x6e, 0x75, 0x6d, 0x73, 0xea, 0x02,
	0x22, 0x47, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x3a, 0x3a, 0x41, 0x64, 0x73, 0x3a, 0x3a, 0x47, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x41, 0x64, 0x73, 0x3a, 0x3a, 0x56, 0x31, 0x37, 0x3a, 0x3a, 0x45, 0x6e,
	0x75, 0x6d, 0x73, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_enums_asset_link_primary_status_proto_rawDescOnce sync.Once
	file_enums_asset_link_primary_status_proto_rawDescData = file_enums_asset_link_primary_status_proto_rawDesc
)

func file_enums_asset_link_primary_status_proto_rawDescGZIP() []byte {
	file_enums_asset_link_primary_status_proto_rawDescOnce.Do(func() {
		file_enums_asset_link_primary_status_proto_rawDescData = protoimpl.X.CompressGZIP(file_enums_asset_link_primary_status_proto_rawDescData)
	})
	return file_enums_asset_link_primary_status_proto_rawDescData
}

var file_enums_asset_link_primary_status_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_enums_asset_link_primary_status_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_enums_asset_link_primary_status_proto_goTypes = []interface{}{
	(AssetLinkPrimaryStatusEnum_AssetLinkPrimaryStatus)(0), // 0: google.ads.googleads.v17.enums.AssetLinkPrimaryStatusEnum.AssetLinkPrimaryStatus
	(*AssetLinkPrimaryStatusEnum)(nil),                     // 1: google.ads.googleads.v17.enums.AssetLinkPrimaryStatusEnum
}
var file_enums_asset_link_primary_status_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_enums_asset_link_primary_status_proto_init() }
func file_enums_asset_link_primary_status_proto_init() {
	if File_enums_asset_link_primary_status_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_enums_asset_link_primary_status_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AssetLinkPrimaryStatusEnum); i {
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
			RawDescriptor: file_enums_asset_link_primary_status_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_enums_asset_link_primary_status_proto_goTypes,
		DependencyIndexes: file_enums_asset_link_primary_status_proto_depIdxs,
		EnumInfos:         file_enums_asset_link_primary_status_proto_enumTypes,
		MessageInfos:      file_enums_asset_link_primary_status_proto_msgTypes,
	}.Build()
	File_enums_asset_link_primary_status_proto = out.File
	file_enums_asset_link_primary_status_proto_rawDesc = nil
	file_enums_asset_link_primary_status_proto_goTypes = nil
	file_enums_asset_link_primary_status_proto_depIdxs = nil
}
