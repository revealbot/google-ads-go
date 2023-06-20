// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.21.12
// source: google/ads/googleads/v14/services/product_link_service.proto

package services

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

// ProductLinkServiceClient is the client API for ProductLinkService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type ProductLinkServiceClient interface {
	// Creates a product link.
	//
	// List of thrown errors:
	//
	//	[AuthenticationError]()
	//	[AuthorizationError]()
	//	[DatabaseError]()
	//	[FieldError]()
	//	[HeaderError]()
	//	[InternalError]()
	//	[MutateError]()
	//	[QuotaError]()
	//	[RequestError]()
	CreateProductLink(ctx context.Context, in *CreateProductLinkRequest, opts ...grpc.CallOption) (*CreateProductLinkResponse, error)
	// Removes a product link.
	//
	// List of thrown errors:
	//
	//	[AuthenticationError]()
	//	[AuthorizationError]()
	//	[FieldMaskError]()
	//	[HeaderError]()
	//	[InternalError]()
	//	[MutateError]()
	//	[QuotaError]()
	//	[RequestError]()
	RemoveProductLink(ctx context.Context, in *RemoveProductLinkRequest, opts ...grpc.CallOption) (*RemoveProductLinkResponse, error)
}

type productLinkServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewProductLinkServiceClient(cc grpc.ClientConnInterface) ProductLinkServiceClient {
	return &productLinkServiceClient{cc}
}

func (c *productLinkServiceClient) CreateProductLink(ctx context.Context, in *CreateProductLinkRequest, opts ...grpc.CallOption) (*CreateProductLinkResponse, error) {
	out := new(CreateProductLinkResponse)
	err := c.cc.Invoke(ctx, "/google.ads.googleads.v14.services.ProductLinkService/CreateProductLink", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *productLinkServiceClient) RemoveProductLink(ctx context.Context, in *RemoveProductLinkRequest, opts ...grpc.CallOption) (*RemoveProductLinkResponse, error) {
	out := new(RemoveProductLinkResponse)
	err := c.cc.Invoke(ctx, "/google.ads.googleads.v14.services.ProductLinkService/RemoveProductLink", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ProductLinkServiceServer is the server API for ProductLinkService service.
// All implementations must embed UnimplementedProductLinkServiceServer
// for forward compatibility
type ProductLinkServiceServer interface {
	// Creates a product link.
	//
	// List of thrown errors:
	//
	//	[AuthenticationError]()
	//	[AuthorizationError]()
	//	[DatabaseError]()
	//	[FieldError]()
	//	[HeaderError]()
	//	[InternalError]()
	//	[MutateError]()
	//	[QuotaError]()
	//	[RequestError]()
	CreateProductLink(context.Context, *CreateProductLinkRequest) (*CreateProductLinkResponse, error)
	// Removes a product link.
	//
	// List of thrown errors:
	//
	//	[AuthenticationError]()
	//	[AuthorizationError]()
	//	[FieldMaskError]()
	//	[HeaderError]()
	//	[InternalError]()
	//	[MutateError]()
	//	[QuotaError]()
	//	[RequestError]()
	RemoveProductLink(context.Context, *RemoveProductLinkRequest) (*RemoveProductLinkResponse, error)
	mustEmbedUnimplementedProductLinkServiceServer()
}

// UnimplementedProductLinkServiceServer must be embedded to have forward compatible implementations.
type UnimplementedProductLinkServiceServer struct {
}

func (UnimplementedProductLinkServiceServer) CreateProductLink(context.Context, *CreateProductLinkRequest) (*CreateProductLinkResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateProductLink not implemented")
}
func (UnimplementedProductLinkServiceServer) RemoveProductLink(context.Context, *RemoveProductLinkRequest) (*RemoveProductLinkResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method RemoveProductLink not implemented")
}
func (UnimplementedProductLinkServiceServer) mustEmbedUnimplementedProductLinkServiceServer() {}

// UnsafeProductLinkServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to ProductLinkServiceServer will
// result in compilation errors.
type UnsafeProductLinkServiceServer interface {
	mustEmbedUnimplementedProductLinkServiceServer()
}

func RegisterProductLinkServiceServer(s grpc.ServiceRegistrar, srv ProductLinkServiceServer) {
	s.RegisterService(&ProductLinkService_ServiceDesc, srv)
}

func _ProductLinkService_CreateProductLink_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateProductLinkRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ProductLinkServiceServer).CreateProductLink(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/google.ads.googleads.v14.services.ProductLinkService/CreateProductLink",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ProductLinkServiceServer).CreateProductLink(ctx, req.(*CreateProductLinkRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ProductLinkService_RemoveProductLink_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RemoveProductLinkRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ProductLinkServiceServer).RemoveProductLink(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/google.ads.googleads.v14.services.ProductLinkService/RemoveProductLink",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ProductLinkServiceServer).RemoveProductLink(ctx, req.(*RemoveProductLinkRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// ProductLinkService_ServiceDesc is the grpc.ServiceDesc for ProductLinkService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var ProductLinkService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "google.ads.googleads.v14.services.ProductLinkService",
	HandlerType: (*ProductLinkServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateProductLink",
			Handler:    _ProductLinkService_CreateProductLink_Handler,
		},
		{
			MethodName: "RemoveProductLink",
			Handler:    _ProductLinkService_RemoveProductLink_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "google/ads/googleads/v14/services/product_link_service.proto",
}
