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
// 	protoc-gen-go v1.28.1
// 	protoc        v3.21.12
// source: google/ads/googleads/v14/errors/asset_link_error.proto

package errors

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

// Enum describing possible asset link errors.
type AssetLinkErrorEnum_AssetLinkError int32

const (
	// Enum unspecified.
	AssetLinkErrorEnum_UNSPECIFIED AssetLinkErrorEnum_AssetLinkError = 0
	// The received error code is not known in this version.
	AssetLinkErrorEnum_UNKNOWN AssetLinkErrorEnum_AssetLinkError = 1
	// Pinning is not supported for the given asset link field.
	AssetLinkErrorEnum_PINNING_UNSUPPORTED AssetLinkErrorEnum_AssetLinkError = 2
	// The given field type is not supported to be added directly through asset
	// links.
	AssetLinkErrorEnum_UNSUPPORTED_FIELD_TYPE AssetLinkErrorEnum_AssetLinkError = 3
	// The given asset's type and the specified field type are incompatible.
	AssetLinkErrorEnum_FIELD_TYPE_INCOMPATIBLE_WITH_ASSET_TYPE AssetLinkErrorEnum_AssetLinkError = 4
	// The specified field type is incompatible with the given campaign type.
	AssetLinkErrorEnum_FIELD_TYPE_INCOMPATIBLE_WITH_CAMPAIGN_TYPE AssetLinkErrorEnum_AssetLinkError = 5
	// The campaign advertising channel type cannot be associated with the given
	// asset due to channel-based restrictions on the asset's fields.
	AssetLinkErrorEnum_INCOMPATIBLE_ADVERTISING_CHANNEL_TYPE AssetLinkErrorEnum_AssetLinkError = 6
	// The image asset provided is not within the dimension constraints
	// specified for the submitted asset field.
	AssetLinkErrorEnum_IMAGE_NOT_WITHIN_SPECIFIED_DIMENSION_RANGE AssetLinkErrorEnum_AssetLinkError = 7
	// The pinned field is not valid for the submitted asset field.
	AssetLinkErrorEnum_INVALID_PINNED_FIELD AssetLinkErrorEnum_AssetLinkError = 8
	// The media bundle asset provided is too large for the submitted asset
	// field.
	AssetLinkErrorEnum_MEDIA_BUNDLE_ASSET_FILE_SIZE_TOO_LARGE AssetLinkErrorEnum_AssetLinkError = 9
	// Not enough assets are available for use with other fields since other
	// assets are pinned to specific fields.
	AssetLinkErrorEnum_NOT_ENOUGH_AVAILABLE_ASSET_LINKS_FOR_VALID_COMBINATION AssetLinkErrorEnum_AssetLinkError = 10
	// Not enough assets with fallback are available. When validating the
	// minimum number of assets, assets without fallback (for example, assets
	// that contain location tag without default value "{LOCATION(City)}") will
	// not be counted.
	AssetLinkErrorEnum_NOT_ENOUGH_AVAILABLE_ASSET_LINKS_WITH_FALLBACK AssetLinkErrorEnum_AssetLinkError = 11
	// This is a combination of the
	// NOT_ENOUGH_AVAILABLE_ASSET_LINKS_FOR_VALID_COMBINATION and
	// NOT_ENOUGH_AVAILABLE_ASSET_LINKS_WITH_FALLBACK errors. Not enough assets
	// with fallback are available since some assets are pinned.
	AssetLinkErrorEnum_NOT_ENOUGH_AVAILABLE_ASSET_LINKS_WITH_FALLBACK_FOR_VALID_COMBINATION AssetLinkErrorEnum_AssetLinkError = 12
	// The YouTube video referenced in the provided asset has been removed.
	AssetLinkErrorEnum_YOUTUBE_VIDEO_REMOVED AssetLinkErrorEnum_AssetLinkError = 13
	// The YouTube video referenced in the provided asset is too long for the
	// field submitted.
	AssetLinkErrorEnum_YOUTUBE_VIDEO_TOO_LONG AssetLinkErrorEnum_AssetLinkError = 14
	// The YouTube video referenced in the provided asset is too short for the
	// field submitted.
	AssetLinkErrorEnum_YOUTUBE_VIDEO_TOO_SHORT AssetLinkErrorEnum_AssetLinkError = 15
	// The specified field type is excluded for given campaign or ad group.
	AssetLinkErrorEnum_EXCLUDED_PARENT_FIELD_TYPE AssetLinkErrorEnum_AssetLinkError = 16
	// The status is invalid for the operation specified.
	AssetLinkErrorEnum_INVALID_STATUS AssetLinkErrorEnum_AssetLinkError = 17
	// The YouTube video referenced in the provided asset has unknown duration.
	// This might be the case for a livestream video or a video being currently
	// uploaded to YouTube. In both cases, the video duration should eventually
	// get resolved.
	AssetLinkErrorEnum_YOUTUBE_VIDEO_DURATION_NOT_DEFINED AssetLinkErrorEnum_AssetLinkError = 18
	// User cannot create automatically created links.
	AssetLinkErrorEnum_CANNOT_CREATE_AUTOMATICALLY_CREATED_LINKS AssetLinkErrorEnum_AssetLinkError = 19
	// Advertiser links cannot link to automatically created asset.
	AssetLinkErrorEnum_CANNOT_LINK_TO_AUTOMATICALLY_CREATED_ASSET AssetLinkErrorEnum_AssetLinkError = 20
	// Automatically created links cannot be changed into advertiser links or
	// the reverse.
	AssetLinkErrorEnum_CANNOT_MODIFY_ASSET_LINK_SOURCE AssetLinkErrorEnum_AssetLinkError = 21
	// Lead Form asset with Location answer type can't be linked to the
	// Customer/Campaign because there are no Location assets.
	AssetLinkErrorEnum_CANNOT_LINK_LOCATION_LEAD_FORM_WITHOUT_LOCATION_ASSET AssetLinkErrorEnum_AssetLinkError = 22
)

// Enum value maps for AssetLinkErrorEnum_AssetLinkError.
var (
	AssetLinkErrorEnum_AssetLinkError_name = map[int32]string{
		0:  "UNSPECIFIED",
		1:  "UNKNOWN",
		2:  "PINNING_UNSUPPORTED",
		3:  "UNSUPPORTED_FIELD_TYPE",
		4:  "FIELD_TYPE_INCOMPATIBLE_WITH_ASSET_TYPE",
		5:  "FIELD_TYPE_INCOMPATIBLE_WITH_CAMPAIGN_TYPE",
		6:  "INCOMPATIBLE_ADVERTISING_CHANNEL_TYPE",
		7:  "IMAGE_NOT_WITHIN_SPECIFIED_DIMENSION_RANGE",
		8:  "INVALID_PINNED_FIELD",
		9:  "MEDIA_BUNDLE_ASSET_FILE_SIZE_TOO_LARGE",
		10: "NOT_ENOUGH_AVAILABLE_ASSET_LINKS_FOR_VALID_COMBINATION",
		11: "NOT_ENOUGH_AVAILABLE_ASSET_LINKS_WITH_FALLBACK",
		12: "NOT_ENOUGH_AVAILABLE_ASSET_LINKS_WITH_FALLBACK_FOR_VALID_COMBINATION",
		13: "YOUTUBE_VIDEO_REMOVED",
		14: "YOUTUBE_VIDEO_TOO_LONG",
		15: "YOUTUBE_VIDEO_TOO_SHORT",
		16: "EXCLUDED_PARENT_FIELD_TYPE",
		17: "INVALID_STATUS",
		18: "YOUTUBE_VIDEO_DURATION_NOT_DEFINED",
		19: "CANNOT_CREATE_AUTOMATICALLY_CREATED_LINKS",
		20: "CANNOT_LINK_TO_AUTOMATICALLY_CREATED_ASSET",
		21: "CANNOT_MODIFY_ASSET_LINK_SOURCE",
		22: "CANNOT_LINK_LOCATION_LEAD_FORM_WITHOUT_LOCATION_ASSET",
	}
	AssetLinkErrorEnum_AssetLinkError_value = map[string]int32{
		"UNSPECIFIED":            0,
		"UNKNOWN":                1,
		"PINNING_UNSUPPORTED":    2,
		"UNSUPPORTED_FIELD_TYPE": 3,
		"FIELD_TYPE_INCOMPATIBLE_WITH_ASSET_TYPE":                              4,
		"FIELD_TYPE_INCOMPATIBLE_WITH_CAMPAIGN_TYPE":                           5,
		"INCOMPATIBLE_ADVERTISING_CHANNEL_TYPE":                                6,
		"IMAGE_NOT_WITHIN_SPECIFIED_DIMENSION_RANGE":                           7,
		"INVALID_PINNED_FIELD":                                                 8,
		"MEDIA_BUNDLE_ASSET_FILE_SIZE_TOO_LARGE":                               9,
		"NOT_ENOUGH_AVAILABLE_ASSET_LINKS_FOR_VALID_COMBINATION":               10,
		"NOT_ENOUGH_AVAILABLE_ASSET_LINKS_WITH_FALLBACK":                       11,
		"NOT_ENOUGH_AVAILABLE_ASSET_LINKS_WITH_FALLBACK_FOR_VALID_COMBINATION": 12,
		"YOUTUBE_VIDEO_REMOVED":                                                13,
		"YOUTUBE_VIDEO_TOO_LONG":                                               14,
		"YOUTUBE_VIDEO_TOO_SHORT":                                              15,
		"EXCLUDED_PARENT_FIELD_TYPE":                                           16,
		"INVALID_STATUS":                                                       17,
		"YOUTUBE_VIDEO_DURATION_NOT_DEFINED":                                   18,
		"CANNOT_CREATE_AUTOMATICALLY_CREATED_LINKS":                            19,
		"CANNOT_LINK_TO_AUTOMATICALLY_CREATED_ASSET":                           20,
		"CANNOT_MODIFY_ASSET_LINK_SOURCE":                                      21,
		"CANNOT_LINK_LOCATION_LEAD_FORM_WITHOUT_LOCATION_ASSET":                22,
	}
)

func (x AssetLinkErrorEnum_AssetLinkError) Enum() *AssetLinkErrorEnum_AssetLinkError {
	p := new(AssetLinkErrorEnum_AssetLinkError)
	*p = x
	return p
}

func (x AssetLinkErrorEnum_AssetLinkError) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (AssetLinkErrorEnum_AssetLinkError) Descriptor() protoreflect.EnumDescriptor {
	return file_errors_asset_link_error_proto_enumTypes[0].Descriptor()
}

func (AssetLinkErrorEnum_AssetLinkError) Type() protoreflect.EnumType {
	return &file_errors_asset_link_error_proto_enumTypes[0]
}

func (x AssetLinkErrorEnum_AssetLinkError) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use AssetLinkErrorEnum_AssetLinkError.Descriptor instead.
func (AssetLinkErrorEnum_AssetLinkError) EnumDescriptor() ([]byte, []int) {
	return file_errors_asset_link_error_proto_rawDescGZIP(), []int{0, 0}
}

// Container for enum describing possible asset link errors.
type AssetLinkErrorEnum struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *AssetLinkErrorEnum) Reset() {
	*x = AssetLinkErrorEnum{}
	if protoimpl.UnsafeEnabled {
		mi := &file_errors_asset_link_error_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AssetLinkErrorEnum) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AssetLinkErrorEnum) ProtoMessage() {}

func (x *AssetLinkErrorEnum) ProtoReflect() protoreflect.Message {
	mi := &file_errors_asset_link_error_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AssetLinkErrorEnum.ProtoReflect.Descriptor instead.
func (*AssetLinkErrorEnum) Descriptor() ([]byte, []int) {
	return file_errors_asset_link_error_proto_rawDescGZIP(), []int{0}
}

var File_errors_asset_link_error_proto protoreflect.FileDescriptor

var file_errors_asset_link_error_proto_rawDesc = []byte{
	0x0a, 0x36, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x64, 0x73, 0x2f, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x61, 0x64, 0x73, 0x2f, 0x76, 0x31, 0x34, 0x2f, 0x65, 0x72, 0x72, 0x6f, 0x72,
	0x73, 0x2f, 0x61, 0x73, 0x73, 0x65, 0x74, 0x5f, 0x6c, 0x69, 0x6e, 0x6b, 0x5f, 0x65, 0x72, 0x72,
	0x6f, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x2e, 0x61, 0x64, 0x73, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61, 0x64, 0x73, 0x2e, 0x76,
	0x31, 0x34, 0x2e, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x73, 0x22, 0xa1, 0x07, 0x0a, 0x12, 0x41, 0x73,
	0x73, 0x65, 0x74, 0x4c, 0x69, 0x6e, 0x6b, 0x45, 0x72, 0x72, 0x6f, 0x72, 0x45, 0x6e, 0x75, 0x6d,
	0x22, 0x8a, 0x07, 0x0a, 0x0e, 0x41, 0x73, 0x73, 0x65, 0x74, 0x4c, 0x69, 0x6e, 0x6b, 0x45, 0x72,
	0x72, 0x6f, 0x72, 0x12, 0x0f, 0x0a, 0x0b, 0x55, 0x4e, 0x53, 0x50, 0x45, 0x43, 0x49, 0x46, 0x49,
	0x45, 0x44, 0x10, 0x00, 0x12, 0x0b, 0x0a, 0x07, 0x55, 0x4e, 0x4b, 0x4e, 0x4f, 0x57, 0x4e, 0x10,
	0x01, 0x12, 0x17, 0x0a, 0x13, 0x50, 0x49, 0x4e, 0x4e, 0x49, 0x4e, 0x47, 0x5f, 0x55, 0x4e, 0x53,
	0x55, 0x50, 0x50, 0x4f, 0x52, 0x54, 0x45, 0x44, 0x10, 0x02, 0x12, 0x1a, 0x0a, 0x16, 0x55, 0x4e,
	0x53, 0x55, 0x50, 0x50, 0x4f, 0x52, 0x54, 0x45, 0x44, 0x5f, 0x46, 0x49, 0x45, 0x4c, 0x44, 0x5f,
	0x54, 0x59, 0x50, 0x45, 0x10, 0x03, 0x12, 0x2b, 0x0a, 0x27, 0x46, 0x49, 0x45, 0x4c, 0x44, 0x5f,
	0x54, 0x59, 0x50, 0x45, 0x5f, 0x49, 0x4e, 0x43, 0x4f, 0x4d, 0x50, 0x41, 0x54, 0x49, 0x42, 0x4c,
	0x45, 0x5f, 0x57, 0x49, 0x54, 0x48, 0x5f, 0x41, 0x53, 0x53, 0x45, 0x54, 0x5f, 0x54, 0x59, 0x50,
	0x45, 0x10, 0x04, 0x12, 0x2e, 0x0a, 0x2a, 0x46, 0x49, 0x45, 0x4c, 0x44, 0x5f, 0x54, 0x59, 0x50,
	0x45, 0x5f, 0x49, 0x4e, 0x43, 0x4f, 0x4d, 0x50, 0x41, 0x54, 0x49, 0x42, 0x4c, 0x45, 0x5f, 0x57,
	0x49, 0x54, 0x48, 0x5f, 0x43, 0x41, 0x4d, 0x50, 0x41, 0x49, 0x47, 0x4e, 0x5f, 0x54, 0x59, 0x50,
	0x45, 0x10, 0x05, 0x12, 0x29, 0x0a, 0x25, 0x49, 0x4e, 0x43, 0x4f, 0x4d, 0x50, 0x41, 0x54, 0x49,
	0x42, 0x4c, 0x45, 0x5f, 0x41, 0x44, 0x56, 0x45, 0x52, 0x54, 0x49, 0x53, 0x49, 0x4e, 0x47, 0x5f,
	0x43, 0x48, 0x41, 0x4e, 0x4e, 0x45, 0x4c, 0x5f, 0x54, 0x59, 0x50, 0x45, 0x10, 0x06, 0x12, 0x2e,
	0x0a, 0x2a, 0x49, 0x4d, 0x41, 0x47, 0x45, 0x5f, 0x4e, 0x4f, 0x54, 0x5f, 0x57, 0x49, 0x54, 0x48,
	0x49, 0x4e, 0x5f, 0x53, 0x50, 0x45, 0x43, 0x49, 0x46, 0x49, 0x45, 0x44, 0x5f, 0x44, 0x49, 0x4d,
	0x45, 0x4e, 0x53, 0x49, 0x4f, 0x4e, 0x5f, 0x52, 0x41, 0x4e, 0x47, 0x45, 0x10, 0x07, 0x12, 0x18,
	0x0a, 0x14, 0x49, 0x4e, 0x56, 0x41, 0x4c, 0x49, 0x44, 0x5f, 0x50, 0x49, 0x4e, 0x4e, 0x45, 0x44,
	0x5f, 0x46, 0x49, 0x45, 0x4c, 0x44, 0x10, 0x08, 0x12, 0x2a, 0x0a, 0x26, 0x4d, 0x45, 0x44, 0x49,
	0x41, 0x5f, 0x42, 0x55, 0x4e, 0x44, 0x4c, 0x45, 0x5f, 0x41, 0x53, 0x53, 0x45, 0x54, 0x5f, 0x46,
	0x49, 0x4c, 0x45, 0x5f, 0x53, 0x49, 0x5a, 0x45, 0x5f, 0x54, 0x4f, 0x4f, 0x5f, 0x4c, 0x41, 0x52,
	0x47, 0x45, 0x10, 0x09, 0x12, 0x3a, 0x0a, 0x36, 0x4e, 0x4f, 0x54, 0x5f, 0x45, 0x4e, 0x4f, 0x55,
	0x47, 0x48, 0x5f, 0x41, 0x56, 0x41, 0x49, 0x4c, 0x41, 0x42, 0x4c, 0x45, 0x5f, 0x41, 0x53, 0x53,
	0x45, 0x54, 0x5f, 0x4c, 0x49, 0x4e, 0x4b, 0x53, 0x5f, 0x46, 0x4f, 0x52, 0x5f, 0x56, 0x41, 0x4c,
	0x49, 0x44, 0x5f, 0x43, 0x4f, 0x4d, 0x42, 0x49, 0x4e, 0x41, 0x54, 0x49, 0x4f, 0x4e, 0x10, 0x0a,
	0x12, 0x32, 0x0a, 0x2e, 0x4e, 0x4f, 0x54, 0x5f, 0x45, 0x4e, 0x4f, 0x55, 0x47, 0x48, 0x5f, 0x41,
	0x56, 0x41, 0x49, 0x4c, 0x41, 0x42, 0x4c, 0x45, 0x5f, 0x41, 0x53, 0x53, 0x45, 0x54, 0x5f, 0x4c,
	0x49, 0x4e, 0x4b, 0x53, 0x5f, 0x57, 0x49, 0x54, 0x48, 0x5f, 0x46, 0x41, 0x4c, 0x4c, 0x42, 0x41,
	0x43, 0x4b, 0x10, 0x0b, 0x12, 0x48, 0x0a, 0x44, 0x4e, 0x4f, 0x54, 0x5f, 0x45, 0x4e, 0x4f, 0x55,
	0x47, 0x48, 0x5f, 0x41, 0x56, 0x41, 0x49, 0x4c, 0x41, 0x42, 0x4c, 0x45, 0x5f, 0x41, 0x53, 0x53,
	0x45, 0x54, 0x5f, 0x4c, 0x49, 0x4e, 0x4b, 0x53, 0x5f, 0x57, 0x49, 0x54, 0x48, 0x5f, 0x46, 0x41,
	0x4c, 0x4c, 0x42, 0x41, 0x43, 0x4b, 0x5f, 0x46, 0x4f, 0x52, 0x5f, 0x56, 0x41, 0x4c, 0x49, 0x44,
	0x5f, 0x43, 0x4f, 0x4d, 0x42, 0x49, 0x4e, 0x41, 0x54, 0x49, 0x4f, 0x4e, 0x10, 0x0c, 0x12, 0x19,
	0x0a, 0x15, 0x59, 0x4f, 0x55, 0x54, 0x55, 0x42, 0x45, 0x5f, 0x56, 0x49, 0x44, 0x45, 0x4f, 0x5f,
	0x52, 0x45, 0x4d, 0x4f, 0x56, 0x45, 0x44, 0x10, 0x0d, 0x12, 0x1a, 0x0a, 0x16, 0x59, 0x4f, 0x55,
	0x54, 0x55, 0x42, 0x45, 0x5f, 0x56, 0x49, 0x44, 0x45, 0x4f, 0x5f, 0x54, 0x4f, 0x4f, 0x5f, 0x4c,
	0x4f, 0x4e, 0x47, 0x10, 0x0e, 0x12, 0x1b, 0x0a, 0x17, 0x59, 0x4f, 0x55, 0x54, 0x55, 0x42, 0x45,
	0x5f, 0x56, 0x49, 0x44, 0x45, 0x4f, 0x5f, 0x54, 0x4f, 0x4f, 0x5f, 0x53, 0x48, 0x4f, 0x52, 0x54,
	0x10, 0x0f, 0x12, 0x1e, 0x0a, 0x1a, 0x45, 0x58, 0x43, 0x4c, 0x55, 0x44, 0x45, 0x44, 0x5f, 0x50,
	0x41, 0x52, 0x45, 0x4e, 0x54, 0x5f, 0x46, 0x49, 0x45, 0x4c, 0x44, 0x5f, 0x54, 0x59, 0x50, 0x45,
	0x10, 0x10, 0x12, 0x12, 0x0a, 0x0e, 0x49, 0x4e, 0x56, 0x41, 0x4c, 0x49, 0x44, 0x5f, 0x53, 0x54,
	0x41, 0x54, 0x55, 0x53, 0x10, 0x11, 0x12, 0x26, 0x0a, 0x22, 0x59, 0x4f, 0x55, 0x54, 0x55, 0x42,
	0x45, 0x5f, 0x56, 0x49, 0x44, 0x45, 0x4f, 0x5f, 0x44, 0x55, 0x52, 0x41, 0x54, 0x49, 0x4f, 0x4e,
	0x5f, 0x4e, 0x4f, 0x54, 0x5f, 0x44, 0x45, 0x46, 0x49, 0x4e, 0x45, 0x44, 0x10, 0x12, 0x12, 0x2d,
	0x0a, 0x29, 0x43, 0x41, 0x4e, 0x4e, 0x4f, 0x54, 0x5f, 0x43, 0x52, 0x45, 0x41, 0x54, 0x45, 0x5f,
	0x41, 0x55, 0x54, 0x4f, 0x4d, 0x41, 0x54, 0x49, 0x43, 0x41, 0x4c, 0x4c, 0x59, 0x5f, 0x43, 0x52,
	0x45, 0x41, 0x54, 0x45, 0x44, 0x5f, 0x4c, 0x49, 0x4e, 0x4b, 0x53, 0x10, 0x13, 0x12, 0x2e, 0x0a,
	0x2a, 0x43, 0x41, 0x4e, 0x4e, 0x4f, 0x54, 0x5f, 0x4c, 0x49, 0x4e, 0x4b, 0x5f, 0x54, 0x4f, 0x5f,
	0x41, 0x55, 0x54, 0x4f, 0x4d, 0x41, 0x54, 0x49, 0x43, 0x41, 0x4c, 0x4c, 0x59, 0x5f, 0x43, 0x52,
	0x45, 0x41, 0x54, 0x45, 0x44, 0x5f, 0x41, 0x53, 0x53, 0x45, 0x54, 0x10, 0x14, 0x12, 0x23, 0x0a,
	0x1f, 0x43, 0x41, 0x4e, 0x4e, 0x4f, 0x54, 0x5f, 0x4d, 0x4f, 0x44, 0x49, 0x46, 0x59, 0x5f, 0x41,
	0x53, 0x53, 0x45, 0x54, 0x5f, 0x4c, 0x49, 0x4e, 0x4b, 0x5f, 0x53, 0x4f, 0x55, 0x52, 0x43, 0x45,
	0x10, 0x15, 0x12, 0x39, 0x0a, 0x35, 0x43, 0x41, 0x4e, 0x4e, 0x4f, 0x54, 0x5f, 0x4c, 0x49, 0x4e,
	0x4b, 0x5f, 0x4c, 0x4f, 0x43, 0x41, 0x54, 0x49, 0x4f, 0x4e, 0x5f, 0x4c, 0x45, 0x41, 0x44, 0x5f,
	0x46, 0x4f, 0x52, 0x4d, 0x5f, 0x57, 0x49, 0x54, 0x48, 0x4f, 0x55, 0x54, 0x5f, 0x4c, 0x4f, 0x43,
	0x41, 0x54, 0x49, 0x4f, 0x4e, 0x5f, 0x41, 0x53, 0x53, 0x45, 0x54, 0x10, 0x16, 0x42, 0xf3, 0x01,
	0x0a, 0x23, 0x63, 0x6f, 0x6d, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x61, 0x64, 0x73,
	0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61, 0x64, 0x73, 0x2e, 0x76, 0x31, 0x34, 0x2e, 0x65,
	0x72, 0x72, 0x6f, 0x72, 0x73, 0x42, 0x13, 0x41, 0x73, 0x73, 0x65, 0x74, 0x4c, 0x69, 0x6e, 0x6b,
	0x45, 0x72, 0x72, 0x6f, 0x72, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a, 0x45, 0x67, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x67, 0x6f, 0x6c, 0x61, 0x6e, 0x67, 0x2e, 0x6f, 0x72, 0x67, 0x2f,
	0x67, 0x65, 0x6e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61,
	0x70, 0x69, 0x73, 0x2f, 0x61, 0x64, 0x73, 0x2f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61, 0x64,
	0x73, 0x2f, 0x76, 0x31, 0x34, 0x2f, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x73, 0x3b, 0x65, 0x72, 0x72,
	0x6f, 0x72, 0x73, 0xa2, 0x02, 0x03, 0x47, 0x41, 0x41, 0xaa, 0x02, 0x1f, 0x47, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x2e, 0x41, 0x64, 0x73, 0x2e, 0x47, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x41, 0x64, 0x73,
	0x2e, 0x56, 0x31, 0x34, 0x2e, 0x45, 0x72, 0x72, 0x6f, 0x72, 0x73, 0xca, 0x02, 0x1f, 0x47, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x5c, 0x41, 0x64, 0x73, 0x5c, 0x47, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x41,
	0x64, 0x73, 0x5c, 0x56, 0x31, 0x34, 0x5c, 0x45, 0x72, 0x72, 0x6f, 0x72, 0x73, 0xea, 0x02, 0x23,
	0x47, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x3a, 0x3a, 0x41, 0x64, 0x73, 0x3a, 0x3a, 0x47, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x41, 0x64, 0x73, 0x3a, 0x3a, 0x56, 0x31, 0x34, 0x3a, 0x3a, 0x45, 0x72, 0x72,
	0x6f, 0x72, 0x73, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_errors_asset_link_error_proto_rawDescOnce sync.Once
	file_errors_asset_link_error_proto_rawDescData = file_errors_asset_link_error_proto_rawDesc
)

func file_errors_asset_link_error_proto_rawDescGZIP() []byte {
	file_errors_asset_link_error_proto_rawDescOnce.Do(func() {
		file_errors_asset_link_error_proto_rawDescData = protoimpl.X.CompressGZIP(file_errors_asset_link_error_proto_rawDescData)
	})
	return file_errors_asset_link_error_proto_rawDescData
}

var file_errors_asset_link_error_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_errors_asset_link_error_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_errors_asset_link_error_proto_goTypes = []interface{}{
	(AssetLinkErrorEnum_AssetLinkError)(0), // 0: google.ads.googleads.v14.errors.AssetLinkErrorEnum.AssetLinkError
	(*AssetLinkErrorEnum)(nil),             // 1: google.ads.googleads.v14.errors.AssetLinkErrorEnum
}
var file_errors_asset_link_error_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_errors_asset_link_error_proto_init() }
func file_errors_asset_link_error_proto_init() {
	if File_errors_asset_link_error_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_errors_asset_link_error_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AssetLinkErrorEnum); i {
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
			RawDescriptor: file_errors_asset_link_error_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_errors_asset_link_error_proto_goTypes,
		DependencyIndexes: file_errors_asset_link_error_proto_depIdxs,
		EnumInfos:         file_errors_asset_link_error_proto_enumTypes,
		MessageInfos:      file_errors_asset_link_error_proto_msgTypes,
	}.Build()
	File_errors_asset_link_error_proto = out.File
	file_errors_asset_link_error_proto_rawDesc = nil
	file_errors_asset_link_error_proto_goTypes = nil
	file_errors_asset_link_error_proto_depIdxs = nil
}
