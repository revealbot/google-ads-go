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
// source: google/ads/googleads/v13/services/campaign_conversion_goal_service.proto

package services

import (
	resources "github.com/revealbot/google-ads-go/resources"
	_ "google.golang.org/genproto/googleapis/api/annotations"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	fieldmaskpb "google.golang.org/protobuf/types/known/fieldmaskpb"
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
// [CampaignConversionGoalService.MutateCampaignConversionGoals][google.ads.googleads.v13.services.CampaignConversionGoalService.MutateCampaignConversionGoals].
type MutateCampaignConversionGoalsRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Required. The ID of the customer whose campaign conversion goals are being
	// modified.
	CustomerId string `protobuf:"bytes,1,opt,name=customer_id,json=customerId,proto3" json:"customer_id,omitempty"`
	// Required. The list of operations to perform on individual campaign
	// conversion goal.
	Operations []*CampaignConversionGoalOperation `protobuf:"bytes,2,rep,name=operations,proto3" json:"operations,omitempty"`
	// If true, the request is validated but not executed. Only errors are
	// returned, not results.
	ValidateOnly bool `protobuf:"varint,3,opt,name=validate_only,json=validateOnly,proto3" json:"validate_only,omitempty"`
}

func (x *MutateCampaignConversionGoalsRequest) Reset() {
	*x = MutateCampaignConversionGoalsRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_services_campaign_conversion_goal_service_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MutateCampaignConversionGoalsRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MutateCampaignConversionGoalsRequest) ProtoMessage() {}

func (x *MutateCampaignConversionGoalsRequest) ProtoReflect() protoreflect.Message {
	mi := &file_services_campaign_conversion_goal_service_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MutateCampaignConversionGoalsRequest.ProtoReflect.Descriptor instead.
func (*MutateCampaignConversionGoalsRequest) Descriptor() ([]byte, []int) {
	return file_services_campaign_conversion_goal_service_proto_rawDescGZIP(), []int{0}
}

func (x *MutateCampaignConversionGoalsRequest) GetCustomerId() string {
	if x != nil {
		return x.CustomerId
	}
	return ""
}

func (x *MutateCampaignConversionGoalsRequest) GetOperations() []*CampaignConversionGoalOperation {
	if x != nil {
		return x.Operations
	}
	return nil
}

func (x *MutateCampaignConversionGoalsRequest) GetValidateOnly() bool {
	if x != nil {
		return x.ValidateOnly
	}
	return false
}

// A single operation (update) on a campaign conversion goal.
type CampaignConversionGoalOperation struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// FieldMask that determines which resource fields are modified in an update.
	UpdateMask *fieldmaskpb.FieldMask `protobuf:"bytes,2,opt,name=update_mask,json=updateMask,proto3" json:"update_mask,omitempty"`
	// The mutate operation.
	//
	// Types that are assignable to Operation:
	//	*CampaignConversionGoalOperation_Update
	Operation isCampaignConversionGoalOperation_Operation `protobuf_oneof:"operation"`
}

func (x *CampaignConversionGoalOperation) Reset() {
	*x = CampaignConversionGoalOperation{}
	if protoimpl.UnsafeEnabled {
		mi := &file_services_campaign_conversion_goal_service_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CampaignConversionGoalOperation) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CampaignConversionGoalOperation) ProtoMessage() {}

func (x *CampaignConversionGoalOperation) ProtoReflect() protoreflect.Message {
	mi := &file_services_campaign_conversion_goal_service_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CampaignConversionGoalOperation.ProtoReflect.Descriptor instead.
func (*CampaignConversionGoalOperation) Descriptor() ([]byte, []int) {
	return file_services_campaign_conversion_goal_service_proto_rawDescGZIP(), []int{1}
}

func (x *CampaignConversionGoalOperation) GetUpdateMask() *fieldmaskpb.FieldMask {
	if x != nil {
		return x.UpdateMask
	}
	return nil
}

func (m *CampaignConversionGoalOperation) GetOperation() isCampaignConversionGoalOperation_Operation {
	if m != nil {
		return m.Operation
	}
	return nil
}

func (x *CampaignConversionGoalOperation) GetUpdate() *resources.CampaignConversionGoal {
	if x, ok := x.GetOperation().(*CampaignConversionGoalOperation_Update); ok {
		return x.Update
	}
	return nil
}

type isCampaignConversionGoalOperation_Operation interface {
	isCampaignConversionGoalOperation_Operation()
}

type CampaignConversionGoalOperation_Update struct {
	// Update operation: The customer conversion goal is expected to have a
	// valid resource name.
	Update *resources.CampaignConversionGoal `protobuf:"bytes,1,opt,name=update,proto3,oneof"`
}

func (*CampaignConversionGoalOperation_Update) isCampaignConversionGoalOperation_Operation() {}

// Response message for a campaign conversion goal mutate.
type MutateCampaignConversionGoalsResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// All results for the mutate.
	Results []*MutateCampaignConversionGoalResult `protobuf:"bytes,1,rep,name=results,proto3" json:"results,omitempty"`
}

func (x *MutateCampaignConversionGoalsResponse) Reset() {
	*x = MutateCampaignConversionGoalsResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_services_campaign_conversion_goal_service_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MutateCampaignConversionGoalsResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MutateCampaignConversionGoalsResponse) ProtoMessage() {}

func (x *MutateCampaignConversionGoalsResponse) ProtoReflect() protoreflect.Message {
	mi := &file_services_campaign_conversion_goal_service_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MutateCampaignConversionGoalsResponse.ProtoReflect.Descriptor instead.
func (*MutateCampaignConversionGoalsResponse) Descriptor() ([]byte, []int) {
	return file_services_campaign_conversion_goal_service_proto_rawDescGZIP(), []int{2}
}

func (x *MutateCampaignConversionGoalsResponse) GetResults() []*MutateCampaignConversionGoalResult {
	if x != nil {
		return x.Results
	}
	return nil
}

// The result for the campaign conversion goal mutate.
type MutateCampaignConversionGoalResult struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Returned for successful operations.
	ResourceName string `protobuf:"bytes,1,opt,name=resource_name,json=resourceName,proto3" json:"resource_name,omitempty"`
}

func (x *MutateCampaignConversionGoalResult) Reset() {
	*x = MutateCampaignConversionGoalResult{}
	if protoimpl.UnsafeEnabled {
		mi := &file_services_campaign_conversion_goal_service_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MutateCampaignConversionGoalResult) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MutateCampaignConversionGoalResult) ProtoMessage() {}

func (x *MutateCampaignConversionGoalResult) ProtoReflect() protoreflect.Message {
	mi := &file_services_campaign_conversion_goal_service_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MutateCampaignConversionGoalResult.ProtoReflect.Descriptor instead.
func (*MutateCampaignConversionGoalResult) Descriptor() ([]byte, []int) {
	return file_services_campaign_conversion_goal_service_proto_rawDescGZIP(), []int{3}
}

func (x *MutateCampaignConversionGoalResult) GetResourceName() string {
	if x != nil {
		return x.ResourceName
	}
	return ""
}

var File_services_campaign_conversion_goal_service_proto protoreflect.FileDescriptor

var file_services_campaign_conversion_goal_service_proto_rawDesc = []byte{
	0x0a, 0x48, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x64, 0x73, 0x2f, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x61, 0x64, 0x73, 0x2f, 0x76, 0x31, 0x33, 0x2f, 0x73, 0x65, 0x72, 0x76, 0x69,
	0x63, 0x65, 0x73, 0x2f, 0x63, 0x61, 0x6d, 0x70, 0x61, 0x69, 0x67, 0x6e, 0x5f, 0x63, 0x6f, 0x6e,
	0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x5f, 0x67, 0x6f, 0x61, 0x6c, 0x5f, 0x73, 0x65, 0x72,
	0x76, 0x69, 0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x21, 0x67, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x2e, 0x61, 0x64, 0x73, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61, 0x64, 0x73,
	0x2e, 0x76, 0x31, 0x33, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x1a, 0x41, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x64, 0x73, 0x2f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x61, 0x64, 0x73, 0x2f, 0x76, 0x31, 0x33, 0x2f, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65,
	0x73, 0x2f, 0x63, 0x61, 0x6d, 0x70, 0x61, 0x69, 0x67, 0x6e, 0x5f, 0x63, 0x6f, 0x6e, 0x76, 0x65,
	0x72, 0x73, 0x69, 0x6f, 0x6e, 0x5f, 0x67, 0x6f, 0x61, 0x6c, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x1a, 0x1c, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x61, 0x6e, 0x6e,
	0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x17,
	0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x63, 0x6c, 0x69, 0x65, 0x6e,
	0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f,
	0x61, 0x70, 0x69, 0x2f, 0x66, 0x69, 0x65, 0x6c, 0x64, 0x5f, 0x62, 0x65, 0x68, 0x61, 0x76, 0x69,
	0x6f, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x19, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x2f, 0x61, 0x70, 0x69, 0x2f, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x1a, 0x20, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x62, 0x75, 0x66, 0x2f, 0x66, 0x69, 0x65, 0x6c, 0x64, 0x5f, 0x6d, 0x61, 0x73, 0x6b, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xda, 0x01, 0x0a, 0x24, 0x4d, 0x75, 0x74, 0x61, 0x74, 0x65,
	0x43, 0x61, 0x6d, 0x70, 0x61, 0x69, 0x67, 0x6e, 0x43, 0x6f, 0x6e, 0x76, 0x65, 0x72, 0x73, 0x69,
	0x6f, 0x6e, 0x47, 0x6f, 0x61, 0x6c, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x24,
	0x0a, 0x0b, 0x63, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x42, 0x03, 0xe0, 0x41, 0x02, 0x52, 0x0a, 0x63, 0x75, 0x73, 0x74, 0x6f, 0x6d,
	0x65, 0x72, 0x49, 0x64, 0x12, 0x67, 0x0a, 0x0a, 0x6f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x42, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x2e, 0x61, 0x64, 0x73, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61, 0x64, 0x73, 0x2e,
	0x76, 0x31, 0x33, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x43, 0x61, 0x6d,
	0x70, 0x61, 0x69, 0x67, 0x6e, 0x43, 0x6f, 0x6e, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x47,
	0x6f, 0x61, 0x6c, 0x4f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x42, 0x03, 0xe0, 0x41,
	0x02, 0x52, 0x0a, 0x6f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x12, 0x23, 0x0a,
	0x0d, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x65, 0x5f, 0x6f, 0x6e, 0x6c, 0x79, 0x18, 0x03,
	0x20, 0x01, 0x28, 0x08, 0x52, 0x0c, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x65, 0x4f, 0x6e,
	0x6c, 0x79, 0x22, 0xc1, 0x01, 0x0a, 0x1f, 0x43, 0x61, 0x6d, 0x70, 0x61, 0x69, 0x67, 0x6e, 0x43,
	0x6f, 0x6e, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x47, 0x6f, 0x61, 0x6c, 0x4f, 0x70, 0x65,
	0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x3b, 0x0a, 0x0b, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65,
	0x5f, 0x6d, 0x61, 0x73, 0x6b, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x46, 0x69,
	0x65, 0x6c, 0x64, 0x4d, 0x61, 0x73, 0x6b, 0x52, 0x0a, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x4d,
	0x61, 0x73, 0x6b, 0x12, 0x54, 0x0a, 0x06, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x3a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x61, 0x64, 0x73,
	0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61, 0x64, 0x73, 0x2e, 0x76, 0x31, 0x33, 0x2e, 0x72,
	0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x73, 0x2e, 0x43, 0x61, 0x6d, 0x70, 0x61, 0x69, 0x67,
	0x6e, 0x43, 0x6f, 0x6e, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x47, 0x6f, 0x61, 0x6c, 0x48,
	0x00, 0x52, 0x06, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x42, 0x0b, 0x0a, 0x09, 0x6f, 0x70, 0x65,
	0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x22, 0x88, 0x01, 0x0a, 0x25, 0x4d, 0x75, 0x74, 0x61, 0x74,
	0x65, 0x43, 0x61, 0x6d, 0x70, 0x61, 0x69, 0x67, 0x6e, 0x43, 0x6f, 0x6e, 0x76, 0x65, 0x72, 0x73,
	0x69, 0x6f, 0x6e, 0x47, 0x6f, 0x61, 0x6c, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x12, 0x5f, 0x0a, 0x07, 0x72, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28,
	0x0b, 0x32, 0x45, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x61, 0x64, 0x73, 0x2e, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61, 0x64, 0x73, 0x2e, 0x76, 0x31, 0x33, 0x2e, 0x73, 0x65, 0x72,
	0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x4d, 0x75, 0x74, 0x61, 0x74, 0x65, 0x43, 0x61, 0x6d, 0x70,
	0x61, 0x69, 0x67, 0x6e, 0x43, 0x6f, 0x6e, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x47, 0x6f,
	0x61, 0x6c, 0x52, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x52, 0x07, 0x72, 0x65, 0x73, 0x75, 0x6c, 0x74,
	0x73, 0x22, 0x7f, 0x0a, 0x22, 0x4d, 0x75, 0x74, 0x61, 0x74, 0x65, 0x43, 0x61, 0x6d, 0x70, 0x61,
	0x69, 0x67, 0x6e, 0x43, 0x6f, 0x6e, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x47, 0x6f, 0x61,
	0x6c, 0x52, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x12, 0x59, 0x0a, 0x0d, 0x72, 0x65, 0x73, 0x6f, 0x75,
	0x72, 0x63, 0x65, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x42, 0x34,
	0xfa, 0x41, 0x31, 0x0a, 0x2f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61, 0x64, 0x73, 0x2e, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61, 0x70, 0x69, 0x73, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x43, 0x61,
	0x6d, 0x70, 0x61, 0x69, 0x67, 0x6e, 0x43, 0x6f, 0x6e, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e,
	0x47, 0x6f, 0x61, 0x6c, 0x52, 0x0c, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x4e, 0x61,
	0x6d, 0x65, 0x32, 0xfe, 0x02, 0x0a, 0x1d, 0x43, 0x61, 0x6d, 0x70, 0x61, 0x69, 0x67, 0x6e, 0x43,
	0x6f, 0x6e, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x47, 0x6f, 0x61, 0x6c, 0x53, 0x65, 0x72,
	0x76, 0x69, 0x63, 0x65, 0x12, 0x95, 0x02, 0x0a, 0x1d, 0x4d, 0x75, 0x74, 0x61, 0x74, 0x65, 0x43,
	0x61, 0x6d, 0x70, 0x61, 0x69, 0x67, 0x6e, 0x43, 0x6f, 0x6e, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f,
	0x6e, 0x47, 0x6f, 0x61, 0x6c, 0x73, 0x12, 0x47, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e,
	0x61, 0x64, 0x73, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61, 0x64, 0x73, 0x2e, 0x76, 0x31,
	0x33, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x4d, 0x75, 0x74, 0x61, 0x74,
	0x65, 0x43, 0x61, 0x6d, 0x70, 0x61, 0x69, 0x67, 0x6e, 0x43, 0x6f, 0x6e, 0x76, 0x65, 0x72, 0x73,
	0x69, 0x6f, 0x6e, 0x47, 0x6f, 0x61, 0x6c, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a,
	0x48, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x61, 0x64, 0x73, 0x2e, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x61, 0x64, 0x73, 0x2e, 0x76, 0x31, 0x33, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69,
	0x63, 0x65, 0x73, 0x2e, 0x4d, 0x75, 0x74, 0x61, 0x74, 0x65, 0x43, 0x61, 0x6d, 0x70, 0x61, 0x69,
	0x67, 0x6e, 0x43, 0x6f, 0x6e, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x47, 0x6f, 0x61, 0x6c,
	0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x61, 0x82, 0xd3, 0xe4, 0x93, 0x02,
	0x42, 0x22, 0x3d, 0x2f, 0x76, 0x31, 0x33, 0x2f, 0x63, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x65, 0x72,
	0x73, 0x2f, 0x7b, 0x63, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x3d, 0x2a,
	0x7d, 0x2f, 0x63, 0x61, 0x6d, 0x70, 0x61, 0x69, 0x67, 0x6e, 0x43, 0x6f, 0x6e, 0x76, 0x65, 0x72,
	0x73, 0x69, 0x6f, 0x6e, 0x47, 0x6f, 0x61, 0x6c, 0x73, 0x3a, 0x6d, 0x75, 0x74, 0x61, 0x74, 0x65,
	0x3a, 0x01, 0x2a, 0xda, 0x41, 0x16, 0x63, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x65, 0x72, 0x5f, 0x69,
	0x64, 0x2c, 0x6f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x1a, 0x45, 0xca, 0x41,
	0x18, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61, 0x64, 0x73, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x61, 0x70, 0x69, 0x73, 0x2e, 0x63, 0x6f, 0x6d, 0xd2, 0x41, 0x27, 0x68, 0x74, 0x74, 0x70,
	0x73, 0x3a, 0x2f, 0x2f, 0x77, 0x77, 0x77, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61, 0x70,
	0x69, 0x73, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x61, 0x75, 0x74, 0x68, 0x2f, 0x61, 0x64, 0x77, 0x6f,
	0x72, 0x64, 0x73, 0x42, 0x8e, 0x02, 0x0a, 0x25, 0x63, 0x6f, 0x6d, 0x2e, 0x67, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x2e, 0x61, 0x64, 0x73, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61, 0x64, 0x73,
	0x2e, 0x76, 0x31, 0x33, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x42, 0x22, 0x43,
	0x61, 0x6d, 0x70, 0x61, 0x69, 0x67, 0x6e, 0x43, 0x6f, 0x6e, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f,
	0x6e, 0x47, 0x6f, 0x61, 0x6c, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x50, 0x72, 0x6f, 0x74,
	0x6f, 0x50, 0x01, 0x5a, 0x49, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x67, 0x6f, 0x6c, 0x61,
	0x6e, 0x67, 0x2e, 0x6f, 0x72, 0x67, 0x2f, 0x67, 0x65, 0x6e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f,
	0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61, 0x70, 0x69, 0x73, 0x2f, 0x61, 0x64, 0x73, 0x2f, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61, 0x64, 0x73, 0x2f, 0x76, 0x31, 0x33, 0x2f, 0x73, 0x65, 0x72,
	0x76, 0x69, 0x63, 0x65, 0x73, 0x3b, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0xa2, 0x02,
	0x03, 0x47, 0x41, 0x41, 0xaa, 0x02, 0x21, 0x47, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x41, 0x64,
	0x73, 0x2e, 0x47, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x41, 0x64, 0x73, 0x2e, 0x56, 0x31, 0x33, 0x2e,
	0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0xca, 0x02, 0x21, 0x47, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x5c, 0x41, 0x64, 0x73, 0x5c, 0x47, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x41, 0x64, 0x73, 0x5c,
	0x56, 0x31, 0x33, 0x5c, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0xea, 0x02, 0x25, 0x47,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x3a, 0x3a, 0x41, 0x64, 0x73, 0x3a, 0x3a, 0x47, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x41, 0x64, 0x73, 0x3a, 0x3a, 0x56, 0x31, 0x33, 0x3a, 0x3a, 0x53, 0x65, 0x72, 0x76,
	0x69, 0x63, 0x65, 0x73, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_services_campaign_conversion_goal_service_proto_rawDescOnce sync.Once
	file_services_campaign_conversion_goal_service_proto_rawDescData = file_services_campaign_conversion_goal_service_proto_rawDesc
)

func file_services_campaign_conversion_goal_service_proto_rawDescGZIP() []byte {
	file_services_campaign_conversion_goal_service_proto_rawDescOnce.Do(func() {
		file_services_campaign_conversion_goal_service_proto_rawDescData = protoimpl.X.CompressGZIP(file_services_campaign_conversion_goal_service_proto_rawDescData)
	})
	return file_services_campaign_conversion_goal_service_proto_rawDescData
}

var file_services_campaign_conversion_goal_service_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_services_campaign_conversion_goal_service_proto_goTypes = []interface{}{
	(*MutateCampaignConversionGoalsRequest)(nil),  // 0: google.ads.googleads.v13.services.MutateCampaignConversionGoalsRequest
	(*CampaignConversionGoalOperation)(nil),       // 1: google.ads.googleads.v13.services.CampaignConversionGoalOperation
	(*MutateCampaignConversionGoalsResponse)(nil), // 2: google.ads.googleads.v13.services.MutateCampaignConversionGoalsResponse
	(*MutateCampaignConversionGoalResult)(nil),    // 3: google.ads.googleads.v13.services.MutateCampaignConversionGoalResult
	(*fieldmaskpb.FieldMask)(nil),                 // 4: google.protobuf.FieldMask
	(*resources.CampaignConversionGoal)(nil),      // 5: google.ads.googleads.v13.resources.CampaignConversionGoal
}
var file_services_campaign_conversion_goal_service_proto_depIdxs = []int32{
	1, // 0: google.ads.googleads.v13.services.MutateCampaignConversionGoalsRequest.operations:type_name -> google.ads.googleads.v13.services.CampaignConversionGoalOperation
	4, // 1: google.ads.googleads.v13.services.CampaignConversionGoalOperation.update_mask:type_name -> google.protobuf.FieldMask
	5, // 2: google.ads.googleads.v13.services.CampaignConversionGoalOperation.update:type_name -> google.ads.googleads.v13.resources.CampaignConversionGoal
	3, // 3: google.ads.googleads.v13.services.MutateCampaignConversionGoalsResponse.results:type_name -> google.ads.googleads.v13.services.MutateCampaignConversionGoalResult
	0, // 4: google.ads.googleads.v13.services.CampaignConversionGoalService.MutateCampaignConversionGoals:input_type -> google.ads.googleads.v13.services.MutateCampaignConversionGoalsRequest
	2, // 5: google.ads.googleads.v13.services.CampaignConversionGoalService.MutateCampaignConversionGoals:output_type -> google.ads.googleads.v13.services.MutateCampaignConversionGoalsResponse
	5, // [5:6] is the sub-list for method output_type
	4, // [4:5] is the sub-list for method input_type
	4, // [4:4] is the sub-list for extension type_name
	4, // [4:4] is the sub-list for extension extendee
	0, // [0:4] is the sub-list for field type_name
}

func init() { file_services_campaign_conversion_goal_service_proto_init() }
func file_services_campaign_conversion_goal_service_proto_init() {
	if File_services_campaign_conversion_goal_service_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_services_campaign_conversion_goal_service_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*MutateCampaignConversionGoalsRequest); i {
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
		file_services_campaign_conversion_goal_service_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CampaignConversionGoalOperation); i {
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
		file_services_campaign_conversion_goal_service_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*MutateCampaignConversionGoalsResponse); i {
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
		file_services_campaign_conversion_goal_service_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*MutateCampaignConversionGoalResult); i {
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
	file_services_campaign_conversion_goal_service_proto_msgTypes[1].OneofWrappers = []interface{}{
		(*CampaignConversionGoalOperation_Update)(nil),
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_services_campaign_conversion_goal_service_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_services_campaign_conversion_goal_service_proto_goTypes,
		DependencyIndexes: file_services_campaign_conversion_goal_service_proto_depIdxs,
		MessageInfos:      file_services_campaign_conversion_goal_service_proto_msgTypes,
	}.Build()
	File_services_campaign_conversion_goal_service_proto = out.File
	file_services_campaign_conversion_goal_service_proto_rawDesc = nil
	file_services_campaign_conversion_goal_service_proto_goTypes = nil
	file_services_campaign_conversion_goal_service_proto_depIdxs = nil
}
