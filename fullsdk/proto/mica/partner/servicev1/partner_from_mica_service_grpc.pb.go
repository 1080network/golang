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
// source: mica/partner/service/v1/partner_from_mica_service.proto

package servicev1

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	userv1 "github.com/1080network/golang/fullsdk/proto/mica/partner/userv1"
	pingv1 "github.com/1080network/golang/fullsdk/proto/micashared/common/pingv1"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

const (
	PartnerFromMicaService_EnrollUserInstrument_FullMethodName = "/mica.partner.service.v1.PartnerFromMicaService/EnrollUserInstrument"
	PartnerFromMicaService_Ping_FullMethodName                 = "/mica.partner.service.v1.PartnerFromMicaService/Ping"
)

// PartnerFromMicaServiceClient is the client API for PartnerFromMicaService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type PartnerFromMicaServiceClient interface {
	// operation to support reverse onboarding. Instead of the Partner calling Mica, Mica will call the Partner with a
	// new enrollment. For example, a User starting at the mica user website could sign up for Netflix assuming that
	// Netflix supported mica and have implemented this API.
	EnrollUserInstrument(ctx context.Context, in *userv1.EnrollUserInstrumentRequest, opts ...grpc.CallOption) (*userv1.EnrollUserInstrumentResponse, error)
	// An operation to ping the server to ensure it's up and running and that the connection is good.
	Ping(ctx context.Context, in *pingv1.PingRequest, opts ...grpc.CallOption) (*pingv1.PingResponse, error)
}

type partnerFromMicaServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewPartnerFromMicaServiceClient(cc grpc.ClientConnInterface) PartnerFromMicaServiceClient {
	return &partnerFromMicaServiceClient{cc}
}

func (c *partnerFromMicaServiceClient) EnrollUserInstrument(ctx context.Context, in *userv1.EnrollUserInstrumentRequest, opts ...grpc.CallOption) (*userv1.EnrollUserInstrumentResponse, error) {
	out := new(userv1.EnrollUserInstrumentResponse)
	err := c.cc.Invoke(ctx, PartnerFromMicaService_EnrollUserInstrument_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *partnerFromMicaServiceClient) Ping(ctx context.Context, in *pingv1.PingRequest, opts ...grpc.CallOption) (*pingv1.PingResponse, error) {
	out := new(pingv1.PingResponse)
	err := c.cc.Invoke(ctx, PartnerFromMicaService_Ping_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// PartnerFromMicaServiceServer is the server API for PartnerFromMicaService service.
// All implementations must embed UnimplementedPartnerFromMicaServiceServer
// for forward compatibility
type PartnerFromMicaServiceServer interface {
	// operation to support reverse onboarding. Instead of the Partner calling Mica, Mica will call the Partner with a
	// new enrollment. For example, a User starting at the mica user website could sign up for Netflix assuming that
	// Netflix supported mica and have implemented this API.
	EnrollUserInstrument(context.Context, *userv1.EnrollUserInstrumentRequest) (*userv1.EnrollUserInstrumentResponse, error)
	// An operation to ping the server to ensure it's up and running and that the connection is good.
	Ping(context.Context, *pingv1.PingRequest) (*pingv1.PingResponse, error)
	mustEmbedUnimplementedPartnerFromMicaServiceServer()
}

// UnimplementedPartnerFromMicaServiceServer must be embedded to have forward compatible implementations.
type UnimplementedPartnerFromMicaServiceServer struct {
}

func (UnimplementedPartnerFromMicaServiceServer) EnrollUserInstrument(context.Context, *userv1.EnrollUserInstrumentRequest) (*userv1.EnrollUserInstrumentResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method EnrollUserInstrument not implemented")
}
func (UnimplementedPartnerFromMicaServiceServer) Ping(context.Context, *pingv1.PingRequest) (*pingv1.PingResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Ping not implemented")
}
func (UnimplementedPartnerFromMicaServiceServer) mustEmbedUnimplementedPartnerFromMicaServiceServer() {
}

// UnsafePartnerFromMicaServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to PartnerFromMicaServiceServer will
// result in compilation errors.
type UnsafePartnerFromMicaServiceServer interface {
	mustEmbedUnimplementedPartnerFromMicaServiceServer()
}

func RegisterPartnerFromMicaServiceServer(s grpc.ServiceRegistrar, srv PartnerFromMicaServiceServer) {
	s.RegisterService(&PartnerFromMicaService_ServiceDesc, srv)
}

func _PartnerFromMicaService_EnrollUserInstrument_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(userv1.EnrollUserInstrumentRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PartnerFromMicaServiceServer).EnrollUserInstrument(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: PartnerFromMicaService_EnrollUserInstrument_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PartnerFromMicaServiceServer).EnrollUserInstrument(ctx, req.(*userv1.EnrollUserInstrumentRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _PartnerFromMicaService_Ping_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(pingv1.PingRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PartnerFromMicaServiceServer).Ping(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: PartnerFromMicaService_Ping_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PartnerFromMicaServiceServer).Ping(ctx, req.(*pingv1.PingRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// PartnerFromMicaService_ServiceDesc is the grpc.ServiceDesc for PartnerFromMicaService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var PartnerFromMicaService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "mica.partner.service.v1.PartnerFromMicaService",
	HandlerType: (*PartnerFromMicaServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "EnrollUserInstrument",
			Handler:    _PartnerFromMicaService_EnrollUserInstrument_Handler,
		},
		{
			MethodName: "Ping",
			Handler:    _PartnerFromMicaService_Ping_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "github.com/1080network/golang/fullsdk/proto/mica/partner/service/v1/partner_from_mica_service.proto",
}