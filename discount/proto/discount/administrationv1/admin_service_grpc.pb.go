// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.21.8
// source: discount/administration/v1/admin_service.proto

package administrationv1

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	pingv1 "github.com/1080network/golang/discount/proto/common/pingv1"
	v1 "github.com/1080network/golang/discount/proto/common/v1"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

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
	GenerateClientMTLSCertificate(ctx context.Context, in *v1.GenerateClientMTLSCertificateRequest, opts ...grpc.CallOption) (*v1.GenerateClientMTLSCertificateResponse, error)
	UpdateClientMTLSCertificate(ctx context.Context, in *v1.UpdateClientMTLSCertificateRequest, opts ...grpc.CallOption) (*v1.UpdateClientMTLSCertificateResponse, error)
	SearchClientMTLSCertificate(ctx context.Context, in *v1.SearchClientMTLSCertificateRequest, opts ...grpc.CallOption) (*v1.SearchClientMTLSCertificateResponse, error)
	// if the provider cannot enable certificate authentication this can configure
	ConfigureApiTokenAuthentication(ctx context.Context, in *v1.ApiTokenConfigurationRequest, opts ...grpc.CallOption) (*v1.ApiTokenConfigurationResponse, error)
	// retrieves the api token configuration including the token itself
	GetApiTokenConfiguration(ctx context.Context, in *v1.GetApiTokenConfigurationRequest, opts ...grpc.CallOption) (*v1.GetApiTokenConfigurationResponse, error)
	// Update the configutation of the token, does not include the token itself
	UpdateAPITokenAuthenticationConfiguration(ctx context.Context, in *v1.UpdateApiTokenConfigurationRequest, opts ...grpc.CallOption) (*v1.UpdateApiTokenConfigurationResponse, error)
	// Updates the values for either the token or the ca certificate or both
	UpdateApiTokenValue(ctx context.Context, in *v1.UpdateApiTokenValueRequest, opts ...grpc.CallOption) (*v1.UpdateApiTokenValueResponse, error)
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
	err := c.cc.Invoke(ctx, "/discount.administration.v1.DiscountAdministrationService/GenerateMTLSCertificate", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *discountAdministrationServiceClient) UpdateMTLSCertificate(ctx context.Context, in *v1.UpdateMTLSCertificateRequest, opts ...grpc.CallOption) (*v1.UpdateMTLSCertificateResponse, error) {
	out := new(v1.UpdateMTLSCertificateResponse)
	err := c.cc.Invoke(ctx, "/discount.administration.v1.DiscountAdministrationService/UpdateMTLSCertificate", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *discountAdministrationServiceClient) SearchMTLSCertificate(ctx context.Context, in *v1.SearchMTLSCertificateRequest, opts ...grpc.CallOption) (*v1.SearchMTLSCertificateResponse, error) {
	out := new(v1.SearchMTLSCertificateResponse)
	err := c.cc.Invoke(ctx, "/discount.administration.v1.DiscountAdministrationService/SearchMTLSCertificate", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *discountAdministrationServiceClient) GetMTLSCertificate(ctx context.Context, in *v1.GetMTLSCertificateRequest, opts ...grpc.CallOption) (*v1.GetMTLSCertificateResponse, error) {
	out := new(v1.GetMTLSCertificateResponse)
	err := c.cc.Invoke(ctx, "/discount.administration.v1.DiscountAdministrationService/GetMTLSCertificate", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *discountAdministrationServiceClient) GenerateClientMTLSCertificate(ctx context.Context, in *v1.GenerateClientMTLSCertificateRequest, opts ...grpc.CallOption) (*v1.GenerateClientMTLSCertificateResponse, error) {
	out := new(v1.GenerateClientMTLSCertificateResponse)
	err := c.cc.Invoke(ctx, "/discount.administration.v1.DiscountAdministrationService/GenerateClientMTLSCertificate", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *discountAdministrationServiceClient) UpdateClientMTLSCertificate(ctx context.Context, in *v1.UpdateClientMTLSCertificateRequest, opts ...grpc.CallOption) (*v1.UpdateClientMTLSCertificateResponse, error) {
	out := new(v1.UpdateClientMTLSCertificateResponse)
	err := c.cc.Invoke(ctx, "/discount.administration.v1.DiscountAdministrationService/UpdateClientMTLSCertificate", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *discountAdministrationServiceClient) SearchClientMTLSCertificate(ctx context.Context, in *v1.SearchClientMTLSCertificateRequest, opts ...grpc.CallOption) (*v1.SearchClientMTLSCertificateResponse, error) {
	out := new(v1.SearchClientMTLSCertificateResponse)
	err := c.cc.Invoke(ctx, "/discount.administration.v1.DiscountAdministrationService/SearchClientMTLSCertificate", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *discountAdministrationServiceClient) ConfigureApiTokenAuthentication(ctx context.Context, in *v1.ApiTokenConfigurationRequest, opts ...grpc.CallOption) (*v1.ApiTokenConfigurationResponse, error) {
	out := new(v1.ApiTokenConfigurationResponse)
	err := c.cc.Invoke(ctx, "/discount.administration.v1.DiscountAdministrationService/ConfigureApiTokenAuthentication", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *discountAdministrationServiceClient) GetApiTokenConfiguration(ctx context.Context, in *v1.GetApiTokenConfigurationRequest, opts ...grpc.CallOption) (*v1.GetApiTokenConfigurationResponse, error) {
	out := new(v1.GetApiTokenConfigurationResponse)
	err := c.cc.Invoke(ctx, "/discount.administration.v1.DiscountAdministrationService/GetApiTokenConfiguration", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *discountAdministrationServiceClient) UpdateAPITokenAuthenticationConfiguration(ctx context.Context, in *v1.UpdateApiTokenConfigurationRequest, opts ...grpc.CallOption) (*v1.UpdateApiTokenConfigurationResponse, error) {
	out := new(v1.UpdateApiTokenConfigurationResponse)
	err := c.cc.Invoke(ctx, "/discount.administration.v1.DiscountAdministrationService/UpdateAPITokenAuthenticationConfiguration", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *discountAdministrationServiceClient) UpdateApiTokenValue(ctx context.Context, in *v1.UpdateApiTokenValueRequest, opts ...grpc.CallOption) (*v1.UpdateApiTokenValueResponse, error) {
	out := new(v1.UpdateApiTokenValueResponse)
	err := c.cc.Invoke(ctx, "/discount.administration.v1.DiscountAdministrationService/UpdateApiTokenValue", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *discountAdministrationServiceClient) PingExternal(ctx context.Context, in *pingv1.PingRequest, opts ...grpc.CallOption) (*pingv1.PingResponse, error) {
	out := new(pingv1.PingResponse)
	err := c.cc.Invoke(ctx, "/discount.administration.v1.DiscountAdministrationService/PingExternal", in, out, opts...)
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
	GenerateClientMTLSCertificate(context.Context, *v1.GenerateClientMTLSCertificateRequest) (*v1.GenerateClientMTLSCertificateResponse, error)
	UpdateClientMTLSCertificate(context.Context, *v1.UpdateClientMTLSCertificateRequest) (*v1.UpdateClientMTLSCertificateResponse, error)
	SearchClientMTLSCertificate(context.Context, *v1.SearchClientMTLSCertificateRequest) (*v1.SearchClientMTLSCertificateResponse, error)
	// if the provider cannot enable certificate authentication this can configure
	ConfigureApiTokenAuthentication(context.Context, *v1.ApiTokenConfigurationRequest) (*v1.ApiTokenConfigurationResponse, error)
	// retrieves the api token configuration including the token itself
	GetApiTokenConfiguration(context.Context, *v1.GetApiTokenConfigurationRequest) (*v1.GetApiTokenConfigurationResponse, error)
	// Update the configutation of the token, does not include the token itself
	UpdateAPITokenAuthenticationConfiguration(context.Context, *v1.UpdateApiTokenConfigurationRequest) (*v1.UpdateApiTokenConfigurationResponse, error)
	// Updates the values for either the token or the ca certificate or both
	UpdateApiTokenValue(context.Context, *v1.UpdateApiTokenValueRequest) (*v1.UpdateApiTokenValueResponse, error)
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
func (UnimplementedDiscountAdministrationServiceServer) GenerateClientMTLSCertificate(context.Context, *v1.GenerateClientMTLSCertificateRequest) (*v1.GenerateClientMTLSCertificateResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GenerateClientMTLSCertificate not implemented")
}
func (UnimplementedDiscountAdministrationServiceServer) UpdateClientMTLSCertificate(context.Context, *v1.UpdateClientMTLSCertificateRequest) (*v1.UpdateClientMTLSCertificateResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateClientMTLSCertificate not implemented")
}
func (UnimplementedDiscountAdministrationServiceServer) SearchClientMTLSCertificate(context.Context, *v1.SearchClientMTLSCertificateRequest) (*v1.SearchClientMTLSCertificateResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SearchClientMTLSCertificate not implemented")
}
func (UnimplementedDiscountAdministrationServiceServer) ConfigureApiTokenAuthentication(context.Context, *v1.ApiTokenConfigurationRequest) (*v1.ApiTokenConfigurationResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ConfigureApiTokenAuthentication not implemented")
}
func (UnimplementedDiscountAdministrationServiceServer) GetApiTokenConfiguration(context.Context, *v1.GetApiTokenConfigurationRequest) (*v1.GetApiTokenConfigurationResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetApiTokenConfiguration not implemented")
}
func (UnimplementedDiscountAdministrationServiceServer) UpdateAPITokenAuthenticationConfiguration(context.Context, *v1.UpdateApiTokenConfigurationRequest) (*v1.UpdateApiTokenConfigurationResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateAPITokenAuthenticationConfiguration not implemented")
}
func (UnimplementedDiscountAdministrationServiceServer) UpdateApiTokenValue(context.Context, *v1.UpdateApiTokenValueRequest) (*v1.UpdateApiTokenValueResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateApiTokenValue not implemented")
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
		FullMethod: "/discount.administration.v1.DiscountAdministrationService/GenerateMTLSCertificate",
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
		FullMethod: "/discount.administration.v1.DiscountAdministrationService/UpdateMTLSCertificate",
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
		FullMethod: "/discount.administration.v1.DiscountAdministrationService/SearchMTLSCertificate",
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
		FullMethod: "/discount.administration.v1.DiscountAdministrationService/GetMTLSCertificate",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DiscountAdministrationServiceServer).GetMTLSCertificate(ctx, req.(*v1.GetMTLSCertificateRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _DiscountAdministrationService_GenerateClientMTLSCertificate_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(v1.GenerateClientMTLSCertificateRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DiscountAdministrationServiceServer).GenerateClientMTLSCertificate(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/discount.administration.v1.DiscountAdministrationService/GenerateClientMTLSCertificate",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DiscountAdministrationServiceServer).GenerateClientMTLSCertificate(ctx, req.(*v1.GenerateClientMTLSCertificateRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _DiscountAdministrationService_UpdateClientMTLSCertificate_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(v1.UpdateClientMTLSCertificateRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DiscountAdministrationServiceServer).UpdateClientMTLSCertificate(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/discount.administration.v1.DiscountAdministrationService/UpdateClientMTLSCertificate",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DiscountAdministrationServiceServer).UpdateClientMTLSCertificate(ctx, req.(*v1.UpdateClientMTLSCertificateRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _DiscountAdministrationService_SearchClientMTLSCertificate_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(v1.SearchClientMTLSCertificateRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DiscountAdministrationServiceServer).SearchClientMTLSCertificate(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/discount.administration.v1.DiscountAdministrationService/SearchClientMTLSCertificate",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DiscountAdministrationServiceServer).SearchClientMTLSCertificate(ctx, req.(*v1.SearchClientMTLSCertificateRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _DiscountAdministrationService_ConfigureApiTokenAuthentication_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(v1.ApiTokenConfigurationRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DiscountAdministrationServiceServer).ConfigureApiTokenAuthentication(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/discount.administration.v1.DiscountAdministrationService/ConfigureApiTokenAuthentication",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DiscountAdministrationServiceServer).ConfigureApiTokenAuthentication(ctx, req.(*v1.ApiTokenConfigurationRequest))
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
		FullMethod: "/discount.administration.v1.DiscountAdministrationService/GetApiTokenConfiguration",
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
		FullMethod: "/discount.administration.v1.DiscountAdministrationService/UpdateAPITokenAuthenticationConfiguration",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DiscountAdministrationServiceServer).UpdateAPITokenAuthenticationConfiguration(ctx, req.(*v1.UpdateApiTokenConfigurationRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _DiscountAdministrationService_UpdateApiTokenValue_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(v1.UpdateApiTokenValueRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DiscountAdministrationServiceServer).UpdateApiTokenValue(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/discount.administration.v1.DiscountAdministrationService/UpdateApiTokenValue",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DiscountAdministrationServiceServer).UpdateApiTokenValue(ctx, req.(*v1.UpdateApiTokenValueRequest))
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
		FullMethod: "/discount.administration.v1.DiscountAdministrationService/PingExternal",
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
	ServiceName: "discount.administration.v1.DiscountAdministrationService",
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
			MethodName: "GenerateClientMTLSCertificate",
			Handler:    _DiscountAdministrationService_GenerateClientMTLSCertificate_Handler,
		},
		{
			MethodName: "UpdateClientMTLSCertificate",
			Handler:    _DiscountAdministrationService_UpdateClientMTLSCertificate_Handler,
		},
		{
			MethodName: "SearchClientMTLSCertificate",
			Handler:    _DiscountAdministrationService_SearchClientMTLSCertificate_Handler,
		},
		{
			MethodName: "ConfigureApiTokenAuthentication",
			Handler:    _DiscountAdministrationService_ConfigureApiTokenAuthentication_Handler,
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
			MethodName: "UpdateApiTokenValue",
			Handler:    _DiscountAdministrationService_UpdateApiTokenValue_Handler,
		},
		{
			MethodName: "PingExternal",
			Handler:    _DiscountAdministrationService_PingExternal_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "discount/administration/v1/admin_service.proto",
}
