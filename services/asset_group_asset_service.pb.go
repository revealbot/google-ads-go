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
// source: google/ads/googleads/v16/services/asset_group_asset_service.proto

package services

import (
	resources "github.com/revealbot/google-ads-go/resources"
	_ "google.golang.org/genproto/googleapis/api/annotations"
	status "google.golang.org/genproto/googleapis/rpc/status"
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
// [AssetGroupAssetService.MutateAssetGroupAssets][google.ads.googleads.v16.services.AssetGroupAssetService.MutateAssetGroupAssets].
type MutateAssetGroupAssetsRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Required. The ID of the customer whose asset group assets are being
	// modified.
	CustomerId string `protobuf:"bytes,1,opt,name=customer_id,json=customerId,proto3" json:"customer_id,omitempty"`
	// Required. The list of operations to perform on individual asset group
	// assets.
	Operations []*AssetGroupAssetOperation `protobuf:"bytes,2,rep,name=operations,proto3" json:"operations,omitempty"`
	// If true, successful operations will be carried out and invalid
	// operations will return errors. If false, all operations will be carried
	// out in one transaction if and only if they are all valid.
	// Default is false.
	PartialFailure bool `protobuf:"varint,3,opt,name=partial_failure,json=partialFailure,proto3" json:"partial_failure,omitempty"`
	// If true, the request is validated but not executed. Only errors are
	// returned, not results.
	ValidateOnly bool `protobuf:"varint,4,opt,name=validate_only,json=validateOnly,proto3" json:"validate_only,omitempty"`
}

func (x *MutateAssetGroupAssetsRequest) Reset() {
	*x = MutateAssetGroupAssetsRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_services_asset_group_asset_service_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MutateAssetGroupAssetsRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MutateAssetGroupAssetsRequest) ProtoMessage() {}

func (x *MutateAssetGroupAssetsRequest) ProtoReflect() protoreflect.Message {
	mi := &file_services_asset_group_asset_service_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MutateAssetGroupAssetsRequest.ProtoReflect.Descriptor instead.
func (*MutateAssetGroupAssetsRequest) Descriptor() ([]byte, []int) {
	return file_services_asset_group_asset_service_proto_rawDescGZIP(), []int{0}
}

func (x *MutateAssetGroupAssetsRequest) GetCustomerId() string {
	if x != nil {
		return x.CustomerId
	}
	return ""
}

func (x *MutateAssetGroupAssetsRequest) GetOperations() []*AssetGroupAssetOperation {
	if x != nil {
		return x.Operations
	}
	return nil
}

func (x *MutateAssetGroupAssetsRequest) GetPartialFailure() bool {
	if x != nil {
		return x.PartialFailure
	}
	return false
}

func (x *MutateAssetGroupAssetsRequest) GetValidateOnly() bool {
	if x != nil {
		return x.ValidateOnly
	}
	return false
}

// A single operation (create, remove) on an asset group asset.
type AssetGroupAssetOperation struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// FieldMask that determines which resource fields are modified in an update.
	UpdateMask *fieldmaskpb.FieldMask `protobuf:"bytes,4,opt,name=update_mask,json=updateMask,proto3" json:"update_mask,omitempty"`
	// The mutate operation.
	//
	// Types that are assignable to Operation:
	//
	//	*AssetGroupAssetOperation_Create
	//	*AssetGroupAssetOperation_Update
	//	*AssetGroupAssetOperation_Remove
	Operation isAssetGroupAssetOperation_Operation `protobuf_oneof:"operation"`
}

func (x *AssetGroupAssetOperation) Reset() {
	*x = AssetGroupAssetOperation{}
	if protoimpl.UnsafeEnabled {
		mi := &file_services_asset_group_asset_service_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AssetGroupAssetOperation) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AssetGroupAssetOperation) ProtoMessage() {}

func (x *AssetGroupAssetOperation) ProtoReflect() protoreflect.Message {
	mi := &file_services_asset_group_asset_service_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AssetGroupAssetOperation.ProtoReflect.Descriptor instead.
func (*AssetGroupAssetOperation) Descriptor() ([]byte, []int) {
	return file_services_asset_group_asset_service_proto_rawDescGZIP(), []int{1}
}

func (x *AssetGroupAssetOperation) GetUpdateMask() *fieldmaskpb.FieldMask {
	if x != nil {
		return x.UpdateMask
	}
	return nil
}

func (m *AssetGroupAssetOperation) GetOperation() isAssetGroupAssetOperation_Operation {
	if m != nil {
		return m.Operation
	}
	return nil
}

func (x *AssetGroupAssetOperation) GetCreate() *resources.AssetGroupAsset {
	if x, ok := x.GetOperation().(*AssetGroupAssetOperation_Create); ok {
		return x.Create
	}
	return nil
}

func (x *AssetGroupAssetOperation) GetUpdate() *resources.AssetGroupAsset {
	if x, ok := x.GetOperation().(*AssetGroupAssetOperation_Update); ok {
		return x.Update
	}
	return nil
}

func (x *AssetGroupAssetOperation) GetRemove() string {
	if x, ok := x.GetOperation().(*AssetGroupAssetOperation_Remove); ok {
		return x.Remove
	}
	return ""
}

type isAssetGroupAssetOperation_Operation interface {
	isAssetGroupAssetOperation_Operation()
}

type AssetGroupAssetOperation_Create struct {
	// Create operation: No resource name is expected for the new asset group
	// asset.
	Create *resources.AssetGroupAsset `protobuf:"bytes,1,opt,name=create,proto3,oneof"`
}

type AssetGroupAssetOperation_Update struct {
	// Update operation: The asset group asset is expected to have a valid
	// resource name.
	Update *resources.AssetGroupAsset `protobuf:"bytes,2,opt,name=update,proto3,oneof"`
}

type AssetGroupAssetOperation_Remove struct {
	// Remove operation: A resource name for the removed asset group asset is
	// expected, in this format:
	// `customers/{customer_id}/assetGroupAssets/{asset_group_id}~{asset_id}~{field_type}`
	Remove string `protobuf:"bytes,3,opt,name=remove,proto3,oneof"`
}

func (*AssetGroupAssetOperation_Create) isAssetGroupAssetOperation_Operation() {}

func (*AssetGroupAssetOperation_Update) isAssetGroupAssetOperation_Operation() {}

func (*AssetGroupAssetOperation_Remove) isAssetGroupAssetOperation_Operation() {}

// Response message for an asset group asset mutate.
type MutateAssetGroupAssetsResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// All results for the mutate.
	Results []*MutateAssetGroupAssetResult `protobuf:"bytes,1,rep,name=results,proto3" json:"results,omitempty"`
	// Errors that pertain to operation failures in the partial failure mode.
	// Returned only when partial_failure = true and all errors occur inside the
	// operations. If any errors occur outside the operations (for example, auth
	// errors), we return an RPC level error.
	PartialFailureError *status.Status `protobuf:"bytes,2,opt,name=partial_failure_error,json=partialFailureError,proto3" json:"partial_failure_error,omitempty"`
}

func (x *MutateAssetGroupAssetsResponse) Reset() {
	*x = MutateAssetGroupAssetsResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_services_asset_group_asset_service_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MutateAssetGroupAssetsResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MutateAssetGroupAssetsResponse) ProtoMessage() {}

func (x *MutateAssetGroupAssetsResponse) ProtoReflect() protoreflect.Message {
	mi := &file_services_asset_group_asset_service_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MutateAssetGroupAssetsResponse.ProtoReflect.Descriptor instead.
func (*MutateAssetGroupAssetsResponse) Descriptor() ([]byte, []int) {
	return file_services_asset_group_asset_service_proto_rawDescGZIP(), []int{2}
}

func (x *MutateAssetGroupAssetsResponse) GetResults() []*MutateAssetGroupAssetResult {
	if x != nil {
		return x.Results
	}
	return nil
}

func (x *MutateAssetGroupAssetsResponse) GetPartialFailureError() *status.Status {
	if x != nil {
		return x.PartialFailureError
	}
	return nil
}

// The result for the asset group asset mutate.
type MutateAssetGroupAssetResult struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Returned for successful operations.
	ResourceName string `protobuf:"bytes,1,opt,name=resource_name,json=resourceName,proto3" json:"resource_name,omitempty"`
}

func (x *MutateAssetGroupAssetResult) Reset() {
	*x = MutateAssetGroupAssetResult{}
	if protoimpl.UnsafeEnabled {
		mi := &file_services_asset_group_asset_service_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MutateAssetGroupAssetResult) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MutateAssetGroupAssetResult) ProtoMessage() {}

func (x *MutateAssetGroupAssetResult) ProtoReflect() protoreflect.Message {
	mi := &file_services_asset_group_asset_service_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MutateAssetGroupAssetResult.ProtoReflect.Descriptor instead.
func (*MutateAssetGroupAssetResult) Descriptor() ([]byte, []int) {
	return file_services_asset_group_asset_service_proto_rawDescGZIP(), []int{3}
}

func (x *MutateAssetGroupAssetResult) GetResourceName() string {
	if x != nil {
		return x.ResourceName
	}
	return ""
}

var File_services_asset_group_asset_service_proto protoreflect.FileDescriptor

var file_services_asset_group_asset_service_proto_rawDesc = []byte{
	0x0a, 0x41, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x64, 0x73, 0x2f, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x61, 0x64, 0x73, 0x2f, 0x76, 0x31, 0x36, 0x2f, 0x73, 0x65, 0x72, 0x76, 0x69,
	0x63, 0x65, 0x73, 0x2f, 0x61, 0x73, 0x73, 0x65, 0x74, 0x5f, 0x67, 0x72, 0x6f, 0x75, 0x70, 0x5f,
	0x61, 0x73, 0x73, 0x65, 0x74, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x12, 0x21, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x61, 0x64, 0x73, 0x2e,
	0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61, 0x64, 0x73, 0x2e, 0x76, 0x31, 0x36, 0x2e, 0x73, 0x65,
	0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x1a, 0x3a, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61,
	0x64, 0x73, 0x2f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61, 0x64, 0x73, 0x2f, 0x76, 0x31, 0x36,
	0x2f, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x73, 0x2f, 0x61, 0x73, 0x73, 0x65, 0x74,
	0x5f, 0x67, 0x72, 0x6f, 0x75, 0x70, 0x5f, 0x61, 0x73, 0x73, 0x65, 0x74, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x1a, 0x1c, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x61,
	0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x1a, 0x17, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x63, 0x6c, 0x69,
	0x65, 0x6e, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x66, 0x69, 0x65, 0x6c, 0x64, 0x5f, 0x62, 0x65, 0x68, 0x61,
	0x76, 0x69, 0x6f, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x19, 0x67, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x20, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x66, 0x69, 0x65, 0x6c, 0x64, 0x5f, 0x6d, 0x61, 0x73,
	0x6b, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x17, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f,
	0x72, 0x70, 0x63, 0x2f, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x22, 0xf5, 0x01, 0x0a, 0x1d, 0x4d, 0x75, 0x74, 0x61, 0x74, 0x65, 0x41, 0x73, 0x73, 0x65, 0x74,
	0x47, 0x72, 0x6f, 0x75, 0x70, 0x41, 0x73, 0x73, 0x65, 0x74, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x12, 0x24, 0x0a, 0x0b, 0x63, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x65, 0x72, 0x5f, 0x69,
	0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x42, 0x03, 0xe0, 0x41, 0x02, 0x52, 0x0a, 0x63, 0x75,
	0x73, 0x74, 0x6f, 0x6d, 0x65, 0x72, 0x49, 0x64, 0x12, 0x60, 0x0a, 0x0a, 0x6f, 0x70, 0x65, 0x72,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x3b, 0x2e, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x61, 0x64, 0x73, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x61, 0x64, 0x73, 0x2e, 0x76, 0x31, 0x36, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73,
	0x2e, 0x41, 0x73, 0x73, 0x65, 0x74, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x41, 0x73, 0x73, 0x65, 0x74,
	0x4f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x42, 0x03, 0xe0, 0x41, 0x02, 0x52, 0x0a,
	0x6f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x12, 0x27, 0x0a, 0x0f, 0x70, 0x61,
	0x72, 0x74, 0x69, 0x61, 0x6c, 0x5f, 0x66, 0x61, 0x69, 0x6c, 0x75, 0x72, 0x65, 0x18, 0x03, 0x20,
	0x01, 0x28, 0x08, 0x52, 0x0e, 0x70, 0x61, 0x72, 0x74, 0x69, 0x61, 0x6c, 0x46, 0x61, 0x69, 0x6c,
	0x75, 0x72, 0x65, 0x12, 0x23, 0x0a, 0x0d, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x65, 0x5f,
	0x6f, 0x6e, 0x6c, 0x79, 0x18, 0x04, 0x20, 0x01, 0x28, 0x08, 0x52, 0x0c, 0x76, 0x61, 0x6c, 0x69,
	0x64, 0x61, 0x74, 0x65, 0x4f, 0x6e, 0x6c, 0x79, 0x22, 0xcb, 0x02, 0x0a, 0x18, 0x41, 0x73, 0x73,
	0x65, 0x74, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x41, 0x73, 0x73, 0x65, 0x74, 0x4f, 0x70, 0x65, 0x72,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x3b, 0x0a, 0x0b, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x5f,
	0x6d, 0x61, 0x73, 0x6b, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x46, 0x69, 0x65,
	0x6c, 0x64, 0x4d, 0x61, 0x73, 0x6b, 0x52, 0x0a, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x4d, 0x61,
	0x73, 0x6b, 0x12, 0x4d, 0x0a, 0x06, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x33, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x61, 0x64, 0x73, 0x2e,
	0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61, 0x64, 0x73, 0x2e, 0x76, 0x31, 0x36, 0x2e, 0x72, 0x65,
	0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x73, 0x2e, 0x41, 0x73, 0x73, 0x65, 0x74, 0x47, 0x72, 0x6f,
	0x75, 0x70, 0x41, 0x73, 0x73, 0x65, 0x74, 0x48, 0x00, 0x52, 0x06, 0x63, 0x72, 0x65, 0x61, 0x74,
	0x65, 0x12, 0x4d, 0x0a, 0x06, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x33, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x61, 0x64, 0x73, 0x2e, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61, 0x64, 0x73, 0x2e, 0x76, 0x31, 0x36, 0x2e, 0x72, 0x65, 0x73,
	0x6f, 0x75, 0x72, 0x63, 0x65, 0x73, 0x2e, 0x41, 0x73, 0x73, 0x65, 0x74, 0x47, 0x72, 0x6f, 0x75,
	0x70, 0x41, 0x73, 0x73, 0x65, 0x74, 0x48, 0x00, 0x52, 0x06, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65,
	0x12, 0x47, 0x0a, 0x06, 0x72, 0x65, 0x6d, 0x6f, 0x76, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09,
	0x42, 0x2d, 0xfa, 0x41, 0x2a, 0x0a, 0x28, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61, 0x64, 0x73,
	0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61, 0x70, 0x69, 0x73, 0x2e, 0x63, 0x6f, 0x6d, 0x2f,
	0x41, 0x73, 0x73, 0x65, 0x74, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x41, 0x73, 0x73, 0x65, 0x74, 0x48,
	0x00, 0x52, 0x06, 0x72, 0x65, 0x6d, 0x6f, 0x76, 0x65, 0x42, 0x0b, 0x0a, 0x09, 0x6f, 0x70, 0x65,
	0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x22, 0xc2, 0x01, 0x0a, 0x1e, 0x4d, 0x75, 0x74, 0x61, 0x74,
	0x65, 0x41, 0x73, 0x73, 0x65, 0x74, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x41, 0x73, 0x73, 0x65, 0x74,
	0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x58, 0x0a, 0x07, 0x72, 0x65, 0x73,
	0x75, 0x6c, 0x74, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x3e, 0x2e, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x2e, 0x61, 0x64, 0x73, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61, 0x64,
	0x73, 0x2e, 0x76, 0x31, 0x36, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x4d,
	0x75, 0x74, 0x61, 0x74, 0x65, 0x41, 0x73, 0x73, 0x65, 0x74, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x41,
	0x73, 0x73, 0x65, 0x74, 0x52, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x52, 0x07, 0x72, 0x65, 0x73, 0x75,
	0x6c, 0x74, 0x73, 0x12, 0x46, 0x0a, 0x15, 0x70, 0x61, 0x72, 0x74, 0x69, 0x61, 0x6c, 0x5f, 0x66,
	0x61, 0x69, 0x6c, 0x75, 0x72, 0x65, 0x5f, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x12, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x72, 0x70, 0x63, 0x2e,
	0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x52, 0x13, 0x70, 0x61, 0x72, 0x74, 0x69, 0x61, 0x6c, 0x46,
	0x61, 0x69, 0x6c, 0x75, 0x72, 0x65, 0x45, 0x72, 0x72, 0x6f, 0x72, 0x22, 0x71, 0x0a, 0x1b, 0x4d,
	0x75, 0x74, 0x61, 0x74, 0x65, 0x41, 0x73, 0x73, 0x65, 0x74, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x41,
	0x73, 0x73, 0x65, 0x74, 0x52, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x12, 0x52, 0x0a, 0x0d, 0x72, 0x65,
	0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x42, 0x2d, 0xfa, 0x41, 0x2a, 0x0a, 0x28, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61, 0x64,
	0x73, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61, 0x70, 0x69, 0x73, 0x2e, 0x63, 0x6f, 0x6d,
	0x2f, 0x41, 0x73, 0x73, 0x65, 0x74, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x41, 0x73, 0x73, 0x65, 0x74,
	0x52, 0x0c, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x4e, 0x61, 0x6d, 0x65, 0x32, 0xdb,
	0x02, 0x0a, 0x16, 0x41, 0x73, 0x73, 0x65, 0x74, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x41, 0x73, 0x73,
	0x65, 0x74, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0xf9, 0x01, 0x0a, 0x16, 0x4d, 0x75,
	0x74, 0x61, 0x74, 0x65, 0x41, 0x73, 0x73, 0x65, 0x74, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x41, 0x73,
	0x73, 0x65, 0x74, 0x73, 0x12, 0x40, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x61, 0x64,
	0x73, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61, 0x64, 0x73, 0x2e, 0x76, 0x31, 0x36, 0x2e,
	0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x4d, 0x75, 0x74, 0x61, 0x74, 0x65, 0x41,
	0x73, 0x73, 0x65, 0x74, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x41, 0x73, 0x73, 0x65, 0x74, 0x73, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x41, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e,
	0x61, 0x64, 0x73, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61, 0x64, 0x73, 0x2e, 0x76, 0x31,
	0x36, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x4d, 0x75, 0x74, 0x61, 0x74,
	0x65, 0x41, 0x73, 0x73, 0x65, 0x74, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x41, 0x73, 0x73, 0x65, 0x74,
	0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x5a, 0x82, 0xd3, 0xe4, 0x93, 0x02,
	0x3b, 0x22, 0x36, 0x2f, 0x76, 0x31, 0x36, 0x2f, 0x63, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x65, 0x72,
	0x73, 0x2f, 0x7b, 0x63, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x3d, 0x2a,
	0x7d, 0x2f, 0x61, 0x73, 0x73, 0x65, 0x74, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x41, 0x73, 0x73, 0x65,
	0x74, 0x73, 0x3a, 0x6d, 0x75, 0x74, 0x61, 0x74, 0x65, 0x3a, 0x01, 0x2a, 0xda, 0x41, 0x16, 0x63,
	0x75, 0x73, 0x74, 0x6f, 0x6d, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x2c, 0x6f, 0x70, 0x65, 0x72, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x73, 0x1a, 0x45, 0xca, 0x41, 0x18, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x61, 0x64, 0x73, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61, 0x70, 0x69, 0x73, 0x2e, 0x63,
	0x6f, 0x6d, 0xd2, 0x41, 0x27, 0x68, 0x74, 0x74, 0x70, 0x73, 0x3a, 0x2f, 0x2f, 0x77, 0x77, 0x77,
	0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61, 0x70, 0x69, 0x73, 0x2e, 0x63, 0x6f, 0x6d, 0x2f,
	0x61, 0x75, 0x74, 0x68, 0x2f, 0x61, 0x64, 0x77, 0x6f, 0x72, 0x64, 0x73, 0x42, 0x87, 0x02, 0x0a,
	0x25, 0x63, 0x6f, 0x6d, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x61, 0x64, 0x73, 0x2e,
	0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61, 0x64, 0x73, 0x2e, 0x76, 0x31, 0x36, 0x2e, 0x73, 0x65,
	0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x42, 0x1b, 0x41, 0x73, 0x73, 0x65, 0x74, 0x47, 0x72, 0x6f,
	0x75, 0x70, 0x41, 0x73, 0x73, 0x65, 0x74, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x50, 0x72,
	0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a, 0x49, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x67, 0x6f,
	0x6c, 0x61, 0x6e, 0x67, 0x2e, 0x6f, 0x72, 0x67, 0x2f, 0x67, 0x65, 0x6e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x2f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61, 0x70, 0x69, 0x73, 0x2f, 0x61, 0x64, 0x73,
	0x2f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61, 0x64, 0x73, 0x2f, 0x76, 0x31, 0x36, 0x2f, 0x73,
	0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x3b, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73,
	0xa2, 0x02, 0x03, 0x47, 0x41, 0x41, 0xaa, 0x02, 0x21, 0x47, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e,
	0x41, 0x64, 0x73, 0x2e, 0x47, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x41, 0x64, 0x73, 0x2e, 0x56, 0x31,
	0x36, 0x2e, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0xca, 0x02, 0x21, 0x47, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x5c, 0x41, 0x64, 0x73, 0x5c, 0x47, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x41, 0x64,
	0x73, 0x5c, 0x56, 0x31, 0x36, 0x5c, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0xea, 0x02,
	0x25, 0x47, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x3a, 0x3a, 0x41, 0x64, 0x73, 0x3a, 0x3a, 0x47, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x41, 0x64, 0x73, 0x3a, 0x3a, 0x56, 0x31, 0x36, 0x3a, 0x3a, 0x53, 0x65,
	0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_services_asset_group_asset_service_proto_rawDescOnce sync.Once
	file_services_asset_group_asset_service_proto_rawDescData = file_services_asset_group_asset_service_proto_rawDesc
)

func file_services_asset_group_asset_service_proto_rawDescGZIP() []byte {
	file_services_asset_group_asset_service_proto_rawDescOnce.Do(func() {
		file_services_asset_group_asset_service_proto_rawDescData = protoimpl.X.CompressGZIP(file_services_asset_group_asset_service_proto_rawDescData)
	})
	return file_services_asset_group_asset_service_proto_rawDescData
}

var file_services_asset_group_asset_service_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_services_asset_group_asset_service_proto_goTypes = []interface{}{
	(*MutateAssetGroupAssetsRequest)(nil),  // 0: google.ads.googleads.v16.services.MutateAssetGroupAssetsRequest
	(*AssetGroupAssetOperation)(nil),       // 1: google.ads.googleads.v16.services.AssetGroupAssetOperation
	(*MutateAssetGroupAssetsResponse)(nil), // 2: google.ads.googleads.v16.services.MutateAssetGroupAssetsResponse
	(*MutateAssetGroupAssetResult)(nil),    // 3: google.ads.googleads.v16.services.MutateAssetGroupAssetResult
	(*fieldmaskpb.FieldMask)(nil),          // 4: google.protobuf.FieldMask
	(*resources.AssetGroupAsset)(nil),      // 5: google.ads.googleads.v16.resources.AssetGroupAsset
	(*status.Status)(nil),                  // 6: google.rpc.Status
}
var file_services_asset_group_asset_service_proto_depIdxs = []int32{
	1, // 0: google.ads.googleads.v16.services.MutateAssetGroupAssetsRequest.operations:type_name -> google.ads.googleads.v16.services.AssetGroupAssetOperation
	4, // 1: google.ads.googleads.v16.services.AssetGroupAssetOperation.update_mask:type_name -> google.protobuf.FieldMask
	5, // 2: google.ads.googleads.v16.services.AssetGroupAssetOperation.create:type_name -> google.ads.googleads.v16.resources.AssetGroupAsset
	5, // 3: google.ads.googleads.v16.services.AssetGroupAssetOperation.update:type_name -> google.ads.googleads.v16.resources.AssetGroupAsset
	3, // 4: google.ads.googleads.v16.services.MutateAssetGroupAssetsResponse.results:type_name -> google.ads.googleads.v16.services.MutateAssetGroupAssetResult
	6, // 5: google.ads.googleads.v16.services.MutateAssetGroupAssetsResponse.partial_failure_error:type_name -> google.rpc.Status
	0, // 6: google.ads.googleads.v16.services.AssetGroupAssetService.MutateAssetGroupAssets:input_type -> google.ads.googleads.v16.services.MutateAssetGroupAssetsRequest
	2, // 7: google.ads.googleads.v16.services.AssetGroupAssetService.MutateAssetGroupAssets:output_type -> google.ads.googleads.v16.services.MutateAssetGroupAssetsResponse
	7, // [7:8] is the sub-list for method output_type
	6, // [6:7] is the sub-list for method input_type
	6, // [6:6] is the sub-list for extension type_name
	6, // [6:6] is the sub-list for extension extendee
	0, // [0:6] is the sub-list for field type_name
}

func init() { file_services_asset_group_asset_service_proto_init() }
func file_services_asset_group_asset_service_proto_init() {
	if File_services_asset_group_asset_service_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_services_asset_group_asset_service_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*MutateAssetGroupAssetsRequest); i {
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
		file_services_asset_group_asset_service_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AssetGroupAssetOperation); i {
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
		file_services_asset_group_asset_service_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*MutateAssetGroupAssetsResponse); i {
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
		file_services_asset_group_asset_service_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*MutateAssetGroupAssetResult); i {
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
	file_services_asset_group_asset_service_proto_msgTypes[1].OneofWrappers = []interface{}{
		(*AssetGroupAssetOperation_Create)(nil),
		(*AssetGroupAssetOperation_Update)(nil),
		(*AssetGroupAssetOperation_Remove)(nil),
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_services_asset_group_asset_service_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_services_asset_group_asset_service_proto_goTypes,
		DependencyIndexes: file_services_asset_group_asset_service_proto_depIdxs,
		MessageInfos:      file_services_asset_group_asset_service_proto_msgTypes,
	}.Build()
	File_services_asset_group_asset_service_proto = out.File
	file_services_asset_group_asset_service_proto_rawDesc = nil
	file_services_asset_group_asset_service_proto_goTypes = nil
	file_services_asset_group_asset_service_proto_depIdxs = nil
}
