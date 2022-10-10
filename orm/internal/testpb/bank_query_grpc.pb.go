// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             (unknown)
// source: testpb/bank_query.proto

package testpb

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

// BankQueryClient is the client API for BankQuery service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type BankQueryClient interface {
	// Get queries the Balance table by its primary key.
	GetBalance(ctx context.Context, in *GetBalanceRequest, opts ...grpc.CallOption) (*GetBalanceResponse, error)
	// ListBalance queries the Balance table using prefix and range queries against defined indexes.
	ListBalance(ctx context.Context, in *ListBalanceRequest, opts ...grpc.CallOption) (*ListBalanceResponse, error)
	// Get queries the Supply table by its primary key.
	GetSupply(ctx context.Context, in *GetSupplyRequest, opts ...grpc.CallOption) (*GetSupplyResponse, error)
	// ListSupply queries the Supply table using prefix and range queries against defined indexes.
	ListSupply(ctx context.Context, in *ListSupplyRequest, opts ...grpc.CallOption) (*ListSupplyResponse, error)
}

type bankQueryClient struct {
	cc grpc.ClientConnInterface
}

func NewBankQueryClient(cc grpc.ClientConnInterface) BankQueryClient {
	return &bankQueryClient{cc}
}

func (c *bankQueryClient) GetBalance(ctx context.Context, in *GetBalanceRequest, opts ...grpc.CallOption) (*GetBalanceResponse, error) {
	out := new(GetBalanceResponse)
	err := c.cc.Invoke(ctx, "/testpb.BankQuery/GetBalance", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *bankQueryClient) ListBalance(ctx context.Context, in *ListBalanceRequest, opts ...grpc.CallOption) (*ListBalanceResponse, error) {
	out := new(ListBalanceResponse)
	err := c.cc.Invoke(ctx, "/testpb.BankQuery/ListBalance", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *bankQueryClient) GetSupply(ctx context.Context, in *GetSupplyRequest, opts ...grpc.CallOption) (*GetSupplyResponse, error) {
	out := new(GetSupplyResponse)
	err := c.cc.Invoke(ctx, "/testpb.BankQuery/GetSupply", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *bankQueryClient) ListSupply(ctx context.Context, in *ListSupplyRequest, opts ...grpc.CallOption) (*ListSupplyResponse, error) {
	out := new(ListSupplyResponse)
	err := c.cc.Invoke(ctx, "/testpb.BankQuery/ListSupply", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// BankQueryServer is the server API for BankQuery service.
// All implementations must embed UnimplementedBankQueryServer
// for forward compatibility
type BankQueryServer interface {
	// Get queries the Balance table by its primary key.
	GetBalance(context.Context, *GetBalanceRequest) (*GetBalanceResponse, error)
	// ListBalance queries the Balance table using prefix and range queries against defined indexes.
	ListBalance(context.Context, *ListBalanceRequest) (*ListBalanceResponse, error)
	// Get queries the Supply table by its primary key.
	GetSupply(context.Context, *GetSupplyRequest) (*GetSupplyResponse, error)
	// ListSupply queries the Supply table using prefix and range queries against defined indexes.
	ListSupply(context.Context, *ListSupplyRequest) (*ListSupplyResponse, error)
	mustEmbedUnimplementedBankQueryServer()
}

// UnimplementedBankQueryServer must be embedded to have forward compatible implementations.
type UnimplementedBankQueryServer struct {
}

func (UnimplementedBankQueryServer) GetBalance(context.Context, *GetBalanceRequest) (*GetBalanceResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetBalance not implemented")
}
func (UnimplementedBankQueryServer) ListBalance(context.Context, *ListBalanceRequest) (*ListBalanceResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListBalance not implemented")
}
func (UnimplementedBankQueryServer) GetSupply(context.Context, *GetSupplyRequest) (*GetSupplyResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetSupply not implemented")
}
func (UnimplementedBankQueryServer) ListSupply(context.Context, *ListSupplyRequest) (*ListSupplyResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListSupply not implemented")
}
func (UnimplementedBankQueryServer) mustEmbedUnimplementedBankQueryServer() {}

// UnsafeBankQueryServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to BankQueryServer will
// result in compilation errors.
type UnsafeBankQueryServer interface {
	mustEmbedUnimplementedBankQueryServer()
}

func RegisterBankQueryServer(s grpc.ServiceRegistrar, srv BankQueryServer) {
	s.RegisterService(&BankQuery_ServiceDesc, srv)
}

func _BankQuery_GetBalance_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetBalanceRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BankQueryServer).GetBalance(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/testpb.BankQuery/GetBalance",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BankQueryServer).GetBalance(ctx, req.(*GetBalanceRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _BankQuery_ListBalance_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListBalanceRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BankQueryServer).ListBalance(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/testpb.BankQuery/ListBalance",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BankQueryServer).ListBalance(ctx, req.(*ListBalanceRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _BankQuery_GetSupply_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetSupplyRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BankQueryServer).GetSupply(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/testpb.BankQuery/GetSupply",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BankQueryServer).GetSupply(ctx, req.(*GetSupplyRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _BankQuery_ListSupply_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListSupplyRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BankQueryServer).ListSupply(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/testpb.BankQuery/ListSupply",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BankQueryServer).ListSupply(ctx, req.(*ListSupplyRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// BankQuery_ServiceDesc is the grpc.ServiceDesc for BankQuery service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var BankQuery_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "testpb.BankQuery",
	HandlerType: (*BankQueryServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetBalance",
			Handler:    _BankQuery_GetBalance_Handler,
		},
		{
			MethodName: "ListBalance",
			Handler:    _BankQuery_ListBalance_Handler,
		},
		{
			MethodName: "GetSupply",
			Handler:    _BankQuery_GetSupply_Handler,
		},
		{
			MethodName: "ListSupply",
			Handler:    _BankQuery_ListSupply_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "testpb/bank_query.proto",
}