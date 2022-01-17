// Code generated by protoc-gen-go. DO NOT EDIT.
// source: google/ads/googleads/v9/enums/app_payment_model_type.proto

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

// Enum describing possible app payment models.
type AppPaymentModelTypeEnum_AppPaymentModelType int32

const (
	// Not specified.
	AppPaymentModelTypeEnum_UNSPECIFIED AppPaymentModelTypeEnum_AppPaymentModelType = 0
	// Used for return value only. Represents value unknown in this version.
	AppPaymentModelTypeEnum_UNKNOWN AppPaymentModelTypeEnum_AppPaymentModelType = 1
	// Represents paid-for apps.
	AppPaymentModelTypeEnum_PAID AppPaymentModelTypeEnum_AppPaymentModelType = 30
)

var AppPaymentModelTypeEnum_AppPaymentModelType_name = map[int32]string{
	0:  "UNSPECIFIED",
	1:  "UNKNOWN",
	30: "PAID",
}

var AppPaymentModelTypeEnum_AppPaymentModelType_value = map[string]int32{
	"UNSPECIFIED": 0,
	"UNKNOWN":     1,
	"PAID":        30,
}

func (x AppPaymentModelTypeEnum_AppPaymentModelType) String() string {
	return proto.EnumName(AppPaymentModelTypeEnum_AppPaymentModelType_name, int32(x))
}

func (AppPaymentModelTypeEnum_AppPaymentModelType) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_328f59e7a6a2ec87, []int{0, 0}
}

// Represents a criterion for targeting paid apps.
type AppPaymentModelTypeEnum struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *AppPaymentModelTypeEnum) Reset()         { *m = AppPaymentModelTypeEnum{} }
func (m *AppPaymentModelTypeEnum) String() string { return proto.CompactTextString(m) }
func (*AppPaymentModelTypeEnum) ProtoMessage()    {}
func (*AppPaymentModelTypeEnum) Descriptor() ([]byte, []int) {
	return fileDescriptor_328f59e7a6a2ec87, []int{0}
}

func (m *AppPaymentModelTypeEnum) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_AppPaymentModelTypeEnum.Unmarshal(m, b)
}
func (m *AppPaymentModelTypeEnum) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_AppPaymentModelTypeEnum.Marshal(b, m, deterministic)
}
func (m *AppPaymentModelTypeEnum) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AppPaymentModelTypeEnum.Merge(m, src)
}
func (m *AppPaymentModelTypeEnum) XXX_Size() int {
	return xxx_messageInfo_AppPaymentModelTypeEnum.Size(m)
}
func (m *AppPaymentModelTypeEnum) XXX_DiscardUnknown() {
	xxx_messageInfo_AppPaymentModelTypeEnum.DiscardUnknown(m)
}

var xxx_messageInfo_AppPaymentModelTypeEnum proto.InternalMessageInfo

func init() {
	proto.RegisterEnum("google.ads.googleads.v9.enums.AppPaymentModelTypeEnum_AppPaymentModelType", AppPaymentModelTypeEnum_AppPaymentModelType_name, AppPaymentModelTypeEnum_AppPaymentModelType_value)
	proto.RegisterType((*AppPaymentModelTypeEnum)(nil), "google.ads.googleads.v9.enums.AppPaymentModelTypeEnum")
}

func init() {
	proto.RegisterFile("google/ads/googleads/v9/enums/app_payment_model_type.proto", fileDescriptor_328f59e7a6a2ec87)
}

var fileDescriptor_328f59e7a6a2ec87 = []byte{
	// 283 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xb2, 0x4a, 0xcf, 0xcf, 0x4f,
	0xcf, 0x49, 0xd5, 0x4f, 0x4c, 0x29, 0xd6, 0x87, 0x30, 0x41, 0xac, 0x32, 0x03, 0xfd, 0xd4, 0xbc,
	0xd2, 0xdc, 0x62, 0xfd, 0xc4, 0x82, 0x82, 0xf8, 0x82, 0xc4, 0xca, 0xdc, 0xd4, 0xbc, 0x92, 0xf8,
	0xdc, 0xfc, 0x94, 0xd4, 0x9c, 0xf8, 0x92, 0xca, 0x82, 0x54, 0xbd, 0x82, 0xa2, 0xfc, 0x92, 0x7c,
	0x21, 0x59, 0x88, 0x06, 0xbd, 0xc4, 0x94, 0x62, 0x3d, 0xb8, 0x5e, 0xbd, 0x32, 0x03, 0x3d, 0xb0,
	0x5e, 0xa5, 0x08, 0x2e, 0x71, 0xc7, 0x82, 0x82, 0x00, 0x88, 0x6e, 0x5f, 0x90, 0xe6, 0x90, 0xca,
	0x82, 0x54, 0xd7, 0xbc, 0xd2, 0x5c, 0x25, 0x5b, 0x2e, 0x61, 0x2c, 0x52, 0x42, 0xfc, 0x5c, 0xdc,
	0xa1, 0x7e, 0xc1, 0x01, 0xae, 0xce, 0x9e, 0x6e, 0x9e, 0xae, 0x2e, 0x02, 0x0c, 0x42, 0xdc, 0x5c,
	0xec, 0xa1, 0x7e, 0xde, 0x7e, 0xfe, 0xe1, 0x7e, 0x02, 0x8c, 0x42, 0x1c, 0x5c, 0x2c, 0x01, 0x8e,
	0x9e, 0x2e, 0x02, 0x72, 0x4e, 0x6f, 0x19, 0xb9, 0x14, 0x93, 0xf3, 0x73, 0xf5, 0xf0, 0xda, 0xef,
	0x24, 0x81, 0xc5, 0x8a, 0x00, 0x90, 0xc3, 0x03, 0x18, 0xa3, 0x9c, 0xa0, 0x5a, 0xd3, 0xf3, 0x73,
	0x12, 0xf3, 0xd2, 0xf5, 0xf2, 0x8b, 0xd2, 0xf5, 0xd3, 0x53, 0xf3, 0xc0, 0xde, 0x82, 0x05, 0x43,
	0x41, 0x66, 0x31, 0x8e, 0x50, 0xb1, 0x06, 0x93, 0x8b, 0x98, 0x98, 0xdd, 0x1d, 0x1d, 0x57, 0x31,
	0xc9, 0xba, 0x43, 0x8c, 0x72, 0x4c, 0x29, 0xd6, 0x83, 0x30, 0x41, 0xac, 0x30, 0x03, 0x3d, 0x90,
	0x4f, 0x8b, 0x4f, 0xc1, 0xe4, 0x63, 0x1c, 0x53, 0x8a, 0x63, 0xe0, 0xf2, 0x31, 0x61, 0x06, 0x31,
	0x60, 0xf9, 0x57, 0x4c, 0x8a, 0x10, 0x41, 0x2b, 0x2b, 0xc7, 0x94, 0x62, 0x2b, 0x2b, 0xb8, 0x0a,
	0x2b, 0xab, 0x30, 0x03, 0x2b, 0x2b, 0xb0, 0x9a, 0x24, 0x36, 0xb0, 0xc3, 0x8c, 0x01, 0x01, 0x00,
	0x00, 0xff, 0xff, 0x2d, 0xef, 0x1e, 0xf9, 0xad, 0x01, 0x00, 0x00,
}
