// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package services

import (
	context "context"
	resources "github.com/shenzhencenter/google-ads-pb/resources"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

// MerchantCenterLinkServiceClient is the client API for MerchantCenterLinkService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type MerchantCenterLinkServiceClient interface {
	// Returns Merchant Center links available for this customer.
	//
	// List of thrown errors:
	//   [AuthenticationError]()
	//   [AuthorizationError]()
	//   [HeaderError]()
	//   [InternalError]()
	//   [QuotaError]()
	//   [RequestError]()
	ListMerchantCenterLinks(ctx context.Context, in *ListMerchantCenterLinksRequest, opts ...grpc.CallOption) (*ListMerchantCenterLinksResponse, error)
	// Returns the Merchant Center link in full detail.
	//
	// List of thrown errors:
	//   [AuthenticationError]()
	//   [AuthorizationError]()
	//   [HeaderError]()
	//   [InternalError]()
	//   [QuotaError]()
	//   [RequestError]()
	GetMerchantCenterLink(ctx context.Context, in *GetMerchantCenterLinkRequest, opts ...grpc.CallOption) (*resources.MerchantCenterLink, error)
	// Updates status or removes a Merchant Center link.
	//
	// List of thrown errors:
	//   [AuthenticationError]()
	//   [AuthorizationError]()
	//   [FieldMaskError]()
	//   [HeaderError]()
	//   [InternalError]()
	//   [QuotaError]()
	//   [RequestError]()
	MutateMerchantCenterLink(ctx context.Context, in *MutateMerchantCenterLinkRequest, opts ...grpc.CallOption) (*MutateMerchantCenterLinkResponse, error)
}

type merchantCenterLinkServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewMerchantCenterLinkServiceClient(cc grpc.ClientConnInterface) MerchantCenterLinkServiceClient {
	return &merchantCenterLinkServiceClient{cc}
}

func (c *merchantCenterLinkServiceClient) ListMerchantCenterLinks(ctx context.Context, in *ListMerchantCenterLinksRequest, opts ...grpc.CallOption) (*ListMerchantCenterLinksResponse, error) {
	out := new(ListMerchantCenterLinksResponse)
	err := c.cc.Invoke(ctx, "/google.ads.googleads.v10.services.MerchantCenterLinkService/ListMerchantCenterLinks", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *merchantCenterLinkServiceClient) GetMerchantCenterLink(ctx context.Context, in *GetMerchantCenterLinkRequest, opts ...grpc.CallOption) (*resources.MerchantCenterLink, error) {
	out := new(resources.MerchantCenterLink)
	err := c.cc.Invoke(ctx, "/google.ads.googleads.v10.services.MerchantCenterLinkService/GetMerchantCenterLink", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *merchantCenterLinkServiceClient) MutateMerchantCenterLink(ctx context.Context, in *MutateMerchantCenterLinkRequest, opts ...grpc.CallOption) (*MutateMerchantCenterLinkResponse, error) {
	out := new(MutateMerchantCenterLinkResponse)
	err := c.cc.Invoke(ctx, "/google.ads.googleads.v10.services.MerchantCenterLinkService/MutateMerchantCenterLink", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// MerchantCenterLinkServiceServer is the server API for MerchantCenterLinkService service.
// All implementations must embed UnimplementedMerchantCenterLinkServiceServer
// for forward compatibility
type MerchantCenterLinkServiceServer interface {
	// Returns Merchant Center links available for this customer.
	//
	// List of thrown errors:
	//   [AuthenticationError]()
	//   [AuthorizationError]()
	//   [HeaderError]()
	//   [InternalError]()
	//   [QuotaError]()
	//   [RequestError]()
	ListMerchantCenterLinks(context.Context, *ListMerchantCenterLinksRequest) (*ListMerchantCenterLinksResponse, error)
	// Returns the Merchant Center link in full detail.
	//
	// List of thrown errors:
	//   [AuthenticationError]()
	//   [AuthorizationError]()
	//   [HeaderError]()
	//   [InternalError]()
	//   [QuotaError]()
	//   [RequestError]()
	GetMerchantCenterLink(context.Context, *GetMerchantCenterLinkRequest) (*resources.MerchantCenterLink, error)
	// Updates status or removes a Merchant Center link.
	//
	// List of thrown errors:
	//   [AuthenticationError]()
	//   [AuthorizationError]()
	//   [FieldMaskError]()
	//   [HeaderError]()
	//   [InternalError]()
	//   [QuotaError]()
	//   [RequestError]()
	MutateMerchantCenterLink(context.Context, *MutateMerchantCenterLinkRequest) (*MutateMerchantCenterLinkResponse, error)
	mustEmbedUnimplementedMerchantCenterLinkServiceServer()
}

// UnimplementedMerchantCenterLinkServiceServer must be embedded to have forward compatible implementations.
type UnimplementedMerchantCenterLinkServiceServer struct {
}

func (UnimplementedMerchantCenterLinkServiceServer) ListMerchantCenterLinks(context.Context, *ListMerchantCenterLinksRequest) (*ListMerchantCenterLinksResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListMerchantCenterLinks not implemented")
}
func (UnimplementedMerchantCenterLinkServiceServer) GetMerchantCenterLink(context.Context, *GetMerchantCenterLinkRequest) (*resources.MerchantCenterLink, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetMerchantCenterLink not implemented")
}
func (UnimplementedMerchantCenterLinkServiceServer) MutateMerchantCenterLink(context.Context, *MutateMerchantCenterLinkRequest) (*MutateMerchantCenterLinkResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method MutateMerchantCenterLink not implemented")
}
func (UnimplementedMerchantCenterLinkServiceServer) mustEmbedUnimplementedMerchantCenterLinkServiceServer() {
}

// UnsafeMerchantCenterLinkServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to MerchantCenterLinkServiceServer will
// result in compilation errors.
type UnsafeMerchantCenterLinkServiceServer interface {
	mustEmbedUnimplementedMerchantCenterLinkServiceServer()
}

func RegisterMerchantCenterLinkServiceServer(s grpc.ServiceRegistrar, srv MerchantCenterLinkServiceServer) {
	s.RegisterService(&MerchantCenterLinkService_ServiceDesc, srv)
}

func _MerchantCenterLinkService_ListMerchantCenterLinks_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListMerchantCenterLinksRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MerchantCenterLinkServiceServer).ListMerchantCenterLinks(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/google.ads.googleads.v10.services.MerchantCenterLinkService/ListMerchantCenterLinks",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MerchantCenterLinkServiceServer).ListMerchantCenterLinks(ctx, req.(*ListMerchantCenterLinksRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _MerchantCenterLinkService_GetMerchantCenterLink_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetMerchantCenterLinkRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MerchantCenterLinkServiceServer).GetMerchantCenterLink(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/google.ads.googleads.v10.services.MerchantCenterLinkService/GetMerchantCenterLink",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MerchantCenterLinkServiceServer).GetMerchantCenterLink(ctx, req.(*GetMerchantCenterLinkRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _MerchantCenterLinkService_MutateMerchantCenterLink_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MutateMerchantCenterLinkRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MerchantCenterLinkServiceServer).MutateMerchantCenterLink(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/google.ads.googleads.v10.services.MerchantCenterLinkService/MutateMerchantCenterLink",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MerchantCenterLinkServiceServer).MutateMerchantCenterLink(ctx, req.(*MutateMerchantCenterLinkRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// MerchantCenterLinkService_ServiceDesc is the grpc.ServiceDesc for MerchantCenterLinkService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var MerchantCenterLinkService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "google.ads.googleads.v10.services.MerchantCenterLinkService",
	HandlerType: (*MerchantCenterLinkServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "ListMerchantCenterLinks",
			Handler:    _MerchantCenterLinkService_ListMerchantCenterLinks_Handler,
		},
		{
			MethodName: "GetMerchantCenterLink",
			Handler:    _MerchantCenterLinkService_GetMerchantCenterLink_Handler,
		},
		{
			MethodName: "MutateMerchantCenterLink",
			Handler:    _MerchantCenterLinkService_MutateMerchantCenterLink_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "google/ads/googleads/v10/services/merchant_center_link_service.proto",
}
