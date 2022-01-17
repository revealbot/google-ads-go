// Code generated by protoc-gen-go. DO NOT EDIT.
// source: google/ads/googleads/v9/enums/vanity_pharma_display_url_mode.proto

package enums

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

// Enum describing possible display modes for vanity pharma URLs.
type VanityPharmaDisplayUrlModeEnum_VanityPharmaDisplayUrlMode int32

const (
	// Not specified.
	VanityPharmaDisplayUrlModeEnum_UNSPECIFIED VanityPharmaDisplayUrlModeEnum_VanityPharmaDisplayUrlMode = 0
	// Used for return value only. Represents value unknown in this version.
	VanityPharmaDisplayUrlModeEnum_UNKNOWN VanityPharmaDisplayUrlModeEnum_VanityPharmaDisplayUrlMode = 1
	// Replace vanity pharma URL with manufacturer website url.
	VanityPharmaDisplayUrlModeEnum_MANUFACTURER_WEBSITE_URL VanityPharmaDisplayUrlModeEnum_VanityPharmaDisplayUrlMode = 2
	// Replace vanity pharma URL with description of the website.
	VanityPharmaDisplayUrlModeEnum_WEBSITE_DESCRIPTION VanityPharmaDisplayUrlModeEnum_VanityPharmaDisplayUrlMode = 3
)

var VanityPharmaDisplayUrlModeEnum_VanityPharmaDisplayUrlMode_name = map[int32]string{
	0: "UNSPECIFIED",
	1: "UNKNOWN",
	2: "MANUFACTURER_WEBSITE_URL",
	3: "WEBSITE_DESCRIPTION",
}

var VanityPharmaDisplayUrlModeEnum_VanityPharmaDisplayUrlMode_value = map[string]int32{
	"UNSPECIFIED":              0,
	"UNKNOWN":                  1,
	"MANUFACTURER_WEBSITE_URL": 2,
	"WEBSITE_DESCRIPTION":      3,
}

func (x VanityPharmaDisplayUrlModeEnum_VanityPharmaDisplayUrlMode) String() string {
	return proto.EnumName(VanityPharmaDisplayUrlModeEnum_VanityPharmaDisplayUrlMode_name, int32(x))
}

func (VanityPharmaDisplayUrlModeEnum_VanityPharmaDisplayUrlMode) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_f1330998531b8af0, []int{0, 0}
}

// The display mode for vanity pharma URLs.
type VanityPharmaDisplayUrlModeEnum struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *VanityPharmaDisplayUrlModeEnum) Reset()         { *m = VanityPharmaDisplayUrlModeEnum{} }
func (m *VanityPharmaDisplayUrlModeEnum) String() string { return proto.CompactTextString(m) }
func (*VanityPharmaDisplayUrlModeEnum) ProtoMessage()    {}
func (*VanityPharmaDisplayUrlModeEnum) Descriptor() ([]byte, []int) {
	return fileDescriptor_f1330998531b8af0, []int{0}
}

func (m *VanityPharmaDisplayUrlModeEnum) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_VanityPharmaDisplayUrlModeEnum.Unmarshal(m, b)
}
func (m *VanityPharmaDisplayUrlModeEnum) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_VanityPharmaDisplayUrlModeEnum.Marshal(b, m, deterministic)
}
func (m *VanityPharmaDisplayUrlModeEnum) XXX_Merge(src proto.Message) {
	xxx_messageInfo_VanityPharmaDisplayUrlModeEnum.Merge(m, src)
}
func (m *VanityPharmaDisplayUrlModeEnum) XXX_Size() int {
	return xxx_messageInfo_VanityPharmaDisplayUrlModeEnum.Size(m)
}
func (m *VanityPharmaDisplayUrlModeEnum) XXX_DiscardUnknown() {
	xxx_messageInfo_VanityPharmaDisplayUrlModeEnum.DiscardUnknown(m)
}

var xxx_messageInfo_VanityPharmaDisplayUrlModeEnum proto.InternalMessageInfo

func init() {
	proto.RegisterEnum("google.ads.googleads.v9.enums.VanityPharmaDisplayUrlModeEnum_VanityPharmaDisplayUrlMode", VanityPharmaDisplayUrlModeEnum_VanityPharmaDisplayUrlMode_name, VanityPharmaDisplayUrlModeEnum_VanityPharmaDisplayUrlMode_value)
	proto.RegisterType((*VanityPharmaDisplayUrlModeEnum)(nil), "google.ads.googleads.v9.enums.VanityPharmaDisplayUrlModeEnum")
}

func init() {
	proto.RegisterFile("google/ads/googleads/v9/enums/vanity_pharma_display_url_mode.proto", fileDescriptor_f1330998531b8af0)
}

var fileDescriptor_f1330998531b8af0 = []byte{
	// 331 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x7c, 0x90, 0x41, 0x4a, 0xf3, 0x40,
	0x1c, 0xc5, 0xbf, 0xa4, 0xf0, 0x09, 0xd3, 0x85, 0x21, 0x2e, 0x14, 0xb1, 0x4a, 0x7b, 0x80, 0x49,
	0xc0, 0xdd, 0xb8, 0x9a, 0xb4, 0xd3, 0x12, 0xb4, 0x69, 0x48, 0x9b, 0x14, 0x24, 0x10, 0x46, 0x27,
	0x8c, 0x85, 0x24, 0x13, 0x33, 0x6d, 0xa1, 0xe7, 0xf0, 0x06, 0x2e, 0x3d, 0x8a, 0x47, 0x71, 0xed,
	0x01, 0x24, 0x33, 0xb6, 0xbb, 0xba, 0x09, 0x8f, 0xff, 0x7b, 0xf9, 0xf1, 0xe6, 0x01, 0x8f, 0x0b,
	0xc1, 0x8b, 0xdc, 0xa1, 0x4c, 0x3a, 0x5a, 0xb6, 0x6a, 0xeb, 0x3a, 0x79, 0xb5, 0x29, 0xa5, 0xb3,
	0xa5, 0xd5, 0x6a, 0xbd, 0xcb, 0xea, 0x17, 0xda, 0x94, 0x34, 0x63, 0x2b, 0x59, 0x17, 0x74, 0x97,
	0x6d, 0x9a, 0x22, 0x2b, 0x05, 0xcb, 0x61, 0xdd, 0x88, 0xb5, 0xb0, 0x7b, 0xfa, 0x47, 0x48, 0x99,
	0x84, 0x07, 0x06, 0xdc, 0xba, 0x50, 0x31, 0x06, 0x6f, 0x06, 0xb8, 0x4e, 0x14, 0x27, 0x54, 0x98,
	0x91, 0xa6, 0xc4, 0x4d, 0x31, 0x15, 0x2c, 0x27, 0xd5, 0xa6, 0x1c, 0xbc, 0x82, 0xcb, 0xe3, 0x09,
	0xfb, 0x14, 0x74, 0xe3, 0x60, 0x1e, 0x92, 0xa1, 0x3f, 0xf6, 0xc9, 0xc8, 0xfa, 0x67, 0x77, 0xc1,
	0x49, 0x1c, 0xdc, 0x07, 0xb3, 0x65, 0x60, 0x19, 0xf6, 0x15, 0xb8, 0x98, 0xe2, 0x20, 0x1e, 0xe3,
	0xe1, 0x22, 0x8e, 0x48, 0x94, 0x2d, 0x89, 0x37, 0xf7, 0x17, 0x24, 0x8b, 0xa3, 0x07, 0xcb, 0xb4,
	0xcf, 0xc1, 0xd9, 0xfe, 0x30, 0x22, 0xf3, 0x61, 0xe4, 0x87, 0x0b, 0x7f, 0x16, 0x58, 0x1d, 0xef,
	0xdb, 0x00, 0xfd, 0x67, 0x51, 0xc2, 0x3f, 0xbb, 0x7b, 0x37, 0xc7, 0x6b, 0x85, 0xed, 0xdb, 0x43,
	0xe3, 0xf1, 0x77, 0x41, 0xc8, 0x45, 0x41, 0x2b, 0x0e, 0x45, 0xc3, 0x1d, 0x9e, 0x57, 0x6a, 0x99,
	0xfd, 0xa2, 0xf5, 0x4a, 0x1e, 0x19, 0xf8, 0x4e, 0x7d, 0xdf, 0xcd, 0xce, 0x04, 0xe3, 0x0f, 0xb3,
	0x37, 0xd1, 0x28, 0xcc, 0x24, 0xd4, 0xb2, 0x55, 0x89, 0x0b, 0xdb, 0x91, 0xe4, 0xe7, 0xde, 0x4f,
	0x31, 0x93, 0xe9, 0xc1, 0x4f, 0x13, 0x37, 0x55, 0xfe, 0x97, 0xd9, 0xd7, 0x47, 0x84, 0x30, 0x93,
	0x08, 0x1d, 0x12, 0x08, 0x25, 0x2e, 0x42, 0x2a, 0xf3, 0xf4, 0x5f, 0x15, 0xbb, 0xfd, 0x09, 0x00,
	0x00, 0xff, 0xff, 0xf6, 0xd0, 0xb9, 0xe3, 0xf8, 0x01, 0x00, 0x00,
}
