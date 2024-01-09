// Copyright (c) 2022 Mica. All rights reserved. All software, including, without limitation, all source
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
// source: mica/discount/administration/v1/admin_service.proto

package administrationv1

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	pingv1 "github.com/1080network/golang/discount/proto/micashared/common/pingv1"
	v1 "github.com/1080network/golang/discount/proto/micashared/common/v1"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

const (
	DiscountAdministrationService_GenerateMTLSCertificate_FullMethodName                   = "/mica.discount.administration.v1.DiscountAdministrationService/GenerateMTLSCertificate"
	DiscountAdministrationService_UpdateMTLSCertificate_FullMethodName                     = "/mica.discount.administration.v1.DiscountAdministrationService/UpdateMTLSCertificate"
	DiscountAdministrationService_SearchMTLSCertificate_FullMethodName                     = "/mica.discount.administration.v1.DiscountAdministrationService/SearchMTLSCertificate"
	DiscountAdministrationService_GetMTLSCertificate_FullMethodName                        = "/mica.discount.administration.v1.DiscountAdministrationService/GetMTLSCertificate"
	DiscountAdministrationService_GetExternalClientSettings_FullMethodName                 = "/mica.discount.administration.v1.DiscountAdministrationService/GetExternalClientSettings"
	DiscountAdministrationService_UpdateExternalClientCallbackAddress_FullMethodName       = "/mica.discount.administration.v1.DiscountAdministrationService/UpdateExternalClientCallbackAddress"
	DiscountAdministrationService_GenerateExternalClientMTLSCertificate_FullMethodName     = "/mica.discount.administration.v1.DiscountAdministrationService/GenerateExternalClientMTLSCertificate"
	DiscountAdministrationService_UpdateExternalClientMTLSCertificate_FullMethodName       = "/mica.discount.administration.v1.DiscountAdministrationService/UpdateExternalClientMTLSCertificate"
	DiscountAdministrationService_SearchExternalClientMTLSCertificate_FullMethodName       = "/mica.discount.administration.v1.DiscountAdministrationService/SearchExternalClientMTLSCertificate"
	DiscountAdministrationService_CreateApiTokenConfiguration_FullMethodName               = "/mica.discount.administration.v1.DiscountAdministrationService/CreateApiTokenConfiguration"
	DiscountAdministrationService_GetApiTokenConfiguration_FullMethodName                  = "/mica.discount.administration.v1.DiscountAdministrationService/GetApiTokenConfiguration"
	DiscountAdministrationService_UpdateAPITokenAuthenticationConfiguration_FullMethodName = "/mica.discount.administration.v1.DiscountAdministrationService/UpdateAPITokenAuthenticationConfiguration"
	DiscountAdministrationService_PingExternal_FullMethodName                              = "/mica.discount.administration.v1.DiscountAdministrationService/PingExternal"
)

// DiscountAdministrationServiceClient is the client API for DiscountAdministrationService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type DiscountAdministrationServiceClient interface {
	// Certificate authentication to call from a provider onto mica
	GenerateMTLSCertificate(ctx context.Context, in *v1.GenerateMTLSCertificateRequest, opts ...grpc.CallOption) (*v1.GenerateMTLSCertificateResponse, error)
	// Update the certificate with a given serial number, only supports enable/disable for now
	UpdateMTLSCertificate(ctx context.Context, in *v1.UpdateMTLSCertificateRequest, opts ...grpc.CallOption) (*v1.UpdateMTLSCertificateResponse, error)
	// Search for certificates and return the ones that match the criteria provided
	SearchMTLSCertificate(ctx context.Context, in *v1.SearchMTLSCertificateRequest, opts ...grpc.CallOption) (*v1.SearchMTLSCertificateResponse, error)
	GetMTLSCertificate(ctx context.Context, in *v1.GetMTLSCertificateRequest, opts ...grpc.CallOption) (*v1.GetMTLSCertificateResponse, error)
	// External authentication mechanisms for Mica to call provider endpoints
	GetExternalClientSettings(ctx context.Context, in *v1.GetExternalClientSettingsRequest, opts ...grpc.CallOption) (*v1.GetExternalClientSettingsResponse, error)
	UpdateExternalClientCallbackAddress(ctx context.Context, in *v1.UpdateExternalClientCallBackAddressRequest, opts ...grpc.CallOption) (*v1.UpdateExternalClientCallBackAddressResponse, error)
	// Client certificates are used when mica needs to call out to a customers environment.
	GenerateExternalClientMTLSCertificate(ctx context.Context, in *v1.GenerateExternalClientMTLSCertificateRequest, opts ...grpc.CallOption) (*v1.GenerateExternalClientMTLSCertificateResponse, error)
	UpdateExternalClientMTLSCertificate(ctx context.Context, in *v1.UpdateExternalClientMTLSCertificateRequest, opts ...grpc.CallOption) (*v1.UpdateExternalClientMTLSCertificateResponse, error)
	SearchExternalClientMTLSCertificate(ctx context.Context, in *v1.SearchExternalClientMTLSCertificateRequest, opts ...grpc.CallOption) (*v1.SearchExternalClientMTLSCertificateResponse, error)
	// creates the token configuration
	CreateApiTokenConfiguration(ctx context.Context, in *v1.CreateApiTokenConfigurationRequest, opts ...grpc.CallOption) (*v1.CreateApiTokenConfigurationResponse, error)
	// if the authentication type is api token use these methods to properly setup the api token
	// retrieves the api token configuration including the token itself
	GetApiTokenConfiguration(ctx context.Context, in *v1.GetApiTokenConfigurationRequest, opts ...grpc.CallOption) (*v1.GetApiTokenConfigurationResponse, error)
	// Update the configuration of the token, the set fields will be modified
	UpdateAPITokenAuthenticationConfiguration(ctx context.Context, in *v1.UpdateApiTokenConfigurationRequest, opts ...grpc.CallOption) (*v1.UpdateApiTokenConfigurationResponse, error)
	// tests the external call to verify proper configuration and connectivity
	PingExternal(ctx context.Context, in *pingv1.PingRequest, opts ...grpc.CallOption) (*pingv1.PingResponse, error)
}

type discountAdministrationServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewDiscountAdministrationServiceClient(cc grpc.ClientConnInterface) DiscountAdministrationServiceClient {
	return &discountAdministrationServiceClient{cc}
}

func (c *discountAdministrationServiceClient) GenerateMTLSCertificate(ctx context.Context, in *v1.GenerateMTLSCertificateRequest, opts ...grpc.CallOption) (*v1.GenerateMTLSCertificateResponse, error) {
	out := new(v1.GenerateMTLSCertificateResponse)
	err := c.cc.Invoke(ctx, DiscountAdministrationService_GenerateMTLSCertificate_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *discountAdministrationServiceClient) UpdateMTLSCertificate(ctx context.Context, in *v1.UpdateMTLSCertificateRequest, opts ...grpc.CallOption) (*v1.UpdateMTLSCertificateResponse, error) {
	out := new(v1.UpdateMTLSCertificateResponse)
	err := c.cc.Invoke(ctx, DiscountAdministrationService_UpdateMTLSCertificate_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *discountAdministrationServiceClient) SearchMTLSCertificate(ctx context.Context, in *v1.SearchMTLSCertificateRequest, opts ...grpc.CallOption) (*v1.SearchMTLSCertificateResponse, error) {
	out := new(v1.SearchMTLSCertificateResponse)
	err := c.cc.Invoke(ctx, DiscountAdministrationService_SearchMTLSCertificate_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *discountAdministrationServiceClient) GetMTLSCertificate(ctx context.Context, in *v1.GetMTLSCertificateRequest, opts ...grpc.CallOption) (*v1.GetMTLSCertificateResponse, error) {
	out := new(v1.GetMTLSCertificateResponse)
	err := c.cc.Invoke(ctx, DiscountAdministrationService_GetMTLSCertificate_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *discountAdministrationServiceClient) GetExternalClientSettings(ctx context.Context, in *v1.GetExternalClientSettingsRequest, opts ...grpc.CallOption) (*v1.GetExternalClientSettingsResponse, error) {
	out := new(v1.GetExternalClientSettingsResponse)
	err := c.cc.Invoke(ctx, DiscountAdministrationService_GetExternalClientSettings_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *discountAdministrationServiceClient) UpdateExternalClientCallbackAddress(ctx context.Context, in *v1.UpdateExternalClientCallBackAddressRequest, opts ...grpc.CallOption) (*v1.UpdateExternalClientCallBackAddressResponse, error) {
	out := new(v1.UpdateExternalClientCallBackAddressResponse)
	err := c.cc.Invoke(ctx, DiscountAdministrationService_UpdateExternalClientCallbackAddress_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *discountAdministrationServiceClient) GenerateExternalClientMTLSCertificate(ctx context.Context, in *v1.GenerateExternalClientMTLSCertificateRequest, opts ...grpc.CallOption) (*v1.GenerateExternalClientMTLSCertificateResponse, error) {
	out := new(v1.GenerateExternalClientMTLSCertificateResponse)
	err := c.cc.Invoke(ctx, DiscountAdministrationService_GenerateExternalClientMTLSCertificate_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *discountAdministrationServiceClient) UpdateExternalClientMTLSCertificate(ctx context.Context, in *v1.UpdateExternalClientMTLSCertificateRequest, opts ...grpc.CallOption) (*v1.UpdateExternalClientMTLSCertificateResponse, error) {
	out := new(v1.UpdateExternalClientMTLSCertificateResponse)
	err := c.cc.Invoke(ctx, DiscountAdministrationService_UpdateExternalClientMTLSCertificate_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *discountAdministrationServiceClient) SearchExternalClientMTLSCertificate(ctx context.Context, in *v1.SearchExternalClientMTLSCertificateRequest, opts ...grpc.CallOption) (*v1.SearchExternalClientMTLSCertificateResponse, error) {
	out := new(v1.SearchExternalClientMTLSCertificateResponse)
	err := c.cc.Invoke(ctx, DiscountAdministrationService_SearchExternalClientMTLSCertificate_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *discountAdministrationServiceClient) CreateApiTokenConfiguration(ctx context.Context, in *v1.CreateApiTokenConfigurationRequest, opts ...grpc.CallOption) (*v1.CreateApiTokenConfigurationResponse, error) {
	out := new(v1.CreateApiTokenConfigurationResponse)
	err := c.cc.Invoke(ctx, DiscountAdministrationService_CreateApiTokenConfiguration_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *discountAdministrationServiceClient) GetApiTokenConfiguration(ctx context.Context, in *v1.GetApiTokenConfigurationRequest, opts ...grpc.CallOption) (*v1.GetApiTokenConfigurationResponse, error) {
	out := new(v1.GetApiTokenConfigurationResponse)
	err := c.cc.Invoke(ctx, DiscountAdministrationService_GetApiTokenConfiguration_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *discountAdministrationServiceClient) UpdateAPITokenAuthenticationConfiguration(ctx context.Context, in *v1.UpdateApiTokenConfigurationRequest, opts ...grpc.CallOption) (*v1.UpdateApiTokenConfigurationResponse, error) {
	out := new(v1.UpdateApiTokenConfigurationResponse)
	err := c.cc.Invoke(ctx, DiscountAdministrationService_UpdateAPITokenAuthenticationConfiguration_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *discountAdministrationServiceClient) PingExternal(ctx context.Context, in *pingv1.PingRequest, opts ...grpc.CallOption) (*pingv1.PingResponse, error) {
	out := new(pingv1.PingResponse)
	err := c.cc.Invoke(ctx, DiscountAdministrationService_PingExternal_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// DiscountAdministrationServiceServer is the server API for DiscountAdministrationService service.
// All implementations must embed UnimplementedDiscountAdministrationServiceServer
// for forward compatibility
type DiscountAdministrationServiceServer interface {
	// Certificate authentication to call from a provider onto mica
	GenerateMTLSCertificate(context.Context, *v1.GenerateMTLSCertificateRequest) (*v1.GenerateMTLSCertificateResponse, error)
	// Update the certificate with a given serial number, only supports enable/disable for now
	UpdateMTLSCertificate(context.Context, *v1.UpdateMTLSCertificateRequest) (*v1.UpdateMTLSCertificateResponse, error)
	// Search for certificates and return the ones that match the criteria provided
	SearchMTLSCertificate(context.Context, *v1.SearchMTLSCertificateRequest) (*v1.SearchMTLSCertificateResponse, error)
	GetMTLSCertificate(context.Context, *v1.GetMTLSCertificateRequest) (*v1.GetMTLSCertificateResponse, error)
	// External authentication mechanisms for Mica to call provider endpoints
	GetExternalClientSettings(context.Context, *v1.GetExternalClientSettingsRequest) (*v1.GetExternalClientSettingsResponse, error)
	UpdateExternalClientCallbackAddress(context.Context, *v1.UpdateExternalClientCallBackAddressRequest) (*v1.UpdateExternalClientCallBackAddressResponse, error)
	// Client certificates are used when mica needs to call out to a customers environment.
	GenerateExternalClientMTLSCertificate(context.Context, *v1.GenerateExternalClientMTLSCertificateRequest) (*v1.GenerateExternalClientMTLSCertificateResponse, error)
	UpdateExternalClientMTLSCertificate(context.Context, *v1.UpdateExternalClientMTLSCertificateRequest) (*v1.UpdateExternalClientMTLSCertificateResponse, error)
	SearchExternalClientMTLSCertificate(context.Context, *v1.SearchExternalClientMTLSCertificateRequest) (*v1.SearchExternalClientMTLSCertificateResponse, error)
	// creates the token configuration
	CreateApiTokenConfiguration(context.Context, *v1.CreateApiTokenConfigurationRequest) (*v1.CreateApiTokenConfigurationResponse, error)
	// if the authentication type is api token use these methods to properly setup the api token
	// retrieves the api token configuration including the token itself
	GetApiTokenConfiguration(context.Context, *v1.GetApiTokenConfigurationRequest) (*v1.GetApiTokenConfigurationResponse, error)
	// Update the configuration of the token, the set fields will be modified
	UpdateAPITokenAuthenticationConfiguration(context.Context, *v1.UpdateApiTokenConfigurationRequest) (*v1.UpdateApiTokenConfigurationResponse, error)
	// tests the external call to verify proper configuration and connectivity
	PingExternal(context.Context, *pingv1.PingRequest) (*pingv1.PingResponse, error)
	mustEmbedUnimplementedDiscountAdministrationServiceServer()
}

// UnimplementedDiscountAdministrationServiceServer must be embedded to have forward compatible implementations.
type UnimplementedDiscountAdministrationServiceServer struct {
}

func (UnimplementedDiscountAdministrationServiceServer) GenerateMTLSCertificate(context.Context, *v1.GenerateMTLSCertificateRequest) (*v1.GenerateMTLSCertificateResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GenerateMTLSCertificate not implemented")
}
func (UnimplementedDiscountAdministrationServiceServer) UpdateMTLSCertificate(context.Context, *v1.UpdateMTLSCertificateRequest) (*v1.UpdateMTLSCertificateResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateMTLSCertificate not implemented")
}
func (UnimplementedDiscountAdministrationServiceServer) SearchMTLSCertificate(context.Context, *v1.SearchMTLSCertificateRequest) (*v1.SearchMTLSCertificateResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SearchMTLSCertificate not implemented")
}
func (UnimplementedDiscountAdministrationServiceServer) GetMTLSCertificate(context.Context, *v1.GetMTLSCertificateRequest) (*v1.GetMTLSCertificateResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetMTLSCertificate not implemented")
}
func (UnimplementedDiscountAdministrationServiceServer) GetExternalClientSettings(context.Context, *v1.GetExternalClientSettingsRequest) (*v1.GetExternalClientSettingsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetExternalClientSettings not implemented")
}
func (UnimplementedDiscountAdministrationServiceServer) UpdateExternalClientCallbackAddress(context.Context, *v1.UpdateExternalClientCallBackAddressRequest) (*v1.UpdateExternalClientCallBackAddressResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateExternalClientCallbackAddress not implemented")
}
func (UnimplementedDiscountAdministrationServiceServer) GenerateExternalClientMTLSCertificate(context.Context, *v1.GenerateExternalClientMTLSCertificateRequest) (*v1.GenerateExternalClientMTLSCertificateResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GenerateExternalClientMTLSCertificate not implemented")
}
func (UnimplementedDiscountAdministrationServiceServer) UpdateExternalClientMTLSCertificate(context.Context, *v1.UpdateExternalClientMTLSCertificateRequest) (*v1.UpdateExternalClientMTLSCertificateResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateExternalClientMTLSCertificate not implemented")
}
func (UnimplementedDiscountAdministrationServiceServer) SearchExternalClientMTLSCertificate(context.Context, *v1.SearchExternalClientMTLSCertificateRequest) (*v1.SearchExternalClientMTLSCertificateResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SearchExternalClientMTLSCertificate not implemented")
}
func (UnimplementedDiscountAdministrationServiceServer) CreateApiTokenConfiguration(context.Context, *v1.CreateApiTokenConfigurationRequest) (*v1.CreateApiTokenConfigurationResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateApiTokenConfiguration not implemented")
}
func (UnimplementedDiscountAdministrationServiceServer) GetApiTokenConfiguration(context.Context, *v1.GetApiTokenConfigurationRequest) (*v1.GetApiTokenConfigurationResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetApiTokenConfiguration not implemented")
}
func (UnimplementedDiscountAdministrationServiceServer) UpdateAPITokenAuthenticationConfiguration(context.Context, *v1.UpdateApiTokenConfigurationRequest) (*v1.UpdateApiTokenConfigurationResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateAPITokenAuthenticationConfiguration not implemented")
}
func (UnimplementedDiscountAdministrationServiceServer) PingExternal(context.Context, *pingv1.PingRequest) (*pingv1.PingResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method PingExternal not implemented")
}
func (UnimplementedDiscountAdministrationServiceServer) mustEmbedUnimplementedDiscountAdministrationServiceServer() {
}

// UnsafeDiscountAdministrationServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to DiscountAdministrationServiceServer will
// result in compilation errors.
type UnsafeDiscountAdministrationServiceServer interface {
	mustEmbedUnimplementedDiscountAdministrationServiceServer()
}

func RegisterDiscountAdministrationServiceServer(s grpc.ServiceRegistrar, srv DiscountAdministrationServiceServer) {
	s.RegisterService(&DiscountAdministrationService_ServiceDesc, srv)
}

func _DiscountAdministrationService_GenerateMTLSCertificate_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(v1.GenerateMTLSCertificateRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DiscountAdministrationServiceServer).GenerateMTLSCertificate(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: DiscountAdministrationService_GenerateMTLSCertificate_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DiscountAdministrationServiceServer).GenerateMTLSCertificate(ctx, req.(*v1.GenerateMTLSCertificateRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _DiscountAdministrationService_UpdateMTLSCertificate_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(v1.UpdateMTLSCertificateRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DiscountAdministrationServiceServer).UpdateMTLSCertificate(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: DiscountAdministrationService_UpdateMTLSCertificate_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DiscountAdministrationServiceServer).UpdateMTLSCertificate(ctx, req.(*v1.UpdateMTLSCertificateRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _DiscountAdministrationService_SearchMTLSCertificate_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(v1.SearchMTLSCertificateRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DiscountAdministrationServiceServer).SearchMTLSCertificate(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: DiscountAdministrationService_SearchMTLSCertificate_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DiscountAdministrationServiceServer).SearchMTLSCertificate(ctx, req.(*v1.SearchMTLSCertificateRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _DiscountAdministrationService_GetMTLSCertificate_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(v1.GetMTLSCertificateRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DiscountAdministrationServiceServer).GetMTLSCertificate(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: DiscountAdministrationService_GetMTLSCertificate_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DiscountAdministrationServiceServer).GetMTLSCertificate(ctx, req.(*v1.GetMTLSCertificateRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _DiscountAdministrationService_GetExternalClientSettings_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(v1.GetExternalClientSettingsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DiscountAdministrationServiceServer).GetExternalClientSettings(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: DiscountAdministrationService_GetExternalClientSettings_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DiscountAdministrationServiceServer).GetExternalClientSettings(ctx, req.(*v1.GetExternalClientSettingsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _DiscountAdministrationService_UpdateExternalClientCallbackAddress_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(v1.UpdateExternalClientCallBackAddressRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DiscountAdministrationServiceServer).UpdateExternalClientCallbackAddress(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: DiscountAdministrationService_UpdateExternalClientCallbackAddress_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DiscountAdministrationServiceServer).UpdateExternalClientCallbackAddress(ctx, req.(*v1.UpdateExternalClientCallBackAddressRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _DiscountAdministrationService_GenerateExternalClientMTLSCertificate_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(v1.GenerateExternalClientMTLSCertificateRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DiscountAdministrationServiceServer).GenerateExternalClientMTLSCertificate(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: DiscountAdministrationService_GenerateExternalClientMTLSCertificate_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DiscountAdministrationServiceServer).GenerateExternalClientMTLSCertificate(ctx, req.(*v1.GenerateExternalClientMTLSCertificateRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _DiscountAdministrationService_UpdateExternalClientMTLSCertificate_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(v1.UpdateExternalClientMTLSCertificateRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DiscountAdministrationServiceServer).UpdateExternalClientMTLSCertificate(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: DiscountAdministrationService_UpdateExternalClientMTLSCertificate_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DiscountAdministrationServiceServer).UpdateExternalClientMTLSCertificate(ctx, req.(*v1.UpdateExternalClientMTLSCertificateRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _DiscountAdministrationService_SearchExternalClientMTLSCertificate_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(v1.SearchExternalClientMTLSCertificateRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DiscountAdministrationServiceServer).SearchExternalClientMTLSCertificate(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: DiscountAdministrationService_SearchExternalClientMTLSCertificate_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DiscountAdministrationServiceServer).SearchExternalClientMTLSCertificate(ctx, req.(*v1.SearchExternalClientMTLSCertificateRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _DiscountAdministrationService_CreateApiTokenConfiguration_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(v1.CreateApiTokenConfigurationRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DiscountAdministrationServiceServer).CreateApiTokenConfiguration(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: DiscountAdministrationService_CreateApiTokenConfiguration_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DiscountAdministrationServiceServer).CreateApiTokenConfiguration(ctx, req.(*v1.CreateApiTokenConfigurationRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _DiscountAdministrationService_GetApiTokenConfiguration_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(v1.GetApiTokenConfigurationRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DiscountAdministrationServiceServer).GetApiTokenConfiguration(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: DiscountAdministrationService_GetApiTokenConfiguration_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DiscountAdministrationServiceServer).GetApiTokenConfiguration(ctx, req.(*v1.GetApiTokenConfigurationRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _DiscountAdministrationService_UpdateAPITokenAuthenticationConfiguration_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(v1.UpdateApiTokenConfigurationRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DiscountAdministrationServiceServer).UpdateAPITokenAuthenticationConfiguration(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: DiscountAdministrationService_UpdateAPITokenAuthenticationConfiguration_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DiscountAdministrationServiceServer).UpdateAPITokenAuthenticationConfiguration(ctx, req.(*v1.UpdateApiTokenConfigurationRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _DiscountAdministrationService_PingExternal_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(pingv1.PingRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DiscountAdministrationServiceServer).PingExternal(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: DiscountAdministrationService_PingExternal_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DiscountAdministrationServiceServer).PingExternal(ctx, req.(*pingv1.PingRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// DiscountAdministrationService_ServiceDesc is the grpc.ServiceDesc for DiscountAdministrationService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var DiscountAdministrationService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "mica.discount.administration.v1.DiscountAdministrationService",
	HandlerType: (*DiscountAdministrationServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GenerateMTLSCertificate",
			Handler:    _DiscountAdministrationService_GenerateMTLSCertificate_Handler,
		},
		{
			MethodName: "UpdateMTLSCertificate",
			Handler:    _DiscountAdministrationService_UpdateMTLSCertificate_Handler,
		},
		{
			MethodName: "SearchMTLSCertificate",
			Handler:    _DiscountAdministrationService_SearchMTLSCertificate_Handler,
		},
		{
			MethodName: "GetMTLSCertificate",
			Handler:    _DiscountAdministrationService_GetMTLSCertificate_Handler,
		},
		{
			MethodName: "GetExternalClientSettings",
			Handler:    _DiscountAdministrationService_GetExternalClientSettings_Handler,
		},
		{
			MethodName: "UpdateExternalClientCallbackAddress",
			Handler:    _DiscountAdministrationService_UpdateExternalClientCallbackAddress_Handler,
		},
		{
			MethodName: "GenerateExternalClientMTLSCertificate",
			Handler:    _DiscountAdministrationService_GenerateExternalClientMTLSCertificate_Handler,
		},
		{
			MethodName: "UpdateExternalClientMTLSCertificate",
			Handler:    _DiscountAdministrationService_UpdateExternalClientMTLSCertificate_Handler,
		},
		{
			MethodName: "SearchExternalClientMTLSCertificate",
			Handler:    _DiscountAdministrationService_SearchExternalClientMTLSCertificate_Handler,
		},
		{
			MethodName: "CreateApiTokenConfiguration",
			Handler:    _DiscountAdministrationService_CreateApiTokenConfiguration_Handler,
		},
		{
			MethodName: "GetApiTokenConfiguration",
			Handler:    _DiscountAdministrationService_GetApiTokenConfiguration_Handler,
		},
		{
			MethodName: "UpdateAPITokenAuthenticationConfiguration",
			Handler:    _DiscountAdministrationService_UpdateAPITokenAuthenticationConfiguration_Handler,
		},
		{
			MethodName: "PingExternal",
			Handler:    _DiscountAdministrationService_PingExternal_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "github.com/1080network/golang/discount/proto/mica/discount/administration/v1/admin_service.proto",
}
