// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.21.9
// source: google/ads/googleads/v12/services/customer_conversion_goal_service.proto

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

// CustomerConversionGoalServiceClient is the client API for CustomerConversionGoalService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type CustomerConversionGoalServiceClient interface {
	// Creates, updates or removes customer conversion goals. Operation statuses
	// are returned.
	MutateCustomerConversionGoals(ctx context.Context, in *MutateCustomerConversionGoalsRequest, opts ...grpc.CallOption) (*MutateCustomerConversionGoalsResponse, error)
}

type customerConversionGoalServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewCustomerConversionGoalServiceClient(cc grpc.ClientConnInterface) CustomerConversionGoalServiceClient {
	return &customerConversionGoalServiceClient{cc}
}

func (c *customerConversionGoalServiceClient) MutateCustomerConversionGoals(ctx context.Context, in *MutateCustomerConversionGoalsRequest, opts ...grpc.CallOption) (*MutateCustomerConversionGoalsResponse, error) {
	out := new(MutateCustomerConversionGoalsResponse)
	err := c.cc.Invoke(ctx, "/google.ads.googleads.v12.services.CustomerConversionGoalService/MutateCustomerConversionGoals", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// CustomerConversionGoalServiceServer is the server API for CustomerConversionGoalService service.
// All implementations must embed UnimplementedCustomerConversionGoalServiceServer
// for forward compatibility
type CustomerConversionGoalServiceServer interface {
	// Creates, updates or removes customer conversion goals. Operation statuses
	// are returned.
	MutateCustomerConversionGoals(context.Context, *MutateCustomerConversionGoalsRequest) (*MutateCustomerConversionGoalsResponse, error)
	mustEmbedUnimplementedCustomerConversionGoalServiceServer()
}

// UnimplementedCustomerConversionGoalServiceServer must be embedded to have forward compatible implementations.
type UnimplementedCustomerConversionGoalServiceServer struct {
}

func (UnimplementedCustomerConversionGoalServiceServer) MutateCustomerConversionGoals(context.Context, *MutateCustomerConversionGoalsRequest) (*MutateCustomerConversionGoalsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method MutateCustomerConversionGoals not implemented")
}
func (UnimplementedCustomerConversionGoalServiceServer) mustEmbedUnimplementedCustomerConversionGoalServiceServer() {
}

// UnsafeCustomerConversionGoalServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to CustomerConversionGoalServiceServer will
// result in compilation errors.
type UnsafeCustomerConversionGoalServiceServer interface {
	mustEmbedUnimplementedCustomerConversionGoalServiceServer()
}

func RegisterCustomerConversionGoalServiceServer(s grpc.ServiceRegistrar, srv CustomerConversionGoalServiceServer) {
	s.RegisterService(&CustomerConversionGoalService_ServiceDesc, srv)
}

func _CustomerConversionGoalService_MutateCustomerConversionGoals_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MutateCustomerConversionGoalsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CustomerConversionGoalServiceServer).MutateCustomerConversionGoals(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/google.ads.googleads.v12.services.CustomerConversionGoalService/MutateCustomerConversionGoals",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CustomerConversionGoalServiceServer).MutateCustomerConversionGoals(ctx, req.(*MutateCustomerConversionGoalsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// CustomerConversionGoalService_ServiceDesc is the grpc.ServiceDesc for CustomerConversionGoalService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var CustomerConversionGoalService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "google.ads.googleads.v12.services.CustomerConversionGoalService",
	HandlerType: (*CustomerConversionGoalServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "MutateCustomerConversionGoals",
			Handler:    _CustomerConversionGoalService_MutateCustomerConversionGoals_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "google/ads/googleads/v12/services/customer_conversion_goal_service.proto",
}
