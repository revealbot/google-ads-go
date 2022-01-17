// Code generated by protoc-gen-go. DO NOT EDIT.
// source: google/ads/googleads/v9/enums/product_channel.proto

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

// Enum describing the locality of a product offer.
type ProductChannelEnum_ProductChannel int32

const (
	// Not specified.
	ProductChannelEnum_UNSPECIFIED ProductChannelEnum_ProductChannel = 0
	// Used for return value only. Represents value unknown in this version.
	ProductChannelEnum_UNKNOWN ProductChannelEnum_ProductChannel = 1
	// The item is sold online.
	ProductChannelEnum_ONLINE ProductChannelEnum_ProductChannel = 2
	// The item is sold in local stores.
	ProductChannelEnum_LOCAL ProductChannelEnum_ProductChannel = 3
)

var ProductChannelEnum_ProductChannel_name = map[int32]string{
	0: "UNSPECIFIED",
	1: "UNKNOWN",
	2: "ONLINE",
	3: "LOCAL",
}

var ProductChannelEnum_ProductChannel_value = map[string]int32{
	"UNSPECIFIED": 0,
	"UNKNOWN":     1,
	"ONLINE":      2,
	"LOCAL":       3,
}

func (x ProductChannelEnum_ProductChannel) String() string {
	return proto.EnumName(ProductChannelEnum_ProductChannel_name, int32(x))
}

func (ProductChannelEnum_ProductChannel) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_dfea4296342ab9ba, []int{0, 0}
}

// Locality of a product offer.
type ProductChannelEnum struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ProductChannelEnum) Reset()         { *m = ProductChannelEnum{} }
func (m *ProductChannelEnum) String() string { return proto.CompactTextString(m) }
func (*ProductChannelEnum) ProtoMessage()    {}
func (*ProductChannelEnum) Descriptor() ([]byte, []int) {
	return fileDescriptor_dfea4296342ab9ba, []int{0}
}

func (m *ProductChannelEnum) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ProductChannelEnum.Unmarshal(m, b)
}
func (m *ProductChannelEnum) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ProductChannelEnum.Marshal(b, m, deterministic)
}
func (m *ProductChannelEnum) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ProductChannelEnum.Merge(m, src)
}
func (m *ProductChannelEnum) XXX_Size() int {
	return xxx_messageInfo_ProductChannelEnum.Size(m)
}
func (m *ProductChannelEnum) XXX_DiscardUnknown() {
	xxx_messageInfo_ProductChannelEnum.DiscardUnknown(m)
}

var xxx_messageInfo_ProductChannelEnum proto.InternalMessageInfo

func init() {
	proto.RegisterEnum("google.ads.googleads.v9.enums.ProductChannelEnum_ProductChannel", ProductChannelEnum_ProductChannel_name, ProductChannelEnum_ProductChannel_value)
	proto.RegisterType((*ProductChannelEnum)(nil), "google.ads.googleads.v9.enums.ProductChannelEnum")
}

func init() {
	proto.RegisterFile("google/ads/googleads/v9/enums/product_channel.proto", fileDescriptor_dfea4296342ab9ba)
}

var fileDescriptor_dfea4296342ab9ba = []byte{
	// 281 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x7c, 0x90, 0xc1, 0x4a, 0xc3, 0x30,
	0x1c, 0xc6, 0x5d, 0x87, 0x13, 0x33, 0xd0, 0x12, 0xcf, 0x3b, 0x6c, 0x0f, 0x90, 0x16, 0x76, 0x8b,
	0xa7, 0xb4, 0xc6, 0x51, 0x2c, 0x69, 0x41, 0x56, 0x41, 0x0b, 0x52, 0x9b, 0x12, 0x85, 0x36, 0x29,
	0xcd, 0xba, 0x07, 0xf2, 0xe8, 0xa3, 0xf8, 0x24, 0xe2, 0x53, 0x48, 0x13, 0x57, 0xd8, 0x41, 0x2f,
	0xe1, 0xe3, 0xff, 0xff, 0x7e, 0xc9, 0xf7, 0x05, 0xac, 0x85, 0x52, 0xa2, 0xae, 0xbc, 0x82, 0x6b,
	0xcf, 0xca, 0x41, 0xed, 0x7d, 0xaf, 0x92, 0x7d, 0xa3, 0xbd, 0xb6, 0x53, 0xbc, 0x2f, 0x77, 0xcf,
	0xe5, 0x6b, 0x21, 0x65, 0x55, 0xa3, 0xb6, 0x53, 0x3b, 0x05, 0x17, 0xd6, 0x89, 0x0a, 0xae, 0xd1,
	0x08, 0xa1, 0xbd, 0x8f, 0x0c, 0xb4, 0x7a, 0x02, 0x30, 0xb5, 0x5c, 0x68, 0x31, 0x2a, 0xfb, 0x66,
	0x45, 0xc1, 0xc5, 0xf1, 0x14, 0x5e, 0x82, 0xf9, 0x96, 0xdd, 0xa7, 0x34, 0x8c, 0x6e, 0x23, 0x7a,
	0xe3, 0x9e, 0xc0, 0x39, 0x38, 0xdb, 0xb2, 0x3b, 0x96, 0x3c, 0x30, 0x77, 0x02, 0x01, 0x98, 0x25,
	0x2c, 0x8e, 0x18, 0x75, 0x1d, 0x78, 0x0e, 0x4e, 0xe3, 0x24, 0x24, 0xb1, 0x3b, 0x0d, 0xbe, 0x26,
	0x60, 0x59, 0xaa, 0x06, 0xfd, 0x1b, 0x21, 0xb8, 0x3a, 0x7e, 0x2a, 0x1d, 0x62, 0xa7, 0x93, 0xc7,
	0xe0, 0x97, 0x12, 0xaa, 0x2e, 0xa4, 0x40, 0xaa, 0x13, 0x9e, 0xa8, 0xa4, 0x29, 0x75, 0x68, 0xdf,
	0xbe, 0xe9, 0x3f, 0x3e, 0xe3, 0xda, 0x9c, 0xef, 0xce, 0x74, 0x43, 0xc8, 0x87, 0xb3, 0xd8, 0xd8,
	0xab, 0x08, 0xd7, 0xc8, 0xca, 0x41, 0x65, 0x3e, 0x1a, 0xca, 0xea, 0xcf, 0xc3, 0x3e, 0x27, 0x5c,
	0xe7, 0xe3, 0x3e, 0xcf, 0xfc, 0xdc, 0xec, 0xbf, 0x9d, 0xa5, 0x1d, 0x62, 0x4c, 0xb8, 0xc6, 0x78,
	0x74, 0x60, 0x9c, 0xf9, 0x18, 0x1b, 0xcf, 0xcb, 0xcc, 0x04, 0x5b, 0xff, 0x04, 0x00, 0x00, 0xff,
	0xff, 0x72, 0x64, 0x0b, 0x62, 0xa4, 0x01, 0x00, 0x00,
}
