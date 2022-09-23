// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.19.4
// source: partner/administration/v1/admin_service.proto

package administrationv1

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	pingv1 "github.com/1080network/golang/partner/proto/common/pingv1"
	v1 "github.com/1080network/golang/partner/proto/common/v1"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

// PartnerAdministrationServiceClient is the client API for PartnerAdministrationService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type PartnerAdministrationServiceClient interface {
	//  Generate a new MTLS certificate.
	GenerateMTLSCertificate(ctx context.Context, in *v1.GenerateMTLSCertificateRequest, opts ...grpc.CallOption) (*v1.GenerateMTLSCertificateResponse, error)
	// Update the certificate with a given serial number, only supports enable/disable for now
	UpdateMTLSCertificate(ctx context.Context, in *v1.UpdateMTLSCertificateRequest, opts ...grpc.CallOption) (*v1.UpdateMTLSCertificateResponse, error)
	//Search for certificates and return the ones that match the criteria provided
	SearchMTLSCertificate(ctx context.Context, in *v1.SearchMTLSCertificateRequest, opts ...grpc.CallOption) (*v1.SearchMTLSCertificateResponse, error)
	GetMTLSCertificate(ctx context.Context, in *v1.GetMTLSCertificateRequest, opts ...grpc.CallOption) (*v1.GetMTLSCertificateResponse, error)
	CreateSingleSignOnConsoleUser(ctx context.Context, in *v1.CreateSingleSignOnConsoleUserRequest, opts ...grpc.CallOption) (*v1.CreateSingleSignOnConsoleUserResponse, error)
	UpdateSingleSignOnConsoleUser(ctx context.Context, in *v1.UpdateSingleSignOnConsoleUserRequest, opts ...grpc.CallOption) (*v1.UpdateSingleSignOnConsoleUserResponse, error)
	SearchSingleSignOnUser(ctx context.Context, in *v1.SearchSingleSignOnConsoleUserRequest, opts ...grpc.CallOption) (*v1.SearchSingleSignOnConsoleUserResponse, error)
	GetSingleSignOnConsoleUser(ctx context.Context, in *v1.GetSingleSignOnConsoleUserRequest, opts ...grpc.CallOption) (*v1.GetSingleSignOnConsoleUserResponse, error)
	PingExternal(ctx context.Context, in *pingv1.PingRequest, opts ...grpc.CallOption) (*pingv1.PingResponse, error)
}

type partnerAdministrationServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewPartnerAdministrationServiceClient(cc grpc.ClientConnInterface) PartnerAdministrationServiceClient {
	return &partnerAdministrationServiceClient{cc}
}

func (c *partnerAdministrationServiceClient) GenerateMTLSCertificate(ctx context.Context, in *v1.GenerateMTLSCertificateRequest, opts ...grpc.CallOption) (*v1.GenerateMTLSCertificateResponse, error) {
	out := new(v1.GenerateMTLSCertificateResponse)
	err := c.cc.Invoke(ctx, "/partner.administration.v1.PartnerAdministrationService/GenerateMTLSCertificate", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *partnerAdministrationServiceClient) UpdateMTLSCertificate(ctx context.Context, in *v1.UpdateMTLSCertificateRequest, opts ...grpc.CallOption) (*v1.UpdateMTLSCertificateResponse, error) {
	out := new(v1.UpdateMTLSCertificateResponse)
	err := c.cc.Invoke(ctx, "/partner.administration.v1.PartnerAdministrationService/UpdateMTLSCertificate", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *partnerAdministrationServiceClient) SearchMTLSCertificate(ctx context.Context, in *v1.SearchMTLSCertificateRequest, opts ...grpc.CallOption) (*v1.SearchMTLSCertificateResponse, error) {
	out := new(v1.SearchMTLSCertificateResponse)
	err := c.cc.Invoke(ctx, "/partner.administration.v1.PartnerAdministrationService/SearchMTLSCertificate", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *partnerAdministrationServiceClient) GetMTLSCertificate(ctx context.Context, in *v1.GetMTLSCertificateRequest, opts ...grpc.CallOption) (*v1.GetMTLSCertificateResponse, error) {
	out := new(v1.GetMTLSCertificateResponse)
	err := c.cc.Invoke(ctx, "/partner.administration.v1.PartnerAdministrationService/GetMTLSCertificate", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *partnerAdministrationServiceClient) CreateSingleSignOnConsoleUser(ctx context.Context, in *v1.CreateSingleSignOnConsoleUserRequest, opts ...grpc.CallOption) (*v1.CreateSingleSignOnConsoleUserResponse, error) {
	out := new(v1.CreateSingleSignOnConsoleUserResponse)
	err := c.cc.Invoke(ctx, "/partner.administration.v1.PartnerAdministrationService/CreateSingleSignOnConsoleUser", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *partnerAdministrationServiceClient) UpdateSingleSignOnConsoleUser(ctx context.Context, in *v1.UpdateSingleSignOnConsoleUserRequest, opts ...grpc.CallOption) (*v1.UpdateSingleSignOnConsoleUserResponse, error) {
	out := new(v1.UpdateSingleSignOnConsoleUserResponse)
	err := c.cc.Invoke(ctx, "/partner.administration.v1.PartnerAdministrationService/UpdateSingleSignOnConsoleUser", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *partnerAdministrationServiceClient) SearchSingleSignOnUser(ctx context.Context, in *v1.SearchSingleSignOnConsoleUserRequest, opts ...grpc.CallOption) (*v1.SearchSingleSignOnConsoleUserResponse, error) {
	out := new(v1.SearchSingleSignOnConsoleUserResponse)
	err := c.cc.Invoke(ctx, "/partner.administration.v1.PartnerAdministrationService/SearchSingleSignOnUser", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *partnerAdministrationServiceClient) GetSingleSignOnConsoleUser(ctx context.Context, in *v1.GetSingleSignOnConsoleUserRequest, opts ...grpc.CallOption) (*v1.GetSingleSignOnConsoleUserResponse, error) {
	out := new(v1.GetSingleSignOnConsoleUserResponse)
	err := c.cc.Invoke(ctx, "/partner.administration.v1.PartnerAdministrationService/GetSingleSignOnConsoleUser", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *partnerAdministrationServiceClient) PingExternal(ctx context.Context, in *pingv1.PingRequest, opts ...grpc.CallOption) (*pingv1.PingResponse, error) {
	out := new(pingv1.PingResponse)
	err := c.cc.Invoke(ctx, "/partner.administration.v1.PartnerAdministrationService/PingExternal", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// PartnerAdministrationServiceServer is the server API for PartnerAdministrationService service.
// All implementations must embed UnimplementedPartnerAdministrationServiceServer
// for forward compatibility
type PartnerAdministrationServiceServer interface {
	//  Generate a new MTLS certificate.
	GenerateMTLSCertificate(context.Context, *v1.GenerateMTLSCertificateRequest) (*v1.GenerateMTLSCertificateResponse, error)
	// Update the certificate with a given serial number, only supports enable/disable for now
	UpdateMTLSCertificate(context.Context, *v1.UpdateMTLSCertificateRequest) (*v1.UpdateMTLSCertificateResponse, error)
	//Search for certificates and return the ones that match the criteria provided
	SearchMTLSCertificate(context.Context, *v1.SearchMTLSCertificateRequest) (*v1.SearchMTLSCertificateResponse, error)
	GetMTLSCertificate(context.Context, *v1.GetMTLSCertificateRequest) (*v1.GetMTLSCertificateResponse, error)
	CreateSingleSignOnConsoleUser(context.Context, *v1.CreateSingleSignOnConsoleUserRequest) (*v1.CreateSingleSignOnConsoleUserResponse, error)
	UpdateSingleSignOnConsoleUser(context.Context, *v1.UpdateSingleSignOnConsoleUserRequest) (*v1.UpdateSingleSignOnConsoleUserResponse, error)
	SearchSingleSignOnUser(context.Context, *v1.SearchSingleSignOnConsoleUserRequest) (*v1.SearchSingleSignOnConsoleUserResponse, error)
	GetSingleSignOnConsoleUser(context.Context, *v1.GetSingleSignOnConsoleUserRequest) (*v1.GetSingleSignOnConsoleUserResponse, error)
	PingExternal(context.Context, *pingv1.PingRequest) (*pingv1.PingResponse, error)
	mustEmbedUnimplementedPartnerAdministrationServiceServer()
}

// UnimplementedPartnerAdministrationServiceServer must be embedded to have forward compatible implementations.
type UnimplementedPartnerAdministrationServiceServer struct {
}

func (UnimplementedPartnerAdministrationServiceServer) GenerateMTLSCertificate(context.Context, *v1.GenerateMTLSCertificateRequest) (*v1.GenerateMTLSCertificateResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GenerateMTLSCertificate not implemented")
}
func (UnimplementedPartnerAdministrationServiceServer) UpdateMTLSCertificate(context.Context, *v1.UpdateMTLSCertificateRequest) (*v1.UpdateMTLSCertificateResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateMTLSCertificate not implemented")
}
func (UnimplementedPartnerAdministrationServiceServer) SearchMTLSCertificate(context.Context, *v1.SearchMTLSCertificateRequest) (*v1.SearchMTLSCertificateResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SearchMTLSCertificate not implemented")
}
func (UnimplementedPartnerAdministrationServiceServer) GetMTLSCertificate(context.Context, *v1.GetMTLSCertificateRequest) (*v1.GetMTLSCertificateResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetMTLSCertificate not implemented")
}
func (UnimplementedPartnerAdministrationServiceServer) CreateSingleSignOnConsoleUser(context.Context, *v1.CreateSingleSignOnConsoleUserRequest) (*v1.CreateSingleSignOnConsoleUserResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateSingleSignOnConsoleUser not implemented")
}
func (UnimplementedPartnerAdministrationServiceServer) UpdateSingleSignOnConsoleUser(context.Context, *v1.UpdateSingleSignOnConsoleUserRequest) (*v1.UpdateSingleSignOnConsoleUserResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateSingleSignOnConsoleUser not implemented")
}
func (UnimplementedPartnerAdministrationServiceServer) SearchSingleSignOnUser(context.Context, *v1.SearchSingleSignOnConsoleUserRequest) (*v1.SearchSingleSignOnConsoleUserResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SearchSingleSignOnUser not implemented")
}
func (UnimplementedPartnerAdministrationServiceServer) GetSingleSignOnConsoleUser(context.Context, *v1.GetSingleSignOnConsoleUserRequest) (*v1.GetSingleSignOnConsoleUserResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetSingleSignOnConsoleUser not implemented")
}
func (UnimplementedPartnerAdministrationServiceServer) PingExternal(context.Context, *pingv1.PingRequest) (*pingv1.PingResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method PingExternal not implemented")
}
func (UnimplementedPartnerAdministrationServiceServer) mustEmbedUnimplementedPartnerAdministrationServiceServer() {
}

// UnsafePartnerAdministrationServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to PartnerAdministrationServiceServer will
// result in compilation errors.
type UnsafePartnerAdministrationServiceServer interface {
	mustEmbedUnimplementedPartnerAdministrationServiceServer()
}

func RegisterPartnerAdministrationServiceServer(s grpc.ServiceRegistrar, srv PartnerAdministrationServiceServer) {
	s.RegisterService(&PartnerAdministrationService_ServiceDesc, srv)
}

func _PartnerAdministrationService_GenerateMTLSCertificate_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(v1.GenerateMTLSCertificateRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PartnerAdministrationServiceServer).GenerateMTLSCertificate(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/partner.administration.v1.PartnerAdministrationService/GenerateMTLSCertificate",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PartnerAdministrationServiceServer).GenerateMTLSCertificate(ctx, req.(*v1.GenerateMTLSCertificateRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _PartnerAdministrationService_UpdateMTLSCertificate_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(v1.UpdateMTLSCertificateRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PartnerAdministrationServiceServer).UpdateMTLSCertificate(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/partner.administration.v1.PartnerAdministrationService/UpdateMTLSCertificate",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PartnerAdministrationServiceServer).UpdateMTLSCertificate(ctx, req.(*v1.UpdateMTLSCertificateRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _PartnerAdministrationService_SearchMTLSCertificate_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(v1.SearchMTLSCertificateRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PartnerAdministrationServiceServer).SearchMTLSCertificate(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/partner.administration.v1.PartnerAdministrationService/SearchMTLSCertificate",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PartnerAdministrationServiceServer).SearchMTLSCertificate(ctx, req.(*v1.SearchMTLSCertificateRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _PartnerAdministrationService_GetMTLSCertificate_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(v1.GetMTLSCertificateRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PartnerAdministrationServiceServer).GetMTLSCertificate(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/partner.administration.v1.PartnerAdministrationService/GetMTLSCertificate",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PartnerAdministrationServiceServer).GetMTLSCertificate(ctx, req.(*v1.GetMTLSCertificateRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _PartnerAdministrationService_CreateSingleSignOnConsoleUser_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(v1.CreateSingleSignOnConsoleUserRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PartnerAdministrationServiceServer).CreateSingleSignOnConsoleUser(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/partner.administration.v1.PartnerAdministrationService/CreateSingleSignOnConsoleUser",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PartnerAdministrationServiceServer).CreateSingleSignOnConsoleUser(ctx, req.(*v1.CreateSingleSignOnConsoleUserRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _PartnerAdministrationService_UpdateSingleSignOnConsoleUser_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(v1.UpdateSingleSignOnConsoleUserRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PartnerAdministrationServiceServer).UpdateSingleSignOnConsoleUser(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/partner.administration.v1.PartnerAdministrationService/UpdateSingleSignOnConsoleUser",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PartnerAdministrationServiceServer).UpdateSingleSignOnConsoleUser(ctx, req.(*v1.UpdateSingleSignOnConsoleUserRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _PartnerAdministrationService_SearchSingleSignOnUser_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(v1.SearchSingleSignOnConsoleUserRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PartnerAdministrationServiceServer).SearchSingleSignOnUser(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/partner.administration.v1.PartnerAdministrationService/SearchSingleSignOnUser",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PartnerAdministrationServiceServer).SearchSingleSignOnUser(ctx, req.(*v1.SearchSingleSignOnConsoleUserRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _PartnerAdministrationService_GetSingleSignOnConsoleUser_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(v1.GetSingleSignOnConsoleUserRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PartnerAdministrationServiceServer).GetSingleSignOnConsoleUser(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/partner.administration.v1.PartnerAdministrationService/GetSingleSignOnConsoleUser",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PartnerAdministrationServiceServer).GetSingleSignOnConsoleUser(ctx, req.(*v1.GetSingleSignOnConsoleUserRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _PartnerAdministrationService_PingExternal_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(pingv1.PingRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PartnerAdministrationServiceServer).PingExternal(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/partner.administration.v1.PartnerAdministrationService/PingExternal",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PartnerAdministrationServiceServer).PingExternal(ctx, req.(*pingv1.PingRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// PartnerAdministrationService_ServiceDesc is the grpc.ServiceDesc for PartnerAdministrationService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var PartnerAdministrationService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "partner.administration.v1.PartnerAdministrationService",
	HandlerType: (*PartnerAdministrationServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GenerateMTLSCertificate",
			Handler:    _PartnerAdministrationService_GenerateMTLSCertificate_Handler,
		},
		{
			MethodName: "UpdateMTLSCertificate",
			Handler:    _PartnerAdministrationService_UpdateMTLSCertificate_Handler,
		},
		{
			MethodName: "SearchMTLSCertificate",
			Handler:    _PartnerAdministrationService_SearchMTLSCertificate_Handler,
		},
		{
			MethodName: "GetMTLSCertificate",
			Handler:    _PartnerAdministrationService_GetMTLSCertificate_Handler,
		},
		{
			MethodName: "CreateSingleSignOnConsoleUser",
			Handler:    _PartnerAdministrationService_CreateSingleSignOnConsoleUser_Handler,
		},
		{
			MethodName: "UpdateSingleSignOnConsoleUser",
			Handler:    _PartnerAdministrationService_UpdateSingleSignOnConsoleUser_Handler,
		},
		{
			MethodName: "SearchSingleSignOnUser",
			Handler:    _PartnerAdministrationService_SearchSingleSignOnUser_Handler,
		},
		{
			MethodName: "GetSingleSignOnConsoleUser",
			Handler:    _PartnerAdministrationService_GetSingleSignOnConsoleUser_Handler,
		},
		{
			MethodName: "PingExternal",
			Handler:    _PartnerAdministrationService_PingExternal_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "partner/administration/v1/admin_service.proto",
}
