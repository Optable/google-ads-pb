// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.21.9
// source: google/ads/googleads/v12/services/customer_extension_setting_service.proto

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

// CustomerExtensionSettingServiceClient is the client API for CustomerExtensionSettingService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type CustomerExtensionSettingServiceClient interface {
	// Creates, updates, or removes customer extension settings. Operation
	// statuses are returned.
	//
	// List of thrown errors:
	//
	//	[AuthenticationError]()
	//	[AuthorizationError]()
	//	[CollectionSizeError]()
	//	[CriterionError]()
	//	[DatabaseError]()
	//	[DateError]()
	//	[DistinctError]()
	//	[ExtensionSettingError]()
	//	[FieldError]()
	//	[HeaderError]()
	//	[IdError]()
	//	[InternalError]()
	//	[ListOperationError]()
	//	[MutateError]()
	//	[NewResourceCreationError]()
	//	[NotEmptyError]()
	//	[NullError]()
	//	[OperatorError]()
	//	[QuotaError]()
	//	[RangeError]()
	//	[RequestError]()
	//	[SizeLimitError]()
	//	[StringFormatError]()
	//	[StringLengthError]()
	//	[UrlFieldError]()
	MutateCustomerExtensionSettings(ctx context.Context, in *MutateCustomerExtensionSettingsRequest, opts ...grpc.CallOption) (*MutateCustomerExtensionSettingsResponse, error)
}

type customerExtensionSettingServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewCustomerExtensionSettingServiceClient(cc grpc.ClientConnInterface) CustomerExtensionSettingServiceClient {
	return &customerExtensionSettingServiceClient{cc}
}

func (c *customerExtensionSettingServiceClient) MutateCustomerExtensionSettings(ctx context.Context, in *MutateCustomerExtensionSettingsRequest, opts ...grpc.CallOption) (*MutateCustomerExtensionSettingsResponse, error) {
	out := new(MutateCustomerExtensionSettingsResponse)
	err := c.cc.Invoke(ctx, "/google.ads.googleads.v12.services.CustomerExtensionSettingService/MutateCustomerExtensionSettings", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// CustomerExtensionSettingServiceServer is the server API for CustomerExtensionSettingService service.
// All implementations must embed UnimplementedCustomerExtensionSettingServiceServer
// for forward compatibility
type CustomerExtensionSettingServiceServer interface {
	// Creates, updates, or removes customer extension settings. Operation
	// statuses are returned.
	//
	// List of thrown errors:
	//
	//	[AuthenticationError]()
	//	[AuthorizationError]()
	//	[CollectionSizeError]()
	//	[CriterionError]()
	//	[DatabaseError]()
	//	[DateError]()
	//	[DistinctError]()
	//	[ExtensionSettingError]()
	//	[FieldError]()
	//	[HeaderError]()
	//	[IdError]()
	//	[InternalError]()
	//	[ListOperationError]()
	//	[MutateError]()
	//	[NewResourceCreationError]()
	//	[NotEmptyError]()
	//	[NullError]()
	//	[OperatorError]()
	//	[QuotaError]()
	//	[RangeError]()
	//	[RequestError]()
	//	[SizeLimitError]()
	//	[StringFormatError]()
	//	[StringLengthError]()
	//	[UrlFieldError]()
	MutateCustomerExtensionSettings(context.Context, *MutateCustomerExtensionSettingsRequest) (*MutateCustomerExtensionSettingsResponse, error)
	mustEmbedUnimplementedCustomerExtensionSettingServiceServer()
}

// UnimplementedCustomerExtensionSettingServiceServer must be embedded to have forward compatible implementations.
type UnimplementedCustomerExtensionSettingServiceServer struct {
}

func (UnimplementedCustomerExtensionSettingServiceServer) MutateCustomerExtensionSettings(context.Context, *MutateCustomerExtensionSettingsRequest) (*MutateCustomerExtensionSettingsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method MutateCustomerExtensionSettings not implemented")
}
func (UnimplementedCustomerExtensionSettingServiceServer) mustEmbedUnimplementedCustomerExtensionSettingServiceServer() {
}

// UnsafeCustomerExtensionSettingServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to CustomerExtensionSettingServiceServer will
// result in compilation errors.
type UnsafeCustomerExtensionSettingServiceServer interface {
	mustEmbedUnimplementedCustomerExtensionSettingServiceServer()
}

func RegisterCustomerExtensionSettingServiceServer(s grpc.ServiceRegistrar, srv CustomerExtensionSettingServiceServer) {
	s.RegisterService(&CustomerExtensionSettingService_ServiceDesc, srv)
}

func _CustomerExtensionSettingService_MutateCustomerExtensionSettings_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MutateCustomerExtensionSettingsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CustomerExtensionSettingServiceServer).MutateCustomerExtensionSettings(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/google.ads.googleads.v12.services.CustomerExtensionSettingService/MutateCustomerExtensionSettings",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CustomerExtensionSettingServiceServer).MutateCustomerExtensionSettings(ctx, req.(*MutateCustomerExtensionSettingsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// CustomerExtensionSettingService_ServiceDesc is the grpc.ServiceDesc for CustomerExtensionSettingService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var CustomerExtensionSettingService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "google.ads.googleads.v12.services.CustomerExtensionSettingService",
	HandlerType: (*CustomerExtensionSettingServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "MutateCustomerExtensionSettings",
			Handler:    _CustomerExtensionSettingService_MutateCustomerExtensionSettings_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "google/ads/googleads/v12/services/customer_extension_setting_service.proto",
}
