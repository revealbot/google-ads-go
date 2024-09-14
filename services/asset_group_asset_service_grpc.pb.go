// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.21.12
// source: google/ads/googleads/v17/services/asset_group_asset_service.proto

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

// AssetGroupAssetServiceClient is the client API for AssetGroupAssetService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type AssetGroupAssetServiceClient interface {
	// Creates, updates or removes asset group assets. Operation statuses are
	// returned.
	MutateAssetGroupAssets(ctx context.Context, in *MutateAssetGroupAssetsRequest, opts ...grpc.CallOption) (*MutateAssetGroupAssetsResponse, error)
}

type assetGroupAssetServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewAssetGroupAssetServiceClient(cc grpc.ClientConnInterface) AssetGroupAssetServiceClient {
	return &assetGroupAssetServiceClient{cc}
}

func (c *assetGroupAssetServiceClient) MutateAssetGroupAssets(ctx context.Context, in *MutateAssetGroupAssetsRequest, opts ...grpc.CallOption) (*MutateAssetGroupAssetsResponse, error) {
	out := new(MutateAssetGroupAssetsResponse)
	err := c.cc.Invoke(ctx, "/google.ads.googleads.v17.services.AssetGroupAssetService/MutateAssetGroupAssets", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// AssetGroupAssetServiceServer is the server API for AssetGroupAssetService service.
// All implementations must embed UnimplementedAssetGroupAssetServiceServer
// for forward compatibility
type AssetGroupAssetServiceServer interface {
	// Creates, updates or removes asset group assets. Operation statuses are
	// returned.
	MutateAssetGroupAssets(context.Context, *MutateAssetGroupAssetsRequest) (*MutateAssetGroupAssetsResponse, error)
	mustEmbedUnimplementedAssetGroupAssetServiceServer()
}

// UnimplementedAssetGroupAssetServiceServer must be embedded to have forward compatible implementations.
type UnimplementedAssetGroupAssetServiceServer struct {
}

func (UnimplementedAssetGroupAssetServiceServer) MutateAssetGroupAssets(context.Context, *MutateAssetGroupAssetsRequest) (*MutateAssetGroupAssetsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method MutateAssetGroupAssets not implemented")
}
func (UnimplementedAssetGroupAssetServiceServer) mustEmbedUnimplementedAssetGroupAssetServiceServer() {
}

// UnsafeAssetGroupAssetServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to AssetGroupAssetServiceServer will
// result in compilation errors.
type UnsafeAssetGroupAssetServiceServer interface {
	mustEmbedUnimplementedAssetGroupAssetServiceServer()
}

func RegisterAssetGroupAssetServiceServer(s grpc.ServiceRegistrar, srv AssetGroupAssetServiceServer) {
	s.RegisterService(&AssetGroupAssetService_ServiceDesc, srv)
}

func _AssetGroupAssetService_MutateAssetGroupAssets_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MutateAssetGroupAssetsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AssetGroupAssetServiceServer).MutateAssetGroupAssets(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/google.ads.googleads.v17.services.AssetGroupAssetService/MutateAssetGroupAssets",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AssetGroupAssetServiceServer).MutateAssetGroupAssets(ctx, req.(*MutateAssetGroupAssetsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// AssetGroupAssetService_ServiceDesc is the grpc.ServiceDesc for AssetGroupAssetService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var AssetGroupAssetService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "google.ads.googleads.v17.services.AssetGroupAssetService",
	HandlerType: (*AssetGroupAssetServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "MutateAssetGroupAssets",
			Handler:    _AssetGroupAssetService_MutateAssetGroupAssets_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "google/ads/googleads/v17/services/asset_group_asset_service.proto",
}
