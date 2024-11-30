// Copyright (c) 2022-2024 Mica. All rights reserved. All software, including, without limitation, all source
// code and object code, is the intellectual property of Mica, Inc. and is protected by copyright, trademark and
// other intellectual property laws (collective "Mica Software"). You may not use, copy, reproduce, download, store,
// post, broadcast, transmit, modify, sell or make available to the public content from the Mica Software without a
// valid license or the prior written approval of Mica, Inc. Mica, its logos, slogans, taglines,
// products, feature names, and other trademarks are trademarks of Mica, Inc. and may not be used without
// permission.

// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v4.24.4
// source: mica/serviceprovider/service/v1/service_provider_to_mica_wip_service.proto

package servicev1

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
	ServiceProviderToMicaWIPService_SetVisibilityStatus_FullMethodName            = "/mica.serviceprovider.service.v1.ServiceProviderToMicaWIPService/SetVisibilityStatus"
	ServiceProviderToMicaWIPService_DiscoverUser_FullMethodName                   = "/mica.serviceprovider.service.v1.ServiceProviderToMicaWIPService/DiscoverUser"
	ServiceProviderToMicaWIPService_GetServiceProviderUUEK_FullMethodName         = "/mica.serviceprovider.service.v1.ServiceProviderToMicaWIPService/GetServiceProviderUUEK"
	ServiceProviderToMicaWIPService_SetPIN_FullMethodName                         = "/mica.serviceprovider.service.v1.ServiceProviderToMicaWIPService/SetPIN"
	ServiceProviderToMicaWIPService_ValidatePIN_FullMethodName                    = "/mica.serviceprovider.service.v1.ServiceProviderToMicaWIPService/ValidatePIN"
	ServiceProviderToMicaWIPService_ResetPIN_FullMethodName                       = "/mica.serviceprovider.service.v1.ServiceProviderToMicaWIPService/ResetPIN"
	ServiceProviderToMicaWIPService_RemovePIN_FullMethodName                      = "/mica.serviceprovider.service.v1.ServiceProviderToMicaWIPService/RemovePIN"
	ServiceProviderToMicaWIPService_GetLinkedAccountBalance_FullMethodName        = "/mica.serviceprovider.service.v1.ServiceProviderToMicaWIPService/GetLinkedAccountBalance"
	ServiceProviderToMicaWIPService_UpdateAccountEnableForPurchase_FullMethodName = "/mica.serviceprovider.service.v1.ServiceProviderToMicaWIPService/UpdateAccountEnableForPurchase"
)

// ServiceProviderToMicaWIPServiceClient is the client API for ServiceProviderToMicaWIPService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type ServiceProviderToMicaWIPServiceClient interface {
	SetVisibilityStatus(ctx context.Context, in *SetVisibilityStatusRequest, opts ...grpc.CallOption) (*SetVisibilityStatusResponse, error)
	DiscoverUser(ctx context.Context, in *DiscoverUserRequest, opts ...grpc.CallOption) (*DiscoverUserResponse, error)
	GetServiceProviderUUEK(ctx context.Context, in *GetServiceProviderUUEKRequest, opts ...grpc.CallOption) (*GetServiceProviderUUEKResponse, error)
	SetPIN(ctx context.Context, in *SetPINRequest, opts ...grpc.CallOption) (*SetPINResponse, error)
	ValidatePIN(ctx context.Context, in *ValidatePINRequest, opts ...grpc.CallOption) (*ValidatePINResponse, error)
	ResetPIN(ctx context.Context, in *ResetPINRequest, opts ...grpc.CallOption) (*ResetPINResponse, error)
	RemovePIN(ctx context.Context, in *RemovePINRequest, opts ...grpc.CallOption) (*RemovePINResponse, error)
	GetLinkedAccountBalance(ctx context.Context, in *GetLinkedAccountBalanceRequest, opts ...grpc.CallOption) (*GetLinkedAccountBalanceResponse, error)
	UpdateAccountEnableForPurchase(ctx context.Context, in *UpdateAccountEnableForPurchaseRequest, opts ...grpc.CallOption) (*UpdateAccountEnableForPurchaseResponse, error)
}

type serviceProviderToMicaWIPServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewServiceProviderToMicaWIPServiceClient(cc grpc.ClientConnInterface) ServiceProviderToMicaWIPServiceClient {
	return &serviceProviderToMicaWIPServiceClient{cc}
}

func (c *serviceProviderToMicaWIPServiceClient) SetVisibilityStatus(ctx context.Context, in *SetVisibilityStatusRequest, opts ...grpc.CallOption) (*SetVisibilityStatusResponse, error) {
	out := new(SetVisibilityStatusResponse)
	err := c.cc.Invoke(ctx, ServiceProviderToMicaWIPService_SetVisibilityStatus_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *serviceProviderToMicaWIPServiceClient) DiscoverUser(ctx context.Context, in *DiscoverUserRequest, opts ...grpc.CallOption) (*DiscoverUserResponse, error) {
	out := new(DiscoverUserResponse)
	err := c.cc.Invoke(ctx, ServiceProviderToMicaWIPService_DiscoverUser_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *serviceProviderToMicaWIPServiceClient) GetServiceProviderUUEK(ctx context.Context, in *GetServiceProviderUUEKRequest, opts ...grpc.CallOption) (*GetServiceProviderUUEKResponse, error) {
	out := new(GetServiceProviderUUEKResponse)
	err := c.cc.Invoke(ctx, ServiceProviderToMicaWIPService_GetServiceProviderUUEK_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *serviceProviderToMicaWIPServiceClient) SetPIN(ctx context.Context, in *SetPINRequest, opts ...grpc.CallOption) (*SetPINResponse, error) {
	out := new(SetPINResponse)
	err := c.cc.Invoke(ctx, ServiceProviderToMicaWIPService_SetPIN_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *serviceProviderToMicaWIPServiceClient) ValidatePIN(ctx context.Context, in *ValidatePINRequest, opts ...grpc.CallOption) (*ValidatePINResponse, error) {
	out := new(ValidatePINResponse)
	err := c.cc.Invoke(ctx, ServiceProviderToMicaWIPService_ValidatePIN_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *serviceProviderToMicaWIPServiceClient) ResetPIN(ctx context.Context, in *ResetPINRequest, opts ...grpc.CallOption) (*ResetPINResponse, error) {
	out := new(ResetPINResponse)
	err := c.cc.Invoke(ctx, ServiceProviderToMicaWIPService_ResetPIN_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *serviceProviderToMicaWIPServiceClient) RemovePIN(ctx context.Context, in *RemovePINRequest, opts ...grpc.CallOption) (*RemovePINResponse, error) {
	out := new(RemovePINResponse)
	err := c.cc.Invoke(ctx, ServiceProviderToMicaWIPService_RemovePIN_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *serviceProviderToMicaWIPServiceClient) GetLinkedAccountBalance(ctx context.Context, in *GetLinkedAccountBalanceRequest, opts ...grpc.CallOption) (*GetLinkedAccountBalanceResponse, error) {
	out := new(GetLinkedAccountBalanceResponse)
	err := c.cc.Invoke(ctx, ServiceProviderToMicaWIPService_GetLinkedAccountBalance_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *serviceProviderToMicaWIPServiceClient) UpdateAccountEnableForPurchase(ctx context.Context, in *UpdateAccountEnableForPurchaseRequest, opts ...grpc.CallOption) (*UpdateAccountEnableForPurchaseResponse, error) {
	out := new(UpdateAccountEnableForPurchaseResponse)
	err := c.cc.Invoke(ctx, ServiceProviderToMicaWIPService_UpdateAccountEnableForPurchase_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ServiceProviderToMicaWIPServiceServer is the server API for ServiceProviderToMicaWIPService service.
// All implementations must embed UnimplementedServiceProviderToMicaWIPServiceServer
// for forward compatibility
type ServiceProviderToMicaWIPServiceServer interface {
	SetVisibilityStatus(context.Context, *SetVisibilityStatusRequest) (*SetVisibilityStatusResponse, error)
	DiscoverUser(context.Context, *DiscoverUserRequest) (*DiscoverUserResponse, error)
	GetServiceProviderUUEK(context.Context, *GetServiceProviderUUEKRequest) (*GetServiceProviderUUEKResponse, error)
	SetPIN(context.Context, *SetPINRequest) (*SetPINResponse, error)
	ValidatePIN(context.Context, *ValidatePINRequest) (*ValidatePINResponse, error)
	ResetPIN(context.Context, *ResetPINRequest) (*ResetPINResponse, error)
	RemovePIN(context.Context, *RemovePINRequest) (*RemovePINResponse, error)
	GetLinkedAccountBalance(context.Context, *GetLinkedAccountBalanceRequest) (*GetLinkedAccountBalanceResponse, error)
	UpdateAccountEnableForPurchase(context.Context, *UpdateAccountEnableForPurchaseRequest) (*UpdateAccountEnableForPurchaseResponse, error)
	mustEmbedUnimplementedServiceProviderToMicaWIPServiceServer()
}

// UnimplementedServiceProviderToMicaWIPServiceServer must be embedded to have forward compatible implementations.
type UnimplementedServiceProviderToMicaWIPServiceServer struct {
}

func (UnimplementedServiceProviderToMicaWIPServiceServer) SetVisibilityStatus(context.Context, *SetVisibilityStatusRequest) (*SetVisibilityStatusResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SetVisibilityStatus not implemented")
}
func (UnimplementedServiceProviderToMicaWIPServiceServer) DiscoverUser(context.Context, *DiscoverUserRequest) (*DiscoverUserResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DiscoverUser not implemented")
}
func (UnimplementedServiceProviderToMicaWIPServiceServer) GetServiceProviderUUEK(context.Context, *GetServiceProviderUUEKRequest) (*GetServiceProviderUUEKResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetServiceProviderUUEK not implemented")
}
func (UnimplementedServiceProviderToMicaWIPServiceServer) SetPIN(context.Context, *SetPINRequest) (*SetPINResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SetPIN not implemented")
}
func (UnimplementedServiceProviderToMicaWIPServiceServer) ValidatePIN(context.Context, *ValidatePINRequest) (*ValidatePINResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ValidatePIN not implemented")
}
func (UnimplementedServiceProviderToMicaWIPServiceServer) ResetPIN(context.Context, *ResetPINRequest) (*ResetPINResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ResetPIN not implemented")
}
func (UnimplementedServiceProviderToMicaWIPServiceServer) RemovePIN(context.Context, *RemovePINRequest) (*RemovePINResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method RemovePIN not implemented")
}
func (UnimplementedServiceProviderToMicaWIPServiceServer) GetLinkedAccountBalance(context.Context, *GetLinkedAccountBalanceRequest) (*GetLinkedAccountBalanceResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetLinkedAccountBalance not implemented")
}
func (UnimplementedServiceProviderToMicaWIPServiceServer) UpdateAccountEnableForPurchase(context.Context, *UpdateAccountEnableForPurchaseRequest) (*UpdateAccountEnableForPurchaseResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateAccountEnableForPurchase not implemented")
}
func (UnimplementedServiceProviderToMicaWIPServiceServer) mustEmbedUnimplementedServiceProviderToMicaWIPServiceServer() {
}

// UnsafeServiceProviderToMicaWIPServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to ServiceProviderToMicaWIPServiceServer will
// result in compilation errors.
type UnsafeServiceProviderToMicaWIPServiceServer interface {
	mustEmbedUnimplementedServiceProviderToMicaWIPServiceServer()
}

func RegisterServiceProviderToMicaWIPServiceServer(s grpc.ServiceRegistrar, srv ServiceProviderToMicaWIPServiceServer) {
	s.RegisterService(&ServiceProviderToMicaWIPService_ServiceDesc, srv)
}

func _ServiceProviderToMicaWIPService_SetVisibilityStatus_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SetVisibilityStatusRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ServiceProviderToMicaWIPServiceServer).SetVisibilityStatus(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ServiceProviderToMicaWIPService_SetVisibilityStatus_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ServiceProviderToMicaWIPServiceServer).SetVisibilityStatus(ctx, req.(*SetVisibilityStatusRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ServiceProviderToMicaWIPService_DiscoverUser_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DiscoverUserRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ServiceProviderToMicaWIPServiceServer).DiscoverUser(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ServiceProviderToMicaWIPService_DiscoverUser_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ServiceProviderToMicaWIPServiceServer).DiscoverUser(ctx, req.(*DiscoverUserRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ServiceProviderToMicaWIPService_GetServiceProviderUUEK_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetServiceProviderUUEKRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ServiceProviderToMicaWIPServiceServer).GetServiceProviderUUEK(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ServiceProviderToMicaWIPService_GetServiceProviderUUEK_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ServiceProviderToMicaWIPServiceServer).GetServiceProviderUUEK(ctx, req.(*GetServiceProviderUUEKRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ServiceProviderToMicaWIPService_SetPIN_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SetPINRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ServiceProviderToMicaWIPServiceServer).SetPIN(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ServiceProviderToMicaWIPService_SetPIN_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ServiceProviderToMicaWIPServiceServer).SetPIN(ctx, req.(*SetPINRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ServiceProviderToMicaWIPService_ValidatePIN_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ValidatePINRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ServiceProviderToMicaWIPServiceServer).ValidatePIN(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ServiceProviderToMicaWIPService_ValidatePIN_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ServiceProviderToMicaWIPServiceServer).ValidatePIN(ctx, req.(*ValidatePINRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ServiceProviderToMicaWIPService_ResetPIN_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ResetPINRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ServiceProviderToMicaWIPServiceServer).ResetPIN(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ServiceProviderToMicaWIPService_ResetPIN_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ServiceProviderToMicaWIPServiceServer).ResetPIN(ctx, req.(*ResetPINRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ServiceProviderToMicaWIPService_RemovePIN_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RemovePINRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ServiceProviderToMicaWIPServiceServer).RemovePIN(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ServiceProviderToMicaWIPService_RemovePIN_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ServiceProviderToMicaWIPServiceServer).RemovePIN(ctx, req.(*RemovePINRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ServiceProviderToMicaWIPService_GetLinkedAccountBalance_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetLinkedAccountBalanceRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ServiceProviderToMicaWIPServiceServer).GetLinkedAccountBalance(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ServiceProviderToMicaWIPService_GetLinkedAccountBalance_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ServiceProviderToMicaWIPServiceServer).GetLinkedAccountBalance(ctx, req.(*GetLinkedAccountBalanceRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ServiceProviderToMicaWIPService_UpdateAccountEnableForPurchase_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateAccountEnableForPurchaseRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ServiceProviderToMicaWIPServiceServer).UpdateAccountEnableForPurchase(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ServiceProviderToMicaWIPService_UpdateAccountEnableForPurchase_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ServiceProviderToMicaWIPServiceServer).UpdateAccountEnableForPurchase(ctx, req.(*UpdateAccountEnableForPurchaseRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// ServiceProviderToMicaWIPService_ServiceDesc is the grpc.ServiceDesc for ServiceProviderToMicaWIPService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var ServiceProviderToMicaWIPService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "mica.serviceprovider.service.v1.ServiceProviderToMicaWIPService",
	HandlerType: (*ServiceProviderToMicaWIPServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "SetVisibilityStatus",
			Handler:    _ServiceProviderToMicaWIPService_SetVisibilityStatus_Handler,
		},
		{
			MethodName: "DiscoverUser",
			Handler:    _ServiceProviderToMicaWIPService_DiscoverUser_Handler,
		},
		{
			MethodName: "GetServiceProviderUUEK",
			Handler:    _ServiceProviderToMicaWIPService_GetServiceProviderUUEK_Handler,
		},
		{
			MethodName: "SetPIN",
			Handler:    _ServiceProviderToMicaWIPService_SetPIN_Handler,
		},
		{
			MethodName: "ValidatePIN",
			Handler:    _ServiceProviderToMicaWIPService_ValidatePIN_Handler,
		},
		{
			MethodName: "ResetPIN",
			Handler:    _ServiceProviderToMicaWIPService_ResetPIN_Handler,
		},
		{
			MethodName: "RemovePIN",
			Handler:    _ServiceProviderToMicaWIPService_RemovePIN_Handler,
		},
		{
			MethodName: "GetLinkedAccountBalance",
			Handler:    _ServiceProviderToMicaWIPService_GetLinkedAccountBalance_Handler,
		},
		{
			MethodName: "UpdateAccountEnableForPurchase",
			Handler:    _ServiceProviderToMicaWIPService_UpdateAccountEnableForPurchase_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "github.com/1080network/golang/fullsdk/proto/mica/serviceprovider/service/v1/service_provider_to_mica_wip_service.proto",
}
