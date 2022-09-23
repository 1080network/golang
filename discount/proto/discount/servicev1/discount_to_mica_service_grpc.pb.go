// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.19.4
// source: discount/service/v1/discount_to_mica_service.proto

package servicev1

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	pingv1 "github.com/1080network/golang/discount/proto/common/pingv1"
	discountdefinitionv1 "github.com/1080network/golang/discount/proto/discount/discountdefinitionv1"
	discountproviderv1 "github.com/1080network/golang/discount/proto/discount/discountproviderv1"
	discountv1 "github.com/1080network/golang/discount/proto/discount/discountv1"
	productgroupv1 "github.com/1080network/golang/discount/proto/discount/productgroupv1"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

// DiscountToMicaServiceClient is the client API for DiscountToMicaService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type DiscountToMicaServiceClient interface {
	GetDiscountProvider(ctx context.Context, in *discountproviderv1.GetDiscountProviderRequest, opts ...grpc.CallOption) (*discountproviderv1.GetDiscountProviderResponse, error)
	UpdateDiscountProvider(ctx context.Context, in *discountproviderv1.UpdateDiscountProviderRequest, opts ...grpc.CallOption) (*discountproviderv1.UpdateDiscountProviderResponse, error)
	CreateDiscountDefinition(ctx context.Context, in *discountdefinitionv1.CreateDiscountDefinitionRequest, opts ...grpc.CallOption) (*discountdefinitionv1.CreateDiscountDefinitionResponse, error)
	GetDiscountDefinition(ctx context.Context, in *discountdefinitionv1.GetDiscountDefinitionRequest, opts ...grpc.CallOption) (*discountdefinitionv1.GetDiscountDefinitionResponse, error)
	UpdateDiscountDefinition(ctx context.Context, in *discountdefinitionv1.UpdateDiscountDefinitionRequest, opts ...grpc.CallOption) (*discountdefinitionv1.UpdateDiscountDefinitionResponse, error)
	RemoveDiscountDefinition(ctx context.Context, in *discountdefinitionv1.RemoveDiscountDefinitionRequest, opts ...grpc.CallOption) (*discountdefinitionv1.RemoveDiscountDefinitionResponse, error)
	CreateDiscount(ctx context.Context, in *discountv1.CreateDiscountRequest, opts ...grpc.CallOption) (*discountv1.CreateDiscountResponse, error)
	GetDiscount(ctx context.Context, in *discountv1.GetDiscountRequest, opts ...grpc.CallOption) (*discountv1.GetDiscountResponse, error)
	CreateProductGroup(ctx context.Context, in *productgroupv1.CreateProductGroupRequest, opts ...grpc.CallOption) (*productgroupv1.CreateProductGroupResponse, error)
	GetProductGroup(ctx context.Context, in *productgroupv1.GetProductGroupRequest, opts ...grpc.CallOption) (*productgroupv1.GetProductGroupResponse, error)
	UpdateProductGroup(ctx context.Context, in *productgroupv1.UpdateProductGroupRequest, opts ...grpc.CallOption) (*productgroupv1.UpdateProductGroupResponse, error)
	RemoveProductGroup(ctx context.Context, in *productgroupv1.RemoveProductGroupRequest, opts ...grpc.CallOption) (*productgroupv1.RemoveProductGroupResponse, error)
	// An operation to ping the server to ensure it's up and running and that the connection is good.
	Ping(ctx context.Context, in *pingv1.PingRequest, opts ...grpc.CallOption) (*pingv1.PingResponse, error)
	// An operation that triggers a ping to the Mica servervice and that in turn triggers a ping to the discount provider service.
	// This is a useful call to ensure that the connectivity and security between the mica an external services is working.
	PingWithRoundTrip(ctx context.Context, in *pingv1.PingRequest, opts ...grpc.CallOption) (*pingv1.PingResponse, error)
}

type discountToMicaServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewDiscountToMicaServiceClient(cc grpc.ClientConnInterface) DiscountToMicaServiceClient {
	return &discountToMicaServiceClient{cc}
}

func (c *discountToMicaServiceClient) GetDiscountProvider(ctx context.Context, in *discountproviderv1.GetDiscountProviderRequest, opts ...grpc.CallOption) (*discountproviderv1.GetDiscountProviderResponse, error) {
	out := new(discountproviderv1.GetDiscountProviderResponse)
	err := c.cc.Invoke(ctx, "/discount.service.v1.DiscountToMicaService/GetDiscountProvider", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *discountToMicaServiceClient) UpdateDiscountProvider(ctx context.Context, in *discountproviderv1.UpdateDiscountProviderRequest, opts ...grpc.CallOption) (*discountproviderv1.UpdateDiscountProviderResponse, error) {
	out := new(discountproviderv1.UpdateDiscountProviderResponse)
	err := c.cc.Invoke(ctx, "/discount.service.v1.DiscountToMicaService/UpdateDiscountProvider", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *discountToMicaServiceClient) CreateDiscountDefinition(ctx context.Context, in *discountdefinitionv1.CreateDiscountDefinitionRequest, opts ...grpc.CallOption) (*discountdefinitionv1.CreateDiscountDefinitionResponse, error) {
	out := new(discountdefinitionv1.CreateDiscountDefinitionResponse)
	err := c.cc.Invoke(ctx, "/discount.service.v1.DiscountToMicaService/CreateDiscountDefinition", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *discountToMicaServiceClient) GetDiscountDefinition(ctx context.Context, in *discountdefinitionv1.GetDiscountDefinitionRequest, opts ...grpc.CallOption) (*discountdefinitionv1.GetDiscountDefinitionResponse, error) {
	out := new(discountdefinitionv1.GetDiscountDefinitionResponse)
	err := c.cc.Invoke(ctx, "/discount.service.v1.DiscountToMicaService/GetDiscountDefinition", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *discountToMicaServiceClient) UpdateDiscountDefinition(ctx context.Context, in *discountdefinitionv1.UpdateDiscountDefinitionRequest, opts ...grpc.CallOption) (*discountdefinitionv1.UpdateDiscountDefinitionResponse, error) {
	out := new(discountdefinitionv1.UpdateDiscountDefinitionResponse)
	err := c.cc.Invoke(ctx, "/discount.service.v1.DiscountToMicaService/UpdateDiscountDefinition", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *discountToMicaServiceClient) RemoveDiscountDefinition(ctx context.Context, in *discountdefinitionv1.RemoveDiscountDefinitionRequest, opts ...grpc.CallOption) (*discountdefinitionv1.RemoveDiscountDefinitionResponse, error) {
	out := new(discountdefinitionv1.RemoveDiscountDefinitionResponse)
	err := c.cc.Invoke(ctx, "/discount.service.v1.DiscountToMicaService/RemoveDiscountDefinition", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *discountToMicaServiceClient) CreateDiscount(ctx context.Context, in *discountv1.CreateDiscountRequest, opts ...grpc.CallOption) (*discountv1.CreateDiscountResponse, error) {
	out := new(discountv1.CreateDiscountResponse)
	err := c.cc.Invoke(ctx, "/discount.service.v1.DiscountToMicaService/CreateDiscount", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *discountToMicaServiceClient) GetDiscount(ctx context.Context, in *discountv1.GetDiscountRequest, opts ...grpc.CallOption) (*discountv1.GetDiscountResponse, error) {
	out := new(discountv1.GetDiscountResponse)
	err := c.cc.Invoke(ctx, "/discount.service.v1.DiscountToMicaService/GetDiscount", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *discountToMicaServiceClient) CreateProductGroup(ctx context.Context, in *productgroupv1.CreateProductGroupRequest, opts ...grpc.CallOption) (*productgroupv1.CreateProductGroupResponse, error) {
	out := new(productgroupv1.CreateProductGroupResponse)
	err := c.cc.Invoke(ctx, "/discount.service.v1.DiscountToMicaService/CreateProductGroup", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *discountToMicaServiceClient) GetProductGroup(ctx context.Context, in *productgroupv1.GetProductGroupRequest, opts ...grpc.CallOption) (*productgroupv1.GetProductGroupResponse, error) {
	out := new(productgroupv1.GetProductGroupResponse)
	err := c.cc.Invoke(ctx, "/discount.service.v1.DiscountToMicaService/GetProductGroup", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *discountToMicaServiceClient) UpdateProductGroup(ctx context.Context, in *productgroupv1.UpdateProductGroupRequest, opts ...grpc.CallOption) (*productgroupv1.UpdateProductGroupResponse, error) {
	out := new(productgroupv1.UpdateProductGroupResponse)
	err := c.cc.Invoke(ctx, "/discount.service.v1.DiscountToMicaService/UpdateProductGroup", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *discountToMicaServiceClient) RemoveProductGroup(ctx context.Context, in *productgroupv1.RemoveProductGroupRequest, opts ...grpc.CallOption) (*productgroupv1.RemoveProductGroupResponse, error) {
	out := new(productgroupv1.RemoveProductGroupResponse)
	err := c.cc.Invoke(ctx, "/discount.service.v1.DiscountToMicaService/RemoveProductGroup", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *discountToMicaServiceClient) Ping(ctx context.Context, in *pingv1.PingRequest, opts ...grpc.CallOption) (*pingv1.PingResponse, error) {
	out := new(pingv1.PingResponse)
	err := c.cc.Invoke(ctx, "/discount.service.v1.DiscountToMicaService/Ping", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *discountToMicaServiceClient) PingWithRoundTrip(ctx context.Context, in *pingv1.PingRequest, opts ...grpc.CallOption) (*pingv1.PingResponse, error) {
	out := new(pingv1.PingResponse)
	err := c.cc.Invoke(ctx, "/discount.service.v1.DiscountToMicaService/PingWithRoundTrip", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// DiscountToMicaServiceServer is the server API for DiscountToMicaService service.
// All implementations must embed UnimplementedDiscountToMicaServiceServer
// for forward compatibility
type DiscountToMicaServiceServer interface {
	GetDiscountProvider(context.Context, *discountproviderv1.GetDiscountProviderRequest) (*discountproviderv1.GetDiscountProviderResponse, error)
	UpdateDiscountProvider(context.Context, *discountproviderv1.UpdateDiscountProviderRequest) (*discountproviderv1.UpdateDiscountProviderResponse, error)
	CreateDiscountDefinition(context.Context, *discountdefinitionv1.CreateDiscountDefinitionRequest) (*discountdefinitionv1.CreateDiscountDefinitionResponse, error)
	GetDiscountDefinition(context.Context, *discountdefinitionv1.GetDiscountDefinitionRequest) (*discountdefinitionv1.GetDiscountDefinitionResponse, error)
	UpdateDiscountDefinition(context.Context, *discountdefinitionv1.UpdateDiscountDefinitionRequest) (*discountdefinitionv1.UpdateDiscountDefinitionResponse, error)
	RemoveDiscountDefinition(context.Context, *discountdefinitionv1.RemoveDiscountDefinitionRequest) (*discountdefinitionv1.RemoveDiscountDefinitionResponse, error)
	CreateDiscount(context.Context, *discountv1.CreateDiscountRequest) (*discountv1.CreateDiscountResponse, error)
	GetDiscount(context.Context, *discountv1.GetDiscountRequest) (*discountv1.GetDiscountResponse, error)
	CreateProductGroup(context.Context, *productgroupv1.CreateProductGroupRequest) (*productgroupv1.CreateProductGroupResponse, error)
	GetProductGroup(context.Context, *productgroupv1.GetProductGroupRequest) (*productgroupv1.GetProductGroupResponse, error)
	UpdateProductGroup(context.Context, *productgroupv1.UpdateProductGroupRequest) (*productgroupv1.UpdateProductGroupResponse, error)
	RemoveProductGroup(context.Context, *productgroupv1.RemoveProductGroupRequest) (*productgroupv1.RemoveProductGroupResponse, error)
	// An operation to ping the server to ensure it's up and running and that the connection is good.
	Ping(context.Context, *pingv1.PingRequest) (*pingv1.PingResponse, error)
	// An operation that triggers a ping to the Mica servervice and that in turn triggers a ping to the discount provider service.
	// This is a useful call to ensure that the connectivity and security between the mica an external services is working.
	PingWithRoundTrip(context.Context, *pingv1.PingRequest) (*pingv1.PingResponse, error)
	mustEmbedUnimplementedDiscountToMicaServiceServer()
}

// UnimplementedDiscountToMicaServiceServer must be embedded to have forward compatible implementations.
type UnimplementedDiscountToMicaServiceServer struct {
}

func (UnimplementedDiscountToMicaServiceServer) GetDiscountProvider(context.Context, *discountproviderv1.GetDiscountProviderRequest) (*discountproviderv1.GetDiscountProviderResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetDiscountProvider not implemented")
}
func (UnimplementedDiscountToMicaServiceServer) UpdateDiscountProvider(context.Context, *discountproviderv1.UpdateDiscountProviderRequest) (*discountproviderv1.UpdateDiscountProviderResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateDiscountProvider not implemented")
}
func (UnimplementedDiscountToMicaServiceServer) CreateDiscountDefinition(context.Context, *discountdefinitionv1.CreateDiscountDefinitionRequest) (*discountdefinitionv1.CreateDiscountDefinitionResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateDiscountDefinition not implemented")
}
func (UnimplementedDiscountToMicaServiceServer) GetDiscountDefinition(context.Context, *discountdefinitionv1.GetDiscountDefinitionRequest) (*discountdefinitionv1.GetDiscountDefinitionResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetDiscountDefinition not implemented")
}
func (UnimplementedDiscountToMicaServiceServer) UpdateDiscountDefinition(context.Context, *discountdefinitionv1.UpdateDiscountDefinitionRequest) (*discountdefinitionv1.UpdateDiscountDefinitionResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateDiscountDefinition not implemented")
}
func (UnimplementedDiscountToMicaServiceServer) RemoveDiscountDefinition(context.Context, *discountdefinitionv1.RemoveDiscountDefinitionRequest) (*discountdefinitionv1.RemoveDiscountDefinitionResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method RemoveDiscountDefinition not implemented")
}
func (UnimplementedDiscountToMicaServiceServer) CreateDiscount(context.Context, *discountv1.CreateDiscountRequest) (*discountv1.CreateDiscountResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateDiscount not implemented")
}
func (UnimplementedDiscountToMicaServiceServer) GetDiscount(context.Context, *discountv1.GetDiscountRequest) (*discountv1.GetDiscountResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetDiscount not implemented")
}
func (UnimplementedDiscountToMicaServiceServer) CreateProductGroup(context.Context, *productgroupv1.CreateProductGroupRequest) (*productgroupv1.CreateProductGroupResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateProductGroup not implemented")
}
func (UnimplementedDiscountToMicaServiceServer) GetProductGroup(context.Context, *productgroupv1.GetProductGroupRequest) (*productgroupv1.GetProductGroupResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetProductGroup not implemented")
}
func (UnimplementedDiscountToMicaServiceServer) UpdateProductGroup(context.Context, *productgroupv1.UpdateProductGroupRequest) (*productgroupv1.UpdateProductGroupResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateProductGroup not implemented")
}
func (UnimplementedDiscountToMicaServiceServer) RemoveProductGroup(context.Context, *productgroupv1.RemoveProductGroupRequest) (*productgroupv1.RemoveProductGroupResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method RemoveProductGroup not implemented")
}
func (UnimplementedDiscountToMicaServiceServer) Ping(context.Context, *pingv1.PingRequest) (*pingv1.PingResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Ping not implemented")
}
func (UnimplementedDiscountToMicaServiceServer) PingWithRoundTrip(context.Context, *pingv1.PingRequest) (*pingv1.PingResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method PingWithRoundTrip not implemented")
}
func (UnimplementedDiscountToMicaServiceServer) mustEmbedUnimplementedDiscountToMicaServiceServer() {}

// UnsafeDiscountToMicaServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to DiscountToMicaServiceServer will
// result in compilation errors.
type UnsafeDiscountToMicaServiceServer interface {
	mustEmbedUnimplementedDiscountToMicaServiceServer()
}

func RegisterDiscountToMicaServiceServer(s grpc.ServiceRegistrar, srv DiscountToMicaServiceServer) {
	s.RegisterService(&DiscountToMicaService_ServiceDesc, srv)
}

func _DiscountToMicaService_GetDiscountProvider_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(discountproviderv1.GetDiscountProviderRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DiscountToMicaServiceServer).GetDiscountProvider(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/discount.service.v1.DiscountToMicaService/GetDiscountProvider",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DiscountToMicaServiceServer).GetDiscountProvider(ctx, req.(*discountproviderv1.GetDiscountProviderRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _DiscountToMicaService_UpdateDiscountProvider_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(discountproviderv1.UpdateDiscountProviderRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DiscountToMicaServiceServer).UpdateDiscountProvider(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/discount.service.v1.DiscountToMicaService/UpdateDiscountProvider",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DiscountToMicaServiceServer).UpdateDiscountProvider(ctx, req.(*discountproviderv1.UpdateDiscountProviderRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _DiscountToMicaService_CreateDiscountDefinition_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(discountdefinitionv1.CreateDiscountDefinitionRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DiscountToMicaServiceServer).CreateDiscountDefinition(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/discount.service.v1.DiscountToMicaService/CreateDiscountDefinition",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DiscountToMicaServiceServer).CreateDiscountDefinition(ctx, req.(*discountdefinitionv1.CreateDiscountDefinitionRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _DiscountToMicaService_GetDiscountDefinition_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(discountdefinitionv1.GetDiscountDefinitionRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DiscountToMicaServiceServer).GetDiscountDefinition(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/discount.service.v1.DiscountToMicaService/GetDiscountDefinition",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DiscountToMicaServiceServer).GetDiscountDefinition(ctx, req.(*discountdefinitionv1.GetDiscountDefinitionRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _DiscountToMicaService_UpdateDiscountDefinition_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(discountdefinitionv1.UpdateDiscountDefinitionRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DiscountToMicaServiceServer).UpdateDiscountDefinition(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/discount.service.v1.DiscountToMicaService/UpdateDiscountDefinition",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DiscountToMicaServiceServer).UpdateDiscountDefinition(ctx, req.(*discountdefinitionv1.UpdateDiscountDefinitionRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _DiscountToMicaService_RemoveDiscountDefinition_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(discountdefinitionv1.RemoveDiscountDefinitionRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DiscountToMicaServiceServer).RemoveDiscountDefinition(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/discount.service.v1.DiscountToMicaService/RemoveDiscountDefinition",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DiscountToMicaServiceServer).RemoveDiscountDefinition(ctx, req.(*discountdefinitionv1.RemoveDiscountDefinitionRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _DiscountToMicaService_CreateDiscount_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(discountv1.CreateDiscountRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DiscountToMicaServiceServer).CreateDiscount(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/discount.service.v1.DiscountToMicaService/CreateDiscount",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DiscountToMicaServiceServer).CreateDiscount(ctx, req.(*discountv1.CreateDiscountRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _DiscountToMicaService_GetDiscount_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(discountv1.GetDiscountRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DiscountToMicaServiceServer).GetDiscount(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/discount.service.v1.DiscountToMicaService/GetDiscount",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DiscountToMicaServiceServer).GetDiscount(ctx, req.(*discountv1.GetDiscountRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _DiscountToMicaService_CreateProductGroup_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(productgroupv1.CreateProductGroupRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DiscountToMicaServiceServer).CreateProductGroup(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/discount.service.v1.DiscountToMicaService/CreateProductGroup",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DiscountToMicaServiceServer).CreateProductGroup(ctx, req.(*productgroupv1.CreateProductGroupRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _DiscountToMicaService_GetProductGroup_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(productgroupv1.GetProductGroupRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DiscountToMicaServiceServer).GetProductGroup(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/discount.service.v1.DiscountToMicaService/GetProductGroup",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DiscountToMicaServiceServer).GetProductGroup(ctx, req.(*productgroupv1.GetProductGroupRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _DiscountToMicaService_UpdateProductGroup_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(productgroupv1.UpdateProductGroupRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DiscountToMicaServiceServer).UpdateProductGroup(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/discount.service.v1.DiscountToMicaService/UpdateProductGroup",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DiscountToMicaServiceServer).UpdateProductGroup(ctx, req.(*productgroupv1.UpdateProductGroupRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _DiscountToMicaService_RemoveProductGroup_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(productgroupv1.RemoveProductGroupRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DiscountToMicaServiceServer).RemoveProductGroup(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/discount.service.v1.DiscountToMicaService/RemoveProductGroup",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DiscountToMicaServiceServer).RemoveProductGroup(ctx, req.(*productgroupv1.RemoveProductGroupRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _DiscountToMicaService_Ping_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(pingv1.PingRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DiscountToMicaServiceServer).Ping(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/discount.service.v1.DiscountToMicaService/Ping",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DiscountToMicaServiceServer).Ping(ctx, req.(*pingv1.PingRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _DiscountToMicaService_PingWithRoundTrip_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(pingv1.PingRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DiscountToMicaServiceServer).PingWithRoundTrip(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/discount.service.v1.DiscountToMicaService/PingWithRoundTrip",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DiscountToMicaServiceServer).PingWithRoundTrip(ctx, req.(*pingv1.PingRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// DiscountToMicaService_ServiceDesc is the grpc.ServiceDesc for DiscountToMicaService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var DiscountToMicaService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "discount.service.v1.DiscountToMicaService",
	HandlerType: (*DiscountToMicaServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetDiscountProvider",
			Handler:    _DiscountToMicaService_GetDiscountProvider_Handler,
		},
		{
			MethodName: "UpdateDiscountProvider",
			Handler:    _DiscountToMicaService_UpdateDiscountProvider_Handler,
		},
		{
			MethodName: "CreateDiscountDefinition",
			Handler:    _DiscountToMicaService_CreateDiscountDefinition_Handler,
		},
		{
			MethodName: "GetDiscountDefinition",
			Handler:    _DiscountToMicaService_GetDiscountDefinition_Handler,
		},
		{
			MethodName: "UpdateDiscountDefinition",
			Handler:    _DiscountToMicaService_UpdateDiscountDefinition_Handler,
		},
		{
			MethodName: "RemoveDiscountDefinition",
			Handler:    _DiscountToMicaService_RemoveDiscountDefinition_Handler,
		},
		{
			MethodName: "CreateDiscount",
			Handler:    _DiscountToMicaService_CreateDiscount_Handler,
		},
		{
			MethodName: "GetDiscount",
			Handler:    _DiscountToMicaService_GetDiscount_Handler,
		},
		{
			MethodName: "CreateProductGroup",
			Handler:    _DiscountToMicaService_CreateProductGroup_Handler,
		},
		{
			MethodName: "GetProductGroup",
			Handler:    _DiscountToMicaService_GetProductGroup_Handler,
		},
		{
			MethodName: "UpdateProductGroup",
			Handler:    _DiscountToMicaService_UpdateProductGroup_Handler,
		},
		{
			MethodName: "RemoveProductGroup",
			Handler:    _DiscountToMicaService_RemoveProductGroup_Handler,
		},
		{
			MethodName: "Ping",
			Handler:    _DiscountToMicaService_Ping_Handler,
		},
		{
			MethodName: "PingWithRoundTrip",
			Handler:    _DiscountToMicaService_PingWithRoundTrip_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "discount/service/v1/discount_to_mica_service.proto",
}
