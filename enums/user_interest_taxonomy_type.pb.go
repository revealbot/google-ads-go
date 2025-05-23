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
// source: google/ads/googleads/v19/enums/user_interest_taxonomy_type.proto

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

// Enum containing the possible UserInterestTaxonomyTypes.
type UserInterestTaxonomyTypeEnum_UserInterestTaxonomyType int32

const (
	// Not specified.
	UserInterestTaxonomyTypeEnum_UNSPECIFIED UserInterestTaxonomyTypeEnum_UserInterestTaxonomyType = 0
	// Used for return value only. Represents value unknown in this version.
	UserInterestTaxonomyTypeEnum_UNKNOWN UserInterestTaxonomyTypeEnum_UserInterestTaxonomyType = 1
	// The affinity for this user interest.
	UserInterestTaxonomyTypeEnum_AFFINITY UserInterestTaxonomyTypeEnum_UserInterestTaxonomyType = 2
	// The market for this user interest.
	UserInterestTaxonomyTypeEnum_IN_MARKET UserInterestTaxonomyTypeEnum_UserInterestTaxonomyType = 3
	// Users known to have installed applications in the specified categories.
	UserInterestTaxonomyTypeEnum_MOBILE_APP_INSTALL_USER UserInterestTaxonomyTypeEnum_UserInterestTaxonomyType = 4
	// The geographical location of the interest-based vertical.
	UserInterestTaxonomyTypeEnum_VERTICAL_GEO UserInterestTaxonomyTypeEnum_UserInterestTaxonomyType = 5
	// User interest criteria for new smart phone users.
	UserInterestTaxonomyTypeEnum_NEW_SMART_PHONE_USER UserInterestTaxonomyTypeEnum_UserInterestTaxonomyType = 6
)

// Enum value maps for UserInterestTaxonomyTypeEnum_UserInterestTaxonomyType.
var (
	UserInterestTaxonomyTypeEnum_UserInterestTaxonomyType_name = map[int32]string{
		0: "UNSPECIFIED",
		1: "UNKNOWN",
		2: "AFFINITY",
		3: "IN_MARKET",
		4: "MOBILE_APP_INSTALL_USER",
		5: "VERTICAL_GEO",
		6: "NEW_SMART_PHONE_USER",
	}
	UserInterestTaxonomyTypeEnum_UserInterestTaxonomyType_value = map[string]int32{
		"UNSPECIFIED":             0,
		"UNKNOWN":                 1,
		"AFFINITY":                2,
		"IN_MARKET":               3,
		"MOBILE_APP_INSTALL_USER": 4,
		"VERTICAL_GEO":            5,
		"NEW_SMART_PHONE_USER":    6,
	}
)

func (x UserInterestTaxonomyTypeEnum_UserInterestTaxonomyType) Enum() *UserInterestTaxonomyTypeEnum_UserInterestTaxonomyType {
	p := new(UserInterestTaxonomyTypeEnum_UserInterestTaxonomyType)
	*p = x
	return p
}

func (x UserInterestTaxonomyTypeEnum_UserInterestTaxonomyType) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (UserInterestTaxonomyTypeEnum_UserInterestTaxonomyType) Descriptor() protoreflect.EnumDescriptor {
	return file_enums_user_interest_taxonomy_type_proto_enumTypes[0].Descriptor()
}

func (UserInterestTaxonomyTypeEnum_UserInterestTaxonomyType) Type() protoreflect.EnumType {
	return &file_enums_user_interest_taxonomy_type_proto_enumTypes[0]
}

func (x UserInterestTaxonomyTypeEnum_UserInterestTaxonomyType) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use UserInterestTaxonomyTypeEnum_UserInterestTaxonomyType.Descriptor instead.
func (UserInterestTaxonomyTypeEnum_UserInterestTaxonomyType) EnumDescriptor() ([]byte, []int) {
	return file_enums_user_interest_taxonomy_type_proto_rawDescGZIP(), []int{0, 0}
}

// Message describing a UserInterestTaxonomyType.
type UserInterestTaxonomyTypeEnum struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *UserInterestTaxonomyTypeEnum) Reset() {
	*x = UserInterestTaxonomyTypeEnum{}
	mi := &file_enums_user_interest_taxonomy_type_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *UserInterestTaxonomyTypeEnum) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UserInterestTaxonomyTypeEnum) ProtoMessage() {}

func (x *UserInterestTaxonomyTypeEnum) ProtoReflect() protoreflect.Message {
	mi := &file_enums_user_interest_taxonomy_type_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UserInterestTaxonomyTypeEnum.ProtoReflect.Descriptor instead.
func (*UserInterestTaxonomyTypeEnum) Descriptor() ([]byte, []int) {
	return file_enums_user_interest_taxonomy_type_proto_rawDescGZIP(), []int{0}
}

var File_enums_user_interest_taxonomy_type_proto protoreflect.FileDescriptor

var file_enums_user_interest_taxonomy_type_proto_rawDesc = string([]byte{
	0x0a, 0x40, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x64, 0x73, 0x2f, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x61, 0x64, 0x73, 0x2f, 0x76, 0x31, 0x39, 0x2f, 0x65, 0x6e, 0x75, 0x6d, 0x73,
	0x2f, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x65, 0x73, 0x74, 0x5f, 0x74,
	0x61, 0x78, 0x6f, 0x6e, 0x6f, 0x6d, 0x79, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x12, 0x1e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x61, 0x64, 0x73, 0x2e, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61, 0x64, 0x73, 0x2e, 0x76, 0x31, 0x39, 0x2e, 0x65, 0x6e, 0x75,
	0x6d, 0x73, 0x22, 0xbf, 0x01, 0x0a, 0x1c, 0x55, 0x73, 0x65, 0x72, 0x49, 0x6e, 0x74, 0x65, 0x72,
	0x65, 0x73, 0x74, 0x54, 0x61, 0x78, 0x6f, 0x6e, 0x6f, 0x6d, 0x79, 0x54, 0x79, 0x70, 0x65, 0x45,
	0x6e, 0x75, 0x6d, 0x22, 0x9e, 0x01, 0x0a, 0x18, 0x55, 0x73, 0x65, 0x72, 0x49, 0x6e, 0x74, 0x65,
	0x72, 0x65, 0x73, 0x74, 0x54, 0x61, 0x78, 0x6f, 0x6e, 0x6f, 0x6d, 0x79, 0x54, 0x79, 0x70, 0x65,
	0x12, 0x0f, 0x0a, 0x0b, 0x55, 0x4e, 0x53, 0x50, 0x45, 0x43, 0x49, 0x46, 0x49, 0x45, 0x44, 0x10,
	0x00, 0x12, 0x0b, 0x0a, 0x07, 0x55, 0x4e, 0x4b, 0x4e, 0x4f, 0x57, 0x4e, 0x10, 0x01, 0x12, 0x0c,
	0x0a, 0x08, 0x41, 0x46, 0x46, 0x49, 0x4e, 0x49, 0x54, 0x59, 0x10, 0x02, 0x12, 0x0d, 0x0a, 0x09,
	0x49, 0x4e, 0x5f, 0x4d, 0x41, 0x52, 0x4b, 0x45, 0x54, 0x10, 0x03, 0x12, 0x1b, 0x0a, 0x17, 0x4d,
	0x4f, 0x42, 0x49, 0x4c, 0x45, 0x5f, 0x41, 0x50, 0x50, 0x5f, 0x49, 0x4e, 0x53, 0x54, 0x41, 0x4c,
	0x4c, 0x5f, 0x55, 0x53, 0x45, 0x52, 0x10, 0x04, 0x12, 0x10, 0x0a, 0x0c, 0x56, 0x45, 0x52, 0x54,
	0x49, 0x43, 0x41, 0x4c, 0x5f, 0x47, 0x45, 0x4f, 0x10, 0x05, 0x12, 0x18, 0x0a, 0x14, 0x4e, 0x45,
	0x57, 0x5f, 0x53, 0x4d, 0x41, 0x52, 0x54, 0x5f, 0x50, 0x48, 0x4f, 0x4e, 0x45, 0x5f, 0x55, 0x53,
	0x45, 0x52, 0x10, 0x06, 0x42, 0xf7, 0x01, 0x0a, 0x22, 0x63, 0x6f, 0x6d, 0x2e, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x2e, 0x61, 0x64, 0x73, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61, 0x64,
	0x73, 0x2e, 0x76, 0x31, 0x39, 0x2e, 0x65, 0x6e, 0x75, 0x6d, 0x73, 0x42, 0x1d, 0x55, 0x73, 0x65,
	0x72, 0x49, 0x6e, 0x74, 0x65, 0x72, 0x65, 0x73, 0x74, 0x54, 0x61, 0x78, 0x6f, 0x6e, 0x6f, 0x6d,
	0x79, 0x54, 0x79, 0x70, 0x65, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a, 0x43, 0x67, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x67, 0x6f, 0x6c, 0x61, 0x6e, 0x67, 0x2e, 0x6f, 0x72, 0x67, 0x2f,
	0x67, 0x65, 0x6e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61,
	0x70, 0x69, 0x73, 0x2f, 0x61, 0x64, 0x73, 0x2f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61, 0x64,
	0x73, 0x2f, 0x76, 0x31, 0x39, 0x2f, 0x65, 0x6e, 0x75, 0x6d, 0x73, 0x3b, 0x65, 0x6e, 0x75, 0x6d,
	0x73, 0xa2, 0x02, 0x03, 0x47, 0x41, 0x41, 0xaa, 0x02, 0x1e, 0x47, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x2e, 0x41, 0x64, 0x73, 0x2e, 0x47, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x41, 0x64, 0x73, 0x2e, 0x56,
	0x31, 0x39, 0x2e, 0x45, 0x6e, 0x75, 0x6d, 0x73, 0xca, 0x02, 0x1e, 0x47, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x5c, 0x41, 0x64, 0x73, 0x5c, 0x47, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x41, 0x64, 0x73, 0x5c,
	0x56, 0x31, 0x39, 0x5c, 0x45, 0x6e, 0x75, 0x6d, 0x73, 0xea, 0x02, 0x22, 0x47, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x3a, 0x3a, 0x41, 0x64, 0x73, 0x3a, 0x3a, 0x47, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x41,
	0x64, 0x73, 0x3a, 0x3a, 0x56, 0x31, 0x39, 0x3a, 0x3a, 0x45, 0x6e, 0x75, 0x6d, 0x73, 0x62, 0x06,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
})

var (
	file_enums_user_interest_taxonomy_type_proto_rawDescOnce sync.Once
	file_enums_user_interest_taxonomy_type_proto_rawDescData []byte
)

func file_enums_user_interest_taxonomy_type_proto_rawDescGZIP() []byte {
	file_enums_user_interest_taxonomy_type_proto_rawDescOnce.Do(func() {
		file_enums_user_interest_taxonomy_type_proto_rawDescData = protoimpl.X.CompressGZIP(unsafe.Slice(unsafe.StringData(file_enums_user_interest_taxonomy_type_proto_rawDesc), len(file_enums_user_interest_taxonomy_type_proto_rawDesc)))
	})
	return file_enums_user_interest_taxonomy_type_proto_rawDescData
}

var file_enums_user_interest_taxonomy_type_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_enums_user_interest_taxonomy_type_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_enums_user_interest_taxonomy_type_proto_goTypes = []any{
	(UserInterestTaxonomyTypeEnum_UserInterestTaxonomyType)(0), // 0: google.ads.googleads.v19.enums.UserInterestTaxonomyTypeEnum.UserInterestTaxonomyType
	(*UserInterestTaxonomyTypeEnum)(nil),                       // 1: google.ads.googleads.v19.enums.UserInterestTaxonomyTypeEnum
}
var file_enums_user_interest_taxonomy_type_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_enums_user_interest_taxonomy_type_proto_init() }
func file_enums_user_interest_taxonomy_type_proto_init() {
	if File_enums_user_interest_taxonomy_type_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: unsafe.Slice(unsafe.StringData(file_enums_user_interest_taxonomy_type_proto_rawDesc), len(file_enums_user_interest_taxonomy_type_proto_rawDesc)),
			NumEnums:      1,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_enums_user_interest_taxonomy_type_proto_goTypes,
		DependencyIndexes: file_enums_user_interest_taxonomy_type_proto_depIdxs,
		EnumInfos:         file_enums_user_interest_taxonomy_type_proto_enumTypes,
		MessageInfos:      file_enums_user_interest_taxonomy_type_proto_msgTypes,
	}.Build()
	File_enums_user_interest_taxonomy_type_proto = out.File
	file_enums_user_interest_taxonomy_type_proto_goTypes = nil
	file_enums_user_interest_taxonomy_type_proto_depIdxs = nil
}
