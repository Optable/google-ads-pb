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
// source: google/ads/googleads/v17/services/campaign_asset_set_service.proto

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
	CampaignAssetSetService_MutateCampaignAssetSets_FullMethodName = "/google.ads.googleads.v17.services.CampaignAssetSetService/MutateCampaignAssetSets"
)

// CampaignAssetSetServiceClient is the client API for CampaignAssetSetService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
//
// Service to manage campaign asset set
type CampaignAssetSetServiceClient interface {
	// Creates, updates or removes campaign asset sets. Operation statuses are
	// returned.
	MutateCampaignAssetSets(ctx context.Context, in *MutateCampaignAssetSetsRequest, opts ...grpc.CallOption) (*MutateCampaignAssetSetsResponse, error)
}

type campaignAssetSetServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewCampaignAssetSetServiceClient(cc grpc.ClientConnInterface) CampaignAssetSetServiceClient {
	return &campaignAssetSetServiceClient{cc}
}

func (c *campaignAssetSetServiceClient) MutateCampaignAssetSets(ctx context.Context, in *MutateCampaignAssetSetsRequest, opts ...grpc.CallOption) (*MutateCampaignAssetSetsResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(MutateCampaignAssetSetsResponse)
	err := c.cc.Invoke(ctx, CampaignAssetSetService_MutateCampaignAssetSets_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// CampaignAssetSetServiceServer is the server API for CampaignAssetSetService service.
// All implementations must embed UnimplementedCampaignAssetSetServiceServer
// for forward compatibility.
//
// Service to manage campaign asset set
type CampaignAssetSetServiceServer interface {
	// Creates, updates or removes campaign asset sets. Operation statuses are
	// returned.
	MutateCampaignAssetSets(context.Context, *MutateCampaignAssetSetsRequest) (*MutateCampaignAssetSetsResponse, error)
	mustEmbedUnimplementedCampaignAssetSetServiceServer()
}

// UnimplementedCampaignAssetSetServiceServer must be embedded to have
// forward compatible implementations.
//
// NOTE: this should be embedded by value instead of pointer to avoid a nil
// pointer dereference when methods are called.
type UnimplementedCampaignAssetSetServiceServer struct{}

func (UnimplementedCampaignAssetSetServiceServer) MutateCampaignAssetSets(context.Context, *MutateCampaignAssetSetsRequest) (*MutateCampaignAssetSetsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method MutateCampaignAssetSets not implemented")
}
func (UnimplementedCampaignAssetSetServiceServer) mustEmbedUnimplementedCampaignAssetSetServiceServer() {
}
func (UnimplementedCampaignAssetSetServiceServer) testEmbeddedByValue() {}

// UnsafeCampaignAssetSetServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to CampaignAssetSetServiceServer will
// result in compilation errors.
type UnsafeCampaignAssetSetServiceServer interface {
	mustEmbedUnimplementedCampaignAssetSetServiceServer()
}

func RegisterCampaignAssetSetServiceServer(s grpc.ServiceRegistrar, srv CampaignAssetSetServiceServer) {
	// If the following call pancis, it indicates UnimplementedCampaignAssetSetServiceServer was
	// embedded by pointer and is nil.  This will cause panics if an
	// unimplemented method is ever invoked, so we test this at initialization
	// time to prevent it from happening at runtime later due to I/O.
	if t, ok := srv.(interface{ testEmbeddedByValue() }); ok {
		t.testEmbeddedByValue()
	}
	s.RegisterService(&CampaignAssetSetService_ServiceDesc, srv)
}

func _CampaignAssetSetService_MutateCampaignAssetSets_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MutateCampaignAssetSetsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CampaignAssetSetServiceServer).MutateCampaignAssetSets(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: CampaignAssetSetService_MutateCampaignAssetSets_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CampaignAssetSetServiceServer).MutateCampaignAssetSets(ctx, req.(*MutateCampaignAssetSetsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// CampaignAssetSetService_ServiceDesc is the grpc.ServiceDesc for CampaignAssetSetService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var CampaignAssetSetService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "google.ads.googleads.v17.services.CampaignAssetSetService",
	HandlerType: (*CampaignAssetSetServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "MutateCampaignAssetSets",
			Handler:    _CampaignAssetSetService_MutateCampaignAssetSets_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "google/ads/googleads/v17/services/campaign_asset_set_service.proto",
}
