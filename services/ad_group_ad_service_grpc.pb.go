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

// AdGroupAdServiceClient is the client API for AdGroupAdService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type AdGroupAdServiceClient interface {
	// Creates, updates, or removes ads. Operation statuses are returned.
	//
	// List of thrown errors:
	//   [AdCustomizerError]()
	//   [AdError]()
	//   [AdGroupAdError]()
	//   [AdSharingError]()
	//   [AdxError]()
	//   [AssetError]()
	//   [AssetLinkError]()
	//   [AuthenticationError]()
	//   [AuthorizationError]()
	//   [CollectionSizeError]()
	//   [ContextError]()
	//   [DatabaseError]()
	//   [DateError]()
	//   [DistinctError]()
	//   [FeedAttributeReferenceError]()
	//   [FieldError]()
	//   [FieldMaskError]()
	//   [FunctionError]()
	//   [FunctionParsingError]()
	//   [HeaderError]()
	//   [IdError]()
	//   [ImageError]()
	//   [InternalError]()
	//   [ListOperationError]()
	//   [MediaBundleError]()
	//   [MediaFileError]()
	//   [MutateError]()
	//   [NewResourceCreationError]()
	//   [NotEmptyError]()
	//   [NullError]()
	//   [OperationAccessDeniedError]()
	//   [OperatorError]()
	//   [PolicyFindingError]()
	//   [PolicyValidationParameterError]()
	//   [PolicyViolationError]()
	//   [QuotaError]()
	//   [RangeError]()
	//   [RequestError]()
	//   [ResourceCountLimitExceededError]()
	//   [SizeLimitError]()
	//   [StringFormatError]()
	//   [StringLengthError]()
	//   [UrlFieldError]()
	MutateAdGroupAds(ctx context.Context, in *MutateAdGroupAdsRequest, opts ...grpc.CallOption) (*MutateAdGroupAdsResponse, error)
}

type adGroupAdServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewAdGroupAdServiceClient(cc grpc.ClientConnInterface) AdGroupAdServiceClient {
	return &adGroupAdServiceClient{cc}
}

func (c *adGroupAdServiceClient) MutateAdGroupAds(ctx context.Context, in *MutateAdGroupAdsRequest, opts ...grpc.CallOption) (*MutateAdGroupAdsResponse, error) {
	out := new(MutateAdGroupAdsResponse)
	err := c.cc.Invoke(ctx, "/google.ads.googleads.v13.services.AdGroupAdService/MutateAdGroupAds", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// AdGroupAdServiceServer is the server API for AdGroupAdService service.
// All implementations must embed UnimplementedAdGroupAdServiceServer
// for forward compatibility
type AdGroupAdServiceServer interface {
	// Creates, updates, or removes ads. Operation statuses are returned.
	//
	// List of thrown errors:
	//   [AdCustomizerError]()
	//   [AdError]()
	//   [AdGroupAdError]()
	//   [AdSharingError]()
	//   [AdxError]()
	//   [AssetError]()
	//   [AssetLinkError]()
	//   [AuthenticationError]()
	//   [AuthorizationError]()
	//   [CollectionSizeError]()
	//   [ContextError]()
	//   [DatabaseError]()
	//   [DateError]()
	//   [DistinctError]()
	//   [FeedAttributeReferenceError]()
	//   [FieldError]()
	//   [FieldMaskError]()
	//   [FunctionError]()
	//   [FunctionParsingError]()
	//   [HeaderError]()
	//   [IdError]()
	//   [ImageError]()
	//   [InternalError]()
	//   [ListOperationError]()
	//   [MediaBundleError]()
	//   [MediaFileError]()
	//   [MutateError]()
	//   [NewResourceCreationError]()
	//   [NotEmptyError]()
	//   [NullError]()
	//   [OperationAccessDeniedError]()
	//   [OperatorError]()
	//   [PolicyFindingError]()
	//   [PolicyValidationParameterError]()
	//   [PolicyViolationError]()
	//   [QuotaError]()
	//   [RangeError]()
	//   [RequestError]()
	//   [ResourceCountLimitExceededError]()
	//   [SizeLimitError]()
	//   [StringFormatError]()
	//   [StringLengthError]()
	//   [UrlFieldError]()
	MutateAdGroupAds(context.Context, *MutateAdGroupAdsRequest) (*MutateAdGroupAdsResponse, error)
	mustEmbedUnimplementedAdGroupAdServiceServer()
}

// UnimplementedAdGroupAdServiceServer must be embedded to have forward compatible implementations.
type UnimplementedAdGroupAdServiceServer struct {
}

func (UnimplementedAdGroupAdServiceServer) MutateAdGroupAds(context.Context, *MutateAdGroupAdsRequest) (*MutateAdGroupAdsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method MutateAdGroupAds not implemented")
}
func (UnimplementedAdGroupAdServiceServer) mustEmbedUnimplementedAdGroupAdServiceServer() {}

// UnsafeAdGroupAdServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to AdGroupAdServiceServer will
// result in compilation errors.
type UnsafeAdGroupAdServiceServer interface {
	mustEmbedUnimplementedAdGroupAdServiceServer()
}

func RegisterAdGroupAdServiceServer(s grpc.ServiceRegistrar, srv AdGroupAdServiceServer) {
	s.RegisterService(&AdGroupAdService_ServiceDesc, srv)
}

func _AdGroupAdService_MutateAdGroupAds_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MutateAdGroupAdsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AdGroupAdServiceServer).MutateAdGroupAds(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/google.ads.googleads.v13.services.AdGroupAdService/MutateAdGroupAds",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AdGroupAdServiceServer).MutateAdGroupAds(ctx, req.(*MutateAdGroupAdsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// AdGroupAdService_ServiceDesc is the grpc.ServiceDesc for AdGroupAdService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var AdGroupAdService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "google.ads.googleads.v13.services.AdGroupAdService",
	HandlerType: (*AdGroupAdServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "MutateAdGroupAds",
			Handler:    _AdGroupAdService_MutateAdGroupAds_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "google/ads/googleads/v13/services/ad_group_ad_service.proto",
}
