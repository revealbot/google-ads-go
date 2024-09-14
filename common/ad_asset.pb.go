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
// source: google/ads/googleads/v17/common/ad_asset.proto

package common

import (
	enums "github.com/revealbot/google-ads-go/enums"
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

// A text asset used inside an ad.
type AdTextAsset struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Asset text.
	Text *string `protobuf:"bytes,4,opt,name=text,proto3,oneof" json:"text,omitempty"`
	// The pinned field of the asset. This restricts the asset to only serve
	// within this field. Multiple assets can be pinned to the same field. An
	// asset that is unpinned or pinned to a different field will not serve in a
	// field where some other asset has been pinned.
	PinnedField enums.ServedAssetFieldTypeEnum_ServedAssetFieldType `protobuf:"varint,2,opt,name=pinned_field,json=pinnedField,proto3,enum=google.ads.googleads.v17.enums.ServedAssetFieldTypeEnum_ServedAssetFieldType" json:"pinned_field,omitempty"`
	// The performance label of this text asset.
	AssetPerformanceLabel enums.AssetPerformanceLabelEnum_AssetPerformanceLabel `protobuf:"varint,5,opt,name=asset_performance_label,json=assetPerformanceLabel,proto3,enum=google.ads.googleads.v17.enums.AssetPerformanceLabelEnum_AssetPerformanceLabel" json:"asset_performance_label,omitempty"`
	// The policy summary of this text asset.
	PolicySummaryInfo *AdAssetPolicySummary `protobuf:"bytes,6,opt,name=policy_summary_info,json=policySummaryInfo,proto3" json:"policy_summary_info,omitempty"`
}

func (x *AdTextAsset) Reset() {
	*x = AdTextAsset{}
	if protoimpl.UnsafeEnabled {
		mi := &file_common_ad_asset_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AdTextAsset) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AdTextAsset) ProtoMessage() {}

func (x *AdTextAsset) ProtoReflect() protoreflect.Message {
	mi := &file_common_ad_asset_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AdTextAsset.ProtoReflect.Descriptor instead.
func (*AdTextAsset) Descriptor() ([]byte, []int) {
	return file_common_ad_asset_proto_rawDescGZIP(), []int{0}
}

func (x *AdTextAsset) GetText() string {
	if x != nil && x.Text != nil {
		return *x.Text
	}
	return ""
}

func (x *AdTextAsset) GetPinnedField() enums.ServedAssetFieldTypeEnum_ServedAssetFieldType {
	if x != nil {
		return x.PinnedField
	}
	return enums.ServedAssetFieldTypeEnum_ServedAssetFieldType(0)
}

func (x *AdTextAsset) GetAssetPerformanceLabel() enums.AssetPerformanceLabelEnum_AssetPerformanceLabel {
	if x != nil {
		return x.AssetPerformanceLabel
	}
	return enums.AssetPerformanceLabelEnum_AssetPerformanceLabel(0)
}

func (x *AdTextAsset) GetPolicySummaryInfo() *AdAssetPolicySummary {
	if x != nil {
		return x.PolicySummaryInfo
	}
	return nil
}

// An image asset used inside an ad.
type AdImageAsset struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The Asset resource name of this image.
	Asset *string `protobuf:"bytes,2,opt,name=asset,proto3,oneof" json:"asset,omitempty"`
}

func (x *AdImageAsset) Reset() {
	*x = AdImageAsset{}
	if protoimpl.UnsafeEnabled {
		mi := &file_common_ad_asset_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AdImageAsset) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AdImageAsset) ProtoMessage() {}

func (x *AdImageAsset) ProtoReflect() protoreflect.Message {
	mi := &file_common_ad_asset_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AdImageAsset.ProtoReflect.Descriptor instead.
func (*AdImageAsset) Descriptor() ([]byte, []int) {
	return file_common_ad_asset_proto_rawDescGZIP(), []int{1}
}

func (x *AdImageAsset) GetAsset() string {
	if x != nil && x.Asset != nil {
		return *x.Asset
	}
	return ""
}

// A video asset used inside an ad.
type AdVideoAsset struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The Asset resource name of this video.
	Asset *string `protobuf:"bytes,2,opt,name=asset,proto3,oneof" json:"asset,omitempty"`
	// Contains info fields for this AdVideoAsset.
	AdVideoAssetInfo *AdVideoAssetInfo `protobuf:"bytes,4,opt,name=ad_video_asset_info,json=adVideoAssetInfo,proto3,oneof" json:"ad_video_asset_info,omitempty"`
}

func (x *AdVideoAsset) Reset() {
	*x = AdVideoAsset{}
	if protoimpl.UnsafeEnabled {
		mi := &file_common_ad_asset_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AdVideoAsset) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AdVideoAsset) ProtoMessage() {}

func (x *AdVideoAsset) ProtoReflect() protoreflect.Message {
	mi := &file_common_ad_asset_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AdVideoAsset.ProtoReflect.Descriptor instead.
func (*AdVideoAsset) Descriptor() ([]byte, []int) {
	return file_common_ad_asset_proto_rawDescGZIP(), []int{2}
}

func (x *AdVideoAsset) GetAsset() string {
	if x != nil && x.Asset != nil {
		return *x.Asset
	}
	return ""
}

func (x *AdVideoAsset) GetAdVideoAssetInfo() *AdVideoAssetInfo {
	if x != nil {
		return x.AdVideoAssetInfo
	}
	return nil
}

// Contains info fields for AdVideoAssets.
type AdVideoAssetInfo struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// List of inventory preferences for this AdVideoAsset. This field can only be
	// set for DiscoveryVideoResponsiveAd. The video assets with an inventory
	// asset preference set will be preferred over other videos from the same ad
	// during serving time. For example, consider this ad being served for a
	// specific inventory. The server will first try to match an eligible video
	// with a matching preference for that inventory. Videos with no preferences
	// are chosen only when a video with matching preference and eligible for a
	// given ad slot can be found.
	AdVideoAssetInventoryPreferences *AdVideoAssetInventoryPreferences `protobuf:"bytes,1,opt,name=ad_video_asset_inventory_preferences,json=adVideoAssetInventoryPreferences,proto3,oneof" json:"ad_video_asset_inventory_preferences,omitempty"`
}

func (x *AdVideoAssetInfo) Reset() {
	*x = AdVideoAssetInfo{}
	if protoimpl.UnsafeEnabled {
		mi := &file_common_ad_asset_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AdVideoAssetInfo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AdVideoAssetInfo) ProtoMessage() {}

func (x *AdVideoAssetInfo) ProtoReflect() protoreflect.Message {
	mi := &file_common_ad_asset_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AdVideoAssetInfo.ProtoReflect.Descriptor instead.
func (*AdVideoAssetInfo) Descriptor() ([]byte, []int) {
	return file_common_ad_asset_proto_rawDescGZIP(), []int{3}
}

func (x *AdVideoAssetInfo) GetAdVideoAssetInventoryPreferences() *AdVideoAssetInventoryPreferences {
	if x != nil {
		return x.AdVideoAssetInventoryPreferences
	}
	return nil
}

// YouTube Video Asset inventory preferences.
type AdVideoAssetInventoryPreferences struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// When true, YouTube Video Asset with this inventory preference will be
	// preferred when choosing a video to serve In Feed.
	InFeedPreference *bool `protobuf:"varint,1,opt,name=in_feed_preference,json=inFeedPreference,proto3,oneof" json:"in_feed_preference,omitempty"`
	// When true, YouTube Video Asset with this inventory preference will be
	// preferred when choosing a video to serve In Stream.
	InStreamPreference *bool `protobuf:"varint,2,opt,name=in_stream_preference,json=inStreamPreference,proto3,oneof" json:"in_stream_preference,omitempty"`
	// When true, YouTube Video Asset with this inventory preference will be
	// preferred when choosing a video to serve on YouTube Shorts.
	ShortsPreference *bool `protobuf:"varint,3,opt,name=shorts_preference,json=shortsPreference,proto3,oneof" json:"shorts_preference,omitempty"`
}

func (x *AdVideoAssetInventoryPreferences) Reset() {
	*x = AdVideoAssetInventoryPreferences{}
	if protoimpl.UnsafeEnabled {
		mi := &file_common_ad_asset_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AdVideoAssetInventoryPreferences) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AdVideoAssetInventoryPreferences) ProtoMessage() {}

func (x *AdVideoAssetInventoryPreferences) ProtoReflect() protoreflect.Message {
	mi := &file_common_ad_asset_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AdVideoAssetInventoryPreferences.ProtoReflect.Descriptor instead.
func (*AdVideoAssetInventoryPreferences) Descriptor() ([]byte, []int) {
	return file_common_ad_asset_proto_rawDescGZIP(), []int{4}
}

func (x *AdVideoAssetInventoryPreferences) GetInFeedPreference() bool {
	if x != nil && x.InFeedPreference != nil {
		return *x.InFeedPreference
	}
	return false
}

func (x *AdVideoAssetInventoryPreferences) GetInStreamPreference() bool {
	if x != nil && x.InStreamPreference != nil {
		return *x.InStreamPreference
	}
	return false
}

func (x *AdVideoAssetInventoryPreferences) GetShortsPreference() bool {
	if x != nil && x.ShortsPreference != nil {
		return *x.ShortsPreference
	}
	return false
}

// A media bundle asset used inside an ad.
type AdMediaBundleAsset struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The Asset resource name of this media bundle.
	Asset *string `protobuf:"bytes,2,opt,name=asset,proto3,oneof" json:"asset,omitempty"`
}

func (x *AdMediaBundleAsset) Reset() {
	*x = AdMediaBundleAsset{}
	if protoimpl.UnsafeEnabled {
		mi := &file_common_ad_asset_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AdMediaBundleAsset) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AdMediaBundleAsset) ProtoMessage() {}

func (x *AdMediaBundleAsset) ProtoReflect() protoreflect.Message {
	mi := &file_common_ad_asset_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AdMediaBundleAsset.ProtoReflect.Descriptor instead.
func (*AdMediaBundleAsset) Descriptor() ([]byte, []int) {
	return file_common_ad_asset_proto_rawDescGZIP(), []int{5}
}

func (x *AdMediaBundleAsset) GetAsset() string {
	if x != nil && x.Asset != nil {
		return *x.Asset
	}
	return ""
}

// A Demand Gen carousel card asset used inside an ad.
type AdDemandGenCarouselCardAsset struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The Asset resource name of this discovery carousel card.
	Asset *string `protobuf:"bytes,1,opt,name=asset,proto3,oneof" json:"asset,omitempty"`
}

func (x *AdDemandGenCarouselCardAsset) Reset() {
	*x = AdDemandGenCarouselCardAsset{}
	if protoimpl.UnsafeEnabled {
		mi := &file_common_ad_asset_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AdDemandGenCarouselCardAsset) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AdDemandGenCarouselCardAsset) ProtoMessage() {}

func (x *AdDemandGenCarouselCardAsset) ProtoReflect() protoreflect.Message {
	mi := &file_common_ad_asset_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AdDemandGenCarouselCardAsset.ProtoReflect.Descriptor instead.
func (*AdDemandGenCarouselCardAsset) Descriptor() ([]byte, []int) {
	return file_common_ad_asset_proto_rawDescGZIP(), []int{6}
}

func (x *AdDemandGenCarouselCardAsset) GetAsset() string {
	if x != nil && x.Asset != nil {
		return *x.Asset
	}
	return ""
}

// A call to action asset used inside an ad.
type AdCallToActionAsset struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The Asset resource name of this call to action asset.
	Asset *string `protobuf:"bytes,1,opt,name=asset,proto3,oneof" json:"asset,omitempty"`
}

func (x *AdCallToActionAsset) Reset() {
	*x = AdCallToActionAsset{}
	if protoimpl.UnsafeEnabled {
		mi := &file_common_ad_asset_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AdCallToActionAsset) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AdCallToActionAsset) ProtoMessage() {}

func (x *AdCallToActionAsset) ProtoReflect() protoreflect.Message {
	mi := &file_common_ad_asset_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AdCallToActionAsset.ProtoReflect.Descriptor instead.
func (*AdCallToActionAsset) Descriptor() ([]byte, []int) {
	return file_common_ad_asset_proto_rawDescGZIP(), []int{7}
}

func (x *AdCallToActionAsset) GetAsset() string {
	if x != nil && x.Asset != nil {
		return *x.Asset
	}
	return ""
}

var File_common_ad_asset_proto protoreflect.FileDescriptor

var file_common_ad_asset_proto_rawDesc = []byte{
	0x0a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x64, 0x73, 0x2f, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x61, 0x64, 0x73, 0x2f, 0x76, 0x31, 0x37, 0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x6f,
	0x6e, 0x2f, 0x61, 0x64, 0x5f, 0x61, 0x73, 0x73, 0x65, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x12, 0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x61, 0x64, 0x73, 0x2e, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x61, 0x64, 0x73, 0x2e, 0x76, 0x31, 0x37, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f,
	0x6e, 0x1a, 0x32, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x64, 0x73, 0x2f, 0x67, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x61, 0x64, 0x73, 0x2f, 0x76, 0x31, 0x37, 0x2f, 0x63, 0x6f, 0x6d, 0x6d,
	0x6f, 0x6e, 0x2f, 0x61, 0x73, 0x73, 0x65, 0x74, 0x5f, 0x70, 0x6f, 0x6c, 0x69, 0x63, 0x79, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x3c, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x64,
	0x73, 0x2f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61, 0x64, 0x73, 0x2f, 0x76, 0x31, 0x37, 0x2f,
	0x65, 0x6e, 0x75, 0x6d, 0x73, 0x2f, 0x61, 0x73, 0x73, 0x65, 0x74, 0x5f, 0x70, 0x65, 0x72, 0x66,
	0x6f, 0x72, 0x6d, 0x61, 0x6e, 0x63, 0x65, 0x5f, 0x6c, 0x61, 0x62, 0x65, 0x6c, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x1a, 0x3c, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x64, 0x73, 0x2f,
	0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61, 0x64, 0x73, 0x2f, 0x76, 0x31, 0x37, 0x2f, 0x65, 0x6e,
	0x75, 0x6d, 0x73, 0x2f, 0x73, 0x65, 0x72, 0x76, 0x65, 0x64, 0x5f, 0x61, 0x73, 0x73, 0x65, 0x74,
	0x5f, 0x66, 0x69, 0x65, 0x6c, 0x64, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x22, 0x92, 0x03, 0x0a, 0x0b, 0x41, 0x64, 0x54, 0x65, 0x78, 0x74, 0x41, 0x73, 0x73, 0x65,
	0x74, 0x12, 0x17, 0x0a, 0x04, 0x74, 0x65, 0x78, 0x74, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x48,
	0x00, 0x52, 0x04, 0x74, 0x65, 0x78, 0x74, 0x88, 0x01, 0x01, 0x12, 0x70, 0x0a, 0x0c, 0x70, 0x69,
	0x6e, 0x6e, 0x65, 0x64, 0x5f, 0x66, 0x69, 0x65, 0x6c, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0e,
	0x32, 0x4d, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x61, 0x64, 0x73, 0x2e, 0x67, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x61, 0x64, 0x73, 0x2e, 0x76, 0x31, 0x37, 0x2e, 0x65, 0x6e, 0x75, 0x6d,
	0x73, 0x2e, 0x53, 0x65, 0x72, 0x76, 0x65, 0x64, 0x41, 0x73, 0x73, 0x65, 0x74, 0x46, 0x69, 0x65,
	0x6c, 0x64, 0x54, 0x79, 0x70, 0x65, 0x45, 0x6e, 0x75, 0x6d, 0x2e, 0x53, 0x65, 0x72, 0x76, 0x65,
	0x64, 0x41, 0x73, 0x73, 0x65, 0x74, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x54, 0x79, 0x70, 0x65, 0x52,
	0x0b, 0x70, 0x69, 0x6e, 0x6e, 0x65, 0x64, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x12, 0x87, 0x01, 0x0a,
	0x17, 0x61, 0x73, 0x73, 0x65, 0x74, 0x5f, 0x70, 0x65, 0x72, 0x66, 0x6f, 0x72, 0x6d, 0x61, 0x6e,
	0x63, 0x65, 0x5f, 0x6c, 0x61, 0x62, 0x65, 0x6c, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x4f,
	0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x61, 0x64, 0x73, 0x2e, 0x67, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x61, 0x64, 0x73, 0x2e, 0x76, 0x31, 0x37, 0x2e, 0x65, 0x6e, 0x75, 0x6d, 0x73, 0x2e,
	0x41, 0x73, 0x73, 0x65, 0x74, 0x50, 0x65, 0x72, 0x66, 0x6f, 0x72, 0x6d, 0x61, 0x6e, 0x63, 0x65,
	0x4c, 0x61, 0x62, 0x65, 0x6c, 0x45, 0x6e, 0x75, 0x6d, 0x2e, 0x41, 0x73, 0x73, 0x65, 0x74, 0x50,
	0x65, 0x72, 0x66, 0x6f, 0x72, 0x6d, 0x61, 0x6e, 0x63, 0x65, 0x4c, 0x61, 0x62, 0x65, 0x6c, 0x52,
	0x15, 0x61, 0x73, 0x73, 0x65, 0x74, 0x50, 0x65, 0x72, 0x66, 0x6f, 0x72, 0x6d, 0x61, 0x6e, 0x63,
	0x65, 0x4c, 0x61, 0x62, 0x65, 0x6c, 0x12, 0x65, 0x0a, 0x13, 0x70, 0x6f, 0x6c, 0x69, 0x63, 0x79,
	0x5f, 0x73, 0x75, 0x6d, 0x6d, 0x61, 0x72, 0x79, 0x5f, 0x69, 0x6e, 0x66, 0x6f, 0x18, 0x06, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x35, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x61, 0x64, 0x73,
	0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61, 0x64, 0x73, 0x2e, 0x76, 0x31, 0x37, 0x2e, 0x63,
	0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x41, 0x64, 0x41, 0x73, 0x73, 0x65, 0x74, 0x50, 0x6f, 0x6c,
	0x69, 0x63, 0x79, 0x53, 0x75, 0x6d, 0x6d, 0x61, 0x72, 0x79, 0x52, 0x11, 0x70, 0x6f, 0x6c, 0x69,
	0x63, 0x79, 0x53, 0x75, 0x6d, 0x6d, 0x61, 0x72, 0x79, 0x49, 0x6e, 0x66, 0x6f, 0x42, 0x07, 0x0a,
	0x05, 0x5f, 0x74, 0x65, 0x78, 0x74, 0x22, 0x33, 0x0a, 0x0c, 0x41, 0x64, 0x49, 0x6d, 0x61, 0x67,
	0x65, 0x41, 0x73, 0x73, 0x65, 0x74, 0x12, 0x19, 0x0a, 0x05, 0x61, 0x73, 0x73, 0x65, 0x74, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x09, 0x48, 0x00, 0x52, 0x05, 0x61, 0x73, 0x73, 0x65, 0x74, 0x88, 0x01,
	0x01, 0x42, 0x08, 0x0a, 0x06, 0x5f, 0x61, 0x73, 0x73, 0x65, 0x74, 0x22, 0xb2, 0x01, 0x0a, 0x0c,
	0x41, 0x64, 0x56, 0x69, 0x64, 0x65, 0x6f, 0x41, 0x73, 0x73, 0x65, 0x74, 0x12, 0x19, 0x0a, 0x05,
	0x61, 0x73, 0x73, 0x65, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x48, 0x00, 0x52, 0x05, 0x61,
	0x73, 0x73, 0x65, 0x74, 0x88, 0x01, 0x01, 0x12, 0x65, 0x0a, 0x13, 0x61, 0x64, 0x5f, 0x76, 0x69,
	0x64, 0x65, 0x6f, 0x5f, 0x61, 0x73, 0x73, 0x65, 0x74, 0x5f, 0x69, 0x6e, 0x66, 0x6f, 0x18, 0x04,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x31, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x61, 0x64,
	0x73, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61, 0x64, 0x73, 0x2e, 0x76, 0x31, 0x37, 0x2e,
	0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x41, 0x64, 0x56, 0x69, 0x64, 0x65, 0x6f, 0x41, 0x73,
	0x73, 0x65, 0x74, 0x49, 0x6e, 0x66, 0x6f, 0x48, 0x01, 0x52, 0x10, 0x61, 0x64, 0x56, 0x69, 0x64,
	0x65, 0x6f, 0x41, 0x73, 0x73, 0x65, 0x74, 0x49, 0x6e, 0x66, 0x6f, 0x88, 0x01, 0x01, 0x42, 0x08,
	0x0a, 0x06, 0x5f, 0x61, 0x73, 0x73, 0x65, 0x74, 0x42, 0x16, 0x0a, 0x14, 0x5f, 0x61, 0x64, 0x5f,
	0x76, 0x69, 0x64, 0x65, 0x6f, 0x5f, 0x61, 0x73, 0x73, 0x65, 0x74, 0x5f, 0x69, 0x6e, 0x66, 0x6f,
	0x22, 0xd4, 0x01, 0x0a, 0x10, 0x41, 0x64, 0x56, 0x69, 0x64, 0x65, 0x6f, 0x41, 0x73, 0x73, 0x65,
	0x74, 0x49, 0x6e, 0x66, 0x6f, 0x12, 0x96, 0x01, 0x0a, 0x24, 0x61, 0x64, 0x5f, 0x76, 0x69, 0x64,
	0x65, 0x6f, 0x5f, 0x61, 0x73, 0x73, 0x65, 0x74, 0x5f, 0x69, 0x6e, 0x76, 0x65, 0x6e, 0x74, 0x6f,
	0x72, 0x79, 0x5f, 0x70, 0x72, 0x65, 0x66, 0x65, 0x72, 0x65, 0x6e, 0x63, 0x65, 0x73, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x41, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x61, 0x64,
	0x73, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61, 0x64, 0x73, 0x2e, 0x76, 0x31, 0x37, 0x2e,
	0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x41, 0x64, 0x56, 0x69, 0x64, 0x65, 0x6f, 0x41, 0x73,
	0x73, 0x65, 0x74, 0x49, 0x6e, 0x76, 0x65, 0x6e, 0x74, 0x6f, 0x72, 0x79, 0x50, 0x72, 0x65, 0x66,
	0x65, 0x72, 0x65, 0x6e, 0x63, 0x65, 0x73, 0x48, 0x00, 0x52, 0x20, 0x61, 0x64, 0x56, 0x69, 0x64,
	0x65, 0x6f, 0x41, 0x73, 0x73, 0x65, 0x74, 0x49, 0x6e, 0x76, 0x65, 0x6e, 0x74, 0x6f, 0x72, 0x79,
	0x50, 0x72, 0x65, 0x66, 0x65, 0x72, 0x65, 0x6e, 0x63, 0x65, 0x73, 0x88, 0x01, 0x01, 0x42, 0x27,
	0x0a, 0x25, 0x5f, 0x61, 0x64, 0x5f, 0x76, 0x69, 0x64, 0x65, 0x6f, 0x5f, 0x61, 0x73, 0x73, 0x65,
	0x74, 0x5f, 0x69, 0x6e, 0x76, 0x65, 0x6e, 0x74, 0x6f, 0x72, 0x79, 0x5f, 0x70, 0x72, 0x65, 0x66,
	0x65, 0x72, 0x65, 0x6e, 0x63, 0x65, 0x73, 0x22, 0x84, 0x02, 0x0a, 0x20, 0x41, 0x64, 0x56, 0x69,
	0x64, 0x65, 0x6f, 0x41, 0x73, 0x73, 0x65, 0x74, 0x49, 0x6e, 0x76, 0x65, 0x6e, 0x74, 0x6f, 0x72,
	0x79, 0x50, 0x72, 0x65, 0x66, 0x65, 0x72, 0x65, 0x6e, 0x63, 0x65, 0x73, 0x12, 0x31, 0x0a, 0x12,
	0x69, 0x6e, 0x5f, 0x66, 0x65, 0x65, 0x64, 0x5f, 0x70, 0x72, 0x65, 0x66, 0x65, 0x72, 0x65, 0x6e,
	0x63, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x08, 0x48, 0x00, 0x52, 0x10, 0x69, 0x6e, 0x46, 0x65,
	0x65, 0x64, 0x50, 0x72, 0x65, 0x66, 0x65, 0x72, 0x65, 0x6e, 0x63, 0x65, 0x88, 0x01, 0x01, 0x12,
	0x35, 0x0a, 0x14, 0x69, 0x6e, 0x5f, 0x73, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x5f, 0x70, 0x72, 0x65,
	0x66, 0x65, 0x72, 0x65, 0x6e, 0x63, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x08, 0x48, 0x01, 0x52,
	0x12, 0x69, 0x6e, 0x53, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x50, 0x72, 0x65, 0x66, 0x65, 0x72, 0x65,
	0x6e, 0x63, 0x65, 0x88, 0x01, 0x01, 0x12, 0x30, 0x0a, 0x11, 0x73, 0x68, 0x6f, 0x72, 0x74, 0x73,
	0x5f, 0x70, 0x72, 0x65, 0x66, 0x65, 0x72, 0x65, 0x6e, 0x63, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28,
	0x08, 0x48, 0x02, 0x52, 0x10, 0x73, 0x68, 0x6f, 0x72, 0x74, 0x73, 0x50, 0x72, 0x65, 0x66, 0x65,
	0x72, 0x65, 0x6e, 0x63, 0x65, 0x88, 0x01, 0x01, 0x42, 0x15, 0x0a, 0x13, 0x5f, 0x69, 0x6e, 0x5f,
	0x66, 0x65, 0x65, 0x64, 0x5f, 0x70, 0x72, 0x65, 0x66, 0x65, 0x72, 0x65, 0x6e, 0x63, 0x65, 0x42,
	0x17, 0x0a, 0x15, 0x5f, 0x69, 0x6e, 0x5f, 0x73, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x5f, 0x70, 0x72,
	0x65, 0x66, 0x65, 0x72, 0x65, 0x6e, 0x63, 0x65, 0x42, 0x14, 0x0a, 0x12, 0x5f, 0x73, 0x68, 0x6f,
	0x72, 0x74, 0x73, 0x5f, 0x70, 0x72, 0x65, 0x66, 0x65, 0x72, 0x65, 0x6e, 0x63, 0x65, 0x22, 0x39,
	0x0a, 0x12, 0x41, 0x64, 0x4d, 0x65, 0x64, 0x69, 0x61, 0x42, 0x75, 0x6e, 0x64, 0x6c, 0x65, 0x41,
	0x73, 0x73, 0x65, 0x74, 0x12, 0x19, 0x0a, 0x05, 0x61, 0x73, 0x73, 0x65, 0x74, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x09, 0x48, 0x00, 0x52, 0x05, 0x61, 0x73, 0x73, 0x65, 0x74, 0x88, 0x01, 0x01, 0x42,
	0x08, 0x0a, 0x06, 0x5f, 0x61, 0x73, 0x73, 0x65, 0x74, 0x22, 0x43, 0x0a, 0x1c, 0x41, 0x64, 0x44,
	0x65, 0x6d, 0x61, 0x6e, 0x64, 0x47, 0x65, 0x6e, 0x43, 0x61, 0x72, 0x6f, 0x75, 0x73, 0x65, 0x6c,
	0x43, 0x61, 0x72, 0x64, 0x41, 0x73, 0x73, 0x65, 0x74, 0x12, 0x19, 0x0a, 0x05, 0x61, 0x73, 0x73,
	0x65, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x48, 0x00, 0x52, 0x05, 0x61, 0x73, 0x73, 0x65,
	0x74, 0x88, 0x01, 0x01, 0x42, 0x08, 0x0a, 0x06, 0x5f, 0x61, 0x73, 0x73, 0x65, 0x74, 0x22, 0x3a,
	0x0a, 0x13, 0x41, 0x64, 0x43, 0x61, 0x6c, 0x6c, 0x54, 0x6f, 0x41, 0x63, 0x74, 0x69, 0x6f, 0x6e,
	0x41, 0x73, 0x73, 0x65, 0x74, 0x12, 0x19, 0x0a, 0x05, 0x61, 0x73, 0x73, 0x65, 0x74, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x48, 0x00, 0x52, 0x05, 0x61, 0x73, 0x73, 0x65, 0x74, 0x88, 0x01, 0x01,
	0x42, 0x08, 0x0a, 0x06, 0x5f, 0x61, 0x73, 0x73, 0x65, 0x74, 0x42, 0xec, 0x01, 0x0a, 0x23, 0x63,
	0x6f, 0x6d, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x61, 0x64, 0x73, 0x2e, 0x67, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x61, 0x64, 0x73, 0x2e, 0x76, 0x31, 0x37, 0x2e, 0x63, 0x6f, 0x6d, 0x6d,
	0x6f, 0x6e, 0x42, 0x0c, 0x41, 0x64, 0x41, 0x73, 0x73, 0x65, 0x74, 0x50, 0x72, 0x6f, 0x74, 0x6f,
	0x50, 0x01, 0x5a, 0x45, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x67, 0x6f, 0x6c, 0x61, 0x6e,
	0x67, 0x2e, 0x6f, 0x72, 0x67, 0x2f, 0x67, 0x65, 0x6e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61, 0x70, 0x69, 0x73, 0x2f, 0x61, 0x64, 0x73, 0x2f, 0x67, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x61, 0x64, 0x73, 0x2f, 0x76, 0x31, 0x37, 0x2f, 0x63, 0x6f, 0x6d, 0x6d,
	0x6f, 0x6e, 0x3b, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0xa2, 0x02, 0x03, 0x47, 0x41, 0x41, 0xaa,
	0x02, 0x1f, 0x47, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x41, 0x64, 0x73, 0x2e, 0x47, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x41, 0x64, 0x73, 0x2e, 0x56, 0x31, 0x37, 0x2e, 0x43, 0x6f, 0x6d, 0x6d, 0x6f,
	0x6e, 0xca, 0x02, 0x1f, 0x47, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x5c, 0x41, 0x64, 0x73, 0x5c, 0x47,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x41, 0x64, 0x73, 0x5c, 0x56, 0x31, 0x37, 0x5c, 0x43, 0x6f, 0x6d,
	0x6d, 0x6f, 0x6e, 0xea, 0x02, 0x23, 0x47, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x3a, 0x3a, 0x41, 0x64,
	0x73, 0x3a, 0x3a, 0x47, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x41, 0x64, 0x73, 0x3a, 0x3a, 0x56, 0x31,
	0x37, 0x3a, 0x3a, 0x43, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x33,
}

var (
	file_common_ad_asset_proto_rawDescOnce sync.Once
	file_common_ad_asset_proto_rawDescData = file_common_ad_asset_proto_rawDesc
)

func file_common_ad_asset_proto_rawDescGZIP() []byte {
	file_common_ad_asset_proto_rawDescOnce.Do(func() {
		file_common_ad_asset_proto_rawDescData = protoimpl.X.CompressGZIP(file_common_ad_asset_proto_rawDescData)
	})
	return file_common_ad_asset_proto_rawDescData
}

var file_common_ad_asset_proto_msgTypes = make([]protoimpl.MessageInfo, 8)
var file_common_ad_asset_proto_goTypes = []interface{}{
	(*AdTextAsset)(nil),                                        // 0: google.ads.googleads.v17.common.AdTextAsset
	(*AdImageAsset)(nil),                                       // 1: google.ads.googleads.v17.common.AdImageAsset
	(*AdVideoAsset)(nil),                                       // 2: google.ads.googleads.v17.common.AdVideoAsset
	(*AdVideoAssetInfo)(nil),                                   // 3: google.ads.googleads.v17.common.AdVideoAssetInfo
	(*AdVideoAssetInventoryPreferences)(nil),                   // 4: google.ads.googleads.v17.common.AdVideoAssetInventoryPreferences
	(*AdMediaBundleAsset)(nil),                                 // 5: google.ads.googleads.v17.common.AdMediaBundleAsset
	(*AdDemandGenCarouselCardAsset)(nil),                       // 6: google.ads.googleads.v17.common.AdDemandGenCarouselCardAsset
	(*AdCallToActionAsset)(nil),                                // 7: google.ads.googleads.v17.common.AdCallToActionAsset
	(enums.ServedAssetFieldTypeEnum_ServedAssetFieldType)(0),   // 8: google.ads.googleads.v17.enums.ServedAssetFieldTypeEnum.ServedAssetFieldType
	(enums.AssetPerformanceLabelEnum_AssetPerformanceLabel)(0), // 9: google.ads.googleads.v17.enums.AssetPerformanceLabelEnum.AssetPerformanceLabel
	(*AdAssetPolicySummary)(nil),                               // 10: google.ads.googleads.v17.common.AdAssetPolicySummary
}
var file_common_ad_asset_proto_depIdxs = []int32{
	8,  // 0: google.ads.googleads.v17.common.AdTextAsset.pinned_field:type_name -> google.ads.googleads.v17.enums.ServedAssetFieldTypeEnum.ServedAssetFieldType
	9,  // 1: google.ads.googleads.v17.common.AdTextAsset.asset_performance_label:type_name -> google.ads.googleads.v17.enums.AssetPerformanceLabelEnum.AssetPerformanceLabel
	10, // 2: google.ads.googleads.v17.common.AdTextAsset.policy_summary_info:type_name -> google.ads.googleads.v17.common.AdAssetPolicySummary
	3,  // 3: google.ads.googleads.v17.common.AdVideoAsset.ad_video_asset_info:type_name -> google.ads.googleads.v17.common.AdVideoAssetInfo
	4,  // 4: google.ads.googleads.v17.common.AdVideoAssetInfo.ad_video_asset_inventory_preferences:type_name -> google.ads.googleads.v17.common.AdVideoAssetInventoryPreferences
	5,  // [5:5] is the sub-list for method output_type
	5,  // [5:5] is the sub-list for method input_type
	5,  // [5:5] is the sub-list for extension type_name
	5,  // [5:5] is the sub-list for extension extendee
	0,  // [0:5] is the sub-list for field type_name
}

func init() { file_common_ad_asset_proto_init() }
func file_common_ad_asset_proto_init() {
	if File_common_ad_asset_proto != nil {
		return
	}
	file_common_asset_policy_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_common_ad_asset_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AdTextAsset); i {
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
		file_common_ad_asset_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AdImageAsset); i {
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
		file_common_ad_asset_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AdVideoAsset); i {
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
		file_common_ad_asset_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AdVideoAssetInfo); i {
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
		file_common_ad_asset_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AdVideoAssetInventoryPreferences); i {
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
		file_common_ad_asset_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AdMediaBundleAsset); i {
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
		file_common_ad_asset_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AdDemandGenCarouselCardAsset); i {
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
		file_common_ad_asset_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AdCallToActionAsset); i {
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
	file_common_ad_asset_proto_msgTypes[0].OneofWrappers = []interface{}{}
	file_common_ad_asset_proto_msgTypes[1].OneofWrappers = []interface{}{}
	file_common_ad_asset_proto_msgTypes[2].OneofWrappers = []interface{}{}
	file_common_ad_asset_proto_msgTypes[3].OneofWrappers = []interface{}{}
	file_common_ad_asset_proto_msgTypes[4].OneofWrappers = []interface{}{}
	file_common_ad_asset_proto_msgTypes[5].OneofWrappers = []interface{}{}
	file_common_ad_asset_proto_msgTypes[6].OneofWrappers = []interface{}{}
	file_common_ad_asset_proto_msgTypes[7].OneofWrappers = []interface{}{}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_common_ad_asset_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   8,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_common_ad_asset_proto_goTypes,
		DependencyIndexes: file_common_ad_asset_proto_depIdxs,
		MessageInfos:      file_common_ad_asset_proto_msgTypes,
	}.Build()
	File_common_ad_asset_proto = out.File
	file_common_ad_asset_proto_rawDesc = nil
	file_common_ad_asset_proto_goTypes = nil
	file_common_ad_asset_proto_depIdxs = nil
}
