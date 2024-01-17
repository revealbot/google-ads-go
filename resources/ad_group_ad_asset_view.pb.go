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
// source: google/ads/googleads/v15/resources/ad_group_ad_asset_view.proto

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

// A link between an AdGroupAd and an Asset.
// Currently we only support AdGroupAdAssetView for AppAds and Responsive Search
// Ads.
type AdGroupAdAssetView struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Output only. The resource name of the ad group ad asset view.
	// Ad group ad asset view resource names have the form (Before V4):
	//
	// `customers/{customer_id}/adGroupAdAssets/{AdGroupAdAsset.ad_group_id}~{AdGroupAdAsset.ad.ad_id}~{AdGroupAdAsset.asset_id}~{AdGroupAdAsset.field_type}`
	//
	// Ad group ad asset view resource names have the form (Beginning from V4):
	//
	// `customers/{customer_id}/adGroupAdAssetViews/{AdGroupAdAsset.ad_group_id}~{AdGroupAdAsset.ad_id}~{AdGroupAdAsset.asset_id}~{AdGroupAdAsset.field_type}`
	ResourceName string `protobuf:"bytes,1,opt,name=resource_name,json=resourceName,proto3" json:"resource_name,omitempty"`
	// Output only. The ad group ad to which the asset is linked.
	AdGroupAd *string `protobuf:"bytes,9,opt,name=ad_group_ad,json=adGroupAd,proto3,oneof" json:"ad_group_ad,omitempty"`
	// Output only. The asset which is linked to the ad group ad.
	Asset *string `protobuf:"bytes,10,opt,name=asset,proto3,oneof" json:"asset,omitempty"`
	// Output only. Role that the asset takes in the ad.
	FieldType enums.AssetFieldTypeEnum_AssetFieldType `protobuf:"varint,2,opt,name=field_type,json=fieldType,proto3,enum=google.ads.googleads.v15.enums.AssetFieldTypeEnum_AssetFieldType" json:"field_type,omitempty"`
	// Output only. The status between the asset and the latest version of the ad.
	// If true, the asset is linked to the latest version of the ad. If false, it
	// means the link once existed but has been removed and is no longer present
	// in the latest version of the ad.
	Enabled *bool `protobuf:"varint,8,opt,name=enabled,proto3,oneof" json:"enabled,omitempty"`
	// Output only. Policy information for the ad group ad asset.
	PolicySummary *AdGroupAdAssetPolicySummary `protobuf:"bytes,3,opt,name=policy_summary,json=policySummary,proto3" json:"policy_summary,omitempty"`
	// Output only. Performance of an asset linkage.
	PerformanceLabel enums.AssetPerformanceLabelEnum_AssetPerformanceLabel `protobuf:"varint,4,opt,name=performance_label,json=performanceLabel,proto3,enum=google.ads.googleads.v15.enums.AssetPerformanceLabelEnum_AssetPerformanceLabel" json:"performance_label,omitempty"`
	// Output only. Pinned field.
	PinnedField enums.ServedAssetFieldTypeEnum_ServedAssetFieldType `protobuf:"varint,11,opt,name=pinned_field,json=pinnedField,proto3,enum=google.ads.googleads.v15.enums.ServedAssetFieldTypeEnum_ServedAssetFieldType" json:"pinned_field,omitempty"`
	// Output only. Source of the ad group ad asset.
	Source enums.AssetSourceEnum_AssetSource `protobuf:"varint,12,opt,name=source,proto3,enum=google.ads.googleads.v15.enums.AssetSourceEnum_AssetSource" json:"source,omitempty"`
}

func (x *AdGroupAdAssetView) Reset() {
	*x = AdGroupAdAssetView{}
	if protoimpl.UnsafeEnabled {
		mi := &file_resources_ad_group_ad_asset_view_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AdGroupAdAssetView) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AdGroupAdAssetView) ProtoMessage() {}

func (x *AdGroupAdAssetView) ProtoReflect() protoreflect.Message {
	mi := &file_resources_ad_group_ad_asset_view_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AdGroupAdAssetView.ProtoReflect.Descriptor instead.
func (*AdGroupAdAssetView) Descriptor() ([]byte, []int) {
	return file_resources_ad_group_ad_asset_view_proto_rawDescGZIP(), []int{0}
}

func (x *AdGroupAdAssetView) GetResourceName() string {
	if x != nil {
		return x.ResourceName
	}
	return ""
}

func (x *AdGroupAdAssetView) GetAdGroupAd() string {
	if x != nil && x.AdGroupAd != nil {
		return *x.AdGroupAd
	}
	return ""
}

func (x *AdGroupAdAssetView) GetAsset() string {
	if x != nil && x.Asset != nil {
		return *x.Asset
	}
	return ""
}

func (x *AdGroupAdAssetView) GetFieldType() enums.AssetFieldTypeEnum_AssetFieldType {
	if x != nil {
		return x.FieldType
	}
	return enums.AssetFieldTypeEnum_AssetFieldType(0)
}

func (x *AdGroupAdAssetView) GetEnabled() bool {
	if x != nil && x.Enabled != nil {
		return *x.Enabled
	}
	return false
}

func (x *AdGroupAdAssetView) GetPolicySummary() *AdGroupAdAssetPolicySummary {
	if x != nil {
		return x.PolicySummary
	}
	return nil
}

func (x *AdGroupAdAssetView) GetPerformanceLabel() enums.AssetPerformanceLabelEnum_AssetPerformanceLabel {
	if x != nil {
		return x.PerformanceLabel
	}
	return enums.AssetPerformanceLabelEnum_AssetPerformanceLabel(0)
}

func (x *AdGroupAdAssetView) GetPinnedField() enums.ServedAssetFieldTypeEnum_ServedAssetFieldType {
	if x != nil {
		return x.PinnedField
	}
	return enums.ServedAssetFieldTypeEnum_ServedAssetFieldType(0)
}

func (x *AdGroupAdAssetView) GetSource() enums.AssetSourceEnum_AssetSource {
	if x != nil {
		return x.Source
	}
	return enums.AssetSourceEnum_AssetSource(0)
}

// Contains policy information for an ad group ad asset.
type AdGroupAdAssetPolicySummary struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Output only. The list of policy findings for the ad group ad asset.
	PolicyTopicEntries []*common.PolicyTopicEntry `protobuf:"bytes,1,rep,name=policy_topic_entries,json=policyTopicEntries,proto3" json:"policy_topic_entries,omitempty"`
	// Output only. Where in the review process this ad group ad asset is.
	ReviewStatus enums.PolicyReviewStatusEnum_PolicyReviewStatus `protobuf:"varint,2,opt,name=review_status,json=reviewStatus,proto3,enum=google.ads.googleads.v15.enums.PolicyReviewStatusEnum_PolicyReviewStatus" json:"review_status,omitempty"`
	// Output only. The overall approval status of this ad group ad asset,
	// calculated based on the status of its individual policy topic entries.
	ApprovalStatus enums.PolicyApprovalStatusEnum_PolicyApprovalStatus `protobuf:"varint,3,opt,name=approval_status,json=approvalStatus,proto3,enum=google.ads.googleads.v15.enums.PolicyApprovalStatusEnum_PolicyApprovalStatus" json:"approval_status,omitempty"`
}

func (x *AdGroupAdAssetPolicySummary) Reset() {
	*x = AdGroupAdAssetPolicySummary{}
	if protoimpl.UnsafeEnabled {
		mi := &file_resources_ad_group_ad_asset_view_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AdGroupAdAssetPolicySummary) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AdGroupAdAssetPolicySummary) ProtoMessage() {}

func (x *AdGroupAdAssetPolicySummary) ProtoReflect() protoreflect.Message {
	mi := &file_resources_ad_group_ad_asset_view_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AdGroupAdAssetPolicySummary.ProtoReflect.Descriptor instead.
func (*AdGroupAdAssetPolicySummary) Descriptor() ([]byte, []int) {
	return file_resources_ad_group_ad_asset_view_proto_rawDescGZIP(), []int{1}
}

func (x *AdGroupAdAssetPolicySummary) GetPolicyTopicEntries() []*common.PolicyTopicEntry {
	if x != nil {
		return x.PolicyTopicEntries
	}
	return nil
}

func (x *AdGroupAdAssetPolicySummary) GetReviewStatus() enums.PolicyReviewStatusEnum_PolicyReviewStatus {
	if x != nil {
		return x.ReviewStatus
	}
	return enums.PolicyReviewStatusEnum_PolicyReviewStatus(0)
}

func (x *AdGroupAdAssetPolicySummary) GetApprovalStatus() enums.PolicyApprovalStatusEnum_PolicyApprovalStatus {
	if x != nil {
		return x.ApprovalStatus
	}
	return enums.PolicyApprovalStatusEnum_PolicyApprovalStatus(0)
}

var File_resources_ad_group_ad_asset_view_proto protoreflect.FileDescriptor

var file_resources_ad_group_ad_asset_view_proto_rawDesc = []byte{
	0x0a, 0x3f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x64, 0x73, 0x2f, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x61, 0x64, 0x73, 0x2f, 0x76, 0x31, 0x35, 0x2f, 0x72, 0x65, 0x73, 0x6f, 0x75,
	0x72, 0x63, 0x65, 0x73, 0x2f, 0x61, 0x64, 0x5f, 0x67, 0x72, 0x6f, 0x75, 0x70, 0x5f, 0x61, 0x64,
	0x5f, 0x61, 0x73, 0x73, 0x65, 0x74, 0x5f, 0x76, 0x69, 0x65, 0x77, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x12, 0x22, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x61, 0x64, 0x73, 0x2e, 0x67, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x61, 0x64, 0x73, 0x2e, 0x76, 0x31, 0x35, 0x2e, 0x72, 0x65, 0x73, 0x6f,
	0x75, 0x72, 0x63, 0x65, 0x73, 0x1a, 0x2c, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x64,
	0x73, 0x2f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61, 0x64, 0x73, 0x2f, 0x76, 0x31, 0x35, 0x2f,
	0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2f, 0x70, 0x6f, 0x6c, 0x69, 0x63, 0x79, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x1a, 0x35, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x64, 0x73, 0x2f,
	0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61, 0x64, 0x73, 0x2f, 0x76, 0x31, 0x35, 0x2f, 0x65, 0x6e,
	0x75, 0x6d, 0x73, 0x2f, 0x61, 0x73, 0x73, 0x65, 0x74, 0x5f, 0x66, 0x69, 0x65, 0x6c, 0x64, 0x5f,
	0x74, 0x79, 0x70, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x3c, 0x67, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x2f, 0x61, 0x64, 0x73, 0x2f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61, 0x64, 0x73,
	0x2f, 0x76, 0x31, 0x35, 0x2f, 0x65, 0x6e, 0x75, 0x6d, 0x73, 0x2f, 0x61, 0x73, 0x73, 0x65, 0x74,
	0x5f, 0x70, 0x65, 0x72, 0x66, 0x6f, 0x72, 0x6d, 0x61, 0x6e, 0x63, 0x65, 0x5f, 0x6c, 0x61, 0x62,
	0x65, 0x6c, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x31, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x2f, 0x61, 0x64, 0x73, 0x2f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61, 0x64, 0x73, 0x2f, 0x76,
	0x31, 0x35, 0x2f, 0x65, 0x6e, 0x75, 0x6d, 0x73, 0x2f, 0x61, 0x73, 0x73, 0x65, 0x74, 0x5f, 0x73,
	0x6f, 0x75, 0x72, 0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x3b, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x2f, 0x61, 0x64, 0x73, 0x2f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61, 0x64,
	0x73, 0x2f, 0x76, 0x31, 0x35, 0x2f, 0x65, 0x6e, 0x75, 0x6d, 0x73, 0x2f, 0x70, 0x6f, 0x6c, 0x69,
	0x63, 0x79, 0x5f, 0x61, 0x70, 0x70, 0x72, 0x6f, 0x76, 0x61, 0x6c, 0x5f, 0x73, 0x74, 0x61, 0x74,
	0x75, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x39, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x2f, 0x61, 0x64, 0x73, 0x2f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61, 0x64, 0x73, 0x2f, 0x76,
	0x31, 0x35, 0x2f, 0x65, 0x6e, 0x75, 0x6d, 0x73, 0x2f, 0x70, 0x6f, 0x6c, 0x69, 0x63, 0x79, 0x5f,
	0x72, 0x65, 0x76, 0x69, 0x65, 0x77, 0x5f, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x1a, 0x3c, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x64, 0x73, 0x2f,
	0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61, 0x64, 0x73, 0x2f, 0x76, 0x31, 0x35, 0x2f, 0x65, 0x6e,
	0x75, 0x6d, 0x73, 0x2f, 0x73, 0x65, 0x72, 0x76, 0x65, 0x64, 0x5f, 0x61, 0x73, 0x73, 0x65, 0x74,
	0x5f, 0x66, 0x69, 0x65, 0x6c, 0x64, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x66, 0x69,
	0x65, 0x6c, 0x64, 0x5f, 0x62, 0x65, 0x68, 0x61, 0x76, 0x69, 0x6f, 0x72, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x1a, 0x19, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x72,
	0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x84, 0x08,
	0x0a, 0x12, 0x41, 0x64, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x41, 0x64, 0x41, 0x73, 0x73, 0x65, 0x74,
	0x56, 0x69, 0x65, 0x77, 0x12, 0x58, 0x0a, 0x0d, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65,
	0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x42, 0x33, 0xe0, 0x41, 0x03,
	0xfa, 0x41, 0x2d, 0x0a, 0x2b, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61, 0x64, 0x73, 0x2e, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61, 0x70, 0x69, 0x73, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x41, 0x64,
	0x47, 0x72, 0x6f, 0x75, 0x70, 0x41, 0x64, 0x41, 0x73, 0x73, 0x65, 0x74, 0x56, 0x69, 0x65, 0x77,
	0x52, 0x0c, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x4f,
	0x0a, 0x0b, 0x61, 0x64, 0x5f, 0x67, 0x72, 0x6f, 0x75, 0x70, 0x5f, 0x61, 0x64, 0x18, 0x09, 0x20,
	0x01, 0x28, 0x09, 0x42, 0x2a, 0xe0, 0x41, 0x03, 0xfa, 0x41, 0x24, 0x0a, 0x22, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x61, 0x64, 0x73, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61, 0x70, 0x69,
	0x73, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x41, 0x64, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x41, 0x64, 0x48,
	0x00, 0x52, 0x09, 0x61, 0x64, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x41, 0x64, 0x88, 0x01, 0x01, 0x12,
	0x41, 0x0a, 0x05, 0x61, 0x73, 0x73, 0x65, 0x74, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x09, 0x42, 0x26,
	0xe0, 0x41, 0x03, 0xfa, 0x41, 0x20, 0x0a, 0x1e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61, 0x64,
	0x73, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61, 0x70, 0x69, 0x73, 0x2e, 0x63, 0x6f, 0x6d,
	0x2f, 0x41, 0x73, 0x73, 0x65, 0x74, 0x48, 0x01, 0x52, 0x05, 0x61, 0x73, 0x73, 0x65, 0x74, 0x88,
	0x01, 0x01, 0x12, 0x65, 0x0a, 0x0a, 0x66, 0x69, 0x65, 0x6c, 0x64, 0x5f, 0x74, 0x79, 0x70, 0x65,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x41, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e,
	0x61, 0x64, 0x73, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61, 0x64, 0x73, 0x2e, 0x76, 0x31,
	0x35, 0x2e, 0x65, 0x6e, 0x75, 0x6d, 0x73, 0x2e, 0x41, 0x73, 0x73, 0x65, 0x74, 0x46, 0x69, 0x65,
	0x6c, 0x64, 0x54, 0x79, 0x70, 0x65, 0x45, 0x6e, 0x75, 0x6d, 0x2e, 0x41, 0x73, 0x73, 0x65, 0x74,
	0x46, 0x69, 0x65, 0x6c, 0x64, 0x54, 0x79, 0x70, 0x65, 0x42, 0x03, 0xe0, 0x41, 0x03, 0x52, 0x09,
	0x66, 0x69, 0x65, 0x6c, 0x64, 0x54, 0x79, 0x70, 0x65, 0x12, 0x22, 0x0a, 0x07, 0x65, 0x6e, 0x61,
	0x62, 0x6c, 0x65, 0x64, 0x18, 0x08, 0x20, 0x01, 0x28, 0x08, 0x42, 0x03, 0xe0, 0x41, 0x03, 0x48,
	0x02, 0x52, 0x07, 0x65, 0x6e, 0x61, 0x62, 0x6c, 0x65, 0x64, 0x88, 0x01, 0x01, 0x12, 0x6b, 0x0a,
	0x0e, 0x70, 0x6f, 0x6c, 0x69, 0x63, 0x79, 0x5f, 0x73, 0x75, 0x6d, 0x6d, 0x61, 0x72, 0x79, 0x18,
	0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x3f, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x61,
	0x64, 0x73, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61, 0x64, 0x73, 0x2e, 0x76, 0x31, 0x35,
	0x2e, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x73, 0x2e, 0x41, 0x64, 0x47, 0x72, 0x6f,
	0x75, 0x70, 0x41, 0x64, 0x41, 0x73, 0x73, 0x65, 0x74, 0x50, 0x6f, 0x6c, 0x69, 0x63, 0x79, 0x53,
	0x75, 0x6d, 0x6d, 0x61, 0x72, 0x79, 0x42, 0x03, 0xe0, 0x41, 0x03, 0x52, 0x0d, 0x70, 0x6f, 0x6c,
	0x69, 0x63, 0x79, 0x53, 0x75, 0x6d, 0x6d, 0x61, 0x72, 0x79, 0x12, 0x81, 0x01, 0x0a, 0x11, 0x70,
	0x65, 0x72, 0x66, 0x6f, 0x72, 0x6d, 0x61, 0x6e, 0x63, 0x65, 0x5f, 0x6c, 0x61, 0x62, 0x65, 0x6c,
	0x18, 0x04, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x4f, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e,
	0x61, 0x64, 0x73, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61, 0x64, 0x73, 0x2e, 0x76, 0x31,
	0x35, 0x2e, 0x65, 0x6e, 0x75, 0x6d, 0x73, 0x2e, 0x41, 0x73, 0x73, 0x65, 0x74, 0x50, 0x65, 0x72,
	0x66, 0x6f, 0x72, 0x6d, 0x61, 0x6e, 0x63, 0x65, 0x4c, 0x61, 0x62, 0x65, 0x6c, 0x45, 0x6e, 0x75,
	0x6d, 0x2e, 0x41, 0x73, 0x73, 0x65, 0x74, 0x50, 0x65, 0x72, 0x66, 0x6f, 0x72, 0x6d, 0x61, 0x6e,
	0x63, 0x65, 0x4c, 0x61, 0x62, 0x65, 0x6c, 0x42, 0x03, 0xe0, 0x41, 0x03, 0x52, 0x10, 0x70, 0x65,
	0x72, 0x66, 0x6f, 0x72, 0x6d, 0x61, 0x6e, 0x63, 0x65, 0x4c, 0x61, 0x62, 0x65, 0x6c, 0x12, 0x75,
	0x0a, 0x0c, 0x70, 0x69, 0x6e, 0x6e, 0x65, 0x64, 0x5f, 0x66, 0x69, 0x65, 0x6c, 0x64, 0x18, 0x0b,
	0x20, 0x01, 0x28, 0x0e, 0x32, 0x4d, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x61, 0x64,
	0x73, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61, 0x64, 0x73, 0x2e, 0x76, 0x31, 0x35, 0x2e,
	0x65, 0x6e, 0x75, 0x6d, 0x73, 0x2e, 0x53, 0x65, 0x72, 0x76, 0x65, 0x64, 0x41, 0x73, 0x73, 0x65,
	0x74, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x54, 0x79, 0x70, 0x65, 0x45, 0x6e, 0x75, 0x6d, 0x2e, 0x53,
	0x65, 0x72, 0x76, 0x65, 0x64, 0x41, 0x73, 0x73, 0x65, 0x74, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x54,
	0x79, 0x70, 0x65, 0x42, 0x03, 0xe0, 0x41, 0x03, 0x52, 0x0b, 0x70, 0x69, 0x6e, 0x6e, 0x65, 0x64,
	0x46, 0x69, 0x65, 0x6c, 0x64, 0x12, 0x58, 0x0a, 0x06, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x18,
	0x0c, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x3b, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x61,
	0x64, 0x73, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61, 0x64, 0x73, 0x2e, 0x76, 0x31, 0x35,
	0x2e, 0x65, 0x6e, 0x75, 0x6d, 0x73, 0x2e, 0x41, 0x73, 0x73, 0x65, 0x74, 0x53, 0x6f, 0x75, 0x72,
	0x63, 0x65, 0x45, 0x6e, 0x75, 0x6d, 0x2e, 0x41, 0x73, 0x73, 0x65, 0x74, 0x53, 0x6f, 0x75, 0x72,
	0x63, 0x65, 0x42, 0x03, 0xe0, 0x41, 0x03, 0x52, 0x06, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x3a,
	0x8c, 0x01, 0xea, 0x41, 0x88, 0x01, 0x0a, 0x2b, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61, 0x64,
	0x73, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61, 0x70, 0x69, 0x73, 0x2e, 0x63, 0x6f, 0x6d,
	0x2f, 0x41, 0x64, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x41, 0x64, 0x41, 0x73, 0x73, 0x65, 0x74, 0x56,
	0x69, 0x65, 0x77, 0x12, 0x59, 0x63, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x65, 0x72, 0x73, 0x2f, 0x7b,
	0x63, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x7d, 0x2f, 0x61, 0x64, 0x47,
	0x72, 0x6f, 0x75, 0x70, 0x41, 0x64, 0x41, 0x73, 0x73, 0x65, 0x74, 0x56, 0x69, 0x65, 0x77, 0x73,
	0x2f, 0x7b, 0x61, 0x64, 0x5f, 0x67, 0x72, 0x6f, 0x75, 0x70, 0x5f, 0x69, 0x64, 0x7d, 0x7e, 0x7b,
	0x61, 0x64, 0x5f, 0x69, 0x64, 0x7d, 0x7e, 0x7b, 0x61, 0x73, 0x73, 0x65, 0x74, 0x5f, 0x69, 0x64,
	0x7d, 0x7e, 0x7b, 0x66, 0x69, 0x65, 0x6c, 0x64, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x7d, 0x42, 0x0e,
	0x0a, 0x0c, 0x5f, 0x61, 0x64, 0x5f, 0x67, 0x72, 0x6f, 0x75, 0x70, 0x5f, 0x61, 0x64, 0x42, 0x08,
	0x0a, 0x06, 0x5f, 0x61, 0x73, 0x73, 0x65, 0x74, 0x42, 0x0a, 0x0a, 0x08, 0x5f, 0x65, 0x6e, 0x61,
	0x62, 0x6c, 0x65, 0x64, 0x22, 0xf9, 0x02, 0x0a, 0x1b, 0x41, 0x64, 0x47, 0x72, 0x6f, 0x75, 0x70,
	0x41, 0x64, 0x41, 0x73, 0x73, 0x65, 0x74, 0x50, 0x6f, 0x6c, 0x69, 0x63, 0x79, 0x53, 0x75, 0x6d,
	0x6d, 0x61, 0x72, 0x79, 0x12, 0x68, 0x0a, 0x14, 0x70, 0x6f, 0x6c, 0x69, 0x63, 0x79, 0x5f, 0x74,
	0x6f, 0x70, 0x69, 0x63, 0x5f, 0x65, 0x6e, 0x74, 0x72, 0x69, 0x65, 0x73, 0x18, 0x01, 0x20, 0x03,
	0x28, 0x0b, 0x32, 0x31, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x61, 0x64, 0x73, 0x2e,
	0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61, 0x64, 0x73, 0x2e, 0x76, 0x31, 0x35, 0x2e, 0x63, 0x6f,
	0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x50, 0x6f, 0x6c, 0x69, 0x63, 0x79, 0x54, 0x6f, 0x70, 0x69, 0x63,
	0x45, 0x6e, 0x74, 0x72, 0x79, 0x42, 0x03, 0xe0, 0x41, 0x03, 0x52, 0x12, 0x70, 0x6f, 0x6c, 0x69,
	0x63, 0x79, 0x54, 0x6f, 0x70, 0x69, 0x63, 0x45, 0x6e, 0x74, 0x72, 0x69, 0x65, 0x73, 0x12, 0x73,
	0x0a, 0x0d, 0x72, 0x65, 0x76, 0x69, 0x65, 0x77, 0x5f, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x49, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x61,
	0x64, 0x73, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61, 0x64, 0x73, 0x2e, 0x76, 0x31, 0x35,
	0x2e, 0x65, 0x6e, 0x75, 0x6d, 0x73, 0x2e, 0x50, 0x6f, 0x6c, 0x69, 0x63, 0x79, 0x52, 0x65, 0x76,
	0x69, 0x65, 0x77, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x45, 0x6e, 0x75, 0x6d, 0x2e, 0x50, 0x6f,
	0x6c, 0x69, 0x63, 0x79, 0x52, 0x65, 0x76, 0x69, 0x65, 0x77, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73,
	0x42, 0x03, 0xe0, 0x41, 0x03, 0x52, 0x0c, 0x72, 0x65, 0x76, 0x69, 0x65, 0x77, 0x53, 0x74, 0x61,
	0x74, 0x75, 0x73, 0x12, 0x7b, 0x0a, 0x0f, 0x61, 0x70, 0x70, 0x72, 0x6f, 0x76, 0x61, 0x6c, 0x5f,
	0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x4d, 0x2e, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x61, 0x64, 0x73, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x61, 0x64, 0x73, 0x2e, 0x76, 0x31, 0x35, 0x2e, 0x65, 0x6e, 0x75, 0x6d, 0x73, 0x2e, 0x50, 0x6f,
	0x6c, 0x69, 0x63, 0x79, 0x41, 0x70, 0x70, 0x72, 0x6f, 0x76, 0x61, 0x6c, 0x53, 0x74, 0x61, 0x74,
	0x75, 0x73, 0x45, 0x6e, 0x75, 0x6d, 0x2e, 0x50, 0x6f, 0x6c, 0x69, 0x63, 0x79, 0x41, 0x70, 0x70,
	0x72, 0x6f, 0x76, 0x61, 0x6c, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x42, 0x03, 0xe0, 0x41, 0x03,
	0x52, 0x0e, 0x61, 0x70, 0x70, 0x72, 0x6f, 0x76, 0x61, 0x6c, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73,
	0x42, 0x89, 0x02, 0x0a, 0x26, 0x63, 0x6f, 0x6d, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e,
	0x61, 0x64, 0x73, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61, 0x64, 0x73, 0x2e, 0x76, 0x31,
	0x35, 0x2e, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x73, 0x42, 0x17, 0x41, 0x64, 0x47,
	0x72, 0x6f, 0x75, 0x70, 0x41, 0x64, 0x41, 0x73, 0x73, 0x65, 0x74, 0x56, 0x69, 0x65, 0x77, 0x50,
	0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a, 0x4b, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x67,
	0x6f, 0x6c, 0x61, 0x6e, 0x67, 0x2e, 0x6f, 0x72, 0x67, 0x2f, 0x67, 0x65, 0x6e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x2f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61, 0x70, 0x69, 0x73, 0x2f, 0x61, 0x64,
	0x73, 0x2f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61, 0x64, 0x73, 0x2f, 0x76, 0x31, 0x35, 0x2f,
	0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x73, 0x3b, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72,
	0x63, 0x65, 0x73, 0xa2, 0x02, 0x03, 0x47, 0x41, 0x41, 0xaa, 0x02, 0x22, 0x47, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x2e, 0x41, 0x64, 0x73, 0x2e, 0x47, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x41, 0x64, 0x73,
	0x2e, 0x56, 0x31, 0x35, 0x2e, 0x52, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x73, 0xca, 0x02,
	0x22, 0x47, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x5c, 0x41, 0x64, 0x73, 0x5c, 0x47, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x41, 0x64, 0x73, 0x5c, 0x56, 0x31, 0x35, 0x5c, 0x52, 0x65, 0x73, 0x6f, 0x75, 0x72,
	0x63, 0x65, 0x73, 0xea, 0x02, 0x26, 0x47, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x3a, 0x3a, 0x41, 0x64,
	0x73, 0x3a, 0x3a, 0x47, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x41, 0x64, 0x73, 0x3a, 0x3a, 0x56, 0x31,
	0x35, 0x3a, 0x3a, 0x52, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x73, 0x62, 0x06, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_resources_ad_group_ad_asset_view_proto_rawDescOnce sync.Once
	file_resources_ad_group_ad_asset_view_proto_rawDescData = file_resources_ad_group_ad_asset_view_proto_rawDesc
)

func file_resources_ad_group_ad_asset_view_proto_rawDescGZIP() []byte {
	file_resources_ad_group_ad_asset_view_proto_rawDescOnce.Do(func() {
		file_resources_ad_group_ad_asset_view_proto_rawDescData = protoimpl.X.CompressGZIP(file_resources_ad_group_ad_asset_view_proto_rawDescData)
	})
	return file_resources_ad_group_ad_asset_view_proto_rawDescData
}

var file_resources_ad_group_ad_asset_view_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_resources_ad_group_ad_asset_view_proto_goTypes = []interface{}{
	(*AdGroupAdAssetView)(nil),                                 // 0: google.ads.googleads.v15.resources.AdGroupAdAssetView
	(*AdGroupAdAssetPolicySummary)(nil),                        // 1: google.ads.googleads.v15.resources.AdGroupAdAssetPolicySummary
	(enums.AssetFieldTypeEnum_AssetFieldType)(0),               // 2: google.ads.googleads.v15.enums.AssetFieldTypeEnum.AssetFieldType
	(enums.AssetPerformanceLabelEnum_AssetPerformanceLabel)(0), // 3: google.ads.googleads.v15.enums.AssetPerformanceLabelEnum.AssetPerformanceLabel
	(enums.ServedAssetFieldTypeEnum_ServedAssetFieldType)(0),   // 4: google.ads.googleads.v15.enums.ServedAssetFieldTypeEnum.ServedAssetFieldType
	(enums.AssetSourceEnum_AssetSource)(0),                     // 5: google.ads.googleads.v15.enums.AssetSourceEnum.AssetSource
	(*common.PolicyTopicEntry)(nil),                            // 6: google.ads.googleads.v15.common.PolicyTopicEntry
	(enums.PolicyReviewStatusEnum_PolicyReviewStatus)(0),       // 7: google.ads.googleads.v15.enums.PolicyReviewStatusEnum.PolicyReviewStatus
	(enums.PolicyApprovalStatusEnum_PolicyApprovalStatus)(0),   // 8: google.ads.googleads.v15.enums.PolicyApprovalStatusEnum.PolicyApprovalStatus
}
var file_resources_ad_group_ad_asset_view_proto_depIdxs = []int32{
	2, // 0: google.ads.googleads.v15.resources.AdGroupAdAssetView.field_type:type_name -> google.ads.googleads.v15.enums.AssetFieldTypeEnum.AssetFieldType
	1, // 1: google.ads.googleads.v15.resources.AdGroupAdAssetView.policy_summary:type_name -> google.ads.googleads.v15.resources.AdGroupAdAssetPolicySummary
	3, // 2: google.ads.googleads.v15.resources.AdGroupAdAssetView.performance_label:type_name -> google.ads.googleads.v15.enums.AssetPerformanceLabelEnum.AssetPerformanceLabel
	4, // 3: google.ads.googleads.v15.resources.AdGroupAdAssetView.pinned_field:type_name -> google.ads.googleads.v15.enums.ServedAssetFieldTypeEnum.ServedAssetFieldType
	5, // 4: google.ads.googleads.v15.resources.AdGroupAdAssetView.source:type_name -> google.ads.googleads.v15.enums.AssetSourceEnum.AssetSource
	6, // 5: google.ads.googleads.v15.resources.AdGroupAdAssetPolicySummary.policy_topic_entries:type_name -> google.ads.googleads.v15.common.PolicyTopicEntry
	7, // 6: google.ads.googleads.v15.resources.AdGroupAdAssetPolicySummary.review_status:type_name -> google.ads.googleads.v15.enums.PolicyReviewStatusEnum.PolicyReviewStatus
	8, // 7: google.ads.googleads.v15.resources.AdGroupAdAssetPolicySummary.approval_status:type_name -> google.ads.googleads.v15.enums.PolicyApprovalStatusEnum.PolicyApprovalStatus
	8, // [8:8] is the sub-list for method output_type
	8, // [8:8] is the sub-list for method input_type
	8, // [8:8] is the sub-list for extension type_name
	8, // [8:8] is the sub-list for extension extendee
	0, // [0:8] is the sub-list for field type_name
}

func init() { file_resources_ad_group_ad_asset_view_proto_init() }
func file_resources_ad_group_ad_asset_view_proto_init() {
	if File_resources_ad_group_ad_asset_view_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_resources_ad_group_ad_asset_view_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AdGroupAdAssetView); i {
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
		file_resources_ad_group_ad_asset_view_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AdGroupAdAssetPolicySummary); i {
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
	file_resources_ad_group_ad_asset_view_proto_msgTypes[0].OneofWrappers = []interface{}{}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_resources_ad_group_ad_asset_view_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_resources_ad_group_ad_asset_view_proto_goTypes,
		DependencyIndexes: file_resources_ad_group_ad_asset_view_proto_depIdxs,
		MessageInfos:      file_resources_ad_group_ad_asset_view_proto_msgTypes,
	}.Build()
	File_resources_ad_group_ad_asset_view_proto = out.File
	file_resources_ad_group_ad_asset_view_proto_rawDesc = nil
	file_resources_ad_group_ad_asset_view_proto_goTypes = nil
	file_resources_ad_group_ad_asset_view_proto_depIdxs = nil
}
