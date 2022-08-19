// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.21.5
// source: google/ads/googleads/v11/services/conversion_upload_service.proto

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

// ConversionUploadServiceClient is the client API for ConversionUploadService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type ConversionUploadServiceClient interface {
	// Processes the given click conversions.
	//
	// List of thrown errors:
	//   [AuthenticationError]()
	//   [AuthorizationError]()
	//   [ConversionUploadError]()
	//   [HeaderError]()
	//   [InternalError]()
	//   [PartialFailureError]()
	//   [QuotaError]()
	//   [RequestError]()
	UploadClickConversions(ctx context.Context, in *UploadClickConversionsRequest, opts ...grpc.CallOption) (*UploadClickConversionsResponse, error)
	// Processes the given call conversions.
	//
	// List of thrown errors:
	//   [AuthenticationError]()
	//   [AuthorizationError]()
	//   [HeaderError]()
	//   [InternalError]()
	//   [PartialFailureError]()
	//   [QuotaError]()
	//   [RequestError]()
	UploadCallConversions(ctx context.Context, in *UploadCallConversionsRequest, opts ...grpc.CallOption) (*UploadCallConversionsResponse, error)
}

type conversionUploadServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewConversionUploadServiceClient(cc grpc.ClientConnInterface) ConversionUploadServiceClient {
	return &conversionUploadServiceClient{cc}
}

func (c *conversionUploadServiceClient) UploadClickConversions(ctx context.Context, in *UploadClickConversionsRequest, opts ...grpc.CallOption) (*UploadClickConversionsResponse, error) {
	out := new(UploadClickConversionsResponse)
	err := c.cc.Invoke(ctx, "/google.ads.googleads.v11.services.ConversionUploadService/UploadClickConversions", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *conversionUploadServiceClient) UploadCallConversions(ctx context.Context, in *UploadCallConversionsRequest, opts ...grpc.CallOption) (*UploadCallConversionsResponse, error) {
	out := new(UploadCallConversionsResponse)
	err := c.cc.Invoke(ctx, "/google.ads.googleads.v11.services.ConversionUploadService/UploadCallConversions", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ConversionUploadServiceServer is the server API for ConversionUploadService service.
// All implementations must embed UnimplementedConversionUploadServiceServer
// for forward compatibility
type ConversionUploadServiceServer interface {
	// Processes the given click conversions.
	//
	// List of thrown errors:
	//   [AuthenticationError]()
	//   [AuthorizationError]()
	//   [ConversionUploadError]()
	//   [HeaderError]()
	//   [InternalError]()
	//   [PartialFailureError]()
	//   [QuotaError]()
	//   [RequestError]()
	UploadClickConversions(context.Context, *UploadClickConversionsRequest) (*UploadClickConversionsResponse, error)
	// Processes the given call conversions.
	//
	// List of thrown errors:
	//   [AuthenticationError]()
	//   [AuthorizationError]()
	//   [HeaderError]()
	//   [InternalError]()
	//   [PartialFailureError]()
	//   [QuotaError]()
	//   [RequestError]()
	UploadCallConversions(context.Context, *UploadCallConversionsRequest) (*UploadCallConversionsResponse, error)
	mustEmbedUnimplementedConversionUploadServiceServer()
}

// UnimplementedConversionUploadServiceServer must be embedded to have forward compatible implementations.
type UnimplementedConversionUploadServiceServer struct {
}

func (UnimplementedConversionUploadServiceServer) UploadClickConversions(context.Context, *UploadClickConversionsRequest) (*UploadClickConversionsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UploadClickConversions not implemented")
}
func (UnimplementedConversionUploadServiceServer) UploadCallConversions(context.Context, *UploadCallConversionsRequest) (*UploadCallConversionsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UploadCallConversions not implemented")
}
func (UnimplementedConversionUploadServiceServer) mustEmbedUnimplementedConversionUploadServiceServer() {
}

// UnsafeConversionUploadServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to ConversionUploadServiceServer will
// result in compilation errors.
type UnsafeConversionUploadServiceServer interface {
	mustEmbedUnimplementedConversionUploadServiceServer()
}

func RegisterConversionUploadServiceServer(s grpc.ServiceRegistrar, srv ConversionUploadServiceServer) {
	s.RegisterService(&ConversionUploadService_ServiceDesc, srv)
}

func _ConversionUploadService_UploadClickConversions_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UploadClickConversionsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ConversionUploadServiceServer).UploadClickConversions(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/google.ads.googleads.v11.services.ConversionUploadService/UploadClickConversions",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ConversionUploadServiceServer).UploadClickConversions(ctx, req.(*UploadClickConversionsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ConversionUploadService_UploadCallConversions_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UploadCallConversionsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ConversionUploadServiceServer).UploadCallConversions(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/google.ads.googleads.v11.services.ConversionUploadService/UploadCallConversions",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ConversionUploadServiceServer).UploadCallConversions(ctx, req.(*UploadCallConversionsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// ConversionUploadService_ServiceDesc is the grpc.ServiceDesc for ConversionUploadService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var ConversionUploadService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "google.ads.googleads.v11.services.ConversionUploadService",
	HandlerType: (*ConversionUploadServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "UploadClickConversions",
			Handler:    _ConversionUploadService_UploadClickConversions_Handler,
		},
		{
			MethodName: "UploadCallConversions",
			Handler:    _ConversionUploadService_UploadCallConversions_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "google/ads/googleads/v11/services/conversion_upload_service.proto",
}
