// Code generated by protoc-gen-go. DO NOT EDIT.
// source: google/ads/googleads/v9/enums/shared_set_status.proto

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

// Enum listing the possible shared set statuses.
type SharedSetStatusEnum_SharedSetStatus int32

const (
	// Not specified.
	SharedSetStatusEnum_UNSPECIFIED SharedSetStatusEnum_SharedSetStatus = 0
	// Used for return value only. Represents value unknown in this version.
	SharedSetStatusEnum_UNKNOWN SharedSetStatusEnum_SharedSetStatus = 1
	// The shared set is enabled.
	SharedSetStatusEnum_ENABLED SharedSetStatusEnum_SharedSetStatus = 2
	// The shared set is removed and can no longer be used.
	SharedSetStatusEnum_REMOVED SharedSetStatusEnum_SharedSetStatus = 3
)

var SharedSetStatusEnum_SharedSetStatus_name = map[int32]string{
	0: "UNSPECIFIED",
	1: "UNKNOWN",
	2: "ENABLED",
	3: "REMOVED",
}

var SharedSetStatusEnum_SharedSetStatus_value = map[string]int32{
	"UNSPECIFIED": 0,
	"UNKNOWN":     1,
	"ENABLED":     2,
	"REMOVED":     3,
}

func (x SharedSetStatusEnum_SharedSetStatus) String() string {
	return proto.EnumName(SharedSetStatusEnum_SharedSetStatus_name, int32(x))
}

func (SharedSetStatusEnum_SharedSetStatus) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_3337f4109e961f02, []int{0, 0}
}

// Container for enum describing types of shared set statuses.
type SharedSetStatusEnum struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *SharedSetStatusEnum) Reset()         { *m = SharedSetStatusEnum{} }
func (m *SharedSetStatusEnum) String() string { return proto.CompactTextString(m) }
func (*SharedSetStatusEnum) ProtoMessage()    {}
func (*SharedSetStatusEnum) Descriptor() ([]byte, []int) {
	return fileDescriptor_3337f4109e961f02, []int{0}
}

func (m *SharedSetStatusEnum) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SharedSetStatusEnum.Unmarshal(m, b)
}
func (m *SharedSetStatusEnum) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SharedSetStatusEnum.Marshal(b, m, deterministic)
}
func (m *SharedSetStatusEnum) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SharedSetStatusEnum.Merge(m, src)
}
func (m *SharedSetStatusEnum) XXX_Size() int {
	return xxx_messageInfo_SharedSetStatusEnum.Size(m)
}
func (m *SharedSetStatusEnum) XXX_DiscardUnknown() {
	xxx_messageInfo_SharedSetStatusEnum.DiscardUnknown(m)
}

var xxx_messageInfo_SharedSetStatusEnum proto.InternalMessageInfo

func init() {
	proto.RegisterEnum("google.ads.googleads.v9.enums.SharedSetStatusEnum_SharedSetStatus", SharedSetStatusEnum_SharedSetStatus_name, SharedSetStatusEnum_SharedSetStatus_value)
	proto.RegisterType((*SharedSetStatusEnum)(nil), "google.ads.googleads.v9.enums.SharedSetStatusEnum")
}

func init() {
	proto.RegisterFile("google/ads/googleads/v9/enums/shared_set_status.proto", fileDescriptor_3337f4109e961f02)
}

var fileDescriptor_3337f4109e961f02 = []byte{
	// 285 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x7c, 0x50, 0xd1, 0x4a, 0xc3, 0x30,
	0x14, 0xb5, 0x1d, 0x28, 0x64, 0x0f, 0x2b, 0xd5, 0xd7, 0x3d, 0x6c, 0x1f, 0x90, 0x16, 0xc4, 0x97,
	0xf8, 0x94, 0xda, 0x38, 0x8a, 0xda, 0x15, 0xcb, 0x2a, 0x48, 0x61, 0x56, 0x13, 0xa2, 0xb0, 0x36,
	0xa3, 0xb7, 0xdd, 0x07, 0xf9, 0xe8, 0xa7, 0xf8, 0x27, 0xfa, 0x15, 0x92, 0xc4, 0xf5, 0x61, 0xa0,
	0x2f, 0xe1, 0xdc, 0x7b, 0xce, 0xc9, 0x3d, 0xf7, 0xa2, 0x0b, 0xa9, 0x94, 0xdc, 0x88, 0xa0, 0xe2,
	0x10, 0x58, 0xa8, 0xd1, 0x2e, 0x0c, 0x44, 0xd3, 0xd7, 0x10, 0xc0, 0x6b, 0xd5, 0x0a, 0xbe, 0x06,
	0xd1, 0xad, 0xa1, 0xab, 0xba, 0x1e, 0xf0, 0xb6, 0x55, 0x9d, 0xf2, 0xa7, 0x56, 0x8b, 0x2b, 0x0e,
	0x78, 0xb0, 0xe1, 0x5d, 0x88, 0x8d, 0x6d, 0xfe, 0x84, 0x4e, 0x73, 0xe3, 0xcc, 0x45, 0x97, 0x1b,
	0x1f, 0x6b, 0xfa, 0x7a, 0x9e, 0xa0, 0xc9, 0x41, 0xdb, 0x9f, 0xa0, 0xf1, 0x2a, 0xcd, 0x33, 0x76,
	0x95, 0x5c, 0x27, 0x2c, 0xf6, 0x8e, 0xfc, 0x31, 0x3a, 0x59, 0xa5, 0x37, 0xe9, 0xf2, 0x21, 0xf5,
	0x1c, 0x5d, 0xb0, 0x94, 0x46, 0xb7, 0x2c, 0xf6, 0x5c, 0x5d, 0xdc, 0xb3, 0xbb, 0x65, 0xc1, 0x62,
	0x6f, 0x14, 0x7d, 0x39, 0x68, 0xf6, 0xa2, 0x6a, 0xfc, 0x6f, 0x8e, 0xe8, 0xec, 0x60, 0x5c, 0xa6,
	0xc3, 0x67, 0xce, 0x63, 0xf4, 0x6b, 0x93, 0x6a, 0x53, 0x35, 0x12, 0xab, 0x56, 0x06, 0x52, 0x34,
	0x66, 0xb5, 0xfd, 0x15, 0xb6, 0x6f, 0xf0, 0xc7, 0x51, 0x2e, 0xcd, 0xfb, 0xee, 0x8e, 0x16, 0x94,
	0x7e, 0xb8, 0xd3, 0x85, 0xfd, 0x8a, 0x72, 0xc0, 0x16, 0x6a, 0x54, 0x84, 0x58, 0x6f, 0x0c, 0x9f,
	0x7b, 0xbe, 0xa4, 0x1c, 0xca, 0x81, 0x2f, 0x8b, 0xb0, 0x34, 0xfc, 0xb7, 0x3b, 0xb3, 0x4d, 0x42,
	0x28, 0x07, 0x42, 0x06, 0x05, 0x21, 0x45, 0x48, 0x88, 0xd1, 0x3c, 0x1f, 0x9b, 0x60, 0xe7, 0x3f,
	0x01, 0x00, 0x00, 0xff, 0xff, 0x95, 0x7a, 0xe8, 0xa3, 0xac, 0x01, 0x00, 0x00,
}
