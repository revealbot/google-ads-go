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
// source: google/ads/googleads/v19/enums/ad_type.proto

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

// The possible types of an ad.
type AdTypeEnum_AdType int32

const (
	// No value has been specified.
	AdTypeEnum_UNSPECIFIED AdTypeEnum_AdType = 0
	// The received value is not known in this version.
	//
	// This is a response-only value.
	AdTypeEnum_UNKNOWN AdTypeEnum_AdType = 1
	// The ad is a text ad.
	AdTypeEnum_TEXT_AD AdTypeEnum_AdType = 2
	// The ad is an expanded text ad.
	AdTypeEnum_EXPANDED_TEXT_AD AdTypeEnum_AdType = 3
	// The ad is an expanded dynamic search ad.
	AdTypeEnum_EXPANDED_DYNAMIC_SEARCH_AD AdTypeEnum_AdType = 7
	// The ad is a hotel ad.
	AdTypeEnum_HOTEL_AD AdTypeEnum_AdType = 8
	// The ad is a Smart Shopping ad.
	AdTypeEnum_SHOPPING_SMART_AD AdTypeEnum_AdType = 9
	// The ad is a standard Shopping ad.
	AdTypeEnum_SHOPPING_PRODUCT_AD AdTypeEnum_AdType = 10
	// The ad is a video ad.
	AdTypeEnum_VIDEO_AD AdTypeEnum_AdType = 12
	// This ad is an Image ad.
	AdTypeEnum_IMAGE_AD AdTypeEnum_AdType = 14
	// The ad is a responsive search ad.
	AdTypeEnum_RESPONSIVE_SEARCH_AD AdTypeEnum_AdType = 15
	// The ad is a legacy responsive display ad.
	AdTypeEnum_LEGACY_RESPONSIVE_DISPLAY_AD AdTypeEnum_AdType = 16
	// The ad is an app ad.
	AdTypeEnum_APP_AD AdTypeEnum_AdType = 17
	// The ad is a legacy app install ad.
	AdTypeEnum_LEGACY_APP_INSTALL_AD AdTypeEnum_AdType = 18
	// The ad is a responsive display ad.
	AdTypeEnum_RESPONSIVE_DISPLAY_AD AdTypeEnum_AdType = 19
	// The ad is a local ad.
	AdTypeEnum_LOCAL_AD AdTypeEnum_AdType = 20
	// The ad is a display upload ad with the HTML5_UPLOAD_AD product type.
	AdTypeEnum_HTML5_UPLOAD_AD AdTypeEnum_AdType = 21
	// The ad is a display upload ad with one of the DYNAMIC_HTML5_* product
	// types.
	AdTypeEnum_DYNAMIC_HTML5_AD AdTypeEnum_AdType = 22
	// The ad is an app engagement ad.
	AdTypeEnum_APP_ENGAGEMENT_AD AdTypeEnum_AdType = 23
	// The ad is a Shopping Comparison Listing ad.
	AdTypeEnum_SHOPPING_COMPARISON_LISTING_AD AdTypeEnum_AdType = 24
	// Video bumper ad.
	AdTypeEnum_VIDEO_BUMPER_AD AdTypeEnum_AdType = 25
	// Video non-skippable in-stream ad.
	AdTypeEnum_VIDEO_NON_SKIPPABLE_IN_STREAM_AD AdTypeEnum_AdType = 26
	// Video TrueView in-stream ad.
	AdTypeEnum_VIDEO_TRUEVIEW_IN_STREAM_AD AdTypeEnum_AdType = 29
	// Video responsive ad.
	AdTypeEnum_VIDEO_RESPONSIVE_AD AdTypeEnum_AdType = 30
	// Smart campaign ad.
	AdTypeEnum_SMART_CAMPAIGN_AD AdTypeEnum_AdType = 31
	// Call ad.
	AdTypeEnum_CALL_AD AdTypeEnum_AdType = 32
	// Universal app pre-registration ad.
	AdTypeEnum_APP_PRE_REGISTRATION_AD AdTypeEnum_AdType = 33
	// In-feed video ad.
	AdTypeEnum_IN_FEED_VIDEO_AD AdTypeEnum_AdType = 34
	// Demand Gen multi asset ad.
	AdTypeEnum_DEMAND_GEN_MULTI_ASSET_AD AdTypeEnum_AdType = 40
	// Demand Gen carousel ad.
	AdTypeEnum_DEMAND_GEN_CAROUSEL_AD AdTypeEnum_AdType = 41
	// Travel ad.
	AdTypeEnum_TRAVEL_AD AdTypeEnum_AdType = 37
	// Demand Gen video responsive ad.
	AdTypeEnum_DEMAND_GEN_VIDEO_RESPONSIVE_AD AdTypeEnum_AdType = 42
	// Demand Gen product ad.
	AdTypeEnum_DEMAND_GEN_PRODUCT_AD AdTypeEnum_AdType = 39
	// YouTube Audio ad.
	AdTypeEnum_YOUTUBE_AUDIO_AD AdTypeEnum_AdType = 44
)

// Enum value maps for AdTypeEnum_AdType.
var (
	AdTypeEnum_AdType_name = map[int32]string{
		0:  "UNSPECIFIED",
		1:  "UNKNOWN",
		2:  "TEXT_AD",
		3:  "EXPANDED_TEXT_AD",
		7:  "EXPANDED_DYNAMIC_SEARCH_AD",
		8:  "HOTEL_AD",
		9:  "SHOPPING_SMART_AD",
		10: "SHOPPING_PRODUCT_AD",
		12: "VIDEO_AD",
		14: "IMAGE_AD",
		15: "RESPONSIVE_SEARCH_AD",
		16: "LEGACY_RESPONSIVE_DISPLAY_AD",
		17: "APP_AD",
		18: "LEGACY_APP_INSTALL_AD",
		19: "RESPONSIVE_DISPLAY_AD",
		20: "LOCAL_AD",
		21: "HTML5_UPLOAD_AD",
		22: "DYNAMIC_HTML5_AD",
		23: "APP_ENGAGEMENT_AD",
		24: "SHOPPING_COMPARISON_LISTING_AD",
		25: "VIDEO_BUMPER_AD",
		26: "VIDEO_NON_SKIPPABLE_IN_STREAM_AD",
		29: "VIDEO_TRUEVIEW_IN_STREAM_AD",
		30: "VIDEO_RESPONSIVE_AD",
		31: "SMART_CAMPAIGN_AD",
		32: "CALL_AD",
		33: "APP_PRE_REGISTRATION_AD",
		34: "IN_FEED_VIDEO_AD",
		40: "DEMAND_GEN_MULTI_ASSET_AD",
		41: "DEMAND_GEN_CAROUSEL_AD",
		37: "TRAVEL_AD",
		42: "DEMAND_GEN_VIDEO_RESPONSIVE_AD",
		39: "DEMAND_GEN_PRODUCT_AD",
		44: "YOUTUBE_AUDIO_AD",
	}
	AdTypeEnum_AdType_value = map[string]int32{
		"UNSPECIFIED":                      0,
		"UNKNOWN":                          1,
		"TEXT_AD":                          2,
		"EXPANDED_TEXT_AD":                 3,
		"EXPANDED_DYNAMIC_SEARCH_AD":       7,
		"HOTEL_AD":                         8,
		"SHOPPING_SMART_AD":                9,
		"SHOPPING_PRODUCT_AD":              10,
		"VIDEO_AD":                         12,
		"IMAGE_AD":                         14,
		"RESPONSIVE_SEARCH_AD":             15,
		"LEGACY_RESPONSIVE_DISPLAY_AD":     16,
		"APP_AD":                           17,
		"LEGACY_APP_INSTALL_AD":            18,
		"RESPONSIVE_DISPLAY_AD":            19,
		"LOCAL_AD":                         20,
		"HTML5_UPLOAD_AD":                  21,
		"DYNAMIC_HTML5_AD":                 22,
		"APP_ENGAGEMENT_AD":                23,
		"SHOPPING_COMPARISON_LISTING_AD":   24,
		"VIDEO_BUMPER_AD":                  25,
		"VIDEO_NON_SKIPPABLE_IN_STREAM_AD": 26,
		"VIDEO_TRUEVIEW_IN_STREAM_AD":      29,
		"VIDEO_RESPONSIVE_AD":              30,
		"SMART_CAMPAIGN_AD":                31,
		"CALL_AD":                          32,
		"APP_PRE_REGISTRATION_AD":          33,
		"IN_FEED_VIDEO_AD":                 34,
		"DEMAND_GEN_MULTI_ASSET_AD":        40,
		"DEMAND_GEN_CAROUSEL_AD":           41,
		"TRAVEL_AD":                        37,
		"DEMAND_GEN_VIDEO_RESPONSIVE_AD":   42,
		"DEMAND_GEN_PRODUCT_AD":            39,
		"YOUTUBE_AUDIO_AD":                 44,
	}
)

func (x AdTypeEnum_AdType) Enum() *AdTypeEnum_AdType {
	p := new(AdTypeEnum_AdType)
	*p = x
	return p
}

func (x AdTypeEnum_AdType) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (AdTypeEnum_AdType) Descriptor() protoreflect.EnumDescriptor {
	return file_enums_ad_type_proto_enumTypes[0].Descriptor()
}

func (AdTypeEnum_AdType) Type() protoreflect.EnumType {
	return &file_enums_ad_type_proto_enumTypes[0]
}

func (x AdTypeEnum_AdType) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use AdTypeEnum_AdType.Descriptor instead.
func (AdTypeEnum_AdType) EnumDescriptor() ([]byte, []int) {
	return file_enums_ad_type_proto_rawDescGZIP(), []int{0, 0}
}

// Container for enum describing possible types of an ad.
type AdTypeEnum struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *AdTypeEnum) Reset() {
	*x = AdTypeEnum{}
	mi := &file_enums_ad_type_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *AdTypeEnum) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AdTypeEnum) ProtoMessage() {}

func (x *AdTypeEnum) ProtoReflect() protoreflect.Message {
	mi := &file_enums_ad_type_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AdTypeEnum.ProtoReflect.Descriptor instead.
func (*AdTypeEnum) Descriptor() ([]byte, []int) {
	return file_enums_ad_type_proto_rawDescGZIP(), []int{0}
}

var File_enums_ad_type_proto protoreflect.FileDescriptor

var file_enums_ad_type_proto_rawDesc = string([]byte{
	0x0a, 0x2c, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x64, 0x73, 0x2f, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x61, 0x64, 0x73, 0x2f, 0x76, 0x31, 0x39, 0x2f, 0x65, 0x6e, 0x75, 0x6d, 0x73,
	0x2f, 0x61, 0x64, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x1e,
	0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x61, 0x64, 0x73, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x61, 0x64, 0x73, 0x2e, 0x76, 0x31, 0x39, 0x2e, 0x65, 0x6e, 0x75, 0x6d, 0x73, 0x22, 0xaf,
	0x06, 0x0a, 0x0a, 0x41, 0x64, 0x54, 0x79, 0x70, 0x65, 0x45, 0x6e, 0x75, 0x6d, 0x22, 0xa0, 0x06,
	0x0a, 0x06, 0x41, 0x64, 0x54, 0x79, 0x70, 0x65, 0x12, 0x0f, 0x0a, 0x0b, 0x55, 0x4e, 0x53, 0x50,
	0x45, 0x43, 0x49, 0x46, 0x49, 0x45, 0x44, 0x10, 0x00, 0x12, 0x0b, 0x0a, 0x07, 0x55, 0x4e, 0x4b,
	0x4e, 0x4f, 0x57, 0x4e, 0x10, 0x01, 0x12, 0x0b, 0x0a, 0x07, 0x54, 0x45, 0x58, 0x54, 0x5f, 0x41,
	0x44, 0x10, 0x02, 0x12, 0x14, 0x0a, 0x10, 0x45, 0x58, 0x50, 0x41, 0x4e, 0x44, 0x45, 0x44, 0x5f,
	0x54, 0x45, 0x58, 0x54, 0x5f, 0x41, 0x44, 0x10, 0x03, 0x12, 0x1e, 0x0a, 0x1a, 0x45, 0x58, 0x50,
	0x41, 0x4e, 0x44, 0x45, 0x44, 0x5f, 0x44, 0x59, 0x4e, 0x41, 0x4d, 0x49, 0x43, 0x5f, 0x53, 0x45,
	0x41, 0x52, 0x43, 0x48, 0x5f, 0x41, 0x44, 0x10, 0x07, 0x12, 0x0c, 0x0a, 0x08, 0x48, 0x4f, 0x54,
	0x45, 0x4c, 0x5f, 0x41, 0x44, 0x10, 0x08, 0x12, 0x15, 0x0a, 0x11, 0x53, 0x48, 0x4f, 0x50, 0x50,
	0x49, 0x4e, 0x47, 0x5f, 0x53, 0x4d, 0x41, 0x52, 0x54, 0x5f, 0x41, 0x44, 0x10, 0x09, 0x12, 0x17,
	0x0a, 0x13, 0x53, 0x48, 0x4f, 0x50, 0x50, 0x49, 0x4e, 0x47, 0x5f, 0x50, 0x52, 0x4f, 0x44, 0x55,
	0x43, 0x54, 0x5f, 0x41, 0x44, 0x10, 0x0a, 0x12, 0x0c, 0x0a, 0x08, 0x56, 0x49, 0x44, 0x45, 0x4f,
	0x5f, 0x41, 0x44, 0x10, 0x0c, 0x12, 0x0c, 0x0a, 0x08, 0x49, 0x4d, 0x41, 0x47, 0x45, 0x5f, 0x41,
	0x44, 0x10, 0x0e, 0x12, 0x18, 0x0a, 0x14, 0x52, 0x45, 0x53, 0x50, 0x4f, 0x4e, 0x53, 0x49, 0x56,
	0x45, 0x5f, 0x53, 0x45, 0x41, 0x52, 0x43, 0x48, 0x5f, 0x41, 0x44, 0x10, 0x0f, 0x12, 0x20, 0x0a,
	0x1c, 0x4c, 0x45, 0x47, 0x41, 0x43, 0x59, 0x5f, 0x52, 0x45, 0x53, 0x50, 0x4f, 0x4e, 0x53, 0x49,
	0x56, 0x45, 0x5f, 0x44, 0x49, 0x53, 0x50, 0x4c, 0x41, 0x59, 0x5f, 0x41, 0x44, 0x10, 0x10, 0x12,
	0x0a, 0x0a, 0x06, 0x41, 0x50, 0x50, 0x5f, 0x41, 0x44, 0x10, 0x11, 0x12, 0x19, 0x0a, 0x15, 0x4c,
	0x45, 0x47, 0x41, 0x43, 0x59, 0x5f, 0x41, 0x50, 0x50, 0x5f, 0x49, 0x4e, 0x53, 0x54, 0x41, 0x4c,
	0x4c, 0x5f, 0x41, 0x44, 0x10, 0x12, 0x12, 0x19, 0x0a, 0x15, 0x52, 0x45, 0x53, 0x50, 0x4f, 0x4e,
	0x53, 0x49, 0x56, 0x45, 0x5f, 0x44, 0x49, 0x53, 0x50, 0x4c, 0x41, 0x59, 0x5f, 0x41, 0x44, 0x10,
	0x13, 0x12, 0x0c, 0x0a, 0x08, 0x4c, 0x4f, 0x43, 0x41, 0x4c, 0x5f, 0x41, 0x44, 0x10, 0x14, 0x12,
	0x13, 0x0a, 0x0f, 0x48, 0x54, 0x4d, 0x4c, 0x35, 0x5f, 0x55, 0x50, 0x4c, 0x4f, 0x41, 0x44, 0x5f,
	0x41, 0x44, 0x10, 0x15, 0x12, 0x14, 0x0a, 0x10, 0x44, 0x59, 0x4e, 0x41, 0x4d, 0x49, 0x43, 0x5f,
	0x48, 0x54, 0x4d, 0x4c, 0x35, 0x5f, 0x41, 0x44, 0x10, 0x16, 0x12, 0x15, 0x0a, 0x11, 0x41, 0x50,
	0x50, 0x5f, 0x45, 0x4e, 0x47, 0x41, 0x47, 0x45, 0x4d, 0x45, 0x4e, 0x54, 0x5f, 0x41, 0x44, 0x10,
	0x17, 0x12, 0x22, 0x0a, 0x1e, 0x53, 0x48, 0x4f, 0x50, 0x50, 0x49, 0x4e, 0x47, 0x5f, 0x43, 0x4f,
	0x4d, 0x50, 0x41, 0x52, 0x49, 0x53, 0x4f, 0x4e, 0x5f, 0x4c, 0x49, 0x53, 0x54, 0x49, 0x4e, 0x47,
	0x5f, 0x41, 0x44, 0x10, 0x18, 0x12, 0x13, 0x0a, 0x0f, 0x56, 0x49, 0x44, 0x45, 0x4f, 0x5f, 0x42,
	0x55, 0x4d, 0x50, 0x45, 0x52, 0x5f, 0x41, 0x44, 0x10, 0x19, 0x12, 0x24, 0x0a, 0x20, 0x56, 0x49,
	0x44, 0x45, 0x4f, 0x5f, 0x4e, 0x4f, 0x4e, 0x5f, 0x53, 0x4b, 0x49, 0x50, 0x50, 0x41, 0x42, 0x4c,
	0x45, 0x5f, 0x49, 0x4e, 0x5f, 0x53, 0x54, 0x52, 0x45, 0x41, 0x4d, 0x5f, 0x41, 0x44, 0x10, 0x1a,
	0x12, 0x1f, 0x0a, 0x1b, 0x56, 0x49, 0x44, 0x45, 0x4f, 0x5f, 0x54, 0x52, 0x55, 0x45, 0x56, 0x49,
	0x45, 0x57, 0x5f, 0x49, 0x4e, 0x5f, 0x53, 0x54, 0x52, 0x45, 0x41, 0x4d, 0x5f, 0x41, 0x44, 0x10,
	0x1d, 0x12, 0x17, 0x0a, 0x13, 0x56, 0x49, 0x44, 0x45, 0x4f, 0x5f, 0x52, 0x45, 0x53, 0x50, 0x4f,
	0x4e, 0x53, 0x49, 0x56, 0x45, 0x5f, 0x41, 0x44, 0x10, 0x1e, 0x12, 0x15, 0x0a, 0x11, 0x53, 0x4d,
	0x41, 0x52, 0x54, 0x5f, 0x43, 0x41, 0x4d, 0x50, 0x41, 0x49, 0x47, 0x4e, 0x5f, 0x41, 0x44, 0x10,
	0x1f, 0x12, 0x0b, 0x0a, 0x07, 0x43, 0x41, 0x4c, 0x4c, 0x5f, 0x41, 0x44, 0x10, 0x20, 0x12, 0x1b,
	0x0a, 0x17, 0x41, 0x50, 0x50, 0x5f, 0x50, 0x52, 0x45, 0x5f, 0x52, 0x45, 0x47, 0x49, 0x53, 0x54,
	0x52, 0x41, 0x54, 0x49, 0x4f, 0x4e, 0x5f, 0x41, 0x44, 0x10, 0x21, 0x12, 0x14, 0x0a, 0x10, 0x49,
	0x4e, 0x5f, 0x46, 0x45, 0x45, 0x44, 0x5f, 0x56, 0x49, 0x44, 0x45, 0x4f, 0x5f, 0x41, 0x44, 0x10,
	0x22, 0x12, 0x1d, 0x0a, 0x19, 0x44, 0x45, 0x4d, 0x41, 0x4e, 0x44, 0x5f, 0x47, 0x45, 0x4e, 0x5f,
	0x4d, 0x55, 0x4c, 0x54, 0x49, 0x5f, 0x41, 0x53, 0x53, 0x45, 0x54, 0x5f, 0x41, 0x44, 0x10, 0x28,
	0x12, 0x1a, 0x0a, 0x16, 0x44, 0x45, 0x4d, 0x41, 0x4e, 0x44, 0x5f, 0x47, 0x45, 0x4e, 0x5f, 0x43,
	0x41, 0x52, 0x4f, 0x55, 0x53, 0x45, 0x4c, 0x5f, 0x41, 0x44, 0x10, 0x29, 0x12, 0x0d, 0x0a, 0x09,
	0x54, 0x52, 0x41, 0x56, 0x45, 0x4c, 0x5f, 0x41, 0x44, 0x10, 0x25, 0x12, 0x22, 0x0a, 0x1e, 0x44,
	0x45, 0x4d, 0x41, 0x4e, 0x44, 0x5f, 0x47, 0x45, 0x4e, 0x5f, 0x56, 0x49, 0x44, 0x45, 0x4f, 0x5f,
	0x52, 0x45, 0x53, 0x50, 0x4f, 0x4e, 0x53, 0x49, 0x56, 0x45, 0x5f, 0x41, 0x44, 0x10, 0x2a, 0x12,
	0x19, 0x0a, 0x15, 0x44, 0x45, 0x4d, 0x41, 0x4e, 0x44, 0x5f, 0x47, 0x45, 0x4e, 0x5f, 0x50, 0x52,
	0x4f, 0x44, 0x55, 0x43, 0x54, 0x5f, 0x41, 0x44, 0x10, 0x27, 0x12, 0x14, 0x0a, 0x10, 0x59, 0x4f,
	0x55, 0x54, 0x55, 0x42, 0x45, 0x5f, 0x41, 0x55, 0x44, 0x49, 0x4f, 0x5f, 0x41, 0x44, 0x10, 0x2c,
	0x42, 0xe5, 0x01, 0x0a, 0x22, 0x63, 0x6f, 0x6d, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e,
	0x61, 0x64, 0x73, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61, 0x64, 0x73, 0x2e, 0x76, 0x31,
	0x39, 0x2e, 0x65, 0x6e, 0x75, 0x6d, 0x73, 0x42, 0x0b, 0x41, 0x64, 0x54, 0x79, 0x70, 0x65, 0x50,
	0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a, 0x43, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x67,
	0x6f, 0x6c, 0x61, 0x6e, 0x67, 0x2e, 0x6f, 0x72, 0x67, 0x2f, 0x67, 0x65, 0x6e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x2f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61, 0x70, 0x69, 0x73, 0x2f, 0x61, 0x64,
	0x73, 0x2f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61, 0x64, 0x73, 0x2f, 0x76, 0x31, 0x39, 0x2f,
	0x65, 0x6e, 0x75, 0x6d, 0x73, 0x3b, 0x65, 0x6e, 0x75, 0x6d, 0x73, 0xa2, 0x02, 0x03, 0x47, 0x41,
	0x41, 0xaa, 0x02, 0x1e, 0x47, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x41, 0x64, 0x73, 0x2e, 0x47,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x41, 0x64, 0x73, 0x2e, 0x56, 0x31, 0x39, 0x2e, 0x45, 0x6e, 0x75,
	0x6d, 0x73, 0xca, 0x02, 0x1e, 0x47, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x5c, 0x41, 0x64, 0x73, 0x5c,
	0x47, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x41, 0x64, 0x73, 0x5c, 0x56, 0x31, 0x39, 0x5c, 0x45, 0x6e,
	0x75, 0x6d, 0x73, 0xea, 0x02, 0x22, 0x47, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x3a, 0x3a, 0x41, 0x64,
	0x73, 0x3a, 0x3a, 0x47, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x41, 0x64, 0x73, 0x3a, 0x3a, 0x56, 0x31,
	0x39, 0x3a, 0x3a, 0x45, 0x6e, 0x75, 0x6d, 0x73, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
})

var (
	file_enums_ad_type_proto_rawDescOnce sync.Once
	file_enums_ad_type_proto_rawDescData []byte
)

func file_enums_ad_type_proto_rawDescGZIP() []byte {
	file_enums_ad_type_proto_rawDescOnce.Do(func() {
		file_enums_ad_type_proto_rawDescData = protoimpl.X.CompressGZIP(unsafe.Slice(unsafe.StringData(file_enums_ad_type_proto_rawDesc), len(file_enums_ad_type_proto_rawDesc)))
	})
	return file_enums_ad_type_proto_rawDescData
}

var file_enums_ad_type_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_enums_ad_type_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_enums_ad_type_proto_goTypes = []any{
	(AdTypeEnum_AdType)(0), // 0: google.ads.googleads.v19.enums.AdTypeEnum.AdType
	(*AdTypeEnum)(nil),     // 1: google.ads.googleads.v19.enums.AdTypeEnum
}
var file_enums_ad_type_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_enums_ad_type_proto_init() }
func file_enums_ad_type_proto_init() {
	if File_enums_ad_type_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: unsafe.Slice(unsafe.StringData(file_enums_ad_type_proto_rawDesc), len(file_enums_ad_type_proto_rawDesc)),
			NumEnums:      1,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_enums_ad_type_proto_goTypes,
		DependencyIndexes: file_enums_ad_type_proto_depIdxs,
		EnumInfos:         file_enums_ad_type_proto_enumTypes,
		MessageInfos:      file_enums_ad_type_proto_msgTypes,
	}.Build()
	File_enums_ad_type_proto = out.File
	file_enums_ad_type_proto_goTypes = nil
	file_enums_ad_type_proto_depIdxs = nil
}
