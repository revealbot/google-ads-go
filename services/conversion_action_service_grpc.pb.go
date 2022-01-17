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

// ConversionActionServiceClient is the client API for ConversionActionService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type ConversionActionServiceClient interface {
	// Returns the requested conversion action.
	//
	// List of thrown errors:
	//   [AuthenticationError]()
	//   [AuthorizationError]()
	//   [HeaderError]()
	//   [InternalError]()
	//   [QuotaError]()
	//   [RequestError]()
	GetConversionAction(ctx context.Context, in *GetConversionActionRequest, opts ...grpc.CallOption) (*resources.ConversionAction, error)
	// Creates, updates or removes conversion actions. Operation statuses are
	// returned.
	//
	// List of thrown errors:
	//   [AuthenticationError]()
	//   [AuthorizationError]()
	//   [ConversionActionError]()
	//   [CurrencyCodeError]()
	//   [DatabaseError]()
	//   [FieldError]()
	//   [FieldMaskError]()
	//   [HeaderError]()
	//   [InternalError]()
	//   [MutateError]()
	//   [NewResourceCreationError]()
	//   [QuotaError]()
	//   [RangeError]()
	//   [RequestError]()
	//   [ResourceCountLimitExceededError]()
	//   [StringLengthError]()
	MutateConversionActions(ctx context.Context, in *MutateConversionActionsRequest, opts ...grpc.CallOption) (*MutateConversionActionsResponse, error)
}

type conversionActionServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewConversionActionServiceClient(cc grpc.ClientConnInterface) ConversionActionServiceClient {
	return &conversionActionServiceClient{cc}
}

func (c *conversionActionServiceClient) GetConversionAction(ctx context.Context, in *GetConversionActionRequest, opts ...grpc.CallOption) (*resources.ConversionAction, error) {
	out := new(resources.ConversionAction)
	err := c.cc.Invoke(ctx, "/google.ads.googleads.v9.services.ConversionActionService/GetConversionAction", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *conversionActionServiceClient) MutateConversionActions(ctx context.Context, in *MutateConversionActionsRequest, opts ...grpc.CallOption) (*MutateConversionActionsResponse, error) {
	out := new(MutateConversionActionsResponse)
	err := c.cc.Invoke(ctx, "/google.ads.googleads.v9.services.ConversionActionService/MutateConversionActions", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ConversionActionServiceServer is the server API for ConversionActionService service.
// All implementations must embed UnimplementedConversionActionServiceServer
// for forward compatibility
type ConversionActionServiceServer interface {
	// Returns the requested conversion action.
	//
	// List of thrown errors:
	//   [AuthenticationError]()
	//   [AuthorizationError]()
	//   [HeaderError]()
	//   [InternalError]()
	//   [QuotaError]()
	//   [RequestError]()
	GetConversionAction(context.Context, *GetConversionActionRequest) (*resources.ConversionAction, error)
	// Creates, updates or removes conversion actions. Operation statuses are
	// returned.
	//
	// List of thrown errors:
	//   [AuthenticationError]()
	//   [AuthorizationError]()
	//   [ConversionActionError]()
	//   [CurrencyCodeError]()
	//   [DatabaseError]()
	//   [FieldError]()
	//   [FieldMaskError]()
	//   [HeaderError]()
	//   [InternalError]()
	//   [MutateError]()
	//   [NewResourceCreationError]()
	//   [QuotaError]()
	//   [RangeError]()
	//   [RequestError]()
	//   [ResourceCountLimitExceededError]()
	//   [StringLengthError]()
	MutateConversionActions(context.Context, *MutateConversionActionsRequest) (*MutateConversionActionsResponse, error)
	mustEmbedUnimplementedConversionActionServiceServer()
}

// UnimplementedConversionActionServiceServer must be embedded to have forward compatible implementations.
type UnimplementedConversionActionServiceServer struct {
}

func (UnimplementedConversionActionServiceServer) GetConversionAction(context.Context, *GetConversionActionRequest) (*resources.ConversionAction, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetConversionAction not implemented")
}
func (UnimplementedConversionActionServiceServer) MutateConversionActions(context.Context, *MutateConversionActionsRequest) (*MutateConversionActionsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method MutateConversionActions not implemented")
}
func (UnimplementedConversionActionServiceServer) mustEmbedUnimplementedConversionActionServiceServer() {
}

// UnsafeConversionActionServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to ConversionActionServiceServer will
// result in compilation errors.
type UnsafeConversionActionServiceServer interface {
	mustEmbedUnimplementedConversionActionServiceServer()
}

func RegisterConversionActionServiceServer(s grpc.ServiceRegistrar, srv ConversionActionServiceServer) {
	s.RegisterService(&ConversionActionService_ServiceDesc, srv)
}

func _ConversionActionService_GetConversionAction_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetConversionActionRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ConversionActionServiceServer).GetConversionAction(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/google.ads.googleads.v9.services.ConversionActionService/GetConversionAction",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ConversionActionServiceServer).GetConversionAction(ctx, req.(*GetConversionActionRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ConversionActionService_MutateConversionActions_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MutateConversionActionsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ConversionActionServiceServer).MutateConversionActions(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/google.ads.googleads.v9.services.ConversionActionService/MutateConversionActions",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ConversionActionServiceServer).MutateConversionActions(ctx, req.(*MutateConversionActionsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// ConversionActionService_ServiceDesc is the grpc.ServiceDesc for ConversionActionService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var ConversionActionService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "google.ads.googleads.v9.services.ConversionActionService",
	HandlerType: (*ConversionActionServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetConversionAction",
			Handler:    _ConversionActionService_GetConversionAction_Handler,
		},
		{
			MethodName: "MutateConversionActions",
			Handler:    _ConversionActionService_MutateConversionActions_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "google/ads/googleads/v9/services/conversion_action_service.proto",
}
