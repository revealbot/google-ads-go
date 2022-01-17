// Code generated by protoc-gen-go. DO NOT EDIT.
// source: google/ads/googleads/v9/errors/feed_attribute_reference_error.proto

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

// Enum describing possible feed attribute reference errors.
type FeedAttributeReferenceErrorEnum_FeedAttributeReferenceError int32

const (
	// Enum unspecified.
	FeedAttributeReferenceErrorEnum_UNSPECIFIED FeedAttributeReferenceErrorEnum_FeedAttributeReferenceError = 0
	// The received error code is not known in this version.
	FeedAttributeReferenceErrorEnum_UNKNOWN FeedAttributeReferenceErrorEnum_FeedAttributeReferenceError = 1
	// A feed referenced by ID has been deleted.
	FeedAttributeReferenceErrorEnum_CANNOT_REFERENCE_DELETED_FEED FeedAttributeReferenceErrorEnum_FeedAttributeReferenceError = 2
	// There is no active feed with the given name.
	FeedAttributeReferenceErrorEnum_INVALID_FEED_NAME FeedAttributeReferenceErrorEnum_FeedAttributeReferenceError = 3
	// There is no feed attribute in an active feed with the given name.
	FeedAttributeReferenceErrorEnum_INVALID_FEED_ATTRIBUTE_NAME FeedAttributeReferenceErrorEnum_FeedAttributeReferenceError = 4
)

var FeedAttributeReferenceErrorEnum_FeedAttributeReferenceError_name = map[int32]string{
	0: "UNSPECIFIED",
	1: "UNKNOWN",
	2: "CANNOT_REFERENCE_DELETED_FEED",
	3: "INVALID_FEED_NAME",
	4: "INVALID_FEED_ATTRIBUTE_NAME",
}

var FeedAttributeReferenceErrorEnum_FeedAttributeReferenceError_value = map[string]int32{
	"UNSPECIFIED":                   0,
	"UNKNOWN":                       1,
	"CANNOT_REFERENCE_DELETED_FEED": 2,
	"INVALID_FEED_NAME":             3,
	"INVALID_FEED_ATTRIBUTE_NAME":   4,
}

func (x FeedAttributeReferenceErrorEnum_FeedAttributeReferenceError) String() string {
	return proto.EnumName(FeedAttributeReferenceErrorEnum_FeedAttributeReferenceError_name, int32(x))
}

func (FeedAttributeReferenceErrorEnum_FeedAttributeReferenceError) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_471a76b8ae1b4ef8, []int{0, 0}
}

// Container for enum describing possible feed attribute reference errors.
type FeedAttributeReferenceErrorEnum struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *FeedAttributeReferenceErrorEnum) Reset()         { *m = FeedAttributeReferenceErrorEnum{} }
func (m *FeedAttributeReferenceErrorEnum) String() string { return proto.CompactTextString(m) }
func (*FeedAttributeReferenceErrorEnum) ProtoMessage()    {}
func (*FeedAttributeReferenceErrorEnum) Descriptor() ([]byte, []int) {
	return fileDescriptor_471a76b8ae1b4ef8, []int{0}
}

func (m *FeedAttributeReferenceErrorEnum) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_FeedAttributeReferenceErrorEnum.Unmarshal(m, b)
}
func (m *FeedAttributeReferenceErrorEnum) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_FeedAttributeReferenceErrorEnum.Marshal(b, m, deterministic)
}
func (m *FeedAttributeReferenceErrorEnum) XXX_Merge(src proto.Message) {
	xxx_messageInfo_FeedAttributeReferenceErrorEnum.Merge(m, src)
}
func (m *FeedAttributeReferenceErrorEnum) XXX_Size() int {
	return xxx_messageInfo_FeedAttributeReferenceErrorEnum.Size(m)
}
func (m *FeedAttributeReferenceErrorEnum) XXX_DiscardUnknown() {
	xxx_messageInfo_FeedAttributeReferenceErrorEnum.DiscardUnknown(m)
}

var xxx_messageInfo_FeedAttributeReferenceErrorEnum proto.InternalMessageInfo

func init() {
	proto.RegisterEnum("google.ads.googleads.v9.errors.FeedAttributeReferenceErrorEnum_FeedAttributeReferenceError", FeedAttributeReferenceErrorEnum_FeedAttributeReferenceError_name, FeedAttributeReferenceErrorEnum_FeedAttributeReferenceError_value)
	proto.RegisterType((*FeedAttributeReferenceErrorEnum)(nil), "google.ads.googleads.v9.errors.FeedAttributeReferenceErrorEnum")
}

func init() {
	proto.RegisterFile("google/ads/googleads/v9/errors/feed_attribute_reference_error.proto", fileDescriptor_471a76b8ae1b4ef8)
}

var fileDescriptor_471a76b8ae1b4ef8 = []byte{
	// 343 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x7c, 0x91, 0xc1, 0x4a, 0xc3, 0x30,
	0x18, 0xc7, 0x6d, 0x27, 0x0a, 0xd9, 0xc1, 0x5a, 0xf0, 0x34, 0xdc, 0xb4, 0x0f, 0x90, 0x16, 0xbc,
	0xc5, 0x53, 0xd6, 0xa6, 0xa3, 0x38, 0xb3, 0x51, 0xbb, 0x0a, 0x52, 0x28, 0xdd, 0x9a, 0x95, 0xc1,
	0xd6, 0x8c, 0xa4, 0xdb, 0xa3, 0xf8, 0x00, 0x1e, 0x7d, 0x04, 0x1f, 0xc1, 0x47, 0xf1, 0x15, 0xbc,
	0x48, 0x1b, 0x3b, 0xf0, 0xe0, 0x4e, 0xf9, 0xf3, 0xe5, 0x97, 0x5f, 0xbe, 0x7c, 0x01, 0x6e, 0xc1,
	0x79, 0xb1, 0x66, 0x76, 0x96, 0x4b, 0x5b, 0xc5, 0x3a, 0xed, 0x1d, 0x9b, 0x09, 0xc1, 0x85, 0xb4,
	0x97, 0x8c, 0xe5, 0x69, 0x56, 0x55, 0x62, 0x35, 0xdf, 0x55, 0x2c, 0x15, 0x6c, 0xc9, 0x04, 0x2b,
	0x17, 0x2c, 0x6d, 0xf6, 0xe1, 0x56, 0xf0, 0x8a, 0x9b, 0x7d, 0x75, 0x12, 0x66, 0xb9, 0x84, 0x07,
	0x09, 0xdc, 0x3b, 0x50, 0x49, 0xac, 0x0f, 0x0d, 0x0c, 0x7c, 0xc6, 0x72, 0xdc, 0x7a, 0xc2, 0x56,
	0x43, 0x6a, 0x80, 0x94, 0xbb, 0x8d, 0xf5, 0xaa, 0x81, 0xde, 0x11, 0xc6, 0xbc, 0x00, 0xdd, 0x19,
	0x7d, 0x9a, 0x12, 0x37, 0xf0, 0x03, 0xe2, 0x19, 0x27, 0x66, 0x17, 0x9c, 0xcf, 0xe8, 0x03, 0x9d,
	0x3c, 0x53, 0x43, 0x33, 0x6f, 0xc1, 0xb5, 0x8b, 0x29, 0x9d, 0x44, 0x69, 0x48, 0x7c, 0x12, 0x12,
	0xea, 0x92, 0xd4, 0x23, 0x63, 0x12, 0x11, 0x2f, 0xf5, 0x09, 0xf1, 0x0c, 0xdd, 0xbc, 0x02, 0x97,
	0x01, 0x8d, 0xf1, 0x38, 0x50, 0x95, 0x94, 0xe2, 0x47, 0x62, 0x74, 0xcc, 0x01, 0xe8, 0xfd, 0x29,
	0xe3, 0x28, 0x0a, 0x83, 0xe1, 0x2c, 0x22, 0x0a, 0x38, 0x1d, 0x7e, 0x6b, 0xc0, 0x5a, 0xf0, 0x0d,
	0x3c, 0xfe, 0xc6, 0xe1, 0xcd, 0x91, 0xe6, 0xa7, 0xf5, 0x94, 0xa6, 0xda, 0x8b, 0xf7, 0xeb, 0x28,
	0xf8, 0x3a, 0x2b, 0x0b, 0xc8, 0x45, 0x61, 0x17, 0xac, 0x6c, 0x66, 0xd8, 0x0e, 0x7f, 0xbb, 0x92,
	0xff, 0xfd, 0xc5, 0xbd, 0x5a, 0xde, 0xf4, 0xce, 0x08, 0xe3, 0x77, 0xbd, 0x3f, 0x52, 0x32, 0x9c,
	0x4b, 0xa8, 0x62, 0x9d, 0x62, 0x07, 0x36, 0x57, 0xca, 0xcf, 0x16, 0x48, 0x70, 0x2e, 0x93, 0x03,
	0x90, 0xc4, 0x4e, 0xa2, 0x80, 0x2f, 0xdd, 0x52, 0x55, 0x84, 0x70, 0x2e, 0x11, 0x3a, 0x20, 0x08,
	0xc5, 0x0e, 0x42, 0x0a, 0x9a, 0x9f, 0x35, 0xdd, 0xdd, 0xfd, 0x04, 0x00, 0x00, 0xff, 0xff, 0xfb,
	0x20, 0xe4, 0x39, 0x28, 0x02, 0x00, 0x00,
}
