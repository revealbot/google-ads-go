// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.21.12
// source: google/ads/googleads/v16/services/google_ads_service.proto

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

// GoogleAdsServiceClient is the client API for GoogleAdsService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type GoogleAdsServiceClient interface {
	// Returns all rows that match the search query.
	//
	// List of thrown errors:
	//
	//	[AuthenticationError]()
	//	[AuthorizationError]()
	//	[ChangeEventError]()
	//	[ChangeStatusError]()
	//	[ClickViewError]()
	//	[HeaderError]()
	//	[InternalError]()
	//	[QueryError]()
	//	[QuotaError]()
	//	[RequestError]()
	Search(ctx context.Context, in *SearchGoogleAdsRequest, opts ...grpc.CallOption) (*SearchGoogleAdsResponse, error)
	// Returns all rows that match the search stream query.
	//
	// List of thrown errors:
	//
	//	[AuthenticationError]()
	//	[AuthorizationError]()
	//	[ChangeEventError]()
	//	[ChangeStatusError]()
	//	[ClickViewError]()
	//	[HeaderError]()
	//	[InternalError]()
	//	[QueryError]()
	//	[QuotaError]()
	//	[RequestError]()
	SearchStream(ctx context.Context, in *SearchGoogleAdsStreamRequest, opts ...grpc.CallOption) (GoogleAdsService_SearchStreamClient, error)
	// Creates, updates, or removes resources. This method supports atomic
	// transactions with multiple types of resources. For example, you can
	// atomically create a campaign and a campaign budget, or perform up to
	// thousands of mutates atomically.
	//
	// This method is essentially a wrapper around a series of mutate methods. The
	// only features it offers over calling those methods directly are:
	//
	// - Atomic transactions
	// - Temp resource names (described below)
	// - Somewhat reduced latency over making a series of mutate calls
	//
	// Note: Only resources that support atomic transactions are included, so this
	// method can't replace all calls to individual services.
	//
	// ## Atomic Transaction Benefits
	//
	// Atomicity makes error handling much easier. If you're making a series of
	// changes and one fails, it can leave your account in an inconsistent state.
	// With atomicity, you either reach the chosen state directly, or the request
	// fails and you can retry.
	//
	// ## Temp Resource Names
	//
	// Temp resource names are a special type of resource name used to create a
	// resource and reference that resource in the same request. For example, if a
	// campaign budget is created with `resource_name` equal to
	// `customers/123/campaignBudgets/-1`, that resource name can be reused in
	// the `Campaign.budget` field in the same request. That way, the two
	// resources are created and linked atomically.
	//
	// To create a temp resource name, put a negative number in the part of the
	// name that the server would normally allocate.
	//
	// Note:
	//
	//   - Resources must be created with a temp name before the name can be reused.
	//     For example, the previous CampaignBudget+Campaign example would fail if
	//     the mutate order was reversed.
	//   - Temp names are not remembered across requests.
	//   - There's no limit to the number of temp names in a request.
	//   - Each temp name must use a unique negative number, even if the resource
	//     types differ.
	//
	// ## Latency
	//
	// It's important to group mutates by resource type or the request may time
	// out and fail. Latency is roughly equal to a series of calls to individual
	// mutate methods, where each change in resource type is a new call. For
	// example, mutating 10 campaigns then 10 ad groups is like 2 calls, while
	// mutating 1 campaign, 1 ad group, 1 campaign, 1 ad group is like 4 calls.
	//
	// List of thrown errors:
	//
	//	[AdCustomizerError]()
	//	[AdError]()
	//	[AdGroupAdError]()
	//	[AdGroupCriterionError]()
	//	[AdGroupError]()
	//	[AssetError]()
	//	[AuthenticationError]()
	//	[AuthorizationError]()
	//	[BiddingError]()
	//	[CampaignBudgetError]()
	//	[CampaignCriterionError]()
	//	[CampaignError]()
	//	[CampaignExperimentError]()
	//	[CampaignSharedSetError]()
	//	[CollectionSizeError]()
	//	[ContextError]()
	//	[ConversionActionError]()
	//	[CriterionError]()
	//	[CustomerFeedError]()
	//	[DatabaseError]()
	//	[DateError]()
	//	[DateRangeError]()
	//	[DistinctError]()
	//	[ExtensionFeedItemError]()
	//	[ExtensionSettingError]()
	//	[FeedAttributeReferenceError]()
	//	[FeedError]()
	//	[FeedItemError]()
	//	[FeedItemSetError]()
	//	[FieldError]()
	//	[FieldMaskError]()
	//	[FunctionParsingError]()
	//	[HeaderError]()
	//	[ImageError]()
	//	[InternalError]()
	//	[KeywordPlanAdGroupKeywordError]()
	//	[KeywordPlanCampaignError]()
	//	[KeywordPlanError]()
	//	[LabelError]()
	//	[ListOperationError]()
	//	[MediaUploadError]()
	//	[MutateError]()
	//	[NewResourceCreationError]()
	//	[NullError]()
	//	[OperationAccessDeniedError]()
	//	[PolicyFindingError]()
	//	[PolicyViolationError]()
	//	[QuotaError]()
	//	[RangeError]()
	//	[RequestError]()
	//	[ResourceCountLimitExceededError]()
	//	[SettingError]()
	//	[SharedSetError]()
	//	[SizeLimitError]()
	//	[StringFormatError]()
	//	[StringLengthError]()
	//	[UrlFieldError]()
	//	[UserListError]()
	//	[YoutubeVideoRegistrationError]()
	Mutate(ctx context.Context, in *MutateGoogleAdsRequest, opts ...grpc.CallOption) (*MutateGoogleAdsResponse, error)
}

type googleAdsServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewGoogleAdsServiceClient(cc grpc.ClientConnInterface) GoogleAdsServiceClient {
	return &googleAdsServiceClient{cc}
}

func (c *googleAdsServiceClient) Search(ctx context.Context, in *SearchGoogleAdsRequest, opts ...grpc.CallOption) (*SearchGoogleAdsResponse, error) {
	out := new(SearchGoogleAdsResponse)
	err := c.cc.Invoke(ctx, "/google.ads.googleads.v16.services.GoogleAdsService/Search", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *googleAdsServiceClient) SearchStream(ctx context.Context, in *SearchGoogleAdsStreamRequest, opts ...grpc.CallOption) (GoogleAdsService_SearchStreamClient, error) {
	stream, err := c.cc.NewStream(ctx, &GoogleAdsService_ServiceDesc.Streams[0], "/google.ads.googleads.v16.services.GoogleAdsService/SearchStream", opts...)
	if err != nil {
		return nil, err
	}
	x := &googleAdsServiceSearchStreamClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type GoogleAdsService_SearchStreamClient interface {
	Recv() (*SearchGoogleAdsStreamResponse, error)
	grpc.ClientStream
}

type googleAdsServiceSearchStreamClient struct {
	grpc.ClientStream
}

func (x *googleAdsServiceSearchStreamClient) Recv() (*SearchGoogleAdsStreamResponse, error) {
	m := new(SearchGoogleAdsStreamResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *googleAdsServiceClient) Mutate(ctx context.Context, in *MutateGoogleAdsRequest, opts ...grpc.CallOption) (*MutateGoogleAdsResponse, error) {
	out := new(MutateGoogleAdsResponse)
	err := c.cc.Invoke(ctx, "/google.ads.googleads.v16.services.GoogleAdsService/Mutate", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// GoogleAdsServiceServer is the server API for GoogleAdsService service.
// All implementations must embed UnimplementedGoogleAdsServiceServer
// for forward compatibility
type GoogleAdsServiceServer interface {
	// Returns all rows that match the search query.
	//
	// List of thrown errors:
	//
	//	[AuthenticationError]()
	//	[AuthorizationError]()
	//	[ChangeEventError]()
	//	[ChangeStatusError]()
	//	[ClickViewError]()
	//	[HeaderError]()
	//	[InternalError]()
	//	[QueryError]()
	//	[QuotaError]()
	//	[RequestError]()
	Search(context.Context, *SearchGoogleAdsRequest) (*SearchGoogleAdsResponse, error)
	// Returns all rows that match the search stream query.
	//
	// List of thrown errors:
	//
	//	[AuthenticationError]()
	//	[AuthorizationError]()
	//	[ChangeEventError]()
	//	[ChangeStatusError]()
	//	[ClickViewError]()
	//	[HeaderError]()
	//	[InternalError]()
	//	[QueryError]()
	//	[QuotaError]()
	//	[RequestError]()
	SearchStream(*SearchGoogleAdsStreamRequest, GoogleAdsService_SearchStreamServer) error
	// Creates, updates, or removes resources. This method supports atomic
	// transactions with multiple types of resources. For example, you can
	// atomically create a campaign and a campaign budget, or perform up to
	// thousands of mutates atomically.
	//
	// This method is essentially a wrapper around a series of mutate methods. The
	// only features it offers over calling those methods directly are:
	//
	// - Atomic transactions
	// - Temp resource names (described below)
	// - Somewhat reduced latency over making a series of mutate calls
	//
	// Note: Only resources that support atomic transactions are included, so this
	// method can't replace all calls to individual services.
	//
	// ## Atomic Transaction Benefits
	//
	// Atomicity makes error handling much easier. If you're making a series of
	// changes and one fails, it can leave your account in an inconsistent state.
	// With atomicity, you either reach the chosen state directly, or the request
	// fails and you can retry.
	//
	// ## Temp Resource Names
	//
	// Temp resource names are a special type of resource name used to create a
	// resource and reference that resource in the same request. For example, if a
	// campaign budget is created with `resource_name` equal to
	// `customers/123/campaignBudgets/-1`, that resource name can be reused in
	// the `Campaign.budget` field in the same request. That way, the two
	// resources are created and linked atomically.
	//
	// To create a temp resource name, put a negative number in the part of the
	// name that the server would normally allocate.
	//
	// Note:
	//
	//   - Resources must be created with a temp name before the name can be reused.
	//     For example, the previous CampaignBudget+Campaign example would fail if
	//     the mutate order was reversed.
	//   - Temp names are not remembered across requests.
	//   - There's no limit to the number of temp names in a request.
	//   - Each temp name must use a unique negative number, even if the resource
	//     types differ.
	//
	// ## Latency
	//
	// It's important to group mutates by resource type or the request may time
	// out and fail. Latency is roughly equal to a series of calls to individual
	// mutate methods, where each change in resource type is a new call. For
	// example, mutating 10 campaigns then 10 ad groups is like 2 calls, while
	// mutating 1 campaign, 1 ad group, 1 campaign, 1 ad group is like 4 calls.
	//
	// List of thrown errors:
	//
	//	[AdCustomizerError]()
	//	[AdError]()
	//	[AdGroupAdError]()
	//	[AdGroupCriterionError]()
	//	[AdGroupError]()
	//	[AssetError]()
	//	[AuthenticationError]()
	//	[AuthorizationError]()
	//	[BiddingError]()
	//	[CampaignBudgetError]()
	//	[CampaignCriterionError]()
	//	[CampaignError]()
	//	[CampaignExperimentError]()
	//	[CampaignSharedSetError]()
	//	[CollectionSizeError]()
	//	[ContextError]()
	//	[ConversionActionError]()
	//	[CriterionError]()
	//	[CustomerFeedError]()
	//	[DatabaseError]()
	//	[DateError]()
	//	[DateRangeError]()
	//	[DistinctError]()
	//	[ExtensionFeedItemError]()
	//	[ExtensionSettingError]()
	//	[FeedAttributeReferenceError]()
	//	[FeedError]()
	//	[FeedItemError]()
	//	[FeedItemSetError]()
	//	[FieldError]()
	//	[FieldMaskError]()
	//	[FunctionParsingError]()
	//	[HeaderError]()
	//	[ImageError]()
	//	[InternalError]()
	//	[KeywordPlanAdGroupKeywordError]()
	//	[KeywordPlanCampaignError]()
	//	[KeywordPlanError]()
	//	[LabelError]()
	//	[ListOperationError]()
	//	[MediaUploadError]()
	//	[MutateError]()
	//	[NewResourceCreationError]()
	//	[NullError]()
	//	[OperationAccessDeniedError]()
	//	[PolicyFindingError]()
	//	[PolicyViolationError]()
	//	[QuotaError]()
	//	[RangeError]()
	//	[RequestError]()
	//	[ResourceCountLimitExceededError]()
	//	[SettingError]()
	//	[SharedSetError]()
	//	[SizeLimitError]()
	//	[StringFormatError]()
	//	[StringLengthError]()
	//	[UrlFieldError]()
	//	[UserListError]()
	//	[YoutubeVideoRegistrationError]()
	Mutate(context.Context, *MutateGoogleAdsRequest) (*MutateGoogleAdsResponse, error)
	mustEmbedUnimplementedGoogleAdsServiceServer()
}

// UnimplementedGoogleAdsServiceServer must be embedded to have forward compatible implementations.
type UnimplementedGoogleAdsServiceServer struct {
}

func (UnimplementedGoogleAdsServiceServer) Search(context.Context, *SearchGoogleAdsRequest) (*SearchGoogleAdsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Search not implemented")
}
func (UnimplementedGoogleAdsServiceServer) SearchStream(*SearchGoogleAdsStreamRequest, GoogleAdsService_SearchStreamServer) error {
	return status.Errorf(codes.Unimplemented, "method SearchStream not implemented")
}
func (UnimplementedGoogleAdsServiceServer) Mutate(context.Context, *MutateGoogleAdsRequest) (*MutateGoogleAdsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Mutate not implemented")
}
func (UnimplementedGoogleAdsServiceServer) mustEmbedUnimplementedGoogleAdsServiceServer() {}

// UnsafeGoogleAdsServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to GoogleAdsServiceServer will
// result in compilation errors.
type UnsafeGoogleAdsServiceServer interface {
	mustEmbedUnimplementedGoogleAdsServiceServer()
}

func RegisterGoogleAdsServiceServer(s grpc.ServiceRegistrar, srv GoogleAdsServiceServer) {
	s.RegisterService(&GoogleAdsService_ServiceDesc, srv)
}

func _GoogleAdsService_Search_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SearchGoogleAdsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GoogleAdsServiceServer).Search(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/google.ads.googleads.v16.services.GoogleAdsService/Search",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GoogleAdsServiceServer).Search(ctx, req.(*SearchGoogleAdsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _GoogleAdsService_SearchStream_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(SearchGoogleAdsStreamRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(GoogleAdsServiceServer).SearchStream(m, &googleAdsServiceSearchStreamServer{stream})
}

type GoogleAdsService_SearchStreamServer interface {
	Send(*SearchGoogleAdsStreamResponse) error
	grpc.ServerStream
}

type googleAdsServiceSearchStreamServer struct {
	grpc.ServerStream
}

func (x *googleAdsServiceSearchStreamServer) Send(m *SearchGoogleAdsStreamResponse) error {
	return x.ServerStream.SendMsg(m)
}

func _GoogleAdsService_Mutate_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MutateGoogleAdsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GoogleAdsServiceServer).Mutate(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/google.ads.googleads.v16.services.GoogleAdsService/Mutate",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GoogleAdsServiceServer).Mutate(ctx, req.(*MutateGoogleAdsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// GoogleAdsService_ServiceDesc is the grpc.ServiceDesc for GoogleAdsService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var GoogleAdsService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "google.ads.googleads.v16.services.GoogleAdsService",
	HandlerType: (*GoogleAdsServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Search",
			Handler:    _GoogleAdsService_Search_Handler,
		},
		{
			MethodName: "Mutate",
			Handler:    _GoogleAdsService_Mutate_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "SearchStream",
			Handler:       _GoogleAdsService_SearchStream_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "google/ads/googleads/v16/services/google_ads_service.proto",
}
