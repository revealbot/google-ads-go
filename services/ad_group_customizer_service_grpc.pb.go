// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.21.12
// source: google/ads/googleads/v15/services/ad_group_customizer_service.proto

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

// AdGroupCustomizerServiceClient is the client API for AdGroupCustomizerService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type AdGroupCustomizerServiceClient interface {
	// Creates, updates or removes ad group customizers. Operation statuses are
	// returned.
	MutateAdGroupCustomizers(ctx context.Context, in *MutateAdGroupCustomizersRequest, opts ...grpc.CallOption) (*MutateAdGroupCustomizersResponse, error)
}

type adGroupCustomizerServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewAdGroupCustomizerServiceClient(cc grpc.ClientConnInterface) AdGroupCustomizerServiceClient {
	return &adGroupCustomizerServiceClient{cc}
}

func (c *adGroupCustomizerServiceClient) MutateAdGroupCustomizers(ctx context.Context, in *MutateAdGroupCustomizersRequest, opts ...grpc.CallOption) (*MutateAdGroupCustomizersResponse, error) {
	out := new(MutateAdGroupCustomizersResponse)
	err := c.cc.Invoke(ctx, "/google.ads.googleads.v15.services.AdGroupCustomizerService/MutateAdGroupCustomizers", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// AdGroupCustomizerServiceServer is the server API for AdGroupCustomizerService service.
// All implementations must embed UnimplementedAdGroupCustomizerServiceServer
// for forward compatibility
type AdGroupCustomizerServiceServer interface {
	// Creates, updates or removes ad group customizers. Operation statuses are
	// returned.
	MutateAdGroupCustomizers(context.Context, *MutateAdGroupCustomizersRequest) (*MutateAdGroupCustomizersResponse, error)
	mustEmbedUnimplementedAdGroupCustomizerServiceServer()
}

// UnimplementedAdGroupCustomizerServiceServer must be embedded to have forward compatible implementations.
type UnimplementedAdGroupCustomizerServiceServer struct {
}

func (UnimplementedAdGroupCustomizerServiceServer) MutateAdGroupCustomizers(context.Context, *MutateAdGroupCustomizersRequest) (*MutateAdGroupCustomizersResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method MutateAdGroupCustomizers not implemented")
}
func (UnimplementedAdGroupCustomizerServiceServer) mustEmbedUnimplementedAdGroupCustomizerServiceServer() {
}

// UnsafeAdGroupCustomizerServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to AdGroupCustomizerServiceServer will
// result in compilation errors.
type UnsafeAdGroupCustomizerServiceServer interface {
	mustEmbedUnimplementedAdGroupCustomizerServiceServer()
}

func RegisterAdGroupCustomizerServiceServer(s grpc.ServiceRegistrar, srv AdGroupCustomizerServiceServer) {
	s.RegisterService(&AdGroupCustomizerService_ServiceDesc, srv)
}

func _AdGroupCustomizerService_MutateAdGroupCustomizers_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MutateAdGroupCustomizersRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AdGroupCustomizerServiceServer).MutateAdGroupCustomizers(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/google.ads.googleads.v15.services.AdGroupCustomizerService/MutateAdGroupCustomizers",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AdGroupCustomizerServiceServer).MutateAdGroupCustomizers(ctx, req.(*MutateAdGroupCustomizersRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// AdGroupCustomizerService_ServiceDesc is the grpc.ServiceDesc for AdGroupCustomizerService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var AdGroupCustomizerService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "google.ads.googleads.v15.services.AdGroupCustomizerService",
	HandlerType: (*AdGroupCustomizerServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "MutateAdGroupCustomizers",
			Handler:    _AdGroupCustomizerService_MutateAdGroupCustomizers_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "google/ads/googleads/v15/services/ad_group_customizer_service.proto",
}
