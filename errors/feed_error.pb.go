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
// source: google/ads/googleads/v11/errors/feed_error.proto

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

// Enum describing possible feed errors.
type FeedErrorEnum_FeedError int32

const (
	// Enum unspecified.
	FeedErrorEnum_UNSPECIFIED FeedErrorEnum_FeedError = 0
	// The received error code is not known in this version.
	FeedErrorEnum_UNKNOWN FeedErrorEnum_FeedError = 1
	// The names of the FeedAttributes must be unique.
	FeedErrorEnum_ATTRIBUTE_NAMES_NOT_UNIQUE FeedErrorEnum_FeedError = 2
	// The attribute list must be an exact copy of the existing list if the
	// attribute ID's are present.
	FeedErrorEnum_ATTRIBUTES_DO_NOT_MATCH_EXISTING_ATTRIBUTES FeedErrorEnum_FeedError = 3
	// Cannot specify USER origin for a system generated feed.
	FeedErrorEnum_CANNOT_SPECIFY_USER_ORIGIN_FOR_SYSTEM_FEED FeedErrorEnum_FeedError = 4
	// Cannot specify GOOGLE origin for a non-system generated feed.
	FeedErrorEnum_CANNOT_SPECIFY_GOOGLE_ORIGIN_FOR_NON_SYSTEM_FEED FeedErrorEnum_FeedError = 5
	// Cannot specify feed attributes for system feed.
	FeedErrorEnum_CANNOT_SPECIFY_FEED_ATTRIBUTES_FOR_SYSTEM_FEED FeedErrorEnum_FeedError = 6
	// Cannot update FeedAttributes on feed with origin GOOGLE.
	FeedErrorEnum_CANNOT_UPDATE_FEED_ATTRIBUTES_WITH_ORIGIN_GOOGLE FeedErrorEnum_FeedError = 7
	// The given ID refers to a removed Feed. Removed Feeds are immutable.
	FeedErrorEnum_FEED_REMOVED FeedErrorEnum_FeedError = 8
	// The origin of the feed is not valid for the client.
	FeedErrorEnum_INVALID_ORIGIN_VALUE FeedErrorEnum_FeedError = 9
	// A user can only create and modify feeds with USER origin.
	FeedErrorEnum_FEED_ORIGIN_IS_NOT_USER FeedErrorEnum_FeedError = 10
	// Invalid auth token for the given email.
	FeedErrorEnum_INVALID_AUTH_TOKEN_FOR_EMAIL FeedErrorEnum_FeedError = 11
	// Invalid email specified.
	FeedErrorEnum_INVALID_EMAIL FeedErrorEnum_FeedError = 12
	// Feed name matches that of another active Feed.
	FeedErrorEnum_DUPLICATE_FEED_NAME FeedErrorEnum_FeedError = 13
	// Name of feed is not allowed.
	FeedErrorEnum_INVALID_FEED_NAME FeedErrorEnum_FeedError = 14
	// Missing OAuthInfo.
	FeedErrorEnum_MISSING_OAUTH_INFO FeedErrorEnum_FeedError = 15
	// New FeedAttributes must not affect the unique key.
	FeedErrorEnum_NEW_ATTRIBUTE_CANNOT_BE_PART_OF_UNIQUE_KEY FeedErrorEnum_FeedError = 16
	// Too many FeedAttributes for a Feed.
	FeedErrorEnum_TOO_MANY_ATTRIBUTES FeedErrorEnum_FeedError = 17
	// The business account is not valid.
	FeedErrorEnum_INVALID_BUSINESS_ACCOUNT FeedErrorEnum_FeedError = 18
	// Business account cannot access Business Profile.
	FeedErrorEnum_BUSINESS_ACCOUNT_CANNOT_ACCESS_LOCATION_ACCOUNT FeedErrorEnum_FeedError = 19
	// Invalid chain ID provided for affiliate location feed.
	FeedErrorEnum_INVALID_AFFILIATE_CHAIN_ID FeedErrorEnum_FeedError = 20
	// There is already a feed with the given system feed generation data.
	FeedErrorEnum_DUPLICATE_SYSTEM_FEED FeedErrorEnum_FeedError = 21
	// An error occurred accessing Business Profile.
	FeedErrorEnum_GMB_ACCESS_ERROR FeedErrorEnum_FeedError = 22
	// A customer cannot have both LOCATION and AFFILIATE_LOCATION feeds.
	FeedErrorEnum_CANNOT_HAVE_LOCATION_AND_AFFILIATE_LOCATION_FEEDS FeedErrorEnum_FeedError = 23
	// Feed-based extension is read-only for this extension type.
	FeedErrorEnum_LEGACY_EXTENSION_TYPE_READ_ONLY FeedErrorEnum_FeedError = 24
)

// Enum value maps for FeedErrorEnum_FeedError.
var (
	FeedErrorEnum_FeedError_name = map[int32]string{
		0:  "UNSPECIFIED",
		1:  "UNKNOWN",
		2:  "ATTRIBUTE_NAMES_NOT_UNIQUE",
		3:  "ATTRIBUTES_DO_NOT_MATCH_EXISTING_ATTRIBUTES",
		4:  "CANNOT_SPECIFY_USER_ORIGIN_FOR_SYSTEM_FEED",
		5:  "CANNOT_SPECIFY_GOOGLE_ORIGIN_FOR_NON_SYSTEM_FEED",
		6:  "CANNOT_SPECIFY_FEED_ATTRIBUTES_FOR_SYSTEM_FEED",
		7:  "CANNOT_UPDATE_FEED_ATTRIBUTES_WITH_ORIGIN_GOOGLE",
		8:  "FEED_REMOVED",
		9:  "INVALID_ORIGIN_VALUE",
		10: "FEED_ORIGIN_IS_NOT_USER",
		11: "INVALID_AUTH_TOKEN_FOR_EMAIL",
		12: "INVALID_EMAIL",
		13: "DUPLICATE_FEED_NAME",
		14: "INVALID_FEED_NAME",
		15: "MISSING_OAUTH_INFO",
		16: "NEW_ATTRIBUTE_CANNOT_BE_PART_OF_UNIQUE_KEY",
		17: "TOO_MANY_ATTRIBUTES",
		18: "INVALID_BUSINESS_ACCOUNT",
		19: "BUSINESS_ACCOUNT_CANNOT_ACCESS_LOCATION_ACCOUNT",
		20: "INVALID_AFFILIATE_CHAIN_ID",
		21: "DUPLICATE_SYSTEM_FEED",
		22: "GMB_ACCESS_ERROR",
		23: "CANNOT_HAVE_LOCATION_AND_AFFILIATE_LOCATION_FEEDS",
		24: "LEGACY_EXTENSION_TYPE_READ_ONLY",
	}
	FeedErrorEnum_FeedError_value = map[string]int32{
		"UNSPECIFIED":                0,
		"UNKNOWN":                    1,
		"ATTRIBUTE_NAMES_NOT_UNIQUE": 2,
		"ATTRIBUTES_DO_NOT_MATCH_EXISTING_ATTRIBUTES":      3,
		"CANNOT_SPECIFY_USER_ORIGIN_FOR_SYSTEM_FEED":       4,
		"CANNOT_SPECIFY_GOOGLE_ORIGIN_FOR_NON_SYSTEM_FEED": 5,
		"CANNOT_SPECIFY_FEED_ATTRIBUTES_FOR_SYSTEM_FEED":   6,
		"CANNOT_UPDATE_FEED_ATTRIBUTES_WITH_ORIGIN_GOOGLE": 7,
		"FEED_REMOVED":                                      8,
		"INVALID_ORIGIN_VALUE":                              9,
		"FEED_ORIGIN_IS_NOT_USER":                           10,
		"INVALID_AUTH_TOKEN_FOR_EMAIL":                      11,
		"INVALID_EMAIL":                                     12,
		"DUPLICATE_FEED_NAME":                               13,
		"INVALID_FEED_NAME":                                 14,
		"MISSING_OAUTH_INFO":                                15,
		"NEW_ATTRIBUTE_CANNOT_BE_PART_OF_UNIQUE_KEY":        16,
		"TOO_MANY_ATTRIBUTES":                               17,
		"INVALID_BUSINESS_ACCOUNT":                          18,
		"BUSINESS_ACCOUNT_CANNOT_ACCESS_LOCATION_ACCOUNT":   19,
		"INVALID_AFFILIATE_CHAIN_ID":                        20,
		"DUPLICATE_SYSTEM_FEED":                             21,
		"GMB_ACCESS_ERROR":                                  22,
		"CANNOT_HAVE_LOCATION_AND_AFFILIATE_LOCATION_FEEDS": 23,
		"LEGACY_EXTENSION_TYPE_READ_ONLY":                   24,
	}
)

func (x FeedErrorEnum_FeedError) Enum() *FeedErrorEnum_FeedError {
	p := new(FeedErrorEnum_FeedError)
	*p = x
	return p
}

func (x FeedErrorEnum_FeedError) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (FeedErrorEnum_FeedError) Descriptor() protoreflect.EnumDescriptor {
	return file_errors_feed_error_proto_enumTypes[0].Descriptor()
}

func (FeedErrorEnum_FeedError) Type() protoreflect.EnumType {
	return &file_errors_feed_error_proto_enumTypes[0]
}

func (x FeedErrorEnum_FeedError) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use FeedErrorEnum_FeedError.Descriptor instead.
func (FeedErrorEnum_FeedError) EnumDescriptor() ([]byte, []int) {
	return file_errors_feed_error_proto_rawDescGZIP(), []int{0, 0}
}

// Container for enum describing possible feed errors.
type FeedErrorEnum struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *FeedErrorEnum) Reset() {
	*x = FeedErrorEnum{}
	if protoimpl.UnsafeEnabled {
		mi := &file_errors_feed_error_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FeedErrorEnum) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FeedErrorEnum) ProtoMessage() {}

func (x *FeedErrorEnum) ProtoReflect() protoreflect.Message {
	mi := &file_errors_feed_error_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FeedErrorEnum.ProtoReflect.Descriptor instead.
func (*FeedErrorEnum) Descriptor() ([]byte, []int) {
	return file_errors_feed_error_proto_rawDescGZIP(), []int{0}
}

var File_errors_feed_error_proto protoreflect.FileDescriptor

var file_errors_feed_error_proto_rawDesc = []byte{
	0x0a, 0x30, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x64, 0x73, 0x2f, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x61, 0x64, 0x73, 0x2f, 0x76, 0x31, 0x31, 0x2f, 0x65, 0x72, 0x72, 0x6f, 0x72,
	0x73, 0x2f, 0x66, 0x65, 0x65, 0x64, 0x5f, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x12, 0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x61, 0x64, 0x73, 0x2e, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61, 0x64, 0x73, 0x2e, 0x76, 0x31, 0x31, 0x2e, 0x65, 0x72, 0x72,
	0x6f, 0x72, 0x73, 0x22, 0xeb, 0x06, 0x0a, 0x0d, 0x46, 0x65, 0x65, 0x64, 0x45, 0x72, 0x72, 0x6f,
	0x72, 0x45, 0x6e, 0x75, 0x6d, 0x22, 0xd9, 0x06, 0x0a, 0x09, 0x46, 0x65, 0x65, 0x64, 0x45, 0x72,
	0x72, 0x6f, 0x72, 0x12, 0x0f, 0x0a, 0x0b, 0x55, 0x4e, 0x53, 0x50, 0x45, 0x43, 0x49, 0x46, 0x49,
	0x45, 0x44, 0x10, 0x00, 0x12, 0x0b, 0x0a, 0x07, 0x55, 0x4e, 0x4b, 0x4e, 0x4f, 0x57, 0x4e, 0x10,
	0x01, 0x12, 0x1e, 0x0a, 0x1a, 0x41, 0x54, 0x54, 0x52, 0x49, 0x42, 0x55, 0x54, 0x45, 0x5f, 0x4e,
	0x41, 0x4d, 0x45, 0x53, 0x5f, 0x4e, 0x4f, 0x54, 0x5f, 0x55, 0x4e, 0x49, 0x51, 0x55, 0x45, 0x10,
	0x02, 0x12, 0x2f, 0x0a, 0x2b, 0x41, 0x54, 0x54, 0x52, 0x49, 0x42, 0x55, 0x54, 0x45, 0x53, 0x5f,
	0x44, 0x4f, 0x5f, 0x4e, 0x4f, 0x54, 0x5f, 0x4d, 0x41, 0x54, 0x43, 0x48, 0x5f, 0x45, 0x58, 0x49,
	0x53, 0x54, 0x49, 0x4e, 0x47, 0x5f, 0x41, 0x54, 0x54, 0x52, 0x49, 0x42, 0x55, 0x54, 0x45, 0x53,
	0x10, 0x03, 0x12, 0x2e, 0x0a, 0x2a, 0x43, 0x41, 0x4e, 0x4e, 0x4f, 0x54, 0x5f, 0x53, 0x50, 0x45,
	0x43, 0x49, 0x46, 0x59, 0x5f, 0x55, 0x53, 0x45, 0x52, 0x5f, 0x4f, 0x52, 0x49, 0x47, 0x49, 0x4e,
	0x5f, 0x46, 0x4f, 0x52, 0x5f, 0x53, 0x59, 0x53, 0x54, 0x45, 0x4d, 0x5f, 0x46, 0x45, 0x45, 0x44,
	0x10, 0x04, 0x12, 0x34, 0x0a, 0x30, 0x43, 0x41, 0x4e, 0x4e, 0x4f, 0x54, 0x5f, 0x53, 0x50, 0x45,
	0x43, 0x49, 0x46, 0x59, 0x5f, 0x47, 0x4f, 0x4f, 0x47, 0x4c, 0x45, 0x5f, 0x4f, 0x52, 0x49, 0x47,
	0x49, 0x4e, 0x5f, 0x46, 0x4f, 0x52, 0x5f, 0x4e, 0x4f, 0x4e, 0x5f, 0x53, 0x59, 0x53, 0x54, 0x45,
	0x4d, 0x5f, 0x46, 0x45, 0x45, 0x44, 0x10, 0x05, 0x12, 0x32, 0x0a, 0x2e, 0x43, 0x41, 0x4e, 0x4e,
	0x4f, 0x54, 0x5f, 0x53, 0x50, 0x45, 0x43, 0x49, 0x46, 0x59, 0x5f, 0x46, 0x45, 0x45, 0x44, 0x5f,
	0x41, 0x54, 0x54, 0x52, 0x49, 0x42, 0x55, 0x54, 0x45, 0x53, 0x5f, 0x46, 0x4f, 0x52, 0x5f, 0x53,
	0x59, 0x53, 0x54, 0x45, 0x4d, 0x5f, 0x46, 0x45, 0x45, 0x44, 0x10, 0x06, 0x12, 0x34, 0x0a, 0x30,
	0x43, 0x41, 0x4e, 0x4e, 0x4f, 0x54, 0x5f, 0x55, 0x50, 0x44, 0x41, 0x54, 0x45, 0x5f, 0x46, 0x45,
	0x45, 0x44, 0x5f, 0x41, 0x54, 0x54, 0x52, 0x49, 0x42, 0x55, 0x54, 0x45, 0x53, 0x5f, 0x57, 0x49,
	0x54, 0x48, 0x5f, 0x4f, 0x52, 0x49, 0x47, 0x49, 0x4e, 0x5f, 0x47, 0x4f, 0x4f, 0x47, 0x4c, 0x45,
	0x10, 0x07, 0x12, 0x10, 0x0a, 0x0c, 0x46, 0x45, 0x45, 0x44, 0x5f, 0x52, 0x45, 0x4d, 0x4f, 0x56,
	0x45, 0x44, 0x10, 0x08, 0x12, 0x18, 0x0a, 0x14, 0x49, 0x4e, 0x56, 0x41, 0x4c, 0x49, 0x44, 0x5f,
	0x4f, 0x52, 0x49, 0x47, 0x49, 0x4e, 0x5f, 0x56, 0x41, 0x4c, 0x55, 0x45, 0x10, 0x09, 0x12, 0x1b,
	0x0a, 0x17, 0x46, 0x45, 0x45, 0x44, 0x5f, 0x4f, 0x52, 0x49, 0x47, 0x49, 0x4e, 0x5f, 0x49, 0x53,
	0x5f, 0x4e, 0x4f, 0x54, 0x5f, 0x55, 0x53, 0x45, 0x52, 0x10, 0x0a, 0x12, 0x20, 0x0a, 0x1c, 0x49,
	0x4e, 0x56, 0x41, 0x4c, 0x49, 0x44, 0x5f, 0x41, 0x55, 0x54, 0x48, 0x5f, 0x54, 0x4f, 0x4b, 0x45,
	0x4e, 0x5f, 0x46, 0x4f, 0x52, 0x5f, 0x45, 0x4d, 0x41, 0x49, 0x4c, 0x10, 0x0b, 0x12, 0x11, 0x0a,
	0x0d, 0x49, 0x4e, 0x56, 0x41, 0x4c, 0x49, 0x44, 0x5f, 0x45, 0x4d, 0x41, 0x49, 0x4c, 0x10, 0x0c,
	0x12, 0x17, 0x0a, 0x13, 0x44, 0x55, 0x50, 0x4c, 0x49, 0x43, 0x41, 0x54, 0x45, 0x5f, 0x46, 0x45,
	0x45, 0x44, 0x5f, 0x4e, 0x41, 0x4d, 0x45, 0x10, 0x0d, 0x12, 0x15, 0x0a, 0x11, 0x49, 0x4e, 0x56,
	0x41, 0x4c, 0x49, 0x44, 0x5f, 0x46, 0x45, 0x45, 0x44, 0x5f, 0x4e, 0x41, 0x4d, 0x45, 0x10, 0x0e,
	0x12, 0x16, 0x0a, 0x12, 0x4d, 0x49, 0x53, 0x53, 0x49, 0x4e, 0x47, 0x5f, 0x4f, 0x41, 0x55, 0x54,
	0x48, 0x5f, 0x49, 0x4e, 0x46, 0x4f, 0x10, 0x0f, 0x12, 0x2e, 0x0a, 0x2a, 0x4e, 0x45, 0x57, 0x5f,
	0x41, 0x54, 0x54, 0x52, 0x49, 0x42, 0x55, 0x54, 0x45, 0x5f, 0x43, 0x41, 0x4e, 0x4e, 0x4f, 0x54,
	0x5f, 0x42, 0x45, 0x5f, 0x50, 0x41, 0x52, 0x54, 0x5f, 0x4f, 0x46, 0x5f, 0x55, 0x4e, 0x49, 0x51,
	0x55, 0x45, 0x5f, 0x4b, 0x45, 0x59, 0x10, 0x10, 0x12, 0x17, 0x0a, 0x13, 0x54, 0x4f, 0x4f, 0x5f,
	0x4d, 0x41, 0x4e, 0x59, 0x5f, 0x41, 0x54, 0x54, 0x52, 0x49, 0x42, 0x55, 0x54, 0x45, 0x53, 0x10,
	0x11, 0x12, 0x1c, 0x0a, 0x18, 0x49, 0x4e, 0x56, 0x41, 0x4c, 0x49, 0x44, 0x5f, 0x42, 0x55, 0x53,
	0x49, 0x4e, 0x45, 0x53, 0x53, 0x5f, 0x41, 0x43, 0x43, 0x4f, 0x55, 0x4e, 0x54, 0x10, 0x12, 0x12,
	0x33, 0x0a, 0x2f, 0x42, 0x55, 0x53, 0x49, 0x4e, 0x45, 0x53, 0x53, 0x5f, 0x41, 0x43, 0x43, 0x4f,
	0x55, 0x4e, 0x54, 0x5f, 0x43, 0x41, 0x4e, 0x4e, 0x4f, 0x54, 0x5f, 0x41, 0x43, 0x43, 0x45, 0x53,
	0x53, 0x5f, 0x4c, 0x4f, 0x43, 0x41, 0x54, 0x49, 0x4f, 0x4e, 0x5f, 0x41, 0x43, 0x43, 0x4f, 0x55,
	0x4e, 0x54, 0x10, 0x13, 0x12, 0x1e, 0x0a, 0x1a, 0x49, 0x4e, 0x56, 0x41, 0x4c, 0x49, 0x44, 0x5f,
	0x41, 0x46, 0x46, 0x49, 0x4c, 0x49, 0x41, 0x54, 0x45, 0x5f, 0x43, 0x48, 0x41, 0x49, 0x4e, 0x5f,
	0x49, 0x44, 0x10, 0x14, 0x12, 0x19, 0x0a, 0x15, 0x44, 0x55, 0x50, 0x4c, 0x49, 0x43, 0x41, 0x54,
	0x45, 0x5f, 0x53, 0x59, 0x53, 0x54, 0x45, 0x4d, 0x5f, 0x46, 0x45, 0x45, 0x44, 0x10, 0x15, 0x12,
	0x14, 0x0a, 0x10, 0x47, 0x4d, 0x42, 0x5f, 0x41, 0x43, 0x43, 0x45, 0x53, 0x53, 0x5f, 0x45, 0x52,
	0x52, 0x4f, 0x52, 0x10, 0x16, 0x12, 0x35, 0x0a, 0x31, 0x43, 0x41, 0x4e, 0x4e, 0x4f, 0x54, 0x5f,
	0x48, 0x41, 0x56, 0x45, 0x5f, 0x4c, 0x4f, 0x43, 0x41, 0x54, 0x49, 0x4f, 0x4e, 0x5f, 0x41, 0x4e,
	0x44, 0x5f, 0x41, 0x46, 0x46, 0x49, 0x4c, 0x49, 0x41, 0x54, 0x45, 0x5f, 0x4c, 0x4f, 0x43, 0x41,
	0x54, 0x49, 0x4f, 0x4e, 0x5f, 0x46, 0x45, 0x45, 0x44, 0x53, 0x10, 0x17, 0x12, 0x23, 0x0a, 0x1f,
	0x4c, 0x45, 0x47, 0x41, 0x43, 0x59, 0x5f, 0x45, 0x58, 0x54, 0x45, 0x4e, 0x53, 0x49, 0x4f, 0x4e,
	0x5f, 0x54, 0x59, 0x50, 0x45, 0x5f, 0x52, 0x45, 0x41, 0x44, 0x5f, 0x4f, 0x4e, 0x4c, 0x59, 0x10,
	0x18, 0x42, 0xee, 0x01, 0x0a, 0x23, 0x63, 0x6f, 0x6d, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x2e, 0x61, 0x64, 0x73, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61, 0x64, 0x73, 0x2e, 0x76,
	0x31, 0x31, 0x2e, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x73, 0x42, 0x0e, 0x46, 0x65, 0x65, 0x64, 0x45,
	0x72, 0x72, 0x6f, 0x72, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a, 0x45, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x2e, 0x67, 0x6f, 0x6c, 0x61, 0x6e, 0x67, 0x2e, 0x6f, 0x72, 0x67, 0x2f, 0x67,
	0x65, 0x6e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61, 0x70,
	0x69, 0x73, 0x2f, 0x61, 0x64, 0x73, 0x2f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61, 0x64, 0x73,
	0x2f, 0x76, 0x31, 0x31, 0x2f, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x73, 0x3b, 0x65, 0x72, 0x72, 0x6f,
	0x72, 0x73, 0xa2, 0x02, 0x03, 0x47, 0x41, 0x41, 0xaa, 0x02, 0x1f, 0x47, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x2e, 0x41, 0x64, 0x73, 0x2e, 0x47, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x41, 0x64, 0x73, 0x2e,
	0x56, 0x31, 0x31, 0x2e, 0x45, 0x72, 0x72, 0x6f, 0x72, 0x73, 0xca, 0x02, 0x1f, 0x47, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x5c, 0x41, 0x64, 0x73, 0x5c, 0x47, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x41, 0x64,
	0x73, 0x5c, 0x56, 0x31, 0x31, 0x5c, 0x45, 0x72, 0x72, 0x6f, 0x72, 0x73, 0xea, 0x02, 0x23, 0x47,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x3a, 0x3a, 0x41, 0x64, 0x73, 0x3a, 0x3a, 0x47, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x41, 0x64, 0x73, 0x3a, 0x3a, 0x56, 0x31, 0x31, 0x3a, 0x3a, 0x45, 0x72, 0x72, 0x6f,
	0x72, 0x73, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_errors_feed_error_proto_rawDescOnce sync.Once
	file_errors_feed_error_proto_rawDescData = file_errors_feed_error_proto_rawDesc
)

func file_errors_feed_error_proto_rawDescGZIP() []byte {
	file_errors_feed_error_proto_rawDescOnce.Do(func() {
		file_errors_feed_error_proto_rawDescData = protoimpl.X.CompressGZIP(file_errors_feed_error_proto_rawDescData)
	})
	return file_errors_feed_error_proto_rawDescData
}

var file_errors_feed_error_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_errors_feed_error_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_errors_feed_error_proto_goTypes = []interface{}{
	(FeedErrorEnum_FeedError)(0), // 0: google.ads.googleads.v11.errors.FeedErrorEnum.FeedError
	(*FeedErrorEnum)(nil),        // 1: google.ads.googleads.v11.errors.FeedErrorEnum
}
var file_errors_feed_error_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_errors_feed_error_proto_init() }
func file_errors_feed_error_proto_init() {
	if File_errors_feed_error_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_errors_feed_error_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*FeedErrorEnum); i {
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
			RawDescriptor: file_errors_feed_error_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_errors_feed_error_proto_goTypes,
		DependencyIndexes: file_errors_feed_error_proto_depIdxs,
		EnumInfos:         file_errors_feed_error_proto_enumTypes,
		MessageInfos:      file_errors_feed_error_proto_msgTypes,
	}.Build()
	File_errors_feed_error_proto = out.File
	file_errors_feed_error_proto_rawDesc = nil
	file_errors_feed_error_proto_goTypes = nil
	file_errors_feed_error_proto_depIdxs = nil
}
