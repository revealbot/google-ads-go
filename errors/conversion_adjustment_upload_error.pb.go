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
// source: google/ads/googleads/v15/errors/conversion_adjustment_upload_error.proto

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

// Enum describing possible conversion adjustment upload errors.
type ConversionAdjustmentUploadErrorEnum_ConversionAdjustmentUploadError int32

const (
	// Not specified.
	ConversionAdjustmentUploadErrorEnum_UNSPECIFIED ConversionAdjustmentUploadErrorEnum_ConversionAdjustmentUploadError = 0
	// Used for return value only. Represents value unknown in this version.
	ConversionAdjustmentUploadErrorEnum_UNKNOWN ConversionAdjustmentUploadErrorEnum_ConversionAdjustmentUploadError = 1
	// Can't import events to a conversion action that was just created. Try
	// importing again in 6 hours.
	ConversionAdjustmentUploadErrorEnum_TOO_RECENT_CONVERSION_ACTION ConversionAdjustmentUploadErrorEnum_ConversionAdjustmentUploadError = 2
	// The conversion was already retracted. This adjustment was not processed.
	ConversionAdjustmentUploadErrorEnum_CONVERSION_ALREADY_RETRACTED ConversionAdjustmentUploadErrorEnum_ConversionAdjustmentUploadError = 4
	// The conversion for this conversion action and conversion identifier can't
	// be found. Make sure your conversion identifiers are associated with the
	// correct conversion action and try again.
	ConversionAdjustmentUploadErrorEnum_CONVERSION_NOT_FOUND ConversionAdjustmentUploadErrorEnum_ConversionAdjustmentUploadError = 5
	// Adjustment can't be made to a conversion that occurred more than 54 days
	// ago.
	ConversionAdjustmentUploadErrorEnum_CONVERSION_EXPIRED ConversionAdjustmentUploadErrorEnum_ConversionAdjustmentUploadError = 6
	// Adjustment has an `adjustment_date_time` that occurred before the
	// associated conversion. Make sure your `adjustment_date_time` is correct
	// and try again.
	ConversionAdjustmentUploadErrorEnum_ADJUSTMENT_PRECEDES_CONVERSION ConversionAdjustmentUploadErrorEnum_ConversionAdjustmentUploadError = 7
	// More recent adjustment `adjustment_date_time` has already been reported
	// for the associated conversion.  Make sure your adjustment
	// `adjustment_date_time` is correct and try again.
	ConversionAdjustmentUploadErrorEnum_MORE_RECENT_RESTATEMENT_FOUND ConversionAdjustmentUploadErrorEnum_ConversionAdjustmentUploadError = 8
	// Adjustment can't be recorded because the conversion occurred too
	// recently. Try adjusting a conversion that occurred at least 24 hours ago.
	ConversionAdjustmentUploadErrorEnum_TOO_RECENT_CONVERSION ConversionAdjustmentUploadErrorEnum_ConversionAdjustmentUploadError = 9
	// Can't make an adjustment to a conversion that is set up to use the
	// default value. Check your conversion action value setting and try again.
	ConversionAdjustmentUploadErrorEnum_CANNOT_RESTATE_CONVERSION_ACTION_THAT_ALWAYS_USES_DEFAULT_CONVERSION_VALUE ConversionAdjustmentUploadErrorEnum_ConversionAdjustmentUploadError = 10
	// Try uploading fewer than 2001 adjustments in a single API request.
	ConversionAdjustmentUploadErrorEnum_TOO_MANY_ADJUSTMENTS_IN_REQUEST ConversionAdjustmentUploadErrorEnum_ConversionAdjustmentUploadError = 11
	// The conversion has already been adjusted the maximum number of times.
	// Make sure you're only making necessary adjustment to existing conversion.
	ConversionAdjustmentUploadErrorEnum_TOO_MANY_ADJUSTMENTS ConversionAdjustmentUploadErrorEnum_ConversionAdjustmentUploadError = 12
	// The conversion has prior a restatement with the same
	// `adjustment_date_time`. Make sure your adjustment has the correct and
	// unique `adjustment_date_time` and try again.
	ConversionAdjustmentUploadErrorEnum_RESTATEMENT_ALREADY_EXISTS ConversionAdjustmentUploadErrorEnum_ConversionAdjustmentUploadError = 13
	// Imported adjustment has a duplicate conversion adjustment with same
	// `adjustment_date_time`. Make sure your adjustment has the correct
	// `adjustment_date_time` and try again.
	ConversionAdjustmentUploadErrorEnum_DUPLICATE_ADJUSTMENT_IN_REQUEST ConversionAdjustmentUploadErrorEnum_ConversionAdjustmentUploadError = 14
	// Make sure you agree to the customer data processing terms in conversion
	// settings and try again.
	ConversionAdjustmentUploadErrorEnum_CUSTOMER_NOT_ACCEPTED_CUSTOMER_DATA_TERMS ConversionAdjustmentUploadErrorEnum_ConversionAdjustmentUploadError = 15
	// Can't use enhanced conversions with the specified conversion action.
	ConversionAdjustmentUploadErrorEnum_CONVERSION_ACTION_NOT_ELIGIBLE_FOR_ENHANCEMENT ConversionAdjustmentUploadErrorEnum_ConversionAdjustmentUploadError = 16
	// Make sure you hash user provided data using SHA-256 and ensure you are
	// normalizing according to the guidelines.
	ConversionAdjustmentUploadErrorEnum_INVALID_USER_IDENTIFIER ConversionAdjustmentUploadErrorEnum_ConversionAdjustmentUploadError = 17
	// Use user provided data such as emails or phone numbers hashed using
	// SHA-256 and try again.
	ConversionAdjustmentUploadErrorEnum_UNSUPPORTED_USER_IDENTIFIER ConversionAdjustmentUploadErrorEnum_ConversionAdjustmentUploadError = 18
	// Cannot set both gclid_date_time_pair and order_id. Use only 1 type and
	// try again.
	ConversionAdjustmentUploadErrorEnum_GCLID_DATE_TIME_PAIR_AND_ORDER_ID_BOTH_SET ConversionAdjustmentUploadErrorEnum_ConversionAdjustmentUploadError = 20
	// Conversion already has enhancements with the same Order ID and conversion
	// action. Make sure your data is correctly configured and try again.
	ConversionAdjustmentUploadErrorEnum_CONVERSION_ALREADY_ENHANCED ConversionAdjustmentUploadErrorEnum_ConversionAdjustmentUploadError = 21
	// Multiple enhancements have the same conversion action and Order ID.  Make
	// sure your data is correctly configured and try again.
	ConversionAdjustmentUploadErrorEnum_DUPLICATE_ENHANCEMENT_IN_REQUEST ConversionAdjustmentUploadErrorEnum_ConversionAdjustmentUploadError = 22
	// Enhanced conversions can't be used for this account because of Google
	// customer data policies. Contact your Google representative.
	ConversionAdjustmentUploadErrorEnum_CUSTOMER_DATA_POLICY_PROHIBITS_ENHANCEMENT ConversionAdjustmentUploadErrorEnum_ConversionAdjustmentUploadError = 23
	// Adjustment for website conversion requires Order ID (ie, transaction ID).
	// Make sure your website tags capture Order IDs and you send the same Order
	// IDs with your adjustment.
	ConversionAdjustmentUploadErrorEnum_MISSING_ORDER_ID_FOR_WEBPAGE ConversionAdjustmentUploadErrorEnum_ConversionAdjustmentUploadError = 24
	// Can't use adjustment with Order IDs containing personally-identifiable
	// information (PII).
	ConversionAdjustmentUploadErrorEnum_ORDER_ID_CONTAINS_PII ConversionAdjustmentUploadErrorEnum_ConversionAdjustmentUploadError = 25
	// The provided job id in the request is not within the allowed range. A job
	// ID must be a positive integer in the range [1, 2^31).
	ConversionAdjustmentUploadErrorEnum_INVALID_JOB_ID ConversionAdjustmentUploadErrorEnum_ConversionAdjustmentUploadError = 26
	// The conversion action specified in the adjustment request cannot be
	// found. Make sure it's available in this account.
	ConversionAdjustmentUploadErrorEnum_NO_CONVERSION_ACTION_FOUND ConversionAdjustmentUploadErrorEnum_ConversionAdjustmentUploadError = 27
	// The type of the conversion action specified in the adjustment request
	// isn't supported for uploading adjustments. A conversion adjustment of
	// type `RETRACTION` or `RESTATEMENT` is only permitted for conversion
	// actions of type `SALESFORCE`, `UPLOAD` or `WEBPAGE`. A conversion
	// adjustment of type `ENHANCEMENT` is only permitted for conversion
	// actions of type `WEBPAGE`.
	ConversionAdjustmentUploadErrorEnum_INVALID_CONVERSION_ACTION_TYPE ConversionAdjustmentUploadErrorEnum_ConversionAdjustmentUploadError = 28
)

// Enum value maps for ConversionAdjustmentUploadErrorEnum_ConversionAdjustmentUploadError.
var (
	ConversionAdjustmentUploadErrorEnum_ConversionAdjustmentUploadError_name = map[int32]string{
		0:  "UNSPECIFIED",
		1:  "UNKNOWN",
		2:  "TOO_RECENT_CONVERSION_ACTION",
		4:  "CONVERSION_ALREADY_RETRACTED",
		5:  "CONVERSION_NOT_FOUND",
		6:  "CONVERSION_EXPIRED",
		7:  "ADJUSTMENT_PRECEDES_CONVERSION",
		8:  "MORE_RECENT_RESTATEMENT_FOUND",
		9:  "TOO_RECENT_CONVERSION",
		10: "CANNOT_RESTATE_CONVERSION_ACTION_THAT_ALWAYS_USES_DEFAULT_CONVERSION_VALUE",
		11: "TOO_MANY_ADJUSTMENTS_IN_REQUEST",
		12: "TOO_MANY_ADJUSTMENTS",
		13: "RESTATEMENT_ALREADY_EXISTS",
		14: "DUPLICATE_ADJUSTMENT_IN_REQUEST",
		15: "CUSTOMER_NOT_ACCEPTED_CUSTOMER_DATA_TERMS",
		16: "CONVERSION_ACTION_NOT_ELIGIBLE_FOR_ENHANCEMENT",
		17: "INVALID_USER_IDENTIFIER",
		18: "UNSUPPORTED_USER_IDENTIFIER",
		20: "GCLID_DATE_TIME_PAIR_AND_ORDER_ID_BOTH_SET",
		21: "CONVERSION_ALREADY_ENHANCED",
		22: "DUPLICATE_ENHANCEMENT_IN_REQUEST",
		23: "CUSTOMER_DATA_POLICY_PROHIBITS_ENHANCEMENT",
		24: "MISSING_ORDER_ID_FOR_WEBPAGE",
		25: "ORDER_ID_CONTAINS_PII",
		26: "INVALID_JOB_ID",
		27: "NO_CONVERSION_ACTION_FOUND",
		28: "INVALID_CONVERSION_ACTION_TYPE",
	}
	ConversionAdjustmentUploadErrorEnum_ConversionAdjustmentUploadError_value = map[string]int32{
		"UNSPECIFIED":                    0,
		"UNKNOWN":                        1,
		"TOO_RECENT_CONVERSION_ACTION":   2,
		"CONVERSION_ALREADY_RETRACTED":   4,
		"CONVERSION_NOT_FOUND":           5,
		"CONVERSION_EXPIRED":             6,
		"ADJUSTMENT_PRECEDES_CONVERSION": 7,
		"MORE_RECENT_RESTATEMENT_FOUND":  8,
		"TOO_RECENT_CONVERSION":          9,
		"CANNOT_RESTATE_CONVERSION_ACTION_THAT_ALWAYS_USES_DEFAULT_CONVERSION_VALUE": 10,
		"TOO_MANY_ADJUSTMENTS_IN_REQUEST":                                            11,
		"TOO_MANY_ADJUSTMENTS":                                                       12,
		"RESTATEMENT_ALREADY_EXISTS":                                                 13,
		"DUPLICATE_ADJUSTMENT_IN_REQUEST":                                            14,
		"CUSTOMER_NOT_ACCEPTED_CUSTOMER_DATA_TERMS":                                  15,
		"CONVERSION_ACTION_NOT_ELIGIBLE_FOR_ENHANCEMENT":                             16,
		"INVALID_USER_IDENTIFIER":                                                    17,
		"UNSUPPORTED_USER_IDENTIFIER":                                                18,
		"GCLID_DATE_TIME_PAIR_AND_ORDER_ID_BOTH_SET":                                 20,
		"CONVERSION_ALREADY_ENHANCED":                                                21,
		"DUPLICATE_ENHANCEMENT_IN_REQUEST":                                           22,
		"CUSTOMER_DATA_POLICY_PROHIBITS_ENHANCEMENT":                                 23,
		"MISSING_ORDER_ID_FOR_WEBPAGE":                                               24,
		"ORDER_ID_CONTAINS_PII":                                                      25,
		"INVALID_JOB_ID":                                                             26,
		"NO_CONVERSION_ACTION_FOUND":                                                 27,
		"INVALID_CONVERSION_ACTION_TYPE":                                             28,
	}
)

func (x ConversionAdjustmentUploadErrorEnum_ConversionAdjustmentUploadError) Enum() *ConversionAdjustmentUploadErrorEnum_ConversionAdjustmentUploadError {
	p := new(ConversionAdjustmentUploadErrorEnum_ConversionAdjustmentUploadError)
	*p = x
	return p
}

func (x ConversionAdjustmentUploadErrorEnum_ConversionAdjustmentUploadError) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (ConversionAdjustmentUploadErrorEnum_ConversionAdjustmentUploadError) Descriptor() protoreflect.EnumDescriptor {
	return file_errors_conversion_adjustment_upload_error_proto_enumTypes[0].Descriptor()
}

func (ConversionAdjustmentUploadErrorEnum_ConversionAdjustmentUploadError) Type() protoreflect.EnumType {
	return &file_errors_conversion_adjustment_upload_error_proto_enumTypes[0]
}

func (x ConversionAdjustmentUploadErrorEnum_ConversionAdjustmentUploadError) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use ConversionAdjustmentUploadErrorEnum_ConversionAdjustmentUploadError.Descriptor instead.
func (ConversionAdjustmentUploadErrorEnum_ConversionAdjustmentUploadError) EnumDescriptor() ([]byte, []int) {
	return file_errors_conversion_adjustment_upload_error_proto_rawDescGZIP(), []int{0, 0}
}

// Container for enum describing possible conversion adjustment upload errors.
type ConversionAdjustmentUploadErrorEnum struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *ConversionAdjustmentUploadErrorEnum) Reset() {
	*x = ConversionAdjustmentUploadErrorEnum{}
	if protoimpl.UnsafeEnabled {
		mi := &file_errors_conversion_adjustment_upload_error_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ConversionAdjustmentUploadErrorEnum) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ConversionAdjustmentUploadErrorEnum) ProtoMessage() {}

func (x *ConversionAdjustmentUploadErrorEnum) ProtoReflect() protoreflect.Message {
	mi := &file_errors_conversion_adjustment_upload_error_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ConversionAdjustmentUploadErrorEnum.ProtoReflect.Descriptor instead.
func (*ConversionAdjustmentUploadErrorEnum) Descriptor() ([]byte, []int) {
	return file_errors_conversion_adjustment_upload_error_proto_rawDescGZIP(), []int{0}
}

var File_errors_conversion_adjustment_upload_error_proto protoreflect.FileDescriptor

var file_errors_conversion_adjustment_upload_error_proto_rawDesc = []byte{
	0x0a, 0x48, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x64, 0x73, 0x2f, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x61, 0x64, 0x73, 0x2f, 0x76, 0x31, 0x35, 0x2f, 0x65, 0x72, 0x72, 0x6f, 0x72,
	0x73, 0x2f, 0x63, 0x6f, 0x6e, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x5f, 0x61, 0x64, 0x6a,
	0x75, 0x73, 0x74, 0x6d, 0x65, 0x6e, 0x74, 0x5f, 0x75, 0x70, 0x6c, 0x6f, 0x61, 0x64, 0x5f, 0x65,
	0x72, 0x72, 0x6f, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x1f, 0x67, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x2e, 0x61, 0x64, 0x73, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61, 0x64, 0x73,
	0x2e, 0x76, 0x31, 0x35, 0x2e, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x73, 0x22, 0xf0, 0x07, 0x0a, 0x23,
	0x43, 0x6f, 0x6e, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x41, 0x64, 0x6a, 0x75, 0x73, 0x74,
	0x6d, 0x65, 0x6e, 0x74, 0x55, 0x70, 0x6c, 0x6f, 0x61, 0x64, 0x45, 0x72, 0x72, 0x6f, 0x72, 0x45,
	0x6e, 0x75, 0x6d, 0x22, 0xc8, 0x07, 0x0a, 0x1f, 0x43, 0x6f, 0x6e, 0x76, 0x65, 0x72, 0x73, 0x69,
	0x6f, 0x6e, 0x41, 0x64, 0x6a, 0x75, 0x73, 0x74, 0x6d, 0x65, 0x6e, 0x74, 0x55, 0x70, 0x6c, 0x6f,
	0x61, 0x64, 0x45, 0x72, 0x72, 0x6f, 0x72, 0x12, 0x0f, 0x0a, 0x0b, 0x55, 0x4e, 0x53, 0x50, 0x45,
	0x43, 0x49, 0x46, 0x49, 0x45, 0x44, 0x10, 0x00, 0x12, 0x0b, 0x0a, 0x07, 0x55, 0x4e, 0x4b, 0x4e,
	0x4f, 0x57, 0x4e, 0x10, 0x01, 0x12, 0x20, 0x0a, 0x1c, 0x54, 0x4f, 0x4f, 0x5f, 0x52, 0x45, 0x43,
	0x45, 0x4e, 0x54, 0x5f, 0x43, 0x4f, 0x4e, 0x56, 0x45, 0x52, 0x53, 0x49, 0x4f, 0x4e, 0x5f, 0x41,
	0x43, 0x54, 0x49, 0x4f, 0x4e, 0x10, 0x02, 0x12, 0x20, 0x0a, 0x1c, 0x43, 0x4f, 0x4e, 0x56, 0x45,
	0x52, 0x53, 0x49, 0x4f, 0x4e, 0x5f, 0x41, 0x4c, 0x52, 0x45, 0x41, 0x44, 0x59, 0x5f, 0x52, 0x45,
	0x54, 0x52, 0x41, 0x43, 0x54, 0x45, 0x44, 0x10, 0x04, 0x12, 0x18, 0x0a, 0x14, 0x43, 0x4f, 0x4e,
	0x56, 0x45, 0x52, 0x53, 0x49, 0x4f, 0x4e, 0x5f, 0x4e, 0x4f, 0x54, 0x5f, 0x46, 0x4f, 0x55, 0x4e,
	0x44, 0x10, 0x05, 0x12, 0x16, 0x0a, 0x12, 0x43, 0x4f, 0x4e, 0x56, 0x45, 0x52, 0x53, 0x49, 0x4f,
	0x4e, 0x5f, 0x45, 0x58, 0x50, 0x49, 0x52, 0x45, 0x44, 0x10, 0x06, 0x12, 0x22, 0x0a, 0x1e, 0x41,
	0x44, 0x4a, 0x55, 0x53, 0x54, 0x4d, 0x45, 0x4e, 0x54, 0x5f, 0x50, 0x52, 0x45, 0x43, 0x45, 0x44,
	0x45, 0x53, 0x5f, 0x43, 0x4f, 0x4e, 0x56, 0x45, 0x52, 0x53, 0x49, 0x4f, 0x4e, 0x10, 0x07, 0x12,
	0x21, 0x0a, 0x1d, 0x4d, 0x4f, 0x52, 0x45, 0x5f, 0x52, 0x45, 0x43, 0x45, 0x4e, 0x54, 0x5f, 0x52,
	0x45, 0x53, 0x54, 0x41, 0x54, 0x45, 0x4d, 0x45, 0x4e, 0x54, 0x5f, 0x46, 0x4f, 0x55, 0x4e, 0x44,
	0x10, 0x08, 0x12, 0x19, 0x0a, 0x15, 0x54, 0x4f, 0x4f, 0x5f, 0x52, 0x45, 0x43, 0x45, 0x4e, 0x54,
	0x5f, 0x43, 0x4f, 0x4e, 0x56, 0x45, 0x52, 0x53, 0x49, 0x4f, 0x4e, 0x10, 0x09, 0x12, 0x4e, 0x0a,
	0x4a, 0x43, 0x41, 0x4e, 0x4e, 0x4f, 0x54, 0x5f, 0x52, 0x45, 0x53, 0x54, 0x41, 0x54, 0x45, 0x5f,
	0x43, 0x4f, 0x4e, 0x56, 0x45, 0x52, 0x53, 0x49, 0x4f, 0x4e, 0x5f, 0x41, 0x43, 0x54, 0x49, 0x4f,
	0x4e, 0x5f, 0x54, 0x48, 0x41, 0x54, 0x5f, 0x41, 0x4c, 0x57, 0x41, 0x59, 0x53, 0x5f, 0x55, 0x53,
	0x45, 0x53, 0x5f, 0x44, 0x45, 0x46, 0x41, 0x55, 0x4c, 0x54, 0x5f, 0x43, 0x4f, 0x4e, 0x56, 0x45,
	0x52, 0x53, 0x49, 0x4f, 0x4e, 0x5f, 0x56, 0x41, 0x4c, 0x55, 0x45, 0x10, 0x0a, 0x12, 0x23, 0x0a,
	0x1f, 0x54, 0x4f, 0x4f, 0x5f, 0x4d, 0x41, 0x4e, 0x59, 0x5f, 0x41, 0x44, 0x4a, 0x55, 0x53, 0x54,
	0x4d, 0x45, 0x4e, 0x54, 0x53, 0x5f, 0x49, 0x4e, 0x5f, 0x52, 0x45, 0x51, 0x55, 0x45, 0x53, 0x54,
	0x10, 0x0b, 0x12, 0x18, 0x0a, 0x14, 0x54, 0x4f, 0x4f, 0x5f, 0x4d, 0x41, 0x4e, 0x59, 0x5f, 0x41,
	0x44, 0x4a, 0x55, 0x53, 0x54, 0x4d, 0x45, 0x4e, 0x54, 0x53, 0x10, 0x0c, 0x12, 0x1e, 0x0a, 0x1a,
	0x52, 0x45, 0x53, 0x54, 0x41, 0x54, 0x45, 0x4d, 0x45, 0x4e, 0x54, 0x5f, 0x41, 0x4c, 0x52, 0x45,
	0x41, 0x44, 0x59, 0x5f, 0x45, 0x58, 0x49, 0x53, 0x54, 0x53, 0x10, 0x0d, 0x12, 0x23, 0x0a, 0x1f,
	0x44, 0x55, 0x50, 0x4c, 0x49, 0x43, 0x41, 0x54, 0x45, 0x5f, 0x41, 0x44, 0x4a, 0x55, 0x53, 0x54,
	0x4d, 0x45, 0x4e, 0x54, 0x5f, 0x49, 0x4e, 0x5f, 0x52, 0x45, 0x51, 0x55, 0x45, 0x53, 0x54, 0x10,
	0x0e, 0x12, 0x2d, 0x0a, 0x29, 0x43, 0x55, 0x53, 0x54, 0x4f, 0x4d, 0x45, 0x52, 0x5f, 0x4e, 0x4f,
	0x54, 0x5f, 0x41, 0x43, 0x43, 0x45, 0x50, 0x54, 0x45, 0x44, 0x5f, 0x43, 0x55, 0x53, 0x54, 0x4f,
	0x4d, 0x45, 0x52, 0x5f, 0x44, 0x41, 0x54, 0x41, 0x5f, 0x54, 0x45, 0x52, 0x4d, 0x53, 0x10, 0x0f,
	0x12, 0x32, 0x0a, 0x2e, 0x43, 0x4f, 0x4e, 0x56, 0x45, 0x52, 0x53, 0x49, 0x4f, 0x4e, 0x5f, 0x41,
	0x43, 0x54, 0x49, 0x4f, 0x4e, 0x5f, 0x4e, 0x4f, 0x54, 0x5f, 0x45, 0x4c, 0x49, 0x47, 0x49, 0x42,
	0x4c, 0x45, 0x5f, 0x46, 0x4f, 0x52, 0x5f, 0x45, 0x4e, 0x48, 0x41, 0x4e, 0x43, 0x45, 0x4d, 0x45,
	0x4e, 0x54, 0x10, 0x10, 0x12, 0x1b, 0x0a, 0x17, 0x49, 0x4e, 0x56, 0x41, 0x4c, 0x49, 0x44, 0x5f,
	0x55, 0x53, 0x45, 0x52, 0x5f, 0x49, 0x44, 0x45, 0x4e, 0x54, 0x49, 0x46, 0x49, 0x45, 0x52, 0x10,
	0x11, 0x12, 0x1f, 0x0a, 0x1b, 0x55, 0x4e, 0x53, 0x55, 0x50, 0x50, 0x4f, 0x52, 0x54, 0x45, 0x44,
	0x5f, 0x55, 0x53, 0x45, 0x52, 0x5f, 0x49, 0x44, 0x45, 0x4e, 0x54, 0x49, 0x46, 0x49, 0x45, 0x52,
	0x10, 0x12, 0x12, 0x2e, 0x0a, 0x2a, 0x47, 0x43, 0x4c, 0x49, 0x44, 0x5f, 0x44, 0x41, 0x54, 0x45,
	0x5f, 0x54, 0x49, 0x4d, 0x45, 0x5f, 0x50, 0x41, 0x49, 0x52, 0x5f, 0x41, 0x4e, 0x44, 0x5f, 0x4f,
	0x52, 0x44, 0x45, 0x52, 0x5f, 0x49, 0x44, 0x5f, 0x42, 0x4f, 0x54, 0x48, 0x5f, 0x53, 0x45, 0x54,
	0x10, 0x14, 0x12, 0x1f, 0x0a, 0x1b, 0x43, 0x4f, 0x4e, 0x56, 0x45, 0x52, 0x53, 0x49, 0x4f, 0x4e,
	0x5f, 0x41, 0x4c, 0x52, 0x45, 0x41, 0x44, 0x59, 0x5f, 0x45, 0x4e, 0x48, 0x41, 0x4e, 0x43, 0x45,
	0x44, 0x10, 0x15, 0x12, 0x24, 0x0a, 0x20, 0x44, 0x55, 0x50, 0x4c, 0x49, 0x43, 0x41, 0x54, 0x45,
	0x5f, 0x45, 0x4e, 0x48, 0x41, 0x4e, 0x43, 0x45, 0x4d, 0x45, 0x4e, 0x54, 0x5f, 0x49, 0x4e, 0x5f,
	0x52, 0x45, 0x51, 0x55, 0x45, 0x53, 0x54, 0x10, 0x16, 0x12, 0x2e, 0x0a, 0x2a, 0x43, 0x55, 0x53,
	0x54, 0x4f, 0x4d, 0x45, 0x52, 0x5f, 0x44, 0x41, 0x54, 0x41, 0x5f, 0x50, 0x4f, 0x4c, 0x49, 0x43,
	0x59, 0x5f, 0x50, 0x52, 0x4f, 0x48, 0x49, 0x42, 0x49, 0x54, 0x53, 0x5f, 0x45, 0x4e, 0x48, 0x41,
	0x4e, 0x43, 0x45, 0x4d, 0x45, 0x4e, 0x54, 0x10, 0x17, 0x12, 0x20, 0x0a, 0x1c, 0x4d, 0x49, 0x53,
	0x53, 0x49, 0x4e, 0x47, 0x5f, 0x4f, 0x52, 0x44, 0x45, 0x52, 0x5f, 0x49, 0x44, 0x5f, 0x46, 0x4f,
	0x52, 0x5f, 0x57, 0x45, 0x42, 0x50, 0x41, 0x47, 0x45, 0x10, 0x18, 0x12, 0x19, 0x0a, 0x15, 0x4f,
	0x52, 0x44, 0x45, 0x52, 0x5f, 0x49, 0x44, 0x5f, 0x43, 0x4f, 0x4e, 0x54, 0x41, 0x49, 0x4e, 0x53,
	0x5f, 0x50, 0x49, 0x49, 0x10, 0x19, 0x12, 0x12, 0x0a, 0x0e, 0x49, 0x4e, 0x56, 0x41, 0x4c, 0x49,
	0x44, 0x5f, 0x4a, 0x4f, 0x42, 0x5f, 0x49, 0x44, 0x10, 0x1a, 0x12, 0x1e, 0x0a, 0x1a, 0x4e, 0x4f,
	0x5f, 0x43, 0x4f, 0x4e, 0x56, 0x45, 0x52, 0x53, 0x49, 0x4f, 0x4e, 0x5f, 0x41, 0x43, 0x54, 0x49,
	0x4f, 0x4e, 0x5f, 0x46, 0x4f, 0x55, 0x4e, 0x44, 0x10, 0x1b, 0x12, 0x22, 0x0a, 0x1e, 0x49, 0x4e,
	0x56, 0x41, 0x4c, 0x49, 0x44, 0x5f, 0x43, 0x4f, 0x4e, 0x56, 0x45, 0x52, 0x53, 0x49, 0x4f, 0x4e,
	0x5f, 0x41, 0x43, 0x54, 0x49, 0x4f, 0x4e, 0x5f, 0x54, 0x59, 0x50, 0x45, 0x10, 0x1c, 0x42, 0x84,
	0x02, 0x0a, 0x23, 0x63, 0x6f, 0x6d, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x61, 0x64,
	0x73, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61, 0x64, 0x73, 0x2e, 0x76, 0x31, 0x35, 0x2e,
	0x65, 0x72, 0x72, 0x6f, 0x72, 0x73, 0x42, 0x24, 0x43, 0x6f, 0x6e, 0x76, 0x65, 0x72, 0x73, 0x69,
	0x6f, 0x6e, 0x41, 0x64, 0x6a, 0x75, 0x73, 0x74, 0x6d, 0x65, 0x6e, 0x74, 0x55, 0x70, 0x6c, 0x6f,
	0x61, 0x64, 0x45, 0x72, 0x72, 0x6f, 0x72, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a, 0x45,
	0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x67, 0x6f, 0x6c, 0x61, 0x6e, 0x67, 0x2e, 0x6f, 0x72,
	0x67, 0x2f, 0x67, 0x65, 0x6e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x67, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x61, 0x70, 0x69, 0x73, 0x2f, 0x61, 0x64, 0x73, 0x2f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x61, 0x64, 0x73, 0x2f, 0x76, 0x31, 0x35, 0x2f, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x73, 0x3b, 0x65,
	0x72, 0x72, 0x6f, 0x72, 0x73, 0xa2, 0x02, 0x03, 0x47, 0x41, 0x41, 0xaa, 0x02, 0x1f, 0x47, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x41, 0x64, 0x73, 0x2e, 0x47, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x41,
	0x64, 0x73, 0x2e, 0x56, 0x31, 0x35, 0x2e, 0x45, 0x72, 0x72, 0x6f, 0x72, 0x73, 0xca, 0x02, 0x1f,
	0x47, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x5c, 0x41, 0x64, 0x73, 0x5c, 0x47, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x41, 0x64, 0x73, 0x5c, 0x56, 0x31, 0x35, 0x5c, 0x45, 0x72, 0x72, 0x6f, 0x72, 0x73, 0xea,
	0x02, 0x23, 0x47, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x3a, 0x3a, 0x41, 0x64, 0x73, 0x3a, 0x3a, 0x47,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x41, 0x64, 0x73, 0x3a, 0x3a, 0x56, 0x31, 0x35, 0x3a, 0x3a, 0x45,
	0x72, 0x72, 0x6f, 0x72, 0x73, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_errors_conversion_adjustment_upload_error_proto_rawDescOnce sync.Once
	file_errors_conversion_adjustment_upload_error_proto_rawDescData = file_errors_conversion_adjustment_upload_error_proto_rawDesc
)

func file_errors_conversion_adjustment_upload_error_proto_rawDescGZIP() []byte {
	file_errors_conversion_adjustment_upload_error_proto_rawDescOnce.Do(func() {
		file_errors_conversion_adjustment_upload_error_proto_rawDescData = protoimpl.X.CompressGZIP(file_errors_conversion_adjustment_upload_error_proto_rawDescData)
	})
	return file_errors_conversion_adjustment_upload_error_proto_rawDescData
}

var file_errors_conversion_adjustment_upload_error_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_errors_conversion_adjustment_upload_error_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_errors_conversion_adjustment_upload_error_proto_goTypes = []interface{}{
	(ConversionAdjustmentUploadErrorEnum_ConversionAdjustmentUploadError)(0), // 0: google.ads.googleads.v15.errors.ConversionAdjustmentUploadErrorEnum.ConversionAdjustmentUploadError
	(*ConversionAdjustmentUploadErrorEnum)(nil),                              // 1: google.ads.googleads.v15.errors.ConversionAdjustmentUploadErrorEnum
}
var file_errors_conversion_adjustment_upload_error_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_errors_conversion_adjustment_upload_error_proto_init() }
func file_errors_conversion_adjustment_upload_error_proto_init() {
	if File_errors_conversion_adjustment_upload_error_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_errors_conversion_adjustment_upload_error_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ConversionAdjustmentUploadErrorEnum); i {
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
			RawDescriptor: file_errors_conversion_adjustment_upload_error_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_errors_conversion_adjustment_upload_error_proto_goTypes,
		DependencyIndexes: file_errors_conversion_adjustment_upload_error_proto_depIdxs,
		EnumInfos:         file_errors_conversion_adjustment_upload_error_proto_enumTypes,
		MessageInfos:      file_errors_conversion_adjustment_upload_error_proto_msgTypes,
	}.Build()
	File_errors_conversion_adjustment_upload_error_proto = out.File
	file_errors_conversion_adjustment_upload_error_proto_rawDesc = nil
	file_errors_conversion_adjustment_upload_error_proto_goTypes = nil
	file_errors_conversion_adjustment_upload_error_proto_depIdxs = nil
}
