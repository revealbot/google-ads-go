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

// CustomerFeedServiceClient is the client API for CustomerFeedService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type CustomerFeedServiceClient interface {
	// Returns the requested customer feed in full detail.
	//
	// List of thrown errors:
	//   [AuthenticationError]()
	//   [AuthorizationError]()
	//   [HeaderError]()
	//   [InternalError]()
	//   [QuotaError]()
	//   [RequestError]()
	GetCustomerFeed(ctx context.Context, in *GetCustomerFeedRequest, opts ...grpc.CallOption) (*resources.CustomerFeed, error)
	// Creates, updates, or removes customer feeds. Operation statuses are
	// returned.
	//
	// List of thrown errors:
	//   [AuthenticationError]()
	//   [AuthorizationError]()
	//   [CollectionSizeError]()
	//   [CustomerFeedError]()
	//   [DatabaseError]()
	//   [DistinctError]()
	//   [FieldError]()
	//   [FieldMaskError]()
	//   [FunctionError]()
	//   [FunctionParsingError]()
	//   [HeaderError]()
	//   [IdError]()
	//   [InternalError]()
	//   [MutateError]()
	//   [NotEmptyError]()
	//   [OperatorError]()
	//   [QuotaError]()
	//   [RangeError]()
	//   [RequestError]()
	//   [SizeLimitError]()
	//   [StringFormatError]()
	//   [StringLengthError]()
	MutateCustomerFeeds(ctx context.Context, in *MutateCustomerFeedsRequest, opts ...grpc.CallOption) (*MutateCustomerFeedsResponse, error)
}

type customerFeedServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewCustomerFeedServiceClient(cc grpc.ClientConnInterface) CustomerFeedServiceClient {
	return &customerFeedServiceClient{cc}
}

func (c *customerFeedServiceClient) GetCustomerFeed(ctx context.Context, in *GetCustomerFeedRequest, opts ...grpc.CallOption) (*resources.CustomerFeed, error) {
	out := new(resources.CustomerFeed)
	err := c.cc.Invoke(ctx, "/google.ads.googleads.v9.services.CustomerFeedService/GetCustomerFeed", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *customerFeedServiceClient) MutateCustomerFeeds(ctx context.Context, in *MutateCustomerFeedsRequest, opts ...grpc.CallOption) (*MutateCustomerFeedsResponse, error) {
	out := new(MutateCustomerFeedsResponse)
	err := c.cc.Invoke(ctx, "/google.ads.googleads.v9.services.CustomerFeedService/MutateCustomerFeeds", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// CustomerFeedServiceServer is the server API for CustomerFeedService service.
// All implementations must embed UnimplementedCustomerFeedServiceServer
// for forward compatibility
type CustomerFeedServiceServer interface {
	// Returns the requested customer feed in full detail.
	//
	// List of thrown errors:
	//   [AuthenticationError]()
	//   [AuthorizationError]()
	//   [HeaderError]()
	//   [InternalError]()
	//   [QuotaError]()
	//   [RequestError]()
	GetCustomerFeed(context.Context, *GetCustomerFeedRequest) (*resources.CustomerFeed, error)
	// Creates, updates, or removes customer feeds. Operation statuses are
	// returned.
	//
	// List of thrown errors:
	//   [AuthenticationError]()
	//   [AuthorizationError]()
	//   [CollectionSizeError]()
	//   [CustomerFeedError]()
	//   [DatabaseError]()
	//   [DistinctError]()
	//   [FieldError]()
	//   [FieldMaskError]()
	//   [FunctionError]()
	//   [FunctionParsingError]()
	//   [HeaderError]()
	//   [IdError]()
	//   [InternalError]()
	//   [MutateError]()
	//   [NotEmptyError]()
	//   [OperatorError]()
	//   [QuotaError]()
	//   [RangeError]()
	//   [RequestError]()
	//   [SizeLimitError]()
	//   [StringFormatError]()
	//   [StringLengthError]()
	MutateCustomerFeeds(context.Context, *MutateCustomerFeedsRequest) (*MutateCustomerFeedsResponse, error)
	mustEmbedUnimplementedCustomerFeedServiceServer()
}

// UnimplementedCustomerFeedServiceServer must be embedded to have forward compatible implementations.
type UnimplementedCustomerFeedServiceServer struct {
}

func (UnimplementedCustomerFeedServiceServer) GetCustomerFeed(context.Context, *GetCustomerFeedRequest) (*resources.CustomerFeed, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetCustomerFeed not implemented")
}
func (UnimplementedCustomerFeedServiceServer) MutateCustomerFeeds(context.Context, *MutateCustomerFeedsRequest) (*MutateCustomerFeedsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method MutateCustomerFeeds not implemented")
}
func (UnimplementedCustomerFeedServiceServer) mustEmbedUnimplementedCustomerFeedServiceServer() {}

// UnsafeCustomerFeedServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to CustomerFeedServiceServer will
// result in compilation errors.
type UnsafeCustomerFeedServiceServer interface {
	mustEmbedUnimplementedCustomerFeedServiceServer()
}

func RegisterCustomerFeedServiceServer(s grpc.ServiceRegistrar, srv CustomerFeedServiceServer) {
	s.RegisterService(&CustomerFeedService_ServiceDesc, srv)
}

func _CustomerFeedService_GetCustomerFeed_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetCustomerFeedRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CustomerFeedServiceServer).GetCustomerFeed(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/google.ads.googleads.v9.services.CustomerFeedService/GetCustomerFeed",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CustomerFeedServiceServer).GetCustomerFeed(ctx, req.(*GetCustomerFeedRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _CustomerFeedService_MutateCustomerFeeds_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MutateCustomerFeedsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CustomerFeedServiceServer).MutateCustomerFeeds(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/google.ads.googleads.v9.services.CustomerFeedService/MutateCustomerFeeds",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CustomerFeedServiceServer).MutateCustomerFeeds(ctx, req.(*MutateCustomerFeedsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// CustomerFeedService_ServiceDesc is the grpc.ServiceDesc for CustomerFeedService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var CustomerFeedService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "google.ads.googleads.v9.services.CustomerFeedService",
	HandlerType: (*CustomerFeedServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetCustomerFeed",
			Handler:    _CustomerFeedService_GetCustomerFeed_Handler,
		},
		{
			MethodName: "MutateCustomerFeeds",
			Handler:    _CustomerFeedService_MutateCustomerFeeds_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "google/ads/googleads/v9/services/customer_feed_service.proto",
}
