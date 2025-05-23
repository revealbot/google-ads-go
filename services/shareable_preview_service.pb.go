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
// source: google/ads/googleads/v19/services/shareable_preview_service.proto

package services

import (
	_ "google.golang.org/genproto/googleapis/api/annotations"
	status "google.golang.org/genproto/googleapis/rpc/status"
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

// Request message for
// [ShareablePreviewService.GenerateShareablePreviews][google.ads.googleads.v19.services.ShareablePreviewService.GenerateShareablePreviews].
type GenerateShareablePreviewsRequest struct {
	state protoimpl.MessageState `protogen:"open.v1"`
	// Required. The customer creating the shareable previews request.
	CustomerId string `protobuf:"bytes,1,opt,name=customer_id,json=customerId,proto3" json:"customer_id,omitempty"`
	// Required. The list of shareable previews to generate.
	ShareablePreviews []*ShareablePreview `protobuf:"bytes,2,rep,name=shareable_previews,json=shareablePreviews,proto3" json:"shareable_previews,omitempty"`
	unknownFields     protoimpl.UnknownFields
	sizeCache         protoimpl.SizeCache
}

func (x *GenerateShareablePreviewsRequest) Reset() {
	*x = GenerateShareablePreviewsRequest{}
	mi := &file_services_shareable_preview_service_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *GenerateShareablePreviewsRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GenerateShareablePreviewsRequest) ProtoMessage() {}

func (x *GenerateShareablePreviewsRequest) ProtoReflect() protoreflect.Message {
	mi := &file_services_shareable_preview_service_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GenerateShareablePreviewsRequest.ProtoReflect.Descriptor instead.
func (*GenerateShareablePreviewsRequest) Descriptor() ([]byte, []int) {
	return file_services_shareable_preview_service_proto_rawDescGZIP(), []int{0}
}

func (x *GenerateShareablePreviewsRequest) GetCustomerId() string {
	if x != nil {
		return x.CustomerId
	}
	return ""
}

func (x *GenerateShareablePreviewsRequest) GetShareablePreviews() []*ShareablePreview {
	if x != nil {
		return x.ShareablePreviews
	}
	return nil
}

// A shareable preview with its identifier.
type ShareablePreview struct {
	state protoimpl.MessageState `protogen:"open.v1"`
	// Required. Asset group of the shareable preview.
	AssetGroupIdentifier *AssetGroupIdentifier `protobuf:"bytes,1,opt,name=asset_group_identifier,json=assetGroupIdentifier,proto3" json:"asset_group_identifier,omitempty"`
	unknownFields        protoimpl.UnknownFields
	sizeCache            protoimpl.SizeCache
}

func (x *ShareablePreview) Reset() {
	*x = ShareablePreview{}
	mi := &file_services_shareable_preview_service_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *ShareablePreview) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ShareablePreview) ProtoMessage() {}

func (x *ShareablePreview) ProtoReflect() protoreflect.Message {
	mi := &file_services_shareable_preview_service_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ShareablePreview.ProtoReflect.Descriptor instead.
func (*ShareablePreview) Descriptor() ([]byte, []int) {
	return file_services_shareable_preview_service_proto_rawDescGZIP(), []int{1}
}

func (x *ShareablePreview) GetAssetGroupIdentifier() *AssetGroupIdentifier {
	if x != nil {
		return x.AssetGroupIdentifier
	}
	return nil
}

// Asset group of the shareable preview.
type AssetGroupIdentifier struct {
	state protoimpl.MessageState `protogen:"open.v1"`
	// Required. The asset group identifier.
	AssetGroupId  int64 `protobuf:"varint,1,opt,name=asset_group_id,json=assetGroupId,proto3" json:"asset_group_id,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *AssetGroupIdentifier) Reset() {
	*x = AssetGroupIdentifier{}
	mi := &file_services_shareable_preview_service_proto_msgTypes[2]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *AssetGroupIdentifier) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AssetGroupIdentifier) ProtoMessage() {}

func (x *AssetGroupIdentifier) ProtoReflect() protoreflect.Message {
	mi := &file_services_shareable_preview_service_proto_msgTypes[2]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AssetGroupIdentifier.ProtoReflect.Descriptor instead.
func (*AssetGroupIdentifier) Descriptor() ([]byte, []int) {
	return file_services_shareable_preview_service_proto_rawDescGZIP(), []int{2}
}

func (x *AssetGroupIdentifier) GetAssetGroupId() int64 {
	if x != nil {
		return x.AssetGroupId
	}
	return 0
}

// Response message for
// [ShareablePreviewService.GenerateShareablePreviews][google.ads.googleads.v19.services.ShareablePreviewService.GenerateShareablePreviews].
type GenerateShareablePreviewsResponse struct {
	state protoimpl.MessageState `protogen:"open.v1"`
	// List of generate shareable preview results.
	Responses     []*ShareablePreviewOrError `protobuf:"bytes,1,rep,name=responses,proto3" json:"responses,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *GenerateShareablePreviewsResponse) Reset() {
	*x = GenerateShareablePreviewsResponse{}
	mi := &file_services_shareable_preview_service_proto_msgTypes[3]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *GenerateShareablePreviewsResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GenerateShareablePreviewsResponse) ProtoMessage() {}

func (x *GenerateShareablePreviewsResponse) ProtoReflect() protoreflect.Message {
	mi := &file_services_shareable_preview_service_proto_msgTypes[3]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GenerateShareablePreviewsResponse.ProtoReflect.Descriptor instead.
func (*GenerateShareablePreviewsResponse) Descriptor() ([]byte, []int) {
	return file_services_shareable_preview_service_proto_rawDescGZIP(), []int{3}
}

func (x *GenerateShareablePreviewsResponse) GetResponses() []*ShareablePreviewOrError {
	if x != nil {
		return x.Responses
	}
	return nil
}

// Result of the generate shareable preview.
type ShareablePreviewOrError struct {
	state protoimpl.MessageState `protogen:"open.v1"`
	// The asset group of the shareable preview.
	AssetGroupIdentifier *AssetGroupIdentifier `protobuf:"bytes,3,opt,name=asset_group_identifier,json=assetGroupIdentifier,proto3" json:"asset_group_identifier,omitempty"`
	// The shareable preview result or error.
	//
	// Types that are valid to be assigned to GenerateShareablePreviewResponse:
	//
	//	*ShareablePreviewOrError_ShareablePreviewResult
	//	*ShareablePreviewOrError_PartialFailureError
	GenerateShareablePreviewResponse isShareablePreviewOrError_GenerateShareablePreviewResponse `protobuf_oneof:"generate_shareable_preview_response"`
	unknownFields                    protoimpl.UnknownFields
	sizeCache                        protoimpl.SizeCache
}

func (x *ShareablePreviewOrError) Reset() {
	*x = ShareablePreviewOrError{}
	mi := &file_services_shareable_preview_service_proto_msgTypes[4]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *ShareablePreviewOrError) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ShareablePreviewOrError) ProtoMessage() {}

func (x *ShareablePreviewOrError) ProtoReflect() protoreflect.Message {
	mi := &file_services_shareable_preview_service_proto_msgTypes[4]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ShareablePreviewOrError.ProtoReflect.Descriptor instead.
func (*ShareablePreviewOrError) Descriptor() ([]byte, []int) {
	return file_services_shareable_preview_service_proto_rawDescGZIP(), []int{4}
}

func (x *ShareablePreviewOrError) GetAssetGroupIdentifier() *AssetGroupIdentifier {
	if x != nil {
		return x.AssetGroupIdentifier
	}
	return nil
}

func (x *ShareablePreviewOrError) GetGenerateShareablePreviewResponse() isShareablePreviewOrError_GenerateShareablePreviewResponse {
	if x != nil {
		return x.GenerateShareablePreviewResponse
	}
	return nil
}

func (x *ShareablePreviewOrError) GetShareablePreviewResult() *ShareablePreviewResult {
	if x != nil {
		if x, ok := x.GenerateShareablePreviewResponse.(*ShareablePreviewOrError_ShareablePreviewResult); ok {
			return x.ShareablePreviewResult
		}
	}
	return nil
}

func (x *ShareablePreviewOrError) GetPartialFailureError() *status.Status {
	if x != nil {
		if x, ok := x.GenerateShareablePreviewResponse.(*ShareablePreviewOrError_PartialFailureError); ok {
			return x.PartialFailureError
		}
	}
	return nil
}

type isShareablePreviewOrError_GenerateShareablePreviewResponse interface {
	isShareablePreviewOrError_GenerateShareablePreviewResponse()
}

type ShareablePreviewOrError_ShareablePreviewResult struct {
	// The shareable preview result.
	ShareablePreviewResult *ShareablePreviewResult `protobuf:"bytes,1,opt,name=shareable_preview_result,json=shareablePreviewResult,proto3,oneof"`
}

type ShareablePreviewOrError_PartialFailureError struct {
	// The shareable preview partial failure error.
	PartialFailureError *status.Status `protobuf:"bytes,2,opt,name=partial_failure_error,json=partialFailureError,proto3,oneof"`
}

func (*ShareablePreviewOrError_ShareablePreviewResult) isShareablePreviewOrError_GenerateShareablePreviewResponse() {
}

func (*ShareablePreviewOrError_PartialFailureError) isShareablePreviewOrError_GenerateShareablePreviewResponse() {
}

// Message to hold a shareable preview result.
type ShareablePreviewResult struct {
	state protoimpl.MessageState `protogen:"open.v1"`
	// The shareable preview URL.
	ShareablePreviewUrl string `protobuf:"bytes,1,opt,name=shareable_preview_url,json=shareablePreviewUrl,proto3" json:"shareable_preview_url,omitempty"`
	// Expiration date time using the ISO-8601 format.
	ExpirationDateTime string `protobuf:"bytes,2,opt,name=expiration_date_time,json=expirationDateTime,proto3" json:"expiration_date_time,omitempty"`
	unknownFields      protoimpl.UnknownFields
	sizeCache          protoimpl.SizeCache
}

func (x *ShareablePreviewResult) Reset() {
	*x = ShareablePreviewResult{}
	mi := &file_services_shareable_preview_service_proto_msgTypes[5]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *ShareablePreviewResult) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ShareablePreviewResult) ProtoMessage() {}

func (x *ShareablePreviewResult) ProtoReflect() protoreflect.Message {
	mi := &file_services_shareable_preview_service_proto_msgTypes[5]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ShareablePreviewResult.ProtoReflect.Descriptor instead.
func (*ShareablePreviewResult) Descriptor() ([]byte, []int) {
	return file_services_shareable_preview_service_proto_rawDescGZIP(), []int{5}
}

func (x *ShareablePreviewResult) GetShareablePreviewUrl() string {
	if x != nil {
		return x.ShareablePreviewUrl
	}
	return ""
}

func (x *ShareablePreviewResult) GetExpirationDateTime() string {
	if x != nil {
		return x.ExpirationDateTime
	}
	return ""
}

var File_services_shareable_preview_service_proto protoreflect.FileDescriptor

var file_services_shareable_preview_service_proto_rawDesc = string([]byte{
	0x0a, 0x41, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x64, 0x73, 0x2f, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x61, 0x64, 0x73, 0x2f, 0x76, 0x31, 0x39, 0x2f, 0x73, 0x65, 0x72, 0x76, 0x69,
	0x63, 0x65, 0x73, 0x2f, 0x73, 0x68, 0x61, 0x72, 0x65, 0x61, 0x62, 0x6c, 0x65, 0x5f, 0x70, 0x72,
	0x65, 0x76, 0x69, 0x65, 0x77, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x12, 0x21, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x61, 0x64, 0x73, 0x2e,
	0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61, 0x64, 0x73, 0x2e, 0x76, 0x31, 0x39, 0x2e, 0x73, 0x65,
	0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x1a, 0x1c, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61,
	0x70, 0x69, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x17, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69,
	0x2f, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1f, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x66, 0x69, 0x65, 0x6c, 0x64, 0x5f,
	0x62, 0x65, 0x68, 0x61, 0x76, 0x69, 0x6f, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x17,
	0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x72, 0x70, 0x63, 0x2f, 0x73, 0x74, 0x61, 0x74, 0x75,
	0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xb1, 0x01, 0x0a, 0x20, 0x47, 0x65, 0x6e, 0x65,
	0x72, 0x61, 0x74, 0x65, 0x53, 0x68, 0x61, 0x72, 0x65, 0x61, 0x62, 0x6c, 0x65, 0x50, 0x72, 0x65,
	0x76, 0x69, 0x65, 0x77, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x24, 0x0a, 0x0b,
	0x63, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x42, 0x03, 0xe0, 0x41, 0x02, 0x52, 0x0a, 0x63, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x65, 0x72,
	0x49, 0x64, 0x12, 0x67, 0x0a, 0x12, 0x73, 0x68, 0x61, 0x72, 0x65, 0x61, 0x62, 0x6c, 0x65, 0x5f,
	0x70, 0x72, 0x65, 0x76, 0x69, 0x65, 0x77, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x33,
	0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x61, 0x64, 0x73, 0x2e, 0x67, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x61, 0x64, 0x73, 0x2e, 0x76, 0x31, 0x39, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63,
	0x65, 0x73, 0x2e, 0x53, 0x68, 0x61, 0x72, 0x65, 0x61, 0x62, 0x6c, 0x65, 0x50, 0x72, 0x65, 0x76,
	0x69, 0x65, 0x77, 0x42, 0x03, 0xe0, 0x41, 0x02, 0x52, 0x11, 0x73, 0x68, 0x61, 0x72, 0x65, 0x61,
	0x62, 0x6c, 0x65, 0x50, 0x72, 0x65, 0x76, 0x69, 0x65, 0x77, 0x73, 0x22, 0x86, 0x01, 0x0a, 0x10,
	0x53, 0x68, 0x61, 0x72, 0x65, 0x61, 0x62, 0x6c, 0x65, 0x50, 0x72, 0x65, 0x76, 0x69, 0x65, 0x77,
	0x12, 0x72, 0x0a, 0x16, 0x61, 0x73, 0x73, 0x65, 0x74, 0x5f, 0x67, 0x72, 0x6f, 0x75, 0x70, 0x5f,
	0x69, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x66, 0x69, 0x65, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x37, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x61, 0x64, 0x73, 0x2e, 0x67, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x61, 0x64, 0x73, 0x2e, 0x76, 0x31, 0x39, 0x2e, 0x73, 0x65, 0x72, 0x76,
	0x69, 0x63, 0x65, 0x73, 0x2e, 0x41, 0x73, 0x73, 0x65, 0x74, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x49,
	0x64, 0x65, 0x6e, 0x74, 0x69, 0x66, 0x69, 0x65, 0x72, 0x42, 0x03, 0xe0, 0x41, 0x02, 0x52, 0x14,
	0x61, 0x73, 0x73, 0x65, 0x74, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x49, 0x64, 0x65, 0x6e, 0x74, 0x69,
	0x66, 0x69, 0x65, 0x72, 0x22, 0x41, 0x0a, 0x14, 0x41, 0x73, 0x73, 0x65, 0x74, 0x47, 0x72, 0x6f,
	0x75, 0x70, 0x49, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x66, 0x69, 0x65, 0x72, 0x12, 0x29, 0x0a, 0x0e,
	0x61, 0x73, 0x73, 0x65, 0x74, 0x5f, 0x67, 0x72, 0x6f, 0x75, 0x70, 0x5f, 0x69, 0x64, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x03, 0x42, 0x03, 0xe0, 0x41, 0x02, 0x52, 0x0c, 0x61, 0x73, 0x73, 0x65, 0x74,
	0x47, 0x72, 0x6f, 0x75, 0x70, 0x49, 0x64, 0x22, 0x7d, 0x0a, 0x21, 0x47, 0x65, 0x6e, 0x65, 0x72,
	0x61, 0x74, 0x65, 0x53, 0x68, 0x61, 0x72, 0x65, 0x61, 0x62, 0x6c, 0x65, 0x50, 0x72, 0x65, 0x76,
	0x69, 0x65, 0x77, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x58, 0x0a, 0x09,
	0x72, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32,
	0x3a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x61, 0x64, 0x73, 0x2e, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x61, 0x64, 0x73, 0x2e, 0x76, 0x31, 0x39, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69,
	0x63, 0x65, 0x73, 0x2e, 0x53, 0x68, 0x61, 0x72, 0x65, 0x61, 0x62, 0x6c, 0x65, 0x50, 0x72, 0x65,
	0x76, 0x69, 0x65, 0x77, 0x4f, 0x72, 0x45, 0x72, 0x72, 0x6f, 0x72, 0x52, 0x09, 0x72, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x73, 0x22, 0xf0, 0x02, 0x0a, 0x17, 0x53, 0x68, 0x61, 0x72, 0x65,
	0x61, 0x62, 0x6c, 0x65, 0x50, 0x72, 0x65, 0x76, 0x69, 0x65, 0x77, 0x4f, 0x72, 0x45, 0x72, 0x72,
	0x6f, 0x72, 0x12, 0x6d, 0x0a, 0x16, 0x61, 0x73, 0x73, 0x65, 0x74, 0x5f, 0x67, 0x72, 0x6f, 0x75,
	0x70, 0x5f, 0x69, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x66, 0x69, 0x65, 0x72, 0x18, 0x03, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x37, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x61, 0x64, 0x73, 0x2e,
	0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61, 0x64, 0x73, 0x2e, 0x76, 0x31, 0x39, 0x2e, 0x73, 0x65,
	0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x41, 0x73, 0x73, 0x65, 0x74, 0x47, 0x72, 0x6f, 0x75,
	0x70, 0x49, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x66, 0x69, 0x65, 0x72, 0x52, 0x14, 0x61, 0x73, 0x73,
	0x65, 0x74, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x49, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x66, 0x69, 0x65,
	0x72, 0x12, 0x75, 0x0a, 0x18, 0x73, 0x68, 0x61, 0x72, 0x65, 0x61, 0x62, 0x6c, 0x65, 0x5f, 0x70,
	0x72, 0x65, 0x76, 0x69, 0x65, 0x77, 0x5f, 0x72, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x39, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x61, 0x64, 0x73,
	0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61, 0x64, 0x73, 0x2e, 0x76, 0x31, 0x39, 0x2e, 0x73,
	0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x53, 0x68, 0x61, 0x72, 0x65, 0x61, 0x62, 0x6c,
	0x65, 0x50, 0x72, 0x65, 0x76, 0x69, 0x65, 0x77, 0x52, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x48, 0x00,
	0x52, 0x16, 0x73, 0x68, 0x61, 0x72, 0x65, 0x61, 0x62, 0x6c, 0x65, 0x50, 0x72, 0x65, 0x76, 0x69,
	0x65, 0x77, 0x52, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x12, 0x48, 0x0a, 0x15, 0x70, 0x61, 0x72, 0x74,
	0x69, 0x61, 0x6c, 0x5f, 0x66, 0x61, 0x69, 0x6c, 0x75, 0x72, 0x65, 0x5f, 0x65, 0x72, 0x72, 0x6f,
	0x72, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x12, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x2e, 0x72, 0x70, 0x63, 0x2e, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x48, 0x00, 0x52, 0x13, 0x70,
	0x61, 0x72, 0x74, 0x69, 0x61, 0x6c, 0x46, 0x61, 0x69, 0x6c, 0x75, 0x72, 0x65, 0x45, 0x72, 0x72,
	0x6f, 0x72, 0x42, 0x25, 0x0a, 0x23, 0x67, 0x65, 0x6e, 0x65, 0x72, 0x61, 0x74, 0x65, 0x5f, 0x73,
	0x68, 0x61, 0x72, 0x65, 0x61, 0x62, 0x6c, 0x65, 0x5f, 0x70, 0x72, 0x65, 0x76, 0x69, 0x65, 0x77,
	0x5f, 0x72, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x7e, 0x0a, 0x16, 0x53, 0x68, 0x61,
	0x72, 0x65, 0x61, 0x62, 0x6c, 0x65, 0x50, 0x72, 0x65, 0x76, 0x69, 0x65, 0x77, 0x52, 0x65, 0x73,
	0x75, 0x6c, 0x74, 0x12, 0x32, 0x0a, 0x15, 0x73, 0x68, 0x61, 0x72, 0x65, 0x61, 0x62, 0x6c, 0x65,
	0x5f, 0x70, 0x72, 0x65, 0x76, 0x69, 0x65, 0x77, 0x5f, 0x75, 0x72, 0x6c, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x13, 0x73, 0x68, 0x61, 0x72, 0x65, 0x61, 0x62, 0x6c, 0x65, 0x50, 0x72, 0x65,
	0x76, 0x69, 0x65, 0x77, 0x55, 0x72, 0x6c, 0x12, 0x30, 0x0a, 0x14, 0x65, 0x78, 0x70, 0x69, 0x72,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x64, 0x61, 0x74, 0x65, 0x5f, 0x74, 0x69, 0x6d, 0x65, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x12, 0x65, 0x78, 0x70, 0x69, 0x72, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x44, 0x61, 0x74, 0x65, 0x54, 0x69, 0x6d, 0x65, 0x32, 0xef, 0x02, 0x0a, 0x17, 0x53, 0x68,
	0x61, 0x72, 0x65, 0x61, 0x62, 0x6c, 0x65, 0x50, 0x72, 0x65, 0x76, 0x69, 0x65, 0x77, 0x53, 0x65,
	0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x8c, 0x02, 0x0a, 0x19, 0x47, 0x65, 0x6e, 0x65, 0x72, 0x61,
	0x74, 0x65, 0x53, 0x68, 0x61, 0x72, 0x65, 0x61, 0x62, 0x6c, 0x65, 0x50, 0x72, 0x65, 0x76, 0x69,
	0x65, 0x77, 0x73, 0x12, 0x43, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x61, 0x64, 0x73,
	0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61, 0x64, 0x73, 0x2e, 0x76, 0x31, 0x39, 0x2e, 0x73,
	0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x47, 0x65, 0x6e, 0x65, 0x72, 0x61, 0x74, 0x65,
	0x53, 0x68, 0x61, 0x72, 0x65, 0x61, 0x62, 0x6c, 0x65, 0x50, 0x72, 0x65, 0x76, 0x69, 0x65, 0x77,
	0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x44, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x2e, 0x61, 0x64, 0x73, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61, 0x64, 0x73, 0x2e,
	0x76, 0x31, 0x39, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x47, 0x65, 0x6e,
	0x65, 0x72, 0x61, 0x74, 0x65, 0x53, 0x68, 0x61, 0x72, 0x65, 0x61, 0x62, 0x6c, 0x65, 0x50, 0x72,
	0x65, 0x76, 0x69, 0x65, 0x77, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x64,
	0xda, 0x41, 0x1e, 0x63, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x2c, 0x73,
	0x68, 0x61, 0x72, 0x65, 0x61, 0x62, 0x6c, 0x65, 0x5f, 0x70, 0x72, 0x65, 0x76, 0x69, 0x65, 0x77,
	0x73, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x3d, 0x3a, 0x01, 0x2a, 0x22, 0x38, 0x2f, 0x76, 0x31, 0x39,
	0x2f, 0x63, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x65, 0x72, 0x73, 0x2f, 0x7b, 0x63, 0x75, 0x73, 0x74,
	0x6f, 0x6d, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x3d, 0x2a, 0x7d, 0x3a, 0x67, 0x65, 0x6e, 0x65, 0x72,
	0x61, 0x74, 0x65, 0x53, 0x68, 0x61, 0x72, 0x65, 0x61, 0x62, 0x6c, 0x65, 0x50, 0x72, 0x65, 0x76,
	0x69, 0x65, 0x77, 0x73, 0x1a, 0x45, 0xca, 0x41, 0x18, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61,
	0x64, 0x73, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61, 0x70, 0x69, 0x73, 0x2e, 0x63, 0x6f,
	0x6d, 0xd2, 0x41, 0x27, 0x68, 0x74, 0x74, 0x70, 0x73, 0x3a, 0x2f, 0x2f, 0x77, 0x77, 0x77, 0x2e,
	0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61, 0x70, 0x69, 0x73, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x61,
	0x75, 0x74, 0x68, 0x2f, 0x61, 0x64, 0x77, 0x6f, 0x72, 0x64, 0x73, 0x42, 0x88, 0x02, 0x0a, 0x25,
	0x63, 0x6f, 0x6d, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x61, 0x64, 0x73, 0x2e, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61, 0x64, 0x73, 0x2e, 0x76, 0x31, 0x39, 0x2e, 0x73, 0x65, 0x72,
	0x76, 0x69, 0x63, 0x65, 0x73, 0x42, 0x1c, 0x53, 0x68, 0x61, 0x72, 0x65, 0x61, 0x62, 0x6c, 0x65,
	0x50, 0x72, 0x65, 0x76, 0x69, 0x65, 0x77, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x50, 0x72,
	0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a, 0x49, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x67, 0x6f,
	0x6c, 0x61, 0x6e, 0x67, 0x2e, 0x6f, 0x72, 0x67, 0x2f, 0x67, 0x65, 0x6e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x2f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61, 0x70, 0x69, 0x73, 0x2f, 0x61, 0x64, 0x73,
	0x2f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61, 0x64, 0x73, 0x2f, 0x76, 0x31, 0x39, 0x2f, 0x73,
	0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x3b, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73,
	0xa2, 0x02, 0x03, 0x47, 0x41, 0x41, 0xaa, 0x02, 0x21, 0x47, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e,
	0x41, 0x64, 0x73, 0x2e, 0x47, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x41, 0x64, 0x73, 0x2e, 0x56, 0x31,
	0x39, 0x2e, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0xca, 0x02, 0x21, 0x47, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x5c, 0x41, 0x64, 0x73, 0x5c, 0x47, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x41, 0x64,
	0x73, 0x5c, 0x56, 0x31, 0x39, 0x5c, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0xea, 0x02,
	0x25, 0x47, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x3a, 0x3a, 0x41, 0x64, 0x73, 0x3a, 0x3a, 0x47, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x41, 0x64, 0x73, 0x3a, 0x3a, 0x56, 0x31, 0x39, 0x3a, 0x3a, 0x53, 0x65,
	0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
})

var (
	file_services_shareable_preview_service_proto_rawDescOnce sync.Once
	file_services_shareable_preview_service_proto_rawDescData []byte
)

func file_services_shareable_preview_service_proto_rawDescGZIP() []byte {
	file_services_shareable_preview_service_proto_rawDescOnce.Do(func() {
		file_services_shareable_preview_service_proto_rawDescData = protoimpl.X.CompressGZIP(unsafe.Slice(unsafe.StringData(file_services_shareable_preview_service_proto_rawDesc), len(file_services_shareable_preview_service_proto_rawDesc)))
	})
	return file_services_shareable_preview_service_proto_rawDescData
}

var file_services_shareable_preview_service_proto_msgTypes = make([]protoimpl.MessageInfo, 6)
var file_services_shareable_preview_service_proto_goTypes = []any{
	(*GenerateShareablePreviewsRequest)(nil),  // 0: google.ads.googleads.v19.services.GenerateShareablePreviewsRequest
	(*ShareablePreview)(nil),                  // 1: google.ads.googleads.v19.services.ShareablePreview
	(*AssetGroupIdentifier)(nil),              // 2: google.ads.googleads.v19.services.AssetGroupIdentifier
	(*GenerateShareablePreviewsResponse)(nil), // 3: google.ads.googleads.v19.services.GenerateShareablePreviewsResponse
	(*ShareablePreviewOrError)(nil),           // 4: google.ads.googleads.v19.services.ShareablePreviewOrError
	(*ShareablePreviewResult)(nil),            // 5: google.ads.googleads.v19.services.ShareablePreviewResult
	(*status.Status)(nil),                     // 6: google.rpc.Status
}
var file_services_shareable_preview_service_proto_depIdxs = []int32{
	1, // 0: google.ads.googleads.v19.services.GenerateShareablePreviewsRequest.shareable_previews:type_name -> google.ads.googleads.v19.services.ShareablePreview
	2, // 1: google.ads.googleads.v19.services.ShareablePreview.asset_group_identifier:type_name -> google.ads.googleads.v19.services.AssetGroupIdentifier
	4, // 2: google.ads.googleads.v19.services.GenerateShareablePreviewsResponse.responses:type_name -> google.ads.googleads.v19.services.ShareablePreviewOrError
	2, // 3: google.ads.googleads.v19.services.ShareablePreviewOrError.asset_group_identifier:type_name -> google.ads.googleads.v19.services.AssetGroupIdentifier
	5, // 4: google.ads.googleads.v19.services.ShareablePreviewOrError.shareable_preview_result:type_name -> google.ads.googleads.v19.services.ShareablePreviewResult
	6, // 5: google.ads.googleads.v19.services.ShareablePreviewOrError.partial_failure_error:type_name -> google.rpc.Status
	0, // 6: google.ads.googleads.v19.services.ShareablePreviewService.GenerateShareablePreviews:input_type -> google.ads.googleads.v19.services.GenerateShareablePreviewsRequest
	3, // 7: google.ads.googleads.v19.services.ShareablePreviewService.GenerateShareablePreviews:output_type -> google.ads.googleads.v19.services.GenerateShareablePreviewsResponse
	7, // [7:8] is the sub-list for method output_type
	6, // [6:7] is the sub-list for method input_type
	6, // [6:6] is the sub-list for extension type_name
	6, // [6:6] is the sub-list for extension extendee
	0, // [0:6] is the sub-list for field type_name
}

func init() { file_services_shareable_preview_service_proto_init() }
func file_services_shareable_preview_service_proto_init() {
	if File_services_shareable_preview_service_proto != nil {
		return
	}
	file_services_shareable_preview_service_proto_msgTypes[4].OneofWrappers = []any{
		(*ShareablePreviewOrError_ShareablePreviewResult)(nil),
		(*ShareablePreviewOrError_PartialFailureError)(nil),
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: unsafe.Slice(unsafe.StringData(file_services_shareable_preview_service_proto_rawDesc), len(file_services_shareable_preview_service_proto_rawDesc)),
			NumEnums:      0,
			NumMessages:   6,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_services_shareable_preview_service_proto_goTypes,
		DependencyIndexes: file_services_shareable_preview_service_proto_depIdxs,
		MessageInfos:      file_services_shareable_preview_service_proto_msgTypes,
	}.Build()
	File_services_shareable_preview_service_proto = out.File
	file_services_shareable_preview_service_proto_goTypes = nil
	file_services_shareable_preview_service_proto_depIdxs = nil
}
