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
// 	protoc        v3.21.7
// source: google/ads/googleads/v12/enums/hotel_rate_type.proto

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

// Enum describing possible hotel rate types.
type HotelRateTypeEnum_HotelRateType int32

const (
	// Not specified.
	HotelRateTypeEnum_UNSPECIFIED HotelRateTypeEnum_HotelRateType = 0
	// The value is unknown in this version.
	HotelRateTypeEnum_UNKNOWN HotelRateTypeEnum_HotelRateType = 1
	// Rate type information is unavailable.
	HotelRateTypeEnum_UNAVAILABLE HotelRateTypeEnum_HotelRateType = 2
	// Rates available to everyone.
	HotelRateTypeEnum_PUBLIC_RATE HotelRateTypeEnum_HotelRateType = 3
	// A membership program rate is available and satisfies basic requirements
	// like having a public rate available. UI treatment will strikethrough the
	// public rate and indicate that a discount is available to the user. For
	// more on Qualified Rates, visit
	// https://developers.google.com/hotels/hotel-ads/dev-guide/qualified-rates
	HotelRateTypeEnum_QUALIFIED_RATE HotelRateTypeEnum_HotelRateType = 4
	// Rates available to users that satisfy some eligibility criteria, for
	// example, all signed-in users, 20% of mobile users, all mobile users in
	// Canada, etc.
	HotelRateTypeEnum_PRIVATE_RATE HotelRateTypeEnum_HotelRateType = 5
)

// Enum value maps for HotelRateTypeEnum_HotelRateType.
var (
	HotelRateTypeEnum_HotelRateType_name = map[int32]string{
		0: "UNSPECIFIED",
		1: "UNKNOWN",
		2: "UNAVAILABLE",
		3: "PUBLIC_RATE",
		4: "QUALIFIED_RATE",
		5: "PRIVATE_RATE",
	}
	HotelRateTypeEnum_HotelRateType_value = map[string]int32{
		"UNSPECIFIED":    0,
		"UNKNOWN":        1,
		"UNAVAILABLE":    2,
		"PUBLIC_RATE":    3,
		"QUALIFIED_RATE": 4,
		"PRIVATE_RATE":   5,
	}
)

func (x HotelRateTypeEnum_HotelRateType) Enum() *HotelRateTypeEnum_HotelRateType {
	p := new(HotelRateTypeEnum_HotelRateType)
	*p = x
	return p
}

func (x HotelRateTypeEnum_HotelRateType) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (HotelRateTypeEnum_HotelRateType) Descriptor() protoreflect.EnumDescriptor {
	return file_enums_hotel_rate_type_proto_enumTypes[0].Descriptor()
}

func (HotelRateTypeEnum_HotelRateType) Type() protoreflect.EnumType {
	return &file_enums_hotel_rate_type_proto_enumTypes[0]
}

func (x HotelRateTypeEnum_HotelRateType) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use HotelRateTypeEnum_HotelRateType.Descriptor instead.
func (HotelRateTypeEnum_HotelRateType) EnumDescriptor() ([]byte, []int) {
	return file_enums_hotel_rate_type_proto_rawDescGZIP(), []int{0, 0}
}

// Container for enum describing possible hotel rate types.
type HotelRateTypeEnum struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *HotelRateTypeEnum) Reset() {
	*x = HotelRateTypeEnum{}
	if protoimpl.UnsafeEnabled {
		mi := &file_enums_hotel_rate_type_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *HotelRateTypeEnum) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*HotelRateTypeEnum) ProtoMessage() {}

func (x *HotelRateTypeEnum) ProtoReflect() protoreflect.Message {
	mi := &file_enums_hotel_rate_type_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use HotelRateTypeEnum.ProtoReflect.Descriptor instead.
func (*HotelRateTypeEnum) Descriptor() ([]byte, []int) {
	return file_enums_hotel_rate_type_proto_rawDescGZIP(), []int{0}
}

var File_enums_hotel_rate_type_proto protoreflect.FileDescriptor

var file_enums_hotel_rate_type_proto_rawDesc = []byte{
	0x0a, 0x34, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x64, 0x73, 0x2f, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x61, 0x64, 0x73, 0x2f, 0x76, 0x31, 0x32, 0x2f, 0x65, 0x6e, 0x75, 0x6d, 0x73,
	0x2f, 0x68, 0x6f, 0x74, 0x65, 0x6c, 0x5f, 0x72, 0x61, 0x74, 0x65, 0x5f, 0x74, 0x79, 0x70, 0x65,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x1e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x61,
	0x64, 0x73, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61, 0x64, 0x73, 0x2e, 0x76, 0x31, 0x32,
	0x2e, 0x65, 0x6e, 0x75, 0x6d, 0x73, 0x22, 0x8a, 0x01, 0x0a, 0x11, 0x48, 0x6f, 0x74, 0x65, 0x6c,
	0x52, 0x61, 0x74, 0x65, 0x54, 0x79, 0x70, 0x65, 0x45, 0x6e, 0x75, 0x6d, 0x22, 0x75, 0x0a, 0x0d,
	0x48, 0x6f, 0x74, 0x65, 0x6c, 0x52, 0x61, 0x74, 0x65, 0x54, 0x79, 0x70, 0x65, 0x12, 0x0f, 0x0a,
	0x0b, 0x55, 0x4e, 0x53, 0x50, 0x45, 0x43, 0x49, 0x46, 0x49, 0x45, 0x44, 0x10, 0x00, 0x12, 0x0b,
	0x0a, 0x07, 0x55, 0x4e, 0x4b, 0x4e, 0x4f, 0x57, 0x4e, 0x10, 0x01, 0x12, 0x0f, 0x0a, 0x0b, 0x55,
	0x4e, 0x41, 0x56, 0x41, 0x49, 0x4c, 0x41, 0x42, 0x4c, 0x45, 0x10, 0x02, 0x12, 0x0f, 0x0a, 0x0b,
	0x50, 0x55, 0x42, 0x4c, 0x49, 0x43, 0x5f, 0x52, 0x41, 0x54, 0x45, 0x10, 0x03, 0x12, 0x12, 0x0a,
	0x0e, 0x51, 0x55, 0x41, 0x4c, 0x49, 0x46, 0x49, 0x45, 0x44, 0x5f, 0x52, 0x41, 0x54, 0x45, 0x10,
	0x04, 0x12, 0x10, 0x0a, 0x0c, 0x50, 0x52, 0x49, 0x56, 0x41, 0x54, 0x45, 0x5f, 0x52, 0x41, 0x54,
	0x45, 0x10, 0x05, 0x42, 0xec, 0x01, 0x0a, 0x22, 0x63, 0x6f, 0x6d, 0x2e, 0x67, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x2e, 0x61, 0x64, 0x73, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61, 0x64, 0x73,
	0x2e, 0x76, 0x31, 0x32, 0x2e, 0x65, 0x6e, 0x75, 0x6d, 0x73, 0x42, 0x12, 0x48, 0x6f, 0x74, 0x65,
	0x6c, 0x52, 0x61, 0x74, 0x65, 0x54, 0x79, 0x70, 0x65, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01,
	0x5a, 0x43, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x67, 0x6f, 0x6c, 0x61, 0x6e, 0x67, 0x2e,
	0x6f, 0x72, 0x67, 0x2f, 0x67, 0x65, 0x6e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x61, 0x70, 0x69, 0x73, 0x2f, 0x61, 0x64, 0x73, 0x2f, 0x67, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x61, 0x64, 0x73, 0x2f, 0x76, 0x31, 0x32, 0x2f, 0x65, 0x6e, 0x75, 0x6d, 0x73, 0x3b,
	0x65, 0x6e, 0x75, 0x6d, 0x73, 0xa2, 0x02, 0x03, 0x47, 0x41, 0x41, 0xaa, 0x02, 0x1e, 0x47, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x41, 0x64, 0x73, 0x2e, 0x47, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x41,
	0x64, 0x73, 0x2e, 0x56, 0x31, 0x32, 0x2e, 0x45, 0x6e, 0x75, 0x6d, 0x73, 0xca, 0x02, 0x1e, 0x47,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x5c, 0x41, 0x64, 0x73, 0x5c, 0x47, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x41, 0x64, 0x73, 0x5c, 0x56, 0x31, 0x32, 0x5c, 0x45, 0x6e, 0x75, 0x6d, 0x73, 0xea, 0x02, 0x22,
	0x47, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x3a, 0x3a, 0x41, 0x64, 0x73, 0x3a, 0x3a, 0x47, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x41, 0x64, 0x73, 0x3a, 0x3a, 0x56, 0x31, 0x32, 0x3a, 0x3a, 0x45, 0x6e, 0x75,
	0x6d, 0x73, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_enums_hotel_rate_type_proto_rawDescOnce sync.Once
	file_enums_hotel_rate_type_proto_rawDescData = file_enums_hotel_rate_type_proto_rawDesc
)

func file_enums_hotel_rate_type_proto_rawDescGZIP() []byte {
	file_enums_hotel_rate_type_proto_rawDescOnce.Do(func() {
		file_enums_hotel_rate_type_proto_rawDescData = protoimpl.X.CompressGZIP(file_enums_hotel_rate_type_proto_rawDescData)
	})
	return file_enums_hotel_rate_type_proto_rawDescData
}

var file_enums_hotel_rate_type_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_enums_hotel_rate_type_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_enums_hotel_rate_type_proto_goTypes = []interface{}{
	(HotelRateTypeEnum_HotelRateType)(0), // 0: google.ads.googleads.v12.enums.HotelRateTypeEnum.HotelRateType
	(*HotelRateTypeEnum)(nil),            // 1: google.ads.googleads.v12.enums.HotelRateTypeEnum
}
var file_enums_hotel_rate_type_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_enums_hotel_rate_type_proto_init() }
func file_enums_hotel_rate_type_proto_init() {
	if File_enums_hotel_rate_type_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_enums_hotel_rate_type_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*HotelRateTypeEnum); i {
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
			RawDescriptor: file_enums_hotel_rate_type_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_enums_hotel_rate_type_proto_goTypes,
		DependencyIndexes: file_enums_hotel_rate_type_proto_depIdxs,
		EnumInfos:         file_enums_hotel_rate_type_proto_enumTypes,
		MessageInfos:      file_enums_hotel_rate_type_proto_msgTypes,
	}.Build()
	File_enums_hotel_rate_type_proto = out.File
	file_enums_hotel_rate_type_proto_rawDesc = nil
	file_enums_hotel_rate_type_proto_goTypes = nil
	file_enums_hotel_rate_type_proto_depIdxs = nil
}
