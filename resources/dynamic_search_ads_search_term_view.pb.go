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
// source: google/ads/googleads/v19/resources/dynamic_search_ads_search_term_view.proto

package resources

import (
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

// A dynamic search ads search term view.
type DynamicSearchAdsSearchTermView struct {
	state protoimpl.MessageState `protogen:"open.v1"`
	// Output only. The resource name of the dynamic search ads search term view.
	// Dynamic search ads search term view resource names have the form:
	//
	// `customers/{customer_id}/dynamicSearchAdsSearchTermViews/{ad_group_id}~{search_term_fingerprint}~{headline_fingerprint}~{landing_page_fingerprint}~{page_url_fingerprint}`
	ResourceName string `protobuf:"bytes,1,opt,name=resource_name,json=resourceName,proto3" json:"resource_name,omitempty"`
	// Output only. Search term
	//
	// This field is read-only.
	SearchTerm *string `protobuf:"bytes,9,opt,name=search_term,json=searchTerm,proto3,oneof" json:"search_term,omitempty"`
	// Output only. The dynamically generated headline of the Dynamic Search Ad.
	//
	// This field is read-only.
	Headline *string `protobuf:"bytes,10,opt,name=headline,proto3,oneof" json:"headline,omitempty"`
	// Output only. The dynamically selected landing page URL of the impression.
	//
	// This field is read-only.
	LandingPage *string `protobuf:"bytes,11,opt,name=landing_page,json=landingPage,proto3,oneof" json:"landing_page,omitempty"`
	// Output only. The URL of page feed item served for the impression.
	//
	// This field is read-only.
	PageUrl *string `protobuf:"bytes,12,opt,name=page_url,json=pageUrl,proto3,oneof" json:"page_url,omitempty"`
	// Output only. True if query matches a negative keyword.
	//
	// This field is read-only.
	HasNegativeKeyword *bool `protobuf:"varint,13,opt,name=has_negative_keyword,json=hasNegativeKeyword,proto3,oneof" json:"has_negative_keyword,omitempty"`
	// Output only. True if query is added to targeted keywords.
	//
	// This field is read-only.
	HasMatchingKeyword *bool `protobuf:"varint,14,opt,name=has_matching_keyword,json=hasMatchingKeyword,proto3,oneof" json:"has_matching_keyword,omitempty"`
	// Output only. True if query matches a negative url.
	//
	// This field is read-only.
	HasNegativeUrl *bool `protobuf:"varint,15,opt,name=has_negative_url,json=hasNegativeUrl,proto3,oneof" json:"has_negative_url,omitempty"`
	unknownFields  protoimpl.UnknownFields
	sizeCache      protoimpl.SizeCache
}

func (x *DynamicSearchAdsSearchTermView) Reset() {
	*x = DynamicSearchAdsSearchTermView{}
	mi := &file_resources_dynamic_search_ads_search_term_view_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *DynamicSearchAdsSearchTermView) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DynamicSearchAdsSearchTermView) ProtoMessage() {}

func (x *DynamicSearchAdsSearchTermView) ProtoReflect() protoreflect.Message {
	mi := &file_resources_dynamic_search_ads_search_term_view_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DynamicSearchAdsSearchTermView.ProtoReflect.Descriptor instead.
func (*DynamicSearchAdsSearchTermView) Descriptor() ([]byte, []int) {
	return file_resources_dynamic_search_ads_search_term_view_proto_rawDescGZIP(), []int{0}
}

func (x *DynamicSearchAdsSearchTermView) GetResourceName() string {
	if x != nil {
		return x.ResourceName
	}
	return ""
}

func (x *DynamicSearchAdsSearchTermView) GetSearchTerm() string {
	if x != nil && x.SearchTerm != nil {
		return *x.SearchTerm
	}
	return ""
}

func (x *DynamicSearchAdsSearchTermView) GetHeadline() string {
	if x != nil && x.Headline != nil {
		return *x.Headline
	}
	return ""
}

func (x *DynamicSearchAdsSearchTermView) GetLandingPage() string {
	if x != nil && x.LandingPage != nil {
		return *x.LandingPage
	}
	return ""
}

func (x *DynamicSearchAdsSearchTermView) GetPageUrl() string {
	if x != nil && x.PageUrl != nil {
		return *x.PageUrl
	}
	return ""
}

func (x *DynamicSearchAdsSearchTermView) GetHasNegativeKeyword() bool {
	if x != nil && x.HasNegativeKeyword != nil {
		return *x.HasNegativeKeyword
	}
	return false
}

func (x *DynamicSearchAdsSearchTermView) GetHasMatchingKeyword() bool {
	if x != nil && x.HasMatchingKeyword != nil {
		return *x.HasMatchingKeyword
	}
	return false
}

func (x *DynamicSearchAdsSearchTermView) GetHasNegativeUrl() bool {
	if x != nil && x.HasNegativeUrl != nil {
		return *x.HasNegativeUrl
	}
	return false
}

var File_resources_dynamic_search_ads_search_term_view_proto protoreflect.FileDescriptor

var file_resources_dynamic_search_ads_search_term_view_proto_rawDesc = string([]byte{
	0x0a, 0x4c, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x64, 0x73, 0x2f, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x61, 0x64, 0x73, 0x2f, 0x76, 0x31, 0x39, 0x2f, 0x72, 0x65, 0x73, 0x6f, 0x75,
	0x72, 0x63, 0x65, 0x73, 0x2f, 0x64, 0x79, 0x6e, 0x61, 0x6d, 0x69, 0x63, 0x5f, 0x73, 0x65, 0x61,
	0x72, 0x63, 0x68, 0x5f, 0x61, 0x64, 0x73, 0x5f, 0x73, 0x65, 0x61, 0x72, 0x63, 0x68, 0x5f, 0x74,
	0x65, 0x72, 0x6d, 0x5f, 0x76, 0x69, 0x65, 0x77, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x22,
	0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x61, 0x64, 0x73, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x61, 0x64, 0x73, 0x2e, 0x76, 0x31, 0x39, 0x2e, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63,
	0x65, 0x73, 0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x66,
	0x69, 0x65, 0x6c, 0x64, 0x5f, 0x62, 0x65, 0x68, 0x61, 0x76, 0x69, 0x6f, 0x72, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x1a, 0x19, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f,
	0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xc2,
	0x06, 0x0a, 0x1e, 0x44, 0x79, 0x6e, 0x61, 0x6d, 0x69, 0x63, 0x53, 0x65, 0x61, 0x72, 0x63, 0x68,
	0x41, 0x64, 0x73, 0x53, 0x65, 0x61, 0x72, 0x63, 0x68, 0x54, 0x65, 0x72, 0x6d, 0x56, 0x69, 0x65,
	0x77, 0x12, 0x64, 0x0a, 0x0d, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x5f, 0x6e, 0x61,
	0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x42, 0x3f, 0xe0, 0x41, 0x03, 0xfa, 0x41, 0x39,
	0x0a, 0x37, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61, 0x64, 0x73, 0x2e, 0x67, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x61, 0x70, 0x69, 0x73, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x44, 0x79, 0x6e, 0x61, 0x6d,
	0x69, 0x63, 0x53, 0x65, 0x61, 0x72, 0x63, 0x68, 0x41, 0x64, 0x73, 0x53, 0x65, 0x61, 0x72, 0x63,
	0x68, 0x54, 0x65, 0x72, 0x6d, 0x56, 0x69, 0x65, 0x77, 0x52, 0x0c, 0x72, 0x65, 0x73, 0x6f, 0x75,
	0x72, 0x63, 0x65, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x29, 0x0a, 0x0b, 0x73, 0x65, 0x61, 0x72, 0x63,
	0x68, 0x5f, 0x74, 0x65, 0x72, 0x6d, 0x18, 0x09, 0x20, 0x01, 0x28, 0x09, 0x42, 0x03, 0xe0, 0x41,
	0x03, 0x48, 0x00, 0x52, 0x0a, 0x73, 0x65, 0x61, 0x72, 0x63, 0x68, 0x54, 0x65, 0x72, 0x6d, 0x88,
	0x01, 0x01, 0x12, 0x24, 0x0a, 0x08, 0x68, 0x65, 0x61, 0x64, 0x6c, 0x69, 0x6e, 0x65, 0x18, 0x0a,
	0x20, 0x01, 0x28, 0x09, 0x42, 0x03, 0xe0, 0x41, 0x03, 0x48, 0x01, 0x52, 0x08, 0x68, 0x65, 0x61,
	0x64, 0x6c, 0x69, 0x6e, 0x65, 0x88, 0x01, 0x01, 0x12, 0x2b, 0x0a, 0x0c, 0x6c, 0x61, 0x6e, 0x64,
	0x69, 0x6e, 0x67, 0x5f, 0x70, 0x61, 0x67, 0x65, 0x18, 0x0b, 0x20, 0x01, 0x28, 0x09, 0x42, 0x03,
	0xe0, 0x41, 0x03, 0x48, 0x02, 0x52, 0x0b, 0x6c, 0x61, 0x6e, 0x64, 0x69, 0x6e, 0x67, 0x50, 0x61,
	0x67, 0x65, 0x88, 0x01, 0x01, 0x12, 0x23, 0x0a, 0x08, 0x70, 0x61, 0x67, 0x65, 0x5f, 0x75, 0x72,
	0x6c, 0x18, 0x0c, 0x20, 0x01, 0x28, 0x09, 0x42, 0x03, 0xe0, 0x41, 0x03, 0x48, 0x03, 0x52, 0x07,
	0x70, 0x61, 0x67, 0x65, 0x55, 0x72, 0x6c, 0x88, 0x01, 0x01, 0x12, 0x3a, 0x0a, 0x14, 0x68, 0x61,
	0x73, 0x5f, 0x6e, 0x65, 0x67, 0x61, 0x74, 0x69, 0x76, 0x65, 0x5f, 0x6b, 0x65, 0x79, 0x77, 0x6f,
	0x72, 0x64, 0x18, 0x0d, 0x20, 0x01, 0x28, 0x08, 0x42, 0x03, 0xe0, 0x41, 0x03, 0x48, 0x04, 0x52,
	0x12, 0x68, 0x61, 0x73, 0x4e, 0x65, 0x67, 0x61, 0x74, 0x69, 0x76, 0x65, 0x4b, 0x65, 0x79, 0x77,
	0x6f, 0x72, 0x64, 0x88, 0x01, 0x01, 0x12, 0x3a, 0x0a, 0x14, 0x68, 0x61, 0x73, 0x5f, 0x6d, 0x61,
	0x74, 0x63, 0x68, 0x69, 0x6e, 0x67, 0x5f, 0x6b, 0x65, 0x79, 0x77, 0x6f, 0x72, 0x64, 0x18, 0x0e,
	0x20, 0x01, 0x28, 0x08, 0x42, 0x03, 0xe0, 0x41, 0x03, 0x48, 0x05, 0x52, 0x12, 0x68, 0x61, 0x73,
	0x4d, 0x61, 0x74, 0x63, 0x68, 0x69, 0x6e, 0x67, 0x4b, 0x65, 0x79, 0x77, 0x6f, 0x72, 0x64, 0x88,
	0x01, 0x01, 0x12, 0x32, 0x0a, 0x10, 0x68, 0x61, 0x73, 0x5f, 0x6e, 0x65, 0x67, 0x61, 0x74, 0x69,
	0x76, 0x65, 0x5f, 0x75, 0x72, 0x6c, 0x18, 0x0f, 0x20, 0x01, 0x28, 0x08, 0x42, 0x03, 0xe0, 0x41,
	0x03, 0x48, 0x06, 0x52, 0x0e, 0x68, 0x61, 0x73, 0x4e, 0x65, 0x67, 0x61, 0x74, 0x69, 0x76, 0x65,
	0x55, 0x72, 0x6c, 0x88, 0x01, 0x01, 0x3a, 0xe8, 0x01, 0xea, 0x41, 0xe4, 0x01, 0x0a, 0x37, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61, 0x64, 0x73, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61,
	0x70, 0x69, 0x73, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x44, 0x79, 0x6e, 0x61, 0x6d, 0x69, 0x63, 0x53,
	0x65, 0x61, 0x72, 0x63, 0x68, 0x41, 0x64, 0x73, 0x53, 0x65, 0x61, 0x72, 0x63, 0x68, 0x54, 0x65,
	0x72, 0x6d, 0x56, 0x69, 0x65, 0x77, 0x12, 0xa8, 0x01, 0x63, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x65,
	0x72, 0x73, 0x2f, 0x7b, 0x63, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x7d,
	0x2f, 0x64, 0x79, 0x6e, 0x61, 0x6d, 0x69, 0x63, 0x53, 0x65, 0x61, 0x72, 0x63, 0x68, 0x41, 0x64,
	0x73, 0x53, 0x65, 0x61, 0x72, 0x63, 0x68, 0x54, 0x65, 0x72, 0x6d, 0x56, 0x69, 0x65, 0x77, 0x73,
	0x2f, 0x7b, 0x61, 0x64, 0x5f, 0x67, 0x72, 0x6f, 0x75, 0x70, 0x5f, 0x69, 0x64, 0x7d, 0x7e, 0x7b,
	0x73, 0x65, 0x61, 0x72, 0x63, 0x68, 0x5f, 0x74, 0x65, 0x72, 0x6d, 0x5f, 0x66, 0x69, 0x6e, 0x67,
	0x65, 0x72, 0x70, 0x72, 0x69, 0x6e, 0x74, 0x7d, 0x7e, 0x7b, 0x68, 0x65, 0x61, 0x64, 0x6c, 0x69,
	0x6e, 0x65, 0x5f, 0x66, 0x69, 0x6e, 0x67, 0x65, 0x72, 0x70, 0x72, 0x69, 0x6e, 0x74, 0x7d, 0x7e,
	0x7b, 0x6c, 0x61, 0x6e, 0x64, 0x69, 0x6e, 0x67, 0x5f, 0x70, 0x61, 0x67, 0x65, 0x5f, 0x66, 0x69,
	0x6e, 0x67, 0x65, 0x72, 0x70, 0x72, 0x69, 0x6e, 0x74, 0x7d, 0x7e, 0x7b, 0x70, 0x61, 0x67, 0x65,
	0x5f, 0x75, 0x72, 0x6c, 0x5f, 0x66, 0x69, 0x6e, 0x67, 0x65, 0x72, 0x70, 0x72, 0x69, 0x6e, 0x74,
	0x7d, 0x42, 0x0e, 0x0a, 0x0c, 0x5f, 0x73, 0x65, 0x61, 0x72, 0x63, 0x68, 0x5f, 0x74, 0x65, 0x72,
	0x6d, 0x42, 0x0b, 0x0a, 0x09, 0x5f, 0x68, 0x65, 0x61, 0x64, 0x6c, 0x69, 0x6e, 0x65, 0x42, 0x0f,
	0x0a, 0x0d, 0x5f, 0x6c, 0x61, 0x6e, 0x64, 0x69, 0x6e, 0x67, 0x5f, 0x70, 0x61, 0x67, 0x65, 0x42,
	0x0b, 0x0a, 0x09, 0x5f, 0x70, 0x61, 0x67, 0x65, 0x5f, 0x75, 0x72, 0x6c, 0x42, 0x17, 0x0a, 0x15,
	0x5f, 0x68, 0x61, 0x73, 0x5f, 0x6e, 0x65, 0x67, 0x61, 0x74, 0x69, 0x76, 0x65, 0x5f, 0x6b, 0x65,
	0x79, 0x77, 0x6f, 0x72, 0x64, 0x42, 0x17, 0x0a, 0x15, 0x5f, 0x68, 0x61, 0x73, 0x5f, 0x6d, 0x61,
	0x74, 0x63, 0x68, 0x69, 0x6e, 0x67, 0x5f, 0x6b, 0x65, 0x79, 0x77, 0x6f, 0x72, 0x64, 0x42, 0x13,
	0x0a, 0x11, 0x5f, 0x68, 0x61, 0x73, 0x5f, 0x6e, 0x65, 0x67, 0x61, 0x74, 0x69, 0x76, 0x65, 0x5f,
	0x75, 0x72, 0x6c, 0x42, 0x95, 0x02, 0x0a, 0x26, 0x63, 0x6f, 0x6d, 0x2e, 0x67, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x2e, 0x61, 0x64, 0x73, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61, 0x64, 0x73,
	0x2e, 0x76, 0x31, 0x39, 0x2e, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x73, 0x42, 0x23,
	0x44, 0x79, 0x6e, 0x61, 0x6d, 0x69, 0x63, 0x53, 0x65, 0x61, 0x72, 0x63, 0x68, 0x41, 0x64, 0x73,
	0x53, 0x65, 0x61, 0x72, 0x63, 0x68, 0x54, 0x65, 0x72, 0x6d, 0x56, 0x69, 0x65, 0x77, 0x50, 0x72,
	0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a, 0x4b, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x67, 0x6f,
	0x6c, 0x61, 0x6e, 0x67, 0x2e, 0x6f, 0x72, 0x67, 0x2f, 0x67, 0x65, 0x6e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x2f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61, 0x70, 0x69, 0x73, 0x2f, 0x61, 0x64, 0x73,
	0x2f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61, 0x64, 0x73, 0x2f, 0x76, 0x31, 0x39, 0x2f, 0x72,
	0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x73, 0x3b, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63,
	0x65, 0x73, 0xa2, 0x02, 0x03, 0x47, 0x41, 0x41, 0xaa, 0x02, 0x22, 0x47, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x2e, 0x41, 0x64, 0x73, 0x2e, 0x47, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x41, 0x64, 0x73, 0x2e,
	0x56, 0x31, 0x39, 0x2e, 0x52, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x73, 0xca, 0x02, 0x22,
	0x47, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x5c, 0x41, 0x64, 0x73, 0x5c, 0x47, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x41, 0x64, 0x73, 0x5c, 0x56, 0x31, 0x39, 0x5c, 0x52, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63,
	0x65, 0x73, 0xea, 0x02, 0x26, 0x47, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x3a, 0x3a, 0x41, 0x64, 0x73,
	0x3a, 0x3a, 0x47, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x41, 0x64, 0x73, 0x3a, 0x3a, 0x56, 0x31, 0x39,
	0x3a, 0x3a, 0x52, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x73, 0x62, 0x06, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x33,
})

var (
	file_resources_dynamic_search_ads_search_term_view_proto_rawDescOnce sync.Once
	file_resources_dynamic_search_ads_search_term_view_proto_rawDescData []byte
)

func file_resources_dynamic_search_ads_search_term_view_proto_rawDescGZIP() []byte {
	file_resources_dynamic_search_ads_search_term_view_proto_rawDescOnce.Do(func() {
		file_resources_dynamic_search_ads_search_term_view_proto_rawDescData = protoimpl.X.CompressGZIP(unsafe.Slice(unsafe.StringData(file_resources_dynamic_search_ads_search_term_view_proto_rawDesc), len(file_resources_dynamic_search_ads_search_term_view_proto_rawDesc)))
	})
	return file_resources_dynamic_search_ads_search_term_view_proto_rawDescData
}

var file_resources_dynamic_search_ads_search_term_view_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_resources_dynamic_search_ads_search_term_view_proto_goTypes = []any{
	(*DynamicSearchAdsSearchTermView)(nil), // 0: google.ads.googleads.v19.resources.DynamicSearchAdsSearchTermView
}
var file_resources_dynamic_search_ads_search_term_view_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_resources_dynamic_search_ads_search_term_view_proto_init() }
func file_resources_dynamic_search_ads_search_term_view_proto_init() {
	if File_resources_dynamic_search_ads_search_term_view_proto != nil {
		return
	}
	file_resources_dynamic_search_ads_search_term_view_proto_msgTypes[0].OneofWrappers = []any{}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: unsafe.Slice(unsafe.StringData(file_resources_dynamic_search_ads_search_term_view_proto_rawDesc), len(file_resources_dynamic_search_ads_search_term_view_proto_rawDesc)),
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_resources_dynamic_search_ads_search_term_view_proto_goTypes,
		DependencyIndexes: file_resources_dynamic_search_ads_search_term_view_proto_depIdxs,
		MessageInfos:      file_resources_dynamic_search_ads_search_term_view_proto_msgTypes,
	}.Build()
	File_resources_dynamic_search_ads_search_term_view_proto = out.File
	file_resources_dynamic_search_ads_search_term_view_proto_goTypes = nil
	file_resources_dynamic_search_ads_search_term_view_proto_depIdxs = nil
}
