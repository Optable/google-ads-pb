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
// - protoc             v3.21.9
// source: google/ads/googleads/v13/services/user_list_service.proto

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
	UserListService_MutateUserLists_FullMethodName = "/google.ads.googleads.v13.services.UserListService/MutateUserLists"
)

// UserListServiceClient is the client API for UserListService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type UserListServiceClient interface {
	// Creates or updates user lists. Operation statuses are returned.
	//
	// List of thrown errors:
	//
	//	[AuthenticationError]()
	//	[AuthorizationError]()
	//	[CollectionSizeError]()
	//	[DatabaseError]()
	//	[DistinctError]()
	//	[FieldError]()
	//	[FieldMaskError]()
	//	[HeaderError]()
	//	[InternalError]()
	//	[MutateError]()
	//	[NewResourceCreationError]()
	//	[NotAllowlistedError]()
	//	[NotEmptyError]()
	//	[OperationAccessDeniedError]()
	//	[QuotaError]()
	//	[RangeError]()
	//	[RequestError]()
	//	[StringFormatError]()
	//	[StringLengthError]()
	//	[UserListError]()
	MutateUserLists(ctx context.Context, in *MutateUserListsRequest, opts ...grpc.CallOption) (*MutateUserListsResponse, error)
}

type userListServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewUserListServiceClient(cc grpc.ClientConnInterface) UserListServiceClient {
	return &userListServiceClient{cc}
}

func (c *userListServiceClient) MutateUserLists(ctx context.Context, in *MutateUserListsRequest, opts ...grpc.CallOption) (*MutateUserListsResponse, error) {
	out := new(MutateUserListsResponse)
	err := c.cc.Invoke(ctx, UserListService_MutateUserLists_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// UserListServiceServer is the server API for UserListService service.
// All implementations must embed UnimplementedUserListServiceServer
// for forward compatibility
type UserListServiceServer interface {
	// Creates or updates user lists. Operation statuses are returned.
	//
	// List of thrown errors:
	//
	//	[AuthenticationError]()
	//	[AuthorizationError]()
	//	[CollectionSizeError]()
	//	[DatabaseError]()
	//	[DistinctError]()
	//	[FieldError]()
	//	[FieldMaskError]()
	//	[HeaderError]()
	//	[InternalError]()
	//	[MutateError]()
	//	[NewResourceCreationError]()
	//	[NotAllowlistedError]()
	//	[NotEmptyError]()
	//	[OperationAccessDeniedError]()
	//	[QuotaError]()
	//	[RangeError]()
	//	[RequestError]()
	//	[StringFormatError]()
	//	[StringLengthError]()
	//	[UserListError]()
	MutateUserLists(context.Context, *MutateUserListsRequest) (*MutateUserListsResponse, error)
	mustEmbedUnimplementedUserListServiceServer()
}

// UnimplementedUserListServiceServer must be embedded to have forward compatible implementations.
type UnimplementedUserListServiceServer struct {
}

func (UnimplementedUserListServiceServer) MutateUserLists(context.Context, *MutateUserListsRequest) (*MutateUserListsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method MutateUserLists not implemented")
}
func (UnimplementedUserListServiceServer) mustEmbedUnimplementedUserListServiceServer() {}

// UnsafeUserListServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to UserListServiceServer will
// result in compilation errors.
type UnsafeUserListServiceServer interface {
	mustEmbedUnimplementedUserListServiceServer()
}

func RegisterUserListServiceServer(s grpc.ServiceRegistrar, srv UserListServiceServer) {
	s.RegisterService(&UserListService_ServiceDesc, srv)
}

func _UserListService_MutateUserLists_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MutateUserListsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserListServiceServer).MutateUserLists(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: UserListService_MutateUserLists_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserListServiceServer).MutateUserLists(ctx, req.(*MutateUserListsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// UserListService_ServiceDesc is the grpc.ServiceDesc for UserListService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var UserListService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "google.ads.googleads.v13.services.UserListService",
	HandlerType: (*UserListServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "MutateUserLists",
			Handler:    _UserListService_MutateUserLists_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "google/ads/googleads/v13/services/user_list_service.proto",
}
