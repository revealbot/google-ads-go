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
// source: google/ads/googleads/v19/resources/bidding_seasonality_adjustment.proto

package resources

import (
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

// Represents a bidding seasonality adjustment. Cannot be used in manager
// accounts.
//
// See "About seasonality adjustments" at
// https://support.google.com/google-ads/answer/10369906.
type BiddingSeasonalityAdjustment struct {
	state protoimpl.MessageState `protogen:"open.v1"`
	// Immutable. The resource name of the seasonality adjustment.
	// Seasonality adjustment resource names have the form:
	//
	// `customers/{customer_id}/biddingSeasonalityAdjustments/{seasonality_adjustment_id}`
	ResourceName string `protobuf:"bytes,1,opt,name=resource_name,json=resourceName,proto3" json:"resource_name,omitempty"`
	// Output only. The ID of the seasonality adjustment.
	SeasonalityAdjustmentId int64 `protobuf:"varint,2,opt,name=seasonality_adjustment_id,json=seasonalityAdjustmentId,proto3" json:"seasonality_adjustment_id,omitempty"`
	// The scope of the seasonality adjustment.
	Scope enums.SeasonalityEventScopeEnum_SeasonalityEventScope `protobuf:"varint,3,opt,name=scope,proto3,enum=google.ads.googleads.v19.enums.SeasonalityEventScopeEnum_SeasonalityEventScope" json:"scope,omitempty"`
	// Output only. The status of the seasonality adjustment.
	Status enums.SeasonalityEventStatusEnum_SeasonalityEventStatus `protobuf:"varint,4,opt,name=status,proto3,enum=google.ads.googleads.v19.enums.SeasonalityEventStatusEnum_SeasonalityEventStatus" json:"status,omitempty"`
	// Required. The inclusive start time of the seasonality adjustment in
	// yyyy-MM-dd HH:mm:ss format.
	//
	// A seasonality adjustment is forward looking and should be used for events
	// that start and end in the future.
	StartDateTime string `protobuf:"bytes,5,opt,name=start_date_time,json=startDateTime,proto3" json:"start_date_time,omitempty"`
	// Required. The exclusive end time of the seasonality adjustment in
	// yyyy-MM-dd HH:mm:ss format.
	//
	// The length of [start_date_time, end_date_time) interval must be
	// within (0, 14 days].
	EndDateTime string `protobuf:"bytes,6,opt,name=end_date_time,json=endDateTime,proto3" json:"end_date_time,omitempty"`
	// The name of the seasonality adjustment. The name can be at most 255
	// characters.
	Name string `protobuf:"bytes,7,opt,name=name,proto3" json:"name,omitempty"`
	// The description of the seasonality adjustment. The description can be at
	// most 2048 characters.
	Description string `protobuf:"bytes,8,opt,name=description,proto3" json:"description,omitempty"`
	// If not specified, all devices will be included in this adjustment.
	// Otherwise, only the specified targeted devices will be included in this
	// adjustment.
	Devices []enums.DeviceEnum_Device `protobuf:"varint,9,rep,packed,name=devices,proto3,enum=google.ads.googleads.v19.enums.DeviceEnum_Device" json:"devices,omitempty"`
	// Conversion rate modifier estimated based on expected conversion rate
	// changes. When this field is unset or set to 1.0 no adjustment will be
	// applied to traffic. The allowed range is 0.1 to 10.0.
	ConversionRateModifier float64 `protobuf:"fixed64,10,opt,name=conversion_rate_modifier,json=conversionRateModifier,proto3" json:"conversion_rate_modifier,omitempty"`
	// The seasonality adjustment will apply to the campaigns listed when the
	// scope of this adjustment is CAMPAIGN. The maximum number of campaigns per
	// event is 2000.
	// Note: a seasonality adjustment with both advertising_channel_types and
	// campaign_ids is not supported.
	Campaigns []string `protobuf:"bytes,11,rep,name=campaigns,proto3" json:"campaigns,omitempty"`
	// The seasonality adjustment will apply to all the campaigns under the listed
	// channels retroactively as well as going forward when the scope of this
	// adjustment is CHANNEL.
	// The supported advertising channel types are DISPLAY, SEARCH and SHOPPING.
	// Note: a seasonality adjustment with both advertising_channel_types and
	// campaign_ids is not supported.
	AdvertisingChannelTypes []enums.AdvertisingChannelTypeEnum_AdvertisingChannelType `protobuf:"varint,12,rep,packed,name=advertising_channel_types,json=advertisingChannelTypes,proto3,enum=google.ads.googleads.v19.enums.AdvertisingChannelTypeEnum_AdvertisingChannelType" json:"advertising_channel_types,omitempty"`
	unknownFields           protoimpl.UnknownFields
	sizeCache               protoimpl.SizeCache
}

func (x *BiddingSeasonalityAdjustment) Reset() {
	*x = BiddingSeasonalityAdjustment{}
	mi := &file_resources_bidding_seasonality_adjustment_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *BiddingSeasonalityAdjustment) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BiddingSeasonalityAdjustment) ProtoMessage() {}

func (x *BiddingSeasonalityAdjustment) ProtoReflect() protoreflect.Message {
	mi := &file_resources_bidding_seasonality_adjustment_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use BiddingSeasonalityAdjustment.ProtoReflect.Descriptor instead.
func (*BiddingSeasonalityAdjustment) Descriptor() ([]byte, []int) {
	return file_resources_bidding_seasonality_adjustment_proto_rawDescGZIP(), []int{0}
}

func (x *BiddingSeasonalityAdjustment) GetResourceName() string {
	if x != nil {
		return x.ResourceName
	}
	return ""
}

func (x *BiddingSeasonalityAdjustment) GetSeasonalityAdjustmentId() int64 {
	if x != nil {
		return x.SeasonalityAdjustmentId
	}
	return 0
}

func (x *BiddingSeasonalityAdjustment) GetScope() enums.SeasonalityEventScopeEnum_SeasonalityEventScope {
	if x != nil {
		return x.Scope
	}
	return enums.SeasonalityEventScopeEnum_SeasonalityEventScope(0)
}

func (x *BiddingSeasonalityAdjustment) GetStatus() enums.SeasonalityEventStatusEnum_SeasonalityEventStatus {
	if x != nil {
		return x.Status
	}
	return enums.SeasonalityEventStatusEnum_SeasonalityEventStatus(0)
}

func (x *BiddingSeasonalityAdjustment) GetStartDateTime() string {
	if x != nil {
		return x.StartDateTime
	}
	return ""
}

func (x *BiddingSeasonalityAdjustment) GetEndDateTime() string {
	if x != nil {
		return x.EndDateTime
	}
	return ""
}

func (x *BiddingSeasonalityAdjustment) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *BiddingSeasonalityAdjustment) GetDescription() string {
	if x != nil {
		return x.Description
	}
	return ""
}

func (x *BiddingSeasonalityAdjustment) GetDevices() []enums.DeviceEnum_Device {
	if x != nil {
		return x.Devices
	}
	return nil
}

func (x *BiddingSeasonalityAdjustment) GetConversionRateModifier() float64 {
	if x != nil {
		return x.ConversionRateModifier
	}
	return 0
}

func (x *BiddingSeasonalityAdjustment) GetCampaigns() []string {
	if x != nil {
		return x.Campaigns
	}
	return nil
}

func (x *BiddingSeasonalityAdjustment) GetAdvertisingChannelTypes() []enums.AdvertisingChannelTypeEnum_AdvertisingChannelType {
	if x != nil {
		return x.AdvertisingChannelTypes
	}
	return nil
}

var File_resources_bidding_seasonality_adjustment_proto protoreflect.FileDescriptor

var file_resources_bidding_seasonality_adjustment_proto_rawDesc = string([]byte{
	0x0a, 0x47, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x64, 0x73, 0x2f, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x61, 0x64, 0x73, 0x2f, 0x76, 0x31, 0x39, 0x2f, 0x72, 0x65, 0x73, 0x6f, 0x75,
	0x72, 0x63, 0x65, 0x73, 0x2f, 0x62, 0x69, 0x64, 0x64, 0x69, 0x6e, 0x67, 0x5f, 0x73, 0x65, 0x61,
	0x73, 0x6f, 0x6e, 0x61, 0x6c, 0x69, 0x74, 0x79, 0x5f, 0x61, 0x64, 0x6a, 0x75, 0x73, 0x74, 0x6d,
	0x65, 0x6e, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x22, 0x67, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x2e, 0x61, 0x64, 0x73, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61, 0x64, 0x73, 0x2e,
	0x76, 0x31, 0x39, 0x2e, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x73, 0x1a, 0x3d, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x64, 0x73, 0x2f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x61, 0x64, 0x73, 0x2f, 0x76, 0x31, 0x39, 0x2f, 0x65, 0x6e, 0x75, 0x6d, 0x73, 0x2f, 0x61, 0x64,
	0x76, 0x65, 0x72, 0x74, 0x69, 0x73, 0x69, 0x6e, 0x67, 0x5f, 0x63, 0x68, 0x61, 0x6e, 0x6e, 0x65,
	0x6c, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x2b, 0x67, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x64, 0x73, 0x2f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61,
	0x64, 0x73, 0x2f, 0x76, 0x31, 0x39, 0x2f, 0x65, 0x6e, 0x75, 0x6d, 0x73, 0x2f, 0x64, 0x65, 0x76,
	0x69, 0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x3c, 0x67, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x2f, 0x61, 0x64, 0x73, 0x2f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61, 0x64, 0x73, 0x2f,
	0x76, 0x31, 0x39, 0x2f, 0x65, 0x6e, 0x75, 0x6d, 0x73, 0x2f, 0x73, 0x65, 0x61, 0x73, 0x6f, 0x6e,
	0x61, 0x6c, 0x69, 0x74, 0x79, 0x5f, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x5f, 0x73, 0x63, 0x6f, 0x70,
	0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x3d, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f,
	0x61, 0x64, 0x73, 0x2f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61, 0x64, 0x73, 0x2f, 0x76, 0x31,
	0x39, 0x2f, 0x65, 0x6e, 0x75, 0x6d, 0x73, 0x2f, 0x73, 0x65, 0x61, 0x73, 0x6f, 0x6e, 0x61, 0x6c,
	0x69, 0x74, 0x79, 0x5f, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x5f, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61,
	0x70, 0x69, 0x2f, 0x66, 0x69, 0x65, 0x6c, 0x64, 0x5f, 0x62, 0x65, 0x68, 0x61, 0x76, 0x69, 0x6f,
	0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x19, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f,
	0x61, 0x70, 0x69, 0x2f, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x22, 0x8f, 0x08, 0x0a, 0x1c, 0x42, 0x69, 0x64, 0x64, 0x69, 0x6e, 0x67, 0x53, 0x65,
	0x61, 0x73, 0x6f, 0x6e, 0x61, 0x6c, 0x69, 0x74, 0x79, 0x41, 0x64, 0x6a, 0x75, 0x73, 0x74, 0x6d,
	0x65, 0x6e, 0x74, 0x12, 0x62, 0x0a, 0x0d, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x5f,
	0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x42, 0x3d, 0xe0, 0x41, 0x05, 0xfa,
	0x41, 0x37, 0x0a, 0x35, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61, 0x64, 0x73, 0x2e, 0x67, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x61, 0x70, 0x69, 0x73, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x42, 0x69, 0x64,
	0x64, 0x69, 0x6e, 0x67, 0x53, 0x65, 0x61, 0x73, 0x6f, 0x6e, 0x61, 0x6c, 0x69, 0x74, 0x79, 0x41,
	0x64, 0x6a, 0x75, 0x73, 0x74, 0x6d, 0x65, 0x6e, 0x74, 0x52, 0x0c, 0x72, 0x65, 0x73, 0x6f, 0x75,
	0x72, 0x63, 0x65, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x3f, 0x0a, 0x19, 0x73, 0x65, 0x61, 0x73, 0x6f,
	0x6e, 0x61, 0x6c, 0x69, 0x74, 0x79, 0x5f, 0x61, 0x64, 0x6a, 0x75, 0x73, 0x74, 0x6d, 0x65, 0x6e,
	0x74, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x42, 0x03, 0xe0, 0x41, 0x03, 0x52,
	0x17, 0x73, 0x65, 0x61, 0x73, 0x6f, 0x6e, 0x61, 0x6c, 0x69, 0x74, 0x79, 0x41, 0x64, 0x6a, 0x75,
	0x73, 0x74, 0x6d, 0x65, 0x6e, 0x74, 0x49, 0x64, 0x12, 0x65, 0x0a, 0x05, 0x73, 0x63, 0x6f, 0x70,
	0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x4f, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x2e, 0x61, 0x64, 0x73, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61, 0x64, 0x73, 0x2e, 0x76,
	0x31, 0x39, 0x2e, 0x65, 0x6e, 0x75, 0x6d, 0x73, 0x2e, 0x53, 0x65, 0x61, 0x73, 0x6f, 0x6e, 0x61,
	0x6c, 0x69, 0x74, 0x79, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x53, 0x63, 0x6f, 0x70, 0x65, 0x45, 0x6e,
	0x75, 0x6d, 0x2e, 0x53, 0x65, 0x61, 0x73, 0x6f, 0x6e, 0x61, 0x6c, 0x69, 0x74, 0x79, 0x45, 0x76,
	0x65, 0x6e, 0x74, 0x53, 0x63, 0x6f, 0x70, 0x65, 0x52, 0x05, 0x73, 0x63, 0x6f, 0x70, 0x65, 0x12,
	0x6e, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0e, 0x32,
	0x51, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x61, 0x64, 0x73, 0x2e, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x61, 0x64, 0x73, 0x2e, 0x76, 0x31, 0x39, 0x2e, 0x65, 0x6e, 0x75, 0x6d, 0x73,
	0x2e, 0x53, 0x65, 0x61, 0x73, 0x6f, 0x6e, 0x61, 0x6c, 0x69, 0x74, 0x79, 0x45, 0x76, 0x65, 0x6e,
	0x74, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x45, 0x6e, 0x75, 0x6d, 0x2e, 0x53, 0x65, 0x61, 0x73,
	0x6f, 0x6e, 0x61, 0x6c, 0x69, 0x74, 0x79, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x53, 0x74, 0x61, 0x74,
	0x75, 0x73, 0x42, 0x03, 0xe0, 0x41, 0x03, 0x52, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12,
	0x2b, 0x0a, 0x0f, 0x73, 0x74, 0x61, 0x72, 0x74, 0x5f, 0x64, 0x61, 0x74, 0x65, 0x5f, 0x74, 0x69,
	0x6d, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x42, 0x03, 0xe0, 0x41, 0x02, 0x52, 0x0d, 0x73,
	0x74, 0x61, 0x72, 0x74, 0x44, 0x61, 0x74, 0x65, 0x54, 0x69, 0x6d, 0x65, 0x12, 0x27, 0x0a, 0x0d,
	0x65, 0x6e, 0x64, 0x5f, 0x64, 0x61, 0x74, 0x65, 0x5f, 0x74, 0x69, 0x6d, 0x65, 0x18, 0x06, 0x20,
	0x01, 0x28, 0x09, 0x42, 0x03, 0xe0, 0x41, 0x02, 0x52, 0x0b, 0x65, 0x6e, 0x64, 0x44, 0x61, 0x74,
	0x65, 0x54, 0x69, 0x6d, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x07, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x20, 0x0a, 0x0b, 0x64, 0x65, 0x73,
	0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x08, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b,
	0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x4b, 0x0a, 0x07, 0x64,
	0x65, 0x76, 0x69, 0x63, 0x65, 0x73, 0x18, 0x09, 0x20, 0x03, 0x28, 0x0e, 0x32, 0x31, 0x2e, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x61, 0x64, 0x73, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x61, 0x64, 0x73, 0x2e, 0x76, 0x31, 0x39, 0x2e, 0x65, 0x6e, 0x75, 0x6d, 0x73, 0x2e, 0x44, 0x65,
	0x76, 0x69, 0x63, 0x65, 0x45, 0x6e, 0x75, 0x6d, 0x2e, 0x44, 0x65, 0x76, 0x69, 0x63, 0x65, 0x52,
	0x07, 0x64, 0x65, 0x76, 0x69, 0x63, 0x65, 0x73, 0x12, 0x38, 0x0a, 0x18, 0x63, 0x6f, 0x6e, 0x76,
	0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x5f, 0x72, 0x61, 0x74, 0x65, 0x5f, 0x6d, 0x6f, 0x64, 0x69,
	0x66, 0x69, 0x65, 0x72, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x01, 0x52, 0x16, 0x63, 0x6f, 0x6e, 0x76,
	0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x52, 0x61, 0x74, 0x65, 0x4d, 0x6f, 0x64, 0x69, 0x66, 0x69,
	0x65, 0x72, 0x12, 0x44, 0x0a, 0x09, 0x63, 0x61, 0x6d, 0x70, 0x61, 0x69, 0x67, 0x6e, 0x73, 0x18,
	0x0b, 0x20, 0x03, 0x28, 0x09, 0x42, 0x26, 0xfa, 0x41, 0x23, 0x0a, 0x21, 0x67, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x61, 0x64, 0x73, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61, 0x70, 0x69, 0x73,
	0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x43, 0x61, 0x6d, 0x70, 0x61, 0x69, 0x67, 0x6e, 0x52, 0x09, 0x63,
	0x61, 0x6d, 0x70, 0x61, 0x69, 0x67, 0x6e, 0x73, 0x12, 0x8d, 0x01, 0x0a, 0x19, 0x61, 0x64, 0x76,
	0x65, 0x72, 0x74, 0x69, 0x73, 0x69, 0x6e, 0x67, 0x5f, 0x63, 0x68, 0x61, 0x6e, 0x6e, 0x65, 0x6c,
	0x5f, 0x74, 0x79, 0x70, 0x65, 0x73, 0x18, 0x0c, 0x20, 0x03, 0x28, 0x0e, 0x32, 0x51, 0x2e, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x61, 0x64, 0x73, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x61, 0x64, 0x73, 0x2e, 0x76, 0x31, 0x39, 0x2e, 0x65, 0x6e, 0x75, 0x6d, 0x73, 0x2e, 0x41, 0x64,
	0x76, 0x65, 0x72, 0x74, 0x69, 0x73, 0x69, 0x6e, 0x67, 0x43, 0x68, 0x61, 0x6e, 0x6e, 0x65, 0x6c,
	0x54, 0x79, 0x70, 0x65, 0x45, 0x6e, 0x75, 0x6d, 0x2e, 0x41, 0x64, 0x76, 0x65, 0x72, 0x74, 0x69,
	0x73, 0x69, 0x6e, 0x67, 0x43, 0x68, 0x61, 0x6e, 0x6e, 0x65, 0x6c, 0x54, 0x79, 0x70, 0x65, 0x52,
	0x17, 0x61, 0x64, 0x76, 0x65, 0x72, 0x74, 0x69, 0x73, 0x69, 0x6e, 0x67, 0x43, 0x68, 0x61, 0x6e,
	0x6e, 0x65, 0x6c, 0x54, 0x79, 0x70, 0x65, 0x73, 0x3a, 0x89, 0x01, 0xea, 0x41, 0x85, 0x01, 0x0a,
	0x35, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61, 0x64, 0x73, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x61, 0x70, 0x69, 0x73, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x42, 0x69, 0x64, 0x64, 0x69, 0x6e,
	0x67, 0x53, 0x65, 0x61, 0x73, 0x6f, 0x6e, 0x61, 0x6c, 0x69, 0x74, 0x79, 0x41, 0x64, 0x6a, 0x75,
	0x73, 0x74, 0x6d, 0x65, 0x6e, 0x74, 0x12, 0x4c, 0x63, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x65, 0x72,
	0x73, 0x2f, 0x7b, 0x63, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x7d, 0x2f,
	0x62, 0x69, 0x64, 0x64, 0x69, 0x6e, 0x67, 0x53, 0x65, 0x61, 0x73, 0x6f, 0x6e, 0x61, 0x6c, 0x69,
	0x74, 0x79, 0x41, 0x64, 0x6a, 0x75, 0x73, 0x74, 0x6d, 0x65, 0x6e, 0x74, 0x73, 0x2f, 0x7b, 0x73,
	0x65, 0x61, 0x73, 0x6f, 0x6e, 0x61, 0x6c, 0x69, 0x74, 0x79, 0x5f, 0x65, 0x76, 0x65, 0x6e, 0x74,
	0x5f, 0x69, 0x64, 0x7d, 0x42, 0x93, 0x02, 0x0a, 0x26, 0x63, 0x6f, 0x6d, 0x2e, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x2e, 0x61, 0x64, 0x73, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61, 0x64,
	0x73, 0x2e, 0x76, 0x31, 0x39, 0x2e, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x73, 0x42,
	0x21, 0x42, 0x69, 0x64, 0x64, 0x69, 0x6e, 0x67, 0x53, 0x65, 0x61, 0x73, 0x6f, 0x6e, 0x61, 0x6c,
	0x69, 0x74, 0x79, 0x41, 0x64, 0x6a, 0x75, 0x73, 0x74, 0x6d, 0x65, 0x6e, 0x74, 0x50, 0x72, 0x6f,
	0x74, 0x6f, 0x50, 0x01, 0x5a, 0x4b, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x67, 0x6f, 0x6c,
	0x61, 0x6e, 0x67, 0x2e, 0x6f, 0x72, 0x67, 0x2f, 0x67, 0x65, 0x6e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x2f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61, 0x70, 0x69, 0x73, 0x2f, 0x61, 0x64, 0x73, 0x2f,
	0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61, 0x64, 0x73, 0x2f, 0x76, 0x31, 0x39, 0x2f, 0x72, 0x65,
	0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x73, 0x3b, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65,
	0x73, 0xa2, 0x02, 0x03, 0x47, 0x41, 0x41, 0xaa, 0x02, 0x22, 0x47, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x2e, 0x41, 0x64, 0x73, 0x2e, 0x47, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x41, 0x64, 0x73, 0x2e, 0x56,
	0x31, 0x39, 0x2e, 0x52, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x73, 0xca, 0x02, 0x22, 0x47,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x5c, 0x41, 0x64, 0x73, 0x5c, 0x47, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x41, 0x64, 0x73, 0x5c, 0x56, 0x31, 0x39, 0x5c, 0x52, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65,
	0x73, 0xea, 0x02, 0x26, 0x47, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x3a, 0x3a, 0x41, 0x64, 0x73, 0x3a,
	0x3a, 0x47, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x41, 0x64, 0x73, 0x3a, 0x3a, 0x56, 0x31, 0x39, 0x3a,
	0x3a, 0x52, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x73, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x33,
})

var (
	file_resources_bidding_seasonality_adjustment_proto_rawDescOnce sync.Once
	file_resources_bidding_seasonality_adjustment_proto_rawDescData []byte
)

func file_resources_bidding_seasonality_adjustment_proto_rawDescGZIP() []byte {
	file_resources_bidding_seasonality_adjustment_proto_rawDescOnce.Do(func() {
		file_resources_bidding_seasonality_adjustment_proto_rawDescData = protoimpl.X.CompressGZIP(unsafe.Slice(unsafe.StringData(file_resources_bidding_seasonality_adjustment_proto_rawDesc), len(file_resources_bidding_seasonality_adjustment_proto_rawDesc)))
	})
	return file_resources_bidding_seasonality_adjustment_proto_rawDescData
}

var file_resources_bidding_seasonality_adjustment_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_resources_bidding_seasonality_adjustment_proto_goTypes = []any{
	(*BiddingSeasonalityAdjustment)(nil),                         // 0: google.ads.googleads.v19.resources.BiddingSeasonalityAdjustment
	(enums.SeasonalityEventScopeEnum_SeasonalityEventScope)(0),   // 1: google.ads.googleads.v19.enums.SeasonalityEventScopeEnum.SeasonalityEventScope
	(enums.SeasonalityEventStatusEnum_SeasonalityEventStatus)(0), // 2: google.ads.googleads.v19.enums.SeasonalityEventStatusEnum.SeasonalityEventStatus
	(enums.DeviceEnum_Device)(0),                                 // 3: google.ads.googleads.v19.enums.DeviceEnum.Device
	(enums.AdvertisingChannelTypeEnum_AdvertisingChannelType)(0), // 4: google.ads.googleads.v19.enums.AdvertisingChannelTypeEnum.AdvertisingChannelType
}
var file_resources_bidding_seasonality_adjustment_proto_depIdxs = []int32{
	1, // 0: google.ads.googleads.v19.resources.BiddingSeasonalityAdjustment.scope:type_name -> google.ads.googleads.v19.enums.SeasonalityEventScopeEnum.SeasonalityEventScope
	2, // 1: google.ads.googleads.v19.resources.BiddingSeasonalityAdjustment.status:type_name -> google.ads.googleads.v19.enums.SeasonalityEventStatusEnum.SeasonalityEventStatus
	3, // 2: google.ads.googleads.v19.resources.BiddingSeasonalityAdjustment.devices:type_name -> google.ads.googleads.v19.enums.DeviceEnum.Device
	4, // 3: google.ads.googleads.v19.resources.BiddingSeasonalityAdjustment.advertising_channel_types:type_name -> google.ads.googleads.v19.enums.AdvertisingChannelTypeEnum.AdvertisingChannelType
	4, // [4:4] is the sub-list for method output_type
	4, // [4:4] is the sub-list for method input_type
	4, // [4:4] is the sub-list for extension type_name
	4, // [4:4] is the sub-list for extension extendee
	0, // [0:4] is the sub-list for field type_name
}

func init() { file_resources_bidding_seasonality_adjustment_proto_init() }
func file_resources_bidding_seasonality_adjustment_proto_init() {
	if File_resources_bidding_seasonality_adjustment_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: unsafe.Slice(unsafe.StringData(file_resources_bidding_seasonality_adjustment_proto_rawDesc), len(file_resources_bidding_seasonality_adjustment_proto_rawDesc)),
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_resources_bidding_seasonality_adjustment_proto_goTypes,
		DependencyIndexes: file_resources_bidding_seasonality_adjustment_proto_depIdxs,
		MessageInfos:      file_resources_bidding_seasonality_adjustment_proto_msgTypes,
	}.Build()
	File_resources_bidding_seasonality_adjustment_proto = out.File
	file_resources_bidding_seasonality_adjustment_proto_goTypes = nil
	file_resources_bidding_seasonality_adjustment_proto_depIdxs = nil
}
