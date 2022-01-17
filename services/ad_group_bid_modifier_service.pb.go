// Code generated by protoc-gen-go. DO NOT EDIT.
// source: google/ads/googleads/v9/services/ad_group_bid_modifier_service.proto

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

// Request message for [AdGroupBidModifierService.GetAdGroupBidModifier][google.ads.googleads.v9.services.AdGroupBidModifierService.GetAdGroupBidModifier].
type GetAdGroupBidModifierRequest struct {
	// The resource name of the ad group bid modifier to fetch.
	ResourceName         string   `protobuf:"bytes,1,opt,name=resource_name,json=resourceName,proto3" json:"resource_name,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetAdGroupBidModifierRequest) Reset()         { *m = GetAdGroupBidModifierRequest{} }
func (m *GetAdGroupBidModifierRequest) String() string { return proto.CompactTextString(m) }
func (*GetAdGroupBidModifierRequest) ProtoMessage()    {}
func (*GetAdGroupBidModifierRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_268835c995d32f7c, []int{0}
}

func (m *GetAdGroupBidModifierRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetAdGroupBidModifierRequest.Unmarshal(m, b)
}
func (m *GetAdGroupBidModifierRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetAdGroupBidModifierRequest.Marshal(b, m, deterministic)
}
func (m *GetAdGroupBidModifierRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetAdGroupBidModifierRequest.Merge(m, src)
}
func (m *GetAdGroupBidModifierRequest) XXX_Size() int {
	return xxx_messageInfo_GetAdGroupBidModifierRequest.Size(m)
}
func (m *GetAdGroupBidModifierRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_GetAdGroupBidModifierRequest.DiscardUnknown(m)
}

var xxx_messageInfo_GetAdGroupBidModifierRequest proto.InternalMessageInfo

func (m *GetAdGroupBidModifierRequest) GetResourceName() string {
	if m != nil {
		return m.ResourceName
	}
	return ""
}

// Request message for [AdGroupBidModifierService.MutateAdGroupBidModifiers][google.ads.googleads.v9.services.AdGroupBidModifierService.MutateAdGroupBidModifiers].
type MutateAdGroupBidModifiersRequest struct {
	// ID of the customer whose ad group bid modifiers are being modified.
	CustomerId string `protobuf:"bytes,1,opt,name=customer_id,json=customerId,proto3" json:"customer_id,omitempty"`
	// The list of operations to perform on individual ad group bid modifiers.
	Operations []*AdGroupBidModifierOperation `protobuf:"bytes,2,rep,name=operations,proto3" json:"operations,omitempty"`
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

func (m *MutateAdGroupBidModifiersRequest) Reset()         { *m = MutateAdGroupBidModifiersRequest{} }
func (m *MutateAdGroupBidModifiersRequest) String() string { return proto.CompactTextString(m) }
func (*MutateAdGroupBidModifiersRequest) ProtoMessage()    {}
func (*MutateAdGroupBidModifiersRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_268835c995d32f7c, []int{1}
}

func (m *MutateAdGroupBidModifiersRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_MutateAdGroupBidModifiersRequest.Unmarshal(m, b)
}
func (m *MutateAdGroupBidModifiersRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_MutateAdGroupBidModifiersRequest.Marshal(b, m, deterministic)
}
func (m *MutateAdGroupBidModifiersRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MutateAdGroupBidModifiersRequest.Merge(m, src)
}
func (m *MutateAdGroupBidModifiersRequest) XXX_Size() int {
	return xxx_messageInfo_MutateAdGroupBidModifiersRequest.Size(m)
}
func (m *MutateAdGroupBidModifiersRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_MutateAdGroupBidModifiersRequest.DiscardUnknown(m)
}

var xxx_messageInfo_MutateAdGroupBidModifiersRequest proto.InternalMessageInfo

func (m *MutateAdGroupBidModifiersRequest) GetCustomerId() string {
	if m != nil {
		return m.CustomerId
	}
	return ""
}

func (m *MutateAdGroupBidModifiersRequest) GetOperations() []*AdGroupBidModifierOperation {
	if m != nil {
		return m.Operations
	}
	return nil
}

func (m *MutateAdGroupBidModifiersRequest) GetPartialFailure() bool {
	if m != nil {
		return m.PartialFailure
	}
	return false
}

func (m *MutateAdGroupBidModifiersRequest) GetValidateOnly() bool {
	if m != nil {
		return m.ValidateOnly
	}
	return false
}

// A single operation (create, remove, update) on an ad group bid modifier.
type AdGroupBidModifierOperation struct {
	// FieldMask that determines which resource fields are modified in an update.
	UpdateMask *field_mask.FieldMask `protobuf:"bytes,4,opt,name=update_mask,json=updateMask,proto3" json:"update_mask,omitempty"`
	// The mutate operation.
	//
	// Types that are valid to be assigned to Operation:
	//	*AdGroupBidModifierOperation_Create
	//	*AdGroupBidModifierOperation_Update
	//	*AdGroupBidModifierOperation_Remove
	Operation            isAdGroupBidModifierOperation_Operation `protobuf_oneof:"operation"`
	XXX_NoUnkeyedLiteral struct{}                                `json:"-"`
	XXX_unrecognized     []byte                                  `json:"-"`
	XXX_sizecache        int32                                   `json:"-"`
}

func (m *AdGroupBidModifierOperation) Reset()         { *m = AdGroupBidModifierOperation{} }
func (m *AdGroupBidModifierOperation) String() string { return proto.CompactTextString(m) }
func (*AdGroupBidModifierOperation) ProtoMessage()    {}
func (*AdGroupBidModifierOperation) Descriptor() ([]byte, []int) {
	return fileDescriptor_268835c995d32f7c, []int{2}
}

func (m *AdGroupBidModifierOperation) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_AdGroupBidModifierOperation.Unmarshal(m, b)
}
func (m *AdGroupBidModifierOperation) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_AdGroupBidModifierOperation.Marshal(b, m, deterministic)
}
func (m *AdGroupBidModifierOperation) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AdGroupBidModifierOperation.Merge(m, src)
}
func (m *AdGroupBidModifierOperation) XXX_Size() int {
	return xxx_messageInfo_AdGroupBidModifierOperation.Size(m)
}
func (m *AdGroupBidModifierOperation) XXX_DiscardUnknown() {
	xxx_messageInfo_AdGroupBidModifierOperation.DiscardUnknown(m)
}

var xxx_messageInfo_AdGroupBidModifierOperation proto.InternalMessageInfo

func (m *AdGroupBidModifierOperation) GetUpdateMask() *field_mask.FieldMask {
	if m != nil {
		return m.UpdateMask
	}
	return nil
}

type isAdGroupBidModifierOperation_Operation interface {
	isAdGroupBidModifierOperation_Operation()
}

type AdGroupBidModifierOperation_Create struct {
	Create *resources.AdGroupBidModifier `protobuf:"bytes,1,opt,name=create,proto3,oneof"`
}

type AdGroupBidModifierOperation_Update struct {
	Update *resources.AdGroupBidModifier `protobuf:"bytes,2,opt,name=update,proto3,oneof"`
}

type AdGroupBidModifierOperation_Remove struct {
	Remove string `protobuf:"bytes,3,opt,name=remove,proto3,oneof"`
}

func (*AdGroupBidModifierOperation_Create) isAdGroupBidModifierOperation_Operation() {}

func (*AdGroupBidModifierOperation_Update) isAdGroupBidModifierOperation_Operation() {}

func (*AdGroupBidModifierOperation_Remove) isAdGroupBidModifierOperation_Operation() {}

func (m *AdGroupBidModifierOperation) GetOperation() isAdGroupBidModifierOperation_Operation {
	if m != nil {
		return m.Operation
	}
	return nil
}

func (m *AdGroupBidModifierOperation) GetCreate() *resources.AdGroupBidModifier {
	if x, ok := m.GetOperation().(*AdGroupBidModifierOperation_Create); ok {
		return x.Create
	}
	return nil
}

func (m *AdGroupBidModifierOperation) GetUpdate() *resources.AdGroupBidModifier {
	if x, ok := m.GetOperation().(*AdGroupBidModifierOperation_Update); ok {
		return x.Update
	}
	return nil
}

func (m *AdGroupBidModifierOperation) GetRemove() string {
	if x, ok := m.GetOperation().(*AdGroupBidModifierOperation_Remove); ok {
		return x.Remove
	}
	return ""
}

// XXX_OneofWrappers is for the internal use of the proto package.
func (*AdGroupBidModifierOperation) XXX_OneofWrappers() []interface{} {
	return []interface{}{
		(*AdGroupBidModifierOperation_Create)(nil),
		(*AdGroupBidModifierOperation_Update)(nil),
		(*AdGroupBidModifierOperation_Remove)(nil),
	}
}

// Response message for ad group bid modifiers mutate.
type MutateAdGroupBidModifiersResponse struct {
	// Errors that pertain to operation failures in the partial failure mode.
	// Returned only when partial_failure = true and all errors occur inside the
	// operations. If any errors occur outside the operations (e.g. auth errors),
	// we return an RPC level error.
	PartialFailureError *status.Status `protobuf:"bytes,3,opt,name=partial_failure_error,json=partialFailureError,proto3" json:"partial_failure_error,omitempty"`
	// All results for the mutate.
	Results              []*MutateAdGroupBidModifierResult `protobuf:"bytes,2,rep,name=results,proto3" json:"results,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                          `json:"-"`
	XXX_unrecognized     []byte                            `json:"-"`
	XXX_sizecache        int32                             `json:"-"`
}

func (m *MutateAdGroupBidModifiersResponse) Reset()         { *m = MutateAdGroupBidModifiersResponse{} }
func (m *MutateAdGroupBidModifiersResponse) String() string { return proto.CompactTextString(m) }
func (*MutateAdGroupBidModifiersResponse) ProtoMessage()    {}
func (*MutateAdGroupBidModifiersResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_268835c995d32f7c, []int{3}
}

func (m *MutateAdGroupBidModifiersResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_MutateAdGroupBidModifiersResponse.Unmarshal(m, b)
}
func (m *MutateAdGroupBidModifiersResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_MutateAdGroupBidModifiersResponse.Marshal(b, m, deterministic)
}
func (m *MutateAdGroupBidModifiersResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MutateAdGroupBidModifiersResponse.Merge(m, src)
}
func (m *MutateAdGroupBidModifiersResponse) XXX_Size() int {
	return xxx_messageInfo_MutateAdGroupBidModifiersResponse.Size(m)
}
func (m *MutateAdGroupBidModifiersResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_MutateAdGroupBidModifiersResponse.DiscardUnknown(m)
}

var xxx_messageInfo_MutateAdGroupBidModifiersResponse proto.InternalMessageInfo

func (m *MutateAdGroupBidModifiersResponse) GetPartialFailureError() *status.Status {
	if m != nil {
		return m.PartialFailureError
	}
	return nil
}

func (m *MutateAdGroupBidModifiersResponse) GetResults() []*MutateAdGroupBidModifierResult {
	if m != nil {
		return m.Results
	}
	return nil
}

// The result for the criterion mutate.
type MutateAdGroupBidModifierResult struct {
	// Returned for successful operations.
	ResourceName         string   `protobuf:"bytes,1,opt,name=resource_name,json=resourceName,proto3" json:"resource_name,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *MutateAdGroupBidModifierResult) Reset()         { *m = MutateAdGroupBidModifierResult{} }
func (m *MutateAdGroupBidModifierResult) String() string { return proto.CompactTextString(m) }
func (*MutateAdGroupBidModifierResult) ProtoMessage()    {}
func (*MutateAdGroupBidModifierResult) Descriptor() ([]byte, []int) {
	return fileDescriptor_268835c995d32f7c, []int{4}
}

func (m *MutateAdGroupBidModifierResult) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_MutateAdGroupBidModifierResult.Unmarshal(m, b)
}
func (m *MutateAdGroupBidModifierResult) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_MutateAdGroupBidModifierResult.Marshal(b, m, deterministic)
}
func (m *MutateAdGroupBidModifierResult) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MutateAdGroupBidModifierResult.Merge(m, src)
}
func (m *MutateAdGroupBidModifierResult) XXX_Size() int {
	return xxx_messageInfo_MutateAdGroupBidModifierResult.Size(m)
}
func (m *MutateAdGroupBidModifierResult) XXX_DiscardUnknown() {
	xxx_messageInfo_MutateAdGroupBidModifierResult.DiscardUnknown(m)
}

var xxx_messageInfo_MutateAdGroupBidModifierResult proto.InternalMessageInfo

func (m *MutateAdGroupBidModifierResult) GetResourceName() string {
	if m != nil {
		return m.ResourceName
	}
	return ""
}

func init() {
	proto.RegisterType((*GetAdGroupBidModifierRequest)(nil), "google.ads.googleads.v9.services.GetAdGroupBidModifierRequest")
	proto.RegisterType((*MutateAdGroupBidModifiersRequest)(nil), "google.ads.googleads.v9.services.MutateAdGroupBidModifiersRequest")
	proto.RegisterType((*AdGroupBidModifierOperation)(nil), "google.ads.googleads.v9.services.AdGroupBidModifierOperation")
	proto.RegisterType((*MutateAdGroupBidModifiersResponse)(nil), "google.ads.googleads.v9.services.MutateAdGroupBidModifiersResponse")
	proto.RegisterType((*MutateAdGroupBidModifierResult)(nil), "google.ads.googleads.v9.services.MutateAdGroupBidModifierResult")
}

func init() {
	proto.RegisterFile("google/ads/googleads/v9/services/ad_group_bid_modifier_service.proto", fileDescriptor_268835c995d32f7c)
}

var fileDescriptor_268835c995d32f7c = []byte{
	// 730 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xa4, 0x95, 0xcd, 0x6e, 0xd3, 0x4a,
	0x14, 0xc7, 0xaf, 0x9d, 0xab, 0xde, 0xdb, 0x49, 0xef, 0x45, 0x1a, 0x54, 0x91, 0x86, 0xaa, 0x04,
	0x53, 0x89, 0x2a, 0x0b, 0x3b, 0x0a, 0xaa, 0x40, 0xae, 0x82, 0x48, 0x4a, 0x9b, 0xb2, 0x28, 0xad,
	0x5c, 0xa9, 0x8b, 0x2a, 0xc8, 0x9a, 0x66, 0x26, 0x96, 0x55, 0xdb, 0x63, 0x66, 0xc6, 0x41, 0x55,
	0xd5, 0x0d, 0x12, 0x4f, 0xc0, 0x1b, 0xc0, 0x8e, 0x17, 0x41, 0x42, 0x62, 0xc5, 0x82, 0x17, 0x60,
	0x03, 0x2b, 0x1e, 0x01, 0xd9, 0xe3, 0x09, 0x6d, 0x53, 0x27, 0xa8, 0xdd, 0x9d, 0x9c, 0xf9, 0xe7,
	0x77, 0x3e, 0x67, 0x0c, 0x9e, 0x7a, 0x94, 0x7a, 0x01, 0xb1, 0x10, 0xe6, 0x96, 0x34, 0x53, 0x6b,
	0xd8, 0xb0, 0x38, 0x61, 0x43, 0xbf, 0x4f, 0xb8, 0x85, 0xb0, 0xeb, 0x31, 0x9a, 0xc4, 0xee, 0xa1,
	0x8f, 0xdd, 0x90, 0x62, 0x7f, 0xe0, 0x13, 0xe6, 0xe6, 0xc7, 0x66, 0xcc, 0xa8, 0xa0, 0xb0, 0x26,
	0xff, 0x6a, 0x22, 0xcc, 0xcd, 0x11, 0xc5, 0x1c, 0x36, 0x4c, 0x45, 0xa9, 0xb6, 0x8a, 0xe2, 0x30,
	0xc2, 0x69, 0xc2, 0x0a, 0x03, 0xc9, 0x00, 0xd5, 0x45, 0xf5, 0xf7, 0xd8, 0xb7, 0x50, 0x14, 0x51,
	0x81, 0x84, 0x4f, 0x23, 0x9e, 0x9f, 0xe6, 0xe1, 0xad, 0xec, 0xd7, 0x61, 0x32, 0xb0, 0x06, 0x3e,
	0x09, 0xb0, 0x1b, 0x22, 0x7e, 0x94, 0x2b, 0x96, 0x2e, 0x2a, 0x5e, 0x31, 0x14, 0xc7, 0x84, 0x29,
	0xc2, 0xad, 0xfc, 0x9c, 0xc5, 0x7d, 0x8b, 0x0b, 0x24, 0x92, 0xfc, 0xc0, 0x58, 0x07, 0x8b, 0x5d,
	0x22, 0xda, 0xb8, 0x9b, 0x66, 0xd6, 0xf1, 0xf1, 0x76, 0x9e, 0x97, 0x43, 0x5e, 0x26, 0x84, 0x0b,
	0x78, 0x0f, 0xfc, 0xa7, 0x2a, 0x70, 0x23, 0x14, 0x92, 0x8a, 0x56, 0xd3, 0x56, 0x66, 0x9d, 0x39,
	0xe5, 0x7c, 0x8e, 0x42, 0x62, 0xfc, 0xd4, 0x40, 0x6d, 0x3b, 0x11, 0x48, 0x90, 0x71, 0x10, 0x57,
	0xa4, 0x3b, 0xa0, 0xdc, 0x4f, 0xb8, 0xa0, 0x21, 0x61, 0xae, 0x8f, 0x73, 0x0e, 0x50, 0xae, 0x67,
	0x18, 0xbe, 0x00, 0x80, 0xc6, 0x84, 0xc9, 0xca, 0x2b, 0x7a, 0xad, 0xb4, 0x52, 0x6e, 0xb6, 0xcc,
	0x69, 0x9d, 0x37, 0xc7, 0x43, 0xee, 0x28, 0x8a, 0x73, 0x06, 0x08, 0xef, 0x83, 0x1b, 0x31, 0x62,
	0xc2, 0x47, 0x81, 0x3b, 0x40, 0x7e, 0x90, 0x30, 0x52, 0x29, 0xd5, 0xb4, 0x95, 0x7f, 0x9d, 0xff,
	0x73, 0xf7, 0xa6, 0xf4, 0xa6, 0x25, 0x0f, 0x51, 0xe0, 0x63, 0x24, 0x88, 0x4b, 0xa3, 0xe0, 0xb8,
	0xf2, 0x77, 0x26, 0x9b, 0x53, 0xce, 0x9d, 0x28, 0x38, 0x36, 0xde, 0xeb, 0xe0, 0xf6, 0x84, 0xc8,
	0x70, 0x0d, 0x94, 0x93, 0x38, 0x43, 0xa4, 0x53, 0xca, 0x10, 0xe5, 0x66, 0x55, 0x55, 0xa3, 0xc6,
	0x64, 0x6e, 0xa6, 0x83, 0xdc, 0x46, 0xfc, 0xc8, 0x01, 0x52, 0x9e, 0xda, 0x70, 0x07, 0xcc, 0xf4,
	0x19, 0x41, 0x42, 0x76, 0xbb, 0xdc, 0x5c, 0x2d, 0xec, 0xc2, 0x68, 0xbb, 0x2e, 0x69, 0xc3, 0xd6,
	0x5f, 0x4e, 0x8e, 0x49, 0x81, 0x12, 0x5f, 0xd1, 0xaf, 0x09, 0x94, 0x18, 0x58, 0x01, 0x33, 0x8c,
	0x84, 0x74, 0x28, 0x7b, 0x38, 0x9b, 0x9e, 0xc8, 0xdf, 0x9d, 0x32, 0x98, 0x1d, 0x35, 0xdd, 0xf8,
	0xa8, 0x81, 0xbb, 0x13, 0x16, 0x83, 0xc7, 0x34, 0xe2, 0x04, 0x6e, 0x82, 0xf9, 0x0b, 0x93, 0x71,
	0x09, 0x63, 0x94, 0x65, 0xec, 0x72, 0x13, 0xaa, 0x64, 0x59, 0xdc, 0x37, 0xf7, 0xb2, 0xe5, 0x75,
	0x6e, 0x9e, 0x9f, 0xd9, 0x46, 0x2a, 0x87, 0x07, 0xe0, 0x1f, 0x46, 0x78, 0x12, 0x08, 0xb5, 0x3d,
	0x4f, 0xa6, 0x6f, 0x4f, 0x51, 0x76, 0x4e, 0x06, 0x72, 0x14, 0xd0, 0xd8, 0x00, 0x4b, 0x93, 0xa5,
	0x7f, 0x74, 0x53, 0x9a, 0x5f, 0x4b, 0x60, 0x61, 0x9c, 0xb0, 0x27, 0xb3, 0x81, 0x9f, 0x35, 0x30,
	0x7f, 0xe9, 0x6d, 0x84, 0x8f, 0xa7, 0x57, 0x32, 0xe9, 0x1a, 0x57, 0xaf, 0x36, 0x70, 0xa3, 0xf5,
	0xfa, 0xcb, 0xb7, 0xb7, 0xfa, 0x43, 0xb8, 0x9a, 0xbe, 0x64, 0x27, 0xe7, 0xca, 0x6b, 0xa9, 0x9b,
	0xcb, 0xad, 0xba, 0x85, 0xc6, 0xa7, 0x6b, 0xd5, 0x4f, 0xe1, 0x77, 0x0d, 0x2c, 0x14, 0x8e, 0x1f,
	0x76, 0xae, 0x3e, 0x1d, 0xf5, 0xa8, 0x54, 0xd7, 0xaf, 0xc5, 0x90, 0xfb, 0x67, 0xac, 0x67, 0x55,
	0xb6, 0x8c, 0x47, 0x69, 0x95, 0xbf, 0xcb, 0x3a, 0x39, 0xf3, 0x5c, 0xb5, 0xea, 0xa7, 0x97, 0x15,
	0x69, 0x87, 0x19, 0xdc, 0xd6, 0xea, 0x9d, 0x37, 0x3a, 0x58, 0xee, 0xd3, 0x70, 0x6a, 0x3e, 0x9d,
	0xa5, 0xc2, 0xf9, 0xef, 0xa6, 0xaf, 0xc2, 0xae, 0x76, 0xb0, 0x95, 0x33, 0x3c, 0x1a, 0xa0, 0xc8,
	0x33, 0x29, 0xf3, 0x2c, 0x8f, 0x44, 0xd9, 0x9b, 0xa1, 0xbe, 0x2d, 0xb1, 0xcf, 0x8b, 0x3f, 0x69,
	0x6b, 0xca, 0x78, 0xa7, 0x97, 0xba, 0xed, 0xf6, 0x07, 0xbd, 0xd6, 0x95, 0xc0, 0x36, 0xe6, 0xa6,
	0x34, 0x53, 0x6b, 0xbf, 0x61, 0xe6, 0x81, 0xf9, 0x27, 0x25, 0xe9, 0xb5, 0x31, 0xef, 0x8d, 0x24,
	0xbd, 0xfd, 0x46, 0x4f, 0x49, 0x7e, 0xe8, 0xcb, 0xd2, 0x6f, 0xdb, 0x6d, 0xcc, 0x6d, 0x7b, 0x24,
	0xb2, 0xed, 0xfd, 0x86, 0x6d, 0x2b, 0xd9, 0xe1, 0x4c, 0x96, 0xe7, 0x83, 0x5f, 0x01, 0x00, 0x00,
	0xff, 0xff, 0x7a, 0x28, 0x07, 0xaf, 0x79, 0x07, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// AdGroupBidModifierServiceClient is the client API for AdGroupBidModifierService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type AdGroupBidModifierServiceClient interface {
	// Returns the requested ad group bid modifier in full detail.
	GetAdGroupBidModifier(ctx context.Context, in *GetAdGroupBidModifierRequest, opts ...grpc.CallOption) (*resources.AdGroupBidModifier, error)
	// Creates, updates, or removes ad group bid modifiers.
	// Operation statuses are returned.
	MutateAdGroupBidModifiers(ctx context.Context, in *MutateAdGroupBidModifiersRequest, opts ...grpc.CallOption) (*MutateAdGroupBidModifiersResponse, error)
}

type adGroupBidModifierServiceClient struct {
	cc *grpc.ClientConn
}

func NewAdGroupBidModifierServiceClient(cc *grpc.ClientConn) AdGroupBidModifierServiceClient {
	return &adGroupBidModifierServiceClient{cc}
}

func (c *adGroupBidModifierServiceClient) GetAdGroupBidModifier(ctx context.Context, in *GetAdGroupBidModifierRequest, opts ...grpc.CallOption) (*resources.AdGroupBidModifier, error) {
	out := new(resources.AdGroupBidModifier)
	err := c.cc.Invoke(ctx, "/google.ads.googleads.v9.services.AdGroupBidModifierService/GetAdGroupBidModifier", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *adGroupBidModifierServiceClient) MutateAdGroupBidModifiers(ctx context.Context, in *MutateAdGroupBidModifiersRequest, opts ...grpc.CallOption) (*MutateAdGroupBidModifiersResponse, error) {
	out := new(MutateAdGroupBidModifiersResponse)
	err := c.cc.Invoke(ctx, "/google.ads.googleads.v9.services.AdGroupBidModifierService/MutateAdGroupBidModifiers", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// AdGroupBidModifierServiceServer is the server API for AdGroupBidModifierService service.
type AdGroupBidModifierServiceServer interface {
	// Returns the requested ad group bid modifier in full detail.
	GetAdGroupBidModifier(context.Context, *GetAdGroupBidModifierRequest) (*resources.AdGroupBidModifier, error)
	// Creates, updates, or removes ad group bid modifiers.
	// Operation statuses are returned.
	MutateAdGroupBidModifiers(context.Context, *MutateAdGroupBidModifiersRequest) (*MutateAdGroupBidModifiersResponse, error)
}

func RegisterAdGroupBidModifierServiceServer(s *grpc.Server, srv AdGroupBidModifierServiceServer) {
	s.RegisterService(&_AdGroupBidModifierService_serviceDesc, srv)
}

func _AdGroupBidModifierService_GetAdGroupBidModifier_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetAdGroupBidModifierRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AdGroupBidModifierServiceServer).GetAdGroupBidModifier(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/google.ads.googleads.v9.services.AdGroupBidModifierService/GetAdGroupBidModifier",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AdGroupBidModifierServiceServer).GetAdGroupBidModifier(ctx, req.(*GetAdGroupBidModifierRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _AdGroupBidModifierService_MutateAdGroupBidModifiers_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MutateAdGroupBidModifiersRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AdGroupBidModifierServiceServer).MutateAdGroupBidModifiers(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/google.ads.googleads.v9.services.AdGroupBidModifierService/MutateAdGroupBidModifiers",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AdGroupBidModifierServiceServer).MutateAdGroupBidModifiers(ctx, req.(*MutateAdGroupBidModifiersRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _AdGroupBidModifierService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "google.ads.googleads.v9.services.AdGroupBidModifierService",
	HandlerType: (*AdGroupBidModifierServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetAdGroupBidModifier",
			Handler:    _AdGroupBidModifierService_GetAdGroupBidModifier_Handler,
		},
		{
			MethodName: "MutateAdGroupBidModifiers",
			Handler:    _AdGroupBidModifierService_MutateAdGroupBidModifiers_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "google/ads/googleads/v9/services/ad_group_bid_modifier_service.proto",
}
