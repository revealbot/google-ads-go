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

// ExperimentArmServiceClient is the client API for ExperimentArmService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type ExperimentArmServiceClient interface {
	// Creates, updates, or removes experiment arms. Operation statuses are
	// returned.
	//
	// List of thrown errors:
	//   [AuthenticationError]()
	//   [AuthorizationError]()
	//   [ExperimentArmError]()
	//   [HeaderError]()
	//   [InternalError]()
	//   [QuotaError]()
	//   [RequestError]()
	MutateExperimentArms(ctx context.Context, in *MutateExperimentArmsRequest, opts ...grpc.CallOption) (*MutateExperimentArmsResponse, error)
}

type experimentArmServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewExperimentArmServiceClient(cc grpc.ClientConnInterface) ExperimentArmServiceClient {
	return &experimentArmServiceClient{cc}
}

func (c *experimentArmServiceClient) MutateExperimentArms(ctx context.Context, in *MutateExperimentArmsRequest, opts ...grpc.CallOption) (*MutateExperimentArmsResponse, error) {
	out := new(MutateExperimentArmsResponse)
	err := c.cc.Invoke(ctx, "/google.ads.googleads.v13.services.ExperimentArmService/MutateExperimentArms", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ExperimentArmServiceServer is the server API for ExperimentArmService service.
// All implementations must embed UnimplementedExperimentArmServiceServer
// for forward compatibility
type ExperimentArmServiceServer interface {
	// Creates, updates, or removes experiment arms. Operation statuses are
	// returned.
	//
	// List of thrown errors:
	//   [AuthenticationError]()
	//   [AuthorizationError]()
	//   [ExperimentArmError]()
	//   [HeaderError]()
	//   [InternalError]()
	//   [QuotaError]()
	//   [RequestError]()
	MutateExperimentArms(context.Context, *MutateExperimentArmsRequest) (*MutateExperimentArmsResponse, error)
	mustEmbedUnimplementedExperimentArmServiceServer()
}

// UnimplementedExperimentArmServiceServer must be embedded to have forward compatible implementations.
type UnimplementedExperimentArmServiceServer struct {
}

func (UnimplementedExperimentArmServiceServer) MutateExperimentArms(context.Context, *MutateExperimentArmsRequest) (*MutateExperimentArmsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method MutateExperimentArms not implemented")
}
func (UnimplementedExperimentArmServiceServer) mustEmbedUnimplementedExperimentArmServiceServer() {}

// UnsafeExperimentArmServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to ExperimentArmServiceServer will
// result in compilation errors.
type UnsafeExperimentArmServiceServer interface {
	mustEmbedUnimplementedExperimentArmServiceServer()
}

func RegisterExperimentArmServiceServer(s grpc.ServiceRegistrar, srv ExperimentArmServiceServer) {
	s.RegisterService(&ExperimentArmService_ServiceDesc, srv)
}

func _ExperimentArmService_MutateExperimentArms_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MutateExperimentArmsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ExperimentArmServiceServer).MutateExperimentArms(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/google.ads.googleads.v13.services.ExperimentArmService/MutateExperimentArms",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ExperimentArmServiceServer).MutateExperimentArms(ctx, req.(*MutateExperimentArmsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// ExperimentArmService_ServiceDesc is the grpc.ServiceDesc for ExperimentArmService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var ExperimentArmService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "google.ads.googleads.v13.services.ExperimentArmService",
	HandlerType: (*ExperimentArmServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "MutateExperimentArms",
			Handler:    _ExperimentArmService_MutateExperimentArms_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "google/ads/googleads/v13/services/experiment_arm_service.proto",
}
