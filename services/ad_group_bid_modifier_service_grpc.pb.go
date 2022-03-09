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

// AdGroupBidModifierServiceClient is the client API for AdGroupBidModifierService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type AdGroupBidModifierServiceClient interface {
	// Creates, updates, or removes ad group bid modifiers.
	// Operation statuses are returned.
	//
	// List of thrown errors:
	//   [AdGroupBidModifierError]()
	//   [AuthenticationError]()
	//   [AuthorizationError]()
	//   [ContextError]()
	//   [CriterionError]()
	//   [DatabaseError]()
	//   [DistinctError]()
	//   [FieldError]()
	//   [FieldMaskError]()
	//   [HeaderError]()
	//   [IdError]()
	//   [InternalError]()
	//   [MutateError]()
	//   [NewResourceCreationError]()
	//   [NotEmptyError]()
	//   [OperatorError]()
	//   [QuotaError]()
	//   [RangeError]()
	//   [RequestError]()
	//   [ResourceCountLimitExceededError]()
	//   [SizeLimitError]()
	//   [StringFormatError]()
	//   [StringLengthError]()
	MutateAdGroupBidModifiers(ctx context.Context, in *MutateAdGroupBidModifiersRequest, opts ...grpc.CallOption) (*MutateAdGroupBidModifiersResponse, error)
}

type adGroupBidModifierServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewAdGroupBidModifierServiceClient(cc grpc.ClientConnInterface) AdGroupBidModifierServiceClient {
	return &adGroupBidModifierServiceClient{cc}
}

func (c *adGroupBidModifierServiceClient) MutateAdGroupBidModifiers(ctx context.Context, in *MutateAdGroupBidModifiersRequest, opts ...grpc.CallOption) (*MutateAdGroupBidModifiersResponse, error) {
	out := new(MutateAdGroupBidModifiersResponse)
	err := c.cc.Invoke(ctx, "/google.ads.googleads.v10.services.AdGroupBidModifierService/MutateAdGroupBidModifiers", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// AdGroupBidModifierServiceServer is the server API for AdGroupBidModifierService service.
// All implementations must embed UnimplementedAdGroupBidModifierServiceServer
// for forward compatibility
type AdGroupBidModifierServiceServer interface {
	// Creates, updates, or removes ad group bid modifiers.
	// Operation statuses are returned.
	//
	// List of thrown errors:
	//   [AdGroupBidModifierError]()
	//   [AuthenticationError]()
	//   [AuthorizationError]()
	//   [ContextError]()
	//   [CriterionError]()
	//   [DatabaseError]()
	//   [DistinctError]()
	//   [FieldError]()
	//   [FieldMaskError]()
	//   [HeaderError]()
	//   [IdError]()
	//   [InternalError]()
	//   [MutateError]()
	//   [NewResourceCreationError]()
	//   [NotEmptyError]()
	//   [OperatorError]()
	//   [QuotaError]()
	//   [RangeError]()
	//   [RequestError]()
	//   [ResourceCountLimitExceededError]()
	//   [SizeLimitError]()
	//   [StringFormatError]()
	//   [StringLengthError]()
	MutateAdGroupBidModifiers(context.Context, *MutateAdGroupBidModifiersRequest) (*MutateAdGroupBidModifiersResponse, error)
	mustEmbedUnimplementedAdGroupBidModifierServiceServer()
}

// UnimplementedAdGroupBidModifierServiceServer must be embedded to have forward compatible implementations.
type UnimplementedAdGroupBidModifierServiceServer struct {
}

func (UnimplementedAdGroupBidModifierServiceServer) MutateAdGroupBidModifiers(context.Context, *MutateAdGroupBidModifiersRequest) (*MutateAdGroupBidModifiersResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method MutateAdGroupBidModifiers not implemented")
}
func (UnimplementedAdGroupBidModifierServiceServer) mustEmbedUnimplementedAdGroupBidModifierServiceServer() {
}

// UnsafeAdGroupBidModifierServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to AdGroupBidModifierServiceServer will
// result in compilation errors.
type UnsafeAdGroupBidModifierServiceServer interface {
	mustEmbedUnimplementedAdGroupBidModifierServiceServer()
}

func RegisterAdGroupBidModifierServiceServer(s grpc.ServiceRegistrar, srv AdGroupBidModifierServiceServer) {
	s.RegisterService(&AdGroupBidModifierService_ServiceDesc, srv)
}

func _AdGroupBidModifierService_MutateAdGroupBidModifiers_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MutateAdGroupBidModifiersRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AdGroupBidModifierServiceServer).MutateAdGroupBidModifiers(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/google.ads.googleads.v10.services.AdGroupBidModifierService/MutateAdGroupBidModifiers",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AdGroupBidModifierServiceServer).MutateAdGroupBidModifiers(ctx, req.(*MutateAdGroupBidModifiersRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// AdGroupBidModifierService_ServiceDesc is the grpc.ServiceDesc for AdGroupBidModifierService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var AdGroupBidModifierService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "google.ads.googleads.v10.services.AdGroupBidModifierService",
	HandlerType: (*AdGroupBidModifierServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "MutateAdGroupBidModifiers",
			Handler:    _AdGroupBidModifierService_MutateAdGroupBidModifiers_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "google/ads/googleads/v10/services/ad_group_bid_modifier_service.proto",
}
