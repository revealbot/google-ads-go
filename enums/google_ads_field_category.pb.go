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
// source: google/ads/googleads/v10/enums/google_ads_field_category.proto

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

// The category of the artifact.
type GoogleAdsFieldCategoryEnum_GoogleAdsFieldCategory int32

const (
	// Unspecified
	GoogleAdsFieldCategoryEnum_UNSPECIFIED GoogleAdsFieldCategoryEnum_GoogleAdsFieldCategory = 0
	// Unknown
	GoogleAdsFieldCategoryEnum_UNKNOWN GoogleAdsFieldCategoryEnum_GoogleAdsFieldCategory = 1
	// The described artifact is a resource.
	GoogleAdsFieldCategoryEnum_RESOURCE GoogleAdsFieldCategoryEnum_GoogleAdsFieldCategory = 2
	// The described artifact is a field and is an attribute of a resource.
	// Including a resource attribute field in a query may segment the query if
	// the resource to which it is attributed segments the resource found in
	// the FROM clause.
	GoogleAdsFieldCategoryEnum_ATTRIBUTE GoogleAdsFieldCategoryEnum_GoogleAdsFieldCategory = 3
	// The described artifact is a field and always segments search queries.
	GoogleAdsFieldCategoryEnum_SEGMENT GoogleAdsFieldCategoryEnum_GoogleAdsFieldCategory = 5
	// The described artifact is a field and is a metric. It never segments
	// search queries.
	GoogleAdsFieldCategoryEnum_METRIC GoogleAdsFieldCategoryEnum_GoogleAdsFieldCategory = 6
)

// Enum value maps for GoogleAdsFieldCategoryEnum_GoogleAdsFieldCategory.
var (
	GoogleAdsFieldCategoryEnum_GoogleAdsFieldCategory_name = map[int32]string{
		0: "UNSPECIFIED",
		1: "UNKNOWN",
		2: "RESOURCE",
		3: "ATTRIBUTE",
		5: "SEGMENT",
		6: "METRIC",
	}
	GoogleAdsFieldCategoryEnum_GoogleAdsFieldCategory_value = map[string]int32{
		"UNSPECIFIED": 0,
		"UNKNOWN":     1,
		"RESOURCE":    2,
		"ATTRIBUTE":   3,
		"SEGMENT":     5,
		"METRIC":      6,
	}
)

func (x GoogleAdsFieldCategoryEnum_GoogleAdsFieldCategory) Enum() *GoogleAdsFieldCategoryEnum_GoogleAdsFieldCategory {
	p := new(GoogleAdsFieldCategoryEnum_GoogleAdsFieldCategory)
	*p = x
	return p
}

func (x GoogleAdsFieldCategoryEnum_GoogleAdsFieldCategory) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (GoogleAdsFieldCategoryEnum_GoogleAdsFieldCategory) Descriptor() protoreflect.EnumDescriptor {
	return file_enums_google_ads_field_category_proto_enumTypes[0].Descriptor()
}

func (GoogleAdsFieldCategoryEnum_GoogleAdsFieldCategory) Type() protoreflect.EnumType {
	return &file_enums_google_ads_field_category_proto_enumTypes[0]
}

func (x GoogleAdsFieldCategoryEnum_GoogleAdsFieldCategory) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use GoogleAdsFieldCategoryEnum_GoogleAdsFieldCategory.Descriptor instead.
func (GoogleAdsFieldCategoryEnum_GoogleAdsFieldCategory) EnumDescriptor() ([]byte, []int) {
	return file_enums_google_ads_field_category_proto_rawDescGZIP(), []int{0, 0}
}

// Container for enum that determines if the described artifact is a resource
// or a field, and if it is a field, when it segments search queries.
type GoogleAdsFieldCategoryEnum struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *GoogleAdsFieldCategoryEnum) Reset() {
	*x = GoogleAdsFieldCategoryEnum{}
	if protoimpl.UnsafeEnabled {
		mi := &file_enums_google_ads_field_category_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GoogleAdsFieldCategoryEnum) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GoogleAdsFieldCategoryEnum) ProtoMessage() {}

func (x *GoogleAdsFieldCategoryEnum) ProtoReflect() protoreflect.Message {
	mi := &file_enums_google_ads_field_category_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GoogleAdsFieldCategoryEnum.ProtoReflect.Descriptor instead.
func (*GoogleAdsFieldCategoryEnum) Descriptor() ([]byte, []int) {
	return file_enums_google_ads_field_category_proto_rawDescGZIP(), []int{0}
}

var File_enums_google_ads_field_category_proto protoreflect.FileDescriptor

var file_enums_google_ads_field_category_proto_rawDesc = []byte{
	0x0a, 0x3e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x64, 0x73, 0x2f, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x61, 0x64, 0x73, 0x2f, 0x76, 0x31, 0x30, 0x2f, 0x65, 0x6e, 0x75, 0x6d, 0x73,
	0x2f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x5f, 0x61, 0x64, 0x73, 0x5f, 0x66, 0x69, 0x65, 0x6c,
	0x64, 0x5f, 0x63, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x12, 0x1e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x61, 0x64, 0x73, 0x2e, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x61, 0x64, 0x73, 0x2e, 0x76, 0x31, 0x30, 0x2e, 0x65, 0x6e, 0x75, 0x6d, 0x73,
	0x1a, 0x1c, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x61, 0x6e, 0x6e,
	0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x8a,
	0x01, 0x0a, 0x1a, 0x47, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x41, 0x64, 0x73, 0x46, 0x69, 0x65, 0x6c,
	0x64, 0x43, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x79, 0x45, 0x6e, 0x75, 0x6d, 0x22, 0x6c, 0x0a,
	0x16, 0x47, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x41, 0x64, 0x73, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x43,
	0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x79, 0x12, 0x0f, 0x0a, 0x0b, 0x55, 0x4e, 0x53, 0x50, 0x45,
	0x43, 0x49, 0x46, 0x49, 0x45, 0x44, 0x10, 0x00, 0x12, 0x0b, 0x0a, 0x07, 0x55, 0x4e, 0x4b, 0x4e,
	0x4f, 0x57, 0x4e, 0x10, 0x01, 0x12, 0x0c, 0x0a, 0x08, 0x52, 0x45, 0x53, 0x4f, 0x55, 0x52, 0x43,
	0x45, 0x10, 0x02, 0x12, 0x0d, 0x0a, 0x09, 0x41, 0x54, 0x54, 0x52, 0x49, 0x42, 0x55, 0x54, 0x45,
	0x10, 0x03, 0x12, 0x0b, 0x0a, 0x07, 0x53, 0x45, 0x47, 0x4d, 0x45, 0x4e, 0x54, 0x10, 0x05, 0x12,
	0x0a, 0x0a, 0x06, 0x4d, 0x45, 0x54, 0x52, 0x49, 0x43, 0x10, 0x06, 0x42, 0xf5, 0x01, 0x0a, 0x22,
	0x63, 0x6f, 0x6d, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x61, 0x64, 0x73, 0x2e, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61, 0x64, 0x73, 0x2e, 0x76, 0x31, 0x30, 0x2e, 0x65, 0x6e, 0x75,
	0x6d, 0x73, 0x42, 0x1b, 0x47, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x41, 0x64, 0x73, 0x46, 0x69, 0x65,
	0x6c, 0x64, 0x43, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x79, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50,
	0x01, 0x5a, 0x43, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x67, 0x6f, 0x6c, 0x61, 0x6e, 0x67,
	0x2e, 0x6f, 0x72, 0x67, 0x2f, 0x67, 0x65, 0x6e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x67, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x61, 0x70, 0x69, 0x73, 0x2f, 0x61, 0x64, 0x73, 0x2f, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x61, 0x64, 0x73, 0x2f, 0x76, 0x31, 0x30, 0x2f, 0x65, 0x6e, 0x75, 0x6d, 0x73,
	0x3b, 0x65, 0x6e, 0x75, 0x6d, 0x73, 0xa2, 0x02, 0x03, 0x47, 0x41, 0x41, 0xaa, 0x02, 0x1e, 0x47,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x41, 0x64, 0x73, 0x2e, 0x47, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x41, 0x64, 0x73, 0x2e, 0x56, 0x31, 0x30, 0x2e, 0x45, 0x6e, 0x75, 0x6d, 0x73, 0xca, 0x02, 0x1e,
	0x47, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x5c, 0x41, 0x64, 0x73, 0x5c, 0x47, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x41, 0x64, 0x73, 0x5c, 0x56, 0x31, 0x30, 0x5c, 0x45, 0x6e, 0x75, 0x6d, 0x73, 0xea, 0x02,
	0x22, 0x47, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x3a, 0x3a, 0x41, 0x64, 0x73, 0x3a, 0x3a, 0x47, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x41, 0x64, 0x73, 0x3a, 0x3a, 0x56, 0x31, 0x30, 0x3a, 0x3a, 0x45, 0x6e,
	0x75, 0x6d, 0x73, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_enums_google_ads_field_category_proto_rawDescOnce sync.Once
	file_enums_google_ads_field_category_proto_rawDescData = file_enums_google_ads_field_category_proto_rawDesc
)

func file_enums_google_ads_field_category_proto_rawDescGZIP() []byte {
	file_enums_google_ads_field_category_proto_rawDescOnce.Do(func() {
		file_enums_google_ads_field_category_proto_rawDescData = protoimpl.X.CompressGZIP(file_enums_google_ads_field_category_proto_rawDescData)
	})
	return file_enums_google_ads_field_category_proto_rawDescData
}

var file_enums_google_ads_field_category_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_enums_google_ads_field_category_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_enums_google_ads_field_category_proto_goTypes = []interface{}{
	(GoogleAdsFieldCategoryEnum_GoogleAdsFieldCategory)(0), // 0: google.ads.googleads.v10.enums.GoogleAdsFieldCategoryEnum.GoogleAdsFieldCategory
	(*GoogleAdsFieldCategoryEnum)(nil),                     // 1: google.ads.googleads.v10.enums.GoogleAdsFieldCategoryEnum
}
var file_enums_google_ads_field_category_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_enums_google_ads_field_category_proto_init() }
func file_enums_google_ads_field_category_proto_init() {
	if File_enums_google_ads_field_category_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_enums_google_ads_field_category_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GoogleAdsFieldCategoryEnum); i {
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
			RawDescriptor: file_enums_google_ads_field_category_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_enums_google_ads_field_category_proto_goTypes,
		DependencyIndexes: file_enums_google_ads_field_category_proto_depIdxs,
		EnumInfos:         file_enums_google_ads_field_category_proto_enumTypes,
		MessageInfos:      file_enums_google_ads_field_category_proto_msgTypes,
	}.Build()
	File_enums_google_ads_field_category_proto = out.File
	file_enums_google_ads_field_category_proto_rawDesc = nil
	file_enums_google_ads_field_category_proto_goTypes = nil
	file_enums_google_ads_field_category_proto_depIdxs = nil
}
