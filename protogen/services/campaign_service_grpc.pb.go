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
// source: google/ads/googleads/v14/services/campaign_service.proto

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
	CampaignService_MutateCampaigns_FullMethodName = "/google.ads.googleads.v14.services.CampaignService/MutateCampaigns"
)

// CampaignServiceClient is the client API for CampaignService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type CampaignServiceClient interface {
	// Creates, updates, or removes campaigns. Operation statuses are returned.
	//
	// List of thrown errors:
	//
	//	[AdxError]()
	//	[AuthenticationError]()
	//	[AuthorizationError]()
	//	[BiddingError]()
	//	[BiddingStrategyError]()
	//	[CampaignBudgetError]()
	//	[CampaignError]()
	//	[ContextError]()
	//	[DatabaseError]()
	//	[DateError]()
	//	[DateRangeError]()
	//	[DistinctError]()
	//	[FieldError]()
	//	[FieldMaskError]()
	//	[HeaderError]()
	//	[IdError]()
	//	[InternalError]()
	//	[ListOperationError]()
	//	[MutateError]()
	//	[NewResourceCreationError]()
	//	[NotAllowlistedError]()
	//	[NotEmptyError]()
	//	[NullError]()
	//	[OperationAccessDeniedError]()
	//	[OperatorError]()
	//	[QuotaError]()
	//	[RangeError]()
	//	[RegionCodeError]()
	//	[RequestError]()
	//	[ResourceCountLimitExceededError]()
	//	[SettingError]()
	//	[SizeLimitError]()
	//	[StringFormatError]()
	//	[StringLengthError]()
	//	[UrlFieldError]()
	MutateCampaigns(ctx context.Context, in *MutateCampaignsRequest, opts ...grpc.CallOption) (*MutateCampaignsResponse, error)
}

type campaignServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewCampaignServiceClient(cc grpc.ClientConnInterface) CampaignServiceClient {
	return &campaignServiceClient{cc}
}

func (c *campaignServiceClient) MutateCampaigns(ctx context.Context, in *MutateCampaignsRequest, opts ...grpc.CallOption) (*MutateCampaignsResponse, error) {
	out := new(MutateCampaignsResponse)
	err := c.cc.Invoke(ctx, CampaignService_MutateCampaigns_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// CampaignServiceServer is the server API for CampaignService service.
// All implementations must embed UnimplementedCampaignServiceServer
// for forward compatibility
type CampaignServiceServer interface {
	// Creates, updates, or removes campaigns. Operation statuses are returned.
	//
	// List of thrown errors:
	//
	//	[AdxError]()
	//	[AuthenticationError]()
	//	[AuthorizationError]()
	//	[BiddingError]()
	//	[BiddingStrategyError]()
	//	[CampaignBudgetError]()
	//	[CampaignError]()
	//	[ContextError]()
	//	[DatabaseError]()
	//	[DateError]()
	//	[DateRangeError]()
	//	[DistinctError]()
	//	[FieldError]()
	//	[FieldMaskError]()
	//	[HeaderError]()
	//	[IdError]()
	//	[InternalError]()
	//	[ListOperationError]()
	//	[MutateError]()
	//	[NewResourceCreationError]()
	//	[NotAllowlistedError]()
	//	[NotEmptyError]()
	//	[NullError]()
	//	[OperationAccessDeniedError]()
	//	[OperatorError]()
	//	[QuotaError]()
	//	[RangeError]()
	//	[RegionCodeError]()
	//	[RequestError]()
	//	[ResourceCountLimitExceededError]()
	//	[SettingError]()
	//	[SizeLimitError]()
	//	[StringFormatError]()
	//	[StringLengthError]()
	//	[UrlFieldError]()
	MutateCampaigns(context.Context, *MutateCampaignsRequest) (*MutateCampaignsResponse, error)
	mustEmbedUnimplementedCampaignServiceServer()
}

// UnimplementedCampaignServiceServer must be embedded to have forward compatible implementations.
type UnimplementedCampaignServiceServer struct {
}

func (UnimplementedCampaignServiceServer) MutateCampaigns(context.Context, *MutateCampaignsRequest) (*MutateCampaignsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method MutateCampaigns not implemented")
}
func (UnimplementedCampaignServiceServer) mustEmbedUnimplementedCampaignServiceServer() {}

// UnsafeCampaignServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to CampaignServiceServer will
// result in compilation errors.
type UnsafeCampaignServiceServer interface {
	mustEmbedUnimplementedCampaignServiceServer()
}

func RegisterCampaignServiceServer(s grpc.ServiceRegistrar, srv CampaignServiceServer) {
	s.RegisterService(&CampaignService_ServiceDesc, srv)
}

func _CampaignService_MutateCampaigns_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MutateCampaignsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CampaignServiceServer).MutateCampaigns(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: CampaignService_MutateCampaigns_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CampaignServiceServer).MutateCampaigns(ctx, req.(*MutateCampaignsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// CampaignService_ServiceDesc is the grpc.ServiceDesc for CampaignService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var CampaignService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "google.ads.googleads.v14.services.CampaignService",
	HandlerType: (*CampaignServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "MutateCampaigns",
			Handler:    _CampaignService_MutateCampaigns_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "google/ads/googleads/v14/services/campaign_service.proto",
}
