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

// RemarketingActionServiceClient is the client API for RemarketingActionService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type RemarketingActionServiceClient interface {
	// Creates or updates remarketing actions. Operation statuses are returned.
	//
	// List of thrown errors:
	//   [AuthenticationError]()
	//   [AuthorizationError]()
	//   [ConversionActionError]()
	//   [HeaderError]()
	//   [InternalError]()
	//   [QuotaError]()
	//   [RequestError]()
	MutateRemarketingActions(ctx context.Context, in *MutateRemarketingActionsRequest, opts ...grpc.CallOption) (*MutateRemarketingActionsResponse, error)
}

type remarketingActionServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewRemarketingActionServiceClient(cc grpc.ClientConnInterface) RemarketingActionServiceClient {
	return &remarketingActionServiceClient{cc}
}

func (c *remarketingActionServiceClient) MutateRemarketingActions(ctx context.Context, in *MutateRemarketingActionsRequest, opts ...grpc.CallOption) (*MutateRemarketingActionsResponse, error) {
	out := new(MutateRemarketingActionsResponse)
	err := c.cc.Invoke(ctx, "/google.ads.googleads.v13.services.RemarketingActionService/MutateRemarketingActions", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// RemarketingActionServiceServer is the server API for RemarketingActionService service.
// All implementations must embed UnimplementedRemarketingActionServiceServer
// for forward compatibility
type RemarketingActionServiceServer interface {
	// Creates or updates remarketing actions. Operation statuses are returned.
	//
	// List of thrown errors:
	//   [AuthenticationError]()
	//   [AuthorizationError]()
	//   [ConversionActionError]()
	//   [HeaderError]()
	//   [InternalError]()
	//   [QuotaError]()
	//   [RequestError]()
	MutateRemarketingActions(context.Context, *MutateRemarketingActionsRequest) (*MutateRemarketingActionsResponse, error)
	mustEmbedUnimplementedRemarketingActionServiceServer()
}

// UnimplementedRemarketingActionServiceServer must be embedded to have forward compatible implementations.
type UnimplementedRemarketingActionServiceServer struct {
}

func (UnimplementedRemarketingActionServiceServer) MutateRemarketingActions(context.Context, *MutateRemarketingActionsRequest) (*MutateRemarketingActionsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method MutateRemarketingActions not implemented")
}
func (UnimplementedRemarketingActionServiceServer) mustEmbedUnimplementedRemarketingActionServiceServer() {
}

// UnsafeRemarketingActionServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to RemarketingActionServiceServer will
// result in compilation errors.
type UnsafeRemarketingActionServiceServer interface {
	mustEmbedUnimplementedRemarketingActionServiceServer()
}

func RegisterRemarketingActionServiceServer(s grpc.ServiceRegistrar, srv RemarketingActionServiceServer) {
	s.RegisterService(&RemarketingActionService_ServiceDesc, srv)
}

func _RemarketingActionService_MutateRemarketingActions_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MutateRemarketingActionsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RemarketingActionServiceServer).MutateRemarketingActions(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/google.ads.googleads.v13.services.RemarketingActionService/MutateRemarketingActions",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RemarketingActionServiceServer).MutateRemarketingActions(ctx, req.(*MutateRemarketingActionsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// RemarketingActionService_ServiceDesc is the grpc.ServiceDesc for RemarketingActionService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var RemarketingActionService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "google.ads.googleads.v13.services.RemarketingActionService",
	HandlerType: (*RemarketingActionServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "MutateRemarketingActions",
			Handler:    _RemarketingActionService_MutateRemarketingActions_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "google/ads/googleads/v13/services/remarketing_action_service.proto",
}
