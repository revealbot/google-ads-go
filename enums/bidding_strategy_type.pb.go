// Copyright 2024 Google LLC
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
// source: google/ads/googleads/v17/enums/bidding_strategy_type.proto

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

// Enum describing possible bidding strategy types.
type BiddingStrategyTypeEnum_BiddingStrategyType int32

const (
	// Not specified.
	BiddingStrategyTypeEnum_UNSPECIFIED BiddingStrategyTypeEnum_BiddingStrategyType = 0
	// Used for return value only. Represents value unknown in this version.
	BiddingStrategyTypeEnum_UNKNOWN BiddingStrategyTypeEnum_BiddingStrategyType = 1
	// Commission is an automatic bidding strategy in which the advertiser pays
	// a certain portion of the conversion value.
	BiddingStrategyTypeEnum_COMMISSION BiddingStrategyTypeEnum_BiddingStrategyType = 16
	// Enhanced CPC is a bidding strategy that raises bids for clicks
	// that seem more likely to lead to a conversion and lowers
	// them for clicks where they seem less likely.
	BiddingStrategyTypeEnum_ENHANCED_CPC BiddingStrategyTypeEnum_BiddingStrategyType = 2
	// Fixed CPM is a manual bidding strategy with a fixed CPM.
	BiddingStrategyTypeEnum_FIXED_CPM BiddingStrategyTypeEnum_BiddingStrategyType = 19
	// Used for return value only. Indicates that a campaign does not have a
	// bidding strategy. This prevents the campaign from serving. For example,
	// a campaign may be attached to a manager bidding strategy and the serving
	// account is subsequently unlinked from the manager account. In this case
	// the campaign will automatically be detached from the now inaccessible
	// manager bidding strategy and transition to the INVALID bidding strategy
	// type.
	BiddingStrategyTypeEnum_INVALID BiddingStrategyTypeEnum_BiddingStrategyType = 17
	// Manual bidding strategy that allows advertiser to set the bid per
	// advertiser-specified action.
	BiddingStrategyTypeEnum_MANUAL_CPA BiddingStrategyTypeEnum_BiddingStrategyType = 18
	// Manual click based bidding where user pays per click.
	BiddingStrategyTypeEnum_MANUAL_CPC BiddingStrategyTypeEnum_BiddingStrategyType = 3
	// Manual impression based bidding
	// where user pays per thousand impressions.
	BiddingStrategyTypeEnum_MANUAL_CPM BiddingStrategyTypeEnum_BiddingStrategyType = 4
	// A bidding strategy that pays a configurable amount per video view.
	BiddingStrategyTypeEnum_MANUAL_CPV BiddingStrategyTypeEnum_BiddingStrategyType = 13
	// A bidding strategy that automatically maximizes number of conversions
	// given a daily budget.
	BiddingStrategyTypeEnum_MAXIMIZE_CONVERSIONS BiddingStrategyTypeEnum_BiddingStrategyType = 10
	// An automated bidding strategy that automatically sets bids to maximize
	// revenue while spending your budget.
	BiddingStrategyTypeEnum_MAXIMIZE_CONVERSION_VALUE BiddingStrategyTypeEnum_BiddingStrategyType = 11
	// Page-One Promoted bidding scheme, which sets max cpc bids to
	// target impressions on page one or page one promoted slots on google.com.
	// This enum value is deprecated.
	BiddingStrategyTypeEnum_PAGE_ONE_PROMOTED BiddingStrategyTypeEnum_BiddingStrategyType = 5
	// Percent Cpc is bidding strategy where bids are a fraction of the
	// advertised price for some good or service.
	BiddingStrategyTypeEnum_PERCENT_CPC BiddingStrategyTypeEnum_BiddingStrategyType = 12
	// Target CPA is an automated bid strategy that sets bids
	// to help get as many conversions as possible
	// at the target cost-per-acquisition (CPA) you set.
	BiddingStrategyTypeEnum_TARGET_CPA BiddingStrategyTypeEnum_BiddingStrategyType = 6
	// Target CPM is an automated bid strategy that sets bids to help get
	// as many impressions as possible at the target cost per one thousand
	// impressions (CPM) you set.
	BiddingStrategyTypeEnum_TARGET_CPM BiddingStrategyTypeEnum_BiddingStrategyType = 14
	// Target CPV is an automated bidding strategy that sets bids to optimize
	// performance given the average target cost per view.
	BiddingStrategyTypeEnum_TARGET_CPV BiddingStrategyTypeEnum_BiddingStrategyType = 20
	// An automated bidding strategy that sets bids so that a certain percentage
	// of search ads are shown at the top of the first page (or other targeted
	// location).
	BiddingStrategyTypeEnum_TARGET_IMPRESSION_SHARE BiddingStrategyTypeEnum_BiddingStrategyType = 15
	// Target Outrank Share is an automated bidding strategy that sets bids
	// based on the target fraction of auctions where the advertiser
	// should outrank a specific competitor.
	// This enum value is deprecated.
	BiddingStrategyTypeEnum_TARGET_OUTRANK_SHARE BiddingStrategyTypeEnum_BiddingStrategyType = 7
	// Target ROAS is an automated bidding strategy
	// that helps you maximize revenue while averaging
	// a specific target Return On Average Spend (ROAS).
	BiddingStrategyTypeEnum_TARGET_ROAS BiddingStrategyTypeEnum_BiddingStrategyType = 8
	// Target Spend is an automated bid strategy that sets your bids
	// to help get as many clicks as possible within your budget.
	BiddingStrategyTypeEnum_TARGET_SPEND BiddingStrategyTypeEnum_BiddingStrategyType = 9
)

// Enum value maps for BiddingStrategyTypeEnum_BiddingStrategyType.
var (
	BiddingStrategyTypeEnum_BiddingStrategyType_name = map[int32]string{
		0:  "UNSPECIFIED",
		1:  "UNKNOWN",
		16: "COMMISSION",
		2:  "ENHANCED_CPC",
		19: "FIXED_CPM",
		17: "INVALID",
		18: "MANUAL_CPA",
		3:  "MANUAL_CPC",
		4:  "MANUAL_CPM",
		13: "MANUAL_CPV",
		10: "MAXIMIZE_CONVERSIONS",
		11: "MAXIMIZE_CONVERSION_VALUE",
		5:  "PAGE_ONE_PROMOTED",
		12: "PERCENT_CPC",
		6:  "TARGET_CPA",
		14: "TARGET_CPM",
		20: "TARGET_CPV",
		15: "TARGET_IMPRESSION_SHARE",
		7:  "TARGET_OUTRANK_SHARE",
		8:  "TARGET_ROAS",
		9:  "TARGET_SPEND",
	}
	BiddingStrategyTypeEnum_BiddingStrategyType_value = map[string]int32{
		"UNSPECIFIED":               0,
		"UNKNOWN":                   1,
		"COMMISSION":                16,
		"ENHANCED_CPC":              2,
		"FIXED_CPM":                 19,
		"INVALID":                   17,
		"MANUAL_CPA":                18,
		"MANUAL_CPC":                3,
		"MANUAL_CPM":                4,
		"MANUAL_CPV":                13,
		"MAXIMIZE_CONVERSIONS":      10,
		"MAXIMIZE_CONVERSION_VALUE": 11,
		"PAGE_ONE_PROMOTED":         5,
		"PERCENT_CPC":               12,
		"TARGET_CPA":                6,
		"TARGET_CPM":                14,
		"TARGET_CPV":                20,
		"TARGET_IMPRESSION_SHARE":   15,
		"TARGET_OUTRANK_SHARE":      7,
		"TARGET_ROAS":               8,
		"TARGET_SPEND":              9,
	}
)

func (x BiddingStrategyTypeEnum_BiddingStrategyType) Enum() *BiddingStrategyTypeEnum_BiddingStrategyType {
	p := new(BiddingStrategyTypeEnum_BiddingStrategyType)
	*p = x
	return p
}

func (x BiddingStrategyTypeEnum_BiddingStrategyType) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (BiddingStrategyTypeEnum_BiddingStrategyType) Descriptor() protoreflect.EnumDescriptor {
	return file_enums_bidding_strategy_type_proto_enumTypes[0].Descriptor()
}

func (BiddingStrategyTypeEnum_BiddingStrategyType) Type() protoreflect.EnumType {
	return &file_enums_bidding_strategy_type_proto_enumTypes[0]
}

func (x BiddingStrategyTypeEnum_BiddingStrategyType) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use BiddingStrategyTypeEnum_BiddingStrategyType.Descriptor instead.
func (BiddingStrategyTypeEnum_BiddingStrategyType) EnumDescriptor() ([]byte, []int) {
	return file_enums_bidding_strategy_type_proto_rawDescGZIP(), []int{0, 0}
}

// Container for enum describing possible bidding strategy types.
type BiddingStrategyTypeEnum struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *BiddingStrategyTypeEnum) Reset() {
	*x = BiddingStrategyTypeEnum{}
	if protoimpl.UnsafeEnabled {
		mi := &file_enums_bidding_strategy_type_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *BiddingStrategyTypeEnum) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BiddingStrategyTypeEnum) ProtoMessage() {}

func (x *BiddingStrategyTypeEnum) ProtoReflect() protoreflect.Message {
	mi := &file_enums_bidding_strategy_type_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use BiddingStrategyTypeEnum.ProtoReflect.Descriptor instead.
func (*BiddingStrategyTypeEnum) Descriptor() ([]byte, []int) {
	return file_enums_bidding_strategy_type_proto_rawDescGZIP(), []int{0}
}

var File_enums_bidding_strategy_type_proto protoreflect.FileDescriptor

var file_enums_bidding_strategy_type_proto_rawDesc = []byte{
	0x0a, 0x3a, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x64, 0x73, 0x2f, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x61, 0x64, 0x73, 0x2f, 0x76, 0x31, 0x37, 0x2f, 0x65, 0x6e, 0x75, 0x6d, 0x73,
	0x2f, 0x62, 0x69, 0x64, 0x64, 0x69, 0x6e, 0x67, 0x5f, 0x73, 0x74, 0x72, 0x61, 0x74, 0x65, 0x67,
	0x79, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x1e, 0x67, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x61, 0x64, 0x73, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61,
	0x64, 0x73, 0x2e, 0x76, 0x31, 0x37, 0x2e, 0x65, 0x6e, 0x75, 0x6d, 0x73, 0x22, 0xb8, 0x03, 0x0a,
	0x17, 0x42, 0x69, 0x64, 0x64, 0x69, 0x6e, 0x67, 0x53, 0x74, 0x72, 0x61, 0x74, 0x65, 0x67, 0x79,
	0x54, 0x79, 0x70, 0x65, 0x45, 0x6e, 0x75, 0x6d, 0x22, 0x9c, 0x03, 0x0a, 0x13, 0x42, 0x69, 0x64,
	0x64, 0x69, 0x6e, 0x67, 0x53, 0x74, 0x72, 0x61, 0x74, 0x65, 0x67, 0x79, 0x54, 0x79, 0x70, 0x65,
	0x12, 0x0f, 0x0a, 0x0b, 0x55, 0x4e, 0x53, 0x50, 0x45, 0x43, 0x49, 0x46, 0x49, 0x45, 0x44, 0x10,
	0x00, 0x12, 0x0b, 0x0a, 0x07, 0x55, 0x4e, 0x4b, 0x4e, 0x4f, 0x57, 0x4e, 0x10, 0x01, 0x12, 0x0e,
	0x0a, 0x0a, 0x43, 0x4f, 0x4d, 0x4d, 0x49, 0x53, 0x53, 0x49, 0x4f, 0x4e, 0x10, 0x10, 0x12, 0x10,
	0x0a, 0x0c, 0x45, 0x4e, 0x48, 0x41, 0x4e, 0x43, 0x45, 0x44, 0x5f, 0x43, 0x50, 0x43, 0x10, 0x02,
	0x12, 0x0d, 0x0a, 0x09, 0x46, 0x49, 0x58, 0x45, 0x44, 0x5f, 0x43, 0x50, 0x4d, 0x10, 0x13, 0x12,
	0x0b, 0x0a, 0x07, 0x49, 0x4e, 0x56, 0x41, 0x4c, 0x49, 0x44, 0x10, 0x11, 0x12, 0x0e, 0x0a, 0x0a,
	0x4d, 0x41, 0x4e, 0x55, 0x41, 0x4c, 0x5f, 0x43, 0x50, 0x41, 0x10, 0x12, 0x12, 0x0e, 0x0a, 0x0a,
	0x4d, 0x41, 0x4e, 0x55, 0x41, 0x4c, 0x5f, 0x43, 0x50, 0x43, 0x10, 0x03, 0x12, 0x0e, 0x0a, 0x0a,
	0x4d, 0x41, 0x4e, 0x55, 0x41, 0x4c, 0x5f, 0x43, 0x50, 0x4d, 0x10, 0x04, 0x12, 0x0e, 0x0a, 0x0a,
	0x4d, 0x41, 0x4e, 0x55, 0x41, 0x4c, 0x5f, 0x43, 0x50, 0x56, 0x10, 0x0d, 0x12, 0x18, 0x0a, 0x14,
	0x4d, 0x41, 0x58, 0x49, 0x4d, 0x49, 0x5a, 0x45, 0x5f, 0x43, 0x4f, 0x4e, 0x56, 0x45, 0x52, 0x53,
	0x49, 0x4f, 0x4e, 0x53, 0x10, 0x0a, 0x12, 0x1d, 0x0a, 0x19, 0x4d, 0x41, 0x58, 0x49, 0x4d, 0x49,
	0x5a, 0x45, 0x5f, 0x43, 0x4f, 0x4e, 0x56, 0x45, 0x52, 0x53, 0x49, 0x4f, 0x4e, 0x5f, 0x56, 0x41,
	0x4c, 0x55, 0x45, 0x10, 0x0b, 0x12, 0x15, 0x0a, 0x11, 0x50, 0x41, 0x47, 0x45, 0x5f, 0x4f, 0x4e,
	0x45, 0x5f, 0x50, 0x52, 0x4f, 0x4d, 0x4f, 0x54, 0x45, 0x44, 0x10, 0x05, 0x12, 0x0f, 0x0a, 0x0b,
	0x50, 0x45, 0x52, 0x43, 0x45, 0x4e, 0x54, 0x5f, 0x43, 0x50, 0x43, 0x10, 0x0c, 0x12, 0x0e, 0x0a,
	0x0a, 0x54, 0x41, 0x52, 0x47, 0x45, 0x54, 0x5f, 0x43, 0x50, 0x41, 0x10, 0x06, 0x12, 0x0e, 0x0a,
	0x0a, 0x54, 0x41, 0x52, 0x47, 0x45, 0x54, 0x5f, 0x43, 0x50, 0x4d, 0x10, 0x0e, 0x12, 0x0e, 0x0a,
	0x0a, 0x54, 0x41, 0x52, 0x47, 0x45, 0x54, 0x5f, 0x43, 0x50, 0x56, 0x10, 0x14, 0x12, 0x1b, 0x0a,
	0x17, 0x54, 0x41, 0x52, 0x47, 0x45, 0x54, 0x5f, 0x49, 0x4d, 0x50, 0x52, 0x45, 0x53, 0x53, 0x49,
	0x4f, 0x4e, 0x5f, 0x53, 0x48, 0x41, 0x52, 0x45, 0x10, 0x0f, 0x12, 0x18, 0x0a, 0x14, 0x54, 0x41,
	0x52, 0x47, 0x45, 0x54, 0x5f, 0x4f, 0x55, 0x54, 0x52, 0x41, 0x4e, 0x4b, 0x5f, 0x53, 0x48, 0x41,
	0x52, 0x45, 0x10, 0x07, 0x12, 0x0f, 0x0a, 0x0b, 0x54, 0x41, 0x52, 0x47, 0x45, 0x54, 0x5f, 0x52,
	0x4f, 0x41, 0x53, 0x10, 0x08, 0x12, 0x10, 0x0a, 0x0c, 0x54, 0x41, 0x52, 0x47, 0x45, 0x54, 0x5f,
	0x53, 0x50, 0x45, 0x4e, 0x44, 0x10, 0x09, 0x42, 0xf2, 0x01, 0x0a, 0x22, 0x63, 0x6f, 0x6d, 0x2e,
	0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x61, 0x64, 0x73, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x61, 0x64, 0x73, 0x2e, 0x76, 0x31, 0x37, 0x2e, 0x65, 0x6e, 0x75, 0x6d, 0x73, 0x42, 0x18,
	0x42, 0x69, 0x64, 0x64, 0x69, 0x6e, 0x67, 0x53, 0x74, 0x72, 0x61, 0x74, 0x65, 0x67, 0x79, 0x54,
	0x79, 0x70, 0x65, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a, 0x43, 0x67, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x2e, 0x67, 0x6f, 0x6c, 0x61, 0x6e, 0x67, 0x2e, 0x6f, 0x72, 0x67, 0x2f, 0x67, 0x65,
	0x6e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61, 0x70, 0x69,
	0x73, 0x2f, 0x61, 0x64, 0x73, 0x2f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61, 0x64, 0x73, 0x2f,
	0x76, 0x31, 0x37, 0x2f, 0x65, 0x6e, 0x75, 0x6d, 0x73, 0x3b, 0x65, 0x6e, 0x75, 0x6d, 0x73, 0xa2,
	0x02, 0x03, 0x47, 0x41, 0x41, 0xaa, 0x02, 0x1e, 0x47, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x41,
	0x64, 0x73, 0x2e, 0x47, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x41, 0x64, 0x73, 0x2e, 0x56, 0x31, 0x37,
	0x2e, 0x45, 0x6e, 0x75, 0x6d, 0x73, 0xca, 0x02, 0x1e, 0x47, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x5c,
	0x41, 0x64, 0x73, 0x5c, 0x47, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x41, 0x64, 0x73, 0x5c, 0x56, 0x31,
	0x37, 0x5c, 0x45, 0x6e, 0x75, 0x6d, 0x73, 0xea, 0x02, 0x22, 0x47, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x3a, 0x3a, 0x41, 0x64, 0x73, 0x3a, 0x3a, 0x47, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x41, 0x64, 0x73,
	0x3a, 0x3a, 0x56, 0x31, 0x37, 0x3a, 0x3a, 0x45, 0x6e, 0x75, 0x6d, 0x73, 0x62, 0x06, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_enums_bidding_strategy_type_proto_rawDescOnce sync.Once
	file_enums_bidding_strategy_type_proto_rawDescData = file_enums_bidding_strategy_type_proto_rawDesc
)

func file_enums_bidding_strategy_type_proto_rawDescGZIP() []byte {
	file_enums_bidding_strategy_type_proto_rawDescOnce.Do(func() {
		file_enums_bidding_strategy_type_proto_rawDescData = protoimpl.X.CompressGZIP(file_enums_bidding_strategy_type_proto_rawDescData)
	})
	return file_enums_bidding_strategy_type_proto_rawDescData
}

var file_enums_bidding_strategy_type_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_enums_bidding_strategy_type_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_enums_bidding_strategy_type_proto_goTypes = []interface{}{
	(BiddingStrategyTypeEnum_BiddingStrategyType)(0), // 0: google.ads.googleads.v17.enums.BiddingStrategyTypeEnum.BiddingStrategyType
	(*BiddingStrategyTypeEnum)(nil),                  // 1: google.ads.googleads.v17.enums.BiddingStrategyTypeEnum
}
var file_enums_bidding_strategy_type_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_enums_bidding_strategy_type_proto_init() }
func file_enums_bidding_strategy_type_proto_init() {
	if File_enums_bidding_strategy_type_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_enums_bidding_strategy_type_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*BiddingStrategyTypeEnum); i {
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
			RawDescriptor: file_enums_bidding_strategy_type_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_enums_bidding_strategy_type_proto_goTypes,
		DependencyIndexes: file_enums_bidding_strategy_type_proto_depIdxs,
		EnumInfos:         file_enums_bidding_strategy_type_proto_enumTypes,
		MessageInfos:      file_enums_bidding_strategy_type_proto_msgTypes,
	}.Build()
	File_enums_bidding_strategy_type_proto = out.File
	file_enums_bidding_strategy_type_proto_rawDesc = nil
	file_enums_bidding_strategy_type_proto_goTypes = nil
	file_enums_bidding_strategy_type_proto_depIdxs = nil
}
