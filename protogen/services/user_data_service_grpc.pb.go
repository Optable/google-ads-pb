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
// source: google/ads/googleads/v15/services/user_data_service.proto

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
	UserDataService_UploadUserData_FullMethodName = "/google.ads.googleads.v15.services.UserDataService/UploadUserData"
)

// UserDataServiceClient is the client API for UserDataService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type UserDataServiceClient interface {
	// Uploads the given user data.
	//
	// List of thrown errors:
	//
	//	[AuthenticationError]()
	//	[AuthorizationError]()
	//	[CollectionSizeError]()
	//	[FieldError]()
	//	[HeaderError]()
	//	[InternalError]()
	//	[MutateError]()
	//	[OfflineUserDataJobError]()
	//	[QuotaError]()
	//	[RequestError]()
	//	[UserDataError]()
	UploadUserData(ctx context.Context, in *UploadUserDataRequest, opts ...grpc.CallOption) (*UploadUserDataResponse, error)
}

type userDataServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewUserDataServiceClient(cc grpc.ClientConnInterface) UserDataServiceClient {
	return &userDataServiceClient{cc}
}

func (c *userDataServiceClient) UploadUserData(ctx context.Context, in *UploadUserDataRequest, opts ...grpc.CallOption) (*UploadUserDataResponse, error) {
	out := new(UploadUserDataResponse)
	err := c.cc.Invoke(ctx, UserDataService_UploadUserData_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// UserDataServiceServer is the server API for UserDataService service.
// All implementations must embed UnimplementedUserDataServiceServer
// for forward compatibility
type UserDataServiceServer interface {
	// Uploads the given user data.
	//
	// List of thrown errors:
	//
	//	[AuthenticationError]()
	//	[AuthorizationError]()
	//	[CollectionSizeError]()
	//	[FieldError]()
	//	[HeaderError]()
	//	[InternalError]()
	//	[MutateError]()
	//	[OfflineUserDataJobError]()
	//	[QuotaError]()
	//	[RequestError]()
	//	[UserDataError]()
	UploadUserData(context.Context, *UploadUserDataRequest) (*UploadUserDataResponse, error)
	mustEmbedUnimplementedUserDataServiceServer()
}

// UnimplementedUserDataServiceServer must be embedded to have forward compatible implementations.
type UnimplementedUserDataServiceServer struct {
}

func (UnimplementedUserDataServiceServer) UploadUserData(context.Context, *UploadUserDataRequest) (*UploadUserDataResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UploadUserData not implemented")
}
func (UnimplementedUserDataServiceServer) mustEmbedUnimplementedUserDataServiceServer() {}

// UnsafeUserDataServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to UserDataServiceServer will
// result in compilation errors.
type UnsafeUserDataServiceServer interface {
	mustEmbedUnimplementedUserDataServiceServer()
}

func RegisterUserDataServiceServer(s grpc.ServiceRegistrar, srv UserDataServiceServer) {
	s.RegisterService(&UserDataService_ServiceDesc, srv)
}

func _UserDataService_UploadUserData_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UploadUserDataRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserDataServiceServer).UploadUserData(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: UserDataService_UploadUserData_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserDataServiceServer).UploadUserData(ctx, req.(*UploadUserDataRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// UserDataService_ServiceDesc is the grpc.ServiceDesc for UserDataService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var UserDataService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "google.ads.googleads.v15.services.UserDataService",
	HandlerType: (*UserDataServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "UploadUserData",
			Handler:    _UserDataService_UploadUserData_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "google/ads/googleads/v15/services/user_data_service.proto",
}
