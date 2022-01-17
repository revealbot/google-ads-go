// Code generated by protoc-gen-go. DO NOT EDIT.
// source: google/ads/googleads/v9/enums/user_list_logical_rule_operator.proto

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

// Enum describing possible user list logical rule operators.
type UserListLogicalRuleOperatorEnum_UserListLogicalRuleOperator int32

const (
	// Not specified.
	UserListLogicalRuleOperatorEnum_UNSPECIFIED UserListLogicalRuleOperatorEnum_UserListLogicalRuleOperator = 0
	// Used for return value only. Represents value unknown in this version.
	UserListLogicalRuleOperatorEnum_UNKNOWN UserListLogicalRuleOperatorEnum_UserListLogicalRuleOperator = 1
	// And - all of the operands.
	UserListLogicalRuleOperatorEnum_ALL UserListLogicalRuleOperatorEnum_UserListLogicalRuleOperator = 2
	// Or - at least one of the operands.
	UserListLogicalRuleOperatorEnum_ANY UserListLogicalRuleOperatorEnum_UserListLogicalRuleOperator = 3
	// Not - none of the operands.
	UserListLogicalRuleOperatorEnum_NONE UserListLogicalRuleOperatorEnum_UserListLogicalRuleOperator = 4
)

var UserListLogicalRuleOperatorEnum_UserListLogicalRuleOperator_name = map[int32]string{
	0: "UNSPECIFIED",
	1: "UNKNOWN",
	2: "ALL",
	3: "ANY",
	4: "NONE",
}

var UserListLogicalRuleOperatorEnum_UserListLogicalRuleOperator_value = map[string]int32{
	"UNSPECIFIED": 0,
	"UNKNOWN":     1,
	"ALL":         2,
	"ANY":         3,
	"NONE":        4,
}

func (x UserListLogicalRuleOperatorEnum_UserListLogicalRuleOperator) String() string {
	return proto.EnumName(UserListLogicalRuleOperatorEnum_UserListLogicalRuleOperator_name, int32(x))
}

func (UserListLogicalRuleOperatorEnum_UserListLogicalRuleOperator) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_f27f55432bba4f7d, []int{0, 0}
}

// The logical operator of the rule.
type UserListLogicalRuleOperatorEnum struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *UserListLogicalRuleOperatorEnum) Reset()         { *m = UserListLogicalRuleOperatorEnum{} }
func (m *UserListLogicalRuleOperatorEnum) String() string { return proto.CompactTextString(m) }
func (*UserListLogicalRuleOperatorEnum) ProtoMessage()    {}
func (*UserListLogicalRuleOperatorEnum) Descriptor() ([]byte, []int) {
	return fileDescriptor_f27f55432bba4f7d, []int{0}
}

func (m *UserListLogicalRuleOperatorEnum) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UserListLogicalRuleOperatorEnum.Unmarshal(m, b)
}
func (m *UserListLogicalRuleOperatorEnum) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UserListLogicalRuleOperatorEnum.Marshal(b, m, deterministic)
}
func (m *UserListLogicalRuleOperatorEnum) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UserListLogicalRuleOperatorEnum.Merge(m, src)
}
func (m *UserListLogicalRuleOperatorEnum) XXX_Size() int {
	return xxx_messageInfo_UserListLogicalRuleOperatorEnum.Size(m)
}
func (m *UserListLogicalRuleOperatorEnum) XXX_DiscardUnknown() {
	xxx_messageInfo_UserListLogicalRuleOperatorEnum.DiscardUnknown(m)
}

var xxx_messageInfo_UserListLogicalRuleOperatorEnum proto.InternalMessageInfo

func init() {
	proto.RegisterEnum("google.ads.googleads.v9.enums.UserListLogicalRuleOperatorEnum_UserListLogicalRuleOperator", UserListLogicalRuleOperatorEnum_UserListLogicalRuleOperator_name, UserListLogicalRuleOperatorEnum_UserListLogicalRuleOperator_value)
	proto.RegisterType((*UserListLogicalRuleOperatorEnum)(nil), "google.ads.googleads.v9.enums.UserListLogicalRuleOperatorEnum")
}

func init() {
	proto.RegisterFile("google/ads/googleads/v9/enums/user_list_logical_rule_operator.proto", fileDescriptor_f27f55432bba4f7d)
}

var fileDescriptor_f27f55432bba4f7d = []byte{
	// 307 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x7c, 0x90, 0xb1, 0x6a, 0xf3, 0x30,
	0x14, 0x85, 0xff, 0x38, 0xe1, 0x4f, 0x51, 0x86, 0x1a, 0xaf, 0x25, 0xb4, 0xc9, 0x03, 0xc8, 0x86,
	0x6e, 0xea, 0xa4, 0xa4, 0x6e, 0x08, 0x35, 0x4a, 0x68, 0x49, 0x42, 0x8b, 0xc1, 0xb8, 0xb1, 0x10,
	0x06, 0xc5, 0x32, 0xba, 0x76, 0x86, 0x3e, 0x4e, 0xc7, 0x3e, 0x4a, 0x1f, 0xa5, 0x7b, 0xf7, 0x62,
	0x29, 0xf1, 0x56, 0x2f, 0xe2, 0xa0, 0x73, 0xf4, 0xe9, 0xde, 0x83, 0xe6, 0x42, 0x29, 0x21, 0xb9,
	0x9f, 0x66, 0xe0, 0x5b, 0xd9, 0xa8, 0x63, 0xe0, 0xf3, 0xa2, 0x3e, 0x80, 0x5f, 0x03, 0xd7, 0x89,
	0xcc, 0xa1, 0x4a, 0xa4, 0x12, 0xf9, 0x3e, 0x95, 0x89, 0xae, 0x25, 0x4f, 0x54, 0xc9, 0x75, 0x5a,
	0x29, 0x8d, 0x4b, 0xad, 0x2a, 0xe5, 0x8d, 0xed, 0x4b, 0x9c, 0x66, 0x80, 0x5b, 0x08, 0x3e, 0x06,
	0xd8, 0x40, 0xa6, 0xef, 0xe8, 0x7a, 0x03, 0x5c, 0x47, 0x39, 0x54, 0x91, 0xa5, 0x3c, 0xd5, 0x92,
	0xaf, 0x4e, 0x8c, 0xb0, 0xa8, 0x0f, 0xd3, 0x1d, 0xba, 0xea, 0x88, 0x78, 0x97, 0x68, 0xb4, 0x61,
	0xcf, 0xeb, 0x70, 0xbe, 0x7c, 0x58, 0x86, 0xf7, 0xee, 0x3f, 0x6f, 0x84, 0x86, 0x1b, 0xf6, 0xc8,
	0x56, 0x3b, 0xe6, 0xf6, 0xbc, 0x21, 0xea, 0xd3, 0x28, 0x72, 0x1d, 0x23, 0xd8, 0x8b, 0xdb, 0xf7,
	0x2e, 0xd0, 0x80, 0xad, 0x58, 0xe8, 0x0e, 0x66, 0x3f, 0x3d, 0x34, 0xd9, 0xab, 0x03, 0xee, 0x9c,
	0x70, 0x76, 0xd3, 0xf1, 0xf9, 0xba, 0x59, 0x71, 0xdd, 0x7b, 0x9d, 0x9d, 0x10, 0x42, 0xc9, 0xb4,
	0x10, 0x58, 0x69, 0xe1, 0x0b, 0x5e, 0x98, 0x02, 0xce, 0xcd, 0x95, 0x39, 0xfc, 0x51, 0xe4, 0x9d,
	0x39, 0x3f, 0x9c, 0xfe, 0x82, 0xd2, 0x4f, 0x67, 0xbc, 0xb0, 0x28, 0x9a, 0x01, 0xb6, 0xb2, 0x51,
	0xdb, 0x00, 0x37, 0x5d, 0xc0, 0xd7, 0xd9, 0x8f, 0x69, 0x06, 0x71, 0xeb, 0xc7, 0xdb, 0x20, 0x36,
	0xfe, 0xb7, 0x33, 0xb1, 0x97, 0x84, 0xd0, 0x0c, 0x08, 0x69, 0x13, 0x84, 0x6c, 0x03, 0x42, 0x4c,
	0xe6, 0xed, 0xbf, 0x19, 0xec, 0xf6, 0x37, 0x00, 0x00, 0xff, 0xff, 0xa6, 0xc4, 0x42, 0x48, 0xe0,
	0x01, 0x00, 0x00,
}
