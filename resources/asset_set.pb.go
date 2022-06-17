// Copyright 2022 Google LLC
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
// 	protoc-gen-go v1.26.0
// 	protoc        v3.19.4
// source: google/ads/googleads/v11/resources/asset_set.proto

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

// An asset set representing a collection of assets.
// Use AssetSetAsset to link an asset to the asset set.
type AssetSet struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Output only. The ID of the asset set.
	Id int64 `protobuf:"varint,6,opt,name=id,proto3" json:"id,omitempty"`
	// Immutable. The resource name of the asset set.
	// Asset set resource names have the form:
	//
	// `customers/{customer_id}/assetSets/{asset_set_id}`
	ResourceName string `protobuf:"bytes,1,opt,name=resource_name,json=resourceName,proto3" json:"resource_name,omitempty"`
	// Required. Name of the asset set. Required. It must have a minimum length of 1 and
	// maximum length of 128.
	Name string `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	// Required. Immutable. The type of the asset set. Required.
	Type enums.AssetSetTypeEnum_AssetSetType `protobuf:"varint,3,opt,name=type,proto3,enum=google.ads.googleads.v11.enums.AssetSetTypeEnum_AssetSetType" json:"type,omitempty"`
	// Output only. The status of the asset set. Read-only.
	Status enums.AssetSetStatusEnum_AssetSetStatus `protobuf:"varint,4,opt,name=status,proto3,enum=google.ads.googleads.v11.enums.AssetSetStatusEnum_AssetSetStatus" json:"status,omitempty"`
	// Merchant ID and Feed Label from Google Merchant Center.
	MerchantCenterFeed *AssetSet_MerchantCenterFeed `protobuf:"bytes,5,opt,name=merchant_center_feed,json=merchantCenterFeed,proto3" json:"merchant_center_feed,omitempty"`
}

func (x *AssetSet) Reset() {
	*x = AssetSet{}
	if protoimpl.UnsafeEnabled {
		mi := &file_resources_asset_set_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AssetSet) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AssetSet) ProtoMessage() {}

func (x *AssetSet) ProtoReflect() protoreflect.Message {
	mi := &file_resources_asset_set_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AssetSet.ProtoReflect.Descriptor instead.
func (*AssetSet) Descriptor() ([]byte, []int) {
	return file_resources_asset_set_proto_rawDescGZIP(), []int{0}
}

func (x *AssetSet) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *AssetSet) GetResourceName() string {
	if x != nil {
		return x.ResourceName
	}
	return ""
}

func (x *AssetSet) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *AssetSet) GetType() enums.AssetSetTypeEnum_AssetSetType {
	if x != nil {
		return x.Type
	}
	return enums.AssetSetTypeEnum_UNSPECIFIED
}

func (x *AssetSet) GetStatus() enums.AssetSetStatusEnum_AssetSetStatus {
	if x != nil {
		return x.Status
	}
	return enums.AssetSetStatusEnum_UNSPECIFIED
}

func (x *AssetSet) GetMerchantCenterFeed() *AssetSet_MerchantCenterFeed {
	if x != nil {
		return x.MerchantCenterFeed
	}
	return nil
}

// Merchant ID and Feed Label from Google Merchant Center.
type AssetSet_MerchantCenterFeed struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Required. Merchant ID from Google Merchant Center
	MerchantId int64 `protobuf:"varint,1,opt,name=merchant_id,json=merchantId,proto3" json:"merchant_id,omitempty"`
	// Optional. Feed Label from Google Merchant Center.
	FeedLabel *string `protobuf:"bytes,2,opt,name=feed_label,json=feedLabel,proto3,oneof" json:"feed_label,omitempty"`
}

func (x *AssetSet_MerchantCenterFeed) Reset() {
	*x = AssetSet_MerchantCenterFeed{}
	if protoimpl.UnsafeEnabled {
		mi := &file_resources_asset_set_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AssetSet_MerchantCenterFeed) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AssetSet_MerchantCenterFeed) ProtoMessage() {}

func (x *AssetSet_MerchantCenterFeed) ProtoReflect() protoreflect.Message {
	mi := &file_resources_asset_set_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AssetSet_MerchantCenterFeed.ProtoReflect.Descriptor instead.
func (*AssetSet_MerchantCenterFeed) Descriptor() ([]byte, []int) {
	return file_resources_asset_set_proto_rawDescGZIP(), []int{0, 0}
}

func (x *AssetSet_MerchantCenterFeed) GetMerchantId() int64 {
	if x != nil {
		return x.MerchantId
	}
	return 0
}

func (x *AssetSet_MerchantCenterFeed) GetFeedLabel() string {
	if x != nil && x.FeedLabel != nil {
		return *x.FeedLabel
	}
	return ""
}

var File_resources_asset_set_proto protoreflect.FileDescriptor

var file_resources_asset_set_proto_rawDesc = []byte{
	0x0a, 0x32, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x64, 0x73, 0x2f, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x61, 0x64, 0x73, 0x2f, 0x76, 0x31, 0x31, 0x2f, 0x72, 0x65, 0x73, 0x6f, 0x75,
	0x72, 0x63, 0x65, 0x73, 0x2f, 0x61, 0x73, 0x73, 0x65, 0x74, 0x5f, 0x73, 0x65, 0x74, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x12, 0x22, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x61, 0x64, 0x73,
	0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61, 0x64, 0x73, 0x2e, 0x76, 0x31, 0x31, 0x2e, 0x72,
	0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x73, 0x1a, 0x35, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x2f, 0x61, 0x64, 0x73, 0x2f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61, 0x64, 0x73, 0x2f, 0x76,
	0x31, 0x31, 0x2f, 0x65, 0x6e, 0x75, 0x6d, 0x73, 0x2f, 0x61, 0x73, 0x73, 0x65, 0x74, 0x5f, 0x73,
	0x65, 0x74, 0x5f, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a,
	0x33, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x64, 0x73, 0x2f, 0x67, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x61, 0x64, 0x73, 0x2f, 0x76, 0x31, 0x31, 0x2f, 0x65, 0x6e, 0x75, 0x6d, 0x73, 0x2f,
	0x61, 0x73, 0x73, 0x65, 0x74, 0x5f, 0x73, 0x65, 0x74, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69,
	0x2f, 0x66, 0x69, 0x65, 0x6c, 0x64, 0x5f, 0x62, 0x65, 0x68, 0x61, 0x76, 0x69, 0x6f, 0x72, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x19, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70,
	0x69, 0x2f, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x22, 0x84, 0x05, 0x0a, 0x08, 0x41, 0x73, 0x73, 0x65, 0x74, 0x53, 0x65, 0x74, 0x12, 0x13, 0x0a,
	0x02, 0x69, 0x64, 0x18, 0x06, 0x20, 0x01, 0x28, 0x03, 0x42, 0x03, 0xe0, 0x41, 0x03, 0x52, 0x02,
	0x69, 0x64, 0x12, 0x4e, 0x0a, 0x0d, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x5f, 0x6e,
	0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x42, 0x29, 0xe0, 0x41, 0x05, 0xfa, 0x41,
	0x23, 0x0a, 0x21, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61, 0x64, 0x73, 0x2e, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x61, 0x70, 0x69, 0x73, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x41, 0x73, 0x73, 0x65,
	0x74, 0x53, 0x65, 0x74, 0x52, 0x0c, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x4e, 0x61,
	0x6d, 0x65, 0x12, 0x17, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09,
	0x42, 0x03, 0xe0, 0x41, 0x02, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x59, 0x0a, 0x04, 0x74,
	0x79, 0x70, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x3d, 0x2e, 0x67, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x2e, 0x61, 0x64, 0x73, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61, 0x64, 0x73,
	0x2e, 0x76, 0x31, 0x31, 0x2e, 0x65, 0x6e, 0x75, 0x6d, 0x73, 0x2e, 0x41, 0x73, 0x73, 0x65, 0x74,
	0x53, 0x65, 0x74, 0x54, 0x79, 0x70, 0x65, 0x45, 0x6e, 0x75, 0x6d, 0x2e, 0x41, 0x73, 0x73, 0x65,
	0x74, 0x53, 0x65, 0x74, 0x54, 0x79, 0x70, 0x65, 0x42, 0x06, 0xe0, 0x41, 0x02, 0xe0, 0x41, 0x05,
	0x52, 0x04, 0x74, 0x79, 0x70, 0x65, 0x12, 0x5e, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73,
	0x18, 0x04, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x41, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e,
	0x61, 0x64, 0x73, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61, 0x64, 0x73, 0x2e, 0x76, 0x31,
	0x31, 0x2e, 0x65, 0x6e, 0x75, 0x6d, 0x73, 0x2e, 0x41, 0x73, 0x73, 0x65, 0x74, 0x53, 0x65, 0x74,
	0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x45, 0x6e, 0x75, 0x6d, 0x2e, 0x41, 0x73, 0x73, 0x65, 0x74,
	0x53, 0x65, 0x74, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x42, 0x03, 0xe0, 0x41, 0x03, 0x52, 0x06,
	0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x71, 0x0a, 0x14, 0x6d, 0x65, 0x72, 0x63, 0x68, 0x61,
	0x6e, 0x74, 0x5f, 0x63, 0x65, 0x6e, 0x74, 0x65, 0x72, 0x5f, 0x66, 0x65, 0x65, 0x64, 0x18, 0x05,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x3f, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x61, 0x64,
	0x73, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61, 0x64, 0x73, 0x2e, 0x76, 0x31, 0x31, 0x2e,
	0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x73, 0x2e, 0x41, 0x73, 0x73, 0x65, 0x74, 0x53,
	0x65, 0x74, 0x2e, 0x4d, 0x65, 0x72, 0x63, 0x68, 0x61, 0x6e, 0x74, 0x43, 0x65, 0x6e, 0x74, 0x65,
	0x72, 0x46, 0x65, 0x65, 0x64, 0x52, 0x12, 0x6d, 0x65, 0x72, 0x63, 0x68, 0x61, 0x6e, 0x74, 0x43,
	0x65, 0x6e, 0x74, 0x65, 0x72, 0x46, 0x65, 0x65, 0x64, 0x1a, 0x72, 0x0a, 0x12, 0x4d, 0x65, 0x72,
	0x63, 0x68, 0x61, 0x6e, 0x74, 0x43, 0x65, 0x6e, 0x74, 0x65, 0x72, 0x46, 0x65, 0x65, 0x64, 0x12,
	0x24, 0x0a, 0x0b, 0x6d, 0x65, 0x72, 0x63, 0x68, 0x61, 0x6e, 0x74, 0x5f, 0x69, 0x64, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x03, 0x42, 0x03, 0xe0, 0x41, 0x02, 0x52, 0x0a, 0x6d, 0x65, 0x72, 0x63, 0x68,
	0x61, 0x6e, 0x74, 0x49, 0x64, 0x12, 0x27, 0x0a, 0x0a, 0x66, 0x65, 0x65, 0x64, 0x5f, 0x6c, 0x61,
	0x62, 0x65, 0x6c, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x42, 0x03, 0xe0, 0x41, 0x01, 0x48, 0x00,
	0x52, 0x09, 0x66, 0x65, 0x65, 0x64, 0x4c, 0x61, 0x62, 0x65, 0x6c, 0x88, 0x01, 0x01, 0x42, 0x0d,
	0x0a, 0x0b, 0x5f, 0x66, 0x65, 0x65, 0x64, 0x5f, 0x6c, 0x61, 0x62, 0x65, 0x6c, 0x3a, 0x58, 0xea,
	0x41, 0x55, 0x0a, 0x21, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61, 0x64, 0x73, 0x2e, 0x67, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x61, 0x70, 0x69, 0x73, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x41, 0x73, 0x73,
	0x65, 0x74, 0x53, 0x65, 0x74, 0x12, 0x30, 0x63, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x65, 0x72, 0x73,
	0x2f, 0x7b, 0x63, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x7d, 0x2f, 0x61,
	0x73, 0x73, 0x65, 0x74, 0x53, 0x65, 0x74, 0x73, 0x2f, 0x7b, 0x61, 0x73, 0x73, 0x65, 0x74, 0x5f,
	0x73, 0x65, 0x74, 0x5f, 0x69, 0x64, 0x7d, 0x42, 0xff, 0x01, 0x0a, 0x26, 0x63, 0x6f, 0x6d, 0x2e,
	0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x61, 0x64, 0x73, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x61, 0x64, 0x73, 0x2e, 0x76, 0x31, 0x31, 0x2e, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63,
	0x65, 0x73, 0x42, 0x0d, 0x41, 0x73, 0x73, 0x65, 0x74, 0x53, 0x65, 0x74, 0x50, 0x72, 0x6f, 0x74,
	0x6f, 0x50, 0x01, 0x5a, 0x4b, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x67, 0x6f, 0x6c, 0x61,
	0x6e, 0x67, 0x2e, 0x6f, 0x72, 0x67, 0x2f, 0x67, 0x65, 0x6e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f,
	0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61, 0x70, 0x69, 0x73, 0x2f, 0x61, 0x64, 0x73, 0x2f, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61, 0x64, 0x73, 0x2f, 0x76, 0x31, 0x31, 0x2f, 0x72, 0x65, 0x73,
	0x6f, 0x75, 0x72, 0x63, 0x65, 0x73, 0x3b, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x73,
	0xa2, 0x02, 0x03, 0x47, 0x41, 0x41, 0xaa, 0x02, 0x22, 0x47, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e,
	0x41, 0x64, 0x73, 0x2e, 0x47, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x41, 0x64, 0x73, 0x2e, 0x56, 0x31,
	0x31, 0x2e, 0x52, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x73, 0xca, 0x02, 0x22, 0x47, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x5c, 0x41, 0x64, 0x73, 0x5c, 0x47, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x41,
	0x64, 0x73, 0x5c, 0x56, 0x31, 0x31, 0x5c, 0x52, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x73,
	0xea, 0x02, 0x26, 0x47, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x3a, 0x3a, 0x41, 0x64, 0x73, 0x3a, 0x3a,
	0x47, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x41, 0x64, 0x73, 0x3a, 0x3a, 0x56, 0x31, 0x31, 0x3a, 0x3a,
	0x52, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x73, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x33,
}

var (
	file_resources_asset_set_proto_rawDescOnce sync.Once
	file_resources_asset_set_proto_rawDescData = file_resources_asset_set_proto_rawDesc
)

func file_resources_asset_set_proto_rawDescGZIP() []byte {
	file_resources_asset_set_proto_rawDescOnce.Do(func() {
		file_resources_asset_set_proto_rawDescData = protoimpl.X.CompressGZIP(file_resources_asset_set_proto_rawDescData)
	})
	return file_resources_asset_set_proto_rawDescData
}

var file_resources_asset_set_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_resources_asset_set_proto_goTypes = []interface{}{
	(*AssetSet)(nil),                             // 0: google.ads.googleads.v11.resources.AssetSet
	(*AssetSet_MerchantCenterFeed)(nil),          // 1: google.ads.googleads.v11.resources.AssetSet.MerchantCenterFeed
	(enums.AssetSetTypeEnum_AssetSetType)(0),     // 2: google.ads.googleads.v11.enums.AssetSetTypeEnum.AssetSetType
	(enums.AssetSetStatusEnum_AssetSetStatus)(0), // 3: google.ads.googleads.v11.enums.AssetSetStatusEnum.AssetSetStatus
}
var file_resources_asset_set_proto_depIdxs = []int32{
	2, // 0: google.ads.googleads.v11.resources.AssetSet.type:type_name -> google.ads.googleads.v11.enums.AssetSetTypeEnum.AssetSetType
	3, // 1: google.ads.googleads.v11.resources.AssetSet.status:type_name -> google.ads.googleads.v11.enums.AssetSetStatusEnum.AssetSetStatus
	1, // 2: google.ads.googleads.v11.resources.AssetSet.merchant_center_feed:type_name -> google.ads.googleads.v11.resources.AssetSet.MerchantCenterFeed
	3, // [3:3] is the sub-list for method output_type
	3, // [3:3] is the sub-list for method input_type
	3, // [3:3] is the sub-list for extension type_name
	3, // [3:3] is the sub-list for extension extendee
	0, // [0:3] is the sub-list for field type_name
}

func init() { file_resources_asset_set_proto_init() }
func file_resources_asset_set_proto_init() {
	if File_resources_asset_set_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_resources_asset_set_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AssetSet); i {
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
		file_resources_asset_set_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AssetSet_MerchantCenterFeed); i {
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
	file_resources_asset_set_proto_msgTypes[1].OneofWrappers = []interface{}{}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_resources_asset_set_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_resources_asset_set_proto_goTypes,
		DependencyIndexes: file_resources_asset_set_proto_depIdxs,
		MessageInfos:      file_resources_asset_set_proto_msgTypes,
	}.Build()
	File_resources_asset_set_proto = out.File
	file_resources_asset_set_proto_rawDesc = nil
	file_resources_asset_set_proto_goTypes = nil
	file_resources_asset_set_proto_depIdxs = nil
}
