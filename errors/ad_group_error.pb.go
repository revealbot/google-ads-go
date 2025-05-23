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
// source: google/ads/googleads/v19/errors/ad_group_error.proto

package errors

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

// Enum describing possible ad group errors.
type AdGroupErrorEnum_AdGroupError int32

const (
	// Enum unspecified.
	AdGroupErrorEnum_UNSPECIFIED AdGroupErrorEnum_AdGroupError = 0
	// The received error code is not known in this version.
	AdGroupErrorEnum_UNKNOWN AdGroupErrorEnum_AdGroupError = 1
	// AdGroup with the same name already exists for the campaign.
	AdGroupErrorEnum_DUPLICATE_ADGROUP_NAME AdGroupErrorEnum_AdGroupError = 2
	// AdGroup name is not valid.
	AdGroupErrorEnum_INVALID_ADGROUP_NAME AdGroupErrorEnum_AdGroupError = 3
	// Advertiser is not allowed to target sites or set site bids that are not
	// on the Google Search Network.
	AdGroupErrorEnum_ADVERTISER_NOT_ON_CONTENT_NETWORK AdGroupErrorEnum_AdGroupError = 5
	// Bid amount is too big.
	AdGroupErrorEnum_BID_TOO_BIG AdGroupErrorEnum_AdGroupError = 6
	// AdGroup bid does not match the campaign's bidding strategy.
	AdGroupErrorEnum_BID_TYPE_AND_BIDDING_STRATEGY_MISMATCH AdGroupErrorEnum_AdGroupError = 7
	// AdGroup name is required for Add.
	AdGroupErrorEnum_MISSING_ADGROUP_NAME AdGroupErrorEnum_AdGroupError = 8
	// No link found between the ad group and the label.
	AdGroupErrorEnum_ADGROUP_LABEL_DOES_NOT_EXIST AdGroupErrorEnum_AdGroupError = 9
	// The label has already been attached to the ad group.
	AdGroupErrorEnum_ADGROUP_LABEL_ALREADY_EXISTS AdGroupErrorEnum_AdGroupError = 10
	// The CriterionTypeGroup is not supported for the content bid dimension.
	AdGroupErrorEnum_INVALID_CONTENT_BID_CRITERION_TYPE_GROUP AdGroupErrorEnum_AdGroupError = 11
	// The ad group type is not compatible with the campaign channel type.
	AdGroupErrorEnum_AD_GROUP_TYPE_NOT_VALID_FOR_ADVERTISING_CHANNEL_TYPE AdGroupErrorEnum_AdGroupError = 12
	// The ad group type is not supported in the country of sale of the
	// campaign.
	AdGroupErrorEnum_ADGROUP_TYPE_NOT_SUPPORTED_FOR_CAMPAIGN_SALES_COUNTRY AdGroupErrorEnum_AdGroupError = 13
	// Ad groups of AdGroupType.SEARCH_DYNAMIC_ADS can only be added to
	// campaigns that have DynamicSearchAdsSetting attached.
	AdGroupErrorEnum_CANNOT_ADD_ADGROUP_OF_TYPE_DSA_TO_CAMPAIGN_WITHOUT_DSA_SETTING AdGroupErrorEnum_AdGroupError = 14
	// Promoted hotels ad groups are only available to customers on the
	// allow-list.
	AdGroupErrorEnum_PROMOTED_HOTEL_AD_GROUPS_NOT_AVAILABLE_FOR_CUSTOMER AdGroupErrorEnum_AdGroupError = 15
	// The field type cannot be excluded because an active ad group-asset link
	// of this type exists.
	AdGroupErrorEnum_INVALID_EXCLUDED_PARENT_ASSET_FIELD_TYPE AdGroupErrorEnum_AdGroupError = 16
	// The asset set type is invalid for setting the
	// excluded_parent_asset_set_types field.
	AdGroupErrorEnum_INVALID_EXCLUDED_PARENT_ASSET_SET_TYPE AdGroupErrorEnum_AdGroupError = 17
	// Cannot add ad groups for the campaign type.
	AdGroupErrorEnum_CANNOT_ADD_AD_GROUP_FOR_CAMPAIGN_TYPE AdGroupErrorEnum_AdGroupError = 18
	// Invalid status for the ad group.
	AdGroupErrorEnum_INVALID_STATUS AdGroupErrorEnum_AdGroupError = 19
)

// Enum value maps for AdGroupErrorEnum_AdGroupError.
var (
	AdGroupErrorEnum_AdGroupError_name = map[int32]string{
		0:  "UNSPECIFIED",
		1:  "UNKNOWN",
		2:  "DUPLICATE_ADGROUP_NAME",
		3:  "INVALID_ADGROUP_NAME",
		5:  "ADVERTISER_NOT_ON_CONTENT_NETWORK",
		6:  "BID_TOO_BIG",
		7:  "BID_TYPE_AND_BIDDING_STRATEGY_MISMATCH",
		8:  "MISSING_ADGROUP_NAME",
		9:  "ADGROUP_LABEL_DOES_NOT_EXIST",
		10: "ADGROUP_LABEL_ALREADY_EXISTS",
		11: "INVALID_CONTENT_BID_CRITERION_TYPE_GROUP",
		12: "AD_GROUP_TYPE_NOT_VALID_FOR_ADVERTISING_CHANNEL_TYPE",
		13: "ADGROUP_TYPE_NOT_SUPPORTED_FOR_CAMPAIGN_SALES_COUNTRY",
		14: "CANNOT_ADD_ADGROUP_OF_TYPE_DSA_TO_CAMPAIGN_WITHOUT_DSA_SETTING",
		15: "PROMOTED_HOTEL_AD_GROUPS_NOT_AVAILABLE_FOR_CUSTOMER",
		16: "INVALID_EXCLUDED_PARENT_ASSET_FIELD_TYPE",
		17: "INVALID_EXCLUDED_PARENT_ASSET_SET_TYPE",
		18: "CANNOT_ADD_AD_GROUP_FOR_CAMPAIGN_TYPE",
		19: "INVALID_STATUS",
	}
	AdGroupErrorEnum_AdGroupError_value = map[string]int32{
		"UNSPECIFIED":                                                    0,
		"UNKNOWN":                                                        1,
		"DUPLICATE_ADGROUP_NAME":                                         2,
		"INVALID_ADGROUP_NAME":                                           3,
		"ADVERTISER_NOT_ON_CONTENT_NETWORK":                              5,
		"BID_TOO_BIG":                                                    6,
		"BID_TYPE_AND_BIDDING_STRATEGY_MISMATCH":                         7,
		"MISSING_ADGROUP_NAME":                                           8,
		"ADGROUP_LABEL_DOES_NOT_EXIST":                                   9,
		"ADGROUP_LABEL_ALREADY_EXISTS":                                   10,
		"INVALID_CONTENT_BID_CRITERION_TYPE_GROUP":                       11,
		"AD_GROUP_TYPE_NOT_VALID_FOR_ADVERTISING_CHANNEL_TYPE":           12,
		"ADGROUP_TYPE_NOT_SUPPORTED_FOR_CAMPAIGN_SALES_COUNTRY":          13,
		"CANNOT_ADD_ADGROUP_OF_TYPE_DSA_TO_CAMPAIGN_WITHOUT_DSA_SETTING": 14,
		"PROMOTED_HOTEL_AD_GROUPS_NOT_AVAILABLE_FOR_CUSTOMER":            15,
		"INVALID_EXCLUDED_PARENT_ASSET_FIELD_TYPE":                       16,
		"INVALID_EXCLUDED_PARENT_ASSET_SET_TYPE":                         17,
		"CANNOT_ADD_AD_GROUP_FOR_CAMPAIGN_TYPE":                          18,
		"INVALID_STATUS":                                                 19,
	}
)

func (x AdGroupErrorEnum_AdGroupError) Enum() *AdGroupErrorEnum_AdGroupError {
	p := new(AdGroupErrorEnum_AdGroupError)
	*p = x
	return p
}

func (x AdGroupErrorEnum_AdGroupError) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (AdGroupErrorEnum_AdGroupError) Descriptor() protoreflect.EnumDescriptor {
	return file_errors_ad_group_error_proto_enumTypes[0].Descriptor()
}

func (AdGroupErrorEnum_AdGroupError) Type() protoreflect.EnumType {
	return &file_errors_ad_group_error_proto_enumTypes[0]
}

func (x AdGroupErrorEnum_AdGroupError) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use AdGroupErrorEnum_AdGroupError.Descriptor instead.
func (AdGroupErrorEnum_AdGroupError) EnumDescriptor() ([]byte, []int) {
	return file_errors_ad_group_error_proto_rawDescGZIP(), []int{0, 0}
}

// Container for enum describing possible ad group errors.
type AdGroupErrorEnum struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *AdGroupErrorEnum) Reset() {
	*x = AdGroupErrorEnum{}
	mi := &file_errors_ad_group_error_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *AdGroupErrorEnum) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AdGroupErrorEnum) ProtoMessage() {}

func (x *AdGroupErrorEnum) ProtoReflect() protoreflect.Message {
	mi := &file_errors_ad_group_error_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AdGroupErrorEnum.ProtoReflect.Descriptor instead.
func (*AdGroupErrorEnum) Descriptor() ([]byte, []int) {
	return file_errors_ad_group_error_proto_rawDescGZIP(), []int{0}
}

var File_errors_ad_group_error_proto protoreflect.FileDescriptor

var file_errors_ad_group_error_proto_rawDesc = string([]byte{
	0x0a, 0x34, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x64, 0x73, 0x2f, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x61, 0x64, 0x73, 0x2f, 0x76, 0x31, 0x39, 0x2f, 0x65, 0x72, 0x72, 0x6f, 0x72,
	0x73, 0x2f, 0x61, 0x64, 0x5f, 0x67, 0x72, 0x6f, 0x75, 0x70, 0x5f, 0x65, 0x72, 0x72, 0x6f, 0x72,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x61,
	0x64, 0x73, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61, 0x64, 0x73, 0x2e, 0x76, 0x31, 0x39,
	0x2e, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x73, 0x22, 0xf2, 0x05, 0x0a, 0x10, 0x41, 0x64, 0x47, 0x72,
	0x6f, 0x75, 0x70, 0x45, 0x72, 0x72, 0x6f, 0x72, 0x45, 0x6e, 0x75, 0x6d, 0x22, 0xdd, 0x05, 0x0a,
	0x0c, 0x41, 0x64, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x45, 0x72, 0x72, 0x6f, 0x72, 0x12, 0x0f, 0x0a,
	0x0b, 0x55, 0x4e, 0x53, 0x50, 0x45, 0x43, 0x49, 0x46, 0x49, 0x45, 0x44, 0x10, 0x00, 0x12, 0x0b,
	0x0a, 0x07, 0x55, 0x4e, 0x4b, 0x4e, 0x4f, 0x57, 0x4e, 0x10, 0x01, 0x12, 0x1a, 0x0a, 0x16, 0x44,
	0x55, 0x50, 0x4c, 0x49, 0x43, 0x41, 0x54, 0x45, 0x5f, 0x41, 0x44, 0x47, 0x52, 0x4f, 0x55, 0x50,
	0x5f, 0x4e, 0x41, 0x4d, 0x45, 0x10, 0x02, 0x12, 0x18, 0x0a, 0x14, 0x49, 0x4e, 0x56, 0x41, 0x4c,
	0x49, 0x44, 0x5f, 0x41, 0x44, 0x47, 0x52, 0x4f, 0x55, 0x50, 0x5f, 0x4e, 0x41, 0x4d, 0x45, 0x10,
	0x03, 0x12, 0x25, 0x0a, 0x21, 0x41, 0x44, 0x56, 0x45, 0x52, 0x54, 0x49, 0x53, 0x45, 0x52, 0x5f,
	0x4e, 0x4f, 0x54, 0x5f, 0x4f, 0x4e, 0x5f, 0x43, 0x4f, 0x4e, 0x54, 0x45, 0x4e, 0x54, 0x5f, 0x4e,
	0x45, 0x54, 0x57, 0x4f, 0x52, 0x4b, 0x10, 0x05, 0x12, 0x0f, 0x0a, 0x0b, 0x42, 0x49, 0x44, 0x5f,
	0x54, 0x4f, 0x4f, 0x5f, 0x42, 0x49, 0x47, 0x10, 0x06, 0x12, 0x2a, 0x0a, 0x26, 0x42, 0x49, 0x44,
	0x5f, 0x54, 0x59, 0x50, 0x45, 0x5f, 0x41, 0x4e, 0x44, 0x5f, 0x42, 0x49, 0x44, 0x44, 0x49, 0x4e,
	0x47, 0x5f, 0x53, 0x54, 0x52, 0x41, 0x54, 0x45, 0x47, 0x59, 0x5f, 0x4d, 0x49, 0x53, 0x4d, 0x41,
	0x54, 0x43, 0x48, 0x10, 0x07, 0x12, 0x18, 0x0a, 0x14, 0x4d, 0x49, 0x53, 0x53, 0x49, 0x4e, 0x47,
	0x5f, 0x41, 0x44, 0x47, 0x52, 0x4f, 0x55, 0x50, 0x5f, 0x4e, 0x41, 0x4d, 0x45, 0x10, 0x08, 0x12,
	0x20, 0x0a, 0x1c, 0x41, 0x44, 0x47, 0x52, 0x4f, 0x55, 0x50, 0x5f, 0x4c, 0x41, 0x42, 0x45, 0x4c,
	0x5f, 0x44, 0x4f, 0x45, 0x53, 0x5f, 0x4e, 0x4f, 0x54, 0x5f, 0x45, 0x58, 0x49, 0x53, 0x54, 0x10,
	0x09, 0x12, 0x20, 0x0a, 0x1c, 0x41, 0x44, 0x47, 0x52, 0x4f, 0x55, 0x50, 0x5f, 0x4c, 0x41, 0x42,
	0x45, 0x4c, 0x5f, 0x41, 0x4c, 0x52, 0x45, 0x41, 0x44, 0x59, 0x5f, 0x45, 0x58, 0x49, 0x53, 0x54,
	0x53, 0x10, 0x0a, 0x12, 0x2c, 0x0a, 0x28, 0x49, 0x4e, 0x56, 0x41, 0x4c, 0x49, 0x44, 0x5f, 0x43,
	0x4f, 0x4e, 0x54, 0x45, 0x4e, 0x54, 0x5f, 0x42, 0x49, 0x44, 0x5f, 0x43, 0x52, 0x49, 0x54, 0x45,
	0x52, 0x49, 0x4f, 0x4e, 0x5f, 0x54, 0x59, 0x50, 0x45, 0x5f, 0x47, 0x52, 0x4f, 0x55, 0x50, 0x10,
	0x0b, 0x12, 0x38, 0x0a, 0x34, 0x41, 0x44, 0x5f, 0x47, 0x52, 0x4f, 0x55, 0x50, 0x5f, 0x54, 0x59,
	0x50, 0x45, 0x5f, 0x4e, 0x4f, 0x54, 0x5f, 0x56, 0x41, 0x4c, 0x49, 0x44, 0x5f, 0x46, 0x4f, 0x52,
	0x5f, 0x41, 0x44, 0x56, 0x45, 0x52, 0x54, 0x49, 0x53, 0x49, 0x4e, 0x47, 0x5f, 0x43, 0x48, 0x41,
	0x4e, 0x4e, 0x45, 0x4c, 0x5f, 0x54, 0x59, 0x50, 0x45, 0x10, 0x0c, 0x12, 0x39, 0x0a, 0x35, 0x41,
	0x44, 0x47, 0x52, 0x4f, 0x55, 0x50, 0x5f, 0x54, 0x59, 0x50, 0x45, 0x5f, 0x4e, 0x4f, 0x54, 0x5f,
	0x53, 0x55, 0x50, 0x50, 0x4f, 0x52, 0x54, 0x45, 0x44, 0x5f, 0x46, 0x4f, 0x52, 0x5f, 0x43, 0x41,
	0x4d, 0x50, 0x41, 0x49, 0x47, 0x4e, 0x5f, 0x53, 0x41, 0x4c, 0x45, 0x53, 0x5f, 0x43, 0x4f, 0x55,
	0x4e, 0x54, 0x52, 0x59, 0x10, 0x0d, 0x12, 0x42, 0x0a, 0x3e, 0x43, 0x41, 0x4e, 0x4e, 0x4f, 0x54,
	0x5f, 0x41, 0x44, 0x44, 0x5f, 0x41, 0x44, 0x47, 0x52, 0x4f, 0x55, 0x50, 0x5f, 0x4f, 0x46, 0x5f,
	0x54, 0x59, 0x50, 0x45, 0x5f, 0x44, 0x53, 0x41, 0x5f, 0x54, 0x4f, 0x5f, 0x43, 0x41, 0x4d, 0x50,
	0x41, 0x49, 0x47, 0x4e, 0x5f, 0x57, 0x49, 0x54, 0x48, 0x4f, 0x55, 0x54, 0x5f, 0x44, 0x53, 0x41,
	0x5f, 0x53, 0x45, 0x54, 0x54, 0x49, 0x4e, 0x47, 0x10, 0x0e, 0x12, 0x37, 0x0a, 0x33, 0x50, 0x52,
	0x4f, 0x4d, 0x4f, 0x54, 0x45, 0x44, 0x5f, 0x48, 0x4f, 0x54, 0x45, 0x4c, 0x5f, 0x41, 0x44, 0x5f,
	0x47, 0x52, 0x4f, 0x55, 0x50, 0x53, 0x5f, 0x4e, 0x4f, 0x54, 0x5f, 0x41, 0x56, 0x41, 0x49, 0x4c,
	0x41, 0x42, 0x4c, 0x45, 0x5f, 0x46, 0x4f, 0x52, 0x5f, 0x43, 0x55, 0x53, 0x54, 0x4f, 0x4d, 0x45,
	0x52, 0x10, 0x0f, 0x12, 0x2c, 0x0a, 0x28, 0x49, 0x4e, 0x56, 0x41, 0x4c, 0x49, 0x44, 0x5f, 0x45,
	0x58, 0x43, 0x4c, 0x55, 0x44, 0x45, 0x44, 0x5f, 0x50, 0x41, 0x52, 0x45, 0x4e, 0x54, 0x5f, 0x41,
	0x53, 0x53, 0x45, 0x54, 0x5f, 0x46, 0x49, 0x45, 0x4c, 0x44, 0x5f, 0x54, 0x59, 0x50, 0x45, 0x10,
	0x10, 0x12, 0x2a, 0x0a, 0x26, 0x49, 0x4e, 0x56, 0x41, 0x4c, 0x49, 0x44, 0x5f, 0x45, 0x58, 0x43,
	0x4c, 0x55, 0x44, 0x45, 0x44, 0x5f, 0x50, 0x41, 0x52, 0x45, 0x4e, 0x54, 0x5f, 0x41, 0x53, 0x53,
	0x45, 0x54, 0x5f, 0x53, 0x45, 0x54, 0x5f, 0x54, 0x59, 0x50, 0x45, 0x10, 0x11, 0x12, 0x29, 0x0a,
	0x25, 0x43, 0x41, 0x4e, 0x4e, 0x4f, 0x54, 0x5f, 0x41, 0x44, 0x44, 0x5f, 0x41, 0x44, 0x5f, 0x47,
	0x52, 0x4f, 0x55, 0x50, 0x5f, 0x46, 0x4f, 0x52, 0x5f, 0x43, 0x41, 0x4d, 0x50, 0x41, 0x49, 0x47,
	0x4e, 0x5f, 0x54, 0x59, 0x50, 0x45, 0x10, 0x12, 0x12, 0x12, 0x0a, 0x0e, 0x49, 0x4e, 0x56, 0x41,
	0x4c, 0x49, 0x44, 0x5f, 0x53, 0x54, 0x41, 0x54, 0x55, 0x53, 0x10, 0x13, 0x42, 0xf1, 0x01, 0x0a,
	0x23, 0x63, 0x6f, 0x6d, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x61, 0x64, 0x73, 0x2e,
	0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61, 0x64, 0x73, 0x2e, 0x76, 0x31, 0x39, 0x2e, 0x65, 0x72,
	0x72, 0x6f, 0x72, 0x73, 0x42, 0x11, 0x41, 0x64, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x45, 0x72, 0x72,
	0x6f, 0x72, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a, 0x45, 0x67, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x2e, 0x67, 0x6f, 0x6c, 0x61, 0x6e, 0x67, 0x2e, 0x6f, 0x72, 0x67, 0x2f, 0x67, 0x65, 0x6e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61, 0x70, 0x69, 0x73,
	0x2f, 0x61, 0x64, 0x73, 0x2f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61, 0x64, 0x73, 0x2f, 0x76,
	0x31, 0x39, 0x2f, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x73, 0x3b, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x73,
	0xa2, 0x02, 0x03, 0x47, 0x41, 0x41, 0xaa, 0x02, 0x1f, 0x47, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e,
	0x41, 0x64, 0x73, 0x2e, 0x47, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x41, 0x64, 0x73, 0x2e, 0x56, 0x31,
	0x39, 0x2e, 0x45, 0x72, 0x72, 0x6f, 0x72, 0x73, 0xca, 0x02, 0x1f, 0x47, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x5c, 0x41, 0x64, 0x73, 0x5c, 0x47, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x41, 0x64, 0x73, 0x5c,
	0x56, 0x31, 0x39, 0x5c, 0x45, 0x72, 0x72, 0x6f, 0x72, 0x73, 0xea, 0x02, 0x23, 0x47, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x3a, 0x3a, 0x41, 0x64, 0x73, 0x3a, 0x3a, 0x47, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x41, 0x64, 0x73, 0x3a, 0x3a, 0x56, 0x31, 0x39, 0x3a, 0x3a, 0x45, 0x72, 0x72, 0x6f, 0x72, 0x73,
	0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
})

var (
	file_errors_ad_group_error_proto_rawDescOnce sync.Once
	file_errors_ad_group_error_proto_rawDescData []byte
)

func file_errors_ad_group_error_proto_rawDescGZIP() []byte {
	file_errors_ad_group_error_proto_rawDescOnce.Do(func() {
		file_errors_ad_group_error_proto_rawDescData = protoimpl.X.CompressGZIP(unsafe.Slice(unsafe.StringData(file_errors_ad_group_error_proto_rawDesc), len(file_errors_ad_group_error_proto_rawDesc)))
	})
	return file_errors_ad_group_error_proto_rawDescData
}

var file_errors_ad_group_error_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_errors_ad_group_error_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_errors_ad_group_error_proto_goTypes = []any{
	(AdGroupErrorEnum_AdGroupError)(0), // 0: google.ads.googleads.v19.errors.AdGroupErrorEnum.AdGroupError
	(*AdGroupErrorEnum)(nil),           // 1: google.ads.googleads.v19.errors.AdGroupErrorEnum
}
var file_errors_ad_group_error_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_errors_ad_group_error_proto_init() }
func file_errors_ad_group_error_proto_init() {
	if File_errors_ad_group_error_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: unsafe.Slice(unsafe.StringData(file_errors_ad_group_error_proto_rawDesc), len(file_errors_ad_group_error_proto_rawDesc)),
			NumEnums:      1,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_errors_ad_group_error_proto_goTypes,
		DependencyIndexes: file_errors_ad_group_error_proto_depIdxs,
		EnumInfos:         file_errors_ad_group_error_proto_enumTypes,
		MessageInfos:      file_errors_ad_group_error_proto_msgTypes,
	}.Build()
	File_errors_ad_group_error_proto = out.File
	file_errors_ad_group_error_proto_goTypes = nil
	file_errors_ad_group_error_proto_depIdxs = nil
}
