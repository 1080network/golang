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
// - protoc             v5.28.3
// source: mica/partner/service/v1/partner_to_mica_wip_service.proto

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
	PartnerToMicaWIPService_CreateRecurringPayment_FullMethodName        = "/mica.partner.service.v1.PartnerToMicaWIPService/CreateRecurringPayment"
	PartnerToMicaWIPService_ActivateRecurringPayment_FullMethodName      = "/mica.partner.service.v1.PartnerToMicaWIPService/ActivateRecurringPayment"
	PartnerToMicaWIPService_CancelRecurringPayment_FullMethodName        = "/mica.partner.service.v1.PartnerToMicaWIPService/CancelRecurringPayment"
	PartnerToMicaWIPService_GetRecurringPayment_FullMethodName           = "/mica.partner.service.v1.PartnerToMicaWIPService/GetRecurringPayment"
	PartnerToMicaWIPService_SearchRecurringPayment_FullMethodName        = "/mica.partner.service.v1.PartnerToMicaWIPService/SearchRecurringPayment"
	PartnerToMicaWIPService_ProvisionRecurringPaymentUUEK_FullMethodName = "/mica.partner.service.v1.PartnerToMicaWIPService/ProvisionRecurringPaymentUUEK"
	PartnerToMicaWIPService_RecurringPaymentObtainValue_FullMethodName   = "/mica.partner.service.v1.PartnerToMicaWIPService/RecurringPaymentObtainValue"
	PartnerToMicaWIPService_RecurringPaymentReturnValue_FullMethodName   = "/mica.partner.service.v1.PartnerToMicaWIPService/RecurringPaymentReturnValue"
	PartnerToMicaWIPService_InitializeWidget_FullMethodName              = "/mica.partner.service.v1.PartnerToMicaWIPService/InitializeWidget"
	PartnerToMicaWIPService_ExchangeSessionKey_FullMethodName            = "/mica.partner.service.v1.PartnerToMicaWIPService/ExchangeSessionKey"
)

// PartnerToMicaWIPServiceClient is the client API for PartnerToMicaWIPService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type PartnerToMicaWIPServiceClient interface {
	// Create a recurring payment to be attached to a partner to service provider link
	CreateRecurringPayment(ctx context.Context, in *CreateRecurringPaymentRequest, opts ...grpc.CallOption) (*CreateRecurringPaymentResponse, error)
	// Activate a recurring payment so that it can be used to make payments
	ActivateRecurringPayment(ctx context.Context, in *ActivateRecurringPaymentRequest, opts ...grpc.CallOption) (*ActivateRecurringPaymentResponse, error)
	// The merchant or the user using the merchants software may wish to cancel a recurring payment
	CancelRecurringPayment(ctx context.Context, in *CancelRecurringPaymentRequest, opts ...grpc.CallOption) (*CancelRecurringPaymentResponse, error)
	// return a recurring payment using a unique identifier as a key
	GetRecurringPayment(ctx context.Context, in *GetRecurringPaymentRequest, opts ...grpc.CallOption) (*GetRecurringPaymentResponse, error)
	// Search for recurring payments by a link key
	SearchRecurringPayment(ctx context.Context, in *SearchRecurringPaymentRequest, opts ...grpc.CallOption) (*SearchRecurringPaymentResponse, error)
	// The user should call this method before calling ExecuteRecurringPaymentObtainValue below
	ProvisionRecurringPaymentUUEK(ctx context.Context, in *ProvisionRecurringPaymentUUEKRequest, opts ...grpc.CallOption) (*ProvisionRecurringPaymentUUEKResponse, error)
	// Note that an attempt to pass a UUEK not minted by the ProvisionRecurringPaymentUUEK method will result in a runtime error
	RecurringPaymentObtainValue(ctx context.Context, in *RecurringPaymentObtainValueRequest, opts ...grpc.CallOption) (*RecurringPaymentObtainValueResponse, error)
	// This method is used to return value to the user's instrument for a recurring payment
	RecurringPaymentReturnValue(ctx context.Context, in *RecurringPaymentReturnValueRequest, opts ...grpc.CallOption) (*RecurringPaymentReturnValueResponse, error)
	// Deprecated: Do not use.
	// Initialize the widget (deprecated)
	InitializeWidget(ctx context.Context, in *InitializeWidgetRequest, opts ...grpc.CallOption) (*InitializeWidgetResponse, error)
	// Deprecated: Do not use.
	// Exchange the session key for a UUEK (deprecated)
	ExchangeSessionKey(ctx context.Context, in *ExchangeSessionKeyRequest, opts ...grpc.CallOption) (*ExchangeSessionKeyResponse, error)
}

type partnerToMicaWIPServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewPartnerToMicaWIPServiceClient(cc grpc.ClientConnInterface) PartnerToMicaWIPServiceClient {
	return &partnerToMicaWIPServiceClient{cc}
}

func (c *partnerToMicaWIPServiceClient) CreateRecurringPayment(ctx context.Context, in *CreateRecurringPaymentRequest, opts ...grpc.CallOption) (*CreateRecurringPaymentResponse, error) {
	out := new(CreateRecurringPaymentResponse)
	err := c.cc.Invoke(ctx, PartnerToMicaWIPService_CreateRecurringPayment_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *partnerToMicaWIPServiceClient) ActivateRecurringPayment(ctx context.Context, in *ActivateRecurringPaymentRequest, opts ...grpc.CallOption) (*ActivateRecurringPaymentResponse, error) {
	out := new(ActivateRecurringPaymentResponse)
	err := c.cc.Invoke(ctx, PartnerToMicaWIPService_ActivateRecurringPayment_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *partnerToMicaWIPServiceClient) CancelRecurringPayment(ctx context.Context, in *CancelRecurringPaymentRequest, opts ...grpc.CallOption) (*CancelRecurringPaymentResponse, error) {
	out := new(CancelRecurringPaymentResponse)
	err := c.cc.Invoke(ctx, PartnerToMicaWIPService_CancelRecurringPayment_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *partnerToMicaWIPServiceClient) GetRecurringPayment(ctx context.Context, in *GetRecurringPaymentRequest, opts ...grpc.CallOption) (*GetRecurringPaymentResponse, error) {
	out := new(GetRecurringPaymentResponse)
	err := c.cc.Invoke(ctx, PartnerToMicaWIPService_GetRecurringPayment_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *partnerToMicaWIPServiceClient) SearchRecurringPayment(ctx context.Context, in *SearchRecurringPaymentRequest, opts ...grpc.CallOption) (*SearchRecurringPaymentResponse, error) {
	out := new(SearchRecurringPaymentResponse)
	err := c.cc.Invoke(ctx, PartnerToMicaWIPService_SearchRecurringPayment_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *partnerToMicaWIPServiceClient) ProvisionRecurringPaymentUUEK(ctx context.Context, in *ProvisionRecurringPaymentUUEKRequest, opts ...grpc.CallOption) (*ProvisionRecurringPaymentUUEKResponse, error) {
	out := new(ProvisionRecurringPaymentUUEKResponse)
	err := c.cc.Invoke(ctx, PartnerToMicaWIPService_ProvisionRecurringPaymentUUEK_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *partnerToMicaWIPServiceClient) RecurringPaymentObtainValue(ctx context.Context, in *RecurringPaymentObtainValueRequest, opts ...grpc.CallOption) (*RecurringPaymentObtainValueResponse, error) {
	out := new(RecurringPaymentObtainValueResponse)
	err := c.cc.Invoke(ctx, PartnerToMicaWIPService_RecurringPaymentObtainValue_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *partnerToMicaWIPServiceClient) RecurringPaymentReturnValue(ctx context.Context, in *RecurringPaymentReturnValueRequest, opts ...grpc.CallOption) (*RecurringPaymentReturnValueResponse, error) {
	out := new(RecurringPaymentReturnValueResponse)
	err := c.cc.Invoke(ctx, PartnerToMicaWIPService_RecurringPaymentReturnValue_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Deprecated: Do not use.
func (c *partnerToMicaWIPServiceClient) InitializeWidget(ctx context.Context, in *InitializeWidgetRequest, opts ...grpc.CallOption) (*InitializeWidgetResponse, error) {
	out := new(InitializeWidgetResponse)
	err := c.cc.Invoke(ctx, PartnerToMicaWIPService_InitializeWidget_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Deprecated: Do not use.
func (c *partnerToMicaWIPServiceClient) ExchangeSessionKey(ctx context.Context, in *ExchangeSessionKeyRequest, opts ...grpc.CallOption) (*ExchangeSessionKeyResponse, error) {
	out := new(ExchangeSessionKeyResponse)
	err := c.cc.Invoke(ctx, PartnerToMicaWIPService_ExchangeSessionKey_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// PartnerToMicaWIPServiceServer is the server API for PartnerToMicaWIPService service.
// All implementations must embed UnimplementedPartnerToMicaWIPServiceServer
// for forward compatibility
type PartnerToMicaWIPServiceServer interface {
	// Create a recurring payment to be attached to a partner to service provider link
	CreateRecurringPayment(context.Context, *CreateRecurringPaymentRequest) (*CreateRecurringPaymentResponse, error)
	// Activate a recurring payment so that it can be used to make payments
	ActivateRecurringPayment(context.Context, *ActivateRecurringPaymentRequest) (*ActivateRecurringPaymentResponse, error)
	// The merchant or the user using the merchants software may wish to cancel a recurring payment
	CancelRecurringPayment(context.Context, *CancelRecurringPaymentRequest) (*CancelRecurringPaymentResponse, error)
	// return a recurring payment using a unique identifier as a key
	GetRecurringPayment(context.Context, *GetRecurringPaymentRequest) (*GetRecurringPaymentResponse, error)
	// Search for recurring payments by a link key
	SearchRecurringPayment(context.Context, *SearchRecurringPaymentRequest) (*SearchRecurringPaymentResponse, error)
	// The user should call this method before calling ExecuteRecurringPaymentObtainValue below
	ProvisionRecurringPaymentUUEK(context.Context, *ProvisionRecurringPaymentUUEKRequest) (*ProvisionRecurringPaymentUUEKResponse, error)
	// Note that an attempt to pass a UUEK not minted by the ProvisionRecurringPaymentUUEK method will result in a runtime error
	RecurringPaymentObtainValue(context.Context, *RecurringPaymentObtainValueRequest) (*RecurringPaymentObtainValueResponse, error)
	// This method is used to return value to the user's instrument for a recurring payment
	RecurringPaymentReturnValue(context.Context, *RecurringPaymentReturnValueRequest) (*RecurringPaymentReturnValueResponse, error)
	// Deprecated: Do not use.
	// Initialize the widget (deprecated)
	InitializeWidget(context.Context, *InitializeWidgetRequest) (*InitializeWidgetResponse, error)
	// Deprecated: Do not use.
	// Exchange the session key for a UUEK (deprecated)
	ExchangeSessionKey(context.Context, *ExchangeSessionKeyRequest) (*ExchangeSessionKeyResponse, error)
	mustEmbedUnimplementedPartnerToMicaWIPServiceServer()
}

// UnimplementedPartnerToMicaWIPServiceServer must be embedded to have forward compatible implementations.
type UnimplementedPartnerToMicaWIPServiceServer struct {
}

func (UnimplementedPartnerToMicaWIPServiceServer) CreateRecurringPayment(context.Context, *CreateRecurringPaymentRequest) (*CreateRecurringPaymentResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateRecurringPayment not implemented")
}
func (UnimplementedPartnerToMicaWIPServiceServer) ActivateRecurringPayment(context.Context, *ActivateRecurringPaymentRequest) (*ActivateRecurringPaymentResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ActivateRecurringPayment not implemented")
}
func (UnimplementedPartnerToMicaWIPServiceServer) CancelRecurringPayment(context.Context, *CancelRecurringPaymentRequest) (*CancelRecurringPaymentResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CancelRecurringPayment not implemented")
}
func (UnimplementedPartnerToMicaWIPServiceServer) GetRecurringPayment(context.Context, *GetRecurringPaymentRequest) (*GetRecurringPaymentResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetRecurringPayment not implemented")
}
func (UnimplementedPartnerToMicaWIPServiceServer) SearchRecurringPayment(context.Context, *SearchRecurringPaymentRequest) (*SearchRecurringPaymentResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SearchRecurringPayment not implemented")
}
func (UnimplementedPartnerToMicaWIPServiceServer) ProvisionRecurringPaymentUUEK(context.Context, *ProvisionRecurringPaymentUUEKRequest) (*ProvisionRecurringPaymentUUEKResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ProvisionRecurringPaymentUUEK not implemented")
}
func (UnimplementedPartnerToMicaWIPServiceServer) RecurringPaymentObtainValue(context.Context, *RecurringPaymentObtainValueRequest) (*RecurringPaymentObtainValueResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method RecurringPaymentObtainValue not implemented")
}
func (UnimplementedPartnerToMicaWIPServiceServer) RecurringPaymentReturnValue(context.Context, *RecurringPaymentReturnValueRequest) (*RecurringPaymentReturnValueResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method RecurringPaymentReturnValue not implemented")
}
func (UnimplementedPartnerToMicaWIPServiceServer) InitializeWidget(context.Context, *InitializeWidgetRequest) (*InitializeWidgetResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method InitializeWidget not implemented")
}
func (UnimplementedPartnerToMicaWIPServiceServer) ExchangeSessionKey(context.Context, *ExchangeSessionKeyRequest) (*ExchangeSessionKeyResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ExchangeSessionKey not implemented")
}
func (UnimplementedPartnerToMicaWIPServiceServer) mustEmbedUnimplementedPartnerToMicaWIPServiceServer() {
}

// UnsafePartnerToMicaWIPServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to PartnerToMicaWIPServiceServer will
// result in compilation errors.
type UnsafePartnerToMicaWIPServiceServer interface {
	mustEmbedUnimplementedPartnerToMicaWIPServiceServer()
}

func RegisterPartnerToMicaWIPServiceServer(s grpc.ServiceRegistrar, srv PartnerToMicaWIPServiceServer) {
	s.RegisterService(&PartnerToMicaWIPService_ServiceDesc, srv)
}

func _PartnerToMicaWIPService_CreateRecurringPayment_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateRecurringPaymentRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PartnerToMicaWIPServiceServer).CreateRecurringPayment(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: PartnerToMicaWIPService_CreateRecurringPayment_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PartnerToMicaWIPServiceServer).CreateRecurringPayment(ctx, req.(*CreateRecurringPaymentRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _PartnerToMicaWIPService_ActivateRecurringPayment_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ActivateRecurringPaymentRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PartnerToMicaWIPServiceServer).ActivateRecurringPayment(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: PartnerToMicaWIPService_ActivateRecurringPayment_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PartnerToMicaWIPServiceServer).ActivateRecurringPayment(ctx, req.(*ActivateRecurringPaymentRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _PartnerToMicaWIPService_CancelRecurringPayment_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CancelRecurringPaymentRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PartnerToMicaWIPServiceServer).CancelRecurringPayment(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: PartnerToMicaWIPService_CancelRecurringPayment_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PartnerToMicaWIPServiceServer).CancelRecurringPayment(ctx, req.(*CancelRecurringPaymentRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _PartnerToMicaWIPService_GetRecurringPayment_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetRecurringPaymentRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PartnerToMicaWIPServiceServer).GetRecurringPayment(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: PartnerToMicaWIPService_GetRecurringPayment_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PartnerToMicaWIPServiceServer).GetRecurringPayment(ctx, req.(*GetRecurringPaymentRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _PartnerToMicaWIPService_SearchRecurringPayment_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SearchRecurringPaymentRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PartnerToMicaWIPServiceServer).SearchRecurringPayment(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: PartnerToMicaWIPService_SearchRecurringPayment_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PartnerToMicaWIPServiceServer).SearchRecurringPayment(ctx, req.(*SearchRecurringPaymentRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _PartnerToMicaWIPService_ProvisionRecurringPaymentUUEK_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ProvisionRecurringPaymentUUEKRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PartnerToMicaWIPServiceServer).ProvisionRecurringPaymentUUEK(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: PartnerToMicaWIPService_ProvisionRecurringPaymentUUEK_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PartnerToMicaWIPServiceServer).ProvisionRecurringPaymentUUEK(ctx, req.(*ProvisionRecurringPaymentUUEKRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _PartnerToMicaWIPService_RecurringPaymentObtainValue_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RecurringPaymentObtainValueRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PartnerToMicaWIPServiceServer).RecurringPaymentObtainValue(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: PartnerToMicaWIPService_RecurringPaymentObtainValue_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PartnerToMicaWIPServiceServer).RecurringPaymentObtainValue(ctx, req.(*RecurringPaymentObtainValueRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _PartnerToMicaWIPService_RecurringPaymentReturnValue_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RecurringPaymentReturnValueRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PartnerToMicaWIPServiceServer).RecurringPaymentReturnValue(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: PartnerToMicaWIPService_RecurringPaymentReturnValue_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PartnerToMicaWIPServiceServer).RecurringPaymentReturnValue(ctx, req.(*RecurringPaymentReturnValueRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _PartnerToMicaWIPService_InitializeWidget_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(InitializeWidgetRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PartnerToMicaWIPServiceServer).InitializeWidget(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: PartnerToMicaWIPService_InitializeWidget_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PartnerToMicaWIPServiceServer).InitializeWidget(ctx, req.(*InitializeWidgetRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _PartnerToMicaWIPService_ExchangeSessionKey_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ExchangeSessionKeyRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PartnerToMicaWIPServiceServer).ExchangeSessionKey(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: PartnerToMicaWIPService_ExchangeSessionKey_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PartnerToMicaWIPServiceServer).ExchangeSessionKey(ctx, req.(*ExchangeSessionKeyRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// PartnerToMicaWIPService_ServiceDesc is the grpc.ServiceDesc for PartnerToMicaWIPService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var PartnerToMicaWIPService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "mica.partner.service.v1.PartnerToMicaWIPService",
	HandlerType: (*PartnerToMicaWIPServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateRecurringPayment",
			Handler:    _PartnerToMicaWIPService_CreateRecurringPayment_Handler,
		},
		{
			MethodName: "ActivateRecurringPayment",
			Handler:    _PartnerToMicaWIPService_ActivateRecurringPayment_Handler,
		},
		{
			MethodName: "CancelRecurringPayment",
			Handler:    _PartnerToMicaWIPService_CancelRecurringPayment_Handler,
		},
		{
			MethodName: "GetRecurringPayment",
			Handler:    _PartnerToMicaWIPService_GetRecurringPayment_Handler,
		},
		{
			MethodName: "SearchRecurringPayment",
			Handler:    _PartnerToMicaWIPService_SearchRecurringPayment_Handler,
		},
		{
			MethodName: "ProvisionRecurringPaymentUUEK",
			Handler:    _PartnerToMicaWIPService_ProvisionRecurringPaymentUUEK_Handler,
		},
		{
			MethodName: "RecurringPaymentObtainValue",
			Handler:    _PartnerToMicaWIPService_RecurringPaymentObtainValue_Handler,
		},
		{
			MethodName: "RecurringPaymentReturnValue",
			Handler:    _PartnerToMicaWIPService_RecurringPaymentReturnValue_Handler,
		},
		{
			MethodName: "InitializeWidget",
			Handler:    _PartnerToMicaWIPService_InitializeWidget_Handler,
		},
		{
			MethodName: "ExchangeSessionKey",
			Handler:    _PartnerToMicaWIPService_ExchangeSessionKey_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "github.com/1080network/golang/fullsdk/proto/mica/partner/service/v1/partner_to_mica_wip_service.proto",
}
