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
// 	protoc-gen-go v1.26.0
// 	protoc        v3.21.12
// source: google/ads/googleads/v14/errors/url_field_error.proto

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

// Enum describing possible url field errors.
type UrlFieldErrorEnum_UrlFieldError int32

const (
	// Enum unspecified.
	UrlFieldErrorEnum_UNSPECIFIED UrlFieldErrorEnum_UrlFieldError = 0
	// The received error code is not known in this version.
	UrlFieldErrorEnum_UNKNOWN UrlFieldErrorEnum_UrlFieldError = 1
	// The tracking url template is invalid.
	UrlFieldErrorEnum_INVALID_TRACKING_URL_TEMPLATE UrlFieldErrorEnum_UrlFieldError = 2
	// The tracking url template contains invalid tag.
	UrlFieldErrorEnum_INVALID_TAG_IN_TRACKING_URL_TEMPLATE UrlFieldErrorEnum_UrlFieldError = 3
	// The tracking url template must contain at least one tag (for example,
	// {lpurl}), This applies only to tracking url template associated with
	// website ads or product ads.
	UrlFieldErrorEnum_MISSING_TRACKING_URL_TEMPLATE_TAG UrlFieldErrorEnum_UrlFieldError = 4
	// The tracking url template must start with a valid protocol (or lpurl
	// tag).
	UrlFieldErrorEnum_MISSING_PROTOCOL_IN_TRACKING_URL_TEMPLATE UrlFieldErrorEnum_UrlFieldError = 5
	// The tracking url template starts with an invalid protocol.
	UrlFieldErrorEnum_INVALID_PROTOCOL_IN_TRACKING_URL_TEMPLATE UrlFieldErrorEnum_UrlFieldError = 6
	// The tracking url template contains illegal characters.
	UrlFieldErrorEnum_MALFORMED_TRACKING_URL_TEMPLATE UrlFieldErrorEnum_UrlFieldError = 7
	// The tracking url template must contain a host name (or lpurl tag).
	UrlFieldErrorEnum_MISSING_HOST_IN_TRACKING_URL_TEMPLATE UrlFieldErrorEnum_UrlFieldError = 8
	// The tracking url template has an invalid or missing top level domain
	// extension.
	UrlFieldErrorEnum_INVALID_TLD_IN_TRACKING_URL_TEMPLATE UrlFieldErrorEnum_UrlFieldError = 9
	// The tracking url template contains nested occurrences of the same
	// conditional tag (for example, {ifmobile:{ifmobile:x}}).
	UrlFieldErrorEnum_REDUNDANT_NESTED_TRACKING_URL_TEMPLATE_TAG UrlFieldErrorEnum_UrlFieldError = 10
	// The final url is invalid.
	UrlFieldErrorEnum_INVALID_FINAL_URL UrlFieldErrorEnum_UrlFieldError = 11
	// The final url contains invalid tag.
	UrlFieldErrorEnum_INVALID_TAG_IN_FINAL_URL UrlFieldErrorEnum_UrlFieldError = 12
	// The final url contains nested occurrences of the same conditional tag
	// (for example, {ifmobile:{ifmobile:x}}).
	UrlFieldErrorEnum_REDUNDANT_NESTED_FINAL_URL_TAG UrlFieldErrorEnum_UrlFieldError = 13
	// The final url must start with a valid protocol.
	UrlFieldErrorEnum_MISSING_PROTOCOL_IN_FINAL_URL UrlFieldErrorEnum_UrlFieldError = 14
	// The final url starts with an invalid protocol.
	UrlFieldErrorEnum_INVALID_PROTOCOL_IN_FINAL_URL UrlFieldErrorEnum_UrlFieldError = 15
	// The final url contains illegal characters.
	UrlFieldErrorEnum_MALFORMED_FINAL_URL UrlFieldErrorEnum_UrlFieldError = 16
	// The final url must contain a host name.
	UrlFieldErrorEnum_MISSING_HOST_IN_FINAL_URL UrlFieldErrorEnum_UrlFieldError = 17
	// The tracking url template has an invalid or missing top level domain
	// extension.
	UrlFieldErrorEnum_INVALID_TLD_IN_FINAL_URL UrlFieldErrorEnum_UrlFieldError = 18
	// The final mobile url is invalid.
	UrlFieldErrorEnum_INVALID_FINAL_MOBILE_URL UrlFieldErrorEnum_UrlFieldError = 19
	// The final mobile url contains invalid tag.
	UrlFieldErrorEnum_INVALID_TAG_IN_FINAL_MOBILE_URL UrlFieldErrorEnum_UrlFieldError = 20
	// The final mobile url contains nested occurrences of the same conditional
	// tag (for example, {ifmobile:{ifmobile:x}}).
	UrlFieldErrorEnum_REDUNDANT_NESTED_FINAL_MOBILE_URL_TAG UrlFieldErrorEnum_UrlFieldError = 21
	// The final mobile url must start with a valid protocol.
	UrlFieldErrorEnum_MISSING_PROTOCOL_IN_FINAL_MOBILE_URL UrlFieldErrorEnum_UrlFieldError = 22
	// The final mobile url starts with an invalid protocol.
	UrlFieldErrorEnum_INVALID_PROTOCOL_IN_FINAL_MOBILE_URL UrlFieldErrorEnum_UrlFieldError = 23
	// The final mobile url contains illegal characters.
	UrlFieldErrorEnum_MALFORMED_FINAL_MOBILE_URL UrlFieldErrorEnum_UrlFieldError = 24
	// The final mobile url must contain a host name.
	UrlFieldErrorEnum_MISSING_HOST_IN_FINAL_MOBILE_URL UrlFieldErrorEnum_UrlFieldError = 25
	// The tracking url template has an invalid or missing top level domain
	// extension.
	UrlFieldErrorEnum_INVALID_TLD_IN_FINAL_MOBILE_URL UrlFieldErrorEnum_UrlFieldError = 26
	// The final app url is invalid.
	UrlFieldErrorEnum_INVALID_FINAL_APP_URL UrlFieldErrorEnum_UrlFieldError = 27
	// The final app url contains invalid tag.
	UrlFieldErrorEnum_INVALID_TAG_IN_FINAL_APP_URL UrlFieldErrorEnum_UrlFieldError = 28
	// The final app url contains nested occurrences of the same conditional tag
	// (for example, {ifmobile:{ifmobile:x}}).
	UrlFieldErrorEnum_REDUNDANT_NESTED_FINAL_APP_URL_TAG UrlFieldErrorEnum_UrlFieldError = 29
	// More than one app url found for the same OS type.
	UrlFieldErrorEnum_MULTIPLE_APP_URLS_FOR_OSTYPE UrlFieldErrorEnum_UrlFieldError = 30
	// The OS type given for an app url is not valid.
	UrlFieldErrorEnum_INVALID_OSTYPE UrlFieldErrorEnum_UrlFieldError = 31
	// The protocol given for an app url is not valid. (For example,
	// "android-app://")
	UrlFieldErrorEnum_INVALID_PROTOCOL_FOR_APP_URL UrlFieldErrorEnum_UrlFieldError = 32
	// The package id (app id) given for an app url is not valid.
	UrlFieldErrorEnum_INVALID_PACKAGE_ID_FOR_APP_URL UrlFieldErrorEnum_UrlFieldError = 33
	// The number of url custom parameters for an resource exceeds the maximum
	// limit allowed.
	UrlFieldErrorEnum_URL_CUSTOM_PARAMETERS_COUNT_EXCEEDS_LIMIT UrlFieldErrorEnum_UrlFieldError = 34
	// An invalid character appears in the parameter key.
	UrlFieldErrorEnum_INVALID_CHARACTERS_IN_URL_CUSTOM_PARAMETER_KEY UrlFieldErrorEnum_UrlFieldError = 39
	// An invalid character appears in the parameter value.
	UrlFieldErrorEnum_INVALID_CHARACTERS_IN_URL_CUSTOM_PARAMETER_VALUE UrlFieldErrorEnum_UrlFieldError = 40
	// The url custom parameter value fails url tag validation.
	UrlFieldErrorEnum_INVALID_TAG_IN_URL_CUSTOM_PARAMETER_VALUE UrlFieldErrorEnum_UrlFieldError = 41
	// The custom parameter contains nested occurrences of the same conditional
	// tag (for example, {ifmobile:{ifmobile:x}}).
	UrlFieldErrorEnum_REDUNDANT_NESTED_URL_CUSTOM_PARAMETER_TAG UrlFieldErrorEnum_UrlFieldError = 42
	// The protocol (http:// or https://) is missing.
	UrlFieldErrorEnum_MISSING_PROTOCOL UrlFieldErrorEnum_UrlFieldError = 43
	// Unsupported protocol in URL. Only http and https are supported.
	UrlFieldErrorEnum_INVALID_PROTOCOL UrlFieldErrorEnum_UrlFieldError = 52
	// The url is invalid.
	UrlFieldErrorEnum_INVALID_URL UrlFieldErrorEnum_UrlFieldError = 44
	// Destination Url is deprecated.
	UrlFieldErrorEnum_DESTINATION_URL_DEPRECATED UrlFieldErrorEnum_UrlFieldError = 45
	// The url contains invalid tag.
	UrlFieldErrorEnum_INVALID_TAG_IN_URL UrlFieldErrorEnum_UrlFieldError = 46
	// The url must contain at least one tag (for example, {lpurl}).
	UrlFieldErrorEnum_MISSING_URL_TAG UrlFieldErrorEnum_UrlFieldError = 47
	// Duplicate url id.
	UrlFieldErrorEnum_DUPLICATE_URL_ID UrlFieldErrorEnum_UrlFieldError = 48
	// Invalid url id.
	UrlFieldErrorEnum_INVALID_URL_ID UrlFieldErrorEnum_UrlFieldError = 49
	// The final url suffix cannot begin with '?' or '&' characters and must be
	// a valid query string.
	UrlFieldErrorEnum_FINAL_URL_SUFFIX_MALFORMED UrlFieldErrorEnum_UrlFieldError = 50
	// The final url suffix cannot contain {lpurl} related or {ignore} tags.
	UrlFieldErrorEnum_INVALID_TAG_IN_FINAL_URL_SUFFIX UrlFieldErrorEnum_UrlFieldError = 51
	// The top level domain is invalid, for example, not a public top level
	// domain listed in publicsuffix.org.
	UrlFieldErrorEnum_INVALID_TOP_LEVEL_DOMAIN UrlFieldErrorEnum_UrlFieldError = 53
	// Malformed top level domain in URL.
	UrlFieldErrorEnum_MALFORMED_TOP_LEVEL_DOMAIN UrlFieldErrorEnum_UrlFieldError = 54
	// Malformed URL.
	UrlFieldErrorEnum_MALFORMED_URL UrlFieldErrorEnum_UrlFieldError = 55
	// No host found in URL.
	UrlFieldErrorEnum_MISSING_HOST UrlFieldErrorEnum_UrlFieldError = 56
	// Custom parameter value cannot be null.
	UrlFieldErrorEnum_NULL_CUSTOM_PARAMETER_VALUE UrlFieldErrorEnum_UrlFieldError = 57
	// Track parameter is not supported.
	UrlFieldErrorEnum_VALUE_TRACK_PARAMETER_NOT_SUPPORTED UrlFieldErrorEnum_UrlFieldError = 58
)

// Enum value maps for UrlFieldErrorEnum_UrlFieldError.
var (
	UrlFieldErrorEnum_UrlFieldError_name = map[int32]string{
		0:  "UNSPECIFIED",
		1:  "UNKNOWN",
		2:  "INVALID_TRACKING_URL_TEMPLATE",
		3:  "INVALID_TAG_IN_TRACKING_URL_TEMPLATE",
		4:  "MISSING_TRACKING_URL_TEMPLATE_TAG",
		5:  "MISSING_PROTOCOL_IN_TRACKING_URL_TEMPLATE",
		6:  "INVALID_PROTOCOL_IN_TRACKING_URL_TEMPLATE",
		7:  "MALFORMED_TRACKING_URL_TEMPLATE",
		8:  "MISSING_HOST_IN_TRACKING_URL_TEMPLATE",
		9:  "INVALID_TLD_IN_TRACKING_URL_TEMPLATE",
		10: "REDUNDANT_NESTED_TRACKING_URL_TEMPLATE_TAG",
		11: "INVALID_FINAL_URL",
		12: "INVALID_TAG_IN_FINAL_URL",
		13: "REDUNDANT_NESTED_FINAL_URL_TAG",
		14: "MISSING_PROTOCOL_IN_FINAL_URL",
		15: "INVALID_PROTOCOL_IN_FINAL_URL",
		16: "MALFORMED_FINAL_URL",
		17: "MISSING_HOST_IN_FINAL_URL",
		18: "INVALID_TLD_IN_FINAL_URL",
		19: "INVALID_FINAL_MOBILE_URL",
		20: "INVALID_TAG_IN_FINAL_MOBILE_URL",
		21: "REDUNDANT_NESTED_FINAL_MOBILE_URL_TAG",
		22: "MISSING_PROTOCOL_IN_FINAL_MOBILE_URL",
		23: "INVALID_PROTOCOL_IN_FINAL_MOBILE_URL",
		24: "MALFORMED_FINAL_MOBILE_URL",
		25: "MISSING_HOST_IN_FINAL_MOBILE_URL",
		26: "INVALID_TLD_IN_FINAL_MOBILE_URL",
		27: "INVALID_FINAL_APP_URL",
		28: "INVALID_TAG_IN_FINAL_APP_URL",
		29: "REDUNDANT_NESTED_FINAL_APP_URL_TAG",
		30: "MULTIPLE_APP_URLS_FOR_OSTYPE",
		31: "INVALID_OSTYPE",
		32: "INVALID_PROTOCOL_FOR_APP_URL",
		33: "INVALID_PACKAGE_ID_FOR_APP_URL",
		34: "URL_CUSTOM_PARAMETERS_COUNT_EXCEEDS_LIMIT",
		39: "INVALID_CHARACTERS_IN_URL_CUSTOM_PARAMETER_KEY",
		40: "INVALID_CHARACTERS_IN_URL_CUSTOM_PARAMETER_VALUE",
		41: "INVALID_TAG_IN_URL_CUSTOM_PARAMETER_VALUE",
		42: "REDUNDANT_NESTED_URL_CUSTOM_PARAMETER_TAG",
		43: "MISSING_PROTOCOL",
		52: "INVALID_PROTOCOL",
		44: "INVALID_URL",
		45: "DESTINATION_URL_DEPRECATED",
		46: "INVALID_TAG_IN_URL",
		47: "MISSING_URL_TAG",
		48: "DUPLICATE_URL_ID",
		49: "INVALID_URL_ID",
		50: "FINAL_URL_SUFFIX_MALFORMED",
		51: "INVALID_TAG_IN_FINAL_URL_SUFFIX",
		53: "INVALID_TOP_LEVEL_DOMAIN",
		54: "MALFORMED_TOP_LEVEL_DOMAIN",
		55: "MALFORMED_URL",
		56: "MISSING_HOST",
		57: "NULL_CUSTOM_PARAMETER_VALUE",
		58: "VALUE_TRACK_PARAMETER_NOT_SUPPORTED",
	}
	UrlFieldErrorEnum_UrlFieldError_value = map[string]int32{
		"UNSPECIFIED":                                      0,
		"UNKNOWN":                                          1,
		"INVALID_TRACKING_URL_TEMPLATE":                    2,
		"INVALID_TAG_IN_TRACKING_URL_TEMPLATE":             3,
		"MISSING_TRACKING_URL_TEMPLATE_TAG":                4,
		"MISSING_PROTOCOL_IN_TRACKING_URL_TEMPLATE":        5,
		"INVALID_PROTOCOL_IN_TRACKING_URL_TEMPLATE":        6,
		"MALFORMED_TRACKING_URL_TEMPLATE":                  7,
		"MISSING_HOST_IN_TRACKING_URL_TEMPLATE":            8,
		"INVALID_TLD_IN_TRACKING_URL_TEMPLATE":             9,
		"REDUNDANT_NESTED_TRACKING_URL_TEMPLATE_TAG":       10,
		"INVALID_FINAL_URL":                                11,
		"INVALID_TAG_IN_FINAL_URL":                         12,
		"REDUNDANT_NESTED_FINAL_URL_TAG":                   13,
		"MISSING_PROTOCOL_IN_FINAL_URL":                    14,
		"INVALID_PROTOCOL_IN_FINAL_URL":                    15,
		"MALFORMED_FINAL_URL":                              16,
		"MISSING_HOST_IN_FINAL_URL":                        17,
		"INVALID_TLD_IN_FINAL_URL":                         18,
		"INVALID_FINAL_MOBILE_URL":                         19,
		"INVALID_TAG_IN_FINAL_MOBILE_URL":                  20,
		"REDUNDANT_NESTED_FINAL_MOBILE_URL_TAG":            21,
		"MISSING_PROTOCOL_IN_FINAL_MOBILE_URL":             22,
		"INVALID_PROTOCOL_IN_FINAL_MOBILE_URL":             23,
		"MALFORMED_FINAL_MOBILE_URL":                       24,
		"MISSING_HOST_IN_FINAL_MOBILE_URL":                 25,
		"INVALID_TLD_IN_FINAL_MOBILE_URL":                  26,
		"INVALID_FINAL_APP_URL":                            27,
		"INVALID_TAG_IN_FINAL_APP_URL":                     28,
		"REDUNDANT_NESTED_FINAL_APP_URL_TAG":               29,
		"MULTIPLE_APP_URLS_FOR_OSTYPE":                     30,
		"INVALID_OSTYPE":                                   31,
		"INVALID_PROTOCOL_FOR_APP_URL":                     32,
		"INVALID_PACKAGE_ID_FOR_APP_URL":                   33,
		"URL_CUSTOM_PARAMETERS_COUNT_EXCEEDS_LIMIT":        34,
		"INVALID_CHARACTERS_IN_URL_CUSTOM_PARAMETER_KEY":   39,
		"INVALID_CHARACTERS_IN_URL_CUSTOM_PARAMETER_VALUE": 40,
		"INVALID_TAG_IN_URL_CUSTOM_PARAMETER_VALUE":        41,
		"REDUNDANT_NESTED_URL_CUSTOM_PARAMETER_TAG":        42,
		"MISSING_PROTOCOL":                                 43,
		"INVALID_PROTOCOL":                                 52,
		"INVALID_URL":                                      44,
		"DESTINATION_URL_DEPRECATED":                       45,
		"INVALID_TAG_IN_URL":                               46,
		"MISSING_URL_TAG":                                  47,
		"DUPLICATE_URL_ID":                                 48,
		"INVALID_URL_ID":                                   49,
		"FINAL_URL_SUFFIX_MALFORMED":                       50,
		"INVALID_TAG_IN_FINAL_URL_SUFFIX":                  51,
		"INVALID_TOP_LEVEL_DOMAIN":                         53,
		"MALFORMED_TOP_LEVEL_DOMAIN":                       54,
		"MALFORMED_URL":                                    55,
		"MISSING_HOST":                                     56,
		"NULL_CUSTOM_PARAMETER_VALUE":                      57,
		"VALUE_TRACK_PARAMETER_NOT_SUPPORTED":              58,
	}
)

func (x UrlFieldErrorEnum_UrlFieldError) Enum() *UrlFieldErrorEnum_UrlFieldError {
	p := new(UrlFieldErrorEnum_UrlFieldError)
	*p = x
	return p
}

func (x UrlFieldErrorEnum_UrlFieldError) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (UrlFieldErrorEnum_UrlFieldError) Descriptor() protoreflect.EnumDescriptor {
	return file_errors_url_field_error_proto_enumTypes[0].Descriptor()
}

func (UrlFieldErrorEnum_UrlFieldError) Type() protoreflect.EnumType {
	return &file_errors_url_field_error_proto_enumTypes[0]
}

func (x UrlFieldErrorEnum_UrlFieldError) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use UrlFieldErrorEnum_UrlFieldError.Descriptor instead.
func (UrlFieldErrorEnum_UrlFieldError) EnumDescriptor() ([]byte, []int) {
	return file_errors_url_field_error_proto_rawDescGZIP(), []int{0, 0}
}

// Container for enum describing possible url field errors.
type UrlFieldErrorEnum struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *UrlFieldErrorEnum) Reset() {
	*x = UrlFieldErrorEnum{}
	if protoimpl.UnsafeEnabled {
		mi := &file_errors_url_field_error_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UrlFieldErrorEnum) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UrlFieldErrorEnum) ProtoMessage() {}

func (x *UrlFieldErrorEnum) ProtoReflect() protoreflect.Message {
	mi := &file_errors_url_field_error_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UrlFieldErrorEnum.ProtoReflect.Descriptor instead.
func (*UrlFieldErrorEnum) Descriptor() ([]byte, []int) {
	return file_errors_url_field_error_proto_rawDescGZIP(), []int{0}
}

var File_errors_url_field_error_proto protoreflect.FileDescriptor

var file_errors_url_field_error_proto_rawDesc = []byte{
	0x0a, 0x35, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x64, 0x73, 0x2f, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x61, 0x64, 0x73, 0x2f, 0x76, 0x31, 0x34, 0x2f, 0x65, 0x72, 0x72, 0x6f, 0x72,
	0x73, 0x2f, 0x75, 0x72, 0x6c, 0x5f, 0x66, 0x69, 0x65, 0x6c, 0x64, 0x5f, 0x65, 0x72, 0x72, 0x6f,
	0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e,
	0x61, 0x64, 0x73, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61, 0x64, 0x73, 0x2e, 0x76, 0x31,
	0x34, 0x2e, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x73, 0x22, 0xdf, 0x0e, 0x0a, 0x11, 0x55, 0x72, 0x6c,
	0x46, 0x69, 0x65, 0x6c, 0x64, 0x45, 0x72, 0x72, 0x6f, 0x72, 0x45, 0x6e, 0x75, 0x6d, 0x22, 0xc9,
	0x0e, 0x0a, 0x0d, 0x55, 0x72, 0x6c, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x45, 0x72, 0x72, 0x6f, 0x72,
	0x12, 0x0f, 0x0a, 0x0b, 0x55, 0x4e, 0x53, 0x50, 0x45, 0x43, 0x49, 0x46, 0x49, 0x45, 0x44, 0x10,
	0x00, 0x12, 0x0b, 0x0a, 0x07, 0x55, 0x4e, 0x4b, 0x4e, 0x4f, 0x57, 0x4e, 0x10, 0x01, 0x12, 0x21,
	0x0a, 0x1d, 0x49, 0x4e, 0x56, 0x41, 0x4c, 0x49, 0x44, 0x5f, 0x54, 0x52, 0x41, 0x43, 0x4b, 0x49,
	0x4e, 0x47, 0x5f, 0x55, 0x52, 0x4c, 0x5f, 0x54, 0x45, 0x4d, 0x50, 0x4c, 0x41, 0x54, 0x45, 0x10,
	0x02, 0x12, 0x28, 0x0a, 0x24, 0x49, 0x4e, 0x56, 0x41, 0x4c, 0x49, 0x44, 0x5f, 0x54, 0x41, 0x47,
	0x5f, 0x49, 0x4e, 0x5f, 0x54, 0x52, 0x41, 0x43, 0x4b, 0x49, 0x4e, 0x47, 0x5f, 0x55, 0x52, 0x4c,
	0x5f, 0x54, 0x45, 0x4d, 0x50, 0x4c, 0x41, 0x54, 0x45, 0x10, 0x03, 0x12, 0x25, 0x0a, 0x21, 0x4d,
	0x49, 0x53, 0x53, 0x49, 0x4e, 0x47, 0x5f, 0x54, 0x52, 0x41, 0x43, 0x4b, 0x49, 0x4e, 0x47, 0x5f,
	0x55, 0x52, 0x4c, 0x5f, 0x54, 0x45, 0x4d, 0x50, 0x4c, 0x41, 0x54, 0x45, 0x5f, 0x54, 0x41, 0x47,
	0x10, 0x04, 0x12, 0x2d, 0x0a, 0x29, 0x4d, 0x49, 0x53, 0x53, 0x49, 0x4e, 0x47, 0x5f, 0x50, 0x52,
	0x4f, 0x54, 0x4f, 0x43, 0x4f, 0x4c, 0x5f, 0x49, 0x4e, 0x5f, 0x54, 0x52, 0x41, 0x43, 0x4b, 0x49,
	0x4e, 0x47, 0x5f, 0x55, 0x52, 0x4c, 0x5f, 0x54, 0x45, 0x4d, 0x50, 0x4c, 0x41, 0x54, 0x45, 0x10,
	0x05, 0x12, 0x2d, 0x0a, 0x29, 0x49, 0x4e, 0x56, 0x41, 0x4c, 0x49, 0x44, 0x5f, 0x50, 0x52, 0x4f,
	0x54, 0x4f, 0x43, 0x4f, 0x4c, 0x5f, 0x49, 0x4e, 0x5f, 0x54, 0x52, 0x41, 0x43, 0x4b, 0x49, 0x4e,
	0x47, 0x5f, 0x55, 0x52, 0x4c, 0x5f, 0x54, 0x45, 0x4d, 0x50, 0x4c, 0x41, 0x54, 0x45, 0x10, 0x06,
	0x12, 0x23, 0x0a, 0x1f, 0x4d, 0x41, 0x4c, 0x46, 0x4f, 0x52, 0x4d, 0x45, 0x44, 0x5f, 0x54, 0x52,
	0x41, 0x43, 0x4b, 0x49, 0x4e, 0x47, 0x5f, 0x55, 0x52, 0x4c, 0x5f, 0x54, 0x45, 0x4d, 0x50, 0x4c,
	0x41, 0x54, 0x45, 0x10, 0x07, 0x12, 0x29, 0x0a, 0x25, 0x4d, 0x49, 0x53, 0x53, 0x49, 0x4e, 0x47,
	0x5f, 0x48, 0x4f, 0x53, 0x54, 0x5f, 0x49, 0x4e, 0x5f, 0x54, 0x52, 0x41, 0x43, 0x4b, 0x49, 0x4e,
	0x47, 0x5f, 0x55, 0x52, 0x4c, 0x5f, 0x54, 0x45, 0x4d, 0x50, 0x4c, 0x41, 0x54, 0x45, 0x10, 0x08,
	0x12, 0x28, 0x0a, 0x24, 0x49, 0x4e, 0x56, 0x41, 0x4c, 0x49, 0x44, 0x5f, 0x54, 0x4c, 0x44, 0x5f,
	0x49, 0x4e, 0x5f, 0x54, 0x52, 0x41, 0x43, 0x4b, 0x49, 0x4e, 0x47, 0x5f, 0x55, 0x52, 0x4c, 0x5f,
	0x54, 0x45, 0x4d, 0x50, 0x4c, 0x41, 0x54, 0x45, 0x10, 0x09, 0x12, 0x2e, 0x0a, 0x2a, 0x52, 0x45,
	0x44, 0x55, 0x4e, 0x44, 0x41, 0x4e, 0x54, 0x5f, 0x4e, 0x45, 0x53, 0x54, 0x45, 0x44, 0x5f, 0x54,
	0x52, 0x41, 0x43, 0x4b, 0x49, 0x4e, 0x47, 0x5f, 0x55, 0x52, 0x4c, 0x5f, 0x54, 0x45, 0x4d, 0x50,
	0x4c, 0x41, 0x54, 0x45, 0x5f, 0x54, 0x41, 0x47, 0x10, 0x0a, 0x12, 0x15, 0x0a, 0x11, 0x49, 0x4e,
	0x56, 0x41, 0x4c, 0x49, 0x44, 0x5f, 0x46, 0x49, 0x4e, 0x41, 0x4c, 0x5f, 0x55, 0x52, 0x4c, 0x10,
	0x0b, 0x12, 0x1c, 0x0a, 0x18, 0x49, 0x4e, 0x56, 0x41, 0x4c, 0x49, 0x44, 0x5f, 0x54, 0x41, 0x47,
	0x5f, 0x49, 0x4e, 0x5f, 0x46, 0x49, 0x4e, 0x41, 0x4c, 0x5f, 0x55, 0x52, 0x4c, 0x10, 0x0c, 0x12,
	0x22, 0x0a, 0x1e, 0x52, 0x45, 0x44, 0x55, 0x4e, 0x44, 0x41, 0x4e, 0x54, 0x5f, 0x4e, 0x45, 0x53,
	0x54, 0x45, 0x44, 0x5f, 0x46, 0x49, 0x4e, 0x41, 0x4c, 0x5f, 0x55, 0x52, 0x4c, 0x5f, 0x54, 0x41,
	0x47, 0x10, 0x0d, 0x12, 0x21, 0x0a, 0x1d, 0x4d, 0x49, 0x53, 0x53, 0x49, 0x4e, 0x47, 0x5f, 0x50,
	0x52, 0x4f, 0x54, 0x4f, 0x43, 0x4f, 0x4c, 0x5f, 0x49, 0x4e, 0x5f, 0x46, 0x49, 0x4e, 0x41, 0x4c,
	0x5f, 0x55, 0x52, 0x4c, 0x10, 0x0e, 0x12, 0x21, 0x0a, 0x1d, 0x49, 0x4e, 0x56, 0x41, 0x4c, 0x49,
	0x44, 0x5f, 0x50, 0x52, 0x4f, 0x54, 0x4f, 0x43, 0x4f, 0x4c, 0x5f, 0x49, 0x4e, 0x5f, 0x46, 0x49,
	0x4e, 0x41, 0x4c, 0x5f, 0x55, 0x52, 0x4c, 0x10, 0x0f, 0x12, 0x17, 0x0a, 0x13, 0x4d, 0x41, 0x4c,
	0x46, 0x4f, 0x52, 0x4d, 0x45, 0x44, 0x5f, 0x46, 0x49, 0x4e, 0x41, 0x4c, 0x5f, 0x55, 0x52, 0x4c,
	0x10, 0x10, 0x12, 0x1d, 0x0a, 0x19, 0x4d, 0x49, 0x53, 0x53, 0x49, 0x4e, 0x47, 0x5f, 0x48, 0x4f,
	0x53, 0x54, 0x5f, 0x49, 0x4e, 0x5f, 0x46, 0x49, 0x4e, 0x41, 0x4c, 0x5f, 0x55, 0x52, 0x4c, 0x10,
	0x11, 0x12, 0x1c, 0x0a, 0x18, 0x49, 0x4e, 0x56, 0x41, 0x4c, 0x49, 0x44, 0x5f, 0x54, 0x4c, 0x44,
	0x5f, 0x49, 0x4e, 0x5f, 0x46, 0x49, 0x4e, 0x41, 0x4c, 0x5f, 0x55, 0x52, 0x4c, 0x10, 0x12, 0x12,
	0x1c, 0x0a, 0x18, 0x49, 0x4e, 0x56, 0x41, 0x4c, 0x49, 0x44, 0x5f, 0x46, 0x49, 0x4e, 0x41, 0x4c,
	0x5f, 0x4d, 0x4f, 0x42, 0x49, 0x4c, 0x45, 0x5f, 0x55, 0x52, 0x4c, 0x10, 0x13, 0x12, 0x23, 0x0a,
	0x1f, 0x49, 0x4e, 0x56, 0x41, 0x4c, 0x49, 0x44, 0x5f, 0x54, 0x41, 0x47, 0x5f, 0x49, 0x4e, 0x5f,
	0x46, 0x49, 0x4e, 0x41, 0x4c, 0x5f, 0x4d, 0x4f, 0x42, 0x49, 0x4c, 0x45, 0x5f, 0x55, 0x52, 0x4c,
	0x10, 0x14, 0x12, 0x29, 0x0a, 0x25, 0x52, 0x45, 0x44, 0x55, 0x4e, 0x44, 0x41, 0x4e, 0x54, 0x5f,
	0x4e, 0x45, 0x53, 0x54, 0x45, 0x44, 0x5f, 0x46, 0x49, 0x4e, 0x41, 0x4c, 0x5f, 0x4d, 0x4f, 0x42,
	0x49, 0x4c, 0x45, 0x5f, 0x55, 0x52, 0x4c, 0x5f, 0x54, 0x41, 0x47, 0x10, 0x15, 0x12, 0x28, 0x0a,
	0x24, 0x4d, 0x49, 0x53, 0x53, 0x49, 0x4e, 0x47, 0x5f, 0x50, 0x52, 0x4f, 0x54, 0x4f, 0x43, 0x4f,
	0x4c, 0x5f, 0x49, 0x4e, 0x5f, 0x46, 0x49, 0x4e, 0x41, 0x4c, 0x5f, 0x4d, 0x4f, 0x42, 0x49, 0x4c,
	0x45, 0x5f, 0x55, 0x52, 0x4c, 0x10, 0x16, 0x12, 0x28, 0x0a, 0x24, 0x49, 0x4e, 0x56, 0x41, 0x4c,
	0x49, 0x44, 0x5f, 0x50, 0x52, 0x4f, 0x54, 0x4f, 0x43, 0x4f, 0x4c, 0x5f, 0x49, 0x4e, 0x5f, 0x46,
	0x49, 0x4e, 0x41, 0x4c, 0x5f, 0x4d, 0x4f, 0x42, 0x49, 0x4c, 0x45, 0x5f, 0x55, 0x52, 0x4c, 0x10,
	0x17, 0x12, 0x1e, 0x0a, 0x1a, 0x4d, 0x41, 0x4c, 0x46, 0x4f, 0x52, 0x4d, 0x45, 0x44, 0x5f, 0x46,
	0x49, 0x4e, 0x41, 0x4c, 0x5f, 0x4d, 0x4f, 0x42, 0x49, 0x4c, 0x45, 0x5f, 0x55, 0x52, 0x4c, 0x10,
	0x18, 0x12, 0x24, 0x0a, 0x20, 0x4d, 0x49, 0x53, 0x53, 0x49, 0x4e, 0x47, 0x5f, 0x48, 0x4f, 0x53,
	0x54, 0x5f, 0x49, 0x4e, 0x5f, 0x46, 0x49, 0x4e, 0x41, 0x4c, 0x5f, 0x4d, 0x4f, 0x42, 0x49, 0x4c,
	0x45, 0x5f, 0x55, 0x52, 0x4c, 0x10, 0x19, 0x12, 0x23, 0x0a, 0x1f, 0x49, 0x4e, 0x56, 0x41, 0x4c,
	0x49, 0x44, 0x5f, 0x54, 0x4c, 0x44, 0x5f, 0x49, 0x4e, 0x5f, 0x46, 0x49, 0x4e, 0x41, 0x4c, 0x5f,
	0x4d, 0x4f, 0x42, 0x49, 0x4c, 0x45, 0x5f, 0x55, 0x52, 0x4c, 0x10, 0x1a, 0x12, 0x19, 0x0a, 0x15,
	0x49, 0x4e, 0x56, 0x41, 0x4c, 0x49, 0x44, 0x5f, 0x46, 0x49, 0x4e, 0x41, 0x4c, 0x5f, 0x41, 0x50,
	0x50, 0x5f, 0x55, 0x52, 0x4c, 0x10, 0x1b, 0x12, 0x20, 0x0a, 0x1c, 0x49, 0x4e, 0x56, 0x41, 0x4c,
	0x49, 0x44, 0x5f, 0x54, 0x41, 0x47, 0x5f, 0x49, 0x4e, 0x5f, 0x46, 0x49, 0x4e, 0x41, 0x4c, 0x5f,
	0x41, 0x50, 0x50, 0x5f, 0x55, 0x52, 0x4c, 0x10, 0x1c, 0x12, 0x26, 0x0a, 0x22, 0x52, 0x45, 0x44,
	0x55, 0x4e, 0x44, 0x41, 0x4e, 0x54, 0x5f, 0x4e, 0x45, 0x53, 0x54, 0x45, 0x44, 0x5f, 0x46, 0x49,
	0x4e, 0x41, 0x4c, 0x5f, 0x41, 0x50, 0x50, 0x5f, 0x55, 0x52, 0x4c, 0x5f, 0x54, 0x41, 0x47, 0x10,
	0x1d, 0x12, 0x20, 0x0a, 0x1c, 0x4d, 0x55, 0x4c, 0x54, 0x49, 0x50, 0x4c, 0x45, 0x5f, 0x41, 0x50,
	0x50, 0x5f, 0x55, 0x52, 0x4c, 0x53, 0x5f, 0x46, 0x4f, 0x52, 0x5f, 0x4f, 0x53, 0x54, 0x59, 0x50,
	0x45, 0x10, 0x1e, 0x12, 0x12, 0x0a, 0x0e, 0x49, 0x4e, 0x56, 0x41, 0x4c, 0x49, 0x44, 0x5f, 0x4f,
	0x53, 0x54, 0x59, 0x50, 0x45, 0x10, 0x1f, 0x12, 0x20, 0x0a, 0x1c, 0x49, 0x4e, 0x56, 0x41, 0x4c,
	0x49, 0x44, 0x5f, 0x50, 0x52, 0x4f, 0x54, 0x4f, 0x43, 0x4f, 0x4c, 0x5f, 0x46, 0x4f, 0x52, 0x5f,
	0x41, 0x50, 0x50, 0x5f, 0x55, 0x52, 0x4c, 0x10, 0x20, 0x12, 0x22, 0x0a, 0x1e, 0x49, 0x4e, 0x56,
	0x41, 0x4c, 0x49, 0x44, 0x5f, 0x50, 0x41, 0x43, 0x4b, 0x41, 0x47, 0x45, 0x5f, 0x49, 0x44, 0x5f,
	0x46, 0x4f, 0x52, 0x5f, 0x41, 0x50, 0x50, 0x5f, 0x55, 0x52, 0x4c, 0x10, 0x21, 0x12, 0x2d, 0x0a,
	0x29, 0x55, 0x52, 0x4c, 0x5f, 0x43, 0x55, 0x53, 0x54, 0x4f, 0x4d, 0x5f, 0x50, 0x41, 0x52, 0x41,
	0x4d, 0x45, 0x54, 0x45, 0x52, 0x53, 0x5f, 0x43, 0x4f, 0x55, 0x4e, 0x54, 0x5f, 0x45, 0x58, 0x43,
	0x45, 0x45, 0x44, 0x53, 0x5f, 0x4c, 0x49, 0x4d, 0x49, 0x54, 0x10, 0x22, 0x12, 0x32, 0x0a, 0x2e,
	0x49, 0x4e, 0x56, 0x41, 0x4c, 0x49, 0x44, 0x5f, 0x43, 0x48, 0x41, 0x52, 0x41, 0x43, 0x54, 0x45,
	0x52, 0x53, 0x5f, 0x49, 0x4e, 0x5f, 0x55, 0x52, 0x4c, 0x5f, 0x43, 0x55, 0x53, 0x54, 0x4f, 0x4d,
	0x5f, 0x50, 0x41, 0x52, 0x41, 0x4d, 0x45, 0x54, 0x45, 0x52, 0x5f, 0x4b, 0x45, 0x59, 0x10, 0x27,
	0x12, 0x34, 0x0a, 0x30, 0x49, 0x4e, 0x56, 0x41, 0x4c, 0x49, 0x44, 0x5f, 0x43, 0x48, 0x41, 0x52,
	0x41, 0x43, 0x54, 0x45, 0x52, 0x53, 0x5f, 0x49, 0x4e, 0x5f, 0x55, 0x52, 0x4c, 0x5f, 0x43, 0x55,
	0x53, 0x54, 0x4f, 0x4d, 0x5f, 0x50, 0x41, 0x52, 0x41, 0x4d, 0x45, 0x54, 0x45, 0x52, 0x5f, 0x56,
	0x41, 0x4c, 0x55, 0x45, 0x10, 0x28, 0x12, 0x2d, 0x0a, 0x29, 0x49, 0x4e, 0x56, 0x41, 0x4c, 0x49,
	0x44, 0x5f, 0x54, 0x41, 0x47, 0x5f, 0x49, 0x4e, 0x5f, 0x55, 0x52, 0x4c, 0x5f, 0x43, 0x55, 0x53,
	0x54, 0x4f, 0x4d, 0x5f, 0x50, 0x41, 0x52, 0x41, 0x4d, 0x45, 0x54, 0x45, 0x52, 0x5f, 0x56, 0x41,
	0x4c, 0x55, 0x45, 0x10, 0x29, 0x12, 0x2d, 0x0a, 0x29, 0x52, 0x45, 0x44, 0x55, 0x4e, 0x44, 0x41,
	0x4e, 0x54, 0x5f, 0x4e, 0x45, 0x53, 0x54, 0x45, 0x44, 0x5f, 0x55, 0x52, 0x4c, 0x5f, 0x43, 0x55,
	0x53, 0x54, 0x4f, 0x4d, 0x5f, 0x50, 0x41, 0x52, 0x41, 0x4d, 0x45, 0x54, 0x45, 0x52, 0x5f, 0x54,
	0x41, 0x47, 0x10, 0x2a, 0x12, 0x14, 0x0a, 0x10, 0x4d, 0x49, 0x53, 0x53, 0x49, 0x4e, 0x47, 0x5f,
	0x50, 0x52, 0x4f, 0x54, 0x4f, 0x43, 0x4f, 0x4c, 0x10, 0x2b, 0x12, 0x14, 0x0a, 0x10, 0x49, 0x4e,
	0x56, 0x41, 0x4c, 0x49, 0x44, 0x5f, 0x50, 0x52, 0x4f, 0x54, 0x4f, 0x43, 0x4f, 0x4c, 0x10, 0x34,
	0x12, 0x0f, 0x0a, 0x0b, 0x49, 0x4e, 0x56, 0x41, 0x4c, 0x49, 0x44, 0x5f, 0x55, 0x52, 0x4c, 0x10,
	0x2c, 0x12, 0x1e, 0x0a, 0x1a, 0x44, 0x45, 0x53, 0x54, 0x49, 0x4e, 0x41, 0x54, 0x49, 0x4f, 0x4e,
	0x5f, 0x55, 0x52, 0x4c, 0x5f, 0x44, 0x45, 0x50, 0x52, 0x45, 0x43, 0x41, 0x54, 0x45, 0x44, 0x10,
	0x2d, 0x12, 0x16, 0x0a, 0x12, 0x49, 0x4e, 0x56, 0x41, 0x4c, 0x49, 0x44, 0x5f, 0x54, 0x41, 0x47,
	0x5f, 0x49, 0x4e, 0x5f, 0x55, 0x52, 0x4c, 0x10, 0x2e, 0x12, 0x13, 0x0a, 0x0f, 0x4d, 0x49, 0x53,
	0x53, 0x49, 0x4e, 0x47, 0x5f, 0x55, 0x52, 0x4c, 0x5f, 0x54, 0x41, 0x47, 0x10, 0x2f, 0x12, 0x14,
	0x0a, 0x10, 0x44, 0x55, 0x50, 0x4c, 0x49, 0x43, 0x41, 0x54, 0x45, 0x5f, 0x55, 0x52, 0x4c, 0x5f,
	0x49, 0x44, 0x10, 0x30, 0x12, 0x12, 0x0a, 0x0e, 0x49, 0x4e, 0x56, 0x41, 0x4c, 0x49, 0x44, 0x5f,
	0x55, 0x52, 0x4c, 0x5f, 0x49, 0x44, 0x10, 0x31, 0x12, 0x1e, 0x0a, 0x1a, 0x46, 0x49, 0x4e, 0x41,
	0x4c, 0x5f, 0x55, 0x52, 0x4c, 0x5f, 0x53, 0x55, 0x46, 0x46, 0x49, 0x58, 0x5f, 0x4d, 0x41, 0x4c,
	0x46, 0x4f, 0x52, 0x4d, 0x45, 0x44, 0x10, 0x32, 0x12, 0x23, 0x0a, 0x1f, 0x49, 0x4e, 0x56, 0x41,
	0x4c, 0x49, 0x44, 0x5f, 0x54, 0x41, 0x47, 0x5f, 0x49, 0x4e, 0x5f, 0x46, 0x49, 0x4e, 0x41, 0x4c,
	0x5f, 0x55, 0x52, 0x4c, 0x5f, 0x53, 0x55, 0x46, 0x46, 0x49, 0x58, 0x10, 0x33, 0x12, 0x1c, 0x0a,
	0x18, 0x49, 0x4e, 0x56, 0x41, 0x4c, 0x49, 0x44, 0x5f, 0x54, 0x4f, 0x50, 0x5f, 0x4c, 0x45, 0x56,
	0x45, 0x4c, 0x5f, 0x44, 0x4f, 0x4d, 0x41, 0x49, 0x4e, 0x10, 0x35, 0x12, 0x1e, 0x0a, 0x1a, 0x4d,
	0x41, 0x4c, 0x46, 0x4f, 0x52, 0x4d, 0x45, 0x44, 0x5f, 0x54, 0x4f, 0x50, 0x5f, 0x4c, 0x45, 0x56,
	0x45, 0x4c, 0x5f, 0x44, 0x4f, 0x4d, 0x41, 0x49, 0x4e, 0x10, 0x36, 0x12, 0x11, 0x0a, 0x0d, 0x4d,
	0x41, 0x4c, 0x46, 0x4f, 0x52, 0x4d, 0x45, 0x44, 0x5f, 0x55, 0x52, 0x4c, 0x10, 0x37, 0x12, 0x10,
	0x0a, 0x0c, 0x4d, 0x49, 0x53, 0x53, 0x49, 0x4e, 0x47, 0x5f, 0x48, 0x4f, 0x53, 0x54, 0x10, 0x38,
	0x12, 0x1f, 0x0a, 0x1b, 0x4e, 0x55, 0x4c, 0x4c, 0x5f, 0x43, 0x55, 0x53, 0x54, 0x4f, 0x4d, 0x5f,
	0x50, 0x41, 0x52, 0x41, 0x4d, 0x45, 0x54, 0x45, 0x52, 0x5f, 0x56, 0x41, 0x4c, 0x55, 0x45, 0x10,
	0x39, 0x12, 0x27, 0x0a, 0x23, 0x56, 0x41, 0x4c, 0x55, 0x45, 0x5f, 0x54, 0x52, 0x41, 0x43, 0x4b,
	0x5f, 0x50, 0x41, 0x52, 0x41, 0x4d, 0x45, 0x54, 0x45, 0x52, 0x5f, 0x4e, 0x4f, 0x54, 0x5f, 0x53,
	0x55, 0x50, 0x50, 0x4f, 0x52, 0x54, 0x45, 0x44, 0x10, 0x3a, 0x42, 0xf2, 0x01, 0x0a, 0x23, 0x63,
	0x6f, 0x6d, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x61, 0x64, 0x73, 0x2e, 0x67, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x61, 0x64, 0x73, 0x2e, 0x76, 0x31, 0x34, 0x2e, 0x65, 0x72, 0x72, 0x6f,
	0x72, 0x73, 0x42, 0x12, 0x55, 0x72, 0x6c, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x45, 0x72, 0x72, 0x6f,
	0x72, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a, 0x45, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x2e, 0x67, 0x6f, 0x6c, 0x61, 0x6e, 0x67, 0x2e, 0x6f, 0x72, 0x67, 0x2f, 0x67, 0x65, 0x6e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61, 0x70, 0x69, 0x73, 0x2f,
	0x61, 0x64, 0x73, 0x2f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61, 0x64, 0x73, 0x2f, 0x76, 0x31,
	0x34, 0x2f, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x73, 0x3b, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x73, 0xa2,
	0x02, 0x03, 0x47, 0x41, 0x41, 0xaa, 0x02, 0x1f, 0x47, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x41,
	0x64, 0x73, 0x2e, 0x47, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x41, 0x64, 0x73, 0x2e, 0x56, 0x31, 0x34,
	0x2e, 0x45, 0x72, 0x72, 0x6f, 0x72, 0x73, 0xca, 0x02, 0x1f, 0x47, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x5c, 0x41, 0x64, 0x73, 0x5c, 0x47, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x41, 0x64, 0x73, 0x5c, 0x56,
	0x31, 0x34, 0x5c, 0x45, 0x72, 0x72, 0x6f, 0x72, 0x73, 0xea, 0x02, 0x23, 0x47, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x3a, 0x3a, 0x41, 0x64, 0x73, 0x3a, 0x3a, 0x47, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x41,
	0x64, 0x73, 0x3a, 0x3a, 0x56, 0x31, 0x34, 0x3a, 0x3a, 0x45, 0x72, 0x72, 0x6f, 0x72, 0x73, 0x62,
	0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_errors_url_field_error_proto_rawDescOnce sync.Once
	file_errors_url_field_error_proto_rawDescData = file_errors_url_field_error_proto_rawDesc
)

func file_errors_url_field_error_proto_rawDescGZIP() []byte {
	file_errors_url_field_error_proto_rawDescOnce.Do(func() {
		file_errors_url_field_error_proto_rawDescData = protoimpl.X.CompressGZIP(file_errors_url_field_error_proto_rawDescData)
	})
	return file_errors_url_field_error_proto_rawDescData
}

var file_errors_url_field_error_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_errors_url_field_error_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_errors_url_field_error_proto_goTypes = []interface{}{
	(UrlFieldErrorEnum_UrlFieldError)(0), // 0: google.ads.googleads.v14.errors.UrlFieldErrorEnum.UrlFieldError
	(*UrlFieldErrorEnum)(nil),            // 1: google.ads.googleads.v14.errors.UrlFieldErrorEnum
}
var file_errors_url_field_error_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_errors_url_field_error_proto_init() }
func file_errors_url_field_error_proto_init() {
	if File_errors_url_field_error_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_errors_url_field_error_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UrlFieldErrorEnum); i {
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
			RawDescriptor: file_errors_url_field_error_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_errors_url_field_error_proto_goTypes,
		DependencyIndexes: file_errors_url_field_error_proto_depIdxs,
		EnumInfos:         file_errors_url_field_error_proto_enumTypes,
		MessageInfos:      file_errors_url_field_error_proto_msgTypes,
	}.Build()
	File_errors_url_field_error_proto = out.File
	file_errors_url_field_error_proto_rawDesc = nil
	file_errors_url_field_error_proto_goTypes = nil
	file_errors_url_field_error_proto_depIdxs = nil
}
