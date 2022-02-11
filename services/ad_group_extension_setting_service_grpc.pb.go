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

// AdGroupExtensionSettingServiceClient is the client API for AdGroupExtensionSettingService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type AdGroupExtensionSettingServiceClient interface {
	// Creates, updates, or removes ad group extension settings. Operation
	// statuses are returned.
	//
	// List of thrown errors:
	//   [AuthenticationError]()
	//   [AuthorizationError]()
	//   [CollectionSizeError]()
	//   [CriterionError]()
	//   [DatabaseError]()
	//   [DateError]()
	//   [DistinctError]()
	//   [ExtensionSettingError]()
	//   [FieldError]()
	//   [FieldMaskError]()
	//   [HeaderError]()
	//   [IdError]()
	//   [InternalError]()
	//   [ListOperationError]()
	//   [MutateError]()
	//   [NewResourceCreationError]()
	//   [NotEmptyError]()
	//   [NullError]()
	//   [OperationAccessDeniedError]()
	//   [OperatorError]()
	//   [QuotaError]()
	//   [RangeError]()
	//   [RequestError]()
	//   [ResourceCountLimitExceededError]()
	//   [SizeLimitError]()
	//   [StringFormatError]()
	//   [StringLengthError]()
	//   [UrlFieldError]()
	MutateAdGroupExtensionSettings(ctx context.Context, in *MutateAdGroupExtensionSettingsRequest, opts ...grpc.CallOption) (*MutateAdGroupExtensionSettingsResponse, error)
}

type adGroupExtensionSettingServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewAdGroupExtensionSettingServiceClient(cc grpc.ClientConnInterface) AdGroupExtensionSettingServiceClient {
	return &adGroupExtensionSettingServiceClient{cc}
}

func (c *adGroupExtensionSettingServiceClient) MutateAdGroupExtensionSettings(ctx context.Context, in *MutateAdGroupExtensionSettingsRequest, opts ...grpc.CallOption) (*MutateAdGroupExtensionSettingsResponse, error) {
	out := new(MutateAdGroupExtensionSettingsResponse)
	err := c.cc.Invoke(ctx, "/google.ads.googleads.v10.services.AdGroupExtensionSettingService/MutateAdGroupExtensionSettings", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// AdGroupExtensionSettingServiceServer is the server API for AdGroupExtensionSettingService service.
// All implementations must embed UnimplementedAdGroupExtensionSettingServiceServer
// for forward compatibility
type AdGroupExtensionSettingServiceServer interface {
	// Creates, updates, or removes ad group extension settings. Operation
	// statuses are returned.
	//
	// List of thrown errors:
	//   [AuthenticationError]()
	//   [AuthorizationError]()
	//   [CollectionSizeError]()
	//   [CriterionError]()
	//   [DatabaseError]()
	//   [DateError]()
	//   [DistinctError]()
	//   [ExtensionSettingError]()
	//   [FieldError]()
	//   [FieldMaskError]()
	//   [HeaderError]()
	//   [IdError]()
	//   [InternalError]()
	//   [ListOperationError]()
	//   [MutateError]()
	//   [NewResourceCreationError]()
	//   [NotEmptyError]()
	//   [NullError]()
	//   [OperationAccessDeniedError]()
	//   [OperatorError]()
	//   [QuotaError]()
	//   [RangeError]()
	//   [RequestError]()
	//   [ResourceCountLimitExceededError]()
	//   [SizeLimitError]()
	//   [StringFormatError]()
	//   [StringLengthError]()
	//   [UrlFieldError]()
	MutateAdGroupExtensionSettings(context.Context, *MutateAdGroupExtensionSettingsRequest) (*MutateAdGroupExtensionSettingsResponse, error)
	mustEmbedUnimplementedAdGroupExtensionSettingServiceServer()
}

// UnimplementedAdGroupExtensionSettingServiceServer must be embedded to have forward compatible implementations.
type UnimplementedAdGroupExtensionSettingServiceServer struct {
}

func (UnimplementedAdGroupExtensionSettingServiceServer) MutateAdGroupExtensionSettings(context.Context, *MutateAdGroupExtensionSettingsRequest) (*MutateAdGroupExtensionSettingsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method MutateAdGroupExtensionSettings not implemented")
}
func (UnimplementedAdGroupExtensionSettingServiceServer) mustEmbedUnimplementedAdGroupExtensionSettingServiceServer() {
}

// UnsafeAdGroupExtensionSettingServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to AdGroupExtensionSettingServiceServer will
// result in compilation errors.
type UnsafeAdGroupExtensionSettingServiceServer interface {
	mustEmbedUnimplementedAdGroupExtensionSettingServiceServer()
}

func RegisterAdGroupExtensionSettingServiceServer(s grpc.ServiceRegistrar, srv AdGroupExtensionSettingServiceServer) {
	s.RegisterService(&AdGroupExtensionSettingService_ServiceDesc, srv)
}

func _AdGroupExtensionSettingService_MutateAdGroupExtensionSettings_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MutateAdGroupExtensionSettingsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AdGroupExtensionSettingServiceServer).MutateAdGroupExtensionSettings(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/google.ads.googleads.v10.services.AdGroupExtensionSettingService/MutateAdGroupExtensionSettings",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AdGroupExtensionSettingServiceServer).MutateAdGroupExtensionSettings(ctx, req.(*MutateAdGroupExtensionSettingsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// AdGroupExtensionSettingService_ServiceDesc is the grpc.ServiceDesc for AdGroupExtensionSettingService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var AdGroupExtensionSettingService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "google.ads.googleads.v10.services.AdGroupExtensionSettingService",
	HandlerType: (*AdGroupExtensionSettingServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "MutateAdGroupExtensionSettings",
			Handler:    _AdGroupExtensionSettingService_MutateAdGroupExtensionSettings_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "google/ads/googleads/v10/services/ad_group_extension_setting_service.proto",
}
