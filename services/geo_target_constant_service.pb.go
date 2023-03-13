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
// 	protoc        v3.21.12
// source: google/ads/googleads/v13/services/geo_target_constant_service.proto

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
// [GeoTargetConstantService.SuggestGeoTargetConstants][google.ads.googleads.v13.services.GeoTargetConstantService.SuggestGeoTargetConstants].
type SuggestGeoTargetConstantsRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// If possible, returned geo targets are translated using this locale. If not,
	// en is used by default. This is also used as a hint for returned geo
	// targets.
	Locale *string `protobuf:"bytes,6,opt,name=locale,proto3,oneof" json:"locale,omitempty"`
	// Returned geo targets are restricted to this country code.
	CountryCode *string `protobuf:"bytes,7,opt,name=country_code,json=countryCode,proto3,oneof" json:"country_code,omitempty"`
	// Required. A selector of geo target constants.
	//
	// Types that are assignable to Query:
	//	*SuggestGeoTargetConstantsRequest_LocationNames_
	//	*SuggestGeoTargetConstantsRequest_GeoTargets_
	Query isSuggestGeoTargetConstantsRequest_Query `protobuf_oneof:"query"`
}

func (x *SuggestGeoTargetConstantsRequest) Reset() {
	*x = SuggestGeoTargetConstantsRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_services_geo_target_constant_service_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SuggestGeoTargetConstantsRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SuggestGeoTargetConstantsRequest) ProtoMessage() {}

func (x *SuggestGeoTargetConstantsRequest) ProtoReflect() protoreflect.Message {
	mi := &file_services_geo_target_constant_service_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SuggestGeoTargetConstantsRequest.ProtoReflect.Descriptor instead.
func (*SuggestGeoTargetConstantsRequest) Descriptor() ([]byte, []int) {
	return file_services_geo_target_constant_service_proto_rawDescGZIP(), []int{0}
}

func (x *SuggestGeoTargetConstantsRequest) GetLocale() string {
	if x != nil && x.Locale != nil {
		return *x.Locale
	}
	return ""
}

func (x *SuggestGeoTargetConstantsRequest) GetCountryCode() string {
	if x != nil && x.CountryCode != nil {
		return *x.CountryCode
	}
	return ""
}

func (m *SuggestGeoTargetConstantsRequest) GetQuery() isSuggestGeoTargetConstantsRequest_Query {
	if m != nil {
		return m.Query
	}
	return nil
}

func (x *SuggestGeoTargetConstantsRequest) GetLocationNames() *SuggestGeoTargetConstantsRequest_LocationNames {
	if x, ok := x.GetQuery().(*SuggestGeoTargetConstantsRequest_LocationNames_); ok {
		return x.LocationNames
	}
	return nil
}

func (x *SuggestGeoTargetConstantsRequest) GetGeoTargets() *SuggestGeoTargetConstantsRequest_GeoTargets {
	if x, ok := x.GetQuery().(*SuggestGeoTargetConstantsRequest_GeoTargets_); ok {
		return x.GeoTargets
	}
	return nil
}

type isSuggestGeoTargetConstantsRequest_Query interface {
	isSuggestGeoTargetConstantsRequest_Query()
}

type SuggestGeoTargetConstantsRequest_LocationNames_ struct {
	// The location names to search by. At most 25 names can be set.
	LocationNames *SuggestGeoTargetConstantsRequest_LocationNames `protobuf:"bytes,1,opt,name=location_names,json=locationNames,proto3,oneof"`
}

type SuggestGeoTargetConstantsRequest_GeoTargets_ struct {
	// The geo target constant resource names to filter by.
	GeoTargets *SuggestGeoTargetConstantsRequest_GeoTargets `protobuf:"bytes,2,opt,name=geo_targets,json=geoTargets,proto3,oneof"`
}

func (*SuggestGeoTargetConstantsRequest_LocationNames_) isSuggestGeoTargetConstantsRequest_Query() {}

func (*SuggestGeoTargetConstantsRequest_GeoTargets_) isSuggestGeoTargetConstantsRequest_Query() {}

// Response message for
// [GeoTargetConstantService.SuggestGeoTargetConstants][google.ads.googleads.v13.services.GeoTargetConstantService.SuggestGeoTargetConstants].
type SuggestGeoTargetConstantsResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Geo target constant suggestions.
	GeoTargetConstantSuggestions []*GeoTargetConstantSuggestion `protobuf:"bytes,1,rep,name=geo_target_constant_suggestions,json=geoTargetConstantSuggestions,proto3" json:"geo_target_constant_suggestions,omitempty"`
}

func (x *SuggestGeoTargetConstantsResponse) Reset() {
	*x = SuggestGeoTargetConstantsResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_services_geo_target_constant_service_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SuggestGeoTargetConstantsResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SuggestGeoTargetConstantsResponse) ProtoMessage() {}

func (x *SuggestGeoTargetConstantsResponse) ProtoReflect() protoreflect.Message {
	mi := &file_services_geo_target_constant_service_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SuggestGeoTargetConstantsResponse.ProtoReflect.Descriptor instead.
func (*SuggestGeoTargetConstantsResponse) Descriptor() ([]byte, []int) {
	return file_services_geo_target_constant_service_proto_rawDescGZIP(), []int{1}
}

func (x *SuggestGeoTargetConstantsResponse) GetGeoTargetConstantSuggestions() []*GeoTargetConstantSuggestion {
	if x != nil {
		return x.GeoTargetConstantSuggestions
	}
	return nil
}

// A geo target constant suggestion.
type GeoTargetConstantSuggestion struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The language this GeoTargetConstantSuggestion is currently translated to.
	// It affects the name of geo target fields. For example, if locale=en, then
	// name=Spain. If locale=es, then name=España. The default locale will be
	// returned if no translation exists for the locale in the request.
	Locale *string `protobuf:"bytes,6,opt,name=locale,proto3,oneof" json:"locale,omitempty"`
	// Approximate user population that will be targeted, rounded to the
	// nearest 100.
	Reach *int64 `protobuf:"varint,7,opt,name=reach,proto3,oneof" json:"reach,omitempty"`
	// If the request searched by location name, this is the location name that
	// matched the geo target.
	SearchTerm *string `protobuf:"bytes,8,opt,name=search_term,json=searchTerm,proto3,oneof" json:"search_term,omitempty"`
	// The GeoTargetConstant result.
	GeoTargetConstant *resources.GeoTargetConstant `protobuf:"bytes,4,opt,name=geo_target_constant,json=geoTargetConstant,proto3" json:"geo_target_constant,omitempty"`
	// The list of parents of the geo target constant.
	GeoTargetConstantParents []*resources.GeoTargetConstant `protobuf:"bytes,5,rep,name=geo_target_constant_parents,json=geoTargetConstantParents,proto3" json:"geo_target_constant_parents,omitempty"`
}

func (x *GeoTargetConstantSuggestion) Reset() {
	*x = GeoTargetConstantSuggestion{}
	if protoimpl.UnsafeEnabled {
		mi := &file_services_geo_target_constant_service_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GeoTargetConstantSuggestion) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GeoTargetConstantSuggestion) ProtoMessage() {}

func (x *GeoTargetConstantSuggestion) ProtoReflect() protoreflect.Message {
	mi := &file_services_geo_target_constant_service_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GeoTargetConstantSuggestion.ProtoReflect.Descriptor instead.
func (*GeoTargetConstantSuggestion) Descriptor() ([]byte, []int) {
	return file_services_geo_target_constant_service_proto_rawDescGZIP(), []int{2}
}

func (x *GeoTargetConstantSuggestion) GetLocale() string {
	if x != nil && x.Locale != nil {
		return *x.Locale
	}
	return ""
}

func (x *GeoTargetConstantSuggestion) GetReach() int64 {
	if x != nil && x.Reach != nil {
		return *x.Reach
	}
	return 0
}

func (x *GeoTargetConstantSuggestion) GetSearchTerm() string {
	if x != nil && x.SearchTerm != nil {
		return *x.SearchTerm
	}
	return ""
}

func (x *GeoTargetConstantSuggestion) GetGeoTargetConstant() *resources.GeoTargetConstant {
	if x != nil {
		return x.GeoTargetConstant
	}
	return nil
}

func (x *GeoTargetConstantSuggestion) GetGeoTargetConstantParents() []*resources.GeoTargetConstant {
	if x != nil {
		return x.GeoTargetConstantParents
	}
	return nil
}

// A list of location names.
type SuggestGeoTargetConstantsRequest_LocationNames struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// A list of location names.
	Names []string `protobuf:"bytes,2,rep,name=names,proto3" json:"names,omitempty"`
}

func (x *SuggestGeoTargetConstantsRequest_LocationNames) Reset() {
	*x = SuggestGeoTargetConstantsRequest_LocationNames{}
	if protoimpl.UnsafeEnabled {
		mi := &file_services_geo_target_constant_service_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SuggestGeoTargetConstantsRequest_LocationNames) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SuggestGeoTargetConstantsRequest_LocationNames) ProtoMessage() {}

func (x *SuggestGeoTargetConstantsRequest_LocationNames) ProtoReflect() protoreflect.Message {
	mi := &file_services_geo_target_constant_service_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SuggestGeoTargetConstantsRequest_LocationNames.ProtoReflect.Descriptor instead.
func (*SuggestGeoTargetConstantsRequest_LocationNames) Descriptor() ([]byte, []int) {
	return file_services_geo_target_constant_service_proto_rawDescGZIP(), []int{0, 0}
}

func (x *SuggestGeoTargetConstantsRequest_LocationNames) GetNames() []string {
	if x != nil {
		return x.Names
	}
	return nil
}

// A list of geo target constant resource names.
type SuggestGeoTargetConstantsRequest_GeoTargets struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// A list of geo target constant resource names.
	GeoTargetConstants []string `protobuf:"bytes,2,rep,name=geo_target_constants,json=geoTargetConstants,proto3" json:"geo_target_constants,omitempty"`
}

func (x *SuggestGeoTargetConstantsRequest_GeoTargets) Reset() {
	*x = SuggestGeoTargetConstantsRequest_GeoTargets{}
	if protoimpl.UnsafeEnabled {
		mi := &file_services_geo_target_constant_service_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SuggestGeoTargetConstantsRequest_GeoTargets) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SuggestGeoTargetConstantsRequest_GeoTargets) ProtoMessage() {}

func (x *SuggestGeoTargetConstantsRequest_GeoTargets) ProtoReflect() protoreflect.Message {
	mi := &file_services_geo_target_constant_service_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SuggestGeoTargetConstantsRequest_GeoTargets.ProtoReflect.Descriptor instead.
func (*SuggestGeoTargetConstantsRequest_GeoTargets) Descriptor() ([]byte, []int) {
	return file_services_geo_target_constant_service_proto_rawDescGZIP(), []int{0, 1}
}

func (x *SuggestGeoTargetConstantsRequest_GeoTargets) GetGeoTargetConstants() []string {
	if x != nil {
		return x.GeoTargetConstants
	}
	return nil
}

var File_services_geo_target_constant_service_proto protoreflect.FileDescriptor

var file_services_geo_target_constant_service_proto_rawDesc = []byte{
	0x0a, 0x43, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x64, 0x73, 0x2f, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x61, 0x64, 0x73, 0x2f, 0x76, 0x31, 0x33, 0x2f, 0x73, 0x65, 0x72, 0x76, 0x69,
	0x63, 0x65, 0x73, 0x2f, 0x67, 0x65, 0x6f, 0x5f, 0x74, 0x61, 0x72, 0x67, 0x65, 0x74, 0x5f, 0x63,
	0x6f, 0x6e, 0x73, 0x74, 0x61, 0x6e, 0x74, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x21, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x61, 0x64,
	0x73, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61, 0x64, 0x73, 0x2e, 0x76, 0x31, 0x33, 0x2e,
	0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x1a, 0x3c, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x2f, 0x61, 0x64, 0x73, 0x2f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61, 0x64, 0x73, 0x2f, 0x76,
	0x31, 0x33, 0x2f, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x73, 0x2f, 0x67, 0x65, 0x6f,
	0x5f, 0x74, 0x61, 0x72, 0x67, 0x65, 0x74, 0x5f, 0x63, 0x6f, 0x6e, 0x73, 0x74, 0x61, 0x6e, 0x74,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1c, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61,
	0x70, 0x69, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x17, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69,
	0x2f, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xe2, 0x03,
	0x0a, 0x20, 0x53, 0x75, 0x67, 0x67, 0x65, 0x73, 0x74, 0x47, 0x65, 0x6f, 0x54, 0x61, 0x72, 0x67,
	0x65, 0x74, 0x43, 0x6f, 0x6e, 0x73, 0x74, 0x61, 0x6e, 0x74, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x12, 0x1b, 0x0a, 0x06, 0x6c, 0x6f, 0x63, 0x61, 0x6c, 0x65, 0x18, 0x06, 0x20, 0x01,
	0x28, 0x09, 0x48, 0x01, 0x52, 0x06, 0x6c, 0x6f, 0x63, 0x61, 0x6c, 0x65, 0x88, 0x01, 0x01, 0x12,
	0x26, 0x0a, 0x0c, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x72, 0x79, 0x5f, 0x63, 0x6f, 0x64, 0x65, 0x18,
	0x07, 0x20, 0x01, 0x28, 0x09, 0x48, 0x02, 0x52, 0x0b, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x72, 0x79,
	0x43, 0x6f, 0x64, 0x65, 0x88, 0x01, 0x01, 0x12, 0x7a, 0x0a, 0x0e, 0x6c, 0x6f, 0x63, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x51, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x61, 0x64, 0x73, 0x2e, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x61, 0x64, 0x73, 0x2e, 0x76, 0x31, 0x33, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69,
	0x63, 0x65, 0x73, 0x2e, 0x53, 0x75, 0x67, 0x67, 0x65, 0x73, 0x74, 0x47, 0x65, 0x6f, 0x54, 0x61,
	0x72, 0x67, 0x65, 0x74, 0x43, 0x6f, 0x6e, 0x73, 0x74, 0x61, 0x6e, 0x74, 0x73, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x2e, 0x4c, 0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x4e, 0x61, 0x6d,
	0x65, 0x73, 0x48, 0x00, 0x52, 0x0d, 0x6c, 0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x4e, 0x61,
	0x6d, 0x65, 0x73, 0x12, 0x71, 0x0a, 0x0b, 0x67, 0x65, 0x6f, 0x5f, 0x74, 0x61, 0x72, 0x67, 0x65,
	0x74, 0x73, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x4e, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x2e, 0x61, 0x64, 0x73, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61, 0x64, 0x73, 0x2e,
	0x76, 0x31, 0x33, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x53, 0x75, 0x67,
	0x67, 0x65, 0x73, 0x74, 0x47, 0x65, 0x6f, 0x54, 0x61, 0x72, 0x67, 0x65, 0x74, 0x43, 0x6f, 0x6e,
	0x73, 0x74, 0x61, 0x6e, 0x74, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x2e, 0x47, 0x65,
	0x6f, 0x54, 0x61, 0x72, 0x67, 0x65, 0x74, 0x73, 0x48, 0x00, 0x52, 0x0a, 0x67, 0x65, 0x6f, 0x54,
	0x61, 0x72, 0x67, 0x65, 0x74, 0x73, 0x1a, 0x25, 0x0a, 0x0d, 0x4c, 0x6f, 0x63, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x4e, 0x61, 0x6d, 0x65, 0x73, 0x12, 0x14, 0x0a, 0x05, 0x6e, 0x61, 0x6d, 0x65, 0x73,
	0x18, 0x02, 0x20, 0x03, 0x28, 0x09, 0x52, 0x05, 0x6e, 0x61, 0x6d, 0x65, 0x73, 0x1a, 0x3e, 0x0a,
	0x0a, 0x47, 0x65, 0x6f, 0x54, 0x61, 0x72, 0x67, 0x65, 0x74, 0x73, 0x12, 0x30, 0x0a, 0x14, 0x67,
	0x65, 0x6f, 0x5f, 0x74, 0x61, 0x72, 0x67, 0x65, 0x74, 0x5f, 0x63, 0x6f, 0x6e, 0x73, 0x74, 0x61,
	0x6e, 0x74, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x09, 0x52, 0x12, 0x67, 0x65, 0x6f, 0x54, 0x61,
	0x72, 0x67, 0x65, 0x74, 0x43, 0x6f, 0x6e, 0x73, 0x74, 0x61, 0x6e, 0x74, 0x73, 0x42, 0x07, 0x0a,
	0x05, 0x71, 0x75, 0x65, 0x72, 0x79, 0x42, 0x09, 0x0a, 0x07, 0x5f, 0x6c, 0x6f, 0x63, 0x61, 0x6c,
	0x65, 0x42, 0x0f, 0x0a, 0x0d, 0x5f, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x72, 0x79, 0x5f, 0x63, 0x6f,
	0x64, 0x65, 0x22, 0xab, 0x01, 0x0a, 0x21, 0x53, 0x75, 0x67, 0x67, 0x65, 0x73, 0x74, 0x47, 0x65,
	0x6f, 0x54, 0x61, 0x72, 0x67, 0x65, 0x74, 0x43, 0x6f, 0x6e, 0x73, 0x74, 0x61, 0x6e, 0x74, 0x73,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x85, 0x01, 0x0a, 0x1f, 0x67, 0x65, 0x6f,
	0x5f, 0x74, 0x61, 0x72, 0x67, 0x65, 0x74, 0x5f, 0x63, 0x6f, 0x6e, 0x73, 0x74, 0x61, 0x6e, 0x74,
	0x5f, 0x73, 0x75, 0x67, 0x67, 0x65, 0x73, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x18, 0x01, 0x20, 0x03,
	0x28, 0x0b, 0x32, 0x3e, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x61, 0x64, 0x73, 0x2e,
	0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61, 0x64, 0x73, 0x2e, 0x76, 0x31, 0x33, 0x2e, 0x73, 0x65,
	0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x47, 0x65, 0x6f, 0x54, 0x61, 0x72, 0x67, 0x65, 0x74,
	0x43, 0x6f, 0x6e, 0x73, 0x74, 0x61, 0x6e, 0x74, 0x53, 0x75, 0x67, 0x67, 0x65, 0x73, 0x74, 0x69,
	0x6f, 0x6e, 0x52, 0x1c, 0x67, 0x65, 0x6f, 0x54, 0x61, 0x72, 0x67, 0x65, 0x74, 0x43, 0x6f, 0x6e,
	0x73, 0x74, 0x61, 0x6e, 0x74, 0x53, 0x75, 0x67, 0x67, 0x65, 0x73, 0x74, 0x69, 0x6f, 0x6e, 0x73,
	0x22, 0xfd, 0x02, 0x0a, 0x1b, 0x47, 0x65, 0x6f, 0x54, 0x61, 0x72, 0x67, 0x65, 0x74, 0x43, 0x6f,
	0x6e, 0x73, 0x74, 0x61, 0x6e, 0x74, 0x53, 0x75, 0x67, 0x67, 0x65, 0x73, 0x74, 0x69, 0x6f, 0x6e,
	0x12, 0x1b, 0x0a, 0x06, 0x6c, 0x6f, 0x63, 0x61, 0x6c, 0x65, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09,
	0x48, 0x00, 0x52, 0x06, 0x6c, 0x6f, 0x63, 0x61, 0x6c, 0x65, 0x88, 0x01, 0x01, 0x12, 0x19, 0x0a,
	0x05, 0x72, 0x65, 0x61, 0x63, 0x68, 0x18, 0x07, 0x20, 0x01, 0x28, 0x03, 0x48, 0x01, 0x52, 0x05,
	0x72, 0x65, 0x61, 0x63, 0x68, 0x88, 0x01, 0x01, 0x12, 0x24, 0x0a, 0x0b, 0x73, 0x65, 0x61, 0x72,
	0x63, 0x68, 0x5f, 0x74, 0x65, 0x72, 0x6d, 0x18, 0x08, 0x20, 0x01, 0x28, 0x09, 0x48, 0x02, 0x52,
	0x0a, 0x73, 0x65, 0x61, 0x72, 0x63, 0x68, 0x54, 0x65, 0x72, 0x6d, 0x88, 0x01, 0x01, 0x12, 0x65,
	0x0a, 0x13, 0x67, 0x65, 0x6f, 0x5f, 0x74, 0x61, 0x72, 0x67, 0x65, 0x74, 0x5f, 0x63, 0x6f, 0x6e,
	0x73, 0x74, 0x61, 0x6e, 0x74, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x35, 0x2e, 0x67, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x61, 0x64, 0x73, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61,
	0x64, 0x73, 0x2e, 0x76, 0x31, 0x33, 0x2e, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x73,
	0x2e, 0x47, 0x65, 0x6f, 0x54, 0x61, 0x72, 0x67, 0x65, 0x74, 0x43, 0x6f, 0x6e, 0x73, 0x74, 0x61,
	0x6e, 0x74, 0x52, 0x11, 0x67, 0x65, 0x6f, 0x54, 0x61, 0x72, 0x67, 0x65, 0x74, 0x43, 0x6f, 0x6e,
	0x73, 0x74, 0x61, 0x6e, 0x74, 0x12, 0x74, 0x0a, 0x1b, 0x67, 0x65, 0x6f, 0x5f, 0x74, 0x61, 0x72,
	0x67, 0x65, 0x74, 0x5f, 0x63, 0x6f, 0x6e, 0x73, 0x74, 0x61, 0x6e, 0x74, 0x5f, 0x70, 0x61, 0x72,
	0x65, 0x6e, 0x74, 0x73, 0x18, 0x05, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x35, 0x2e, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x2e, 0x61, 0x64, 0x73, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61, 0x64,
	0x73, 0x2e, 0x76, 0x31, 0x33, 0x2e, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x73, 0x2e,
	0x47, 0x65, 0x6f, 0x54, 0x61, 0x72, 0x67, 0x65, 0x74, 0x43, 0x6f, 0x6e, 0x73, 0x74, 0x61, 0x6e,
	0x74, 0x52, 0x18, 0x67, 0x65, 0x6f, 0x54, 0x61, 0x72, 0x67, 0x65, 0x74, 0x43, 0x6f, 0x6e, 0x73,
	0x74, 0x61, 0x6e, 0x74, 0x50, 0x61, 0x72, 0x65, 0x6e, 0x74, 0x73, 0x42, 0x09, 0x0a, 0x07, 0x5f,
	0x6c, 0x6f, 0x63, 0x61, 0x6c, 0x65, 0x42, 0x08, 0x0a, 0x06, 0x5f, 0x72, 0x65, 0x61, 0x63, 0x68,
	0x42, 0x0e, 0x0a, 0x0c, 0x5f, 0x73, 0x65, 0x61, 0x72, 0x63, 0x68, 0x5f, 0x74, 0x65, 0x72, 0x6d,
	0x32, 0xb6, 0x02, 0x0a, 0x18, 0x47, 0x65, 0x6f, 0x54, 0x61, 0x72, 0x67, 0x65, 0x74, 0x43, 0x6f,
	0x6e, 0x73, 0x74, 0x61, 0x6e, 0x74, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0xd2, 0x01,
	0x0a, 0x19, 0x53, 0x75, 0x67, 0x67, 0x65, 0x73, 0x74, 0x47, 0x65, 0x6f, 0x54, 0x61, 0x72, 0x67,
	0x65, 0x74, 0x43, 0x6f, 0x6e, 0x73, 0x74, 0x61, 0x6e, 0x74, 0x73, 0x12, 0x43, 0x2e, 0x67, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x61, 0x64, 0x73, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61,
	0x64, 0x73, 0x2e, 0x76, 0x31, 0x33, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e,
	0x53, 0x75, 0x67, 0x67, 0x65, 0x73, 0x74, 0x47, 0x65, 0x6f, 0x54, 0x61, 0x72, 0x67, 0x65, 0x74,
	0x43, 0x6f, 0x6e, 0x73, 0x74, 0x61, 0x6e, 0x74, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x1a, 0x44, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x61, 0x64, 0x73, 0x2e, 0x67, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x61, 0x64, 0x73, 0x2e, 0x76, 0x31, 0x33, 0x2e, 0x73, 0x65, 0x72, 0x76,
	0x69, 0x63, 0x65, 0x73, 0x2e, 0x53, 0x75, 0x67, 0x67, 0x65, 0x73, 0x74, 0x47, 0x65, 0x6f, 0x54,
	0x61, 0x72, 0x67, 0x65, 0x74, 0x43, 0x6f, 0x6e, 0x73, 0x74, 0x61, 0x6e, 0x74, 0x73, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x2a, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x24, 0x22, 0x1f,
	0x2f, 0x76, 0x31, 0x33, 0x2f, 0x67, 0x65, 0x6f, 0x54, 0x61, 0x72, 0x67, 0x65, 0x74, 0x43, 0x6f,
	0x6e, 0x73, 0x74, 0x61, 0x6e, 0x74, 0x73, 0x3a, 0x73, 0x75, 0x67, 0x67, 0x65, 0x73, 0x74, 0x3a,
	0x01, 0x2a, 0x1a, 0x45, 0xca, 0x41, 0x18, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61, 0x64, 0x73,
	0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61, 0x70, 0x69, 0x73, 0x2e, 0x63, 0x6f, 0x6d, 0xd2,
	0x41, 0x27, 0x68, 0x74, 0x74, 0x70, 0x73, 0x3a, 0x2f, 0x2f, 0x77, 0x77, 0x77, 0x2e, 0x67, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x61, 0x70, 0x69, 0x73, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x61, 0x75, 0x74,
	0x68, 0x2f, 0x61, 0x64, 0x77, 0x6f, 0x72, 0x64, 0x73, 0x42, 0x89, 0x02, 0x0a, 0x25, 0x63, 0x6f,
	0x6d, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x61, 0x64, 0x73, 0x2e, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x61, 0x64, 0x73, 0x2e, 0x76, 0x31, 0x33, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69,
	0x63, 0x65, 0x73, 0x42, 0x1d, 0x47, 0x65, 0x6f, 0x54, 0x61, 0x72, 0x67, 0x65, 0x74, 0x43, 0x6f,
	0x6e, 0x73, 0x74, 0x61, 0x6e, 0x74, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x50, 0x72, 0x6f,
	0x74, 0x6f, 0x50, 0x01, 0x5a, 0x49, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x67, 0x6f, 0x6c,
	0x61, 0x6e, 0x67, 0x2e, 0x6f, 0x72, 0x67, 0x2f, 0x67, 0x65, 0x6e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x2f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61, 0x70, 0x69, 0x73, 0x2f, 0x61, 0x64, 0x73, 0x2f,
	0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61, 0x64, 0x73, 0x2f, 0x76, 0x31, 0x33, 0x2f, 0x73, 0x65,
	0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x3b, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0xa2,
	0x02, 0x03, 0x47, 0x41, 0x41, 0xaa, 0x02, 0x21, 0x47, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x41,
	0x64, 0x73, 0x2e, 0x47, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x41, 0x64, 0x73, 0x2e, 0x56, 0x31, 0x33,
	0x2e, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0xca, 0x02, 0x21, 0x47, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x5c, 0x41, 0x64, 0x73, 0x5c, 0x47, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x41, 0x64, 0x73,
	0x5c, 0x56, 0x31, 0x33, 0x5c, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0xea, 0x02, 0x25,
	0x47, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x3a, 0x3a, 0x41, 0x64, 0x73, 0x3a, 0x3a, 0x47, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x41, 0x64, 0x73, 0x3a, 0x3a, 0x56, 0x31, 0x33, 0x3a, 0x3a, 0x53, 0x65, 0x72,
	0x76, 0x69, 0x63, 0x65, 0x73, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_services_geo_target_constant_service_proto_rawDescOnce sync.Once
	file_services_geo_target_constant_service_proto_rawDescData = file_services_geo_target_constant_service_proto_rawDesc
)

func file_services_geo_target_constant_service_proto_rawDescGZIP() []byte {
	file_services_geo_target_constant_service_proto_rawDescOnce.Do(func() {
		file_services_geo_target_constant_service_proto_rawDescData = protoimpl.X.CompressGZIP(file_services_geo_target_constant_service_proto_rawDescData)
	})
	return file_services_geo_target_constant_service_proto_rawDescData
}

var file_services_geo_target_constant_service_proto_msgTypes = make([]protoimpl.MessageInfo, 5)
var file_services_geo_target_constant_service_proto_goTypes = []interface{}{
	(*SuggestGeoTargetConstantsRequest)(nil),               // 0: google.ads.googleads.v13.services.SuggestGeoTargetConstantsRequest
	(*SuggestGeoTargetConstantsResponse)(nil),              // 1: google.ads.googleads.v13.services.SuggestGeoTargetConstantsResponse
	(*GeoTargetConstantSuggestion)(nil),                    // 2: google.ads.googleads.v13.services.GeoTargetConstantSuggestion
	(*SuggestGeoTargetConstantsRequest_LocationNames)(nil), // 3: google.ads.googleads.v13.services.SuggestGeoTargetConstantsRequest.LocationNames
	(*SuggestGeoTargetConstantsRequest_GeoTargets)(nil),    // 4: google.ads.googleads.v13.services.SuggestGeoTargetConstantsRequest.GeoTargets
	(*resources.GeoTargetConstant)(nil),                    // 5: google.ads.googleads.v13.resources.GeoTargetConstant
}
var file_services_geo_target_constant_service_proto_depIdxs = []int32{
	3, // 0: google.ads.googleads.v13.services.SuggestGeoTargetConstantsRequest.location_names:type_name -> google.ads.googleads.v13.services.SuggestGeoTargetConstantsRequest.LocationNames
	4, // 1: google.ads.googleads.v13.services.SuggestGeoTargetConstantsRequest.geo_targets:type_name -> google.ads.googleads.v13.services.SuggestGeoTargetConstantsRequest.GeoTargets
	2, // 2: google.ads.googleads.v13.services.SuggestGeoTargetConstantsResponse.geo_target_constant_suggestions:type_name -> google.ads.googleads.v13.services.GeoTargetConstantSuggestion
	5, // 3: google.ads.googleads.v13.services.GeoTargetConstantSuggestion.geo_target_constant:type_name -> google.ads.googleads.v13.resources.GeoTargetConstant
	5, // 4: google.ads.googleads.v13.services.GeoTargetConstantSuggestion.geo_target_constant_parents:type_name -> google.ads.googleads.v13.resources.GeoTargetConstant
	0, // 5: google.ads.googleads.v13.services.GeoTargetConstantService.SuggestGeoTargetConstants:input_type -> google.ads.googleads.v13.services.SuggestGeoTargetConstantsRequest
	1, // 6: google.ads.googleads.v13.services.GeoTargetConstantService.SuggestGeoTargetConstants:output_type -> google.ads.googleads.v13.services.SuggestGeoTargetConstantsResponse
	6, // [6:7] is the sub-list for method output_type
	5, // [5:6] is the sub-list for method input_type
	5, // [5:5] is the sub-list for extension type_name
	5, // [5:5] is the sub-list for extension extendee
	0, // [0:5] is the sub-list for field type_name
}

func init() { file_services_geo_target_constant_service_proto_init() }
func file_services_geo_target_constant_service_proto_init() {
	if File_services_geo_target_constant_service_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_services_geo_target_constant_service_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SuggestGeoTargetConstantsRequest); i {
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
		file_services_geo_target_constant_service_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SuggestGeoTargetConstantsResponse); i {
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
		file_services_geo_target_constant_service_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GeoTargetConstantSuggestion); i {
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
		file_services_geo_target_constant_service_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SuggestGeoTargetConstantsRequest_LocationNames); i {
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
		file_services_geo_target_constant_service_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SuggestGeoTargetConstantsRequest_GeoTargets); i {
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
	file_services_geo_target_constant_service_proto_msgTypes[0].OneofWrappers = []interface{}{
		(*SuggestGeoTargetConstantsRequest_LocationNames_)(nil),
		(*SuggestGeoTargetConstantsRequest_GeoTargets_)(nil),
	}
	file_services_geo_target_constant_service_proto_msgTypes[2].OneofWrappers = []interface{}{}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_services_geo_target_constant_service_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   5,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_services_geo_target_constant_service_proto_goTypes,
		DependencyIndexes: file_services_geo_target_constant_service_proto_depIdxs,
		MessageInfos:      file_services_geo_target_constant_service_proto_msgTypes,
	}.Build()
	File_services_geo_target_constant_service_proto = out.File
	file_services_geo_target_constant_service_proto_rawDesc = nil
	file_services_geo_target_constant_service_proto_goTypes = nil
	file_services_geo_target_constant_service_proto_depIdxs = nil
}
