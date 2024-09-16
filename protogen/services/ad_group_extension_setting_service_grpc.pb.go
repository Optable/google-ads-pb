// Copyright 2024 Google LLC
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.5.1
// - protoc             v3.21.1
// source: google/ads/googleads/v17/services/ad_group_extension_setting_service.proto

package services

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.64.0 or later.
const _ = grpc.SupportPackageIsVersion9

const (
	AdGroupExtensionSettingService_MutateAdGroupExtensionSettings_FullMethodName = "/google.ads.googleads.v17.services.AdGroupExtensionSettingService/MutateAdGroupExtensionSettings"
)

// AdGroupExtensionSettingServiceClient is the client API for AdGroupExtensionSettingService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
//
// Service to manage ad group extension settings.
type AdGroupExtensionSettingServiceClient interface {
	// Creates, updates, or removes ad group extension settings. Operation
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
	//	[FieldMaskError]()
	//	[HeaderError]()
	//	[IdError]()
	//	[InternalError]()
	//	[ListOperationError]()
	//	[MutateError]()
	//	[NewResourceCreationError]()
	//	[NotEmptyError]()
	//	[NullError]()
	//	[OperationAccessDeniedError]()
	//	[OperatorError]()
	//	[QuotaError]()
	//	[RangeError]()
	//	[RequestError]()
	//	[ResourceCountLimitExceededError]()
	//	[SizeLimitError]()
	//	[StringFormatError]()
	//	[StringLengthError]()
	//	[UrlFieldError]()
	MutateAdGroupExtensionSettings(ctx context.Context, in *MutateAdGroupExtensionSettingsRequest, opts ...grpc.CallOption) (*MutateAdGroupExtensionSettingsResponse, error)
}

type adGroupExtensionSettingServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewAdGroupExtensionSettingServiceClient(cc grpc.ClientConnInterface) AdGroupExtensionSettingServiceClient {
	return &adGroupExtensionSettingServiceClient{cc}
}

func (c *adGroupExtensionSettingServiceClient) MutateAdGroupExtensionSettings(ctx context.Context, in *MutateAdGroupExtensionSettingsRequest, opts ...grpc.CallOption) (*MutateAdGroupExtensionSettingsResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(MutateAdGroupExtensionSettingsResponse)
	err := c.cc.Invoke(ctx, AdGroupExtensionSettingService_MutateAdGroupExtensionSettings_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// AdGroupExtensionSettingServiceServer is the server API for AdGroupExtensionSettingService service.
// All implementations must embed UnimplementedAdGroupExtensionSettingServiceServer
// for forward compatibility.
//
// Service to manage ad group extension settings.
type AdGroupExtensionSettingServiceServer interface {
	// Creates, updates, or removes ad group extension settings. Operation
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
	//	[FieldMaskError]()
	//	[HeaderError]()
	//	[IdError]()
	//	[InternalError]()
	//	[ListOperationError]()
	//	[MutateError]()
	//	[NewResourceCreationError]()
	//	[NotEmptyError]()
	//	[NullError]()
	//	[OperationAccessDeniedError]()
	//	[OperatorError]()
	//	[QuotaError]()
	//	[RangeError]()
	//	[RequestError]()
	//	[ResourceCountLimitExceededError]()
	//	[SizeLimitError]()
	//	[StringFormatError]()
	//	[StringLengthError]()
	//	[UrlFieldError]()
	MutateAdGroupExtensionSettings(context.Context, *MutateAdGroupExtensionSettingsRequest) (*MutateAdGroupExtensionSettingsResponse, error)
	mustEmbedUnimplementedAdGroupExtensionSettingServiceServer()
}

// UnimplementedAdGroupExtensionSettingServiceServer must be embedded to have
// forward compatible implementations.
//
// NOTE: this should be embedded by value instead of pointer to avoid a nil
// pointer dereference when methods are called.
type UnimplementedAdGroupExtensionSettingServiceServer struct{}

func (UnimplementedAdGroupExtensionSettingServiceServer) MutateAdGroupExtensionSettings(context.Context, *MutateAdGroupExtensionSettingsRequest) (*MutateAdGroupExtensionSettingsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method MutateAdGroupExtensionSettings not implemented")
}
func (UnimplementedAdGroupExtensionSettingServiceServer) mustEmbedUnimplementedAdGroupExtensionSettingServiceServer() {
}
func (UnimplementedAdGroupExtensionSettingServiceServer) testEmbeddedByValue() {}

// UnsafeAdGroupExtensionSettingServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to AdGroupExtensionSettingServiceServer will
// result in compilation errors.
type UnsafeAdGroupExtensionSettingServiceServer interface {
	mustEmbedUnimplementedAdGroupExtensionSettingServiceServer()
}

func RegisterAdGroupExtensionSettingServiceServer(s grpc.ServiceRegistrar, srv AdGroupExtensionSettingServiceServer) {
	// If the following call pancis, it indicates UnimplementedAdGroupExtensionSettingServiceServer was
	// embedded by pointer and is nil.  This will cause panics if an
	// unimplemented method is ever invoked, so we test this at initialization
	// time to prevent it from happening at runtime later due to I/O.
	if t, ok := srv.(interface{ testEmbeddedByValue() }); ok {
		t.testEmbeddedByValue()
	}
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
		FullMethod: AdGroupExtensionSettingService_MutateAdGroupExtensionSettings_FullMethodName,
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
	ServiceName: "google.ads.googleads.v17.services.AdGroupExtensionSettingService",
	HandlerType: (*AdGroupExtensionSettingServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "MutateAdGroupExtensionSettings",
			Handler:    _AdGroupExtensionSettingService_MutateAdGroupExtensionSettings_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "google/ads/googleads/v17/services/ad_group_extension_setting_service.proto",
}
