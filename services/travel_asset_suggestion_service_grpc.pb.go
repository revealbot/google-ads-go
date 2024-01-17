// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.21.12
// source: google/ads/googleads/v15/services/travel_asset_suggestion_service.proto

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

// TravelAssetSuggestionServiceClient is the client API for TravelAssetSuggestionService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type TravelAssetSuggestionServiceClient interface {
	// Returns Travel Asset suggestions. Asset
	// suggestions are returned on a best-effort basis. There are no guarantees
	// that all possible asset types will be returned for any given hotel
	// property.
	SuggestTravelAssets(ctx context.Context, in *SuggestTravelAssetsRequest, opts ...grpc.CallOption) (*SuggestTravelAssetsResponse, error)
}

type travelAssetSuggestionServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewTravelAssetSuggestionServiceClient(cc grpc.ClientConnInterface) TravelAssetSuggestionServiceClient {
	return &travelAssetSuggestionServiceClient{cc}
}

func (c *travelAssetSuggestionServiceClient) SuggestTravelAssets(ctx context.Context, in *SuggestTravelAssetsRequest, opts ...grpc.CallOption) (*SuggestTravelAssetsResponse, error) {
	out := new(SuggestTravelAssetsResponse)
	err := c.cc.Invoke(ctx, "/google.ads.googleads.v15.services.TravelAssetSuggestionService/SuggestTravelAssets", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// TravelAssetSuggestionServiceServer is the server API for TravelAssetSuggestionService service.
// All implementations must embed UnimplementedTravelAssetSuggestionServiceServer
// for forward compatibility
type TravelAssetSuggestionServiceServer interface {
	// Returns Travel Asset suggestions. Asset
	// suggestions are returned on a best-effort basis. There are no guarantees
	// that all possible asset types will be returned for any given hotel
	// property.
	SuggestTravelAssets(context.Context, *SuggestTravelAssetsRequest) (*SuggestTravelAssetsResponse, error)
	mustEmbedUnimplementedTravelAssetSuggestionServiceServer()
}

// UnimplementedTravelAssetSuggestionServiceServer must be embedded to have forward compatible implementations.
type UnimplementedTravelAssetSuggestionServiceServer struct {
}

func (UnimplementedTravelAssetSuggestionServiceServer) SuggestTravelAssets(context.Context, *SuggestTravelAssetsRequest) (*SuggestTravelAssetsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SuggestTravelAssets not implemented")
}
func (UnimplementedTravelAssetSuggestionServiceServer) mustEmbedUnimplementedTravelAssetSuggestionServiceServer() {
}

// UnsafeTravelAssetSuggestionServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to TravelAssetSuggestionServiceServer will
// result in compilation errors.
type UnsafeTravelAssetSuggestionServiceServer interface {
	mustEmbedUnimplementedTravelAssetSuggestionServiceServer()
}

func RegisterTravelAssetSuggestionServiceServer(s grpc.ServiceRegistrar, srv TravelAssetSuggestionServiceServer) {
	s.RegisterService(&TravelAssetSuggestionService_ServiceDesc, srv)
}

func _TravelAssetSuggestionService_SuggestTravelAssets_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SuggestTravelAssetsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TravelAssetSuggestionServiceServer).SuggestTravelAssets(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/google.ads.googleads.v15.services.TravelAssetSuggestionService/SuggestTravelAssets",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TravelAssetSuggestionServiceServer).SuggestTravelAssets(ctx, req.(*SuggestTravelAssetsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// TravelAssetSuggestionService_ServiceDesc is the grpc.ServiceDesc for TravelAssetSuggestionService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var TravelAssetSuggestionService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "google.ads.googleads.v15.services.TravelAssetSuggestionService",
	HandlerType: (*TravelAssetSuggestionServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "SuggestTravelAssets",
			Handler:    _TravelAssetSuggestionService_SuggestTravelAssets_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "google/ads/googleads/v15/services/travel_asset_suggestion_service.proto",
}
