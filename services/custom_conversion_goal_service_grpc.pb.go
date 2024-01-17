// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.21.12
// source: google/ads/googleads/v15/services/custom_conversion_goal_service.proto

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

// CustomConversionGoalServiceClient is the client API for CustomConversionGoalService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type CustomConversionGoalServiceClient interface {
	// Creates, updates or removes custom conversion goals. Operation statuses
	// are returned.
	MutateCustomConversionGoals(ctx context.Context, in *MutateCustomConversionGoalsRequest, opts ...grpc.CallOption) (*MutateCustomConversionGoalsResponse, error)
}

type customConversionGoalServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewCustomConversionGoalServiceClient(cc grpc.ClientConnInterface) CustomConversionGoalServiceClient {
	return &customConversionGoalServiceClient{cc}
}

func (c *customConversionGoalServiceClient) MutateCustomConversionGoals(ctx context.Context, in *MutateCustomConversionGoalsRequest, opts ...grpc.CallOption) (*MutateCustomConversionGoalsResponse, error) {
	out := new(MutateCustomConversionGoalsResponse)
	err := c.cc.Invoke(ctx, "/google.ads.googleads.v15.services.CustomConversionGoalService/MutateCustomConversionGoals", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// CustomConversionGoalServiceServer is the server API for CustomConversionGoalService service.
// All implementations must embed UnimplementedCustomConversionGoalServiceServer
// for forward compatibility
type CustomConversionGoalServiceServer interface {
	// Creates, updates or removes custom conversion goals. Operation statuses
	// are returned.
	MutateCustomConversionGoals(context.Context, *MutateCustomConversionGoalsRequest) (*MutateCustomConversionGoalsResponse, error)
	mustEmbedUnimplementedCustomConversionGoalServiceServer()
}

// UnimplementedCustomConversionGoalServiceServer must be embedded to have forward compatible implementations.
type UnimplementedCustomConversionGoalServiceServer struct {
}

func (UnimplementedCustomConversionGoalServiceServer) MutateCustomConversionGoals(context.Context, *MutateCustomConversionGoalsRequest) (*MutateCustomConversionGoalsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method MutateCustomConversionGoals not implemented")
}
func (UnimplementedCustomConversionGoalServiceServer) mustEmbedUnimplementedCustomConversionGoalServiceServer() {
}

// UnsafeCustomConversionGoalServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to CustomConversionGoalServiceServer will
// result in compilation errors.
type UnsafeCustomConversionGoalServiceServer interface {
	mustEmbedUnimplementedCustomConversionGoalServiceServer()
}

func RegisterCustomConversionGoalServiceServer(s grpc.ServiceRegistrar, srv CustomConversionGoalServiceServer) {
	s.RegisterService(&CustomConversionGoalService_ServiceDesc, srv)
}

func _CustomConversionGoalService_MutateCustomConversionGoals_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MutateCustomConversionGoalsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CustomConversionGoalServiceServer).MutateCustomConversionGoals(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/google.ads.googleads.v15.services.CustomConversionGoalService/MutateCustomConversionGoals",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CustomConversionGoalServiceServer).MutateCustomConversionGoals(ctx, req.(*MutateCustomConversionGoalsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// CustomConversionGoalService_ServiceDesc is the grpc.ServiceDesc for CustomConversionGoalService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var CustomConversionGoalService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "google.ads.googleads.v15.services.CustomConversionGoalService",
	HandlerType: (*CustomConversionGoalServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "MutateCustomConversionGoals",
			Handler:    _CustomConversionGoalService_MutateCustomConversionGoals_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "google/ads/googleads/v15/services/custom_conversion_goal_service.proto",
}
