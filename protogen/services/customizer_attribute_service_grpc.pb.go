// Copyright 2023 Google LLC
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
// - protoc-gen-go-grpc v1.3.0
// - protoc             v3.21.1
// source: google/ads/googleads/v14/services/customizer_attribute_service.proto

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

const (
	CustomizerAttributeService_MutateCustomizerAttributes_FullMethodName = "/google.ads.googleads.v14.services.CustomizerAttributeService/MutateCustomizerAttributes"
)

// CustomizerAttributeServiceClient is the client API for CustomizerAttributeService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type CustomizerAttributeServiceClient interface {
	// Creates, updates or removes customizer attributes. Operation statuses are
	// returned.
	MutateCustomizerAttributes(ctx context.Context, in *MutateCustomizerAttributesRequest, opts ...grpc.CallOption) (*MutateCustomizerAttributesResponse, error)
}

type customizerAttributeServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewCustomizerAttributeServiceClient(cc grpc.ClientConnInterface) CustomizerAttributeServiceClient {
	return &customizerAttributeServiceClient{cc}
}

func (c *customizerAttributeServiceClient) MutateCustomizerAttributes(ctx context.Context, in *MutateCustomizerAttributesRequest, opts ...grpc.CallOption) (*MutateCustomizerAttributesResponse, error) {
	out := new(MutateCustomizerAttributesResponse)
	err := c.cc.Invoke(ctx, CustomizerAttributeService_MutateCustomizerAttributes_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// CustomizerAttributeServiceServer is the server API for CustomizerAttributeService service.
// All implementations must embed UnimplementedCustomizerAttributeServiceServer
// for forward compatibility
type CustomizerAttributeServiceServer interface {
	// Creates, updates or removes customizer attributes. Operation statuses are
	// returned.
	MutateCustomizerAttributes(context.Context, *MutateCustomizerAttributesRequest) (*MutateCustomizerAttributesResponse, error)
	mustEmbedUnimplementedCustomizerAttributeServiceServer()
}

// UnimplementedCustomizerAttributeServiceServer must be embedded to have forward compatible implementations.
type UnimplementedCustomizerAttributeServiceServer struct {
}

func (UnimplementedCustomizerAttributeServiceServer) MutateCustomizerAttributes(context.Context, *MutateCustomizerAttributesRequest) (*MutateCustomizerAttributesResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method MutateCustomizerAttributes not implemented")
}
func (UnimplementedCustomizerAttributeServiceServer) mustEmbedUnimplementedCustomizerAttributeServiceServer() {
}

// UnsafeCustomizerAttributeServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to CustomizerAttributeServiceServer will
// result in compilation errors.
type UnsafeCustomizerAttributeServiceServer interface {
	mustEmbedUnimplementedCustomizerAttributeServiceServer()
}

func RegisterCustomizerAttributeServiceServer(s grpc.ServiceRegistrar, srv CustomizerAttributeServiceServer) {
	s.RegisterService(&CustomizerAttributeService_ServiceDesc, srv)
}

func _CustomizerAttributeService_MutateCustomizerAttributes_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MutateCustomizerAttributesRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CustomizerAttributeServiceServer).MutateCustomizerAttributes(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: CustomizerAttributeService_MutateCustomizerAttributes_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CustomizerAttributeServiceServer).MutateCustomizerAttributes(ctx, req.(*MutateCustomizerAttributesRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// CustomizerAttributeService_ServiceDesc is the grpc.ServiceDesc for CustomizerAttributeService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var CustomizerAttributeService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "google.ads.googleads.v14.services.CustomizerAttributeService",
	HandlerType: (*CustomizerAttributeServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "MutateCustomizerAttributes",
			Handler:    _CustomizerAttributeService_MutateCustomizerAttributes_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "google/ads/googleads/v14/services/customizer_attribute_service.proto",
}
