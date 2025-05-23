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
// source: google/ads/googleads/v19/resources/hotel_reconciliation.proto

package resources

import (
	enums "github.com/revealbot/google-ads-go/enums"
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

// A hotel reconciliation. It contains conversion information from Hotel
// bookings to reconcile with advertiser records. These rows may be updated
// or canceled before billing through Bulk Uploads.
type HotelReconciliation struct {
	state protoimpl.MessageState `protogen:"open.v1"`
	// Immutable. The resource name of the hotel reconciliation.
	// Hotel reconciliation resource names have the form:
	//
	// `customers/{customer_id}/hotelReconciliations/{commission_id}`
	ResourceName string `protobuf:"bytes,1,opt,name=resource_name,json=resourceName,proto3" json:"resource_name,omitempty"`
	// Required. Output only. The commission ID is Google's ID for this booking.
	// Every booking event is assigned a Commission ID to help you match it to a
	// guest stay.
	CommissionId string `protobuf:"bytes,2,opt,name=commission_id,json=commissionId,proto3" json:"commission_id,omitempty"`
	// Output only. The order ID is the identifier for this booking as provided in
	// the 'transaction_id' parameter of the conversion tracking tag.
	OrderId string `protobuf:"bytes,3,opt,name=order_id,json=orderId,proto3" json:"order_id,omitempty"`
	// Output only. The resource name for the Campaign associated with the
	// conversion.
	Campaign string `protobuf:"bytes,11,opt,name=campaign,proto3" json:"campaign,omitempty"`
	// Output only. Identifier for the Hotel Center account which provides the
	// rates for the Hotel campaign.
	HotelCenterId int64 `protobuf:"varint,4,opt,name=hotel_center_id,json=hotelCenterId,proto3" json:"hotel_center_id,omitempty"`
	// Output only. Unique identifier for the booked property, as provided in the
	// Hotel Center feed. The hotel ID comes from the 'ID' parameter of the
	// conversion tracking tag.
	HotelId string `protobuf:"bytes,5,opt,name=hotel_id,json=hotelId,proto3" json:"hotel_id,omitempty"`
	// Output only. Check-in date recorded when the booking is made. If the
	// check-in date is modified at reconciliation, the revised date will then
	// take the place of the original date in this column. Format is YYYY-MM-DD.
	CheckInDate string `protobuf:"bytes,6,opt,name=check_in_date,json=checkInDate,proto3" json:"check_in_date,omitempty"`
	// Output only. Check-out date recorded when the booking is made. If the
	// check-in date is modified at reconciliation, the revised date will then
	// take the place of the original date in this column. Format is YYYY-MM-DD.
	CheckOutDate string `protobuf:"bytes,7,opt,name=check_out_date,json=checkOutDate,proto3" json:"check_out_date,omitempty"`
	// Required. Output only. Reconciled value is the final value of a booking as
	// paid by the guest. If original booking value changes for any reason, such
	// as itinerary changes or room upsells, the reconciled value should be the
	// full final amount collected. If a booking is canceled, the reconciled value
	// should include the value of any cancellation fees or non-refundable nights
	// charged. Value is in millionths of the base unit currency. For example,
	// $12.35 would be represented as 12350000. Currency unit is in the default
	// customer currency.
	ReconciledValueMicros int64 `protobuf:"varint,8,opt,name=reconciled_value_micros,json=reconciledValueMicros,proto3" json:"reconciled_value_micros,omitempty"`
	// Output only. Whether a given booking has been billed. Once billed, a
	// booking can't be modified.
	Billed bool `protobuf:"varint,9,opt,name=billed,proto3" json:"billed,omitempty"`
	// Required. Output only. Current status of a booking with regards to
	// reconciliation and billing. Bookings should be reconciled within 45 days
	// after the check-out date. Any booking not reconciled within 45 days will be
	// billed at its original value.
	Status        enums.HotelReconciliationStatusEnum_HotelReconciliationStatus `protobuf:"varint,10,opt,name=status,proto3,enum=google.ads.googleads.v19.enums.HotelReconciliationStatusEnum_HotelReconciliationStatus" json:"status,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *HotelReconciliation) Reset() {
	*x = HotelReconciliation{}
	mi := &file_resources_hotel_reconciliation_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *HotelReconciliation) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*HotelReconciliation) ProtoMessage() {}

func (x *HotelReconciliation) ProtoReflect() protoreflect.Message {
	mi := &file_resources_hotel_reconciliation_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use HotelReconciliation.ProtoReflect.Descriptor instead.
func (*HotelReconciliation) Descriptor() ([]byte, []int) {
	return file_resources_hotel_reconciliation_proto_rawDescGZIP(), []int{0}
}

func (x *HotelReconciliation) GetResourceName() string {
	if x != nil {
		return x.ResourceName
	}
	return ""
}

func (x *HotelReconciliation) GetCommissionId() string {
	if x != nil {
		return x.CommissionId
	}
	return ""
}

func (x *HotelReconciliation) GetOrderId() string {
	if x != nil {
		return x.OrderId
	}
	return ""
}

func (x *HotelReconciliation) GetCampaign() string {
	if x != nil {
		return x.Campaign
	}
	return ""
}

func (x *HotelReconciliation) GetHotelCenterId() int64 {
	if x != nil {
		return x.HotelCenterId
	}
	return 0
}

func (x *HotelReconciliation) GetHotelId() string {
	if x != nil {
		return x.HotelId
	}
	return ""
}

func (x *HotelReconciliation) GetCheckInDate() string {
	if x != nil {
		return x.CheckInDate
	}
	return ""
}

func (x *HotelReconciliation) GetCheckOutDate() string {
	if x != nil {
		return x.CheckOutDate
	}
	return ""
}

func (x *HotelReconciliation) GetReconciledValueMicros() int64 {
	if x != nil {
		return x.ReconciledValueMicros
	}
	return 0
}

func (x *HotelReconciliation) GetBilled() bool {
	if x != nil {
		return x.Billed
	}
	return false
}

func (x *HotelReconciliation) GetStatus() enums.HotelReconciliationStatusEnum_HotelReconciliationStatus {
	if x != nil {
		return x.Status
	}
	return enums.HotelReconciliationStatusEnum_HotelReconciliationStatus(0)
}

var File_resources_hotel_reconciliation_proto protoreflect.FileDescriptor

var file_resources_hotel_reconciliation_proto_rawDesc = string([]byte{
	0x0a, 0x3d, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x64, 0x73, 0x2f, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x61, 0x64, 0x73, 0x2f, 0x76, 0x31, 0x39, 0x2f, 0x72, 0x65, 0x73, 0x6f, 0x75,
	0x72, 0x63, 0x65, 0x73, 0x2f, 0x68, 0x6f, 0x74, 0x65, 0x6c, 0x5f, 0x72, 0x65, 0x63, 0x6f, 0x6e,
	0x63, 0x69, 0x6c, 0x69, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12,
	0x22, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x61, 0x64, 0x73, 0x2e, 0x67, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x61, 0x64, 0x73, 0x2e, 0x76, 0x31, 0x39, 0x2e, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72,
	0x63, 0x65, 0x73, 0x1a, 0x40, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x64, 0x73, 0x2f,
	0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61, 0x64, 0x73, 0x2f, 0x76, 0x31, 0x39, 0x2f, 0x65, 0x6e,
	0x75, 0x6d, 0x73, 0x2f, 0x68, 0x6f, 0x74, 0x65, 0x6c, 0x5f, 0x72, 0x65, 0x63, 0x6f, 0x6e, 0x63,
	0x69, 0x6c, 0x69, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70,
	0x69, 0x2f, 0x66, 0x69, 0x65, 0x6c, 0x64, 0x5f, 0x62, 0x65, 0x68, 0x61, 0x76, 0x69, 0x6f, 0x72,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x19, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61,
	0x70, 0x69, 0x2f, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x22, 0xec, 0x05, 0x0a, 0x13, 0x48, 0x6f, 0x74, 0x65, 0x6c, 0x52, 0x65, 0x63, 0x6f, 0x6e,
	0x63, 0x69, 0x6c, 0x69, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x59, 0x0a, 0x0d, 0x72, 0x65, 0x73,
	0x6f, 0x75, 0x72, 0x63, 0x65, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x42, 0x34, 0xe0, 0x41, 0x05, 0xfa, 0x41, 0x2e, 0x0a, 0x2c, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x61, 0x64, 0x73, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61, 0x70, 0x69, 0x73, 0x2e, 0x63,
	0x6f, 0x6d, 0x2f, 0x48, 0x6f, 0x74, 0x65, 0x6c, 0x52, 0x65, 0x63, 0x6f, 0x6e, 0x63, 0x69, 0x6c,
	0x69, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x0c, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65,
	0x4e, 0x61, 0x6d, 0x65, 0x12, 0x2b, 0x0a, 0x0d, 0x63, 0x6f, 0x6d, 0x6d, 0x69, 0x73, 0x73, 0x69,
	0x6f, 0x6e, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x42, 0x06, 0xe0, 0x41, 0x02,
	0xe0, 0x41, 0x03, 0x52, 0x0c, 0x63, 0x6f, 0x6d, 0x6d, 0x69, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x49,
	0x64, 0x12, 0x1e, 0x0a, 0x08, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x03, 0x20,
	0x01, 0x28, 0x09, 0x42, 0x03, 0xe0, 0x41, 0x03, 0x52, 0x07, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x49,
	0x64, 0x12, 0x45, 0x0a, 0x08, 0x63, 0x61, 0x6d, 0x70, 0x61, 0x69, 0x67, 0x6e, 0x18, 0x0b, 0x20,
	0x01, 0x28, 0x09, 0x42, 0x29, 0xe0, 0x41, 0x03, 0xfa, 0x41, 0x23, 0x0a, 0x21, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x61, 0x64, 0x73, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61, 0x70, 0x69,
	0x73, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x43, 0x61, 0x6d, 0x70, 0x61, 0x69, 0x67, 0x6e, 0x52, 0x08,
	0x63, 0x61, 0x6d, 0x70, 0x61, 0x69, 0x67, 0x6e, 0x12, 0x2b, 0x0a, 0x0f, 0x68, 0x6f, 0x74, 0x65,
	0x6c, 0x5f, 0x63, 0x65, 0x6e, 0x74, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x04, 0x20, 0x01, 0x28,
	0x03, 0x42, 0x03, 0xe0, 0x41, 0x03, 0x52, 0x0d, 0x68, 0x6f, 0x74, 0x65, 0x6c, 0x43, 0x65, 0x6e,
	0x74, 0x65, 0x72, 0x49, 0x64, 0x12, 0x1e, 0x0a, 0x08, 0x68, 0x6f, 0x74, 0x65, 0x6c, 0x5f, 0x69,
	0x64, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x42, 0x03, 0xe0, 0x41, 0x03, 0x52, 0x07, 0x68, 0x6f,
	0x74, 0x65, 0x6c, 0x49, 0x64, 0x12, 0x27, 0x0a, 0x0d, 0x63, 0x68, 0x65, 0x63, 0x6b, 0x5f, 0x69,
	0x6e, 0x5f, 0x64, 0x61, 0x74, 0x65, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x42, 0x03, 0xe0, 0x41,
	0x03, 0x52, 0x0b, 0x63, 0x68, 0x65, 0x63, 0x6b, 0x49, 0x6e, 0x44, 0x61, 0x74, 0x65, 0x12, 0x29,
	0x0a, 0x0e, 0x63, 0x68, 0x65, 0x63, 0x6b, 0x5f, 0x6f, 0x75, 0x74, 0x5f, 0x64, 0x61, 0x74, 0x65,
	0x18, 0x07, 0x20, 0x01, 0x28, 0x09, 0x42, 0x03, 0xe0, 0x41, 0x03, 0x52, 0x0c, 0x63, 0x68, 0x65,
	0x63, 0x6b, 0x4f, 0x75, 0x74, 0x44, 0x61, 0x74, 0x65, 0x12, 0x3e, 0x0a, 0x17, 0x72, 0x65, 0x63,
	0x6f, 0x6e, 0x63, 0x69, 0x6c, 0x65, 0x64, 0x5f, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x5f, 0x6d, 0x69,
	0x63, 0x72, 0x6f, 0x73, 0x18, 0x08, 0x20, 0x01, 0x28, 0x03, 0x42, 0x06, 0xe0, 0x41, 0x02, 0xe0,
	0x41, 0x03, 0x52, 0x15, 0x72, 0x65, 0x63, 0x6f, 0x6e, 0x63, 0x69, 0x6c, 0x65, 0x64, 0x56, 0x61,
	0x6c, 0x75, 0x65, 0x4d, 0x69, 0x63, 0x72, 0x6f, 0x73, 0x12, 0x1b, 0x0a, 0x06, 0x62, 0x69, 0x6c,
	0x6c, 0x65, 0x64, 0x18, 0x09, 0x20, 0x01, 0x28, 0x08, 0x42, 0x03, 0xe0, 0x41, 0x03, 0x52, 0x06,
	0x62, 0x69, 0x6c, 0x6c, 0x65, 0x64, 0x12, 0x77, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73,
	0x18, 0x0a, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x57, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e,
	0x61, 0x64, 0x73, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61, 0x64, 0x73, 0x2e, 0x76, 0x31,
	0x39, 0x2e, 0x65, 0x6e, 0x75, 0x6d, 0x73, 0x2e, 0x48, 0x6f, 0x74, 0x65, 0x6c, 0x52, 0x65, 0x63,
	0x6f, 0x6e, 0x63, 0x69, 0x6c, 0x69, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x53, 0x74, 0x61, 0x74, 0x75,
	0x73, 0x45, 0x6e, 0x75, 0x6d, 0x2e, 0x48, 0x6f, 0x74, 0x65, 0x6c, 0x52, 0x65, 0x63, 0x6f, 0x6e,
	0x63, 0x69, 0x6c, 0x69, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x42,
	0x06, 0xe0, 0x41, 0x02, 0xe0, 0x41, 0x03, 0x52, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x3a,
	0x6f, 0xea, 0x41, 0x6c, 0x0a, 0x2c, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61, 0x64, 0x73, 0x2e,
	0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61, 0x70, 0x69, 0x73, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x48,
	0x6f, 0x74, 0x65, 0x6c, 0x52, 0x65, 0x63, 0x6f, 0x6e, 0x63, 0x69, 0x6c, 0x69, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x12, 0x3c, 0x63, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x65, 0x72, 0x73, 0x2f, 0x7b, 0x63,
	0x75, 0x73, 0x74, 0x6f, 0x6d, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x7d, 0x2f, 0x68, 0x6f, 0x74, 0x65,
	0x6c, 0x52, 0x65, 0x63, 0x6f, 0x6e, 0x63, 0x69, 0x6c, 0x69, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73,
	0x2f, 0x7b, 0x63, 0x6f, 0x6d, 0x6d, 0x69, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x5f, 0x69, 0x64, 0x7d,
	0x42, 0x8a, 0x02, 0x0a, 0x26, 0x63, 0x6f, 0x6d, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e,
	0x61, 0x64, 0x73, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61, 0x64, 0x73, 0x2e, 0x76, 0x31,
	0x39, 0x2e, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x73, 0x42, 0x18, 0x48, 0x6f, 0x74,
	0x65, 0x6c, 0x52, 0x65, 0x63, 0x6f, 0x6e, 0x63, 0x69, 0x6c, 0x69, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a, 0x4b, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e,
	0x67, 0x6f, 0x6c, 0x61, 0x6e, 0x67, 0x2e, 0x6f, 0x72, 0x67, 0x2f, 0x67, 0x65, 0x6e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x2f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61, 0x70, 0x69, 0x73, 0x2f, 0x61,
	0x64, 0x73, 0x2f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61, 0x64, 0x73, 0x2f, 0x76, 0x31, 0x39,
	0x2f, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x73, 0x3b, 0x72, 0x65, 0x73, 0x6f, 0x75,
	0x72, 0x63, 0x65, 0x73, 0xa2, 0x02, 0x03, 0x47, 0x41, 0x41, 0xaa, 0x02, 0x22, 0x47, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x2e, 0x41, 0x64, 0x73, 0x2e, 0x47, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x41, 0x64,
	0x73, 0x2e, 0x56, 0x31, 0x39, 0x2e, 0x52, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x73, 0xca,
	0x02, 0x22, 0x47, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x5c, 0x41, 0x64, 0x73, 0x5c, 0x47, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x41, 0x64, 0x73, 0x5c, 0x56, 0x31, 0x39, 0x5c, 0x52, 0x65, 0x73, 0x6f, 0x75,
	0x72, 0x63, 0x65, 0x73, 0xea, 0x02, 0x26, 0x47, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x3a, 0x3a, 0x41,
	0x64, 0x73, 0x3a, 0x3a, 0x47, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x41, 0x64, 0x73, 0x3a, 0x3a, 0x56,
	0x31, 0x39, 0x3a, 0x3a, 0x52, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x73, 0x62, 0x06, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x33,
})

var (
	file_resources_hotel_reconciliation_proto_rawDescOnce sync.Once
	file_resources_hotel_reconciliation_proto_rawDescData []byte
)

func file_resources_hotel_reconciliation_proto_rawDescGZIP() []byte {
	file_resources_hotel_reconciliation_proto_rawDescOnce.Do(func() {
		file_resources_hotel_reconciliation_proto_rawDescData = protoimpl.X.CompressGZIP(unsafe.Slice(unsafe.StringData(file_resources_hotel_reconciliation_proto_rawDesc), len(file_resources_hotel_reconciliation_proto_rawDesc)))
	})
	return file_resources_hotel_reconciliation_proto_rawDescData
}

var file_resources_hotel_reconciliation_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_resources_hotel_reconciliation_proto_goTypes = []any{
	(*HotelReconciliation)(nil),                                        // 0: google.ads.googleads.v19.resources.HotelReconciliation
	(enums.HotelReconciliationStatusEnum_HotelReconciliationStatus)(0), // 1: google.ads.googleads.v19.enums.HotelReconciliationStatusEnum.HotelReconciliationStatus
}
var file_resources_hotel_reconciliation_proto_depIdxs = []int32{
	1, // 0: google.ads.googleads.v19.resources.HotelReconciliation.status:type_name -> google.ads.googleads.v19.enums.HotelReconciliationStatusEnum.HotelReconciliationStatus
	1, // [1:1] is the sub-list for method output_type
	1, // [1:1] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_resources_hotel_reconciliation_proto_init() }
func file_resources_hotel_reconciliation_proto_init() {
	if File_resources_hotel_reconciliation_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: unsafe.Slice(unsafe.StringData(file_resources_hotel_reconciliation_proto_rawDesc), len(file_resources_hotel_reconciliation_proto_rawDesc)),
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_resources_hotel_reconciliation_proto_goTypes,
		DependencyIndexes: file_resources_hotel_reconciliation_proto_depIdxs,
		MessageInfos:      file_resources_hotel_reconciliation_proto_msgTypes,
	}.Build()
	File_resources_hotel_reconciliation_proto = out.File
	file_resources_hotel_reconciliation_proto_goTypes = nil
	file_resources_hotel_reconciliation_proto_depIdxs = nil
}
