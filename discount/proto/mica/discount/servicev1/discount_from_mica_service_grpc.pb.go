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
// source: mica/discount/service/v1/discount_from_mica_service.proto

package servicev1

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	discountv1 "github.com/1080network/golang/discount/proto/mica/discount/discountv1"
	pingv1 "github.com/1080network/golang/discount/proto/micashared/common/pingv1"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

const (
	DiscountFromMicaService_ApplyDiscountNotification_FullMethodName = "/mica.discount.service.v1.DiscountFromMicaService/ApplyDiscountNotification"
	DiscountFromMicaService_Ping_FullMethodName                      = "/mica.discount.service.v1.DiscountFromMicaService/Ping"
)

// DiscountFromMicaServiceClient is the client API for DiscountFromMicaService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type DiscountFromMicaServiceClient interface {
	ApplyDiscountNotification(ctx context.Context, in *discountv1.ApplyDiscountNotificationRequest, opts ...grpc.CallOption) (*discountv1.ApplyDiscountNotificationResponse, error)
	// An operation to ping the server to ensure it's up and running and that the connection is good.
	Ping(ctx context.Context, in *pingv1.PingRequest, opts ...grpc.CallOption) (*pingv1.PingResponse, error)
}

type discountFromMicaServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewDiscountFromMicaServiceClient(cc grpc.ClientConnInterface) DiscountFromMicaServiceClient {
	return &discountFromMicaServiceClient{cc}
}

func (c *discountFromMicaServiceClient) ApplyDiscountNotification(ctx context.Context, in *discountv1.ApplyDiscountNotificationRequest, opts ...grpc.CallOption) (*discountv1.ApplyDiscountNotificationResponse, error) {
	out := new(discountv1.ApplyDiscountNotificationResponse)
	err := c.cc.Invoke(ctx, DiscountFromMicaService_ApplyDiscountNotification_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *discountFromMicaServiceClient) Ping(ctx context.Context, in *pingv1.PingRequest, opts ...grpc.CallOption) (*pingv1.PingResponse, error) {
	out := new(pingv1.PingResponse)
	err := c.cc.Invoke(ctx, DiscountFromMicaService_Ping_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// DiscountFromMicaServiceServer is the server API for DiscountFromMicaService service.
// All implementations must embed UnimplementedDiscountFromMicaServiceServer
// for forward compatibility
type DiscountFromMicaServiceServer interface {
	ApplyDiscountNotification(context.Context, *discountv1.ApplyDiscountNotificationRequest) (*discountv1.ApplyDiscountNotificationResponse, error)
	// An operation to ping the server to ensure it's up and running and that the connection is good.
	Ping(context.Context, *pingv1.PingRequest) (*pingv1.PingResponse, error)
	mustEmbedUnimplementedDiscountFromMicaServiceServer()
}

// UnimplementedDiscountFromMicaServiceServer must be embedded to have forward compatible implementations.
type UnimplementedDiscountFromMicaServiceServer struct {
}

func (UnimplementedDiscountFromMicaServiceServer) ApplyDiscountNotification(context.Context, *discountv1.ApplyDiscountNotificationRequest) (*discountv1.ApplyDiscountNotificationResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ApplyDiscountNotification not implemented")
}
func (UnimplementedDiscountFromMicaServiceServer) Ping(context.Context, *pingv1.PingRequest) (*pingv1.PingResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Ping not implemented")
}
func (UnimplementedDiscountFromMicaServiceServer) mustEmbedUnimplementedDiscountFromMicaServiceServer() {
}

// UnsafeDiscountFromMicaServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to DiscountFromMicaServiceServer will
// result in compilation errors.
type UnsafeDiscountFromMicaServiceServer interface {
	mustEmbedUnimplementedDiscountFromMicaServiceServer()
}

func RegisterDiscountFromMicaServiceServer(s grpc.ServiceRegistrar, srv DiscountFromMicaServiceServer) {
	s.RegisterService(&DiscountFromMicaService_ServiceDesc, srv)
}

func _DiscountFromMicaService_ApplyDiscountNotification_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(discountv1.ApplyDiscountNotificationRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DiscountFromMicaServiceServer).ApplyDiscountNotification(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: DiscountFromMicaService_ApplyDiscountNotification_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DiscountFromMicaServiceServer).ApplyDiscountNotification(ctx, req.(*discountv1.ApplyDiscountNotificationRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _DiscountFromMicaService_Ping_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(pingv1.PingRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DiscountFromMicaServiceServer).Ping(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: DiscountFromMicaService_Ping_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DiscountFromMicaServiceServer).Ping(ctx, req.(*pingv1.PingRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// DiscountFromMicaService_ServiceDesc is the grpc.ServiceDesc for DiscountFromMicaService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var DiscountFromMicaService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "mica.discount.service.v1.DiscountFromMicaService",
	HandlerType: (*DiscountFromMicaServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "ApplyDiscountNotification",
			Handler:    _DiscountFromMicaService_ApplyDiscountNotification_Handler,
		},
		{
			MethodName: "Ping",
			Handler:    _DiscountFromMicaService_Ping_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "github.com/1080network/golang/discount/proto/mica/discount/service/v1/discount_from_mica_service.proto",
}
