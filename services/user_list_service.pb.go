// Code generated by protoc-gen-go. DO NOT EDIT.
// source: google/ads/googleads/v0/services/user_list_service.proto

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

// Request message for [UserListService.GetUserList][google.ads.googleads.v0.services.UserListService.GetUserList].
type GetUserListRequest struct {
	// The resource name of the user list to fetch.
	ResourceName         string   `protobuf:"bytes,1,opt,name=resource_name,json=resourceName,proto3" json:"resource_name,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetUserListRequest) Reset()         { *m = GetUserListRequest{} }
func (m *GetUserListRequest) String() string { return proto.CompactTextString(m) }
func (*GetUserListRequest) ProtoMessage()    {}
func (*GetUserListRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_29727dde55d320ce, []int{0}
}

func (m *GetUserListRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetUserListRequest.Unmarshal(m, b)
}
func (m *GetUserListRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetUserListRequest.Marshal(b, m, deterministic)
}
func (m *GetUserListRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetUserListRequest.Merge(m, src)
}
func (m *GetUserListRequest) XXX_Size() int {
	return xxx_messageInfo_GetUserListRequest.Size(m)
}
func (m *GetUserListRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_GetUserListRequest.DiscardUnknown(m)
}

var xxx_messageInfo_GetUserListRequest proto.InternalMessageInfo

func (m *GetUserListRequest) GetResourceName() string {
	if m != nil {
		return m.ResourceName
	}
	return ""
}

// Request message for [UserListService.MutateUserLists][google.ads.googleads.v0.services.UserListService.MutateUserLists].
type MutateUserListsRequest struct {
	// The ID of the customer whose user lists are being modified.
	CustomerId string `protobuf:"bytes,1,opt,name=customer_id,json=customerId,proto3" json:"customer_id,omitempty"`
	// The list of operations to perform on individual user lists.
	Operations []*UserListOperation `protobuf:"bytes,2,rep,name=operations,proto3" json:"operations,omitempty"`
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

func (m *MutateUserListsRequest) Reset()         { *m = MutateUserListsRequest{} }
func (m *MutateUserListsRequest) String() string { return proto.CompactTextString(m) }
func (*MutateUserListsRequest) ProtoMessage()    {}
func (*MutateUserListsRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_29727dde55d320ce, []int{1}
}

func (m *MutateUserListsRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_MutateUserListsRequest.Unmarshal(m, b)
}
func (m *MutateUserListsRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_MutateUserListsRequest.Marshal(b, m, deterministic)
}
func (m *MutateUserListsRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MutateUserListsRequest.Merge(m, src)
}
func (m *MutateUserListsRequest) XXX_Size() int {
	return xxx_messageInfo_MutateUserListsRequest.Size(m)
}
func (m *MutateUserListsRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_MutateUserListsRequest.DiscardUnknown(m)
}

var xxx_messageInfo_MutateUserListsRequest proto.InternalMessageInfo

func (m *MutateUserListsRequest) GetCustomerId() string {
	if m != nil {
		return m.CustomerId
	}
	return ""
}

func (m *MutateUserListsRequest) GetOperations() []*UserListOperation {
	if m != nil {
		return m.Operations
	}
	return nil
}

func (m *MutateUserListsRequest) GetPartialFailure() bool {
	if m != nil {
		return m.PartialFailure
	}
	return false
}

func (m *MutateUserListsRequest) GetValidateOnly() bool {
	if m != nil {
		return m.ValidateOnly
	}
	return false
}

// A single operation (create, update) on a user list.
type UserListOperation struct {
	// FieldMask that determines which resource fields are modified in an update.
	UpdateMask *field_mask.FieldMask `protobuf:"bytes,4,opt,name=update_mask,json=updateMask,proto3" json:"update_mask,omitempty"`
	// The mutate operation.
	//
	// Types that are valid to be assigned to Operation:
	//	*UserListOperation_Create
	//	*UserListOperation_Update
	//	*UserListOperation_Remove
	Operation            isUserListOperation_Operation `protobuf_oneof:"operation"`
	XXX_NoUnkeyedLiteral struct{}                      `json:"-"`
	XXX_unrecognized     []byte                        `json:"-"`
	XXX_sizecache        int32                         `json:"-"`
}

func (m *UserListOperation) Reset()         { *m = UserListOperation{} }
func (m *UserListOperation) String() string { return proto.CompactTextString(m) }
func (*UserListOperation) ProtoMessage()    {}
func (*UserListOperation) Descriptor() ([]byte, []int) {
	return fileDescriptor_29727dde55d320ce, []int{2}
}

func (m *UserListOperation) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UserListOperation.Unmarshal(m, b)
}
func (m *UserListOperation) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UserListOperation.Marshal(b, m, deterministic)
}
func (m *UserListOperation) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UserListOperation.Merge(m, src)
}
func (m *UserListOperation) XXX_Size() int {
	return xxx_messageInfo_UserListOperation.Size(m)
}
func (m *UserListOperation) XXX_DiscardUnknown() {
	xxx_messageInfo_UserListOperation.DiscardUnknown(m)
}

var xxx_messageInfo_UserListOperation proto.InternalMessageInfo

func (m *UserListOperation) GetUpdateMask() *field_mask.FieldMask {
	if m != nil {
		return m.UpdateMask
	}
	return nil
}

type isUserListOperation_Operation interface {
	isUserListOperation_Operation()
}

type UserListOperation_Create struct {
	Create *resources.UserList `protobuf:"bytes,1,opt,name=create,proto3,oneof"`
}

type UserListOperation_Update struct {
	Update *resources.UserList `protobuf:"bytes,2,opt,name=update,proto3,oneof"`
}

type UserListOperation_Remove struct {
	Remove string `protobuf:"bytes,3,opt,name=remove,proto3,oneof"`
}

func (*UserListOperation_Create) isUserListOperation_Operation() {}

func (*UserListOperation_Update) isUserListOperation_Operation() {}

func (*UserListOperation_Remove) isUserListOperation_Operation() {}

func (m *UserListOperation) GetOperation() isUserListOperation_Operation {
	if m != nil {
		return m.Operation
	}
	return nil
}

func (m *UserListOperation) GetCreate() *resources.UserList {
	if x, ok := m.GetOperation().(*UserListOperation_Create); ok {
		return x.Create
	}
	return nil
}

func (m *UserListOperation) GetUpdate() *resources.UserList {
	if x, ok := m.GetOperation().(*UserListOperation_Update); ok {
		return x.Update
	}
	return nil
}

func (m *UserListOperation) GetRemove() string {
	if x, ok := m.GetOperation().(*UserListOperation_Remove); ok {
		return x.Remove
	}
	return ""
}

// XXX_OneofWrappers is for the internal use of the proto package.
func (*UserListOperation) XXX_OneofWrappers() []interface{} {
	return []interface{}{
		(*UserListOperation_Create)(nil),
		(*UserListOperation_Update)(nil),
		(*UserListOperation_Remove)(nil),
	}
}

// Response message for user list mutate.
type MutateUserListsResponse struct {
	// Errors that pertain to operation failures in the partial failure mode.
	// Returned only when partial_failure = true and all errors occur inside the
	// operations. If any errors occur outside the operations (e.g. auth errors),
	// we return an RPC level error.
	PartialFailureError *status.Status `protobuf:"bytes,3,opt,name=partial_failure_error,json=partialFailureError,proto3" json:"partial_failure_error,omitempty"`
	// All results for the mutate.
	Results              []*MutateUserListResult `protobuf:"bytes,2,rep,name=results,proto3" json:"results,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                `json:"-"`
	XXX_unrecognized     []byte                  `json:"-"`
	XXX_sizecache        int32                   `json:"-"`
}

func (m *MutateUserListsResponse) Reset()         { *m = MutateUserListsResponse{} }
func (m *MutateUserListsResponse) String() string { return proto.CompactTextString(m) }
func (*MutateUserListsResponse) ProtoMessage()    {}
func (*MutateUserListsResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_29727dde55d320ce, []int{3}
}

func (m *MutateUserListsResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_MutateUserListsResponse.Unmarshal(m, b)
}
func (m *MutateUserListsResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_MutateUserListsResponse.Marshal(b, m, deterministic)
}
func (m *MutateUserListsResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MutateUserListsResponse.Merge(m, src)
}
func (m *MutateUserListsResponse) XXX_Size() int {
	return xxx_messageInfo_MutateUserListsResponse.Size(m)
}
func (m *MutateUserListsResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_MutateUserListsResponse.DiscardUnknown(m)
}

var xxx_messageInfo_MutateUserListsResponse proto.InternalMessageInfo

func (m *MutateUserListsResponse) GetPartialFailureError() *status.Status {
	if m != nil {
		return m.PartialFailureError
	}
	return nil
}

func (m *MutateUserListsResponse) GetResults() []*MutateUserListResult {
	if m != nil {
		return m.Results
	}
	return nil
}

// The result for the user list mutate.
type MutateUserListResult struct {
	// Returned for successful operations.
	ResourceName         string   `protobuf:"bytes,1,opt,name=resource_name,json=resourceName,proto3" json:"resource_name,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *MutateUserListResult) Reset()         { *m = MutateUserListResult{} }
func (m *MutateUserListResult) String() string { return proto.CompactTextString(m) }
func (*MutateUserListResult) ProtoMessage()    {}
func (*MutateUserListResult) Descriptor() ([]byte, []int) {
	return fileDescriptor_29727dde55d320ce, []int{4}
}

func (m *MutateUserListResult) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_MutateUserListResult.Unmarshal(m, b)
}
func (m *MutateUserListResult) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_MutateUserListResult.Marshal(b, m, deterministic)
}
func (m *MutateUserListResult) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MutateUserListResult.Merge(m, src)
}
func (m *MutateUserListResult) XXX_Size() int {
	return xxx_messageInfo_MutateUserListResult.Size(m)
}
func (m *MutateUserListResult) XXX_DiscardUnknown() {
	xxx_messageInfo_MutateUserListResult.DiscardUnknown(m)
}

var xxx_messageInfo_MutateUserListResult proto.InternalMessageInfo

func (m *MutateUserListResult) GetResourceName() string {
	if m != nil {
		return m.ResourceName
	}
	return ""
}

func init() {
	proto.RegisterType((*GetUserListRequest)(nil), "google.ads.googleads.v0.services.GetUserListRequest")
	proto.RegisterType((*MutateUserListsRequest)(nil), "google.ads.googleads.v0.services.MutateUserListsRequest")
	proto.RegisterType((*UserListOperation)(nil), "google.ads.googleads.v0.services.UserListOperation")
	proto.RegisterType((*MutateUserListsResponse)(nil), "google.ads.googleads.v0.services.MutateUserListsResponse")
	proto.RegisterType((*MutateUserListResult)(nil), "google.ads.googleads.v0.services.MutateUserListResult")
}

func init() {
	proto.RegisterFile("google/ads/googleads/v0/services/user_list_service.proto", fileDescriptor_29727dde55d320ce)
}

var fileDescriptor_29727dde55d320ce = []byte{
	// 708 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x9c, 0x54, 0x4f, 0x6b, 0xd4, 0x40,
	0x14, 0x37, 0xa9, 0x54, 0x3b, 0xa9, 0x16, 0xc7, 0x6a, 0x97, 0x45, 0x74, 0x89, 0x05, 0xcb, 0x16,
	0x27, 0xeb, 0xae, 0x48, 0x9b, 0xd2, 0xc3, 0x16, 0xfa, 0x47, 0xb0, 0xb6, 0xa4, 0xd8, 0x83, 0x2c,
	0x84, 0xe9, 0x66, 0xba, 0x84, 0x26, 0x99, 0x38, 0x33, 0x59, 0x29, 0xa5, 0x17, 0xc1, 0x4f, 0xe0,
	0x37, 0x10, 0xbc, 0x78, 0xf5, 0x13, 0x78, 0xf5, 0xea, 0xd5, 0xa3, 0x27, 0xbf, 0x82, 0x08, 0x92,
	0x99, 0x4c, 0xda, 0xdd, 0x5a, 0xd6, 0xf6, 0xf6, 0xe6, 0xcd, 0xef, 0xf7, 0x7b, 0x6f, 0xde, 0x9f,
	0x01, 0x0b, 0x3d, 0x4a, 0x7b, 0x11, 0x71, 0x70, 0xc0, 0x1d, 0x65, 0xe6, 0x56, 0xbf, 0xe1, 0x70,
	0xc2, 0xfa, 0x61, 0x97, 0x70, 0x27, 0xe3, 0x84, 0xf9, 0x51, 0xc8, 0x85, 0x5f, 0xb8, 0x50, 0xca,
	0xa8, 0xa0, 0xb0, 0xa6, 0xe0, 0x08, 0x07, 0x1c, 0x95, 0x4c, 0xd4, 0x6f, 0x20, 0xcd, 0xac, 0x3e,
	0x39, 0x4f, 0x9b, 0x11, 0x4e, 0x33, 0x36, 0x20, 0xae, 0x44, 0xab, 0xf7, 0x34, 0x25, 0x0d, 0x1d,
	0x9c, 0x24, 0x54, 0x60, 0x11, 0xd2, 0x84, 0x17, 0xb7, 0x45, 0x48, 0x47, 0x9e, 0xf6, 0xb2, 0x7d,
	0x67, 0x3f, 0x24, 0x51, 0xe0, 0xc7, 0x98, 0x1f, 0x14, 0x88, 0xfb, 0xc3, 0x88, 0xb7, 0x0c, 0xa7,
	0x29, 0x61, 0x5a, 0x61, 0xa6, 0xb8, 0x67, 0x69, 0xd7, 0xe1, 0x02, 0x8b, 0xac, 0xb8, 0xb0, 0x17,
	0x01, 0x5c, 0x27, 0xe2, 0x15, 0x27, 0xec, 0x45, 0xc8, 0x85, 0x47, 0xde, 0x64, 0x84, 0x0b, 0xf8,
	0x10, 0xdc, 0xd0, 0xb9, 0xfa, 0x09, 0x8e, 0x49, 0xc5, 0xa8, 0x19, 0x73, 0x13, 0xde, 0xa4, 0x76,
	0xbe, 0xc4, 0x31, 0xb1, 0x7f, 0x18, 0xe0, 0xee, 0x66, 0x26, 0xb0, 0x20, 0x9a, 0xce, 0x35, 0xff,
	0x01, 0xb0, 0xba, 0x19, 0x17, 0x34, 0x26, 0xcc, 0x0f, 0x83, 0x82, 0x0d, 0xb4, 0xeb, 0x79, 0x00,
	0x77, 0x00, 0xa0, 0x29, 0x61, 0xea, 0x95, 0x15, 0xb3, 0x36, 0x36, 0x67, 0x35, 0x5b, 0x68, 0x54,
	0x65, 0x91, 0x0e, 0xb4, 0xa5, 0xb9, 0xde, 0x29, 0x19, 0xf8, 0x08, 0x4c, 0xa5, 0x98, 0x89, 0x10,
	0x47, 0xfe, 0x3e, 0x0e, 0xa3, 0x8c, 0x91, 0xca, 0x58, 0xcd, 0x98, 0xbb, 0xee, 0xdd, 0x2c, 0xdc,
	0x6b, 0xca, 0x9b, 0x3f, 0xaf, 0x8f, 0xa3, 0x30, 0xc0, 0x82, 0xf8, 0x34, 0x89, 0x0e, 0x2b, 0x57,
	0x25, 0x6c, 0x52, 0x3b, 0xb7, 0x92, 0xe8, 0xd0, 0x7e, 0x6f, 0x82, 0x5b, 0x67, 0xe2, 0xc1, 0x25,
	0x60, 0x65, 0xa9, 0x24, 0xe6, 0xd5, 0x97, 0x44, 0xab, 0x59, 0xd5, 0x99, 0xeb, 0xf2, 0xa3, 0xb5,
	0xbc, 0x41, 0x9b, 0x98, 0x1f, 0x78, 0x40, 0xc1, 0x73, 0x1b, 0xae, 0x82, 0xf1, 0x2e, 0x23, 0x58,
	0xa8, 0x7a, 0x5a, 0xcd, 0xf9, 0x73, 0x5f, 0x5c, 0x4e, 0x4a, 0xf9, 0xe4, 0x8d, 0x2b, 0x5e, 0x41,
	0xce, 0x65, 0x94, 0x68, 0xc5, 0xbc, 0x94, 0x8c, 0x22, 0xc3, 0x0a, 0x18, 0x67, 0x24, 0xa6, 0x7d,
	0x55, 0xa5, 0x89, 0xfc, 0x46, 0x9d, 0x57, 0x2c, 0x30, 0x51, 0x96, 0xd5, 0xfe, 0x62, 0x80, 0x99,
	0x33, 0x6d, 0xe6, 0x29, 0x4d, 0x38, 0x81, 0x6b, 0xe0, 0xce, 0x50, 0xc5, 0x7d, 0xc2, 0x18, 0x65,
	0x52, 0xd1, 0x6a, 0x42, 0x9d, 0x18, 0x4b, 0xbb, 0x68, 0x47, 0x8e, 0x9d, 0x77, 0x7b, 0xb0, 0x17,
	0xab, 0x39, 0x1c, 0x6e, 0x83, 0x6b, 0x8c, 0xf0, 0x2c, 0x12, 0x7a, 0x16, 0x9e, 0x8d, 0x9e, 0x85,
	0xc1, 0x9c, 0x3c, 0x49, 0xf7, 0xb4, 0x8c, 0xbd, 0x04, 0xa6, 0xff, 0x05, 0xf8, 0xaf, 0xc9, 0x6e,
	0xfe, 0x31, 0xc1, 0x94, 0xe6, 0xed, 0xa8, 0x78, 0xf0, 0x93, 0x01, 0xac, 0x53, 0x9b, 0x02, 0x9f,
	0x8e, 0xce, 0xf0, 0xec, 0x62, 0x55, 0x2f, 0xd2, 0x2a, 0xbb, 0xf5, 0xee, 0xfb, 0xcf, 0x0f, 0xe6,
	0x63, 0x38, 0x9f, 0xff, 0x1d, 0x47, 0x03, 0x69, 0x2f, 0xeb, 0x5d, 0xe2, 0x4e, 0x5d, 0x7e, 0x26,
	0xb2, 0x2f, 0x4e, 0xfd, 0x18, 0x7e, 0x35, 0xc0, 0xd4, 0x50, 0xbb, 0xe0, 0xc2, 0x45, 0xab, 0xa9,
	0x17, 0xb9, 0xba, 0x78, 0x09, 0xa6, 0x9a, 0x0d, 0x7b, 0x51, 0x66, 0xdf, 0xb2, 0x51, 0x9e, 0xfd,
	0x49, 0xba, 0x47, 0xa7, 0x3e, 0x86, 0xe5, 0xfa, 0xf1, 0x49, 0xf2, 0x6e, 0x2c, 0x85, 0x5c, 0xa3,
	0xbe, 0xf2, 0xdb, 0x00, 0xb3, 0x5d, 0x1a, 0x8f, 0x8c, 0xbd, 0x32, 0x3d, 0xd4, 0xa5, 0xed, 0x7c,
	0xff, 0xb6, 0x8d, 0xd7, 0x1b, 0x05, 0xb3, 0x47, 0x23, 0x9c, 0xf4, 0x10, 0x65, 0x3d, 0xa7, 0x47,
	0x12, 0xb9, 0x9d, 0xfa, 0x47, 0x4e, 0x43, 0x7e, 0xfe, 0xe7, 0xbf, 0xa4, 0x8d, 0x8f, 0xe6, 0xd8,
	0x7a, 0xbb, 0xfd, 0xd9, 0xac, 0xad, 0x2b, 0xc1, 0x76, 0xc0, 0x91, 0x32, 0x73, 0x6b, 0xb7, 0x81,
	0x8a, 0xc0, 0xfc, 0x9b, 0x86, 0x74, 0xda, 0x01, 0xef, 0x94, 0x90, 0xce, 0x6e, 0xa3, 0xa3, 0x21,
	0xbf, 0xcc, 0x59, 0xe5, 0x77, 0xdd, 0x76, 0xc0, 0x5d, 0xb7, 0x04, 0xb9, 0xee, 0x6e, 0xc3, 0x75,
	0x35, 0x6c, 0x6f, 0x5c, 0xe6, 0xd9, 0xfa, 0x1b, 0x00, 0x00, 0xff, 0xff, 0xd4, 0x79, 0xcc, 0x65,
	0xa3, 0x06, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// UserListServiceClient is the client API for UserListService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type UserListServiceClient interface {
	// Returns the requested user list.
	GetUserList(ctx context.Context, in *GetUserListRequest, opts ...grpc.CallOption) (*resources.UserList, error)
	// Creates or updates user lists. Operation statuses are returned.
	MutateUserLists(ctx context.Context, in *MutateUserListsRequest, opts ...grpc.CallOption) (*MutateUserListsResponse, error)
}

type userListServiceClient struct {
	cc *grpc.ClientConn
}

func NewUserListServiceClient(cc *grpc.ClientConn) UserListServiceClient {
	return &userListServiceClient{cc}
}

func (c *userListServiceClient) GetUserList(ctx context.Context, in *GetUserListRequest, opts ...grpc.CallOption) (*resources.UserList, error) {
	out := new(resources.UserList)
	err := c.cc.Invoke(ctx, "/google.ads.googleads.v0.services.UserListService/GetUserList", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userListServiceClient) MutateUserLists(ctx context.Context, in *MutateUserListsRequest, opts ...grpc.CallOption) (*MutateUserListsResponse, error) {
	out := new(MutateUserListsResponse)
	err := c.cc.Invoke(ctx, "/google.ads.googleads.v0.services.UserListService/MutateUserLists", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// UserListServiceServer is the server API for UserListService service.
type UserListServiceServer interface {
	// Returns the requested user list.
	GetUserList(context.Context, *GetUserListRequest) (*resources.UserList, error)
	// Creates or updates user lists. Operation statuses are returned.
	MutateUserLists(context.Context, *MutateUserListsRequest) (*MutateUserListsResponse, error)
}

func RegisterUserListServiceServer(s *grpc.Server, srv UserListServiceServer) {
	s.RegisterService(&_UserListService_serviceDesc, srv)
}

func _UserListService_GetUserList_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetUserListRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserListServiceServer).GetUserList(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/google.ads.googleads.v0.services.UserListService/GetUserList",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserListServiceServer).GetUserList(ctx, req.(*GetUserListRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserListService_MutateUserLists_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MutateUserListsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserListServiceServer).MutateUserLists(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/google.ads.googleads.v0.services.UserListService/MutateUserLists",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserListServiceServer).MutateUserLists(ctx, req.(*MutateUserListsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _UserListService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "google.ads.googleads.v0.services.UserListService",
	HandlerType: (*UserListServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetUserList",
			Handler:    _UserListService_GetUserList_Handler,
		},
		{
			MethodName: "MutateUserLists",
			Handler:    _UserListService_MutateUserLists_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "google/ads/googleads/v0/services/user_list_service.proto",
}
