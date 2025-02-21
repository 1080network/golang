// Copyright (c) 2024 Mica. All rights reserved. All software, including, without limitation, all source
// code and object code, is the intellectual property of Mica, Inc. and is protected by copyright, trademark and
// other intellectual property laws (collective "Mica Software"). You may not use, copy, reproduce, download, store,
// post, broadcast, transmit, modify, sell or make available to the public content from the Mica Software without a
// valid license or the prior written approval of Mica, Inc. Mica, its logos, slogans, taglines,
// products, feature names, and other trademarks are trademarks of Mica, Inc. and may not be used without
// permission.

// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v5.28.3
// source: mica/valueexchange/service/v1/value_exchange_to_mica.proto

package servicev1

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	pingv1 "github.com/1080network/golang/fullsdk/proto/micashared/common/pingv1"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

const (
	ValueProviderToMicaService_GetValueExchangeProvider_FullMethodName    = "/mica.valueexchange.service.v1.ValueProviderToMicaService/GetValueExchangeProvider"
	ValueProviderToMicaService_RegisterAccount_FullMethodName             = "/mica.valueexchange.service.v1.ValueProviderToMicaService/RegisterAccount"
	ValueProviderToMicaService_GetAccount_FullMethodName                  = "/mica.valueexchange.service.v1.ValueProviderToMicaService/GetAccount"
	ValueProviderToMicaService_UpdateBalance_FullMethodName               = "/mica.valueexchange.service.v1.ValueProviderToMicaService/UpdateBalance"
	ValueProviderToMicaService_RemoveAccount_FullMethodName               = "/mica.valueexchange.service.v1.ValueProviderToMicaService/RemoveAccount"
	ValueProviderToMicaService_ProvisionAccountLinkingCode_FullMethodName = "/mica.valueexchange.service.v1.ValueProviderToMicaService/ProvisionAccountLinkingCode"
	ValueProviderToMicaService_Ping_FullMethodName                        = "/mica.valueexchange.service.v1.ValueProviderToMicaService/Ping"
	ValueProviderToMicaService_PingWithRoundTrip_FullMethodName           = "/mica.valueexchange.service.v1.ValueProviderToMicaService/PingWithRoundTrip"
)

// ValueProviderToMicaServiceClient is the client API for ValueProviderToMicaService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type ValueProviderToMicaServiceClient interface {
	// Value Exchange Provider management
	GetValueExchangeProvider(ctx context.Context, in *GetValueExchangeProviderRequest, opts ...grpc.CallOption) (*GetValueExchangeProviderResponse, error)
	RegisterAccount(ctx context.Context, in *RegisterAccountRequest, opts ...grpc.CallOption) (*RegisterAccountResponse, error)
	GetAccount(ctx context.Context, in *GetAccountRequest, opts ...grpc.CallOption) (*GetAccountResponse, error)
	UpdateBalance(ctx context.Context, in *UpdateBalanceRequest, opts ...grpc.CallOption) (*UpdateBalanceResponse, error)
	RemoveAccount(ctx context.Context, in *RemoveAccountRequest, opts ...grpc.CallOption) (*RemoveAccountResponse, error)
	ProvisionAccountLinkingCode(ctx context.Context, in *ProvisionAccountLinkingCodeRequest, opts ...grpc.CallOption) (*ProvisionAccountLinkingCodeResponse, error)
	// An operation to ping the server to ensure it's up and running and that the connection is good.
	Ping(ctx context.Context, in *pingv1.PingRequest, opts ...grpc.CallOption) (*pingv1.PingResponse, error)
	// An operation that triggers a ping to the Mica servervice and that in turn triggers a ping to the discount provider service.
	// This is a useful call to ensure that the connectivity and security between the mica an external services is working.
	PingWithRoundTrip(ctx context.Context, in *pingv1.PingRequest, opts ...grpc.CallOption) (*pingv1.PingResponse, error)
}

type valueProviderToMicaServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewValueProviderToMicaServiceClient(cc grpc.ClientConnInterface) ValueProviderToMicaServiceClient {
	return &valueProviderToMicaServiceClient{cc}
}

func (c *valueProviderToMicaServiceClient) GetValueExchangeProvider(ctx context.Context, in *GetValueExchangeProviderRequest, opts ...grpc.CallOption) (*GetValueExchangeProviderResponse, error) {
	out := new(GetValueExchangeProviderResponse)
	err := c.cc.Invoke(ctx, ValueProviderToMicaService_GetValueExchangeProvider_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *valueProviderToMicaServiceClient) RegisterAccount(ctx context.Context, in *RegisterAccountRequest, opts ...grpc.CallOption) (*RegisterAccountResponse, error) {
	out := new(RegisterAccountResponse)
	err := c.cc.Invoke(ctx, ValueProviderToMicaService_RegisterAccount_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *valueProviderToMicaServiceClient) GetAccount(ctx context.Context, in *GetAccountRequest, opts ...grpc.CallOption) (*GetAccountResponse, error) {
	out := new(GetAccountResponse)
	err := c.cc.Invoke(ctx, ValueProviderToMicaService_GetAccount_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *valueProviderToMicaServiceClient) UpdateBalance(ctx context.Context, in *UpdateBalanceRequest, opts ...grpc.CallOption) (*UpdateBalanceResponse, error) {
	out := new(UpdateBalanceResponse)
	err := c.cc.Invoke(ctx, ValueProviderToMicaService_UpdateBalance_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *valueProviderToMicaServiceClient) RemoveAccount(ctx context.Context, in *RemoveAccountRequest, opts ...grpc.CallOption) (*RemoveAccountResponse, error) {
	out := new(RemoveAccountResponse)
	err := c.cc.Invoke(ctx, ValueProviderToMicaService_RemoveAccount_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *valueProviderToMicaServiceClient) ProvisionAccountLinkingCode(ctx context.Context, in *ProvisionAccountLinkingCodeRequest, opts ...grpc.CallOption) (*ProvisionAccountLinkingCodeResponse, error) {
	out := new(ProvisionAccountLinkingCodeResponse)
	err := c.cc.Invoke(ctx, ValueProviderToMicaService_ProvisionAccountLinkingCode_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *valueProviderToMicaServiceClient) Ping(ctx context.Context, in *pingv1.PingRequest, opts ...grpc.CallOption) (*pingv1.PingResponse, error) {
	out := new(pingv1.PingResponse)
	err := c.cc.Invoke(ctx, ValueProviderToMicaService_Ping_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *valueProviderToMicaServiceClient) PingWithRoundTrip(ctx context.Context, in *pingv1.PingRequest, opts ...grpc.CallOption) (*pingv1.PingResponse, error) {
	out := new(pingv1.PingResponse)
	err := c.cc.Invoke(ctx, ValueProviderToMicaService_PingWithRoundTrip_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ValueProviderToMicaServiceServer is the server API for ValueProviderToMicaService service.
// All implementations must embed UnimplementedValueProviderToMicaServiceServer
// for forward compatibility
type ValueProviderToMicaServiceServer interface {
	// Value Exchange Provider management
	GetValueExchangeProvider(context.Context, *GetValueExchangeProviderRequest) (*GetValueExchangeProviderResponse, error)
	RegisterAccount(context.Context, *RegisterAccountRequest) (*RegisterAccountResponse, error)
	GetAccount(context.Context, *GetAccountRequest) (*GetAccountResponse, error)
	UpdateBalance(context.Context, *UpdateBalanceRequest) (*UpdateBalanceResponse, error)
	RemoveAccount(context.Context, *RemoveAccountRequest) (*RemoveAccountResponse, error)
	ProvisionAccountLinkingCode(context.Context, *ProvisionAccountLinkingCodeRequest) (*ProvisionAccountLinkingCodeResponse, error)
	// An operation to ping the server to ensure it's up and running and that the connection is good.
	Ping(context.Context, *pingv1.PingRequest) (*pingv1.PingResponse, error)
	// An operation that triggers a ping to the Mica servervice and that in turn triggers a ping to the discount provider service.
	// This is a useful call to ensure that the connectivity and security between the mica an external services is working.
	PingWithRoundTrip(context.Context, *pingv1.PingRequest) (*pingv1.PingResponse, error)
	mustEmbedUnimplementedValueProviderToMicaServiceServer()
}

// UnimplementedValueProviderToMicaServiceServer must be embedded to have forward compatible implementations.
type UnimplementedValueProviderToMicaServiceServer struct {
}

func (UnimplementedValueProviderToMicaServiceServer) GetValueExchangeProvider(context.Context, *GetValueExchangeProviderRequest) (*GetValueExchangeProviderResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetValueExchangeProvider not implemented")
}
func (UnimplementedValueProviderToMicaServiceServer) RegisterAccount(context.Context, *RegisterAccountRequest) (*RegisterAccountResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method RegisterAccount not implemented")
}
func (UnimplementedValueProviderToMicaServiceServer) GetAccount(context.Context, *GetAccountRequest) (*GetAccountResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetAccount not implemented")
}
func (UnimplementedValueProviderToMicaServiceServer) UpdateBalance(context.Context, *UpdateBalanceRequest) (*UpdateBalanceResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateBalance not implemented")
}
func (UnimplementedValueProviderToMicaServiceServer) RemoveAccount(context.Context, *RemoveAccountRequest) (*RemoveAccountResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method RemoveAccount not implemented")
}
func (UnimplementedValueProviderToMicaServiceServer) ProvisionAccountLinkingCode(context.Context, *ProvisionAccountLinkingCodeRequest) (*ProvisionAccountLinkingCodeResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ProvisionAccountLinkingCode not implemented")
}
func (UnimplementedValueProviderToMicaServiceServer) Ping(context.Context, *pingv1.PingRequest) (*pingv1.PingResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Ping not implemented")
}
func (UnimplementedValueProviderToMicaServiceServer) PingWithRoundTrip(context.Context, *pingv1.PingRequest) (*pingv1.PingResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method PingWithRoundTrip not implemented")
}
func (UnimplementedValueProviderToMicaServiceServer) mustEmbedUnimplementedValueProviderToMicaServiceServer() {
}

// UnsafeValueProviderToMicaServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to ValueProviderToMicaServiceServer will
// result in compilation errors.
type UnsafeValueProviderToMicaServiceServer interface {
	mustEmbedUnimplementedValueProviderToMicaServiceServer()
}

func RegisterValueProviderToMicaServiceServer(s grpc.ServiceRegistrar, srv ValueProviderToMicaServiceServer) {
	s.RegisterService(&ValueProviderToMicaService_ServiceDesc, srv)
}

func _ValueProviderToMicaService_GetValueExchangeProvider_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetValueExchangeProviderRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ValueProviderToMicaServiceServer).GetValueExchangeProvider(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ValueProviderToMicaService_GetValueExchangeProvider_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ValueProviderToMicaServiceServer).GetValueExchangeProvider(ctx, req.(*GetValueExchangeProviderRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ValueProviderToMicaService_RegisterAccount_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RegisterAccountRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ValueProviderToMicaServiceServer).RegisterAccount(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ValueProviderToMicaService_RegisterAccount_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ValueProviderToMicaServiceServer).RegisterAccount(ctx, req.(*RegisterAccountRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ValueProviderToMicaService_GetAccount_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetAccountRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ValueProviderToMicaServiceServer).GetAccount(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ValueProviderToMicaService_GetAccount_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ValueProviderToMicaServiceServer).GetAccount(ctx, req.(*GetAccountRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ValueProviderToMicaService_UpdateBalance_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateBalanceRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ValueProviderToMicaServiceServer).UpdateBalance(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ValueProviderToMicaService_UpdateBalance_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ValueProviderToMicaServiceServer).UpdateBalance(ctx, req.(*UpdateBalanceRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ValueProviderToMicaService_RemoveAccount_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RemoveAccountRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ValueProviderToMicaServiceServer).RemoveAccount(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ValueProviderToMicaService_RemoveAccount_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ValueProviderToMicaServiceServer).RemoveAccount(ctx, req.(*RemoveAccountRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ValueProviderToMicaService_ProvisionAccountLinkingCode_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ProvisionAccountLinkingCodeRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ValueProviderToMicaServiceServer).ProvisionAccountLinkingCode(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ValueProviderToMicaService_ProvisionAccountLinkingCode_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ValueProviderToMicaServiceServer).ProvisionAccountLinkingCode(ctx, req.(*ProvisionAccountLinkingCodeRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ValueProviderToMicaService_Ping_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(pingv1.PingRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ValueProviderToMicaServiceServer).Ping(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ValueProviderToMicaService_Ping_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ValueProviderToMicaServiceServer).Ping(ctx, req.(*pingv1.PingRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ValueProviderToMicaService_PingWithRoundTrip_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(pingv1.PingRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ValueProviderToMicaServiceServer).PingWithRoundTrip(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ValueProviderToMicaService_PingWithRoundTrip_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ValueProviderToMicaServiceServer).PingWithRoundTrip(ctx, req.(*pingv1.PingRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// ValueProviderToMicaService_ServiceDesc is the grpc.ServiceDesc for ValueProviderToMicaService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var ValueProviderToMicaService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "mica.valueexchange.service.v1.ValueProviderToMicaService",
	HandlerType: (*ValueProviderToMicaServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetValueExchangeProvider",
			Handler:    _ValueProviderToMicaService_GetValueExchangeProvider_Handler,
		},
		{
			MethodName: "RegisterAccount",
			Handler:    _ValueProviderToMicaService_RegisterAccount_Handler,
		},
		{
			MethodName: "GetAccount",
			Handler:    _ValueProviderToMicaService_GetAccount_Handler,
		},
		{
			MethodName: "UpdateBalance",
			Handler:    _ValueProviderToMicaService_UpdateBalance_Handler,
		},
		{
			MethodName: "RemoveAccount",
			Handler:    _ValueProviderToMicaService_RemoveAccount_Handler,
		},
		{
			MethodName: "ProvisionAccountLinkingCode",
			Handler:    _ValueProviderToMicaService_ProvisionAccountLinkingCode_Handler,
		},
		{
			MethodName: "Ping",
			Handler:    _ValueProviderToMicaService_Ping_Handler,
		},
		{
			MethodName: "PingWithRoundTrip",
			Handler:    _ValueProviderToMicaService_PingWithRoundTrip_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "github.com/1080network/golang/fullsdk/proto/mica/valueexchange/service/v1/value_exchange_to_mica.proto",
}
