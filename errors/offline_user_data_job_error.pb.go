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
// source: google/ads/googleads/v10/errors/offline_user_data_job_error.proto

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

// Enum describing possible request errors.
type OfflineUserDataJobErrorEnum_OfflineUserDataJobError int32

const (
	// Enum unspecified.
	OfflineUserDataJobErrorEnum_UNSPECIFIED OfflineUserDataJobErrorEnum_OfflineUserDataJobError = 0
	// The received error code is not known in this version.
	OfflineUserDataJobErrorEnum_UNKNOWN OfflineUserDataJobErrorEnum_OfflineUserDataJobError = 1
	// The user list ID provided for the job is invalid.
	OfflineUserDataJobErrorEnum_INVALID_USER_LIST_ID OfflineUserDataJobErrorEnum_OfflineUserDataJobError = 3
	// Type of the user list is not applicable for the job.
	OfflineUserDataJobErrorEnum_INVALID_USER_LIST_TYPE OfflineUserDataJobErrorEnum_OfflineUserDataJobError = 4
	// Customer is not allowisted for using user ID in upload data.
	OfflineUserDataJobErrorEnum_NOT_ON_ALLOWLIST_FOR_USER_ID OfflineUserDataJobErrorEnum_OfflineUserDataJobError = 33
	// Upload data is not compatible with the upload key type of the associated
	// user list.
	OfflineUserDataJobErrorEnum_INCOMPATIBLE_UPLOAD_KEY_TYPE OfflineUserDataJobErrorEnum_OfflineUserDataJobError = 6
	// The user identifier is missing valid data.
	OfflineUserDataJobErrorEnum_MISSING_USER_IDENTIFIER OfflineUserDataJobErrorEnum_OfflineUserDataJobError = 7
	// The mobile ID is malformed.
	OfflineUserDataJobErrorEnum_INVALID_MOBILE_ID_FORMAT OfflineUserDataJobErrorEnum_OfflineUserDataJobError = 8
	// Maximum number of user identifiers allowed per request is 100,000 and per
	// operation is 20.
	OfflineUserDataJobErrorEnum_TOO_MANY_USER_IDENTIFIERS OfflineUserDataJobErrorEnum_OfflineUserDataJobError = 9
	// Customer is not on the allow-list for store sales direct data.
	OfflineUserDataJobErrorEnum_NOT_ON_ALLOWLIST_FOR_STORE_SALES_DIRECT OfflineUserDataJobErrorEnum_OfflineUserDataJobError = 31
	// Customer is not on the allow-list for unified store sales data.
	OfflineUserDataJobErrorEnum_NOT_ON_ALLOWLIST_FOR_UNIFIED_STORE_SALES OfflineUserDataJobErrorEnum_OfflineUserDataJobError = 32
	// The partner ID in store sales direct metadata is invalid.
	OfflineUserDataJobErrorEnum_INVALID_PARTNER_ID OfflineUserDataJobErrorEnum_OfflineUserDataJobError = 11
	// The data in user identifier should not be encoded.
	OfflineUserDataJobErrorEnum_INVALID_ENCODING OfflineUserDataJobErrorEnum_OfflineUserDataJobError = 12
	// The country code is invalid.
	OfflineUserDataJobErrorEnum_INVALID_COUNTRY_CODE OfflineUserDataJobErrorEnum_OfflineUserDataJobError = 13
	// Incompatible user identifier when using third_party_user_id for store
	// sales direct first party data or not using third_party_user_id for store
	// sales third party data.
	OfflineUserDataJobErrorEnum_INCOMPATIBLE_USER_IDENTIFIER OfflineUserDataJobErrorEnum_OfflineUserDataJobError = 14
	// A transaction time in the future is not allowed.
	OfflineUserDataJobErrorEnum_FUTURE_TRANSACTION_TIME OfflineUserDataJobErrorEnum_OfflineUserDataJobError = 15
	// The conversion_action specified in transaction_attributes is used to
	// report conversions to a conversion action configured in Google Ads. This
	// error indicates there is no such conversion action in the account.
	OfflineUserDataJobErrorEnum_INVALID_CONVERSION_ACTION OfflineUserDataJobErrorEnum_OfflineUserDataJobError = 16
	// Mobile ID is not supported for store sales direct data.
	OfflineUserDataJobErrorEnum_MOBILE_ID_NOT_SUPPORTED OfflineUserDataJobErrorEnum_OfflineUserDataJobError = 17
	// When a remove-all operation is provided, it has to be the first operation
	// of the operation list.
	OfflineUserDataJobErrorEnum_INVALID_OPERATION_ORDER OfflineUserDataJobErrorEnum_OfflineUserDataJobError = 18
	// Mixing creation and removal of offline data in the same job is not
	// allowed.
	OfflineUserDataJobErrorEnum_CONFLICTING_OPERATION OfflineUserDataJobErrorEnum_OfflineUserDataJobError = 19
	// The external update ID already exists.
	OfflineUserDataJobErrorEnum_EXTERNAL_UPDATE_ID_ALREADY_EXISTS OfflineUserDataJobErrorEnum_OfflineUserDataJobError = 21
	// Once the upload job is started, new operations cannot be added.
	OfflineUserDataJobErrorEnum_JOB_ALREADY_STARTED OfflineUserDataJobErrorEnum_OfflineUserDataJobError = 22
	// Remove operation is not allowed for store sales direct updates.
	OfflineUserDataJobErrorEnum_REMOVE_NOT_SUPPORTED OfflineUserDataJobErrorEnum_OfflineUserDataJobError = 23
	// Remove-all is not supported for certain offline user data job types.
	OfflineUserDataJobErrorEnum_REMOVE_ALL_NOT_SUPPORTED OfflineUserDataJobErrorEnum_OfflineUserDataJobError = 24
	// The SHA256 encoded value is malformed.
	OfflineUserDataJobErrorEnum_INVALID_SHA256_FORMAT OfflineUserDataJobErrorEnum_OfflineUserDataJobError = 25
	// The custom key specified is not enabled for the unified store sales
	// upload.
	OfflineUserDataJobErrorEnum_CUSTOM_KEY_DISABLED OfflineUserDataJobErrorEnum_OfflineUserDataJobError = 26
	// The custom key specified is not predefined through the Google Ads UI.
	OfflineUserDataJobErrorEnum_CUSTOM_KEY_NOT_PREDEFINED OfflineUserDataJobErrorEnum_OfflineUserDataJobError = 27
	// The custom key specified is not set in the upload.
	OfflineUserDataJobErrorEnum_CUSTOM_KEY_NOT_SET OfflineUserDataJobErrorEnum_OfflineUserDataJobError = 29
	// The customer has not accepted the customer data terms in the conversion
	// settings page.
	OfflineUserDataJobErrorEnum_CUSTOMER_NOT_ACCEPTED_CUSTOMER_DATA_TERMS OfflineUserDataJobErrorEnum_OfflineUserDataJobError = 30
	// User attributes cannot be uploaded into a user list.
	OfflineUserDataJobErrorEnum_ATTRIBUTES_NOT_APPLICABLE_FOR_CUSTOMER_MATCH_USER_LIST OfflineUserDataJobErrorEnum_OfflineUserDataJobError = 34
	// Lifetime value bucket must be a number from 1-10, except for remove
	// operation where 0 will be accepted.
	OfflineUserDataJobErrorEnum_LIFETIME_VALUE_BUCKET_NOT_IN_RANGE OfflineUserDataJobErrorEnum_OfflineUserDataJobError = 35
	// Identifiers not supported for Customer Match attributes. User attributes
	// can only be provided with contact info (email, phone, address) user
	// identifiers.
	OfflineUserDataJobErrorEnum_INCOMPATIBLE_USER_IDENTIFIER_FOR_ATTRIBUTES OfflineUserDataJobErrorEnum_OfflineUserDataJobError = 36
	// A time in the future is not allowed.
	OfflineUserDataJobErrorEnum_FUTURE_TIME_NOT_ALLOWED OfflineUserDataJobErrorEnum_OfflineUserDataJobError = 37
	// Last purchase date time cannot be less than acquisition date time.
	OfflineUserDataJobErrorEnum_LAST_PURCHASE_TIME_LESS_THAN_ACQUISITION_TIME OfflineUserDataJobErrorEnum_OfflineUserDataJobError = 38
	// Only emails are accepted as user identifiers for shopping loyalty match.
	// {-- api.dev/not-precedent: The identifier is not limited to ids, but
	// also include other user info eg. phone numbers.}
	OfflineUserDataJobErrorEnum_CUSTOMER_IDENTIFIER_NOT_ALLOWED OfflineUserDataJobErrorEnum_OfflineUserDataJobError = 39
	// Provided item ID is invalid.
	OfflineUserDataJobErrorEnum_INVALID_ITEM_ID OfflineUserDataJobErrorEnum_OfflineUserDataJobError = 40
)

// Enum value maps for OfflineUserDataJobErrorEnum_OfflineUserDataJobError.
var (
	OfflineUserDataJobErrorEnum_OfflineUserDataJobError_name = map[int32]string{
		0:  "UNSPECIFIED",
		1:  "UNKNOWN",
		3:  "INVALID_USER_LIST_ID",
		4:  "INVALID_USER_LIST_TYPE",
		33: "NOT_ON_ALLOWLIST_FOR_USER_ID",
		6:  "INCOMPATIBLE_UPLOAD_KEY_TYPE",
		7:  "MISSING_USER_IDENTIFIER",
		8:  "INVALID_MOBILE_ID_FORMAT",
		9:  "TOO_MANY_USER_IDENTIFIERS",
		31: "NOT_ON_ALLOWLIST_FOR_STORE_SALES_DIRECT",
		32: "NOT_ON_ALLOWLIST_FOR_UNIFIED_STORE_SALES",
		11: "INVALID_PARTNER_ID",
		12: "INVALID_ENCODING",
		13: "INVALID_COUNTRY_CODE",
		14: "INCOMPATIBLE_USER_IDENTIFIER",
		15: "FUTURE_TRANSACTION_TIME",
		16: "INVALID_CONVERSION_ACTION",
		17: "MOBILE_ID_NOT_SUPPORTED",
		18: "INVALID_OPERATION_ORDER",
		19: "CONFLICTING_OPERATION",
		21: "EXTERNAL_UPDATE_ID_ALREADY_EXISTS",
		22: "JOB_ALREADY_STARTED",
		23: "REMOVE_NOT_SUPPORTED",
		24: "REMOVE_ALL_NOT_SUPPORTED",
		25: "INVALID_SHA256_FORMAT",
		26: "CUSTOM_KEY_DISABLED",
		27: "CUSTOM_KEY_NOT_PREDEFINED",
		29: "CUSTOM_KEY_NOT_SET",
		30: "CUSTOMER_NOT_ACCEPTED_CUSTOMER_DATA_TERMS",
		34: "ATTRIBUTES_NOT_APPLICABLE_FOR_CUSTOMER_MATCH_USER_LIST",
		35: "LIFETIME_VALUE_BUCKET_NOT_IN_RANGE",
		36: "INCOMPATIBLE_USER_IDENTIFIER_FOR_ATTRIBUTES",
		37: "FUTURE_TIME_NOT_ALLOWED",
		38: "LAST_PURCHASE_TIME_LESS_THAN_ACQUISITION_TIME",
		39: "CUSTOMER_IDENTIFIER_NOT_ALLOWED",
		40: "INVALID_ITEM_ID",
	}
	OfflineUserDataJobErrorEnum_OfflineUserDataJobError_value = map[string]int32{
		"UNSPECIFIED":                                            0,
		"UNKNOWN":                                                1,
		"INVALID_USER_LIST_ID":                                   3,
		"INVALID_USER_LIST_TYPE":                                 4,
		"NOT_ON_ALLOWLIST_FOR_USER_ID":                           33,
		"INCOMPATIBLE_UPLOAD_KEY_TYPE":                           6,
		"MISSING_USER_IDENTIFIER":                                7,
		"INVALID_MOBILE_ID_FORMAT":                               8,
		"TOO_MANY_USER_IDENTIFIERS":                              9,
		"NOT_ON_ALLOWLIST_FOR_STORE_SALES_DIRECT":                31,
		"NOT_ON_ALLOWLIST_FOR_UNIFIED_STORE_SALES":               32,
		"INVALID_PARTNER_ID":                                     11,
		"INVALID_ENCODING":                                       12,
		"INVALID_COUNTRY_CODE":                                   13,
		"INCOMPATIBLE_USER_IDENTIFIER":                           14,
		"FUTURE_TRANSACTION_TIME":                                15,
		"INVALID_CONVERSION_ACTION":                              16,
		"MOBILE_ID_NOT_SUPPORTED":                                17,
		"INVALID_OPERATION_ORDER":                                18,
		"CONFLICTING_OPERATION":                                  19,
		"EXTERNAL_UPDATE_ID_ALREADY_EXISTS":                      21,
		"JOB_ALREADY_STARTED":                                    22,
		"REMOVE_NOT_SUPPORTED":                                   23,
		"REMOVE_ALL_NOT_SUPPORTED":                               24,
		"INVALID_SHA256_FORMAT":                                  25,
		"CUSTOM_KEY_DISABLED":                                    26,
		"CUSTOM_KEY_NOT_PREDEFINED":                              27,
		"CUSTOM_KEY_NOT_SET":                                     29,
		"CUSTOMER_NOT_ACCEPTED_CUSTOMER_DATA_TERMS":              30,
		"ATTRIBUTES_NOT_APPLICABLE_FOR_CUSTOMER_MATCH_USER_LIST": 34,
		"LIFETIME_VALUE_BUCKET_NOT_IN_RANGE":                     35,
		"INCOMPATIBLE_USER_IDENTIFIER_FOR_ATTRIBUTES":            36,
		"FUTURE_TIME_NOT_ALLOWED":                                37,
		"LAST_PURCHASE_TIME_LESS_THAN_ACQUISITION_TIME":          38,
		"CUSTOMER_IDENTIFIER_NOT_ALLOWED":                        39,
		"INVALID_ITEM_ID":                                        40,
	}
)

func (x OfflineUserDataJobErrorEnum_OfflineUserDataJobError) Enum() *OfflineUserDataJobErrorEnum_OfflineUserDataJobError {
	p := new(OfflineUserDataJobErrorEnum_OfflineUserDataJobError)
	*p = x
	return p
}

func (x OfflineUserDataJobErrorEnum_OfflineUserDataJobError) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (OfflineUserDataJobErrorEnum_OfflineUserDataJobError) Descriptor() protoreflect.EnumDescriptor {
	return file_errors_offline_user_data_job_error_proto_enumTypes[0].Descriptor()
}

func (OfflineUserDataJobErrorEnum_OfflineUserDataJobError) Type() protoreflect.EnumType {
	return &file_errors_offline_user_data_job_error_proto_enumTypes[0]
}

func (x OfflineUserDataJobErrorEnum_OfflineUserDataJobError) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use OfflineUserDataJobErrorEnum_OfflineUserDataJobError.Descriptor instead.
func (OfflineUserDataJobErrorEnum_OfflineUserDataJobError) EnumDescriptor() ([]byte, []int) {
	return file_errors_offline_user_data_job_error_proto_rawDescGZIP(), []int{0, 0}
}

// Container for enum describing possible offline user data job errors.
type OfflineUserDataJobErrorEnum struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *OfflineUserDataJobErrorEnum) Reset() {
	*x = OfflineUserDataJobErrorEnum{}
	if protoimpl.UnsafeEnabled {
		mi := &file_errors_offline_user_data_job_error_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *OfflineUserDataJobErrorEnum) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*OfflineUserDataJobErrorEnum) ProtoMessage() {}

func (x *OfflineUserDataJobErrorEnum) ProtoReflect() protoreflect.Message {
	mi := &file_errors_offline_user_data_job_error_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use OfflineUserDataJobErrorEnum.ProtoReflect.Descriptor instead.
func (*OfflineUserDataJobErrorEnum) Descriptor() ([]byte, []int) {
	return file_errors_offline_user_data_job_error_proto_rawDescGZIP(), []int{0}
}

var File_errors_offline_user_data_job_error_proto protoreflect.FileDescriptor

var file_errors_offline_user_data_job_error_proto_rawDesc = []byte{
	0x0a, 0x41, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x64, 0x73, 0x2f, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x61, 0x64, 0x73, 0x2f, 0x76, 0x31, 0x30, 0x2f, 0x65, 0x72, 0x72, 0x6f, 0x72,
	0x73, 0x2f, 0x6f, 0x66, 0x66, 0x6c, 0x69, 0x6e, 0x65, 0x5f, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x64,
	0x61, 0x74, 0x61, 0x5f, 0x6a, 0x6f, 0x62, 0x5f, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x12, 0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x61, 0x64, 0x73, 0x2e,
	0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61, 0x64, 0x73, 0x2e, 0x76, 0x31, 0x30, 0x2e, 0x65, 0x72,
	0x72, 0x6f, 0x72, 0x73, 0x1a, 0x1c, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69,
	0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x22, 0xb2, 0x09, 0x0a, 0x1b, 0x4f, 0x66, 0x66, 0x6c, 0x69, 0x6e, 0x65, 0x55, 0x73,
	0x65, 0x72, 0x44, 0x61, 0x74, 0x61, 0x4a, 0x6f, 0x62, 0x45, 0x72, 0x72, 0x6f, 0x72, 0x45, 0x6e,
	0x75, 0x6d, 0x22, 0x92, 0x09, 0x0a, 0x17, 0x4f, 0x66, 0x66, 0x6c, 0x69, 0x6e, 0x65, 0x55, 0x73,
	0x65, 0x72, 0x44, 0x61, 0x74, 0x61, 0x4a, 0x6f, 0x62, 0x45, 0x72, 0x72, 0x6f, 0x72, 0x12, 0x0f,
	0x0a, 0x0b, 0x55, 0x4e, 0x53, 0x50, 0x45, 0x43, 0x49, 0x46, 0x49, 0x45, 0x44, 0x10, 0x00, 0x12,
	0x0b, 0x0a, 0x07, 0x55, 0x4e, 0x4b, 0x4e, 0x4f, 0x57, 0x4e, 0x10, 0x01, 0x12, 0x18, 0x0a, 0x14,
	0x49, 0x4e, 0x56, 0x41, 0x4c, 0x49, 0x44, 0x5f, 0x55, 0x53, 0x45, 0x52, 0x5f, 0x4c, 0x49, 0x53,
	0x54, 0x5f, 0x49, 0x44, 0x10, 0x03, 0x12, 0x1a, 0x0a, 0x16, 0x49, 0x4e, 0x56, 0x41, 0x4c, 0x49,
	0x44, 0x5f, 0x55, 0x53, 0x45, 0x52, 0x5f, 0x4c, 0x49, 0x53, 0x54, 0x5f, 0x54, 0x59, 0x50, 0x45,
	0x10, 0x04, 0x12, 0x20, 0x0a, 0x1c, 0x4e, 0x4f, 0x54, 0x5f, 0x4f, 0x4e, 0x5f, 0x41, 0x4c, 0x4c,
	0x4f, 0x57, 0x4c, 0x49, 0x53, 0x54, 0x5f, 0x46, 0x4f, 0x52, 0x5f, 0x55, 0x53, 0x45, 0x52, 0x5f,
	0x49, 0x44, 0x10, 0x21, 0x12, 0x20, 0x0a, 0x1c, 0x49, 0x4e, 0x43, 0x4f, 0x4d, 0x50, 0x41, 0x54,
	0x49, 0x42, 0x4c, 0x45, 0x5f, 0x55, 0x50, 0x4c, 0x4f, 0x41, 0x44, 0x5f, 0x4b, 0x45, 0x59, 0x5f,
	0x54, 0x59, 0x50, 0x45, 0x10, 0x06, 0x12, 0x1b, 0x0a, 0x17, 0x4d, 0x49, 0x53, 0x53, 0x49, 0x4e,
	0x47, 0x5f, 0x55, 0x53, 0x45, 0x52, 0x5f, 0x49, 0x44, 0x45, 0x4e, 0x54, 0x49, 0x46, 0x49, 0x45,
	0x52, 0x10, 0x07, 0x12, 0x1c, 0x0a, 0x18, 0x49, 0x4e, 0x56, 0x41, 0x4c, 0x49, 0x44, 0x5f, 0x4d,
	0x4f, 0x42, 0x49, 0x4c, 0x45, 0x5f, 0x49, 0x44, 0x5f, 0x46, 0x4f, 0x52, 0x4d, 0x41, 0x54, 0x10,
	0x08, 0x12, 0x1d, 0x0a, 0x19, 0x54, 0x4f, 0x4f, 0x5f, 0x4d, 0x41, 0x4e, 0x59, 0x5f, 0x55, 0x53,
	0x45, 0x52, 0x5f, 0x49, 0x44, 0x45, 0x4e, 0x54, 0x49, 0x46, 0x49, 0x45, 0x52, 0x53, 0x10, 0x09,
	0x12, 0x2b, 0x0a, 0x27, 0x4e, 0x4f, 0x54, 0x5f, 0x4f, 0x4e, 0x5f, 0x41, 0x4c, 0x4c, 0x4f, 0x57,
	0x4c, 0x49, 0x53, 0x54, 0x5f, 0x46, 0x4f, 0x52, 0x5f, 0x53, 0x54, 0x4f, 0x52, 0x45, 0x5f, 0x53,
	0x41, 0x4c, 0x45, 0x53, 0x5f, 0x44, 0x49, 0x52, 0x45, 0x43, 0x54, 0x10, 0x1f, 0x12, 0x2c, 0x0a,
	0x28, 0x4e, 0x4f, 0x54, 0x5f, 0x4f, 0x4e, 0x5f, 0x41, 0x4c, 0x4c, 0x4f, 0x57, 0x4c, 0x49, 0x53,
	0x54, 0x5f, 0x46, 0x4f, 0x52, 0x5f, 0x55, 0x4e, 0x49, 0x46, 0x49, 0x45, 0x44, 0x5f, 0x53, 0x54,
	0x4f, 0x52, 0x45, 0x5f, 0x53, 0x41, 0x4c, 0x45, 0x53, 0x10, 0x20, 0x12, 0x16, 0x0a, 0x12, 0x49,
	0x4e, 0x56, 0x41, 0x4c, 0x49, 0x44, 0x5f, 0x50, 0x41, 0x52, 0x54, 0x4e, 0x45, 0x52, 0x5f, 0x49,
	0x44, 0x10, 0x0b, 0x12, 0x14, 0x0a, 0x10, 0x49, 0x4e, 0x56, 0x41, 0x4c, 0x49, 0x44, 0x5f, 0x45,
	0x4e, 0x43, 0x4f, 0x44, 0x49, 0x4e, 0x47, 0x10, 0x0c, 0x12, 0x18, 0x0a, 0x14, 0x49, 0x4e, 0x56,
	0x41, 0x4c, 0x49, 0x44, 0x5f, 0x43, 0x4f, 0x55, 0x4e, 0x54, 0x52, 0x59, 0x5f, 0x43, 0x4f, 0x44,
	0x45, 0x10, 0x0d, 0x12, 0x20, 0x0a, 0x1c, 0x49, 0x4e, 0x43, 0x4f, 0x4d, 0x50, 0x41, 0x54, 0x49,
	0x42, 0x4c, 0x45, 0x5f, 0x55, 0x53, 0x45, 0x52, 0x5f, 0x49, 0x44, 0x45, 0x4e, 0x54, 0x49, 0x46,
	0x49, 0x45, 0x52, 0x10, 0x0e, 0x12, 0x1b, 0x0a, 0x17, 0x46, 0x55, 0x54, 0x55, 0x52, 0x45, 0x5f,
	0x54, 0x52, 0x41, 0x4e, 0x53, 0x41, 0x43, 0x54, 0x49, 0x4f, 0x4e, 0x5f, 0x54, 0x49, 0x4d, 0x45,
	0x10, 0x0f, 0x12, 0x1d, 0x0a, 0x19, 0x49, 0x4e, 0x56, 0x41, 0x4c, 0x49, 0x44, 0x5f, 0x43, 0x4f,
	0x4e, 0x56, 0x45, 0x52, 0x53, 0x49, 0x4f, 0x4e, 0x5f, 0x41, 0x43, 0x54, 0x49, 0x4f, 0x4e, 0x10,
	0x10, 0x12, 0x1b, 0x0a, 0x17, 0x4d, 0x4f, 0x42, 0x49, 0x4c, 0x45, 0x5f, 0x49, 0x44, 0x5f, 0x4e,
	0x4f, 0x54, 0x5f, 0x53, 0x55, 0x50, 0x50, 0x4f, 0x52, 0x54, 0x45, 0x44, 0x10, 0x11, 0x12, 0x1b,
	0x0a, 0x17, 0x49, 0x4e, 0x56, 0x41, 0x4c, 0x49, 0x44, 0x5f, 0x4f, 0x50, 0x45, 0x52, 0x41, 0x54,
	0x49, 0x4f, 0x4e, 0x5f, 0x4f, 0x52, 0x44, 0x45, 0x52, 0x10, 0x12, 0x12, 0x19, 0x0a, 0x15, 0x43,
	0x4f, 0x4e, 0x46, 0x4c, 0x49, 0x43, 0x54, 0x49, 0x4e, 0x47, 0x5f, 0x4f, 0x50, 0x45, 0x52, 0x41,
	0x54, 0x49, 0x4f, 0x4e, 0x10, 0x13, 0x12, 0x25, 0x0a, 0x21, 0x45, 0x58, 0x54, 0x45, 0x52, 0x4e,
	0x41, 0x4c, 0x5f, 0x55, 0x50, 0x44, 0x41, 0x54, 0x45, 0x5f, 0x49, 0x44, 0x5f, 0x41, 0x4c, 0x52,
	0x45, 0x41, 0x44, 0x59, 0x5f, 0x45, 0x58, 0x49, 0x53, 0x54, 0x53, 0x10, 0x15, 0x12, 0x17, 0x0a,
	0x13, 0x4a, 0x4f, 0x42, 0x5f, 0x41, 0x4c, 0x52, 0x45, 0x41, 0x44, 0x59, 0x5f, 0x53, 0x54, 0x41,
	0x52, 0x54, 0x45, 0x44, 0x10, 0x16, 0x12, 0x18, 0x0a, 0x14, 0x52, 0x45, 0x4d, 0x4f, 0x56, 0x45,
	0x5f, 0x4e, 0x4f, 0x54, 0x5f, 0x53, 0x55, 0x50, 0x50, 0x4f, 0x52, 0x54, 0x45, 0x44, 0x10, 0x17,
	0x12, 0x1c, 0x0a, 0x18, 0x52, 0x45, 0x4d, 0x4f, 0x56, 0x45, 0x5f, 0x41, 0x4c, 0x4c, 0x5f, 0x4e,
	0x4f, 0x54, 0x5f, 0x53, 0x55, 0x50, 0x50, 0x4f, 0x52, 0x54, 0x45, 0x44, 0x10, 0x18, 0x12, 0x19,
	0x0a, 0x15, 0x49, 0x4e, 0x56, 0x41, 0x4c, 0x49, 0x44, 0x5f, 0x53, 0x48, 0x41, 0x32, 0x35, 0x36,
	0x5f, 0x46, 0x4f, 0x52, 0x4d, 0x41, 0x54, 0x10, 0x19, 0x12, 0x17, 0x0a, 0x13, 0x43, 0x55, 0x53,
	0x54, 0x4f, 0x4d, 0x5f, 0x4b, 0x45, 0x59, 0x5f, 0x44, 0x49, 0x53, 0x41, 0x42, 0x4c, 0x45, 0x44,
	0x10, 0x1a, 0x12, 0x1d, 0x0a, 0x19, 0x43, 0x55, 0x53, 0x54, 0x4f, 0x4d, 0x5f, 0x4b, 0x45, 0x59,
	0x5f, 0x4e, 0x4f, 0x54, 0x5f, 0x50, 0x52, 0x45, 0x44, 0x45, 0x46, 0x49, 0x4e, 0x45, 0x44, 0x10,
	0x1b, 0x12, 0x16, 0x0a, 0x12, 0x43, 0x55, 0x53, 0x54, 0x4f, 0x4d, 0x5f, 0x4b, 0x45, 0x59, 0x5f,
	0x4e, 0x4f, 0x54, 0x5f, 0x53, 0x45, 0x54, 0x10, 0x1d, 0x12, 0x2d, 0x0a, 0x29, 0x43, 0x55, 0x53,
	0x54, 0x4f, 0x4d, 0x45, 0x52, 0x5f, 0x4e, 0x4f, 0x54, 0x5f, 0x41, 0x43, 0x43, 0x45, 0x50, 0x54,
	0x45, 0x44, 0x5f, 0x43, 0x55, 0x53, 0x54, 0x4f, 0x4d, 0x45, 0x52, 0x5f, 0x44, 0x41, 0x54, 0x41,
	0x5f, 0x54, 0x45, 0x52, 0x4d, 0x53, 0x10, 0x1e, 0x12, 0x3a, 0x0a, 0x36, 0x41, 0x54, 0x54, 0x52,
	0x49, 0x42, 0x55, 0x54, 0x45, 0x53, 0x5f, 0x4e, 0x4f, 0x54, 0x5f, 0x41, 0x50, 0x50, 0x4c, 0x49,
	0x43, 0x41, 0x42, 0x4c, 0x45, 0x5f, 0x46, 0x4f, 0x52, 0x5f, 0x43, 0x55, 0x53, 0x54, 0x4f, 0x4d,
	0x45, 0x52, 0x5f, 0x4d, 0x41, 0x54, 0x43, 0x48, 0x5f, 0x55, 0x53, 0x45, 0x52, 0x5f, 0x4c, 0x49,
	0x53, 0x54, 0x10, 0x22, 0x12, 0x26, 0x0a, 0x22, 0x4c, 0x49, 0x46, 0x45, 0x54, 0x49, 0x4d, 0x45,
	0x5f, 0x56, 0x41, 0x4c, 0x55, 0x45, 0x5f, 0x42, 0x55, 0x43, 0x4b, 0x45, 0x54, 0x5f, 0x4e, 0x4f,
	0x54, 0x5f, 0x49, 0x4e, 0x5f, 0x52, 0x41, 0x4e, 0x47, 0x45, 0x10, 0x23, 0x12, 0x2f, 0x0a, 0x2b,
	0x49, 0x4e, 0x43, 0x4f, 0x4d, 0x50, 0x41, 0x54, 0x49, 0x42, 0x4c, 0x45, 0x5f, 0x55, 0x53, 0x45,
	0x52, 0x5f, 0x49, 0x44, 0x45, 0x4e, 0x54, 0x49, 0x46, 0x49, 0x45, 0x52, 0x5f, 0x46, 0x4f, 0x52,
	0x5f, 0x41, 0x54, 0x54, 0x52, 0x49, 0x42, 0x55, 0x54, 0x45, 0x53, 0x10, 0x24, 0x12, 0x1b, 0x0a,
	0x17, 0x46, 0x55, 0x54, 0x55, 0x52, 0x45, 0x5f, 0x54, 0x49, 0x4d, 0x45, 0x5f, 0x4e, 0x4f, 0x54,
	0x5f, 0x41, 0x4c, 0x4c, 0x4f, 0x57, 0x45, 0x44, 0x10, 0x25, 0x12, 0x31, 0x0a, 0x2d, 0x4c, 0x41,
	0x53, 0x54, 0x5f, 0x50, 0x55, 0x52, 0x43, 0x48, 0x41, 0x53, 0x45, 0x5f, 0x54, 0x49, 0x4d, 0x45,
	0x5f, 0x4c, 0x45, 0x53, 0x53, 0x5f, 0x54, 0x48, 0x41, 0x4e, 0x5f, 0x41, 0x43, 0x51, 0x55, 0x49,
	0x53, 0x49, 0x54, 0x49, 0x4f, 0x4e, 0x5f, 0x54, 0x49, 0x4d, 0x45, 0x10, 0x26, 0x12, 0x23, 0x0a,
	0x1f, 0x43, 0x55, 0x53, 0x54, 0x4f, 0x4d, 0x45, 0x52, 0x5f, 0x49, 0x44, 0x45, 0x4e, 0x54, 0x49,
	0x46, 0x49, 0x45, 0x52, 0x5f, 0x4e, 0x4f, 0x54, 0x5f, 0x41, 0x4c, 0x4c, 0x4f, 0x57, 0x45, 0x44,
	0x10, 0x27, 0x12, 0x13, 0x0a, 0x0f, 0x49, 0x4e, 0x56, 0x41, 0x4c, 0x49, 0x44, 0x5f, 0x49, 0x54,
	0x45, 0x4d, 0x5f, 0x49, 0x44, 0x10, 0x28, 0x42, 0xfc, 0x01, 0x0a, 0x23, 0x63, 0x6f, 0x6d, 0x2e,
	0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x61, 0x64, 0x73, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x61, 0x64, 0x73, 0x2e, 0x76, 0x31, 0x30, 0x2e, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x73, 0x42,
	0x1c, 0x4f, 0x66, 0x66, 0x6c, 0x69, 0x6e, 0x65, 0x55, 0x73, 0x65, 0x72, 0x44, 0x61, 0x74, 0x61,
	0x4a, 0x6f, 0x62, 0x45, 0x72, 0x72, 0x6f, 0x72, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a,
	0x45, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x67, 0x6f, 0x6c, 0x61, 0x6e, 0x67, 0x2e, 0x6f,
	0x72, 0x67, 0x2f, 0x67, 0x65, 0x6e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x67, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x61, 0x70, 0x69, 0x73, 0x2f, 0x61, 0x64, 0x73, 0x2f, 0x67, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x61, 0x64, 0x73, 0x2f, 0x76, 0x31, 0x30, 0x2f, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x73, 0x3b,
	0x65, 0x72, 0x72, 0x6f, 0x72, 0x73, 0xa2, 0x02, 0x03, 0x47, 0x41, 0x41, 0xaa, 0x02, 0x1f, 0x47,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x41, 0x64, 0x73, 0x2e, 0x47, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x41, 0x64, 0x73, 0x2e, 0x56, 0x31, 0x30, 0x2e, 0x45, 0x72, 0x72, 0x6f, 0x72, 0x73, 0xca, 0x02,
	0x1f, 0x47, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x5c, 0x41, 0x64, 0x73, 0x5c, 0x47, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x41, 0x64, 0x73, 0x5c, 0x56, 0x31, 0x30, 0x5c, 0x45, 0x72, 0x72, 0x6f, 0x72, 0x73,
	0xea, 0x02, 0x23, 0x47, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x3a, 0x3a, 0x41, 0x64, 0x73, 0x3a, 0x3a,
	0x47, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x41, 0x64, 0x73, 0x3a, 0x3a, 0x56, 0x31, 0x30, 0x3a, 0x3a,
	0x45, 0x72, 0x72, 0x6f, 0x72, 0x73, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_errors_offline_user_data_job_error_proto_rawDescOnce sync.Once
	file_errors_offline_user_data_job_error_proto_rawDescData = file_errors_offline_user_data_job_error_proto_rawDesc
)

func file_errors_offline_user_data_job_error_proto_rawDescGZIP() []byte {
	file_errors_offline_user_data_job_error_proto_rawDescOnce.Do(func() {
		file_errors_offline_user_data_job_error_proto_rawDescData = protoimpl.X.CompressGZIP(file_errors_offline_user_data_job_error_proto_rawDescData)
	})
	return file_errors_offline_user_data_job_error_proto_rawDescData
}

var file_errors_offline_user_data_job_error_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_errors_offline_user_data_job_error_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_errors_offline_user_data_job_error_proto_goTypes = []interface{}{
	(OfflineUserDataJobErrorEnum_OfflineUserDataJobError)(0), // 0: google.ads.googleads.v10.errors.OfflineUserDataJobErrorEnum.OfflineUserDataJobError
	(*OfflineUserDataJobErrorEnum)(nil),                      // 1: google.ads.googleads.v10.errors.OfflineUserDataJobErrorEnum
}
var file_errors_offline_user_data_job_error_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_errors_offline_user_data_job_error_proto_init() }
func file_errors_offline_user_data_job_error_proto_init() {
	if File_errors_offline_user_data_job_error_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_errors_offline_user_data_job_error_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*OfflineUserDataJobErrorEnum); i {
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
			RawDescriptor: file_errors_offline_user_data_job_error_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_errors_offline_user_data_job_error_proto_goTypes,
		DependencyIndexes: file_errors_offline_user_data_job_error_proto_depIdxs,
		EnumInfos:         file_errors_offline_user_data_job_error_proto_enumTypes,
		MessageInfos:      file_errors_offline_user_data_job_error_proto_msgTypes,
	}.Build()
	File_errors_offline_user_data_job_error_proto = out.File
	file_errors_offline_user_data_job_error_proto_rawDesc = nil
	file_errors_offline_user_data_job_error_proto_goTypes = nil
	file_errors_offline_user_data_job_error_proto_depIdxs = nil
}
