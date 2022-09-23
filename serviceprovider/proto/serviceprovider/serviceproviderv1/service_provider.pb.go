// Copyright (c) 2022 Mica. All rights reserved. All software, including, without limitation, all source
// code and object code, is the intellectual property of Mica, Inc. and is protected by copyright, trademark and
// other intellectual property laws (collective "Mica Software"). You may not use, copy, reproduce, download, store,
// post, broadcast, transmit, modify, sell or make available to the public content from the Mica Software without a
// valid license or the prior written approval of Mica, Inc. Mica, its logos, slogans, taglines,
// products, feature names, and other trademarks are trademarks of Mica, Inc. and may not be used without
// permission.

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.0
// 	protoc        v3.19.4
// source: serviceprovider/serviceprovider/v1/service_provider.proto

package serviceproviderv1

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	timestamppb "google.golang.org/protobuf/types/known/timestamppb"
	accounttypev1 "github.com/1080network/golang/serviceprovider/proto/common/enums/accounttypev1"
	currencyv1 "github.com/1080network/golang/serviceprovider/proto/common/enums/currencyv1"
	v1 "github.com/1080network/golang/serviceprovider/proto/common/v1"
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
	return file_serviceprovider_serviceprovider_v1_service_provider_proto_enumTypes[0].Descriptor()
}

func (GetServiceProviderResponse_Status) Type() protoreflect.EnumType {
	return &file_serviceprovider_serviceprovider_v1_service_provider_proto_enumTypes[0]
}

func (x GetServiceProviderResponse_Status) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use GetServiceProviderResponse_Status.Descriptor instead.
func (GetServiceProviderResponse_Status) EnumDescriptor() ([]byte, []int) {
	return file_serviceprovider_serviceprovider_v1_service_provider_proto_rawDescGZIP(), []int{3, 0}
}

type ServiceProviderAccount struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ServiceProviderAccountKey string                    `protobuf:"bytes,1,opt,name=service_provider_account_key,json=serviceProviderAccountKey,proto3" json:"service_provider_account_key,omitempty"`
	Created                   *timestamppb.Timestamp    `protobuf:"bytes,2,opt,name=created,proto3" json:"created,omitempty"`
	Updated                   *timestamppb.Timestamp    `protobuf:"bytes,3,opt,name=updated,proto3" json:"updated,omitempty"`
	Version                   int64                     `protobuf:"varint,4,opt,name=version,proto3" json:"version,omitempty"`
	AccountType               accounttypev1.AccountType `protobuf:"varint,5,opt,name=account_type,json=accountType,proto3,enum=common.enums.accounttype.v1.AccountType" json:"account_type,omitempty"`
	Currency                  currencyv1.Currency       `protobuf:"varint,6,opt,name=currency,proto3,enum=common.enums.currency.v1.Currency" json:"currency,omitempty"`
}

func (x *ServiceProviderAccount) Reset() {
	*x = ServiceProviderAccount{}
	if protoimpl.UnsafeEnabled {
		mi := &file_serviceprovider_serviceprovider_v1_service_provider_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ServiceProviderAccount) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ServiceProviderAccount) ProtoMessage() {}

func (x *ServiceProviderAccount) ProtoReflect() protoreflect.Message {
	mi := &file_serviceprovider_serviceprovider_v1_service_provider_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ServiceProviderAccount.ProtoReflect.Descriptor instead.
func (*ServiceProviderAccount) Descriptor() ([]byte, []int) {
	return file_serviceprovider_serviceprovider_v1_service_provider_proto_rawDescGZIP(), []int{0}
}

func (x *ServiceProviderAccount) GetServiceProviderAccountKey() string {
	if x != nil {
		return x.ServiceProviderAccountKey
	}
	return ""
}

func (x *ServiceProviderAccount) GetCreated() *timestamppb.Timestamp {
	if x != nil {
		return x.Created
	}
	return nil
}

func (x *ServiceProviderAccount) GetUpdated() *timestamppb.Timestamp {
	if x != nil {
		return x.Updated
	}
	return nil
}

func (x *ServiceProviderAccount) GetVersion() int64 {
	if x != nil {
		return x.Version
	}
	return 0
}

func (x *ServiceProviderAccount) GetAccountType() accounttypev1.AccountType {
	if x != nil {
		return x.AccountType
	}
	return accounttypev1.AccountType(0)
}

func (x *ServiceProviderAccount) GetCurrency() currencyv1.Currency {
	if x != nil {
		return x.Currency
	}
	return currencyv1.Currency(0)
}

type ServiceProvider struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ServiceProviderKey      string                    `protobuf:"bytes,1,opt,name=service_provider_key,json=serviceProviderKey,proto3" json:"service_provider_key,omitempty"`
	Created                 *timestamppb.Timestamp    `protobuf:"bytes,2,opt,name=created,proto3" json:"created,omitempty"`
	Updated                 *timestamppb.Timestamp    `protobuf:"bytes,3,opt,name=updated,proto3" json:"updated,omitempty"`
	Version                 int64                     `protobuf:"varint,4,opt,name=version,proto3" json:"version,omitempty"`
	Name                    string                    `protobuf:"bytes,5,opt,name=name,proto3" json:"name,omitempty"`
	ServiceProviderUrl      string                    `protobuf:"bytes,6,opt,name=service_provider_url,json=serviceProviderUrl,proto3" json:"service_provider_url,omitempty"`
	ServiceProviderAccounts []*ServiceProviderAccount `protobuf:"bytes,7,rep,name=service_provider_accounts,json=serviceProviderAccounts,proto3" json:"service_provider_accounts,omitempty"`
}

func (x *ServiceProvider) Reset() {
	*x = ServiceProvider{}
	if protoimpl.UnsafeEnabled {
		mi := &file_serviceprovider_serviceprovider_v1_service_provider_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ServiceProvider) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ServiceProvider) ProtoMessage() {}

func (x *ServiceProvider) ProtoReflect() protoreflect.Message {
	mi := &file_serviceprovider_serviceprovider_v1_service_provider_proto_msgTypes[1]
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
	return file_serviceprovider_serviceprovider_v1_service_provider_proto_rawDescGZIP(), []int{1}
}

func (x *ServiceProvider) GetServiceProviderKey() string {
	if x != nil {
		return x.ServiceProviderKey
	}
	return ""
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

func (x *ServiceProvider) GetVersion() int64 {
	if x != nil {
		return x.Version
	}
	return 0
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

func (x *ServiceProvider) GetServiceProviderAccounts() []*ServiceProviderAccount {
	if x != nil {
		return x.ServiceProviderAccounts
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
		mi := &file_serviceprovider_serviceprovider_v1_service_provider_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetServiceProviderRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetServiceProviderRequest) ProtoMessage() {}

func (x *GetServiceProviderRequest) ProtoReflect() protoreflect.Message {
	mi := &file_serviceprovider_serviceprovider_v1_service_provider_proto_msgTypes[2]
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
	return file_serviceprovider_serviceprovider_v1_service_provider_proto_rawDescGZIP(), []int{2}
}

type GetServiceProviderResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Status          GetServiceProviderResponse_Status `protobuf:"varint,1,opt,name=status,proto3,enum=serviceprovider.serviceprovider.v1.GetServiceProviderResponse_Status" json:"status,omitempty"`
	Error           *v1.Error                         `protobuf:"bytes,2,opt,name=error,proto3" json:"error,omitempty"`
	ServiceProvider *ServiceProvider                  `protobuf:"bytes,3,opt,name=service_provider,json=serviceProvider,proto3" json:"service_provider,omitempty"`
}

func (x *GetServiceProviderResponse) Reset() {
	*x = GetServiceProviderResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_serviceprovider_serviceprovider_v1_service_provider_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetServiceProviderResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetServiceProviderResponse) ProtoMessage() {}

func (x *GetServiceProviderResponse) ProtoReflect() protoreflect.Message {
	mi := &file_serviceprovider_serviceprovider_v1_service_provider_proto_msgTypes[3]
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
	return file_serviceprovider_serviceprovider_v1_service_provider_proto_rawDescGZIP(), []int{3}
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

var File_serviceprovider_serviceprovider_v1_service_provider_proto protoreflect.FileDescriptor

var file_serviceprovider_serviceprovider_v1_service_provider_proto_rawDesc = []byte{
	0x0a, 0x39, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x70, 0x72, 0x6f, 0x76, 0x69, 0x64, 0x65,
	0x72, 0x2f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x70, 0x72, 0x6f, 0x76, 0x69, 0x64, 0x65,
	0x72, 0x2f, 0x76, 0x31, 0x2f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x5f, 0x70, 0x72, 0x6f,
	0x76, 0x69, 0x64, 0x65, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x22, 0x73, 0x65, 0x72,
	0x76, 0x69, 0x63, 0x65, 0x70, 0x72, 0x6f, 0x76, 0x69, 0x64, 0x65, 0x72, 0x2e, 0x73, 0x65, 0x72,
	0x76, 0x69, 0x63, 0x65, 0x70, 0x72, 0x6f, 0x76, 0x69, 0x64, 0x65, 0x72, 0x2e, 0x76, 0x31, 0x1a,
	0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2f, 0x65, 0x6e, 0x75, 0x6d, 0x73, 0x2f, 0x61, 0x63,
	0x63, 0x6f, 0x75, 0x6e, 0x74, 0x74, 0x79, 0x70, 0x65, 0x2f, 0x76, 0x31, 0x2f, 0x61, 0x63, 0x63,
	0x6f, 0x75, 0x6e, 0x74, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a,
	0x27, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2f, 0x65, 0x6e, 0x75, 0x6d, 0x73, 0x2f, 0x63, 0x75,
	0x72, 0x72, 0x65, 0x6e, 0x63, 0x79, 0x2f, 0x76, 0x31, 0x2f, 0x63, 0x75, 0x72, 0x72, 0x65, 0x6e,
	0x63, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x15, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e,
	0x2f, 0x76, 0x31, 0x2f, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a,
	0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66,
	0x2f, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x22, 0xec, 0x02, 0x0a, 0x16, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x50, 0x72, 0x6f, 0x76,
	0x69, 0x64, 0x65, 0x72, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x12, 0x3f, 0x0a, 0x1c, 0x73,
	0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x5f, 0x70, 0x72, 0x6f, 0x76, 0x69, 0x64, 0x65, 0x72, 0x5f,
	0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x5f, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x19, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x50, 0x72, 0x6f, 0x76, 0x69, 0x64,
	0x65, 0x72, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x4b, 0x65, 0x79, 0x12, 0x34, 0x0a, 0x07,
	0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e,
	0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e,
	0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x07, 0x63, 0x72, 0x65, 0x61, 0x74,
	0x65, 0x64, 0x12, 0x34, 0x0a, 0x07, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x18, 0x03, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52,
	0x07, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x12, 0x18, 0x0a, 0x07, 0x76, 0x65, 0x72, 0x73,
	0x69, 0x6f, 0x6e, 0x18, 0x04, 0x20, 0x01, 0x28, 0x03, 0x52, 0x07, 0x76, 0x65, 0x72, 0x73, 0x69,
	0x6f, 0x6e, 0x12, 0x4b, 0x0a, 0x0c, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x5f, 0x74, 0x79,
	0x70, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x28, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f,
	0x6e, 0x2e, 0x65, 0x6e, 0x75, 0x6d, 0x73, 0x2e, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x74,
	0x79, 0x70, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x54, 0x79,
	0x70, 0x65, 0x52, 0x0b, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x54, 0x79, 0x70, 0x65, 0x12,
	0x3e, 0x0a, 0x08, 0x63, 0x75, 0x72, 0x72, 0x65, 0x6e, 0x63, 0x79, 0x18, 0x06, 0x20, 0x01, 0x28,
	0x0e, 0x32, 0x22, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x65, 0x6e, 0x75, 0x6d, 0x73,
	0x2e, 0x63, 0x75, 0x72, 0x72, 0x65, 0x6e, 0x63, 0x79, 0x2e, 0x76, 0x31, 0x2e, 0x43, 0x75, 0x72,
	0x72, 0x65, 0x6e, 0x63, 0x79, 0x52, 0x08, 0x63, 0x75, 0x72, 0x72, 0x65, 0x6e, 0x63, 0x79, 0x22,
	0x87, 0x03, 0x0a, 0x0f, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x50, 0x72, 0x6f, 0x76, 0x69,
	0x64, 0x65, 0x72, 0x12, 0x30, 0x0a, 0x14, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x5f, 0x70,
	0x72, 0x6f, 0x76, 0x69, 0x64, 0x65, 0x72, 0x5f, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x12, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x50, 0x72, 0x6f, 0x76, 0x69, 0x64,
	0x65, 0x72, 0x4b, 0x65, 0x79, 0x12, 0x34, 0x0a, 0x07, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61,
	0x6d, 0x70, 0x52, 0x07, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x12, 0x34, 0x0a, 0x07, 0x75,
	0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54,
	0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x07, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65,
	0x64, 0x12, 0x18, 0x0a, 0x07, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x18, 0x04, 0x20, 0x01,
	0x28, 0x03, 0x52, 0x07, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x12, 0x12, 0x0a, 0x04, 0x6e,
	0x61, 0x6d, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12,
	0x30, 0x0a, 0x14, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x5f, 0x70, 0x72, 0x6f, 0x76, 0x69,
	0x64, 0x65, 0x72, 0x5f, 0x75, 0x72, 0x6c, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x12, 0x73,
	0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x50, 0x72, 0x6f, 0x76, 0x69, 0x64, 0x65, 0x72, 0x55, 0x72,
	0x6c, 0x12, 0x76, 0x0a, 0x19, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x5f, 0x70, 0x72, 0x6f,
	0x76, 0x69, 0x64, 0x65, 0x72, 0x5f, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x73, 0x18, 0x07,
	0x20, 0x03, 0x28, 0x0b, 0x32, 0x3a, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x70, 0x72,
	0x6f, 0x76, 0x69, 0x64, 0x65, 0x72, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x70, 0x72,
	0x6f, 0x76, 0x69, 0x64, 0x65, 0x72, 0x2e, 0x76, 0x31, 0x2e, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63,
	0x65, 0x50, 0x72, 0x6f, 0x76, 0x69, 0x64, 0x65, 0x72, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74,
	0x52, 0x17, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x50, 0x72, 0x6f, 0x76, 0x69, 0x64, 0x65,
	0x72, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x73, 0x22, 0x1b, 0x0a, 0x19, 0x47, 0x65, 0x74,
	0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x50, 0x72, 0x6f, 0x76, 0x69, 0x64, 0x65, 0x72, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x22, 0xe1, 0x02, 0x0a, 0x1a, 0x47, 0x65, 0x74, 0x53, 0x65,
	0x72, 0x76, 0x69, 0x63, 0x65, 0x50, 0x72, 0x6f, 0x76, 0x69, 0x64, 0x65, 0x72, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x5d, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x45, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x70,
	0x72, 0x6f, 0x76, 0x69, 0x64, 0x65, 0x72, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x70,
	0x72, 0x6f, 0x76, 0x69, 0x64, 0x65, 0x72, 0x2e, 0x76, 0x31, 0x2e, 0x47, 0x65, 0x74, 0x53, 0x65,
	0x72, 0x76, 0x69, 0x63, 0x65, 0x50, 0x72, 0x6f, 0x76, 0x69, 0x64, 0x65, 0x72, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x2e, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x52, 0x06, 0x73, 0x74,
	0x61, 0x74, 0x75, 0x73, 0x12, 0x26, 0x0a, 0x05, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x10, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x76, 0x31, 0x2e,
	0x45, 0x72, 0x72, 0x6f, 0x72, 0x52, 0x05, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x12, 0x5e, 0x0a, 0x10,
	0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x5f, 0x70, 0x72, 0x6f, 0x76, 0x69, 0x64, 0x65, 0x72,
	0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x33, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65,
	0x70, 0x72, 0x6f, 0x76, 0x69, 0x64, 0x65, 0x72, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65,
	0x70, 0x72, 0x6f, 0x76, 0x69, 0x64, 0x65, 0x72, 0x2e, 0x76, 0x31, 0x2e, 0x53, 0x65, 0x72, 0x76,
	0x69, 0x63, 0x65, 0x50, 0x72, 0x6f, 0x76, 0x69, 0x64, 0x65, 0x72, 0x52, 0x0f, 0x73, 0x65, 0x72,
	0x76, 0x69, 0x63, 0x65, 0x50, 0x72, 0x6f, 0x76, 0x69, 0x64, 0x65, 0x72, 0x22, 0x5c, 0x0a, 0x06,
	0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x16, 0x0a, 0x12, 0x53, 0x54, 0x41, 0x54, 0x55, 0x53,
	0x5f, 0x55, 0x4e, 0x53, 0x50, 0x45, 0x43, 0x49, 0x46, 0x49, 0x45, 0x44, 0x10, 0x00, 0x12, 0x12,
	0x0a, 0x0e, 0x53, 0x54, 0x41, 0x54, 0x55, 0x53, 0x5f, 0x53, 0x55, 0x43, 0x43, 0x45, 0x53, 0x53,
	0x10, 0x01, 0x12, 0x10, 0x0a, 0x0c, 0x53, 0x54, 0x41, 0x54, 0x55, 0x53, 0x5f, 0x45, 0x52, 0x52,
	0x4f, 0x52, 0x10, 0x02, 0x12, 0x14, 0x0a, 0x10, 0x53, 0x54, 0x41, 0x54, 0x55, 0x53, 0x5f, 0x4e,
	0x4f, 0x54, 0x5f, 0x46, 0x4f, 0x55, 0x4e, 0x44, 0x10, 0x03, 0x42, 0x77, 0x0a, 0x2a, 0x69, 0x6f,
	0x2e, 0x6d, 0x69, 0x63, 0x61, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x70, 0x72, 0x6f,
	0x76, 0x69, 0x64, 0x65, 0x72, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x70, 0x72, 0x6f,
	0x76, 0x69, 0x64, 0x65, 0x72, 0x2e, 0x76, 0x31, 0x42, 0x14, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63,
	0x65, 0x50, 0x72, 0x6f, 0x76, 0x69, 0x64, 0x65, 0x72, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x5a, 0x2c,
	0x6d, 0x69, 0x63, 0x61, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x73, 0x65, 0x72, 0x76, 0x69,
	0x63, 0x65, 0x70, 0x72, 0x6f, 0x76, 0x69, 0x64, 0x65, 0x72, 0x2f, 0x73, 0x65, 0x72, 0x76, 0x69,
	0x63, 0x65, 0x70, 0x72, 0x6f, 0x76, 0x69, 0x64, 0x65, 0x72, 0x76, 0x31, 0xa2, 0x02, 0x04, 0x4d,
	0x49, 0x43, 0x41, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_serviceprovider_serviceprovider_v1_service_provider_proto_rawDescOnce sync.Once
	file_serviceprovider_serviceprovider_v1_service_provider_proto_rawDescData = file_serviceprovider_serviceprovider_v1_service_provider_proto_rawDesc
)

func file_serviceprovider_serviceprovider_v1_service_provider_proto_rawDescGZIP() []byte {
	file_serviceprovider_serviceprovider_v1_service_provider_proto_rawDescOnce.Do(func() {
		file_serviceprovider_serviceprovider_v1_service_provider_proto_rawDescData = protoimpl.X.CompressGZIP(file_serviceprovider_serviceprovider_v1_service_provider_proto_rawDescData)
	})
	return file_serviceprovider_serviceprovider_v1_service_provider_proto_rawDescData
}

var file_serviceprovider_serviceprovider_v1_service_provider_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_serviceprovider_serviceprovider_v1_service_provider_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_serviceprovider_serviceprovider_v1_service_provider_proto_goTypes = []interface{}{
	(GetServiceProviderResponse_Status)(0), // 0: serviceprovider.serviceprovider.v1.GetServiceProviderResponse.Status
	(*ServiceProviderAccount)(nil),         // 1: serviceprovider.serviceprovider.v1.ServiceProviderAccount
	(*ServiceProvider)(nil),                // 2: serviceprovider.serviceprovider.v1.ServiceProvider
	(*GetServiceProviderRequest)(nil),      // 3: serviceprovider.serviceprovider.v1.GetServiceProviderRequest
	(*GetServiceProviderResponse)(nil),     // 4: serviceprovider.serviceprovider.v1.GetServiceProviderResponse
	(*timestamppb.Timestamp)(nil),          // 5: google.protobuf.Timestamp
	(accounttypev1.AccountType)(0),         // 6: common.enums.accounttype.v1.AccountType
	(currencyv1.Currency)(0),               // 7: common.enums.currency.v1.Currency
	(*v1.Error)(nil),                       // 8: common.v1.Error
}
var file_serviceprovider_serviceprovider_v1_service_provider_proto_depIdxs = []int32{
	5,  // 0: serviceprovider.serviceprovider.v1.ServiceProviderAccount.created:type_name -> google.protobuf.Timestamp
	5,  // 1: serviceprovider.serviceprovider.v1.ServiceProviderAccount.updated:type_name -> google.protobuf.Timestamp
	6,  // 2: serviceprovider.serviceprovider.v1.ServiceProviderAccount.account_type:type_name -> common.enums.accounttype.v1.AccountType
	7,  // 3: serviceprovider.serviceprovider.v1.ServiceProviderAccount.currency:type_name -> common.enums.currency.v1.Currency
	5,  // 4: serviceprovider.serviceprovider.v1.ServiceProvider.created:type_name -> google.protobuf.Timestamp
	5,  // 5: serviceprovider.serviceprovider.v1.ServiceProvider.updated:type_name -> google.protobuf.Timestamp
	1,  // 6: serviceprovider.serviceprovider.v1.ServiceProvider.service_provider_accounts:type_name -> serviceprovider.serviceprovider.v1.ServiceProviderAccount
	0,  // 7: serviceprovider.serviceprovider.v1.GetServiceProviderResponse.status:type_name -> serviceprovider.serviceprovider.v1.GetServiceProviderResponse.Status
	8,  // 8: serviceprovider.serviceprovider.v1.GetServiceProviderResponse.error:type_name -> common.v1.Error
	2,  // 9: serviceprovider.serviceprovider.v1.GetServiceProviderResponse.service_provider:type_name -> serviceprovider.serviceprovider.v1.ServiceProvider
	10, // [10:10] is the sub-list for method output_type
	10, // [10:10] is the sub-list for method input_type
	10, // [10:10] is the sub-list for extension type_name
	10, // [10:10] is the sub-list for extension extendee
	0,  // [0:10] is the sub-list for field type_name
}

func init() { file_serviceprovider_serviceprovider_v1_service_provider_proto_init() }
func file_serviceprovider_serviceprovider_v1_service_provider_proto_init() {
	if File_serviceprovider_serviceprovider_v1_service_provider_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_serviceprovider_serviceprovider_v1_service_provider_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ServiceProviderAccount); i {
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
		file_serviceprovider_serviceprovider_v1_service_provider_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
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
		file_serviceprovider_serviceprovider_v1_service_provider_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
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
		file_serviceprovider_serviceprovider_v1_service_provider_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
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
			RawDescriptor: file_serviceprovider_serviceprovider_v1_service_provider_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_serviceprovider_serviceprovider_v1_service_provider_proto_goTypes,
		DependencyIndexes: file_serviceprovider_serviceprovider_v1_service_provider_proto_depIdxs,
		EnumInfos:         file_serviceprovider_serviceprovider_v1_service_provider_proto_enumTypes,
		MessageInfos:      file_serviceprovider_serviceprovider_v1_service_provider_proto_msgTypes,
	}.Build()
	File_serviceprovider_serviceprovider_v1_service_provider_proto = out.File
	file_serviceprovider_serviceprovider_v1_service_provider_proto_rawDesc = nil
	file_serviceprovider_serviceprovider_v1_service_provider_proto_goTypes = nil
	file_serviceprovider_serviceprovider_v1_service_provider_proto_depIdxs = nil
}
