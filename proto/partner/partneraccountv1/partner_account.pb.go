// Copyright (c) 2022 Mica, Inc. All rights reserved. All software, including, without limitation, all source
// code and object code, is the intellectual property of Mica, Inc. and is protected by copyright, trademark and
// other intellectual property laws (collective "Mica Software"). You may not use, copy, reproduce, download, store,
// post, broadcast, transmit, modify, sell or make available to the public content from the Mica Software without a
// valid license or the prior written approval of Mica, Inc. Mica, its logos, slogans, taglines,
// products, feature names, and other trademarks are trademarks of Mica, Inc. and may not be used without
// permission.

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.1
// 	protoc        v3.19.0
// source: partner/partneraccount/v1/partner_account.proto

package partneraccountv1

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	timestamppb "google.golang.org/protobuf/types/known/timestamppb"
	accounttypev1 "mica/proto/common/enums/accounttypev1"
	currencyv1 "mica/proto/common/enums/currencyv1"
	v1 "mica/proto/common/v1"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type GetPartnerAccountResponse_Status int32

const (
	GetPartnerAccountResponse_STATUS_UNSPECIFIED GetPartnerAccountResponse_Status = 0
	GetPartnerAccountResponse_STATUS_SUCCESS     GetPartnerAccountResponse_Status = 1
	GetPartnerAccountResponse_STATUS_ERROR       GetPartnerAccountResponse_Status = 2
	GetPartnerAccountResponse_STATUS_NOT_FOUND   GetPartnerAccountResponse_Status = 3
)

// Enum value maps for GetPartnerAccountResponse_Status.
var (
	GetPartnerAccountResponse_Status_name = map[int32]string{
		0: "STATUS_UNSPECIFIED",
		1: "STATUS_SUCCESS",
		2: "STATUS_ERROR",
		3: "STATUS_NOT_FOUND",
	}
	GetPartnerAccountResponse_Status_value = map[string]int32{
		"STATUS_UNSPECIFIED": 0,
		"STATUS_SUCCESS":     1,
		"STATUS_ERROR":       2,
		"STATUS_NOT_FOUND":   3,
	}
)

func (x GetPartnerAccountResponse_Status) Enum() *GetPartnerAccountResponse_Status {
	p := new(GetPartnerAccountResponse_Status)
	*p = x
	return p
}

func (x GetPartnerAccountResponse_Status) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (GetPartnerAccountResponse_Status) Descriptor() protoreflect.EnumDescriptor {
	return file_partner_partneraccount_v1_partner_account_proto_enumTypes[0].Descriptor()
}

func (GetPartnerAccountResponse_Status) Type() protoreflect.EnumType {
	return &file_partner_partneraccount_v1_partner_account_proto_enumTypes[0]
}

func (x GetPartnerAccountResponse_Status) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use GetPartnerAccountResponse_Status.Descriptor instead.
func (GetPartnerAccountResponse_Status) EnumDescriptor() ([]byte, []int) {
	return file_partner_partneraccount_v1_partner_account_proto_rawDescGZIP(), []int{2, 0}
}

type SearchPartnerAccountResponse_Status int32

const (
	SearchPartnerAccountResponse_STATUS_UNSPECIFIED SearchPartnerAccountResponse_Status = 0
	SearchPartnerAccountResponse_STATUS_SUCCESS     SearchPartnerAccountResponse_Status = 1
	SearchPartnerAccountResponse_STATUS_ERROR       SearchPartnerAccountResponse_Status = 2
)

// Enum value maps for SearchPartnerAccountResponse_Status.
var (
	SearchPartnerAccountResponse_Status_name = map[int32]string{
		0: "STATUS_UNSPECIFIED",
		1: "STATUS_SUCCESS",
		2: "STATUS_ERROR",
	}
	SearchPartnerAccountResponse_Status_value = map[string]int32{
		"STATUS_UNSPECIFIED": 0,
		"STATUS_SUCCESS":     1,
		"STATUS_ERROR":       2,
	}
)

func (x SearchPartnerAccountResponse_Status) Enum() *SearchPartnerAccountResponse_Status {
	p := new(SearchPartnerAccountResponse_Status)
	*p = x
	return p
}

func (x SearchPartnerAccountResponse_Status) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (SearchPartnerAccountResponse_Status) Descriptor() protoreflect.EnumDescriptor {
	return file_partner_partneraccount_v1_partner_account_proto_enumTypes[1].Descriptor()
}

func (SearchPartnerAccountResponse_Status) Type() protoreflect.EnumType {
	return &file_partner_partneraccount_v1_partner_account_proto_enumTypes[1]
}

func (x SearchPartnerAccountResponse_Status) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use SearchPartnerAccountResponse_Status.Descriptor instead.
func (SearchPartnerAccountResponse_Status) EnumDescriptor() ([]byte, []int) {
	return file_partner_partneraccount_v1_partner_account_proto_rawDescGZIP(), []int{4, 0}
}

type PartnerAccount struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	PartnerAccountKey string                    `protobuf:"bytes,1,opt,name=partner_account_key,json=partnerAccountKey,proto3" json:"partner_account_key,omitempty"`
	Created           *timestamppb.Timestamp    `protobuf:"bytes,2,opt,name=created,proto3" json:"created,omitempty"`
	Updated           *timestamppb.Timestamp    `protobuf:"bytes,3,opt,name=updated,proto3" json:"updated,omitempty"`
	Version           int64                     `protobuf:"varint,4,opt,name=version,proto3" json:"version,omitempty"`
	PartnerKey        string                    `protobuf:"bytes,5,opt,name=partner_key,json=partnerKey,proto3" json:"partner_key,omitempty"`
	AccountType       accounttypev1.AccountType `protobuf:"varint,6,opt,name=account_type,json=accountType,proto3,enum=common.enums.accounttype.v1.AccountType" json:"account_type,omitempty"`
	Currency          currencyv1.Currency       `protobuf:"varint,7,opt,name=currency,proto3,enum=common.enums.currency.v1.Currency" json:"currency,omitempty"`
}

func (x *PartnerAccount) Reset() {
	*x = PartnerAccount{}
	if protoimpl.UnsafeEnabled {
		mi := &file_partner_partneraccount_v1_partner_account_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PartnerAccount) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PartnerAccount) ProtoMessage() {}

func (x *PartnerAccount) ProtoReflect() protoreflect.Message {
	mi := &file_partner_partneraccount_v1_partner_account_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PartnerAccount.ProtoReflect.Descriptor instead.
func (*PartnerAccount) Descriptor() ([]byte, []int) {
	return file_partner_partneraccount_v1_partner_account_proto_rawDescGZIP(), []int{0}
}

func (x *PartnerAccount) GetPartnerAccountKey() string {
	if x != nil {
		return x.PartnerAccountKey
	}
	return ""
}

func (x *PartnerAccount) GetCreated() *timestamppb.Timestamp {
	if x != nil {
		return x.Created
	}
	return nil
}

func (x *PartnerAccount) GetUpdated() *timestamppb.Timestamp {
	if x != nil {
		return x.Updated
	}
	return nil
}

func (x *PartnerAccount) GetVersion() int64 {
	if x != nil {
		return x.Version
	}
	return 0
}

func (x *PartnerAccount) GetPartnerKey() string {
	if x != nil {
		return x.PartnerKey
	}
	return ""
}

func (x *PartnerAccount) GetAccountType() accounttypev1.AccountType {
	if x != nil {
		return x.AccountType
	}
	return accounttypev1.AccountType(0)
}

func (x *PartnerAccount) GetCurrency() currencyv1.Currency {
	if x != nil {
		return x.Currency
	}
	return currencyv1.Currency(0)
}

type GetPartnerAccountRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	PartnerAccountKey string `protobuf:"bytes,1,opt,name=partner_account_key,json=partnerAccountKey,proto3" json:"partner_account_key,omitempty"`
}

func (x *GetPartnerAccountRequest) Reset() {
	*x = GetPartnerAccountRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_partner_partneraccount_v1_partner_account_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetPartnerAccountRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetPartnerAccountRequest) ProtoMessage() {}

func (x *GetPartnerAccountRequest) ProtoReflect() protoreflect.Message {
	mi := &file_partner_partneraccount_v1_partner_account_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetPartnerAccountRequest.ProtoReflect.Descriptor instead.
func (*GetPartnerAccountRequest) Descriptor() ([]byte, []int) {
	return file_partner_partneraccount_v1_partner_account_proto_rawDescGZIP(), []int{1}
}

func (x *GetPartnerAccountRequest) GetPartnerAccountKey() string {
	if x != nil {
		return x.PartnerAccountKey
	}
	return ""
}

type GetPartnerAccountResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Status         GetPartnerAccountResponse_Status `protobuf:"varint,1,opt,name=status,proto3,enum=partner.partneraccount.v1.GetPartnerAccountResponse_Status" json:"status,omitempty"`
	Error          *v1.Error                        `protobuf:"bytes,2,opt,name=error,proto3" json:"error,omitempty"`
	PartnerAccount *PartnerAccount                  `protobuf:"bytes,3,opt,name=partner_account,json=partnerAccount,proto3" json:"partner_account,omitempty"`
}

func (x *GetPartnerAccountResponse) Reset() {
	*x = GetPartnerAccountResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_partner_partneraccount_v1_partner_account_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetPartnerAccountResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetPartnerAccountResponse) ProtoMessage() {}

func (x *GetPartnerAccountResponse) ProtoReflect() protoreflect.Message {
	mi := &file_partner_partneraccount_v1_partner_account_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetPartnerAccountResponse.ProtoReflect.Descriptor instead.
func (*GetPartnerAccountResponse) Descriptor() ([]byte, []int) {
	return file_partner_partneraccount_v1_partner_account_proto_rawDescGZIP(), []int{2}
}

func (x *GetPartnerAccountResponse) GetStatus() GetPartnerAccountResponse_Status {
	if x != nil {
		return x.Status
	}
	return GetPartnerAccountResponse_STATUS_UNSPECIFIED
}

func (x *GetPartnerAccountResponse) GetError() *v1.Error {
	if x != nil {
		return x.Error
	}
	return nil
}

func (x *GetPartnerAccountResponse) GetPartnerAccount() *PartnerAccount {
	if x != nil {
		return x.PartnerAccount
	}
	return nil
}

type SearchPartnerAccountRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	AccountType accounttypev1.AccountType `protobuf:"varint,1,opt,name=account_type,json=accountType,proto3,enum=common.enums.accounttype.v1.AccountType" json:"account_type,omitempty"`
	Currency    currencyv1.Currency       `protobuf:"varint,2,opt,name=currency,proto3,enum=common.enums.currency.v1.Currency" json:"currency,omitempty"`
}

func (x *SearchPartnerAccountRequest) Reset() {
	*x = SearchPartnerAccountRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_partner_partneraccount_v1_partner_account_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SearchPartnerAccountRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SearchPartnerAccountRequest) ProtoMessage() {}

func (x *SearchPartnerAccountRequest) ProtoReflect() protoreflect.Message {
	mi := &file_partner_partneraccount_v1_partner_account_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SearchPartnerAccountRequest.ProtoReflect.Descriptor instead.
func (*SearchPartnerAccountRequest) Descriptor() ([]byte, []int) {
	return file_partner_partneraccount_v1_partner_account_proto_rawDescGZIP(), []int{3}
}

func (x *SearchPartnerAccountRequest) GetAccountType() accounttypev1.AccountType {
	if x != nil {
		return x.AccountType
	}
	return accounttypev1.AccountType(0)
}

func (x *SearchPartnerAccountRequest) GetCurrency() currencyv1.Currency {
	if x != nil {
		return x.Currency
	}
	return currencyv1.Currency(0)
}

type SearchPartnerAccountResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Status          SearchPartnerAccountResponse_Status `protobuf:"varint,1,opt,name=status,proto3,enum=partner.partneraccount.v1.SearchPartnerAccountResponse_Status" json:"status,omitempty"`
	Error           *v1.Error                           `protobuf:"bytes,2,opt,name=error,proto3" json:"error,omitempty"`
	PartnerAccounts []*PartnerAccount                   `protobuf:"bytes,3,rep,name=partner_accounts,json=partnerAccounts,proto3" json:"partner_accounts,omitempty"`
}

func (x *SearchPartnerAccountResponse) Reset() {
	*x = SearchPartnerAccountResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_partner_partneraccount_v1_partner_account_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SearchPartnerAccountResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SearchPartnerAccountResponse) ProtoMessage() {}

func (x *SearchPartnerAccountResponse) ProtoReflect() protoreflect.Message {
	mi := &file_partner_partneraccount_v1_partner_account_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SearchPartnerAccountResponse.ProtoReflect.Descriptor instead.
func (*SearchPartnerAccountResponse) Descriptor() ([]byte, []int) {
	return file_partner_partneraccount_v1_partner_account_proto_rawDescGZIP(), []int{4}
}

func (x *SearchPartnerAccountResponse) GetStatus() SearchPartnerAccountResponse_Status {
	if x != nil {
		return x.Status
	}
	return SearchPartnerAccountResponse_STATUS_UNSPECIFIED
}

func (x *SearchPartnerAccountResponse) GetError() *v1.Error {
	if x != nil {
		return x.Error
	}
	return nil
}

func (x *SearchPartnerAccountResponse) GetPartnerAccounts() []*PartnerAccount {
	if x != nil {
		return x.PartnerAccounts
	}
	return nil
}

var File_partner_partneraccount_v1_partner_account_proto protoreflect.FileDescriptor

var file_partner_partneraccount_v1_partner_account_proto_rawDesc = []byte{
	0x0a, 0x2f, 0x70, 0x61, 0x72, 0x74, 0x6e, 0x65, 0x72, 0x2f, 0x70, 0x61, 0x72, 0x74, 0x6e, 0x65,
	0x72, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x2f, 0x76, 0x31, 0x2f, 0x70, 0x61, 0x72, 0x74,
	0x6e, 0x65, 0x72, 0x5f, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x12, 0x19, 0x70, 0x61, 0x72, 0x74, 0x6e, 0x65, 0x72, 0x2e, 0x70, 0x61, 0x72, 0x74, 0x6e,
	0x65, 0x72, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x2e, 0x76, 0x31, 0x1a, 0x2e, 0x63, 0x6f,
	0x6d, 0x6d, 0x6f, 0x6e, 0x2f, 0x65, 0x6e, 0x75, 0x6d, 0x73, 0x2f, 0x61, 0x63, 0x63, 0x6f, 0x75,
	0x6e, 0x74, 0x74, 0x79, 0x70, 0x65, 0x2f, 0x76, 0x31, 0x2f, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e,
	0x74, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x27, 0x63, 0x6f,
	0x6d, 0x6d, 0x6f, 0x6e, 0x2f, 0x65, 0x6e, 0x75, 0x6d, 0x73, 0x2f, 0x63, 0x75, 0x72, 0x72, 0x65,
	0x6e, 0x63, 0x79, 0x2f, 0x76, 0x31, 0x2f, 0x63, 0x75, 0x72, 0x72, 0x65, 0x6e, 0x63, 0x79, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x15, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2f, 0x76, 0x31,
	0x2f, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1f, 0x67, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x74, 0x69,
	0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xf4, 0x02,
	0x0a, 0x0e, 0x50, 0x61, 0x72, 0x74, 0x6e, 0x65, 0x72, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74,
	0x12, 0x2e, 0x0a, 0x13, 0x70, 0x61, 0x72, 0x74, 0x6e, 0x65, 0x72, 0x5f, 0x61, 0x63, 0x63, 0x6f,
	0x75, 0x6e, 0x74, 0x5f, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x11, 0x70,
	0x61, 0x72, 0x74, 0x6e, 0x65, 0x72, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x4b, 0x65, 0x79,
	0x12, 0x34, 0x0a, 0x07, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x07, 0x63,
	0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x12, 0x34, 0x0a, 0x07, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65,
	0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74,
	0x61, 0x6d, 0x70, 0x52, 0x07, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x12, 0x18, 0x0a, 0x07,
	0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x18, 0x04, 0x20, 0x01, 0x28, 0x03, 0x52, 0x07, 0x76,
	0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x12, 0x1f, 0x0a, 0x0b, 0x70, 0x61, 0x72, 0x74, 0x6e, 0x65,
	0x72, 0x5f, 0x6b, 0x65, 0x79, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x70, 0x61, 0x72,
	0x74, 0x6e, 0x65, 0x72, 0x4b, 0x65, 0x79, 0x12, 0x4b, 0x0a, 0x0c, 0x61, 0x63, 0x63, 0x6f, 0x75,
	0x6e, 0x74, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x18, 0x06, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x28, 0x2e,
	0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x65, 0x6e, 0x75, 0x6d, 0x73, 0x2e, 0x61, 0x63, 0x63,
	0x6f, 0x75, 0x6e, 0x74, 0x74, 0x79, 0x70, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x41, 0x63, 0x63, 0x6f,
	0x75, 0x6e, 0x74, 0x54, 0x79, 0x70, 0x65, 0x52, 0x0b, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74,
	0x54, 0x79, 0x70, 0x65, 0x12, 0x3e, 0x0a, 0x08, 0x63, 0x75, 0x72, 0x72, 0x65, 0x6e, 0x63, 0x79,
	0x18, 0x07, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x22, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e,
	0x65, 0x6e, 0x75, 0x6d, 0x73, 0x2e, 0x63, 0x75, 0x72, 0x72, 0x65, 0x6e, 0x63, 0x79, 0x2e, 0x76,
	0x31, 0x2e, 0x43, 0x75, 0x72, 0x72, 0x65, 0x6e, 0x63, 0x79, 0x52, 0x08, 0x63, 0x75, 0x72, 0x72,
	0x65, 0x6e, 0x63, 0x79, 0x22, 0x4a, 0x0a, 0x18, 0x47, 0x65, 0x74, 0x50, 0x61, 0x72, 0x74, 0x6e,
	0x65, 0x72, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x12, 0x2e, 0x0a, 0x13, 0x70, 0x61, 0x72, 0x74, 0x6e, 0x65, 0x72, 0x5f, 0x61, 0x63, 0x63, 0x6f,
	0x75, 0x6e, 0x74, 0x5f, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x11, 0x70,
	0x61, 0x72, 0x74, 0x6e, 0x65, 0x72, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x4b, 0x65, 0x79,
	0x22, 0xca, 0x02, 0x0a, 0x19, 0x47, 0x65, 0x74, 0x50, 0x61, 0x72, 0x74, 0x6e, 0x65, 0x72, 0x41,
	0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x53,
	0x0a, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x3b,
	0x2e, 0x70, 0x61, 0x72, 0x74, 0x6e, 0x65, 0x72, 0x2e, 0x70, 0x61, 0x72, 0x74, 0x6e, 0x65, 0x72,
	0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x2e, 0x76, 0x31, 0x2e, 0x47, 0x65, 0x74, 0x50, 0x61,
	0x72, 0x74, 0x6e, 0x65, 0x72, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x2e, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x52, 0x06, 0x73, 0x74, 0x61,
	0x74, 0x75, 0x73, 0x12, 0x26, 0x0a, 0x05, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x10, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x76, 0x31, 0x2e, 0x45,
	0x72, 0x72, 0x6f, 0x72, 0x52, 0x05, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x12, 0x52, 0x0a, 0x0f, 0x70,
	0x61, 0x72, 0x74, 0x6e, 0x65, 0x72, 0x5f, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x03,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x29, 0x2e, 0x70, 0x61, 0x72, 0x74, 0x6e, 0x65, 0x72, 0x2e, 0x70,
	0x61, 0x72, 0x74, 0x6e, 0x65, 0x72, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x2e, 0x76, 0x31,
	0x2e, 0x50, 0x61, 0x72, 0x74, 0x6e, 0x65, 0x72, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x52,
	0x0e, 0x70, 0x61, 0x72, 0x74, 0x6e, 0x65, 0x72, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x22,
	0x5c, 0x0a, 0x06, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x16, 0x0a, 0x12, 0x53, 0x54, 0x41,
	0x54, 0x55, 0x53, 0x5f, 0x55, 0x4e, 0x53, 0x50, 0x45, 0x43, 0x49, 0x46, 0x49, 0x45, 0x44, 0x10,
	0x00, 0x12, 0x12, 0x0a, 0x0e, 0x53, 0x54, 0x41, 0x54, 0x55, 0x53, 0x5f, 0x53, 0x55, 0x43, 0x43,
	0x45, 0x53, 0x53, 0x10, 0x01, 0x12, 0x10, 0x0a, 0x0c, 0x53, 0x54, 0x41, 0x54, 0x55, 0x53, 0x5f,
	0x45, 0x52, 0x52, 0x4f, 0x52, 0x10, 0x02, 0x12, 0x14, 0x0a, 0x10, 0x53, 0x54, 0x41, 0x54, 0x55,
	0x53, 0x5f, 0x4e, 0x4f, 0x54, 0x5f, 0x46, 0x4f, 0x55, 0x4e, 0x44, 0x10, 0x03, 0x22, 0xaa, 0x01,
	0x0a, 0x1b, 0x53, 0x65, 0x61, 0x72, 0x63, 0x68, 0x50, 0x61, 0x72, 0x74, 0x6e, 0x65, 0x72, 0x41,
	0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x4b, 0x0a,
	0x0c, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x0e, 0x32, 0x28, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x65, 0x6e, 0x75,
	0x6d, 0x73, 0x2e, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x74, 0x79, 0x70, 0x65, 0x2e, 0x76,
	0x31, 0x2e, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x54, 0x79, 0x70, 0x65, 0x52, 0x0b, 0x61,
	0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x54, 0x79, 0x70, 0x65, 0x12, 0x3e, 0x0a, 0x08, 0x63, 0x75,
	0x72, 0x72, 0x65, 0x6e, 0x63, 0x79, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x22, 0x2e, 0x63,
	0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x65, 0x6e, 0x75, 0x6d, 0x73, 0x2e, 0x63, 0x75, 0x72, 0x72,
	0x65, 0x6e, 0x63, 0x79, 0x2e, 0x76, 0x31, 0x2e, 0x43, 0x75, 0x72, 0x72, 0x65, 0x6e, 0x63, 0x79,
	0x52, 0x08, 0x63, 0x75, 0x72, 0x72, 0x65, 0x6e, 0x63, 0x79, 0x22, 0xbc, 0x02, 0x0a, 0x1c, 0x53,
	0x65, 0x61, 0x72, 0x63, 0x68, 0x50, 0x61, 0x72, 0x74, 0x6e, 0x65, 0x72, 0x41, 0x63, 0x63, 0x6f,
	0x75, 0x6e, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x56, 0x0a, 0x06, 0x73,
	0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x3e, 0x2e, 0x70, 0x61,
	0x72, 0x74, 0x6e, 0x65, 0x72, 0x2e, 0x70, 0x61, 0x72, 0x74, 0x6e, 0x65, 0x72, 0x61, 0x63, 0x63,
	0x6f, 0x75, 0x6e, 0x74, 0x2e, 0x76, 0x31, 0x2e, 0x53, 0x65, 0x61, 0x72, 0x63, 0x68, 0x50, 0x61,
	0x72, 0x74, 0x6e, 0x65, 0x72, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x2e, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x52, 0x06, 0x73, 0x74, 0x61,
	0x74, 0x75, 0x73, 0x12, 0x26, 0x0a, 0x05, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x10, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x76, 0x31, 0x2e, 0x45,
	0x72, 0x72, 0x6f, 0x72, 0x52, 0x05, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x12, 0x54, 0x0a, 0x10, 0x70,
	0x61, 0x72, 0x74, 0x6e, 0x65, 0x72, 0x5f, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x73, 0x18,
	0x03, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x29, 0x2e, 0x70, 0x61, 0x72, 0x74, 0x6e, 0x65, 0x72, 0x2e,
	0x70, 0x61, 0x72, 0x74, 0x6e, 0x65, 0x72, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x2e, 0x76,
	0x31, 0x2e, 0x50, 0x61, 0x72, 0x74, 0x6e, 0x65, 0x72, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74,
	0x52, 0x0f, 0x70, 0x61, 0x72, 0x74, 0x6e, 0x65, 0x72, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74,
	0x73, 0x22, 0x46, 0x0a, 0x06, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x16, 0x0a, 0x12, 0x53,
	0x54, 0x41, 0x54, 0x55, 0x53, 0x5f, 0x55, 0x4e, 0x53, 0x50, 0x45, 0x43, 0x49, 0x46, 0x49, 0x45,
	0x44, 0x10, 0x00, 0x12, 0x12, 0x0a, 0x0e, 0x53, 0x54, 0x41, 0x54, 0x55, 0x53, 0x5f, 0x53, 0x55,
	0x43, 0x43, 0x45, 0x53, 0x53, 0x10, 0x01, 0x12, 0x10, 0x0a, 0x0c, 0x53, 0x54, 0x41, 0x54, 0x55,
	0x53, 0x5f, 0x45, 0x52, 0x52, 0x4f, 0x52, 0x10, 0x02, 0x42, 0x69, 0x0a, 0x26, 0x6e, 0x65, 0x74,
	0x77, 0x6f, 0x72, 0x6b, 0x2e, 0x6d, 0x69, 0x63, 0x61, 0x2e, 0x70, 0x61, 0x72, 0x74, 0x6e, 0x65,
	0x72, 0x2e, 0x70, 0x61, 0x72, 0x74, 0x6e, 0x65, 0x72, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74,
	0x2e, 0x76, 0x31, 0x42, 0x13, 0x50, 0x61, 0x72, 0x74, 0x6e, 0x65, 0x72, 0x41, 0x63, 0x63, 0x6f,
	0x75, 0x6e, 0x74, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x5a, 0x23, 0x6d, 0x69, 0x63, 0x61, 0x2f, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x70, 0x61, 0x72, 0x74, 0x6e, 0x65, 0x72, 0x2f, 0x70, 0x61, 0x72,
	0x74, 0x6e, 0x65, 0x72, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x76, 0x31, 0xa2, 0x02, 0x04,
	0x4d, 0x49, 0x43, 0x41, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_partner_partneraccount_v1_partner_account_proto_rawDescOnce sync.Once
	file_partner_partneraccount_v1_partner_account_proto_rawDescData = file_partner_partneraccount_v1_partner_account_proto_rawDesc
)

func file_partner_partneraccount_v1_partner_account_proto_rawDescGZIP() []byte {
	file_partner_partneraccount_v1_partner_account_proto_rawDescOnce.Do(func() {
		file_partner_partneraccount_v1_partner_account_proto_rawDescData = protoimpl.X.CompressGZIP(file_partner_partneraccount_v1_partner_account_proto_rawDescData)
	})
	return file_partner_partneraccount_v1_partner_account_proto_rawDescData
}

var file_partner_partneraccount_v1_partner_account_proto_enumTypes = make([]protoimpl.EnumInfo, 2)
var file_partner_partneraccount_v1_partner_account_proto_msgTypes = make([]protoimpl.MessageInfo, 5)
var file_partner_partneraccount_v1_partner_account_proto_goTypes = []interface{}{
	(GetPartnerAccountResponse_Status)(0),    // 0: partner.partneraccount.v1.GetPartnerAccountResponse.Status
	(SearchPartnerAccountResponse_Status)(0), // 1: partner.partneraccount.v1.SearchPartnerAccountResponse.Status
	(*PartnerAccount)(nil),                   // 2: partner.partneraccount.v1.PartnerAccount
	(*GetPartnerAccountRequest)(nil),         // 3: partner.partneraccount.v1.GetPartnerAccountRequest
	(*GetPartnerAccountResponse)(nil),        // 4: partner.partneraccount.v1.GetPartnerAccountResponse
	(*SearchPartnerAccountRequest)(nil),      // 5: partner.partneraccount.v1.SearchPartnerAccountRequest
	(*SearchPartnerAccountResponse)(nil),     // 6: partner.partneraccount.v1.SearchPartnerAccountResponse
	(*timestamppb.Timestamp)(nil),            // 7: google.protobuf.Timestamp
	(accounttypev1.AccountType)(0),           // 8: common.enums.accounttype.v1.AccountType
	(currencyv1.Currency)(0),                 // 9: common.enums.currency.v1.Currency
	(*v1.Error)(nil),                         // 10: common.v1.Error
}
var file_partner_partneraccount_v1_partner_account_proto_depIdxs = []int32{
	7,  // 0: partner.partneraccount.v1.PartnerAccount.created:type_name -> google.protobuf.Timestamp
	7,  // 1: partner.partneraccount.v1.PartnerAccount.updated:type_name -> google.protobuf.Timestamp
	8,  // 2: partner.partneraccount.v1.PartnerAccount.account_type:type_name -> common.enums.accounttype.v1.AccountType
	9,  // 3: partner.partneraccount.v1.PartnerAccount.currency:type_name -> common.enums.currency.v1.Currency
	0,  // 4: partner.partneraccount.v1.GetPartnerAccountResponse.status:type_name -> partner.partneraccount.v1.GetPartnerAccountResponse.Status
	10, // 5: partner.partneraccount.v1.GetPartnerAccountResponse.error:type_name -> common.v1.Error
	2,  // 6: partner.partneraccount.v1.GetPartnerAccountResponse.partner_account:type_name -> partner.partneraccount.v1.PartnerAccount
	8,  // 7: partner.partneraccount.v1.SearchPartnerAccountRequest.account_type:type_name -> common.enums.accounttype.v1.AccountType
	9,  // 8: partner.partneraccount.v1.SearchPartnerAccountRequest.currency:type_name -> common.enums.currency.v1.Currency
	1,  // 9: partner.partneraccount.v1.SearchPartnerAccountResponse.status:type_name -> partner.partneraccount.v1.SearchPartnerAccountResponse.Status
	10, // 10: partner.partneraccount.v1.SearchPartnerAccountResponse.error:type_name -> common.v1.Error
	2,  // 11: partner.partneraccount.v1.SearchPartnerAccountResponse.partner_accounts:type_name -> partner.partneraccount.v1.PartnerAccount
	12, // [12:12] is the sub-list for method output_type
	12, // [12:12] is the sub-list for method input_type
	12, // [12:12] is the sub-list for extension type_name
	12, // [12:12] is the sub-list for extension extendee
	0,  // [0:12] is the sub-list for field type_name
}

func init() { file_partner_partneraccount_v1_partner_account_proto_init() }
func file_partner_partneraccount_v1_partner_account_proto_init() {
	if File_partner_partneraccount_v1_partner_account_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_partner_partneraccount_v1_partner_account_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PartnerAccount); i {
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
		file_partner_partneraccount_v1_partner_account_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetPartnerAccountRequest); i {
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
		file_partner_partneraccount_v1_partner_account_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetPartnerAccountResponse); i {
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
		file_partner_partneraccount_v1_partner_account_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SearchPartnerAccountRequest); i {
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
		file_partner_partneraccount_v1_partner_account_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SearchPartnerAccountResponse); i {
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
			RawDescriptor: file_partner_partneraccount_v1_partner_account_proto_rawDesc,
			NumEnums:      2,
			NumMessages:   5,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_partner_partneraccount_v1_partner_account_proto_goTypes,
		DependencyIndexes: file_partner_partneraccount_v1_partner_account_proto_depIdxs,
		EnumInfos:         file_partner_partneraccount_v1_partner_account_proto_enumTypes,
		MessageInfos:      file_partner_partneraccount_v1_partner_account_proto_msgTypes,
	}.Build()
	File_partner_partneraccount_v1_partner_account_proto = out.File
	file_partner_partneraccount_v1_partner_account_proto_rawDesc = nil
	file_partner_partneraccount_v1_partner_account_proto_goTypes = nil
	file_partner_partneraccount_v1_partner_account_proto_depIdxs = nil
}
