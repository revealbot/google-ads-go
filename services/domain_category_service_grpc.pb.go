// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package services

import (
	context "context"
	resources "github.com/revealbot/google-ads-go/resources"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

// DomainCategoryServiceClient is the client API for DomainCategoryService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type DomainCategoryServiceClient interface {
	// Returns the requested domain category.
	//
	// List of thrown errors:
	//   [AuthenticationError]()
	//   [AuthorizationError]()
	//   [HeaderError]()
	//   [InternalError]()
	//   [QuotaError]()
	//   [RequestError]()
	GetDomainCategory(ctx context.Context, in *GetDomainCategoryRequest, opts ...grpc.CallOption) (*resources.DomainCategory, error)
}

type domainCategoryServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewDomainCategoryServiceClient(cc grpc.ClientConnInterface) DomainCategoryServiceClient {
	return &domainCategoryServiceClient{cc}
}

func (c *domainCategoryServiceClient) GetDomainCategory(ctx context.Context, in *GetDomainCategoryRequest, opts ...grpc.CallOption) (*resources.DomainCategory, error) {
	out := new(resources.DomainCategory)
	err := c.cc.Invoke(ctx, "/google.ads.googleads.v9.services.DomainCategoryService/GetDomainCategory", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// DomainCategoryServiceServer is the server API for DomainCategoryService service.
// All implementations must embed UnimplementedDomainCategoryServiceServer
// for forward compatibility
type DomainCategoryServiceServer interface {
	// Returns the requested domain category.
	//
	// List of thrown errors:
	//   [AuthenticationError]()
	//   [AuthorizationError]()
	//   [HeaderError]()
	//   [InternalError]()
	//   [QuotaError]()
	//   [RequestError]()
	GetDomainCategory(context.Context, *GetDomainCategoryRequest) (*resources.DomainCategory, error)
	mustEmbedUnimplementedDomainCategoryServiceServer()
}

// UnimplementedDomainCategoryServiceServer must be embedded to have forward compatible implementations.
type UnimplementedDomainCategoryServiceServer struct {
}

func (UnimplementedDomainCategoryServiceServer) GetDomainCategory(context.Context, *GetDomainCategoryRequest) (*resources.DomainCategory, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetDomainCategory not implemented")
}
func (UnimplementedDomainCategoryServiceServer) mustEmbedUnimplementedDomainCategoryServiceServer() {}

// UnsafeDomainCategoryServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to DomainCategoryServiceServer will
// result in compilation errors.
type UnsafeDomainCategoryServiceServer interface {
	mustEmbedUnimplementedDomainCategoryServiceServer()
}

func RegisterDomainCategoryServiceServer(s grpc.ServiceRegistrar, srv DomainCategoryServiceServer) {
	s.RegisterService(&DomainCategoryService_ServiceDesc, srv)
}

func _DomainCategoryService_GetDomainCategory_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetDomainCategoryRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DomainCategoryServiceServer).GetDomainCategory(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/google.ads.googleads.v9.services.DomainCategoryService/GetDomainCategory",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DomainCategoryServiceServer).GetDomainCategory(ctx, req.(*GetDomainCategoryRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// DomainCategoryService_ServiceDesc is the grpc.ServiceDesc for DomainCategoryService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var DomainCategoryService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "google.ads.googleads.v9.services.DomainCategoryService",
	HandlerType: (*DomainCategoryServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetDomainCategory",
			Handler:    _DomainCategoryService_GetDomainCategory_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "google/ads/googleads/v9/services/domain_category_service.proto",
}
