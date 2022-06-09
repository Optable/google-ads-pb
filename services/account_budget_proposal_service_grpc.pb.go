// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

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

// AccountBudgetProposalServiceClient is the client API for AccountBudgetProposalService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type AccountBudgetProposalServiceClient interface {
	// Creates, updates, or removes account budget proposals.  Operation statuses
	// are returned.
	//
	// List of thrown errors:
	//   [AccountBudgetProposalError]()
	//   [AuthenticationError]()
	//   [AuthorizationError]()
	//   [DatabaseError]()
	//   [DateError]()
	//   [FieldError]()
	//   [FieldMaskError]()
	//   [HeaderError]()
	//   [InternalError]()
	//   [MutateError]()
	//   [QuotaError]()
	//   [RequestError]()
	//   [StringLengthError]()
	MutateAccountBudgetProposal(ctx context.Context, in *MutateAccountBudgetProposalRequest, opts ...grpc.CallOption) (*MutateAccountBudgetProposalResponse, error)
}

type accountBudgetProposalServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewAccountBudgetProposalServiceClient(cc grpc.ClientConnInterface) AccountBudgetProposalServiceClient {
	return &accountBudgetProposalServiceClient{cc}
}

func (c *accountBudgetProposalServiceClient) MutateAccountBudgetProposal(ctx context.Context, in *MutateAccountBudgetProposalRequest, opts ...grpc.CallOption) (*MutateAccountBudgetProposalResponse, error) {
	out := new(MutateAccountBudgetProposalResponse)
	err := c.cc.Invoke(ctx, "/google.ads.googleads.v11.services.AccountBudgetProposalService/MutateAccountBudgetProposal", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// AccountBudgetProposalServiceServer is the server API for AccountBudgetProposalService service.
// All implementations must embed UnimplementedAccountBudgetProposalServiceServer
// for forward compatibility
type AccountBudgetProposalServiceServer interface {
	// Creates, updates, or removes account budget proposals.  Operation statuses
	// are returned.
	//
	// List of thrown errors:
	//   [AccountBudgetProposalError]()
	//   [AuthenticationError]()
	//   [AuthorizationError]()
	//   [DatabaseError]()
	//   [DateError]()
	//   [FieldError]()
	//   [FieldMaskError]()
	//   [HeaderError]()
	//   [InternalError]()
	//   [MutateError]()
	//   [QuotaError]()
	//   [RequestError]()
	//   [StringLengthError]()
	MutateAccountBudgetProposal(context.Context, *MutateAccountBudgetProposalRequest) (*MutateAccountBudgetProposalResponse, error)
	mustEmbedUnimplementedAccountBudgetProposalServiceServer()
}

// UnimplementedAccountBudgetProposalServiceServer must be embedded to have forward compatible implementations.
type UnimplementedAccountBudgetProposalServiceServer struct {
}

func (UnimplementedAccountBudgetProposalServiceServer) MutateAccountBudgetProposal(context.Context, *MutateAccountBudgetProposalRequest) (*MutateAccountBudgetProposalResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method MutateAccountBudgetProposal not implemented")
}
func (UnimplementedAccountBudgetProposalServiceServer) mustEmbedUnimplementedAccountBudgetProposalServiceServer() {
}

// UnsafeAccountBudgetProposalServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to AccountBudgetProposalServiceServer will
// result in compilation errors.
type UnsafeAccountBudgetProposalServiceServer interface {
	mustEmbedUnimplementedAccountBudgetProposalServiceServer()
}

func RegisterAccountBudgetProposalServiceServer(s grpc.ServiceRegistrar, srv AccountBudgetProposalServiceServer) {
	s.RegisterService(&AccountBudgetProposalService_ServiceDesc, srv)
}

func _AccountBudgetProposalService_MutateAccountBudgetProposal_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MutateAccountBudgetProposalRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AccountBudgetProposalServiceServer).MutateAccountBudgetProposal(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/google.ads.googleads.v11.services.AccountBudgetProposalService/MutateAccountBudgetProposal",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AccountBudgetProposalServiceServer).MutateAccountBudgetProposal(ctx, req.(*MutateAccountBudgetProposalRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// AccountBudgetProposalService_ServiceDesc is the grpc.ServiceDesc for AccountBudgetProposalService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var AccountBudgetProposalService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "google.ads.googleads.v11.services.AccountBudgetProposalService",
	HandlerType: (*AccountBudgetProposalServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "MutateAccountBudgetProposal",
			Handler:    _AccountBudgetProposalService_MutateAccountBudgetProposal_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "google/ads/googleads/v11/services/account_budget_proposal_service.proto",
}
