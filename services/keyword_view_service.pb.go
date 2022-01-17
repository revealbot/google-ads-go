// Code generated by protoc-gen-go. DO NOT EDIT.
// source: google/ads/googleads/v9/services/keyword_view_service.proto

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

// Request message for [KeywordViewService.GetKeywordView][google.ads.googleads.v9.services.KeywordViewService.GetKeywordView].
type GetKeywordViewRequest struct {
	// The resource name of the keyword view to fetch.
	ResourceName         string   `protobuf:"bytes,1,opt,name=resource_name,json=resourceName,proto3" json:"resource_name,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetKeywordViewRequest) Reset()         { *m = GetKeywordViewRequest{} }
func (m *GetKeywordViewRequest) String() string { return proto.CompactTextString(m) }
func (*GetKeywordViewRequest) ProtoMessage()    {}
func (*GetKeywordViewRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_f154d6d131e5f762, []int{0}
}

func (m *GetKeywordViewRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetKeywordViewRequest.Unmarshal(m, b)
}
func (m *GetKeywordViewRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetKeywordViewRequest.Marshal(b, m, deterministic)
}
func (m *GetKeywordViewRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetKeywordViewRequest.Merge(m, src)
}
func (m *GetKeywordViewRequest) XXX_Size() int {
	return xxx_messageInfo_GetKeywordViewRequest.Size(m)
}
func (m *GetKeywordViewRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_GetKeywordViewRequest.DiscardUnknown(m)
}

var xxx_messageInfo_GetKeywordViewRequest proto.InternalMessageInfo

func (m *GetKeywordViewRequest) GetResourceName() string {
	if m != nil {
		return m.ResourceName
	}
	return ""
}

func init() {
	proto.RegisterType((*GetKeywordViewRequest)(nil), "google.ads.googleads.v9.services.GetKeywordViewRequest")
}

func init() {
	proto.RegisterFile("google/ads/googleads/v9/services/keyword_view_service.proto", fileDescriptor_f154d6d131e5f762)
}

var fileDescriptor_f154d6d131e5f762 = []byte{
	// 364 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x84, 0x92, 0xb1, 0x4a, 0xfb, 0x50,
	0x14, 0xc6, 0x49, 0xfe, 0xf0, 0x07, 0x83, 0x3a, 0x04, 0x44, 0x29, 0x0e, 0xa5, 0x76, 0x90, 0x0e,
	0xe7, 0x06, 0x15, 0x85, 0x5b, 0x1d, 0xd2, 0xa5, 0x82, 0x20, 0xa5, 0x42, 0x06, 0x09, 0x94, 0xd8,
	0x1c, 0x42, 0xb0, 0xc9, 0xad, 0xf7, 0xa4, 0x29, 0x22, 0x2e, 0xbe, 0x82, 0x6f, 0xe0, 0xe8, 0xe6,
	0x5b, 0x88, 0xab, 0xaf, 0xe0, 0xe4, 0x43, 0x88, 0xa4, 0xb7, 0x37, 0xb4, 0x6a, 0xe9, 0xf6, 0x71,
	0xf2, 0xfd, 0xbe, 0x9c, 0xf3, 0x25, 0x56, 0x33, 0x12, 0x22, 0x1a, 0x20, 0x0b, 0x42, 0x62, 0x4a,
	0x16, 0x2a, 0x77, 0x18, 0xa1, 0xcc, 0xe3, 0x3e, 0x12, 0xbb, 0xc6, 0xdb, 0xb1, 0x90, 0x61, 0x2f,
	0x8f, 0x71, 0xdc, 0x9b, 0x4e, 0x61, 0x28, 0x45, 0x26, 0xec, 0xaa, 0x22, 0x20, 0x08, 0x09, 0x4a,
	0x18, 0x72, 0x07, 0x34, 0x5c, 0x39, 0x58, 0x14, 0x2f, 0x91, 0xc4, 0x48, 0xfe, 0xcc, 0x57, 0xb9,
	0x95, 0x6d, 0x4d, 0x0d, 0x63, 0x16, 0xa4, 0xa9, 0xc8, 0x82, 0x2c, 0x16, 0x29, 0xa9, 0xa7, 0xb5,
	0x63, 0x6b, 0xa3, 0x8d, 0xd9, 0x99, 0xc2, 0xbc, 0x18, 0xc7, 0x5d, 0xbc, 0x19, 0x21, 0x65, 0xf6,
	0x8e, 0xb5, 0xa6, 0x63, 0x7b, 0x69, 0x90, 0xe0, 0x96, 0x51, 0x35, 0x76, 0x57, 0xba, 0xab, 0x7a,
	0x78, 0x1e, 0x24, 0xb8, 0xf7, 0x6a, 0x58, 0xf6, 0x0c, 0x7b, 0xa1, 0x36, 0xb5, 0x5f, 0x0c, 0x6b,
	0x7d, 0x3e, 0xd5, 0x3e, 0x82, 0x65, 0xe7, 0xc1, 0x9f, 0x7b, 0x54, 0x60, 0x21, 0x58, 0x5e, 0x0d,
	0x33, 0x58, 0xed, 0xf0, 0xe1, 0xfd, 0xe3, 0xd1, 0x74, 0x6c, 0x28, 0x8a, 0xb9, 0x9b, 0x3b, 0xe1,
	0xa4, 0x3f, 0xa2, 0x4c, 0x24, 0x28, 0x89, 0x35, 0x74, 0x53, 0x05, 0x43, 0xac, 0x71, 0xdf, 0xfa,
	0x32, 0xac, 0x7a, 0x5f, 0x24, 0x4b, 0xd7, 0x6c, 0x6d, 0xfe, 0x3e, 0xb8, 0x53, 0x54, 0xd9, 0x31,
	0x2e, 0x4f, 0xa7, 0x70, 0x24, 0x06, 0x41, 0x1a, 0x81, 0x90, 0x11, 0x8b, 0x30, 0x9d, 0x14, 0xad,
	0x3f, 0xd8, 0x30, 0xa6, 0xc5, 0xbf, 0x47, 0x53, 0x8b, 0x27, 0xf3, 0x5f, 0xdb, 0x75, 0x9f, 0xcd,
	0x6a, 0x5b, 0x05, 0xba, 0x21, 0x81, 0x92, 0x85, 0xf2, 0x1c, 0x98, 0xbe, 0x98, 0xde, 0xb4, 0xc5,
	0x77, 0x43, 0xf2, 0x4b, 0x8b, 0xef, 0x39, 0xbe, 0xb6, 0x7c, 0x9a, 0x75, 0x35, 0xe7, 0xdc, 0x0d,
	0x89, 0xf3, 0xd2, 0xc4, 0xb9, 0xe7, 0x70, 0xae, 0x6d, 0x57, 0xff, 0x27, 0x7b, 0xee, 0x7f, 0x07,
	0x00, 0x00, 0xff, 0xff, 0xca, 0xa5, 0xfe, 0xd8, 0xc5, 0x02, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// KeywordViewServiceClient is the client API for KeywordViewService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type KeywordViewServiceClient interface {
	// Returns the requested keyword view in full detail.
	GetKeywordView(ctx context.Context, in *GetKeywordViewRequest, opts ...grpc.CallOption) (*resources.KeywordView, error)
}

type keywordViewServiceClient struct {
	cc *grpc.ClientConn
}

func NewKeywordViewServiceClient(cc *grpc.ClientConn) KeywordViewServiceClient {
	return &keywordViewServiceClient{cc}
}

func (c *keywordViewServiceClient) GetKeywordView(ctx context.Context, in *GetKeywordViewRequest, opts ...grpc.CallOption) (*resources.KeywordView, error) {
	out := new(resources.KeywordView)
	err := c.cc.Invoke(ctx, "/google.ads.googleads.v9.services.KeywordViewService/GetKeywordView", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// KeywordViewServiceServer is the server API for KeywordViewService service.
type KeywordViewServiceServer interface {
	// Returns the requested keyword view in full detail.
	GetKeywordView(context.Context, *GetKeywordViewRequest) (*resources.KeywordView, error)
}

func RegisterKeywordViewServiceServer(s *grpc.Server, srv KeywordViewServiceServer) {
	s.RegisterService(&_KeywordViewService_serviceDesc, srv)
}

func _KeywordViewService_GetKeywordView_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetKeywordViewRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(KeywordViewServiceServer).GetKeywordView(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/google.ads.googleads.v9.services.KeywordViewService/GetKeywordView",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(KeywordViewServiceServer).GetKeywordView(ctx, req.(*GetKeywordViewRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _KeywordViewService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "google.ads.googleads.v9.services.KeywordViewService",
	HandlerType: (*KeywordViewServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetKeywordView",
			Handler:    _KeywordViewService_GetKeywordView_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "google/ads/googleads/v9/services/keyword_view_service.proto",
}
