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
// source: google/ads/googleads/v14/services/shared_criterion_service.proto

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
	SharedCriterionService_MutateSharedCriteria_FullMethodName = "/google.ads.googleads.v14.services.SharedCriterionService/MutateSharedCriteria"
)

// SharedCriterionServiceClient is the client API for SharedCriterionService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type SharedCriterionServiceClient interface {
	// Creates or removes shared criteria. Operation statuses are returned.
	//
	// List of thrown errors:
	//
	//	[AuthenticationError]()
	//	[AuthorizationError]()
	//	[CriterionError]()
	//	[DatabaseError]()
	//	[DistinctError]()
	//	[FieldError]()
	//	[HeaderError]()
	//	[IdError]()
	//	[InternalError]()
	//	[MutateError]()
	//	[NotEmptyError]()
	//	[NullError]()
	//	[OperatorError]()
	//	[QuotaError]()
	//	[RangeError]()
	//	[RequestError]()
	//	[ResourceCountLimitExceededError]()
	//	[SizeLimitError]()
	//	[StringFormatError]()
	//	[StringLengthError]()
	MutateSharedCriteria(ctx context.Context, in *MutateSharedCriteriaRequest, opts ...grpc.CallOption) (*MutateSharedCriteriaResponse, error)
}

type sharedCriterionServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewSharedCriterionServiceClient(cc grpc.ClientConnInterface) SharedCriterionServiceClient {
	return &sharedCriterionServiceClient{cc}
}

func (c *sharedCriterionServiceClient) MutateSharedCriteria(ctx context.Context, in *MutateSharedCriteriaRequest, opts ...grpc.CallOption) (*MutateSharedCriteriaResponse, error) {
	out := new(MutateSharedCriteriaResponse)
	err := c.cc.Invoke(ctx, SharedCriterionService_MutateSharedCriteria_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// SharedCriterionServiceServer is the server API for SharedCriterionService service.
// All implementations must embed UnimplementedSharedCriterionServiceServer
// for forward compatibility
type SharedCriterionServiceServer interface {
	// Creates or removes shared criteria. Operation statuses are returned.
	//
	// List of thrown errors:
	//
	//	[AuthenticationError]()
	//	[AuthorizationError]()
	//	[CriterionError]()
	//	[DatabaseError]()
	//	[DistinctError]()
	//	[FieldError]()
	//	[HeaderError]()
	//	[IdError]()
	//	[InternalError]()
	//	[MutateError]()
	//	[NotEmptyError]()
	//	[NullError]()
	//	[OperatorError]()
	//	[QuotaError]()
	//	[RangeError]()
	//	[RequestError]()
	//	[ResourceCountLimitExceededError]()
	//	[SizeLimitError]()
	//	[StringFormatError]()
	//	[StringLengthError]()
	MutateSharedCriteria(context.Context, *MutateSharedCriteriaRequest) (*MutateSharedCriteriaResponse, error)
	mustEmbedUnimplementedSharedCriterionServiceServer()
}

// UnimplementedSharedCriterionServiceServer must be embedded to have forward compatible implementations.
type UnimplementedSharedCriterionServiceServer struct {
}

func (UnimplementedSharedCriterionServiceServer) MutateSharedCriteria(context.Context, *MutateSharedCriteriaRequest) (*MutateSharedCriteriaResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method MutateSharedCriteria not implemented")
}
func (UnimplementedSharedCriterionServiceServer) mustEmbedUnimplementedSharedCriterionServiceServer() {
}

// UnsafeSharedCriterionServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to SharedCriterionServiceServer will
// result in compilation errors.
type UnsafeSharedCriterionServiceServer interface {
	mustEmbedUnimplementedSharedCriterionServiceServer()
}

func RegisterSharedCriterionServiceServer(s grpc.ServiceRegistrar, srv SharedCriterionServiceServer) {
	s.RegisterService(&SharedCriterionService_ServiceDesc, srv)
}

func _SharedCriterionService_MutateSharedCriteria_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MutateSharedCriteriaRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SharedCriterionServiceServer).MutateSharedCriteria(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: SharedCriterionService_MutateSharedCriteria_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SharedCriterionServiceServer).MutateSharedCriteria(ctx, req.(*MutateSharedCriteriaRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// SharedCriterionService_ServiceDesc is the grpc.ServiceDesc for SharedCriterionService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var SharedCriterionService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "google.ads.googleads.v14.services.SharedCriterionService",
	HandlerType: (*SharedCriterionServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "MutateSharedCriteria",
			Handler:    _SharedCriterionService_MutateSharedCriteria_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "google/ads/googleads/v14/services/shared_criterion_service.proto",
}
