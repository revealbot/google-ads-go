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
// source: google/ads/googleads/v10/enums/hotel_price_bucket.proto

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

// Enum describing possible hotel price buckets.
type HotelPriceBucketEnum_HotelPriceBucket int32

const (
	// Not specified.
	HotelPriceBucketEnum_UNSPECIFIED HotelPriceBucketEnum_HotelPriceBucket = 0
	// The value is unknown in this version.
	HotelPriceBucketEnum_UNKNOWN HotelPriceBucketEnum_HotelPriceBucket = 1
	// Uniquely lowest price. Partner has the lowest price, and no other
	// partners are within a small variance of that price.
	HotelPriceBucketEnum_LOWEST_UNIQUE HotelPriceBucketEnum_HotelPriceBucket = 2
	// Tied for lowest price. Partner is within a small variance of the lowest
	// price.
	HotelPriceBucketEnum_LOWEST_TIED HotelPriceBucketEnum_HotelPriceBucket = 3
	// Not lowest price. Partner is not within a small variance of the lowest
	// price.
	HotelPriceBucketEnum_NOT_LOWEST HotelPriceBucketEnum_HotelPriceBucket = 4
	// Partner was the only one shown.
	HotelPriceBucketEnum_ONLY_PARTNER_SHOWN HotelPriceBucketEnum_HotelPriceBucket = 5
)

// Enum value maps for HotelPriceBucketEnum_HotelPriceBucket.
var (
	HotelPriceBucketEnum_HotelPriceBucket_name = map[int32]string{
		0: "UNSPECIFIED",
		1: "UNKNOWN",
		2: "LOWEST_UNIQUE",
		3: "LOWEST_TIED",
		4: "NOT_LOWEST",
		5: "ONLY_PARTNER_SHOWN",
	}
	HotelPriceBucketEnum_HotelPriceBucket_value = map[string]int32{
		"UNSPECIFIED":        0,
		"UNKNOWN":            1,
		"LOWEST_UNIQUE":      2,
		"LOWEST_TIED":        3,
		"NOT_LOWEST":         4,
		"ONLY_PARTNER_SHOWN": 5,
	}
)

func (x HotelPriceBucketEnum_HotelPriceBucket) Enum() *HotelPriceBucketEnum_HotelPriceBucket {
	p := new(HotelPriceBucketEnum_HotelPriceBucket)
	*p = x
	return p
}

func (x HotelPriceBucketEnum_HotelPriceBucket) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (HotelPriceBucketEnum_HotelPriceBucket) Descriptor() protoreflect.EnumDescriptor {
	return file_enums_hotel_price_bucket_proto_enumTypes[0].Descriptor()
}

func (HotelPriceBucketEnum_HotelPriceBucket) Type() protoreflect.EnumType {
	return &file_enums_hotel_price_bucket_proto_enumTypes[0]
}

func (x HotelPriceBucketEnum_HotelPriceBucket) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use HotelPriceBucketEnum_HotelPriceBucket.Descriptor instead.
func (HotelPriceBucketEnum_HotelPriceBucket) EnumDescriptor() ([]byte, []int) {
	return file_enums_hotel_price_bucket_proto_rawDescGZIP(), []int{0, 0}
}

// Container for enum describing hotel price bucket for a hotel itinerary.
type HotelPriceBucketEnum struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *HotelPriceBucketEnum) Reset() {
	*x = HotelPriceBucketEnum{}
	if protoimpl.UnsafeEnabled {
		mi := &file_enums_hotel_price_bucket_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *HotelPriceBucketEnum) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*HotelPriceBucketEnum) ProtoMessage() {}

func (x *HotelPriceBucketEnum) ProtoReflect() protoreflect.Message {
	mi := &file_enums_hotel_price_bucket_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use HotelPriceBucketEnum.ProtoReflect.Descriptor instead.
func (*HotelPriceBucketEnum) Descriptor() ([]byte, []int) {
	return file_enums_hotel_price_bucket_proto_rawDescGZIP(), []int{0}
}

var File_enums_hotel_price_bucket_proto protoreflect.FileDescriptor

var file_enums_hotel_price_bucket_proto_rawDesc = []byte{
	0x0a, 0x37, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x64, 0x73, 0x2f, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x61, 0x64, 0x73, 0x2f, 0x76, 0x31, 0x30, 0x2f, 0x65, 0x6e, 0x75, 0x6d, 0x73,
	0x2f, 0x68, 0x6f, 0x74, 0x65, 0x6c, 0x5f, 0x70, 0x72, 0x69, 0x63, 0x65, 0x5f, 0x62, 0x75, 0x63,
	0x6b, 0x65, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x1e, 0x67, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x2e, 0x61, 0x64, 0x73, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61, 0x64, 0x73, 0x2e,
	0x76, 0x31, 0x30, 0x2e, 0x65, 0x6e, 0x75, 0x6d, 0x73, 0x1a, 0x1c, 0x67, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x94, 0x01, 0x0a, 0x14, 0x48, 0x6f, 0x74, 0x65,
	0x6c, 0x50, 0x72, 0x69, 0x63, 0x65, 0x42, 0x75, 0x63, 0x6b, 0x65, 0x74, 0x45, 0x6e, 0x75, 0x6d,
	0x22, 0x7c, 0x0a, 0x10, 0x48, 0x6f, 0x74, 0x65, 0x6c, 0x50, 0x72, 0x69, 0x63, 0x65, 0x42, 0x75,
	0x63, 0x6b, 0x65, 0x74, 0x12, 0x0f, 0x0a, 0x0b, 0x55, 0x4e, 0x53, 0x50, 0x45, 0x43, 0x49, 0x46,
	0x49, 0x45, 0x44, 0x10, 0x00, 0x12, 0x0b, 0x0a, 0x07, 0x55, 0x4e, 0x4b, 0x4e, 0x4f, 0x57, 0x4e,
	0x10, 0x01, 0x12, 0x11, 0x0a, 0x0d, 0x4c, 0x4f, 0x57, 0x45, 0x53, 0x54, 0x5f, 0x55, 0x4e, 0x49,
	0x51, 0x55, 0x45, 0x10, 0x02, 0x12, 0x0f, 0x0a, 0x0b, 0x4c, 0x4f, 0x57, 0x45, 0x53, 0x54, 0x5f,
	0x54, 0x49, 0x45, 0x44, 0x10, 0x03, 0x12, 0x0e, 0x0a, 0x0a, 0x4e, 0x4f, 0x54, 0x5f, 0x4c, 0x4f,
	0x57, 0x45, 0x53, 0x54, 0x10, 0x04, 0x12, 0x16, 0x0a, 0x12, 0x4f, 0x4e, 0x4c, 0x59, 0x5f, 0x50,
	0x41, 0x52, 0x54, 0x4e, 0x45, 0x52, 0x5f, 0x53, 0x48, 0x4f, 0x57, 0x4e, 0x10, 0x05, 0x42, 0xef,
	0x01, 0x0a, 0x22, 0x63, 0x6f, 0x6d, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x61, 0x64,
	0x73, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61, 0x64, 0x73, 0x2e, 0x76, 0x31, 0x30, 0x2e,
	0x65, 0x6e, 0x75, 0x6d, 0x73, 0x42, 0x15, 0x48, 0x6f, 0x74, 0x65, 0x6c, 0x50, 0x72, 0x69, 0x63,
	0x65, 0x42, 0x75, 0x63, 0x6b, 0x65, 0x74, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a, 0x43,
	0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x67, 0x6f, 0x6c, 0x61, 0x6e, 0x67, 0x2e, 0x6f, 0x72,
	0x67, 0x2f, 0x67, 0x65, 0x6e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x67, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x61, 0x70, 0x69, 0x73, 0x2f, 0x61, 0x64, 0x73, 0x2f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x61, 0x64, 0x73, 0x2f, 0x76, 0x31, 0x30, 0x2f, 0x65, 0x6e, 0x75, 0x6d, 0x73, 0x3b, 0x65, 0x6e,
	0x75, 0x6d, 0x73, 0xa2, 0x02, 0x03, 0x47, 0x41, 0x41, 0xaa, 0x02, 0x1e, 0x47, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x2e, 0x41, 0x64, 0x73, 0x2e, 0x47, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x41, 0x64, 0x73,
	0x2e, 0x56, 0x31, 0x30, 0x2e, 0x45, 0x6e, 0x75, 0x6d, 0x73, 0xca, 0x02, 0x1e, 0x47, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x5c, 0x41, 0x64, 0x73, 0x5c, 0x47, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x41, 0x64,
	0x73, 0x5c, 0x56, 0x31, 0x30, 0x5c, 0x45, 0x6e, 0x75, 0x6d, 0x73, 0xea, 0x02, 0x22, 0x47, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x3a, 0x3a, 0x41, 0x64, 0x73, 0x3a, 0x3a, 0x47, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x41, 0x64, 0x73, 0x3a, 0x3a, 0x56, 0x31, 0x30, 0x3a, 0x3a, 0x45, 0x6e, 0x75, 0x6d, 0x73,
	0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_enums_hotel_price_bucket_proto_rawDescOnce sync.Once
	file_enums_hotel_price_bucket_proto_rawDescData = file_enums_hotel_price_bucket_proto_rawDesc
)

func file_enums_hotel_price_bucket_proto_rawDescGZIP() []byte {
	file_enums_hotel_price_bucket_proto_rawDescOnce.Do(func() {
		file_enums_hotel_price_bucket_proto_rawDescData = protoimpl.X.CompressGZIP(file_enums_hotel_price_bucket_proto_rawDescData)
	})
	return file_enums_hotel_price_bucket_proto_rawDescData
}

var file_enums_hotel_price_bucket_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_enums_hotel_price_bucket_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_enums_hotel_price_bucket_proto_goTypes = []interface{}{
	(HotelPriceBucketEnum_HotelPriceBucket)(0), // 0: google.ads.googleads.v10.enums.HotelPriceBucketEnum.HotelPriceBucket
	(*HotelPriceBucketEnum)(nil),               // 1: google.ads.googleads.v10.enums.HotelPriceBucketEnum
}
var file_enums_hotel_price_bucket_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_enums_hotel_price_bucket_proto_init() }
func file_enums_hotel_price_bucket_proto_init() {
	if File_enums_hotel_price_bucket_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_enums_hotel_price_bucket_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*HotelPriceBucketEnum); i {
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
			RawDescriptor: file_enums_hotel_price_bucket_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_enums_hotel_price_bucket_proto_goTypes,
		DependencyIndexes: file_enums_hotel_price_bucket_proto_depIdxs,
		EnumInfos:         file_enums_hotel_price_bucket_proto_enumTypes,
		MessageInfos:      file_enums_hotel_price_bucket_proto_msgTypes,
	}.Build()
	File_enums_hotel_price_bucket_proto = out.File
	file_enums_hotel_price_bucket_proto_rawDesc = nil
	file_enums_hotel_price_bucket_proto_goTypes = nil
	file_enums_hotel_price_bucket_proto_depIdxs = nil
}
