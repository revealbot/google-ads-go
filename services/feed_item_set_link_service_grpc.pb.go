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

// FeedItemSetLinkServiceClient is the client API for FeedItemSetLinkService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type FeedItemSetLinkServiceClient interface {
	// Returns the requested feed item set link in full detail.
	//
	// List of thrown errors:
	//   [AuthenticationError]()
	//   [AuthorizationError]()
	//   [HeaderError]()
	//   [InternalError]()
	//   [QuotaError]()
	//   [RequestError]()
	GetFeedItemSetLink(ctx context.Context, in *GetFeedItemSetLinkRequest, opts ...grpc.CallOption) (*resources.FeedItemSetLink, error)
	// Creates, updates, or removes feed item set links.
	//
	// List of thrown errors:
	//   [AuthenticationError]()
	//   [AuthorizationError]()
	//   [HeaderError]()
	//   [InternalError]()
	//   [QuotaError]()
	//   [RequestError]()
	MutateFeedItemSetLinks(ctx context.Context, in *MutateFeedItemSetLinksRequest, opts ...grpc.CallOption) (*MutateFeedItemSetLinksResponse, error)
}

type feedItemSetLinkServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewFeedItemSetLinkServiceClient(cc grpc.ClientConnInterface) FeedItemSetLinkServiceClient {
	return &feedItemSetLinkServiceClient{cc}
}

func (c *feedItemSetLinkServiceClient) GetFeedItemSetLink(ctx context.Context, in *GetFeedItemSetLinkRequest, opts ...grpc.CallOption) (*resources.FeedItemSetLink, error) {
	out := new(resources.FeedItemSetLink)
	err := c.cc.Invoke(ctx, "/google.ads.googleads.v9.services.FeedItemSetLinkService/GetFeedItemSetLink", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *feedItemSetLinkServiceClient) MutateFeedItemSetLinks(ctx context.Context, in *MutateFeedItemSetLinksRequest, opts ...grpc.CallOption) (*MutateFeedItemSetLinksResponse, error) {
	out := new(MutateFeedItemSetLinksResponse)
	err := c.cc.Invoke(ctx, "/google.ads.googleads.v9.services.FeedItemSetLinkService/MutateFeedItemSetLinks", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// FeedItemSetLinkServiceServer is the server API for FeedItemSetLinkService service.
// All implementations must embed UnimplementedFeedItemSetLinkServiceServer
// for forward compatibility
type FeedItemSetLinkServiceServer interface {
	// Returns the requested feed item set link in full detail.
	//
	// List of thrown errors:
	//   [AuthenticationError]()
	//   [AuthorizationError]()
	//   [HeaderError]()
	//   [InternalError]()
	//   [QuotaError]()
	//   [RequestError]()
	GetFeedItemSetLink(context.Context, *GetFeedItemSetLinkRequest) (*resources.FeedItemSetLink, error)
	// Creates, updates, or removes feed item set links.
	//
	// List of thrown errors:
	//   [AuthenticationError]()
	//   [AuthorizationError]()
	//   [HeaderError]()
	//   [InternalError]()
	//   [QuotaError]()
	//   [RequestError]()
	MutateFeedItemSetLinks(context.Context, *MutateFeedItemSetLinksRequest) (*MutateFeedItemSetLinksResponse, error)
	mustEmbedUnimplementedFeedItemSetLinkServiceServer()
}

// UnimplementedFeedItemSetLinkServiceServer must be embedded to have forward compatible implementations.
type UnimplementedFeedItemSetLinkServiceServer struct {
}

func (UnimplementedFeedItemSetLinkServiceServer) GetFeedItemSetLink(context.Context, *GetFeedItemSetLinkRequest) (*resources.FeedItemSetLink, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetFeedItemSetLink not implemented")
}
func (UnimplementedFeedItemSetLinkServiceServer) MutateFeedItemSetLinks(context.Context, *MutateFeedItemSetLinksRequest) (*MutateFeedItemSetLinksResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method MutateFeedItemSetLinks not implemented")
}
func (UnimplementedFeedItemSetLinkServiceServer) mustEmbedUnimplementedFeedItemSetLinkServiceServer() {
}

// UnsafeFeedItemSetLinkServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to FeedItemSetLinkServiceServer will
// result in compilation errors.
type UnsafeFeedItemSetLinkServiceServer interface {
	mustEmbedUnimplementedFeedItemSetLinkServiceServer()
}

func RegisterFeedItemSetLinkServiceServer(s grpc.ServiceRegistrar, srv FeedItemSetLinkServiceServer) {
	s.RegisterService(&FeedItemSetLinkService_ServiceDesc, srv)
}

func _FeedItemSetLinkService_GetFeedItemSetLink_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetFeedItemSetLinkRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(FeedItemSetLinkServiceServer).GetFeedItemSetLink(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/google.ads.googleads.v9.services.FeedItemSetLinkService/GetFeedItemSetLink",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(FeedItemSetLinkServiceServer).GetFeedItemSetLink(ctx, req.(*GetFeedItemSetLinkRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _FeedItemSetLinkService_MutateFeedItemSetLinks_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MutateFeedItemSetLinksRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(FeedItemSetLinkServiceServer).MutateFeedItemSetLinks(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/google.ads.googleads.v9.services.FeedItemSetLinkService/MutateFeedItemSetLinks",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(FeedItemSetLinkServiceServer).MutateFeedItemSetLinks(ctx, req.(*MutateFeedItemSetLinksRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// FeedItemSetLinkService_ServiceDesc is the grpc.ServiceDesc for FeedItemSetLinkService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var FeedItemSetLinkService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "google.ads.googleads.v9.services.FeedItemSetLinkService",
	HandlerType: (*FeedItemSetLinkServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetFeedItemSetLink",
			Handler:    _FeedItemSetLinkService_GetFeedItemSetLink_Handler,
		},
		{
			MethodName: "MutateFeedItemSetLinks",
			Handler:    _FeedItemSetLinkService_MutateFeedItemSetLinks_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "google/ads/googleads/v9/services/feed_item_set_link_service.proto",
}
