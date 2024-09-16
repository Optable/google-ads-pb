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
// source: google/ads/googleads/v17/services/keyword_plan_idea_service.proto

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
	KeywordPlanIdeaService_GenerateKeywordIdeas_FullMethodName             = "/google.ads.googleads.v17.services.KeywordPlanIdeaService/GenerateKeywordIdeas"
	KeywordPlanIdeaService_GenerateKeywordHistoricalMetrics_FullMethodName = "/google.ads.googleads.v17.services.KeywordPlanIdeaService/GenerateKeywordHistoricalMetrics"
	KeywordPlanIdeaService_GenerateAdGroupThemes_FullMethodName            = "/google.ads.googleads.v17.services.KeywordPlanIdeaService/GenerateAdGroupThemes"
	KeywordPlanIdeaService_GenerateKeywordForecastMetrics_FullMethodName   = "/google.ads.googleads.v17.services.KeywordPlanIdeaService/GenerateKeywordForecastMetrics"
)

// KeywordPlanIdeaServiceClient is the client API for KeywordPlanIdeaService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
//
// Service to generate keyword ideas.
type KeywordPlanIdeaServiceClient interface {
	// Returns a list of keyword ideas.
	//
	// List of thrown errors:
	//
	//	[AuthenticationError]()
	//	[AuthorizationError]()
	//	[CollectionSizeError]()
	//	[HeaderError]()
	//	[InternalError]()
	//	[KeywordPlanIdeaError]()
	//	[QuotaError]()
	//	[RequestError]()
	GenerateKeywordIdeas(ctx context.Context, in *GenerateKeywordIdeasRequest, opts ...grpc.CallOption) (*GenerateKeywordIdeaResponse, error)
	// Returns a list of keyword historical metrics.
	//
	// List of thrown errors:
	//
	//	[AuthenticationError]()
	//	[AuthorizationError]()
	//	[CollectionSizeError]()
	//	[HeaderError]()
	//	[InternalError]()
	//	[QuotaError]()
	//	[RequestError]()
	GenerateKeywordHistoricalMetrics(ctx context.Context, in *GenerateKeywordHistoricalMetricsRequest, opts ...grpc.CallOption) (*GenerateKeywordHistoricalMetricsResponse, error)
	// Returns a list of suggested AdGroups and suggested modifications
	// (text, match type) for the given keywords.
	//
	// List of thrown errors:
	//
	//	[AuthenticationError]()
	//	[AuthorizationError]()
	//	[CollectionSizeError]()
	//	[HeaderError]()
	//	[InternalError]()
	//	[QuotaError]()
	//	[RequestError]()
	GenerateAdGroupThemes(ctx context.Context, in *GenerateAdGroupThemesRequest, opts ...grpc.CallOption) (*GenerateAdGroupThemesResponse, error)
	// Returns metrics (such as impressions, clicks, total cost) of a keyword
	// forecast for the given campaign.
	//
	// List of thrown errors:
	//
	//	[AuthenticationError]()
	//	[AuthorizationError]()
	//	[CollectionSizeError]()
	//	[HeaderError]()
	//	[InternalError]()
	//	[QuotaError]()
	//	[RequestError]()
	GenerateKeywordForecastMetrics(ctx context.Context, in *GenerateKeywordForecastMetricsRequest, opts ...grpc.CallOption) (*GenerateKeywordForecastMetricsResponse, error)
}

type keywordPlanIdeaServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewKeywordPlanIdeaServiceClient(cc grpc.ClientConnInterface) KeywordPlanIdeaServiceClient {
	return &keywordPlanIdeaServiceClient{cc}
}

func (c *keywordPlanIdeaServiceClient) GenerateKeywordIdeas(ctx context.Context, in *GenerateKeywordIdeasRequest, opts ...grpc.CallOption) (*GenerateKeywordIdeaResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(GenerateKeywordIdeaResponse)
	err := c.cc.Invoke(ctx, KeywordPlanIdeaService_GenerateKeywordIdeas_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *keywordPlanIdeaServiceClient) GenerateKeywordHistoricalMetrics(ctx context.Context, in *GenerateKeywordHistoricalMetricsRequest, opts ...grpc.CallOption) (*GenerateKeywordHistoricalMetricsResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(GenerateKeywordHistoricalMetricsResponse)
	err := c.cc.Invoke(ctx, KeywordPlanIdeaService_GenerateKeywordHistoricalMetrics_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *keywordPlanIdeaServiceClient) GenerateAdGroupThemes(ctx context.Context, in *GenerateAdGroupThemesRequest, opts ...grpc.CallOption) (*GenerateAdGroupThemesResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(GenerateAdGroupThemesResponse)
	err := c.cc.Invoke(ctx, KeywordPlanIdeaService_GenerateAdGroupThemes_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *keywordPlanIdeaServiceClient) GenerateKeywordForecastMetrics(ctx context.Context, in *GenerateKeywordForecastMetricsRequest, opts ...grpc.CallOption) (*GenerateKeywordForecastMetricsResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(GenerateKeywordForecastMetricsResponse)
	err := c.cc.Invoke(ctx, KeywordPlanIdeaService_GenerateKeywordForecastMetrics_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// KeywordPlanIdeaServiceServer is the server API for KeywordPlanIdeaService service.
// All implementations must embed UnimplementedKeywordPlanIdeaServiceServer
// for forward compatibility.
//
// Service to generate keyword ideas.
type KeywordPlanIdeaServiceServer interface {
	// Returns a list of keyword ideas.
	//
	// List of thrown errors:
	//
	//	[AuthenticationError]()
	//	[AuthorizationError]()
	//	[CollectionSizeError]()
	//	[HeaderError]()
	//	[InternalError]()
	//	[KeywordPlanIdeaError]()
	//	[QuotaError]()
	//	[RequestError]()
	GenerateKeywordIdeas(context.Context, *GenerateKeywordIdeasRequest) (*GenerateKeywordIdeaResponse, error)
	// Returns a list of keyword historical metrics.
	//
	// List of thrown errors:
	//
	//	[AuthenticationError]()
	//	[AuthorizationError]()
	//	[CollectionSizeError]()
	//	[HeaderError]()
	//	[InternalError]()
	//	[QuotaError]()
	//	[RequestError]()
	GenerateKeywordHistoricalMetrics(context.Context, *GenerateKeywordHistoricalMetricsRequest) (*GenerateKeywordHistoricalMetricsResponse, error)
	// Returns a list of suggested AdGroups and suggested modifications
	// (text, match type) for the given keywords.
	//
	// List of thrown errors:
	//
	//	[AuthenticationError]()
	//	[AuthorizationError]()
	//	[CollectionSizeError]()
	//	[HeaderError]()
	//	[InternalError]()
	//	[QuotaError]()
	//	[RequestError]()
	GenerateAdGroupThemes(context.Context, *GenerateAdGroupThemesRequest) (*GenerateAdGroupThemesResponse, error)
	// Returns metrics (such as impressions, clicks, total cost) of a keyword
	// forecast for the given campaign.
	//
	// List of thrown errors:
	//
	//	[AuthenticationError]()
	//	[AuthorizationError]()
	//	[CollectionSizeError]()
	//	[HeaderError]()
	//	[InternalError]()
	//	[QuotaError]()
	//	[RequestError]()
	GenerateKeywordForecastMetrics(context.Context, *GenerateKeywordForecastMetricsRequest) (*GenerateKeywordForecastMetricsResponse, error)
	mustEmbedUnimplementedKeywordPlanIdeaServiceServer()
}

// UnimplementedKeywordPlanIdeaServiceServer must be embedded to have
// forward compatible implementations.
//
// NOTE: this should be embedded by value instead of pointer to avoid a nil
// pointer dereference when methods are called.
type UnimplementedKeywordPlanIdeaServiceServer struct{}

func (UnimplementedKeywordPlanIdeaServiceServer) GenerateKeywordIdeas(context.Context, *GenerateKeywordIdeasRequest) (*GenerateKeywordIdeaResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GenerateKeywordIdeas not implemented")
}
func (UnimplementedKeywordPlanIdeaServiceServer) GenerateKeywordHistoricalMetrics(context.Context, *GenerateKeywordHistoricalMetricsRequest) (*GenerateKeywordHistoricalMetricsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GenerateKeywordHistoricalMetrics not implemented")
}
func (UnimplementedKeywordPlanIdeaServiceServer) GenerateAdGroupThemes(context.Context, *GenerateAdGroupThemesRequest) (*GenerateAdGroupThemesResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GenerateAdGroupThemes not implemented")
}
func (UnimplementedKeywordPlanIdeaServiceServer) GenerateKeywordForecastMetrics(context.Context, *GenerateKeywordForecastMetricsRequest) (*GenerateKeywordForecastMetricsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GenerateKeywordForecastMetrics not implemented")
}
func (UnimplementedKeywordPlanIdeaServiceServer) mustEmbedUnimplementedKeywordPlanIdeaServiceServer() {
}
func (UnimplementedKeywordPlanIdeaServiceServer) testEmbeddedByValue() {}

// UnsafeKeywordPlanIdeaServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to KeywordPlanIdeaServiceServer will
// result in compilation errors.
type UnsafeKeywordPlanIdeaServiceServer interface {
	mustEmbedUnimplementedKeywordPlanIdeaServiceServer()
}

func RegisterKeywordPlanIdeaServiceServer(s grpc.ServiceRegistrar, srv KeywordPlanIdeaServiceServer) {
	// If the following call pancis, it indicates UnimplementedKeywordPlanIdeaServiceServer was
	// embedded by pointer and is nil.  This will cause panics if an
	// unimplemented method is ever invoked, so we test this at initialization
	// time to prevent it from happening at runtime later due to I/O.
	if t, ok := srv.(interface{ testEmbeddedByValue() }); ok {
		t.testEmbeddedByValue()
	}
	s.RegisterService(&KeywordPlanIdeaService_ServiceDesc, srv)
}

func _KeywordPlanIdeaService_GenerateKeywordIdeas_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GenerateKeywordIdeasRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(KeywordPlanIdeaServiceServer).GenerateKeywordIdeas(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: KeywordPlanIdeaService_GenerateKeywordIdeas_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(KeywordPlanIdeaServiceServer).GenerateKeywordIdeas(ctx, req.(*GenerateKeywordIdeasRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _KeywordPlanIdeaService_GenerateKeywordHistoricalMetrics_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GenerateKeywordHistoricalMetricsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(KeywordPlanIdeaServiceServer).GenerateKeywordHistoricalMetrics(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: KeywordPlanIdeaService_GenerateKeywordHistoricalMetrics_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(KeywordPlanIdeaServiceServer).GenerateKeywordHistoricalMetrics(ctx, req.(*GenerateKeywordHistoricalMetricsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _KeywordPlanIdeaService_GenerateAdGroupThemes_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GenerateAdGroupThemesRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(KeywordPlanIdeaServiceServer).GenerateAdGroupThemes(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: KeywordPlanIdeaService_GenerateAdGroupThemes_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(KeywordPlanIdeaServiceServer).GenerateAdGroupThemes(ctx, req.(*GenerateAdGroupThemesRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _KeywordPlanIdeaService_GenerateKeywordForecastMetrics_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GenerateKeywordForecastMetricsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(KeywordPlanIdeaServiceServer).GenerateKeywordForecastMetrics(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: KeywordPlanIdeaService_GenerateKeywordForecastMetrics_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(KeywordPlanIdeaServiceServer).GenerateKeywordForecastMetrics(ctx, req.(*GenerateKeywordForecastMetricsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// KeywordPlanIdeaService_ServiceDesc is the grpc.ServiceDesc for KeywordPlanIdeaService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var KeywordPlanIdeaService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "google.ads.googleads.v17.services.KeywordPlanIdeaService",
	HandlerType: (*KeywordPlanIdeaServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GenerateKeywordIdeas",
			Handler:    _KeywordPlanIdeaService_GenerateKeywordIdeas_Handler,
		},
		{
			MethodName: "GenerateKeywordHistoricalMetrics",
			Handler:    _KeywordPlanIdeaService_GenerateKeywordHistoricalMetrics_Handler,
		},
		{
			MethodName: "GenerateAdGroupThemes",
			Handler:    _KeywordPlanIdeaService_GenerateAdGroupThemes_Handler,
		},
		{
			MethodName: "GenerateKeywordForecastMetrics",
			Handler:    _KeywordPlanIdeaService_GenerateKeywordForecastMetrics_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "google/ads/googleads/v17/services/keyword_plan_idea_service.proto",
}
