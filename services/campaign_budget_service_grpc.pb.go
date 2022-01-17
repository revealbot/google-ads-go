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

// CampaignBudgetServiceClient is the client API for CampaignBudgetService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type CampaignBudgetServiceClient interface {
	// Returns the requested Campaign Budget in full detail.
	//
	// List of thrown errors:
	//   [AuthenticationError]()
	//   [AuthorizationError]()
	//   [HeaderError]()
	//   [InternalError]()
	//   [QuotaError]()
	//   [RequestError]()
	GetCampaignBudget(ctx context.Context, in *GetCampaignBudgetRequest, opts ...grpc.CallOption) (*resources.CampaignBudget, error)
	// Creates, updates, or removes campaign budgets. Operation statuses are
	// returned.
	//
	// List of thrown errors:
	//   [AuthenticationError]()
	//   [AuthorizationError]()
	//   [CampaignBudgetError]()
	//   [DatabaseError]()
	//   [DistinctError]()
	//   [FieldError]()
	//   [FieldMaskError]()
	//   [HeaderError]()
	//   [InternalError]()
	//   [MutateError]()
	//   [NewResourceCreationError]()
	//   [OperationAccessDeniedError]()
	//   [QuotaError]()
	//   [RangeError]()
	//   [RequestError]()
	//   [ResourceCountLimitExceededError]()
	//   [StringLengthError]()
	MutateCampaignBudgets(ctx context.Context, in *MutateCampaignBudgetsRequest, opts ...grpc.CallOption) (*MutateCampaignBudgetsResponse, error)
}

type campaignBudgetServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewCampaignBudgetServiceClient(cc grpc.ClientConnInterface) CampaignBudgetServiceClient {
	return &campaignBudgetServiceClient{cc}
}

func (c *campaignBudgetServiceClient) GetCampaignBudget(ctx context.Context, in *GetCampaignBudgetRequest, opts ...grpc.CallOption) (*resources.CampaignBudget, error) {
	out := new(resources.CampaignBudget)
	err := c.cc.Invoke(ctx, "/google.ads.googleads.v9.services.CampaignBudgetService/GetCampaignBudget", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *campaignBudgetServiceClient) MutateCampaignBudgets(ctx context.Context, in *MutateCampaignBudgetsRequest, opts ...grpc.CallOption) (*MutateCampaignBudgetsResponse, error) {
	out := new(MutateCampaignBudgetsResponse)
	err := c.cc.Invoke(ctx, "/google.ads.googleads.v9.services.CampaignBudgetService/MutateCampaignBudgets", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// CampaignBudgetServiceServer is the server API for CampaignBudgetService service.
// All implementations must embed UnimplementedCampaignBudgetServiceServer
// for forward compatibility
type CampaignBudgetServiceServer interface {
	// Returns the requested Campaign Budget in full detail.
	//
	// List of thrown errors:
	//   [AuthenticationError]()
	//   [AuthorizationError]()
	//   [HeaderError]()
	//   [InternalError]()
	//   [QuotaError]()
	//   [RequestError]()
	GetCampaignBudget(context.Context, *GetCampaignBudgetRequest) (*resources.CampaignBudget, error)
	// Creates, updates, or removes campaign budgets. Operation statuses are
	// returned.
	//
	// List of thrown errors:
	//   [AuthenticationError]()
	//   [AuthorizationError]()
	//   [CampaignBudgetError]()
	//   [DatabaseError]()
	//   [DistinctError]()
	//   [FieldError]()
	//   [FieldMaskError]()
	//   [HeaderError]()
	//   [InternalError]()
	//   [MutateError]()
	//   [NewResourceCreationError]()
	//   [OperationAccessDeniedError]()
	//   [QuotaError]()
	//   [RangeError]()
	//   [RequestError]()
	//   [ResourceCountLimitExceededError]()
	//   [StringLengthError]()
	MutateCampaignBudgets(context.Context, *MutateCampaignBudgetsRequest) (*MutateCampaignBudgetsResponse, error)
	mustEmbedUnimplementedCampaignBudgetServiceServer()
}

// UnimplementedCampaignBudgetServiceServer must be embedded to have forward compatible implementations.
type UnimplementedCampaignBudgetServiceServer struct {
}

func (UnimplementedCampaignBudgetServiceServer) GetCampaignBudget(context.Context, *GetCampaignBudgetRequest) (*resources.CampaignBudget, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetCampaignBudget not implemented")
}
func (UnimplementedCampaignBudgetServiceServer) MutateCampaignBudgets(context.Context, *MutateCampaignBudgetsRequest) (*MutateCampaignBudgetsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method MutateCampaignBudgets not implemented")
}
func (UnimplementedCampaignBudgetServiceServer) mustEmbedUnimplementedCampaignBudgetServiceServer() {}

// UnsafeCampaignBudgetServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to CampaignBudgetServiceServer will
// result in compilation errors.
type UnsafeCampaignBudgetServiceServer interface {
	mustEmbedUnimplementedCampaignBudgetServiceServer()
}

func RegisterCampaignBudgetServiceServer(s grpc.ServiceRegistrar, srv CampaignBudgetServiceServer) {
	s.RegisterService(&CampaignBudgetService_ServiceDesc, srv)
}

func _CampaignBudgetService_GetCampaignBudget_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetCampaignBudgetRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CampaignBudgetServiceServer).GetCampaignBudget(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/google.ads.googleads.v9.services.CampaignBudgetService/GetCampaignBudget",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CampaignBudgetServiceServer).GetCampaignBudget(ctx, req.(*GetCampaignBudgetRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _CampaignBudgetService_MutateCampaignBudgets_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MutateCampaignBudgetsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CampaignBudgetServiceServer).MutateCampaignBudgets(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/google.ads.googleads.v9.services.CampaignBudgetService/MutateCampaignBudgets",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CampaignBudgetServiceServer).MutateCampaignBudgets(ctx, req.(*MutateCampaignBudgetsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// CampaignBudgetService_ServiceDesc is the grpc.ServiceDesc for CampaignBudgetService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var CampaignBudgetService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "google.ads.googleads.v9.services.CampaignBudgetService",
	HandlerType: (*CampaignBudgetServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetCampaignBudget",
			Handler:    _CampaignBudgetService_GetCampaignBudget_Handler,
		},
		{
			MethodName: "MutateCampaignBudgets",
			Handler:    _CampaignBudgetService_MutateCampaignBudgets_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "google/ads/googleads/v9/services/campaign_budget_service.proto",
}
