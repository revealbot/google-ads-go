// Code generated by protoc-gen-go. DO NOT EDIT.
// source: google/ads/googleads/v9/errors/ad_sharing_error.proto

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

// Enum describing possible ad sharing errors.
type AdSharingErrorEnum_AdSharingError int32

const (
	// Enum unspecified.
	AdSharingErrorEnum_UNSPECIFIED AdSharingErrorEnum_AdSharingError = 0
	// The received error code is not known in this version.
	AdSharingErrorEnum_UNKNOWN AdSharingErrorEnum_AdSharingError = 1
	// Error resulting in attempting to add an Ad to an AdGroup that already
	// contains the Ad.
	AdSharingErrorEnum_AD_GROUP_ALREADY_CONTAINS_AD AdSharingErrorEnum_AdSharingError = 2
	// Ad is not compatible with the AdGroup it is being shared with.
	AdSharingErrorEnum_INCOMPATIBLE_AD_UNDER_AD_GROUP AdSharingErrorEnum_AdSharingError = 3
	// Cannot add AdGroupAd on inactive Ad.
	AdSharingErrorEnum_CANNOT_SHARE_INACTIVE_AD AdSharingErrorEnum_AdSharingError = 4
)

var AdSharingErrorEnum_AdSharingError_name = map[int32]string{
	0: "UNSPECIFIED",
	1: "UNKNOWN",
	2: "AD_GROUP_ALREADY_CONTAINS_AD",
	3: "INCOMPATIBLE_AD_UNDER_AD_GROUP",
	4: "CANNOT_SHARE_INACTIVE_AD",
}

var AdSharingErrorEnum_AdSharingError_value = map[string]int32{
	"UNSPECIFIED":                    0,
	"UNKNOWN":                        1,
	"AD_GROUP_ALREADY_CONTAINS_AD":   2,
	"INCOMPATIBLE_AD_UNDER_AD_GROUP": 3,
	"CANNOT_SHARE_INACTIVE_AD":       4,
}

func (x AdSharingErrorEnum_AdSharingError) String() string {
	return proto.EnumName(AdSharingErrorEnum_AdSharingError_name, int32(x))
}

func (AdSharingErrorEnum_AdSharingError) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_59119a77dc307d0a, []int{0, 0}
}

// Container for enum describing possible ad sharing errors.
type AdSharingErrorEnum struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *AdSharingErrorEnum) Reset()         { *m = AdSharingErrorEnum{} }
func (m *AdSharingErrorEnum) String() string { return proto.CompactTextString(m) }
func (*AdSharingErrorEnum) ProtoMessage()    {}
func (*AdSharingErrorEnum) Descriptor() ([]byte, []int) {
	return fileDescriptor_59119a77dc307d0a, []int{0}
}

func (m *AdSharingErrorEnum) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_AdSharingErrorEnum.Unmarshal(m, b)
}
func (m *AdSharingErrorEnum) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_AdSharingErrorEnum.Marshal(b, m, deterministic)
}
func (m *AdSharingErrorEnum) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AdSharingErrorEnum.Merge(m, src)
}
func (m *AdSharingErrorEnum) XXX_Size() int {
	return xxx_messageInfo_AdSharingErrorEnum.Size(m)
}
func (m *AdSharingErrorEnum) XXX_DiscardUnknown() {
	xxx_messageInfo_AdSharingErrorEnum.DiscardUnknown(m)
}

var xxx_messageInfo_AdSharingErrorEnum proto.InternalMessageInfo

func init() {
	proto.RegisterEnum("google.ads.googleads.v9.errors.AdSharingErrorEnum_AdSharingError", AdSharingErrorEnum_AdSharingError_name, AdSharingErrorEnum_AdSharingError_value)
	proto.RegisterType((*AdSharingErrorEnum)(nil), "google.ads.googleads.v9.errors.AdSharingErrorEnum")
}

func init() {
	proto.RegisterFile("google/ads/googleads/v9/errors/ad_sharing_error.proto", fileDescriptor_59119a77dc307d0a)
}

var fileDescriptor_59119a77dc307d0a = []byte{
	// 342 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x7c, 0x90, 0xcd, 0x4a, 0xeb, 0x40,
	0x1c, 0xc5, 0x6f, 0xd2, 0xcb, 0xbd, 0x30, 0x85, 0x7b, 0x43, 0xdc, 0xb8, 0x28, 0x45, 0xf2, 0x00,
	0x93, 0x80, 0xb8, 0x19, 0x57, 0xff, 0x26, 0x63, 0x0d, 0xd6, 0x49, 0xc8, 0x97, 0x28, 0x81, 0x21,
	0x9a, 0x32, 0x16, 0xda, 0x4c, 0xc9, 0x68, 0x1f, 0xc4, 0x47, 0x70, 0xa7, 0x8f, 0xe2, 0xa3, 0xb8,
	0xf0, 0x19, 0x24, 0x19, 0x5b, 0xe8, 0x42, 0x57, 0x73, 0x38, 0xfc, 0xce, 0xcc, 0x99, 0x83, 0x4e,
	0x84, 0x94, 0x62, 0x39, 0x77, 0xab, 0x5a, 0xb9, 0x5a, 0x76, 0x6a, 0xe3, 0xb9, 0xf3, 0xb6, 0x95,
	0xad, 0x72, 0xab, 0x9a, 0xab, 0xfb, 0xaa, 0x5d, 0x34, 0x82, 0xf7, 0x0e, 0x5e, 0xb7, 0xf2, 0x41,
	0xda, 0x63, 0xcd, 0xe2, 0xaa, 0x56, 0x78, 0x17, 0xc3, 0x1b, 0x0f, 0xeb, 0x98, 0xf3, 0x62, 0x20,
	0x1b, 0xea, 0x54, 0x27, 0x69, 0xe7, 0xd1, 0xe6, 0x71, 0xe5, 0x3c, 0x19, 0xe8, 0xdf, 0xbe, 0x6d,
	0xff, 0x47, 0xc3, 0x9c, 0xa5, 0x31, 0xf5, 0xc3, 0xb3, 0x90, 0x06, 0xd6, 0x2f, 0x7b, 0x88, 0xfe,
	0xe6, 0xec, 0x82, 0x45, 0x57, 0xcc, 0x32, 0xec, 0x23, 0x34, 0x82, 0x80, 0x4f, 0x93, 0x28, 0x8f,
	0x39, 0xcc, 0x12, 0x0a, 0xc1, 0x35, 0xf7, 0x23, 0x96, 0x41, 0xc8, 0x52, 0x0e, 0x81, 0x65, 0xda,
	0x0e, 0x1a, 0x87, 0xcc, 0x8f, 0x2e, 0x63, 0xc8, 0xc2, 0xc9, 0x8c, 0x72, 0x08, 0x78, 0xce, 0x02,
	0x9a, 0xf0, 0x6d, 0xce, 0x1a, 0xd8, 0x23, 0x74, 0xe8, 0x03, 0x63, 0x51, 0xc6, 0xd3, 0x73, 0x48,
	0x28, 0x0f, 0x19, 0xf8, 0x59, 0x58, 0x74, 0xb0, 0xf5, 0x7b, 0xf2, 0x61, 0x20, 0xe7, 0x4e, 0xae,
	0xf0, 0xcf, 0x5f, 0x9a, 0x1c, 0xec, 0x17, 0x8f, 0xbb, 0x1d, 0x62, 0xe3, 0x26, 0xf8, 0x8a, 0x09,
	0xb9, 0xac, 0x1a, 0x81, 0x65, 0x2b, 0x5c, 0x31, 0x6f, 0xfa, 0x95, 0xb6, 0x83, 0xae, 0x17, 0xea,
	0xbb, 0x7d, 0x4f, 0xf5, 0xf1, 0x6c, 0x0e, 0xa6, 0x00, 0xaf, 0xe6, 0x78, 0xaa, 0x2f, 0x83, 0x5a,
	0x61, 0x2d, 0x3b, 0x55, 0x78, 0xb8, 0x7f, 0x52, 0xbd, 0x6d, 0x81, 0x12, 0x6a, 0x55, 0xee, 0x80,
	0xb2, 0xf0, 0x4a, 0x0d, 0xbc, 0x9b, 0x8e, 0x76, 0x09, 0x81, 0x5a, 0x11, 0xb2, 0x43, 0x08, 0x29,
	0x3c, 0x42, 0x34, 0x74, 0xfb, 0xa7, 0x6f, 0x77, 0xfc, 0x19, 0x00, 0x00, 0xff, 0xff, 0xae, 0x0b,
	0x1c, 0x40, 0xfc, 0x01, 0x00, 0x00,
}
