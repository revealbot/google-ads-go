// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

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

// AdGroupFeedServiceClient is the client API for AdGroupFeedService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type AdGroupFeedServiceClient interface {
	// Creates, updates, or removes ad group feeds. Operation statuses are
	// returned.
	//
	// List of thrown errors:
	//   [AdGroupFeedError]()
	//   [AuthenticationError]()
	//   [AuthorizationError]()
	//   [CollectionSizeError]()
	//   [DatabaseError]()
	//   [DistinctError]()
	//   [FieldError]()
	//   [FunctionError]()
	//   [FunctionParsingError]()
	//   [HeaderError]()
	//   [IdError]()
	//   [InternalError]()
	//   [MutateError]()
	//   [NotEmptyError]()
	//   [NullError]()
	//   [OperatorError]()
	//   [QuotaError]()
	//   [RangeError]()
	//   [RequestError]()
	//   [SizeLimitError]()
	//   [StringFormatError]()
	//   [StringLengthError]()
	MutateAdGroupFeeds(ctx context.Context, in *MutateAdGroupFeedsRequest, opts ...grpc.CallOption) (*MutateAdGroupFeedsResponse, error)
}

type adGroupFeedServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewAdGroupFeedServiceClient(cc grpc.ClientConnInterface) AdGroupFeedServiceClient {
	return &adGroupFeedServiceClient{cc}
}

func (c *adGroupFeedServiceClient) MutateAdGroupFeeds(ctx context.Context, in *MutateAdGroupFeedsRequest, opts ...grpc.CallOption) (*MutateAdGroupFeedsResponse, error) {
	out := new(MutateAdGroupFeedsResponse)
	err := c.cc.Invoke(ctx, "/google.ads.googleads.v13.services.AdGroupFeedService/MutateAdGroupFeeds", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// AdGroupFeedServiceServer is the server API for AdGroupFeedService service.
// All implementations must embed UnimplementedAdGroupFeedServiceServer
// for forward compatibility
type AdGroupFeedServiceServer interface {
	// Creates, updates, or removes ad group feeds. Operation statuses are
	// returned.
	//
	// List of thrown errors:
	//   [AdGroupFeedError]()
	//   [AuthenticationError]()
	//   [AuthorizationError]()
	//   [CollectionSizeError]()
	//   [DatabaseError]()
	//   [DistinctError]()
	//   [FieldError]()
	//   [FunctionError]()
	//   [FunctionParsingError]()
	//   [HeaderError]()
	//   [IdError]()
	//   [InternalError]()
	//   [MutateError]()
	//   [NotEmptyError]()
	//   [NullError]()
	//   [OperatorError]()
	//   [QuotaError]()
	//   [RangeError]()
	//   [RequestError]()
	//   [SizeLimitError]()
	//   [StringFormatError]()
	//   [StringLengthError]()
	MutateAdGroupFeeds(context.Context, *MutateAdGroupFeedsRequest) (*MutateAdGroupFeedsResponse, error)
	mustEmbedUnimplementedAdGroupFeedServiceServer()
}

// UnimplementedAdGroupFeedServiceServer must be embedded to have forward compatible implementations.
type UnimplementedAdGroupFeedServiceServer struct {
}

func (UnimplementedAdGroupFeedServiceServer) MutateAdGroupFeeds(context.Context, *MutateAdGroupFeedsRequest) (*MutateAdGroupFeedsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method MutateAdGroupFeeds not implemented")
}
func (UnimplementedAdGroupFeedServiceServer) mustEmbedUnimplementedAdGroupFeedServiceServer() {}

// UnsafeAdGroupFeedServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to AdGroupFeedServiceServer will
// result in compilation errors.
type UnsafeAdGroupFeedServiceServer interface {
	mustEmbedUnimplementedAdGroupFeedServiceServer()
}

func RegisterAdGroupFeedServiceServer(s grpc.ServiceRegistrar, srv AdGroupFeedServiceServer) {
	s.RegisterService(&AdGroupFeedService_ServiceDesc, srv)
}

func _AdGroupFeedService_MutateAdGroupFeeds_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MutateAdGroupFeedsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AdGroupFeedServiceServer).MutateAdGroupFeeds(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/google.ads.googleads.v13.services.AdGroupFeedService/MutateAdGroupFeeds",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AdGroupFeedServiceServer).MutateAdGroupFeeds(ctx, req.(*MutateAdGroupFeedsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// AdGroupFeedService_ServiceDesc is the grpc.ServiceDesc for AdGroupFeedService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var AdGroupFeedService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "google.ads.googleads.v13.services.AdGroupFeedService",
	HandlerType: (*AdGroupFeedServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "MutateAdGroupFeeds",
			Handler:    _AdGroupFeedService_MutateAdGroupFeeds_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "google/ads/googleads/v13/services/ad_group_feed_service.proto",
}
