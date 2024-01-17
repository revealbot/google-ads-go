// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.21.12
// source: google/ads/googleads/v15/services/product_link_invitation_service.proto

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

// ProductLinkInvitationServiceClient is the client API for ProductLinkInvitationService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type ProductLinkInvitationServiceClient interface {
	// Update a product link invitation.
	UpdateProductLinkInvitation(ctx context.Context, in *UpdateProductLinkInvitationRequest, opts ...grpc.CallOption) (*UpdateProductLinkInvitationResponse, error)
}

type productLinkInvitationServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewProductLinkInvitationServiceClient(cc grpc.ClientConnInterface) ProductLinkInvitationServiceClient {
	return &productLinkInvitationServiceClient{cc}
}

func (c *productLinkInvitationServiceClient) UpdateProductLinkInvitation(ctx context.Context, in *UpdateProductLinkInvitationRequest, opts ...grpc.CallOption) (*UpdateProductLinkInvitationResponse, error) {
	out := new(UpdateProductLinkInvitationResponse)
	err := c.cc.Invoke(ctx, "/google.ads.googleads.v15.services.ProductLinkInvitationService/UpdateProductLinkInvitation", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ProductLinkInvitationServiceServer is the server API for ProductLinkInvitationService service.
// All implementations must embed UnimplementedProductLinkInvitationServiceServer
// for forward compatibility
type ProductLinkInvitationServiceServer interface {
	// Update a product link invitation.
	UpdateProductLinkInvitation(context.Context, *UpdateProductLinkInvitationRequest) (*UpdateProductLinkInvitationResponse, error)
	mustEmbedUnimplementedProductLinkInvitationServiceServer()
}

// UnimplementedProductLinkInvitationServiceServer must be embedded to have forward compatible implementations.
type UnimplementedProductLinkInvitationServiceServer struct {
}

func (UnimplementedProductLinkInvitationServiceServer) UpdateProductLinkInvitation(context.Context, *UpdateProductLinkInvitationRequest) (*UpdateProductLinkInvitationResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateProductLinkInvitation not implemented")
}
func (UnimplementedProductLinkInvitationServiceServer) mustEmbedUnimplementedProductLinkInvitationServiceServer() {
}

// UnsafeProductLinkInvitationServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to ProductLinkInvitationServiceServer will
// result in compilation errors.
type UnsafeProductLinkInvitationServiceServer interface {
	mustEmbedUnimplementedProductLinkInvitationServiceServer()
}

func RegisterProductLinkInvitationServiceServer(s grpc.ServiceRegistrar, srv ProductLinkInvitationServiceServer) {
	s.RegisterService(&ProductLinkInvitationService_ServiceDesc, srv)
}

func _ProductLinkInvitationService_UpdateProductLinkInvitation_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateProductLinkInvitationRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ProductLinkInvitationServiceServer).UpdateProductLinkInvitation(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/google.ads.googleads.v15.services.ProductLinkInvitationService/UpdateProductLinkInvitation",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ProductLinkInvitationServiceServer).UpdateProductLinkInvitation(ctx, req.(*UpdateProductLinkInvitationRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// ProductLinkInvitationService_ServiceDesc is the grpc.ServiceDesc for ProductLinkInvitationService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var ProductLinkInvitationService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "google.ads.googleads.v15.services.ProductLinkInvitationService",
	HandlerType: (*ProductLinkInvitationServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "UpdateProductLinkInvitation",
			Handler:    _ProductLinkInvitationService_UpdateProductLinkInvitation_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "google/ads/googleads/v15/services/product_link_invitation_service.proto",
}
