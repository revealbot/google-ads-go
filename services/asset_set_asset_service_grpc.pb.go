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

// AssetSetAssetServiceClient is the client API for AssetSetAssetService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type AssetSetAssetServiceClient interface {
	// Creates, updates or removes asset set assets. Operation statuses are
	// returned.
	MutateAssetSetAssets(ctx context.Context, in *MutateAssetSetAssetsRequest, opts ...grpc.CallOption) (*MutateAssetSetAssetsResponse, error)
}

type assetSetAssetServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewAssetSetAssetServiceClient(cc grpc.ClientConnInterface) AssetSetAssetServiceClient {
	return &assetSetAssetServiceClient{cc}
}

func (c *assetSetAssetServiceClient) MutateAssetSetAssets(ctx context.Context, in *MutateAssetSetAssetsRequest, opts ...grpc.CallOption) (*MutateAssetSetAssetsResponse, error) {
	out := new(MutateAssetSetAssetsResponse)
	err := c.cc.Invoke(ctx, "/google.ads.googleads.v9.services.AssetSetAssetService/MutateAssetSetAssets", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// AssetSetAssetServiceServer is the server API for AssetSetAssetService service.
// All implementations must embed UnimplementedAssetSetAssetServiceServer
// for forward compatibility
type AssetSetAssetServiceServer interface {
	// Creates, updates or removes asset set assets. Operation statuses are
	// returned.
	MutateAssetSetAssets(context.Context, *MutateAssetSetAssetsRequest) (*MutateAssetSetAssetsResponse, error)
	mustEmbedUnimplementedAssetSetAssetServiceServer()
}

// UnimplementedAssetSetAssetServiceServer must be embedded to have forward compatible implementations.
type UnimplementedAssetSetAssetServiceServer struct {
}

func (UnimplementedAssetSetAssetServiceServer) MutateAssetSetAssets(context.Context, *MutateAssetSetAssetsRequest) (*MutateAssetSetAssetsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method MutateAssetSetAssets not implemented")
}
func (UnimplementedAssetSetAssetServiceServer) mustEmbedUnimplementedAssetSetAssetServiceServer() {}

// UnsafeAssetSetAssetServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to AssetSetAssetServiceServer will
// result in compilation errors.
type UnsafeAssetSetAssetServiceServer interface {
	mustEmbedUnimplementedAssetSetAssetServiceServer()
}

func RegisterAssetSetAssetServiceServer(s grpc.ServiceRegistrar, srv AssetSetAssetServiceServer) {
	s.RegisterService(&AssetSetAssetService_ServiceDesc, srv)
}

func _AssetSetAssetService_MutateAssetSetAssets_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MutateAssetSetAssetsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AssetSetAssetServiceServer).MutateAssetSetAssets(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/google.ads.googleads.v9.services.AssetSetAssetService/MutateAssetSetAssets",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AssetSetAssetServiceServer).MutateAssetSetAssets(ctx, req.(*MutateAssetSetAssetsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// AssetSetAssetService_ServiceDesc is the grpc.ServiceDesc for AssetSetAssetService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var AssetSetAssetService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "google.ads.googleads.v9.services.AssetSetAssetService",
	HandlerType: (*AssetSetAssetServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "MutateAssetSetAssets",
			Handler:    _AssetSetAssetService_MutateAssetSetAssets_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "google/ads/googleads/v9/services/asset_set_asset_service.proto",
}
