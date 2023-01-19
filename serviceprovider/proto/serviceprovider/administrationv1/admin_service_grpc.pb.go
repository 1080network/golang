// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.21.8
// source: serviceprovider/administration/v1/admin_service.proto

package administrationv1

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	pingv1 "github.com/1080network/golang/serviceprovider/proto/common/pingv1"
	v1 "github.com/1080network/golang/serviceprovider/proto/common/v1"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

// ServiceProviderAdministrationServiceClient is the client API for ServiceProviderAdministrationService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type ServiceProviderAdministrationServiceClient interface {
	// Generate a new MTLS certificate.
	GenerateMTLSCertificate(ctx context.Context, in *v1.GenerateMTLSCertificateRequest, opts ...grpc.CallOption) (*v1.GenerateMTLSCertificateResponse, error)
	// Update the certificate with a given serial number, only supports enable/disable for now
	UpdateMTLSCertificate(ctx context.Context, in *v1.UpdateMTLSCertificateRequest, opts ...grpc.CallOption) (*v1.UpdateMTLSCertificateResponse, error)
	// Search for certificates and return the ones that match the criteria provided
	SearchMTLSCertificate(ctx context.Context, in *v1.SearchMTLSCertificateRequest, opts ...grpc.CallOption) (*v1.SearchMTLSCertificateResponse, error)
	GetMTLSCertificate(ctx context.Context, in *v1.GetMTLSCertificateRequest, opts ...grpc.CallOption) (*v1.GetMTLSCertificateResponse, error)
	CreateSingleSignOnConsoleUser(ctx context.Context, in *v1.CreateSingleSignOnConsoleUserRequest, opts ...grpc.CallOption) (*v1.CreateSingleSignOnConsoleUserResponse, error)
	UpdateSingleSignOnConsoleUser(ctx context.Context, in *v1.UpdateSingleSignOnConsoleUserRequest, opts ...grpc.CallOption) (*v1.UpdateSingleSignOnConsoleUserResponse, error)
	SearchSingleSignOnUser(ctx context.Context, in *v1.SearchSingleSignOnConsoleUserRequest, opts ...grpc.CallOption) (*v1.SearchSingleSignOnConsoleUserResponse, error)
	GetSingleSignOnConsoleUser(ctx context.Context, in *v1.GetSingleSignOnConsoleUserRequest, opts ...grpc.CallOption) (*v1.GetSingleSignOnConsoleUserResponse, error)
	// External authentication mechanisms for Mica to call provider endpoints
	GetExternalClientSettings(ctx context.Context, in *v1.GetExternalClientSettingsRequest, opts ...grpc.CallOption) (*v1.GetExternalClientSettingsResponse, error)
	UpdateExternalClientCallbackAddress(ctx context.Context, in *v1.UpdateExternalClientCallBackAddressRequest, opts ...grpc.CallOption) (*v1.UpdateExternalClientCallBackAddressResponse, error)
	// Client certificates are used when mica needs to call out to a customers environment.
	GenerateExternalClientMTLSCertificate(ctx context.Context, in *v1.GenerateExternalClientMTLSCertificateRequest, opts ...grpc.CallOption) (*v1.GenerateExternalClientMTLSCertificateResponse, error)
	UpdateExternalClientMTLSCertificate(ctx context.Context, in *v1.UpdateExternalClientMTLSCertificateRequest, opts ...grpc.CallOption) (*v1.UpdateExternalClientMTLSCertificateResponse, error)
	SearchExternalClientMTLSCertificate(ctx context.Context, in *v1.SearchExternalClientMTLSCertificateRequest, opts ...grpc.CallOption) (*v1.SearchExternalClientMTLSCertificateResponse, error)
	// tests the external call to verify proper configuration and connectivity
	PingExternal(ctx context.Context, in *pingv1.PingRequest, opts ...grpc.CallOption) (*pingv1.PingResponse, error)
}

type serviceProviderAdministrationServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewServiceProviderAdministrationServiceClient(cc grpc.ClientConnInterface) ServiceProviderAdministrationServiceClient {
	return &serviceProviderAdministrationServiceClient{cc}
}

func (c *serviceProviderAdministrationServiceClient) GenerateMTLSCertificate(ctx context.Context, in *v1.GenerateMTLSCertificateRequest, opts ...grpc.CallOption) (*v1.GenerateMTLSCertificateResponse, error) {
	out := new(v1.GenerateMTLSCertificateResponse)
	err := c.cc.Invoke(ctx, "/serviceprovider.administration.v1.ServiceProviderAdministrationService/GenerateMTLSCertificate", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *serviceProviderAdministrationServiceClient) UpdateMTLSCertificate(ctx context.Context, in *v1.UpdateMTLSCertificateRequest, opts ...grpc.CallOption) (*v1.UpdateMTLSCertificateResponse, error) {
	out := new(v1.UpdateMTLSCertificateResponse)
	err := c.cc.Invoke(ctx, "/serviceprovider.administration.v1.ServiceProviderAdministrationService/UpdateMTLSCertificate", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *serviceProviderAdministrationServiceClient) SearchMTLSCertificate(ctx context.Context, in *v1.SearchMTLSCertificateRequest, opts ...grpc.CallOption) (*v1.SearchMTLSCertificateResponse, error) {
	out := new(v1.SearchMTLSCertificateResponse)
	err := c.cc.Invoke(ctx, "/serviceprovider.administration.v1.ServiceProviderAdministrationService/SearchMTLSCertificate", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *serviceProviderAdministrationServiceClient) GetMTLSCertificate(ctx context.Context, in *v1.GetMTLSCertificateRequest, opts ...grpc.CallOption) (*v1.GetMTLSCertificateResponse, error) {
	out := new(v1.GetMTLSCertificateResponse)
	err := c.cc.Invoke(ctx, "/serviceprovider.administration.v1.ServiceProviderAdministrationService/GetMTLSCertificate", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *serviceProviderAdministrationServiceClient) CreateSingleSignOnConsoleUser(ctx context.Context, in *v1.CreateSingleSignOnConsoleUserRequest, opts ...grpc.CallOption) (*v1.CreateSingleSignOnConsoleUserResponse, error) {
	out := new(v1.CreateSingleSignOnConsoleUserResponse)
	err := c.cc.Invoke(ctx, "/serviceprovider.administration.v1.ServiceProviderAdministrationService/CreateSingleSignOnConsoleUser", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *serviceProviderAdministrationServiceClient) UpdateSingleSignOnConsoleUser(ctx context.Context, in *v1.UpdateSingleSignOnConsoleUserRequest, opts ...grpc.CallOption) (*v1.UpdateSingleSignOnConsoleUserResponse, error) {
	out := new(v1.UpdateSingleSignOnConsoleUserResponse)
	err := c.cc.Invoke(ctx, "/serviceprovider.administration.v1.ServiceProviderAdministrationService/UpdateSingleSignOnConsoleUser", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *serviceProviderAdministrationServiceClient) SearchSingleSignOnUser(ctx context.Context, in *v1.SearchSingleSignOnConsoleUserRequest, opts ...grpc.CallOption) (*v1.SearchSingleSignOnConsoleUserResponse, error) {
	out := new(v1.SearchSingleSignOnConsoleUserResponse)
	err := c.cc.Invoke(ctx, "/serviceprovider.administration.v1.ServiceProviderAdministrationService/SearchSingleSignOnUser", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *serviceProviderAdministrationServiceClient) GetSingleSignOnConsoleUser(ctx context.Context, in *v1.GetSingleSignOnConsoleUserRequest, opts ...grpc.CallOption) (*v1.GetSingleSignOnConsoleUserResponse, error) {
	out := new(v1.GetSingleSignOnConsoleUserResponse)
	err := c.cc.Invoke(ctx, "/serviceprovider.administration.v1.ServiceProviderAdministrationService/GetSingleSignOnConsoleUser", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *serviceProviderAdministrationServiceClient) GetExternalClientSettings(ctx context.Context, in *v1.GetExternalClientSettingsRequest, opts ...grpc.CallOption) (*v1.GetExternalClientSettingsResponse, error) {
	out := new(v1.GetExternalClientSettingsResponse)
	err := c.cc.Invoke(ctx, "/serviceprovider.administration.v1.ServiceProviderAdministrationService/GetExternalClientSettings", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *serviceProviderAdministrationServiceClient) UpdateExternalClientCallbackAddress(ctx context.Context, in *v1.UpdateExternalClientCallBackAddressRequest, opts ...grpc.CallOption) (*v1.UpdateExternalClientCallBackAddressResponse, error) {
	out := new(v1.UpdateExternalClientCallBackAddressResponse)
	err := c.cc.Invoke(ctx, "/serviceprovider.administration.v1.ServiceProviderAdministrationService/UpdateExternalClientCallbackAddress", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *serviceProviderAdministrationServiceClient) GenerateExternalClientMTLSCertificate(ctx context.Context, in *v1.GenerateExternalClientMTLSCertificateRequest, opts ...grpc.CallOption) (*v1.GenerateExternalClientMTLSCertificateResponse, error) {
	out := new(v1.GenerateExternalClientMTLSCertificateResponse)
	err := c.cc.Invoke(ctx, "/serviceprovider.administration.v1.ServiceProviderAdministrationService/GenerateExternalClientMTLSCertificate", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *serviceProviderAdministrationServiceClient) UpdateExternalClientMTLSCertificate(ctx context.Context, in *v1.UpdateExternalClientMTLSCertificateRequest, opts ...grpc.CallOption) (*v1.UpdateExternalClientMTLSCertificateResponse, error) {
	out := new(v1.UpdateExternalClientMTLSCertificateResponse)
	err := c.cc.Invoke(ctx, "/serviceprovider.administration.v1.ServiceProviderAdministrationService/UpdateExternalClientMTLSCertificate", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *serviceProviderAdministrationServiceClient) SearchExternalClientMTLSCertificate(ctx context.Context, in *v1.SearchExternalClientMTLSCertificateRequest, opts ...grpc.CallOption) (*v1.SearchExternalClientMTLSCertificateResponse, error) {
	out := new(v1.SearchExternalClientMTLSCertificateResponse)
	err := c.cc.Invoke(ctx, "/serviceprovider.administration.v1.ServiceProviderAdministrationService/SearchExternalClientMTLSCertificate", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *serviceProviderAdministrationServiceClient) PingExternal(ctx context.Context, in *pingv1.PingRequest, opts ...grpc.CallOption) (*pingv1.PingResponse, error) {
	out := new(pingv1.PingResponse)
	err := c.cc.Invoke(ctx, "/serviceprovider.administration.v1.ServiceProviderAdministrationService/PingExternal", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ServiceProviderAdministrationServiceServer is the server API for ServiceProviderAdministrationService service.
// All implementations must embed UnimplementedServiceProviderAdministrationServiceServer
// for forward compatibility
type ServiceProviderAdministrationServiceServer interface {
	// Generate a new MTLS certificate.
	GenerateMTLSCertificate(context.Context, *v1.GenerateMTLSCertificateRequest) (*v1.GenerateMTLSCertificateResponse, error)
	// Update the certificate with a given serial number, only supports enable/disable for now
	UpdateMTLSCertificate(context.Context, *v1.UpdateMTLSCertificateRequest) (*v1.UpdateMTLSCertificateResponse, error)
	// Search for certificates and return the ones that match the criteria provided
	SearchMTLSCertificate(context.Context, *v1.SearchMTLSCertificateRequest) (*v1.SearchMTLSCertificateResponse, error)
	GetMTLSCertificate(context.Context, *v1.GetMTLSCertificateRequest) (*v1.GetMTLSCertificateResponse, error)
	CreateSingleSignOnConsoleUser(context.Context, *v1.CreateSingleSignOnConsoleUserRequest) (*v1.CreateSingleSignOnConsoleUserResponse, error)
	UpdateSingleSignOnConsoleUser(context.Context, *v1.UpdateSingleSignOnConsoleUserRequest) (*v1.UpdateSingleSignOnConsoleUserResponse, error)
	SearchSingleSignOnUser(context.Context, *v1.SearchSingleSignOnConsoleUserRequest) (*v1.SearchSingleSignOnConsoleUserResponse, error)
	GetSingleSignOnConsoleUser(context.Context, *v1.GetSingleSignOnConsoleUserRequest) (*v1.GetSingleSignOnConsoleUserResponse, error)
	// External authentication mechanisms for Mica to call provider endpoints
	GetExternalClientSettings(context.Context, *v1.GetExternalClientSettingsRequest) (*v1.GetExternalClientSettingsResponse, error)
	UpdateExternalClientCallbackAddress(context.Context, *v1.UpdateExternalClientCallBackAddressRequest) (*v1.UpdateExternalClientCallBackAddressResponse, error)
	// Client certificates are used when mica needs to call out to a customers environment.
	GenerateExternalClientMTLSCertificate(context.Context, *v1.GenerateExternalClientMTLSCertificateRequest) (*v1.GenerateExternalClientMTLSCertificateResponse, error)
	UpdateExternalClientMTLSCertificate(context.Context, *v1.UpdateExternalClientMTLSCertificateRequest) (*v1.UpdateExternalClientMTLSCertificateResponse, error)
	SearchExternalClientMTLSCertificate(context.Context, *v1.SearchExternalClientMTLSCertificateRequest) (*v1.SearchExternalClientMTLSCertificateResponse, error)
	// tests the external call to verify proper configuration and connectivity
	PingExternal(context.Context, *pingv1.PingRequest) (*pingv1.PingResponse, error)
	mustEmbedUnimplementedServiceProviderAdministrationServiceServer()
}

// UnimplementedServiceProviderAdministrationServiceServer must be embedded to have forward compatible implementations.
type UnimplementedServiceProviderAdministrationServiceServer struct {
}

func (UnimplementedServiceProviderAdministrationServiceServer) GenerateMTLSCertificate(context.Context, *v1.GenerateMTLSCertificateRequest) (*v1.GenerateMTLSCertificateResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GenerateMTLSCertificate not implemented")
}
func (UnimplementedServiceProviderAdministrationServiceServer) UpdateMTLSCertificate(context.Context, *v1.UpdateMTLSCertificateRequest) (*v1.UpdateMTLSCertificateResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateMTLSCertificate not implemented")
}
func (UnimplementedServiceProviderAdministrationServiceServer) SearchMTLSCertificate(context.Context, *v1.SearchMTLSCertificateRequest) (*v1.SearchMTLSCertificateResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SearchMTLSCertificate not implemented")
}
func (UnimplementedServiceProviderAdministrationServiceServer) GetMTLSCertificate(context.Context, *v1.GetMTLSCertificateRequest) (*v1.GetMTLSCertificateResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetMTLSCertificate not implemented")
}
func (UnimplementedServiceProviderAdministrationServiceServer) CreateSingleSignOnConsoleUser(context.Context, *v1.CreateSingleSignOnConsoleUserRequest) (*v1.CreateSingleSignOnConsoleUserResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateSingleSignOnConsoleUser not implemented")
}
func (UnimplementedServiceProviderAdministrationServiceServer) UpdateSingleSignOnConsoleUser(context.Context, *v1.UpdateSingleSignOnConsoleUserRequest) (*v1.UpdateSingleSignOnConsoleUserResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateSingleSignOnConsoleUser not implemented")
}
func (UnimplementedServiceProviderAdministrationServiceServer) SearchSingleSignOnUser(context.Context, *v1.SearchSingleSignOnConsoleUserRequest) (*v1.SearchSingleSignOnConsoleUserResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SearchSingleSignOnUser not implemented")
}
func (UnimplementedServiceProviderAdministrationServiceServer) GetSingleSignOnConsoleUser(context.Context, *v1.GetSingleSignOnConsoleUserRequest) (*v1.GetSingleSignOnConsoleUserResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetSingleSignOnConsoleUser not implemented")
}
func (UnimplementedServiceProviderAdministrationServiceServer) GetExternalClientSettings(context.Context, *v1.GetExternalClientSettingsRequest) (*v1.GetExternalClientSettingsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetExternalClientSettings not implemented")
}
func (UnimplementedServiceProviderAdministrationServiceServer) UpdateExternalClientCallbackAddress(context.Context, *v1.UpdateExternalClientCallBackAddressRequest) (*v1.UpdateExternalClientCallBackAddressResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateExternalClientCallbackAddress not implemented")
}
func (UnimplementedServiceProviderAdministrationServiceServer) GenerateExternalClientMTLSCertificate(context.Context, *v1.GenerateExternalClientMTLSCertificateRequest) (*v1.GenerateExternalClientMTLSCertificateResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GenerateExternalClientMTLSCertificate not implemented")
}
func (UnimplementedServiceProviderAdministrationServiceServer) UpdateExternalClientMTLSCertificate(context.Context, *v1.UpdateExternalClientMTLSCertificateRequest) (*v1.UpdateExternalClientMTLSCertificateResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateExternalClientMTLSCertificate not implemented")
}
func (UnimplementedServiceProviderAdministrationServiceServer) SearchExternalClientMTLSCertificate(context.Context, *v1.SearchExternalClientMTLSCertificateRequest) (*v1.SearchExternalClientMTLSCertificateResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SearchExternalClientMTLSCertificate not implemented")
}
func (UnimplementedServiceProviderAdministrationServiceServer) PingExternal(context.Context, *pingv1.PingRequest) (*pingv1.PingResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method PingExternal not implemented")
}
func (UnimplementedServiceProviderAdministrationServiceServer) mustEmbedUnimplementedServiceProviderAdministrationServiceServer() {
}

// UnsafeServiceProviderAdministrationServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to ServiceProviderAdministrationServiceServer will
// result in compilation errors.
type UnsafeServiceProviderAdministrationServiceServer interface {
	mustEmbedUnimplementedServiceProviderAdministrationServiceServer()
}

func RegisterServiceProviderAdministrationServiceServer(s grpc.ServiceRegistrar, srv ServiceProviderAdministrationServiceServer) {
	s.RegisterService(&ServiceProviderAdministrationService_ServiceDesc, srv)
}

func _ServiceProviderAdministrationService_GenerateMTLSCertificate_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(v1.GenerateMTLSCertificateRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ServiceProviderAdministrationServiceServer).GenerateMTLSCertificate(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/serviceprovider.administration.v1.ServiceProviderAdministrationService/GenerateMTLSCertificate",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ServiceProviderAdministrationServiceServer).GenerateMTLSCertificate(ctx, req.(*v1.GenerateMTLSCertificateRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ServiceProviderAdministrationService_UpdateMTLSCertificate_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(v1.UpdateMTLSCertificateRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ServiceProviderAdministrationServiceServer).UpdateMTLSCertificate(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/serviceprovider.administration.v1.ServiceProviderAdministrationService/UpdateMTLSCertificate",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ServiceProviderAdministrationServiceServer).UpdateMTLSCertificate(ctx, req.(*v1.UpdateMTLSCertificateRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ServiceProviderAdministrationService_SearchMTLSCertificate_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(v1.SearchMTLSCertificateRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ServiceProviderAdministrationServiceServer).SearchMTLSCertificate(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/serviceprovider.administration.v1.ServiceProviderAdministrationService/SearchMTLSCertificate",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ServiceProviderAdministrationServiceServer).SearchMTLSCertificate(ctx, req.(*v1.SearchMTLSCertificateRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ServiceProviderAdministrationService_GetMTLSCertificate_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(v1.GetMTLSCertificateRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ServiceProviderAdministrationServiceServer).GetMTLSCertificate(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/serviceprovider.administration.v1.ServiceProviderAdministrationService/GetMTLSCertificate",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ServiceProviderAdministrationServiceServer).GetMTLSCertificate(ctx, req.(*v1.GetMTLSCertificateRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ServiceProviderAdministrationService_CreateSingleSignOnConsoleUser_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(v1.CreateSingleSignOnConsoleUserRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ServiceProviderAdministrationServiceServer).CreateSingleSignOnConsoleUser(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/serviceprovider.administration.v1.ServiceProviderAdministrationService/CreateSingleSignOnConsoleUser",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ServiceProviderAdministrationServiceServer).CreateSingleSignOnConsoleUser(ctx, req.(*v1.CreateSingleSignOnConsoleUserRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ServiceProviderAdministrationService_UpdateSingleSignOnConsoleUser_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(v1.UpdateSingleSignOnConsoleUserRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ServiceProviderAdministrationServiceServer).UpdateSingleSignOnConsoleUser(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/serviceprovider.administration.v1.ServiceProviderAdministrationService/UpdateSingleSignOnConsoleUser",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ServiceProviderAdministrationServiceServer).UpdateSingleSignOnConsoleUser(ctx, req.(*v1.UpdateSingleSignOnConsoleUserRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ServiceProviderAdministrationService_SearchSingleSignOnUser_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(v1.SearchSingleSignOnConsoleUserRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ServiceProviderAdministrationServiceServer).SearchSingleSignOnUser(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/serviceprovider.administration.v1.ServiceProviderAdministrationService/SearchSingleSignOnUser",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ServiceProviderAdministrationServiceServer).SearchSingleSignOnUser(ctx, req.(*v1.SearchSingleSignOnConsoleUserRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ServiceProviderAdministrationService_GetSingleSignOnConsoleUser_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(v1.GetSingleSignOnConsoleUserRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ServiceProviderAdministrationServiceServer).GetSingleSignOnConsoleUser(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/serviceprovider.administration.v1.ServiceProviderAdministrationService/GetSingleSignOnConsoleUser",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ServiceProviderAdministrationServiceServer).GetSingleSignOnConsoleUser(ctx, req.(*v1.GetSingleSignOnConsoleUserRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ServiceProviderAdministrationService_GetExternalClientSettings_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(v1.GetExternalClientSettingsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ServiceProviderAdministrationServiceServer).GetExternalClientSettings(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/serviceprovider.administration.v1.ServiceProviderAdministrationService/GetExternalClientSettings",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ServiceProviderAdministrationServiceServer).GetExternalClientSettings(ctx, req.(*v1.GetExternalClientSettingsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ServiceProviderAdministrationService_UpdateExternalClientCallbackAddress_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(v1.UpdateExternalClientCallBackAddressRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ServiceProviderAdministrationServiceServer).UpdateExternalClientCallbackAddress(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/serviceprovider.administration.v1.ServiceProviderAdministrationService/UpdateExternalClientCallbackAddress",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ServiceProviderAdministrationServiceServer).UpdateExternalClientCallbackAddress(ctx, req.(*v1.UpdateExternalClientCallBackAddressRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ServiceProviderAdministrationService_GenerateExternalClientMTLSCertificate_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(v1.GenerateExternalClientMTLSCertificateRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ServiceProviderAdministrationServiceServer).GenerateExternalClientMTLSCertificate(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/serviceprovider.administration.v1.ServiceProviderAdministrationService/GenerateExternalClientMTLSCertificate",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ServiceProviderAdministrationServiceServer).GenerateExternalClientMTLSCertificate(ctx, req.(*v1.GenerateExternalClientMTLSCertificateRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ServiceProviderAdministrationService_UpdateExternalClientMTLSCertificate_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(v1.UpdateExternalClientMTLSCertificateRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ServiceProviderAdministrationServiceServer).UpdateExternalClientMTLSCertificate(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/serviceprovider.administration.v1.ServiceProviderAdministrationService/UpdateExternalClientMTLSCertificate",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ServiceProviderAdministrationServiceServer).UpdateExternalClientMTLSCertificate(ctx, req.(*v1.UpdateExternalClientMTLSCertificateRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ServiceProviderAdministrationService_SearchExternalClientMTLSCertificate_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(v1.SearchExternalClientMTLSCertificateRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ServiceProviderAdministrationServiceServer).SearchExternalClientMTLSCertificate(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/serviceprovider.administration.v1.ServiceProviderAdministrationService/SearchExternalClientMTLSCertificate",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ServiceProviderAdministrationServiceServer).SearchExternalClientMTLSCertificate(ctx, req.(*v1.SearchExternalClientMTLSCertificateRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ServiceProviderAdministrationService_PingExternal_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(pingv1.PingRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ServiceProviderAdministrationServiceServer).PingExternal(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/serviceprovider.administration.v1.ServiceProviderAdministrationService/PingExternal",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ServiceProviderAdministrationServiceServer).PingExternal(ctx, req.(*pingv1.PingRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// ServiceProviderAdministrationService_ServiceDesc is the grpc.ServiceDesc for ServiceProviderAdministrationService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var ServiceProviderAdministrationService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "serviceprovider.administration.v1.ServiceProviderAdministrationService",
	HandlerType: (*ServiceProviderAdministrationServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GenerateMTLSCertificate",
			Handler:    _ServiceProviderAdministrationService_GenerateMTLSCertificate_Handler,
		},
		{
			MethodName: "UpdateMTLSCertificate",
			Handler:    _ServiceProviderAdministrationService_UpdateMTLSCertificate_Handler,
		},
		{
			MethodName: "SearchMTLSCertificate",
			Handler:    _ServiceProviderAdministrationService_SearchMTLSCertificate_Handler,
		},
		{
			MethodName: "GetMTLSCertificate",
			Handler:    _ServiceProviderAdministrationService_GetMTLSCertificate_Handler,
		},
		{
			MethodName: "CreateSingleSignOnConsoleUser",
			Handler:    _ServiceProviderAdministrationService_CreateSingleSignOnConsoleUser_Handler,
		},
		{
			MethodName: "UpdateSingleSignOnConsoleUser",
			Handler:    _ServiceProviderAdministrationService_UpdateSingleSignOnConsoleUser_Handler,
		},
		{
			MethodName: "SearchSingleSignOnUser",
			Handler:    _ServiceProviderAdministrationService_SearchSingleSignOnUser_Handler,
		},
		{
			MethodName: "GetSingleSignOnConsoleUser",
			Handler:    _ServiceProviderAdministrationService_GetSingleSignOnConsoleUser_Handler,
		},
		{
			MethodName: "GetExternalClientSettings",
			Handler:    _ServiceProviderAdministrationService_GetExternalClientSettings_Handler,
		},
		{
			MethodName: "UpdateExternalClientCallbackAddress",
			Handler:    _ServiceProviderAdministrationService_UpdateExternalClientCallbackAddress_Handler,
		},
		{
			MethodName: "GenerateExternalClientMTLSCertificate",
			Handler:    _ServiceProviderAdministrationService_GenerateExternalClientMTLSCertificate_Handler,
		},
		{
			MethodName: "UpdateExternalClientMTLSCertificate",
			Handler:    _ServiceProviderAdministrationService_UpdateExternalClientMTLSCertificate_Handler,
		},
		{
			MethodName: "SearchExternalClientMTLSCertificate",
			Handler:    _ServiceProviderAdministrationService_SearchExternalClientMTLSCertificate_Handler,
		},
		{
			MethodName: "PingExternal",
			Handler:    _ServiceProviderAdministrationService_PingExternal_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "serviceprovider/administration/v1/admin_service.proto",
}
