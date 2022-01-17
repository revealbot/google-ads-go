// Code generated by protoc-gen-go. DO NOT EDIT.
// source: google/ads/googleads/v9/errors/database_error.proto

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

// Enum describing possible database errors.
type DatabaseErrorEnum_DatabaseError int32

const (
	// Enum unspecified.
	DatabaseErrorEnum_UNSPECIFIED DatabaseErrorEnum_DatabaseError = 0
	// The received error code is not known in this version.
	DatabaseErrorEnum_UNKNOWN DatabaseErrorEnum_DatabaseError = 1
	// Multiple requests were attempting to modify the same resource at once.
	// Please retry the request.
	DatabaseErrorEnum_CONCURRENT_MODIFICATION DatabaseErrorEnum_DatabaseError = 2
)

var DatabaseErrorEnum_DatabaseError_name = map[int32]string{
	0: "UNSPECIFIED",
	1: "UNKNOWN",
	2: "CONCURRENT_MODIFICATION",
}

var DatabaseErrorEnum_DatabaseError_value = map[string]int32{
	"UNSPECIFIED":             0,
	"UNKNOWN":                 1,
	"CONCURRENT_MODIFICATION": 2,
}

func (x DatabaseErrorEnum_DatabaseError) String() string {
	return proto.EnumName(DatabaseErrorEnum_DatabaseError_name, int32(x))
}

func (DatabaseErrorEnum_DatabaseError) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_5af5a1ab517e0304, []int{0, 0}
}

// Container for enum describing possible database errors.
type DatabaseErrorEnum struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *DatabaseErrorEnum) Reset()         { *m = DatabaseErrorEnum{} }
func (m *DatabaseErrorEnum) String() string { return proto.CompactTextString(m) }
func (*DatabaseErrorEnum) ProtoMessage()    {}
func (*DatabaseErrorEnum) Descriptor() ([]byte, []int) {
	return fileDescriptor_5af5a1ab517e0304, []int{0}
}

func (m *DatabaseErrorEnum) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DatabaseErrorEnum.Unmarshal(m, b)
}
func (m *DatabaseErrorEnum) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DatabaseErrorEnum.Marshal(b, m, deterministic)
}
func (m *DatabaseErrorEnum) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DatabaseErrorEnum.Merge(m, src)
}
func (m *DatabaseErrorEnum) XXX_Size() int {
	return xxx_messageInfo_DatabaseErrorEnum.Size(m)
}
func (m *DatabaseErrorEnum) XXX_DiscardUnknown() {
	xxx_messageInfo_DatabaseErrorEnum.DiscardUnknown(m)
}

var xxx_messageInfo_DatabaseErrorEnum proto.InternalMessageInfo

func init() {
	proto.RegisterEnum("google.ads.googleads.v9.errors.DatabaseErrorEnum_DatabaseError", DatabaseErrorEnum_DatabaseError_name, DatabaseErrorEnum_DatabaseError_value)
	proto.RegisterType((*DatabaseErrorEnum)(nil), "google.ads.googleads.v9.errors.DatabaseErrorEnum")
}

func init() {
	proto.RegisterFile("google/ads/googleads/v9/errors/database_error.proto", fileDescriptor_5af5a1ab517e0304)
}

var fileDescriptor_5af5a1ab517e0304 = []byte{
	// 285 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x32, 0x4e, 0xcf, 0xcf, 0x4f,
	0xcf, 0x49, 0xd5, 0x4f, 0x4c, 0x29, 0xd6, 0x87, 0x30, 0x41, 0xac, 0x32, 0x03, 0xfd, 0xd4, 0xa2,
	0xa2, 0xfc, 0xa2, 0x62, 0xfd, 0x94, 0xc4, 0x92, 0xc4, 0xa4, 0xc4, 0xe2, 0xd4, 0x78, 0x30, 0x5f,
	0xaf, 0xa0, 0x28, 0xbf, 0x24, 0x5f, 0x48, 0x0e, 0xa2, 0x52, 0x2f, 0x31, 0xa5, 0x58, 0x0f, 0xae,
	0x49, 0xaf, 0xcc, 0x40, 0x0f, 0xa2, 0x49, 0x29, 0x9e, 0x4b, 0xd0, 0x05, 0xaa, 0xcf, 0x15, 0x24,
	0xe2, 0x9a, 0x57, 0x9a, 0xab, 0xe4, 0xc5, 0xc5, 0x8b, 0x22, 0x28, 0xc4, 0xcf, 0xc5, 0x1d, 0xea,
	0x17, 0x1c, 0xe0, 0xea, 0xec, 0xe9, 0xe6, 0xe9, 0xea, 0x22, 0xc0, 0x20, 0xc4, 0xcd, 0xc5, 0x1e,
	0xea, 0xe7, 0xed, 0xe7, 0x1f, 0xee, 0x27, 0xc0, 0x28, 0x24, 0xcd, 0x25, 0xee, 0xec, 0xef, 0xe7,
	0x1c, 0x1a, 0x14, 0xe4, 0xea, 0x17, 0x12, 0xef, 0xeb, 0xef, 0xe2, 0xe9, 0xe6, 0xe9, 0xec, 0x18,
	0xe2, 0xe9, 0xef, 0x27, 0xc0, 0xe4, 0xf4, 0x96, 0x91, 0x4b, 0x29, 0x39, 0x3f, 0x57, 0x0f, 0xbf,
	0x3b, 0x9c, 0x84, 0x50, 0x2c, 0x0c, 0x00, 0xb9, 0x3d, 0x80, 0x31, 0xca, 0x05, 0xaa, 0x2b, 0x3d,
	0x3f, 0x27, 0x31, 0x2f, 0x5d, 0x2f, 0xbf, 0x28, 0x5d, 0x3f, 0x3d, 0x35, 0x0f, 0xec, 0x33, 0x58,
	0x10, 0x14, 0x64, 0x16, 0xe3, 0x0a, 0x11, 0x6b, 0x08, 0xb5, 0x88, 0x89, 0xd9, 0xdd, 0xd1, 0x71,
	0x15, 0x93, 0x9c, 0x3b, 0xc4, 0x30, 0xc7, 0x94, 0x62, 0x3d, 0x08, 0x13, 0xc4, 0x0a, 0x33, 0xd0,
	0x03, 0x5b, 0x59, 0x7c, 0x0a, 0xa6, 0x20, 0xc6, 0x31, 0xa5, 0x38, 0x06, 0xae, 0x20, 0x26, 0xcc,
	0x20, 0x06, 0xa2, 0xe0, 0x15, 0x93, 0x12, 0x44, 0xd4, 0xca, 0xca, 0x31, 0xa5, 0xd8, 0xca, 0x0a,
	0xae, 0xc4, 0xca, 0x2a, 0xcc, 0xc0, 0xca, 0x0a, 0xa2, 0x28, 0x89, 0x0d, 0xec, 0x3a, 0x63, 0x40,
	0x00, 0x00, 0x00, 0xff, 0xff, 0xe7, 0x01, 0x1f, 0xac, 0xae, 0x01, 0x00, 0x00,
}
