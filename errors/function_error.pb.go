// Code generated by protoc-gen-go. DO NOT EDIT.
// source: google/ads/googleads/v9/errors/function_error.proto

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

// Enum describing possible function errors.
type FunctionErrorEnum_FunctionError int32

const (
	// Enum unspecified.
	FunctionErrorEnum_UNSPECIFIED FunctionErrorEnum_FunctionError = 0
	// The received error code is not known in this version.
	FunctionErrorEnum_UNKNOWN FunctionErrorEnum_FunctionError = 1
	// The format of the function is not recognized as a supported function
	// format.
	FunctionErrorEnum_INVALID_FUNCTION_FORMAT FunctionErrorEnum_FunctionError = 2
	// Operand data types do not match.
	FunctionErrorEnum_DATA_TYPE_MISMATCH FunctionErrorEnum_FunctionError = 3
	// The operands cannot be used together in a conjunction.
	FunctionErrorEnum_INVALID_CONJUNCTION_OPERANDS FunctionErrorEnum_FunctionError = 4
	// Invalid numer of Operands.
	FunctionErrorEnum_INVALID_NUMBER_OF_OPERANDS FunctionErrorEnum_FunctionError = 5
	// Operand Type not supported.
	FunctionErrorEnum_INVALID_OPERAND_TYPE FunctionErrorEnum_FunctionError = 6
	// Operator not supported.
	FunctionErrorEnum_INVALID_OPERATOR FunctionErrorEnum_FunctionError = 7
	// Request context type not supported.
	FunctionErrorEnum_INVALID_REQUEST_CONTEXT_TYPE FunctionErrorEnum_FunctionError = 8
	// The matching function is not allowed for call placeholders
	FunctionErrorEnum_INVALID_FUNCTION_FOR_CALL_PLACEHOLDER FunctionErrorEnum_FunctionError = 9
	// The matching function is not allowed for the specified placeholder
	FunctionErrorEnum_INVALID_FUNCTION_FOR_PLACEHOLDER FunctionErrorEnum_FunctionError = 10
	// Invalid operand.
	FunctionErrorEnum_INVALID_OPERAND FunctionErrorEnum_FunctionError = 11
	// Missing value for the constant operand.
	FunctionErrorEnum_MISSING_CONSTANT_OPERAND_VALUE FunctionErrorEnum_FunctionError = 12
	// The value of the constant operand is invalid.
	FunctionErrorEnum_INVALID_CONSTANT_OPERAND_VALUE FunctionErrorEnum_FunctionError = 13
	// Invalid function nesting.
	FunctionErrorEnum_INVALID_NESTING FunctionErrorEnum_FunctionError = 14
	// The Feed ID was different from another Feed ID in the same function.
	FunctionErrorEnum_MULTIPLE_FEED_IDS_NOT_SUPPORTED FunctionErrorEnum_FunctionError = 15
	// The matching function is invalid for use with a feed with a fixed schema.
	FunctionErrorEnum_INVALID_FUNCTION_FOR_FEED_WITH_FIXED_SCHEMA FunctionErrorEnum_FunctionError = 16
	// Invalid attribute name.
	FunctionErrorEnum_INVALID_ATTRIBUTE_NAME FunctionErrorEnum_FunctionError = 17
)

var FunctionErrorEnum_FunctionError_name = map[int32]string{
	0:  "UNSPECIFIED",
	1:  "UNKNOWN",
	2:  "INVALID_FUNCTION_FORMAT",
	3:  "DATA_TYPE_MISMATCH",
	4:  "INVALID_CONJUNCTION_OPERANDS",
	5:  "INVALID_NUMBER_OF_OPERANDS",
	6:  "INVALID_OPERAND_TYPE",
	7:  "INVALID_OPERATOR",
	8:  "INVALID_REQUEST_CONTEXT_TYPE",
	9:  "INVALID_FUNCTION_FOR_CALL_PLACEHOLDER",
	10: "INVALID_FUNCTION_FOR_PLACEHOLDER",
	11: "INVALID_OPERAND",
	12: "MISSING_CONSTANT_OPERAND_VALUE",
	13: "INVALID_CONSTANT_OPERAND_VALUE",
	14: "INVALID_NESTING",
	15: "MULTIPLE_FEED_IDS_NOT_SUPPORTED",
	16: "INVALID_FUNCTION_FOR_FEED_WITH_FIXED_SCHEMA",
	17: "INVALID_ATTRIBUTE_NAME",
}

var FunctionErrorEnum_FunctionError_value = map[string]int32{
	"UNSPECIFIED":                                 0,
	"UNKNOWN":                                     1,
	"INVALID_FUNCTION_FORMAT":                     2,
	"DATA_TYPE_MISMATCH":                          3,
	"INVALID_CONJUNCTION_OPERANDS":                4,
	"INVALID_NUMBER_OF_OPERANDS":                  5,
	"INVALID_OPERAND_TYPE":                        6,
	"INVALID_OPERATOR":                            7,
	"INVALID_REQUEST_CONTEXT_TYPE":                8,
	"INVALID_FUNCTION_FOR_CALL_PLACEHOLDER":       9,
	"INVALID_FUNCTION_FOR_PLACEHOLDER":            10,
	"INVALID_OPERAND":                             11,
	"MISSING_CONSTANT_OPERAND_VALUE":              12,
	"INVALID_CONSTANT_OPERAND_VALUE":              13,
	"INVALID_NESTING":                             14,
	"MULTIPLE_FEED_IDS_NOT_SUPPORTED":             15,
	"INVALID_FUNCTION_FOR_FEED_WITH_FIXED_SCHEMA": 16,
	"INVALID_ATTRIBUTE_NAME":                      17,
}

func (x FunctionErrorEnum_FunctionError) String() string {
	return proto.EnumName(FunctionErrorEnum_FunctionError_name, int32(x))
}

func (FunctionErrorEnum_FunctionError) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_855fd6ada22aea64, []int{0, 0}
}

// Container for enum describing possible function errors.
type FunctionErrorEnum struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *FunctionErrorEnum) Reset()         { *m = FunctionErrorEnum{} }
func (m *FunctionErrorEnum) String() string { return proto.CompactTextString(m) }
func (*FunctionErrorEnum) ProtoMessage()    {}
func (*FunctionErrorEnum) Descriptor() ([]byte, []int) {
	return fileDescriptor_855fd6ada22aea64, []int{0}
}

func (m *FunctionErrorEnum) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_FunctionErrorEnum.Unmarshal(m, b)
}
func (m *FunctionErrorEnum) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_FunctionErrorEnum.Marshal(b, m, deterministic)
}
func (m *FunctionErrorEnum) XXX_Merge(src proto.Message) {
	xxx_messageInfo_FunctionErrorEnum.Merge(m, src)
}
func (m *FunctionErrorEnum) XXX_Size() int {
	return xxx_messageInfo_FunctionErrorEnum.Size(m)
}
func (m *FunctionErrorEnum) XXX_DiscardUnknown() {
	xxx_messageInfo_FunctionErrorEnum.DiscardUnknown(m)
}

var xxx_messageInfo_FunctionErrorEnum proto.InternalMessageInfo

func init() {
	proto.RegisterEnum("google.ads.googleads.v9.errors.FunctionErrorEnum_FunctionError", FunctionErrorEnum_FunctionError_name, FunctionErrorEnum_FunctionError_value)
	proto.RegisterType((*FunctionErrorEnum)(nil), "google.ads.googleads.v9.errors.FunctionErrorEnum")
}

func init() {
	proto.RegisterFile("google/ads/googleads/v9/errors/function_error.proto", fileDescriptor_855fd6ada22aea64)
}

var fileDescriptor_855fd6ada22aea64 = []byte{
	// 525 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x7c, 0x92, 0xc1, 0x6e, 0xd3, 0x4e,
	0x10, 0xc6, 0xff, 0x49, 0xf3, 0x6f, 0x61, 0x42, 0xc9, 0x76, 0xa9, 0x0a, 0x2a, 0x28, 0x54, 0x01,
	0x0e, 0x08, 0xc9, 0x8e, 0xd4, 0x9b, 0x39, 0x6d, 0xec, 0x71, 0xb2, 0x60, 0xaf, 0x8d, 0x77, 0x9d,
	0x16, 0x14, 0x69, 0x15, 0x9a, 0x60, 0x55, 0x6a, 0xe3, 0x2a, 0x6e, 0xfb, 0x40, 0x1c, 0xe1, 0x29,
	0xb8, 0xf2, 0x28, 0x48, 0xbc, 0x03, 0xb2, 0x1d, 0x87, 0x04, 0x05, 0x4e, 0x1e, 0xcf, 0xfc, 0xbe,
	0xfd, 0x66, 0x57, 0x1f, 0x1c, 0x27, 0x69, 0x9a, 0x5c, 0x4c, 0xcd, 0xf1, 0x24, 0x33, 0xcb, 0x32,
	0xaf, 0x6e, 0xbb, 0xe6, 0x74, 0x3e, 0x4f, 0xe7, 0x99, 0xf9, 0xe9, 0x66, 0x76, 0x76, 0x7d, 0x9e,
	0xce, 0x74, 0xf1, 0x6f, 0x5c, 0xcd, 0xd3, 0xeb, 0x94, 0xb6, 0x4b, 0xd2, 0x18, 0x4f, 0x32, 0x63,
	0x29, 0x32, 0x6e, 0xbb, 0x46, 0x29, 0xea, 0x7c, 0x6b, 0xc0, 0x9e, 0xbb, 0x10, 0x62, 0xde, 0xc2,
	0xd9, 0xcd, 0x65, 0xe7, 0x6b, 0x03, 0x76, 0xd7, 0xba, 0xb4, 0x05, 0xcd, 0x58, 0xc8, 0x10, 0x6d,
	0xee, 0x72, 0x74, 0xc8, 0x7f, 0xb4, 0x09, 0x3b, 0xb1, 0x78, 0x2b, 0x82, 0x13, 0x41, 0x6a, 0xf4,
	0x31, 0x3c, 0xe4, 0x62, 0xc8, 0x3c, 0xee, 0x68, 0x37, 0x16, 0xb6, 0xe2, 0x81, 0xd0, 0x6e, 0x10,
	0xf9, 0x4c, 0x91, 0x3a, 0x3d, 0x00, 0xea, 0x30, 0xc5, 0xb4, 0x7a, 0x1f, 0xa2, 0xf6, 0xb9, 0xf4,
	0x99, 0xb2, 0x07, 0x64, 0x8b, 0x1e, 0xc1, 0x93, 0x4a, 0x64, 0x07, 0xe2, 0x4d, 0xa5, 0x0b, 0x42,
	0x8c, 0x98, 0x70, 0x24, 0x69, 0xd0, 0x36, 0x1c, 0x56, 0x84, 0x88, 0xfd, 0x1e, 0x46, 0x3a, 0x70,
	0x7f, 0xcf, 0xff, 0xa7, 0x8f, 0x60, 0xbf, 0x9a, 0x2f, 0xba, 0x85, 0x09, 0xd9, 0xa6, 0xfb, 0x40,
	0xd6, 0x26, 0x2a, 0x88, 0xc8, 0xce, 0xaa, 0x63, 0x84, 0xef, 0x62, 0x94, 0x2a, 0x77, 0x56, 0x78,
	0xaa, 0x4a, 0xdd, 0x1d, 0xfa, 0x12, 0x5e, 0x6c, 0xba, 0x88, 0xb6, 0x99, 0xe7, 0xe9, 0xd0, 0x63,
	0x36, 0x0e, 0x02, 0xcf, 0xc1, 0x88, 0xdc, 0xa5, 0xcf, 0xe1, 0x68, 0x23, 0xba, 0x4a, 0x01, 0x7d,
	0x00, 0xad, 0x3f, 0x56, 0x24, 0x4d, 0xda, 0x81, 0xb6, 0xcf, 0xa5, 0xe4, 0xa2, 0x9f, 0xfb, 0x4b,
	0xc5, 0x84, 0x5a, 0x5e, 0x60, 0xc8, 0xbc, 0x18, 0xc9, 0xbd, 0x9c, 0x59, 0x79, 0x9d, 0x4d, 0xcc,
	0xee, 0xea, 0xe1, 0x02, 0xa5, 0xe2, 0xa2, 0x4f, 0xee, 0xd3, 0x67, 0xf0, 0xd4, 0x8f, 0x3d, 0xc5,
	0x43, 0x0f, 0xb5, 0x8b, 0xe8, 0x68, 0xee, 0x48, 0x2d, 0x02, 0xa5, 0x65, 0x1c, 0x86, 0x41, 0xa4,
	0xd0, 0x21, 0x2d, 0x6a, 0xc2, 0xab, 0x8d, 0xcb, 0x17, 0x82, 0x13, 0xae, 0x06, 0xda, 0xe5, 0xa7,
	0xe8, 0x68, 0x69, 0x0f, 0xd0, 0x67, 0x84, 0xd0, 0x43, 0x38, 0xa8, 0x04, 0x4c, 0xa9, 0x88, 0xf7,
	0x62, 0x85, 0x5a, 0x30, 0x1f, 0xc9, 0x5e, 0xef, 0x67, 0x0d, 0x3a, 0x67, 0xe9, 0xa5, 0xf1, 0xef,
	0xa8, 0xf5, 0xe8, 0x5a, 0xa2, 0xc2, 0x3c, 0x9e, 0x61, 0xed, 0x83, 0xb3, 0x50, 0x25, 0xe9, 0xc5,
	0x78, 0x96, 0x18, 0xe9, 0x3c, 0x31, 0x93, 0xe9, 0xac, 0x08, 0x6f, 0x95, 0xf2, 0xab, 0xf3, 0xec,
	0x6f, 0xa1, 0x7f, 0x5d, 0x7e, 0x3e, 0xd7, 0xb7, 0xfa, 0x8c, 0x7d, 0xa9, 0xb7, 0xfb, 0xe5, 0x61,
	0x6c, 0x92, 0x19, 0x65, 0x99, 0x57, 0xc3, 0xae, 0x51, 0x58, 0x66, 0xdf, 0x2b, 0x60, 0xc4, 0x26,
	0xd9, 0x68, 0x09, 0x8c, 0x86, 0xdd, 0x51, 0x09, 0xfc, 0xa8, 0x77, 0xca, 0xae, 0x65, 0xb1, 0x49,
	0x66, 0x59, 0x4b, 0xc4, 0xb2, 0x86, 0x5d, 0xcb, 0x2a, 0xa1, 0x8f, 0xdb, 0xc5, 0x76, 0xc7, 0xbf,
	0x02, 0x00, 0x00, 0xff, 0xff, 0x8a, 0x50, 0x2b, 0x06, 0x91, 0x03, 0x00, 0x00,
}
