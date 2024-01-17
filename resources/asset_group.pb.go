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
// source: google/ads/googleads/v15/resources/asset_group.proto

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

// An asset group.
// AssetGroupAsset is used to link an asset to the asset group.
// AssetGroupSignal is used to associate a signal to an asset group.
type AssetGroup struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Immutable. The resource name of the asset group.
	// Asset group resource names have the form:
	//
	// `customers/{customer_id}/assetGroups/{asset_group_id}`
	ResourceName string `protobuf:"bytes,1,opt,name=resource_name,json=resourceName,proto3" json:"resource_name,omitempty"`
	// Output only. The ID of the asset group.
	Id int64 `protobuf:"varint,9,opt,name=id,proto3" json:"id,omitempty"`
	// Immutable. The campaign with which this asset group is associated.
	// The asset which is linked to the asset group.
	Campaign string `protobuf:"bytes,2,opt,name=campaign,proto3" json:"campaign,omitempty"`
	// Required. Name of the asset group. Required. It must have a minimum length
	// of 1 and maximum length of 128. It must be unique under a campaign.
	Name string `protobuf:"bytes,3,opt,name=name,proto3" json:"name,omitempty"`
	// A list of final URLs after all cross domain redirects. In performance max,
	// by default, the urls are eligible for expansion unless opted out.
	FinalUrls []string `protobuf:"bytes,4,rep,name=final_urls,json=finalUrls,proto3" json:"final_urls,omitempty"`
	// A list of final mobile URLs after all cross domain redirects. In
	// performance max, by default, the urls are eligible for expansion
	// unless opted out.
	FinalMobileUrls []string `protobuf:"bytes,5,rep,name=final_mobile_urls,json=finalMobileUrls,proto3" json:"final_mobile_urls,omitempty"`
	// The status of the asset group.
	Status enums.AssetGroupStatusEnum_AssetGroupStatus `protobuf:"varint,6,opt,name=status,proto3,enum=google.ads.googleads.v15.enums.AssetGroupStatusEnum_AssetGroupStatus" json:"status,omitempty"`
	// Output only. The primary status of the asset group. Provides insights into
	// why an asset group is not serving or not serving optimally.
	PrimaryStatus enums.AssetGroupPrimaryStatusEnum_AssetGroupPrimaryStatus `protobuf:"varint,11,opt,name=primary_status,json=primaryStatus,proto3,enum=google.ads.googleads.v15.enums.AssetGroupPrimaryStatusEnum_AssetGroupPrimaryStatus" json:"primary_status,omitempty"`
	// Output only. Provides reasons into why an asset group is not serving or not
	// serving optimally. It will be empty when the asset group is serving without
	// issues.
	PrimaryStatusReasons []enums.AssetGroupPrimaryStatusReasonEnum_AssetGroupPrimaryStatusReason `protobuf:"varint,12,rep,packed,name=primary_status_reasons,json=primaryStatusReasons,proto3,enum=google.ads.googleads.v15.enums.AssetGroupPrimaryStatusReasonEnum_AssetGroupPrimaryStatusReason" json:"primary_status_reasons,omitempty"`
	// First part of text that may appear appended to the url displayed in
	// the ad.
	Path1 string `protobuf:"bytes,7,opt,name=path1,proto3" json:"path1,omitempty"`
	// Second part of text that may appear appended to the url displayed in
	// the ad. This field can only be set when path1 is set.
	Path2 string `protobuf:"bytes,8,opt,name=path2,proto3" json:"path2,omitempty"`
	// Output only. Overall ad strength of this asset group.
	AdStrength enums.AdStrengthEnum_AdStrength `protobuf:"varint,10,opt,name=ad_strength,json=adStrength,proto3,enum=google.ads.googleads.v15.enums.AdStrengthEnum_AdStrength" json:"ad_strength,omitempty"`
}

func (x *AssetGroup) Reset() {
	*x = AssetGroup{}
	if protoimpl.UnsafeEnabled {
		mi := &file_resources_asset_group_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AssetGroup) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AssetGroup) ProtoMessage() {}

func (x *AssetGroup) ProtoReflect() protoreflect.Message {
	mi := &file_resources_asset_group_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AssetGroup.ProtoReflect.Descriptor instead.
func (*AssetGroup) Descriptor() ([]byte, []int) {
	return file_resources_asset_group_proto_rawDescGZIP(), []int{0}
}

func (x *AssetGroup) GetResourceName() string {
	if x != nil {
		return x.ResourceName
	}
	return ""
}

func (x *AssetGroup) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *AssetGroup) GetCampaign() string {
	if x != nil {
		return x.Campaign
	}
	return ""
}

func (x *AssetGroup) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *AssetGroup) GetFinalUrls() []string {
	if x != nil {
		return x.FinalUrls
	}
	return nil
}

func (x *AssetGroup) GetFinalMobileUrls() []string {
	if x != nil {
		return x.FinalMobileUrls
	}
	return nil
}

func (x *AssetGroup) GetStatus() enums.AssetGroupStatusEnum_AssetGroupStatus {
	if x != nil {
		return x.Status
	}
	return enums.AssetGroupStatusEnum_AssetGroupStatus(0)
}

func (x *AssetGroup) GetPrimaryStatus() enums.AssetGroupPrimaryStatusEnum_AssetGroupPrimaryStatus {
	if x != nil {
		return x.PrimaryStatus
	}
	return enums.AssetGroupPrimaryStatusEnum_AssetGroupPrimaryStatus(0)
}

func (x *AssetGroup) GetPrimaryStatusReasons() []enums.AssetGroupPrimaryStatusReasonEnum_AssetGroupPrimaryStatusReason {
	if x != nil {
		return x.PrimaryStatusReasons
	}
	return nil
}

func (x *AssetGroup) GetPath1() string {
	if x != nil {
		return x.Path1
	}
	return ""
}

func (x *AssetGroup) GetPath2() string {
	if x != nil {
		return x.Path2
	}
	return ""
}

func (x *AssetGroup) GetAdStrength() enums.AdStrengthEnum_AdStrength {
	if x != nil {
		return x.AdStrength
	}
	return enums.AdStrengthEnum_AdStrength(0)
}

var File_resources_asset_group_proto protoreflect.FileDescriptor

var file_resources_asset_group_proto_rawDesc = []byte{
	0x0a, 0x34, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x64, 0x73, 0x2f, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x61, 0x64, 0x73, 0x2f, 0x76, 0x31, 0x35, 0x2f, 0x72, 0x65, 0x73, 0x6f, 0x75,
	0x72, 0x63, 0x65, 0x73, 0x2f, 0x61, 0x73, 0x73, 0x65, 0x74, 0x5f, 0x67, 0x72, 0x6f, 0x75, 0x70,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x22, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x61,
	0x64, 0x73, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61, 0x64, 0x73, 0x2e, 0x76, 0x31, 0x35,
	0x2e, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x73, 0x1a, 0x30, 0x67, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x2f, 0x61, 0x64, 0x73, 0x2f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61, 0x64, 0x73,
	0x2f, 0x76, 0x31, 0x35, 0x2f, 0x65, 0x6e, 0x75, 0x6d, 0x73, 0x2f, 0x61, 0x64, 0x5f, 0x73, 0x74,
	0x72, 0x65, 0x6e, 0x67, 0x74, 0x68, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x3f, 0x67, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x64, 0x73, 0x2f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61,
	0x64, 0x73, 0x2f, 0x76, 0x31, 0x35, 0x2f, 0x65, 0x6e, 0x75, 0x6d, 0x73, 0x2f, 0x61, 0x73, 0x73,
	0x65, 0x74, 0x5f, 0x67, 0x72, 0x6f, 0x75, 0x70, 0x5f, 0x70, 0x72, 0x69, 0x6d, 0x61, 0x72, 0x79,
	0x5f, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x46, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x64, 0x73, 0x2f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x61, 0x64, 0x73, 0x2f, 0x76, 0x31, 0x35, 0x2f, 0x65, 0x6e, 0x75, 0x6d, 0x73, 0x2f, 0x61, 0x73,
	0x73, 0x65, 0x74, 0x5f, 0x67, 0x72, 0x6f, 0x75, 0x70, 0x5f, 0x70, 0x72, 0x69, 0x6d, 0x61, 0x72,
	0x79, 0x5f, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x5f, 0x72, 0x65, 0x61, 0x73, 0x6f, 0x6e, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x37, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x64,
	0x73, 0x2f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61, 0x64, 0x73, 0x2f, 0x76, 0x31, 0x35, 0x2f,
	0x65, 0x6e, 0x75, 0x6d, 0x73, 0x2f, 0x61, 0x73, 0x73, 0x65, 0x74, 0x5f, 0x67, 0x72, 0x6f, 0x75,
	0x70, 0x5f, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1f,
	0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x66, 0x69, 0x65, 0x6c, 0x64,
	0x5f, 0x62, 0x65, 0x68, 0x61, 0x76, 0x69, 0x6f, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a,
	0x19, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x72, 0x65, 0x73, 0x6f,
	0x75, 0x72, 0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x88, 0x07, 0x0a, 0x0a, 0x41,
	0x73, 0x73, 0x65, 0x74, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x12, 0x50, 0x0a, 0x0d, 0x72, 0x65, 0x73,
	0x6f, 0x75, 0x72, 0x63, 0x65, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x42, 0x2b, 0xe0, 0x41, 0x05, 0xfa, 0x41, 0x25, 0x0a, 0x23, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x61, 0x64, 0x73, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61, 0x70, 0x69, 0x73, 0x2e, 0x63,
	0x6f, 0x6d, 0x2f, 0x41, 0x73, 0x73, 0x65, 0x74, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x52, 0x0c, 0x72,
	0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x13, 0x0a, 0x02, 0x69,
	0x64, 0x18, 0x09, 0x20, 0x01, 0x28, 0x03, 0x42, 0x03, 0xe0, 0x41, 0x03, 0x52, 0x02, 0x69, 0x64,
	0x12, 0x45, 0x0a, 0x08, 0x63, 0x61, 0x6d, 0x70, 0x61, 0x69, 0x67, 0x6e, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x09, 0x42, 0x29, 0xe0, 0x41, 0x05, 0xfa, 0x41, 0x23, 0x0a, 0x21, 0x67, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x61, 0x64, 0x73, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61, 0x70, 0x69, 0x73,
	0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x43, 0x61, 0x6d, 0x70, 0x61, 0x69, 0x67, 0x6e, 0x52, 0x08, 0x63,
	0x61, 0x6d, 0x70, 0x61, 0x69, 0x67, 0x6e, 0x12, 0x17, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18,
	0x03, 0x20, 0x01, 0x28, 0x09, 0x42, 0x03, 0xe0, 0x41, 0x02, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65,
	0x12, 0x1d, 0x0a, 0x0a, 0x66, 0x69, 0x6e, 0x61, 0x6c, 0x5f, 0x75, 0x72, 0x6c, 0x73, 0x18, 0x04,
	0x20, 0x03, 0x28, 0x09, 0x52, 0x09, 0x66, 0x69, 0x6e, 0x61, 0x6c, 0x55, 0x72, 0x6c, 0x73, 0x12,
	0x2a, 0x0a, 0x11, 0x66, 0x69, 0x6e, 0x61, 0x6c, 0x5f, 0x6d, 0x6f, 0x62, 0x69, 0x6c, 0x65, 0x5f,
	0x75, 0x72, 0x6c, 0x73, 0x18, 0x05, 0x20, 0x03, 0x28, 0x09, 0x52, 0x0f, 0x66, 0x69, 0x6e, 0x61,
	0x6c, 0x4d, 0x6f, 0x62, 0x69, 0x6c, 0x65, 0x55, 0x72, 0x6c, 0x73, 0x12, 0x5d, 0x0a, 0x06, 0x73,
	0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x06, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x45, 0x2e, 0x67, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x61, 0x64, 0x73, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61,
	0x64, 0x73, 0x2e, 0x76, 0x31, 0x35, 0x2e, 0x65, 0x6e, 0x75, 0x6d, 0x73, 0x2e, 0x41, 0x73, 0x73,
	0x65, 0x74, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x45, 0x6e, 0x75,
	0x6d, 0x2e, 0x41, 0x73, 0x73, 0x65, 0x74, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x53, 0x74, 0x61, 0x74,
	0x75, 0x73, 0x52, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x7f, 0x0a, 0x0e, 0x70, 0x72,
	0x69, 0x6d, 0x61, 0x72, 0x79, 0x5f, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x0b, 0x20, 0x01,
	0x28, 0x0e, 0x32, 0x53, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x61, 0x64, 0x73, 0x2e,
	0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61, 0x64, 0x73, 0x2e, 0x76, 0x31, 0x35, 0x2e, 0x65, 0x6e,
	0x75, 0x6d, 0x73, 0x2e, 0x41, 0x73, 0x73, 0x65, 0x74, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x50, 0x72,
	0x69, 0x6d, 0x61, 0x72, 0x79, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x45, 0x6e, 0x75, 0x6d, 0x2e,
	0x41, 0x73, 0x73, 0x65, 0x74, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x50, 0x72, 0x69, 0x6d, 0x61, 0x72,
	0x79, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x42, 0x03, 0xe0, 0x41, 0x03, 0x52, 0x0d, 0x70, 0x72,
	0x69, 0x6d, 0x61, 0x72, 0x79, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x9a, 0x01, 0x0a, 0x16,
	0x70, 0x72, 0x69, 0x6d, 0x61, 0x72, 0x79, 0x5f, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x5f, 0x72,
	0x65, 0x61, 0x73, 0x6f, 0x6e, 0x73, 0x18, 0x0c, 0x20, 0x03, 0x28, 0x0e, 0x32, 0x5f, 0x2e, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x61, 0x64, 0x73, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x61, 0x64, 0x73, 0x2e, 0x76, 0x31, 0x35, 0x2e, 0x65, 0x6e, 0x75, 0x6d, 0x73, 0x2e, 0x41, 0x73,
	0x73, 0x65, 0x74, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x50, 0x72, 0x69, 0x6d, 0x61, 0x72, 0x79, 0x53,
	0x74, 0x61, 0x74, 0x75, 0x73, 0x52, 0x65, 0x61, 0x73, 0x6f, 0x6e, 0x45, 0x6e, 0x75, 0x6d, 0x2e,
	0x41, 0x73, 0x73, 0x65, 0x74, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x50, 0x72, 0x69, 0x6d, 0x61, 0x72,
	0x79, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x52, 0x65, 0x61, 0x73, 0x6f, 0x6e, 0x42, 0x03, 0xe0,
	0x41, 0x03, 0x52, 0x14, 0x70, 0x72, 0x69, 0x6d, 0x61, 0x72, 0x79, 0x53, 0x74, 0x61, 0x74, 0x75,
	0x73, 0x52, 0x65, 0x61, 0x73, 0x6f, 0x6e, 0x73, 0x12, 0x14, 0x0a, 0x05, 0x70, 0x61, 0x74, 0x68,
	0x31, 0x18, 0x07, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x70, 0x61, 0x74, 0x68, 0x31, 0x12, 0x14,
	0x0a, 0x05, 0x70, 0x61, 0x74, 0x68, 0x32, 0x18, 0x08, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x70,
	0x61, 0x74, 0x68, 0x32, 0x12, 0x5f, 0x0a, 0x0b, 0x61, 0x64, 0x5f, 0x73, 0x74, 0x72, 0x65, 0x6e,
	0x67, 0x74, 0x68, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x39, 0x2e, 0x67, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x2e, 0x61, 0x64, 0x73, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61, 0x64, 0x73,
	0x2e, 0x76, 0x31, 0x35, 0x2e, 0x65, 0x6e, 0x75, 0x6d, 0x73, 0x2e, 0x41, 0x64, 0x53, 0x74, 0x72,
	0x65, 0x6e, 0x67, 0x74, 0x68, 0x45, 0x6e, 0x75, 0x6d, 0x2e, 0x41, 0x64, 0x53, 0x74, 0x72, 0x65,
	0x6e, 0x67, 0x74, 0x68, 0x42, 0x03, 0xe0, 0x41, 0x03, 0x52, 0x0a, 0x61, 0x64, 0x53, 0x74, 0x72,
	0x65, 0x6e, 0x67, 0x74, 0x68, 0x3a, 0x5e, 0xea, 0x41, 0x5b, 0x0a, 0x23, 0x67, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x61, 0x64, 0x73, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61, 0x70, 0x69, 0x73,
	0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x41, 0x73, 0x73, 0x65, 0x74, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x12,
	0x34, 0x63, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x65, 0x72, 0x73, 0x2f, 0x7b, 0x63, 0x75, 0x73, 0x74,
	0x6f, 0x6d, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x7d, 0x2f, 0x61, 0x73, 0x73, 0x65, 0x74, 0x47, 0x72,
	0x6f, 0x75, 0x70, 0x73, 0x2f, 0x7b, 0x61, 0x73, 0x73, 0x65, 0x74, 0x5f, 0x67, 0x72, 0x6f, 0x75,
	0x70, 0x5f, 0x69, 0x64, 0x7d, 0x42, 0x81, 0x02, 0x0a, 0x26, 0x63, 0x6f, 0x6d, 0x2e, 0x67, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x61, 0x64, 0x73, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61,
	0x64, 0x73, 0x2e, 0x76, 0x31, 0x35, 0x2e, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x73,
	0x42, 0x0f, 0x41, 0x73, 0x73, 0x65, 0x74, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x50, 0x72, 0x6f, 0x74,
	0x6f, 0x50, 0x01, 0x5a, 0x4b, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x67, 0x6f, 0x6c, 0x61,
	0x6e, 0x67, 0x2e, 0x6f, 0x72, 0x67, 0x2f, 0x67, 0x65, 0x6e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f,
	0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61, 0x70, 0x69, 0x73, 0x2f, 0x61, 0x64, 0x73, 0x2f, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61, 0x64, 0x73, 0x2f, 0x76, 0x31, 0x35, 0x2f, 0x72, 0x65, 0x73,
	0x6f, 0x75, 0x72, 0x63, 0x65, 0x73, 0x3b, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x73,
	0xa2, 0x02, 0x03, 0x47, 0x41, 0x41, 0xaa, 0x02, 0x22, 0x47, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e,
	0x41, 0x64, 0x73, 0x2e, 0x47, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x41, 0x64, 0x73, 0x2e, 0x56, 0x31,
	0x35, 0x2e, 0x52, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x73, 0xca, 0x02, 0x22, 0x47, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x5c, 0x41, 0x64, 0x73, 0x5c, 0x47, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x41,
	0x64, 0x73, 0x5c, 0x56, 0x31, 0x35, 0x5c, 0x52, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x73,
	0xea, 0x02, 0x26, 0x47, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x3a, 0x3a, 0x41, 0x64, 0x73, 0x3a, 0x3a,
	0x47, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x41, 0x64, 0x73, 0x3a, 0x3a, 0x56, 0x31, 0x35, 0x3a, 0x3a,
	0x52, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x73, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x33,
}

var (
	file_resources_asset_group_proto_rawDescOnce sync.Once
	file_resources_asset_group_proto_rawDescData = file_resources_asset_group_proto_rawDesc
)

func file_resources_asset_group_proto_rawDescGZIP() []byte {
	file_resources_asset_group_proto_rawDescOnce.Do(func() {
		file_resources_asset_group_proto_rawDescData = protoimpl.X.CompressGZIP(file_resources_asset_group_proto_rawDescData)
	})
	return file_resources_asset_group_proto_rawDescData
}

var file_resources_asset_group_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_resources_asset_group_proto_goTypes = []interface{}{
	(*AssetGroup)(nil), // 0: google.ads.googleads.v15.resources.AssetGroup
	(enums.AssetGroupStatusEnum_AssetGroupStatus)(0),                           // 1: google.ads.googleads.v15.enums.AssetGroupStatusEnum.AssetGroupStatus
	(enums.AssetGroupPrimaryStatusEnum_AssetGroupPrimaryStatus)(0),             // 2: google.ads.googleads.v15.enums.AssetGroupPrimaryStatusEnum.AssetGroupPrimaryStatus
	(enums.AssetGroupPrimaryStatusReasonEnum_AssetGroupPrimaryStatusReason)(0), // 3: google.ads.googleads.v15.enums.AssetGroupPrimaryStatusReasonEnum.AssetGroupPrimaryStatusReason
	(enums.AdStrengthEnum_AdStrength)(0),                                       // 4: google.ads.googleads.v15.enums.AdStrengthEnum.AdStrength
}
var file_resources_asset_group_proto_depIdxs = []int32{
	1, // 0: google.ads.googleads.v15.resources.AssetGroup.status:type_name -> google.ads.googleads.v15.enums.AssetGroupStatusEnum.AssetGroupStatus
	2, // 1: google.ads.googleads.v15.resources.AssetGroup.primary_status:type_name -> google.ads.googleads.v15.enums.AssetGroupPrimaryStatusEnum.AssetGroupPrimaryStatus
	3, // 2: google.ads.googleads.v15.resources.AssetGroup.primary_status_reasons:type_name -> google.ads.googleads.v15.enums.AssetGroupPrimaryStatusReasonEnum.AssetGroupPrimaryStatusReason
	4, // 3: google.ads.googleads.v15.resources.AssetGroup.ad_strength:type_name -> google.ads.googleads.v15.enums.AdStrengthEnum.AdStrength
	4, // [4:4] is the sub-list for method output_type
	4, // [4:4] is the sub-list for method input_type
	4, // [4:4] is the sub-list for extension type_name
	4, // [4:4] is the sub-list for extension extendee
	0, // [0:4] is the sub-list for field type_name
}

func init() { file_resources_asset_group_proto_init() }
func file_resources_asset_group_proto_init() {
	if File_resources_asset_group_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_resources_asset_group_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AssetGroup); i {
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
			RawDescriptor: file_resources_asset_group_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_resources_asset_group_proto_goTypes,
		DependencyIndexes: file_resources_asset_group_proto_depIdxs,
		MessageInfos:      file_resources_asset_group_proto_msgTypes,
	}.Build()
	File_resources_asset_group_proto = out.File
	file_resources_asset_group_proto_rawDesc = nil
	file_resources_asset_group_proto_goTypes = nil
	file_resources_asset_group_proto_depIdxs = nil
}
