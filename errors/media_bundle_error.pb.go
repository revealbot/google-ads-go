// Copyright 2021 Google LLC
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
// 	protoc        v3.17.3
// source: google/ads/googleads/v9/errors/media_bundle_error.proto

package errors

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

// Enum describing possible media bundle errors.
type MediaBundleErrorEnum_MediaBundleError int32

const (
	// Enum unspecified.
	MediaBundleErrorEnum_UNSPECIFIED MediaBundleErrorEnum_MediaBundleError = 0
	// The received error code is not known in this version.
	MediaBundleErrorEnum_UNKNOWN MediaBundleErrorEnum_MediaBundleError = 1
	// There was a problem with the request.
	MediaBundleErrorEnum_BAD_REQUEST MediaBundleErrorEnum_MediaBundleError = 3
	// HTML5 ads using DoubleClick Studio created ZIP files are not supported.
	MediaBundleErrorEnum_DOUBLECLICK_BUNDLE_NOT_ALLOWED MediaBundleErrorEnum_MediaBundleError = 4
	// Cannot reference URL external to the media bundle.
	MediaBundleErrorEnum_EXTERNAL_URL_NOT_ALLOWED MediaBundleErrorEnum_MediaBundleError = 5
	// Media bundle file is too large.
	MediaBundleErrorEnum_FILE_TOO_LARGE MediaBundleErrorEnum_MediaBundleError = 6
	// ZIP file from Google Web Designer is not published.
	MediaBundleErrorEnum_GOOGLE_WEB_DESIGNER_ZIP_FILE_NOT_PUBLISHED MediaBundleErrorEnum_MediaBundleError = 7
	// Input was invalid.
	MediaBundleErrorEnum_INVALID_INPUT MediaBundleErrorEnum_MediaBundleError = 8
	// There was a problem with the media bundle.
	MediaBundleErrorEnum_INVALID_MEDIA_BUNDLE MediaBundleErrorEnum_MediaBundleError = 9
	// There was a problem with one or more of the media bundle entries.
	MediaBundleErrorEnum_INVALID_MEDIA_BUNDLE_ENTRY MediaBundleErrorEnum_MediaBundleError = 10
	// The media bundle contains a file with an unknown mime type
	MediaBundleErrorEnum_INVALID_MIME_TYPE MediaBundleErrorEnum_MediaBundleError = 11
	// The media bundle contain an invalid asset path.
	MediaBundleErrorEnum_INVALID_PATH MediaBundleErrorEnum_MediaBundleError = 12
	// HTML5 ad is trying to reference an asset not in .ZIP file
	MediaBundleErrorEnum_INVALID_URL_REFERENCE MediaBundleErrorEnum_MediaBundleError = 13
	// Media data is too large.
	MediaBundleErrorEnum_MEDIA_DATA_TOO_LARGE MediaBundleErrorEnum_MediaBundleError = 14
	// The media bundle contains no primary entry.
	MediaBundleErrorEnum_MISSING_PRIMARY_MEDIA_BUNDLE_ENTRY MediaBundleErrorEnum_MediaBundleError = 15
	// There was an error on the server.
	MediaBundleErrorEnum_SERVER_ERROR MediaBundleErrorEnum_MediaBundleError = 16
	// The image could not be stored.
	MediaBundleErrorEnum_STORAGE_ERROR MediaBundleErrorEnum_MediaBundleError = 17
	// Media bundle created with the Swiffy tool is not allowed.
	MediaBundleErrorEnum_SWIFFY_BUNDLE_NOT_ALLOWED MediaBundleErrorEnum_MediaBundleError = 18
	// The media bundle contains too many files.
	MediaBundleErrorEnum_TOO_MANY_FILES MediaBundleErrorEnum_MediaBundleError = 19
	// The media bundle is not of legal dimensions.
	MediaBundleErrorEnum_UNEXPECTED_SIZE MediaBundleErrorEnum_MediaBundleError = 20
	// Google Web Designer not created for "Google Ads" environment.
	MediaBundleErrorEnum_UNSUPPORTED_GOOGLE_WEB_DESIGNER_ENVIRONMENT MediaBundleErrorEnum_MediaBundleError = 21
	// Unsupported HTML5 feature in HTML5 asset.
	MediaBundleErrorEnum_UNSUPPORTED_HTML5_FEATURE MediaBundleErrorEnum_MediaBundleError = 22
	// URL in HTML5 entry is not ssl compliant.
	MediaBundleErrorEnum_URL_IN_MEDIA_BUNDLE_NOT_SSL_COMPLIANT MediaBundleErrorEnum_MediaBundleError = 23
	// Custom exits not allowed in HTML5 entry.
	MediaBundleErrorEnum_CUSTOM_EXIT_NOT_ALLOWED MediaBundleErrorEnum_MediaBundleError = 24
)

// Enum value maps for MediaBundleErrorEnum_MediaBundleError.
var (
	MediaBundleErrorEnum_MediaBundleError_name = map[int32]string{
		0:  "UNSPECIFIED",
		1:  "UNKNOWN",
		3:  "BAD_REQUEST",
		4:  "DOUBLECLICK_BUNDLE_NOT_ALLOWED",
		5:  "EXTERNAL_URL_NOT_ALLOWED",
		6:  "FILE_TOO_LARGE",
		7:  "GOOGLE_WEB_DESIGNER_ZIP_FILE_NOT_PUBLISHED",
		8:  "INVALID_INPUT",
		9:  "INVALID_MEDIA_BUNDLE",
		10: "INVALID_MEDIA_BUNDLE_ENTRY",
		11: "INVALID_MIME_TYPE",
		12: "INVALID_PATH",
		13: "INVALID_URL_REFERENCE",
		14: "MEDIA_DATA_TOO_LARGE",
		15: "MISSING_PRIMARY_MEDIA_BUNDLE_ENTRY",
		16: "SERVER_ERROR",
		17: "STORAGE_ERROR",
		18: "SWIFFY_BUNDLE_NOT_ALLOWED",
		19: "TOO_MANY_FILES",
		20: "UNEXPECTED_SIZE",
		21: "UNSUPPORTED_GOOGLE_WEB_DESIGNER_ENVIRONMENT",
		22: "UNSUPPORTED_HTML5_FEATURE",
		23: "URL_IN_MEDIA_BUNDLE_NOT_SSL_COMPLIANT",
		24: "CUSTOM_EXIT_NOT_ALLOWED",
	}
	MediaBundleErrorEnum_MediaBundleError_value = map[string]int32{
		"UNSPECIFIED":                    0,
		"UNKNOWN":                        1,
		"BAD_REQUEST":                    3,
		"DOUBLECLICK_BUNDLE_NOT_ALLOWED": 4,
		"EXTERNAL_URL_NOT_ALLOWED":       5,
		"FILE_TOO_LARGE":                 6,
		"GOOGLE_WEB_DESIGNER_ZIP_FILE_NOT_PUBLISHED": 7,
		"INVALID_INPUT":                               8,
		"INVALID_MEDIA_BUNDLE":                        9,
		"INVALID_MEDIA_BUNDLE_ENTRY":                  10,
		"INVALID_MIME_TYPE":                           11,
		"INVALID_PATH":                                12,
		"INVALID_URL_REFERENCE":                       13,
		"MEDIA_DATA_TOO_LARGE":                        14,
		"MISSING_PRIMARY_MEDIA_BUNDLE_ENTRY":          15,
		"SERVER_ERROR":                                16,
		"STORAGE_ERROR":                               17,
		"SWIFFY_BUNDLE_NOT_ALLOWED":                   18,
		"TOO_MANY_FILES":                              19,
		"UNEXPECTED_SIZE":                             20,
		"UNSUPPORTED_GOOGLE_WEB_DESIGNER_ENVIRONMENT": 21,
		"UNSUPPORTED_HTML5_FEATURE":                   22,
		"URL_IN_MEDIA_BUNDLE_NOT_SSL_COMPLIANT":       23,
		"CUSTOM_EXIT_NOT_ALLOWED":                     24,
	}
)

func (x MediaBundleErrorEnum_MediaBundleError) Enum() *MediaBundleErrorEnum_MediaBundleError {
	p := new(MediaBundleErrorEnum_MediaBundleError)
	*p = x
	return p
}

func (x MediaBundleErrorEnum_MediaBundleError) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (MediaBundleErrorEnum_MediaBundleError) Descriptor() protoreflect.EnumDescriptor {
	return file_errors_media_bundle_error_proto_enumTypes[0].Descriptor()
}

func (MediaBundleErrorEnum_MediaBundleError) Type() protoreflect.EnumType {
	return &file_errors_media_bundle_error_proto_enumTypes[0]
}

func (x MediaBundleErrorEnum_MediaBundleError) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use MediaBundleErrorEnum_MediaBundleError.Descriptor instead.
func (MediaBundleErrorEnum_MediaBundleError) EnumDescriptor() ([]byte, []int) {
	return file_errors_media_bundle_error_proto_rawDescGZIP(), []int{0, 0}
}

// Container for enum describing possible media bundle errors.
type MediaBundleErrorEnum struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *MediaBundleErrorEnum) Reset() {
	*x = MediaBundleErrorEnum{}
	if protoimpl.UnsafeEnabled {
		mi := &file_errors_media_bundle_error_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MediaBundleErrorEnum) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MediaBundleErrorEnum) ProtoMessage() {}

func (x *MediaBundleErrorEnum) ProtoReflect() protoreflect.Message {
	mi := &file_errors_media_bundle_error_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MediaBundleErrorEnum.ProtoReflect.Descriptor instead.
func (*MediaBundleErrorEnum) Descriptor() ([]byte, []int) {
	return file_errors_media_bundle_error_proto_rawDescGZIP(), []int{0}
}

var File_errors_media_bundle_error_proto protoreflect.FileDescriptor

var file_errors_media_bundle_error_proto_rawDesc = []byte{
	0x0a, 0x37, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x64, 0x73, 0x2f, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x61, 0x64, 0x73, 0x2f, 0x76, 0x39, 0x2f, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x73,
	0x2f, 0x6d, 0x65, 0x64, 0x69, 0x61, 0x5f, 0x62, 0x75, 0x6e, 0x64, 0x6c, 0x65, 0x5f, 0x65, 0x72,
	0x72, 0x6f, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x1e, 0x67, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x2e, 0x61, 0x64, 0x73, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61, 0x64, 0x73, 0x2e,
	0x76, 0x39, 0x2e, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x73, 0x1a, 0x1c, 0x67, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xb8, 0x05, 0x0a, 0x14, 0x4d, 0x65, 0x64, 0x69,
	0x61, 0x42, 0x75, 0x6e, 0x64, 0x6c, 0x65, 0x45, 0x72, 0x72, 0x6f, 0x72, 0x45, 0x6e, 0x75, 0x6d,
	0x22, 0x9f, 0x05, 0x0a, 0x10, 0x4d, 0x65, 0x64, 0x69, 0x61, 0x42, 0x75, 0x6e, 0x64, 0x6c, 0x65,
	0x45, 0x72, 0x72, 0x6f, 0x72, 0x12, 0x0f, 0x0a, 0x0b, 0x55, 0x4e, 0x53, 0x50, 0x45, 0x43, 0x49,
	0x46, 0x49, 0x45, 0x44, 0x10, 0x00, 0x12, 0x0b, 0x0a, 0x07, 0x55, 0x4e, 0x4b, 0x4e, 0x4f, 0x57,
	0x4e, 0x10, 0x01, 0x12, 0x0f, 0x0a, 0x0b, 0x42, 0x41, 0x44, 0x5f, 0x52, 0x45, 0x51, 0x55, 0x45,
	0x53, 0x54, 0x10, 0x03, 0x12, 0x22, 0x0a, 0x1e, 0x44, 0x4f, 0x55, 0x42, 0x4c, 0x45, 0x43, 0x4c,
	0x49, 0x43, 0x4b, 0x5f, 0x42, 0x55, 0x4e, 0x44, 0x4c, 0x45, 0x5f, 0x4e, 0x4f, 0x54, 0x5f, 0x41,
	0x4c, 0x4c, 0x4f, 0x57, 0x45, 0x44, 0x10, 0x04, 0x12, 0x1c, 0x0a, 0x18, 0x45, 0x58, 0x54, 0x45,
	0x52, 0x4e, 0x41, 0x4c, 0x5f, 0x55, 0x52, 0x4c, 0x5f, 0x4e, 0x4f, 0x54, 0x5f, 0x41, 0x4c, 0x4c,
	0x4f, 0x57, 0x45, 0x44, 0x10, 0x05, 0x12, 0x12, 0x0a, 0x0e, 0x46, 0x49, 0x4c, 0x45, 0x5f, 0x54,
	0x4f, 0x4f, 0x5f, 0x4c, 0x41, 0x52, 0x47, 0x45, 0x10, 0x06, 0x12, 0x2e, 0x0a, 0x2a, 0x47, 0x4f,
	0x4f, 0x47, 0x4c, 0x45, 0x5f, 0x57, 0x45, 0x42, 0x5f, 0x44, 0x45, 0x53, 0x49, 0x47, 0x4e, 0x45,
	0x52, 0x5f, 0x5a, 0x49, 0x50, 0x5f, 0x46, 0x49, 0x4c, 0x45, 0x5f, 0x4e, 0x4f, 0x54, 0x5f, 0x50,
	0x55, 0x42, 0x4c, 0x49, 0x53, 0x48, 0x45, 0x44, 0x10, 0x07, 0x12, 0x11, 0x0a, 0x0d, 0x49, 0x4e,
	0x56, 0x41, 0x4c, 0x49, 0x44, 0x5f, 0x49, 0x4e, 0x50, 0x55, 0x54, 0x10, 0x08, 0x12, 0x18, 0x0a,
	0x14, 0x49, 0x4e, 0x56, 0x41, 0x4c, 0x49, 0x44, 0x5f, 0x4d, 0x45, 0x44, 0x49, 0x41, 0x5f, 0x42,
	0x55, 0x4e, 0x44, 0x4c, 0x45, 0x10, 0x09, 0x12, 0x1e, 0x0a, 0x1a, 0x49, 0x4e, 0x56, 0x41, 0x4c,
	0x49, 0x44, 0x5f, 0x4d, 0x45, 0x44, 0x49, 0x41, 0x5f, 0x42, 0x55, 0x4e, 0x44, 0x4c, 0x45, 0x5f,
	0x45, 0x4e, 0x54, 0x52, 0x59, 0x10, 0x0a, 0x12, 0x15, 0x0a, 0x11, 0x49, 0x4e, 0x56, 0x41, 0x4c,
	0x49, 0x44, 0x5f, 0x4d, 0x49, 0x4d, 0x45, 0x5f, 0x54, 0x59, 0x50, 0x45, 0x10, 0x0b, 0x12, 0x10,
	0x0a, 0x0c, 0x49, 0x4e, 0x56, 0x41, 0x4c, 0x49, 0x44, 0x5f, 0x50, 0x41, 0x54, 0x48, 0x10, 0x0c,
	0x12, 0x19, 0x0a, 0x15, 0x49, 0x4e, 0x56, 0x41, 0x4c, 0x49, 0x44, 0x5f, 0x55, 0x52, 0x4c, 0x5f,
	0x52, 0x45, 0x46, 0x45, 0x52, 0x45, 0x4e, 0x43, 0x45, 0x10, 0x0d, 0x12, 0x18, 0x0a, 0x14, 0x4d,
	0x45, 0x44, 0x49, 0x41, 0x5f, 0x44, 0x41, 0x54, 0x41, 0x5f, 0x54, 0x4f, 0x4f, 0x5f, 0x4c, 0x41,
	0x52, 0x47, 0x45, 0x10, 0x0e, 0x12, 0x26, 0x0a, 0x22, 0x4d, 0x49, 0x53, 0x53, 0x49, 0x4e, 0x47,
	0x5f, 0x50, 0x52, 0x49, 0x4d, 0x41, 0x52, 0x59, 0x5f, 0x4d, 0x45, 0x44, 0x49, 0x41, 0x5f, 0x42,
	0x55, 0x4e, 0x44, 0x4c, 0x45, 0x5f, 0x45, 0x4e, 0x54, 0x52, 0x59, 0x10, 0x0f, 0x12, 0x10, 0x0a,
	0x0c, 0x53, 0x45, 0x52, 0x56, 0x45, 0x52, 0x5f, 0x45, 0x52, 0x52, 0x4f, 0x52, 0x10, 0x10, 0x12,
	0x11, 0x0a, 0x0d, 0x53, 0x54, 0x4f, 0x52, 0x41, 0x47, 0x45, 0x5f, 0x45, 0x52, 0x52, 0x4f, 0x52,
	0x10, 0x11, 0x12, 0x1d, 0x0a, 0x19, 0x53, 0x57, 0x49, 0x46, 0x46, 0x59, 0x5f, 0x42, 0x55, 0x4e,
	0x44, 0x4c, 0x45, 0x5f, 0x4e, 0x4f, 0x54, 0x5f, 0x41, 0x4c, 0x4c, 0x4f, 0x57, 0x45, 0x44, 0x10,
	0x12, 0x12, 0x12, 0x0a, 0x0e, 0x54, 0x4f, 0x4f, 0x5f, 0x4d, 0x41, 0x4e, 0x59, 0x5f, 0x46, 0x49,
	0x4c, 0x45, 0x53, 0x10, 0x13, 0x12, 0x13, 0x0a, 0x0f, 0x55, 0x4e, 0x45, 0x58, 0x50, 0x45, 0x43,
	0x54, 0x45, 0x44, 0x5f, 0x53, 0x49, 0x5a, 0x45, 0x10, 0x14, 0x12, 0x2f, 0x0a, 0x2b, 0x55, 0x4e,
	0x53, 0x55, 0x50, 0x50, 0x4f, 0x52, 0x54, 0x45, 0x44, 0x5f, 0x47, 0x4f, 0x4f, 0x47, 0x4c, 0x45,
	0x5f, 0x57, 0x45, 0x42, 0x5f, 0x44, 0x45, 0x53, 0x49, 0x47, 0x4e, 0x45, 0x52, 0x5f, 0x45, 0x4e,
	0x56, 0x49, 0x52, 0x4f, 0x4e, 0x4d, 0x45, 0x4e, 0x54, 0x10, 0x15, 0x12, 0x1d, 0x0a, 0x19, 0x55,
	0x4e, 0x53, 0x55, 0x50, 0x50, 0x4f, 0x52, 0x54, 0x45, 0x44, 0x5f, 0x48, 0x54, 0x4d, 0x4c, 0x35,
	0x5f, 0x46, 0x45, 0x41, 0x54, 0x55, 0x52, 0x45, 0x10, 0x16, 0x12, 0x29, 0x0a, 0x25, 0x55, 0x52,
	0x4c, 0x5f, 0x49, 0x4e, 0x5f, 0x4d, 0x45, 0x44, 0x49, 0x41, 0x5f, 0x42, 0x55, 0x4e, 0x44, 0x4c,
	0x45, 0x5f, 0x4e, 0x4f, 0x54, 0x5f, 0x53, 0x53, 0x4c, 0x5f, 0x43, 0x4f, 0x4d, 0x50, 0x4c, 0x49,
	0x41, 0x4e, 0x54, 0x10, 0x17, 0x12, 0x1b, 0x0a, 0x17, 0x43, 0x55, 0x53, 0x54, 0x4f, 0x4d, 0x5f,
	0x45, 0x58, 0x49, 0x54, 0x5f, 0x4e, 0x4f, 0x54, 0x5f, 0x41, 0x4c, 0x4c, 0x4f, 0x57, 0x45, 0x44,
	0x10, 0x18, 0x42, 0xf0, 0x01, 0x0a, 0x22, 0x63, 0x6f, 0x6d, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x2e, 0x61, 0x64, 0x73, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61, 0x64, 0x73, 0x2e,
	0x76, 0x39, 0x2e, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x73, 0x42, 0x15, 0x4d, 0x65, 0x64, 0x69, 0x61,
	0x42, 0x75, 0x6e, 0x64, 0x6c, 0x65, 0x45, 0x72, 0x72, 0x6f, 0x72, 0x50, 0x72, 0x6f, 0x74, 0x6f,
	0x50, 0x01, 0x5a, 0x44, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x67, 0x6f, 0x6c, 0x61, 0x6e,
	0x67, 0x2e, 0x6f, 0x72, 0x67, 0x2f, 0x67, 0x65, 0x6e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61, 0x70, 0x69, 0x73, 0x2f, 0x61, 0x64, 0x73, 0x2f, 0x67, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x61, 0x64, 0x73, 0x2f, 0x76, 0x39, 0x2f, 0x65, 0x72, 0x72, 0x6f, 0x72,
	0x73, 0x3b, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x73, 0xa2, 0x02, 0x03, 0x47, 0x41, 0x41, 0xaa, 0x02,
	0x1e, 0x47, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x41, 0x64, 0x73, 0x2e, 0x47, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x41, 0x64, 0x73, 0x2e, 0x56, 0x39, 0x2e, 0x45, 0x72, 0x72, 0x6f, 0x72, 0x73, 0xca,
	0x02, 0x1e, 0x47, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x5c, 0x41, 0x64, 0x73, 0x5c, 0x47, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x41, 0x64, 0x73, 0x5c, 0x56, 0x39, 0x5c, 0x45, 0x72, 0x72, 0x6f, 0x72, 0x73,
	0xea, 0x02, 0x22, 0x47, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x3a, 0x3a, 0x41, 0x64, 0x73, 0x3a, 0x3a,
	0x47, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x41, 0x64, 0x73, 0x3a, 0x3a, 0x56, 0x39, 0x3a, 0x3a, 0x45,
	0x72, 0x72, 0x6f, 0x72, 0x73, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_errors_media_bundle_error_proto_rawDescOnce sync.Once
	file_errors_media_bundle_error_proto_rawDescData = file_errors_media_bundle_error_proto_rawDesc
)

func file_errors_media_bundle_error_proto_rawDescGZIP() []byte {
	file_errors_media_bundle_error_proto_rawDescOnce.Do(func() {
		file_errors_media_bundle_error_proto_rawDescData = protoimpl.X.CompressGZIP(file_errors_media_bundle_error_proto_rawDescData)
	})
	return file_errors_media_bundle_error_proto_rawDescData
}

var file_errors_media_bundle_error_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_errors_media_bundle_error_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_errors_media_bundle_error_proto_goTypes = []interface{}{
	(MediaBundleErrorEnum_MediaBundleError)(0), // 0: google.ads.googleads.v9.errors.MediaBundleErrorEnum.MediaBundleError
	(*MediaBundleErrorEnum)(nil),               // 1: google.ads.googleads.v9.errors.MediaBundleErrorEnum
}
var file_errors_media_bundle_error_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_errors_media_bundle_error_proto_init() }
func file_errors_media_bundle_error_proto_init() {
	if File_errors_media_bundle_error_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_errors_media_bundle_error_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*MediaBundleErrorEnum); i {
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
			RawDescriptor: file_errors_media_bundle_error_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_errors_media_bundle_error_proto_goTypes,
		DependencyIndexes: file_errors_media_bundle_error_proto_depIdxs,
		EnumInfos:         file_errors_media_bundle_error_proto_enumTypes,
		MessageInfos:      file_errors_media_bundle_error_proto_msgTypes,
	}.Build()
	File_errors_media_bundle_error_proto = out.File
	file_errors_media_bundle_error_proto_rawDesc = nil
	file_errors_media_bundle_error_proto_goTypes = nil
	file_errors_media_bundle_error_proto_depIdxs = nil
}
