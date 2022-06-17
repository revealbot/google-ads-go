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

// MediaFileServiceClient is the client API for MediaFileService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type MediaFileServiceClient interface {
	// Creates media files. Operation statuses are returned.
	//
	// List of thrown errors:
	//   [AuthenticationError]()
	//   [AuthorizationError]()
	//   [DatabaseError]()
	//   [DistinctError]()
	//   [FieldError]()
	//   [HeaderError]()
	//   [IdError]()
	//   [ImageError]()
	//   [InternalError]()
	//   [MediaBundleError]()
	//   [MediaFileError]()
	//   [NewResourceCreationError]()
	//   [NotEmptyError]()
	//   [NullError]()
	//   [OperatorError]()
	//   [QuotaError]()
	//   [RangeError]()
	//   [RequestError]()
	//   [SizeLimitError]()
	//   [StringFormatError]()
	//   [StringLengthError]()
	MutateMediaFiles(ctx context.Context, in *MutateMediaFilesRequest, opts ...grpc.CallOption) (*MutateMediaFilesResponse, error)
}

type mediaFileServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewMediaFileServiceClient(cc grpc.ClientConnInterface) MediaFileServiceClient {
	return &mediaFileServiceClient{cc}
}

func (c *mediaFileServiceClient) MutateMediaFiles(ctx context.Context, in *MutateMediaFilesRequest, opts ...grpc.CallOption) (*MutateMediaFilesResponse, error) {
	out := new(MutateMediaFilesResponse)
	err := c.cc.Invoke(ctx, "/google.ads.googleads.v11.services.MediaFileService/MutateMediaFiles", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// MediaFileServiceServer is the server API for MediaFileService service.
// All implementations must embed UnimplementedMediaFileServiceServer
// for forward compatibility
type MediaFileServiceServer interface {
	// Creates media files. Operation statuses are returned.
	//
	// List of thrown errors:
	//   [AuthenticationError]()
	//   [AuthorizationError]()
	//   [DatabaseError]()
	//   [DistinctError]()
	//   [FieldError]()
	//   [HeaderError]()
	//   [IdError]()
	//   [ImageError]()
	//   [InternalError]()
	//   [MediaBundleError]()
	//   [MediaFileError]()
	//   [NewResourceCreationError]()
	//   [NotEmptyError]()
	//   [NullError]()
	//   [OperatorError]()
	//   [QuotaError]()
	//   [RangeError]()
	//   [RequestError]()
	//   [SizeLimitError]()
	//   [StringFormatError]()
	//   [StringLengthError]()
	MutateMediaFiles(context.Context, *MutateMediaFilesRequest) (*MutateMediaFilesResponse, error)
	mustEmbedUnimplementedMediaFileServiceServer()
}

// UnimplementedMediaFileServiceServer must be embedded to have forward compatible implementations.
type UnimplementedMediaFileServiceServer struct {
}

func (UnimplementedMediaFileServiceServer) MutateMediaFiles(context.Context, *MutateMediaFilesRequest) (*MutateMediaFilesResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method MutateMediaFiles not implemented")
}
func (UnimplementedMediaFileServiceServer) mustEmbedUnimplementedMediaFileServiceServer() {}

// UnsafeMediaFileServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to MediaFileServiceServer will
// result in compilation errors.
type UnsafeMediaFileServiceServer interface {
	mustEmbedUnimplementedMediaFileServiceServer()
}

func RegisterMediaFileServiceServer(s grpc.ServiceRegistrar, srv MediaFileServiceServer) {
	s.RegisterService(&MediaFileService_ServiceDesc, srv)
}

func _MediaFileService_MutateMediaFiles_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MutateMediaFilesRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MediaFileServiceServer).MutateMediaFiles(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/google.ads.googleads.v11.services.MediaFileService/MutateMediaFiles",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MediaFileServiceServer).MutateMediaFiles(ctx, req.(*MutateMediaFilesRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// MediaFileService_ServiceDesc is the grpc.ServiceDesc for MediaFileService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var MediaFileService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "google.ads.googleads.v11.services.MediaFileService",
	HandlerType: (*MediaFileServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "MutateMediaFiles",
			Handler:    _MediaFileService_MutateMediaFiles_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "google/ads/googleads/v11/services/media_file_service.proto",
}
