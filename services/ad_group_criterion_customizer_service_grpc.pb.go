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

// AdGroupCriterionCustomizerServiceClient is the client API for AdGroupCriterionCustomizerService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type AdGroupCriterionCustomizerServiceClient interface {
	// Creates, updates or removes ad group criterion customizers. Operation
	// statuses are returned.
	MutateAdGroupCriterionCustomizers(ctx context.Context, in *MutateAdGroupCriterionCustomizersRequest, opts ...grpc.CallOption) (*MutateAdGroupCriterionCustomizersResponse, error)
}

type adGroupCriterionCustomizerServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewAdGroupCriterionCustomizerServiceClient(cc grpc.ClientConnInterface) AdGroupCriterionCustomizerServiceClient {
	return &adGroupCriterionCustomizerServiceClient{cc}
}

func (c *adGroupCriterionCustomizerServiceClient) MutateAdGroupCriterionCustomizers(ctx context.Context, in *MutateAdGroupCriterionCustomizersRequest, opts ...grpc.CallOption) (*MutateAdGroupCriterionCustomizersResponse, error) {
	out := new(MutateAdGroupCriterionCustomizersResponse)
	err := c.cc.Invoke(ctx, "/google.ads.googleads.v12.services.AdGroupCriterionCustomizerService/MutateAdGroupCriterionCustomizers", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// AdGroupCriterionCustomizerServiceServer is the server API for AdGroupCriterionCustomizerService service.
// All implementations must embed UnimplementedAdGroupCriterionCustomizerServiceServer
// for forward compatibility
type AdGroupCriterionCustomizerServiceServer interface {
	// Creates, updates or removes ad group criterion customizers. Operation
	// statuses are returned.
	MutateAdGroupCriterionCustomizers(context.Context, *MutateAdGroupCriterionCustomizersRequest) (*MutateAdGroupCriterionCustomizersResponse, error)
	mustEmbedUnimplementedAdGroupCriterionCustomizerServiceServer()
}

// UnimplementedAdGroupCriterionCustomizerServiceServer must be embedded to have forward compatible implementations.
type UnimplementedAdGroupCriterionCustomizerServiceServer struct {
}

func (UnimplementedAdGroupCriterionCustomizerServiceServer) MutateAdGroupCriterionCustomizers(context.Context, *MutateAdGroupCriterionCustomizersRequest) (*MutateAdGroupCriterionCustomizersResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method MutateAdGroupCriterionCustomizers not implemented")
}
func (UnimplementedAdGroupCriterionCustomizerServiceServer) mustEmbedUnimplementedAdGroupCriterionCustomizerServiceServer() {
}

// UnsafeAdGroupCriterionCustomizerServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to AdGroupCriterionCustomizerServiceServer will
// result in compilation errors.
type UnsafeAdGroupCriterionCustomizerServiceServer interface {
	mustEmbedUnimplementedAdGroupCriterionCustomizerServiceServer()
}

func RegisterAdGroupCriterionCustomizerServiceServer(s grpc.ServiceRegistrar, srv AdGroupCriterionCustomizerServiceServer) {
	s.RegisterService(&AdGroupCriterionCustomizerService_ServiceDesc, srv)
}

func _AdGroupCriterionCustomizerService_MutateAdGroupCriterionCustomizers_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MutateAdGroupCriterionCustomizersRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AdGroupCriterionCustomizerServiceServer).MutateAdGroupCriterionCustomizers(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/google.ads.googleads.v12.services.AdGroupCriterionCustomizerService/MutateAdGroupCriterionCustomizers",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AdGroupCriterionCustomizerServiceServer).MutateAdGroupCriterionCustomizers(ctx, req.(*MutateAdGroupCriterionCustomizersRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// AdGroupCriterionCustomizerService_ServiceDesc is the grpc.ServiceDesc for AdGroupCriterionCustomizerService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var AdGroupCriterionCustomizerService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "google.ads.googleads.v12.services.AdGroupCriterionCustomizerService",
	HandlerType: (*AdGroupCriterionCustomizerServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "MutateAdGroupCriterionCustomizers",
			Handler:    _AdGroupCriterionCustomizerService_MutateAdGroupCriterionCustomizers_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "google/ads/googleads/v12/services/ad_group_criterion_customizer_service.proto",
}
