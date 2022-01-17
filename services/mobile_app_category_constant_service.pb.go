// Code generated by protoc-gen-go. DO NOT EDIT.
// source: google/ads/googleads/v0/services/mobile_app_category_constant_service.proto

package services

import (
	context "context"
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	resources "github.com/revealbot/google-ads-go/resources"
	_ "google.golang.org/genproto/googleapis/api/annotations"
	grpc "google.golang.org/grpc"
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

// Request message for
// [MobileAppCategoryConstantService.GetMobileAppCategoryConstant][google.ads.googleads.v0.services.MobileAppCategoryConstantService.GetMobileAppCategoryConstant].
type GetMobileAppCategoryConstantRequest struct {
	// Resource name of the mobile app category constant to fetch.
	ResourceName         string   `protobuf:"bytes,1,opt,name=resource_name,json=resourceName,proto3" json:"resource_name,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetMobileAppCategoryConstantRequest) Reset()         { *m = GetMobileAppCategoryConstantRequest{} }
func (m *GetMobileAppCategoryConstantRequest) String() string { return proto.CompactTextString(m) }
func (*GetMobileAppCategoryConstantRequest) ProtoMessage()    {}
func (*GetMobileAppCategoryConstantRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_5e2e2f517adeb75e, []int{0}
}

func (m *GetMobileAppCategoryConstantRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetMobileAppCategoryConstantRequest.Unmarshal(m, b)
}
func (m *GetMobileAppCategoryConstantRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetMobileAppCategoryConstantRequest.Marshal(b, m, deterministic)
}
func (m *GetMobileAppCategoryConstantRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetMobileAppCategoryConstantRequest.Merge(m, src)
}
func (m *GetMobileAppCategoryConstantRequest) XXX_Size() int {
	return xxx_messageInfo_GetMobileAppCategoryConstantRequest.Size(m)
}
func (m *GetMobileAppCategoryConstantRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_GetMobileAppCategoryConstantRequest.DiscardUnknown(m)
}

var xxx_messageInfo_GetMobileAppCategoryConstantRequest proto.InternalMessageInfo

func (m *GetMobileAppCategoryConstantRequest) GetResourceName() string {
	if m != nil {
		return m.ResourceName
	}
	return ""
}

func init() {
	proto.RegisterType((*GetMobileAppCategoryConstantRequest)(nil), "google.ads.googleads.v0.services.GetMobileAppCategoryConstantRequest")
}

func init() {
	proto.RegisterFile("google/ads/googleads/v0/services/mobile_app_category_constant_service.proto", fileDescriptor_5e2e2f517adeb75e)
}

var fileDescriptor_5e2e2f517adeb75e = []byte{
	// 369 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x92, 0x31, 0x4b, 0xc3, 0x40,
	0x1c, 0xc5, 0x49, 0x04, 0xc1, 0xa0, 0x4b, 0x26, 0x29, 0x1d, 0x42, 0x5b, 0x41, 0x1c, 0x2e, 0xa1,
	0x2e, 0x72, 0xea, 0x90, 0x56, 0xa9, 0x28, 0x4a, 0xa9, 0xd0, 0x41, 0x02, 0xe1, 0x9a, 0x1c, 0x21,
	0xd0, 0xdc, 0x9d, 0xf9, 0x5f, 0x0b, 0x22, 0x2e, 0x9d, 0xdd, 0xfc, 0x06, 0x8e, 0x7e, 0x14, 0x57,
	0x67, 0x37, 0x27, 0x3f, 0x85, 0xa4, 0x97, 0x0b, 0x38, 0xc4, 0xb8, 0x3d, 0x2e, 0x2f, 0xbf, 0x77,
	0xff, 0x77, 0x7f, 0xeb, 0x2a, 0xe1, 0x3c, 0x99, 0x53, 0x97, 0xc4, 0xe0, 0x2a, 0x59, 0xa8, 0xa5,
	0xe7, 0x02, 0xcd, 0x97, 0x69, 0x44, 0xc1, 0xcd, 0xf8, 0x2c, 0x9d, 0xd3, 0x90, 0x08, 0x11, 0x46,
	0x44, 0xd2, 0x84, 0xe7, 0x0f, 0x61, 0xc4, 0x19, 0x48, 0xc2, 0x64, 0x58, 0xba, 0x90, 0xc8, 0xb9,
	0xe4, 0xb6, 0xa3, 0x08, 0x88, 0xc4, 0x80, 0x2a, 0x18, 0x5a, 0x7a, 0x48, 0xc3, 0x5a, 0x67, 0x75,
	0x71, 0x39, 0x05, 0xbe, 0xc8, 0x9b, 0xf2, 0x54, 0x4e, 0xab, 0xad, 0x29, 0x22, 0x75, 0x09, 0x63,
	0x5c, 0x12, 0x99, 0x72, 0x06, 0xea, 0x6b, 0xe7, 0xd2, 0xea, 0x8e, 0xa8, 0xbc, 0x5e, 0x63, 0x7c,
	0x21, 0x86, 0x25, 0x64, 0x58, 0x32, 0x26, 0xf4, 0x7e, 0x41, 0x41, 0xda, 0x5d, 0x6b, 0x47, 0x87,
	0x86, 0x8c, 0x64, 0x74, 0xd7, 0x70, 0x8c, 0xfd, 0xad, 0xc9, 0xb6, 0x3e, 0xbc, 0x21, 0x19, 0xed,
	0xaf, 0x4c, 0xcb, 0xa9, 0x25, 0xdd, 0xaa, 0xa9, 0xec, 0x4f, 0xc3, 0x6a, 0xff, 0x95, 0x68, 0x9f,
	0xa3, 0xa6, 0x62, 0xd0, 0x3f, 0x6e, 0xdc, 0x3a, 0xa9, 0xc5, 0x54, 0xed, 0xa1, 0x5a, 0x48, 0xe7,
	0x68, 0xf5, 0xf1, 0xf5, 0x62, 0xf6, 0x6d, 0xaf, 0xa8, 0xfb, 0xf1, 0xd7, 0xe8, 0xa7, 0x59, 0xdd,
	0x5f, 0xe0, 0x1e, 0x3c, 0x0d, 0x9e, 0x4d, 0xab, 0x17, 0xf1, 0xac, 0x71, 0x88, 0xc1, 0x5e, 0x53,
	0x55, 0xe3, 0xe2, 0x81, 0xc6, 0xc6, 0xdd, 0x45, 0x89, 0x4a, 0xf8, 0x9c, 0xb0, 0x04, 0xf1, 0x3c,
	0x71, 0x13, 0xca, 0xd6, 0xcf, 0xa7, 0xd7, 0x42, 0xa4, 0x50, 0xbf, 0x94, 0xc7, 0x5a, 0xbc, 0x9a,
	0x1b, 0x23, 0xdf, 0x7f, 0x33, 0x9d, 0x91, 0x02, 0xfa, 0x31, 0x20, 0x25, 0x0b, 0x35, 0xf5, 0x50,
	0x19, 0x0c, 0xef, 0xda, 0x12, 0xf8, 0x31, 0x04, 0x95, 0x25, 0x98, 0x7a, 0x81, 0xb6, 0x7c, 0x9b,
	0x3d, 0x75, 0x8e, 0xb1, 0x1f, 0x03, 0xc6, 0x95, 0x09, 0xe3, 0xa9, 0x87, 0xb1, 0xb6, 0xcd, 0x36,
	0xd7, 0xf7, 0x3c, 0xfc, 0x09, 0x00, 0x00, 0xff, 0xff, 0xb4, 0x6e, 0xb6, 0x64, 0x3b, 0x03, 0x00,
	0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// MobileAppCategoryConstantServiceClient is the client API for MobileAppCategoryConstantService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type MobileAppCategoryConstantServiceClient interface {
	// Returns the requested mobile app category constant.
	GetMobileAppCategoryConstant(ctx context.Context, in *GetMobileAppCategoryConstantRequest, opts ...grpc.CallOption) (*resources.MobileAppCategoryConstant, error)
}

type mobileAppCategoryConstantServiceClient struct {
	cc *grpc.ClientConn
}

func NewMobileAppCategoryConstantServiceClient(cc *grpc.ClientConn) MobileAppCategoryConstantServiceClient {
	return &mobileAppCategoryConstantServiceClient{cc}
}

func (c *mobileAppCategoryConstantServiceClient) GetMobileAppCategoryConstant(ctx context.Context, in *GetMobileAppCategoryConstantRequest, opts ...grpc.CallOption) (*resources.MobileAppCategoryConstant, error) {
	out := new(resources.MobileAppCategoryConstant)
	err := c.cc.Invoke(ctx, "/google.ads.googleads.v0.services.MobileAppCategoryConstantService/GetMobileAppCategoryConstant", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// MobileAppCategoryConstantServiceServer is the server API for MobileAppCategoryConstantService service.
type MobileAppCategoryConstantServiceServer interface {
	// Returns the requested mobile app category constant.
	GetMobileAppCategoryConstant(context.Context, *GetMobileAppCategoryConstantRequest) (*resources.MobileAppCategoryConstant, error)
}

func RegisterMobileAppCategoryConstantServiceServer(s *grpc.Server, srv MobileAppCategoryConstantServiceServer) {
	s.RegisterService(&_MobileAppCategoryConstantService_serviceDesc, srv)
}

func _MobileAppCategoryConstantService_GetMobileAppCategoryConstant_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetMobileAppCategoryConstantRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MobileAppCategoryConstantServiceServer).GetMobileAppCategoryConstant(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/google.ads.googleads.v0.services.MobileAppCategoryConstantService/GetMobileAppCategoryConstant",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MobileAppCategoryConstantServiceServer).GetMobileAppCategoryConstant(ctx, req.(*GetMobileAppCategoryConstantRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _MobileAppCategoryConstantService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "google.ads.googleads.v0.services.MobileAppCategoryConstantService",
	HandlerType: (*MobileAppCategoryConstantServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetMobileAppCategoryConstant",
			Handler:    _MobileAppCategoryConstantService_GetMobileAppCategoryConstant_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "google/ads/googleads/v0/services/mobile_app_category_constant_service.proto",
}
