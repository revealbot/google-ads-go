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
// 	protoc        v3.21.12
// source: google/ads/googleads/v13/enums/criterion_type.proto

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

// Enum describing possible criterion types.
type CriterionTypeEnum_CriterionType int32

const (
	// Not specified.
	CriterionTypeEnum_UNSPECIFIED CriterionTypeEnum_CriterionType = 0
	// Used for return value only. Represents value unknown in this version.
	CriterionTypeEnum_UNKNOWN CriterionTypeEnum_CriterionType = 1
	// Keyword, for example, 'mars cruise'.
	CriterionTypeEnum_KEYWORD CriterionTypeEnum_CriterionType = 2
	// Placement, also known as Website, for example, 'www.flowers4sale.com'
	CriterionTypeEnum_PLACEMENT CriterionTypeEnum_CriterionType = 3
	// Mobile application categories to target.
	CriterionTypeEnum_MOBILE_APP_CATEGORY CriterionTypeEnum_CriterionType = 4
	// Mobile applications to target.
	CriterionTypeEnum_MOBILE_APPLICATION CriterionTypeEnum_CriterionType = 5
	// Devices to target.
	CriterionTypeEnum_DEVICE CriterionTypeEnum_CriterionType = 6
	// Locations to target.
	CriterionTypeEnum_LOCATION CriterionTypeEnum_CriterionType = 7
	// Listing groups to target.
	CriterionTypeEnum_LISTING_GROUP CriterionTypeEnum_CriterionType = 8
	// Ad Schedule.
	CriterionTypeEnum_AD_SCHEDULE CriterionTypeEnum_CriterionType = 9
	// Age range.
	CriterionTypeEnum_AGE_RANGE CriterionTypeEnum_CriterionType = 10
	// Gender.
	CriterionTypeEnum_GENDER CriterionTypeEnum_CriterionType = 11
	// Income Range.
	CriterionTypeEnum_INCOME_RANGE CriterionTypeEnum_CriterionType = 12
	// Parental status.
	CriterionTypeEnum_PARENTAL_STATUS CriterionTypeEnum_CriterionType = 13
	// YouTube Video.
	CriterionTypeEnum_YOUTUBE_VIDEO CriterionTypeEnum_CriterionType = 14
	// YouTube Channel.
	CriterionTypeEnum_YOUTUBE_CHANNEL CriterionTypeEnum_CriterionType = 15
	// User list.
	CriterionTypeEnum_USER_LIST CriterionTypeEnum_CriterionType = 16
	// Proximity.
	CriterionTypeEnum_PROXIMITY CriterionTypeEnum_CriterionType = 17
	// A topic target on the display network (for example, "Pets & Animals").
	CriterionTypeEnum_TOPIC CriterionTypeEnum_CriterionType = 18
	// Listing scope to target.
	CriterionTypeEnum_LISTING_SCOPE CriterionTypeEnum_CriterionType = 19
	// Language.
	CriterionTypeEnum_LANGUAGE CriterionTypeEnum_CriterionType = 20
	// IpBlock.
	CriterionTypeEnum_IP_BLOCK CriterionTypeEnum_CriterionType = 21
	// Content Label for category exclusion.
	CriterionTypeEnum_CONTENT_LABEL CriterionTypeEnum_CriterionType = 22
	// Carrier.
	CriterionTypeEnum_CARRIER CriterionTypeEnum_CriterionType = 23
	// A category the user is interested in.
	CriterionTypeEnum_USER_INTEREST CriterionTypeEnum_CriterionType = 24
	// Webpage criterion for dynamic search ads.
	CriterionTypeEnum_WEBPAGE CriterionTypeEnum_CriterionType = 25
	// Operating system version.
	CriterionTypeEnum_OPERATING_SYSTEM_VERSION CriterionTypeEnum_CriterionType = 26
	// App payment model.
	CriterionTypeEnum_APP_PAYMENT_MODEL CriterionTypeEnum_CriterionType = 27
	// Mobile device.
	CriterionTypeEnum_MOBILE_DEVICE CriterionTypeEnum_CriterionType = 28
	// Custom affinity.
	CriterionTypeEnum_CUSTOM_AFFINITY CriterionTypeEnum_CriterionType = 29
	// Custom intent.
	CriterionTypeEnum_CUSTOM_INTENT CriterionTypeEnum_CriterionType = 30
	// Location group.
	CriterionTypeEnum_LOCATION_GROUP CriterionTypeEnum_CriterionType = 31
	// Custom audience
	CriterionTypeEnum_CUSTOM_AUDIENCE CriterionTypeEnum_CriterionType = 32
	// Combined audience
	CriterionTypeEnum_COMBINED_AUDIENCE CriterionTypeEnum_CriterionType = 33
	// Smart Campaign keyword theme
	CriterionTypeEnum_KEYWORD_THEME CriterionTypeEnum_CriterionType = 34
	// Audience
	CriterionTypeEnum_AUDIENCE CriterionTypeEnum_CriterionType = 35
	// Local Services Ads Service ID.
	CriterionTypeEnum_LOCAL_SERVICE_ID CriterionTypeEnum_CriterionType = 37
)

// Enum value maps for CriterionTypeEnum_CriterionType.
var (
	CriterionTypeEnum_CriterionType_name = map[int32]string{
		0:  "UNSPECIFIED",
		1:  "UNKNOWN",
		2:  "KEYWORD",
		3:  "PLACEMENT",
		4:  "MOBILE_APP_CATEGORY",
		5:  "MOBILE_APPLICATION",
		6:  "DEVICE",
		7:  "LOCATION",
		8:  "LISTING_GROUP",
		9:  "AD_SCHEDULE",
		10: "AGE_RANGE",
		11: "GENDER",
		12: "INCOME_RANGE",
		13: "PARENTAL_STATUS",
		14: "YOUTUBE_VIDEO",
		15: "YOUTUBE_CHANNEL",
		16: "USER_LIST",
		17: "PROXIMITY",
		18: "TOPIC",
		19: "LISTING_SCOPE",
		20: "LANGUAGE",
		21: "IP_BLOCK",
		22: "CONTENT_LABEL",
		23: "CARRIER",
		24: "USER_INTEREST",
		25: "WEBPAGE",
		26: "OPERATING_SYSTEM_VERSION",
		27: "APP_PAYMENT_MODEL",
		28: "MOBILE_DEVICE",
		29: "CUSTOM_AFFINITY",
		30: "CUSTOM_INTENT",
		31: "LOCATION_GROUP",
		32: "CUSTOM_AUDIENCE",
		33: "COMBINED_AUDIENCE",
		34: "KEYWORD_THEME",
		35: "AUDIENCE",
		37: "LOCAL_SERVICE_ID",
	}
	CriterionTypeEnum_CriterionType_value = map[string]int32{
		"UNSPECIFIED":              0,
		"UNKNOWN":                  1,
		"KEYWORD":                  2,
		"PLACEMENT":                3,
		"MOBILE_APP_CATEGORY":      4,
		"MOBILE_APPLICATION":       5,
		"DEVICE":                   6,
		"LOCATION":                 7,
		"LISTING_GROUP":            8,
		"AD_SCHEDULE":              9,
		"AGE_RANGE":                10,
		"GENDER":                   11,
		"INCOME_RANGE":             12,
		"PARENTAL_STATUS":          13,
		"YOUTUBE_VIDEO":            14,
		"YOUTUBE_CHANNEL":          15,
		"USER_LIST":                16,
		"PROXIMITY":                17,
		"TOPIC":                    18,
		"LISTING_SCOPE":            19,
		"LANGUAGE":                 20,
		"IP_BLOCK":                 21,
		"CONTENT_LABEL":            22,
		"CARRIER":                  23,
		"USER_INTEREST":            24,
		"WEBPAGE":                  25,
		"OPERATING_SYSTEM_VERSION": 26,
		"APP_PAYMENT_MODEL":        27,
		"MOBILE_DEVICE":            28,
		"CUSTOM_AFFINITY":          29,
		"CUSTOM_INTENT":            30,
		"LOCATION_GROUP":           31,
		"CUSTOM_AUDIENCE":          32,
		"COMBINED_AUDIENCE":        33,
		"KEYWORD_THEME":            34,
		"AUDIENCE":                 35,
		"LOCAL_SERVICE_ID":         37,
	}
)

func (x CriterionTypeEnum_CriterionType) Enum() *CriterionTypeEnum_CriterionType {
	p := new(CriterionTypeEnum_CriterionType)
	*p = x
	return p
}

func (x CriterionTypeEnum_CriterionType) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (CriterionTypeEnum_CriterionType) Descriptor() protoreflect.EnumDescriptor {
	return file_enums_criterion_type_proto_enumTypes[0].Descriptor()
}

func (CriterionTypeEnum_CriterionType) Type() protoreflect.EnumType {
	return &file_enums_criterion_type_proto_enumTypes[0]
}

func (x CriterionTypeEnum_CriterionType) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use CriterionTypeEnum_CriterionType.Descriptor instead.
func (CriterionTypeEnum_CriterionType) EnumDescriptor() ([]byte, []int) {
	return file_enums_criterion_type_proto_rawDescGZIP(), []int{0, 0}
}

// The possible types of a criterion.
type CriterionTypeEnum struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *CriterionTypeEnum) Reset() {
	*x = CriterionTypeEnum{}
	if protoimpl.UnsafeEnabled {
		mi := &file_enums_criterion_type_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CriterionTypeEnum) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CriterionTypeEnum) ProtoMessage() {}

func (x *CriterionTypeEnum) ProtoReflect() protoreflect.Message {
	mi := &file_enums_criterion_type_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CriterionTypeEnum.ProtoReflect.Descriptor instead.
func (*CriterionTypeEnum) Descriptor() ([]byte, []int) {
	return file_enums_criterion_type_proto_rawDescGZIP(), []int{0}
}

var File_enums_criterion_type_proto protoreflect.FileDescriptor

var file_enums_criterion_type_proto_rawDesc = []byte{
	0x0a, 0x33, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x64, 0x73, 0x2f, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x61, 0x64, 0x73, 0x2f, 0x76, 0x31, 0x33, 0x2f, 0x65, 0x6e, 0x75, 0x6d, 0x73,
	0x2f, 0x63, 0x72, 0x69, 0x74, 0x65, 0x72, 0x69, 0x6f, 0x6e, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x1e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x61, 0x64,
	0x73, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61, 0x64, 0x73, 0x2e, 0x76, 0x31, 0x33, 0x2e,
	0x65, 0x6e, 0x75, 0x6d, 0x73, 0x22, 0xb7, 0x05, 0x0a, 0x11, 0x43, 0x72, 0x69, 0x74, 0x65, 0x72,
	0x69, 0x6f, 0x6e, 0x54, 0x79, 0x70, 0x65, 0x45, 0x6e, 0x75, 0x6d, 0x22, 0xa1, 0x05, 0x0a, 0x0d,
	0x43, 0x72, 0x69, 0x74, 0x65, 0x72, 0x69, 0x6f, 0x6e, 0x54, 0x79, 0x70, 0x65, 0x12, 0x0f, 0x0a,
	0x0b, 0x55, 0x4e, 0x53, 0x50, 0x45, 0x43, 0x49, 0x46, 0x49, 0x45, 0x44, 0x10, 0x00, 0x12, 0x0b,
	0x0a, 0x07, 0x55, 0x4e, 0x4b, 0x4e, 0x4f, 0x57, 0x4e, 0x10, 0x01, 0x12, 0x0b, 0x0a, 0x07, 0x4b,
	0x45, 0x59, 0x57, 0x4f, 0x52, 0x44, 0x10, 0x02, 0x12, 0x0d, 0x0a, 0x09, 0x50, 0x4c, 0x41, 0x43,
	0x45, 0x4d, 0x45, 0x4e, 0x54, 0x10, 0x03, 0x12, 0x17, 0x0a, 0x13, 0x4d, 0x4f, 0x42, 0x49, 0x4c,
	0x45, 0x5f, 0x41, 0x50, 0x50, 0x5f, 0x43, 0x41, 0x54, 0x45, 0x47, 0x4f, 0x52, 0x59, 0x10, 0x04,
	0x12, 0x16, 0x0a, 0x12, 0x4d, 0x4f, 0x42, 0x49, 0x4c, 0x45, 0x5f, 0x41, 0x50, 0x50, 0x4c, 0x49,
	0x43, 0x41, 0x54, 0x49, 0x4f, 0x4e, 0x10, 0x05, 0x12, 0x0a, 0x0a, 0x06, 0x44, 0x45, 0x56, 0x49,
	0x43, 0x45, 0x10, 0x06, 0x12, 0x0c, 0x0a, 0x08, 0x4c, 0x4f, 0x43, 0x41, 0x54, 0x49, 0x4f, 0x4e,
	0x10, 0x07, 0x12, 0x11, 0x0a, 0x0d, 0x4c, 0x49, 0x53, 0x54, 0x49, 0x4e, 0x47, 0x5f, 0x47, 0x52,
	0x4f, 0x55, 0x50, 0x10, 0x08, 0x12, 0x0f, 0x0a, 0x0b, 0x41, 0x44, 0x5f, 0x53, 0x43, 0x48, 0x45,
	0x44, 0x55, 0x4c, 0x45, 0x10, 0x09, 0x12, 0x0d, 0x0a, 0x09, 0x41, 0x47, 0x45, 0x5f, 0x52, 0x41,
	0x4e, 0x47, 0x45, 0x10, 0x0a, 0x12, 0x0a, 0x0a, 0x06, 0x47, 0x45, 0x4e, 0x44, 0x45, 0x52, 0x10,
	0x0b, 0x12, 0x10, 0x0a, 0x0c, 0x49, 0x4e, 0x43, 0x4f, 0x4d, 0x45, 0x5f, 0x52, 0x41, 0x4e, 0x47,
	0x45, 0x10, 0x0c, 0x12, 0x13, 0x0a, 0x0f, 0x50, 0x41, 0x52, 0x45, 0x4e, 0x54, 0x41, 0x4c, 0x5f,
	0x53, 0x54, 0x41, 0x54, 0x55, 0x53, 0x10, 0x0d, 0x12, 0x11, 0x0a, 0x0d, 0x59, 0x4f, 0x55, 0x54,
	0x55, 0x42, 0x45, 0x5f, 0x56, 0x49, 0x44, 0x45, 0x4f, 0x10, 0x0e, 0x12, 0x13, 0x0a, 0x0f, 0x59,
	0x4f, 0x55, 0x54, 0x55, 0x42, 0x45, 0x5f, 0x43, 0x48, 0x41, 0x4e, 0x4e, 0x45, 0x4c, 0x10, 0x0f,
	0x12, 0x0d, 0x0a, 0x09, 0x55, 0x53, 0x45, 0x52, 0x5f, 0x4c, 0x49, 0x53, 0x54, 0x10, 0x10, 0x12,
	0x0d, 0x0a, 0x09, 0x50, 0x52, 0x4f, 0x58, 0x49, 0x4d, 0x49, 0x54, 0x59, 0x10, 0x11, 0x12, 0x09,
	0x0a, 0x05, 0x54, 0x4f, 0x50, 0x49, 0x43, 0x10, 0x12, 0x12, 0x11, 0x0a, 0x0d, 0x4c, 0x49, 0x53,
	0x54, 0x49, 0x4e, 0x47, 0x5f, 0x53, 0x43, 0x4f, 0x50, 0x45, 0x10, 0x13, 0x12, 0x0c, 0x0a, 0x08,
	0x4c, 0x41, 0x4e, 0x47, 0x55, 0x41, 0x47, 0x45, 0x10, 0x14, 0x12, 0x0c, 0x0a, 0x08, 0x49, 0x50,
	0x5f, 0x42, 0x4c, 0x4f, 0x43, 0x4b, 0x10, 0x15, 0x12, 0x11, 0x0a, 0x0d, 0x43, 0x4f, 0x4e, 0x54,
	0x45, 0x4e, 0x54, 0x5f, 0x4c, 0x41, 0x42, 0x45, 0x4c, 0x10, 0x16, 0x12, 0x0b, 0x0a, 0x07, 0x43,
	0x41, 0x52, 0x52, 0x49, 0x45, 0x52, 0x10, 0x17, 0x12, 0x11, 0x0a, 0x0d, 0x55, 0x53, 0x45, 0x52,
	0x5f, 0x49, 0x4e, 0x54, 0x45, 0x52, 0x45, 0x53, 0x54, 0x10, 0x18, 0x12, 0x0b, 0x0a, 0x07, 0x57,
	0x45, 0x42, 0x50, 0x41, 0x47, 0x45, 0x10, 0x19, 0x12, 0x1c, 0x0a, 0x18, 0x4f, 0x50, 0x45, 0x52,
	0x41, 0x54, 0x49, 0x4e, 0x47, 0x5f, 0x53, 0x59, 0x53, 0x54, 0x45, 0x4d, 0x5f, 0x56, 0x45, 0x52,
	0x53, 0x49, 0x4f, 0x4e, 0x10, 0x1a, 0x12, 0x15, 0x0a, 0x11, 0x41, 0x50, 0x50, 0x5f, 0x50, 0x41,
	0x59, 0x4d, 0x45, 0x4e, 0x54, 0x5f, 0x4d, 0x4f, 0x44, 0x45, 0x4c, 0x10, 0x1b, 0x12, 0x11, 0x0a,
	0x0d, 0x4d, 0x4f, 0x42, 0x49, 0x4c, 0x45, 0x5f, 0x44, 0x45, 0x56, 0x49, 0x43, 0x45, 0x10, 0x1c,
	0x12, 0x13, 0x0a, 0x0f, 0x43, 0x55, 0x53, 0x54, 0x4f, 0x4d, 0x5f, 0x41, 0x46, 0x46, 0x49, 0x4e,
	0x49, 0x54, 0x59, 0x10, 0x1d, 0x12, 0x11, 0x0a, 0x0d, 0x43, 0x55, 0x53, 0x54, 0x4f, 0x4d, 0x5f,
	0x49, 0x4e, 0x54, 0x45, 0x4e, 0x54, 0x10, 0x1e, 0x12, 0x12, 0x0a, 0x0e, 0x4c, 0x4f, 0x43, 0x41,
	0x54, 0x49, 0x4f, 0x4e, 0x5f, 0x47, 0x52, 0x4f, 0x55, 0x50, 0x10, 0x1f, 0x12, 0x13, 0x0a, 0x0f,
	0x43, 0x55, 0x53, 0x54, 0x4f, 0x4d, 0x5f, 0x41, 0x55, 0x44, 0x49, 0x45, 0x4e, 0x43, 0x45, 0x10,
	0x20, 0x12, 0x15, 0x0a, 0x11, 0x43, 0x4f, 0x4d, 0x42, 0x49, 0x4e, 0x45, 0x44, 0x5f, 0x41, 0x55,
	0x44, 0x49, 0x45, 0x4e, 0x43, 0x45, 0x10, 0x21, 0x12, 0x11, 0x0a, 0x0d, 0x4b, 0x45, 0x59, 0x57,
	0x4f, 0x52, 0x44, 0x5f, 0x54, 0x48, 0x45, 0x4d, 0x45, 0x10, 0x22, 0x12, 0x0c, 0x0a, 0x08, 0x41,
	0x55, 0x44, 0x49, 0x45, 0x4e, 0x43, 0x45, 0x10, 0x23, 0x12, 0x14, 0x0a, 0x10, 0x4c, 0x4f, 0x43,
	0x41, 0x4c, 0x5f, 0x53, 0x45, 0x52, 0x56, 0x49, 0x43, 0x45, 0x5f, 0x49, 0x44, 0x10, 0x25, 0x42,
	0xec, 0x01, 0x0a, 0x22, 0x63, 0x6f, 0x6d, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x61,
	0x64, 0x73, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61, 0x64, 0x73, 0x2e, 0x76, 0x31, 0x33,
	0x2e, 0x65, 0x6e, 0x75, 0x6d, 0x73, 0x42, 0x12, 0x43, 0x72, 0x69, 0x74, 0x65, 0x72, 0x69, 0x6f,
	0x6e, 0x54, 0x79, 0x70, 0x65, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a, 0x43, 0x67, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x67, 0x6f, 0x6c, 0x61, 0x6e, 0x67, 0x2e, 0x6f, 0x72, 0x67, 0x2f,
	0x67, 0x65, 0x6e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61,
	0x70, 0x69, 0x73, 0x2f, 0x61, 0x64, 0x73, 0x2f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61, 0x64,
	0x73, 0x2f, 0x76, 0x31, 0x33, 0x2f, 0x65, 0x6e, 0x75, 0x6d, 0x73, 0x3b, 0x65, 0x6e, 0x75, 0x6d,
	0x73, 0xa2, 0x02, 0x03, 0x47, 0x41, 0x41, 0xaa, 0x02, 0x1e, 0x47, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x2e, 0x41, 0x64, 0x73, 0x2e, 0x47, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x41, 0x64, 0x73, 0x2e, 0x56,
	0x31, 0x33, 0x2e, 0x45, 0x6e, 0x75, 0x6d, 0x73, 0xca, 0x02, 0x1e, 0x47, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x5c, 0x41, 0x64, 0x73, 0x5c, 0x47, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x41, 0x64, 0x73, 0x5c,
	0x56, 0x31, 0x33, 0x5c, 0x45, 0x6e, 0x75, 0x6d, 0x73, 0xea, 0x02, 0x22, 0x47, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x3a, 0x3a, 0x41, 0x64, 0x73, 0x3a, 0x3a, 0x47, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x41,
	0x64, 0x73, 0x3a, 0x3a, 0x56, 0x31, 0x33, 0x3a, 0x3a, 0x45, 0x6e, 0x75, 0x6d, 0x73, 0x62, 0x06,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_enums_criterion_type_proto_rawDescOnce sync.Once
	file_enums_criterion_type_proto_rawDescData = file_enums_criterion_type_proto_rawDesc
)

func file_enums_criterion_type_proto_rawDescGZIP() []byte {
	file_enums_criterion_type_proto_rawDescOnce.Do(func() {
		file_enums_criterion_type_proto_rawDescData = protoimpl.X.CompressGZIP(file_enums_criterion_type_proto_rawDescData)
	})
	return file_enums_criterion_type_proto_rawDescData
}

var file_enums_criterion_type_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_enums_criterion_type_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_enums_criterion_type_proto_goTypes = []interface{}{
	(CriterionTypeEnum_CriterionType)(0), // 0: google.ads.googleads.v13.enums.CriterionTypeEnum.CriterionType
	(*CriterionTypeEnum)(nil),            // 1: google.ads.googleads.v13.enums.CriterionTypeEnum
}
var file_enums_criterion_type_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_enums_criterion_type_proto_init() }
func file_enums_criterion_type_proto_init() {
	if File_enums_criterion_type_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_enums_criterion_type_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CriterionTypeEnum); i {
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
			RawDescriptor: file_enums_criterion_type_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_enums_criterion_type_proto_goTypes,
		DependencyIndexes: file_enums_criterion_type_proto_depIdxs,
		EnumInfos:         file_enums_criterion_type_proto_enumTypes,
		MessageInfos:      file_enums_criterion_type_proto_msgTypes,
	}.Build()
	File_enums_criterion_type_proto = out.File
	file_enums_criterion_type_proto_rawDesc = nil
	file_enums_criterion_type_proto_goTypes = nil
	file_enums_criterion_type_proto_depIdxs = nil
}
