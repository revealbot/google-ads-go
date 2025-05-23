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
// source: google/ads/googleads/v19/services/bidding_seasonality_adjustment_service.proto

package services

import (
	enums "github.com/revealbot/google-ads-go/enums"
	resources "github.com/revealbot/google-ads-go/resources"
	_ "google.golang.org/genproto/googleapis/api/annotations"
	status "google.golang.org/genproto/googleapis/rpc/status"
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
// [BiddingSeasonalityAdjustmentService.MutateBiddingSeasonalityAdjustments][google.ads.googleads.v19.services.BiddingSeasonalityAdjustmentService.MutateBiddingSeasonalityAdjustments].
type MutateBiddingSeasonalityAdjustmentsRequest struct {
	state protoimpl.MessageState `protogen:"open.v1"`
	// Required. ID of the customer whose seasonality adjustments are being
	// modified.
	CustomerId string `protobuf:"bytes,1,opt,name=customer_id,json=customerId,proto3" json:"customer_id,omitempty"`
	// Required. The list of operations to perform on individual seasonality
	// adjustments.
	Operations []*BiddingSeasonalityAdjustmentOperation `protobuf:"bytes,2,rep,name=operations,proto3" json:"operations,omitempty"`
	// If true, successful operations will be carried out and invalid
	// operations will return errors. If false, all operations will be carried
	// out in one transaction if and only if they are all valid.
	// Default is false.
	PartialFailure bool `protobuf:"varint,3,opt,name=partial_failure,json=partialFailure,proto3" json:"partial_failure,omitempty"`
	// If true, the request is validated but not executed. Only errors are
	// returned, not results.
	ValidateOnly bool `protobuf:"varint,4,opt,name=validate_only,json=validateOnly,proto3" json:"validate_only,omitempty"`
	// The response content type setting. Determines whether the mutable resource
	// or just the resource name should be returned post mutation.
	ResponseContentType enums.ResponseContentTypeEnum_ResponseContentType `protobuf:"varint,5,opt,name=response_content_type,json=responseContentType,proto3,enum=google.ads.googleads.v19.enums.ResponseContentTypeEnum_ResponseContentType" json:"response_content_type,omitempty"`
	unknownFields       protoimpl.UnknownFields
	sizeCache           protoimpl.SizeCache
}

func (x *MutateBiddingSeasonalityAdjustmentsRequest) Reset() {
	*x = MutateBiddingSeasonalityAdjustmentsRequest{}
	mi := &file_services_bidding_seasonality_adjustment_service_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *MutateBiddingSeasonalityAdjustmentsRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MutateBiddingSeasonalityAdjustmentsRequest) ProtoMessage() {}

func (x *MutateBiddingSeasonalityAdjustmentsRequest) ProtoReflect() protoreflect.Message {
	mi := &file_services_bidding_seasonality_adjustment_service_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MutateBiddingSeasonalityAdjustmentsRequest.ProtoReflect.Descriptor instead.
func (*MutateBiddingSeasonalityAdjustmentsRequest) Descriptor() ([]byte, []int) {
	return file_services_bidding_seasonality_adjustment_service_proto_rawDescGZIP(), []int{0}
}

func (x *MutateBiddingSeasonalityAdjustmentsRequest) GetCustomerId() string {
	if x != nil {
		return x.CustomerId
	}
	return ""
}

func (x *MutateBiddingSeasonalityAdjustmentsRequest) GetOperations() []*BiddingSeasonalityAdjustmentOperation {
	if x != nil {
		return x.Operations
	}
	return nil
}

func (x *MutateBiddingSeasonalityAdjustmentsRequest) GetPartialFailure() bool {
	if x != nil {
		return x.PartialFailure
	}
	return false
}

func (x *MutateBiddingSeasonalityAdjustmentsRequest) GetValidateOnly() bool {
	if x != nil {
		return x.ValidateOnly
	}
	return false
}

func (x *MutateBiddingSeasonalityAdjustmentsRequest) GetResponseContentType() enums.ResponseContentTypeEnum_ResponseContentType {
	if x != nil {
		return x.ResponseContentType
	}
	return enums.ResponseContentTypeEnum_ResponseContentType(0)
}

// A single operation (create, remove, update) on a seasonality adjustment.
type BiddingSeasonalityAdjustmentOperation struct {
	state protoimpl.MessageState `protogen:"open.v1"`
	// FieldMask that determines which resource fields are modified in an update.
	UpdateMask *fieldmaskpb.FieldMask `protobuf:"bytes,4,opt,name=update_mask,json=updateMask,proto3" json:"update_mask,omitempty"`
	// The mutate operation.
	//
	// Types that are valid to be assigned to Operation:
	//
	//	*BiddingSeasonalityAdjustmentOperation_Create
	//	*BiddingSeasonalityAdjustmentOperation_Update
	//	*BiddingSeasonalityAdjustmentOperation_Remove
	Operation     isBiddingSeasonalityAdjustmentOperation_Operation `protobuf_oneof:"operation"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *BiddingSeasonalityAdjustmentOperation) Reset() {
	*x = BiddingSeasonalityAdjustmentOperation{}
	mi := &file_services_bidding_seasonality_adjustment_service_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *BiddingSeasonalityAdjustmentOperation) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BiddingSeasonalityAdjustmentOperation) ProtoMessage() {}

func (x *BiddingSeasonalityAdjustmentOperation) ProtoReflect() protoreflect.Message {
	mi := &file_services_bidding_seasonality_adjustment_service_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use BiddingSeasonalityAdjustmentOperation.ProtoReflect.Descriptor instead.
func (*BiddingSeasonalityAdjustmentOperation) Descriptor() ([]byte, []int) {
	return file_services_bidding_seasonality_adjustment_service_proto_rawDescGZIP(), []int{1}
}

func (x *BiddingSeasonalityAdjustmentOperation) GetUpdateMask() *fieldmaskpb.FieldMask {
	if x != nil {
		return x.UpdateMask
	}
	return nil
}

func (x *BiddingSeasonalityAdjustmentOperation) GetOperation() isBiddingSeasonalityAdjustmentOperation_Operation {
	if x != nil {
		return x.Operation
	}
	return nil
}

func (x *BiddingSeasonalityAdjustmentOperation) GetCreate() *resources.BiddingSeasonalityAdjustment {
	if x != nil {
		if x, ok := x.Operation.(*BiddingSeasonalityAdjustmentOperation_Create); ok {
			return x.Create
		}
	}
	return nil
}

func (x *BiddingSeasonalityAdjustmentOperation) GetUpdate() *resources.BiddingSeasonalityAdjustment {
	if x != nil {
		if x, ok := x.Operation.(*BiddingSeasonalityAdjustmentOperation_Update); ok {
			return x.Update
		}
	}
	return nil
}

func (x *BiddingSeasonalityAdjustmentOperation) GetRemove() string {
	if x != nil {
		if x, ok := x.Operation.(*BiddingSeasonalityAdjustmentOperation_Remove); ok {
			return x.Remove
		}
	}
	return ""
}

type isBiddingSeasonalityAdjustmentOperation_Operation interface {
	isBiddingSeasonalityAdjustmentOperation_Operation()
}

type BiddingSeasonalityAdjustmentOperation_Create struct {
	// Create operation: No resource name is expected for the new seasonality
	// adjustment.
	Create *resources.BiddingSeasonalityAdjustment `protobuf:"bytes,1,opt,name=create,proto3,oneof"`
}

type BiddingSeasonalityAdjustmentOperation_Update struct {
	// Update operation: The seasonality adjustment is expected to have a valid
	// resource name.
	Update *resources.BiddingSeasonalityAdjustment `protobuf:"bytes,2,opt,name=update,proto3,oneof"`
}

type BiddingSeasonalityAdjustmentOperation_Remove struct {
	// Remove operation: A resource name for the removed seasonality adjustment
	// is expected, in this format:
	//
	// `customers/{customer_id}/biddingSeasonalityAdjustments/{seasonality_adjustment_id}`
	Remove string `protobuf:"bytes,3,opt,name=remove,proto3,oneof"`
}

func (*BiddingSeasonalityAdjustmentOperation_Create) isBiddingSeasonalityAdjustmentOperation_Operation() {
}

func (*BiddingSeasonalityAdjustmentOperation_Update) isBiddingSeasonalityAdjustmentOperation_Operation() {
}

func (*BiddingSeasonalityAdjustmentOperation_Remove) isBiddingSeasonalityAdjustmentOperation_Operation() {
}

// Response message for seasonality adjustments mutate.
type MutateBiddingSeasonalityAdjustmentsResponse struct {
	state protoimpl.MessageState `protogen:"open.v1"`
	// Errors that pertain to operation failures in the partial failure mode.
	// Returned only when partial_failure = true and all errors occur inside the
	// operations. If any errors occur outside the operations (for example, auth
	// errors), we return an RPC level error.
	PartialFailureError *status.Status `protobuf:"bytes,3,opt,name=partial_failure_error,json=partialFailureError,proto3" json:"partial_failure_error,omitempty"`
	// All results for the mutate.
	Results       []*MutateBiddingSeasonalityAdjustmentsResult `protobuf:"bytes,2,rep,name=results,proto3" json:"results,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *MutateBiddingSeasonalityAdjustmentsResponse) Reset() {
	*x = MutateBiddingSeasonalityAdjustmentsResponse{}
	mi := &file_services_bidding_seasonality_adjustment_service_proto_msgTypes[2]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *MutateBiddingSeasonalityAdjustmentsResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MutateBiddingSeasonalityAdjustmentsResponse) ProtoMessage() {}

func (x *MutateBiddingSeasonalityAdjustmentsResponse) ProtoReflect() protoreflect.Message {
	mi := &file_services_bidding_seasonality_adjustment_service_proto_msgTypes[2]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MutateBiddingSeasonalityAdjustmentsResponse.ProtoReflect.Descriptor instead.
func (*MutateBiddingSeasonalityAdjustmentsResponse) Descriptor() ([]byte, []int) {
	return file_services_bidding_seasonality_adjustment_service_proto_rawDescGZIP(), []int{2}
}

func (x *MutateBiddingSeasonalityAdjustmentsResponse) GetPartialFailureError() *status.Status {
	if x != nil {
		return x.PartialFailureError
	}
	return nil
}

func (x *MutateBiddingSeasonalityAdjustmentsResponse) GetResults() []*MutateBiddingSeasonalityAdjustmentsResult {
	if x != nil {
		return x.Results
	}
	return nil
}

// The result for the seasonality adjustment mutate.
type MutateBiddingSeasonalityAdjustmentsResult struct {
	state protoimpl.MessageState `protogen:"open.v1"`
	// Returned for successful operations.
	ResourceName string `protobuf:"bytes,1,opt,name=resource_name,json=resourceName,proto3" json:"resource_name,omitempty"`
	// The mutated bidding seasonality adjustment with only mutable fields after
	// mutate. The field will only be returned when response_content_type is set
	// to "MUTABLE_RESOURCE".
	BiddingSeasonalityAdjustment *resources.BiddingSeasonalityAdjustment `protobuf:"bytes,2,opt,name=bidding_seasonality_adjustment,json=biddingSeasonalityAdjustment,proto3" json:"bidding_seasonality_adjustment,omitempty"`
	unknownFields                protoimpl.UnknownFields
	sizeCache                    protoimpl.SizeCache
}

func (x *MutateBiddingSeasonalityAdjustmentsResult) Reset() {
	*x = MutateBiddingSeasonalityAdjustmentsResult{}
	mi := &file_services_bidding_seasonality_adjustment_service_proto_msgTypes[3]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *MutateBiddingSeasonalityAdjustmentsResult) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MutateBiddingSeasonalityAdjustmentsResult) ProtoMessage() {}

func (x *MutateBiddingSeasonalityAdjustmentsResult) ProtoReflect() protoreflect.Message {
	mi := &file_services_bidding_seasonality_adjustment_service_proto_msgTypes[3]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MutateBiddingSeasonalityAdjustmentsResult.ProtoReflect.Descriptor instead.
func (*MutateBiddingSeasonalityAdjustmentsResult) Descriptor() ([]byte, []int) {
	return file_services_bidding_seasonality_adjustment_service_proto_rawDescGZIP(), []int{3}
}

func (x *MutateBiddingSeasonalityAdjustmentsResult) GetResourceName() string {
	if x != nil {
		return x.ResourceName
	}
	return ""
}

func (x *MutateBiddingSeasonalityAdjustmentsResult) GetBiddingSeasonalityAdjustment() *resources.BiddingSeasonalityAdjustment {
	if x != nil {
		return x.BiddingSeasonalityAdjustment
	}
	return nil
}

var File_services_bidding_seasonality_adjustment_service_proto protoreflect.FileDescriptor

var file_services_bidding_seasonality_adjustment_service_proto_rawDesc = string([]byte{
	0x0a, 0x4e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x64, 0x73, 0x2f, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x61, 0x64, 0x73, 0x2f, 0x76, 0x31, 0x39, 0x2f, 0x73, 0x65, 0x72, 0x76, 0x69,
	0x63, 0x65, 0x73, 0x2f, 0x62, 0x69, 0x64, 0x64, 0x69, 0x6e, 0x67, 0x5f, 0x73, 0x65, 0x61, 0x73,
	0x6f, 0x6e, 0x61, 0x6c, 0x69, 0x74, 0x79, 0x5f, 0x61, 0x64, 0x6a, 0x75, 0x73, 0x74, 0x6d, 0x65,
	0x6e, 0x74, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x12, 0x21, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x61, 0x64, 0x73, 0x2e, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x61, 0x64, 0x73, 0x2e, 0x76, 0x31, 0x39, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69,
	0x63, 0x65, 0x73, 0x1a, 0x3a, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x64, 0x73, 0x2f,
	0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61, 0x64, 0x73, 0x2f, 0x76, 0x31, 0x39, 0x2f, 0x65, 0x6e,
	0x75, 0x6d, 0x73, 0x2f, 0x72, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x5f, 0x63, 0x6f, 0x6e,
	0x74, 0x65, 0x6e, 0x74, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a,
	0x47, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x64, 0x73, 0x2f, 0x67, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x61, 0x64, 0x73, 0x2f, 0x76, 0x31, 0x39, 0x2f, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72,
	0x63, 0x65, 0x73, 0x2f, 0x62, 0x69, 0x64, 0x64, 0x69, 0x6e, 0x67, 0x5f, 0x73, 0x65, 0x61, 0x73,
	0x6f, 0x6e, 0x61, 0x6c, 0x69, 0x74, 0x79, 0x5f, 0x61, 0x64, 0x6a, 0x75, 0x73, 0x74, 0x6d, 0x65,
	0x6e, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1c, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x2f, 0x61, 0x70, 0x69, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x17, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61,
	0x70, 0x69, 0x2f, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a,
	0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x66, 0x69, 0x65, 0x6c,
	0x64, 0x5f, 0x62, 0x65, 0x68, 0x61, 0x76, 0x69, 0x6f, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x1a, 0x19, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x72, 0x65, 0x73,
	0x6f, 0x75, 0x72, 0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x20, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x66, 0x69, 0x65,
	0x6c, 0x64, 0x5f, 0x6d, 0x61, 0x73, 0x6b, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x17, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x72, 0x70, 0x63, 0x2f, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x90, 0x03, 0x0a, 0x2a, 0x4d, 0x75, 0x74, 0x61, 0x74,
	0x65, 0x42, 0x69, 0x64, 0x64, 0x69, 0x6e, 0x67, 0x53, 0x65, 0x61, 0x73, 0x6f, 0x6e, 0x61, 0x6c,
	0x69, 0x74, 0x79, 0x41, 0x64, 0x6a, 0x75, 0x73, 0x74, 0x6d, 0x65, 0x6e, 0x74, 0x73, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x24, 0x0a, 0x0b, 0x63, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x65,
	0x72, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x42, 0x03, 0xe0, 0x41, 0x02, 0x52,
	0x0a, 0x63, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x65, 0x72, 0x49, 0x64, 0x12, 0x6d, 0x0a, 0x0a, 0x6f,
	0x70, 0x65, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32,
	0x48, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x61, 0x64, 0x73, 0x2e, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x61, 0x64, 0x73, 0x2e, 0x76, 0x31, 0x39, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69,
	0x63, 0x65, 0x73, 0x2e, 0x42, 0x69, 0x64, 0x64, 0x69, 0x6e, 0x67, 0x53, 0x65, 0x61, 0x73, 0x6f,
	0x6e, 0x61, 0x6c, 0x69, 0x74, 0x79, 0x41, 0x64, 0x6a, 0x75, 0x73, 0x74, 0x6d, 0x65, 0x6e, 0x74,
	0x4f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x42, 0x03, 0xe0, 0x41, 0x02, 0x52, 0x0a,
	0x6f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x12, 0x27, 0x0a, 0x0f, 0x70, 0x61,
	0x72, 0x74, 0x69, 0x61, 0x6c, 0x5f, 0x66, 0x61, 0x69, 0x6c, 0x75, 0x72, 0x65, 0x18, 0x03, 0x20,
	0x01, 0x28, 0x08, 0x52, 0x0e, 0x70, 0x61, 0x72, 0x74, 0x69, 0x61, 0x6c, 0x46, 0x61, 0x69, 0x6c,
	0x75, 0x72, 0x65, 0x12, 0x23, 0x0a, 0x0d, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x65, 0x5f,
	0x6f, 0x6e, 0x6c, 0x79, 0x18, 0x04, 0x20, 0x01, 0x28, 0x08, 0x52, 0x0c, 0x76, 0x61, 0x6c, 0x69,
	0x64, 0x61, 0x74, 0x65, 0x4f, 0x6e, 0x6c, 0x79, 0x12, 0x7f, 0x0a, 0x15, 0x72, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x5f, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x5f, 0x74, 0x79, 0x70,
	0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x4b, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x2e, 0x61, 0x64, 0x73, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61, 0x64, 0x73, 0x2e, 0x76,
	0x31, 0x39, 0x2e, 0x65, 0x6e, 0x75, 0x6d, 0x73, 0x2e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x43, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x54, 0x79, 0x70, 0x65, 0x45, 0x6e, 0x75, 0x6d,
	0x2e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x43, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74,
	0x54, 0x79, 0x70, 0x65, 0x52, 0x13, 0x72, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x43, 0x6f,
	0x6e, 0x74, 0x65, 0x6e, 0x74, 0x54, 0x79, 0x70, 0x65, 0x22, 0xff, 0x02, 0x0a, 0x25, 0x42, 0x69,
	0x64, 0x64, 0x69, 0x6e, 0x67, 0x53, 0x65, 0x61, 0x73, 0x6f, 0x6e, 0x61, 0x6c, 0x69, 0x74, 0x79,
	0x41, 0x64, 0x6a, 0x75, 0x73, 0x74, 0x6d, 0x65, 0x6e, 0x74, 0x4f, 0x70, 0x65, 0x72, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x12, 0x3b, 0x0a, 0x0b, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x5f, 0x6d, 0x61,
	0x73, 0x6b, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x46, 0x69, 0x65, 0x6c, 0x64,
	0x4d, 0x61, 0x73, 0x6b, 0x52, 0x0a, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x4d, 0x61, 0x73, 0x6b,
	0x12, 0x5a, 0x0a, 0x06, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x40, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x61, 0x64, 0x73, 0x2e, 0x67, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x61, 0x64, 0x73, 0x2e, 0x76, 0x31, 0x39, 0x2e, 0x72, 0x65, 0x73, 0x6f,
	0x75, 0x72, 0x63, 0x65, 0x73, 0x2e, 0x42, 0x69, 0x64, 0x64, 0x69, 0x6e, 0x67, 0x53, 0x65, 0x61,
	0x73, 0x6f, 0x6e, 0x61, 0x6c, 0x69, 0x74, 0x79, 0x41, 0x64, 0x6a, 0x75, 0x73, 0x74, 0x6d, 0x65,
	0x6e, 0x74, 0x48, 0x00, 0x52, 0x06, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x12, 0x5a, 0x0a, 0x06,
	0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x40, 0x2e, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x61, 0x64, 0x73, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x61, 0x64, 0x73, 0x2e, 0x76, 0x31, 0x39, 0x2e, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65,
	0x73, 0x2e, 0x42, 0x69, 0x64, 0x64, 0x69, 0x6e, 0x67, 0x53, 0x65, 0x61, 0x73, 0x6f, 0x6e, 0x61,
	0x6c, 0x69, 0x74, 0x79, 0x41, 0x64, 0x6a, 0x75, 0x73, 0x74, 0x6d, 0x65, 0x6e, 0x74, 0x48, 0x00,
	0x52, 0x06, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x12, 0x54, 0x0a, 0x06, 0x72, 0x65, 0x6d, 0x6f,
	0x76, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x42, 0x3a, 0xfa, 0x41, 0x37, 0x0a, 0x35, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61, 0x64, 0x73, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61,
	0x70, 0x69, 0x73, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x42, 0x69, 0x64, 0x64, 0x69, 0x6e, 0x67, 0x53,
	0x65, 0x61, 0x73, 0x6f, 0x6e, 0x61, 0x6c, 0x69, 0x74, 0x79, 0x41, 0x64, 0x6a, 0x75, 0x73, 0x74,
	0x6d, 0x65, 0x6e, 0x74, 0x48, 0x00, 0x52, 0x06, 0x72, 0x65, 0x6d, 0x6f, 0x76, 0x65, 0x42, 0x0b,
	0x0a, 0x09, 0x6f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x22, 0xdd, 0x01, 0x0a, 0x2b,
	0x4d, 0x75, 0x74, 0x61, 0x74, 0x65, 0x42, 0x69, 0x64, 0x64, 0x69, 0x6e, 0x67, 0x53, 0x65, 0x61,
	0x73, 0x6f, 0x6e, 0x61, 0x6c, 0x69, 0x74, 0x79, 0x41, 0x64, 0x6a, 0x75, 0x73, 0x74, 0x6d, 0x65,
	0x6e, 0x74, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x46, 0x0a, 0x15, 0x70,
	0x61, 0x72, 0x74, 0x69, 0x61, 0x6c, 0x5f, 0x66, 0x61, 0x69, 0x6c, 0x75, 0x72, 0x65, 0x5f, 0x65,
	0x72, 0x72, 0x6f, 0x72, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x12, 0x2e, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x2e, 0x72, 0x70, 0x63, 0x2e, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x52, 0x13,
	0x70, 0x61, 0x72, 0x74, 0x69, 0x61, 0x6c, 0x46, 0x61, 0x69, 0x6c, 0x75, 0x72, 0x65, 0x45, 0x72,
	0x72, 0x6f, 0x72, 0x12, 0x66, 0x0a, 0x07, 0x72, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x73, 0x18, 0x02,
	0x20, 0x03, 0x28, 0x0b, 0x32, 0x4c, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x61, 0x64,
	0x73, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61, 0x64, 0x73, 0x2e, 0x76, 0x31, 0x39, 0x2e,
	0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x4d, 0x75, 0x74, 0x61, 0x74, 0x65, 0x42,
	0x69, 0x64, 0x64, 0x69, 0x6e, 0x67, 0x53, 0x65, 0x61, 0x73, 0x6f, 0x6e, 0x61, 0x6c, 0x69, 0x74,
	0x79, 0x41, 0x64, 0x6a, 0x75, 0x73, 0x74, 0x6d, 0x65, 0x6e, 0x74, 0x73, 0x52, 0x65, 0x73, 0x75,
	0x6c, 0x74, 0x52, 0x07, 0x72, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x73, 0x22, 0x95, 0x02, 0x0a, 0x29,
	0x4d, 0x75, 0x74, 0x61, 0x74, 0x65, 0x42, 0x69, 0x64, 0x64, 0x69, 0x6e, 0x67, 0x53, 0x65, 0x61,
	0x73, 0x6f, 0x6e, 0x61, 0x6c, 0x69, 0x74, 0x79, 0x41, 0x64, 0x6a, 0x75, 0x73, 0x74, 0x6d, 0x65,
	0x6e, 0x74, 0x73, 0x52, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x12, 0x5f, 0x0a, 0x0d, 0x72, 0x65, 0x73,
	0x6f, 0x75, 0x72, 0x63, 0x65, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x42, 0x3a, 0xfa, 0x41, 0x37, 0x0a, 0x35, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61, 0x64, 0x73,
	0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61, 0x70, 0x69, 0x73, 0x2e, 0x63, 0x6f, 0x6d, 0x2f,
	0x42, 0x69, 0x64, 0x64, 0x69, 0x6e, 0x67, 0x53, 0x65, 0x61, 0x73, 0x6f, 0x6e, 0x61, 0x6c, 0x69,
	0x74, 0x79, 0x41, 0x64, 0x6a, 0x75, 0x73, 0x74, 0x6d, 0x65, 0x6e, 0x74, 0x52, 0x0c, 0x72, 0x65,
	0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x86, 0x01, 0x0a, 0x1e, 0x62,
	0x69, 0x64, 0x64, 0x69, 0x6e, 0x67, 0x5f, 0x73, 0x65, 0x61, 0x73, 0x6f, 0x6e, 0x61, 0x6c, 0x69,
	0x74, 0x79, 0x5f, 0x61, 0x64, 0x6a, 0x75, 0x73, 0x74, 0x6d, 0x65, 0x6e, 0x74, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x40, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x61, 0x64, 0x73,
	0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61, 0x64, 0x73, 0x2e, 0x76, 0x31, 0x39, 0x2e, 0x72,
	0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x73, 0x2e, 0x42, 0x69, 0x64, 0x64, 0x69, 0x6e, 0x67,
	0x53, 0x65, 0x61, 0x73, 0x6f, 0x6e, 0x61, 0x6c, 0x69, 0x74, 0x79, 0x41, 0x64, 0x6a, 0x75, 0x73,
	0x74, 0x6d, 0x65, 0x6e, 0x74, 0x52, 0x1c, 0x62, 0x69, 0x64, 0x64, 0x69, 0x6e, 0x67, 0x53, 0x65,
	0x61, 0x73, 0x6f, 0x6e, 0x61, 0x6c, 0x69, 0x74, 0x79, 0x41, 0x64, 0x6a, 0x75, 0x73, 0x74, 0x6d,
	0x65, 0x6e, 0x74, 0x32, 0x9c, 0x03, 0x0a, 0x23, 0x42, 0x69, 0x64, 0x64, 0x69, 0x6e, 0x67, 0x53,
	0x65, 0x61, 0x73, 0x6f, 0x6e, 0x61, 0x6c, 0x69, 0x74, 0x79, 0x41, 0x64, 0x6a, 0x75, 0x73, 0x74,
	0x6d, 0x65, 0x6e, 0x74, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0xad, 0x02, 0x0a, 0x23,
	0x4d, 0x75, 0x74, 0x61, 0x74, 0x65, 0x42, 0x69, 0x64, 0x64, 0x69, 0x6e, 0x67, 0x53, 0x65, 0x61,
	0x73, 0x6f, 0x6e, 0x61, 0x6c, 0x69, 0x74, 0x79, 0x41, 0x64, 0x6a, 0x75, 0x73, 0x74, 0x6d, 0x65,
	0x6e, 0x74, 0x73, 0x12, 0x4d, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x61, 0x64, 0x73,
	0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61, 0x64, 0x73, 0x2e, 0x76, 0x31, 0x39, 0x2e, 0x73,
	0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x4d, 0x75, 0x74, 0x61, 0x74, 0x65, 0x42, 0x69,
	0x64, 0x64, 0x69, 0x6e, 0x67, 0x53, 0x65, 0x61, 0x73, 0x6f, 0x6e, 0x61, 0x6c, 0x69, 0x74, 0x79,
	0x41, 0x64, 0x6a, 0x75, 0x73, 0x74, 0x6d, 0x65, 0x6e, 0x74, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x1a, 0x4e, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x61, 0x64, 0x73, 0x2e,
	0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61, 0x64, 0x73, 0x2e, 0x76, 0x31, 0x39, 0x2e, 0x73, 0x65,
	0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x4d, 0x75, 0x74, 0x61, 0x74, 0x65, 0x42, 0x69, 0x64,
	0x64, 0x69, 0x6e, 0x67, 0x53, 0x65, 0x61, 0x73, 0x6f, 0x6e, 0x61, 0x6c, 0x69, 0x74, 0x79, 0x41,
	0x64, 0x6a, 0x75, 0x73, 0x74, 0x6d, 0x65, 0x6e, 0x74, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x22, 0x67, 0xda, 0x41, 0x16, 0x63, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x65, 0x72, 0x5f,
	0x69, 0x64, 0x2c, 0x6f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x82, 0xd3, 0xe4,
	0x93, 0x02, 0x48, 0x3a, 0x01, 0x2a, 0x22, 0x43, 0x2f, 0x76, 0x31, 0x39, 0x2f, 0x63, 0x75, 0x73,
	0x74, 0x6f, 0x6d, 0x65, 0x72, 0x73, 0x2f, 0x7b, 0x63, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x65, 0x72,
	0x5f, 0x69, 0x64, 0x3d, 0x2a, 0x7d, 0x2f, 0x62, 0x69, 0x64, 0x64, 0x69, 0x6e, 0x67, 0x53, 0x65,
	0x61, 0x73, 0x6f, 0x6e, 0x61, 0x6c, 0x69, 0x74, 0x79, 0x41, 0x64, 0x6a, 0x75, 0x73, 0x74, 0x6d,
	0x65, 0x6e, 0x74, 0x73, 0x3a, 0x6d, 0x75, 0x74, 0x61, 0x74, 0x65, 0x1a, 0x45, 0xca, 0x41, 0x18,
	0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61, 0x64, 0x73, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x61, 0x70, 0x69, 0x73, 0x2e, 0x63, 0x6f, 0x6d, 0xd2, 0x41, 0x27, 0x68, 0x74, 0x74, 0x70, 0x73,
	0x3a, 0x2f, 0x2f, 0x77, 0x77, 0x77, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61, 0x70, 0x69,
	0x73, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x61, 0x75, 0x74, 0x68, 0x2f, 0x61, 0x64, 0x77, 0x6f, 0x72,
	0x64, 0x73, 0x42, 0x94, 0x02, 0x0a, 0x25, 0x63, 0x6f, 0x6d, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x2e, 0x61, 0x64, 0x73, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61, 0x64, 0x73, 0x2e,
	0x76, 0x31, 0x39, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x42, 0x28, 0x42, 0x69,
	0x64, 0x64, 0x69, 0x6e, 0x67, 0x53, 0x65, 0x61, 0x73, 0x6f, 0x6e, 0x61, 0x6c, 0x69, 0x74, 0x79,
	0x41, 0x64, 0x6a, 0x75, 0x73, 0x74, 0x6d, 0x65, 0x6e, 0x74, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63,
	0x65, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a, 0x49, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x2e, 0x67, 0x6f, 0x6c, 0x61, 0x6e, 0x67, 0x2e, 0x6f, 0x72, 0x67, 0x2f, 0x67, 0x65, 0x6e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61, 0x70, 0x69, 0x73, 0x2f,
	0x61, 0x64, 0x73, 0x2f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61, 0x64, 0x73, 0x2f, 0x76, 0x31,
	0x39, 0x2f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x3b, 0x73, 0x65, 0x72, 0x76, 0x69,
	0x63, 0x65, 0x73, 0xa2, 0x02, 0x03, 0x47, 0x41, 0x41, 0xaa, 0x02, 0x21, 0x47, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x2e, 0x41, 0x64, 0x73, 0x2e, 0x47, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x41, 0x64, 0x73,
	0x2e, 0x56, 0x31, 0x39, 0x2e, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0xca, 0x02, 0x21,
	0x47, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x5c, 0x41, 0x64, 0x73, 0x5c, 0x47, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x41, 0x64, 0x73, 0x5c, 0x56, 0x31, 0x39, 0x5c, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65,
	0x73, 0xea, 0x02, 0x25, 0x47, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x3a, 0x3a, 0x41, 0x64, 0x73, 0x3a,
	0x3a, 0x47, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x41, 0x64, 0x73, 0x3a, 0x3a, 0x56, 0x31, 0x39, 0x3a,
	0x3a, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x33,
})

var (
	file_services_bidding_seasonality_adjustment_service_proto_rawDescOnce sync.Once
	file_services_bidding_seasonality_adjustment_service_proto_rawDescData []byte
)

func file_services_bidding_seasonality_adjustment_service_proto_rawDescGZIP() []byte {
	file_services_bidding_seasonality_adjustment_service_proto_rawDescOnce.Do(func() {
		file_services_bidding_seasonality_adjustment_service_proto_rawDescData = protoimpl.X.CompressGZIP(unsafe.Slice(unsafe.StringData(file_services_bidding_seasonality_adjustment_service_proto_rawDesc), len(file_services_bidding_seasonality_adjustment_service_proto_rawDesc)))
	})
	return file_services_bidding_seasonality_adjustment_service_proto_rawDescData
}

var file_services_bidding_seasonality_adjustment_service_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_services_bidding_seasonality_adjustment_service_proto_goTypes = []any{
	(*MutateBiddingSeasonalityAdjustmentsRequest)(nil),     // 0: google.ads.googleads.v19.services.MutateBiddingSeasonalityAdjustmentsRequest
	(*BiddingSeasonalityAdjustmentOperation)(nil),          // 1: google.ads.googleads.v19.services.BiddingSeasonalityAdjustmentOperation
	(*MutateBiddingSeasonalityAdjustmentsResponse)(nil),    // 2: google.ads.googleads.v19.services.MutateBiddingSeasonalityAdjustmentsResponse
	(*MutateBiddingSeasonalityAdjustmentsResult)(nil),      // 3: google.ads.googleads.v19.services.MutateBiddingSeasonalityAdjustmentsResult
	(enums.ResponseContentTypeEnum_ResponseContentType)(0), // 4: google.ads.googleads.v19.enums.ResponseContentTypeEnum.ResponseContentType
	(*fieldmaskpb.FieldMask)(nil),                          // 5: google.protobuf.FieldMask
	(*resources.BiddingSeasonalityAdjustment)(nil),         // 6: google.ads.googleads.v19.resources.BiddingSeasonalityAdjustment
	(*status.Status)(nil),                                  // 7: google.rpc.Status
}
var file_services_bidding_seasonality_adjustment_service_proto_depIdxs = []int32{
	1, // 0: google.ads.googleads.v19.services.MutateBiddingSeasonalityAdjustmentsRequest.operations:type_name -> google.ads.googleads.v19.services.BiddingSeasonalityAdjustmentOperation
	4, // 1: google.ads.googleads.v19.services.MutateBiddingSeasonalityAdjustmentsRequest.response_content_type:type_name -> google.ads.googleads.v19.enums.ResponseContentTypeEnum.ResponseContentType
	5, // 2: google.ads.googleads.v19.services.BiddingSeasonalityAdjustmentOperation.update_mask:type_name -> google.protobuf.FieldMask
	6, // 3: google.ads.googleads.v19.services.BiddingSeasonalityAdjustmentOperation.create:type_name -> google.ads.googleads.v19.resources.BiddingSeasonalityAdjustment
	6, // 4: google.ads.googleads.v19.services.BiddingSeasonalityAdjustmentOperation.update:type_name -> google.ads.googleads.v19.resources.BiddingSeasonalityAdjustment
	7, // 5: google.ads.googleads.v19.services.MutateBiddingSeasonalityAdjustmentsResponse.partial_failure_error:type_name -> google.rpc.Status
	3, // 6: google.ads.googleads.v19.services.MutateBiddingSeasonalityAdjustmentsResponse.results:type_name -> google.ads.googleads.v19.services.MutateBiddingSeasonalityAdjustmentsResult
	6, // 7: google.ads.googleads.v19.services.MutateBiddingSeasonalityAdjustmentsResult.bidding_seasonality_adjustment:type_name -> google.ads.googleads.v19.resources.BiddingSeasonalityAdjustment
	0, // 8: google.ads.googleads.v19.services.BiddingSeasonalityAdjustmentService.MutateBiddingSeasonalityAdjustments:input_type -> google.ads.googleads.v19.services.MutateBiddingSeasonalityAdjustmentsRequest
	2, // 9: google.ads.googleads.v19.services.BiddingSeasonalityAdjustmentService.MutateBiddingSeasonalityAdjustments:output_type -> google.ads.googleads.v19.services.MutateBiddingSeasonalityAdjustmentsResponse
	9, // [9:10] is the sub-list for method output_type
	8, // [8:9] is the sub-list for method input_type
	8, // [8:8] is the sub-list for extension type_name
	8, // [8:8] is the sub-list for extension extendee
	0, // [0:8] is the sub-list for field type_name
}

func init() {
	file_services_bidding_seasonality_adjustment_service_proto_init()
}
func file_services_bidding_seasonality_adjustment_service_proto_init() {
	if File_services_bidding_seasonality_adjustment_service_proto != nil {
		return
	}
	file_services_bidding_seasonality_adjustment_service_proto_msgTypes[1].OneofWrappers = []any{
		(*BiddingSeasonalityAdjustmentOperation_Create)(nil),
		(*BiddingSeasonalityAdjustmentOperation_Update)(nil),
		(*BiddingSeasonalityAdjustmentOperation_Remove)(nil),
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: unsafe.Slice(unsafe.StringData(file_services_bidding_seasonality_adjustment_service_proto_rawDesc), len(file_services_bidding_seasonality_adjustment_service_proto_rawDesc)),
			NumEnums:      0,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_services_bidding_seasonality_adjustment_service_proto_goTypes,
		DependencyIndexes: file_services_bidding_seasonality_adjustment_service_proto_depIdxs,
		MessageInfos:      file_services_bidding_seasonality_adjustment_service_proto_msgTypes,
	}.Build()
	File_services_bidding_seasonality_adjustment_service_proto = out.File
	file_services_bidding_seasonality_adjustment_service_proto_goTypes = nil
	file_services_bidding_seasonality_adjustment_service_proto_depIdxs = nil
}
