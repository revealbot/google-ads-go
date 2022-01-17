// Code generated by protoc-gen-go. DO NOT EDIT.
// source: google/ads/googleads/v9/enums/policy_review_status.proto

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

// The possible policy review statuses.
type PolicyReviewStatusEnum_PolicyReviewStatus int32

const (
	// No value has been specified.
	PolicyReviewStatusEnum_UNSPECIFIED PolicyReviewStatusEnum_PolicyReviewStatus = 0
	// The received value is not known in this version.
	//
	// This is a response-only value.
	PolicyReviewStatusEnum_UNKNOWN PolicyReviewStatusEnum_PolicyReviewStatus = 1
	// Currently under review.
	PolicyReviewStatusEnum_REVIEW_IN_PROGRESS PolicyReviewStatusEnum_PolicyReviewStatus = 2
	// Primary review complete. Other reviews may be continuing.
	PolicyReviewStatusEnum_REVIEWED PolicyReviewStatusEnum_PolicyReviewStatus = 3
	// The resource has been resubmitted for approval or its policy decision has
	// been appealed.
	PolicyReviewStatusEnum_UNDER_APPEAL PolicyReviewStatusEnum_PolicyReviewStatus = 4
)

var PolicyReviewStatusEnum_PolicyReviewStatus_name = map[int32]string{
	0: "UNSPECIFIED",
	1: "UNKNOWN",
	2: "REVIEW_IN_PROGRESS",
	3: "REVIEWED",
	4: "UNDER_APPEAL",
}

var PolicyReviewStatusEnum_PolicyReviewStatus_value = map[string]int32{
	"UNSPECIFIED":        0,
	"UNKNOWN":            1,
	"REVIEW_IN_PROGRESS": 2,
	"REVIEWED":           3,
	"UNDER_APPEAL":       4,
}

func (x PolicyReviewStatusEnum_PolicyReviewStatus) String() string {
	return proto.EnumName(PolicyReviewStatusEnum_PolicyReviewStatus_name, int32(x))
}

func (PolicyReviewStatusEnum_PolicyReviewStatus) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_96f10d189faa48a6, []int{0, 0}
}

// Container for enum describing possible policy review statuses.
type PolicyReviewStatusEnum struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *PolicyReviewStatusEnum) Reset()         { *m = PolicyReviewStatusEnum{} }
func (m *PolicyReviewStatusEnum) String() string { return proto.CompactTextString(m) }
func (*PolicyReviewStatusEnum) ProtoMessage()    {}
func (*PolicyReviewStatusEnum) Descriptor() ([]byte, []int) {
	return fileDescriptor_96f10d189faa48a6, []int{0}
}

func (m *PolicyReviewStatusEnum) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PolicyReviewStatusEnum.Unmarshal(m, b)
}
func (m *PolicyReviewStatusEnum) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PolicyReviewStatusEnum.Marshal(b, m, deterministic)
}
func (m *PolicyReviewStatusEnum) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PolicyReviewStatusEnum.Merge(m, src)
}
func (m *PolicyReviewStatusEnum) XXX_Size() int {
	return xxx_messageInfo_PolicyReviewStatusEnum.Size(m)
}
func (m *PolicyReviewStatusEnum) XXX_DiscardUnknown() {
	xxx_messageInfo_PolicyReviewStatusEnum.DiscardUnknown(m)
}

var xxx_messageInfo_PolicyReviewStatusEnum proto.InternalMessageInfo

func init() {
	proto.RegisterEnum("google.ads.googleads.v9.enums.PolicyReviewStatusEnum_PolicyReviewStatus", PolicyReviewStatusEnum_PolicyReviewStatus_name, PolicyReviewStatusEnum_PolicyReviewStatus_value)
	proto.RegisterType((*PolicyReviewStatusEnum)(nil), "google.ads.googleads.v9.enums.PolicyReviewStatusEnum")
}

func init() {
	proto.RegisterFile("google/ads/googleads/v9/enums/policy_review_status.proto", fileDescriptor_96f10d189faa48a6)
}

var fileDescriptor_96f10d189faa48a6 = []byte{
	// 317 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x7c, 0x90, 0xd1, 0x4a, 0xf3, 0x30,
	0x1c, 0xc5, 0xbf, 0x76, 0x1f, 0x2a, 0xd9, 0xc0, 0x90, 0x8b, 0x79, 0xb5, 0x8b, 0xed, 0x01, 0xd2,
	0x82, 0x37, 0x12, 0xaf, 0x32, 0x17, 0x47, 0x51, 0xb2, 0xd0, 0xb2, 0x0e, 0xa4, 0x50, 0xea, 0x5a,
	0xc2, 0x64, 0x6b, 0xc6, 0xb2, 0x56, 0xbc, 0xf7, 0x49, 0xbc, 0xf4, 0x51, 0x7c, 0x14, 0xf1, 0x21,
	0xa4, 0x89, 0xeb, 0xcd, 0xd0, 0x9b, 0x72, 0xf8, 0x9f, 0xf3, 0x2b, 0x27, 0x07, 0x5c, 0x49, 0xa5,
	0xe4, 0xba, 0xf0, 0xb2, 0x5c, 0x7b, 0x56, 0x36, 0xaa, 0xf6, 0xbd, 0xa2, 0xac, 0x36, 0xda, 0xdb,
	0xaa, 0xf5, 0x6a, 0xf9, 0x92, 0xee, 0x8a, 0x7a, 0x55, 0x3c, 0xa7, 0x7a, 0x9f, 0xed, 0x2b, 0x8d,
	0xb7, 0x3b, 0xb5, 0x57, 0x68, 0x60, 0xe3, 0x38, 0xcb, 0x35, 0x6e, 0x49, 0x5c, 0xfb, 0xd8, 0x90,
	0xa3, 0x57, 0x07, 0xf4, 0x85, 0xa1, 0x43, 0x03, 0x47, 0x86, 0x65, 0x65, 0xb5, 0x19, 0x3d, 0x01,
	0x74, 0xec, 0xa0, 0x73, 0xd0, 0x9d, 0xf3, 0x48, 0xb0, 0x9b, 0xe0, 0x36, 0x60, 0x13, 0xf8, 0x0f,
	0x75, 0xc1, 0xe9, 0x9c, 0xdf, 0xf1, 0xd9, 0x82, 0x43, 0x07, 0xf5, 0x01, 0x0a, 0x59, 0x1c, 0xb0,
	0x45, 0x1a, 0xf0, 0x54, 0x84, 0xb3, 0x69, 0xc8, 0xa2, 0x08, 0xba, 0xa8, 0x07, 0xce, 0xec, 0x9d,
	0x4d, 0x60, 0x07, 0x41, 0xd0, 0x9b, 0xf3, 0x09, 0x0b, 0x53, 0x2a, 0x04, 0xa3, 0xf7, 0xf0, 0xff,
	0xf8, 0xcb, 0x01, 0xc3, 0xa5, 0xda, 0xe0, 0x3f, 0xcb, 0x8e, 0x2f, 0x8e, 0xfb, 0x88, 0xe6, 0x91,
	0xc2, 0x79, 0x18, 0xff, 0x90, 0x52, 0xad, 0xb3, 0x52, 0x62, 0xb5, 0x93, 0x9e, 0x2c, 0x4a, 0x33,
	0xc1, 0x61, 0xb0, 0xed, 0x4a, 0xff, 0xb2, 0xdf, 0xb5, 0xf9, 0xbe, 0xb9, 0x9d, 0x29, 0xa5, 0xef,
	0xee, 0x60, 0x6a, 0x7f, 0x45, 0x73, 0x8d, 0xad, 0x6c, 0x54, 0xec, 0xe3, 0x66, 0x15, 0xfd, 0x71,
	0xf0, 0x13, 0x9a, 0xeb, 0xa4, 0xf5, 0x93, 0xd8, 0x4f, 0x8c, 0xff, 0xe9, 0x0e, 0xed, 0x91, 0x10,
	0x9a, 0x6b, 0x42, 0xda, 0x04, 0x21, 0xb1, 0x4f, 0x88, 0xc9, 0x3c, 0x9e, 0x98, 0x62, 0x97, 0xdf,
	0x01, 0x00, 0x00, 0xff, 0xff, 0xa7, 0xd5, 0xb7, 0xc7, 0xd7, 0x01, 0x00, 0x00,
}
