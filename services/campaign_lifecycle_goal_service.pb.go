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
// source: google/ads/googleads/v19/services/campaign_lifecycle_goal_service.proto

package services

import (
	resources "github.com/revealbot/google-ads-go/resources"
	_ "google.golang.org/genproto/googleapis/api/annotations"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	fieldmaskpb "google.golang.org/protobuf/types/known/fieldmaskpb"
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

// Request message for
// [CampaignLifecycleGoalService.configureCampaignLifecycleGoals][].
type ConfigureCampaignLifecycleGoalsRequest struct {
	state protoimpl.MessageState `protogen:"open.v1"`
	// Required. The ID of the customer performing the upload.
	CustomerId string `protobuf:"bytes,1,opt,name=customer_id,json=customerId,proto3" json:"customer_id,omitempty"`
	// Required. The operation to perform campaign lifecycle goal update.
	Operation *CampaignLifecycleGoalOperation `protobuf:"bytes,2,opt,name=operation,proto3" json:"operation,omitempty"`
	// Optional. If true, the request is validated but not executed. Only errors
	// are returned, not results.
	ValidateOnly  bool `protobuf:"varint,3,opt,name=validate_only,json=validateOnly,proto3" json:"validate_only,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *ConfigureCampaignLifecycleGoalsRequest) Reset() {
	*x = ConfigureCampaignLifecycleGoalsRequest{}
	mi := &file_services_campaign_lifecycle_goal_service_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *ConfigureCampaignLifecycleGoalsRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ConfigureCampaignLifecycleGoalsRequest) ProtoMessage() {}

func (x *ConfigureCampaignLifecycleGoalsRequest) ProtoReflect() protoreflect.Message {
	mi := &file_services_campaign_lifecycle_goal_service_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ConfigureCampaignLifecycleGoalsRequest.ProtoReflect.Descriptor instead.
func (*ConfigureCampaignLifecycleGoalsRequest) Descriptor() ([]byte, []int) {
	return file_services_campaign_lifecycle_goal_service_proto_rawDescGZIP(), []int{0}
}

func (x *ConfigureCampaignLifecycleGoalsRequest) GetCustomerId() string {
	if x != nil {
		return x.CustomerId
	}
	return ""
}

func (x *ConfigureCampaignLifecycleGoalsRequest) GetOperation() *CampaignLifecycleGoalOperation {
	if x != nil {
		return x.Operation
	}
	return nil
}

func (x *ConfigureCampaignLifecycleGoalsRequest) GetValidateOnly() bool {
	if x != nil {
		return x.ValidateOnly
	}
	return false
}

// A single operation on a campaign lifecycle goal.
type CampaignLifecycleGoalOperation struct {
	state protoimpl.MessageState `protogen:"open.v1"`
	// Optional. FieldMask that determines which resource fields are modified in
	// an update.
	UpdateMask *fieldmaskpb.FieldMask `protobuf:"bytes,2,opt,name=update_mask,json=updateMask,proto3" json:"update_mask,omitempty"`
	// The mutate operation.
	//
	// Types that are valid to be assigned to Operation:
	//
	//	*CampaignLifecycleGoalOperation_Create
	//	*CampaignLifecycleGoalOperation_Update
	Operation     isCampaignLifecycleGoalOperation_Operation `protobuf_oneof:"operation"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *CampaignLifecycleGoalOperation) Reset() {
	*x = CampaignLifecycleGoalOperation{}
	mi := &file_services_campaign_lifecycle_goal_service_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *CampaignLifecycleGoalOperation) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CampaignLifecycleGoalOperation) ProtoMessage() {}

func (x *CampaignLifecycleGoalOperation) ProtoReflect() protoreflect.Message {
	mi := &file_services_campaign_lifecycle_goal_service_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CampaignLifecycleGoalOperation.ProtoReflect.Descriptor instead.
func (*CampaignLifecycleGoalOperation) Descriptor() ([]byte, []int) {
	return file_services_campaign_lifecycle_goal_service_proto_rawDescGZIP(), []int{1}
}

func (x *CampaignLifecycleGoalOperation) GetUpdateMask() *fieldmaskpb.FieldMask {
	if x != nil {
		return x.UpdateMask
	}
	return nil
}

func (x *CampaignLifecycleGoalOperation) GetOperation() isCampaignLifecycleGoalOperation_Operation {
	if x != nil {
		return x.Operation
	}
	return nil
}

func (x *CampaignLifecycleGoalOperation) GetCreate() *resources.CampaignLifecycleGoal {
	if x != nil {
		if x, ok := x.Operation.(*CampaignLifecycleGoalOperation_Create); ok {
			return x.Create
		}
	}
	return nil
}

func (x *CampaignLifecycleGoalOperation) GetUpdate() *resources.CampaignLifecycleGoal {
	if x != nil {
		if x, ok := x.Operation.(*CampaignLifecycleGoalOperation_Update); ok {
			return x.Update
		}
	}
	return nil
}

type isCampaignLifecycleGoalOperation_Operation interface {
	isCampaignLifecycleGoalOperation_Operation()
}

type CampaignLifecycleGoalOperation_Create struct {
	// Create operation: Create a new campaign lifecycle goal. The campaign
	// field should be set for this operation.
	Create *resources.CampaignLifecycleGoal `protobuf:"bytes,1,opt,name=create,proto3,oneof"`
}

type CampaignLifecycleGoalOperation_Update struct {
	// Update operation: Update an existing campaign lifecycle goal. The
	// campaign field should not be set for this operation.
	Update *resources.CampaignLifecycleGoal `protobuf:"bytes,3,opt,name=update,proto3,oneof"`
}

func (*CampaignLifecycleGoalOperation_Create) isCampaignLifecycleGoalOperation_Operation() {}

func (*CampaignLifecycleGoalOperation_Update) isCampaignLifecycleGoalOperation_Operation() {}

// Response message for
// [CampaignLifecycleGoalService.configureCampaignLifecycleGoals][].
type ConfigureCampaignLifecycleGoalsResponse struct {
	state protoimpl.MessageState `protogen:"open.v1"`
	// Result for the campaign lifecycle goal configuration.
	Result        *ConfigureCampaignLifecycleGoalsResult `protobuf:"bytes,1,opt,name=result,proto3" json:"result,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *ConfigureCampaignLifecycleGoalsResponse) Reset() {
	*x = ConfigureCampaignLifecycleGoalsResponse{}
	mi := &file_services_campaign_lifecycle_goal_service_proto_msgTypes[2]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *ConfigureCampaignLifecycleGoalsResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ConfigureCampaignLifecycleGoalsResponse) ProtoMessage() {}

func (x *ConfigureCampaignLifecycleGoalsResponse) ProtoReflect() protoreflect.Message {
	mi := &file_services_campaign_lifecycle_goal_service_proto_msgTypes[2]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ConfigureCampaignLifecycleGoalsResponse.ProtoReflect.Descriptor instead.
func (*ConfigureCampaignLifecycleGoalsResponse) Descriptor() ([]byte, []int) {
	return file_services_campaign_lifecycle_goal_service_proto_rawDescGZIP(), []int{2}
}

func (x *ConfigureCampaignLifecycleGoalsResponse) GetResult() *ConfigureCampaignLifecycleGoalsResult {
	if x != nil {
		return x.Result
	}
	return nil
}

// The result for the campaign lifecycle goal configuration.
type ConfigureCampaignLifecycleGoalsResult struct {
	state protoimpl.MessageState `protogen:"open.v1"`
	// Returned for the successful operation.
	ResourceName  string `protobuf:"bytes,1,opt,name=resource_name,json=resourceName,proto3" json:"resource_name,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *ConfigureCampaignLifecycleGoalsResult) Reset() {
	*x = ConfigureCampaignLifecycleGoalsResult{}
	mi := &file_services_campaign_lifecycle_goal_service_proto_msgTypes[3]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *ConfigureCampaignLifecycleGoalsResult) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ConfigureCampaignLifecycleGoalsResult) ProtoMessage() {}

func (x *ConfigureCampaignLifecycleGoalsResult) ProtoReflect() protoreflect.Message {
	mi := &file_services_campaign_lifecycle_goal_service_proto_msgTypes[3]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ConfigureCampaignLifecycleGoalsResult.ProtoReflect.Descriptor instead.
func (*ConfigureCampaignLifecycleGoalsResult) Descriptor() ([]byte, []int) {
	return file_services_campaign_lifecycle_goal_service_proto_rawDescGZIP(), []int{3}
}

func (x *ConfigureCampaignLifecycleGoalsResult) GetResourceName() string {
	if x != nil {
		return x.ResourceName
	}
	return ""
}

var File_services_campaign_lifecycle_goal_service_proto protoreflect.FileDescriptor

var file_services_campaign_lifecycle_goal_service_proto_rawDesc = string([]byte{
	0x0a, 0x47, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x64, 0x73, 0x2f, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x61, 0x64, 0x73, 0x2f, 0x76, 0x31, 0x39, 0x2f, 0x73, 0x65, 0x72, 0x76, 0x69,
	0x63, 0x65, 0x73, 0x2f, 0x63, 0x61, 0x6d, 0x70, 0x61, 0x69, 0x67, 0x6e, 0x5f, 0x6c, 0x69, 0x66,
	0x65, 0x63, 0x79, 0x63, 0x6c, 0x65, 0x5f, 0x67, 0x6f, 0x61, 0x6c, 0x5f, 0x73, 0x65, 0x72, 0x76,
	0x69, 0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x21, 0x67, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x2e, 0x61, 0x64, 0x73, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61, 0x64, 0x73, 0x2e,
	0x76, 0x31, 0x39, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x1a, 0x40, 0x67, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x64, 0x73, 0x2f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61,
	0x64, 0x73, 0x2f, 0x76, 0x31, 0x39, 0x2f, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x73,
	0x2f, 0x63, 0x61, 0x6d, 0x70, 0x61, 0x69, 0x67, 0x6e, 0x5f, 0x6c, 0x69, 0x66, 0x65, 0x63, 0x79,
	0x63, 0x6c, 0x65, 0x5f, 0x67, 0x6f, 0x61, 0x6c, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1c,
	0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x17, 0x67, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70,
	0x69, 0x2f, 0x66, 0x69, 0x65, 0x6c, 0x64, 0x5f, 0x62, 0x65, 0x68, 0x61, 0x76, 0x69, 0x6f, 0x72,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x19, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61,
	0x70, 0x69, 0x2f, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x1a, 0x20, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62,
	0x75, 0x66, 0x2f, 0x66, 0x69, 0x65, 0x6c, 0x64, 0x5f, 0x6d, 0x61, 0x73, 0x6b, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x22, 0xde, 0x01, 0x0a, 0x26, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x75, 0x72,
	0x65, 0x43, 0x61, 0x6d, 0x70, 0x61, 0x69, 0x67, 0x6e, 0x4c, 0x69, 0x66, 0x65, 0x63, 0x79, 0x63,
	0x6c, 0x65, 0x47, 0x6f, 0x61, 0x6c, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x24,
	0x0a, 0x0b, 0x63, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x42, 0x03, 0xe0, 0x41, 0x02, 0x52, 0x0a, 0x63, 0x75, 0x73, 0x74, 0x6f, 0x6d,
	0x65, 0x72, 0x49, 0x64, 0x12, 0x64, 0x0a, 0x09, 0x6f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x41, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x2e, 0x61, 0x64, 0x73, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61, 0x64, 0x73, 0x2e, 0x76,
	0x31, 0x39, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x43, 0x61, 0x6d, 0x70,
	0x61, 0x69, 0x67, 0x6e, 0x4c, 0x69, 0x66, 0x65, 0x63, 0x79, 0x63, 0x6c, 0x65, 0x47, 0x6f, 0x61,
	0x6c, 0x4f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x42, 0x03, 0xe0, 0x41, 0x02, 0x52,
	0x09, 0x6f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x28, 0x0a, 0x0d, 0x76, 0x61,
	0x6c, 0x69, 0x64, 0x61, 0x74, 0x65, 0x5f, 0x6f, 0x6e, 0x6c, 0x79, 0x18, 0x03, 0x20, 0x01, 0x28,
	0x08, 0x42, 0x03, 0xe0, 0x41, 0x01, 0x52, 0x0c, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x65,
	0x4f, 0x6e, 0x6c, 0x79, 0x22, 0x99, 0x02, 0x0a, 0x1e, 0x43, 0x61, 0x6d, 0x70, 0x61, 0x69, 0x67,
	0x6e, 0x4c, 0x69, 0x66, 0x65, 0x63, 0x79, 0x63, 0x6c, 0x65, 0x47, 0x6f, 0x61, 0x6c, 0x4f, 0x70,
	0x65, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x40, 0x0a, 0x0b, 0x75, 0x70, 0x64, 0x61, 0x74,
	0x65, 0x5f, 0x6d, 0x61, 0x73, 0x6b, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x46,
	0x69, 0x65, 0x6c, 0x64, 0x4d, 0x61, 0x73, 0x6b, 0x42, 0x03, 0xe0, 0x41, 0x01, 0x52, 0x0a, 0x75,
	0x70, 0x64, 0x61, 0x74, 0x65, 0x4d, 0x61, 0x73, 0x6b, 0x12, 0x53, 0x0a, 0x06, 0x63, 0x72, 0x65,
	0x61, 0x74, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x39, 0x2e, 0x67, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x2e, 0x61, 0x64, 0x73, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61, 0x64, 0x73,
	0x2e, 0x76, 0x31, 0x39, 0x2e, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x73, 0x2e, 0x43,
	0x61, 0x6d, 0x70, 0x61, 0x69, 0x67, 0x6e, 0x4c, 0x69, 0x66, 0x65, 0x63, 0x79, 0x63, 0x6c, 0x65,
	0x47, 0x6f, 0x61, 0x6c, 0x48, 0x00, 0x52, 0x06, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x12, 0x53,
	0x0a, 0x06, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x39,
	0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x61, 0x64, 0x73, 0x2e, 0x67, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x61, 0x64, 0x73, 0x2e, 0x76, 0x31, 0x39, 0x2e, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72,
	0x63, 0x65, 0x73, 0x2e, 0x43, 0x61, 0x6d, 0x70, 0x61, 0x69, 0x67, 0x6e, 0x4c, 0x69, 0x66, 0x65,
	0x63, 0x79, 0x63, 0x6c, 0x65, 0x47, 0x6f, 0x61, 0x6c, 0x48, 0x00, 0x52, 0x06, 0x75, 0x70, 0x64,
	0x61, 0x74, 0x65, 0x42, 0x0b, 0x0a, 0x09, 0x6f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x22, 0x8b, 0x01, 0x0a, 0x27, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x75, 0x72, 0x65, 0x43, 0x61,
	0x6d, 0x70, 0x61, 0x69, 0x67, 0x6e, 0x4c, 0x69, 0x66, 0x65, 0x63, 0x79, 0x63, 0x6c, 0x65, 0x47,
	0x6f, 0x61, 0x6c, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x60, 0x0a, 0x06,
	0x72, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x48, 0x2e, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x61, 0x64, 0x73, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x61, 0x64, 0x73, 0x2e, 0x76, 0x31, 0x39, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73,
	0x2e, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x75, 0x72, 0x65, 0x43, 0x61, 0x6d, 0x70, 0x61, 0x69,
	0x67, 0x6e, 0x4c, 0x69, 0x66, 0x65, 0x63, 0x79, 0x63, 0x6c, 0x65, 0x47, 0x6f, 0x61, 0x6c, 0x73,
	0x52, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x52, 0x06, 0x72, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x22, 0x81,
	0x01, 0x0a, 0x25, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x75, 0x72, 0x65, 0x43, 0x61, 0x6d, 0x70,
	0x61, 0x69, 0x67, 0x6e, 0x4c, 0x69, 0x66, 0x65, 0x63, 0x79, 0x63, 0x6c, 0x65, 0x47, 0x6f, 0x61,
	0x6c, 0x73, 0x52, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x12, 0x58, 0x0a, 0x0d, 0x72, 0x65, 0x73, 0x6f,
	0x75, 0x72, 0x63, 0x65, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x42,
	0x33, 0xfa, 0x41, 0x30, 0x0a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61, 0x64, 0x73, 0x2e,
	0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61, 0x70, 0x69, 0x73, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x43,
	0x61, 0x6d, 0x70, 0x61, 0x69, 0x67, 0x6e, 0x4c, 0x69, 0x66, 0x65, 0x63, 0x79, 0x63, 0x6c, 0x65,
	0x47, 0x6f, 0x61, 0x6c, 0x52, 0x0c, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x4e, 0x61,
	0x6d, 0x65, 0x32, 0x99, 0x03, 0x0a, 0x1c, 0x43, 0x61, 0x6d, 0x70, 0x61, 0x69, 0x67, 0x6e, 0x4c,
	0x69, 0x66, 0x65, 0x63, 0x79, 0x63, 0x6c, 0x65, 0x47, 0x6f, 0x61, 0x6c, 0x53, 0x65, 0x72, 0x76,
	0x69, 0x63, 0x65, 0x12, 0xb1, 0x02, 0x0a, 0x1f, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x75, 0x72,
	0x65, 0x43, 0x61, 0x6d, 0x70, 0x61, 0x69, 0x67, 0x6e, 0x4c, 0x69, 0x66, 0x65, 0x63, 0x79, 0x63,
	0x6c, 0x65, 0x47, 0x6f, 0x61, 0x6c, 0x73, 0x12, 0x49, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x2e, 0x61, 0x64, 0x73, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61, 0x64, 0x73, 0x2e, 0x76,
	0x31, 0x39, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x43, 0x6f, 0x6e, 0x66,
	0x69, 0x67, 0x75, 0x72, 0x65, 0x43, 0x61, 0x6d, 0x70, 0x61, 0x69, 0x67, 0x6e, 0x4c, 0x69, 0x66,
	0x65, 0x63, 0x79, 0x63, 0x6c, 0x65, 0x47, 0x6f, 0x61, 0x6c, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x1a, 0x4a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x61, 0x64, 0x73, 0x2e,
	0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61, 0x64, 0x73, 0x2e, 0x76, 0x31, 0x39, 0x2e, 0x73, 0x65,
	0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x75, 0x72, 0x65,
	0x43, 0x61, 0x6d, 0x70, 0x61, 0x69, 0x67, 0x6e, 0x4c, 0x69, 0x66, 0x65, 0x63, 0x79, 0x63, 0x6c,
	0x65, 0x47, 0x6f, 0x61, 0x6c, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x77,
	0xda, 0x41, 0x15, 0x63, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x2c, 0x6f,
	0x70, 0x65, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x59, 0x3a, 0x01,
	0x2a, 0x22, 0x54, 0x2f, 0x76, 0x31, 0x39, 0x2f, 0x63, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x65, 0x72,
	0x73, 0x2f, 0x7b, 0x63, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x3d, 0x2a,
	0x7d, 0x2f, 0x63, 0x61, 0x6d, 0x70, 0x61, 0x69, 0x67, 0x6e, 0x4c, 0x69, 0x66, 0x65, 0x63, 0x79,
	0x63, 0x6c, 0x65, 0x47, 0x6f, 0x61, 0x6c, 0x3a, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x75, 0x72,
	0x65, 0x43, 0x61, 0x6d, 0x70, 0x61, 0x69, 0x67, 0x6e, 0x4c, 0x69, 0x66, 0x65, 0x63, 0x79, 0x63,
	0x6c, 0x65, 0x47, 0x6f, 0x61, 0x6c, 0x73, 0x1a, 0x45, 0xca, 0x41, 0x18, 0x67, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x61, 0x64, 0x73, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61, 0x70, 0x69, 0x73,
	0x2e, 0x63, 0x6f, 0x6d, 0xd2, 0x41, 0x27, 0x68, 0x74, 0x74, 0x70, 0x73, 0x3a, 0x2f, 0x2f, 0x77,
	0x77, 0x77, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61, 0x70, 0x69, 0x73, 0x2e, 0x63, 0x6f,
	0x6d, 0x2f, 0x61, 0x75, 0x74, 0x68, 0x2f, 0x61, 0x64, 0x77, 0x6f, 0x72, 0x64, 0x73, 0x42, 0x8d,
	0x02, 0x0a, 0x25, 0x63, 0x6f, 0x6d, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x61, 0x64,
	0x73, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61, 0x64, 0x73, 0x2e, 0x76, 0x31, 0x39, 0x2e,
	0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x42, 0x21, 0x43, 0x61, 0x6d, 0x70, 0x61, 0x69,
	0x67, 0x6e, 0x4c, 0x69, 0x66, 0x65, 0x63, 0x79, 0x63, 0x6c, 0x65, 0x47, 0x6f, 0x61, 0x6c, 0x53,
	0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a, 0x49, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x67, 0x6f, 0x6c, 0x61, 0x6e, 0x67, 0x2e, 0x6f, 0x72, 0x67,
	0x2f, 0x67, 0x65, 0x6e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x61, 0x70, 0x69, 0x73, 0x2f, 0x61, 0x64, 0x73, 0x2f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61,
	0x64, 0x73, 0x2f, 0x76, 0x31, 0x39, 0x2f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x3b,
	0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0xa2, 0x02, 0x03, 0x47, 0x41, 0x41, 0xaa, 0x02,
	0x21, 0x47, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x41, 0x64, 0x73, 0x2e, 0x47, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x41, 0x64, 0x73, 0x2e, 0x56, 0x31, 0x39, 0x2e, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63,
	0x65, 0x73, 0xca, 0x02, 0x21, 0x47, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x5c, 0x41, 0x64, 0x73, 0x5c,
	0x47, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x41, 0x64, 0x73, 0x5c, 0x56, 0x31, 0x39, 0x5c, 0x53, 0x65,
	0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0xea, 0x02, 0x25, 0x47, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x3a,
	0x3a, 0x41, 0x64, 0x73, 0x3a, 0x3a, 0x47, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x41, 0x64, 0x73, 0x3a,
	0x3a, 0x56, 0x31, 0x39, 0x3a, 0x3a, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x62, 0x06,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
})

var (
	file_services_campaign_lifecycle_goal_service_proto_rawDescOnce sync.Once
	file_services_campaign_lifecycle_goal_service_proto_rawDescData []byte
)

func file_services_campaign_lifecycle_goal_service_proto_rawDescGZIP() []byte {
	file_services_campaign_lifecycle_goal_service_proto_rawDescOnce.Do(func() {
		file_services_campaign_lifecycle_goal_service_proto_rawDescData = protoimpl.X.CompressGZIP(unsafe.Slice(unsafe.StringData(file_services_campaign_lifecycle_goal_service_proto_rawDesc), len(file_services_campaign_lifecycle_goal_service_proto_rawDesc)))
	})
	return file_services_campaign_lifecycle_goal_service_proto_rawDescData
}

var file_services_campaign_lifecycle_goal_service_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_services_campaign_lifecycle_goal_service_proto_goTypes = []any{
	(*ConfigureCampaignLifecycleGoalsRequest)(nil),  // 0: google.ads.googleads.v19.services.ConfigureCampaignLifecycleGoalsRequest
	(*CampaignLifecycleGoalOperation)(nil),          // 1: google.ads.googleads.v19.services.CampaignLifecycleGoalOperation
	(*ConfigureCampaignLifecycleGoalsResponse)(nil), // 2: google.ads.googleads.v19.services.ConfigureCampaignLifecycleGoalsResponse
	(*ConfigureCampaignLifecycleGoalsResult)(nil),   // 3: google.ads.googleads.v19.services.ConfigureCampaignLifecycleGoalsResult
	(*fieldmaskpb.FieldMask)(nil),                   // 4: google.protobuf.FieldMask
	(*resources.CampaignLifecycleGoal)(nil),         // 5: google.ads.googleads.v19.resources.CampaignLifecycleGoal
}
var file_services_campaign_lifecycle_goal_service_proto_depIdxs = []int32{
	1, // 0: google.ads.googleads.v19.services.ConfigureCampaignLifecycleGoalsRequest.operation:type_name -> google.ads.googleads.v19.services.CampaignLifecycleGoalOperation
	4, // 1: google.ads.googleads.v19.services.CampaignLifecycleGoalOperation.update_mask:type_name -> google.protobuf.FieldMask
	5, // 2: google.ads.googleads.v19.services.CampaignLifecycleGoalOperation.create:type_name -> google.ads.googleads.v19.resources.CampaignLifecycleGoal
	5, // 3: google.ads.googleads.v19.services.CampaignLifecycleGoalOperation.update:type_name -> google.ads.googleads.v19.resources.CampaignLifecycleGoal
	3, // 4: google.ads.googleads.v19.services.ConfigureCampaignLifecycleGoalsResponse.result:type_name -> google.ads.googleads.v19.services.ConfigureCampaignLifecycleGoalsResult
	0, // 5: google.ads.googleads.v19.services.CampaignLifecycleGoalService.ConfigureCampaignLifecycleGoals:input_type -> google.ads.googleads.v19.services.ConfigureCampaignLifecycleGoalsRequest
	2, // 6: google.ads.googleads.v19.services.CampaignLifecycleGoalService.ConfigureCampaignLifecycleGoals:output_type -> google.ads.googleads.v19.services.ConfigureCampaignLifecycleGoalsResponse
	6, // [6:7] is the sub-list for method output_type
	5, // [5:6] is the sub-list for method input_type
	5, // [5:5] is the sub-list for extension type_name
	5, // [5:5] is the sub-list for extension extendee
	0, // [0:5] is the sub-list for field type_name
}

func init() { file_services_campaign_lifecycle_goal_service_proto_init() }
func file_services_campaign_lifecycle_goal_service_proto_init() {
	if File_services_campaign_lifecycle_goal_service_proto != nil {
		return
	}
	file_services_campaign_lifecycle_goal_service_proto_msgTypes[1].OneofWrappers = []any{
		(*CampaignLifecycleGoalOperation_Create)(nil),
		(*CampaignLifecycleGoalOperation_Update)(nil),
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: unsafe.Slice(unsafe.StringData(file_services_campaign_lifecycle_goal_service_proto_rawDesc), len(file_services_campaign_lifecycle_goal_service_proto_rawDesc)),
			NumEnums:      0,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_services_campaign_lifecycle_goal_service_proto_goTypes,
		DependencyIndexes: file_services_campaign_lifecycle_goal_service_proto_depIdxs,
		MessageInfos:      file_services_campaign_lifecycle_goal_service_proto_msgTypes,
	}.Build()
	File_services_campaign_lifecycle_goal_service_proto = out.File
	file_services_campaign_lifecycle_goal_service_proto_goTypes = nil
	file_services_campaign_lifecycle_goal_service_proto_depIdxs = nil
}
