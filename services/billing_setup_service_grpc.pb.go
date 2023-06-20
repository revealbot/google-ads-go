// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.21.12
// source: google/ads/googleads/v14/services/billing_setup_service.proto

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

// BillingSetupServiceClient is the client API for BillingSetupService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type BillingSetupServiceClient interface {
	// Creates a billing setup, or cancels an existing billing setup.
	//
	// List of thrown errors:
	//
	//	[AuthenticationError]()
	//	[AuthorizationError]()
	//	[BillingSetupError]()
	//	[DateError]()
	//	[FieldError]()
	//	[HeaderError]()
	//	[InternalError]()
	//	[MutateError]()
	//	[QuotaError]()
	//	[RequestError]()
	MutateBillingSetup(ctx context.Context, in *MutateBillingSetupRequest, opts ...grpc.CallOption) (*MutateBillingSetupResponse, error)
}

type billingSetupServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewBillingSetupServiceClient(cc grpc.ClientConnInterface) BillingSetupServiceClient {
	return &billingSetupServiceClient{cc}
}

func (c *billingSetupServiceClient) MutateBillingSetup(ctx context.Context, in *MutateBillingSetupRequest, opts ...grpc.CallOption) (*MutateBillingSetupResponse, error) {
	out := new(MutateBillingSetupResponse)
	err := c.cc.Invoke(ctx, "/google.ads.googleads.v14.services.BillingSetupService/MutateBillingSetup", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// BillingSetupServiceServer is the server API for BillingSetupService service.
// All implementations must embed UnimplementedBillingSetupServiceServer
// for forward compatibility
type BillingSetupServiceServer interface {
	// Creates a billing setup, or cancels an existing billing setup.
	//
	// List of thrown errors:
	//
	//	[AuthenticationError]()
	//	[AuthorizationError]()
	//	[BillingSetupError]()
	//	[DateError]()
	//	[FieldError]()
	//	[HeaderError]()
	//	[InternalError]()
	//	[MutateError]()
	//	[QuotaError]()
	//	[RequestError]()
	MutateBillingSetup(context.Context, *MutateBillingSetupRequest) (*MutateBillingSetupResponse, error)
	mustEmbedUnimplementedBillingSetupServiceServer()
}

// UnimplementedBillingSetupServiceServer must be embedded to have forward compatible implementations.
type UnimplementedBillingSetupServiceServer struct {
}

func (UnimplementedBillingSetupServiceServer) MutateBillingSetup(context.Context, *MutateBillingSetupRequest) (*MutateBillingSetupResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method MutateBillingSetup not implemented")
}
func (UnimplementedBillingSetupServiceServer) mustEmbedUnimplementedBillingSetupServiceServer() {}

// UnsafeBillingSetupServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to BillingSetupServiceServer will
// result in compilation errors.
type UnsafeBillingSetupServiceServer interface {
	mustEmbedUnimplementedBillingSetupServiceServer()
}

func RegisterBillingSetupServiceServer(s grpc.ServiceRegistrar, srv BillingSetupServiceServer) {
	s.RegisterService(&BillingSetupService_ServiceDesc, srv)
}

func _BillingSetupService_MutateBillingSetup_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MutateBillingSetupRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BillingSetupServiceServer).MutateBillingSetup(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/google.ads.googleads.v14.services.BillingSetupService/MutateBillingSetup",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BillingSetupServiceServer).MutateBillingSetup(ctx, req.(*MutateBillingSetupRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// BillingSetupService_ServiceDesc is the grpc.ServiceDesc for BillingSetupService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var BillingSetupService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "google.ads.googleads.v14.services.BillingSetupService",
	HandlerType: (*BillingSetupServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "MutateBillingSetup",
			Handler:    _BillingSetupService_MutateBillingSetup_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "google/ads/googleads/v14/services/billing_setup_service.proto",
}
