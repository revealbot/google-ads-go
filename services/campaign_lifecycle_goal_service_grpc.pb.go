// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.21.12
// source: google/ads/googleads/v15/services/campaign_lifecycle_goal_service.proto

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

// CampaignLifecycleServiceClient is the client API for CampaignLifecycleService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type CampaignLifecycleServiceClient interface {
	// Process the given campaign lifecycle configurations.
	//
	// List of thrown errors:
	//
	//	[AuthenticationError]()
	//	[AuthorizationError]()
	//	[CampaignLifecycleGoalConfigError]()
	//	[HeaderError]()
	//	[InternalError]()
	//	[QuotaError]()
	//	[RequestError]()
	ConfigureCampaignLifecycleGoals(ctx context.Context, in *ConfigureCampaignLifecycleGoalsRequest, opts ...grpc.CallOption) (*ConfigureCampaignLifecycleGoalsResponse, error)
}

type campaignLifecycleServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewCampaignLifecycleServiceClient(cc grpc.ClientConnInterface) CampaignLifecycleServiceClient {
	return &campaignLifecycleServiceClient{cc}
}

func (c *campaignLifecycleServiceClient) ConfigureCampaignLifecycleGoals(ctx context.Context, in *ConfigureCampaignLifecycleGoalsRequest, opts ...grpc.CallOption) (*ConfigureCampaignLifecycleGoalsResponse, error) {
	out := new(ConfigureCampaignLifecycleGoalsResponse)
	err := c.cc.Invoke(ctx, "/google.ads.googleads.v15.services.CampaignLifecycleService/ConfigureCampaignLifecycleGoals", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// CampaignLifecycleServiceServer is the server API for CampaignLifecycleService service.
// All implementations must embed UnimplementedCampaignLifecycleServiceServer
// for forward compatibility
type CampaignLifecycleServiceServer interface {
	// Process the given campaign lifecycle configurations.
	//
	// List of thrown errors:
	//
	//	[AuthenticationError]()
	//	[AuthorizationError]()
	//	[CampaignLifecycleGoalConfigError]()
	//	[HeaderError]()
	//	[InternalError]()
	//	[QuotaError]()
	//	[RequestError]()
	ConfigureCampaignLifecycleGoals(context.Context, *ConfigureCampaignLifecycleGoalsRequest) (*ConfigureCampaignLifecycleGoalsResponse, error)
	mustEmbedUnimplementedCampaignLifecycleServiceServer()
}

// UnimplementedCampaignLifecycleServiceServer must be embedded to have forward compatible implementations.
type UnimplementedCampaignLifecycleServiceServer struct {
}

func (UnimplementedCampaignLifecycleServiceServer) ConfigureCampaignLifecycleGoals(context.Context, *ConfigureCampaignLifecycleGoalsRequest) (*ConfigureCampaignLifecycleGoalsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ConfigureCampaignLifecycleGoals not implemented")
}
func (UnimplementedCampaignLifecycleServiceServer) mustEmbedUnimplementedCampaignLifecycleServiceServer() {
}

// UnsafeCampaignLifecycleServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to CampaignLifecycleServiceServer will
// result in compilation errors.
type UnsafeCampaignLifecycleServiceServer interface {
	mustEmbedUnimplementedCampaignLifecycleServiceServer()
}

func RegisterCampaignLifecycleServiceServer(s grpc.ServiceRegistrar, srv CampaignLifecycleServiceServer) {
	s.RegisterService(&CampaignLifecycleService_ServiceDesc, srv)
}

func _CampaignLifecycleService_ConfigureCampaignLifecycleGoals_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ConfigureCampaignLifecycleGoalsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CampaignLifecycleServiceServer).ConfigureCampaignLifecycleGoals(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/google.ads.googleads.v15.services.CampaignLifecycleService/ConfigureCampaignLifecycleGoals",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CampaignLifecycleServiceServer).ConfigureCampaignLifecycleGoals(ctx, req.(*ConfigureCampaignLifecycleGoalsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// CampaignLifecycleService_ServiceDesc is the grpc.ServiceDesc for CampaignLifecycleService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var CampaignLifecycleService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "google.ads.googleads.v15.services.CampaignLifecycleService",
	HandlerType: (*CampaignLifecycleServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "ConfigureCampaignLifecycleGoals",
			Handler:    _CampaignLifecycleService_ConfigureCampaignLifecycleGoals_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "google/ads/googleads/v15/services/campaign_lifecycle_goal_service.proto",
}