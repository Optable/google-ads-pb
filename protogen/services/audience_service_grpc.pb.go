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
// source: google/ads/googleads/v14/services/audience_service.proto

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
	AudienceService_MutateAudiences_FullMethodName = "/google.ads.googleads.v14.services.AudienceService/MutateAudiences"
)

// AudienceServiceClient is the client API for AudienceService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type AudienceServiceClient interface {
	// Creates audiences. Operation statuses are returned.
	//
	// List of thrown errors:
	//
	//	[AudienceError]()
	MutateAudiences(ctx context.Context, in *MutateAudiencesRequest, opts ...grpc.CallOption) (*MutateAudiencesResponse, error)
}

type audienceServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewAudienceServiceClient(cc grpc.ClientConnInterface) AudienceServiceClient {
	return &audienceServiceClient{cc}
}

func (c *audienceServiceClient) MutateAudiences(ctx context.Context, in *MutateAudiencesRequest, opts ...grpc.CallOption) (*MutateAudiencesResponse, error) {
	out := new(MutateAudiencesResponse)
	err := c.cc.Invoke(ctx, AudienceService_MutateAudiences_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// AudienceServiceServer is the server API for AudienceService service.
// All implementations must embed UnimplementedAudienceServiceServer
// for forward compatibility
type AudienceServiceServer interface {
	// Creates audiences. Operation statuses are returned.
	//
	// List of thrown errors:
	//
	//	[AudienceError]()
	MutateAudiences(context.Context, *MutateAudiencesRequest) (*MutateAudiencesResponse, error)
	mustEmbedUnimplementedAudienceServiceServer()
}

// UnimplementedAudienceServiceServer must be embedded to have forward compatible implementations.
type UnimplementedAudienceServiceServer struct {
}

func (UnimplementedAudienceServiceServer) MutateAudiences(context.Context, *MutateAudiencesRequest) (*MutateAudiencesResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method MutateAudiences not implemented")
}
func (UnimplementedAudienceServiceServer) mustEmbedUnimplementedAudienceServiceServer() {}

// UnsafeAudienceServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to AudienceServiceServer will
// result in compilation errors.
type UnsafeAudienceServiceServer interface {
	mustEmbedUnimplementedAudienceServiceServer()
}

func RegisterAudienceServiceServer(s grpc.ServiceRegistrar, srv AudienceServiceServer) {
	s.RegisterService(&AudienceService_ServiceDesc, srv)
}

func _AudienceService_MutateAudiences_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MutateAudiencesRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AudienceServiceServer).MutateAudiences(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: AudienceService_MutateAudiences_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AudienceServiceServer).MutateAudiences(ctx, req.(*MutateAudiencesRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// AudienceService_ServiceDesc is the grpc.ServiceDesc for AudienceService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var AudienceService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "google.ads.googleads.v14.services.AudienceService",
	HandlerType: (*AudienceServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "MutateAudiences",
			Handler:    _AudienceService_MutateAudiences_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "google/ads/googleads/v14/services/audience_service.proto",
}
