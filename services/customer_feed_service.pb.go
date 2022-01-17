// Code generated by protoc-gen-go. DO NOT EDIT.
// source: google/ads/googleads/v9/services/customer_feed_service.proto

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

// Request message for [CustomerFeedService.GetCustomerFeed][google.ads.googleads.v9.services.CustomerFeedService.GetCustomerFeed].
type GetCustomerFeedRequest struct {
	// The resource name of the customer feed to fetch.
	ResourceName         string   `protobuf:"bytes,1,opt,name=resource_name,json=resourceName,proto3" json:"resource_name,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetCustomerFeedRequest) Reset()         { *m = GetCustomerFeedRequest{} }
func (m *GetCustomerFeedRequest) String() string { return proto.CompactTextString(m) }
func (*GetCustomerFeedRequest) ProtoMessage()    {}
func (*GetCustomerFeedRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_4fe93d5e76fdc55d, []int{0}
}

func (m *GetCustomerFeedRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetCustomerFeedRequest.Unmarshal(m, b)
}
func (m *GetCustomerFeedRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetCustomerFeedRequest.Marshal(b, m, deterministic)
}
func (m *GetCustomerFeedRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetCustomerFeedRequest.Merge(m, src)
}
func (m *GetCustomerFeedRequest) XXX_Size() int {
	return xxx_messageInfo_GetCustomerFeedRequest.Size(m)
}
func (m *GetCustomerFeedRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_GetCustomerFeedRequest.DiscardUnknown(m)
}

var xxx_messageInfo_GetCustomerFeedRequest proto.InternalMessageInfo

func (m *GetCustomerFeedRequest) GetResourceName() string {
	if m != nil {
		return m.ResourceName
	}
	return ""
}

// Request message for [CustomerFeedService.MutateCustomerFeeds][google.ads.googleads.v9.services.CustomerFeedService.MutateCustomerFeeds].
type MutateCustomerFeedsRequest struct {
	// The ID of the customer whose customer feeds are being modified.
	CustomerId string `protobuf:"bytes,1,opt,name=customer_id,json=customerId,proto3" json:"customer_id,omitempty"`
	// The list of operations to perform on individual customer feeds.
	Operations []*CustomerFeedOperation `protobuf:"bytes,2,rep,name=operations,proto3" json:"operations,omitempty"`
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

func (m *MutateCustomerFeedsRequest) Reset()         { *m = MutateCustomerFeedsRequest{} }
func (m *MutateCustomerFeedsRequest) String() string { return proto.CompactTextString(m) }
func (*MutateCustomerFeedsRequest) ProtoMessage()    {}
func (*MutateCustomerFeedsRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_4fe93d5e76fdc55d, []int{1}
}

func (m *MutateCustomerFeedsRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_MutateCustomerFeedsRequest.Unmarshal(m, b)
}
func (m *MutateCustomerFeedsRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_MutateCustomerFeedsRequest.Marshal(b, m, deterministic)
}
func (m *MutateCustomerFeedsRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MutateCustomerFeedsRequest.Merge(m, src)
}
func (m *MutateCustomerFeedsRequest) XXX_Size() int {
	return xxx_messageInfo_MutateCustomerFeedsRequest.Size(m)
}
func (m *MutateCustomerFeedsRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_MutateCustomerFeedsRequest.DiscardUnknown(m)
}

var xxx_messageInfo_MutateCustomerFeedsRequest proto.InternalMessageInfo

func (m *MutateCustomerFeedsRequest) GetCustomerId() string {
	if m != nil {
		return m.CustomerId
	}
	return ""
}

func (m *MutateCustomerFeedsRequest) GetOperations() []*CustomerFeedOperation {
	if m != nil {
		return m.Operations
	}
	return nil
}

func (m *MutateCustomerFeedsRequest) GetPartialFailure() bool {
	if m != nil {
		return m.PartialFailure
	}
	return false
}

func (m *MutateCustomerFeedsRequest) GetValidateOnly() bool {
	if m != nil {
		return m.ValidateOnly
	}
	return false
}

// A single operation (create, update, remove) on a customer feed.
type CustomerFeedOperation struct {
	// FieldMask that determines which resource fields are modified in an update.
	UpdateMask *field_mask.FieldMask `protobuf:"bytes,4,opt,name=update_mask,json=updateMask,proto3" json:"update_mask,omitempty"`
	// The mutate operation.
	//
	// Types that are valid to be assigned to Operation:
	//	*CustomerFeedOperation_Create
	//	*CustomerFeedOperation_Update
	//	*CustomerFeedOperation_Remove
	Operation            isCustomerFeedOperation_Operation `protobuf_oneof:"operation"`
	XXX_NoUnkeyedLiteral struct{}                          `json:"-"`
	XXX_unrecognized     []byte                            `json:"-"`
	XXX_sizecache        int32                             `json:"-"`
}

func (m *CustomerFeedOperation) Reset()         { *m = CustomerFeedOperation{} }
func (m *CustomerFeedOperation) String() string { return proto.CompactTextString(m) }
func (*CustomerFeedOperation) ProtoMessage()    {}
func (*CustomerFeedOperation) Descriptor() ([]byte, []int) {
	return fileDescriptor_4fe93d5e76fdc55d, []int{2}
}

func (m *CustomerFeedOperation) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CustomerFeedOperation.Unmarshal(m, b)
}
func (m *CustomerFeedOperation) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CustomerFeedOperation.Marshal(b, m, deterministic)
}
func (m *CustomerFeedOperation) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CustomerFeedOperation.Merge(m, src)
}
func (m *CustomerFeedOperation) XXX_Size() int {
	return xxx_messageInfo_CustomerFeedOperation.Size(m)
}
func (m *CustomerFeedOperation) XXX_DiscardUnknown() {
	xxx_messageInfo_CustomerFeedOperation.DiscardUnknown(m)
}

var xxx_messageInfo_CustomerFeedOperation proto.InternalMessageInfo

func (m *CustomerFeedOperation) GetUpdateMask() *field_mask.FieldMask {
	if m != nil {
		return m.UpdateMask
	}
	return nil
}

type isCustomerFeedOperation_Operation interface {
	isCustomerFeedOperation_Operation()
}

type CustomerFeedOperation_Create struct {
	Create *resources.CustomerFeed `protobuf:"bytes,1,opt,name=create,proto3,oneof"`
}

type CustomerFeedOperation_Update struct {
	Update *resources.CustomerFeed `protobuf:"bytes,2,opt,name=update,proto3,oneof"`
}

type CustomerFeedOperation_Remove struct {
	Remove string `protobuf:"bytes,3,opt,name=remove,proto3,oneof"`
}

func (*CustomerFeedOperation_Create) isCustomerFeedOperation_Operation() {}

func (*CustomerFeedOperation_Update) isCustomerFeedOperation_Operation() {}

func (*CustomerFeedOperation_Remove) isCustomerFeedOperation_Operation() {}

func (m *CustomerFeedOperation) GetOperation() isCustomerFeedOperation_Operation {
	if m != nil {
		return m.Operation
	}
	return nil
}

func (m *CustomerFeedOperation) GetCreate() *resources.CustomerFeed {
	if x, ok := m.GetOperation().(*CustomerFeedOperation_Create); ok {
		return x.Create
	}
	return nil
}

func (m *CustomerFeedOperation) GetUpdate() *resources.CustomerFeed {
	if x, ok := m.GetOperation().(*CustomerFeedOperation_Update); ok {
		return x.Update
	}
	return nil
}

func (m *CustomerFeedOperation) GetRemove() string {
	if x, ok := m.GetOperation().(*CustomerFeedOperation_Remove); ok {
		return x.Remove
	}
	return ""
}

// XXX_OneofWrappers is for the internal use of the proto package.
func (*CustomerFeedOperation) XXX_OneofWrappers() []interface{} {
	return []interface{}{
		(*CustomerFeedOperation_Create)(nil),
		(*CustomerFeedOperation_Update)(nil),
		(*CustomerFeedOperation_Remove)(nil),
	}
}

// Response message for a customer feed mutate.
type MutateCustomerFeedsResponse struct {
	// Errors that pertain to operation failures in the partial failure mode.
	// Returned only when partial_failure = true and all errors occur inside the
	// operations. If any errors occur outside the operations (e.g. auth errors),
	// we return an RPC level error.
	PartialFailureError *status.Status `protobuf:"bytes,3,opt,name=partial_failure_error,json=partialFailureError,proto3" json:"partial_failure_error,omitempty"`
	// All results for the mutate.
	Results              []*MutateCustomerFeedResult `protobuf:"bytes,2,rep,name=results,proto3" json:"results,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                    `json:"-"`
	XXX_unrecognized     []byte                      `json:"-"`
	XXX_sizecache        int32                       `json:"-"`
}

func (m *MutateCustomerFeedsResponse) Reset()         { *m = MutateCustomerFeedsResponse{} }
func (m *MutateCustomerFeedsResponse) String() string { return proto.CompactTextString(m) }
func (*MutateCustomerFeedsResponse) ProtoMessage()    {}
func (*MutateCustomerFeedsResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_4fe93d5e76fdc55d, []int{3}
}

func (m *MutateCustomerFeedsResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_MutateCustomerFeedsResponse.Unmarshal(m, b)
}
func (m *MutateCustomerFeedsResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_MutateCustomerFeedsResponse.Marshal(b, m, deterministic)
}
func (m *MutateCustomerFeedsResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MutateCustomerFeedsResponse.Merge(m, src)
}
func (m *MutateCustomerFeedsResponse) XXX_Size() int {
	return xxx_messageInfo_MutateCustomerFeedsResponse.Size(m)
}
func (m *MutateCustomerFeedsResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_MutateCustomerFeedsResponse.DiscardUnknown(m)
}

var xxx_messageInfo_MutateCustomerFeedsResponse proto.InternalMessageInfo

func (m *MutateCustomerFeedsResponse) GetPartialFailureError() *status.Status {
	if m != nil {
		return m.PartialFailureError
	}
	return nil
}

func (m *MutateCustomerFeedsResponse) GetResults() []*MutateCustomerFeedResult {
	if m != nil {
		return m.Results
	}
	return nil
}

// The result for the customer feed mutate.
type MutateCustomerFeedResult struct {
	// Returned for successful operations.
	ResourceName         string   `protobuf:"bytes,1,opt,name=resource_name,json=resourceName,proto3" json:"resource_name,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *MutateCustomerFeedResult) Reset()         { *m = MutateCustomerFeedResult{} }
func (m *MutateCustomerFeedResult) String() string { return proto.CompactTextString(m) }
func (*MutateCustomerFeedResult) ProtoMessage()    {}
func (*MutateCustomerFeedResult) Descriptor() ([]byte, []int) {
	return fileDescriptor_4fe93d5e76fdc55d, []int{4}
}

func (m *MutateCustomerFeedResult) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_MutateCustomerFeedResult.Unmarshal(m, b)
}
func (m *MutateCustomerFeedResult) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_MutateCustomerFeedResult.Marshal(b, m, deterministic)
}
func (m *MutateCustomerFeedResult) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MutateCustomerFeedResult.Merge(m, src)
}
func (m *MutateCustomerFeedResult) XXX_Size() int {
	return xxx_messageInfo_MutateCustomerFeedResult.Size(m)
}
func (m *MutateCustomerFeedResult) XXX_DiscardUnknown() {
	xxx_messageInfo_MutateCustomerFeedResult.DiscardUnknown(m)
}

var xxx_messageInfo_MutateCustomerFeedResult proto.InternalMessageInfo

func (m *MutateCustomerFeedResult) GetResourceName() string {
	if m != nil {
		return m.ResourceName
	}
	return ""
}

func init() {
	proto.RegisterType((*GetCustomerFeedRequest)(nil), "google.ads.googleads.v9.services.GetCustomerFeedRequest")
	proto.RegisterType((*MutateCustomerFeedsRequest)(nil), "google.ads.googleads.v9.services.MutateCustomerFeedsRequest")
	proto.RegisterType((*CustomerFeedOperation)(nil), "google.ads.googleads.v9.services.CustomerFeedOperation")
	proto.RegisterType((*MutateCustomerFeedsResponse)(nil), "google.ads.googleads.v9.services.MutateCustomerFeedsResponse")
	proto.RegisterType((*MutateCustomerFeedResult)(nil), "google.ads.googleads.v9.services.MutateCustomerFeedResult")
}

func init() {
	proto.RegisterFile("google/ads/googleads/v9/services/customer_feed_service.proto", fileDescriptor_4fe93d5e76fdc55d)
}

var fileDescriptor_4fe93d5e76fdc55d = []byte{
	// 711 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x9c, 0x55, 0xc1, 0x6e, 0xd3, 0x4c,
	0x10, 0xfe, 0xed, 0xfc, 0xea, 0xff, 0x77, 0x5d, 0xa8, 0xb4, 0x55, 0xc1, 0x0a, 0x08, 0x22, 0x53,
	0x89, 0x2a, 0x07, 0x6f, 0x08, 0x42, 0x45, 0x6e, 0x23, 0x94, 0x22, 0xd2, 0xf6, 0x50, 0x5a, 0xb9,
	0xa8, 0x48, 0x28, 0x92, 0xb5, 0x8d, 0x37, 0x91, 0x55, 0xdb, 0x6b, 0x76, 0xd7, 0x41, 0x55, 0xd5,
	0x0b, 0x6f, 0x80, 0x78, 0x02, 0x38, 0x72, 0xe7, 0xc4, 0x1b, 0x70, 0x43, 0xbc, 0x42, 0x4f, 0xbc,
	0x04, 0xc8, 0x5e, 0xaf, 0x49, 0xda, 0x44, 0x81, 0xde, 0xc6, 0xb3, 0xdf, 0x7c, 0x33, 0xb3, 0xdf,
	0xec, 0x18, 0x6c, 0x0c, 0x28, 0x1d, 0x84, 0x04, 0x61, 0x9f, 0x23, 0x69, 0x66, 0xd6, 0xb0, 0x81,
	0x38, 0x61, 0xc3, 0xa0, 0x47, 0x38, 0xea, 0xa5, 0x5c, 0xd0, 0x88, 0x30, 0xaf, 0x4f, 0x88, 0xef,
	0x15, 0x6e, 0x3b, 0x61, 0x54, 0x50, 0x58, 0x93, 0x21, 0x36, 0xf6, 0xb9, 0x5d, 0x46, 0xdb, 0xc3,
	0x86, 0xad, 0xa2, 0xab, 0x8f, 0xa6, 0xf1, 0x33, 0xc2, 0x69, 0xca, 0x2e, 0x25, 0x90, 0xc4, 0xd5,
	0xdb, 0x2a, 0x2c, 0x09, 0x10, 0x8e, 0x63, 0x2a, 0xb0, 0x08, 0x68, 0xcc, 0x8b, 0xd3, 0x22, 0x2d,
	0xca, 0xbf, 0x8e, 0xd2, 0x3e, 0xea, 0x07, 0x24, 0xf4, 0xbd, 0x08, 0xf3, 0xe3, 0x02, 0x71, 0xe7,
	0x22, 0xe2, 0x0d, 0xc3, 0x49, 0x42, 0x98, 0x62, 0xb8, 0x59, 0x9c, 0xb3, 0xa4, 0x87, 0xb8, 0xc0,
	0x22, 0x2d, 0x0e, 0xac, 0x16, 0xb8, 0xb1, 0x45, 0xc4, 0xd3, 0xa2, 0xa4, 0x0e, 0x21, 0xbe, 0x4b,
	0x5e, 0xa7, 0x84, 0x0b, 0x78, 0x0f, 0x5c, 0x53, 0x35, 0x7b, 0x31, 0x8e, 0x88, 0xa9, 0xd5, 0xb4,
	0xd5, 0x79, 0x77, 0x41, 0x39, 0x9f, 0xe3, 0x88, 0x58, 0xe7, 0x1a, 0xa8, 0xee, 0xa6, 0x02, 0x0b,
	0x32, 0x4a, 0xc1, 0x15, 0xc7, 0x5d, 0x60, 0x94, 0xdd, 0x06, 0x7e, 0xc1, 0x00, 0x94, 0x6b, 0xc7,
	0x87, 0x2f, 0x01, 0xa0, 0x09, 0x61, 0xb2, 0x5b, 0x53, 0xaf, 0x55, 0x56, 0x8d, 0xe6, 0x9a, 0x3d,
	0xeb, 0x96, 0xed, 0xd1, 0x64, 0x7b, 0x2a, 0xde, 0x1d, 0xa1, 0x82, 0xf7, 0xc1, 0x62, 0x82, 0x99,
	0x08, 0x70, 0xe8, 0xf5, 0x71, 0x10, 0xa6, 0x8c, 0x98, 0x95, 0x9a, 0xb6, 0xfa, 0xbf, 0x7b, 0xbd,
	0x70, 0x77, 0xa4, 0x37, 0x6b, 0x73, 0x88, 0xc3, 0xc0, 0xc7, 0x82, 0x78, 0x34, 0x0e, 0x4f, 0xcc,
	0x7f, 0x73, 0xd8, 0x82, 0x72, 0xee, 0xc5, 0xe1, 0x89, 0xf5, 0x4e, 0x07, 0xcb, 0x13, 0x73, 0xc2,
	0x75, 0x60, 0xa4, 0x49, 0x1e, 0x9c, 0xa9, 0x91, 0x07, 0x1b, 0xcd, 0xaa, 0xea, 0x40, 0xc9, 0x61,
	0x77, 0x32, 0xc1, 0x76, 0x31, 0x3f, 0x76, 0x81, 0x84, 0x67, 0x36, 0xdc, 0x01, 0x73, 0x3d, 0x46,
	0xb0, 0x90, 0x77, 0x6b, 0x34, 0xd1, 0xd4, 0xce, 0xcb, 0xe9, 0x19, 0x6b, 0x7d, 0xfb, 0x1f, 0xb7,
	0x20, 0xc8, 0xa8, 0x24, 0xb1, 0xa9, 0x5f, 0x99, 0x4a, 0x12, 0x40, 0x13, 0xcc, 0x31, 0x12, 0xd1,
	0xa1, 0xbc, 0xb1, 0xf9, 0xec, 0x44, 0x7e, 0x6f, 0x1a, 0x60, 0xbe, 0xbc, 0x62, 0xeb, 0x8b, 0x06,
	0x6e, 0x4d, 0x94, 0x9e, 0x27, 0x34, 0xe6, 0x04, 0x76, 0xc0, 0xf2, 0x05, 0x05, 0x3c, 0xc2, 0x18,
	0x65, 0x39, 0xab, 0xd1, 0x84, 0xaa, 0x40, 0x96, 0xf4, 0xec, 0x83, 0x7c, 0x24, 0xdd, 0xa5, 0x71,
	0x6d, 0x9e, 0x65, 0x70, 0xf8, 0x02, 0xfc, 0xc7, 0x08, 0x4f, 0x43, 0xa1, 0xe6, 0xc3, 0x99, 0x3d,
	0x1f, 0x97, 0xeb, 0x72, 0x73, 0x0a, 0x57, 0x51, 0x59, 0x4f, 0x80, 0x39, 0x0d, 0xf4, 0x47, 0x93,
	0xdf, 0xfc, 0x50, 0x01, 0x4b, 0xa3, 0xb1, 0x07, 0x32, 0x37, 0xfc, 0xac, 0x81, 0xc5, 0x0b, 0x2f,
	0x0a, 0x3e, 0x9e, 0x5d, 0xf1, 0xe4, 0x47, 0x58, 0xfd, 0x5b, 0x19, 0xad, 0xb5, 0xb7, 0xdf, 0xcf,
	0xdf, 0xeb, 0x0f, 0x20, 0xca, 0x76, 0xce, 0xe9, 0x58, 0x1b, 0x2d, 0xf5, 0xee, 0x38, 0xaa, 0x97,
	0x4b, 0x28, 0xd7, 0x0c, 0xd5, 0xcf, 0xe0, 0x37, 0x0d, 0x2c, 0x4d, 0x90, 0x13, 0x6e, 0x5c, 0xe5,
	0xb6, 0xd5, 0x02, 0xa8, 0xb6, 0xae, 0x18, 0x2d, 0x67, 0xc8, 0x6a, 0xe5, 0xdd, 0xac, 0x59, 0xcd,
	0xac, 0x9b, 0xdf, 0xe5, 0x9f, 0x8e, 0x2c, 0x95, 0x56, 0xfd, 0x6c, 0xbc, 0x19, 0x27, 0xca, 0x09,
	0x1d, 0xad, 0xbe, 0xf9, 0x53, 0x03, 0x2b, 0x3d, 0x1a, 0xcd, 0xac, 0x61, 0xd3, 0x9c, 0xa0, 0xe4,
	0x7e, 0xf6, 0x76, 0xf7, 0xb5, 0x57, 0xdb, 0x45, 0xf4, 0x80, 0x86, 0x38, 0x1e, 0xd8, 0x94, 0x0d,
	0xd0, 0x80, 0xc4, 0xf9, 0xcb, 0x56, 0x1b, 0x3e, 0x09, 0xf8, 0xf4, 0x1f, 0xca, 0xba, 0x32, 0x3e,
	0xea, 0x95, 0xad, 0x76, 0xfb, 0x93, 0x5e, 0xdb, 0x92, 0x84, 0x6d, 0x9f, 0xdb, 0xd2, 0xcc, 0xac,
	0xc3, 0x86, 0x5d, 0x24, 0xe6, 0x5f, 0x15, 0xa4, 0xdb, 0xf6, 0x79, 0xb7, 0x84, 0x74, 0x0f, 0x1b,
	0x5d, 0x05, 0xf9, 0xa1, 0xaf, 0x48, 0xbf, 0xe3, 0xb4, 0x7d, 0xee, 0x38, 0x25, 0xc8, 0x71, 0x0e,
	0x1b, 0x8e, 0xa3, 0x60, 0x47, 0x73, 0x79, 0x9d, 0x0f, 0x7f, 0x05, 0x00, 0x00, 0xff, 0xff, 0x00,
	0x9c, 0x26, 0xb8, 0xf7, 0x06, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// CustomerFeedServiceClient is the client API for CustomerFeedService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type CustomerFeedServiceClient interface {
	// Returns the requested customer feed in full detail.
	GetCustomerFeed(ctx context.Context, in *GetCustomerFeedRequest, opts ...grpc.CallOption) (*resources.CustomerFeed, error)
	// Creates, updates, or removes customer feeds. Operation statuses are
	// returned.
	MutateCustomerFeeds(ctx context.Context, in *MutateCustomerFeedsRequest, opts ...grpc.CallOption) (*MutateCustomerFeedsResponse, error)
}

type customerFeedServiceClient struct {
	cc *grpc.ClientConn
}

func NewCustomerFeedServiceClient(cc *grpc.ClientConn) CustomerFeedServiceClient {
	return &customerFeedServiceClient{cc}
}

func (c *customerFeedServiceClient) GetCustomerFeed(ctx context.Context, in *GetCustomerFeedRequest, opts ...grpc.CallOption) (*resources.CustomerFeed, error) {
	out := new(resources.CustomerFeed)
	err := c.cc.Invoke(ctx, "/google.ads.googleads.v9.services.CustomerFeedService/GetCustomerFeed", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *customerFeedServiceClient) MutateCustomerFeeds(ctx context.Context, in *MutateCustomerFeedsRequest, opts ...grpc.CallOption) (*MutateCustomerFeedsResponse, error) {
	out := new(MutateCustomerFeedsResponse)
	err := c.cc.Invoke(ctx, "/google.ads.googleads.v9.services.CustomerFeedService/MutateCustomerFeeds", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// CustomerFeedServiceServer is the server API for CustomerFeedService service.
type CustomerFeedServiceServer interface {
	// Returns the requested customer feed in full detail.
	GetCustomerFeed(context.Context, *GetCustomerFeedRequest) (*resources.CustomerFeed, error)
	// Creates, updates, or removes customer feeds. Operation statuses are
	// returned.
	MutateCustomerFeeds(context.Context, *MutateCustomerFeedsRequest) (*MutateCustomerFeedsResponse, error)
}

func RegisterCustomerFeedServiceServer(s *grpc.Server, srv CustomerFeedServiceServer) {
	s.RegisterService(&_CustomerFeedService_serviceDesc, srv)
}

func _CustomerFeedService_GetCustomerFeed_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetCustomerFeedRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CustomerFeedServiceServer).GetCustomerFeed(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/google.ads.googleads.v9.services.CustomerFeedService/GetCustomerFeed",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CustomerFeedServiceServer).GetCustomerFeed(ctx, req.(*GetCustomerFeedRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _CustomerFeedService_MutateCustomerFeeds_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MutateCustomerFeedsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CustomerFeedServiceServer).MutateCustomerFeeds(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/google.ads.googleads.v9.services.CustomerFeedService/MutateCustomerFeeds",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CustomerFeedServiceServer).MutateCustomerFeeds(ctx, req.(*MutateCustomerFeedsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _CustomerFeedService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "google.ads.googleads.v9.services.CustomerFeedService",
	HandlerType: (*CustomerFeedServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetCustomerFeed",
			Handler:    _CustomerFeedService_GetCustomerFeed_Handler,
		},
		{
			MethodName: "MutateCustomerFeeds",
			Handler:    _CustomerFeedService_MutateCustomerFeeds_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "google/ads/googleads/v9/services/customer_feed_service.proto",
}
