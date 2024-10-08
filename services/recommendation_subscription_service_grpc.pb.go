// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.21.12
// source: google/ads/googleads/v17/services/recommendation_subscription_service.proto

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

// RecommendationSubscriptionServiceClient is the client API for RecommendationSubscriptionService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type RecommendationSubscriptionServiceClient interface {
	// Mutates given subscription with corresponding apply parameters.
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
	//	[RecommendationError]()
	//	[RequestError]()
	//	[UrlFieldError]()
	MutateRecommendationSubscription(ctx context.Context, in *MutateRecommendationSubscriptionRequest, opts ...grpc.CallOption) (*MutateRecommendationSubscriptionResponse, error)
}

type recommendationSubscriptionServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewRecommendationSubscriptionServiceClient(cc grpc.ClientConnInterface) RecommendationSubscriptionServiceClient {
	return &recommendationSubscriptionServiceClient{cc}
}

func (c *recommendationSubscriptionServiceClient) MutateRecommendationSubscription(ctx context.Context, in *MutateRecommendationSubscriptionRequest, opts ...grpc.CallOption) (*MutateRecommendationSubscriptionResponse, error) {
	out := new(MutateRecommendationSubscriptionResponse)
	err := c.cc.Invoke(ctx, "/google.ads.googleads.v17.services.RecommendationSubscriptionService/MutateRecommendationSubscription", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// RecommendationSubscriptionServiceServer is the server API for RecommendationSubscriptionService service.
// All implementations must embed UnimplementedRecommendationSubscriptionServiceServer
// for forward compatibility
type RecommendationSubscriptionServiceServer interface {
	// Mutates given subscription with corresponding apply parameters.
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
	//	[RecommendationError]()
	//	[RequestError]()
	//	[UrlFieldError]()
	MutateRecommendationSubscription(context.Context, *MutateRecommendationSubscriptionRequest) (*MutateRecommendationSubscriptionResponse, error)
	mustEmbedUnimplementedRecommendationSubscriptionServiceServer()
}

// UnimplementedRecommendationSubscriptionServiceServer must be embedded to have forward compatible implementations.
type UnimplementedRecommendationSubscriptionServiceServer struct {
}

func (UnimplementedRecommendationSubscriptionServiceServer) MutateRecommendationSubscription(context.Context, *MutateRecommendationSubscriptionRequest) (*MutateRecommendationSubscriptionResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method MutateRecommendationSubscription not implemented")
}
func (UnimplementedRecommendationSubscriptionServiceServer) mustEmbedUnimplementedRecommendationSubscriptionServiceServer() {
}

// UnsafeRecommendationSubscriptionServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to RecommendationSubscriptionServiceServer will
// result in compilation errors.
type UnsafeRecommendationSubscriptionServiceServer interface {
	mustEmbedUnimplementedRecommendationSubscriptionServiceServer()
}

func RegisterRecommendationSubscriptionServiceServer(s grpc.ServiceRegistrar, srv RecommendationSubscriptionServiceServer) {
	s.RegisterService(&RecommendationSubscriptionService_ServiceDesc, srv)
}

func _RecommendationSubscriptionService_MutateRecommendationSubscription_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MutateRecommendationSubscriptionRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RecommendationSubscriptionServiceServer).MutateRecommendationSubscription(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/google.ads.googleads.v17.services.RecommendationSubscriptionService/MutateRecommendationSubscription",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RecommendationSubscriptionServiceServer).MutateRecommendationSubscription(ctx, req.(*MutateRecommendationSubscriptionRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// RecommendationSubscriptionService_ServiceDesc is the grpc.ServiceDesc for RecommendationSubscriptionService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var RecommendationSubscriptionService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "google.ads.googleads.v17.services.RecommendationSubscriptionService",
	HandlerType: (*RecommendationSubscriptionServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "MutateRecommendationSubscription",
			Handler:    _RecommendationSubscriptionService_MutateRecommendationSubscription_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "google/ads/googleads/v17/services/recommendation_subscription_service.proto",
}
