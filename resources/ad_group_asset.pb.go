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
// source: google/ads/googleads/v14/resources/ad_group_asset.proto

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

// A link between an ad group and an asset.
type AdGroupAsset struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Immutable. The resource name of the ad group asset.
	// AdGroupAsset resource names have the form:
	//
	// `customers/{customer_id}/adGroupAssets/{ad_group_id}~{asset_id}~{field_type}`
	ResourceName string `protobuf:"bytes,1,opt,name=resource_name,json=resourceName,proto3" json:"resource_name,omitempty"`
	// Required. Immutable. The ad group to which the asset is linked.
	AdGroup string `protobuf:"bytes,2,opt,name=ad_group,json=adGroup,proto3" json:"ad_group,omitempty"`
	// Required. Immutable. The asset which is linked to the ad group.
	Asset string `protobuf:"bytes,3,opt,name=asset,proto3" json:"asset,omitempty"`
	// Required. Immutable. Role that the asset takes under the linked ad group.
	FieldType enums.AssetFieldTypeEnum_AssetFieldType `protobuf:"varint,4,opt,name=field_type,json=fieldType,proto3,enum=google.ads.googleads.v14.enums.AssetFieldTypeEnum_AssetFieldType" json:"field_type,omitempty"`
	// Output only. Source of the adgroup asset link.
	Source enums.AssetSourceEnum_AssetSource `protobuf:"varint,6,opt,name=source,proto3,enum=google.ads.googleads.v14.enums.AssetSourceEnum_AssetSource" json:"source,omitempty"`
	// Status of the ad group asset.
	Status enums.AssetLinkStatusEnum_AssetLinkStatus `protobuf:"varint,5,opt,name=status,proto3,enum=google.ads.googleads.v14.enums.AssetLinkStatusEnum_AssetLinkStatus" json:"status,omitempty"`
	// Output only. Provides the PrimaryStatus of this asset link.
	// Primary status is meant essentially to differentiate between the plain
	// "status" field, which has advertiser set values of enabled, paused, or
	// removed.  The primary status takes into account other signals (for assets
	// its mainly policy and quality approvals) to come up with a more
	// comprehensive status to indicate its serving state.
	PrimaryStatus enums.AssetLinkPrimaryStatusEnum_AssetLinkPrimaryStatus `protobuf:"varint,7,opt,name=primary_status,json=primaryStatus,proto3,enum=google.ads.googleads.v14.enums.AssetLinkPrimaryStatusEnum_AssetLinkPrimaryStatus" json:"primary_status,omitempty"`
	// Output only. Provides the details of the primary status and its associated
	// reasons.
	PrimaryStatusDetails []*common.AssetLinkPrimaryStatusDetails `protobuf:"bytes,8,rep,name=primary_status_details,json=primaryStatusDetails,proto3" json:"primary_status_details,omitempty"`
	// Output only. Provides a list of reasons for why an asset is not serving or
	// not serving at full capacity.
	PrimaryStatusReasons []enums.AssetLinkPrimaryStatusReasonEnum_AssetLinkPrimaryStatusReason `protobuf:"varint,9,rep,packed,name=primary_status_reasons,json=primaryStatusReasons,proto3,enum=google.ads.googleads.v14.enums.AssetLinkPrimaryStatusReasonEnum_AssetLinkPrimaryStatusReason" json:"primary_status_reasons,omitempty"`
}

func (x *AdGroupAsset) Reset() {
	*x = AdGroupAsset{}
	if protoimpl.UnsafeEnabled {
		mi := &file_resources_ad_group_asset_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AdGroupAsset) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AdGroupAsset) ProtoMessage() {}

func (x *AdGroupAsset) ProtoReflect() protoreflect.Message {
	mi := &file_resources_ad_group_asset_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AdGroupAsset.ProtoReflect.Descriptor instead.
func (*AdGroupAsset) Descriptor() ([]byte, []int) {
	return file_resources_ad_group_asset_proto_rawDescGZIP(), []int{0}
}

func (x *AdGroupAsset) GetResourceName() string {
	if x != nil {
		return x.ResourceName
	}
	return ""
}

func (x *AdGroupAsset) GetAdGroup() string {
	if x != nil {
		return x.AdGroup
	}
	return ""
}

func (x *AdGroupAsset) GetAsset() string {
	if x != nil {
		return x.Asset
	}
	return ""
}

func (x *AdGroupAsset) GetFieldType() enums.AssetFieldTypeEnum_AssetFieldType {
	if x != nil {
		return x.FieldType
	}
	return enums.AssetFieldTypeEnum_AssetFieldType(0)
}

func (x *AdGroupAsset) GetSource() enums.AssetSourceEnum_AssetSource {
	if x != nil {
		return x.Source
	}
	return enums.AssetSourceEnum_AssetSource(0)
}

func (x *AdGroupAsset) GetStatus() enums.AssetLinkStatusEnum_AssetLinkStatus {
	if x != nil {
		return x.Status
	}
	return enums.AssetLinkStatusEnum_AssetLinkStatus(0)
}

func (x *AdGroupAsset) GetPrimaryStatus() enums.AssetLinkPrimaryStatusEnum_AssetLinkPrimaryStatus {
	if x != nil {
		return x.PrimaryStatus
	}
	return enums.AssetLinkPrimaryStatusEnum_AssetLinkPrimaryStatus(0)
}

func (x *AdGroupAsset) GetPrimaryStatusDetails() []*common.AssetLinkPrimaryStatusDetails {
	if x != nil {
		return x.PrimaryStatusDetails
	}
	return nil
}

func (x *AdGroupAsset) GetPrimaryStatusReasons() []enums.AssetLinkPrimaryStatusReasonEnum_AssetLinkPrimaryStatusReason {
	if x != nil {
		return x.PrimaryStatusReasons
	}
	return nil
}

var File_resources_ad_group_asset_proto protoreflect.FileDescriptor

var file_resources_ad_group_asset_proto_rawDesc = []byte{
	0x0a, 0x37, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x64, 0x73, 0x2f, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x61, 0x64, 0x73, 0x2f, 0x76, 0x31, 0x34, 0x2f, 0x72, 0x65, 0x73, 0x6f, 0x75,
	0x72, 0x63, 0x65, 0x73, 0x2f, 0x61, 0x64, 0x5f, 0x67, 0x72, 0x6f, 0x75, 0x70, 0x5f, 0x61, 0x73,
	0x73, 0x65, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x22, 0x67, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x2e, 0x61, 0x64, 0x73, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61, 0x64, 0x73, 0x2e,
	0x76, 0x31, 0x34, 0x2e, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x73, 0x1a, 0x32, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x64, 0x73, 0x2f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x61, 0x64, 0x73, 0x2f, 0x76, 0x31, 0x34, 0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2f, 0x61,
	0x73, 0x73, 0x65, 0x74, 0x5f, 0x70, 0x6f, 0x6c, 0x69, 0x63, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x1a, 0x35, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x64, 0x73, 0x2f, 0x67, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x61, 0x64, 0x73, 0x2f, 0x76, 0x31, 0x34, 0x2f, 0x65, 0x6e, 0x75, 0x6d,
	0x73, 0x2f, 0x61, 0x73, 0x73, 0x65, 0x74, 0x5f, 0x66, 0x69, 0x65, 0x6c, 0x64, 0x5f, 0x74, 0x79,
	0x70, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x3e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x2f, 0x61, 0x64, 0x73, 0x2f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61, 0x64, 0x73, 0x2f, 0x76,
	0x31, 0x34, 0x2f, 0x65, 0x6e, 0x75, 0x6d, 0x73, 0x2f, 0x61, 0x73, 0x73, 0x65, 0x74, 0x5f, 0x6c,
	0x69, 0x6e, 0x6b, 0x5f, 0x70, 0x72, 0x69, 0x6d, 0x61, 0x72, 0x79, 0x5f, 0x73, 0x74, 0x61, 0x74,
	0x75, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x45, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x2f, 0x61, 0x64, 0x73, 0x2f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61, 0x64, 0x73, 0x2f, 0x76,
	0x31, 0x34, 0x2f, 0x65, 0x6e, 0x75, 0x6d, 0x73, 0x2f, 0x61, 0x73, 0x73, 0x65, 0x74, 0x5f, 0x6c,
	0x69, 0x6e, 0x6b, 0x5f, 0x70, 0x72, 0x69, 0x6d, 0x61, 0x72, 0x79, 0x5f, 0x73, 0x74, 0x61, 0x74,
	0x75, 0x73, 0x5f, 0x72, 0x65, 0x61, 0x73, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a,
	0x36, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x64, 0x73, 0x2f, 0x67, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x61, 0x64, 0x73, 0x2f, 0x76, 0x31, 0x34, 0x2f, 0x65, 0x6e, 0x75, 0x6d, 0x73, 0x2f,
	0x61, 0x73, 0x73, 0x65, 0x74, 0x5f, 0x6c, 0x69, 0x6e, 0x6b, 0x5f, 0x73, 0x74, 0x61, 0x74, 0x75,
	0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x31, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f,
	0x61, 0x64, 0x73, 0x2f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61, 0x64, 0x73, 0x2f, 0x76, 0x31,
	0x34, 0x2f, 0x65, 0x6e, 0x75, 0x6d, 0x73, 0x2f, 0x61, 0x73, 0x73, 0x65, 0x74, 0x5f, 0x73, 0x6f,
	0x75, 0x72, 0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x66, 0x69, 0x65, 0x6c, 0x64, 0x5f, 0x62, 0x65, 0x68,
	0x61, 0x76, 0x69, 0x6f, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x19, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x9a, 0x08, 0x0a, 0x0c, 0x41, 0x64, 0x47, 0x72, 0x6f,
	0x75, 0x70, 0x41, 0x73, 0x73, 0x65, 0x74, 0x12, 0x52, 0x0a, 0x0d, 0x72, 0x65, 0x73, 0x6f, 0x75,
	0x72, 0x63, 0x65, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x42, 0x2d,
	0xe0, 0x41, 0x05, 0xfa, 0x41, 0x27, 0x0a, 0x25, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61, 0x64,
	0x73, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61, 0x70, 0x69, 0x73, 0x2e, 0x63, 0x6f, 0x6d,
	0x2f, 0x41, 0x64, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x41, 0x73, 0x73, 0x65, 0x74, 0x52, 0x0c, 0x72,
	0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x46, 0x0a, 0x08, 0x61,
	0x64, 0x5f, 0x67, 0x72, 0x6f, 0x75, 0x70, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x42, 0x2b, 0xe0,
	0x41, 0x02, 0xe0, 0x41, 0x05, 0xfa, 0x41, 0x22, 0x0a, 0x20, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x61, 0x64, 0x73, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61, 0x70, 0x69, 0x73, 0x2e, 0x63,
	0x6f, 0x6d, 0x2f, 0x41, 0x64, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x52, 0x07, 0x61, 0x64, 0x47, 0x72,
	0x6f, 0x75, 0x70, 0x12, 0x3f, 0x0a, 0x05, 0x61, 0x73, 0x73, 0x65, 0x74, 0x18, 0x03, 0x20, 0x01,
	0x28, 0x09, 0x42, 0x29, 0xe0, 0x41, 0x02, 0xe0, 0x41, 0x05, 0xfa, 0x41, 0x20, 0x0a, 0x1e, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61, 0x64, 0x73, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61,
	0x70, 0x69, 0x73, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x41, 0x73, 0x73, 0x65, 0x74, 0x52, 0x05, 0x61,
	0x73, 0x73, 0x65, 0x74, 0x12, 0x68, 0x0a, 0x0a, 0x66, 0x69, 0x65, 0x6c, 0x64, 0x5f, 0x74, 0x79,
	0x70, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x41, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x2e, 0x61, 0x64, 0x73, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61, 0x64, 0x73, 0x2e,
	0x76, 0x31, 0x34, 0x2e, 0x65, 0x6e, 0x75, 0x6d, 0x73, 0x2e, 0x41, 0x73, 0x73, 0x65, 0x74, 0x46,
	0x69, 0x65, 0x6c, 0x64, 0x54, 0x79, 0x70, 0x65, 0x45, 0x6e, 0x75, 0x6d, 0x2e, 0x41, 0x73, 0x73,
	0x65, 0x74, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x54, 0x79, 0x70, 0x65, 0x42, 0x06, 0xe0, 0x41, 0x02,
	0xe0, 0x41, 0x05, 0x52, 0x09, 0x66, 0x69, 0x65, 0x6c, 0x64, 0x54, 0x79, 0x70, 0x65, 0x12, 0x58,
	0x0a, 0x06, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x18, 0x06, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x3b,
	0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x61, 0x64, 0x73, 0x2e, 0x67, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x61, 0x64, 0x73, 0x2e, 0x76, 0x31, 0x34, 0x2e, 0x65, 0x6e, 0x75, 0x6d, 0x73, 0x2e,
	0x41, 0x73, 0x73, 0x65, 0x74, 0x53, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x45, 0x6e, 0x75, 0x6d, 0x2e,
	0x41, 0x73, 0x73, 0x65, 0x74, 0x53, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x42, 0x03, 0xe0, 0x41, 0x03,
	0x52, 0x06, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x12, 0x5b, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x74,
	0x75, 0x73, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x43, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x2e, 0x61, 0x64, 0x73, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61, 0x64, 0x73, 0x2e,
	0x76, 0x31, 0x34, 0x2e, 0x65, 0x6e, 0x75, 0x6d, 0x73, 0x2e, 0x41, 0x73, 0x73, 0x65, 0x74, 0x4c,
	0x69, 0x6e, 0x6b, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x45, 0x6e, 0x75, 0x6d, 0x2e, 0x41, 0x73,
	0x73, 0x65, 0x74, 0x4c, 0x69, 0x6e, 0x6b, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x52, 0x06, 0x73,
	0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x7d, 0x0a, 0x0e, 0x70, 0x72, 0x69, 0x6d, 0x61, 0x72, 0x79,
	0x5f, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x07, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x51, 0x2e,
	0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x61, 0x64, 0x73, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x61, 0x64, 0x73, 0x2e, 0x76, 0x31, 0x34, 0x2e, 0x65, 0x6e, 0x75, 0x6d, 0x73, 0x2e, 0x41,
	0x73, 0x73, 0x65, 0x74, 0x4c, 0x69, 0x6e, 0x6b, 0x50, 0x72, 0x69, 0x6d, 0x61, 0x72, 0x79, 0x53,
	0x74, 0x61, 0x74, 0x75, 0x73, 0x45, 0x6e, 0x75, 0x6d, 0x2e, 0x41, 0x73, 0x73, 0x65, 0x74, 0x4c,
	0x69, 0x6e, 0x6b, 0x50, 0x72, 0x69, 0x6d, 0x61, 0x72, 0x79, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73,
	0x42, 0x03, 0xe0, 0x41, 0x03, 0x52, 0x0d, 0x70, 0x72, 0x69, 0x6d, 0x61, 0x72, 0x79, 0x53, 0x74,
	0x61, 0x74, 0x75, 0x73, 0x12, 0x79, 0x0a, 0x16, 0x70, 0x72, 0x69, 0x6d, 0x61, 0x72, 0x79, 0x5f,
	0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x5f, 0x64, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x73, 0x18, 0x08,
	0x20, 0x03, 0x28, 0x0b, 0x32, 0x3e, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x61, 0x64,
	0x73, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61, 0x64, 0x73, 0x2e, 0x76, 0x31, 0x34, 0x2e,
	0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x41, 0x73, 0x73, 0x65, 0x74, 0x4c, 0x69, 0x6e, 0x6b,
	0x50, 0x72, 0x69, 0x6d, 0x61, 0x72, 0x79, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x44, 0x65, 0x74,
	0x61, 0x69, 0x6c, 0x73, 0x42, 0x03, 0xe0, 0x41, 0x03, 0x52, 0x14, 0x70, 0x72, 0x69, 0x6d, 0x61,
	0x72, 0x79, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x44, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x73, 0x12,
	0x98, 0x01, 0x0a, 0x16, 0x70, 0x72, 0x69, 0x6d, 0x61, 0x72, 0x79, 0x5f, 0x73, 0x74, 0x61, 0x74,
	0x75, 0x73, 0x5f, 0x72, 0x65, 0x61, 0x73, 0x6f, 0x6e, 0x73, 0x18, 0x09, 0x20, 0x03, 0x28, 0x0e,
	0x32, 0x5d, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x61, 0x64, 0x73, 0x2e, 0x67, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x61, 0x64, 0x73, 0x2e, 0x76, 0x31, 0x34, 0x2e, 0x65, 0x6e, 0x75, 0x6d,
	0x73, 0x2e, 0x41, 0x73, 0x73, 0x65, 0x74, 0x4c, 0x69, 0x6e, 0x6b, 0x50, 0x72, 0x69, 0x6d, 0x61,
	0x72, 0x79, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x52, 0x65, 0x61, 0x73, 0x6f, 0x6e, 0x45, 0x6e,
	0x75, 0x6d, 0x2e, 0x41, 0x73, 0x73, 0x65, 0x74, 0x4c, 0x69, 0x6e, 0x6b, 0x50, 0x72, 0x69, 0x6d,
	0x61, 0x72, 0x79, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x52, 0x65, 0x61, 0x73, 0x6f, 0x6e, 0x42,
	0x03, 0xe0, 0x41, 0x03, 0x52, 0x14, 0x70, 0x72, 0x69, 0x6d, 0x61, 0x72, 0x79, 0x53, 0x74, 0x61,
	0x74, 0x75, 0x73, 0x52, 0x65, 0x61, 0x73, 0x6f, 0x6e, 0x73, 0x3a, 0x77, 0xea, 0x41, 0x74, 0x0a,
	0x25, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61, 0x64, 0x73, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x61, 0x70, 0x69, 0x73, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x41, 0x64, 0x47, 0x72, 0x6f, 0x75,
	0x70, 0x41, 0x73, 0x73, 0x65, 0x74, 0x12, 0x4b, 0x63, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x65, 0x72,
	0x73, 0x2f, 0x7b, 0x63, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x7d, 0x2f,
	0x61, 0x64, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x41, 0x73, 0x73, 0x65, 0x74, 0x73, 0x2f, 0x7b, 0x61,
	0x64, 0x5f, 0x67, 0x72, 0x6f, 0x75, 0x70, 0x5f, 0x69, 0x64, 0x7d, 0x7e, 0x7b, 0x61, 0x73, 0x73,
	0x65, 0x74, 0x5f, 0x69, 0x64, 0x7d, 0x7e, 0x7b, 0x66, 0x69, 0x65, 0x6c, 0x64, 0x5f, 0x74, 0x79,
	0x70, 0x65, 0x7d, 0x42, 0x83, 0x02, 0x0a, 0x26, 0x63, 0x6f, 0x6d, 0x2e, 0x67, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x2e, 0x61, 0x64, 0x73, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61, 0x64, 0x73,
	0x2e, 0x76, 0x31, 0x34, 0x2e, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x73, 0x42, 0x11,
	0x41, 0x64, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x41, 0x73, 0x73, 0x65, 0x74, 0x50, 0x72, 0x6f, 0x74,
	0x6f, 0x50, 0x01, 0x5a, 0x4b, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x67, 0x6f, 0x6c, 0x61,
	0x6e, 0x67, 0x2e, 0x6f, 0x72, 0x67, 0x2f, 0x67, 0x65, 0x6e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f,
	0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61, 0x70, 0x69, 0x73, 0x2f, 0x61, 0x64, 0x73, 0x2f, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61, 0x64, 0x73, 0x2f, 0x76, 0x31, 0x34, 0x2f, 0x72, 0x65, 0x73,
	0x6f, 0x75, 0x72, 0x63, 0x65, 0x73, 0x3b, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x73,
	0xa2, 0x02, 0x03, 0x47, 0x41, 0x41, 0xaa, 0x02, 0x22, 0x47, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e,
	0x41, 0x64, 0x73, 0x2e, 0x47, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x41, 0x64, 0x73, 0x2e, 0x56, 0x31,
	0x34, 0x2e, 0x52, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x73, 0xca, 0x02, 0x22, 0x47, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x5c, 0x41, 0x64, 0x73, 0x5c, 0x47, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x41,
	0x64, 0x73, 0x5c, 0x56, 0x31, 0x34, 0x5c, 0x52, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x73,
	0xea, 0x02, 0x26, 0x47, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x3a, 0x3a, 0x41, 0x64, 0x73, 0x3a, 0x3a,
	0x47, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x41, 0x64, 0x73, 0x3a, 0x3a, 0x56, 0x31, 0x34, 0x3a, 0x3a,
	0x52, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x73, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x33,
}

var (
	file_resources_ad_group_asset_proto_rawDescOnce sync.Once
	file_resources_ad_group_asset_proto_rawDescData = file_resources_ad_group_asset_proto_rawDesc
)

func file_resources_ad_group_asset_proto_rawDescGZIP() []byte {
	file_resources_ad_group_asset_proto_rawDescOnce.Do(func() {
		file_resources_ad_group_asset_proto_rawDescData = protoimpl.X.CompressGZIP(file_resources_ad_group_asset_proto_rawDescData)
	})
	return file_resources_ad_group_asset_proto_rawDescData
}

var file_resources_ad_group_asset_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_resources_ad_group_asset_proto_goTypes = []interface{}{
	(*AdGroupAsset)(nil),                                                     // 0: google.ads.googleads.v14.resources.AdGroupAsset
	(enums.AssetFieldTypeEnum_AssetFieldType)(0),                             // 1: google.ads.googleads.v14.enums.AssetFieldTypeEnum.AssetFieldType
	(enums.AssetSourceEnum_AssetSource)(0),                                   // 2: google.ads.googleads.v14.enums.AssetSourceEnum.AssetSource
	(enums.AssetLinkStatusEnum_AssetLinkStatus)(0),                           // 3: google.ads.googleads.v14.enums.AssetLinkStatusEnum.AssetLinkStatus
	(enums.AssetLinkPrimaryStatusEnum_AssetLinkPrimaryStatus)(0),             // 4: google.ads.googleads.v14.enums.AssetLinkPrimaryStatusEnum.AssetLinkPrimaryStatus
	(*common.AssetLinkPrimaryStatusDetails)(nil),                             // 5: google.ads.googleads.v14.common.AssetLinkPrimaryStatusDetails
	(enums.AssetLinkPrimaryStatusReasonEnum_AssetLinkPrimaryStatusReason)(0), // 6: google.ads.googleads.v14.enums.AssetLinkPrimaryStatusReasonEnum.AssetLinkPrimaryStatusReason
}
var file_resources_ad_group_asset_proto_depIdxs = []int32{
	1, // 0: google.ads.googleads.v14.resources.AdGroupAsset.field_type:type_name -> google.ads.googleads.v14.enums.AssetFieldTypeEnum.AssetFieldType
	2, // 1: google.ads.googleads.v14.resources.AdGroupAsset.source:type_name -> google.ads.googleads.v14.enums.AssetSourceEnum.AssetSource
	3, // 2: google.ads.googleads.v14.resources.AdGroupAsset.status:type_name -> google.ads.googleads.v14.enums.AssetLinkStatusEnum.AssetLinkStatus
	4, // 3: google.ads.googleads.v14.resources.AdGroupAsset.primary_status:type_name -> google.ads.googleads.v14.enums.AssetLinkPrimaryStatusEnum.AssetLinkPrimaryStatus
	5, // 4: google.ads.googleads.v14.resources.AdGroupAsset.primary_status_details:type_name -> google.ads.googleads.v14.common.AssetLinkPrimaryStatusDetails
	6, // 5: google.ads.googleads.v14.resources.AdGroupAsset.primary_status_reasons:type_name -> google.ads.googleads.v14.enums.AssetLinkPrimaryStatusReasonEnum.AssetLinkPrimaryStatusReason
	6, // [6:6] is the sub-list for method output_type
	6, // [6:6] is the sub-list for method input_type
	6, // [6:6] is the sub-list for extension type_name
	6, // [6:6] is the sub-list for extension extendee
	0, // [0:6] is the sub-list for field type_name
}

func init() { file_resources_ad_group_asset_proto_init() }
func file_resources_ad_group_asset_proto_init() {
	if File_resources_ad_group_asset_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_resources_ad_group_asset_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AdGroupAsset); i {
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
			RawDescriptor: file_resources_ad_group_asset_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_resources_ad_group_asset_proto_goTypes,
		DependencyIndexes: file_resources_ad_group_asset_proto_depIdxs,
		MessageInfos:      file_resources_ad_group_asset_proto_msgTypes,
	}.Build()
	File_resources_ad_group_asset_proto = out.File
	file_resources_ad_group_asset_proto_rawDesc = nil
	file_resources_ad_group_asset_proto_goTypes = nil
	file_resources_ad_group_asset_proto_depIdxs = nil
}
