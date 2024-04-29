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
// source: google/ads/googleads/v16/resources/feed_item_target.proto

package resources

import (
	common "github.com/revealbot/google-ads-go/common"
	enums "github.com/revealbot/google-ads-go/enums"
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

// A feed item target.
type FeedItemTarget struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Immutable. The resource name of the feed item target.
	// Feed item target resource names have the form:
	// `customers/{customer_id}/feedItemTargets/{feed_id}~{feed_item_id}~{feed_item_target_type}~{feed_item_target_id}`
	ResourceName string `protobuf:"bytes,1,opt,name=resource_name,json=resourceName,proto3" json:"resource_name,omitempty"`
	// Immutable. The feed item to which this feed item target belongs.
	FeedItem *string `protobuf:"bytes,12,opt,name=feed_item,json=feedItem,proto3,oneof" json:"feed_item,omitempty"`
	// Output only. The target type of this feed item target. This field is
	// read-only.
	FeedItemTargetType enums.FeedItemTargetTypeEnum_FeedItemTargetType `protobuf:"varint,3,opt,name=feed_item_target_type,json=feedItemTargetType,proto3,enum=google.ads.googleads.v16.enums.FeedItemTargetTypeEnum_FeedItemTargetType" json:"feed_item_target_type,omitempty"`
	// Output only. The ID of the targeted resource. This field is read-only.
	FeedItemTargetId *int64 `protobuf:"varint,13,opt,name=feed_item_target_id,json=feedItemTargetId,proto3,oneof" json:"feed_item_target_id,omitempty"`
	// Output only. Status of the feed item target.
	// This field is read-only.
	Status enums.FeedItemTargetStatusEnum_FeedItemTargetStatus `protobuf:"varint,11,opt,name=status,proto3,enum=google.ads.googleads.v16.enums.FeedItemTargetStatusEnum_FeedItemTargetStatus" json:"status,omitempty"`
	// The targeted resource.
	//
	// Types that are assignable to Target:
	//
	//	*FeedItemTarget_Campaign
	//	*FeedItemTarget_AdGroup
	//	*FeedItemTarget_Keyword
	//	*FeedItemTarget_GeoTargetConstant
	//	*FeedItemTarget_Device
	//	*FeedItemTarget_AdSchedule
	Target isFeedItemTarget_Target `protobuf_oneof:"target"`
}

func (x *FeedItemTarget) Reset() {
	*x = FeedItemTarget{}
	if protoimpl.UnsafeEnabled {
		mi := &file_resources_feed_item_target_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FeedItemTarget) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FeedItemTarget) ProtoMessage() {}

func (x *FeedItemTarget) ProtoReflect() protoreflect.Message {
	mi := &file_resources_feed_item_target_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FeedItemTarget.ProtoReflect.Descriptor instead.
func (*FeedItemTarget) Descriptor() ([]byte, []int) {
	return file_resources_feed_item_target_proto_rawDescGZIP(), []int{0}
}

func (x *FeedItemTarget) GetResourceName() string {
	if x != nil {
		return x.ResourceName
	}
	return ""
}

func (x *FeedItemTarget) GetFeedItem() string {
	if x != nil && x.FeedItem != nil {
		return *x.FeedItem
	}
	return ""
}

func (x *FeedItemTarget) GetFeedItemTargetType() enums.FeedItemTargetTypeEnum_FeedItemTargetType {
	if x != nil {
		return x.FeedItemTargetType
	}
	return enums.FeedItemTargetTypeEnum_FeedItemTargetType(0)
}

func (x *FeedItemTarget) GetFeedItemTargetId() int64 {
	if x != nil && x.FeedItemTargetId != nil {
		return *x.FeedItemTargetId
	}
	return 0
}

func (x *FeedItemTarget) GetStatus() enums.FeedItemTargetStatusEnum_FeedItemTargetStatus {
	if x != nil {
		return x.Status
	}
	return enums.FeedItemTargetStatusEnum_FeedItemTargetStatus(0)
}

func (m *FeedItemTarget) GetTarget() isFeedItemTarget_Target {
	if m != nil {
		return m.Target
	}
	return nil
}

func (x *FeedItemTarget) GetCampaign() string {
	if x, ok := x.GetTarget().(*FeedItemTarget_Campaign); ok {
		return x.Campaign
	}
	return ""
}

func (x *FeedItemTarget) GetAdGroup() string {
	if x, ok := x.GetTarget().(*FeedItemTarget_AdGroup); ok {
		return x.AdGroup
	}
	return ""
}

func (x *FeedItemTarget) GetKeyword() *common.KeywordInfo {
	if x, ok := x.GetTarget().(*FeedItemTarget_Keyword); ok {
		return x.Keyword
	}
	return nil
}

func (x *FeedItemTarget) GetGeoTargetConstant() string {
	if x, ok := x.GetTarget().(*FeedItemTarget_GeoTargetConstant); ok {
		return x.GeoTargetConstant
	}
	return ""
}

func (x *FeedItemTarget) GetDevice() enums.FeedItemTargetDeviceEnum_FeedItemTargetDevice {
	if x, ok := x.GetTarget().(*FeedItemTarget_Device); ok {
		return x.Device
	}
	return enums.FeedItemTargetDeviceEnum_FeedItemTargetDevice(0)
}

func (x *FeedItemTarget) GetAdSchedule() *common.AdScheduleInfo {
	if x, ok := x.GetTarget().(*FeedItemTarget_AdSchedule); ok {
		return x.AdSchedule
	}
	return nil
}

type isFeedItemTarget_Target interface {
	isFeedItemTarget_Target()
}

type FeedItemTarget_Campaign struct {
	// Immutable. The targeted campaign.
	Campaign string `protobuf:"bytes,14,opt,name=campaign,proto3,oneof"`
}

type FeedItemTarget_AdGroup struct {
	// Immutable. The targeted ad group.
	AdGroup string `protobuf:"bytes,15,opt,name=ad_group,json=adGroup,proto3,oneof"`
}

type FeedItemTarget_Keyword struct {
	// Immutable. The targeted keyword.
	Keyword *common.KeywordInfo `protobuf:"bytes,7,opt,name=keyword,proto3,oneof"`
}

type FeedItemTarget_GeoTargetConstant struct {
	// Immutable. The targeted geo target constant resource name.
	GeoTargetConstant string `protobuf:"bytes,16,opt,name=geo_target_constant,json=geoTargetConstant,proto3,oneof"`
}

type FeedItemTarget_Device struct {
	// Immutable. The targeted device.
	Device enums.FeedItemTargetDeviceEnum_FeedItemTargetDevice `protobuf:"varint,9,opt,name=device,proto3,enum=google.ads.googleads.v16.enums.FeedItemTargetDeviceEnum_FeedItemTargetDevice,oneof"`
}

type FeedItemTarget_AdSchedule struct {
	// Immutable. The targeted schedule.
	AdSchedule *common.AdScheduleInfo `protobuf:"bytes,10,opt,name=ad_schedule,json=adSchedule,proto3,oneof"`
}

func (*FeedItemTarget_Campaign) isFeedItemTarget_Target() {}

func (*FeedItemTarget_AdGroup) isFeedItemTarget_Target() {}

func (*FeedItemTarget_Keyword) isFeedItemTarget_Target() {}

func (*FeedItemTarget_GeoTargetConstant) isFeedItemTarget_Target() {}

func (*FeedItemTarget_Device) isFeedItemTarget_Target() {}

func (*FeedItemTarget_AdSchedule) isFeedItemTarget_Target() {}

var File_resources_feed_item_target_proto protoreflect.FileDescriptor

var file_resources_feed_item_target_proto_rawDesc = []byte{
	0x0a, 0x39, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x64, 0x73, 0x2f, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x61, 0x64, 0x73, 0x2f, 0x76, 0x31, 0x36, 0x2f, 0x72, 0x65, 0x73, 0x6f, 0x75,
	0x72, 0x63, 0x65, 0x73, 0x2f, 0x66, 0x65, 0x65, 0x64, 0x5f, 0x69, 0x74, 0x65, 0x6d, 0x5f, 0x74,
	0x61, 0x72, 0x67, 0x65, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x22, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x2e, 0x61, 0x64, 0x73, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61, 0x64,
	0x73, 0x2e, 0x76, 0x31, 0x36, 0x2e, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x73, 0x1a,
	0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x64, 0x73, 0x2f, 0x67, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x61, 0x64, 0x73, 0x2f, 0x76, 0x31, 0x36, 0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e,
	0x2f, 0x63, 0x72, 0x69, 0x74, 0x65, 0x72, 0x69, 0x61, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a,
	0x3c, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x64, 0x73, 0x2f, 0x67, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x61, 0x64, 0x73, 0x2f, 0x76, 0x31, 0x36, 0x2f, 0x65, 0x6e, 0x75, 0x6d, 0x73, 0x2f,
	0x66, 0x65, 0x65, 0x64, 0x5f, 0x69, 0x74, 0x65, 0x6d, 0x5f, 0x74, 0x61, 0x72, 0x67, 0x65, 0x74,
	0x5f, 0x64, 0x65, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x3c, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x64, 0x73, 0x2f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x61, 0x64, 0x73, 0x2f, 0x76, 0x31, 0x36, 0x2f, 0x65, 0x6e, 0x75, 0x6d, 0x73, 0x2f, 0x66, 0x65,
	0x65, 0x64, 0x5f, 0x69, 0x74, 0x65, 0x6d, 0x5f, 0x74, 0x61, 0x72, 0x67, 0x65, 0x74, 0x5f, 0x73,
	0x74, 0x61, 0x74, 0x75, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x3a, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x2f, 0x61, 0x64, 0x73, 0x2f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61, 0x64,
	0x73, 0x2f, 0x76, 0x31, 0x36, 0x2f, 0x65, 0x6e, 0x75, 0x6d, 0x73, 0x2f, 0x66, 0x65, 0x65, 0x64,
	0x5f, 0x69, 0x74, 0x65, 0x6d, 0x5f, 0x74, 0x61, 0x72, 0x67, 0x65, 0x74, 0x5f, 0x74, 0x79, 0x70,
	0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f,
	0x61, 0x70, 0x69, 0x2f, 0x66, 0x69, 0x65, 0x6c, 0x64, 0x5f, 0x62, 0x65, 0x68, 0x61, 0x76, 0x69,
	0x6f, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x19, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x2f, 0x61, 0x70, 0x69, 0x2f, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x22, 0xb8, 0x09, 0x0a, 0x0e, 0x46, 0x65, 0x65, 0x64, 0x49, 0x74, 0x65, 0x6d,
	0x54, 0x61, 0x72, 0x67, 0x65, 0x74, 0x12, 0x54, 0x0a, 0x0d, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72,
	0x63, 0x65, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x42, 0x2f, 0xe0,
	0x41, 0x05, 0xfa, 0x41, 0x29, 0x0a, 0x27, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61, 0x64, 0x73,
	0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61, 0x70, 0x69, 0x73, 0x2e, 0x63, 0x6f, 0x6d, 0x2f,
	0x46, 0x65, 0x65, 0x64, 0x49, 0x74, 0x65, 0x6d, 0x54, 0x61, 0x72, 0x67, 0x65, 0x74, 0x52, 0x0c,
	0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x4b, 0x0a, 0x09,
	0x66, 0x65, 0x65, 0x64, 0x5f, 0x69, 0x74, 0x65, 0x6d, 0x18, 0x0c, 0x20, 0x01, 0x28, 0x09, 0x42,
	0x29, 0xe0, 0x41, 0x05, 0xfa, 0x41, 0x23, 0x0a, 0x21, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61,
	0x64, 0x73, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61, 0x70, 0x69, 0x73, 0x2e, 0x63, 0x6f,
	0x6d, 0x2f, 0x46, 0x65, 0x65, 0x64, 0x49, 0x74, 0x65, 0x6d, 0x48, 0x01, 0x52, 0x08, 0x66, 0x65,
	0x65, 0x64, 0x49, 0x74, 0x65, 0x6d, 0x88, 0x01, 0x01, 0x12, 0x81, 0x01, 0x0a, 0x15, 0x66, 0x65,
	0x65, 0x64, 0x5f, 0x69, 0x74, 0x65, 0x6d, 0x5f, 0x74, 0x61, 0x72, 0x67, 0x65, 0x74, 0x5f, 0x74,
	0x79, 0x70, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x49, 0x2e, 0x67, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x2e, 0x61, 0x64, 0x73, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61, 0x64, 0x73,
	0x2e, 0x76, 0x31, 0x36, 0x2e, 0x65, 0x6e, 0x75, 0x6d, 0x73, 0x2e, 0x46, 0x65, 0x65, 0x64, 0x49,
	0x74, 0x65, 0x6d, 0x54, 0x61, 0x72, 0x67, 0x65, 0x74, 0x54, 0x79, 0x70, 0x65, 0x45, 0x6e, 0x75,
	0x6d, 0x2e, 0x46, 0x65, 0x65, 0x64, 0x49, 0x74, 0x65, 0x6d, 0x54, 0x61, 0x72, 0x67, 0x65, 0x74,
	0x54, 0x79, 0x70, 0x65, 0x42, 0x03, 0xe0, 0x41, 0x03, 0x52, 0x12, 0x66, 0x65, 0x65, 0x64, 0x49,
	0x74, 0x65, 0x6d, 0x54, 0x61, 0x72, 0x67, 0x65, 0x74, 0x54, 0x79, 0x70, 0x65, 0x12, 0x37, 0x0a,
	0x13, 0x66, 0x65, 0x65, 0x64, 0x5f, 0x69, 0x74, 0x65, 0x6d, 0x5f, 0x74, 0x61, 0x72, 0x67, 0x65,
	0x74, 0x5f, 0x69, 0x64, 0x18, 0x0d, 0x20, 0x01, 0x28, 0x03, 0x42, 0x03, 0xe0, 0x41, 0x03, 0x48,
	0x02, 0x52, 0x10, 0x66, 0x65, 0x65, 0x64, 0x49, 0x74, 0x65, 0x6d, 0x54, 0x61, 0x72, 0x67, 0x65,
	0x74, 0x49, 0x64, 0x88, 0x01, 0x01, 0x12, 0x6a, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73,
	0x18, 0x0b, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x4d, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e,
	0x61, 0x64, 0x73, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61, 0x64, 0x73, 0x2e, 0x76, 0x31,
	0x36, 0x2e, 0x65, 0x6e, 0x75, 0x6d, 0x73, 0x2e, 0x46, 0x65, 0x65, 0x64, 0x49, 0x74, 0x65, 0x6d,
	0x54, 0x61, 0x72, 0x67, 0x65, 0x74, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x45, 0x6e, 0x75, 0x6d,
	0x2e, 0x46, 0x65, 0x65, 0x64, 0x49, 0x74, 0x65, 0x6d, 0x54, 0x61, 0x72, 0x67, 0x65, 0x74, 0x53,
	0x74, 0x61, 0x74, 0x75, 0x73, 0x42, 0x03, 0xe0, 0x41, 0x03, 0x52, 0x06, 0x73, 0x74, 0x61, 0x74,
	0x75, 0x73, 0x12, 0x47, 0x0a, 0x08, 0x63, 0x61, 0x6d, 0x70, 0x61, 0x69, 0x67, 0x6e, 0x18, 0x0e,
	0x20, 0x01, 0x28, 0x09, 0x42, 0x29, 0xe0, 0x41, 0x05, 0xfa, 0x41, 0x23, 0x0a, 0x21, 0x67, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x61, 0x64, 0x73, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61, 0x70,
	0x69, 0x73, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x43, 0x61, 0x6d, 0x70, 0x61, 0x69, 0x67, 0x6e, 0x48,
	0x00, 0x52, 0x08, 0x63, 0x61, 0x6d, 0x70, 0x61, 0x69, 0x67, 0x6e, 0x12, 0x45, 0x0a, 0x08, 0x61,
	0x64, 0x5f, 0x67, 0x72, 0x6f, 0x75, 0x70, 0x18, 0x0f, 0x20, 0x01, 0x28, 0x09, 0x42, 0x28, 0xe0,
	0x41, 0x05, 0xfa, 0x41, 0x22, 0x0a, 0x20, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61, 0x64, 0x73,
	0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61, 0x70, 0x69, 0x73, 0x2e, 0x63, 0x6f, 0x6d, 0x2f,
	0x41, 0x64, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x48, 0x00, 0x52, 0x07, 0x61, 0x64, 0x47, 0x72, 0x6f,
	0x75, 0x70, 0x12, 0x4d, 0x0a, 0x07, 0x6b, 0x65, 0x79, 0x77, 0x6f, 0x72, 0x64, 0x18, 0x07, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x2c, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x61, 0x64, 0x73,
	0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61, 0x64, 0x73, 0x2e, 0x76, 0x31, 0x36, 0x2e, 0x63,
	0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x4b, 0x65, 0x79, 0x77, 0x6f, 0x72, 0x64, 0x49, 0x6e, 0x66,
	0x6f, 0x42, 0x03, 0xe0, 0x41, 0x05, 0x48, 0x00, 0x52, 0x07, 0x6b, 0x65, 0x79, 0x77, 0x6f, 0x72,
	0x64, 0x12, 0x64, 0x0a, 0x13, 0x67, 0x65, 0x6f, 0x5f, 0x74, 0x61, 0x72, 0x67, 0x65, 0x74, 0x5f,
	0x63, 0x6f, 0x6e, 0x73, 0x74, 0x61, 0x6e, 0x74, 0x18, 0x10, 0x20, 0x01, 0x28, 0x09, 0x42, 0x32,
	0xe0, 0x41, 0x05, 0xfa, 0x41, 0x2c, 0x0a, 0x2a, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61, 0x64,
	0x73, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61, 0x70, 0x69, 0x73, 0x2e, 0x63, 0x6f, 0x6d,
	0x2f, 0x47, 0x65, 0x6f, 0x54, 0x61, 0x72, 0x67, 0x65, 0x74, 0x43, 0x6f, 0x6e, 0x73, 0x74, 0x61,
	0x6e, 0x74, 0x48, 0x00, 0x52, 0x11, 0x67, 0x65, 0x6f, 0x54, 0x61, 0x72, 0x67, 0x65, 0x74, 0x43,
	0x6f, 0x6e, 0x73, 0x74, 0x61, 0x6e, 0x74, 0x12, 0x6c, 0x0a, 0x06, 0x64, 0x65, 0x76, 0x69, 0x63,
	0x65, 0x18, 0x09, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x4d, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x2e, 0x61, 0x64, 0x73, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61, 0x64, 0x73, 0x2e, 0x76,
	0x31, 0x36, 0x2e, 0x65, 0x6e, 0x75, 0x6d, 0x73, 0x2e, 0x46, 0x65, 0x65, 0x64, 0x49, 0x74, 0x65,
	0x6d, 0x54, 0x61, 0x72, 0x67, 0x65, 0x74, 0x44, 0x65, 0x76, 0x69, 0x63, 0x65, 0x45, 0x6e, 0x75,
	0x6d, 0x2e, 0x46, 0x65, 0x65, 0x64, 0x49, 0x74, 0x65, 0x6d, 0x54, 0x61, 0x72, 0x67, 0x65, 0x74,
	0x44, 0x65, 0x76, 0x69, 0x63, 0x65, 0x42, 0x03, 0xe0, 0x41, 0x05, 0x48, 0x00, 0x52, 0x06, 0x64,
	0x65, 0x76, 0x69, 0x63, 0x65, 0x12, 0x57, 0x0a, 0x0b, 0x61, 0x64, 0x5f, 0x73, 0x63, 0x68, 0x65,
	0x64, 0x75, 0x6c, 0x65, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x2f, 0x2e, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x2e, 0x61, 0x64, 0x73, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61, 0x64,
	0x73, 0x2e, 0x76, 0x31, 0x36, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x41, 0x64, 0x53,
	0x63, 0x68, 0x65, 0x64, 0x75, 0x6c, 0x65, 0x49, 0x6e, 0x66, 0x6f, 0x42, 0x03, 0xe0, 0x41, 0x05,
	0x48, 0x00, 0x52, 0x0a, 0x61, 0x64, 0x53, 0x63, 0x68, 0x65, 0x64, 0x75, 0x6c, 0x65, 0x3a, 0x9d,
	0x01, 0xea, 0x41, 0x99, 0x01, 0x0a, 0x27, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61, 0x64, 0x73,
	0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61, 0x70, 0x69, 0x73, 0x2e, 0x63, 0x6f, 0x6d, 0x2f,
	0x46, 0x65, 0x65, 0x64, 0x49, 0x74, 0x65, 0x6d, 0x54, 0x61, 0x72, 0x67, 0x65, 0x74, 0x12, 0x6e,
	0x63, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x65, 0x72, 0x73, 0x2f, 0x7b, 0x63, 0x75, 0x73, 0x74, 0x6f,
	0x6d, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x7d, 0x2f, 0x66, 0x65, 0x65, 0x64, 0x49, 0x74, 0x65, 0x6d,
	0x54, 0x61, 0x72, 0x67, 0x65, 0x74, 0x73, 0x2f, 0x7b, 0x66, 0x65, 0x65, 0x64, 0x5f, 0x69, 0x64,
	0x7d, 0x7e, 0x7b, 0x66, 0x65, 0x65, 0x64, 0x5f, 0x69, 0x74, 0x65, 0x6d, 0x5f, 0x69, 0x64, 0x7d,
	0x7e, 0x7b, 0x66, 0x65, 0x65, 0x64, 0x5f, 0x69, 0x74, 0x65, 0x6d, 0x5f, 0x74, 0x61, 0x72, 0x67,
	0x65, 0x74, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x7d, 0x7e, 0x7b, 0x66, 0x65, 0x65, 0x64, 0x5f, 0x69,
	0x74, 0x65, 0x6d, 0x5f, 0x74, 0x61, 0x72, 0x67, 0x65, 0x74, 0x5f, 0x69, 0x64, 0x7d, 0x42, 0x08,
	0x0a, 0x06, 0x74, 0x61, 0x72, 0x67, 0x65, 0x74, 0x42, 0x0c, 0x0a, 0x0a, 0x5f, 0x66, 0x65, 0x65,
	0x64, 0x5f, 0x69, 0x74, 0x65, 0x6d, 0x42, 0x16, 0x0a, 0x14, 0x5f, 0x66, 0x65, 0x65, 0x64, 0x5f,
	0x69, 0x74, 0x65, 0x6d, 0x5f, 0x74, 0x61, 0x72, 0x67, 0x65, 0x74, 0x5f, 0x69, 0x64, 0x42, 0x85,
	0x02, 0x0a, 0x26, 0x63, 0x6f, 0x6d, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x61, 0x64,
	0x73, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61, 0x64, 0x73, 0x2e, 0x76, 0x31, 0x36, 0x2e,
	0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x73, 0x42, 0x13, 0x46, 0x65, 0x65, 0x64, 0x49,
	0x74, 0x65, 0x6d, 0x54, 0x61, 0x72, 0x67, 0x65, 0x74, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01,
	0x5a, 0x4b, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x67, 0x6f, 0x6c, 0x61, 0x6e, 0x67, 0x2e,
	0x6f, 0x72, 0x67, 0x2f, 0x67, 0x65, 0x6e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x61, 0x70, 0x69, 0x73, 0x2f, 0x61, 0x64, 0x73, 0x2f, 0x67, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x61, 0x64, 0x73, 0x2f, 0x76, 0x31, 0x36, 0x2f, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72,
	0x63, 0x65, 0x73, 0x3b, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x73, 0xa2, 0x02, 0x03,
	0x47, 0x41, 0x41, 0xaa, 0x02, 0x22, 0x47, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x41, 0x64, 0x73,
	0x2e, 0x47, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x41, 0x64, 0x73, 0x2e, 0x56, 0x31, 0x36, 0x2e, 0x52,
	0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x73, 0xca, 0x02, 0x22, 0x47, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x5c, 0x41, 0x64, 0x73, 0x5c, 0x47, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x41, 0x64, 0x73, 0x5c,
	0x56, 0x31, 0x36, 0x5c, 0x52, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x73, 0xea, 0x02, 0x26,
	0x47, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x3a, 0x3a, 0x41, 0x64, 0x73, 0x3a, 0x3a, 0x47, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x41, 0x64, 0x73, 0x3a, 0x3a, 0x56, 0x31, 0x36, 0x3a, 0x3a, 0x52, 0x65, 0x73,
	0x6f, 0x75, 0x72, 0x63, 0x65, 0x73, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_resources_feed_item_target_proto_rawDescOnce sync.Once
	file_resources_feed_item_target_proto_rawDescData = file_resources_feed_item_target_proto_rawDesc
)

func file_resources_feed_item_target_proto_rawDescGZIP() []byte {
	file_resources_feed_item_target_proto_rawDescOnce.Do(func() {
		file_resources_feed_item_target_proto_rawDescData = protoimpl.X.CompressGZIP(file_resources_feed_item_target_proto_rawDescData)
	})
	return file_resources_feed_item_target_proto_rawDescData
}

var file_resources_feed_item_target_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_resources_feed_item_target_proto_goTypes = []interface{}{
	(*FeedItemTarget)(nil),                                   // 0: google.ads.googleads.v16.resources.FeedItemTarget
	(enums.FeedItemTargetTypeEnum_FeedItemTargetType)(0),     // 1: google.ads.googleads.v16.enums.FeedItemTargetTypeEnum.FeedItemTargetType
	(enums.FeedItemTargetStatusEnum_FeedItemTargetStatus)(0), // 2: google.ads.googleads.v16.enums.FeedItemTargetStatusEnum.FeedItemTargetStatus
	(*common.KeywordInfo)(nil),                               // 3: google.ads.googleads.v16.common.KeywordInfo
	(enums.FeedItemTargetDeviceEnum_FeedItemTargetDevice)(0), // 4: google.ads.googleads.v16.enums.FeedItemTargetDeviceEnum.FeedItemTargetDevice
	(*common.AdScheduleInfo)(nil),                            // 5: google.ads.googleads.v16.common.AdScheduleInfo
}
var file_resources_feed_item_target_proto_depIdxs = []int32{
	1, // 0: google.ads.googleads.v16.resources.FeedItemTarget.feed_item_target_type:type_name -> google.ads.googleads.v16.enums.FeedItemTargetTypeEnum.FeedItemTargetType
	2, // 1: google.ads.googleads.v16.resources.FeedItemTarget.status:type_name -> google.ads.googleads.v16.enums.FeedItemTargetStatusEnum.FeedItemTargetStatus
	3, // 2: google.ads.googleads.v16.resources.FeedItemTarget.keyword:type_name -> google.ads.googleads.v16.common.KeywordInfo
	4, // 3: google.ads.googleads.v16.resources.FeedItemTarget.device:type_name -> google.ads.googleads.v16.enums.FeedItemTargetDeviceEnum.FeedItemTargetDevice
	5, // 4: google.ads.googleads.v16.resources.FeedItemTarget.ad_schedule:type_name -> google.ads.googleads.v16.common.AdScheduleInfo
	5, // [5:5] is the sub-list for method output_type
	5, // [5:5] is the sub-list for method input_type
	5, // [5:5] is the sub-list for extension type_name
	5, // [5:5] is the sub-list for extension extendee
	0, // [0:5] is the sub-list for field type_name
}

func init() { file_resources_feed_item_target_proto_init() }
func file_resources_feed_item_target_proto_init() {
	if File_resources_feed_item_target_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_resources_feed_item_target_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*FeedItemTarget); i {
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
	file_resources_feed_item_target_proto_msgTypes[0].OneofWrappers = []interface{}{
		(*FeedItemTarget_Campaign)(nil),
		(*FeedItemTarget_AdGroup)(nil),
		(*FeedItemTarget_Keyword)(nil),
		(*FeedItemTarget_GeoTargetConstant)(nil),
		(*FeedItemTarget_Device)(nil),
		(*FeedItemTarget_AdSchedule)(nil),
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_resources_feed_item_target_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_resources_feed_item_target_proto_goTypes,
		DependencyIndexes: file_resources_feed_item_target_proto_depIdxs,
		MessageInfos:      file_resources_feed_item_target_proto_msgTypes,
	}.Build()
	File_resources_feed_item_target_proto = out.File
	file_resources_feed_item_target_proto_rawDesc = nil
	file_resources_feed_item_target_proto_goTypes = nil
	file_resources_feed_item_target_proto_depIdxs = nil
}
