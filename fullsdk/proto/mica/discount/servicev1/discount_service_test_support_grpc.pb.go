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
// source: mica/discount/service/v1/discount_service_test_support.proto

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
	DiscountTestSupportService_TestDetermineDiscount_FullMethodName = "/mica.discount.service.v1.DiscountTestSupportService/TestDetermineDiscount"
	DiscountTestSupportService_TestApplyDiscount_FullMethodName     = "/mica.discount.service.v1.DiscountTestSupportService/TestApplyDiscount"
	DiscountTestSupportService_TestReleaseDiscount_FullMethodName   = "/mica.discount.service.v1.DiscountTestSupportService/TestReleaseDiscount"
)

// DiscountTestSupportServiceClient is the client API for DiscountTestSupportService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type DiscountTestSupportServiceClient interface {
	// Determine which discount(s) can be applied to this transaction
	TestDetermineDiscount(ctx context.Context, in *DetermineDiscountRequest, opts ...grpc.CallOption) (*DetermineDiscountResponse, error)
	// Apply or release a given discount
	TestApplyDiscount(ctx context.Context, in *ApplyDiscountRequest, opts ...grpc.CallOption) (*ApplyDiscountResponse, error)
	TestReleaseDiscount(ctx context.Context, in *ReleaseDiscountRequest, opts ...grpc.CallOption) (*ReleaseDiscountResponse, error)
}

type discountTestSupportServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewDiscountTestSupportServiceClient(cc grpc.ClientConnInterface) DiscountTestSupportServiceClient {
	return &discountTestSupportServiceClient{cc}
}

func (c *discountTestSupportServiceClient) TestDetermineDiscount(ctx context.Context, in *DetermineDiscountRequest, opts ...grpc.CallOption) (*DetermineDiscountResponse, error) {
	out := new(DetermineDiscountResponse)
	err := c.cc.Invoke(ctx, DiscountTestSupportService_TestDetermineDiscount_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *discountTestSupportServiceClient) TestApplyDiscount(ctx context.Context, in *ApplyDiscountRequest, opts ...grpc.CallOption) (*ApplyDiscountResponse, error) {
	out := new(ApplyDiscountResponse)
	err := c.cc.Invoke(ctx, DiscountTestSupportService_TestApplyDiscount_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *discountTestSupportServiceClient) TestReleaseDiscount(ctx context.Context, in *ReleaseDiscountRequest, opts ...grpc.CallOption) (*ReleaseDiscountResponse, error) {
	out := new(ReleaseDiscountResponse)
	err := c.cc.Invoke(ctx, DiscountTestSupportService_TestReleaseDiscount_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// DiscountTestSupportServiceServer is the server API for DiscountTestSupportService service.
// All implementations must embed UnimplementedDiscountTestSupportServiceServer
// for forward compatibility
type DiscountTestSupportServiceServer interface {
	// Determine which discount(s) can be applied to this transaction
	TestDetermineDiscount(context.Context, *DetermineDiscountRequest) (*DetermineDiscountResponse, error)
	// Apply or release a given discount
	TestApplyDiscount(context.Context, *ApplyDiscountRequest) (*ApplyDiscountResponse, error)
	TestReleaseDiscount(context.Context, *ReleaseDiscountRequest) (*ReleaseDiscountResponse, error)
	mustEmbedUnimplementedDiscountTestSupportServiceServer()
}

// UnimplementedDiscountTestSupportServiceServer must be embedded to have forward compatible implementations.
type UnimplementedDiscountTestSupportServiceServer struct {
}

func (UnimplementedDiscountTestSupportServiceServer) TestDetermineDiscount(context.Context, *DetermineDiscountRequest) (*DetermineDiscountResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method TestDetermineDiscount not implemented")
}
func (UnimplementedDiscountTestSupportServiceServer) TestApplyDiscount(context.Context, *ApplyDiscountRequest) (*ApplyDiscountResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method TestApplyDiscount not implemented")
}
func (UnimplementedDiscountTestSupportServiceServer) TestReleaseDiscount(context.Context, *ReleaseDiscountRequest) (*ReleaseDiscountResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method TestReleaseDiscount not implemented")
}
func (UnimplementedDiscountTestSupportServiceServer) mustEmbedUnimplementedDiscountTestSupportServiceServer() {
}

// UnsafeDiscountTestSupportServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to DiscountTestSupportServiceServer will
// result in compilation errors.
type UnsafeDiscountTestSupportServiceServer interface {
	mustEmbedUnimplementedDiscountTestSupportServiceServer()
}

func RegisterDiscountTestSupportServiceServer(s grpc.ServiceRegistrar, srv DiscountTestSupportServiceServer) {
	s.RegisterService(&DiscountTestSupportService_ServiceDesc, srv)
}

func _DiscountTestSupportService_TestDetermineDiscount_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DetermineDiscountRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DiscountTestSupportServiceServer).TestDetermineDiscount(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: DiscountTestSupportService_TestDetermineDiscount_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DiscountTestSupportServiceServer).TestDetermineDiscount(ctx, req.(*DetermineDiscountRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _DiscountTestSupportService_TestApplyDiscount_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ApplyDiscountRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DiscountTestSupportServiceServer).TestApplyDiscount(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: DiscountTestSupportService_TestApplyDiscount_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DiscountTestSupportServiceServer).TestApplyDiscount(ctx, req.(*ApplyDiscountRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _DiscountTestSupportService_TestReleaseDiscount_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ReleaseDiscountRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DiscountTestSupportServiceServer).TestReleaseDiscount(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: DiscountTestSupportService_TestReleaseDiscount_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DiscountTestSupportServiceServer).TestReleaseDiscount(ctx, req.(*ReleaseDiscountRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// DiscountTestSupportService_ServiceDesc is the grpc.ServiceDesc for DiscountTestSupportService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var DiscountTestSupportService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "mica.discount.service.v1.DiscountTestSupportService",
	HandlerType: (*DiscountTestSupportServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "TestDetermineDiscount",
			Handler:    _DiscountTestSupportService_TestDetermineDiscount_Handler,
		},
		{
			MethodName: "TestApplyDiscount",
			Handler:    _DiscountTestSupportService_TestApplyDiscount_Handler,
		},
		{
			MethodName: "TestReleaseDiscount",
			Handler:    _DiscountTestSupportService_TestReleaseDiscount_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "github.com/1080network/golang/fullsdk/proto/mica/discount/service/v1/discount_service_test_support.proto",
}
