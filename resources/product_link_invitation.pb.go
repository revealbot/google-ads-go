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
// source: google/ads/googleads/v16/resources/product_link_invitation.proto

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

// Represents an invitation for data sharing connection between a Google Ads
// account and another account.
type ProductLinkInvitation struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Immutable. The resource name of a product link invitation.
	// Product link invitation resource names have the form:
	//
	// `customers/{customer_id}/productLinkInvitations/{product_link_invitation_id}`
	ResourceName string `protobuf:"bytes,1,opt,name=resource_name,json=resourceName,proto3" json:"resource_name,omitempty"`
	// Output only. The ID of the product link invitation.
	// This field is read only.
	ProductLinkInvitationId int64 `protobuf:"varint,2,opt,name=product_link_invitation_id,json=productLinkInvitationId,proto3" json:"product_link_invitation_id,omitempty"`
	// Output only. The status of the product link invitation.
	// This field is read only.
	Status enums.ProductLinkInvitationStatusEnum_ProductLinkInvitationStatus `protobuf:"varint,3,opt,name=status,proto3,enum=google.ads.googleads.v16.enums.ProductLinkInvitationStatusEnum_ProductLinkInvitationStatus" json:"status,omitempty"`
	// Output only. The type of the invited account.
	// This field is read only and can be used for filtering invitations with
	// {@code GoogleAdsService.SearchGoogleAdsRequest}.
	Type enums.LinkedProductTypeEnum_LinkedProductType `protobuf:"varint,6,opt,name=type,proto3,enum=google.ads.googleads.v16.enums.LinkedProductTypeEnum_LinkedProductType" json:"type,omitempty"`
	// An account invited to link to this Google Ads account.
	//
	// Types that are assignable to InvitedAccount:
	//
	//	*ProductLinkInvitation_HotelCenter
	//	*ProductLinkInvitation_MerchantCenter
	//	*ProductLinkInvitation_AdvertisingPartner
	InvitedAccount isProductLinkInvitation_InvitedAccount `protobuf_oneof:"invited_account"`
}

func (x *ProductLinkInvitation) Reset() {
	*x = ProductLinkInvitation{}
	if protoimpl.UnsafeEnabled {
		mi := &file_resources_product_link_invitation_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ProductLinkInvitation) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ProductLinkInvitation) ProtoMessage() {}

func (x *ProductLinkInvitation) ProtoReflect() protoreflect.Message {
	mi := &file_resources_product_link_invitation_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ProductLinkInvitation.ProtoReflect.Descriptor instead.
func (*ProductLinkInvitation) Descriptor() ([]byte, []int) {
	return file_resources_product_link_invitation_proto_rawDescGZIP(), []int{0}
}

func (x *ProductLinkInvitation) GetResourceName() string {
	if x != nil {
		return x.ResourceName
	}
	return ""
}

func (x *ProductLinkInvitation) GetProductLinkInvitationId() int64 {
	if x != nil {
		return x.ProductLinkInvitationId
	}
	return 0
}

func (x *ProductLinkInvitation) GetStatus() enums.ProductLinkInvitationStatusEnum_ProductLinkInvitationStatus {
	if x != nil {
		return x.Status
	}
	return enums.ProductLinkInvitationStatusEnum_ProductLinkInvitationStatus(0)
}

func (x *ProductLinkInvitation) GetType() enums.LinkedProductTypeEnum_LinkedProductType {
	if x != nil {
		return x.Type
	}
	return enums.LinkedProductTypeEnum_LinkedProductType(0)
}

func (m *ProductLinkInvitation) GetInvitedAccount() isProductLinkInvitation_InvitedAccount {
	if m != nil {
		return m.InvitedAccount
	}
	return nil
}

func (x *ProductLinkInvitation) GetHotelCenter() *HotelCenterLinkInvitationIdentifier {
	if x, ok := x.GetInvitedAccount().(*ProductLinkInvitation_HotelCenter); ok {
		return x.HotelCenter
	}
	return nil
}

func (x *ProductLinkInvitation) GetMerchantCenter() *MerchantCenterLinkInvitationIdentifier {
	if x, ok := x.GetInvitedAccount().(*ProductLinkInvitation_MerchantCenter); ok {
		return x.MerchantCenter
	}
	return nil
}

func (x *ProductLinkInvitation) GetAdvertisingPartner() *AdvertisingPartnerLinkInvitationIdentifier {
	if x, ok := x.GetInvitedAccount().(*ProductLinkInvitation_AdvertisingPartner); ok {
		return x.AdvertisingPartner
	}
	return nil
}

type isProductLinkInvitation_InvitedAccount interface {
	isProductLinkInvitation_InvitedAccount()
}

type ProductLinkInvitation_HotelCenter struct {
	// Output only. Hotel link invitation.
	HotelCenter *HotelCenterLinkInvitationIdentifier `protobuf:"bytes,4,opt,name=hotel_center,json=hotelCenter,proto3,oneof"`
}

type ProductLinkInvitation_MerchantCenter struct {
	// Output only. Merchant Center link invitation.
	MerchantCenter *MerchantCenterLinkInvitationIdentifier `protobuf:"bytes,5,opt,name=merchant_center,json=merchantCenter,proto3,oneof"`
}

type ProductLinkInvitation_AdvertisingPartner struct {
	// Output only. Advertising Partner link invitation.
	AdvertisingPartner *AdvertisingPartnerLinkInvitationIdentifier `protobuf:"bytes,7,opt,name=advertising_partner,json=advertisingPartner,proto3,oneof"`
}

func (*ProductLinkInvitation_HotelCenter) isProductLinkInvitation_InvitedAccount() {}

func (*ProductLinkInvitation_MerchantCenter) isProductLinkInvitation_InvitedAccount() {}

func (*ProductLinkInvitation_AdvertisingPartner) isProductLinkInvitation_InvitedAccount() {}

// The identifier for Hotel account.
type HotelCenterLinkInvitationIdentifier struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Output only. The hotel center id of the hotel account.
	// This field is read only
	HotelCenterId int64 `protobuf:"varint,1,opt,name=hotel_center_id,json=hotelCenterId,proto3" json:"hotel_center_id,omitempty"`
}

func (x *HotelCenterLinkInvitationIdentifier) Reset() {
	*x = HotelCenterLinkInvitationIdentifier{}
	if protoimpl.UnsafeEnabled {
		mi := &file_resources_product_link_invitation_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *HotelCenterLinkInvitationIdentifier) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*HotelCenterLinkInvitationIdentifier) ProtoMessage() {}

func (x *HotelCenterLinkInvitationIdentifier) ProtoReflect() protoreflect.Message {
	mi := &file_resources_product_link_invitation_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use HotelCenterLinkInvitationIdentifier.ProtoReflect.Descriptor instead.
func (*HotelCenterLinkInvitationIdentifier) Descriptor() ([]byte, []int) {
	return file_resources_product_link_invitation_proto_rawDescGZIP(), []int{1}
}

func (x *HotelCenterLinkInvitationIdentifier) GetHotelCenterId() int64 {
	if x != nil {
		return x.HotelCenterId
	}
	return 0
}

// The identifier for Merchant Center Account.
type MerchantCenterLinkInvitationIdentifier struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Output only. The Merchant Center id of the Merchant account.
	// This field is read only
	MerchantCenterId int64 `protobuf:"varint,1,opt,name=merchant_center_id,json=merchantCenterId,proto3" json:"merchant_center_id,omitempty"`
}

func (x *MerchantCenterLinkInvitationIdentifier) Reset() {
	*x = MerchantCenterLinkInvitationIdentifier{}
	if protoimpl.UnsafeEnabled {
		mi := &file_resources_product_link_invitation_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MerchantCenterLinkInvitationIdentifier) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MerchantCenterLinkInvitationIdentifier) ProtoMessage() {}

func (x *MerchantCenterLinkInvitationIdentifier) ProtoReflect() protoreflect.Message {
	mi := &file_resources_product_link_invitation_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MerchantCenterLinkInvitationIdentifier.ProtoReflect.Descriptor instead.
func (*MerchantCenterLinkInvitationIdentifier) Descriptor() ([]byte, []int) {
	return file_resources_product_link_invitation_proto_rawDescGZIP(), []int{2}
}

func (x *MerchantCenterLinkInvitationIdentifier) GetMerchantCenterId() int64 {
	if x != nil {
		return x.MerchantCenterId
	}
	return 0
}

// The identifier for the Advertising Partner Google Ads account.
type AdvertisingPartnerLinkInvitationIdentifier struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Immutable. The resource name of the advertising partner Google Ads account.
	// This field is read only.
	Customer *string `protobuf:"bytes,1,opt,name=customer,proto3,oneof" json:"customer,omitempty"`
}

func (x *AdvertisingPartnerLinkInvitationIdentifier) Reset() {
	*x = AdvertisingPartnerLinkInvitationIdentifier{}
	if protoimpl.UnsafeEnabled {
		mi := &file_resources_product_link_invitation_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AdvertisingPartnerLinkInvitationIdentifier) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AdvertisingPartnerLinkInvitationIdentifier) ProtoMessage() {}

func (x *AdvertisingPartnerLinkInvitationIdentifier) ProtoReflect() protoreflect.Message {
	mi := &file_resources_product_link_invitation_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AdvertisingPartnerLinkInvitationIdentifier.ProtoReflect.Descriptor instead.
func (*AdvertisingPartnerLinkInvitationIdentifier) Descriptor() ([]byte, []int) {
	return file_resources_product_link_invitation_proto_rawDescGZIP(), []int{3}
}

func (x *AdvertisingPartnerLinkInvitationIdentifier) GetCustomer() string {
	if x != nil && x.Customer != nil {
		return *x.Customer
	}
	return ""
}

var File_resources_product_link_invitation_proto protoreflect.FileDescriptor

var file_resources_product_link_invitation_proto_rawDesc = []byte{
	0x0a, 0x40, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x64, 0x73, 0x2f, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x61, 0x64, 0x73, 0x2f, 0x76, 0x31, 0x36, 0x2f, 0x72, 0x65, 0x73, 0x6f, 0x75,
	0x72, 0x63, 0x65, 0x73, 0x2f, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x5f, 0x6c, 0x69, 0x6e,
	0x6b, 0x5f, 0x69, 0x6e, 0x76, 0x69, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x12, 0x22, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x61, 0x64, 0x73, 0x2e, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61, 0x64, 0x73, 0x2e, 0x76, 0x31, 0x36, 0x2e, 0x72, 0x65, 0x73,
	0x6f, 0x75, 0x72, 0x63, 0x65, 0x73, 0x1a, 0x38, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61,
	0x64, 0x73, 0x2f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61, 0x64, 0x73, 0x2f, 0x76, 0x31, 0x36,
	0x2f, 0x65, 0x6e, 0x75, 0x6d, 0x73, 0x2f, 0x6c, 0x69, 0x6e, 0x6b, 0x65, 0x64, 0x5f, 0x70, 0x72,
	0x6f, 0x64, 0x75, 0x63, 0x74, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x1a, 0x43, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x64, 0x73, 0x2f, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x61, 0x64, 0x73, 0x2f, 0x76, 0x31, 0x36, 0x2f, 0x65, 0x6e, 0x75, 0x6d, 0x73,
	0x2f, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x5f, 0x6c, 0x69, 0x6e, 0x6b, 0x5f, 0x69, 0x6e,
	0x76, 0x69, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70,
	0x69, 0x2f, 0x66, 0x69, 0x65, 0x6c, 0x64, 0x5f, 0x62, 0x65, 0x68, 0x61, 0x76, 0x69, 0x6f, 0x72,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x19, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61,
	0x70, 0x69, 0x2f, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x22, 0x9b, 0x07, 0x0a, 0x15, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x4c, 0x69, 0x6e,
	0x6b, 0x49, 0x6e, 0x76, 0x69, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x5b, 0x0a, 0x0d, 0x72,
	0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x42, 0x36, 0xe0, 0x41, 0x05, 0xfa, 0x41, 0x30, 0x0a, 0x2e, 0x67, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x61, 0x64, 0x73, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61, 0x70, 0x69, 0x73,
	0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x4c, 0x69, 0x6e, 0x6b,
	0x49, 0x6e, 0x76, 0x69, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x0c, 0x72, 0x65, 0x73, 0x6f,
	0x75, 0x72, 0x63, 0x65, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x40, 0x0a, 0x1a, 0x70, 0x72, 0x6f, 0x64,
	0x75, 0x63, 0x74, 0x5f, 0x6c, 0x69, 0x6e, 0x6b, 0x5f, 0x69, 0x6e, 0x76, 0x69, 0x74, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x42, 0x03, 0xe0, 0x41,
	0x03, 0x52, 0x17, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x4c, 0x69, 0x6e, 0x6b, 0x49, 0x6e,
	0x76, 0x69, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x49, 0x64, 0x12, 0x78, 0x0a, 0x06, 0x73, 0x74,
	0x61, 0x74, 0x75, 0x73, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x5b, 0x2e, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x2e, 0x61, 0x64, 0x73, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61, 0x64,
	0x73, 0x2e, 0x76, 0x31, 0x36, 0x2e, 0x65, 0x6e, 0x75, 0x6d, 0x73, 0x2e, 0x50, 0x72, 0x6f, 0x64,
	0x75, 0x63, 0x74, 0x4c, 0x69, 0x6e, 0x6b, 0x49, 0x6e, 0x76, 0x69, 0x74, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x45, 0x6e, 0x75, 0x6d, 0x2e, 0x50, 0x72, 0x6f, 0x64,
	0x75, 0x63, 0x74, 0x4c, 0x69, 0x6e, 0x6b, 0x49, 0x6e, 0x76, 0x69, 0x74, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x42, 0x03, 0xe0, 0x41, 0x03, 0x52, 0x06, 0x73, 0x74,
	0x61, 0x74, 0x75, 0x73, 0x12, 0x60, 0x0a, 0x04, 0x74, 0x79, 0x70, 0x65, 0x18, 0x06, 0x20, 0x01,
	0x28, 0x0e, 0x32, 0x47, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x61, 0x64, 0x73, 0x2e,
	0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61, 0x64, 0x73, 0x2e, 0x76, 0x31, 0x36, 0x2e, 0x65, 0x6e,
	0x75, 0x6d, 0x73, 0x2e, 0x4c, 0x69, 0x6e, 0x6b, 0x65, 0x64, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63,
	0x74, 0x54, 0x79, 0x70, 0x65, 0x45, 0x6e, 0x75, 0x6d, 0x2e, 0x4c, 0x69, 0x6e, 0x6b, 0x65, 0x64,
	0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x54, 0x79, 0x70, 0x65, 0x42, 0x03, 0xe0, 0x41, 0x03,
	0x52, 0x04, 0x74, 0x79, 0x70, 0x65, 0x12, 0x71, 0x0a, 0x0c, 0x68, 0x6f, 0x74, 0x65, 0x6c, 0x5f,
	0x63, 0x65, 0x6e, 0x74, 0x65, 0x72, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x47, 0x2e, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x61, 0x64, 0x73, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x61, 0x64, 0x73, 0x2e, 0x76, 0x31, 0x36, 0x2e, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65,
	0x73, 0x2e, 0x48, 0x6f, 0x74, 0x65, 0x6c, 0x43, 0x65, 0x6e, 0x74, 0x65, 0x72, 0x4c, 0x69, 0x6e,
	0x6b, 0x49, 0x6e, 0x76, 0x69, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x49, 0x64, 0x65, 0x6e, 0x74,
	0x69, 0x66, 0x69, 0x65, 0x72, 0x42, 0x03, 0xe0, 0x41, 0x03, 0x48, 0x00, 0x52, 0x0b, 0x68, 0x6f,
	0x74, 0x65, 0x6c, 0x43, 0x65, 0x6e, 0x74, 0x65, 0x72, 0x12, 0x7a, 0x0a, 0x0f, 0x6d, 0x65, 0x72,
	0x63, 0x68, 0x61, 0x6e, 0x74, 0x5f, 0x63, 0x65, 0x6e, 0x74, 0x65, 0x72, 0x18, 0x05, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x4a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x61, 0x64, 0x73, 0x2e,
	0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61, 0x64, 0x73, 0x2e, 0x76, 0x31, 0x36, 0x2e, 0x72, 0x65,
	0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x73, 0x2e, 0x4d, 0x65, 0x72, 0x63, 0x68, 0x61, 0x6e, 0x74,
	0x43, 0x65, 0x6e, 0x74, 0x65, 0x72, 0x4c, 0x69, 0x6e, 0x6b, 0x49, 0x6e, 0x76, 0x69, 0x74, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x49, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x66, 0x69, 0x65, 0x72, 0x42, 0x03,
	0xe0, 0x41, 0x03, 0x48, 0x00, 0x52, 0x0e, 0x6d, 0x65, 0x72, 0x63, 0x68, 0x61, 0x6e, 0x74, 0x43,
	0x65, 0x6e, 0x74, 0x65, 0x72, 0x12, 0x86, 0x01, 0x0a, 0x13, 0x61, 0x64, 0x76, 0x65, 0x72, 0x74,
	0x69, 0x73, 0x69, 0x6e, 0x67, 0x5f, 0x70, 0x61, 0x72, 0x74, 0x6e, 0x65, 0x72, 0x18, 0x07, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x4e, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x61, 0x64, 0x73,
	0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61, 0x64, 0x73, 0x2e, 0x76, 0x31, 0x36, 0x2e, 0x72,
	0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x73, 0x2e, 0x41, 0x64, 0x76, 0x65, 0x72, 0x74, 0x69,
	0x73, 0x69, 0x6e, 0x67, 0x50, 0x61, 0x72, 0x74, 0x6e, 0x65, 0x72, 0x4c, 0x69, 0x6e, 0x6b, 0x49,
	0x6e, 0x76, 0x69, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x49, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x66,
	0x69, 0x65, 0x72, 0x42, 0x03, 0xe0, 0x41, 0x03, 0x48, 0x00, 0x52, 0x12, 0x61, 0x64, 0x76, 0x65,
	0x72, 0x74, 0x69, 0x73, 0x69, 0x6e, 0x67, 0x50, 0x61, 0x72, 0x74, 0x6e, 0x65, 0x72, 0x3a, 0x7c,
	0xea, 0x41, 0x79, 0x0a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61, 0x64, 0x73, 0x2e, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61, 0x70, 0x69, 0x73, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x50, 0x72,
	0x6f, 0x64, 0x75, 0x63, 0x74, 0x4c, 0x69, 0x6e, 0x6b, 0x49, 0x6e, 0x76, 0x69, 0x74, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x12, 0x47, 0x63, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x65, 0x72, 0x73, 0x2f, 0x7b,
	0x63, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x7d, 0x2f, 0x70, 0x72, 0x6f,
	0x64, 0x75, 0x63, 0x74, 0x4c, 0x69, 0x6e, 0x6b, 0x49, 0x6e, 0x76, 0x69, 0x74, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x73, 0x2f, 0x7b, 0x63, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x65, 0x72, 0x5f, 0x69, 0x6e,
	0x76, 0x69, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x69, 0x64, 0x7d, 0x42, 0x11, 0x0a, 0x0f,
	0x69, 0x6e, 0x76, 0x69, 0x74, 0x65, 0x64, 0x5f, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x22,
	0x52, 0x0a, 0x23, 0x48, 0x6f, 0x74, 0x65, 0x6c, 0x43, 0x65, 0x6e, 0x74, 0x65, 0x72, 0x4c, 0x69,
	0x6e, 0x6b, 0x49, 0x6e, 0x76, 0x69, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x49, 0x64, 0x65, 0x6e,
	0x74, 0x69, 0x66, 0x69, 0x65, 0x72, 0x12, 0x2b, 0x0a, 0x0f, 0x68, 0x6f, 0x74, 0x65, 0x6c, 0x5f,
	0x63, 0x65, 0x6e, 0x74, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x42,
	0x03, 0xe0, 0x41, 0x03, 0x52, 0x0d, 0x68, 0x6f, 0x74, 0x65, 0x6c, 0x43, 0x65, 0x6e, 0x74, 0x65,
	0x72, 0x49, 0x64, 0x22, 0x5b, 0x0a, 0x26, 0x4d, 0x65, 0x72, 0x63, 0x68, 0x61, 0x6e, 0x74, 0x43,
	0x65, 0x6e, 0x74, 0x65, 0x72, 0x4c, 0x69, 0x6e, 0x6b, 0x49, 0x6e, 0x76, 0x69, 0x74, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x49, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x66, 0x69, 0x65, 0x72, 0x12, 0x31, 0x0a,
	0x12, 0x6d, 0x65, 0x72, 0x63, 0x68, 0x61, 0x6e, 0x74, 0x5f, 0x63, 0x65, 0x6e, 0x74, 0x65, 0x72,
	0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x42, 0x03, 0xe0, 0x41, 0x03, 0x52, 0x10,
	0x6d, 0x65, 0x72, 0x63, 0x68, 0x61, 0x6e, 0x74, 0x43, 0x65, 0x6e, 0x74, 0x65, 0x72, 0x49, 0x64,
	0x22, 0x85, 0x01, 0x0a, 0x2a, 0x41, 0x64, 0x76, 0x65, 0x72, 0x74, 0x69, 0x73, 0x69, 0x6e, 0x67,
	0x50, 0x61, 0x72, 0x74, 0x6e, 0x65, 0x72, 0x4c, 0x69, 0x6e, 0x6b, 0x49, 0x6e, 0x76, 0x69, 0x74,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x49, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x66, 0x69, 0x65, 0x72, 0x12,
	0x4a, 0x0a, 0x08, 0x63, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x65, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x42, 0x29, 0xe0, 0x41, 0x05, 0xfa, 0x41, 0x23, 0x0a, 0x21, 0x67, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x61, 0x64, 0x73, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61, 0x70, 0x69, 0x73, 0x2e,
	0x63, 0x6f, 0x6d, 0x2f, 0x43, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x65, 0x72, 0x48, 0x00, 0x52, 0x08,
	0x63, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x65, 0x72, 0x88, 0x01, 0x01, 0x42, 0x0b, 0x0a, 0x09, 0x5f,
	0x63, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x65, 0x72, 0x42, 0x8c, 0x02, 0x0a, 0x26, 0x63, 0x6f, 0x6d,
	0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x61, 0x64, 0x73, 0x2e, 0x67, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x61, 0x64, 0x73, 0x2e, 0x76, 0x31, 0x36, 0x2e, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72,
	0x63, 0x65, 0x73, 0x42, 0x1a, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x4c, 0x69, 0x6e, 0x6b,
	0x49, 0x6e, 0x76, 0x69, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50,
	0x01, 0x5a, 0x4b, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x67, 0x6f, 0x6c, 0x61, 0x6e, 0x67,
	0x2e, 0x6f, 0x72, 0x67, 0x2f, 0x67, 0x65, 0x6e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x67, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x61, 0x70, 0x69, 0x73, 0x2f, 0x61, 0x64, 0x73, 0x2f, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x61, 0x64, 0x73, 0x2f, 0x76, 0x31, 0x36, 0x2f, 0x72, 0x65, 0x73, 0x6f, 0x75,
	0x72, 0x63, 0x65, 0x73, 0x3b, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x73, 0xa2, 0x02,
	0x03, 0x47, 0x41, 0x41, 0xaa, 0x02, 0x22, 0x47, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x41, 0x64,
	0x73, 0x2e, 0x47, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x41, 0x64, 0x73, 0x2e, 0x56, 0x31, 0x36, 0x2e,
	0x52, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x73, 0xca, 0x02, 0x22, 0x47, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x5c, 0x41, 0x64, 0x73, 0x5c, 0x47, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x41, 0x64, 0x73,
	0x5c, 0x56, 0x31, 0x36, 0x5c, 0x52, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x73, 0xea, 0x02,
	0x26, 0x47, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x3a, 0x3a, 0x41, 0x64, 0x73, 0x3a, 0x3a, 0x47, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x41, 0x64, 0x73, 0x3a, 0x3a, 0x56, 0x31, 0x36, 0x3a, 0x3a, 0x52, 0x65,
	0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x73, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_resources_product_link_invitation_proto_rawDescOnce sync.Once
	file_resources_product_link_invitation_proto_rawDescData = file_resources_product_link_invitation_proto_rawDesc
)

func file_resources_product_link_invitation_proto_rawDescGZIP() []byte {
	file_resources_product_link_invitation_proto_rawDescOnce.Do(func() {
		file_resources_product_link_invitation_proto_rawDescData = protoimpl.X.CompressGZIP(file_resources_product_link_invitation_proto_rawDescData)
	})
	return file_resources_product_link_invitation_proto_rawDescData
}

var file_resources_product_link_invitation_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_resources_product_link_invitation_proto_goTypes = []interface{}{
	(*ProductLinkInvitation)(nil),                                          // 0: google.ads.googleads.v16.resources.ProductLinkInvitation
	(*HotelCenterLinkInvitationIdentifier)(nil),                            // 1: google.ads.googleads.v16.resources.HotelCenterLinkInvitationIdentifier
	(*MerchantCenterLinkInvitationIdentifier)(nil),                         // 2: google.ads.googleads.v16.resources.MerchantCenterLinkInvitationIdentifier
	(*AdvertisingPartnerLinkInvitationIdentifier)(nil),                     // 3: google.ads.googleads.v16.resources.AdvertisingPartnerLinkInvitationIdentifier
	(enums.ProductLinkInvitationStatusEnum_ProductLinkInvitationStatus)(0), // 4: google.ads.googleads.v16.enums.ProductLinkInvitationStatusEnum.ProductLinkInvitationStatus
	(enums.LinkedProductTypeEnum_LinkedProductType)(0),                     // 5: google.ads.googleads.v16.enums.LinkedProductTypeEnum.LinkedProductType
}
var file_resources_product_link_invitation_proto_depIdxs = []int32{
	4, // 0: google.ads.googleads.v16.resources.ProductLinkInvitation.status:type_name -> google.ads.googleads.v16.enums.ProductLinkInvitationStatusEnum.ProductLinkInvitationStatus
	5, // 1: google.ads.googleads.v16.resources.ProductLinkInvitation.type:type_name -> google.ads.googleads.v16.enums.LinkedProductTypeEnum.LinkedProductType
	1, // 2: google.ads.googleads.v16.resources.ProductLinkInvitation.hotel_center:type_name -> google.ads.googleads.v16.resources.HotelCenterLinkInvitationIdentifier
	2, // 3: google.ads.googleads.v16.resources.ProductLinkInvitation.merchant_center:type_name -> google.ads.googleads.v16.resources.MerchantCenterLinkInvitationIdentifier
	3, // 4: google.ads.googleads.v16.resources.ProductLinkInvitation.advertising_partner:type_name -> google.ads.googleads.v16.resources.AdvertisingPartnerLinkInvitationIdentifier
	5, // [5:5] is the sub-list for method output_type
	5, // [5:5] is the sub-list for method input_type
	5, // [5:5] is the sub-list for extension type_name
	5, // [5:5] is the sub-list for extension extendee
	0, // [0:5] is the sub-list for field type_name
}

func init() { file_resources_product_link_invitation_proto_init() }
func file_resources_product_link_invitation_proto_init() {
	if File_resources_product_link_invitation_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_resources_product_link_invitation_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ProductLinkInvitation); i {
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
		file_resources_product_link_invitation_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*HotelCenterLinkInvitationIdentifier); i {
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
		file_resources_product_link_invitation_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*MerchantCenterLinkInvitationIdentifier); i {
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
		file_resources_product_link_invitation_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AdvertisingPartnerLinkInvitationIdentifier); i {
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
	file_resources_product_link_invitation_proto_msgTypes[0].OneofWrappers = []interface{}{
		(*ProductLinkInvitation_HotelCenter)(nil),
		(*ProductLinkInvitation_MerchantCenter)(nil),
		(*ProductLinkInvitation_AdvertisingPartner)(nil),
	}
	file_resources_product_link_invitation_proto_msgTypes[3].OneofWrappers = []interface{}{}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_resources_product_link_invitation_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_resources_product_link_invitation_proto_goTypes,
		DependencyIndexes: file_resources_product_link_invitation_proto_depIdxs,
		MessageInfos:      file_resources_product_link_invitation_proto_msgTypes,
	}.Build()
	File_resources_product_link_invitation_proto = out.File
	file_resources_product_link_invitation_proto_rawDesc = nil
	file_resources_product_link_invitation_proto_goTypes = nil
	file_resources_product_link_invitation_proto_depIdxs = nil
}
