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
// source: google/ads/googleads/v14/services/feed_mapping_service.proto

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
	FeedMappingService_MutateFeedMappings_FullMethodName = "/google.ads.googleads.v14.services.FeedMappingService/MutateFeedMappings"
)

// FeedMappingServiceClient is the client API for FeedMappingService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type FeedMappingServiceClient interface {
	// Creates or removes feed mappings. Operation statuses are
	// returned.
	//
	// List of thrown errors:
	//
	//	[AuthenticationError]()
	//	[AuthorizationError]()
	//	[DatabaseError]()
	//	[DistinctError]()
	//	[FeedMappingError]()
	//	[FieldError]()
	//	[HeaderError]()
	//	[IdError]()
	//	[InternalError]()
	//	[MutateError]()
	//	[NotEmptyError]()
	//	[OperationAccessDeniedError]()
	//	[OperatorError]()
	//	[QuotaError]()
	//	[RangeError]()
	//	[RequestError]()
	//	[SizeLimitError]()
	//	[StringFormatError]()
	//	[StringLengthError]()
	MutateFeedMappings(ctx context.Context, in *MutateFeedMappingsRequest, opts ...grpc.CallOption) (*MutateFeedMappingsResponse, error)
}

type feedMappingServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewFeedMappingServiceClient(cc grpc.ClientConnInterface) FeedMappingServiceClient {
	return &feedMappingServiceClient{cc}
}

func (c *feedMappingServiceClient) MutateFeedMappings(ctx context.Context, in *MutateFeedMappingsRequest, opts ...grpc.CallOption) (*MutateFeedMappingsResponse, error) {
	out := new(MutateFeedMappingsResponse)
	err := c.cc.Invoke(ctx, FeedMappingService_MutateFeedMappings_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// FeedMappingServiceServer is the server API for FeedMappingService service.
// All implementations must embed UnimplementedFeedMappingServiceServer
// for forward compatibility
type FeedMappingServiceServer interface {
	// Creates or removes feed mappings. Operation statuses are
	// returned.
	//
	// List of thrown errors:
	//
	//	[AuthenticationError]()
	//	[AuthorizationError]()
	//	[DatabaseError]()
	//	[DistinctError]()
	//	[FeedMappingError]()
	//	[FieldError]()
	//	[HeaderError]()
	//	[IdError]()
	//	[InternalError]()
	//	[MutateError]()
	//	[NotEmptyError]()
	//	[OperationAccessDeniedError]()
	//	[OperatorError]()
	//	[QuotaError]()
	//	[RangeError]()
	//	[RequestError]()
	//	[SizeLimitError]()
	//	[StringFormatError]()
	//	[StringLengthError]()
	MutateFeedMappings(context.Context, *MutateFeedMappingsRequest) (*MutateFeedMappingsResponse, error)
	mustEmbedUnimplementedFeedMappingServiceServer()
}

// UnimplementedFeedMappingServiceServer must be embedded to have forward compatible implementations.
type UnimplementedFeedMappingServiceServer struct {
}

func (UnimplementedFeedMappingServiceServer) MutateFeedMappings(context.Context, *MutateFeedMappingsRequest) (*MutateFeedMappingsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method MutateFeedMappings not implemented")
}
func (UnimplementedFeedMappingServiceServer) mustEmbedUnimplementedFeedMappingServiceServer() {}

// UnsafeFeedMappingServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to FeedMappingServiceServer will
// result in compilation errors.
type UnsafeFeedMappingServiceServer interface {
	mustEmbedUnimplementedFeedMappingServiceServer()
}

func RegisterFeedMappingServiceServer(s grpc.ServiceRegistrar, srv FeedMappingServiceServer) {
	s.RegisterService(&FeedMappingService_ServiceDesc, srv)
}

func _FeedMappingService_MutateFeedMappings_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MutateFeedMappingsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(FeedMappingServiceServer).MutateFeedMappings(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: FeedMappingService_MutateFeedMappings_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(FeedMappingServiceServer).MutateFeedMappings(ctx, req.(*MutateFeedMappingsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// FeedMappingService_ServiceDesc is the grpc.ServiceDesc for FeedMappingService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var FeedMappingService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "google.ads.googleads.v14.services.FeedMappingService",
	HandlerType: (*FeedMappingServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "MutateFeedMappings",
			Handler:    _FeedMappingService_MutateFeedMappings_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "google/ads/googleads/v14/services/feed_mapping_service.proto",
}
