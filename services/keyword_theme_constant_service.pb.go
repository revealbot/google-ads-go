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
// source: google/ads/googleads/v16/services/keyword_theme_constant_service.proto

package services

import (
	resources "github.com/revealbot/google-ads-go/resources"
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

// Request message for
// [KeywordThemeConstantService.SuggestKeywordThemeConstants][google.ads.googleads.v16.services.KeywordThemeConstantService.SuggestKeywordThemeConstants].
type SuggestKeywordThemeConstantsRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The query text of a keyword theme that will be used to map to similar
	// keyword themes. For example, "plumber" or "roofer".
	QueryText string `protobuf:"bytes,1,opt,name=query_text,json=queryText,proto3" json:"query_text,omitempty"`
	// Upper-case, two-letter country code as defined by ISO-3166. This for
	// refining the scope of the query, default to 'US' if not set.
	CountryCode string `protobuf:"bytes,2,opt,name=country_code,json=countryCode,proto3" json:"country_code,omitempty"`
	// The two letter language code for get corresponding keyword theme for
	// refining the scope of the query, default to 'en' if not set.
	LanguageCode string `protobuf:"bytes,3,opt,name=language_code,json=languageCode,proto3" json:"language_code,omitempty"`
}

func (x *SuggestKeywordThemeConstantsRequest) Reset() {
	*x = SuggestKeywordThemeConstantsRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_services_keyword_theme_constant_service_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SuggestKeywordThemeConstantsRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SuggestKeywordThemeConstantsRequest) ProtoMessage() {}

func (x *SuggestKeywordThemeConstantsRequest) ProtoReflect() protoreflect.Message {
	mi := &file_services_keyword_theme_constant_service_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SuggestKeywordThemeConstantsRequest.ProtoReflect.Descriptor instead.
func (*SuggestKeywordThemeConstantsRequest) Descriptor() ([]byte, []int) {
	return file_services_keyword_theme_constant_service_proto_rawDescGZIP(), []int{0}
}

func (x *SuggestKeywordThemeConstantsRequest) GetQueryText() string {
	if x != nil {
		return x.QueryText
	}
	return ""
}

func (x *SuggestKeywordThemeConstantsRequest) GetCountryCode() string {
	if x != nil {
		return x.CountryCode
	}
	return ""
}

func (x *SuggestKeywordThemeConstantsRequest) GetLanguageCode() string {
	if x != nil {
		return x.LanguageCode
	}
	return ""
}

// Response message for
// [KeywordThemeConstantService.SuggestKeywordThemeConstants][google.ads.googleads.v16.services.KeywordThemeConstantService.SuggestKeywordThemeConstants].
type SuggestKeywordThemeConstantsResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Smart Campaign keyword theme suggestions.
	KeywordThemeConstants []*resources.KeywordThemeConstant `protobuf:"bytes,1,rep,name=keyword_theme_constants,json=keywordThemeConstants,proto3" json:"keyword_theme_constants,omitempty"`
}

func (x *SuggestKeywordThemeConstantsResponse) Reset() {
	*x = SuggestKeywordThemeConstantsResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_services_keyword_theme_constant_service_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SuggestKeywordThemeConstantsResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SuggestKeywordThemeConstantsResponse) ProtoMessage() {}

func (x *SuggestKeywordThemeConstantsResponse) ProtoReflect() protoreflect.Message {
	mi := &file_services_keyword_theme_constant_service_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SuggestKeywordThemeConstantsResponse.ProtoReflect.Descriptor instead.
func (*SuggestKeywordThemeConstantsResponse) Descriptor() ([]byte, []int) {
	return file_services_keyword_theme_constant_service_proto_rawDescGZIP(), []int{1}
}

func (x *SuggestKeywordThemeConstantsResponse) GetKeywordThemeConstants() []*resources.KeywordThemeConstant {
	if x != nil {
		return x.KeywordThemeConstants
	}
	return nil
}

var File_services_keyword_theme_constant_service_proto protoreflect.FileDescriptor

var file_services_keyword_theme_constant_service_proto_rawDesc = []byte{
	0x0a, 0x46, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x64, 0x73, 0x2f, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x61, 0x64, 0x73, 0x2f, 0x76, 0x31, 0x36, 0x2f, 0x73, 0x65, 0x72, 0x76, 0x69,
	0x63, 0x65, 0x73, 0x2f, 0x6b, 0x65, 0x79, 0x77, 0x6f, 0x72, 0x64, 0x5f, 0x74, 0x68, 0x65, 0x6d,
	0x65, 0x5f, 0x63, 0x6f, 0x6e, 0x73, 0x74, 0x61, 0x6e, 0x74, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69,
	0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x21, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x2e, 0x61, 0x64, 0x73, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61, 0x64, 0x73, 0x2e, 0x76,
	0x31, 0x36, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x1a, 0x3f, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x2f, 0x61, 0x64, 0x73, 0x2f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61, 0x64,
	0x73, 0x2f, 0x76, 0x31, 0x36, 0x2f, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x73, 0x2f,
	0x6b, 0x65, 0x79, 0x77, 0x6f, 0x72, 0x64, 0x5f, 0x74, 0x68, 0x65, 0x6d, 0x65, 0x5f, 0x63, 0x6f,
	0x6e, 0x73, 0x74, 0x61, 0x6e, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1c, 0x67, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x17, 0x67, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x22, 0x8c, 0x01, 0x0a, 0x23, 0x53, 0x75, 0x67, 0x67, 0x65, 0x73, 0x74, 0x4b,
	0x65, 0x79, 0x77, 0x6f, 0x72, 0x64, 0x54, 0x68, 0x65, 0x6d, 0x65, 0x43, 0x6f, 0x6e, 0x73, 0x74,
	0x61, 0x6e, 0x74, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1d, 0x0a, 0x0a, 0x71,
	0x75, 0x65, 0x72, 0x79, 0x5f, 0x74, 0x65, 0x78, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x09, 0x71, 0x75, 0x65, 0x72, 0x79, 0x54, 0x65, 0x78, 0x74, 0x12, 0x21, 0x0a, 0x0c, 0x63, 0x6f,
	0x75, 0x6e, 0x74, 0x72, 0x79, 0x5f, 0x63, 0x6f, 0x64, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x0b, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x72, 0x79, 0x43, 0x6f, 0x64, 0x65, 0x12, 0x23, 0x0a,
	0x0d, 0x6c, 0x61, 0x6e, 0x67, 0x75, 0x61, 0x67, 0x65, 0x5f, 0x63, 0x6f, 0x64, 0x65, 0x18, 0x03,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x6c, 0x61, 0x6e, 0x67, 0x75, 0x61, 0x67, 0x65, 0x43, 0x6f,
	0x64, 0x65, 0x22, 0x98, 0x01, 0x0a, 0x24, 0x53, 0x75, 0x67, 0x67, 0x65, 0x73, 0x74, 0x4b, 0x65,
	0x79, 0x77, 0x6f, 0x72, 0x64, 0x54, 0x68, 0x65, 0x6d, 0x65, 0x43, 0x6f, 0x6e, 0x73, 0x74, 0x61,
	0x6e, 0x74, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x70, 0x0a, 0x17, 0x6b,
	0x65, 0x79, 0x77, 0x6f, 0x72, 0x64, 0x5f, 0x74, 0x68, 0x65, 0x6d, 0x65, 0x5f, 0x63, 0x6f, 0x6e,
	0x73, 0x74, 0x61, 0x6e, 0x74, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x38, 0x2e, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x61, 0x64, 0x73, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x61, 0x64, 0x73, 0x2e, 0x76, 0x31, 0x36, 0x2e, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65,
	0x73, 0x2e, 0x4b, 0x65, 0x79, 0x77, 0x6f, 0x72, 0x64, 0x54, 0x68, 0x65, 0x6d, 0x65, 0x43, 0x6f,
	0x6e, 0x73, 0x74, 0x61, 0x6e, 0x74, 0x52, 0x15, 0x6b, 0x65, 0x79, 0x77, 0x6f, 0x72, 0x64, 0x54,
	0x68, 0x65, 0x6d, 0x65, 0x43, 0x6f, 0x6e, 0x73, 0x74, 0x61, 0x6e, 0x74, 0x73, 0x32, 0xc5, 0x02,
	0x0a, 0x1b, 0x4b, 0x65, 0x79, 0x77, 0x6f, 0x72, 0x64, 0x54, 0x68, 0x65, 0x6d, 0x65, 0x43, 0x6f,
	0x6e, 0x73, 0x74, 0x61, 0x6e, 0x74, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0xde, 0x01,
	0x0a, 0x1c, 0x53, 0x75, 0x67, 0x67, 0x65, 0x73, 0x74, 0x4b, 0x65, 0x79, 0x77, 0x6f, 0x72, 0x64,
	0x54, 0x68, 0x65, 0x6d, 0x65, 0x43, 0x6f, 0x6e, 0x73, 0x74, 0x61, 0x6e, 0x74, 0x73, 0x12, 0x46,
	0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x61, 0x64, 0x73, 0x2e, 0x67, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x61, 0x64, 0x73, 0x2e, 0x76, 0x31, 0x36, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63,
	0x65, 0x73, 0x2e, 0x53, 0x75, 0x67, 0x67, 0x65, 0x73, 0x74, 0x4b, 0x65, 0x79, 0x77, 0x6f, 0x72,
	0x64, 0x54, 0x68, 0x65, 0x6d, 0x65, 0x43, 0x6f, 0x6e, 0x73, 0x74, 0x61, 0x6e, 0x74, 0x73, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x47, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e,
	0x61, 0x64, 0x73, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61, 0x64, 0x73, 0x2e, 0x76, 0x31,
	0x36, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x53, 0x75, 0x67, 0x67, 0x65,
	0x73, 0x74, 0x4b, 0x65, 0x79, 0x77, 0x6f, 0x72, 0x64, 0x54, 0x68, 0x65, 0x6d, 0x65, 0x43, 0x6f,
	0x6e, 0x73, 0x74, 0x61, 0x6e, 0x74, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22,
	0x2d, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x27, 0x22, 0x22, 0x2f, 0x76, 0x31, 0x36, 0x2f, 0x6b, 0x65,
	0x79, 0x77, 0x6f, 0x72, 0x64, 0x54, 0x68, 0x65, 0x6d, 0x65, 0x43, 0x6f, 0x6e, 0x73, 0x74, 0x61,
	0x6e, 0x74, 0x73, 0x3a, 0x73, 0x75, 0x67, 0x67, 0x65, 0x73, 0x74, 0x3a, 0x01, 0x2a, 0x1a, 0x45,
	0xca, 0x41, 0x18, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61, 0x64, 0x73, 0x2e, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x61, 0x70, 0x69, 0x73, 0x2e, 0x63, 0x6f, 0x6d, 0xd2, 0x41, 0x27, 0x68, 0x74,
	0x74, 0x70, 0x73, 0x3a, 0x2f, 0x2f, 0x77, 0x77, 0x77, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x61, 0x70, 0x69, 0x73, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x61, 0x75, 0x74, 0x68, 0x2f, 0x61, 0x64,
	0x77, 0x6f, 0x72, 0x64, 0x73, 0x42, 0x8c, 0x02, 0x0a, 0x25, 0x63, 0x6f, 0x6d, 0x2e, 0x67, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x61, 0x64, 0x73, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61,
	0x64, 0x73, 0x2e, 0x76, 0x31, 0x36, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x42,
	0x20, 0x4b, 0x65, 0x79, 0x77, 0x6f, 0x72, 0x64, 0x54, 0x68, 0x65, 0x6d, 0x65, 0x43, 0x6f, 0x6e,
	0x73, 0x74, 0x61, 0x6e, 0x74, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x50, 0x72, 0x6f, 0x74,
	0x6f, 0x50, 0x01, 0x5a, 0x49, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x67, 0x6f, 0x6c, 0x61,
	0x6e, 0x67, 0x2e, 0x6f, 0x72, 0x67, 0x2f, 0x67, 0x65, 0x6e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f,
	0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61, 0x70, 0x69, 0x73, 0x2f, 0x61, 0x64, 0x73, 0x2f, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61, 0x64, 0x73, 0x2f, 0x76, 0x31, 0x36, 0x2f, 0x73, 0x65, 0x72,
	0x76, 0x69, 0x63, 0x65, 0x73, 0x3b, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0xa2, 0x02,
	0x03, 0x47, 0x41, 0x41, 0xaa, 0x02, 0x21, 0x47, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x41, 0x64,
	0x73, 0x2e, 0x47, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x41, 0x64, 0x73, 0x2e, 0x56, 0x31, 0x36, 0x2e,
	0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0xca, 0x02, 0x21, 0x47, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x5c, 0x41, 0x64, 0x73, 0x5c, 0x47, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x41, 0x64, 0x73, 0x5c,
	0x56, 0x31, 0x36, 0x5c, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0xea, 0x02, 0x25, 0x47,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x3a, 0x3a, 0x41, 0x64, 0x73, 0x3a, 0x3a, 0x47, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x41, 0x64, 0x73, 0x3a, 0x3a, 0x56, 0x31, 0x36, 0x3a, 0x3a, 0x53, 0x65, 0x72, 0x76,
	0x69, 0x63, 0x65, 0x73, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_services_keyword_theme_constant_service_proto_rawDescOnce sync.Once
	file_services_keyword_theme_constant_service_proto_rawDescData = file_services_keyword_theme_constant_service_proto_rawDesc
)

func file_services_keyword_theme_constant_service_proto_rawDescGZIP() []byte {
	file_services_keyword_theme_constant_service_proto_rawDescOnce.Do(func() {
		file_services_keyword_theme_constant_service_proto_rawDescData = protoimpl.X.CompressGZIP(file_services_keyword_theme_constant_service_proto_rawDescData)
	})
	return file_services_keyword_theme_constant_service_proto_rawDescData
}

var file_services_keyword_theme_constant_service_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_services_keyword_theme_constant_service_proto_goTypes = []interface{}{
	(*SuggestKeywordThemeConstantsRequest)(nil),  // 0: google.ads.googleads.v16.services.SuggestKeywordThemeConstantsRequest
	(*SuggestKeywordThemeConstantsResponse)(nil), // 1: google.ads.googleads.v16.services.SuggestKeywordThemeConstantsResponse
	(*resources.KeywordThemeConstant)(nil),       // 2: google.ads.googleads.v16.resources.KeywordThemeConstant
}
var file_services_keyword_theme_constant_service_proto_depIdxs = []int32{
	2, // 0: google.ads.googleads.v16.services.SuggestKeywordThemeConstantsResponse.keyword_theme_constants:type_name -> google.ads.googleads.v16.resources.KeywordThemeConstant
	0, // 1: google.ads.googleads.v16.services.KeywordThemeConstantService.SuggestKeywordThemeConstants:input_type -> google.ads.googleads.v16.services.SuggestKeywordThemeConstantsRequest
	1, // 2: google.ads.googleads.v16.services.KeywordThemeConstantService.SuggestKeywordThemeConstants:output_type -> google.ads.googleads.v16.services.SuggestKeywordThemeConstantsResponse
	2, // [2:3] is the sub-list for method output_type
	1, // [1:2] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_services_keyword_theme_constant_service_proto_init() }
func file_services_keyword_theme_constant_service_proto_init() {
	if File_services_keyword_theme_constant_service_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_services_keyword_theme_constant_service_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SuggestKeywordThemeConstantsRequest); i {
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
		file_services_keyword_theme_constant_service_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SuggestKeywordThemeConstantsResponse); i {
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
			RawDescriptor: file_services_keyword_theme_constant_service_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_services_keyword_theme_constant_service_proto_goTypes,
		DependencyIndexes: file_services_keyword_theme_constant_service_proto_depIdxs,
		MessageInfos:      file_services_keyword_theme_constant_service_proto_msgTypes,
	}.Build()
	File_services_keyword_theme_constant_service_proto = out.File
	file_services_keyword_theme_constant_service_proto_rawDesc = nil
	file_services_keyword_theme_constant_service_proto_goTypes = nil
	file_services_keyword_theme_constant_service_proto_depIdxs = nil
}
