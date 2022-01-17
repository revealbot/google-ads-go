// Code generated by protoc-gen-go. DO NOT EDIT.
// source: google/ads/googleads/v0/services/ad_group_criterion_service.proto

package services

import (
	context "context"
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	_ "github.com/golang/protobuf/ptypes/wrappers"
	resources "github.com/revealbot/google-ads-go/resources"
	_ "google.golang.org/genproto/googleapis/api/annotations"
	status "google.golang.org/genproto/googleapis/rpc/status"
	field_mask "google.golang.org/genproto/protobuf/field_mask"
	grpc "google.golang.org/grpc"
	math "math"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion3 // please upgrade the proto package

// Request message for [AdGroupCriterionService.GetAdGroupCriterion][google.ads.googleads.v0.services.AdGroupCriterionService.GetAdGroupCriterion].
type GetAdGroupCriterionRequest struct {
	// The resource name of the criterion to fetch.
	ResourceName         string   `protobuf:"bytes,1,opt,name=resource_name,json=resourceName,proto3" json:"resource_name,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetAdGroupCriterionRequest) Reset()         { *m = GetAdGroupCriterionRequest{} }
func (m *GetAdGroupCriterionRequest) String() string { return proto.CompactTextString(m) }
func (*GetAdGroupCriterionRequest) ProtoMessage()    {}
func (*GetAdGroupCriterionRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_df48940ee39c701e, []int{0}
}

func (m *GetAdGroupCriterionRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetAdGroupCriterionRequest.Unmarshal(m, b)
}
func (m *GetAdGroupCriterionRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetAdGroupCriterionRequest.Marshal(b, m, deterministic)
}
func (m *GetAdGroupCriterionRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetAdGroupCriterionRequest.Merge(m, src)
}
func (m *GetAdGroupCriterionRequest) XXX_Size() int {
	return xxx_messageInfo_GetAdGroupCriterionRequest.Size(m)
}
func (m *GetAdGroupCriterionRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_GetAdGroupCriterionRequest.DiscardUnknown(m)
}

var xxx_messageInfo_GetAdGroupCriterionRequest proto.InternalMessageInfo

func (m *GetAdGroupCriterionRequest) GetResourceName() string {
	if m != nil {
		return m.ResourceName
	}
	return ""
}

// Request message for [AdGroupCriterionService.MutateAdGroupCriteria][google.ads.googleads.v0.services.AdGroupCriterionService.MutateAdGroupCriteria].
type MutateAdGroupCriteriaRequest struct {
	// ID of the customer whose criteria are being modified.
	CustomerId string `protobuf:"bytes,1,opt,name=customer_id,json=customerId,proto3" json:"customer_id,omitempty"`
	// The list of operations to perform on individual criteria.
	Operations []*AdGroupCriterionOperation `protobuf:"bytes,2,rep,name=operations,proto3" json:"operations,omitempty"`
	// If true, successful operations will be carried out and invalid
	// operations will return errors. If false, all operations will be carried
	// out in one transaction if and only if they are all valid.
	// Default is false.
	PartialFailure bool `protobuf:"varint,3,opt,name=partial_failure,json=partialFailure,proto3" json:"partial_failure,omitempty"`
	// If true, the request is validated but not executed. Only errors are
	// returned, not results.
	ValidateOnly         bool     `protobuf:"varint,4,opt,name=validate_only,json=validateOnly,proto3" json:"validate_only,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *MutateAdGroupCriteriaRequest) Reset()         { *m = MutateAdGroupCriteriaRequest{} }
func (m *MutateAdGroupCriteriaRequest) String() string { return proto.CompactTextString(m) }
func (*MutateAdGroupCriteriaRequest) ProtoMessage()    {}
func (*MutateAdGroupCriteriaRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_df48940ee39c701e, []int{1}
}

func (m *MutateAdGroupCriteriaRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_MutateAdGroupCriteriaRequest.Unmarshal(m, b)
}
func (m *MutateAdGroupCriteriaRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_MutateAdGroupCriteriaRequest.Marshal(b, m, deterministic)
}
func (m *MutateAdGroupCriteriaRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MutateAdGroupCriteriaRequest.Merge(m, src)
}
func (m *MutateAdGroupCriteriaRequest) XXX_Size() int {
	return xxx_messageInfo_MutateAdGroupCriteriaRequest.Size(m)
}
func (m *MutateAdGroupCriteriaRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_MutateAdGroupCriteriaRequest.DiscardUnknown(m)
}

var xxx_messageInfo_MutateAdGroupCriteriaRequest proto.InternalMessageInfo

func (m *MutateAdGroupCriteriaRequest) GetCustomerId() string {
	if m != nil {
		return m.CustomerId
	}
	return ""
}

func (m *MutateAdGroupCriteriaRequest) GetOperations() []*AdGroupCriterionOperation {
	if m != nil {
		return m.Operations
	}
	return nil
}

func (m *MutateAdGroupCriteriaRequest) GetPartialFailure() bool {
	if m != nil {
		return m.PartialFailure
	}
	return false
}

func (m *MutateAdGroupCriteriaRequest) GetValidateOnly() bool {
	if m != nil {
		return m.ValidateOnly
	}
	return false
}

// A single operation (create, remove, update) on an ad group criterion.
type AdGroupCriterionOperation struct {
	// FieldMask that determines which resource fields are modified in an update.
	UpdateMask *field_mask.FieldMask `protobuf:"bytes,4,opt,name=update_mask,json=updateMask,proto3" json:"update_mask,omitempty"`
	// The mutate operation.
	//
	// Types that are valid to be assigned to Operation:
	//	*AdGroupCriterionOperation_Create
	//	*AdGroupCriterionOperation_Update
	//	*AdGroupCriterionOperation_Remove
	Operation            isAdGroupCriterionOperation_Operation `protobuf_oneof:"operation"`
	XXX_NoUnkeyedLiteral struct{}                              `json:"-"`
	XXX_unrecognized     []byte                                `json:"-"`
	XXX_sizecache        int32                                 `json:"-"`
}

func (m *AdGroupCriterionOperation) Reset()         { *m = AdGroupCriterionOperation{} }
func (m *AdGroupCriterionOperation) String() string { return proto.CompactTextString(m) }
func (*AdGroupCriterionOperation) ProtoMessage()    {}
func (*AdGroupCriterionOperation) Descriptor() ([]byte, []int) {
	return fileDescriptor_df48940ee39c701e, []int{2}
}

func (m *AdGroupCriterionOperation) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_AdGroupCriterionOperation.Unmarshal(m, b)
}
func (m *AdGroupCriterionOperation) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_AdGroupCriterionOperation.Marshal(b, m, deterministic)
}
func (m *AdGroupCriterionOperation) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AdGroupCriterionOperation.Merge(m, src)
}
func (m *AdGroupCriterionOperation) XXX_Size() int {
	return xxx_messageInfo_AdGroupCriterionOperation.Size(m)
}
func (m *AdGroupCriterionOperation) XXX_DiscardUnknown() {
	xxx_messageInfo_AdGroupCriterionOperation.DiscardUnknown(m)
}

var xxx_messageInfo_AdGroupCriterionOperation proto.InternalMessageInfo

func (m *AdGroupCriterionOperation) GetUpdateMask() *field_mask.FieldMask {
	if m != nil {
		return m.UpdateMask
	}
	return nil
}

type isAdGroupCriterionOperation_Operation interface {
	isAdGroupCriterionOperation_Operation()
}

type AdGroupCriterionOperation_Create struct {
	Create *resources.AdGroupCriterion `protobuf:"bytes,1,opt,name=create,proto3,oneof"`
}

type AdGroupCriterionOperation_Update struct {
	Update *resources.AdGroupCriterion `protobuf:"bytes,2,opt,name=update,proto3,oneof"`
}

type AdGroupCriterionOperation_Remove struct {
	Remove string `protobuf:"bytes,3,opt,name=remove,proto3,oneof"`
}

func (*AdGroupCriterionOperation_Create) isAdGroupCriterionOperation_Operation() {}

func (*AdGroupCriterionOperation_Update) isAdGroupCriterionOperation_Operation() {}

func (*AdGroupCriterionOperation_Remove) isAdGroupCriterionOperation_Operation() {}

func (m *AdGroupCriterionOperation) GetOperation() isAdGroupCriterionOperation_Operation {
	if m != nil {
		return m.Operation
	}
	return nil
}

func (m *AdGroupCriterionOperation) GetCreate() *resources.AdGroupCriterion {
	if x, ok := m.GetOperation().(*AdGroupCriterionOperation_Create); ok {
		return x.Create
	}
	return nil
}

func (m *AdGroupCriterionOperation) GetUpdate() *resources.AdGroupCriterion {
	if x, ok := m.GetOperation().(*AdGroupCriterionOperation_Update); ok {
		return x.Update
	}
	return nil
}

func (m *AdGroupCriterionOperation) GetRemove() string {
	if x, ok := m.GetOperation().(*AdGroupCriterionOperation_Remove); ok {
		return x.Remove
	}
	return ""
}

// XXX_OneofWrappers is for the internal use of the proto package.
func (*AdGroupCriterionOperation) XXX_OneofWrappers() []interface{} {
	return []interface{}{
		(*AdGroupCriterionOperation_Create)(nil),
		(*AdGroupCriterionOperation_Update)(nil),
		(*AdGroupCriterionOperation_Remove)(nil),
	}
}

// Response message for an ad group criterion mutate.
type MutateAdGroupCriteriaResponse struct {
	// Errors that pertain to operation failures in the partial failure mode.
	// Returned only when partial_failure = true and all errors occur inside the
	// operations. If any errors occur outside the operations (e.g. auth errors),
	// we return an RPC level error.
	PartialFailureError *status.Status `protobuf:"bytes,3,opt,name=partial_failure_error,json=partialFailureError,proto3" json:"partial_failure_error,omitempty"`
	// All results for the mutate.
	Results              []*MutateAdGroupCriterionResult `protobuf:"bytes,2,rep,name=results,proto3" json:"results,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                        `json:"-"`
	XXX_unrecognized     []byte                          `json:"-"`
	XXX_sizecache        int32                           `json:"-"`
}

func (m *MutateAdGroupCriteriaResponse) Reset()         { *m = MutateAdGroupCriteriaResponse{} }
func (m *MutateAdGroupCriteriaResponse) String() string { return proto.CompactTextString(m) }
func (*MutateAdGroupCriteriaResponse) ProtoMessage()    {}
func (*MutateAdGroupCriteriaResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_df48940ee39c701e, []int{3}
}

func (m *MutateAdGroupCriteriaResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_MutateAdGroupCriteriaResponse.Unmarshal(m, b)
}
func (m *MutateAdGroupCriteriaResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_MutateAdGroupCriteriaResponse.Marshal(b, m, deterministic)
}
func (m *MutateAdGroupCriteriaResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MutateAdGroupCriteriaResponse.Merge(m, src)
}
func (m *MutateAdGroupCriteriaResponse) XXX_Size() int {
	return xxx_messageInfo_MutateAdGroupCriteriaResponse.Size(m)
}
func (m *MutateAdGroupCriteriaResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_MutateAdGroupCriteriaResponse.DiscardUnknown(m)
}

var xxx_messageInfo_MutateAdGroupCriteriaResponse proto.InternalMessageInfo

func (m *MutateAdGroupCriteriaResponse) GetPartialFailureError() *status.Status {
	if m != nil {
		return m.PartialFailureError
	}
	return nil
}

func (m *MutateAdGroupCriteriaResponse) GetResults() []*MutateAdGroupCriterionResult {
	if m != nil {
		return m.Results
	}
	return nil
}

// The result for the criterion mutate.
type MutateAdGroupCriterionResult struct {
	// Returned for successful operations.
	ResourceName         string   `protobuf:"bytes,1,opt,name=resource_name,json=resourceName,proto3" json:"resource_name,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *MutateAdGroupCriterionResult) Reset()         { *m = MutateAdGroupCriterionResult{} }
func (m *MutateAdGroupCriterionResult) String() string { return proto.CompactTextString(m) }
func (*MutateAdGroupCriterionResult) ProtoMessage()    {}
func (*MutateAdGroupCriterionResult) Descriptor() ([]byte, []int) {
	return fileDescriptor_df48940ee39c701e, []int{4}
}

func (m *MutateAdGroupCriterionResult) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_MutateAdGroupCriterionResult.Unmarshal(m, b)
}
func (m *MutateAdGroupCriterionResult) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_MutateAdGroupCriterionResult.Marshal(b, m, deterministic)
}
func (m *MutateAdGroupCriterionResult) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MutateAdGroupCriterionResult.Merge(m, src)
}
func (m *MutateAdGroupCriterionResult) XXX_Size() int {
	return xxx_messageInfo_MutateAdGroupCriterionResult.Size(m)
}
func (m *MutateAdGroupCriterionResult) XXX_DiscardUnknown() {
	xxx_messageInfo_MutateAdGroupCriterionResult.DiscardUnknown(m)
}

var xxx_messageInfo_MutateAdGroupCriterionResult proto.InternalMessageInfo

func (m *MutateAdGroupCriterionResult) GetResourceName() string {
	if m != nil {
		return m.ResourceName
	}
	return ""
}

func init() {
	proto.RegisterType((*GetAdGroupCriterionRequest)(nil), "google.ads.googleads.v0.services.GetAdGroupCriterionRequest")
	proto.RegisterType((*MutateAdGroupCriteriaRequest)(nil), "google.ads.googleads.v0.services.MutateAdGroupCriteriaRequest")
	proto.RegisterType((*AdGroupCriterionOperation)(nil), "google.ads.googleads.v0.services.AdGroupCriterionOperation")
	proto.RegisterType((*MutateAdGroupCriteriaResponse)(nil), "google.ads.googleads.v0.services.MutateAdGroupCriteriaResponse")
	proto.RegisterType((*MutateAdGroupCriterionResult)(nil), "google.ads.googleads.v0.services.MutateAdGroupCriterionResult")
}

func init() {
	proto.RegisterFile("google/ads/googleads/v0/services/ad_group_criterion_service.proto", fileDescriptor_df48940ee39c701e)
}

var fileDescriptor_df48940ee39c701e = []byte{
	// 722 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xa4, 0x95, 0x4f, 0x4f, 0xd4, 0x4c,
	0x1c, 0xc7, 0x9f, 0x76, 0x9f, 0xf0, 0x3c, 0xcc, 0xf2, 0x3c, 0x26, 0x43, 0x08, 0xeb, 0x06, 0x75,
	0x53, 0x49, 0x24, 0x7b, 0x68, 0xd7, 0xc5, 0x0b, 0x45, 0x24, 0x85, 0xc8, 0xe2, 0x01, 0x21, 0x25,
	0x21, 0x46, 0x37, 0x69, 0x86, 0x76, 0x68, 0x1a, 0xda, 0x4e, 0x9d, 0x99, 0xae, 0x21, 0x84, 0x8b,
	0xbe, 0x04, 0xcf, 0x5e, 0x3c, 0xfa, 0x16, 0x3c, 0x19, 0x6f, 0x5e, 0x3d, 0x79, 0xf7, 0xc4, 0xab,
	0x30, 0xd3, 0xe9, 0xac, 0xec, 0xba, 0xcd, 0x1a, 0xbc, 0xfd, 0xfa, 0x9b, 0xef, 0x7c, 0xe6, 0xf7,
	0x6f, 0xa6, 0xc0, 0x09, 0x09, 0x09, 0x63, 0x6c, 0xa1, 0x80, 0x59, 0xd2, 0x14, 0xd6, 0xa0, 0x63,
	0x31, 0x4c, 0x07, 0x91, 0x8f, 0x99, 0x85, 0x02, 0x2f, 0xa4, 0x24, 0xcf, 0x3c, 0x9f, 0x46, 0x1c,
	0xd3, 0x88, 0xa4, 0x5e, 0xb9, 0x66, 0x66, 0x94, 0x70, 0x02, 0x5b, 0x72, 0x9f, 0x89, 0x02, 0x66,
	0x0e, 0x11, 0xe6, 0xa0, 0x63, 0x2a, 0x44, 0xd3, 0xae, 0x3a, 0x84, 0x62, 0x46, 0x72, 0x3a, 0xf9,
	0x14, 0x49, 0x6f, 0x2e, 0xa9, 0xbd, 0x59, 0x64, 0xa1, 0x34, 0x25, 0x1c, 0xf1, 0x88, 0xa4, 0xac,
	0x5c, 0x2d, 0xcf, 0xb6, 0x8a, 0xaf, 0xe3, 0xfc, 0xc4, 0x3a, 0x89, 0x70, 0x1c, 0x78, 0x09, 0x62,
	0xa7, 0xa5, 0xe2, 0xf6, 0xb8, 0xe2, 0x15, 0x45, 0x59, 0x86, 0xa9, 0x22, 0x2c, 0x96, 0xeb, 0x34,
	0xf3, 0x2d, 0xc6, 0x11, 0xcf, 0xcb, 0x05, 0xc3, 0x01, 0xcd, 0x1e, 0xe6, 0x4e, 0xd0, 0x13, 0x61,
	0x6d, 0xab, 0xa8, 0x5c, 0xfc, 0x32, 0xc7, 0x8c, 0xc3, 0xbb, 0xe0, 0x3f, 0x15, 0xbc, 0x97, 0xa2,
	0x04, 0x37, 0xb4, 0x96, 0xb6, 0x32, 0xeb, 0xce, 0x29, 0xe7, 0x53, 0x94, 0x60, 0xe3, 0x52, 0x03,
	0x4b, 0x7b, 0x39, 0x47, 0x1c, 0x8f, 0x62, 0x90, 0xa2, 0xdc, 0x01, 0x75, 0x3f, 0x67, 0x9c, 0x24,
	0x98, 0x7a, 0x51, 0x50, 0x32, 0x80, 0x72, 0x3d, 0x09, 0xe0, 0x0b, 0x00, 0x48, 0x86, 0xa9, 0xcc,
	0xb9, 0xa1, 0xb7, 0x6a, 0x2b, 0xf5, 0xee, 0xba, 0x39, 0xad, 0xe0, 0xe6, 0x78, 0xd4, 0xfb, 0x8a,
	0xe1, 0x5e, 0xc1, 0xc1, 0x7b, 0xe0, 0x46, 0x86, 0x28, 0x8f, 0x50, 0xec, 0x9d, 0xa0, 0x28, 0xce,
	0x29, 0x6e, 0xd4, 0x5a, 0xda, 0xca, 0xbf, 0xee, 0xff, 0xa5, 0x7b, 0x47, 0x7a, 0x45, 0xb2, 0x03,
	0x14, 0x47, 0x01, 0xe2, 0xd8, 0x23, 0x69, 0x7c, 0xd6, 0xf8, 0xbb, 0x90, 0xcd, 0x29, 0xe7, 0x7e,
	0x1a, 0x9f, 0x19, 0xef, 0x74, 0x70, 0xb3, 0xf2, 0x5c, 0xb8, 0x0e, 0xea, 0x79, 0x56, 0x00, 0x44,
	0x6f, 0x0a, 0x40, 0xbd, 0xdb, 0x54, 0x99, 0xa8, 0xe6, 0x98, 0x3b, 0xa2, 0x7d, 0x7b, 0x88, 0x9d,
	0xba, 0x40, 0xca, 0x85, 0x0d, 0xf7, 0xc0, 0x8c, 0x4f, 0x31, 0xe2, 0xb2, 0xca, 0xf5, 0xee, 0x6a,
	0x65, 0x05, 0x86, 0x03, 0xf5, 0x4b, 0x09, 0x76, 0xff, 0x72, 0x4b, 0x88, 0xc0, 0x49, 0x78, 0x43,
	0xff, 0x23, 0x9c, 0x84, 0xc0, 0x06, 0x98, 0xa1, 0x38, 0x21, 0x03, 0x59, 0xbd, 0x59, 0xb1, 0x22,
	0xbf, 0xb7, 0xea, 0x60, 0x76, 0x58, 0x6e, 0xe3, 0x93, 0x06, 0x6e, 0x55, 0x0c, 0x03, 0xcb, 0x48,
	0xca, 0x30, 0xdc, 0x01, 0x0b, 0x63, 0xfd, 0xf0, 0x30, 0xa5, 0x84, 0x16, 0xdc, 0x7a, 0x17, 0xaa,
	0x30, 0x69, 0xe6, 0x9b, 0x87, 0xc5, 0xa8, 0xba, 0xf3, 0xa3, 0x9d, 0x7a, 0x2c, 0xe4, 0xf0, 0x19,
	0xf8, 0x87, 0x62, 0x96, 0xc7, 0x5c, 0x4d, 0xcc, 0xa3, 0xe9, 0x13, 0x33, 0x29, 0x32, 0x31, 0xed,
	0x02, 0xe3, 0x2a, 0x9c, 0xb1, 0x3d, 0x79, 0x9e, 0x95, 0xf0, 0xb7, 0x6e, 0x45, 0xf7, 0x63, 0x0d,
	0x2c, 0x8e, 0xef, 0x3f, 0x94, 0x71, 0xc0, 0xcf, 0x1a, 0x98, 0x9f, 0x70, 0xeb, 0xe0, 0xc3, 0xe9,
	0x19, 0x54, 0x5f, 0xd6, 0xe6, 0x75, 0x1a, 0x6c, 0xac, 0xbd, 0xfe, 0xfa, 0xfd, 0xad, 0xbe, 0x0a,
	0xef, 0x8b, 0x87, 0xea, 0x7c, 0x24, 0xad, 0x0d, 0x75, 0x43, 0x99, 0xd5, 0xb6, 0xd0, 0x68, 0x37,
	0xad, 0xf6, 0x05, 0xfc, 0xa6, 0x81, 0x85, 0x89, 0xad, 0x86, 0xd7, 0xec, 0x84, 0x7a, 0x30, 0x9a,
	0x9b, 0xd7, 0xde, 0x2f, 0x67, 0xcc, 0xd8, 0x2c, 0xb2, 0x5a, 0x33, 0x1e, 0x88, 0xac, 0x7e, 0xa6,
	0x71, 0x7e, 0xe5, 0x19, 0xda, 0x68, 0x5f, 0x8c, 0x27, 0x65, 0x27, 0x05, 0xd4, 0xd6, 0xda, 0x5b,
	0x6f, 0x74, 0xb0, 0xec, 0x93, 0x64, 0x6a, 0x1c, 0x5b, 0x4b, 0x15, 0x3d, 0x3e, 0x10, 0x77, 0xfd,
	0x40, 0x7b, 0xbe, 0x5b, 0x12, 0x42, 0x12, 0xa3, 0x34, 0x34, 0x09, 0x0d, 0xad, 0x10, 0xa7, 0xc5,
	0x4b, 0xa0, 0x7e, 0x12, 0x59, 0xc4, 0xaa, 0x7f, 0x4c, 0xeb, 0xca, 0x78, 0xaf, 0xd7, 0x7a, 0x8e,
	0xf3, 0x41, 0x6f, 0xf5, 0x24, 0xd0, 0x09, 0x98, 0x29, 0x4d, 0x61, 0x1d, 0x75, 0xcc, 0xf2, 0x60,
	0xf6, 0x45, 0x49, 0xfa, 0x4e, 0xc0, 0xfa, 0x43, 0x49, 0xff, 0xa8, 0xd3, 0x57, 0x92, 0x4b, 0x7d,
	0x59, 0xfa, 0x6d, 0xdb, 0x09, 0x98, 0x6d, 0x0f, 0x45, 0xb6, 0x7d, 0xd4, 0xb1, 0x6d, 0x25, 0x3b,
	0x9e, 0x29, 0xe2, 0x5c, 0xfd, 0x11, 0x00, 0x00, 0xff, 0xff, 0xb6, 0xb7, 0x0b, 0x04, 0x3f, 0x07,
	0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// AdGroupCriterionServiceClient is the client API for AdGroupCriterionService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type AdGroupCriterionServiceClient interface {
	// Returns the requested criterion in full detail.
	GetAdGroupCriterion(ctx context.Context, in *GetAdGroupCriterionRequest, opts ...grpc.CallOption) (*resources.AdGroupCriterion, error)
	// Creates, updates, or removes criteria. Operation statuses are returned.
	MutateAdGroupCriteria(ctx context.Context, in *MutateAdGroupCriteriaRequest, opts ...grpc.CallOption) (*MutateAdGroupCriteriaResponse, error)
}

type adGroupCriterionServiceClient struct {
	cc *grpc.ClientConn
}

func NewAdGroupCriterionServiceClient(cc *grpc.ClientConn) AdGroupCriterionServiceClient {
	return &adGroupCriterionServiceClient{cc}
}

func (c *adGroupCriterionServiceClient) GetAdGroupCriterion(ctx context.Context, in *GetAdGroupCriterionRequest, opts ...grpc.CallOption) (*resources.AdGroupCriterion, error) {
	out := new(resources.AdGroupCriterion)
	err := c.cc.Invoke(ctx, "/google.ads.googleads.v0.services.AdGroupCriterionService/GetAdGroupCriterion", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *adGroupCriterionServiceClient) MutateAdGroupCriteria(ctx context.Context, in *MutateAdGroupCriteriaRequest, opts ...grpc.CallOption) (*MutateAdGroupCriteriaResponse, error) {
	out := new(MutateAdGroupCriteriaResponse)
	err := c.cc.Invoke(ctx, "/google.ads.googleads.v0.services.AdGroupCriterionService/MutateAdGroupCriteria", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// AdGroupCriterionServiceServer is the server API for AdGroupCriterionService service.
type AdGroupCriterionServiceServer interface {
	// Returns the requested criterion in full detail.
	GetAdGroupCriterion(context.Context, *GetAdGroupCriterionRequest) (*resources.AdGroupCriterion, error)
	// Creates, updates, or removes criteria. Operation statuses are returned.
	MutateAdGroupCriteria(context.Context, *MutateAdGroupCriteriaRequest) (*MutateAdGroupCriteriaResponse, error)
}

func RegisterAdGroupCriterionServiceServer(s *grpc.Server, srv AdGroupCriterionServiceServer) {
	s.RegisterService(&_AdGroupCriterionService_serviceDesc, srv)
}

func _AdGroupCriterionService_GetAdGroupCriterion_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetAdGroupCriterionRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AdGroupCriterionServiceServer).GetAdGroupCriterion(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/google.ads.googleads.v0.services.AdGroupCriterionService/GetAdGroupCriterion",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AdGroupCriterionServiceServer).GetAdGroupCriterion(ctx, req.(*GetAdGroupCriterionRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _AdGroupCriterionService_MutateAdGroupCriteria_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MutateAdGroupCriteriaRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AdGroupCriterionServiceServer).MutateAdGroupCriteria(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/google.ads.googleads.v0.services.AdGroupCriterionService/MutateAdGroupCriteria",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AdGroupCriterionServiceServer).MutateAdGroupCriteria(ctx, req.(*MutateAdGroupCriteriaRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _AdGroupCriterionService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "google.ads.googleads.v0.services.AdGroupCriterionService",
	HandlerType: (*AdGroupCriterionServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetAdGroupCriterion",
			Handler:    _AdGroupCriterionService_GetAdGroupCriterion_Handler,
		},
		{
			MethodName: "MutateAdGroupCriteria",
			Handler:    _AdGroupCriterionService_MutateAdGroupCriteria_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "google/ads/googleads/v0/services/ad_group_criterion_service.proto",
}
