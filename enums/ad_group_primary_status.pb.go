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
// source: google/ads/googleads/v19/enums/ad_group_primary_status.proto

package enums

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

// The possible statuses of an AdGroup.
type AdGroupPrimaryStatusEnum_AdGroupPrimaryStatus int32

const (
	// Not specified.
	AdGroupPrimaryStatusEnum_UNSPECIFIED AdGroupPrimaryStatusEnum_AdGroupPrimaryStatus = 0
	// Used for return value only. Represents value unknown in this version.
	AdGroupPrimaryStatusEnum_UNKNOWN AdGroupPrimaryStatusEnum_AdGroupPrimaryStatus = 1
	// The ad group is eligible to serve.
	AdGroupPrimaryStatusEnum_ELIGIBLE AdGroupPrimaryStatusEnum_AdGroupPrimaryStatus = 2
	// The ad group is paused.
	AdGroupPrimaryStatusEnum_PAUSED AdGroupPrimaryStatusEnum_AdGroupPrimaryStatus = 3
	// The ad group is removed.
	AdGroupPrimaryStatusEnum_REMOVED AdGroupPrimaryStatusEnum_AdGroupPrimaryStatus = 4
	// The ad group may serve in the future.
	AdGroupPrimaryStatusEnum_PENDING AdGroupPrimaryStatusEnum_AdGroupPrimaryStatus = 5
	// The ad group is not eligible to serve.
	AdGroupPrimaryStatusEnum_NOT_ELIGIBLE AdGroupPrimaryStatusEnum_AdGroupPrimaryStatus = 6
	// The ad group has limited servability.
	AdGroupPrimaryStatusEnum_LIMITED AdGroupPrimaryStatusEnum_AdGroupPrimaryStatus = 7
)

// Enum value maps for AdGroupPrimaryStatusEnum_AdGroupPrimaryStatus.
var (
	AdGroupPrimaryStatusEnum_AdGroupPrimaryStatus_name = map[int32]string{
		0: "UNSPECIFIED",
		1: "UNKNOWN",
		2: "ELIGIBLE",
		3: "PAUSED",
		4: "REMOVED",
		5: "PENDING",
		6: "NOT_ELIGIBLE",
		7: "LIMITED",
	}
	AdGroupPrimaryStatusEnum_AdGroupPrimaryStatus_value = map[string]int32{
		"UNSPECIFIED":  0,
		"UNKNOWN":      1,
		"ELIGIBLE":     2,
		"PAUSED":       3,
		"REMOVED":      4,
		"PENDING":      5,
		"NOT_ELIGIBLE": 6,
		"LIMITED":      7,
	}
)

func (x AdGroupPrimaryStatusEnum_AdGroupPrimaryStatus) Enum() *AdGroupPrimaryStatusEnum_AdGroupPrimaryStatus {
	p := new(AdGroupPrimaryStatusEnum_AdGroupPrimaryStatus)
	*p = x
	return p
}

func (x AdGroupPrimaryStatusEnum_AdGroupPrimaryStatus) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (AdGroupPrimaryStatusEnum_AdGroupPrimaryStatus) Descriptor() protoreflect.EnumDescriptor {
	return file_enums_ad_group_primary_status_proto_enumTypes[0].Descriptor()
}

func (AdGroupPrimaryStatusEnum_AdGroupPrimaryStatus) Type() protoreflect.EnumType {
	return &file_enums_ad_group_primary_status_proto_enumTypes[0]
}

func (x AdGroupPrimaryStatusEnum_AdGroupPrimaryStatus) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use AdGroupPrimaryStatusEnum_AdGroupPrimaryStatus.Descriptor instead.
func (AdGroupPrimaryStatusEnum_AdGroupPrimaryStatus) EnumDescriptor() ([]byte, []int) {
	return file_enums_ad_group_primary_status_proto_rawDescGZIP(), []int{0, 0}
}

// Ad Group Primary Status. Provides insight into why an ad group is not serving
// or not serving optimally.
type AdGroupPrimaryStatusEnum struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *AdGroupPrimaryStatusEnum) Reset() {
	*x = AdGroupPrimaryStatusEnum{}
	mi := &file_enums_ad_group_primary_status_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *AdGroupPrimaryStatusEnum) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AdGroupPrimaryStatusEnum) ProtoMessage() {}

func (x *AdGroupPrimaryStatusEnum) ProtoReflect() protoreflect.Message {
	mi := &file_enums_ad_group_primary_status_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AdGroupPrimaryStatusEnum.ProtoReflect.Descriptor instead.
func (*AdGroupPrimaryStatusEnum) Descriptor() ([]byte, []int) {
	return file_enums_ad_group_primary_status_proto_rawDescGZIP(), []int{0}
}

var File_enums_ad_group_primary_status_proto protoreflect.FileDescriptor

var file_enums_ad_group_primary_status_proto_rawDesc = string([]byte{
	0x0a, 0x3c, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x64, 0x73, 0x2f, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x61, 0x64, 0x73, 0x2f, 0x76, 0x31, 0x39, 0x2f, 0x65, 0x6e, 0x75, 0x6d, 0x73,
	0x2f, 0x61, 0x64, 0x5f, 0x67, 0x72, 0x6f, 0x75, 0x70, 0x5f, 0x70, 0x72, 0x69, 0x6d, 0x61, 0x72,
	0x79, 0x5f, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x1e,
	0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x61, 0x64, 0x73, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x61, 0x64, 0x73, 0x2e, 0x76, 0x31, 0x39, 0x2e, 0x65, 0x6e, 0x75, 0x6d, 0x73, 0x22, 0xa4,
	0x01, 0x0a, 0x18, 0x41, 0x64, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x50, 0x72, 0x69, 0x6d, 0x61, 0x72,
	0x79, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x45, 0x6e, 0x75, 0x6d, 0x22, 0x87, 0x01, 0x0a, 0x14,
	0x41, 0x64, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x50, 0x72, 0x69, 0x6d, 0x61, 0x72, 0x79, 0x53, 0x74,
	0x61, 0x74, 0x75, 0x73, 0x12, 0x0f, 0x0a, 0x0b, 0x55, 0x4e, 0x53, 0x50, 0x45, 0x43, 0x49, 0x46,
	0x49, 0x45, 0x44, 0x10, 0x00, 0x12, 0x0b, 0x0a, 0x07, 0x55, 0x4e, 0x4b, 0x4e, 0x4f, 0x57, 0x4e,
	0x10, 0x01, 0x12, 0x0c, 0x0a, 0x08, 0x45, 0x4c, 0x49, 0x47, 0x49, 0x42, 0x4c, 0x45, 0x10, 0x02,
	0x12, 0x0a, 0x0a, 0x06, 0x50, 0x41, 0x55, 0x53, 0x45, 0x44, 0x10, 0x03, 0x12, 0x0b, 0x0a, 0x07,
	0x52, 0x45, 0x4d, 0x4f, 0x56, 0x45, 0x44, 0x10, 0x04, 0x12, 0x0b, 0x0a, 0x07, 0x50, 0x45, 0x4e,
	0x44, 0x49, 0x4e, 0x47, 0x10, 0x05, 0x12, 0x10, 0x0a, 0x0c, 0x4e, 0x4f, 0x54, 0x5f, 0x45, 0x4c,
	0x49, 0x47, 0x49, 0x42, 0x4c, 0x45, 0x10, 0x06, 0x12, 0x0b, 0x0a, 0x07, 0x4c, 0x49, 0x4d, 0x49,
	0x54, 0x45, 0x44, 0x10, 0x07, 0x42, 0xf3, 0x01, 0x0a, 0x22, 0x63, 0x6f, 0x6d, 0x2e, 0x67, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x61, 0x64, 0x73, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61,
	0x64, 0x73, 0x2e, 0x76, 0x31, 0x39, 0x2e, 0x65, 0x6e, 0x75, 0x6d, 0x73, 0x42, 0x19, 0x41, 0x64,
	0x47, 0x72, 0x6f, 0x75, 0x70, 0x50, 0x72, 0x69, 0x6d, 0x61, 0x72, 0x79, 0x53, 0x74, 0x61, 0x74,
	0x75, 0x73, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a, 0x43, 0x67, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x2e, 0x67, 0x6f, 0x6c, 0x61, 0x6e, 0x67, 0x2e, 0x6f, 0x72, 0x67, 0x2f, 0x67, 0x65, 0x6e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61, 0x70, 0x69, 0x73,
	0x2f, 0x61, 0x64, 0x73, 0x2f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61, 0x64, 0x73, 0x2f, 0x76,
	0x31, 0x39, 0x2f, 0x65, 0x6e, 0x75, 0x6d, 0x73, 0x3b, 0x65, 0x6e, 0x75, 0x6d, 0x73, 0xa2, 0x02,
	0x03, 0x47, 0x41, 0x41, 0xaa, 0x02, 0x1e, 0x47, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x41, 0x64,
	0x73, 0x2e, 0x47, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x41, 0x64, 0x73, 0x2e, 0x56, 0x31, 0x39, 0x2e,
	0x45, 0x6e, 0x75, 0x6d, 0x73, 0xca, 0x02, 0x1e, 0x47, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x5c, 0x41,
	0x64, 0x73, 0x5c, 0x47, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x41, 0x64, 0x73, 0x5c, 0x56, 0x31, 0x39,
	0x5c, 0x45, 0x6e, 0x75, 0x6d, 0x73, 0xea, 0x02, 0x22, 0x47, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x3a,
	0x3a, 0x41, 0x64, 0x73, 0x3a, 0x3a, 0x47, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x41, 0x64, 0x73, 0x3a,
	0x3a, 0x56, 0x31, 0x39, 0x3a, 0x3a, 0x45, 0x6e, 0x75, 0x6d, 0x73, 0x62, 0x06, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x33,
})

var (
	file_enums_ad_group_primary_status_proto_rawDescOnce sync.Once
	file_enums_ad_group_primary_status_proto_rawDescData []byte
)

func file_enums_ad_group_primary_status_proto_rawDescGZIP() []byte {
	file_enums_ad_group_primary_status_proto_rawDescOnce.Do(func() {
		file_enums_ad_group_primary_status_proto_rawDescData = protoimpl.X.CompressGZIP(unsafe.Slice(unsafe.StringData(file_enums_ad_group_primary_status_proto_rawDesc), len(file_enums_ad_group_primary_status_proto_rawDesc)))
	})
	return file_enums_ad_group_primary_status_proto_rawDescData
}

var file_enums_ad_group_primary_status_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_enums_ad_group_primary_status_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_enums_ad_group_primary_status_proto_goTypes = []any{
	(AdGroupPrimaryStatusEnum_AdGroupPrimaryStatus)(0), // 0: google.ads.googleads.v19.enums.AdGroupPrimaryStatusEnum.AdGroupPrimaryStatus
	(*AdGroupPrimaryStatusEnum)(nil),                   // 1: google.ads.googleads.v19.enums.AdGroupPrimaryStatusEnum
}
var file_enums_ad_group_primary_status_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_enums_ad_group_primary_status_proto_init() }
func file_enums_ad_group_primary_status_proto_init() {
	if File_enums_ad_group_primary_status_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: unsafe.Slice(unsafe.StringData(file_enums_ad_group_primary_status_proto_rawDesc), len(file_enums_ad_group_primary_status_proto_rawDesc)),
			NumEnums:      1,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_enums_ad_group_primary_status_proto_goTypes,
		DependencyIndexes: file_enums_ad_group_primary_status_proto_depIdxs,
		EnumInfos:         file_enums_ad_group_primary_status_proto_enumTypes,
		MessageInfos:      file_enums_ad_group_primary_status_proto_msgTypes,
	}.Build()
	File_enums_ad_group_primary_status_proto = out.File
	file_enums_ad_group_primary_status_proto_goTypes = nil
	file_enums_ad_group_primary_status_proto_depIdxs = nil
}
