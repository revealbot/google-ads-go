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
// source: google/ads/googleads/v19/enums/bidding_strategy_system_status.proto

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

// The possible system statuses of a BiddingStrategy.
type BiddingStrategySystemStatusEnum_BiddingStrategySystemStatus int32

const (
	// Signals that an unexpected error occurred, for example, no bidding
	// strategy type was found, or no status information was found.
	BiddingStrategySystemStatusEnum_UNSPECIFIED BiddingStrategySystemStatusEnum_BiddingStrategySystemStatus = 0
	// Used for return value only. Represents value unknown in this version.
	BiddingStrategySystemStatusEnum_UNKNOWN BiddingStrategySystemStatusEnum_BiddingStrategySystemStatus = 1
	// The bid strategy is active, and AdWords cannot find any specific issues
	// with the strategy.
	BiddingStrategySystemStatusEnum_ENABLED BiddingStrategySystemStatusEnum_BiddingStrategySystemStatus = 2
	// The bid strategy is learning because it has been recently created or
	// recently reactivated.
	BiddingStrategySystemStatusEnum_LEARNING_NEW BiddingStrategySystemStatusEnum_BiddingStrategySystemStatus = 3
	// The bid strategy is learning because of a recent setting change.
	BiddingStrategySystemStatusEnum_LEARNING_SETTING_CHANGE BiddingStrategySystemStatusEnum_BiddingStrategySystemStatus = 4
	// The bid strategy is learning because of a recent budget change.
	BiddingStrategySystemStatusEnum_LEARNING_BUDGET_CHANGE BiddingStrategySystemStatusEnum_BiddingStrategySystemStatus = 5
	// The bid strategy is learning because of recent change in number of
	// campaigns, ad groups or keywords attached to it.
	BiddingStrategySystemStatusEnum_LEARNING_COMPOSITION_CHANGE BiddingStrategySystemStatusEnum_BiddingStrategySystemStatus = 6
	// The bid strategy depends on conversion reporting and the customer
	// recently modified conversion types that were relevant to the
	// bid strategy.
	BiddingStrategySystemStatusEnum_LEARNING_CONVERSION_TYPE_CHANGE BiddingStrategySystemStatusEnum_BiddingStrategySystemStatus = 7
	// The bid strategy depends on conversion reporting and the customer
	// recently changed their conversion settings.
	BiddingStrategySystemStatusEnum_LEARNING_CONVERSION_SETTING_CHANGE BiddingStrategySystemStatusEnum_BiddingStrategySystemStatus = 8
	// The bid strategy is limited by its bid ceiling.
	BiddingStrategySystemStatusEnum_LIMITED_BY_CPC_BID_CEILING BiddingStrategySystemStatusEnum_BiddingStrategySystemStatus = 9
	// The bid strategy is limited by its bid floor.
	BiddingStrategySystemStatusEnum_LIMITED_BY_CPC_BID_FLOOR BiddingStrategySystemStatusEnum_BiddingStrategySystemStatus = 10
	// The bid strategy is limited because there was not enough conversion
	// traffic over the past weeks.
	BiddingStrategySystemStatusEnum_LIMITED_BY_DATA BiddingStrategySystemStatusEnum_BiddingStrategySystemStatus = 11
	// A significant fraction of keywords in this bid strategy are limited by
	// budget.
	BiddingStrategySystemStatusEnum_LIMITED_BY_BUDGET BiddingStrategySystemStatusEnum_BiddingStrategySystemStatus = 12
	// The bid strategy cannot reach its target spend because its spend has
	// been de-prioritized.
	BiddingStrategySystemStatusEnum_LIMITED_BY_LOW_PRIORITY_SPEND BiddingStrategySystemStatusEnum_BiddingStrategySystemStatus = 13
	// A significant fraction of keywords in this bid strategy have a low
	// Quality Score.
	BiddingStrategySystemStatusEnum_LIMITED_BY_LOW_QUALITY BiddingStrategySystemStatusEnum_BiddingStrategySystemStatus = 14
	// The bid strategy cannot fully spend its budget because of narrow
	// targeting.
	BiddingStrategySystemStatusEnum_LIMITED_BY_INVENTORY BiddingStrategySystemStatusEnum_BiddingStrategySystemStatus = 15
	// Missing conversion tracking (no pings present) and/or remarketing lists
	// for SSC.
	BiddingStrategySystemStatusEnum_MISCONFIGURED_ZERO_ELIGIBILITY BiddingStrategySystemStatusEnum_BiddingStrategySystemStatus = 16
	// The bid strategy depends on conversion reporting and the customer is
	// lacking conversion types that might be reported against this strategy.
	BiddingStrategySystemStatusEnum_MISCONFIGURED_CONVERSION_TYPES BiddingStrategySystemStatusEnum_BiddingStrategySystemStatus = 17
	// The bid strategy depends on conversion reporting and the customer's
	// conversion settings are misconfigured.
	BiddingStrategySystemStatusEnum_MISCONFIGURED_CONVERSION_SETTINGS BiddingStrategySystemStatusEnum_BiddingStrategySystemStatus = 18
	// There are campaigns outside the bid strategy that share budgets with
	// campaigns included in the strategy.
	BiddingStrategySystemStatusEnum_MISCONFIGURED_SHARED_BUDGET BiddingStrategySystemStatusEnum_BiddingStrategySystemStatus = 19
	// The campaign has an invalid strategy type and is not serving.
	BiddingStrategySystemStatusEnum_MISCONFIGURED_STRATEGY_TYPE BiddingStrategySystemStatusEnum_BiddingStrategySystemStatus = 20
	// The bid strategy is not active. Either there are no active campaigns,
	// ad groups or keywords attached to the bid strategy. Or there are no
	// active budgets connected to the bid strategy.
	BiddingStrategySystemStatusEnum_PAUSED BiddingStrategySystemStatusEnum_BiddingStrategySystemStatus = 21
	// This bid strategy currently does not support status reporting.
	BiddingStrategySystemStatusEnum_UNAVAILABLE BiddingStrategySystemStatusEnum_BiddingStrategySystemStatus = 22
	// There were multiple LEARNING_* system statuses for this bid strategy
	// during the time in question.
	BiddingStrategySystemStatusEnum_MULTIPLE_LEARNING BiddingStrategySystemStatusEnum_BiddingStrategySystemStatus = 23
	// There were multiple LIMITED_* system statuses for this bid strategy
	// during the time in question.
	BiddingStrategySystemStatusEnum_MULTIPLE_LIMITED BiddingStrategySystemStatusEnum_BiddingStrategySystemStatus = 24
	// There were multiple MISCONFIGURED_* system statuses for this bid strategy
	// during the time in question.
	BiddingStrategySystemStatusEnum_MULTIPLE_MISCONFIGURED BiddingStrategySystemStatusEnum_BiddingStrategySystemStatus = 25
	// There were multiple system statuses for this bid strategy during the
	// time in question.
	BiddingStrategySystemStatusEnum_MULTIPLE BiddingStrategySystemStatusEnum_BiddingStrategySystemStatus = 26
)

// Enum value maps for BiddingStrategySystemStatusEnum_BiddingStrategySystemStatus.
var (
	BiddingStrategySystemStatusEnum_BiddingStrategySystemStatus_name = map[int32]string{
		0:  "UNSPECIFIED",
		1:  "UNKNOWN",
		2:  "ENABLED",
		3:  "LEARNING_NEW",
		4:  "LEARNING_SETTING_CHANGE",
		5:  "LEARNING_BUDGET_CHANGE",
		6:  "LEARNING_COMPOSITION_CHANGE",
		7:  "LEARNING_CONVERSION_TYPE_CHANGE",
		8:  "LEARNING_CONVERSION_SETTING_CHANGE",
		9:  "LIMITED_BY_CPC_BID_CEILING",
		10: "LIMITED_BY_CPC_BID_FLOOR",
		11: "LIMITED_BY_DATA",
		12: "LIMITED_BY_BUDGET",
		13: "LIMITED_BY_LOW_PRIORITY_SPEND",
		14: "LIMITED_BY_LOW_QUALITY",
		15: "LIMITED_BY_INVENTORY",
		16: "MISCONFIGURED_ZERO_ELIGIBILITY",
		17: "MISCONFIGURED_CONVERSION_TYPES",
		18: "MISCONFIGURED_CONVERSION_SETTINGS",
		19: "MISCONFIGURED_SHARED_BUDGET",
		20: "MISCONFIGURED_STRATEGY_TYPE",
		21: "PAUSED",
		22: "UNAVAILABLE",
		23: "MULTIPLE_LEARNING",
		24: "MULTIPLE_LIMITED",
		25: "MULTIPLE_MISCONFIGURED",
		26: "MULTIPLE",
	}
	BiddingStrategySystemStatusEnum_BiddingStrategySystemStatus_value = map[string]int32{
		"UNSPECIFIED":                        0,
		"UNKNOWN":                            1,
		"ENABLED":                            2,
		"LEARNING_NEW":                       3,
		"LEARNING_SETTING_CHANGE":            4,
		"LEARNING_BUDGET_CHANGE":             5,
		"LEARNING_COMPOSITION_CHANGE":        6,
		"LEARNING_CONVERSION_TYPE_CHANGE":    7,
		"LEARNING_CONVERSION_SETTING_CHANGE": 8,
		"LIMITED_BY_CPC_BID_CEILING":         9,
		"LIMITED_BY_CPC_BID_FLOOR":           10,
		"LIMITED_BY_DATA":                    11,
		"LIMITED_BY_BUDGET":                  12,
		"LIMITED_BY_LOW_PRIORITY_SPEND":      13,
		"LIMITED_BY_LOW_QUALITY":             14,
		"LIMITED_BY_INVENTORY":               15,
		"MISCONFIGURED_ZERO_ELIGIBILITY":     16,
		"MISCONFIGURED_CONVERSION_TYPES":     17,
		"MISCONFIGURED_CONVERSION_SETTINGS":  18,
		"MISCONFIGURED_SHARED_BUDGET":        19,
		"MISCONFIGURED_STRATEGY_TYPE":        20,
		"PAUSED":                             21,
		"UNAVAILABLE":                        22,
		"MULTIPLE_LEARNING":                  23,
		"MULTIPLE_LIMITED":                   24,
		"MULTIPLE_MISCONFIGURED":             25,
		"MULTIPLE":                           26,
	}
)

func (x BiddingStrategySystemStatusEnum_BiddingStrategySystemStatus) Enum() *BiddingStrategySystemStatusEnum_BiddingStrategySystemStatus {
	p := new(BiddingStrategySystemStatusEnum_BiddingStrategySystemStatus)
	*p = x
	return p
}

func (x BiddingStrategySystemStatusEnum_BiddingStrategySystemStatus) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (BiddingStrategySystemStatusEnum_BiddingStrategySystemStatus) Descriptor() protoreflect.EnumDescriptor {
	return file_enums_bidding_strategy_system_status_proto_enumTypes[0].Descriptor()
}

func (BiddingStrategySystemStatusEnum_BiddingStrategySystemStatus) Type() protoreflect.EnumType {
	return &file_enums_bidding_strategy_system_status_proto_enumTypes[0]
}

func (x BiddingStrategySystemStatusEnum_BiddingStrategySystemStatus) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use BiddingStrategySystemStatusEnum_BiddingStrategySystemStatus.Descriptor instead.
func (BiddingStrategySystemStatusEnum_BiddingStrategySystemStatus) EnumDescriptor() ([]byte, []int) {
	return file_enums_bidding_strategy_system_status_proto_rawDescGZIP(), []int{0, 0}
}

// Message describing BiddingStrategy system statuses.
type BiddingStrategySystemStatusEnum struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *BiddingStrategySystemStatusEnum) Reset() {
	*x = BiddingStrategySystemStatusEnum{}
	mi := &file_enums_bidding_strategy_system_status_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *BiddingStrategySystemStatusEnum) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BiddingStrategySystemStatusEnum) ProtoMessage() {}

func (x *BiddingStrategySystemStatusEnum) ProtoReflect() protoreflect.Message {
	mi := &file_enums_bidding_strategy_system_status_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use BiddingStrategySystemStatusEnum.ProtoReflect.Descriptor instead.
func (*BiddingStrategySystemStatusEnum) Descriptor() ([]byte, []int) {
	return file_enums_bidding_strategy_system_status_proto_rawDescGZIP(), []int{0}
}

var File_enums_bidding_strategy_system_status_proto protoreflect.FileDescriptor

var file_enums_bidding_strategy_system_status_proto_rawDesc = string([]byte{
	0x0a, 0x43, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x64, 0x73, 0x2f, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x61, 0x64, 0x73, 0x2f, 0x76, 0x31, 0x39, 0x2f, 0x65, 0x6e, 0x75, 0x6d, 0x73,
	0x2f, 0x62, 0x69, 0x64, 0x64, 0x69, 0x6e, 0x67, 0x5f, 0x73, 0x74, 0x72, 0x61, 0x74, 0x65, 0x67,
	0x79, 0x5f, 0x73, 0x79, 0x73, 0x74, 0x65, 0x6d, 0x5f, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x1e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x61, 0x64,
	0x73, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61, 0x64, 0x73, 0x2e, 0x76, 0x31, 0x39, 0x2e,
	0x65, 0x6e, 0x75, 0x6d, 0x73, 0x22, 0x8d, 0x06, 0x0a, 0x1f, 0x42, 0x69, 0x64, 0x64, 0x69, 0x6e,
	0x67, 0x53, 0x74, 0x72, 0x61, 0x74, 0x65, 0x67, 0x79, 0x53, 0x79, 0x73, 0x74, 0x65, 0x6d, 0x53,
	0x74, 0x61, 0x74, 0x75, 0x73, 0x45, 0x6e, 0x75, 0x6d, 0x22, 0xe9, 0x05, 0x0a, 0x1b, 0x42, 0x69,
	0x64, 0x64, 0x69, 0x6e, 0x67, 0x53, 0x74, 0x72, 0x61, 0x74, 0x65, 0x67, 0x79, 0x53, 0x79, 0x73,
	0x74, 0x65, 0x6d, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x0f, 0x0a, 0x0b, 0x55, 0x4e, 0x53,
	0x50, 0x45, 0x43, 0x49, 0x46, 0x49, 0x45, 0x44, 0x10, 0x00, 0x12, 0x0b, 0x0a, 0x07, 0x55, 0x4e,
	0x4b, 0x4e, 0x4f, 0x57, 0x4e, 0x10, 0x01, 0x12, 0x0b, 0x0a, 0x07, 0x45, 0x4e, 0x41, 0x42, 0x4c,
	0x45, 0x44, 0x10, 0x02, 0x12, 0x10, 0x0a, 0x0c, 0x4c, 0x45, 0x41, 0x52, 0x4e, 0x49, 0x4e, 0x47,
	0x5f, 0x4e, 0x45, 0x57, 0x10, 0x03, 0x12, 0x1b, 0x0a, 0x17, 0x4c, 0x45, 0x41, 0x52, 0x4e, 0x49,
	0x4e, 0x47, 0x5f, 0x53, 0x45, 0x54, 0x54, 0x49, 0x4e, 0x47, 0x5f, 0x43, 0x48, 0x41, 0x4e, 0x47,
	0x45, 0x10, 0x04, 0x12, 0x1a, 0x0a, 0x16, 0x4c, 0x45, 0x41, 0x52, 0x4e, 0x49, 0x4e, 0x47, 0x5f,
	0x42, 0x55, 0x44, 0x47, 0x45, 0x54, 0x5f, 0x43, 0x48, 0x41, 0x4e, 0x47, 0x45, 0x10, 0x05, 0x12,
	0x1f, 0x0a, 0x1b, 0x4c, 0x45, 0x41, 0x52, 0x4e, 0x49, 0x4e, 0x47, 0x5f, 0x43, 0x4f, 0x4d, 0x50,
	0x4f, 0x53, 0x49, 0x54, 0x49, 0x4f, 0x4e, 0x5f, 0x43, 0x48, 0x41, 0x4e, 0x47, 0x45, 0x10, 0x06,
	0x12, 0x23, 0x0a, 0x1f, 0x4c, 0x45, 0x41, 0x52, 0x4e, 0x49, 0x4e, 0x47, 0x5f, 0x43, 0x4f, 0x4e,
	0x56, 0x45, 0x52, 0x53, 0x49, 0x4f, 0x4e, 0x5f, 0x54, 0x59, 0x50, 0x45, 0x5f, 0x43, 0x48, 0x41,
	0x4e, 0x47, 0x45, 0x10, 0x07, 0x12, 0x26, 0x0a, 0x22, 0x4c, 0x45, 0x41, 0x52, 0x4e, 0x49, 0x4e,
	0x47, 0x5f, 0x43, 0x4f, 0x4e, 0x56, 0x45, 0x52, 0x53, 0x49, 0x4f, 0x4e, 0x5f, 0x53, 0x45, 0x54,
	0x54, 0x49, 0x4e, 0x47, 0x5f, 0x43, 0x48, 0x41, 0x4e, 0x47, 0x45, 0x10, 0x08, 0x12, 0x1e, 0x0a,
	0x1a, 0x4c, 0x49, 0x4d, 0x49, 0x54, 0x45, 0x44, 0x5f, 0x42, 0x59, 0x5f, 0x43, 0x50, 0x43, 0x5f,
	0x42, 0x49, 0x44, 0x5f, 0x43, 0x45, 0x49, 0x4c, 0x49, 0x4e, 0x47, 0x10, 0x09, 0x12, 0x1c, 0x0a,
	0x18, 0x4c, 0x49, 0x4d, 0x49, 0x54, 0x45, 0x44, 0x5f, 0x42, 0x59, 0x5f, 0x43, 0x50, 0x43, 0x5f,
	0x42, 0x49, 0x44, 0x5f, 0x46, 0x4c, 0x4f, 0x4f, 0x52, 0x10, 0x0a, 0x12, 0x13, 0x0a, 0x0f, 0x4c,
	0x49, 0x4d, 0x49, 0x54, 0x45, 0x44, 0x5f, 0x42, 0x59, 0x5f, 0x44, 0x41, 0x54, 0x41, 0x10, 0x0b,
	0x12, 0x15, 0x0a, 0x11, 0x4c, 0x49, 0x4d, 0x49, 0x54, 0x45, 0x44, 0x5f, 0x42, 0x59, 0x5f, 0x42,
	0x55, 0x44, 0x47, 0x45, 0x54, 0x10, 0x0c, 0x12, 0x21, 0x0a, 0x1d, 0x4c, 0x49, 0x4d, 0x49, 0x54,
	0x45, 0x44, 0x5f, 0x42, 0x59, 0x5f, 0x4c, 0x4f, 0x57, 0x5f, 0x50, 0x52, 0x49, 0x4f, 0x52, 0x49,
	0x54, 0x59, 0x5f, 0x53, 0x50, 0x45, 0x4e, 0x44, 0x10, 0x0d, 0x12, 0x1a, 0x0a, 0x16, 0x4c, 0x49,
	0x4d, 0x49, 0x54, 0x45, 0x44, 0x5f, 0x42, 0x59, 0x5f, 0x4c, 0x4f, 0x57, 0x5f, 0x51, 0x55, 0x41,
	0x4c, 0x49, 0x54, 0x59, 0x10, 0x0e, 0x12, 0x18, 0x0a, 0x14, 0x4c, 0x49, 0x4d, 0x49, 0x54, 0x45,
	0x44, 0x5f, 0x42, 0x59, 0x5f, 0x49, 0x4e, 0x56, 0x45, 0x4e, 0x54, 0x4f, 0x52, 0x59, 0x10, 0x0f,
	0x12, 0x22, 0x0a, 0x1e, 0x4d, 0x49, 0x53, 0x43, 0x4f, 0x4e, 0x46, 0x49, 0x47, 0x55, 0x52, 0x45,
	0x44, 0x5f, 0x5a, 0x45, 0x52, 0x4f, 0x5f, 0x45, 0x4c, 0x49, 0x47, 0x49, 0x42, 0x49, 0x4c, 0x49,
	0x54, 0x59, 0x10, 0x10, 0x12, 0x22, 0x0a, 0x1e, 0x4d, 0x49, 0x53, 0x43, 0x4f, 0x4e, 0x46, 0x49,
	0x47, 0x55, 0x52, 0x45, 0x44, 0x5f, 0x43, 0x4f, 0x4e, 0x56, 0x45, 0x52, 0x53, 0x49, 0x4f, 0x4e,
	0x5f, 0x54, 0x59, 0x50, 0x45, 0x53, 0x10, 0x11, 0x12, 0x25, 0x0a, 0x21, 0x4d, 0x49, 0x53, 0x43,
	0x4f, 0x4e, 0x46, 0x49, 0x47, 0x55, 0x52, 0x45, 0x44, 0x5f, 0x43, 0x4f, 0x4e, 0x56, 0x45, 0x52,
	0x53, 0x49, 0x4f, 0x4e, 0x5f, 0x53, 0x45, 0x54, 0x54, 0x49, 0x4e, 0x47, 0x53, 0x10, 0x12, 0x12,
	0x1f, 0x0a, 0x1b, 0x4d, 0x49, 0x53, 0x43, 0x4f, 0x4e, 0x46, 0x49, 0x47, 0x55, 0x52, 0x45, 0x44,
	0x5f, 0x53, 0x48, 0x41, 0x52, 0x45, 0x44, 0x5f, 0x42, 0x55, 0x44, 0x47, 0x45, 0x54, 0x10, 0x13,
	0x12, 0x1f, 0x0a, 0x1b, 0x4d, 0x49, 0x53, 0x43, 0x4f, 0x4e, 0x46, 0x49, 0x47, 0x55, 0x52, 0x45,
	0x44, 0x5f, 0x53, 0x54, 0x52, 0x41, 0x54, 0x45, 0x47, 0x59, 0x5f, 0x54, 0x59, 0x50, 0x45, 0x10,
	0x14, 0x12, 0x0a, 0x0a, 0x06, 0x50, 0x41, 0x55, 0x53, 0x45, 0x44, 0x10, 0x15, 0x12, 0x0f, 0x0a,
	0x0b, 0x55, 0x4e, 0x41, 0x56, 0x41, 0x49, 0x4c, 0x41, 0x42, 0x4c, 0x45, 0x10, 0x16, 0x12, 0x15,
	0x0a, 0x11, 0x4d, 0x55, 0x4c, 0x54, 0x49, 0x50, 0x4c, 0x45, 0x5f, 0x4c, 0x45, 0x41, 0x52, 0x4e,
	0x49, 0x4e, 0x47, 0x10, 0x17, 0x12, 0x14, 0x0a, 0x10, 0x4d, 0x55, 0x4c, 0x54, 0x49, 0x50, 0x4c,
	0x45, 0x5f, 0x4c, 0x49, 0x4d, 0x49, 0x54, 0x45, 0x44, 0x10, 0x18, 0x12, 0x1a, 0x0a, 0x16, 0x4d,
	0x55, 0x4c, 0x54, 0x49, 0x50, 0x4c, 0x45, 0x5f, 0x4d, 0x49, 0x53, 0x43, 0x4f, 0x4e, 0x46, 0x49,
	0x47, 0x55, 0x52, 0x45, 0x44, 0x10, 0x19, 0x12, 0x0c, 0x0a, 0x08, 0x4d, 0x55, 0x4c, 0x54, 0x49,
	0x50, 0x4c, 0x45, 0x10, 0x1a, 0x42, 0xfa, 0x01, 0x0a, 0x22, 0x63, 0x6f, 0x6d, 0x2e, 0x67, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x61, 0x64, 0x73, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61,
	0x64, 0x73, 0x2e, 0x76, 0x31, 0x39, 0x2e, 0x65, 0x6e, 0x75, 0x6d, 0x73, 0x42, 0x20, 0x42, 0x69,
	0x64, 0x64, 0x69, 0x6e, 0x67, 0x53, 0x74, 0x72, 0x61, 0x74, 0x65, 0x67, 0x79, 0x53, 0x79, 0x73,
	0x74, 0x65, 0x6d, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01,
	0x5a, 0x43, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x67, 0x6f, 0x6c, 0x61, 0x6e, 0x67, 0x2e,
	0x6f, 0x72, 0x67, 0x2f, 0x67, 0x65, 0x6e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x61, 0x70, 0x69, 0x73, 0x2f, 0x61, 0x64, 0x73, 0x2f, 0x67, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x61, 0x64, 0x73, 0x2f, 0x76, 0x31, 0x39, 0x2f, 0x65, 0x6e, 0x75, 0x6d, 0x73, 0x3b,
	0x65, 0x6e, 0x75, 0x6d, 0x73, 0xa2, 0x02, 0x03, 0x47, 0x41, 0x41, 0xaa, 0x02, 0x1e, 0x47, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x41, 0x64, 0x73, 0x2e, 0x47, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x41,
	0x64, 0x73, 0x2e, 0x56, 0x31, 0x39, 0x2e, 0x45, 0x6e, 0x75, 0x6d, 0x73, 0xca, 0x02, 0x1e, 0x47,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x5c, 0x41, 0x64, 0x73, 0x5c, 0x47, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x41, 0x64, 0x73, 0x5c, 0x56, 0x31, 0x39, 0x5c, 0x45, 0x6e, 0x75, 0x6d, 0x73, 0xea, 0x02, 0x22,
	0x47, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x3a, 0x3a, 0x41, 0x64, 0x73, 0x3a, 0x3a, 0x47, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x41, 0x64, 0x73, 0x3a, 0x3a, 0x56, 0x31, 0x39, 0x3a, 0x3a, 0x45, 0x6e, 0x75,
	0x6d, 0x73, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
})

var (
	file_enums_bidding_strategy_system_status_proto_rawDescOnce sync.Once
	file_enums_bidding_strategy_system_status_proto_rawDescData []byte
)

func file_enums_bidding_strategy_system_status_proto_rawDescGZIP() []byte {
	file_enums_bidding_strategy_system_status_proto_rawDescOnce.Do(func() {
		file_enums_bidding_strategy_system_status_proto_rawDescData = protoimpl.X.CompressGZIP(unsafe.Slice(unsafe.StringData(file_enums_bidding_strategy_system_status_proto_rawDesc), len(file_enums_bidding_strategy_system_status_proto_rawDesc)))
	})
	return file_enums_bidding_strategy_system_status_proto_rawDescData
}

var file_enums_bidding_strategy_system_status_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_enums_bidding_strategy_system_status_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_enums_bidding_strategy_system_status_proto_goTypes = []any{
	(BiddingStrategySystemStatusEnum_BiddingStrategySystemStatus)(0), // 0: google.ads.googleads.v19.enums.BiddingStrategySystemStatusEnum.BiddingStrategySystemStatus
	(*BiddingStrategySystemStatusEnum)(nil),                          // 1: google.ads.googleads.v19.enums.BiddingStrategySystemStatusEnum
}
var file_enums_bidding_strategy_system_status_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_enums_bidding_strategy_system_status_proto_init() }
func file_enums_bidding_strategy_system_status_proto_init() {
	if File_enums_bidding_strategy_system_status_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: unsafe.Slice(unsafe.StringData(file_enums_bidding_strategy_system_status_proto_rawDesc), len(file_enums_bidding_strategy_system_status_proto_rawDesc)),
			NumEnums:      1,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_enums_bidding_strategy_system_status_proto_goTypes,
		DependencyIndexes: file_enums_bidding_strategy_system_status_proto_depIdxs,
		EnumInfos:         file_enums_bidding_strategy_system_status_proto_enumTypes,
		MessageInfos:      file_enums_bidding_strategy_system_status_proto_msgTypes,
	}.Build()
	File_enums_bidding_strategy_system_status_proto = out.File
	file_enums_bidding_strategy_system_status_proto_goTypes = nil
	file_enums_bidding_strategy_system_status_proto_depIdxs = nil
}
