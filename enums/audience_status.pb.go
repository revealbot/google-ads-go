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
// source: google/ads/googleads/v19/enums/audience_status.proto

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

// Enum containing possible audience status types.
type AudienceStatusEnum_AudienceStatus int32

const (
	// Not specified.
	AudienceStatusEnum_UNSPECIFIED AudienceStatusEnum_AudienceStatus = 0
	// Used for return value only. Represents value unknown in this version.
	AudienceStatusEnum_UNKNOWN AudienceStatusEnum_AudienceStatus = 1
	// Enabled status - audience is enabled and can be targeted.
	AudienceStatusEnum_ENABLED AudienceStatusEnum_AudienceStatus = 2
	// Removed status - audience is removed and cannot be used for
	// targeting.
	AudienceStatusEnum_REMOVED AudienceStatusEnum_AudienceStatus = 3
)

// Enum value maps for AudienceStatusEnum_AudienceStatus.
var (
	AudienceStatusEnum_AudienceStatus_name = map[int32]string{
		0: "UNSPECIFIED",
		1: "UNKNOWN",
		2: "ENABLED",
		3: "REMOVED",
	}
	AudienceStatusEnum_AudienceStatus_value = map[string]int32{
		"UNSPECIFIED": 0,
		"UNKNOWN":     1,
		"ENABLED":     2,
		"REMOVED":     3,
	}
)

func (x AudienceStatusEnum_AudienceStatus) Enum() *AudienceStatusEnum_AudienceStatus {
	p := new(AudienceStatusEnum_AudienceStatus)
	*p = x
	return p
}

func (x AudienceStatusEnum_AudienceStatus) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (AudienceStatusEnum_AudienceStatus) Descriptor() protoreflect.EnumDescriptor {
	return file_enums_audience_status_proto_enumTypes[0].Descriptor()
}

func (AudienceStatusEnum_AudienceStatus) Type() protoreflect.EnumType {
	return &file_enums_audience_status_proto_enumTypes[0]
}

func (x AudienceStatusEnum_AudienceStatus) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use AudienceStatusEnum_AudienceStatus.Descriptor instead.
func (AudienceStatusEnum_AudienceStatus) EnumDescriptor() ([]byte, []int) {
	return file_enums_audience_status_proto_rawDescGZIP(), []int{0, 0}
}

// The status of audience.
type AudienceStatusEnum struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *AudienceStatusEnum) Reset() {
	*x = AudienceStatusEnum{}
	mi := &file_enums_audience_status_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *AudienceStatusEnum) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AudienceStatusEnum) ProtoMessage() {}

func (x *AudienceStatusEnum) ProtoReflect() protoreflect.Message {
	mi := &file_enums_audience_status_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AudienceStatusEnum.ProtoReflect.Descriptor instead.
func (*AudienceStatusEnum) Descriptor() ([]byte, []int) {
	return file_enums_audience_status_proto_rawDescGZIP(), []int{0}
}

var File_enums_audience_status_proto protoreflect.FileDescriptor

var file_enums_audience_status_proto_rawDesc = string([]byte{
	0x0a, 0x34, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x64, 0x73, 0x2f, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x61, 0x64, 0x73, 0x2f, 0x76, 0x31, 0x39, 0x2f, 0x65, 0x6e, 0x75, 0x6d, 0x73,
	0x2f, 0x61, 0x75, 0x64, 0x69, 0x65, 0x6e, 0x63, 0x65, 0x5f, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x1e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x61,
	0x64, 0x73, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61, 0x64, 0x73, 0x2e, 0x76, 0x31, 0x39,
	0x2e, 0x65, 0x6e, 0x75, 0x6d, 0x73, 0x22, 0x5e, 0x0a, 0x12, 0x41, 0x75, 0x64, 0x69, 0x65, 0x6e,
	0x63, 0x65, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x45, 0x6e, 0x75, 0x6d, 0x22, 0x48, 0x0a, 0x0e,
	0x41, 0x75, 0x64, 0x69, 0x65, 0x6e, 0x63, 0x65, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x0f,
	0x0a, 0x0b, 0x55, 0x4e, 0x53, 0x50, 0x45, 0x43, 0x49, 0x46, 0x49, 0x45, 0x44, 0x10, 0x00, 0x12,
	0x0b, 0x0a, 0x07, 0x55, 0x4e, 0x4b, 0x4e, 0x4f, 0x57, 0x4e, 0x10, 0x01, 0x12, 0x0b, 0x0a, 0x07,
	0x45, 0x4e, 0x41, 0x42, 0x4c, 0x45, 0x44, 0x10, 0x02, 0x12, 0x0b, 0x0a, 0x07, 0x52, 0x45, 0x4d,
	0x4f, 0x56, 0x45, 0x44, 0x10, 0x03, 0x42, 0xed, 0x01, 0x0a, 0x22, 0x63, 0x6f, 0x6d, 0x2e, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x61, 0x64, 0x73, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x61, 0x64, 0x73, 0x2e, 0x76, 0x31, 0x39, 0x2e, 0x65, 0x6e, 0x75, 0x6d, 0x73, 0x42, 0x13, 0x41,
	0x75, 0x64, 0x69, 0x65, 0x6e, 0x63, 0x65, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x50, 0x72, 0x6f,
	0x74, 0x6f, 0x50, 0x01, 0x5a, 0x43, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x67, 0x6f, 0x6c,
	0x61, 0x6e, 0x67, 0x2e, 0x6f, 0x72, 0x67, 0x2f, 0x67, 0x65, 0x6e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x2f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61, 0x70, 0x69, 0x73, 0x2f, 0x61, 0x64, 0x73, 0x2f,
	0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61, 0x64, 0x73, 0x2f, 0x76, 0x31, 0x39, 0x2f, 0x65, 0x6e,
	0x75, 0x6d, 0x73, 0x3b, 0x65, 0x6e, 0x75, 0x6d, 0x73, 0xa2, 0x02, 0x03, 0x47, 0x41, 0x41, 0xaa,
	0x02, 0x1e, 0x47, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x41, 0x64, 0x73, 0x2e, 0x47, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x41, 0x64, 0x73, 0x2e, 0x56, 0x31, 0x39, 0x2e, 0x45, 0x6e, 0x75, 0x6d, 0x73,
	0xca, 0x02, 0x1e, 0x47, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x5c, 0x41, 0x64, 0x73, 0x5c, 0x47, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x41, 0x64, 0x73, 0x5c, 0x56, 0x31, 0x39, 0x5c, 0x45, 0x6e, 0x75, 0x6d,
	0x73, 0xea, 0x02, 0x22, 0x47, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x3a, 0x3a, 0x41, 0x64, 0x73, 0x3a,
	0x3a, 0x47, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x41, 0x64, 0x73, 0x3a, 0x3a, 0x56, 0x31, 0x39, 0x3a,
	0x3a, 0x45, 0x6e, 0x75, 0x6d, 0x73, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
})

var (
	file_enums_audience_status_proto_rawDescOnce sync.Once
	file_enums_audience_status_proto_rawDescData []byte
)

func file_enums_audience_status_proto_rawDescGZIP() []byte {
	file_enums_audience_status_proto_rawDescOnce.Do(func() {
		file_enums_audience_status_proto_rawDescData = protoimpl.X.CompressGZIP(unsafe.Slice(unsafe.StringData(file_enums_audience_status_proto_rawDesc), len(file_enums_audience_status_proto_rawDesc)))
	})
	return file_enums_audience_status_proto_rawDescData
}

var file_enums_audience_status_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_enums_audience_status_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_enums_audience_status_proto_goTypes = []any{
	(AudienceStatusEnum_AudienceStatus)(0), // 0: google.ads.googleads.v19.enums.AudienceStatusEnum.AudienceStatus
	(*AudienceStatusEnum)(nil),             // 1: google.ads.googleads.v19.enums.AudienceStatusEnum
}
var file_enums_audience_status_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_enums_audience_status_proto_init() }
func file_enums_audience_status_proto_init() {
	if File_enums_audience_status_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: unsafe.Slice(unsafe.StringData(file_enums_audience_status_proto_rawDesc), len(file_enums_audience_status_proto_rawDesc)),
			NumEnums:      1,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_enums_audience_status_proto_goTypes,
		DependencyIndexes: file_enums_audience_status_proto_depIdxs,
		EnumInfos:         file_enums_audience_status_proto_enumTypes,
		MessageInfos:      file_enums_audience_status_proto_msgTypes,
	}.Build()
	File_enums_audience_status_proto = out.File
	file_enums_audience_status_proto_goTypes = nil
	file_enums_audience_status_proto_depIdxs = nil
}
