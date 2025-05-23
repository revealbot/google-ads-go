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
// source: google/ads/googleads/v19/enums/offline_conversion_diagnostic_status_enum.proto

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

// Possible statuses of the offline ingestion setup.
type OfflineConversionDiagnosticStatusEnum_OfflineConversionDiagnosticStatus int32

const (
	// Not specified.
	OfflineConversionDiagnosticStatusEnum_UNSPECIFIED OfflineConversionDiagnosticStatusEnum_OfflineConversionDiagnosticStatus = 0
	// Used for return value only. Represents value unknown in this version.
	OfflineConversionDiagnosticStatusEnum_UNKNOWN OfflineConversionDiagnosticStatusEnum_OfflineConversionDiagnosticStatus = 1
	// Your offline data ingestion setup is active and optimal for downstream
	// processing.
	OfflineConversionDiagnosticStatusEnum_EXCELLENT OfflineConversionDiagnosticStatusEnum_OfflineConversionDiagnosticStatus = 2
	// Your offline ingestion setup is active, but there are further
	// improvements you could make. See alerts.
	OfflineConversionDiagnosticStatusEnum_GOOD OfflineConversionDiagnosticStatusEnum_OfflineConversionDiagnosticStatus = 3
	// Your offline ingestion setup is active, but there are errors that require
	// your attention. See alerts.
	OfflineConversionDiagnosticStatusEnum_NEEDS_ATTENTION OfflineConversionDiagnosticStatusEnum_OfflineConversionDiagnosticStatus = 4
	// Your offline ingestion setup has not received data in the last 28 days,
	// there may be something wrong.
	OfflineConversionDiagnosticStatusEnum_NO_RECENT_UPLOAD OfflineConversionDiagnosticStatusEnum_OfflineConversionDiagnosticStatus = 6
)

// Enum value maps for OfflineConversionDiagnosticStatusEnum_OfflineConversionDiagnosticStatus.
var (
	OfflineConversionDiagnosticStatusEnum_OfflineConversionDiagnosticStatus_name = map[int32]string{
		0: "UNSPECIFIED",
		1: "UNKNOWN",
		2: "EXCELLENT",
		3: "GOOD",
		4: "NEEDS_ATTENTION",
		6: "NO_RECENT_UPLOAD",
	}
	OfflineConversionDiagnosticStatusEnum_OfflineConversionDiagnosticStatus_value = map[string]int32{
		"UNSPECIFIED":      0,
		"UNKNOWN":          1,
		"EXCELLENT":        2,
		"GOOD":             3,
		"NEEDS_ATTENTION":  4,
		"NO_RECENT_UPLOAD": 6,
	}
)

func (x OfflineConversionDiagnosticStatusEnum_OfflineConversionDiagnosticStatus) Enum() *OfflineConversionDiagnosticStatusEnum_OfflineConversionDiagnosticStatus {
	p := new(OfflineConversionDiagnosticStatusEnum_OfflineConversionDiagnosticStatus)
	*p = x
	return p
}

func (x OfflineConversionDiagnosticStatusEnum_OfflineConversionDiagnosticStatus) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (OfflineConversionDiagnosticStatusEnum_OfflineConversionDiagnosticStatus) Descriptor() protoreflect.EnumDescriptor {
	return file_enums_offline_conversion_diagnostic_status_enum_proto_enumTypes[0].Descriptor()
}

func (OfflineConversionDiagnosticStatusEnum_OfflineConversionDiagnosticStatus) Type() protoreflect.EnumType {
	return &file_enums_offline_conversion_diagnostic_status_enum_proto_enumTypes[0]
}

func (x OfflineConversionDiagnosticStatusEnum_OfflineConversionDiagnosticStatus) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use OfflineConversionDiagnosticStatusEnum_OfflineConversionDiagnosticStatus.Descriptor instead.
func (OfflineConversionDiagnosticStatusEnum_OfflineConversionDiagnosticStatus) EnumDescriptor() ([]byte, []int) {
	return file_enums_offline_conversion_diagnostic_status_enum_proto_rawDescGZIP(), []int{0, 0}
}

// All possible statuses for oci diagnostics.
type OfflineConversionDiagnosticStatusEnum struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *OfflineConversionDiagnosticStatusEnum) Reset() {
	*x = OfflineConversionDiagnosticStatusEnum{}
	mi := &file_enums_offline_conversion_diagnostic_status_enum_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *OfflineConversionDiagnosticStatusEnum) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*OfflineConversionDiagnosticStatusEnum) ProtoMessage() {}

func (x *OfflineConversionDiagnosticStatusEnum) ProtoReflect() protoreflect.Message {
	mi := &file_enums_offline_conversion_diagnostic_status_enum_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use OfflineConversionDiagnosticStatusEnum.ProtoReflect.Descriptor instead.
func (*OfflineConversionDiagnosticStatusEnum) Descriptor() ([]byte, []int) {
	return file_enums_offline_conversion_diagnostic_status_enum_proto_rawDescGZIP(), []int{0}
}

var File_enums_offline_conversion_diagnostic_status_enum_proto protoreflect.FileDescriptor

var file_enums_offline_conversion_diagnostic_status_enum_proto_rawDesc = string([]byte{
	0x0a, 0x4e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x64, 0x73, 0x2f, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x61, 0x64, 0x73, 0x2f, 0x76, 0x31, 0x39, 0x2f, 0x65, 0x6e, 0x75, 0x6d, 0x73,
	0x2f, 0x6f, 0x66, 0x66, 0x6c, 0x69, 0x6e, 0x65, 0x5f, 0x63, 0x6f, 0x6e, 0x76, 0x65, 0x72, 0x73,
	0x69, 0x6f, 0x6e, 0x5f, 0x64, 0x69, 0x61, 0x67, 0x6e, 0x6f, 0x73, 0x74, 0x69, 0x63, 0x5f, 0x73,
	0x74, 0x61, 0x74, 0x75, 0x73, 0x5f, 0x65, 0x6e, 0x75, 0x6d, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x12, 0x1e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x61, 0x64, 0x73, 0x2e, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x61, 0x64, 0x73, 0x2e, 0x76, 0x31, 0x39, 0x2e, 0x65, 0x6e, 0x75, 0x6d, 0x73,
	0x22, 0xaf, 0x01, 0x0a, 0x25, 0x4f, 0x66, 0x66, 0x6c, 0x69, 0x6e, 0x65, 0x43, 0x6f, 0x6e, 0x76,
	0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x44, 0x69, 0x61, 0x67, 0x6e, 0x6f, 0x73, 0x74, 0x69, 0x63,
	0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x45, 0x6e, 0x75, 0x6d, 0x22, 0x85, 0x01, 0x0a, 0x21, 0x4f,
	0x66, 0x66, 0x6c, 0x69, 0x6e, 0x65, 0x43, 0x6f, 0x6e, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e,
	0x44, 0x69, 0x61, 0x67, 0x6e, 0x6f, 0x73, 0x74, 0x69, 0x63, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73,
	0x12, 0x0f, 0x0a, 0x0b, 0x55, 0x4e, 0x53, 0x50, 0x45, 0x43, 0x49, 0x46, 0x49, 0x45, 0x44, 0x10,
	0x00, 0x12, 0x0b, 0x0a, 0x07, 0x55, 0x4e, 0x4b, 0x4e, 0x4f, 0x57, 0x4e, 0x10, 0x01, 0x12, 0x0d,
	0x0a, 0x09, 0x45, 0x58, 0x43, 0x45, 0x4c, 0x4c, 0x45, 0x4e, 0x54, 0x10, 0x02, 0x12, 0x08, 0x0a,
	0x04, 0x47, 0x4f, 0x4f, 0x44, 0x10, 0x03, 0x12, 0x13, 0x0a, 0x0f, 0x4e, 0x45, 0x45, 0x44, 0x53,
	0x5f, 0x41, 0x54, 0x54, 0x45, 0x4e, 0x54, 0x49, 0x4f, 0x4e, 0x10, 0x04, 0x12, 0x14, 0x0a, 0x10,
	0x4e, 0x4f, 0x5f, 0x52, 0x45, 0x43, 0x45, 0x4e, 0x54, 0x5f, 0x55, 0x50, 0x4c, 0x4f, 0x41, 0x44,
	0x10, 0x06, 0x42, 0x84, 0x02, 0x0a, 0x22, 0x63, 0x6f, 0x6d, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x2e, 0x61, 0x64, 0x73, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61, 0x64, 0x73, 0x2e,
	0x76, 0x31, 0x39, 0x2e, 0x65, 0x6e, 0x75, 0x6d, 0x73, 0x42, 0x2a, 0x4f, 0x66, 0x66, 0x6c, 0x69,
	0x6e, 0x65, 0x43, 0x6f, 0x6e, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x44, 0x69, 0x61, 0x67,
	0x6e, 0x6f, 0x73, 0x74, 0x69, 0x63, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x45, 0x6e, 0x75, 0x6d,
	0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a, 0x43, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e,
	0x67, 0x6f, 0x6c, 0x61, 0x6e, 0x67, 0x2e, 0x6f, 0x72, 0x67, 0x2f, 0x67, 0x65, 0x6e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x2f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61, 0x70, 0x69, 0x73, 0x2f, 0x61,
	0x64, 0x73, 0x2f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61, 0x64, 0x73, 0x2f, 0x76, 0x31, 0x39,
	0x2f, 0x65, 0x6e, 0x75, 0x6d, 0x73, 0x3b, 0x65, 0x6e, 0x75, 0x6d, 0x73, 0xa2, 0x02, 0x03, 0x47,
	0x41, 0x41, 0xaa, 0x02, 0x1e, 0x47, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x41, 0x64, 0x73, 0x2e,
	0x47, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x41, 0x64, 0x73, 0x2e, 0x56, 0x31, 0x39, 0x2e, 0x45, 0x6e,
	0x75, 0x6d, 0x73, 0xca, 0x02, 0x1e, 0x47, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x5c, 0x41, 0x64, 0x73,
	0x5c, 0x47, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x41, 0x64, 0x73, 0x5c, 0x56, 0x31, 0x39, 0x5c, 0x45,
	0x6e, 0x75, 0x6d, 0x73, 0xea, 0x02, 0x22, 0x47, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x3a, 0x3a, 0x41,
	0x64, 0x73, 0x3a, 0x3a, 0x47, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x41, 0x64, 0x73, 0x3a, 0x3a, 0x56,
	0x31, 0x39, 0x3a, 0x3a, 0x45, 0x6e, 0x75, 0x6d, 0x73, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x33,
})

var (
	file_enums_offline_conversion_diagnostic_status_enum_proto_rawDescOnce sync.Once
	file_enums_offline_conversion_diagnostic_status_enum_proto_rawDescData []byte
)

func file_enums_offline_conversion_diagnostic_status_enum_proto_rawDescGZIP() []byte {
	file_enums_offline_conversion_diagnostic_status_enum_proto_rawDescOnce.Do(func() {
		file_enums_offline_conversion_diagnostic_status_enum_proto_rawDescData = protoimpl.X.CompressGZIP(unsafe.Slice(unsafe.StringData(file_enums_offline_conversion_diagnostic_status_enum_proto_rawDesc), len(file_enums_offline_conversion_diagnostic_status_enum_proto_rawDesc)))
	})
	return file_enums_offline_conversion_diagnostic_status_enum_proto_rawDescData
}

var file_enums_offline_conversion_diagnostic_status_enum_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_enums_offline_conversion_diagnostic_status_enum_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_enums_offline_conversion_diagnostic_status_enum_proto_goTypes = []any{
	(OfflineConversionDiagnosticStatusEnum_OfflineConversionDiagnosticStatus)(0), // 0: google.ads.googleads.v19.enums.OfflineConversionDiagnosticStatusEnum.OfflineConversionDiagnosticStatus
	(*OfflineConversionDiagnosticStatusEnum)(nil),                                // 1: google.ads.googleads.v19.enums.OfflineConversionDiagnosticStatusEnum
}
var file_enums_offline_conversion_diagnostic_status_enum_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() {
	file_enums_offline_conversion_diagnostic_status_enum_proto_init()
}
func file_enums_offline_conversion_diagnostic_status_enum_proto_init() {
	if File_enums_offline_conversion_diagnostic_status_enum_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: unsafe.Slice(unsafe.StringData(file_enums_offline_conversion_diagnostic_status_enum_proto_rawDesc), len(file_enums_offline_conversion_diagnostic_status_enum_proto_rawDesc)),
			NumEnums:      1,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_enums_offline_conversion_diagnostic_status_enum_proto_goTypes,
		DependencyIndexes: file_enums_offline_conversion_diagnostic_status_enum_proto_depIdxs,
		EnumInfos:         file_enums_offline_conversion_diagnostic_status_enum_proto_enumTypes,
		MessageInfos:      file_enums_offline_conversion_diagnostic_status_enum_proto_msgTypes,
	}.Build()
	File_enums_offline_conversion_diagnostic_status_enum_proto = out.File
	file_enums_offline_conversion_diagnostic_status_enum_proto_goTypes = nil
	file_enums_offline_conversion_diagnostic_status_enum_proto_depIdxs = nil
}
