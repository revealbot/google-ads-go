// Code generated by protoc-gen-go. DO NOT EDIT.
// source: google/ads/googleads/v9/errors/distinct_error.proto

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

// Enum describing possible distinct errors.
type DistinctErrorEnum_DistinctError int32

const (
	// Enum unspecified.
	DistinctErrorEnum_UNSPECIFIED DistinctErrorEnum_DistinctError = 0
	// The received error code is not known in this version.
	DistinctErrorEnum_UNKNOWN DistinctErrorEnum_DistinctError = 1
	// Duplicate element.
	DistinctErrorEnum_DUPLICATE_ELEMENT DistinctErrorEnum_DistinctError = 2
	// Duplicate type.
	DistinctErrorEnum_DUPLICATE_TYPE DistinctErrorEnum_DistinctError = 3
)

var DistinctErrorEnum_DistinctError_name = map[int32]string{
	0: "UNSPECIFIED",
	1: "UNKNOWN",
	2: "DUPLICATE_ELEMENT",
	3: "DUPLICATE_TYPE",
}

var DistinctErrorEnum_DistinctError_value = map[string]int32{
	"UNSPECIFIED":       0,
	"UNKNOWN":           1,
	"DUPLICATE_ELEMENT": 2,
	"DUPLICATE_TYPE":    3,
}

func (x DistinctErrorEnum_DistinctError) String() string {
	return proto.EnumName(DistinctErrorEnum_DistinctError_name, int32(x))
}

func (DistinctErrorEnum_DistinctError) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_f425fa1f40ddce89, []int{0, 0}
}

// Container for enum describing possible distinct errors.
type DistinctErrorEnum struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *DistinctErrorEnum) Reset()         { *m = DistinctErrorEnum{} }
func (m *DistinctErrorEnum) String() string { return proto.CompactTextString(m) }
func (*DistinctErrorEnum) ProtoMessage()    {}
func (*DistinctErrorEnum) Descriptor() ([]byte, []int) {
	return fileDescriptor_f425fa1f40ddce89, []int{0}
}

func (m *DistinctErrorEnum) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DistinctErrorEnum.Unmarshal(m, b)
}
func (m *DistinctErrorEnum) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DistinctErrorEnum.Marshal(b, m, deterministic)
}
func (m *DistinctErrorEnum) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DistinctErrorEnum.Merge(m, src)
}
func (m *DistinctErrorEnum) XXX_Size() int {
	return xxx_messageInfo_DistinctErrorEnum.Size(m)
}
func (m *DistinctErrorEnum) XXX_DiscardUnknown() {
	xxx_messageInfo_DistinctErrorEnum.DiscardUnknown(m)
}

var xxx_messageInfo_DistinctErrorEnum proto.InternalMessageInfo

func init() {
	proto.RegisterEnum("google.ads.googleads.v9.errors.DistinctErrorEnum_DistinctError", DistinctErrorEnum_DistinctError_name, DistinctErrorEnum_DistinctError_value)
	proto.RegisterType((*DistinctErrorEnum)(nil), "google.ads.googleads.v9.errors.DistinctErrorEnum")
}

func init() {
	proto.RegisterFile("google/ads/googleads/v9/errors/distinct_error.proto", fileDescriptor_f425fa1f40ddce89)
}

var fileDescriptor_f425fa1f40ddce89 = []byte{
	// 288 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x7c, 0x90, 0xdd, 0x4a, 0xc3, 0x30,
	0x1c, 0xc5, 0x5d, 0x07, 0x0a, 0x19, 0x6a, 0x17, 0xf0, 0x76, 0x17, 0x7d, 0x80, 0xb4, 0xb0, 0xbb,
	0x78, 0x95, 0xad, 0x71, 0x14, 0x67, 0x0c, 0xd8, 0xd6, 0x0f, 0x0a, 0xa3, 0x2e, 0x25, 0x14, 0xd6,
	0x66, 0x24, 0x75, 0x0f, 0xe4, 0xa5, 0x8f, 0xe2, 0xa3, 0x08, 0xbe, 0x83, 0xb4, 0xd9, 0x2a, 0xbb,
	0xd0, 0xab, 0x9c, 0x1c, 0x7e, 0xe7, 0xff, 0x05, 0xa6, 0x52, 0x29, 0xb9, 0x29, 0xfc, 0x5c, 0x18,
	0xdf, 0xca, 0x56, 0xed, 0x02, 0xbf, 0xd0, 0x5a, 0x69, 0xe3, 0x8b, 0xd2, 0x34, 0x65, 0xbd, 0x6e,
	0x56, 0xdd, 0x1f, 0x6d, 0xb5, 0x6a, 0x14, 0x9c, 0x58, 0x12, 0xe5, 0xc2, 0xa0, 0x3e, 0x84, 0x76,
	0x01, 0xb2, 0x21, 0xaf, 0x02, 0xe3, 0x70, 0x9f, 0xa3, 0xad, 0x43, 0xeb, 0xb7, 0xca, 0x7b, 0x02,
	0xe7, 0x47, 0x26, 0xbc, 0x04, 0xa3, 0x84, 0x3d, 0x70, 0x3a, 0x8f, 0x6e, 0x22, 0x1a, 0xba, 0x27,
	0x70, 0x04, 0xce, 0x12, 0x76, 0xcb, 0xee, 0x1f, 0x99, 0x3b, 0x80, 0x57, 0x60, 0x1c, 0x26, 0x7c,
	0x19, 0xcd, 0x49, 0x4c, 0x57, 0x74, 0x49, 0xef, 0x28, 0x8b, 0x5d, 0x07, 0x42, 0x70, 0xf1, 0x6b,
	0xc7, 0xcf, 0x9c, 0xba, 0xc3, 0xd9, 0xf7, 0x00, 0x78, 0x6b, 0x55, 0xa1, 0xff, 0xa7, 0x9a, 0xc1,
	0xa3, 0xf6, 0xbc, 0xdd, 0x84, 0x0f, 0x5e, 0xc2, 0x7d, 0x4a, 0xaa, 0x4d, 0x5e, 0x4b, 0xa4, 0xb4,
	0xf4, 0x65, 0x51, 0x77, 0x7b, 0x1e, 0x0e, 0xb2, 0x2d, 0xcd, 0x5f, 0xf7, 0xb9, 0xb6, 0xcf, 0xbb,
	0x33, 0x5c, 0x10, 0xf2, 0xe1, 0x4c, 0x16, 0xb6, 0x18, 0x11, 0x06, 0x59, 0xd9, 0xaa, 0x34, 0x40,
	0x5d, 0x4b, 0xf3, 0x79, 0x00, 0x32, 0x22, 0x4c, 0xd6, 0x03, 0x59, 0x1a, 0x64, 0x16, 0xf8, 0x72,
	0x3c, 0xeb, 0x62, 0x4c, 0x84, 0xc1, 0xb8, 0x47, 0x30, 0x4e, 0x03, 0x8c, 0x2d, 0xf4, 0x7a, 0xda,
	0x4d, 0x37, 0xfd, 0x09, 0x00, 0x00, 0xff, 0xff, 0x34, 0xfd, 0x53, 0x0e, 0xbc, 0x01, 0x00, 0x00,
}
