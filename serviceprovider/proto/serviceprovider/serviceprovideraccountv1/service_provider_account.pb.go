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
// 	protoc        v3.21.8
// source: serviceprovider/serviceprovideraccount/v1/service_provider_account.proto

package serviceprovideraccountv1

import (
	_ "github.com/envoyproxy/protoc-gen-validate/validate"
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

type GetServiceProviderAccountResponse_Status int32

const (
	GetServiceProviderAccountResponse_STATUS_UNSPECIFIED GetServiceProviderAccountResponse_Status = 0
	GetServiceProviderAccountResponse_STATUS_SUCCESS     GetServiceProviderAccountResponse_Status = 1
	GetServiceProviderAccountResponse_STATUS_ERROR       GetServiceProviderAccountResponse_Status = 2
	GetServiceProviderAccountResponse_STATUS_NOT_FOUND   GetServiceProviderAccountResponse_Status = 3
)

// Enum value maps for GetServiceProviderAccountResponse_Status.
var (
	GetServiceProviderAccountResponse_Status_name = map[int32]string{
		0: "STATUS_UNSPECIFIED",
		1: "STATUS_SUCCESS",
		2: "STATUS_ERROR",
		3: "STATUS_NOT_FOUND",
	}
	GetServiceProviderAccountResponse_Status_value = map[string]int32{
		"STATUS_UNSPECIFIED": 0,
		"STATUS_SUCCESS":     1,
		"STATUS_ERROR":       2,
		"STATUS_NOT_FOUND":   3,
	}
)

func (x GetServiceProviderAccountResponse_Status) Enum() *GetServiceProviderAccountResponse_Status {
	p := new(GetServiceProviderAccountResponse_Status)
	*p = x
	return p
}

func (x GetServiceProviderAccountResponse_Status) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (GetServiceProviderAccountResponse_Status) Descriptor() protoreflect.EnumDescriptor {
	return file_serviceprovider_serviceprovideraccount_v1_service_provider_account_proto_enumTypes[0].Descriptor()
}

func (GetServiceProviderAccountResponse_Status) Type() protoreflect.EnumType {
	return &file_serviceprovider_serviceprovideraccount_v1_service_provider_account_proto_enumTypes[0]
}

func (x GetServiceProviderAccountResponse_Status) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use GetServiceProviderAccountResponse_Status.Descriptor instead.
func (GetServiceProviderAccountResponse_Status) EnumDescriptor() ([]byte, []int) {
	return file_serviceprovider_serviceprovideraccount_v1_service_provider_account_proto_rawDescGZIP(), []int{2, 0}
}

type SearchServiceProviderAccountResponse_Status int32

const (
	SearchServiceProviderAccountResponse_STATUS_UNSPECIFIED SearchServiceProviderAccountResponse_Status = 0
	SearchServiceProviderAccountResponse_STATUS_SUCCESS     SearchServiceProviderAccountResponse_Status = 1
	SearchServiceProviderAccountResponse_STATUS_ERROR       SearchServiceProviderAccountResponse_Status = 2
)

// Enum value maps for SearchServiceProviderAccountResponse_Status.
var (
	SearchServiceProviderAccountResponse_Status_name = map[int32]string{
		0: "STATUS_UNSPECIFIED",
		1: "STATUS_SUCCESS",
		2: "STATUS_ERROR",
	}
	SearchServiceProviderAccountResponse_Status_value = map[string]int32{
		"STATUS_UNSPECIFIED": 0,
		"STATUS_SUCCESS":     1,
		"STATUS_ERROR":       2,
	}
)

func (x SearchServiceProviderAccountResponse_Status) Enum() *SearchServiceProviderAccountResponse_Status {
	p := new(SearchServiceProviderAccountResponse_Status)
	*p = x
	return p
}

func (x SearchServiceProviderAccountResponse_Status) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (SearchServiceProviderAccountResponse_Status) Descriptor() protoreflect.EnumDescriptor {
	return file_serviceprovider_serviceprovideraccount_v1_service_provider_account_proto_enumTypes[1].Descriptor()
}

func (SearchServiceProviderAccountResponse_Status) Type() protoreflect.EnumType {
	return &file_serviceprovider_serviceprovideraccount_v1_service_provider_account_proto_enumTypes[1]
}

func (x SearchServiceProviderAccountResponse_Status) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use SearchServiceProviderAccountResponse_Status.Descriptor instead.
func (SearchServiceProviderAccountResponse_Status) EnumDescriptor() ([]byte, []int) {
	return file_serviceprovider_serviceprovideraccount_v1_service_provider_account_proto_rawDescGZIP(), []int{4, 0}
}

type ServiceProviderAccount struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ServiceProviderAccountKey string                    `protobuf:"bytes,1,opt,name=service_provider_account_key,json=serviceProviderAccountKey,proto3" json:"service_provider_account_key,omitempty"`
	Version                   int64                     `protobuf:"varint,2,opt,name=version,proto3" json:"version,omitempty"`
	Created                   *timestamppb.Timestamp    `protobuf:"bytes,3,opt,name=created,proto3" json:"created,omitempty"`
	Updated                   *timestamppb.Timestamp    `protobuf:"bytes,4,opt,name=updated,proto3" json:"updated,omitempty"`
	ServiceProviderKey        string                    `protobuf:"bytes,5,opt,name=service_provider_key,json=serviceProviderKey,proto3" json:"service_provider_key,omitempty"`
	AccountType               accounttypev1.AccountType `protobuf:"varint,6,opt,name=account_type,json=accountType,proto3,enum=common.enums.accounttype.v1.AccountType" json:"account_type,omitempty"`
	Currency                  currencyv1.Currency       `protobuf:"varint,7,opt,name=currency,proto3,enum=common.enums.currency.v1.Currency" json:"currency,omitempty"`
}

func (x *ServiceProviderAccount) Reset() {
	*x = ServiceProviderAccount{}
	if protoimpl.UnsafeEnabled {
		mi := &file_serviceprovider_serviceprovideraccount_v1_service_provider_account_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ServiceProviderAccount) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ServiceProviderAccount) ProtoMessage() {}

func (x *ServiceProviderAccount) ProtoReflect() protoreflect.Message {
	mi := &file_serviceprovider_serviceprovideraccount_v1_service_provider_account_proto_msgTypes[0]
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
	return file_serviceprovider_serviceprovideraccount_v1_service_provider_account_proto_rawDescGZIP(), []int{0}
}

func (x *ServiceProviderAccount) GetServiceProviderAccountKey() string {
	if x != nil {
		return x.ServiceProviderAccountKey
	}
	return ""
}

func (x *ServiceProviderAccount) GetVersion() int64 {
	if x != nil {
		return x.Version
	}
	return 0
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

func (x *ServiceProviderAccount) GetServiceProviderKey() string {
	if x != nil {
		return x.ServiceProviderKey
	}
	return ""
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

type GetServiceProviderAccountRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ServiceProviderAccountKey string `protobuf:"bytes,1,opt,name=service_provider_account_key,json=serviceProviderAccountKey,proto3" json:"service_provider_account_key,omitempty"`
}

func (x *GetServiceProviderAccountRequest) Reset() {
	*x = GetServiceProviderAccountRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_serviceprovider_serviceprovideraccount_v1_service_provider_account_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetServiceProviderAccountRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetServiceProviderAccountRequest) ProtoMessage() {}

func (x *GetServiceProviderAccountRequest) ProtoReflect() protoreflect.Message {
	mi := &file_serviceprovider_serviceprovideraccount_v1_service_provider_account_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetServiceProviderAccountRequest.ProtoReflect.Descriptor instead.
func (*GetServiceProviderAccountRequest) Descriptor() ([]byte, []int) {
	return file_serviceprovider_serviceprovideraccount_v1_service_provider_account_proto_rawDescGZIP(), []int{1}
}

func (x *GetServiceProviderAccountRequest) GetServiceProviderAccountKey() string {
	if x != nil {
		return x.ServiceProviderAccountKey
	}
	return ""
}

type GetServiceProviderAccountResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Status                 GetServiceProviderAccountResponse_Status `protobuf:"varint,1,opt,name=status,proto3,enum=serviceprovider.serviceprovideraccount.v1.GetServiceProviderAccountResponse_Status" json:"status,omitempty"`
	Error                  *v1.Error                                `protobuf:"bytes,2,opt,name=error,proto3" json:"error,omitempty"`
	ServiceProviderAccount *ServiceProviderAccount                  `protobuf:"bytes,3,opt,name=service_provider_account,json=serviceProviderAccount,proto3" json:"service_provider_account,omitempty"`
}

func (x *GetServiceProviderAccountResponse) Reset() {
	*x = GetServiceProviderAccountResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_serviceprovider_serviceprovideraccount_v1_service_provider_account_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetServiceProviderAccountResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetServiceProviderAccountResponse) ProtoMessage() {}

func (x *GetServiceProviderAccountResponse) ProtoReflect() protoreflect.Message {
	mi := &file_serviceprovider_serviceprovideraccount_v1_service_provider_account_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetServiceProviderAccountResponse.ProtoReflect.Descriptor instead.
func (*GetServiceProviderAccountResponse) Descriptor() ([]byte, []int) {
	return file_serviceprovider_serviceprovideraccount_v1_service_provider_account_proto_rawDescGZIP(), []int{2}
}

func (x *GetServiceProviderAccountResponse) GetStatus() GetServiceProviderAccountResponse_Status {
	if x != nil {
		return x.Status
	}
	return GetServiceProviderAccountResponse_STATUS_UNSPECIFIED
}

func (x *GetServiceProviderAccountResponse) GetError() *v1.Error {
	if x != nil {
		return x.Error
	}
	return nil
}

func (x *GetServiceProviderAccountResponse) GetServiceProviderAccount() *ServiceProviderAccount {
	if x != nil {
		return x.ServiceProviderAccount
	}
	return nil
}

type SearchServiceProviderAccountRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ServiceProviderKey string                    `protobuf:"bytes,1,opt,name=service_provider_key,json=serviceProviderKey,proto3" json:"service_provider_key,omitempty"`
	AccountType        accounttypev1.AccountType `protobuf:"varint,2,opt,name=account_type,json=accountType,proto3,enum=common.enums.accounttype.v1.AccountType" json:"account_type,omitempty"`
	Currency           currencyv1.Currency       `protobuf:"varint,3,opt,name=currency,proto3,enum=common.enums.currency.v1.Currency" json:"currency,omitempty"`
}

func (x *SearchServiceProviderAccountRequest) Reset() {
	*x = SearchServiceProviderAccountRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_serviceprovider_serviceprovideraccount_v1_service_provider_account_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SearchServiceProviderAccountRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SearchServiceProviderAccountRequest) ProtoMessage() {}

func (x *SearchServiceProviderAccountRequest) ProtoReflect() protoreflect.Message {
	mi := &file_serviceprovider_serviceprovideraccount_v1_service_provider_account_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SearchServiceProviderAccountRequest.ProtoReflect.Descriptor instead.
func (*SearchServiceProviderAccountRequest) Descriptor() ([]byte, []int) {
	return file_serviceprovider_serviceprovideraccount_v1_service_provider_account_proto_rawDescGZIP(), []int{3}
}

func (x *SearchServiceProviderAccountRequest) GetServiceProviderKey() string {
	if x != nil {
		return x.ServiceProviderKey
	}
	return ""
}

func (x *SearchServiceProviderAccountRequest) GetAccountType() accounttypev1.AccountType {
	if x != nil {
		return x.AccountType
	}
	return accounttypev1.AccountType(0)
}

func (x *SearchServiceProviderAccountRequest) GetCurrency() currencyv1.Currency {
	if x != nil {
		return x.Currency
	}
	return currencyv1.Currency(0)
}

type SearchServiceProviderAccountResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Status                  SearchServiceProviderAccountResponse_Status `protobuf:"varint,1,opt,name=status,proto3,enum=serviceprovider.serviceprovideraccount.v1.SearchServiceProviderAccountResponse_Status" json:"status,omitempty"`
	Error                   *v1.Error                                   `protobuf:"bytes,2,opt,name=error,proto3" json:"error,omitempty"`
	ServiceProviderAccounts []*ServiceProviderAccount                   `protobuf:"bytes,3,rep,name=service_provider_accounts,json=serviceProviderAccounts,proto3" json:"service_provider_accounts,omitempty"`
}

func (x *SearchServiceProviderAccountResponse) Reset() {
	*x = SearchServiceProviderAccountResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_serviceprovider_serviceprovideraccount_v1_service_provider_account_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SearchServiceProviderAccountResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SearchServiceProviderAccountResponse) ProtoMessage() {}

func (x *SearchServiceProviderAccountResponse) ProtoReflect() protoreflect.Message {
	mi := &file_serviceprovider_serviceprovideraccount_v1_service_provider_account_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SearchServiceProviderAccountResponse.ProtoReflect.Descriptor instead.
func (*SearchServiceProviderAccountResponse) Descriptor() ([]byte, []int) {
	return file_serviceprovider_serviceprovideraccount_v1_service_provider_account_proto_rawDescGZIP(), []int{4}
}

func (x *SearchServiceProviderAccountResponse) GetStatus() SearchServiceProviderAccountResponse_Status {
	if x != nil {
		return x.Status
	}
	return SearchServiceProviderAccountResponse_STATUS_UNSPECIFIED
}

func (x *SearchServiceProviderAccountResponse) GetError() *v1.Error {
	if x != nil {
		return x.Error
	}
	return nil
}

func (x *SearchServiceProviderAccountResponse) GetServiceProviderAccounts() []*ServiceProviderAccount {
	if x != nil {
		return x.ServiceProviderAccounts
	}
	return nil
}

var File_serviceprovider_serviceprovideraccount_v1_service_provider_account_proto protoreflect.FileDescriptor

var file_serviceprovider_serviceprovideraccount_v1_service_provider_account_proto_rawDesc = []byte{
	0x0a, 0x48, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x70, 0x72, 0x6f, 0x76, 0x69, 0x64, 0x65,
	0x72, 0x2f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x70, 0x72, 0x6f, 0x76, 0x69, 0x64, 0x65,
	0x72, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x2f, 0x76, 0x31, 0x2f, 0x73, 0x65, 0x72, 0x76,
	0x69, 0x63, 0x65, 0x5f, 0x70, 0x72, 0x6f, 0x76, 0x69, 0x64, 0x65, 0x72, 0x5f, 0x61, 0x63, 0x63,
	0x6f, 0x75, 0x6e, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x29, 0x73, 0x65, 0x72, 0x76,
	0x69, 0x63, 0x65, 0x70, 0x72, 0x6f, 0x76, 0x69, 0x64, 0x65, 0x72, 0x2e, 0x73, 0x65, 0x72, 0x76,
	0x69, 0x63, 0x65, 0x70, 0x72, 0x6f, 0x76, 0x69, 0x64, 0x65, 0x72, 0x61, 0x63, 0x63, 0x6f, 0x75,
	0x6e, 0x74, 0x2e, 0x76, 0x31, 0x1a, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2f, 0x65, 0x6e,
	0x75, 0x6d, 0x73, 0x2f, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x74, 0x79, 0x70, 0x65, 0x2f,
	0x76, 0x31, 0x2f, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x27, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2f, 0x65, 0x6e,
	0x75, 0x6d, 0x73, 0x2f, 0x63, 0x75, 0x72, 0x72, 0x65, 0x6e, 0x63, 0x79, 0x2f, 0x76, 0x31, 0x2f,
	0x63, 0x75, 0x72, 0x72, 0x65, 0x6e, 0x63, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x15,
	0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2f, 0x76, 0x31, 0x2f, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x17, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x65,
	0x2f, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22,
	0xbd, 0x03, 0x0a, 0x16, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x50, 0x72, 0x6f, 0x76, 0x69,
	0x64, 0x65, 0x72, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x12, 0x4a, 0x0a, 0x1c, 0x73, 0x65,
	0x72, 0x76, 0x69, 0x63, 0x65, 0x5f, 0x70, 0x72, 0x6f, 0x76, 0x69, 0x64, 0x65, 0x72, 0x5f, 0x61,
	0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x5f, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x42, 0x09, 0xfa, 0x42, 0x06, 0x72, 0x04, 0x10, 0x1e, 0x18, 0x32, 0x52, 0x19, 0x73, 0x65, 0x72,
	0x76, 0x69, 0x63, 0x65, 0x50, 0x72, 0x6f, 0x76, 0x69, 0x64, 0x65, 0x72, 0x41, 0x63, 0x63, 0x6f,
	0x75, 0x6e, 0x74, 0x4b, 0x65, 0x79, 0x12, 0x21, 0x0a, 0x07, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f,
	0x6e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x42, 0x07, 0xfa, 0x42, 0x04, 0x22, 0x02, 0x28, 0x00,
	0x52, 0x07, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x12, 0x34, 0x0a, 0x07, 0x63, 0x72, 0x65,
	0x61, 0x74, 0x65, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d,
	0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x07, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x12,
	0x34, 0x0a, 0x07, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62,
	0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x07, 0x75, 0x70,
	0x64, 0x61, 0x74, 0x65, 0x64, 0x12, 0x3b, 0x0a, 0x14, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65,
	0x5f, 0x70, 0x72, 0x6f, 0x76, 0x69, 0x64, 0x65, 0x72, 0x5f, 0x6b, 0x65, 0x79, 0x18, 0x05, 0x20,
	0x01, 0x28, 0x09, 0x42, 0x09, 0xfa, 0x42, 0x06, 0x72, 0x04, 0x10, 0x1e, 0x18, 0x32, 0x52, 0x12,
	0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x50, 0x72, 0x6f, 0x76, 0x69, 0x64, 0x65, 0x72, 0x4b,
	0x65, 0x79, 0x12, 0x4b, 0x0a, 0x0c, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x5f, 0x74, 0x79,
	0x70, 0x65, 0x18, 0x06, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x28, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f,
	0x6e, 0x2e, 0x65, 0x6e, 0x75, 0x6d, 0x73, 0x2e, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x74,
	0x79, 0x70, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x54, 0x79,
	0x70, 0x65, 0x52, 0x0b, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x54, 0x79, 0x70, 0x65, 0x12,
	0x3e, 0x0a, 0x08, 0x63, 0x75, 0x72, 0x72, 0x65, 0x6e, 0x63, 0x79, 0x18, 0x07, 0x20, 0x01, 0x28,
	0x0e, 0x32, 0x22, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x65, 0x6e, 0x75, 0x6d, 0x73,
	0x2e, 0x63, 0x75, 0x72, 0x72, 0x65, 0x6e, 0x63, 0x79, 0x2e, 0x76, 0x31, 0x2e, 0x43, 0x75, 0x72,
	0x72, 0x65, 0x6e, 0x63, 0x79, 0x52, 0x08, 0x63, 0x75, 0x72, 0x72, 0x65, 0x6e, 0x63, 0x79, 0x22,
	0x6e, 0x0a, 0x20, 0x47, 0x65, 0x74, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x50, 0x72, 0x6f,
	0x76, 0x69, 0x64, 0x65, 0x72, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x12, 0x4a, 0x0a, 0x1c, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x5f, 0x70,
	0x72, 0x6f, 0x76, 0x69, 0x64, 0x65, 0x72, 0x5f, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x5f,
	0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x42, 0x09, 0xfa, 0x42, 0x06, 0x72, 0x04,
	0x10, 0x1e, 0x18, 0x32, 0x52, 0x19, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x50, 0x72, 0x6f,
	0x76, 0x69, 0x64, 0x65, 0x72, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x4b, 0x65, 0x79, 0x22,
	0x93, 0x03, 0x0a, 0x21, 0x47, 0x65, 0x74, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x50, 0x72,
	0x6f, 0x76, 0x69, 0x64, 0x65, 0x72, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x6b, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x53, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x70,
	0x72, 0x6f, 0x76, 0x69, 0x64, 0x65, 0x72, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x70,
	0x72, 0x6f, 0x76, 0x69, 0x64, 0x65, 0x72, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x2e, 0x76,
	0x31, 0x2e, 0x47, 0x65, 0x74, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x50, 0x72, 0x6f, 0x76,
	0x69, 0x64, 0x65, 0x72, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x2e, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x52, 0x06, 0x73, 0x74, 0x61, 0x74,
	0x75, 0x73, 0x12, 0x26, 0x0a, 0x05, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x10, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x76, 0x31, 0x2e, 0x45, 0x72,
	0x72, 0x6f, 0x72, 0x52, 0x05, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x12, 0x7b, 0x0a, 0x18, 0x73, 0x65,
	0x72, 0x76, 0x69, 0x63, 0x65, 0x5f, 0x70, 0x72, 0x6f, 0x76, 0x69, 0x64, 0x65, 0x72, 0x5f, 0x61,
	0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x41, 0x2e, 0x73,
	0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x70, 0x72, 0x6f, 0x76, 0x69, 0x64, 0x65, 0x72, 0x2e, 0x73,
	0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x70, 0x72, 0x6f, 0x76, 0x69, 0x64, 0x65, 0x72, 0x61, 0x63,
	0x63, 0x6f, 0x75, 0x6e, 0x74, 0x2e, 0x76, 0x31, 0x2e, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65,
	0x50, 0x72, 0x6f, 0x76, 0x69, 0x64, 0x65, 0x72, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x52,
	0x16, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x50, 0x72, 0x6f, 0x76, 0x69, 0x64, 0x65, 0x72,
	0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x22, 0x5c, 0x0a, 0x06, 0x53, 0x74, 0x61, 0x74, 0x75,
	0x73, 0x12, 0x16, 0x0a, 0x12, 0x53, 0x54, 0x41, 0x54, 0x55, 0x53, 0x5f, 0x55, 0x4e, 0x53, 0x50,
	0x45, 0x43, 0x49, 0x46, 0x49, 0x45, 0x44, 0x10, 0x00, 0x12, 0x12, 0x0a, 0x0e, 0x53, 0x54, 0x41,
	0x54, 0x55, 0x53, 0x5f, 0x53, 0x55, 0x43, 0x43, 0x45, 0x53, 0x53, 0x10, 0x01, 0x12, 0x10, 0x0a,
	0x0c, 0x53, 0x54, 0x41, 0x54, 0x55, 0x53, 0x5f, 0x45, 0x52, 0x52, 0x4f, 0x52, 0x10, 0x02, 0x12,
	0x14, 0x0a, 0x10, 0x53, 0x54, 0x41, 0x54, 0x55, 0x53, 0x5f, 0x4e, 0x4f, 0x54, 0x5f, 0x46, 0x4f,
	0x55, 0x4e, 0x44, 0x10, 0x03, 0x22, 0xef, 0x01, 0x0a, 0x23, 0x53, 0x65, 0x61, 0x72, 0x63, 0x68,
	0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x50, 0x72, 0x6f, 0x76, 0x69, 0x64, 0x65, 0x72, 0x41,
	0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x3b, 0x0a,
	0x14, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x5f, 0x70, 0x72, 0x6f, 0x76, 0x69, 0x64, 0x65,
	0x72, 0x5f, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x42, 0x09, 0xfa, 0x42, 0x06,
	0x72, 0x04, 0x10, 0x1e, 0x18, 0x32, 0x52, 0x12, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x50,
	0x72, 0x6f, 0x76, 0x69, 0x64, 0x65, 0x72, 0x4b, 0x65, 0x79, 0x12, 0x4b, 0x0a, 0x0c, 0x61, 0x63,
	0x63, 0x6f, 0x75, 0x6e, 0x74, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0e,
	0x32, 0x28, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x65, 0x6e, 0x75, 0x6d, 0x73, 0x2e,
	0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x74, 0x79, 0x70, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x41,
	0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x54, 0x79, 0x70, 0x65, 0x52, 0x0b, 0x61, 0x63, 0x63, 0x6f,
	0x75, 0x6e, 0x74, 0x54, 0x79, 0x70, 0x65, 0x12, 0x3e, 0x0a, 0x08, 0x63, 0x75, 0x72, 0x72, 0x65,
	0x6e, 0x63, 0x79, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x22, 0x2e, 0x63, 0x6f, 0x6d, 0x6d,
	0x6f, 0x6e, 0x2e, 0x65, 0x6e, 0x75, 0x6d, 0x73, 0x2e, 0x63, 0x75, 0x72, 0x72, 0x65, 0x6e, 0x63,
	0x79, 0x2e, 0x76, 0x31, 0x2e, 0x43, 0x75, 0x72, 0x72, 0x65, 0x6e, 0x63, 0x79, 0x52, 0x08, 0x63,
	0x75, 0x72, 0x72, 0x65, 0x6e, 0x63, 0x79, 0x22, 0x85, 0x03, 0x0a, 0x24, 0x53, 0x65, 0x61, 0x72,
	0x63, 0x68, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x50, 0x72, 0x6f, 0x76, 0x69, 0x64, 0x65,
	0x72, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x12, 0x6e, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0e,
	0x32, 0x56, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x70, 0x72, 0x6f, 0x76, 0x69, 0x64,
	0x65, 0x72, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x70, 0x72, 0x6f, 0x76, 0x69, 0x64,
	0x65, 0x72, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x2e, 0x76, 0x31, 0x2e, 0x53, 0x65, 0x61,
	0x72, 0x63, 0x68, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x50, 0x72, 0x6f, 0x76, 0x69, 0x64,
	0x65, 0x72, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x2e, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x52, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73,
	0x12, 0x26, 0x0a, 0x05, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x10, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x76, 0x31, 0x2e, 0x45, 0x72, 0x72, 0x6f,
	0x72, 0x52, 0x05, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x12, 0x7d, 0x0a, 0x19, 0x73, 0x65, 0x72, 0x76,
	0x69, 0x63, 0x65, 0x5f, 0x70, 0x72, 0x6f, 0x76, 0x69, 0x64, 0x65, 0x72, 0x5f, 0x61, 0x63, 0x63,
	0x6f, 0x75, 0x6e, 0x74, 0x73, 0x18, 0x03, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x41, 0x2e, 0x73, 0x65,
	0x72, 0x76, 0x69, 0x63, 0x65, 0x70, 0x72, 0x6f, 0x76, 0x69, 0x64, 0x65, 0x72, 0x2e, 0x73, 0x65,
	0x72, 0x76, 0x69, 0x63, 0x65, 0x70, 0x72, 0x6f, 0x76, 0x69, 0x64, 0x65, 0x72, 0x61, 0x63, 0x63,
	0x6f, 0x75, 0x6e, 0x74, 0x2e, 0x76, 0x31, 0x2e, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x50,
	0x72, 0x6f, 0x76, 0x69, 0x64, 0x65, 0x72, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x52, 0x17,
	0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x50, 0x72, 0x6f, 0x76, 0x69, 0x64, 0x65, 0x72, 0x41,
	0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x73, 0x22, 0x46, 0x0a, 0x06, 0x53, 0x74, 0x61, 0x74, 0x75,
	0x73, 0x12, 0x16, 0x0a, 0x12, 0x53, 0x54, 0x41, 0x54, 0x55, 0x53, 0x5f, 0x55, 0x4e, 0x53, 0x50,
	0x45, 0x43, 0x49, 0x46, 0x49, 0x45, 0x44, 0x10, 0x00, 0x12, 0x12, 0x0a, 0x0e, 0x53, 0x54, 0x41,
	0x54, 0x55, 0x53, 0x5f, 0x53, 0x55, 0x43, 0x43, 0x45, 0x53, 0x53, 0x10, 0x01, 0x12, 0x10, 0x0a,
	0x0c, 0x53, 0x54, 0x41, 0x54, 0x55, 0x53, 0x5f, 0x45, 0x52, 0x52, 0x4f, 0x52, 0x10, 0x02, 0x42,
	0x8c, 0x01, 0x0a, 0x31, 0x69, 0x6f, 0x2e, 0x6d, 0x69, 0x63, 0x61, 0x2e, 0x73, 0x65, 0x72, 0x76,
	0x69, 0x63, 0x65, 0x70, 0x72, 0x6f, 0x76, 0x69, 0x64, 0x65, 0x72, 0x2e, 0x73, 0x65, 0x72, 0x76,
	0x69, 0x63, 0x65, 0x70, 0x72, 0x6f, 0x76, 0x69, 0x64, 0x65, 0x72, 0x61, 0x63, 0x63, 0x6f, 0x75,
	0x6e, 0x74, 0x2e, 0x76, 0x31, 0x42, 0x1b, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x50, 0x72,
	0x6f, 0x76, 0x69, 0x64, 0x65, 0x72, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x50, 0x72, 0x6f,
	0x74, 0x6f, 0x5a, 0x33, 0x6d, 0x69, 0x63, 0x61, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x73,
	0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x70, 0x72, 0x6f, 0x76, 0x69, 0x64, 0x65, 0x72, 0x2f, 0x73,
	0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x70, 0x72, 0x6f, 0x76, 0x69, 0x64, 0x65, 0x72, 0x61, 0x63,
	0x63, 0x6f, 0x75, 0x6e, 0x74, 0x76, 0x31, 0xa2, 0x02, 0x04, 0x4d, 0x49, 0x43, 0x41, 0x62, 0x06,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_serviceprovider_serviceprovideraccount_v1_service_provider_account_proto_rawDescOnce sync.Once
	file_serviceprovider_serviceprovideraccount_v1_service_provider_account_proto_rawDescData = file_serviceprovider_serviceprovideraccount_v1_service_provider_account_proto_rawDesc
)

func file_serviceprovider_serviceprovideraccount_v1_service_provider_account_proto_rawDescGZIP() []byte {
	file_serviceprovider_serviceprovideraccount_v1_service_provider_account_proto_rawDescOnce.Do(func() {
		file_serviceprovider_serviceprovideraccount_v1_service_provider_account_proto_rawDescData = protoimpl.X.CompressGZIP(file_serviceprovider_serviceprovideraccount_v1_service_provider_account_proto_rawDescData)
	})
	return file_serviceprovider_serviceprovideraccount_v1_service_provider_account_proto_rawDescData
}

var file_serviceprovider_serviceprovideraccount_v1_service_provider_account_proto_enumTypes = make([]protoimpl.EnumInfo, 2)
var file_serviceprovider_serviceprovideraccount_v1_service_provider_account_proto_msgTypes = make([]protoimpl.MessageInfo, 5)
var file_serviceprovider_serviceprovideraccount_v1_service_provider_account_proto_goTypes = []interface{}{
	(GetServiceProviderAccountResponse_Status)(0),    // 0: serviceprovider.serviceprovideraccount.v1.GetServiceProviderAccountResponse.Status
	(SearchServiceProviderAccountResponse_Status)(0), // 1: serviceprovider.serviceprovideraccount.v1.SearchServiceProviderAccountResponse.Status
	(*ServiceProviderAccount)(nil),                   // 2: serviceprovider.serviceprovideraccount.v1.ServiceProviderAccount
	(*GetServiceProviderAccountRequest)(nil),         // 3: serviceprovider.serviceprovideraccount.v1.GetServiceProviderAccountRequest
	(*GetServiceProviderAccountResponse)(nil),        // 4: serviceprovider.serviceprovideraccount.v1.GetServiceProviderAccountResponse
	(*SearchServiceProviderAccountRequest)(nil),      // 5: serviceprovider.serviceprovideraccount.v1.SearchServiceProviderAccountRequest
	(*SearchServiceProviderAccountResponse)(nil),     // 6: serviceprovider.serviceprovideraccount.v1.SearchServiceProviderAccountResponse
	(*timestamppb.Timestamp)(nil),                    // 7: google.protobuf.Timestamp
	(accounttypev1.AccountType)(0),                   // 8: common.enums.accounttype.v1.AccountType
	(currencyv1.Currency)(0),                         // 9: common.enums.currency.v1.Currency
	(*v1.Error)(nil),                                 // 10: common.v1.Error
}
var file_serviceprovider_serviceprovideraccount_v1_service_provider_account_proto_depIdxs = []int32{
	7,  // 0: serviceprovider.serviceprovideraccount.v1.ServiceProviderAccount.created:type_name -> google.protobuf.Timestamp
	7,  // 1: serviceprovider.serviceprovideraccount.v1.ServiceProviderAccount.updated:type_name -> google.protobuf.Timestamp
	8,  // 2: serviceprovider.serviceprovideraccount.v1.ServiceProviderAccount.account_type:type_name -> common.enums.accounttype.v1.AccountType
	9,  // 3: serviceprovider.serviceprovideraccount.v1.ServiceProviderAccount.currency:type_name -> common.enums.currency.v1.Currency
	0,  // 4: serviceprovider.serviceprovideraccount.v1.GetServiceProviderAccountResponse.status:type_name -> serviceprovider.serviceprovideraccount.v1.GetServiceProviderAccountResponse.Status
	10, // 5: serviceprovider.serviceprovideraccount.v1.GetServiceProviderAccountResponse.error:type_name -> common.v1.Error
	2,  // 6: serviceprovider.serviceprovideraccount.v1.GetServiceProviderAccountResponse.service_provider_account:type_name -> serviceprovider.serviceprovideraccount.v1.ServiceProviderAccount
	8,  // 7: serviceprovider.serviceprovideraccount.v1.SearchServiceProviderAccountRequest.account_type:type_name -> common.enums.accounttype.v1.AccountType
	9,  // 8: serviceprovider.serviceprovideraccount.v1.SearchServiceProviderAccountRequest.currency:type_name -> common.enums.currency.v1.Currency
	1,  // 9: serviceprovider.serviceprovideraccount.v1.SearchServiceProviderAccountResponse.status:type_name -> serviceprovider.serviceprovideraccount.v1.SearchServiceProviderAccountResponse.Status
	10, // 10: serviceprovider.serviceprovideraccount.v1.SearchServiceProviderAccountResponse.error:type_name -> common.v1.Error
	2,  // 11: serviceprovider.serviceprovideraccount.v1.SearchServiceProviderAccountResponse.service_provider_accounts:type_name -> serviceprovider.serviceprovideraccount.v1.ServiceProviderAccount
	12, // [12:12] is the sub-list for method output_type
	12, // [12:12] is the sub-list for method input_type
	12, // [12:12] is the sub-list for extension type_name
	12, // [12:12] is the sub-list for extension extendee
	0,  // [0:12] is the sub-list for field type_name
}

func init() { file_serviceprovider_serviceprovideraccount_v1_service_provider_account_proto_init() }
func file_serviceprovider_serviceprovideraccount_v1_service_provider_account_proto_init() {
	if File_serviceprovider_serviceprovideraccount_v1_service_provider_account_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_serviceprovider_serviceprovideraccount_v1_service_provider_account_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
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
		file_serviceprovider_serviceprovideraccount_v1_service_provider_account_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetServiceProviderAccountRequest); i {
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
		file_serviceprovider_serviceprovideraccount_v1_service_provider_account_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetServiceProviderAccountResponse); i {
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
		file_serviceprovider_serviceprovideraccount_v1_service_provider_account_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SearchServiceProviderAccountRequest); i {
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
		file_serviceprovider_serviceprovideraccount_v1_service_provider_account_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SearchServiceProviderAccountResponse); i {
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
			RawDescriptor: file_serviceprovider_serviceprovideraccount_v1_service_provider_account_proto_rawDesc,
			NumEnums:      2,
			NumMessages:   5,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_serviceprovider_serviceprovideraccount_v1_service_provider_account_proto_goTypes,
		DependencyIndexes: file_serviceprovider_serviceprovideraccount_v1_service_provider_account_proto_depIdxs,
		EnumInfos:         file_serviceprovider_serviceprovideraccount_v1_service_provider_account_proto_enumTypes,
		MessageInfos:      file_serviceprovider_serviceprovideraccount_v1_service_provider_account_proto_msgTypes,
	}.Build()
	File_serviceprovider_serviceprovideraccount_v1_service_provider_account_proto = out.File
	file_serviceprovider_serviceprovideraccount_v1_service_provider_account_proto_rawDesc = nil
	file_serviceprovider_serviceprovideraccount_v1_service_provider_account_proto_goTypes = nil
	file_serviceprovider_serviceprovideraccount_v1_service_provider_account_proto_depIdxs = nil
}
