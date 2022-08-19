// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.21.5
// source: google/ads/googleads/v11/services/keyword_theme_constant_service.proto

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

// KeywordThemeConstantServiceClient is the client API for KeywordThemeConstantService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type KeywordThemeConstantServiceClient interface {
	// Returns KeywordThemeConstant suggestions by keyword themes.
	//
	// List of thrown errors:
	//   [AuthenticationError]()
	//   [AuthorizationError]()
	//   [HeaderError]()
	//   [InternalError]()
	//   [QuotaError]()
	//   [RequestError]()
	SuggestKeywordThemeConstants(ctx context.Context, in *SuggestKeywordThemeConstantsRequest, opts ...grpc.CallOption) (*SuggestKeywordThemeConstantsResponse, error)
}

type keywordThemeConstantServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewKeywordThemeConstantServiceClient(cc grpc.ClientConnInterface) KeywordThemeConstantServiceClient {
	return &keywordThemeConstantServiceClient{cc}
}

func (c *keywordThemeConstantServiceClient) SuggestKeywordThemeConstants(ctx context.Context, in *SuggestKeywordThemeConstantsRequest, opts ...grpc.CallOption) (*SuggestKeywordThemeConstantsResponse, error) {
	out := new(SuggestKeywordThemeConstantsResponse)
	err := c.cc.Invoke(ctx, "/google.ads.googleads.v11.services.KeywordThemeConstantService/SuggestKeywordThemeConstants", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// KeywordThemeConstantServiceServer is the server API for KeywordThemeConstantService service.
// All implementations must embed UnimplementedKeywordThemeConstantServiceServer
// for forward compatibility
type KeywordThemeConstantServiceServer interface {
	// Returns KeywordThemeConstant suggestions by keyword themes.
	//
	// List of thrown errors:
	//   [AuthenticationError]()
	//   [AuthorizationError]()
	//   [HeaderError]()
	//   [InternalError]()
	//   [QuotaError]()
	//   [RequestError]()
	SuggestKeywordThemeConstants(context.Context, *SuggestKeywordThemeConstantsRequest) (*SuggestKeywordThemeConstantsResponse, error)
	mustEmbedUnimplementedKeywordThemeConstantServiceServer()
}

// UnimplementedKeywordThemeConstantServiceServer must be embedded to have forward compatible implementations.
type UnimplementedKeywordThemeConstantServiceServer struct {
}

func (UnimplementedKeywordThemeConstantServiceServer) SuggestKeywordThemeConstants(context.Context, *SuggestKeywordThemeConstantsRequest) (*SuggestKeywordThemeConstantsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SuggestKeywordThemeConstants not implemented")
}
func (UnimplementedKeywordThemeConstantServiceServer) mustEmbedUnimplementedKeywordThemeConstantServiceServer() {
}

// UnsafeKeywordThemeConstantServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to KeywordThemeConstantServiceServer will
// result in compilation errors.
type UnsafeKeywordThemeConstantServiceServer interface {
	mustEmbedUnimplementedKeywordThemeConstantServiceServer()
}

func RegisterKeywordThemeConstantServiceServer(s grpc.ServiceRegistrar, srv KeywordThemeConstantServiceServer) {
	s.RegisterService(&KeywordThemeConstantService_ServiceDesc, srv)
}

func _KeywordThemeConstantService_SuggestKeywordThemeConstants_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SuggestKeywordThemeConstantsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(KeywordThemeConstantServiceServer).SuggestKeywordThemeConstants(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/google.ads.googleads.v11.services.KeywordThemeConstantService/SuggestKeywordThemeConstants",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(KeywordThemeConstantServiceServer).SuggestKeywordThemeConstants(ctx, req.(*SuggestKeywordThemeConstantsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// KeywordThemeConstantService_ServiceDesc is the grpc.ServiceDesc for KeywordThemeConstantService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var KeywordThemeConstantService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "google.ads.googleads.v11.services.KeywordThemeConstantService",
	HandlerType: (*KeywordThemeConstantServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "SuggestKeywordThemeConstants",
			Handler:    _KeywordThemeConstantService_SuggestKeywordThemeConstants_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "google/ads/googleads/v11/services/keyword_theme_constant_service.proto",
}
