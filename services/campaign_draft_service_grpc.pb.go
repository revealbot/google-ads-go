// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.21.12
// source: google/ads/googleads/v16/services/campaign_draft_service.proto

package services

import (
	longrunningpb "cloud.google.com/go/longrunning/autogen/longrunningpb"
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

// CampaignDraftServiceClient is the client API for CampaignDraftService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type CampaignDraftServiceClient interface {
	// Creates, updates, or removes campaign drafts. Operation statuses are
	// returned.
	//
	// List of thrown errors:
	//
	//	[AuthenticationError]()
	//	[AuthorizationError]()
	//	[CampaignDraftError]()
	//	[DatabaseError]()
	//	[FieldError]()
	//	[HeaderError]()
	//	[InternalError]()
	//	[MutateError]()
	//	[QuotaError]()
	//	[RequestError]()
	MutateCampaignDrafts(ctx context.Context, in *MutateCampaignDraftsRequest, opts ...grpc.CallOption) (*MutateCampaignDraftsResponse, error)
	// Promotes the changes in a draft back to the base campaign.
	//
	// This method returns a Long Running Operation (LRO) indicating if the
	// Promote is done. Use [Operations.GetOperation] to poll the LRO until it
	// is done. Only a done status is returned in the response. See the status
	// in the Campaign Draft resource to determine if the promotion was
	// successful. If the LRO failed, use
	// [CampaignDraftService.ListCampaignDraftAsyncErrors][google.ads.googleads.v16.services.CampaignDraftService.ListCampaignDraftAsyncErrors]
	// to view the list of error reasons.
	//
	// List of thrown errors:
	//
	//	[AuthenticationError]()
	//	[AuthorizationError]()
	//	[CampaignDraftError]()
	//	[HeaderError]()
	//	[InternalError]()
	//	[QuotaError]()
	//	[RequestError]()
	PromoteCampaignDraft(ctx context.Context, in *PromoteCampaignDraftRequest, opts ...grpc.CallOption) (*longrunningpb.Operation, error)
	// Returns all errors that occurred during CampaignDraft promote. Throws an
	// error if called before campaign draft is promoted.
	// Supports standard list paging.
	//
	// List of thrown errors:
	//
	//	[AuthenticationError]()
	//	[AuthorizationError]()
	//	[HeaderError]()
	//	[InternalError]()
	//	[QuotaError]()
	//	[RequestError]()
	ListCampaignDraftAsyncErrors(ctx context.Context, in *ListCampaignDraftAsyncErrorsRequest, opts ...grpc.CallOption) (*ListCampaignDraftAsyncErrorsResponse, error)
}

type campaignDraftServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewCampaignDraftServiceClient(cc grpc.ClientConnInterface) CampaignDraftServiceClient {
	return &campaignDraftServiceClient{cc}
}

func (c *campaignDraftServiceClient) MutateCampaignDrafts(ctx context.Context, in *MutateCampaignDraftsRequest, opts ...grpc.CallOption) (*MutateCampaignDraftsResponse, error) {
	out := new(MutateCampaignDraftsResponse)
	err := c.cc.Invoke(ctx, "/google.ads.googleads.v16.services.CampaignDraftService/MutateCampaignDrafts", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *campaignDraftServiceClient) PromoteCampaignDraft(ctx context.Context, in *PromoteCampaignDraftRequest, opts ...grpc.CallOption) (*longrunningpb.Operation, error) {
	out := new(longrunningpb.Operation)
	err := c.cc.Invoke(ctx, "/google.ads.googleads.v16.services.CampaignDraftService/PromoteCampaignDraft", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *campaignDraftServiceClient) ListCampaignDraftAsyncErrors(ctx context.Context, in *ListCampaignDraftAsyncErrorsRequest, opts ...grpc.CallOption) (*ListCampaignDraftAsyncErrorsResponse, error) {
	out := new(ListCampaignDraftAsyncErrorsResponse)
	err := c.cc.Invoke(ctx, "/google.ads.googleads.v16.services.CampaignDraftService/ListCampaignDraftAsyncErrors", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// CampaignDraftServiceServer is the server API for CampaignDraftService service.
// All implementations must embed UnimplementedCampaignDraftServiceServer
// for forward compatibility
type CampaignDraftServiceServer interface {
	// Creates, updates, or removes campaign drafts. Operation statuses are
	// returned.
	//
	// List of thrown errors:
	//
	//	[AuthenticationError]()
	//	[AuthorizationError]()
	//	[CampaignDraftError]()
	//	[DatabaseError]()
	//	[FieldError]()
	//	[HeaderError]()
	//	[InternalError]()
	//	[MutateError]()
	//	[QuotaError]()
	//	[RequestError]()
	MutateCampaignDrafts(context.Context, *MutateCampaignDraftsRequest) (*MutateCampaignDraftsResponse, error)
	// Promotes the changes in a draft back to the base campaign.
	//
	// This method returns a Long Running Operation (LRO) indicating if the
	// Promote is done. Use [Operations.GetOperation] to poll the LRO until it
	// is done. Only a done status is returned in the response. See the status
	// in the Campaign Draft resource to determine if the promotion was
	// successful. If the LRO failed, use
	// [CampaignDraftService.ListCampaignDraftAsyncErrors][google.ads.googleads.v16.services.CampaignDraftService.ListCampaignDraftAsyncErrors]
	// to view the list of error reasons.
	//
	// List of thrown errors:
	//
	//	[AuthenticationError]()
	//	[AuthorizationError]()
	//	[CampaignDraftError]()
	//	[HeaderError]()
	//	[InternalError]()
	//	[QuotaError]()
	//	[RequestError]()
	PromoteCampaignDraft(context.Context, *PromoteCampaignDraftRequest) (*longrunningpb.Operation, error)
	// Returns all errors that occurred during CampaignDraft promote. Throws an
	// error if called before campaign draft is promoted.
	// Supports standard list paging.
	//
	// List of thrown errors:
	//
	//	[AuthenticationError]()
	//	[AuthorizationError]()
	//	[HeaderError]()
	//	[InternalError]()
	//	[QuotaError]()
	//	[RequestError]()
	ListCampaignDraftAsyncErrors(context.Context, *ListCampaignDraftAsyncErrorsRequest) (*ListCampaignDraftAsyncErrorsResponse, error)
	mustEmbedUnimplementedCampaignDraftServiceServer()
}

// UnimplementedCampaignDraftServiceServer must be embedded to have forward compatible implementations.
type UnimplementedCampaignDraftServiceServer struct {
}

func (UnimplementedCampaignDraftServiceServer) MutateCampaignDrafts(context.Context, *MutateCampaignDraftsRequest) (*MutateCampaignDraftsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method MutateCampaignDrafts not implemented")
}
func (UnimplementedCampaignDraftServiceServer) PromoteCampaignDraft(context.Context, *PromoteCampaignDraftRequest) (*longrunningpb.Operation, error) {
	return nil, status.Errorf(codes.Unimplemented, "method PromoteCampaignDraft not implemented")
}
func (UnimplementedCampaignDraftServiceServer) ListCampaignDraftAsyncErrors(context.Context, *ListCampaignDraftAsyncErrorsRequest) (*ListCampaignDraftAsyncErrorsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListCampaignDraftAsyncErrors not implemented")
}
func (UnimplementedCampaignDraftServiceServer) mustEmbedUnimplementedCampaignDraftServiceServer() {}

// UnsafeCampaignDraftServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to CampaignDraftServiceServer will
// result in compilation errors.
type UnsafeCampaignDraftServiceServer interface {
	mustEmbedUnimplementedCampaignDraftServiceServer()
}

func RegisterCampaignDraftServiceServer(s grpc.ServiceRegistrar, srv CampaignDraftServiceServer) {
	s.RegisterService(&CampaignDraftService_ServiceDesc, srv)
}

func _CampaignDraftService_MutateCampaignDrafts_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MutateCampaignDraftsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CampaignDraftServiceServer).MutateCampaignDrafts(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/google.ads.googleads.v16.services.CampaignDraftService/MutateCampaignDrafts",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CampaignDraftServiceServer).MutateCampaignDrafts(ctx, req.(*MutateCampaignDraftsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _CampaignDraftService_PromoteCampaignDraft_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PromoteCampaignDraftRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CampaignDraftServiceServer).PromoteCampaignDraft(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/google.ads.googleads.v16.services.CampaignDraftService/PromoteCampaignDraft",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CampaignDraftServiceServer).PromoteCampaignDraft(ctx, req.(*PromoteCampaignDraftRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _CampaignDraftService_ListCampaignDraftAsyncErrors_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListCampaignDraftAsyncErrorsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CampaignDraftServiceServer).ListCampaignDraftAsyncErrors(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/google.ads.googleads.v16.services.CampaignDraftService/ListCampaignDraftAsyncErrors",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CampaignDraftServiceServer).ListCampaignDraftAsyncErrors(ctx, req.(*ListCampaignDraftAsyncErrorsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// CampaignDraftService_ServiceDesc is the grpc.ServiceDesc for CampaignDraftService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var CampaignDraftService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "google.ads.googleads.v16.services.CampaignDraftService",
	HandlerType: (*CampaignDraftServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "MutateCampaignDrafts",
			Handler:    _CampaignDraftService_MutateCampaignDrafts_Handler,
		},
		{
			MethodName: "PromoteCampaignDraft",
			Handler:    _CampaignDraftService_PromoteCampaignDraft_Handler,
		},
		{
			MethodName: "ListCampaignDraftAsyncErrors",
			Handler:    _CampaignDraftService_ListCampaignDraftAsyncErrors_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "google/ads/googleads/v16/services/campaign_draft_service.proto",
}
