// Copyright (c) 2022-2024 Mica. All rights reserved. All software, including, without limitation, all source
// code and object code, is the intellectual property of Mica, Inc. and is protected by copyright, trademark and
// other intellectual property laws (collective "Mica Software"). You may not use, copy, reproduce, download, store,
// post, broadcast, transmit, modify, sell or make available to the public content from the Mica Software without a
// valid license or the prior written approval of Mica, Inc. Mica, its logos, slogans, taglines,
// products, feature names, and other trademarks are trademarks of Mica, Inc. and may not be used without
// permission.

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.34.1
// 	protoc        v4.24.4
// source: mica/serviceprovider/serviceprovider/v1/service_provider.proto

package serviceproviderv1

import (
	_ "github.com/envoyproxy/protoc-gen-validate/validate"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	timestamppb "google.golang.org/protobuf/types/known/timestamppb"
	v1 "github.com/1080network/golang/fullsdk/proto/micashared/common/v1"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type GetServiceProviderResponse_Status int32

const (
	GetServiceProviderResponse_STATUS_UNSPECIFIED GetServiceProviderResponse_Status = 0
	GetServiceProviderResponse_STATUS_SUCCESS     GetServiceProviderResponse_Status = 1
	GetServiceProviderResponse_STATUS_ERROR       GetServiceProviderResponse_Status = 2
	GetServiceProviderResponse_STATUS_NOT_FOUND   GetServiceProviderResponse_Status = 3
)

// Enum value maps for GetServiceProviderResponse_Status.
var (
	GetServiceProviderResponse_Status_name = map[int32]string{
		0: "STATUS_UNSPECIFIED",
		1: "STATUS_SUCCESS",
		2: "STATUS_ERROR",
		3: "STATUS_NOT_FOUND",
	}
	GetServiceProviderResponse_Status_value = map[string]int32{
		"STATUS_UNSPECIFIED": 0,
		"STATUS_SUCCESS":     1,
		"STATUS_ERROR":       2,
		"STATUS_NOT_FOUND":   3,
	}
)

func (x GetServiceProviderResponse_Status) Enum() *GetServiceProviderResponse_Status {
	p := new(GetServiceProviderResponse_Status)
	*p = x
	return p
}

func (x GetServiceProviderResponse_Status) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (GetServiceProviderResponse_Status) Descriptor() protoreflect.EnumDescriptor {
	return file_mica_serviceprovider_serviceprovider_v1_service_provider_proto_enumTypes[0].Descriptor()
}

func (GetServiceProviderResponse_Status) Type() protoreflect.EnumType {
	return &file_mica_serviceprovider_serviceprovider_v1_service_provider_proto_enumTypes[0]
}

func (x GetServiceProviderResponse_Status) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use GetServiceProviderResponse_Status.Descriptor instead.
func (GetServiceProviderResponse_Status) EnumDescriptor() ([]byte, []int) {
	return file_mica_serviceprovider_serviceprovider_v1_service_provider_proto_rawDescGZIP(), []int{2, 0}
}

type ServiceProvider struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ServiceProviderKey string                 `protobuf:"bytes,1,opt,name=service_provider_key,json=serviceProviderKey,proto3" json:"service_provider_key,omitempty"`
	Version            int64                  `protobuf:"varint,2,opt,name=version,proto3" json:"version,omitempty"`
	Created            *timestamppb.Timestamp `protobuf:"bytes,3,opt,name=created,proto3" json:"created,omitempty"`
	Updated            *timestamppb.Timestamp `protobuf:"bytes,4,opt,name=updated,proto3" json:"updated,omitempty"`
	Name               string                 `protobuf:"bytes,5,opt,name=name,proto3" json:"name,omitempty"`
	ServiceProviderUrl string                 `protobuf:"bytes,6,opt,name=service_provider_url,json=serviceProviderUrl,proto3" json:"service_provider_url,omitempty"`
	OperatingAccount   *v1.BankAccountDetail  `protobuf:"bytes,8,opt,name=operating_account,json=operatingAccount,proto3" json:"operating_account,omitempty"`
	RevenueAccount     *v1.BankAccountDetail  `protobuf:"bytes,9,opt,name=revenue_account,json=revenueAccount,proto3" json:"revenue_account,omitempty"`
	NetworkGroup       string                 `protobuf:"bytes,11,opt,name=network_group,json=networkGroup,proto3" json:"network_group,omitempty"`
	Features           []string               `protobuf:"bytes,12,rep,name=features,proto3" json:"features,omitempty"`
}

func (x *ServiceProvider) Reset() {
	*x = ServiceProvider{}
	if protoimpl.UnsafeEnabled {
		mi := &file_mica_serviceprovider_serviceprovider_v1_service_provider_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ServiceProvider) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ServiceProvider) ProtoMessage() {}

func (x *ServiceProvider) ProtoReflect() protoreflect.Message {
	mi := &file_mica_serviceprovider_serviceprovider_v1_service_provider_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ServiceProvider.ProtoReflect.Descriptor instead.
func (*ServiceProvider) Descriptor() ([]byte, []int) {
	return file_mica_serviceprovider_serviceprovider_v1_service_provider_proto_rawDescGZIP(), []int{0}
}

func (x *ServiceProvider) GetServiceProviderKey() string {
	if x != nil {
		return x.ServiceProviderKey
	}
	return ""
}

func (x *ServiceProvider) GetVersion() int64 {
	if x != nil {
		return x.Version
	}
	return 0
}

func (x *ServiceProvider) GetCreated() *timestamppb.Timestamp {
	if x != nil {
		return x.Created
	}
	return nil
}

func (x *ServiceProvider) GetUpdated() *timestamppb.Timestamp {
	if x != nil {
		return x.Updated
	}
	return nil
}

func (x *ServiceProvider) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *ServiceProvider) GetServiceProviderUrl() string {
	if x != nil {
		return x.ServiceProviderUrl
	}
	return ""
}

func (x *ServiceProvider) GetOperatingAccount() *v1.BankAccountDetail {
	if x != nil {
		return x.OperatingAccount
	}
	return nil
}

func (x *ServiceProvider) GetRevenueAccount() *v1.BankAccountDetail {
	if x != nil {
		return x.RevenueAccount
	}
	return nil
}

func (x *ServiceProvider) GetNetworkGroup() string {
	if x != nil {
		return x.NetworkGroup
	}
	return ""
}

func (x *ServiceProvider) GetFeatures() []string {
	if x != nil {
		return x.Features
	}
	return nil
}

type GetServiceProviderRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *GetServiceProviderRequest) Reset() {
	*x = GetServiceProviderRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_mica_serviceprovider_serviceprovider_v1_service_provider_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetServiceProviderRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetServiceProviderRequest) ProtoMessage() {}

func (x *GetServiceProviderRequest) ProtoReflect() protoreflect.Message {
	mi := &file_mica_serviceprovider_serviceprovider_v1_service_provider_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetServiceProviderRequest.ProtoReflect.Descriptor instead.
func (*GetServiceProviderRequest) Descriptor() ([]byte, []int) {
	return file_mica_serviceprovider_serviceprovider_v1_service_provider_proto_rawDescGZIP(), []int{1}
}

type GetServiceProviderResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Status          GetServiceProviderResponse_Status `protobuf:"varint,1,opt,name=status,proto3,enum=mica.serviceprovider.serviceprovider.v1.GetServiceProviderResponse_Status" json:"status,omitempty"`
	Error           *v1.Error                         `protobuf:"bytes,2,opt,name=error,proto3" json:"error,omitempty"`
	ServiceProvider *ServiceProvider                  `protobuf:"bytes,3,opt,name=service_provider,json=serviceProvider,proto3" json:"service_provider,omitempty"`
}

func (x *GetServiceProviderResponse) Reset() {
	*x = GetServiceProviderResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_mica_serviceprovider_serviceprovider_v1_service_provider_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetServiceProviderResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetServiceProviderResponse) ProtoMessage() {}

func (x *GetServiceProviderResponse) ProtoReflect() protoreflect.Message {
	mi := &file_mica_serviceprovider_serviceprovider_v1_service_provider_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetServiceProviderResponse.ProtoReflect.Descriptor instead.
func (*GetServiceProviderResponse) Descriptor() ([]byte, []int) {
	return file_mica_serviceprovider_serviceprovider_v1_service_provider_proto_rawDescGZIP(), []int{2}
}

func (x *GetServiceProviderResponse) GetStatus() GetServiceProviderResponse_Status {
	if x != nil {
		return x.Status
	}
	return GetServiceProviderResponse_STATUS_UNSPECIFIED
}

func (x *GetServiceProviderResponse) GetError() *v1.Error {
	if x != nil {
		return x.Error
	}
	return nil
}

func (x *GetServiceProviderResponse) GetServiceProvider() *ServiceProvider {
	if x != nil {
		return x.ServiceProvider
	}
	return nil
}

var File_mica_serviceprovider_serviceprovider_v1_service_provider_proto protoreflect.FileDescriptor

var file_mica_serviceprovider_serviceprovider_v1_service_provider_proto_rawDesc = []byte{
	0x0a, 0x3e, 0x6d, 0x69, 0x63, 0x61, 0x2f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x70, 0x72,
	0x6f, 0x76, 0x69, 0x64, 0x65, 0x72, 0x2f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x70, 0x72,
	0x6f, 0x76, 0x69, 0x64, 0x65, 0x72, 0x2f, 0x76, 0x31, 0x2f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63,
	0x65, 0x5f, 0x70, 0x72, 0x6f, 0x76, 0x69, 0x64, 0x65, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x12, 0x27, 0x6d, 0x69, 0x63, 0x61, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x70, 0x72,
	0x6f, 0x76, 0x69, 0x64, 0x65, 0x72, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x70, 0x72,
	0x6f, 0x76, 0x69, 0x64, 0x65, 0x72, 0x2e, 0x76, 0x31, 0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x74, 0x69, 0x6d, 0x65, 0x73,
	0x74, 0x61, 0x6d, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x2e, 0x6d, 0x69, 0x63, 0x61,
	0x73, 0x68, 0x61, 0x72, 0x65, 0x64, 0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2f, 0x76, 0x31,
	0x2f, 0x62, 0x61, 0x6e, 0x6b, 0x5f, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x5f, 0x64, 0x65,
	0x74, 0x61, 0x69, 0x6c, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x20, 0x6d, 0x69, 0x63, 0x61,
	0x73, 0x68, 0x61, 0x72, 0x65, 0x64, 0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2f, 0x76, 0x31,
	0x2f, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x17, 0x76, 0x61,
	0x6c, 0x69, 0x64, 0x61, 0x74, 0x65, 0x2f, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x65, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x92, 0x04, 0x0a, 0x0f, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63,
	0x65, 0x50, 0x72, 0x6f, 0x76, 0x69, 0x64, 0x65, 0x72, 0x12, 0x3b, 0x0a, 0x14, 0x73, 0x65, 0x72,
	0x76, 0x69, 0x63, 0x65, 0x5f, 0x70, 0x72, 0x6f, 0x76, 0x69, 0x64, 0x65, 0x72, 0x5f, 0x6b, 0x65,
	0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x42, 0x09, 0xfa, 0x42, 0x06, 0x72, 0x04, 0x10, 0x1e,
	0x18, 0x32, 0x52, 0x12, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x50, 0x72, 0x6f, 0x76, 0x69,
	0x64, 0x65, 0x72, 0x4b, 0x65, 0x79, 0x12, 0x21, 0x0a, 0x07, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f,
	0x6e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x42, 0x07, 0xfa, 0x42, 0x04, 0x22, 0x02, 0x28, 0x00,
	0x52, 0x07, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x12, 0x34, 0x0a, 0x07, 0x63, 0x72, 0x65,
	0x61, 0x74, 0x65, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d,
	0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x07, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x12,
	0x34, 0x0a, 0x07, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62,
	0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x07, 0x75, 0x70,
	0x64, 0x61, 0x74, 0x65, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x05, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x30, 0x0a, 0x14, 0x73, 0x65, 0x72,
	0x76, 0x69, 0x63, 0x65, 0x5f, 0x70, 0x72, 0x6f, 0x76, 0x69, 0x64, 0x65, 0x72, 0x5f, 0x75, 0x72,
	0x6c, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x12, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65,
	0x50, 0x72, 0x6f, 0x76, 0x69, 0x64, 0x65, 0x72, 0x55, 0x72, 0x6c, 0x12, 0x54, 0x0a, 0x11, 0x6f,
	0x70, 0x65, 0x72, 0x61, 0x74, 0x69, 0x6e, 0x67, 0x5f, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74,
	0x18, 0x08, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x27, 0x2e, 0x6d, 0x69, 0x63, 0x61, 0x73, 0x68, 0x61,
	0x72, 0x65, 0x64, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x76, 0x31, 0x2e, 0x42, 0x61,
	0x6e, 0x6b, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x44, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x52,
	0x10, 0x6f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x69, 0x6e, 0x67, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e,
	0x74, 0x12, 0x50, 0x0a, 0x0f, 0x72, 0x65, 0x76, 0x65, 0x6e, 0x75, 0x65, 0x5f, 0x61, 0x63, 0x63,
	0x6f, 0x75, 0x6e, 0x74, 0x18, 0x09, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x27, 0x2e, 0x6d, 0x69, 0x63,
	0x61, 0x73, 0x68, 0x61, 0x72, 0x65, 0x64, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x76,
	0x31, 0x2e, 0x42, 0x61, 0x6e, 0x6b, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x44, 0x65, 0x74,
	0x61, 0x69, 0x6c, 0x52, 0x0e, 0x72, 0x65, 0x76, 0x65, 0x6e, 0x75, 0x65, 0x41, 0x63, 0x63, 0x6f,
	0x75, 0x6e, 0x74, 0x12, 0x23, 0x0a, 0x0d, 0x6e, 0x65, 0x74, 0x77, 0x6f, 0x72, 0x6b, 0x5f, 0x67,
	0x72, 0x6f, 0x75, 0x70, 0x18, 0x0b, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x6e, 0x65, 0x74, 0x77,
	0x6f, 0x72, 0x6b, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x12, 0x1a, 0x0a, 0x08, 0x66, 0x65, 0x61, 0x74,
	0x75, 0x72, 0x65, 0x73, 0x18, 0x0c, 0x20, 0x03, 0x28, 0x09, 0x52, 0x08, 0x66, 0x65, 0x61, 0x74,
	0x75, 0x72, 0x65, 0x73, 0x4a, 0x04, 0x08, 0x0a, 0x10, 0x0b, 0x22, 0x1b, 0x0a, 0x19, 0x47, 0x65,
	0x74, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x50, 0x72, 0x6f, 0x76, 0x69, 0x64, 0x65, 0x72,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x22, 0xf6, 0x02, 0x0a, 0x1a, 0x47, 0x65, 0x74, 0x53,
	0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x50, 0x72, 0x6f, 0x76, 0x69, 0x64, 0x65, 0x72, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x62, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x4a, 0x2e, 0x6d, 0x69, 0x63, 0x61, 0x2e, 0x73, 0x65,
	0x72, 0x76, 0x69, 0x63, 0x65, 0x70, 0x72, 0x6f, 0x76, 0x69, 0x64, 0x65, 0x72, 0x2e, 0x73, 0x65,
	0x72, 0x76, 0x69, 0x63, 0x65, 0x70, 0x72, 0x6f, 0x76, 0x69, 0x64, 0x65, 0x72, 0x2e, 0x76, 0x31,
	0x2e, 0x47, 0x65, 0x74, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x50, 0x72, 0x6f, 0x76, 0x69,
	0x64, 0x65, 0x72, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x2e, 0x53, 0x74, 0x61, 0x74,
	0x75, 0x73, 0x52, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x31, 0x0a, 0x05, 0x65, 0x72,
	0x72, 0x6f, 0x72, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1b, 0x2e, 0x6d, 0x69, 0x63, 0x61,
	0x73, 0x68, 0x61, 0x72, 0x65, 0x64, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x76, 0x31,
	0x2e, 0x45, 0x72, 0x72, 0x6f, 0x72, 0x52, 0x05, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x12, 0x63, 0x0a,
	0x10, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x5f, 0x70, 0x72, 0x6f, 0x76, 0x69, 0x64, 0x65,
	0x72, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x38, 0x2e, 0x6d, 0x69, 0x63, 0x61, 0x2e, 0x73,
	0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x70, 0x72, 0x6f, 0x76, 0x69, 0x64, 0x65, 0x72, 0x2e, 0x73,
	0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x70, 0x72, 0x6f, 0x76, 0x69, 0x64, 0x65, 0x72, 0x2e, 0x76,
	0x31, 0x2e, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x50, 0x72, 0x6f, 0x76, 0x69, 0x64, 0x65,
	0x72, 0x52, 0x0f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x50, 0x72, 0x6f, 0x76, 0x69, 0x64,
	0x65, 0x72, 0x22, 0x5c, 0x0a, 0x06, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x16, 0x0a, 0x12,
	0x53, 0x54, 0x41, 0x54, 0x55, 0x53, 0x5f, 0x55, 0x4e, 0x53, 0x50, 0x45, 0x43, 0x49, 0x46, 0x49,
	0x45, 0x44, 0x10, 0x00, 0x12, 0x12, 0x0a, 0x0e, 0x53, 0x54, 0x41, 0x54, 0x55, 0x53, 0x5f, 0x53,
	0x55, 0x43, 0x43, 0x45, 0x53, 0x53, 0x10, 0x01, 0x12, 0x10, 0x0a, 0x0c, 0x53, 0x54, 0x41, 0x54,
	0x55, 0x53, 0x5f, 0x45, 0x52, 0x52, 0x4f, 0x52, 0x10, 0x02, 0x12, 0x14, 0x0a, 0x10, 0x53, 0x54,
	0x41, 0x54, 0x55, 0x53, 0x5f, 0x4e, 0x4f, 0x54, 0x5f, 0x46, 0x4f, 0x55, 0x4e, 0x44, 0x10, 0x03,
	0x42, 0x71, 0x0a, 0x2a, 0x69, 0x6f, 0x2e, 0x6d, 0x69, 0x63, 0x61, 0x2e, 0x73, 0x65, 0x72, 0x76,
	0x69, 0x63, 0x65, 0x70, 0x72, 0x6f, 0x76, 0x69, 0x64, 0x65, 0x72, 0x2e, 0x73, 0x65, 0x72, 0x76,
	0x69, 0x63, 0x65, 0x70, 0x72, 0x6f, 0x76, 0x69, 0x64, 0x65, 0x72, 0x2e, 0x76, 0x31, 0x42, 0x14,
	0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x50, 0x72, 0x6f, 0x76, 0x69, 0x64, 0x65, 0x72, 0x50,
	0x72, 0x6f, 0x74, 0x6f, 0x5a, 0x26, 0x6d, 0x69, 0x63, 0x61, 0x2f, 0x73, 0x65, 0x72, 0x76, 0x69,
	0x63, 0x65, 0x70, 0x72, 0x6f, 0x76, 0x69, 0x64, 0x65, 0x72, 0x2f, 0x73, 0x65, 0x72, 0x76, 0x69,
	0x63, 0x65, 0x70, 0x72, 0x6f, 0x76, 0x69, 0x64, 0x65, 0x72, 0x76, 0x31, 0xa2, 0x02, 0x04, 0x4d,
	0x49, 0x43, 0x41, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_mica_serviceprovider_serviceprovider_v1_service_provider_proto_rawDescOnce sync.Once
	file_mica_serviceprovider_serviceprovider_v1_service_provider_proto_rawDescData = file_mica_serviceprovider_serviceprovider_v1_service_provider_proto_rawDesc
)

func file_mica_serviceprovider_serviceprovider_v1_service_provider_proto_rawDescGZIP() []byte {
	file_mica_serviceprovider_serviceprovider_v1_service_provider_proto_rawDescOnce.Do(func() {
		file_mica_serviceprovider_serviceprovider_v1_service_provider_proto_rawDescData = protoimpl.X.CompressGZIP(file_mica_serviceprovider_serviceprovider_v1_service_provider_proto_rawDescData)
	})
	return file_mica_serviceprovider_serviceprovider_v1_service_provider_proto_rawDescData
}

var file_mica_serviceprovider_serviceprovider_v1_service_provider_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_mica_serviceprovider_serviceprovider_v1_service_provider_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_mica_serviceprovider_serviceprovider_v1_service_provider_proto_goTypes = []interface{}{
	(GetServiceProviderResponse_Status)(0), // 0: mica.serviceprovider.serviceprovider.v1.GetServiceProviderResponse.Status
	(*ServiceProvider)(nil),                // 1: mica.serviceprovider.serviceprovider.v1.ServiceProvider
	(*GetServiceProviderRequest)(nil),      // 2: mica.serviceprovider.serviceprovider.v1.GetServiceProviderRequest
	(*GetServiceProviderResponse)(nil),     // 3: mica.serviceprovider.serviceprovider.v1.GetServiceProviderResponse
	(*timestamppb.Timestamp)(nil),          // 4: google.protobuf.Timestamp
	(*v1.BankAccountDetail)(nil),           // 5: micashared.common.v1.BankAccountDetail
	(*v1.Error)(nil),                       // 6: micashared.common.v1.Error
}
var file_mica_serviceprovider_serviceprovider_v1_service_provider_proto_depIdxs = []int32{
	4, // 0: mica.serviceprovider.serviceprovider.v1.ServiceProvider.created:type_name -> google.protobuf.Timestamp
	4, // 1: mica.serviceprovider.serviceprovider.v1.ServiceProvider.updated:type_name -> google.protobuf.Timestamp
	5, // 2: mica.serviceprovider.serviceprovider.v1.ServiceProvider.operating_account:type_name -> micashared.common.v1.BankAccountDetail
	5, // 3: mica.serviceprovider.serviceprovider.v1.ServiceProvider.revenue_account:type_name -> micashared.common.v1.BankAccountDetail
	0, // 4: mica.serviceprovider.serviceprovider.v1.GetServiceProviderResponse.status:type_name -> mica.serviceprovider.serviceprovider.v1.GetServiceProviderResponse.Status
	6, // 5: mica.serviceprovider.serviceprovider.v1.GetServiceProviderResponse.error:type_name -> micashared.common.v1.Error
	1, // 6: mica.serviceprovider.serviceprovider.v1.GetServiceProviderResponse.service_provider:type_name -> mica.serviceprovider.serviceprovider.v1.ServiceProvider
	7, // [7:7] is the sub-list for method output_type
	7, // [7:7] is the sub-list for method input_type
	7, // [7:7] is the sub-list for extension type_name
	7, // [7:7] is the sub-list for extension extendee
	0, // [0:7] is the sub-list for field type_name
}

func init() { file_mica_serviceprovider_serviceprovider_v1_service_provider_proto_init() }
func file_mica_serviceprovider_serviceprovider_v1_service_provider_proto_init() {
	if File_mica_serviceprovider_serviceprovider_v1_service_provider_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_mica_serviceprovider_serviceprovider_v1_service_provider_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ServiceProvider); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_mica_serviceprovider_serviceprovider_v1_service_provider_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetServiceProviderRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_mica_serviceprovider_serviceprovider_v1_service_provider_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetServiceProviderResponse); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_mica_serviceprovider_serviceprovider_v1_service_provider_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_mica_serviceprovider_serviceprovider_v1_service_provider_proto_goTypes,
		DependencyIndexes: file_mica_serviceprovider_serviceprovider_v1_service_provider_proto_depIdxs,
		EnumInfos:         file_mica_serviceprovider_serviceprovider_v1_service_provider_proto_enumTypes,
		MessageInfos:      file_mica_serviceprovider_serviceprovider_v1_service_provider_proto_msgTypes,
	}.Build()
	File_mica_serviceprovider_serviceprovider_v1_service_provider_proto = out.File
	file_mica_serviceprovider_serviceprovider_v1_service_provider_proto_rawDesc = nil
	file_mica_serviceprovider_serviceprovider_v1_service_provider_proto_goTypes = nil
	file_mica_serviceprovider_serviceprovider_v1_service_provider_proto_depIdxs = nil
}
