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
// 	protoc-gen-go v1.26.0
// 	protoc        v3.21.12
// source: google/ads/googleads/v14/enums/offline_event_upload_client_enum.proto

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

// Next id: 5
type OfflineEventUploadClientEnum_OfflineEventUploadClient int32

const (
	// Not specified.
	OfflineEventUploadClientEnum_UNSPECIFIED OfflineEventUploadClientEnum_OfflineEventUploadClient = 0
	// Used for return value only. Represents value unknown in this version.
	OfflineEventUploadClientEnum_UNKNOWN OfflineEventUploadClientEnum_OfflineEventUploadClient = 1
	// Google Ads API.
	OfflineEventUploadClientEnum_GOOGLE_ADS_API OfflineEventUploadClientEnum_OfflineEventUploadClient = 2
	// Google Ads web client, which could include multiple sources like Ads UI,
	// SFTP, etc.
	OfflineEventUploadClientEnum_GOOGLE_ADS_WEB_CLIENT OfflineEventUploadClientEnum_OfflineEventUploadClient = 3
	// Connection platform.
	OfflineEventUploadClientEnum_ADS_DATA_CONNECTOR OfflineEventUploadClientEnum_OfflineEventUploadClient = 4
)

// Enum value maps for OfflineEventUploadClientEnum_OfflineEventUploadClient.
var (
	OfflineEventUploadClientEnum_OfflineEventUploadClient_name = map[int32]string{
		0: "UNSPECIFIED",
		1: "UNKNOWN",
		2: "GOOGLE_ADS_API",
		3: "GOOGLE_ADS_WEB_CLIENT",
		4: "ADS_DATA_CONNECTOR",
	}
	OfflineEventUploadClientEnum_OfflineEventUploadClient_value = map[string]int32{
		"UNSPECIFIED":           0,
		"UNKNOWN":               1,
		"GOOGLE_ADS_API":        2,
		"GOOGLE_ADS_WEB_CLIENT": 3,
		"ADS_DATA_CONNECTOR":    4,
	}
)

func (x OfflineEventUploadClientEnum_OfflineEventUploadClient) Enum() *OfflineEventUploadClientEnum_OfflineEventUploadClient {
	p := new(OfflineEventUploadClientEnum_OfflineEventUploadClient)
	*p = x
	return p
}

func (x OfflineEventUploadClientEnum_OfflineEventUploadClient) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (OfflineEventUploadClientEnum_OfflineEventUploadClient) Descriptor() protoreflect.EnumDescriptor {
	return file_enums_offline_event_upload_client_enum_proto_enumTypes[0].Descriptor()
}

func (OfflineEventUploadClientEnum_OfflineEventUploadClient) Type() protoreflect.EnumType {
	return &file_enums_offline_event_upload_client_enum_proto_enumTypes[0]
}

func (x OfflineEventUploadClientEnum_OfflineEventUploadClient) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use OfflineEventUploadClientEnum_OfflineEventUploadClient.Descriptor instead.
func (OfflineEventUploadClientEnum_OfflineEventUploadClient) EnumDescriptor() ([]byte, []int) {
	return file_enums_offline_event_upload_client_enum_proto_rawDescGZIP(), []int{0, 0}
}

// All possible clients for an offline upload event.
type OfflineEventUploadClientEnum struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *OfflineEventUploadClientEnum) Reset() {
	*x = OfflineEventUploadClientEnum{}
	if protoimpl.UnsafeEnabled {
		mi := &file_enums_offline_event_upload_client_enum_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *OfflineEventUploadClientEnum) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*OfflineEventUploadClientEnum) ProtoMessage() {}

func (x *OfflineEventUploadClientEnum) ProtoReflect() protoreflect.Message {
	mi := &file_enums_offline_event_upload_client_enum_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use OfflineEventUploadClientEnum.ProtoReflect.Descriptor instead.
func (*OfflineEventUploadClientEnum) Descriptor() ([]byte, []int) {
	return file_enums_offline_event_upload_client_enum_proto_rawDescGZIP(), []int{0}
}

var File_enums_offline_event_upload_client_enum_proto protoreflect.FileDescriptor

var file_enums_offline_event_upload_client_enum_proto_rawDesc = []byte{
	0x0a, 0x45, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x64, 0x73, 0x2f, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x61, 0x64, 0x73, 0x2f, 0x76, 0x31, 0x34, 0x2f, 0x65, 0x6e, 0x75, 0x6d, 0x73,
	0x2f, 0x6f, 0x66, 0x66, 0x6c, 0x69, 0x6e, 0x65, 0x5f, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x5f, 0x75,
	0x70, 0x6c, 0x6f, 0x61, 0x64, 0x5f, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x5f, 0x65, 0x6e, 0x75,
	0x6d, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x1e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e,
	0x61, 0x64, 0x73, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61, 0x64, 0x73, 0x2e, 0x76, 0x31,
	0x34, 0x2e, 0x65, 0x6e, 0x75, 0x6d, 0x73, 0x22, 0x9f, 0x01, 0x0a, 0x1c, 0x4f, 0x66, 0x66, 0x6c,
	0x69, 0x6e, 0x65, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x55, 0x70, 0x6c, 0x6f, 0x61, 0x64, 0x43, 0x6c,
	0x69, 0x65, 0x6e, 0x74, 0x45, 0x6e, 0x75, 0x6d, 0x22, 0x7f, 0x0a, 0x18, 0x4f, 0x66, 0x66, 0x6c,
	0x69, 0x6e, 0x65, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x55, 0x70, 0x6c, 0x6f, 0x61, 0x64, 0x43, 0x6c,
	0x69, 0x65, 0x6e, 0x74, 0x12, 0x0f, 0x0a, 0x0b, 0x55, 0x4e, 0x53, 0x50, 0x45, 0x43, 0x49, 0x46,
	0x49, 0x45, 0x44, 0x10, 0x00, 0x12, 0x0b, 0x0a, 0x07, 0x55, 0x4e, 0x4b, 0x4e, 0x4f, 0x57, 0x4e,
	0x10, 0x01, 0x12, 0x12, 0x0a, 0x0e, 0x47, 0x4f, 0x4f, 0x47, 0x4c, 0x45, 0x5f, 0x41, 0x44, 0x53,
	0x5f, 0x41, 0x50, 0x49, 0x10, 0x02, 0x12, 0x19, 0x0a, 0x15, 0x47, 0x4f, 0x4f, 0x47, 0x4c, 0x45,
	0x5f, 0x41, 0x44, 0x53, 0x5f, 0x57, 0x45, 0x42, 0x5f, 0x43, 0x4c, 0x49, 0x45, 0x4e, 0x54, 0x10,
	0x03, 0x12, 0x16, 0x0a, 0x12, 0x41, 0x44, 0x53, 0x5f, 0x44, 0x41, 0x54, 0x41, 0x5f, 0x43, 0x4f,
	0x4e, 0x4e, 0x45, 0x43, 0x54, 0x4f, 0x52, 0x10, 0x04, 0x42, 0xfb, 0x01, 0x0a, 0x22, 0x63, 0x6f,
	0x6d, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x61, 0x64, 0x73, 0x2e, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x61, 0x64, 0x73, 0x2e, 0x76, 0x31, 0x34, 0x2e, 0x65, 0x6e, 0x75, 0x6d, 0x73,
	0x42, 0x21, 0x4f, 0x66, 0x66, 0x6c, 0x69, 0x6e, 0x65, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x55, 0x70,
	0x6c, 0x6f, 0x61, 0x64, 0x43, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x45, 0x6e, 0x75, 0x6d, 0x50, 0x72,
	0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a, 0x43, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x67, 0x6f,
	0x6c, 0x61, 0x6e, 0x67, 0x2e, 0x6f, 0x72, 0x67, 0x2f, 0x67, 0x65, 0x6e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x2f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61, 0x70, 0x69, 0x73, 0x2f, 0x61, 0x64, 0x73,
	0x2f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61, 0x64, 0x73, 0x2f, 0x76, 0x31, 0x34, 0x2f, 0x65,
	0x6e, 0x75, 0x6d, 0x73, 0x3b, 0x65, 0x6e, 0x75, 0x6d, 0x73, 0xa2, 0x02, 0x03, 0x47, 0x41, 0x41,
	0xaa, 0x02, 0x1e, 0x47, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x41, 0x64, 0x73, 0x2e, 0x47, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x41, 0x64, 0x73, 0x2e, 0x56, 0x31, 0x34, 0x2e, 0x45, 0x6e, 0x75, 0x6d,
	0x73, 0xca, 0x02, 0x1e, 0x47, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x5c, 0x41, 0x64, 0x73, 0x5c, 0x47,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x41, 0x64, 0x73, 0x5c, 0x56, 0x31, 0x34, 0x5c, 0x45, 0x6e, 0x75,
	0x6d, 0x73, 0xea, 0x02, 0x22, 0x47, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x3a, 0x3a, 0x41, 0x64, 0x73,
	0x3a, 0x3a, 0x47, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x41, 0x64, 0x73, 0x3a, 0x3a, 0x56, 0x31, 0x34,
	0x3a, 0x3a, 0x45, 0x6e, 0x75, 0x6d, 0x73, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_enums_offline_event_upload_client_enum_proto_rawDescOnce sync.Once
	file_enums_offline_event_upload_client_enum_proto_rawDescData = file_enums_offline_event_upload_client_enum_proto_rawDesc
)

func file_enums_offline_event_upload_client_enum_proto_rawDescGZIP() []byte {
	file_enums_offline_event_upload_client_enum_proto_rawDescOnce.Do(func() {
		file_enums_offline_event_upload_client_enum_proto_rawDescData = protoimpl.X.CompressGZIP(file_enums_offline_event_upload_client_enum_proto_rawDescData)
	})
	return file_enums_offline_event_upload_client_enum_proto_rawDescData
}

var file_enums_offline_event_upload_client_enum_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_enums_offline_event_upload_client_enum_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_enums_offline_event_upload_client_enum_proto_goTypes = []interface{}{
	(OfflineEventUploadClientEnum_OfflineEventUploadClient)(0), // 0: google.ads.googleads.v14.enums.OfflineEventUploadClientEnum.OfflineEventUploadClient
	(*OfflineEventUploadClientEnum)(nil),                       // 1: google.ads.googleads.v14.enums.OfflineEventUploadClientEnum
}
var file_enums_offline_event_upload_client_enum_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_enums_offline_event_upload_client_enum_proto_init() }
func file_enums_offline_event_upload_client_enum_proto_init() {
	if File_enums_offline_event_upload_client_enum_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_enums_offline_event_upload_client_enum_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*OfflineEventUploadClientEnum); i {
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
			RawDescriptor: file_enums_offline_event_upload_client_enum_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_enums_offline_event_upload_client_enum_proto_goTypes,
		DependencyIndexes: file_enums_offline_event_upload_client_enum_proto_depIdxs,
		EnumInfos:         file_enums_offline_event_upload_client_enum_proto_enumTypes,
		MessageInfos:      file_enums_offline_event_upload_client_enum_proto_msgTypes,
	}.Build()
	File_enums_offline_event_upload_client_enum_proto = out.File
	file_enums_offline_event_upload_client_enum_proto_rawDesc = nil
	file_enums_offline_event_upload_client_enum_proto_goTypes = nil
	file_enums_offline_event_upload_client_enum_proto_depIdxs = nil
}
