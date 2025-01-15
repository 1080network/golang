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
// 	protoc        v5.28.3
// source: mica/partner/account/v1/account.proto

package accountv1

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
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

type SearchABAAccountResponse_Status int32

const (
	SearchABAAccountResponse_STATUS_UNSPECIFIED SearchABAAccountResponse_Status = 0
	SearchABAAccountResponse_STATUS_SUCCESS     SearchABAAccountResponse_Status = 1
	SearchABAAccountResponse_STATUS_ERROR       SearchABAAccountResponse_Status = 2
	SearchABAAccountResponse_STATUS_NOT_FOUND   SearchABAAccountResponse_Status = 3
)

// Enum value maps for SearchABAAccountResponse_Status.
var (
	SearchABAAccountResponse_Status_name = map[int32]string{
		0: "STATUS_UNSPECIFIED",
		1: "STATUS_SUCCESS",
		2: "STATUS_ERROR",
		3: "STATUS_NOT_FOUND",
	}
	SearchABAAccountResponse_Status_value = map[string]int32{
		"STATUS_UNSPECIFIED": 0,
		"STATUS_SUCCESS":     1,
		"STATUS_ERROR":       2,
		"STATUS_NOT_FOUND":   3,
	}
)

func (x SearchABAAccountResponse_Status) Enum() *SearchABAAccountResponse_Status {
	p := new(SearchABAAccountResponse_Status)
	*p = x
	return p
}

func (x SearchABAAccountResponse_Status) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (SearchABAAccountResponse_Status) Descriptor() protoreflect.EnumDescriptor {
	return file_mica_partner_account_v1_account_proto_enumTypes[0].Descriptor()
}

func (SearchABAAccountResponse_Status) Type() protoreflect.EnumType {
	return &file_mica_partner_account_v1_account_proto_enumTypes[0]
}

func (x SearchABAAccountResponse_Status) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use SearchABAAccountResponse_Status.Descriptor instead.
func (SearchABAAccountResponse_Status) EnumDescriptor() ([]byte, []int) {
	return file_mica_partner_account_v1_account_proto_rawDescGZIP(), []int{1, 0}
}

type CreateABAAccountResponse_Status int32

const (
	CreateABAAccountResponse_STATUS_UNSPECIFIED                CreateABAAccountResponse_Status = 0
	CreateABAAccountResponse_STATUS_SUCCESS                    CreateABAAccountResponse_Status = 1
	CreateABAAccountResponse_STATUS_DUPLICATE_ACCOUNT          CreateABAAccountResponse_Status = 2
	CreateABAAccountResponse_STATUS_ROUTING_NUMBER_NOT_ALLOWED CreateABAAccountResponse_Status = 3
)

// Enum value maps for CreateABAAccountResponse_Status.
var (
	CreateABAAccountResponse_Status_name = map[int32]string{
		0: "STATUS_UNSPECIFIED",
		1: "STATUS_SUCCESS",
		2: "STATUS_DUPLICATE_ACCOUNT",
		3: "STATUS_ROUTING_NUMBER_NOT_ALLOWED",
	}
	CreateABAAccountResponse_Status_value = map[string]int32{
		"STATUS_UNSPECIFIED":                0,
		"STATUS_SUCCESS":                    1,
		"STATUS_DUPLICATE_ACCOUNT":          2,
		"STATUS_ROUTING_NUMBER_NOT_ALLOWED": 3,
	}
)

func (x CreateABAAccountResponse_Status) Enum() *CreateABAAccountResponse_Status {
	p := new(CreateABAAccountResponse_Status)
	*p = x
	return p
}

func (x CreateABAAccountResponse_Status) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (CreateABAAccountResponse_Status) Descriptor() protoreflect.EnumDescriptor {
	return file_mica_partner_account_v1_account_proto_enumTypes[1].Descriptor()
}

func (CreateABAAccountResponse_Status) Type() protoreflect.EnumType {
	return &file_mica_partner_account_v1_account_proto_enumTypes[1]
}

func (x CreateABAAccountResponse_Status) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use CreateABAAccountResponse_Status.Descriptor instead.
func (CreateABAAccountResponse_Status) EnumDescriptor() ([]byte, []int) {
	return file_mica_partner_account_v1_account_proto_rawDescGZIP(), []int{3, 0}
}

type GetABAAccountResponse_Status int32

const (
	GetABAAccountResponse_STATUS_UNSPECIFIED GetABAAccountResponse_Status = 0
	GetABAAccountResponse_STATUS_SUCCESS     GetABAAccountResponse_Status = 1
	GetABAAccountResponse_STATUS_NOT_FOUND   GetABAAccountResponse_Status = 3
)

// Enum value maps for GetABAAccountResponse_Status.
var (
	GetABAAccountResponse_Status_name = map[int32]string{
		0: "STATUS_UNSPECIFIED",
		1: "STATUS_SUCCESS",
		3: "STATUS_NOT_FOUND",
	}
	GetABAAccountResponse_Status_value = map[string]int32{
		"STATUS_UNSPECIFIED": 0,
		"STATUS_SUCCESS":     1,
		"STATUS_NOT_FOUND":   3,
	}
)

func (x GetABAAccountResponse_Status) Enum() *GetABAAccountResponse_Status {
	p := new(GetABAAccountResponse_Status)
	*p = x
	return p
}

func (x GetABAAccountResponse_Status) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (GetABAAccountResponse_Status) Descriptor() protoreflect.EnumDescriptor {
	return file_mica_partner_account_v1_account_proto_enumTypes[2].Descriptor()
}

func (GetABAAccountResponse_Status) Type() protoreflect.EnumType {
	return &file_mica_partner_account_v1_account_proto_enumTypes[2]
}

func (x GetABAAccountResponse_Status) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use GetABAAccountResponse_Status.Descriptor instead.
func (GetABAAccountResponse_Status) EnumDescriptor() ([]byte, []int) {
	return file_mica_partner_account_v1_account_proto_rawDescGZIP(), []int{5, 0}
}

// request to search ABA accounts available in the partner instance
type SearchABAAccountRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// if set the search will filter by bank name
	BankName string `protobuf:"bytes,1,opt,name=bank_name,json=bankName,proto3" json:"bank_name,omitempty"`
}

func (x *SearchABAAccountRequest) Reset() {
	*x = SearchABAAccountRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_mica_partner_account_v1_account_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SearchABAAccountRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SearchABAAccountRequest) ProtoMessage() {}

func (x *SearchABAAccountRequest) ProtoReflect() protoreflect.Message {
	mi := &file_mica_partner_account_v1_account_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SearchABAAccountRequest.ProtoReflect.Descriptor instead.
func (*SearchABAAccountRequest) Descriptor() ([]byte, []int) {
	return file_mica_partner_account_v1_account_proto_rawDescGZIP(), []int{0}
}

func (x *SearchABAAccountRequest) GetBankName() string {
	if x != nil {
		return x.BankName
	}
	return ""
}

type SearchABAAccountResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Status      SearchABAAccountResponse_Status `protobuf:"varint,1,opt,name=status,proto3,enum=mica.partner.account.v1.SearchABAAccountResponse_Status" json:"status,omitempty"`
	Error       *v1.Error                       `protobuf:"bytes,2,opt,name=error,proto3" json:"error,omitempty"`
	AbaAccounts []*v1.ABAAccountEntry           `protobuf:"bytes,3,rep,name=aba_accounts,json=abaAccounts,proto3" json:"aba_accounts,omitempty"`
}

func (x *SearchABAAccountResponse) Reset() {
	*x = SearchABAAccountResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_mica_partner_account_v1_account_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SearchABAAccountResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SearchABAAccountResponse) ProtoMessage() {}

func (x *SearchABAAccountResponse) ProtoReflect() protoreflect.Message {
	mi := &file_mica_partner_account_v1_account_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SearchABAAccountResponse.ProtoReflect.Descriptor instead.
func (*SearchABAAccountResponse) Descriptor() ([]byte, []int) {
	return file_mica_partner_account_v1_account_proto_rawDescGZIP(), []int{1}
}

func (x *SearchABAAccountResponse) GetStatus() SearchABAAccountResponse_Status {
	if x != nil {
		return x.Status
	}
	return SearchABAAccountResponse_STATUS_UNSPECIFIED
}

func (x *SearchABAAccountResponse) GetError() *v1.Error {
	if x != nil {
		return x.Error
	}
	return nil
}

func (x *SearchABAAccountResponse) GetAbaAccounts() []*v1.ABAAccountEntry {
	if x != nil {
		return x.AbaAccounts
	}
	return nil
}

type CreateABAAccountRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Account *v1.ABAAccountNumber `protobuf:"bytes,1,opt,name=account,proto3" json:"account,omitempty"`
}

func (x *CreateABAAccountRequest) Reset() {
	*x = CreateABAAccountRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_mica_partner_account_v1_account_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateABAAccountRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateABAAccountRequest) ProtoMessage() {}

func (x *CreateABAAccountRequest) ProtoReflect() protoreflect.Message {
	mi := &file_mica_partner_account_v1_account_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateABAAccountRequest.ProtoReflect.Descriptor instead.
func (*CreateABAAccountRequest) Descriptor() ([]byte, []int) {
	return file_mica_partner_account_v1_account_proto_rawDescGZIP(), []int{2}
}

func (x *CreateABAAccountRequest) GetAccount() *v1.ABAAccountNumber {
	if x != nil {
		return x.Account
	}
	return nil
}

type CreateABAAccountResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Status     CreateABAAccountResponse_Status `protobuf:"varint,1,opt,name=status,proto3,enum=mica.partner.account.v1.CreateABAAccountResponse_Status" json:"status,omitempty"`
	Error      *v1.Error                       `protobuf:"bytes,2,opt,name=error,proto3" json:"error,omitempty"`
	AccountKey string                          `protobuf:"bytes,3,opt,name=account_key,json=accountKey,proto3" json:"account_key,omitempty"`
}

func (x *CreateABAAccountResponse) Reset() {
	*x = CreateABAAccountResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_mica_partner_account_v1_account_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateABAAccountResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateABAAccountResponse) ProtoMessage() {}

func (x *CreateABAAccountResponse) ProtoReflect() protoreflect.Message {
	mi := &file_mica_partner_account_v1_account_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateABAAccountResponse.ProtoReflect.Descriptor instead.
func (*CreateABAAccountResponse) Descriptor() ([]byte, []int) {
	return file_mica_partner_account_v1_account_proto_rawDescGZIP(), []int{3}
}

func (x *CreateABAAccountResponse) GetStatus() CreateABAAccountResponse_Status {
	if x != nil {
		return x.Status
	}
	return CreateABAAccountResponse_STATUS_UNSPECIFIED
}

func (x *CreateABAAccountResponse) GetError() *v1.Error {
	if x != nil {
		return x.Error
	}
	return nil
}

func (x *CreateABAAccountResponse) GetAccountKey() string {
	if x != nil {
		return x.AccountKey
	}
	return ""
}

type GetABAAccountRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	AccountKey string `protobuf:"bytes,1,opt,name=account_key,json=accountKey,proto3" json:"account_key,omitempty"`
}

func (x *GetABAAccountRequest) Reset() {
	*x = GetABAAccountRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_mica_partner_account_v1_account_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetABAAccountRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetABAAccountRequest) ProtoMessage() {}

func (x *GetABAAccountRequest) ProtoReflect() protoreflect.Message {
	mi := &file_mica_partner_account_v1_account_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetABAAccountRequest.ProtoReflect.Descriptor instead.
func (*GetABAAccountRequest) Descriptor() ([]byte, []int) {
	return file_mica_partner_account_v1_account_proto_rawDescGZIP(), []int{4}
}

func (x *GetABAAccountRequest) GetAccountKey() string {
	if x != nil {
		return x.AccountKey
	}
	return ""
}

type GetABAAccountResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Status  GetABAAccountResponse_Status `protobuf:"varint,1,opt,name=status,proto3,enum=mica.partner.account.v1.GetABAAccountResponse_Status" json:"status,omitempty"`
	Error   *v1.Error                    `protobuf:"bytes,2,opt,name=error,proto3" json:"error,omitempty"`
	Account *v1.ABAAccountNumber         `protobuf:"bytes,3,opt,name=account,proto3" json:"account,omitempty"`
}

func (x *GetABAAccountResponse) Reset() {
	*x = GetABAAccountResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_mica_partner_account_v1_account_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetABAAccountResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetABAAccountResponse) ProtoMessage() {}

func (x *GetABAAccountResponse) ProtoReflect() protoreflect.Message {
	mi := &file_mica_partner_account_v1_account_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetABAAccountResponse.ProtoReflect.Descriptor instead.
func (*GetABAAccountResponse) Descriptor() ([]byte, []int) {
	return file_mica_partner_account_v1_account_proto_rawDescGZIP(), []int{5}
}

func (x *GetABAAccountResponse) GetStatus() GetABAAccountResponse_Status {
	if x != nil {
		return x.Status
	}
	return GetABAAccountResponse_STATUS_UNSPECIFIED
}

func (x *GetABAAccountResponse) GetError() *v1.Error {
	if x != nil {
		return x.Error
	}
	return nil
}

func (x *GetABAAccountResponse) GetAccount() *v1.ABAAccountNumber {
	if x != nil {
		return x.Account
	}
	return nil
}

var File_mica_partner_account_v1_account_proto protoreflect.FileDescriptor

var file_mica_partner_account_v1_account_proto_rawDesc = []byte{
	0x0a, 0x25, 0x6d, 0x69, 0x63, 0x61, 0x2f, 0x70, 0x61, 0x72, 0x74, 0x6e, 0x65, 0x72, 0x2f, 0x61,
	0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x2f, 0x76, 0x31, 0x2f, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e,
	0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x17, 0x6d, 0x69, 0x63, 0x61, 0x2e, 0x70, 0x61,
	0x72, 0x74, 0x6e, 0x65, 0x72, 0x2e, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x2e, 0x76, 0x31,
	0x1a, 0x2e, 0x6d, 0x69, 0x63, 0x61, 0x73, 0x68, 0x61, 0x72, 0x65, 0x64, 0x2f, 0x63, 0x6f, 0x6d,
	0x6d, 0x6f, 0x6e, 0x2f, 0x76, 0x31, 0x2f, 0x62, 0x61, 0x6e, 0x6b, 0x5f, 0x61, 0x63, 0x63, 0x6f,
	0x75, 0x6e, 0x74, 0x5f, 0x64, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x1a, 0x20, 0x6d, 0x69, 0x63, 0x61, 0x73, 0x68, 0x61, 0x72, 0x65, 0x64, 0x2f, 0x63, 0x6f, 0x6d,
	0x6d, 0x6f, 0x6e, 0x2f, 0x76, 0x31, 0x2f, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x22, 0x36, 0x0a, 0x17, 0x53, 0x65, 0x61, 0x72, 0x63, 0x68, 0x41, 0x42, 0x41, 0x41,
	0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1b, 0x0a,
	0x09, 0x62, 0x61, 0x6e, 0x6b, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x08, 0x62, 0x61, 0x6e, 0x6b, 0x4e, 0x61, 0x6d, 0x65, 0x22, 0xc7, 0x02, 0x0a, 0x18, 0x53,
	0x65, 0x61, 0x72, 0x63, 0x68, 0x41, 0x42, 0x41, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x50, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75,
	0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x38, 0x2e, 0x6d, 0x69, 0x63, 0x61, 0x2e, 0x70,
	0x61, 0x72, 0x74, 0x6e, 0x65, 0x72, 0x2e, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x2e, 0x76,
	0x31, 0x2e, 0x53, 0x65, 0x61, 0x72, 0x63, 0x68, 0x41, 0x42, 0x41, 0x41, 0x63, 0x63, 0x6f, 0x75,
	0x6e, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x2e, 0x53, 0x74, 0x61, 0x74, 0x75,
	0x73, 0x52, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x31, 0x0a, 0x05, 0x65, 0x72, 0x72,
	0x6f, 0x72, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1b, 0x2e, 0x6d, 0x69, 0x63, 0x61, 0x73,
	0x68, 0x61, 0x72, 0x65, 0x64, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x76, 0x31, 0x2e,
	0x45, 0x72, 0x72, 0x6f, 0x72, 0x52, 0x05, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x12, 0x48, 0x0a, 0x0c,
	0x61, 0x62, 0x61, 0x5f, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x73, 0x18, 0x03, 0x20, 0x03,
	0x28, 0x0b, 0x32, 0x25, 0x2e, 0x6d, 0x69, 0x63, 0x61, 0x73, 0x68, 0x61, 0x72, 0x65, 0x64, 0x2e,
	0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x76, 0x31, 0x2e, 0x41, 0x42, 0x41, 0x41, 0x63, 0x63,
	0x6f, 0x75, 0x6e, 0x74, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x0b, 0x61, 0x62, 0x61, 0x41, 0x63,
	0x63, 0x6f, 0x75, 0x6e, 0x74, 0x73, 0x22, 0x5c, 0x0a, 0x06, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73,
	0x12, 0x16, 0x0a, 0x12, 0x53, 0x54, 0x41, 0x54, 0x55, 0x53, 0x5f, 0x55, 0x4e, 0x53, 0x50, 0x45,
	0x43, 0x49, 0x46, 0x49, 0x45, 0x44, 0x10, 0x00, 0x12, 0x12, 0x0a, 0x0e, 0x53, 0x54, 0x41, 0x54,
	0x55, 0x53, 0x5f, 0x53, 0x55, 0x43, 0x43, 0x45, 0x53, 0x53, 0x10, 0x01, 0x12, 0x10, 0x0a, 0x0c,
	0x53, 0x54, 0x41, 0x54, 0x55, 0x53, 0x5f, 0x45, 0x52, 0x52, 0x4f, 0x52, 0x10, 0x02, 0x12, 0x14,
	0x0a, 0x10, 0x53, 0x54, 0x41, 0x54, 0x55, 0x53, 0x5f, 0x4e, 0x4f, 0x54, 0x5f, 0x46, 0x4f, 0x55,
	0x4e, 0x44, 0x10, 0x03, 0x22, 0x5b, 0x0a, 0x17, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x41, 0x42,
	0x41, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12,
	0x40, 0x0a, 0x07, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x26, 0x2e, 0x6d, 0x69, 0x63, 0x61, 0x73, 0x68, 0x61, 0x72, 0x65, 0x64, 0x2e, 0x63, 0x6f,
	0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x76, 0x31, 0x2e, 0x41, 0x42, 0x41, 0x41, 0x63, 0x63, 0x6f, 0x75,
	0x6e, 0x74, 0x4e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x52, 0x07, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e,
	0x74, 0x22, 0xbb, 0x02, 0x0a, 0x18, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x41, 0x42, 0x41, 0x41,
	0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x50,
	0x0a, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x38,
	0x2e, 0x6d, 0x69, 0x63, 0x61, 0x2e, 0x70, 0x61, 0x72, 0x74, 0x6e, 0x65, 0x72, 0x2e, 0x61, 0x63,
	0x63, 0x6f, 0x75, 0x6e, 0x74, 0x2e, 0x76, 0x31, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x41,
	0x42, 0x41, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x2e, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x52, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73,
	0x12, 0x31, 0x0a, 0x05, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x1b, 0x2e, 0x6d, 0x69, 0x63, 0x61, 0x73, 0x68, 0x61, 0x72, 0x65, 0x64, 0x2e, 0x63, 0x6f, 0x6d,
	0x6d, 0x6f, 0x6e, 0x2e, 0x76, 0x31, 0x2e, 0x45, 0x72, 0x72, 0x6f, 0x72, 0x52, 0x05, 0x65, 0x72,
	0x72, 0x6f, 0x72, 0x12, 0x1f, 0x0a, 0x0b, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x5f, 0x6b,
	0x65, 0x79, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e,
	0x74, 0x4b, 0x65, 0x79, 0x22, 0x79, 0x0a, 0x06, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x16,
	0x0a, 0x12, 0x53, 0x54, 0x41, 0x54, 0x55, 0x53, 0x5f, 0x55, 0x4e, 0x53, 0x50, 0x45, 0x43, 0x49,
	0x46, 0x49, 0x45, 0x44, 0x10, 0x00, 0x12, 0x12, 0x0a, 0x0e, 0x53, 0x54, 0x41, 0x54, 0x55, 0x53,
	0x5f, 0x53, 0x55, 0x43, 0x43, 0x45, 0x53, 0x53, 0x10, 0x01, 0x12, 0x1c, 0x0a, 0x18, 0x53, 0x54,
	0x41, 0x54, 0x55, 0x53, 0x5f, 0x44, 0x55, 0x50, 0x4c, 0x49, 0x43, 0x41, 0x54, 0x45, 0x5f, 0x41,
	0x43, 0x43, 0x4f, 0x55, 0x4e, 0x54, 0x10, 0x02, 0x12, 0x25, 0x0a, 0x21, 0x53, 0x54, 0x41, 0x54,
	0x55, 0x53, 0x5f, 0x52, 0x4f, 0x55, 0x54, 0x49, 0x4e, 0x47, 0x5f, 0x4e, 0x55, 0x4d, 0x42, 0x45,
	0x52, 0x5f, 0x4e, 0x4f, 0x54, 0x5f, 0x41, 0x4c, 0x4c, 0x4f, 0x57, 0x45, 0x44, 0x10, 0x03, 0x22,
	0x37, 0x0a, 0x14, 0x47, 0x65, 0x74, 0x41, 0x42, 0x41, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1f, 0x0a, 0x0b, 0x61, 0x63, 0x63, 0x6f, 0x75,
	0x6e, 0x74, 0x5f, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x61, 0x63,
	0x63, 0x6f, 0x75, 0x6e, 0x74, 0x4b, 0x65, 0x79, 0x22, 0xa7, 0x02, 0x0a, 0x15, 0x47, 0x65, 0x74,
	0x41, 0x42, 0x41, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x12, 0x4d, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x0e, 0x32, 0x35, 0x2e, 0x6d, 0x69, 0x63, 0x61, 0x2e, 0x70, 0x61, 0x72, 0x74, 0x6e, 0x65,
	0x72, 0x2e, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x2e, 0x76, 0x31, 0x2e, 0x47, 0x65, 0x74,
	0x41, 0x42, 0x41, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x2e, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x52, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75,
	0x73, 0x12, 0x31, 0x0a, 0x05, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x1b, 0x2e, 0x6d, 0x69, 0x63, 0x61, 0x73, 0x68, 0x61, 0x72, 0x65, 0x64, 0x2e, 0x63, 0x6f,
	0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x76, 0x31, 0x2e, 0x45, 0x72, 0x72, 0x6f, 0x72, 0x52, 0x05, 0x65,
	0x72, 0x72, 0x6f, 0x72, 0x12, 0x40, 0x0a, 0x07, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x18,
	0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x26, 0x2e, 0x6d, 0x69, 0x63, 0x61, 0x73, 0x68, 0x61, 0x72,
	0x65, 0x64, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x76, 0x31, 0x2e, 0x41, 0x42, 0x41,
	0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x4e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x52, 0x07, 0x61,
	0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x22, 0x4a, 0x0a, 0x06, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73,
	0x12, 0x16, 0x0a, 0x12, 0x53, 0x54, 0x41, 0x54, 0x55, 0x53, 0x5f, 0x55, 0x4e, 0x53, 0x50, 0x45,
	0x43, 0x49, 0x46, 0x49, 0x45, 0x44, 0x10, 0x00, 0x12, 0x12, 0x0a, 0x0e, 0x53, 0x54, 0x41, 0x54,
	0x55, 0x53, 0x5f, 0x53, 0x55, 0x43, 0x43, 0x45, 0x53, 0x53, 0x10, 0x01, 0x12, 0x14, 0x0a, 0x10,
	0x53, 0x54, 0x41, 0x54, 0x55, 0x53, 0x5f, 0x4e, 0x4f, 0x54, 0x5f, 0x46, 0x4f, 0x55, 0x4e, 0x44,
	0x10, 0x03, 0x42, 0x49, 0x0a, 0x1a, 0x69, 0x6f, 0x2e, 0x6d, 0x69, 0x63, 0x61, 0x2e, 0x70, 0x61,
	0x72, 0x74, 0x6e, 0x65, 0x72, 0x2e, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x2e, 0x76, 0x31,
	0x42, 0x0c, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x5a, 0x16,
	0x6d, 0x69, 0x63, 0x61, 0x2f, 0x70, 0x61, 0x72, 0x74, 0x6e, 0x65, 0x72, 0x2f, 0x61, 0x63, 0x63,
	0x6f, 0x75, 0x6e, 0x74, 0x76, 0x31, 0xa2, 0x02, 0x04, 0x4d, 0x49, 0x43, 0x41, 0x62, 0x06, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_mica_partner_account_v1_account_proto_rawDescOnce sync.Once
	file_mica_partner_account_v1_account_proto_rawDescData = file_mica_partner_account_v1_account_proto_rawDesc
)

func file_mica_partner_account_v1_account_proto_rawDescGZIP() []byte {
	file_mica_partner_account_v1_account_proto_rawDescOnce.Do(func() {
		file_mica_partner_account_v1_account_proto_rawDescData = protoimpl.X.CompressGZIP(file_mica_partner_account_v1_account_proto_rawDescData)
	})
	return file_mica_partner_account_v1_account_proto_rawDescData
}

var file_mica_partner_account_v1_account_proto_enumTypes = make([]protoimpl.EnumInfo, 3)
var file_mica_partner_account_v1_account_proto_msgTypes = make([]protoimpl.MessageInfo, 6)
var file_mica_partner_account_v1_account_proto_goTypes = []interface{}{
	(SearchABAAccountResponse_Status)(0), // 0: mica.partner.account.v1.SearchABAAccountResponse.Status
	(CreateABAAccountResponse_Status)(0), // 1: mica.partner.account.v1.CreateABAAccountResponse.Status
	(GetABAAccountResponse_Status)(0),    // 2: mica.partner.account.v1.GetABAAccountResponse.Status
	(*SearchABAAccountRequest)(nil),      // 3: mica.partner.account.v1.SearchABAAccountRequest
	(*SearchABAAccountResponse)(nil),     // 4: mica.partner.account.v1.SearchABAAccountResponse
	(*CreateABAAccountRequest)(nil),      // 5: mica.partner.account.v1.CreateABAAccountRequest
	(*CreateABAAccountResponse)(nil),     // 6: mica.partner.account.v1.CreateABAAccountResponse
	(*GetABAAccountRequest)(nil),         // 7: mica.partner.account.v1.GetABAAccountRequest
	(*GetABAAccountResponse)(nil),        // 8: mica.partner.account.v1.GetABAAccountResponse
	(*v1.Error)(nil),                     // 9: micashared.common.v1.Error
	(*v1.ABAAccountEntry)(nil),           // 10: micashared.common.v1.ABAAccountEntry
	(*v1.ABAAccountNumber)(nil),          // 11: micashared.common.v1.ABAAccountNumber
}
var file_mica_partner_account_v1_account_proto_depIdxs = []int32{
	0,  // 0: mica.partner.account.v1.SearchABAAccountResponse.status:type_name -> mica.partner.account.v1.SearchABAAccountResponse.Status
	9,  // 1: mica.partner.account.v1.SearchABAAccountResponse.error:type_name -> micashared.common.v1.Error
	10, // 2: mica.partner.account.v1.SearchABAAccountResponse.aba_accounts:type_name -> micashared.common.v1.ABAAccountEntry
	11, // 3: mica.partner.account.v1.CreateABAAccountRequest.account:type_name -> micashared.common.v1.ABAAccountNumber
	1,  // 4: mica.partner.account.v1.CreateABAAccountResponse.status:type_name -> mica.partner.account.v1.CreateABAAccountResponse.Status
	9,  // 5: mica.partner.account.v1.CreateABAAccountResponse.error:type_name -> micashared.common.v1.Error
	2,  // 6: mica.partner.account.v1.GetABAAccountResponse.status:type_name -> mica.partner.account.v1.GetABAAccountResponse.Status
	9,  // 7: mica.partner.account.v1.GetABAAccountResponse.error:type_name -> micashared.common.v1.Error
	11, // 8: mica.partner.account.v1.GetABAAccountResponse.account:type_name -> micashared.common.v1.ABAAccountNumber
	9,  // [9:9] is the sub-list for method output_type
	9,  // [9:9] is the sub-list for method input_type
	9,  // [9:9] is the sub-list for extension type_name
	9,  // [9:9] is the sub-list for extension extendee
	0,  // [0:9] is the sub-list for field type_name
}

func init() { file_mica_partner_account_v1_account_proto_init() }
func file_mica_partner_account_v1_account_proto_init() {
	if File_mica_partner_account_v1_account_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_mica_partner_account_v1_account_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SearchABAAccountRequest); i {
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
		file_mica_partner_account_v1_account_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SearchABAAccountResponse); i {
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
		file_mica_partner_account_v1_account_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateABAAccountRequest); i {
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
		file_mica_partner_account_v1_account_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateABAAccountResponse); i {
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
		file_mica_partner_account_v1_account_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetABAAccountRequest); i {
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
		file_mica_partner_account_v1_account_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetABAAccountResponse); i {
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
			RawDescriptor: file_mica_partner_account_v1_account_proto_rawDesc,
			NumEnums:      3,
			NumMessages:   6,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_mica_partner_account_v1_account_proto_goTypes,
		DependencyIndexes: file_mica_partner_account_v1_account_proto_depIdxs,
		EnumInfos:         file_mica_partner_account_v1_account_proto_enumTypes,
		MessageInfos:      file_mica_partner_account_v1_account_proto_msgTypes,
	}.Build()
	File_mica_partner_account_v1_account_proto = out.File
	file_mica_partner_account_v1_account_proto_rawDesc = nil
	file_mica_partner_account_v1_account_proto_goTypes = nil
	file_mica_partner_account_v1_account_proto_depIdxs = nil
}
