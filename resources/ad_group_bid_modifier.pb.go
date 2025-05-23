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
// source: google/ads/googleads/v19/resources/ad_group_bid_modifier.proto

package resources

import (
	common "github.com/revealbot/google-ads-go/common"
	enums "github.com/revealbot/google-ads-go/enums"
	_ "google.golang.org/genproto/googleapis/api/annotations"
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

// Represents an ad group bid modifier.
type AdGroupBidModifier struct {
	state protoimpl.MessageState `protogen:"open.v1"`
	// Immutable. The resource name of the ad group bid modifier.
	// Ad group bid modifier resource names have the form:
	//
	// `customers/{customer_id}/adGroupBidModifiers/{ad_group_id}~{criterion_id}`
	ResourceName string `protobuf:"bytes,1,opt,name=resource_name,json=resourceName,proto3" json:"resource_name,omitempty"`
	// Immutable. The ad group to which this criterion belongs.
	AdGroup *string `protobuf:"bytes,13,opt,name=ad_group,json=adGroup,proto3,oneof" json:"ad_group,omitempty"`
	// Output only. The ID of the criterion to bid modify.
	//
	// This field is ignored for mutates.
	CriterionId *int64 `protobuf:"varint,14,opt,name=criterion_id,json=criterionId,proto3,oneof" json:"criterion_id,omitempty"`
	// The modifier for the bid when the criterion matches. The modifier must be
	// in the range: 0.1 - 10.0. The range is 1.0 - 6.0 for PreferredContent.
	// Use 0 to opt out of a Device type.
	BidModifier *float64 `protobuf:"fixed64,15,opt,name=bid_modifier,json=bidModifier,proto3,oneof" json:"bid_modifier,omitempty"`
	// Output only. The base ad group from which this draft/trial adgroup bid
	// modifier was created. If ad_group is a base ad group then this field will
	// be equal to ad_group. If the ad group was created in the draft or trial and
	// has no corresponding base ad group, then this field will be null. This
	// field is readonly.
	BaseAdGroup *string `protobuf:"bytes,16,opt,name=base_ad_group,json=baseAdGroup,proto3,oneof" json:"base_ad_group,omitempty"`
	// Output only. Bid modifier source.
	BidModifierSource enums.BidModifierSourceEnum_BidModifierSource `protobuf:"varint,10,opt,name=bid_modifier_source,json=bidModifierSource,proto3,enum=google.ads.googleads.v19.enums.BidModifierSourceEnum_BidModifierSource" json:"bid_modifier_source,omitempty"`
	// The criterion of this ad group bid modifier.
	//
	// Required in create operations starting in V5.
	//
	// Types that are valid to be assigned to Criterion:
	//
	//	*AdGroupBidModifier_HotelDateSelectionType
	//	*AdGroupBidModifier_HotelAdvanceBookingWindow
	//	*AdGroupBidModifier_HotelLengthOfStay
	//	*AdGroupBidModifier_HotelCheckInDay
	//	*AdGroupBidModifier_Device
	//	*AdGroupBidModifier_HotelCheckInDateRange
	Criterion     isAdGroupBidModifier_Criterion `protobuf_oneof:"criterion"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *AdGroupBidModifier) Reset() {
	*x = AdGroupBidModifier{}
	mi := &file_resources_ad_group_bid_modifier_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *AdGroupBidModifier) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AdGroupBidModifier) ProtoMessage() {}

func (x *AdGroupBidModifier) ProtoReflect() protoreflect.Message {
	mi := &file_resources_ad_group_bid_modifier_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AdGroupBidModifier.ProtoReflect.Descriptor instead.
func (*AdGroupBidModifier) Descriptor() ([]byte, []int) {
	return file_resources_ad_group_bid_modifier_proto_rawDescGZIP(), []int{0}
}

func (x *AdGroupBidModifier) GetResourceName() string {
	if x != nil {
		return x.ResourceName
	}
	return ""
}

func (x *AdGroupBidModifier) GetAdGroup() string {
	if x != nil && x.AdGroup != nil {
		return *x.AdGroup
	}
	return ""
}

func (x *AdGroupBidModifier) GetCriterionId() int64 {
	if x != nil && x.CriterionId != nil {
		return *x.CriterionId
	}
	return 0
}

func (x *AdGroupBidModifier) GetBidModifier() float64 {
	if x != nil && x.BidModifier != nil {
		return *x.BidModifier
	}
	return 0
}

func (x *AdGroupBidModifier) GetBaseAdGroup() string {
	if x != nil && x.BaseAdGroup != nil {
		return *x.BaseAdGroup
	}
	return ""
}

func (x *AdGroupBidModifier) GetBidModifierSource() enums.BidModifierSourceEnum_BidModifierSource {
	if x != nil {
		return x.BidModifierSource
	}
	return enums.BidModifierSourceEnum_BidModifierSource(0)
}

func (x *AdGroupBidModifier) GetCriterion() isAdGroupBidModifier_Criterion {
	if x != nil {
		return x.Criterion
	}
	return nil
}

func (x *AdGroupBidModifier) GetHotelDateSelectionType() *common.HotelDateSelectionTypeInfo {
	if x != nil {
		if x, ok := x.Criterion.(*AdGroupBidModifier_HotelDateSelectionType); ok {
			return x.HotelDateSelectionType
		}
	}
	return nil
}

func (x *AdGroupBidModifier) GetHotelAdvanceBookingWindow() *common.HotelAdvanceBookingWindowInfo {
	if x != nil {
		if x, ok := x.Criterion.(*AdGroupBidModifier_HotelAdvanceBookingWindow); ok {
			return x.HotelAdvanceBookingWindow
		}
	}
	return nil
}

func (x *AdGroupBidModifier) GetHotelLengthOfStay() *common.HotelLengthOfStayInfo {
	if x != nil {
		if x, ok := x.Criterion.(*AdGroupBidModifier_HotelLengthOfStay); ok {
			return x.HotelLengthOfStay
		}
	}
	return nil
}

func (x *AdGroupBidModifier) GetHotelCheckInDay() *common.HotelCheckInDayInfo {
	if x != nil {
		if x, ok := x.Criterion.(*AdGroupBidModifier_HotelCheckInDay); ok {
			return x.HotelCheckInDay
		}
	}
	return nil
}

func (x *AdGroupBidModifier) GetDevice() *common.DeviceInfo {
	if x != nil {
		if x, ok := x.Criterion.(*AdGroupBidModifier_Device); ok {
			return x.Device
		}
	}
	return nil
}

func (x *AdGroupBidModifier) GetHotelCheckInDateRange() *common.HotelCheckInDateRangeInfo {
	if x != nil {
		if x, ok := x.Criterion.(*AdGroupBidModifier_HotelCheckInDateRange); ok {
			return x.HotelCheckInDateRange
		}
	}
	return nil
}

type isAdGroupBidModifier_Criterion interface {
	isAdGroupBidModifier_Criterion()
}

type AdGroupBidModifier_HotelDateSelectionType struct {
	// Immutable. Criterion for hotel date selection (default dates versus user
	// selected).
	HotelDateSelectionType *common.HotelDateSelectionTypeInfo `protobuf:"bytes,5,opt,name=hotel_date_selection_type,json=hotelDateSelectionType,proto3,oneof"`
}

type AdGroupBidModifier_HotelAdvanceBookingWindow struct {
	// Immutable. Criterion for number of days prior to the stay the booking is
	// being made.
	HotelAdvanceBookingWindow *common.HotelAdvanceBookingWindowInfo `protobuf:"bytes,6,opt,name=hotel_advance_booking_window,json=hotelAdvanceBookingWindow,proto3,oneof"`
}

type AdGroupBidModifier_HotelLengthOfStay struct {
	// Immutable. Criterion for length of hotel stay in nights.
	HotelLengthOfStay *common.HotelLengthOfStayInfo `protobuf:"bytes,7,opt,name=hotel_length_of_stay,json=hotelLengthOfStay,proto3,oneof"`
}

type AdGroupBidModifier_HotelCheckInDay struct {
	// Immutable. Criterion for day of the week the booking is for.
	HotelCheckInDay *common.HotelCheckInDayInfo `protobuf:"bytes,8,opt,name=hotel_check_in_day,json=hotelCheckInDay,proto3,oneof"`
}

type AdGroupBidModifier_Device struct {
	// Immutable. A device criterion.
	Device *common.DeviceInfo `protobuf:"bytes,11,opt,name=device,proto3,oneof"`
}

type AdGroupBidModifier_HotelCheckInDateRange struct {
	// Immutable. Criterion for a hotel check-in date range.
	HotelCheckInDateRange *common.HotelCheckInDateRangeInfo `protobuf:"bytes,17,opt,name=hotel_check_in_date_range,json=hotelCheckInDateRange,proto3,oneof"`
}

func (*AdGroupBidModifier_HotelDateSelectionType) isAdGroupBidModifier_Criterion() {}

func (*AdGroupBidModifier_HotelAdvanceBookingWindow) isAdGroupBidModifier_Criterion() {}

func (*AdGroupBidModifier_HotelLengthOfStay) isAdGroupBidModifier_Criterion() {}

func (*AdGroupBidModifier_HotelCheckInDay) isAdGroupBidModifier_Criterion() {}

func (*AdGroupBidModifier_Device) isAdGroupBidModifier_Criterion() {}

func (*AdGroupBidModifier_HotelCheckInDateRange) isAdGroupBidModifier_Criterion() {}

var File_resources_ad_group_bid_modifier_proto protoreflect.FileDescriptor

var file_resources_ad_group_bid_modifier_proto_rawDesc = string([]byte{
	0x0a, 0x3e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x64, 0x73, 0x2f, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x61, 0x64, 0x73, 0x2f, 0x76, 0x31, 0x39, 0x2f, 0x72, 0x65, 0x73, 0x6f, 0x75,
	0x72, 0x63, 0x65, 0x73, 0x2f, 0x61, 0x64, 0x5f, 0x67, 0x72, 0x6f, 0x75, 0x70, 0x5f, 0x62, 0x69,
	0x64, 0x5f, 0x6d, 0x6f, 0x64, 0x69, 0x66, 0x69, 0x65, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x12, 0x22, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x61, 0x64, 0x73, 0x2e, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x61, 0x64, 0x73, 0x2e, 0x76, 0x31, 0x39, 0x2e, 0x72, 0x65, 0x73, 0x6f, 0x75,
	0x72, 0x63, 0x65, 0x73, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x64, 0x73,
	0x2f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61, 0x64, 0x73, 0x2f, 0x76, 0x31, 0x39, 0x2f, 0x63,
	0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2f, 0x63, 0x72, 0x69, 0x74, 0x65, 0x72, 0x69, 0x61, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x38, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x64, 0x73,
	0x2f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61, 0x64, 0x73, 0x2f, 0x76, 0x31, 0x39, 0x2f, 0x65,
	0x6e, 0x75, 0x6d, 0x73, 0x2f, 0x62, 0x69, 0x64, 0x5f, 0x6d, 0x6f, 0x64, 0x69, 0x66, 0x69, 0x65,
	0x72, 0x5f, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1f,
	0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x66, 0x69, 0x65, 0x6c, 0x64,
	0x5f, 0x62, 0x65, 0x68, 0x61, 0x76, 0x69, 0x6f, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a,
	0x19, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x72, 0x65, 0x73, 0x6f,
	0x75, 0x72, 0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xd3, 0x0a, 0x0a, 0x12, 0x41,
	0x64, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x42, 0x69, 0x64, 0x4d, 0x6f, 0x64, 0x69, 0x66, 0x69, 0x65,
	0x72, 0x12, 0x58, 0x0a, 0x0d, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x5f, 0x6e, 0x61,
	0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x42, 0x33, 0xe0, 0x41, 0x05, 0xfa, 0x41, 0x2d,
	0x0a, 0x2b, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61, 0x64, 0x73, 0x2e, 0x67, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x61, 0x70, 0x69, 0x73, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x41, 0x64, 0x47, 0x72, 0x6f,
	0x75, 0x70, 0x42, 0x69, 0x64, 0x4d, 0x6f, 0x64, 0x69, 0x66, 0x69, 0x65, 0x72, 0x52, 0x0c, 0x72,
	0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x48, 0x0a, 0x08, 0x61,
	0x64, 0x5f, 0x67, 0x72, 0x6f, 0x75, 0x70, 0x18, 0x0d, 0x20, 0x01, 0x28, 0x09, 0x42, 0x28, 0xe0,
	0x41, 0x05, 0xfa, 0x41, 0x22, 0x0a, 0x20, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61, 0x64, 0x73,
	0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61, 0x70, 0x69, 0x73, 0x2e, 0x63, 0x6f, 0x6d, 0x2f,
	0x41, 0x64, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x48, 0x01, 0x52, 0x07, 0x61, 0x64, 0x47, 0x72, 0x6f,
	0x75, 0x70, 0x88, 0x01, 0x01, 0x12, 0x2b, 0x0a, 0x0c, 0x63, 0x72, 0x69, 0x74, 0x65, 0x72, 0x69,
	0x6f, 0x6e, 0x5f, 0x69, 0x64, 0x18, 0x0e, 0x20, 0x01, 0x28, 0x03, 0x42, 0x03, 0xe0, 0x41, 0x03,
	0x48, 0x02, 0x52, 0x0b, 0x63, 0x72, 0x69, 0x74, 0x65, 0x72, 0x69, 0x6f, 0x6e, 0x49, 0x64, 0x88,
	0x01, 0x01, 0x12, 0x26, 0x0a, 0x0c, 0x62, 0x69, 0x64, 0x5f, 0x6d, 0x6f, 0x64, 0x69, 0x66, 0x69,
	0x65, 0x72, 0x18, 0x0f, 0x20, 0x01, 0x28, 0x01, 0x48, 0x03, 0x52, 0x0b, 0x62, 0x69, 0x64, 0x4d,
	0x6f, 0x64, 0x69, 0x66, 0x69, 0x65, 0x72, 0x88, 0x01, 0x01, 0x12, 0x51, 0x0a, 0x0d, 0x62, 0x61,
	0x73, 0x65, 0x5f, 0x61, 0x64, 0x5f, 0x67, 0x72, 0x6f, 0x75, 0x70, 0x18, 0x10, 0x20, 0x01, 0x28,
	0x09, 0x42, 0x28, 0xe0, 0x41, 0x03, 0xfa, 0x41, 0x22, 0x0a, 0x20, 0x67, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x61, 0x64, 0x73, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61, 0x70, 0x69, 0x73, 0x2e,
	0x63, 0x6f, 0x6d, 0x2f, 0x41, 0x64, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x48, 0x04, 0x52, 0x0b, 0x62,
	0x61, 0x73, 0x65, 0x41, 0x64, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x88, 0x01, 0x01, 0x12, 0x7c, 0x0a,
	0x13, 0x62, 0x69, 0x64, 0x5f, 0x6d, 0x6f, 0x64, 0x69, 0x66, 0x69, 0x65, 0x72, 0x5f, 0x73, 0x6f,
	0x75, 0x72, 0x63, 0x65, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x47, 0x2e, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x2e, 0x61, 0x64, 0x73, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61, 0x64,
	0x73, 0x2e, 0x76, 0x31, 0x39, 0x2e, 0x65, 0x6e, 0x75, 0x6d, 0x73, 0x2e, 0x42, 0x69, 0x64, 0x4d,
	0x6f, 0x64, 0x69, 0x66, 0x69, 0x65, 0x72, 0x53, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x45, 0x6e, 0x75,
	0x6d, 0x2e, 0x42, 0x69, 0x64, 0x4d, 0x6f, 0x64, 0x69, 0x66, 0x69, 0x65, 0x72, 0x53, 0x6f, 0x75,
	0x72, 0x63, 0x65, 0x42, 0x03, 0xe0, 0x41, 0x03, 0x52, 0x11, 0x62, 0x69, 0x64, 0x4d, 0x6f, 0x64,
	0x69, 0x66, 0x69, 0x65, 0x72, 0x53, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x12, 0x7d, 0x0a, 0x19, 0x68,
	0x6f, 0x74, 0x65, 0x6c, 0x5f, 0x64, 0x61, 0x74, 0x65, 0x5f, 0x73, 0x65, 0x6c, 0x65, 0x63, 0x74,
	0x69, 0x6f, 0x6e, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x3b,
	0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x61, 0x64, 0x73, 0x2e, 0x67, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x61, 0x64, 0x73, 0x2e, 0x76, 0x31, 0x39, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e,
	0x2e, 0x48, 0x6f, 0x74, 0x65, 0x6c, 0x44, 0x61, 0x74, 0x65, 0x53, 0x65, 0x6c, 0x65, 0x63, 0x74,
	0x69, 0x6f, 0x6e, 0x54, 0x79, 0x70, 0x65, 0x49, 0x6e, 0x66, 0x6f, 0x42, 0x03, 0xe0, 0x41, 0x05,
	0x48, 0x00, 0x52, 0x16, 0x68, 0x6f, 0x74, 0x65, 0x6c, 0x44, 0x61, 0x74, 0x65, 0x53, 0x65, 0x6c,
	0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x54, 0x79, 0x70, 0x65, 0x12, 0x86, 0x01, 0x0a, 0x1c, 0x68,
	0x6f, 0x74, 0x65, 0x6c, 0x5f, 0x61, 0x64, 0x76, 0x61, 0x6e, 0x63, 0x65, 0x5f, 0x62, 0x6f, 0x6f,
	0x6b, 0x69, 0x6e, 0x67, 0x5f, 0x77, 0x69, 0x6e, 0x64, 0x6f, 0x77, 0x18, 0x06, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x3e, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x61, 0x64, 0x73, 0x2e, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61, 0x64, 0x73, 0x2e, 0x76, 0x31, 0x39, 0x2e, 0x63, 0x6f, 0x6d,
	0x6d, 0x6f, 0x6e, 0x2e, 0x48, 0x6f, 0x74, 0x65, 0x6c, 0x41, 0x64, 0x76, 0x61, 0x6e, 0x63, 0x65,
	0x42, 0x6f, 0x6f, 0x6b, 0x69, 0x6e, 0x67, 0x57, 0x69, 0x6e, 0x64, 0x6f, 0x77, 0x49, 0x6e, 0x66,
	0x6f, 0x42, 0x03, 0xe0, 0x41, 0x05, 0x48, 0x00, 0x52, 0x19, 0x68, 0x6f, 0x74, 0x65, 0x6c, 0x41,
	0x64, 0x76, 0x61, 0x6e, 0x63, 0x65, 0x42, 0x6f, 0x6f, 0x6b, 0x69, 0x6e, 0x67, 0x57, 0x69, 0x6e,
	0x64, 0x6f, 0x77, 0x12, 0x6e, 0x0a, 0x14, 0x68, 0x6f, 0x74, 0x65, 0x6c, 0x5f, 0x6c, 0x65, 0x6e,
	0x67, 0x74, 0x68, 0x5f, 0x6f, 0x66, 0x5f, 0x73, 0x74, 0x61, 0x79, 0x18, 0x07, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x36, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x61, 0x64, 0x73, 0x2e, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61, 0x64, 0x73, 0x2e, 0x76, 0x31, 0x39, 0x2e, 0x63, 0x6f, 0x6d,
	0x6d, 0x6f, 0x6e, 0x2e, 0x48, 0x6f, 0x74, 0x65, 0x6c, 0x4c, 0x65, 0x6e, 0x67, 0x74, 0x68, 0x4f,
	0x66, 0x53, 0x74, 0x61, 0x79, 0x49, 0x6e, 0x66, 0x6f, 0x42, 0x03, 0xe0, 0x41, 0x05, 0x48, 0x00,
	0x52, 0x11, 0x68, 0x6f, 0x74, 0x65, 0x6c, 0x4c, 0x65, 0x6e, 0x67, 0x74, 0x68, 0x4f, 0x66, 0x53,
	0x74, 0x61, 0x79, 0x12, 0x68, 0x0a, 0x12, 0x68, 0x6f, 0x74, 0x65, 0x6c, 0x5f, 0x63, 0x68, 0x65,
	0x63, 0x6b, 0x5f, 0x69, 0x6e, 0x5f, 0x64, 0x61, 0x79, 0x18, 0x08, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x34, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x61, 0x64, 0x73, 0x2e, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x61, 0x64, 0x73, 0x2e, 0x76, 0x31, 0x39, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f,
	0x6e, 0x2e, 0x48, 0x6f, 0x74, 0x65, 0x6c, 0x43, 0x68, 0x65, 0x63, 0x6b, 0x49, 0x6e, 0x44, 0x61,
	0x79, 0x49, 0x6e, 0x66, 0x6f, 0x42, 0x03, 0xe0, 0x41, 0x05, 0x48, 0x00, 0x52, 0x0f, 0x68, 0x6f,
	0x74, 0x65, 0x6c, 0x43, 0x68, 0x65, 0x63, 0x6b, 0x49, 0x6e, 0x44, 0x61, 0x79, 0x12, 0x4a, 0x0a,
	0x06, 0x64, 0x65, 0x76, 0x69, 0x63, 0x65, 0x18, 0x0b, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x2b, 0x2e,
	0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x61, 0x64, 0x73, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x61, 0x64, 0x73, 0x2e, 0x76, 0x31, 0x39, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e,
	0x44, 0x65, 0x76, 0x69, 0x63, 0x65, 0x49, 0x6e, 0x66, 0x6f, 0x42, 0x03, 0xe0, 0x41, 0x05, 0x48,
	0x00, 0x52, 0x06, 0x64, 0x65, 0x76, 0x69, 0x63, 0x65, 0x12, 0x7b, 0x0a, 0x19, 0x68, 0x6f, 0x74,
	0x65, 0x6c, 0x5f, 0x63, 0x68, 0x65, 0x63, 0x6b, 0x5f, 0x69, 0x6e, 0x5f, 0x64, 0x61, 0x74, 0x65,
	0x5f, 0x72, 0x61, 0x6e, 0x67, 0x65, 0x18, 0x11, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x3a, 0x2e, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x61, 0x64, 0x73, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x61, 0x64, 0x73, 0x2e, 0x76, 0x31, 0x39, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x48,
	0x6f, 0x74, 0x65, 0x6c, 0x43, 0x68, 0x65, 0x63, 0x6b, 0x49, 0x6e, 0x44, 0x61, 0x74, 0x65, 0x52,
	0x61, 0x6e, 0x67, 0x65, 0x49, 0x6e, 0x66, 0x6f, 0x42, 0x03, 0xe0, 0x41, 0x05, 0x48, 0x00, 0x52,
	0x15, 0x68, 0x6f, 0x74, 0x65, 0x6c, 0x43, 0x68, 0x65, 0x63, 0x6b, 0x49, 0x6e, 0x44, 0x61, 0x74,
	0x65, 0x52, 0x61, 0x6e, 0x67, 0x65, 0x3a, 0x7a, 0xea, 0x41, 0x77, 0x0a, 0x2b, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x61, 0x64, 0x73, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61, 0x70, 0x69,
	0x73, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x41, 0x64, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x42, 0x69, 0x64,
	0x4d, 0x6f, 0x64, 0x69, 0x66, 0x69, 0x65, 0x72, 0x12, 0x48, 0x63, 0x75, 0x73, 0x74, 0x6f, 0x6d,
	0x65, 0x72, 0x73, 0x2f, 0x7b, 0x63, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x65, 0x72, 0x5f, 0x69, 0x64,
	0x7d, 0x2f, 0x61, 0x64, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x42, 0x69, 0x64, 0x4d, 0x6f, 0x64, 0x69,
	0x66, 0x69, 0x65, 0x72, 0x73, 0x2f, 0x7b, 0x61, 0x64, 0x5f, 0x67, 0x72, 0x6f, 0x75, 0x70, 0x5f,
	0x69, 0x64, 0x7d, 0x7e, 0x7b, 0x63, 0x72, 0x69, 0x74, 0x65, 0x72, 0x69, 0x6f, 0x6e, 0x5f, 0x69,
	0x64, 0x7d, 0x42, 0x0b, 0x0a, 0x09, 0x63, 0x72, 0x69, 0x74, 0x65, 0x72, 0x69, 0x6f, 0x6e, 0x42,
	0x0b, 0x0a, 0x09, 0x5f, 0x61, 0x64, 0x5f, 0x67, 0x72, 0x6f, 0x75, 0x70, 0x42, 0x0f, 0x0a, 0x0d,
	0x5f, 0x63, 0x72, 0x69, 0x74, 0x65, 0x72, 0x69, 0x6f, 0x6e, 0x5f, 0x69, 0x64, 0x42, 0x0f, 0x0a,
	0x0d, 0x5f, 0x62, 0x69, 0x64, 0x5f, 0x6d, 0x6f, 0x64, 0x69, 0x66, 0x69, 0x65, 0x72, 0x42, 0x10,
	0x0a, 0x0e, 0x5f, 0x62, 0x61, 0x73, 0x65, 0x5f, 0x61, 0x64, 0x5f, 0x67, 0x72, 0x6f, 0x75, 0x70,
	0x42, 0x89, 0x02, 0x0a, 0x26, 0x63, 0x6f, 0x6d, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e,
	0x61, 0x64, 0x73, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61, 0x64, 0x73, 0x2e, 0x76, 0x31,
	0x39, 0x2e, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x73, 0x42, 0x17, 0x41, 0x64, 0x47,
	0x72, 0x6f, 0x75, 0x70, 0x42, 0x69, 0x64, 0x4d, 0x6f, 0x64, 0x69, 0x66, 0x69, 0x65, 0x72, 0x50,
	0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a, 0x4b, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x67,
	0x6f, 0x6c, 0x61, 0x6e, 0x67, 0x2e, 0x6f, 0x72, 0x67, 0x2f, 0x67, 0x65, 0x6e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x2f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61, 0x70, 0x69, 0x73, 0x2f, 0x61, 0x64,
	0x73, 0x2f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61, 0x64, 0x73, 0x2f, 0x76, 0x31, 0x39, 0x2f,
	0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x73, 0x3b, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72,
	0x63, 0x65, 0x73, 0xa2, 0x02, 0x03, 0x47, 0x41, 0x41, 0xaa, 0x02, 0x22, 0x47, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x2e, 0x41, 0x64, 0x73, 0x2e, 0x47, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x41, 0x64, 0x73,
	0x2e, 0x56, 0x31, 0x39, 0x2e, 0x52, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x73, 0xca, 0x02,
	0x22, 0x47, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x5c, 0x41, 0x64, 0x73, 0x5c, 0x47, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x41, 0x64, 0x73, 0x5c, 0x56, 0x31, 0x39, 0x5c, 0x52, 0x65, 0x73, 0x6f, 0x75, 0x72,
	0x63, 0x65, 0x73, 0xea, 0x02, 0x26, 0x47, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x3a, 0x3a, 0x41, 0x64,
	0x73, 0x3a, 0x3a, 0x47, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x41, 0x64, 0x73, 0x3a, 0x3a, 0x56, 0x31,
	0x39, 0x3a, 0x3a, 0x52, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x73, 0x62, 0x06, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x33,
})

var (
	file_resources_ad_group_bid_modifier_proto_rawDescOnce sync.Once
	file_resources_ad_group_bid_modifier_proto_rawDescData []byte
)

func file_resources_ad_group_bid_modifier_proto_rawDescGZIP() []byte {
	file_resources_ad_group_bid_modifier_proto_rawDescOnce.Do(func() {
		file_resources_ad_group_bid_modifier_proto_rawDescData = protoimpl.X.CompressGZIP(unsafe.Slice(unsafe.StringData(file_resources_ad_group_bid_modifier_proto_rawDesc), len(file_resources_ad_group_bid_modifier_proto_rawDesc)))
	})
	return file_resources_ad_group_bid_modifier_proto_rawDescData
}

var file_resources_ad_group_bid_modifier_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_resources_ad_group_bid_modifier_proto_goTypes = []any{
	(*AdGroupBidModifier)(nil),                         // 0: google.ads.googleads.v19.resources.AdGroupBidModifier
	(enums.BidModifierSourceEnum_BidModifierSource)(0), // 1: google.ads.googleads.v19.enums.BidModifierSourceEnum.BidModifierSource
	(*common.HotelDateSelectionTypeInfo)(nil),          // 2: google.ads.googleads.v19.common.HotelDateSelectionTypeInfo
	(*common.HotelAdvanceBookingWindowInfo)(nil),       // 3: google.ads.googleads.v19.common.HotelAdvanceBookingWindowInfo
	(*common.HotelLengthOfStayInfo)(nil),               // 4: google.ads.googleads.v19.common.HotelLengthOfStayInfo
	(*common.HotelCheckInDayInfo)(nil),                 // 5: google.ads.googleads.v19.common.HotelCheckInDayInfo
	(*common.DeviceInfo)(nil),                          // 6: google.ads.googleads.v19.common.DeviceInfo
	(*common.HotelCheckInDateRangeInfo)(nil),           // 7: google.ads.googleads.v19.common.HotelCheckInDateRangeInfo
}
var file_resources_ad_group_bid_modifier_proto_depIdxs = []int32{
	1, // 0: google.ads.googleads.v19.resources.AdGroupBidModifier.bid_modifier_source:type_name -> google.ads.googleads.v19.enums.BidModifierSourceEnum.BidModifierSource
	2, // 1: google.ads.googleads.v19.resources.AdGroupBidModifier.hotel_date_selection_type:type_name -> google.ads.googleads.v19.common.HotelDateSelectionTypeInfo
	3, // 2: google.ads.googleads.v19.resources.AdGroupBidModifier.hotel_advance_booking_window:type_name -> google.ads.googleads.v19.common.HotelAdvanceBookingWindowInfo
	4, // 3: google.ads.googleads.v19.resources.AdGroupBidModifier.hotel_length_of_stay:type_name -> google.ads.googleads.v19.common.HotelLengthOfStayInfo
	5, // 4: google.ads.googleads.v19.resources.AdGroupBidModifier.hotel_check_in_day:type_name -> google.ads.googleads.v19.common.HotelCheckInDayInfo
	6, // 5: google.ads.googleads.v19.resources.AdGroupBidModifier.device:type_name -> google.ads.googleads.v19.common.DeviceInfo
	7, // 6: google.ads.googleads.v19.resources.AdGroupBidModifier.hotel_check_in_date_range:type_name -> google.ads.googleads.v19.common.HotelCheckInDateRangeInfo
	7, // [7:7] is the sub-list for method output_type
	7, // [7:7] is the sub-list for method input_type
	7, // [7:7] is the sub-list for extension type_name
	7, // [7:7] is the sub-list for extension extendee
	0, // [0:7] is the sub-list for field type_name
}

func init() { file_resources_ad_group_bid_modifier_proto_init() }
func file_resources_ad_group_bid_modifier_proto_init() {
	if File_resources_ad_group_bid_modifier_proto != nil {
		return
	}
	file_resources_ad_group_bid_modifier_proto_msgTypes[0].OneofWrappers = []any{
		(*AdGroupBidModifier_HotelDateSelectionType)(nil),
		(*AdGroupBidModifier_HotelAdvanceBookingWindow)(nil),
		(*AdGroupBidModifier_HotelLengthOfStay)(nil),
		(*AdGroupBidModifier_HotelCheckInDay)(nil),
		(*AdGroupBidModifier_Device)(nil),
		(*AdGroupBidModifier_HotelCheckInDateRange)(nil),
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: unsafe.Slice(unsafe.StringData(file_resources_ad_group_bid_modifier_proto_rawDesc), len(file_resources_ad_group_bid_modifier_proto_rawDesc)),
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_resources_ad_group_bid_modifier_proto_goTypes,
		DependencyIndexes: file_resources_ad_group_bid_modifier_proto_depIdxs,
		MessageInfos:      file_resources_ad_group_bid_modifier_proto_msgTypes,
	}.Build()
	File_resources_ad_group_bid_modifier_proto = out.File
	file_resources_ad_group_bid_modifier_proto_goTypes = nil
	file_resources_ad_group_bid_modifier_proto_depIdxs = nil
}
