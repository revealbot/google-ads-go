// Code generated by protoc-gen-go. DO NOT EDIT.
// source: google/ads/googleads/v9/common/real_time_bidding_setting.proto

package common

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	wrappers "github.com/golang/protobuf/ptypes/wrappers"
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

// Settings for Real-Time Bidding, a feature only available for campaigns
// targeting the Ad Exchange network.
type RealTimeBiddingSetting struct {
	// Whether the campaign is opted in to real-time bidding.
	OptIn                *wrappers.BoolValue `protobuf:"bytes,1,opt,name=opt_in,json=optIn,proto3" json:"opt_in,omitempty"`
	XXX_NoUnkeyedLiteral struct{}            `json:"-"`
	XXX_unrecognized     []byte              `json:"-"`
	XXX_sizecache        int32               `json:"-"`
}

func (m *RealTimeBiddingSetting) Reset()         { *m = RealTimeBiddingSetting{} }
func (m *RealTimeBiddingSetting) String() string { return proto.CompactTextString(m) }
func (*RealTimeBiddingSetting) ProtoMessage()    {}
func (*RealTimeBiddingSetting) Descriptor() ([]byte, []int) {
	return fileDescriptor_b54af3522cdfdf23, []int{0}
}

func (m *RealTimeBiddingSetting) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RealTimeBiddingSetting.Unmarshal(m, b)
}
func (m *RealTimeBiddingSetting) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RealTimeBiddingSetting.Marshal(b, m, deterministic)
}
func (m *RealTimeBiddingSetting) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RealTimeBiddingSetting.Merge(m, src)
}
func (m *RealTimeBiddingSetting) XXX_Size() int {
	return xxx_messageInfo_RealTimeBiddingSetting.Size(m)
}
func (m *RealTimeBiddingSetting) XXX_DiscardUnknown() {
	xxx_messageInfo_RealTimeBiddingSetting.DiscardUnknown(m)
}

var xxx_messageInfo_RealTimeBiddingSetting proto.InternalMessageInfo

func (m *RealTimeBiddingSetting) GetOptIn() *wrappers.BoolValue {
	if m != nil {
		return m.OptIn
	}
	return nil
}

func init() {
	proto.RegisterType((*RealTimeBiddingSetting)(nil), "google.ads.googleads.v9.common.RealTimeBiddingSetting")
}

func init() {
	proto.RegisterFile("google/ads/googleads/v9/common/real_time_bidding_setting.proto", fileDescriptor_b54af3522cdfdf23)
}

var fileDescriptor_b54af3522cdfdf23 = []byte{
	// 290 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x7c, 0x90, 0xc1, 0x4a, 0xf4, 0x30,
	0x14, 0x85, 0xe9, 0xfc, 0xfc, 0xb3, 0xa8, 0xbb, 0x59, 0x88, 0x8c, 0x30, 0x48, 0x57, 0xae, 0x92,
	0xaa, 0xbb, 0x08, 0x42, 0xab, 0x30, 0x88, 0x9b, 0x61, 0x94, 0x2e, 0xa4, 0x50, 0xd2, 0x49, 0x0c,
	0x81, 0x34, 0x37, 0x34, 0x99, 0xf1, 0x7d, 0x5c, 0xfa, 0x28, 0x3e, 0x8a, 0x0f, 0xe0, 0x5a, 0xda,
	0xdb, 0x76, 0xa5, 0xae, 0x72, 0xc8, 0xfd, 0xce, 0x39, 0x97, 0x1b, 0xdf, 0x28, 0x00, 0x65, 0x24,
	0xe5, 0xc2, 0x53, 0x94, 0x9d, 0x3a, 0xa4, 0x74, 0x07, 0x4d, 0x03, 0x96, 0xb6, 0x92, 0x9b, 0x2a,
	0xe8, 0x46, 0x56, 0xb5, 0x16, 0x42, 0x5b, 0x55, 0x79, 0x19, 0x82, 0xb6, 0x8a, 0xb8, 0x16, 0x02,
	0x2c, 0x56, 0x68, 0x22, 0x5c, 0x78, 0x32, 0xf9, 0xc9, 0x21, 0x25, 0xe8, 0x5f, 0x0e, 0x73, 0xda,
	0xd3, 0xf5, 0xfe, 0x85, 0xbe, 0xb6, 0xdc, 0x39, 0xd9, 0x7a, 0xf4, 0x27, 0x0f, 0xf1, 0xf1, 0x56,
	0x72, 0xf3, 0xa4, 0x1b, 0x99, 0x63, 0xc1, 0x23, 0xe6, 0x2f, 0x2e, 0xe2, 0x39, 0xb8, 0x50, 0x69,
	0x7b, 0x12, 0x9d, 0x45, 0xe7, 0x47, 0x97, 0xcb, 0x21, 0x9f, 0x8c, 0x51, 0x24, 0x07, 0x30, 0x05,
	0x37, 0x7b, 0xb9, 0xfd, 0x0f, 0x2e, 0xdc, 0xdb, 0xfc, 0x2b, 0x8a, 0x93, 0x1d, 0x34, 0xe4, 0xef,
	0x9d, 0xf2, 0xd3, 0x9f, 0x1b, 0x37, 0x5d, 0xee, 0x26, 0x7a, 0xbe, 0x1b, 0xec, 0x0a, 0x0c, 0xb7,
	0x8a, 0x40, 0xab, 0xa8, 0x92, 0xb6, 0x6f, 0x1d, 0x4f, 0xe4, 0xb4, 0xff, 0xed, 0x62, 0xd7, 0xf8,
	0xbc, 0xcd, 0xfe, 0xad, 0xb3, 0xec, 0x7d, 0xb6, 0x5a, 0x63, 0x58, 0x26, 0x3c, 0x41, 0xd9, 0xa9,
	0x22, 0x25, 0xb7, 0x3d, 0xf6, 0x31, 0x02, 0x65, 0x26, 0x7c, 0x39, 0x01, 0x65, 0x91, 0x96, 0x08,
	0x7c, 0xce, 0x12, 0xfc, 0x65, 0x2c, 0x13, 0x9e, 0xb1, 0x09, 0x61, 0xac, 0x48, 0x19, 0x43, 0xa8,
	0x9e, 0xf7, 0xdb, 0x5d, 0x7d, 0x07, 0x00, 0x00, 0xff, 0xff, 0x15, 0x19, 0xa5, 0x86, 0xce, 0x01,
	0x00, 0x00,
}
