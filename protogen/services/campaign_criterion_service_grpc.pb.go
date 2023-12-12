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
// source: google/ads/googleads/v14/services/campaign_criterion_service.proto

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
	CampaignCriterionService_MutateCampaignCriteria_FullMethodName = "/google.ads.googleads.v14.services.CampaignCriterionService/MutateCampaignCriteria"
)

// CampaignCriterionServiceClient is the client API for CampaignCriterionService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type CampaignCriterionServiceClient interface {
	// Creates, updates, or removes criteria. Operation statuses are returned.
	//
	// List of thrown errors:
	//
	//	[AdxError]()
	//	[AuthenticationError]()
	//	[AuthorizationError]()
	//	[CampaignCriterionError]()
	//	[CollectionSizeError]()
	//	[ContextError]()
	//	[CriterionError]()
	//	[DatabaseError]()
	//	[DistinctError]()
	//	[FieldError]()
	//	[FieldMaskError]()
	//	[FunctionError]()
	//	[HeaderError]()
	//	[IdError]()
	//	[InternalError]()
	//	[MutateError]()
	//	[NewResourceCreationError]()
	//	[NotEmptyError]()
	//	[NullError]()
	//	[OperationAccessDeniedError]()
	//	[OperatorError]()
	//	[QuotaError]()
	//	[RangeError]()
	//	[RegionCodeError]()
	//	[RequestError]()
	//	[ResourceCountLimitExceededError]()
	//	[SizeLimitError]()
	//	[StringFormatError]()
	//	[StringLengthError]()
	MutateCampaignCriteria(ctx context.Context, in *MutateCampaignCriteriaRequest, opts ...grpc.CallOption) (*MutateCampaignCriteriaResponse, error)
}

type campaignCriterionServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewCampaignCriterionServiceClient(cc grpc.ClientConnInterface) CampaignCriterionServiceClient {
	return &campaignCriterionServiceClient{cc}
}

func (c *campaignCriterionServiceClient) MutateCampaignCriteria(ctx context.Context, in *MutateCampaignCriteriaRequest, opts ...grpc.CallOption) (*MutateCampaignCriteriaResponse, error) {
	out := new(MutateCampaignCriteriaResponse)
	err := c.cc.Invoke(ctx, CampaignCriterionService_MutateCampaignCriteria_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// CampaignCriterionServiceServer is the server API for CampaignCriterionService service.
// All implementations must embed UnimplementedCampaignCriterionServiceServer
// for forward compatibility
type CampaignCriterionServiceServer interface {
	// Creates, updates, or removes criteria. Operation statuses are returned.
	//
	// List of thrown errors:
	//
	//	[AdxError]()
	//	[AuthenticationError]()
	//	[AuthorizationError]()
	//	[CampaignCriterionError]()
	//	[CollectionSizeError]()
	//	[ContextError]()
	//	[CriterionError]()
	//	[DatabaseError]()
	//	[DistinctError]()
	//	[FieldError]()
	//	[FieldMaskError]()
	//	[FunctionError]()
	//	[HeaderError]()
	//	[IdError]()
	//	[InternalError]()
	//	[MutateError]()
	//	[NewResourceCreationError]()
	//	[NotEmptyError]()
	//	[NullError]()
	//	[OperationAccessDeniedError]()
	//	[OperatorError]()
	//	[QuotaError]()
	//	[RangeError]()
	//	[RegionCodeError]()
	//	[RequestError]()
	//	[ResourceCountLimitExceededError]()
	//	[SizeLimitError]()
	//	[StringFormatError]()
	//	[StringLengthError]()
	MutateCampaignCriteria(context.Context, *MutateCampaignCriteriaRequest) (*MutateCampaignCriteriaResponse, error)
	mustEmbedUnimplementedCampaignCriterionServiceServer()
}

// UnimplementedCampaignCriterionServiceServer must be embedded to have forward compatible implementations.
type UnimplementedCampaignCriterionServiceServer struct {
}

func (UnimplementedCampaignCriterionServiceServer) MutateCampaignCriteria(context.Context, *MutateCampaignCriteriaRequest) (*MutateCampaignCriteriaResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method MutateCampaignCriteria not implemented")
}
func (UnimplementedCampaignCriterionServiceServer) mustEmbedUnimplementedCampaignCriterionServiceServer() {
}

// UnsafeCampaignCriterionServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to CampaignCriterionServiceServer will
// result in compilation errors.
type UnsafeCampaignCriterionServiceServer interface {
	mustEmbedUnimplementedCampaignCriterionServiceServer()
}

func RegisterCampaignCriterionServiceServer(s grpc.ServiceRegistrar, srv CampaignCriterionServiceServer) {
	s.RegisterService(&CampaignCriterionService_ServiceDesc, srv)
}

func _CampaignCriterionService_MutateCampaignCriteria_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MutateCampaignCriteriaRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CampaignCriterionServiceServer).MutateCampaignCriteria(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: CampaignCriterionService_MutateCampaignCriteria_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CampaignCriterionServiceServer).MutateCampaignCriteria(ctx, req.(*MutateCampaignCriteriaRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// CampaignCriterionService_ServiceDesc is the grpc.ServiceDesc for CampaignCriterionService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var CampaignCriterionService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "google.ads.googleads.v14.services.CampaignCriterionService",
	HandlerType: (*CampaignCriterionServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "MutateCampaignCriteria",
			Handler:    _CampaignCriterionService_MutateCampaignCriteria_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "google/ads/googleads/v14/services/campaign_criterion_service.proto",
}
