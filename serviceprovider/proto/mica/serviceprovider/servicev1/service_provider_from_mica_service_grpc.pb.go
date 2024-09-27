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
// source: mica/serviceprovider/service/v1/service_provider_from_mica_service.proto

package servicev1

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	instrumentv1 "github.com/1080network/golang/serviceprovider/proto/mica/serviceprovider/instrumentv1"
	userv1 "github.com/1080network/golang/serviceprovider/proto/mica/serviceprovider/userv1"
	valuev1 "github.com/1080network/golang/serviceprovider/proto/mica/serviceprovider/valuev1"
	pingv1 "github.com/1080network/golang/serviceprovider/proto/micashared/common/pingv1"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

const (
	ServiceProviderFromMicaService_EnrollUserInstrument_FullMethodName = "/mica.serviceprovider.service.v1.ServiceProviderFromMicaService/EnrollUserInstrument"
	ServiceProviderFromMicaService_RetrieveTransaction_FullMethodName  = "/mica.serviceprovider.service.v1.ServiceProviderFromMicaService/RetrieveTransaction"
	ServiceProviderFromMicaService_ObtainValue_FullMethodName          = "/mica.serviceprovider.service.v1.ServiceProviderFromMicaService/ObtainValue"
	ServiceProviderFromMicaService_ReverseObtainValue_FullMethodName   = "/mica.serviceprovider.service.v1.ServiceProviderFromMicaService/ReverseObtainValue"
	ServiceProviderFromMicaService_ReturnValue_FullMethodName          = "/mica.serviceprovider.service.v1.ServiceProviderFromMicaService/ReturnValue"
	ServiceProviderFromMicaService_ReverseReturnValue_FullMethodName   = "/mica.serviceprovider.service.v1.ServiceProviderFromMicaService/ReverseReturnValue"
	ServiceProviderFromMicaService_HoldValue_FullMethodName            = "/mica.serviceprovider.service.v1.ServiceProviderFromMicaService/HoldValue"
	ServiceProviderFromMicaService_AmendHoldValue_FullMethodName       = "/mica.serviceprovider.service.v1.ServiceProviderFromMicaService/AmendHoldValue"
	ServiceProviderFromMicaService_ReleaseHoldValue_FullMethodName     = "/mica.serviceprovider.service.v1.ServiceProviderFromMicaService/ReleaseHoldValue"
	ServiceProviderFromMicaService_ObtainHoldValue_FullMethodName      = "/mica.serviceprovider.service.v1.ServiceProviderFromMicaService/ObtainHoldValue"
	ServiceProviderFromMicaService_ReceiveValue_FullMethodName         = "/mica.serviceprovider.service.v1.ServiceProviderFromMicaService/ReceiveValue"
	ServiceProviderFromMicaService_ValueAdvice_FullMethodName          = "/mica.serviceprovider.service.v1.ServiceProviderFromMicaService/ValueAdvice"
	ServiceProviderFromMicaService_Ping_FullMethodName                 = "/mica.serviceprovider.service.v1.ServiceProviderFromMicaService/Ping"
)

// ServiceProviderFromMicaServiceClient is the client API for ServiceProviderFromMicaService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type ServiceProviderFromMicaServiceClient interface {
	// When a User initiates an enrollment at a Partner, mica will call this operation in order to
	// pass the SP the matching code. The SP will then send the matching code to their User through
	// whatever channel they already use to message their users.
	EnrollUserInstrument(ctx context.Context, in *userv1.EnrollUserInstrumentRequest, opts ...grpc.CallOption) (*userv1.EnrollUserInstrumentResponse, error)
	// Operation that can be used to retrieve a very simple transaction history for a given
	// instrument and date range.
	RetrieveTransaction(ctx context.Context, in *instrumentv1.RetrieveTransactionRequest, opts ...grpc.CallOption) (*instrumentv1.RetrieveTransactionResponse, error)
	// Operation to obtain value from a given instrument. Along with a common Value object (see
	// above), it includes an approval type of either FULL or PARTIAL.
	ObtainValue(ctx context.Context, in *valuev1.ObtainValueRequest, opts ...grpc.CallOption) (*valuev1.ObtainValueResponse, error)
	// Reverse a ObtainValue transaction
	ReverseObtainValue(ctx context.Context, in *valuev1.ReverseValueRequest, opts ...grpc.CallOption) (*valuev1.ReverseValueResponse, error)
	// Operation to return value to a given instrument.
	ReturnValue(ctx context.Context, in *valuev1.ReturnValueRequest, opts ...grpc.CallOption) (*valuev1.ReturnValueResponse, error)
	// Reverse a ReturnValue transaction
	ReverseReturnValue(ctx context.Context, in *valuev1.ReverseValueRequest, opts ...grpc.CallOption) (*valuev1.ReverseValueResponse, error)
	// Hold a value for a given instrument to be used at a later time
	HoldValue(ctx context.Context, in *valuev1.HoldValueRequest, opts ...grpc.CallOption) (*valuev1.HoldValueResponse, error)
	// Release a previously placed hold
	AmendHoldValue(ctx context.Context, in *valuev1.AmendHoldValueRequest, opts ...grpc.CallOption) (*valuev1.AmendHoldValueResponse, error)
	// Release a value that was previously held
	ReleaseHoldValue(ctx context.Context, in *valuev1.ReleaseHoldValueRequest, opts ...grpc.CallOption) (*valuev1.ReleaseHoldValueResponse, error)
	// Obtain a value that was previously held
	ObtainHoldValue(ctx context.Context, in *valuev1.ObtainHoldValueRequest, opts ...grpc.CallOption) (*valuev1.ObtainHoldValueResponse, error)
	ReceiveValue(ctx context.Context, in *valuev1.ReceiveValueRequest, opts ...grpc.CallOption) (*valuev1.ReceiveValueResponse, error)
	ValueAdvice(ctx context.Context, in *valuev1.ValueAdviceRequest, opts ...grpc.CallOption) (*valuev1.ValueAdviceResponse, error)
	// An operation to ping the server to ensure it's up and running and that the connection is good.
	Ping(ctx context.Context, in *pingv1.PingRequest, opts ...grpc.CallOption) (*pingv1.PingResponse, error)
}

type serviceProviderFromMicaServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewServiceProviderFromMicaServiceClient(cc grpc.ClientConnInterface) ServiceProviderFromMicaServiceClient {
	return &serviceProviderFromMicaServiceClient{cc}
}

func (c *serviceProviderFromMicaServiceClient) EnrollUserInstrument(ctx context.Context, in *userv1.EnrollUserInstrumentRequest, opts ...grpc.CallOption) (*userv1.EnrollUserInstrumentResponse, error) {
	out := new(userv1.EnrollUserInstrumentResponse)
	err := c.cc.Invoke(ctx, ServiceProviderFromMicaService_EnrollUserInstrument_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *serviceProviderFromMicaServiceClient) RetrieveTransaction(ctx context.Context, in *instrumentv1.RetrieveTransactionRequest, opts ...grpc.CallOption) (*instrumentv1.RetrieveTransactionResponse, error) {
	out := new(instrumentv1.RetrieveTransactionResponse)
	err := c.cc.Invoke(ctx, ServiceProviderFromMicaService_RetrieveTransaction_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *serviceProviderFromMicaServiceClient) ObtainValue(ctx context.Context, in *valuev1.ObtainValueRequest, opts ...grpc.CallOption) (*valuev1.ObtainValueResponse, error) {
	out := new(valuev1.ObtainValueResponse)
	err := c.cc.Invoke(ctx, ServiceProviderFromMicaService_ObtainValue_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *serviceProviderFromMicaServiceClient) ReverseObtainValue(ctx context.Context, in *valuev1.ReverseValueRequest, opts ...grpc.CallOption) (*valuev1.ReverseValueResponse, error) {
	out := new(valuev1.ReverseValueResponse)
	err := c.cc.Invoke(ctx, ServiceProviderFromMicaService_ReverseObtainValue_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *serviceProviderFromMicaServiceClient) ReturnValue(ctx context.Context, in *valuev1.ReturnValueRequest, opts ...grpc.CallOption) (*valuev1.ReturnValueResponse, error) {
	out := new(valuev1.ReturnValueResponse)
	err := c.cc.Invoke(ctx, ServiceProviderFromMicaService_ReturnValue_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *serviceProviderFromMicaServiceClient) ReverseReturnValue(ctx context.Context, in *valuev1.ReverseValueRequest, opts ...grpc.CallOption) (*valuev1.ReverseValueResponse, error) {
	out := new(valuev1.ReverseValueResponse)
	err := c.cc.Invoke(ctx, ServiceProviderFromMicaService_ReverseReturnValue_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *serviceProviderFromMicaServiceClient) HoldValue(ctx context.Context, in *valuev1.HoldValueRequest, opts ...grpc.CallOption) (*valuev1.HoldValueResponse, error) {
	out := new(valuev1.HoldValueResponse)
	err := c.cc.Invoke(ctx, ServiceProviderFromMicaService_HoldValue_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *serviceProviderFromMicaServiceClient) AmendHoldValue(ctx context.Context, in *valuev1.AmendHoldValueRequest, opts ...grpc.CallOption) (*valuev1.AmendHoldValueResponse, error) {
	out := new(valuev1.AmendHoldValueResponse)
	err := c.cc.Invoke(ctx, ServiceProviderFromMicaService_AmendHoldValue_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *serviceProviderFromMicaServiceClient) ReleaseHoldValue(ctx context.Context, in *valuev1.ReleaseHoldValueRequest, opts ...grpc.CallOption) (*valuev1.ReleaseHoldValueResponse, error) {
	out := new(valuev1.ReleaseHoldValueResponse)
	err := c.cc.Invoke(ctx, ServiceProviderFromMicaService_ReleaseHoldValue_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *serviceProviderFromMicaServiceClient) ObtainHoldValue(ctx context.Context, in *valuev1.ObtainHoldValueRequest, opts ...grpc.CallOption) (*valuev1.ObtainHoldValueResponse, error) {
	out := new(valuev1.ObtainHoldValueResponse)
	err := c.cc.Invoke(ctx, ServiceProviderFromMicaService_ObtainHoldValue_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *serviceProviderFromMicaServiceClient) ReceiveValue(ctx context.Context, in *valuev1.ReceiveValueRequest, opts ...grpc.CallOption) (*valuev1.ReceiveValueResponse, error) {
	out := new(valuev1.ReceiveValueResponse)
	err := c.cc.Invoke(ctx, ServiceProviderFromMicaService_ReceiveValue_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *serviceProviderFromMicaServiceClient) ValueAdvice(ctx context.Context, in *valuev1.ValueAdviceRequest, opts ...grpc.CallOption) (*valuev1.ValueAdviceResponse, error) {
	out := new(valuev1.ValueAdviceResponse)
	err := c.cc.Invoke(ctx, ServiceProviderFromMicaService_ValueAdvice_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *serviceProviderFromMicaServiceClient) Ping(ctx context.Context, in *pingv1.PingRequest, opts ...grpc.CallOption) (*pingv1.PingResponse, error) {
	out := new(pingv1.PingResponse)
	err := c.cc.Invoke(ctx, ServiceProviderFromMicaService_Ping_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ServiceProviderFromMicaServiceServer is the server API for ServiceProviderFromMicaService service.
// All implementations must embed UnimplementedServiceProviderFromMicaServiceServer
// for forward compatibility
type ServiceProviderFromMicaServiceServer interface {
	// When a User initiates an enrollment at a Partner, mica will call this operation in order to
	// pass the SP the matching code. The SP will then send the matching code to their User through
	// whatever channel they already use to message their users.
	EnrollUserInstrument(context.Context, *userv1.EnrollUserInstrumentRequest) (*userv1.EnrollUserInstrumentResponse, error)
	// Operation that can be used to retrieve a very simple transaction history for a given
	// instrument and date range.
	RetrieveTransaction(context.Context, *instrumentv1.RetrieveTransactionRequest) (*instrumentv1.RetrieveTransactionResponse, error)
	// Operation to obtain value from a given instrument. Along with a common Value object (see
	// above), it includes an approval type of either FULL or PARTIAL.
	ObtainValue(context.Context, *valuev1.ObtainValueRequest) (*valuev1.ObtainValueResponse, error)
	// Reverse a ObtainValue transaction
	ReverseObtainValue(context.Context, *valuev1.ReverseValueRequest) (*valuev1.ReverseValueResponse, error)
	// Operation to return value to a given instrument.
	ReturnValue(context.Context, *valuev1.ReturnValueRequest) (*valuev1.ReturnValueResponse, error)
	// Reverse a ReturnValue transaction
	ReverseReturnValue(context.Context, *valuev1.ReverseValueRequest) (*valuev1.ReverseValueResponse, error)
	// Hold a value for a given instrument to be used at a later time
	HoldValue(context.Context, *valuev1.HoldValueRequest) (*valuev1.HoldValueResponse, error)
	// Release a previously placed hold
	AmendHoldValue(context.Context, *valuev1.AmendHoldValueRequest) (*valuev1.AmendHoldValueResponse, error)
	// Release a value that was previously held
	ReleaseHoldValue(context.Context, *valuev1.ReleaseHoldValueRequest) (*valuev1.ReleaseHoldValueResponse, error)
	// Obtain a value that was previously held
	ObtainHoldValue(context.Context, *valuev1.ObtainHoldValueRequest) (*valuev1.ObtainHoldValueResponse, error)
	ReceiveValue(context.Context, *valuev1.ReceiveValueRequest) (*valuev1.ReceiveValueResponse, error)
	ValueAdvice(context.Context, *valuev1.ValueAdviceRequest) (*valuev1.ValueAdviceResponse, error)
	// An operation to ping the server to ensure it's up and running and that the connection is good.
	Ping(context.Context, *pingv1.PingRequest) (*pingv1.PingResponse, error)
	mustEmbedUnimplementedServiceProviderFromMicaServiceServer()
}

// UnimplementedServiceProviderFromMicaServiceServer must be embedded to have forward compatible implementations.
type UnimplementedServiceProviderFromMicaServiceServer struct {
}

func (UnimplementedServiceProviderFromMicaServiceServer) EnrollUserInstrument(context.Context, *userv1.EnrollUserInstrumentRequest) (*userv1.EnrollUserInstrumentResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method EnrollUserInstrument not implemented")
}
func (UnimplementedServiceProviderFromMicaServiceServer) RetrieveTransaction(context.Context, *instrumentv1.RetrieveTransactionRequest) (*instrumentv1.RetrieveTransactionResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method RetrieveTransaction not implemented")
}
func (UnimplementedServiceProviderFromMicaServiceServer) ObtainValue(context.Context, *valuev1.ObtainValueRequest) (*valuev1.ObtainValueResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ObtainValue not implemented")
}
func (UnimplementedServiceProviderFromMicaServiceServer) ReverseObtainValue(context.Context, *valuev1.ReverseValueRequest) (*valuev1.ReverseValueResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ReverseObtainValue not implemented")
}
func (UnimplementedServiceProviderFromMicaServiceServer) ReturnValue(context.Context, *valuev1.ReturnValueRequest) (*valuev1.ReturnValueResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ReturnValue not implemented")
}
func (UnimplementedServiceProviderFromMicaServiceServer) ReverseReturnValue(context.Context, *valuev1.ReverseValueRequest) (*valuev1.ReverseValueResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ReverseReturnValue not implemented")
}
func (UnimplementedServiceProviderFromMicaServiceServer) HoldValue(context.Context, *valuev1.HoldValueRequest) (*valuev1.HoldValueResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method HoldValue not implemented")
}
func (UnimplementedServiceProviderFromMicaServiceServer) AmendHoldValue(context.Context, *valuev1.AmendHoldValueRequest) (*valuev1.AmendHoldValueResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AmendHoldValue not implemented")
}
func (UnimplementedServiceProviderFromMicaServiceServer) ReleaseHoldValue(context.Context, *valuev1.ReleaseHoldValueRequest) (*valuev1.ReleaseHoldValueResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ReleaseHoldValue not implemented")
}
func (UnimplementedServiceProviderFromMicaServiceServer) ObtainHoldValue(context.Context, *valuev1.ObtainHoldValueRequest) (*valuev1.ObtainHoldValueResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ObtainHoldValue not implemented")
}
func (UnimplementedServiceProviderFromMicaServiceServer) ReceiveValue(context.Context, *valuev1.ReceiveValueRequest) (*valuev1.ReceiveValueResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ReceiveValue not implemented")
}
func (UnimplementedServiceProviderFromMicaServiceServer) ValueAdvice(context.Context, *valuev1.ValueAdviceRequest) (*valuev1.ValueAdviceResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ValueAdvice not implemented")
}
func (UnimplementedServiceProviderFromMicaServiceServer) Ping(context.Context, *pingv1.PingRequest) (*pingv1.PingResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Ping not implemented")
}
func (UnimplementedServiceProviderFromMicaServiceServer) mustEmbedUnimplementedServiceProviderFromMicaServiceServer() {
}

// UnsafeServiceProviderFromMicaServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to ServiceProviderFromMicaServiceServer will
// result in compilation errors.
type UnsafeServiceProviderFromMicaServiceServer interface {
	mustEmbedUnimplementedServiceProviderFromMicaServiceServer()
}

func RegisterServiceProviderFromMicaServiceServer(s grpc.ServiceRegistrar, srv ServiceProviderFromMicaServiceServer) {
	s.RegisterService(&ServiceProviderFromMicaService_ServiceDesc, srv)
}

func _ServiceProviderFromMicaService_EnrollUserInstrument_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(userv1.EnrollUserInstrumentRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ServiceProviderFromMicaServiceServer).EnrollUserInstrument(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ServiceProviderFromMicaService_EnrollUserInstrument_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ServiceProviderFromMicaServiceServer).EnrollUserInstrument(ctx, req.(*userv1.EnrollUserInstrumentRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ServiceProviderFromMicaService_RetrieveTransaction_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(instrumentv1.RetrieveTransactionRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ServiceProviderFromMicaServiceServer).RetrieveTransaction(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ServiceProviderFromMicaService_RetrieveTransaction_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ServiceProviderFromMicaServiceServer).RetrieveTransaction(ctx, req.(*instrumentv1.RetrieveTransactionRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ServiceProviderFromMicaService_ObtainValue_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(valuev1.ObtainValueRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ServiceProviderFromMicaServiceServer).ObtainValue(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ServiceProviderFromMicaService_ObtainValue_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ServiceProviderFromMicaServiceServer).ObtainValue(ctx, req.(*valuev1.ObtainValueRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ServiceProviderFromMicaService_ReverseObtainValue_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(valuev1.ReverseValueRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ServiceProviderFromMicaServiceServer).ReverseObtainValue(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ServiceProviderFromMicaService_ReverseObtainValue_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ServiceProviderFromMicaServiceServer).ReverseObtainValue(ctx, req.(*valuev1.ReverseValueRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ServiceProviderFromMicaService_ReturnValue_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(valuev1.ReturnValueRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ServiceProviderFromMicaServiceServer).ReturnValue(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ServiceProviderFromMicaService_ReturnValue_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ServiceProviderFromMicaServiceServer).ReturnValue(ctx, req.(*valuev1.ReturnValueRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ServiceProviderFromMicaService_ReverseReturnValue_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(valuev1.ReverseValueRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ServiceProviderFromMicaServiceServer).ReverseReturnValue(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ServiceProviderFromMicaService_ReverseReturnValue_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ServiceProviderFromMicaServiceServer).ReverseReturnValue(ctx, req.(*valuev1.ReverseValueRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ServiceProviderFromMicaService_HoldValue_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(valuev1.HoldValueRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ServiceProviderFromMicaServiceServer).HoldValue(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ServiceProviderFromMicaService_HoldValue_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ServiceProviderFromMicaServiceServer).HoldValue(ctx, req.(*valuev1.HoldValueRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ServiceProviderFromMicaService_AmendHoldValue_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(valuev1.AmendHoldValueRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ServiceProviderFromMicaServiceServer).AmendHoldValue(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ServiceProviderFromMicaService_AmendHoldValue_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ServiceProviderFromMicaServiceServer).AmendHoldValue(ctx, req.(*valuev1.AmendHoldValueRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ServiceProviderFromMicaService_ReleaseHoldValue_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(valuev1.ReleaseHoldValueRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ServiceProviderFromMicaServiceServer).ReleaseHoldValue(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ServiceProviderFromMicaService_ReleaseHoldValue_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ServiceProviderFromMicaServiceServer).ReleaseHoldValue(ctx, req.(*valuev1.ReleaseHoldValueRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ServiceProviderFromMicaService_ObtainHoldValue_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(valuev1.ObtainHoldValueRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ServiceProviderFromMicaServiceServer).ObtainHoldValue(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ServiceProviderFromMicaService_ObtainHoldValue_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ServiceProviderFromMicaServiceServer).ObtainHoldValue(ctx, req.(*valuev1.ObtainHoldValueRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ServiceProviderFromMicaService_ReceiveValue_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(valuev1.ReceiveValueRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ServiceProviderFromMicaServiceServer).ReceiveValue(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ServiceProviderFromMicaService_ReceiveValue_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ServiceProviderFromMicaServiceServer).ReceiveValue(ctx, req.(*valuev1.ReceiveValueRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ServiceProviderFromMicaService_ValueAdvice_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(valuev1.ValueAdviceRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ServiceProviderFromMicaServiceServer).ValueAdvice(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ServiceProviderFromMicaService_ValueAdvice_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ServiceProviderFromMicaServiceServer).ValueAdvice(ctx, req.(*valuev1.ValueAdviceRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ServiceProviderFromMicaService_Ping_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(pingv1.PingRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ServiceProviderFromMicaServiceServer).Ping(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ServiceProviderFromMicaService_Ping_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ServiceProviderFromMicaServiceServer).Ping(ctx, req.(*pingv1.PingRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// ServiceProviderFromMicaService_ServiceDesc is the grpc.ServiceDesc for ServiceProviderFromMicaService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var ServiceProviderFromMicaService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "mica.serviceprovider.service.v1.ServiceProviderFromMicaService",
	HandlerType: (*ServiceProviderFromMicaServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "EnrollUserInstrument",
			Handler:    _ServiceProviderFromMicaService_EnrollUserInstrument_Handler,
		},
		{
			MethodName: "RetrieveTransaction",
			Handler:    _ServiceProviderFromMicaService_RetrieveTransaction_Handler,
		},
		{
			MethodName: "ObtainValue",
			Handler:    _ServiceProviderFromMicaService_ObtainValue_Handler,
		},
		{
			MethodName: "ReverseObtainValue",
			Handler:    _ServiceProviderFromMicaService_ReverseObtainValue_Handler,
		},
		{
			MethodName: "ReturnValue",
			Handler:    _ServiceProviderFromMicaService_ReturnValue_Handler,
		},
		{
			MethodName: "ReverseReturnValue",
			Handler:    _ServiceProviderFromMicaService_ReverseReturnValue_Handler,
		},
		{
			MethodName: "HoldValue",
			Handler:    _ServiceProviderFromMicaService_HoldValue_Handler,
		},
		{
			MethodName: "AmendHoldValue",
			Handler:    _ServiceProviderFromMicaService_AmendHoldValue_Handler,
		},
		{
			MethodName: "ReleaseHoldValue",
			Handler:    _ServiceProviderFromMicaService_ReleaseHoldValue_Handler,
		},
		{
			MethodName: "ObtainHoldValue",
			Handler:    _ServiceProviderFromMicaService_ObtainHoldValue_Handler,
		},
		{
			MethodName: "ReceiveValue",
			Handler:    _ServiceProviderFromMicaService_ReceiveValue_Handler,
		},
		{
			MethodName: "ValueAdvice",
			Handler:    _ServiceProviderFromMicaService_ValueAdvice_Handler,
		},
		{
			MethodName: "Ping",
			Handler:    _ServiceProviderFromMicaService_Ping_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "github.com/1080network/golang/serviceprovider/proto/mica/serviceprovider/service/v1/service_provider_from_mica_service.proto",
}
