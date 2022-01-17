// Code generated by protoc-gen-go. DO NOT EDIT.
// source: google/ads/googleads/v9/errors/campaign_shared_set_error.proto

package errors

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
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

// Enum describing possible campaign shared set errors.
type CampaignSharedSetErrorEnum_CampaignSharedSetError int32

const (
	// Enum unspecified.
	CampaignSharedSetErrorEnum_UNSPECIFIED CampaignSharedSetErrorEnum_CampaignSharedSetError = 0
	// The received error code is not known in this version.
	CampaignSharedSetErrorEnum_UNKNOWN CampaignSharedSetErrorEnum_CampaignSharedSetError = 1
	// The shared set belongs to another customer and permission isn't granted.
	CampaignSharedSetErrorEnum_SHARED_SET_ACCESS_DENIED CampaignSharedSetErrorEnum_CampaignSharedSetError = 2
)

var CampaignSharedSetErrorEnum_CampaignSharedSetError_name = map[int32]string{
	0: "UNSPECIFIED",
	1: "UNKNOWN",
	2: "SHARED_SET_ACCESS_DENIED",
}

var CampaignSharedSetErrorEnum_CampaignSharedSetError_value = map[string]int32{
	"UNSPECIFIED":              0,
	"UNKNOWN":                  1,
	"SHARED_SET_ACCESS_DENIED": 2,
}

func (x CampaignSharedSetErrorEnum_CampaignSharedSetError) String() string {
	return proto.EnumName(CampaignSharedSetErrorEnum_CampaignSharedSetError_name, int32(x))
}

func (CampaignSharedSetErrorEnum_CampaignSharedSetError) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_ac0c0aeef322f689, []int{0, 0}
}

// Container for enum describing possible campaign shared set errors.
type CampaignSharedSetErrorEnum struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CampaignSharedSetErrorEnum) Reset()         { *m = CampaignSharedSetErrorEnum{} }
func (m *CampaignSharedSetErrorEnum) String() string { return proto.CompactTextString(m) }
func (*CampaignSharedSetErrorEnum) ProtoMessage()    {}
func (*CampaignSharedSetErrorEnum) Descriptor() ([]byte, []int) {
	return fileDescriptor_ac0c0aeef322f689, []int{0}
}

func (m *CampaignSharedSetErrorEnum) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CampaignSharedSetErrorEnum.Unmarshal(m, b)
}
func (m *CampaignSharedSetErrorEnum) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CampaignSharedSetErrorEnum.Marshal(b, m, deterministic)
}
func (m *CampaignSharedSetErrorEnum) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CampaignSharedSetErrorEnum.Merge(m, src)
}
func (m *CampaignSharedSetErrorEnum) XXX_Size() int {
	return xxx_messageInfo_CampaignSharedSetErrorEnum.Size(m)
}
func (m *CampaignSharedSetErrorEnum) XXX_DiscardUnknown() {
	xxx_messageInfo_CampaignSharedSetErrorEnum.DiscardUnknown(m)
}

var xxx_messageInfo_CampaignSharedSetErrorEnum proto.InternalMessageInfo

func init() {
	proto.RegisterEnum("google.ads.googleads.v9.errors.CampaignSharedSetErrorEnum_CampaignSharedSetError", CampaignSharedSetErrorEnum_CampaignSharedSetError_name, CampaignSharedSetErrorEnum_CampaignSharedSetError_value)
	proto.RegisterType((*CampaignSharedSetErrorEnum)(nil), "google.ads.googleads.v9.errors.CampaignSharedSetErrorEnum")
}

func init() {
	proto.RegisterFile("google/ads/googleads/v9/errors/campaign_shared_set_error.proto", fileDescriptor_ac0c0aeef322f689)
}

var fileDescriptor_ac0c0aeef322f689 = []byte{
	// 299 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x7c, 0x90, 0xdf, 0x4a, 0xf3, 0x30,
	0x18, 0xc6, 0xbf, 0xf6, 0x03, 0x85, 0xec, 0xc0, 0xd1, 0x03, 0x11, 0x95, 0x1d, 0xf4, 0x02, 0xd2,
	0x82, 0x67, 0x11, 0x84, 0xac, 0x8d, 0x73, 0x08, 0xb5, 0x98, 0xad, 0x82, 0x14, 0x42, 0x5c, 0x42,
	0x1c, 0xac, 0x4d, 0x49, 0xea, 0x2e, 0xc8, 0x43, 0x2f, 0xc5, 0x4b, 0xf1, 0x02, 0x3c, 0x96, 0x36,
	0x5b, 0x8f, 0xa6, 0x47, 0x79, 0xc8, 0xfb, 0x7b, 0x9e, 0xf7, 0x0f, 0xb8, 0x51, 0x5a, 0xab, 0x8d,
	0x8c, 0xb8, 0xb0, 0x91, 0x93, 0x9d, 0xda, 0xc6, 0x91, 0x34, 0x46, 0x1b, 0x1b, 0xad, 0x78, 0xd5,
	0xf0, 0xb5, 0xaa, 0x99, 0x7d, 0xe5, 0x46, 0x0a, 0x66, 0x65, 0xcb, 0xfa, 0x12, 0x6c, 0x8c, 0x6e,
	0x75, 0x30, 0x71, 0x26, 0xc8, 0x85, 0x85, 0x83, 0x1f, 0x6e, 0x63, 0xe8, 0xfc, 0xa1, 0x01, 0xe7,
	0xc9, 0x2e, 0x82, 0xf6, 0x09, 0x54, 0xb6, 0xa4, 0x2b, 0x91, 0xfa, 0xad, 0x0a, 0x17, 0xe0, 0xf4,
	0x70, 0x35, 0x38, 0x01, 0xa3, 0x65, 0x46, 0x73, 0x92, 0xcc, 0x6f, 0xe7, 0x24, 0x1d, 0xff, 0x0b,
	0x46, 0xe0, 0x78, 0x99, 0xdd, 0x67, 0x0f, 0x4f, 0xd9, 0xd8, 0x0b, 0x2e, 0xc1, 0x19, 0xbd, 0xc3,
	0x8f, 0x24, 0x65, 0x94, 0x2c, 0x18, 0x4e, 0x12, 0x42, 0x29, 0x4b, 0x49, 0xd6, 0xa1, 0xfe, 0xf4,
	0xdb, 0x03, 0xe1, 0x4a, 0x57, 0xf0, 0xef, 0xd1, 0xa6, 0x17, 0x87, 0x5b, 0xe7, 0xdd, 0x5e, 0xb9,
	0xf7, 0x9c, 0xee, 0xec, 0x4a, 0x6f, 0x78, 0xad, 0xa0, 0x36, 0x2a, 0x52, 0xb2, 0xee, 0xb7, 0xde,
	0x5f, 0xaa, 0x59, 0xdb, 0xdf, 0x0e, 0x77, 0xed, 0x9e, 0x77, 0xff, 0xff, 0x0c, 0xe3, 0x0f, 0x7f,
	0x32, 0x73, 0x61, 0x58, 0x58, 0xe8, 0x64, 0xa7, 0x8a, 0x18, 0xf6, 0x2d, 0xed, 0xe7, 0x1e, 0x28,
	0xb1, 0xb0, 0xe5, 0x00, 0x94, 0x45, 0x5c, 0x3a, 0xe0, 0xcb, 0x0f, 0xdd, 0x2f, 0x42, 0x58, 0x58,
	0x84, 0x06, 0x04, 0xa1, 0x22, 0x46, 0xc8, 0x41, 0x2f, 0x47, 0xfd, 0x74, 0x57, 0x3f, 0x01, 0x00,
	0x00, 0xff, 0xff, 0x9e, 0x0f, 0x4e, 0xdc, 0xd5, 0x01, 0x00, 0x00,
}
