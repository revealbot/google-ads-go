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
// 	protoc        v3.19.4
// source: google/ads/googleads/v10/enums/sk_ad_network_user_type.proto

package enums

import (
	_ "google.golang.org/genproto/googleapis/api/annotations"
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

// Enumerates SkAdNetwork user types
type SkAdNetworkUserTypeEnum_SkAdNetworkUserType int32

const (
	// Not specified.
	SkAdNetworkUserTypeEnum_UNSPECIFIED SkAdNetworkUserTypeEnum_SkAdNetworkUserType = 0
	// The value is unknown in this version.
	SkAdNetworkUserTypeEnum_UNKNOWN SkAdNetworkUserTypeEnum_SkAdNetworkUserType = 1
	// The value was not present in the postback or we do not have this data for
	// other reasons.
	SkAdNetworkUserTypeEnum_UNAVAILABLE SkAdNetworkUserTypeEnum_SkAdNetworkUserType = 2
	// The user installed the app for the first time.
	SkAdNetworkUserTypeEnum_NEW_INSTALLER SkAdNetworkUserTypeEnum_SkAdNetworkUserType = 3
	// The user has previously installed the app.
	SkAdNetworkUserTypeEnum_REINSTALLER SkAdNetworkUserTypeEnum_SkAdNetworkUserType = 4
)

// Enum value maps for SkAdNetworkUserTypeEnum_SkAdNetworkUserType.
var (
	SkAdNetworkUserTypeEnum_SkAdNetworkUserType_name = map[int32]string{
		0: "UNSPECIFIED",
		1: "UNKNOWN",
		2: "UNAVAILABLE",
		3: "NEW_INSTALLER",
		4: "REINSTALLER",
	}
	SkAdNetworkUserTypeEnum_SkAdNetworkUserType_value = map[string]int32{
		"UNSPECIFIED":   0,
		"UNKNOWN":       1,
		"UNAVAILABLE":   2,
		"NEW_INSTALLER": 3,
		"REINSTALLER":   4,
	}
)

func (x SkAdNetworkUserTypeEnum_SkAdNetworkUserType) Enum() *SkAdNetworkUserTypeEnum_SkAdNetworkUserType {
	p := new(SkAdNetworkUserTypeEnum_SkAdNetworkUserType)
	*p = x
	return p
}

func (x SkAdNetworkUserTypeEnum_SkAdNetworkUserType) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (SkAdNetworkUserTypeEnum_SkAdNetworkUserType) Descriptor() protoreflect.EnumDescriptor {
	return file_enums_sk_ad_network_user_type_proto_enumTypes[0].Descriptor()
}

func (SkAdNetworkUserTypeEnum_SkAdNetworkUserType) Type() protoreflect.EnumType {
	return &file_enums_sk_ad_network_user_type_proto_enumTypes[0]
}

func (x SkAdNetworkUserTypeEnum_SkAdNetworkUserType) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use SkAdNetworkUserTypeEnum_SkAdNetworkUserType.Descriptor instead.
func (SkAdNetworkUserTypeEnum_SkAdNetworkUserType) EnumDescriptor() ([]byte, []int) {
	return file_enums_sk_ad_network_user_type_proto_rawDescGZIP(), []int{0, 0}
}

// Container for enumeration of SkAdNetwork user types.
type SkAdNetworkUserTypeEnum struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *SkAdNetworkUserTypeEnum) Reset() {
	*x = SkAdNetworkUserTypeEnum{}
	if protoimpl.UnsafeEnabled {
		mi := &file_enums_sk_ad_network_user_type_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SkAdNetworkUserTypeEnum) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SkAdNetworkUserTypeEnum) ProtoMessage() {}

func (x *SkAdNetworkUserTypeEnum) ProtoReflect() protoreflect.Message {
	mi := &file_enums_sk_ad_network_user_type_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SkAdNetworkUserTypeEnum.ProtoReflect.Descriptor instead.
func (*SkAdNetworkUserTypeEnum) Descriptor() ([]byte, []int) {
	return file_enums_sk_ad_network_user_type_proto_rawDescGZIP(), []int{0}
}

var File_enums_sk_ad_network_user_type_proto protoreflect.FileDescriptor

var file_enums_sk_ad_network_user_type_proto_rawDesc = []byte{
	0x0a, 0x3c, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x64, 0x73, 0x2f, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x61, 0x64, 0x73, 0x2f, 0x76, 0x31, 0x30, 0x2f, 0x65, 0x6e, 0x75, 0x6d, 0x73,
	0x2f, 0x73, 0x6b, 0x5f, 0x61, 0x64, 0x5f, 0x6e, 0x65, 0x74, 0x77, 0x6f, 0x72, 0x6b, 0x5f, 0x75,
	0x73, 0x65, 0x72, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x1e,
	0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x61, 0x64, 0x73, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x61, 0x64, 0x73, 0x2e, 0x76, 0x31, 0x30, 0x2e, 0x65, 0x6e, 0x75, 0x6d, 0x73, 0x1a, 0x1c,
	0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x83, 0x01, 0x0a,
	0x17, 0x53, 0x6b, 0x41, 0x64, 0x4e, 0x65, 0x74, 0x77, 0x6f, 0x72, 0x6b, 0x55, 0x73, 0x65, 0x72,
	0x54, 0x79, 0x70, 0x65, 0x45, 0x6e, 0x75, 0x6d, 0x22, 0x68, 0x0a, 0x13, 0x53, 0x6b, 0x41, 0x64,
	0x4e, 0x65, 0x74, 0x77, 0x6f, 0x72, 0x6b, 0x55, 0x73, 0x65, 0x72, 0x54, 0x79, 0x70, 0x65, 0x12,
	0x0f, 0x0a, 0x0b, 0x55, 0x4e, 0x53, 0x50, 0x45, 0x43, 0x49, 0x46, 0x49, 0x45, 0x44, 0x10, 0x00,
	0x12, 0x0b, 0x0a, 0x07, 0x55, 0x4e, 0x4b, 0x4e, 0x4f, 0x57, 0x4e, 0x10, 0x01, 0x12, 0x0f, 0x0a,
	0x0b, 0x55, 0x4e, 0x41, 0x56, 0x41, 0x49, 0x4c, 0x41, 0x42, 0x4c, 0x45, 0x10, 0x02, 0x12, 0x11,
	0x0a, 0x0d, 0x4e, 0x45, 0x57, 0x5f, 0x49, 0x4e, 0x53, 0x54, 0x41, 0x4c, 0x4c, 0x45, 0x52, 0x10,
	0x03, 0x12, 0x0f, 0x0a, 0x0b, 0x52, 0x45, 0x49, 0x4e, 0x53, 0x54, 0x41, 0x4c, 0x4c, 0x45, 0x52,
	0x10, 0x04, 0x42, 0xf2, 0x01, 0x0a, 0x22, 0x63, 0x6f, 0x6d, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x2e, 0x61, 0x64, 0x73, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61, 0x64, 0x73, 0x2e,
	0x76, 0x31, 0x30, 0x2e, 0x65, 0x6e, 0x75, 0x6d, 0x73, 0x42, 0x18, 0x53, 0x6b, 0x41, 0x64, 0x4e,
	0x65, 0x74, 0x77, 0x6f, 0x72, 0x6b, 0x55, 0x73, 0x65, 0x72, 0x54, 0x79, 0x70, 0x65, 0x50, 0x72,
	0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a, 0x43, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x67, 0x6f,
	0x6c, 0x61, 0x6e, 0x67, 0x2e, 0x6f, 0x72, 0x67, 0x2f, 0x67, 0x65, 0x6e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x2f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61, 0x70, 0x69, 0x73, 0x2f, 0x61, 0x64, 0x73,
	0x2f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61, 0x64, 0x73, 0x2f, 0x76, 0x31, 0x30, 0x2f, 0x65,
	0x6e, 0x75, 0x6d, 0x73, 0x3b, 0x65, 0x6e, 0x75, 0x6d, 0x73, 0xa2, 0x02, 0x03, 0x47, 0x41, 0x41,
	0xaa, 0x02, 0x1e, 0x47, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x41, 0x64, 0x73, 0x2e, 0x47, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x41, 0x64, 0x73, 0x2e, 0x56, 0x31, 0x30, 0x2e, 0x45, 0x6e, 0x75, 0x6d,
	0x73, 0xca, 0x02, 0x1e, 0x47, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x5c, 0x41, 0x64, 0x73, 0x5c, 0x47,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x41, 0x64, 0x73, 0x5c, 0x56, 0x31, 0x30, 0x5c, 0x45, 0x6e, 0x75,
	0x6d, 0x73, 0xea, 0x02, 0x22, 0x47, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x3a, 0x3a, 0x41, 0x64, 0x73,
	0x3a, 0x3a, 0x47, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x41, 0x64, 0x73, 0x3a, 0x3a, 0x56, 0x31, 0x30,
	0x3a, 0x3a, 0x45, 0x6e, 0x75, 0x6d, 0x73, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_enums_sk_ad_network_user_type_proto_rawDescOnce sync.Once
	file_enums_sk_ad_network_user_type_proto_rawDescData = file_enums_sk_ad_network_user_type_proto_rawDesc
)

func file_enums_sk_ad_network_user_type_proto_rawDescGZIP() []byte {
	file_enums_sk_ad_network_user_type_proto_rawDescOnce.Do(func() {
		file_enums_sk_ad_network_user_type_proto_rawDescData = protoimpl.X.CompressGZIP(file_enums_sk_ad_network_user_type_proto_rawDescData)
	})
	return file_enums_sk_ad_network_user_type_proto_rawDescData
}

var file_enums_sk_ad_network_user_type_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_enums_sk_ad_network_user_type_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_enums_sk_ad_network_user_type_proto_goTypes = []interface{}{
	(SkAdNetworkUserTypeEnum_SkAdNetworkUserType)(0), // 0: google.ads.googleads.v10.enums.SkAdNetworkUserTypeEnum.SkAdNetworkUserType
	(*SkAdNetworkUserTypeEnum)(nil),                  // 1: google.ads.googleads.v10.enums.SkAdNetworkUserTypeEnum
}
var file_enums_sk_ad_network_user_type_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_enums_sk_ad_network_user_type_proto_init() }
func file_enums_sk_ad_network_user_type_proto_init() {
	if File_enums_sk_ad_network_user_type_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_enums_sk_ad_network_user_type_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SkAdNetworkUserTypeEnum); i {
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
			RawDescriptor: file_enums_sk_ad_network_user_type_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_enums_sk_ad_network_user_type_proto_goTypes,
		DependencyIndexes: file_enums_sk_ad_network_user_type_proto_depIdxs,
		EnumInfos:         file_enums_sk_ad_network_user_type_proto_enumTypes,
		MessageInfos:      file_enums_sk_ad_network_user_type_proto_msgTypes,
	}.Build()
	File_enums_sk_ad_network_user_type_proto = out.File
	file_enums_sk_ad_network_user_type_proto_rawDesc = nil
	file_enums_sk_ad_network_user_type_proto_goTypes = nil
	file_enums_sk_ad_network_user_type_proto_depIdxs = nil
}
