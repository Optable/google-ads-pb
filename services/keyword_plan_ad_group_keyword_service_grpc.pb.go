// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.21.12
// source: google/ads/googleads/v12/services/keyword_plan_ad_group_keyword_service.proto

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

// KeywordPlanAdGroupKeywordServiceClient is the client API for KeywordPlanAdGroupKeywordService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type KeywordPlanAdGroupKeywordServiceClient interface {
	// Creates, updates, or removes Keyword Plan ad group keywords. Operation
	// statuses are returned.
	//
	// List of thrown errors:
	//   [AuthenticationError]()
	//   [AuthorizationError]()
	//   [DatabaseError]()
	//   [FieldError]()
	//   [HeaderError]()
	//   [InternalError]()
	//   [KeywordPlanAdGroupKeywordError]()
	//   [KeywordPlanError]()
	//   [MutateError]()
	//   [QuotaError]()
	//   [RequestError]()
	//   [ResourceCountLimitExceededError]()
	MutateKeywordPlanAdGroupKeywords(ctx context.Context, in *MutateKeywordPlanAdGroupKeywordsRequest, opts ...grpc.CallOption) (*MutateKeywordPlanAdGroupKeywordsResponse, error)
}

type keywordPlanAdGroupKeywordServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewKeywordPlanAdGroupKeywordServiceClient(cc grpc.ClientConnInterface) KeywordPlanAdGroupKeywordServiceClient {
	return &keywordPlanAdGroupKeywordServiceClient{cc}
}

func (c *keywordPlanAdGroupKeywordServiceClient) MutateKeywordPlanAdGroupKeywords(ctx context.Context, in *MutateKeywordPlanAdGroupKeywordsRequest, opts ...grpc.CallOption) (*MutateKeywordPlanAdGroupKeywordsResponse, error) {
	out := new(MutateKeywordPlanAdGroupKeywordsResponse)
	err := c.cc.Invoke(ctx, "/google.ads.googleads.v12.services.KeywordPlanAdGroupKeywordService/MutateKeywordPlanAdGroupKeywords", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// KeywordPlanAdGroupKeywordServiceServer is the server API for KeywordPlanAdGroupKeywordService service.
// All implementations must embed UnimplementedKeywordPlanAdGroupKeywordServiceServer
// for forward compatibility
type KeywordPlanAdGroupKeywordServiceServer interface {
	// Creates, updates, or removes Keyword Plan ad group keywords. Operation
	// statuses are returned.
	//
	// List of thrown errors:
	//   [AuthenticationError]()
	//   [AuthorizationError]()
	//   [DatabaseError]()
	//   [FieldError]()
	//   [HeaderError]()
	//   [InternalError]()
	//   [KeywordPlanAdGroupKeywordError]()
	//   [KeywordPlanError]()
	//   [MutateError]()
	//   [QuotaError]()
	//   [RequestError]()
	//   [ResourceCountLimitExceededError]()
	MutateKeywordPlanAdGroupKeywords(context.Context, *MutateKeywordPlanAdGroupKeywordsRequest) (*MutateKeywordPlanAdGroupKeywordsResponse, error)
	mustEmbedUnimplementedKeywordPlanAdGroupKeywordServiceServer()
}

// UnimplementedKeywordPlanAdGroupKeywordServiceServer must be embedded to have forward compatible implementations.
type UnimplementedKeywordPlanAdGroupKeywordServiceServer struct {
}

func (UnimplementedKeywordPlanAdGroupKeywordServiceServer) MutateKeywordPlanAdGroupKeywords(context.Context, *MutateKeywordPlanAdGroupKeywordsRequest) (*MutateKeywordPlanAdGroupKeywordsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method MutateKeywordPlanAdGroupKeywords not implemented")
}
func (UnimplementedKeywordPlanAdGroupKeywordServiceServer) mustEmbedUnimplementedKeywordPlanAdGroupKeywordServiceServer() {
}

// UnsafeKeywordPlanAdGroupKeywordServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to KeywordPlanAdGroupKeywordServiceServer will
// result in compilation errors.
type UnsafeKeywordPlanAdGroupKeywordServiceServer interface {
	mustEmbedUnimplementedKeywordPlanAdGroupKeywordServiceServer()
}

func RegisterKeywordPlanAdGroupKeywordServiceServer(s grpc.ServiceRegistrar, srv KeywordPlanAdGroupKeywordServiceServer) {
	s.RegisterService(&KeywordPlanAdGroupKeywordService_ServiceDesc, srv)
}

func _KeywordPlanAdGroupKeywordService_MutateKeywordPlanAdGroupKeywords_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MutateKeywordPlanAdGroupKeywordsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(KeywordPlanAdGroupKeywordServiceServer).MutateKeywordPlanAdGroupKeywords(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/google.ads.googleads.v12.services.KeywordPlanAdGroupKeywordService/MutateKeywordPlanAdGroupKeywords",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(KeywordPlanAdGroupKeywordServiceServer).MutateKeywordPlanAdGroupKeywords(ctx, req.(*MutateKeywordPlanAdGroupKeywordsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// KeywordPlanAdGroupKeywordService_ServiceDesc is the grpc.ServiceDesc for KeywordPlanAdGroupKeywordService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var KeywordPlanAdGroupKeywordService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "google.ads.googleads.v12.services.KeywordPlanAdGroupKeywordService",
	HandlerType: (*KeywordPlanAdGroupKeywordServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "MutateKeywordPlanAdGroupKeywords",
			Handler:    _KeywordPlanAdGroupKeywordService_MutateKeywordPlanAdGroupKeywords_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "google/ads/googleads/v12/services/keyword_plan_ad_group_keyword_service.proto",
}
