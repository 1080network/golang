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
// - protoc             v3.21.12
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
	ServiceProviderToMicaWIPService_SetVisibilityStatus_FullMethodName               = "/mica.serviceprovider.service.v1.ServiceProviderToMicaWIPService/SetVisibilityStatus"
	ServiceProviderToMicaWIPService_DiscoverUser_FullMethodName                      = "/mica.serviceprovider.service.v1.ServiceProviderToMicaWIPService/DiscoverUser"
	ServiceProviderToMicaWIPService_GetServiceProviderUUEK_FullMethodName            = "/mica.serviceprovider.service.v1.ServiceProviderToMicaWIPService/GetServiceProviderUUEK"
	ServiceProviderToMicaWIPService_ProvisionEnrollmentValidationCode_FullMethodName = "/mica.serviceprovider.service.v1.ServiceProviderToMicaWIPService/ProvisionEnrollmentValidationCode"
)

// ServiceProviderToMicaWIPServiceClient is the client API for ServiceProviderToMicaWIPService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type ServiceProviderToMicaWIPServiceClient interface {
	SetVisibilityStatus(ctx context.Context, in *SetVisibilityStatusRequest, opts ...grpc.CallOption) (*SetVisibilityStatusResponse, error)
	DiscoverUser(ctx context.Context, in *DiscoverUserRequest, opts ...grpc.CallOption) (*DiscoverUserResponse, error)
	GetServiceProviderUUEK(ctx context.Context, in *GetServiceProviderUUEKRequest, opts ...grpc.CallOption) (*GetServiceProviderUUEKResponse, error)
	// <editor-fold desc="Service Provider intiated enrollment authentication">
	ProvisionEnrollmentValidationCode(ctx context.Context, in *ProvisionEnrollmentValidationCodeRequest, opts ...grpc.CallOption) (*ProvisionEnrollmentValidationCodeResponse, error)
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

func (c *serviceProviderToMicaWIPServiceClient) ProvisionEnrollmentValidationCode(ctx context.Context, in *ProvisionEnrollmentValidationCodeRequest, opts ...grpc.CallOption) (*ProvisionEnrollmentValidationCodeResponse, error) {
	out := new(ProvisionEnrollmentValidationCodeResponse)
	err := c.cc.Invoke(ctx, ServiceProviderToMicaWIPService_ProvisionEnrollmentValidationCode_FullMethodName, in, out, opts...)
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
	// <editor-fold desc="Service Provider intiated enrollment authentication">
	ProvisionEnrollmentValidationCode(context.Context, *ProvisionEnrollmentValidationCodeRequest) (*ProvisionEnrollmentValidationCodeResponse, error)
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
func (UnimplementedServiceProviderToMicaWIPServiceServer) ProvisionEnrollmentValidationCode(context.Context, *ProvisionEnrollmentValidationCodeRequest) (*ProvisionEnrollmentValidationCodeResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ProvisionEnrollmentValidationCode not implemented")
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

func _ServiceProviderToMicaWIPService_ProvisionEnrollmentValidationCode_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ProvisionEnrollmentValidationCodeRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ServiceProviderToMicaWIPServiceServer).ProvisionEnrollmentValidationCode(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ServiceProviderToMicaWIPService_ProvisionEnrollmentValidationCode_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ServiceProviderToMicaWIPServiceServer).ProvisionEnrollmentValidationCode(ctx, req.(*ProvisionEnrollmentValidationCodeRequest))
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
			MethodName: "ProvisionEnrollmentValidationCode",
			Handler:    _ServiceProviderToMicaWIPService_ProvisionEnrollmentValidationCode_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "github.com/1080network/golang/serviceprovider/proto/mica/serviceprovider/service/v1/service_provider_to_mica_wip_service.proto",
}