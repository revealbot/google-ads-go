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
// source: google/ads/googleads/v17/resources/smart_campaign_setting.proto

package resources

import (
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

// Settings for configuring Smart campaigns.
type SmartCampaignSetting struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Immutable. The resource name of the Smart campaign setting.
	// Smart campaign setting resource names have the form:
	//
	// `customers/{customer_id}/smartCampaignSettings/{campaign_id}`
	ResourceName string `protobuf:"bytes,1,opt,name=resource_name,json=resourceName,proto3" json:"resource_name,omitempty"`
	// Output only. The campaign to which these settings apply.
	Campaign string `protobuf:"bytes,2,opt,name=campaign,proto3" json:"campaign,omitempty"`
	// Phone number and country code.
	PhoneNumber *SmartCampaignSetting_PhoneNumber `protobuf:"bytes,3,opt,name=phone_number,json=phoneNumber,proto3" json:"phone_number,omitempty"`
	// The language code to advertise in from the set of
	// [supported language codes]
	// (https://developers.google.com/google-ads/api/reference/data/codes-formats#languages).
	AdvertisingLanguageCode string `protobuf:"bytes,7,opt,name=advertising_language_code,json=advertisingLanguageCode,proto3" json:"advertising_language_code,omitempty"`
	// The landing page of this campaign.
	//
	// Types that are assignable to LandingPage:
	//
	//	*SmartCampaignSetting_FinalUrl
	//	*SmartCampaignSetting_AdOptimizedBusinessProfileSetting_
	LandingPage isSmartCampaignSetting_LandingPage `protobuf_oneof:"landing_page"`
	// The business setting of this campaign.
	//
	// Types that are assignable to BusinessSetting:
	//
	//	*SmartCampaignSetting_BusinessName
	//	*SmartCampaignSetting_BusinessProfileLocation
	BusinessSetting isSmartCampaignSetting_BusinessSetting `protobuf_oneof:"business_setting"`
}

func (x *SmartCampaignSetting) Reset() {
	*x = SmartCampaignSetting{}
	if protoimpl.UnsafeEnabled {
		mi := &file_resources_smart_campaign_setting_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SmartCampaignSetting) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SmartCampaignSetting) ProtoMessage() {}

func (x *SmartCampaignSetting) ProtoReflect() protoreflect.Message {
	mi := &file_resources_smart_campaign_setting_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SmartCampaignSetting.ProtoReflect.Descriptor instead.
func (*SmartCampaignSetting) Descriptor() ([]byte, []int) {
	return file_resources_smart_campaign_setting_proto_rawDescGZIP(), []int{0}
}

func (x *SmartCampaignSetting) GetResourceName() string {
	if x != nil {
		return x.ResourceName
	}
	return ""
}

func (x *SmartCampaignSetting) GetCampaign() string {
	if x != nil {
		return x.Campaign
	}
	return ""
}

func (x *SmartCampaignSetting) GetPhoneNumber() *SmartCampaignSetting_PhoneNumber {
	if x != nil {
		return x.PhoneNumber
	}
	return nil
}

func (x *SmartCampaignSetting) GetAdvertisingLanguageCode() string {
	if x != nil {
		return x.AdvertisingLanguageCode
	}
	return ""
}

func (m *SmartCampaignSetting) GetLandingPage() isSmartCampaignSetting_LandingPage {
	if m != nil {
		return m.LandingPage
	}
	return nil
}

func (x *SmartCampaignSetting) GetFinalUrl() string {
	if x, ok := x.GetLandingPage().(*SmartCampaignSetting_FinalUrl); ok {
		return x.FinalUrl
	}
	return ""
}

func (x *SmartCampaignSetting) GetAdOptimizedBusinessProfileSetting() *SmartCampaignSetting_AdOptimizedBusinessProfileSetting {
	if x, ok := x.GetLandingPage().(*SmartCampaignSetting_AdOptimizedBusinessProfileSetting_); ok {
		return x.AdOptimizedBusinessProfileSetting
	}
	return nil
}

func (m *SmartCampaignSetting) GetBusinessSetting() isSmartCampaignSetting_BusinessSetting {
	if m != nil {
		return m.BusinessSetting
	}
	return nil
}

func (x *SmartCampaignSetting) GetBusinessName() string {
	if x, ok := x.GetBusinessSetting().(*SmartCampaignSetting_BusinessName); ok {
		return x.BusinessName
	}
	return ""
}

func (x *SmartCampaignSetting) GetBusinessProfileLocation() string {
	if x, ok := x.GetBusinessSetting().(*SmartCampaignSetting_BusinessProfileLocation); ok {
		return x.BusinessProfileLocation
	}
	return ""
}

type isSmartCampaignSetting_LandingPage interface {
	isSmartCampaignSetting_LandingPage()
}

type SmartCampaignSetting_FinalUrl struct {
	// The user-provided landing page URL for this Campaign.
	FinalUrl string `protobuf:"bytes,8,opt,name=final_url,json=finalUrl,proto3,oneof"`
}

type SmartCampaignSetting_AdOptimizedBusinessProfileSetting_ struct {
	// Settings for configuring a business profile optimized for ads as this
	// campaign's landing page.  This campaign must be linked to a business
	// profile to use this option.  For more information on this feature,
	// consult https://support.google.com/google-ads/answer/9827068.
	AdOptimizedBusinessProfileSetting *SmartCampaignSetting_AdOptimizedBusinessProfileSetting `protobuf:"bytes,9,opt,name=ad_optimized_business_profile_setting,json=adOptimizedBusinessProfileSetting,proto3,oneof"`
}

func (*SmartCampaignSetting_FinalUrl) isSmartCampaignSetting_LandingPage() {}

func (*SmartCampaignSetting_AdOptimizedBusinessProfileSetting_) isSmartCampaignSetting_LandingPage() {
}

type isSmartCampaignSetting_BusinessSetting interface {
	isSmartCampaignSetting_BusinessSetting()
}

type SmartCampaignSetting_BusinessName struct {
	// The name of the business.
	BusinessName string `protobuf:"bytes,5,opt,name=business_name,json=businessName,proto3,oneof"`
}

type SmartCampaignSetting_BusinessProfileLocation struct {
	// The resource name of a Business Profile location.
	// Business Profile location resource names can be fetched through the
	// Business Profile API and adhere to the following format:
	// `locations/{locationId}`.
	//
	// See the [Business Profile API]
	// (https://developers.google.com/my-business/reference/businessinformation/rest/v1/accounts.locations)
	// for additional details.
	BusinessProfileLocation string `protobuf:"bytes,10,opt,name=business_profile_location,json=businessProfileLocation,proto3,oneof"`
}

func (*SmartCampaignSetting_BusinessName) isSmartCampaignSetting_BusinessSetting() {}

func (*SmartCampaignSetting_BusinessProfileLocation) isSmartCampaignSetting_BusinessSetting() {}

// Phone number and country code in smart campaign settings.
type SmartCampaignSetting_PhoneNumber struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Phone number of the smart campaign.
	PhoneNumber *string `protobuf:"bytes,1,opt,name=phone_number,json=phoneNumber,proto3,oneof" json:"phone_number,omitempty"`
	// Upper-case, two-letter country code as defined by ISO-3166.
	CountryCode *string `protobuf:"bytes,2,opt,name=country_code,json=countryCode,proto3,oneof" json:"country_code,omitempty"`
}

func (x *SmartCampaignSetting_PhoneNumber) Reset() {
	*x = SmartCampaignSetting_PhoneNumber{}
	if protoimpl.UnsafeEnabled {
		mi := &file_resources_smart_campaign_setting_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SmartCampaignSetting_PhoneNumber) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SmartCampaignSetting_PhoneNumber) ProtoMessage() {}

func (x *SmartCampaignSetting_PhoneNumber) ProtoReflect() protoreflect.Message {
	mi := &file_resources_smart_campaign_setting_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SmartCampaignSetting_PhoneNumber.ProtoReflect.Descriptor instead.
func (*SmartCampaignSetting_PhoneNumber) Descriptor() ([]byte, []int) {
	return file_resources_smart_campaign_setting_proto_rawDescGZIP(), []int{0, 0}
}

func (x *SmartCampaignSetting_PhoneNumber) GetPhoneNumber() string {
	if x != nil && x.PhoneNumber != nil {
		return *x.PhoneNumber
	}
	return ""
}

func (x *SmartCampaignSetting_PhoneNumber) GetCountryCode() string {
	if x != nil && x.CountryCode != nil {
		return *x.CountryCode
	}
	return ""
}

// Settings for configuring a business profile optimized for ads as this
// campaign's landing page.
type SmartCampaignSetting_AdOptimizedBusinessProfileSetting struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Enabling a lead form on your business profile enables prospective
	// customers to contact your business by filling out a simple form,
	// and you'll receive their information through email.
	IncludeLeadForm *bool `protobuf:"varint,1,opt,name=include_lead_form,json=includeLeadForm,proto3,oneof" json:"include_lead_form,omitempty"`
}

func (x *SmartCampaignSetting_AdOptimizedBusinessProfileSetting) Reset() {
	*x = SmartCampaignSetting_AdOptimizedBusinessProfileSetting{}
	if protoimpl.UnsafeEnabled {
		mi := &file_resources_smart_campaign_setting_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SmartCampaignSetting_AdOptimizedBusinessProfileSetting) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SmartCampaignSetting_AdOptimizedBusinessProfileSetting) ProtoMessage() {}

func (x *SmartCampaignSetting_AdOptimizedBusinessProfileSetting) ProtoReflect() protoreflect.Message {
	mi := &file_resources_smart_campaign_setting_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SmartCampaignSetting_AdOptimizedBusinessProfileSetting.ProtoReflect.Descriptor instead.
func (*SmartCampaignSetting_AdOptimizedBusinessProfileSetting) Descriptor() ([]byte, []int) {
	return file_resources_smart_campaign_setting_proto_rawDescGZIP(), []int{0, 1}
}

func (x *SmartCampaignSetting_AdOptimizedBusinessProfileSetting) GetIncludeLeadForm() bool {
	if x != nil && x.IncludeLeadForm != nil {
		return *x.IncludeLeadForm
	}
	return false
}

var File_resources_smart_campaign_setting_proto protoreflect.FileDescriptor

var file_resources_smart_campaign_setting_proto_rawDesc = []byte{
	0x0a, 0x3f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x64, 0x73, 0x2f, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x61, 0x64, 0x73, 0x2f, 0x76, 0x31, 0x37, 0x2f, 0x72, 0x65, 0x73, 0x6f, 0x75,
	0x72, 0x63, 0x65, 0x73, 0x2f, 0x73, 0x6d, 0x61, 0x72, 0x74, 0x5f, 0x63, 0x61, 0x6d, 0x70, 0x61,
	0x69, 0x67, 0x6e, 0x5f, 0x73, 0x65, 0x74, 0x74, 0x69, 0x6e, 0x67, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x12, 0x22, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x61, 0x64, 0x73, 0x2e, 0x67, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x61, 0x64, 0x73, 0x2e, 0x76, 0x31, 0x37, 0x2e, 0x72, 0x65, 0x73, 0x6f,
	0x75, 0x72, 0x63, 0x65, 0x73, 0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70,
	0x69, 0x2f, 0x66, 0x69, 0x65, 0x6c, 0x64, 0x5f, 0x62, 0x65, 0x68, 0x61, 0x76, 0x69, 0x6f, 0x72,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x19, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61,
	0x70, 0x69, 0x2f, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x22, 0x95, 0x08, 0x0a, 0x14, 0x53, 0x6d, 0x61, 0x72, 0x74, 0x43, 0x61, 0x6d, 0x70, 0x61,
	0x69, 0x67, 0x6e, 0x53, 0x65, 0x74, 0x74, 0x69, 0x6e, 0x67, 0x12, 0x5a, 0x0a, 0x0d, 0x72, 0x65,
	0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x42, 0x35, 0xe0, 0x41, 0x05, 0xfa, 0x41, 0x2f, 0x0a, 0x2d, 0x67, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x61, 0x64, 0x73, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61, 0x70, 0x69, 0x73, 0x2e,
	0x63, 0x6f, 0x6d, 0x2f, 0x53, 0x6d, 0x61, 0x72, 0x74, 0x43, 0x61, 0x6d, 0x70, 0x61, 0x69, 0x67,
	0x6e, 0x53, 0x65, 0x74, 0x74, 0x69, 0x6e, 0x67, 0x52, 0x0c, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72,
	0x63, 0x65, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x45, 0x0a, 0x08, 0x63, 0x61, 0x6d, 0x70, 0x61, 0x69,
	0x67, 0x6e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x42, 0x29, 0xe0, 0x41, 0x03, 0xfa, 0x41, 0x23,
	0x0a, 0x21, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61, 0x64, 0x73, 0x2e, 0x67, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x61, 0x70, 0x69, 0x73, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x43, 0x61, 0x6d, 0x70, 0x61,
	0x69, 0x67, 0x6e, 0x52, 0x08, 0x63, 0x61, 0x6d, 0x70, 0x61, 0x69, 0x67, 0x6e, 0x12, 0x67, 0x0a,
	0x0c, 0x70, 0x68, 0x6f, 0x6e, 0x65, 0x5f, 0x6e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x18, 0x03, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x44, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x61, 0x64, 0x73,
	0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61, 0x64, 0x73, 0x2e, 0x76, 0x31, 0x37, 0x2e, 0x72,
	0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x73, 0x2e, 0x53, 0x6d, 0x61, 0x72, 0x74, 0x43, 0x61,
	0x6d, 0x70, 0x61, 0x69, 0x67, 0x6e, 0x53, 0x65, 0x74, 0x74, 0x69, 0x6e, 0x67, 0x2e, 0x50, 0x68,
	0x6f, 0x6e, 0x65, 0x4e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x52, 0x0b, 0x70, 0x68, 0x6f, 0x6e, 0x65,
	0x4e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x12, 0x3a, 0x0a, 0x19, 0x61, 0x64, 0x76, 0x65, 0x72, 0x74,
	0x69, 0x73, 0x69, 0x6e, 0x67, 0x5f, 0x6c, 0x61, 0x6e, 0x67, 0x75, 0x61, 0x67, 0x65, 0x5f, 0x63,
	0x6f, 0x64, 0x65, 0x18, 0x07, 0x20, 0x01, 0x28, 0x09, 0x52, 0x17, 0x61, 0x64, 0x76, 0x65, 0x72,
	0x74, 0x69, 0x73, 0x69, 0x6e, 0x67, 0x4c, 0x61, 0x6e, 0x67, 0x75, 0x61, 0x67, 0x65, 0x43, 0x6f,
	0x64, 0x65, 0x12, 0x1d, 0x0a, 0x09, 0x66, 0x69, 0x6e, 0x61, 0x6c, 0x5f, 0x75, 0x72, 0x6c, 0x18,
	0x08, 0x20, 0x01, 0x28, 0x09, 0x48, 0x00, 0x52, 0x08, 0x66, 0x69, 0x6e, 0x61, 0x6c, 0x55, 0x72,
	0x6c, 0x12, 0xae, 0x01, 0x0a, 0x25, 0x61, 0x64, 0x5f, 0x6f, 0x70, 0x74, 0x69, 0x6d, 0x69, 0x7a,
	0x65, 0x64, 0x5f, 0x62, 0x75, 0x73, 0x69, 0x6e, 0x65, 0x73, 0x73, 0x5f, 0x70, 0x72, 0x6f, 0x66,
	0x69, 0x6c, 0x65, 0x5f, 0x73, 0x65, 0x74, 0x74, 0x69, 0x6e, 0x67, 0x18, 0x09, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x5a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x61, 0x64, 0x73, 0x2e, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61, 0x64, 0x73, 0x2e, 0x76, 0x31, 0x37, 0x2e, 0x72, 0x65, 0x73,
	0x6f, 0x75, 0x72, 0x63, 0x65, 0x73, 0x2e, 0x53, 0x6d, 0x61, 0x72, 0x74, 0x43, 0x61, 0x6d, 0x70,
	0x61, 0x69, 0x67, 0x6e, 0x53, 0x65, 0x74, 0x74, 0x69, 0x6e, 0x67, 0x2e, 0x41, 0x64, 0x4f, 0x70,
	0x74, 0x69, 0x6d, 0x69, 0x7a, 0x65, 0x64, 0x42, 0x75, 0x73, 0x69, 0x6e, 0x65, 0x73, 0x73, 0x50,
	0x72, 0x6f, 0x66, 0x69, 0x6c, 0x65, 0x53, 0x65, 0x74, 0x74, 0x69, 0x6e, 0x67, 0x48, 0x00, 0x52,
	0x21, 0x61, 0x64, 0x4f, 0x70, 0x74, 0x69, 0x6d, 0x69, 0x7a, 0x65, 0x64, 0x42, 0x75, 0x73, 0x69,
	0x6e, 0x65, 0x73, 0x73, 0x50, 0x72, 0x6f, 0x66, 0x69, 0x6c, 0x65, 0x53, 0x65, 0x74, 0x74, 0x69,
	0x6e, 0x67, 0x12, 0x25, 0x0a, 0x0d, 0x62, 0x75, 0x73, 0x69, 0x6e, 0x65, 0x73, 0x73, 0x5f, 0x6e,
	0x61, 0x6d, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x48, 0x01, 0x52, 0x0c, 0x62, 0x75, 0x73,
	0x69, 0x6e, 0x65, 0x73, 0x73, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x3c, 0x0a, 0x19, 0x62, 0x75, 0x73,
	0x69, 0x6e, 0x65, 0x73, 0x73, 0x5f, 0x70, 0x72, 0x6f, 0x66, 0x69, 0x6c, 0x65, 0x5f, 0x6c, 0x6f,
	0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x09, 0x48, 0x01, 0x52, 0x17,
	0x62, 0x75, 0x73, 0x69, 0x6e, 0x65, 0x73, 0x73, 0x50, 0x72, 0x6f, 0x66, 0x69, 0x6c, 0x65, 0x4c,
	0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x1a, 0x7f, 0x0a, 0x0b, 0x50, 0x68, 0x6f, 0x6e, 0x65,
	0x4e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x12, 0x26, 0x0a, 0x0c, 0x70, 0x68, 0x6f, 0x6e, 0x65, 0x5f,
	0x6e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x48, 0x00, 0x52, 0x0b,
	0x70, 0x68, 0x6f, 0x6e, 0x65, 0x4e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x88, 0x01, 0x01, 0x12, 0x26,
	0x0a, 0x0c, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x72, 0x79, 0x5f, 0x63, 0x6f, 0x64, 0x65, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x09, 0x48, 0x01, 0x52, 0x0b, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x72, 0x79, 0x43,
	0x6f, 0x64, 0x65, 0x88, 0x01, 0x01, 0x42, 0x0f, 0x0a, 0x0d, 0x5f, 0x70, 0x68, 0x6f, 0x6e, 0x65,
	0x5f, 0x6e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x42, 0x0f, 0x0a, 0x0d, 0x5f, 0x63, 0x6f, 0x75, 0x6e,
	0x74, 0x72, 0x79, 0x5f, 0x63, 0x6f, 0x64, 0x65, 0x1a, 0x6a, 0x0a, 0x21, 0x41, 0x64, 0x4f, 0x70,
	0x74, 0x69, 0x6d, 0x69, 0x7a, 0x65, 0x64, 0x42, 0x75, 0x73, 0x69, 0x6e, 0x65, 0x73, 0x73, 0x50,
	0x72, 0x6f, 0x66, 0x69, 0x6c, 0x65, 0x53, 0x65, 0x74, 0x74, 0x69, 0x6e, 0x67, 0x12, 0x2f, 0x0a,
	0x11, 0x69, 0x6e, 0x63, 0x6c, 0x75, 0x64, 0x65, 0x5f, 0x6c, 0x65, 0x61, 0x64, 0x5f, 0x66, 0x6f,
	0x72, 0x6d, 0x18, 0x01, 0x20, 0x01, 0x28, 0x08, 0x48, 0x00, 0x52, 0x0f, 0x69, 0x6e, 0x63, 0x6c,
	0x75, 0x64, 0x65, 0x4c, 0x65, 0x61, 0x64, 0x46, 0x6f, 0x72, 0x6d, 0x88, 0x01, 0x01, 0x42, 0x14,
	0x0a, 0x12, 0x5f, 0x69, 0x6e, 0x63, 0x6c, 0x75, 0x64, 0x65, 0x5f, 0x6c, 0x65, 0x61, 0x64, 0x5f,
	0x66, 0x6f, 0x72, 0x6d, 0x3a, 0x6f, 0xea, 0x41, 0x6c, 0x0a, 0x2d, 0x67, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x61, 0x64, 0x73, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61, 0x70, 0x69, 0x73, 0x2e,
	0x63, 0x6f, 0x6d, 0x2f, 0x53, 0x6d, 0x61, 0x72, 0x74, 0x43, 0x61, 0x6d, 0x70, 0x61, 0x69, 0x67,
	0x6e, 0x53, 0x65, 0x74, 0x74, 0x69, 0x6e, 0x67, 0x12, 0x3b, 0x63, 0x75, 0x73, 0x74, 0x6f, 0x6d,
	0x65, 0x72, 0x73, 0x2f, 0x7b, 0x63, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x65, 0x72, 0x5f, 0x69, 0x64,
	0x7d, 0x2f, 0x73, 0x6d, 0x61, 0x72, 0x74, 0x43, 0x61, 0x6d, 0x70, 0x61, 0x69, 0x67, 0x6e, 0x53,
	0x65, 0x74, 0x74, 0x69, 0x6e, 0x67, 0x73, 0x2f, 0x7b, 0x63, 0x61, 0x6d, 0x70, 0x61, 0x69, 0x67,
	0x6e, 0x5f, 0x69, 0x64, 0x7d, 0x42, 0x0e, 0x0a, 0x0c, 0x6c, 0x61, 0x6e, 0x64, 0x69, 0x6e, 0x67,
	0x5f, 0x70, 0x61, 0x67, 0x65, 0x42, 0x12, 0x0a, 0x10, 0x62, 0x75, 0x73, 0x69, 0x6e, 0x65, 0x73,
	0x73, 0x5f, 0x73, 0x65, 0x74, 0x74, 0x69, 0x6e, 0x67, 0x42, 0x8b, 0x02, 0x0a, 0x26, 0x63, 0x6f,
	0x6d, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x61, 0x64, 0x73, 0x2e, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x61, 0x64, 0x73, 0x2e, 0x76, 0x31, 0x37, 0x2e, 0x72, 0x65, 0x73, 0x6f, 0x75,
	0x72, 0x63, 0x65, 0x73, 0x42, 0x19, 0x53, 0x6d, 0x61, 0x72, 0x74, 0x43, 0x61, 0x6d, 0x70, 0x61,
	0x69, 0x67, 0x6e, 0x53, 0x65, 0x74, 0x74, 0x69, 0x6e, 0x67, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50,
	0x01, 0x5a, 0x4b, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x67, 0x6f, 0x6c, 0x61, 0x6e, 0x67,
	0x2e, 0x6f, 0x72, 0x67, 0x2f, 0x67, 0x65, 0x6e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x67, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x61, 0x70, 0x69, 0x73, 0x2f, 0x61, 0x64, 0x73, 0x2f, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x61, 0x64, 0x73, 0x2f, 0x76, 0x31, 0x37, 0x2f, 0x72, 0x65, 0x73, 0x6f, 0x75,
	0x72, 0x63, 0x65, 0x73, 0x3b, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x73, 0xa2, 0x02,
	0x03, 0x47, 0x41, 0x41, 0xaa, 0x02, 0x22, 0x47, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x41, 0x64,
	0x73, 0x2e, 0x47, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x41, 0x64, 0x73, 0x2e, 0x56, 0x31, 0x37, 0x2e,
	0x52, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x73, 0xca, 0x02, 0x22, 0x47, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x5c, 0x41, 0x64, 0x73, 0x5c, 0x47, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x41, 0x64, 0x73,
	0x5c, 0x56, 0x31, 0x37, 0x5c, 0x52, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x73, 0xea, 0x02,
	0x26, 0x47, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x3a, 0x3a, 0x41, 0x64, 0x73, 0x3a, 0x3a, 0x47, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x41, 0x64, 0x73, 0x3a, 0x3a, 0x56, 0x31, 0x37, 0x3a, 0x3a, 0x52, 0x65,
	0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x73, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_resources_smart_campaign_setting_proto_rawDescOnce sync.Once
	file_resources_smart_campaign_setting_proto_rawDescData = file_resources_smart_campaign_setting_proto_rawDesc
)

func file_resources_smart_campaign_setting_proto_rawDescGZIP() []byte {
	file_resources_smart_campaign_setting_proto_rawDescOnce.Do(func() {
		file_resources_smart_campaign_setting_proto_rawDescData = protoimpl.X.CompressGZIP(file_resources_smart_campaign_setting_proto_rawDescData)
	})
	return file_resources_smart_campaign_setting_proto_rawDescData
}

var file_resources_smart_campaign_setting_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_resources_smart_campaign_setting_proto_goTypes = []interface{}{
	(*SmartCampaignSetting)(nil),                                   // 0: google.ads.googleads.v17.resources.SmartCampaignSetting
	(*SmartCampaignSetting_PhoneNumber)(nil),                       // 1: google.ads.googleads.v17.resources.SmartCampaignSetting.PhoneNumber
	(*SmartCampaignSetting_AdOptimizedBusinessProfileSetting)(nil), // 2: google.ads.googleads.v17.resources.SmartCampaignSetting.AdOptimizedBusinessProfileSetting
}
var file_resources_smart_campaign_setting_proto_depIdxs = []int32{
	1, // 0: google.ads.googleads.v17.resources.SmartCampaignSetting.phone_number:type_name -> google.ads.googleads.v17.resources.SmartCampaignSetting.PhoneNumber
	2, // 1: google.ads.googleads.v17.resources.SmartCampaignSetting.ad_optimized_business_profile_setting:type_name -> google.ads.googleads.v17.resources.SmartCampaignSetting.AdOptimizedBusinessProfileSetting
	2, // [2:2] is the sub-list for method output_type
	2, // [2:2] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_resources_smart_campaign_setting_proto_init() }
func file_resources_smart_campaign_setting_proto_init() {
	if File_resources_smart_campaign_setting_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_resources_smart_campaign_setting_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SmartCampaignSetting); i {
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
		file_resources_smart_campaign_setting_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SmartCampaignSetting_PhoneNumber); i {
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
		file_resources_smart_campaign_setting_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SmartCampaignSetting_AdOptimizedBusinessProfileSetting); i {
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
	file_resources_smart_campaign_setting_proto_msgTypes[0].OneofWrappers = []interface{}{
		(*SmartCampaignSetting_FinalUrl)(nil),
		(*SmartCampaignSetting_AdOptimizedBusinessProfileSetting_)(nil),
		(*SmartCampaignSetting_BusinessName)(nil),
		(*SmartCampaignSetting_BusinessProfileLocation)(nil),
	}
	file_resources_smart_campaign_setting_proto_msgTypes[1].OneofWrappers = []interface{}{}
	file_resources_smart_campaign_setting_proto_msgTypes[2].OneofWrappers = []interface{}{}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_resources_smart_campaign_setting_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_resources_smart_campaign_setting_proto_goTypes,
		DependencyIndexes: file_resources_smart_campaign_setting_proto_depIdxs,
		MessageInfos:      file_resources_smart_campaign_setting_proto_msgTypes,
	}.Build()
	File_resources_smart_campaign_setting_proto = out.File
	file_resources_smart_campaign_setting_proto_rawDesc = nil
	file_resources_smart_campaign_setting_proto_goTypes = nil
	file_resources_smart_campaign_setting_proto_depIdxs = nil
}
