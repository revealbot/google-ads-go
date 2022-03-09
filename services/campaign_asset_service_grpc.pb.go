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

// CampaignAssetServiceClient is the client API for CampaignAssetService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type CampaignAssetServiceClient interface {
	// Creates, updates, or removes campaign assets. Operation statuses are
	// returned.
	//
	// List of thrown errors:
	//   [AssetLinkError]()
	//   [AuthenticationError]()
	//   [AuthorizationError]()
	//   [ContextError]()
	//   [DatabaseError]()
	//   [FieldError]()
	//   [HeaderError]()
	//   [InternalError]()
	//   [MutateError]()
	//   [NotAllowlistedError]()
	//   [QuotaError]()
	//   [RequestError]()
	MutateCampaignAssets(ctx context.Context, in *MutateCampaignAssetsRequest, opts ...grpc.CallOption) (*MutateCampaignAssetsResponse, error)
}

type campaignAssetServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewCampaignAssetServiceClient(cc grpc.ClientConnInterface) CampaignAssetServiceClient {
	return &campaignAssetServiceClient{cc}
}

func (c *campaignAssetServiceClient) MutateCampaignAssets(ctx context.Context, in *MutateCampaignAssetsRequest, opts ...grpc.CallOption) (*MutateCampaignAssetsResponse, error) {
	out := new(MutateCampaignAssetsResponse)
	err := c.cc.Invoke(ctx, "/google.ads.googleads.v10.services.CampaignAssetService/MutateCampaignAssets", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// CampaignAssetServiceServer is the server API for CampaignAssetService service.
// All implementations must embed UnimplementedCampaignAssetServiceServer
// for forward compatibility
type CampaignAssetServiceServer interface {
	// Creates, updates, or removes campaign assets. Operation statuses are
	// returned.
	//
	// List of thrown errors:
	//   [AssetLinkError]()
	//   [AuthenticationError]()
	//   [AuthorizationError]()
	//   [ContextError]()
	//   [DatabaseError]()
	//   [FieldError]()
	//   [HeaderError]()
	//   [InternalError]()
	//   [MutateError]()
	//   [NotAllowlistedError]()
	//   [QuotaError]()
	//   [RequestError]()
	MutateCampaignAssets(context.Context, *MutateCampaignAssetsRequest) (*MutateCampaignAssetsResponse, error)
	mustEmbedUnimplementedCampaignAssetServiceServer()
}

// UnimplementedCampaignAssetServiceServer must be embedded to have forward compatible implementations.
type UnimplementedCampaignAssetServiceServer struct {
}

func (UnimplementedCampaignAssetServiceServer) MutateCampaignAssets(context.Context, *MutateCampaignAssetsRequest) (*MutateCampaignAssetsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method MutateCampaignAssets not implemented")
}
func (UnimplementedCampaignAssetServiceServer) mustEmbedUnimplementedCampaignAssetServiceServer() {}

// UnsafeCampaignAssetServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to CampaignAssetServiceServer will
// result in compilation errors.
type UnsafeCampaignAssetServiceServer interface {
	mustEmbedUnimplementedCampaignAssetServiceServer()
}

func RegisterCampaignAssetServiceServer(s grpc.ServiceRegistrar, srv CampaignAssetServiceServer) {
	s.RegisterService(&CampaignAssetService_ServiceDesc, srv)
}

func _CampaignAssetService_MutateCampaignAssets_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MutateCampaignAssetsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CampaignAssetServiceServer).MutateCampaignAssets(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/google.ads.googleads.v10.services.CampaignAssetService/MutateCampaignAssets",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CampaignAssetServiceServer).MutateCampaignAssets(ctx, req.(*MutateCampaignAssetsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// CampaignAssetService_ServiceDesc is the grpc.ServiceDesc for CampaignAssetService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var CampaignAssetService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "google.ads.googleads.v10.services.CampaignAssetService",
	HandlerType: (*CampaignAssetServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "MutateCampaignAssets",
			Handler:    _CampaignAssetService_MutateCampaignAssets_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "google/ads/googleads/v10/services/campaign_asset_service.proto",
}
