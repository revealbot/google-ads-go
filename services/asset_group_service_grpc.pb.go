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

// AssetGroupServiceClient is the client API for AssetGroupService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type AssetGroupServiceClient interface {
	// Creates, updates or removes asset groups. Operation statuses are
	// returned.
	MutateAssetGroups(ctx context.Context, in *MutateAssetGroupsRequest, opts ...grpc.CallOption) (*MutateAssetGroupsResponse, error)
}

type assetGroupServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewAssetGroupServiceClient(cc grpc.ClientConnInterface) AssetGroupServiceClient {
	return &assetGroupServiceClient{cc}
}

func (c *assetGroupServiceClient) MutateAssetGroups(ctx context.Context, in *MutateAssetGroupsRequest, opts ...grpc.CallOption) (*MutateAssetGroupsResponse, error) {
	out := new(MutateAssetGroupsResponse)
	err := c.cc.Invoke(ctx, "/google.ads.googleads.v13.services.AssetGroupService/MutateAssetGroups", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// AssetGroupServiceServer is the server API for AssetGroupService service.
// All implementations must embed UnimplementedAssetGroupServiceServer
// for forward compatibility
type AssetGroupServiceServer interface {
	// Creates, updates or removes asset groups. Operation statuses are
	// returned.
	MutateAssetGroups(context.Context, *MutateAssetGroupsRequest) (*MutateAssetGroupsResponse, error)
	mustEmbedUnimplementedAssetGroupServiceServer()
}

// UnimplementedAssetGroupServiceServer must be embedded to have forward compatible implementations.
type UnimplementedAssetGroupServiceServer struct {
}

func (UnimplementedAssetGroupServiceServer) MutateAssetGroups(context.Context, *MutateAssetGroupsRequest) (*MutateAssetGroupsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method MutateAssetGroups not implemented")
}
func (UnimplementedAssetGroupServiceServer) mustEmbedUnimplementedAssetGroupServiceServer() {}

// UnsafeAssetGroupServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to AssetGroupServiceServer will
// result in compilation errors.
type UnsafeAssetGroupServiceServer interface {
	mustEmbedUnimplementedAssetGroupServiceServer()
}

func RegisterAssetGroupServiceServer(s grpc.ServiceRegistrar, srv AssetGroupServiceServer) {
	s.RegisterService(&AssetGroupService_ServiceDesc, srv)
}

func _AssetGroupService_MutateAssetGroups_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MutateAssetGroupsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AssetGroupServiceServer).MutateAssetGroups(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/google.ads.googleads.v13.services.AssetGroupService/MutateAssetGroups",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AssetGroupServiceServer).MutateAssetGroups(ctx, req.(*MutateAssetGroupsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// AssetGroupService_ServiceDesc is the grpc.ServiceDesc for AssetGroupService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var AssetGroupService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "google.ads.googleads.v13.services.AssetGroupService",
	HandlerType: (*AssetGroupServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "MutateAssetGroups",
			Handler:    _AssetGroupService_MutateAssetGroups_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "google/ads/googleads/v13/services/asset_group_service.proto",
}
