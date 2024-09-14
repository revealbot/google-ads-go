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
// source: google/ads/googleads/v17/resources/keyword_plan_campaign.proto

package resources

import (
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

// A Keyword Plan campaign.
// Max number of keyword plan campaigns per plan allowed: 1.
type KeywordPlanCampaign struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Immutable. The resource name of the Keyword Plan campaign.
	// KeywordPlanCampaign resource names have the form:
	//
	// `customers/{customer_id}/keywordPlanCampaigns/{kp_campaign_id}`
	ResourceName string `protobuf:"bytes,1,opt,name=resource_name,json=resourceName,proto3" json:"resource_name,omitempty"`
	// The keyword plan this campaign belongs to.
	KeywordPlan *string `protobuf:"bytes,9,opt,name=keyword_plan,json=keywordPlan,proto3,oneof" json:"keyword_plan,omitempty"`
	// Output only. The ID of the Keyword Plan campaign.
	Id *int64 `protobuf:"varint,10,opt,name=id,proto3,oneof" json:"id,omitempty"`
	// The name of the Keyword Plan campaign.
	//
	// This field is required and should not be empty when creating Keyword Plan
	// campaigns.
	Name *string `protobuf:"bytes,11,opt,name=name,proto3,oneof" json:"name,omitempty"`
	// The languages targeted for the Keyword Plan campaign.
	// Max allowed: 1.
	LanguageConstants []string `protobuf:"bytes,12,rep,name=language_constants,json=languageConstants,proto3" json:"language_constants,omitempty"`
	// Targeting network.
	//
	// This field is required and should not be empty when creating Keyword Plan
	// campaigns.
	KeywordPlanNetwork enums.KeywordPlanNetworkEnum_KeywordPlanNetwork `protobuf:"varint,6,opt,name=keyword_plan_network,json=keywordPlanNetwork,proto3,enum=google.ads.googleads.v17.enums.KeywordPlanNetworkEnum_KeywordPlanNetwork" json:"keyword_plan_network,omitempty"`
	// A default max cpc bid in micros, and in the account currency, for all ad
	// groups under the campaign.
	//
	// This field is required and should not be empty when creating Keyword Plan
	// campaigns.
	CpcBidMicros *int64 `protobuf:"varint,13,opt,name=cpc_bid_micros,json=cpcBidMicros,proto3,oneof" json:"cpc_bid_micros,omitempty"`
	// The geo targets.
	// Max number allowed: 20.
	GeoTargets []*KeywordPlanGeoTarget `protobuf:"bytes,8,rep,name=geo_targets,json=geoTargets,proto3" json:"geo_targets,omitempty"`
}

func (x *KeywordPlanCampaign) Reset() {
	*x = KeywordPlanCampaign{}
	if protoimpl.UnsafeEnabled {
		mi := &file_resources_keyword_plan_campaign_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *KeywordPlanCampaign) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*KeywordPlanCampaign) ProtoMessage() {}

func (x *KeywordPlanCampaign) ProtoReflect() protoreflect.Message {
	mi := &file_resources_keyword_plan_campaign_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use KeywordPlanCampaign.ProtoReflect.Descriptor instead.
func (*KeywordPlanCampaign) Descriptor() ([]byte, []int) {
	return file_resources_keyword_plan_campaign_proto_rawDescGZIP(), []int{0}
}

func (x *KeywordPlanCampaign) GetResourceName() string {
	if x != nil {
		return x.ResourceName
	}
	return ""
}

func (x *KeywordPlanCampaign) GetKeywordPlan() string {
	if x != nil && x.KeywordPlan != nil {
		return *x.KeywordPlan
	}
	return ""
}

func (x *KeywordPlanCampaign) GetId() int64 {
	if x != nil && x.Id != nil {
		return *x.Id
	}
	return 0
}

func (x *KeywordPlanCampaign) GetName() string {
	if x != nil && x.Name != nil {
		return *x.Name
	}
	return ""
}

func (x *KeywordPlanCampaign) GetLanguageConstants() []string {
	if x != nil {
		return x.LanguageConstants
	}
	return nil
}

func (x *KeywordPlanCampaign) GetKeywordPlanNetwork() enums.KeywordPlanNetworkEnum_KeywordPlanNetwork {
	if x != nil {
		return x.KeywordPlanNetwork
	}
	return enums.KeywordPlanNetworkEnum_KeywordPlanNetwork(0)
}

func (x *KeywordPlanCampaign) GetCpcBidMicros() int64 {
	if x != nil && x.CpcBidMicros != nil {
		return *x.CpcBidMicros
	}
	return 0
}

func (x *KeywordPlanCampaign) GetGeoTargets() []*KeywordPlanGeoTarget {
	if x != nil {
		return x.GeoTargets
	}
	return nil
}

// A geo target.
type KeywordPlanGeoTarget struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Required. The resource name of the geo target.
	GeoTargetConstant *string `protobuf:"bytes,2,opt,name=geo_target_constant,json=geoTargetConstant,proto3,oneof" json:"geo_target_constant,omitempty"`
}

func (x *KeywordPlanGeoTarget) Reset() {
	*x = KeywordPlanGeoTarget{}
	if protoimpl.UnsafeEnabled {
		mi := &file_resources_keyword_plan_campaign_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *KeywordPlanGeoTarget) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*KeywordPlanGeoTarget) ProtoMessage() {}

func (x *KeywordPlanGeoTarget) ProtoReflect() protoreflect.Message {
	mi := &file_resources_keyword_plan_campaign_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use KeywordPlanGeoTarget.ProtoReflect.Descriptor instead.
func (*KeywordPlanGeoTarget) Descriptor() ([]byte, []int) {
	return file_resources_keyword_plan_campaign_proto_rawDescGZIP(), []int{1}
}

func (x *KeywordPlanGeoTarget) GetGeoTargetConstant() string {
	if x != nil && x.GeoTargetConstant != nil {
		return *x.GeoTargetConstant
	}
	return ""
}

var File_resources_keyword_plan_campaign_proto protoreflect.FileDescriptor

var file_resources_keyword_plan_campaign_proto_rawDesc = []byte{
	0x0a, 0x3e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x64, 0x73, 0x2f, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x61, 0x64, 0x73, 0x2f, 0x76, 0x31, 0x37, 0x2f, 0x72, 0x65, 0x73, 0x6f, 0x75,
	0x72, 0x63, 0x65, 0x73, 0x2f, 0x6b, 0x65, 0x79, 0x77, 0x6f, 0x72, 0x64, 0x5f, 0x70, 0x6c, 0x61,
	0x6e, 0x5f, 0x63, 0x61, 0x6d, 0x70, 0x61, 0x69, 0x67, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x12, 0x22, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x61, 0x64, 0x73, 0x2e, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x61, 0x64, 0x73, 0x2e, 0x76, 0x31, 0x37, 0x2e, 0x72, 0x65, 0x73, 0x6f, 0x75,
	0x72, 0x63, 0x65, 0x73, 0x1a, 0x39, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x64, 0x73,
	0x2f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61, 0x64, 0x73, 0x2f, 0x76, 0x31, 0x37, 0x2f, 0x65,
	0x6e, 0x75, 0x6d, 0x73, 0x2f, 0x6b, 0x65, 0x79, 0x77, 0x6f, 0x72, 0x64, 0x5f, 0x70, 0x6c, 0x61,
	0x6e, 0x5f, 0x6e, 0x65, 0x74, 0x77, 0x6f, 0x72, 0x6b, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a,
	0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x66, 0x69, 0x65, 0x6c,
	0x64, 0x5f, 0x62, 0x65, 0x68, 0x61, 0x76, 0x69, 0x6f, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x1a, 0x19, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x72, 0x65, 0x73,
	0x6f, 0x75, 0x72, 0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x88, 0x06, 0x0a, 0x13,
	0x4b, 0x65, 0x79, 0x77, 0x6f, 0x72, 0x64, 0x50, 0x6c, 0x61, 0x6e, 0x43, 0x61, 0x6d, 0x70, 0x61,
	0x69, 0x67, 0x6e, 0x12, 0x59, 0x0a, 0x0d, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x5f,
	0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x42, 0x34, 0xe0, 0x41, 0x05, 0xfa,
	0x41, 0x2e, 0x0a, 0x2c, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61, 0x64, 0x73, 0x2e, 0x67, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x61, 0x70, 0x69, 0x73, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x4b, 0x65, 0x79,
	0x77, 0x6f, 0x72, 0x64, 0x50, 0x6c, 0x61, 0x6e, 0x43, 0x61, 0x6d, 0x70, 0x61, 0x69, 0x67, 0x6e,
	0x52, 0x0c, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x51,
	0x0a, 0x0c, 0x6b, 0x65, 0x79, 0x77, 0x6f, 0x72, 0x64, 0x5f, 0x70, 0x6c, 0x61, 0x6e, 0x18, 0x09,
	0x20, 0x01, 0x28, 0x09, 0x42, 0x29, 0xfa, 0x41, 0x26, 0x0a, 0x24, 0x67, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x61, 0x64, 0x73, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61, 0x70, 0x69, 0x73, 0x2e,
	0x63, 0x6f, 0x6d, 0x2f, 0x4b, 0x65, 0x79, 0x77, 0x6f, 0x72, 0x64, 0x50, 0x6c, 0x61, 0x6e, 0x48,
	0x00, 0x52, 0x0b, 0x6b, 0x65, 0x79, 0x77, 0x6f, 0x72, 0x64, 0x50, 0x6c, 0x61, 0x6e, 0x88, 0x01,
	0x01, 0x12, 0x18, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x03, 0x42, 0x03, 0xe0,
	0x41, 0x03, 0x48, 0x01, 0x52, 0x02, 0x69, 0x64, 0x88, 0x01, 0x01, 0x12, 0x17, 0x0a, 0x04, 0x6e,
	0x61, 0x6d, 0x65, 0x18, 0x0b, 0x20, 0x01, 0x28, 0x09, 0x48, 0x02, 0x52, 0x04, 0x6e, 0x61, 0x6d,
	0x65, 0x88, 0x01, 0x01, 0x12, 0x5d, 0x0a, 0x12, 0x6c, 0x61, 0x6e, 0x67, 0x75, 0x61, 0x67, 0x65,
	0x5f, 0x63, 0x6f, 0x6e, 0x73, 0x74, 0x61, 0x6e, 0x74, 0x73, 0x18, 0x0c, 0x20, 0x03, 0x28, 0x09,
	0x42, 0x2e, 0xfa, 0x41, 0x2b, 0x0a, 0x29, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61, 0x64, 0x73,
	0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61, 0x70, 0x69, 0x73, 0x2e, 0x63, 0x6f, 0x6d, 0x2f,
	0x4c, 0x61, 0x6e, 0x67, 0x75, 0x61, 0x67, 0x65, 0x43, 0x6f, 0x6e, 0x73, 0x74, 0x61, 0x6e, 0x74,
	0x52, 0x11, 0x6c, 0x61, 0x6e, 0x67, 0x75, 0x61, 0x67, 0x65, 0x43, 0x6f, 0x6e, 0x73, 0x74, 0x61,
	0x6e, 0x74, 0x73, 0x12, 0x7b, 0x0a, 0x14, 0x6b, 0x65, 0x79, 0x77, 0x6f, 0x72, 0x64, 0x5f, 0x70,
	0x6c, 0x61, 0x6e, 0x5f, 0x6e, 0x65, 0x74, 0x77, 0x6f, 0x72, 0x6b, 0x18, 0x06, 0x20, 0x01, 0x28,
	0x0e, 0x32, 0x49, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x61, 0x64, 0x73, 0x2e, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61, 0x64, 0x73, 0x2e, 0x76, 0x31, 0x37, 0x2e, 0x65, 0x6e, 0x75,
	0x6d, 0x73, 0x2e, 0x4b, 0x65, 0x79, 0x77, 0x6f, 0x72, 0x64, 0x50, 0x6c, 0x61, 0x6e, 0x4e, 0x65,
	0x74, 0x77, 0x6f, 0x72, 0x6b, 0x45, 0x6e, 0x75, 0x6d, 0x2e, 0x4b, 0x65, 0x79, 0x77, 0x6f, 0x72,
	0x64, 0x50, 0x6c, 0x61, 0x6e, 0x4e, 0x65, 0x74, 0x77, 0x6f, 0x72, 0x6b, 0x52, 0x12, 0x6b, 0x65,
	0x79, 0x77, 0x6f, 0x72, 0x64, 0x50, 0x6c, 0x61, 0x6e, 0x4e, 0x65, 0x74, 0x77, 0x6f, 0x72, 0x6b,
	0x12, 0x29, 0x0a, 0x0e, 0x63, 0x70, 0x63, 0x5f, 0x62, 0x69, 0x64, 0x5f, 0x6d, 0x69, 0x63, 0x72,
	0x6f, 0x73, 0x18, 0x0d, 0x20, 0x01, 0x28, 0x03, 0x48, 0x03, 0x52, 0x0c, 0x63, 0x70, 0x63, 0x42,
	0x69, 0x64, 0x4d, 0x69, 0x63, 0x72, 0x6f, 0x73, 0x88, 0x01, 0x01, 0x12, 0x59, 0x0a, 0x0b, 0x67,
	0x65, 0x6f, 0x5f, 0x74, 0x61, 0x72, 0x67, 0x65, 0x74, 0x73, 0x18, 0x08, 0x20, 0x03, 0x28, 0x0b,
	0x32, 0x38, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x61, 0x64, 0x73, 0x2e, 0x67, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x61, 0x64, 0x73, 0x2e, 0x76, 0x31, 0x37, 0x2e, 0x72, 0x65, 0x73, 0x6f,
	0x75, 0x72, 0x63, 0x65, 0x73, 0x2e, 0x4b, 0x65, 0x79, 0x77, 0x6f, 0x72, 0x64, 0x50, 0x6c, 0x61,
	0x6e, 0x47, 0x65, 0x6f, 0x54, 0x61, 0x72, 0x67, 0x65, 0x74, 0x52, 0x0a, 0x67, 0x65, 0x6f, 0x54,
	0x61, 0x72, 0x67, 0x65, 0x74, 0x73, 0x3a, 0x7a, 0xea, 0x41, 0x77, 0x0a, 0x2c, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x61, 0x64, 0x73, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61, 0x70, 0x69,
	0x73, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x4b, 0x65, 0x79, 0x77, 0x6f, 0x72, 0x64, 0x50, 0x6c, 0x61,
	0x6e, 0x43, 0x61, 0x6d, 0x70, 0x61, 0x69, 0x67, 0x6e, 0x12, 0x47, 0x63, 0x75, 0x73, 0x74, 0x6f,
	0x6d, 0x65, 0x72, 0x73, 0x2f, 0x7b, 0x63, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x65, 0x72, 0x5f, 0x69,
	0x64, 0x7d, 0x2f, 0x6b, 0x65, 0x79, 0x77, 0x6f, 0x72, 0x64, 0x50, 0x6c, 0x61, 0x6e, 0x43, 0x61,
	0x6d, 0x70, 0x61, 0x69, 0x67, 0x6e, 0x73, 0x2f, 0x7b, 0x6b, 0x65, 0x79, 0x77, 0x6f, 0x72, 0x64,
	0x5f, 0x70, 0x6c, 0x61, 0x6e, 0x5f, 0x63, 0x61, 0x6d, 0x70, 0x61, 0x69, 0x67, 0x6e, 0x5f, 0x69,
	0x64, 0x7d, 0x42, 0x0f, 0x0a, 0x0d, 0x5f, 0x6b, 0x65, 0x79, 0x77, 0x6f, 0x72, 0x64, 0x5f, 0x70,
	0x6c, 0x61, 0x6e, 0x42, 0x05, 0x0a, 0x03, 0x5f, 0x69, 0x64, 0x42, 0x07, 0x0a, 0x05, 0x5f, 0x6e,
	0x61, 0x6d, 0x65, 0x42, 0x11, 0x0a, 0x0f, 0x5f, 0x63, 0x70, 0x63, 0x5f, 0x62, 0x69, 0x64, 0x5f,
	0x6d, 0x69, 0x63, 0x72, 0x6f, 0x73, 0x22, 0x94, 0x01, 0x0a, 0x14, 0x4b, 0x65, 0x79, 0x77, 0x6f,
	0x72, 0x64, 0x50, 0x6c, 0x61, 0x6e, 0x47, 0x65, 0x6f, 0x54, 0x61, 0x72, 0x67, 0x65, 0x74, 0x12,
	0x64, 0x0a, 0x13, 0x67, 0x65, 0x6f, 0x5f, 0x74, 0x61, 0x72, 0x67, 0x65, 0x74, 0x5f, 0x63, 0x6f,
	0x6e, 0x73, 0x74, 0x61, 0x6e, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x42, 0x2f, 0xfa, 0x41,
	0x2c, 0x0a, 0x2a, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61, 0x64, 0x73, 0x2e, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x61, 0x70, 0x69, 0x73, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x47, 0x65, 0x6f, 0x54,
	0x61, 0x72, 0x67, 0x65, 0x74, 0x43, 0x6f, 0x6e, 0x73, 0x74, 0x61, 0x6e, 0x74, 0x48, 0x00, 0x52,
	0x11, 0x67, 0x65, 0x6f, 0x54, 0x61, 0x72, 0x67, 0x65, 0x74, 0x43, 0x6f, 0x6e, 0x73, 0x74, 0x61,
	0x6e, 0x74, 0x88, 0x01, 0x01, 0x42, 0x16, 0x0a, 0x14, 0x5f, 0x67, 0x65, 0x6f, 0x5f, 0x74, 0x61,
	0x72, 0x67, 0x65, 0x74, 0x5f, 0x63, 0x6f, 0x6e, 0x73, 0x74, 0x61, 0x6e, 0x74, 0x42, 0x8a, 0x02,
	0x0a, 0x26, 0x63, 0x6f, 0x6d, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x61, 0x64, 0x73,
	0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61, 0x64, 0x73, 0x2e, 0x76, 0x31, 0x37, 0x2e, 0x72,
	0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x73, 0x42, 0x18, 0x4b, 0x65, 0x79, 0x77, 0x6f, 0x72,
	0x64, 0x50, 0x6c, 0x61, 0x6e, 0x43, 0x61, 0x6d, 0x70, 0x61, 0x69, 0x67, 0x6e, 0x50, 0x72, 0x6f,
	0x74, 0x6f, 0x50, 0x01, 0x5a, 0x4b, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x67, 0x6f, 0x6c,
	0x61, 0x6e, 0x67, 0x2e, 0x6f, 0x72, 0x67, 0x2f, 0x67, 0x65, 0x6e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x2f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61, 0x70, 0x69, 0x73, 0x2f, 0x61, 0x64, 0x73, 0x2f,
	0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61, 0x64, 0x73, 0x2f, 0x76, 0x31, 0x37, 0x2f, 0x72, 0x65,
	0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x73, 0x3b, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65,
	0x73, 0xa2, 0x02, 0x03, 0x47, 0x41, 0x41, 0xaa, 0x02, 0x22, 0x47, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x2e, 0x41, 0x64, 0x73, 0x2e, 0x47, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x41, 0x64, 0x73, 0x2e, 0x56,
	0x31, 0x37, 0x2e, 0x52, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x73, 0xca, 0x02, 0x22, 0x47,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x5c, 0x41, 0x64, 0x73, 0x5c, 0x47, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x41, 0x64, 0x73, 0x5c, 0x56, 0x31, 0x37, 0x5c, 0x52, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65,
	0x73, 0xea, 0x02, 0x26, 0x47, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x3a, 0x3a, 0x41, 0x64, 0x73, 0x3a,
	0x3a, 0x47, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x41, 0x64, 0x73, 0x3a, 0x3a, 0x56, 0x31, 0x37, 0x3a,
	0x3a, 0x52, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x73, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x33,
}

var (
	file_resources_keyword_plan_campaign_proto_rawDescOnce sync.Once
	file_resources_keyword_plan_campaign_proto_rawDescData = file_resources_keyword_plan_campaign_proto_rawDesc
)

func file_resources_keyword_plan_campaign_proto_rawDescGZIP() []byte {
	file_resources_keyword_plan_campaign_proto_rawDescOnce.Do(func() {
		file_resources_keyword_plan_campaign_proto_rawDescData = protoimpl.X.CompressGZIP(file_resources_keyword_plan_campaign_proto_rawDescData)
	})
	return file_resources_keyword_plan_campaign_proto_rawDescData
}

var file_resources_keyword_plan_campaign_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_resources_keyword_plan_campaign_proto_goTypes = []interface{}{
	(*KeywordPlanCampaign)(nil),                          // 0: google.ads.googleads.v17.resources.KeywordPlanCampaign
	(*KeywordPlanGeoTarget)(nil),                         // 1: google.ads.googleads.v17.resources.KeywordPlanGeoTarget
	(enums.KeywordPlanNetworkEnum_KeywordPlanNetwork)(0), // 2: google.ads.googleads.v17.enums.KeywordPlanNetworkEnum.KeywordPlanNetwork
}
var file_resources_keyword_plan_campaign_proto_depIdxs = []int32{
	2, // 0: google.ads.googleads.v17.resources.KeywordPlanCampaign.keyword_plan_network:type_name -> google.ads.googleads.v17.enums.KeywordPlanNetworkEnum.KeywordPlanNetwork
	1, // 1: google.ads.googleads.v17.resources.KeywordPlanCampaign.geo_targets:type_name -> google.ads.googleads.v17.resources.KeywordPlanGeoTarget
	2, // [2:2] is the sub-list for method output_type
	2, // [2:2] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_resources_keyword_plan_campaign_proto_init() }
func file_resources_keyword_plan_campaign_proto_init() {
	if File_resources_keyword_plan_campaign_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_resources_keyword_plan_campaign_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*KeywordPlanCampaign); i {
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
		file_resources_keyword_plan_campaign_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*KeywordPlanGeoTarget); i {
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
	file_resources_keyword_plan_campaign_proto_msgTypes[0].OneofWrappers = []interface{}{}
	file_resources_keyword_plan_campaign_proto_msgTypes[1].OneofWrappers = []interface{}{}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_resources_keyword_plan_campaign_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_resources_keyword_plan_campaign_proto_goTypes,
		DependencyIndexes: file_resources_keyword_plan_campaign_proto_depIdxs,
		MessageInfos:      file_resources_keyword_plan_campaign_proto_msgTypes,
	}.Build()
	File_resources_keyword_plan_campaign_proto = out.File
	file_resources_keyword_plan_campaign_proto_rawDesc = nil
	file_resources_keyword_plan_campaign_proto_goTypes = nil
	file_resources_keyword_plan_campaign_proto_depIdxs = nil
}
