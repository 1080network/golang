// Copyright (c) 2021 1080 Network, Inc. All rights reserved. All software, including, without limitation, all source
// code and object code, is the intellectual property of 1080 Network, Inc. and is protected by copyright, trademark and
// other intellectual property laws (collective "1080 Software"). You may not use, copy, reproduce, download, store,
// post, broadcast, transmit, modify, sell or make available to the public content from the 1080 Software without a
// valid license or the prior written approval of 1080 Network, Inc. 1080 Network, its logos, slogans, taglines,
// products, feature names, and other trademarks are trademarks of 1080 Network, Inc. and may not be used without
// permission.

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.23.0
// 	protoc        v3.14.0
// source: sp/service/v1/sp_from_ten80_service.proto

package servicev1

import (
	context "context"
	proto "github.com/golang/protobuf/proto"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	pingv1 "ten80/proto/common/pingv1"
	fundsv1 "ten80/proto/sp/fundsv1"
	instrumentv1 "ten80/proto/sp/instrumentv1"
	userv1 "ten80/proto/sp/userv1"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

// This is a compile-time assertion that a sufficiently up-to-date version
// of the legacy proto package is being used.
const _ = proto.ProtoPackageIsVersion4

var File_sp_service_v1_sp_from_ten80_service_proto protoreflect.FileDescriptor

var file_sp_service_v1_sp_from_ten80_service_proto_rawDesc = []byte{
	0x0a, 0x29, 0x73, 0x70, 0x2f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2f, 0x76, 0x31, 0x2f,
	0x73, 0x70, 0x5f, 0x66, 0x72, 0x6f, 0x6d, 0x5f, 0x74, 0x65, 0x6e, 0x38, 0x30, 0x5f, 0x73, 0x65,
	0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0d, 0x73, 0x70, 0x2e,
	0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x76, 0x31, 0x1a, 0x19, 0x63, 0x6f, 0x6d, 0x6d,
	0x6f, 0x6e, 0x2f, 0x70, 0x69, 0x6e, 0x67, 0x2f, 0x76, 0x31, 0x2f, 0x70, 0x69, 0x6e, 0x67, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x17, 0x73, 0x70, 0x2f, 0x66, 0x75, 0x6e, 0x64, 0x73, 0x2f,
	0x76, 0x31, 0x2f, 0x66, 0x75, 0x6e, 0x64, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x21,
	0x73, 0x70, 0x2f, 0x69, 0x6e, 0x73, 0x74, 0x72, 0x75, 0x6d, 0x65, 0x6e, 0x74, 0x2f, 0x76, 0x31,
	0x2f, 0x69, 0x6e, 0x73, 0x74, 0x72, 0x75, 0x6d, 0x65, 0x6e, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x1a, 0x15, 0x73, 0x70, 0x2f, 0x75, 0x73, 0x65, 0x72, 0x2f, 0x76, 0x31, 0x2f, 0x75, 0x73,
	0x65, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x32, 0xe3, 0x03, 0x0a, 0x11, 0x53, 0x50, 0x46,
	0x72, 0x6f, 0x6d, 0x31, 0x30, 0x38, 0x30, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x6b,
	0x0a, 0x14, 0x45, 0x6e, 0x72, 0x6f, 0x6c, 0x6c, 0x55, 0x73, 0x65, 0x72, 0x49, 0x6e, 0x73, 0x74,
	0x72, 0x75, 0x6d, 0x65, 0x6e, 0x74, 0x12, 0x27, 0x2e, 0x73, 0x70, 0x2e, 0x75, 0x73, 0x65, 0x72,
	0x2e, 0x76, 0x31, 0x2e, 0x45, 0x6e, 0x72, 0x6f, 0x6c, 0x6c, 0x55, 0x73, 0x65, 0x72, 0x49, 0x6e,
	0x73, 0x74, 0x72, 0x75, 0x6d, 0x65, 0x6e, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a,
	0x28, 0x2e, 0x73, 0x70, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x2e, 0x76, 0x31, 0x2e, 0x45, 0x6e, 0x72,
	0x6f, 0x6c, 0x6c, 0x55, 0x73, 0x65, 0x72, 0x49, 0x6e, 0x73, 0x74, 0x72, 0x75, 0x6d, 0x65, 0x6e,
	0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x74, 0x0a, 0x13, 0x52,
	0x65, 0x74, 0x72, 0x69, 0x65, 0x76, 0x65, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69,
	0x6f, 0x6e, 0x12, 0x2c, 0x2e, 0x73, 0x70, 0x2e, 0x69, 0x6e, 0x73, 0x74, 0x72, 0x75, 0x6d, 0x65,
	0x6e, 0x74, 0x2e, 0x76, 0x31, 0x2e, 0x52, 0x65, 0x74, 0x72, 0x69, 0x65, 0x76, 0x65, 0x54, 0x72,
	0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x1a, 0x2d, 0x2e, 0x73, 0x70, 0x2e, 0x69, 0x6e, 0x73, 0x74, 0x72, 0x75, 0x6d, 0x65, 0x6e, 0x74,
	0x2e, 0x76, 0x31, 0x2e, 0x52, 0x65, 0x74, 0x72, 0x69, 0x65, 0x76, 0x65, 0x54, 0x72, 0x61, 0x6e,
	0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22,
	0x00, 0x12, 0x52, 0x0a, 0x0b, 0x4f, 0x62, 0x74, 0x61, 0x69, 0x6e, 0x46, 0x75, 0x6e, 0x64, 0x73,
	0x12, 0x1f, 0x2e, 0x73, 0x70, 0x2e, 0x66, 0x75, 0x6e, 0x64, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x4f,
	0x62, 0x74, 0x61, 0x69, 0x6e, 0x46, 0x75, 0x6e, 0x64, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x1a, 0x20, 0x2e, 0x73, 0x70, 0x2e, 0x66, 0x75, 0x6e, 0x64, 0x73, 0x2e, 0x76, 0x31, 0x2e,
	0x4f, 0x62, 0x74, 0x61, 0x69, 0x6e, 0x46, 0x75, 0x6e, 0x64, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x52, 0x0a, 0x0b, 0x52, 0x65, 0x74, 0x75, 0x72, 0x6e, 0x46,
	0x75, 0x6e, 0x64, 0x73, 0x12, 0x1f, 0x2e, 0x73, 0x70, 0x2e, 0x66, 0x75, 0x6e, 0x64, 0x73, 0x2e,
	0x76, 0x31, 0x2e, 0x52, 0x65, 0x74, 0x75, 0x72, 0x6e, 0x46, 0x75, 0x6e, 0x64, 0x73, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x20, 0x2e, 0x73, 0x70, 0x2e, 0x66, 0x75, 0x6e, 0x64, 0x73,
	0x2e, 0x76, 0x31, 0x2e, 0x52, 0x65, 0x74, 0x75, 0x72, 0x6e, 0x46, 0x75, 0x6e, 0x64, 0x73, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x43, 0x0a, 0x04, 0x50, 0x69, 0x6e,
	0x67, 0x12, 0x1b, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x70, 0x69, 0x6e, 0x67, 0x2e,
	0x76, 0x31, 0x2e, 0x50, 0x69, 0x6e, 0x67, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1c,
	0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x70, 0x69, 0x6e, 0x67, 0x2e, 0x76, 0x31, 0x2e,
	0x50, 0x69, 0x6e, 0x67, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x42, 0x4c,
	0x0a, 0x18, 0x6e, 0x65, 0x74, 0x77, 0x6f, 0x72, 0x6b, 0x2e, 0x74, 0x65, 0x6e, 0x38, 0x30, 0x2e,
	0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x76, 0x31, 0x42, 0x0e, 0x53, 0x50, 0x53, 0x65,
	0x72, 0x76, 0x69, 0x63, 0x65, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x5a, 0x18, 0x74, 0x65, 0x6e, 0x38,
	0x30, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x73, 0x70, 0x2f, 0x73, 0x65, 0x72, 0x76, 0x69,
	0x63, 0x65, 0x76, 0x31, 0xa2, 0x02, 0x05, 0x54, 0x45, 0x4e, 0x38, 0x30, 0x62, 0x06, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x33,
}

var file_sp_service_v1_sp_from_ten80_service_proto_goTypes = []interface{}{
	(*userv1.EnrollUserInstrumentRequest)(nil),       // 0: sp.user.v1.EnrollUserInstrumentRequest
	(*instrumentv1.RetrieveTransactionRequest)(nil),  // 1: sp.instrument.v1.RetrieveTransactionRequest
	(*fundsv1.ObtainFundsRequest)(nil),               // 2: sp.funds.v1.ObtainFundsRequest
	(*fundsv1.ReturnFundsRequest)(nil),               // 3: sp.funds.v1.ReturnFundsRequest
	(*pingv1.PingRequest)(nil),                       // 4: common.ping.v1.PingRequest
	(*userv1.EnrollUserInstrumentResponse)(nil),      // 5: sp.user.v1.EnrollUserInstrumentResponse
	(*instrumentv1.RetrieveTransactionResponse)(nil), // 6: sp.instrument.v1.RetrieveTransactionResponse
	(*fundsv1.ObtainFundsResponse)(nil),              // 7: sp.funds.v1.ObtainFundsResponse
	(*fundsv1.ReturnFundsResponse)(nil),              // 8: sp.funds.v1.ReturnFundsResponse
	(*pingv1.PingResponse)(nil),                      // 9: common.ping.v1.PingResponse
}
var file_sp_service_v1_sp_from_ten80_service_proto_depIdxs = []int32{
	0, // 0: sp.service.v1.SPFrom1080Service.EnrollUserInstrument:input_type -> sp.user.v1.EnrollUserInstrumentRequest
	1, // 1: sp.service.v1.SPFrom1080Service.RetrieveTransaction:input_type -> sp.instrument.v1.RetrieveTransactionRequest
	2, // 2: sp.service.v1.SPFrom1080Service.ObtainFunds:input_type -> sp.funds.v1.ObtainFundsRequest
	3, // 3: sp.service.v1.SPFrom1080Service.ReturnFunds:input_type -> sp.funds.v1.ReturnFundsRequest
	4, // 4: sp.service.v1.SPFrom1080Service.Ping:input_type -> common.ping.v1.PingRequest
	5, // 5: sp.service.v1.SPFrom1080Service.EnrollUserInstrument:output_type -> sp.user.v1.EnrollUserInstrumentResponse
	6, // 6: sp.service.v1.SPFrom1080Service.RetrieveTransaction:output_type -> sp.instrument.v1.RetrieveTransactionResponse
	7, // 7: sp.service.v1.SPFrom1080Service.ObtainFunds:output_type -> sp.funds.v1.ObtainFundsResponse
	8, // 8: sp.service.v1.SPFrom1080Service.ReturnFunds:output_type -> sp.funds.v1.ReturnFundsResponse
	9, // 9: sp.service.v1.SPFrom1080Service.Ping:output_type -> common.ping.v1.PingResponse
	5, // [5:10] is the sub-list for method output_type
	0, // [0:5] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_sp_service_v1_sp_from_ten80_service_proto_init() }
func file_sp_service_v1_sp_from_ten80_service_proto_init() {
	if File_sp_service_v1_sp_from_ten80_service_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_sp_service_v1_sp_from_ten80_service_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   0,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_sp_service_v1_sp_from_ten80_service_proto_goTypes,
		DependencyIndexes: file_sp_service_v1_sp_from_ten80_service_proto_depIdxs,
	}.Build()
	File_sp_service_v1_sp_from_ten80_service_proto = out.File
	file_sp_service_v1_sp_from_ten80_service_proto_rawDesc = nil
	file_sp_service_v1_sp_from_ten80_service_proto_goTypes = nil
	file_sp_service_v1_sp_from_ten80_service_proto_depIdxs = nil
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// SPFrom1080ServiceClient is the client API for SPFrom1080Service service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type SPFrom1080ServiceClient interface {
	EnrollUserInstrument(ctx context.Context, in *userv1.EnrollUserInstrumentRequest, opts ...grpc.CallOption) (*userv1.EnrollUserInstrumentResponse, error)
	RetrieveTransaction(ctx context.Context, in *instrumentv1.RetrieveTransactionRequest, opts ...grpc.CallOption) (*instrumentv1.RetrieveTransactionResponse, error)
	ObtainFunds(ctx context.Context, in *fundsv1.ObtainFundsRequest, opts ...grpc.CallOption) (*fundsv1.ObtainFundsResponse, error)
	ReturnFunds(ctx context.Context, in *fundsv1.ReturnFundsRequest, opts ...grpc.CallOption) (*fundsv1.ReturnFundsResponse, error)
	Ping(ctx context.Context, in *pingv1.PingRequest, opts ...grpc.CallOption) (*pingv1.PingResponse, error)
}

type sPFrom1080ServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewSPFrom1080ServiceClient(cc grpc.ClientConnInterface) SPFrom1080ServiceClient {
	return &sPFrom1080ServiceClient{cc}
}

func (c *sPFrom1080ServiceClient) EnrollUserInstrument(ctx context.Context, in *userv1.EnrollUserInstrumentRequest, opts ...grpc.CallOption) (*userv1.EnrollUserInstrumentResponse, error) {
	out := new(userv1.EnrollUserInstrumentResponse)
	err := c.cc.Invoke(ctx, "/sp.service.v1.SPFrom1080Service/EnrollUserInstrument", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *sPFrom1080ServiceClient) RetrieveTransaction(ctx context.Context, in *instrumentv1.RetrieveTransactionRequest, opts ...grpc.CallOption) (*instrumentv1.RetrieveTransactionResponse, error) {
	out := new(instrumentv1.RetrieveTransactionResponse)
	err := c.cc.Invoke(ctx, "/sp.service.v1.SPFrom1080Service/RetrieveTransaction", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *sPFrom1080ServiceClient) ObtainFunds(ctx context.Context, in *fundsv1.ObtainFundsRequest, opts ...grpc.CallOption) (*fundsv1.ObtainFundsResponse, error) {
	out := new(fundsv1.ObtainFundsResponse)
	err := c.cc.Invoke(ctx, "/sp.service.v1.SPFrom1080Service/ObtainFunds", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *sPFrom1080ServiceClient) ReturnFunds(ctx context.Context, in *fundsv1.ReturnFundsRequest, opts ...grpc.CallOption) (*fundsv1.ReturnFundsResponse, error) {
	out := new(fundsv1.ReturnFundsResponse)
	err := c.cc.Invoke(ctx, "/sp.service.v1.SPFrom1080Service/ReturnFunds", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *sPFrom1080ServiceClient) Ping(ctx context.Context, in *pingv1.PingRequest, opts ...grpc.CallOption) (*pingv1.PingResponse, error) {
	out := new(pingv1.PingResponse)
	err := c.cc.Invoke(ctx, "/sp.service.v1.SPFrom1080Service/Ping", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// SPFrom1080ServiceServer is the server API for SPFrom1080Service service.
type SPFrom1080ServiceServer interface {
	EnrollUserInstrument(context.Context, *userv1.EnrollUserInstrumentRequest) (*userv1.EnrollUserInstrumentResponse, error)
	RetrieveTransaction(context.Context, *instrumentv1.RetrieveTransactionRequest) (*instrumentv1.RetrieveTransactionResponse, error)
	ObtainFunds(context.Context, *fundsv1.ObtainFundsRequest) (*fundsv1.ObtainFundsResponse, error)
	ReturnFunds(context.Context, *fundsv1.ReturnFundsRequest) (*fundsv1.ReturnFundsResponse, error)
	Ping(context.Context, *pingv1.PingRequest) (*pingv1.PingResponse, error)
}

// UnimplementedSPFrom1080ServiceServer can be embedded to have forward compatible implementations.
type UnimplementedSPFrom1080ServiceServer struct {
}

func (*UnimplementedSPFrom1080ServiceServer) EnrollUserInstrument(context.Context, *userv1.EnrollUserInstrumentRequest) (*userv1.EnrollUserInstrumentResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method EnrollUserInstrument not implemented")
}
func (*UnimplementedSPFrom1080ServiceServer) RetrieveTransaction(context.Context, *instrumentv1.RetrieveTransactionRequest) (*instrumentv1.RetrieveTransactionResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method RetrieveTransaction not implemented")
}
func (*UnimplementedSPFrom1080ServiceServer) ObtainFunds(context.Context, *fundsv1.ObtainFundsRequest) (*fundsv1.ObtainFundsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ObtainFunds not implemented")
}
func (*UnimplementedSPFrom1080ServiceServer) ReturnFunds(context.Context, *fundsv1.ReturnFundsRequest) (*fundsv1.ReturnFundsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ReturnFunds not implemented")
}
func (*UnimplementedSPFrom1080ServiceServer) Ping(context.Context, *pingv1.PingRequest) (*pingv1.PingResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Ping not implemented")
}

func RegisterSPFrom1080ServiceServer(s *grpc.Server, srv SPFrom1080ServiceServer) {
	s.RegisterService(&_SPFrom1080Service_serviceDesc, srv)
}

func _SPFrom1080Service_EnrollUserInstrument_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(userv1.EnrollUserInstrumentRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SPFrom1080ServiceServer).EnrollUserInstrument(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/sp.service.v1.SPFrom1080Service/EnrollUserInstrument",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SPFrom1080ServiceServer).EnrollUserInstrument(ctx, req.(*userv1.EnrollUserInstrumentRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _SPFrom1080Service_RetrieveTransaction_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(instrumentv1.RetrieveTransactionRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SPFrom1080ServiceServer).RetrieveTransaction(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/sp.service.v1.SPFrom1080Service/RetrieveTransaction",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SPFrom1080ServiceServer).RetrieveTransaction(ctx, req.(*instrumentv1.RetrieveTransactionRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _SPFrom1080Service_ObtainFunds_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(fundsv1.ObtainFundsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SPFrom1080ServiceServer).ObtainFunds(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/sp.service.v1.SPFrom1080Service/ObtainFunds",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SPFrom1080ServiceServer).ObtainFunds(ctx, req.(*fundsv1.ObtainFundsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _SPFrom1080Service_ReturnFunds_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(fundsv1.ReturnFundsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SPFrom1080ServiceServer).ReturnFunds(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/sp.service.v1.SPFrom1080Service/ReturnFunds",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SPFrom1080ServiceServer).ReturnFunds(ctx, req.(*fundsv1.ReturnFundsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _SPFrom1080Service_Ping_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(pingv1.PingRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SPFrom1080ServiceServer).Ping(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/sp.service.v1.SPFrom1080Service/Ping",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SPFrom1080ServiceServer).Ping(ctx, req.(*pingv1.PingRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _SPFrom1080Service_serviceDesc = grpc.ServiceDesc{
	ServiceName: "sp.service.v1.SPFrom1080Service",
	HandlerType: (*SPFrom1080ServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "EnrollUserInstrument",
			Handler:    _SPFrom1080Service_EnrollUserInstrument_Handler,
		},
		{
			MethodName: "RetrieveTransaction",
			Handler:    _SPFrom1080Service_RetrieveTransaction_Handler,
		},
		{
			MethodName: "ObtainFunds",
			Handler:    _SPFrom1080Service_ObtainFunds_Handler,
		},
		{
			MethodName: "ReturnFunds",
			Handler:    _SPFrom1080Service_ReturnFunds_Handler,
		},
		{
			MethodName: "Ping",
			Handler:    _SPFrom1080Service_Ping_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "sp/service/v1/sp_from_ten80_service.proto",
}
