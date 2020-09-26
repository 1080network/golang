// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.23.0
// 	protoc        v3.12.2
// source: partner/partner/v1/partner.proto

package partnerv1

import (
	proto "github.com/golang/protobuf/proto"
	timestamp "github.com/golang/protobuf/ptypes/timestamp"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
	accounttypev1 "ten80/proto/common/enums/accounttypev1"
	currencyv1 "ten80/proto/common/enums/currencyv1"
	v1 "ten80/proto/common/v1"
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

type GetPartnerResponse_Status int32

const (
	GetPartnerResponse_STATUS_UNSPECIFIED GetPartnerResponse_Status = 0
	GetPartnerResponse_STATUS_SUCCESS     GetPartnerResponse_Status = 1
	GetPartnerResponse_STATUS_ERROR       GetPartnerResponse_Status = 2
	GetPartnerResponse_STATUS_NOT_FOUND   GetPartnerResponse_Status = 3
)

// Enum value maps for GetPartnerResponse_Status.
var (
	GetPartnerResponse_Status_name = map[int32]string{
		0: "STATUS_UNSPECIFIED",
		1: "STATUS_SUCCESS",
		2: "STATUS_ERROR",
		3: "STATUS_NOT_FOUND",
	}
	GetPartnerResponse_Status_value = map[string]int32{
		"STATUS_UNSPECIFIED": 0,
		"STATUS_SUCCESS":     1,
		"STATUS_ERROR":       2,
		"STATUS_NOT_FOUND":   3,
	}
)

func (x GetPartnerResponse_Status) Enum() *GetPartnerResponse_Status {
	p := new(GetPartnerResponse_Status)
	*p = x
	return p
}

func (x GetPartnerResponse_Status) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (GetPartnerResponse_Status) Descriptor() protoreflect.EnumDescriptor {
	return file_partner_partner_v1_partner_proto_enumTypes[0].Descriptor()
}

func (GetPartnerResponse_Status) Type() protoreflect.EnumType {
	return &file_partner_partner_v1_partner_proto_enumTypes[0]
}

func (x GetPartnerResponse_Status) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use GetPartnerResponse_Status.Descriptor instead.
func (GetPartnerResponse_Status) EnumDescriptor() ([]byte, []int) {
	return file_partner_partner_v1_partner_proto_rawDescGZIP(), []int{3, 0}
}

type PartnerAccount struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	PartnerAccountKey string                    `protobuf:"bytes,1,opt,name=partner_account_key,json=partnerAccountKey,proto3" json:"partner_account_key,omitempty"`
	Created           *timestamp.Timestamp      `protobuf:"bytes,2,opt,name=created,proto3" json:"created,omitempty"`
	Updated           *timestamp.Timestamp      `protobuf:"bytes,3,opt,name=updated,proto3" json:"updated,omitempty"`
	Version           int64                     `protobuf:"varint,4,opt,name=version,proto3" json:"version,omitempty"`
	PartnerKey        string                    `protobuf:"bytes,5,opt,name=partner_key,json=partnerKey,proto3" json:"partner_key,omitempty"`
	AccountType       accounttypev1.AccountType `protobuf:"varint,6,opt,name=account_type,json=accountType,proto3,enum=common.enums.accounttype.v1.AccountType" json:"account_type,omitempty"`
	Currency          currencyv1.Currency       `protobuf:"varint,7,opt,name=currency,proto3,enum=common.enums.currency.v1.Currency" json:"currency,omitempty"`
}

func (x *PartnerAccount) Reset() {
	*x = PartnerAccount{}
	if protoimpl.UnsafeEnabled {
		mi := &file_partner_partner_v1_partner_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PartnerAccount) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PartnerAccount) ProtoMessage() {}

func (x *PartnerAccount) ProtoReflect() protoreflect.Message {
	mi := &file_partner_partner_v1_partner_proto_msgTypes[0]
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
	return file_partner_partner_v1_partner_proto_rawDescGZIP(), []int{0}
}

func (x *PartnerAccount) GetPartnerAccountKey() string {
	if x != nil {
		return x.PartnerAccountKey
	}
	return ""
}

func (x *PartnerAccount) GetCreated() *timestamp.Timestamp {
	if x != nil {
		return x.Created
	}
	return nil
}

func (x *PartnerAccount) GetUpdated() *timestamp.Timestamp {
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
	return accounttypev1.AccountType_ACCOUNT_TYPE_UNSPECIFIED
}

func (x *PartnerAccount) GetCurrency() currencyv1.Currency {
	if x != nil {
		return x.Currency
	}
	return currencyv1.Currency_CURRENCY_UNSPECIFIED
}

type Partner struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	PartnerKey      string               `protobuf:"bytes,1,opt,name=partner_key,json=partnerKey,proto3" json:"partner_key,omitempty"`
	Created         *timestamp.Timestamp `protobuf:"bytes,2,opt,name=created,proto3" json:"created,omitempty"`
	Updated         *timestamp.Timestamp `protobuf:"bytes,3,opt,name=updated,proto3" json:"updated,omitempty"`
	Version         int64                `protobuf:"varint,4,opt,name=version,proto3" json:"version,omitempty"`
	Name            string               `protobuf:"bytes,5,opt,name=name,proto3" json:"name,omitempty"`
	TaxIdentifier   string               `protobuf:"bytes,6,opt,name=tax_identifier,json=taxIdentifier,proto3" json:"tax_identifier,omitempty"`
	Address         *v1.Address          `protobuf:"bytes,7,opt,name=address,proto3" json:"address,omitempty"`
	PartnerUrl      string               `protobuf:"bytes,8,opt,name=partner_url,json=partnerUrl,proto3" json:"partner_url,omitempty"`
	PartnerAccounts []*PartnerAccount    `protobuf:"bytes,9,rep,name=partner_accounts,json=partnerAccounts,proto3" json:"partner_accounts,omitempty"`
}

func (x *Partner) Reset() {
	*x = Partner{}
	if protoimpl.UnsafeEnabled {
		mi := &file_partner_partner_v1_partner_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Partner) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Partner) ProtoMessage() {}

func (x *Partner) ProtoReflect() protoreflect.Message {
	mi := &file_partner_partner_v1_partner_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Partner.ProtoReflect.Descriptor instead.
func (*Partner) Descriptor() ([]byte, []int) {
	return file_partner_partner_v1_partner_proto_rawDescGZIP(), []int{1}
}

func (x *Partner) GetPartnerKey() string {
	if x != nil {
		return x.PartnerKey
	}
	return ""
}

func (x *Partner) GetCreated() *timestamp.Timestamp {
	if x != nil {
		return x.Created
	}
	return nil
}

func (x *Partner) GetUpdated() *timestamp.Timestamp {
	if x != nil {
		return x.Updated
	}
	return nil
}

func (x *Partner) GetVersion() int64 {
	if x != nil {
		return x.Version
	}
	return 0
}

func (x *Partner) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *Partner) GetTaxIdentifier() string {
	if x != nil {
		return x.TaxIdentifier
	}
	return ""
}

func (x *Partner) GetAddress() *v1.Address {
	if x != nil {
		return x.Address
	}
	return nil
}

func (x *Partner) GetPartnerUrl() string {
	if x != nil {
		return x.PartnerUrl
	}
	return ""
}

func (x *Partner) GetPartnerAccounts() []*PartnerAccount {
	if x != nil {
		return x.PartnerAccounts
	}
	return nil
}

type GetPartnerRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *GetPartnerRequest) Reset() {
	*x = GetPartnerRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_partner_partner_v1_partner_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetPartnerRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetPartnerRequest) ProtoMessage() {}

func (x *GetPartnerRequest) ProtoReflect() protoreflect.Message {
	mi := &file_partner_partner_v1_partner_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetPartnerRequest.ProtoReflect.Descriptor instead.
func (*GetPartnerRequest) Descriptor() ([]byte, []int) {
	return file_partner_partner_v1_partner_proto_rawDescGZIP(), []int{2}
}

type GetPartnerResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Status  GetPartnerResponse_Status `protobuf:"varint,1,opt,name=status,proto3,enum=partner.partner.v1.GetPartnerResponse_Status" json:"status,omitempty"`
	Error   *v1.Error                 `protobuf:"bytes,2,opt,name=error,proto3" json:"error,omitempty"`
	Partner *Partner                  `protobuf:"bytes,3,opt,name=partner,proto3" json:"partner,omitempty"`
}

func (x *GetPartnerResponse) Reset() {
	*x = GetPartnerResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_partner_partner_v1_partner_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetPartnerResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetPartnerResponse) ProtoMessage() {}

func (x *GetPartnerResponse) ProtoReflect() protoreflect.Message {
	mi := &file_partner_partner_v1_partner_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetPartnerResponse.ProtoReflect.Descriptor instead.
func (*GetPartnerResponse) Descriptor() ([]byte, []int) {
	return file_partner_partner_v1_partner_proto_rawDescGZIP(), []int{3}
}

func (x *GetPartnerResponse) GetStatus() GetPartnerResponse_Status {
	if x != nil {
		return x.Status
	}
	return GetPartnerResponse_STATUS_UNSPECIFIED
}

func (x *GetPartnerResponse) GetError() *v1.Error {
	if x != nil {
		return x.Error
	}
	return nil
}

func (x *GetPartnerResponse) GetPartner() *Partner {
	if x != nil {
		return x.Partner
	}
	return nil
}

var File_partner_partner_v1_partner_proto protoreflect.FileDescriptor

var file_partner_partner_v1_partner_proto_rawDesc = []byte{
	0x0a, 0x20, 0x70, 0x61, 0x72, 0x74, 0x6e, 0x65, 0x72, 0x2f, 0x70, 0x61, 0x72, 0x74, 0x6e, 0x65,
	0x72, 0x2f, 0x76, 0x31, 0x2f, 0x70, 0x61, 0x72, 0x74, 0x6e, 0x65, 0x72, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x12, 0x12, 0x70, 0x61, 0x72, 0x74, 0x6e, 0x65, 0x72, 0x2e, 0x70, 0x61, 0x72, 0x74,
	0x6e, 0x65, 0x72, 0x2e, 0x76, 0x31, 0x1a, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2f, 0x65,
	0x6e, 0x75, 0x6d, 0x73, 0x2f, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x74, 0x79, 0x70, 0x65,
	0x2f, 0x76, 0x31, 0x2f, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x5f, 0x74, 0x79, 0x70, 0x65,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x27, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2f, 0x65,
	0x6e, 0x75, 0x6d, 0x73, 0x2f, 0x63, 0x75, 0x72, 0x72, 0x65, 0x6e, 0x63, 0x79, 0x2f, 0x76, 0x31,
	0x2f, 0x63, 0x75, 0x72, 0x72, 0x65, 0x6e, 0x63, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a,
	0x17, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2f, 0x76, 0x31, 0x2f, 0x61, 0x64, 0x64, 0x72, 0x65,
	0x73, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x15, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e,
	0x2f, 0x76, 0x31, 0x2f, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a,
	0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66,
	0x2f, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x22, 0xf4, 0x02, 0x0a, 0x0e, 0x50, 0x61, 0x72, 0x74, 0x6e, 0x65, 0x72, 0x41, 0x63, 0x63, 0x6f,
	0x75, 0x6e, 0x74, 0x12, 0x2e, 0x0a, 0x13, 0x70, 0x61, 0x72, 0x74, 0x6e, 0x65, 0x72, 0x5f, 0x61,
	0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x5f, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x11, 0x70, 0x61, 0x72, 0x74, 0x6e, 0x65, 0x72, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74,
	0x4b, 0x65, 0x79, 0x12, 0x34, 0x0a, 0x07, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70,
	0x52, 0x07, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x12, 0x34, 0x0a, 0x07, 0x75, 0x70, 0x64,
	0x61, 0x74, 0x65, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d,
	0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x07, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x12,
	0x18, 0x0a, 0x07, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x18, 0x04, 0x20, 0x01, 0x28, 0x03,
	0x52, 0x07, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x12, 0x1f, 0x0a, 0x0b, 0x70, 0x61, 0x72,
	0x74, 0x6e, 0x65, 0x72, 0x5f, 0x6b, 0x65, 0x79, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a,
	0x70, 0x61, 0x72, 0x74, 0x6e, 0x65, 0x72, 0x4b, 0x65, 0x79, 0x12, 0x4b, 0x0a, 0x0c, 0x61, 0x63,
	0x63, 0x6f, 0x75, 0x6e, 0x74, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x18, 0x06, 0x20, 0x01, 0x28, 0x0e,
	0x32, 0x28, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x65, 0x6e, 0x75, 0x6d, 0x73, 0x2e,
	0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x74, 0x79, 0x70, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x41,
	0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x54, 0x79, 0x70, 0x65, 0x52, 0x0b, 0x61, 0x63, 0x63, 0x6f,
	0x75, 0x6e, 0x74, 0x54, 0x79, 0x70, 0x65, 0x12, 0x3e, 0x0a, 0x08, 0x63, 0x75, 0x72, 0x72, 0x65,
	0x6e, 0x63, 0x79, 0x18, 0x07, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x22, 0x2e, 0x63, 0x6f, 0x6d, 0x6d,
	0x6f, 0x6e, 0x2e, 0x65, 0x6e, 0x75, 0x6d, 0x73, 0x2e, 0x63, 0x75, 0x72, 0x72, 0x65, 0x6e, 0x63,
	0x79, 0x2e, 0x76, 0x31, 0x2e, 0x43, 0x75, 0x72, 0x72, 0x65, 0x6e, 0x63, 0x79, 0x52, 0x08, 0x63,
	0x75, 0x72, 0x72, 0x65, 0x6e, 0x63, 0x79, 0x22, 0x89, 0x03, 0x0a, 0x07, 0x50, 0x61, 0x72, 0x74,
	0x6e, 0x65, 0x72, 0x12, 0x1f, 0x0a, 0x0b, 0x70, 0x61, 0x72, 0x74, 0x6e, 0x65, 0x72, 0x5f, 0x6b,
	0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x70, 0x61, 0x72, 0x74, 0x6e, 0x65,
	0x72, 0x4b, 0x65, 0x79, 0x12, 0x34, 0x0a, 0x07, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d,
	0x70, 0x52, 0x07, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x12, 0x34, 0x0a, 0x07, 0x75, 0x70,
	0x64, 0x61, 0x74, 0x65, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69,
	0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x07, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64,
	0x12, 0x18, 0x0a, 0x07, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x18, 0x04, 0x20, 0x01, 0x28,
	0x03, 0x52, 0x07, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61,
	0x6d, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x25,
	0x0a, 0x0e, 0x74, 0x61, 0x78, 0x5f, 0x69, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x66, 0x69, 0x65, 0x72,
	0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0d, 0x74, 0x61, 0x78, 0x49, 0x64, 0x65, 0x6e, 0x74,
	0x69, 0x66, 0x69, 0x65, 0x72, 0x12, 0x2c, 0x0a, 0x07, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73,
	0x18, 0x07, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x12, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e,
	0x76, 0x31, 0x2e, 0x41, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x52, 0x07, 0x61, 0x64, 0x64, 0x72,
	0x65, 0x73, 0x73, 0x12, 0x1f, 0x0a, 0x0b, 0x70, 0x61, 0x72, 0x74, 0x6e, 0x65, 0x72, 0x5f, 0x75,
	0x72, 0x6c, 0x18, 0x08, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x70, 0x61, 0x72, 0x74, 0x6e, 0x65,
	0x72, 0x55, 0x72, 0x6c, 0x12, 0x4d, 0x0a, 0x10, 0x70, 0x61, 0x72, 0x74, 0x6e, 0x65, 0x72, 0x5f,
	0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x73, 0x18, 0x09, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x22,
	0x2e, 0x70, 0x61, 0x72, 0x74, 0x6e, 0x65, 0x72, 0x2e, 0x70, 0x61, 0x72, 0x74, 0x6e, 0x65, 0x72,
	0x2e, 0x76, 0x31, 0x2e, 0x50, 0x61, 0x72, 0x74, 0x6e, 0x65, 0x72, 0x41, 0x63, 0x63, 0x6f, 0x75,
	0x6e, 0x74, 0x52, 0x0f, 0x70, 0x61, 0x72, 0x74, 0x6e, 0x65, 0x72, 0x41, 0x63, 0x63, 0x6f, 0x75,
	0x6e, 0x74, 0x73, 0x22, 0x13, 0x0a, 0x11, 0x47, 0x65, 0x74, 0x50, 0x61, 0x72, 0x74, 0x6e, 0x65,
	0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x22, 0x98, 0x02, 0x0a, 0x12, 0x47, 0x65, 0x74,
	0x50, 0x61, 0x72, 0x74, 0x6e, 0x65, 0x72, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12,
	0x45, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0e, 0x32,
	0x2d, 0x2e, 0x70, 0x61, 0x72, 0x74, 0x6e, 0x65, 0x72, 0x2e, 0x70, 0x61, 0x72, 0x74, 0x6e, 0x65,
	0x72, 0x2e, 0x76, 0x31, 0x2e, 0x47, 0x65, 0x74, 0x50, 0x61, 0x72, 0x74, 0x6e, 0x65, 0x72, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x2e, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x52, 0x06,
	0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x26, 0x0a, 0x05, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x10, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x76,
	0x31, 0x2e, 0x45, 0x72, 0x72, 0x6f, 0x72, 0x52, 0x05, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x12, 0x35,
	0x0a, 0x07, 0x70, 0x61, 0x72, 0x74, 0x6e, 0x65, 0x72, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x1b, 0x2e, 0x70, 0x61, 0x72, 0x74, 0x6e, 0x65, 0x72, 0x2e, 0x70, 0x61, 0x72, 0x74, 0x6e, 0x65,
	0x72, 0x2e, 0x76, 0x31, 0x2e, 0x50, 0x61, 0x72, 0x74, 0x6e, 0x65, 0x72, 0x52, 0x07, 0x70, 0x61,
	0x72, 0x74, 0x6e, 0x65, 0x72, 0x22, 0x5c, 0x0a, 0x06, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12,
	0x16, 0x0a, 0x12, 0x53, 0x54, 0x41, 0x54, 0x55, 0x53, 0x5f, 0x55, 0x4e, 0x53, 0x50, 0x45, 0x43,
	0x49, 0x46, 0x49, 0x45, 0x44, 0x10, 0x00, 0x12, 0x12, 0x0a, 0x0e, 0x53, 0x54, 0x41, 0x54, 0x55,
	0x53, 0x5f, 0x53, 0x55, 0x43, 0x43, 0x45, 0x53, 0x53, 0x10, 0x01, 0x12, 0x10, 0x0a, 0x0c, 0x53,
	0x54, 0x41, 0x54, 0x55, 0x53, 0x5f, 0x45, 0x52, 0x52, 0x4f, 0x52, 0x10, 0x02, 0x12, 0x14, 0x0a,
	0x10, 0x53, 0x54, 0x41, 0x54, 0x55, 0x53, 0x5f, 0x4e, 0x4f, 0x54, 0x5f, 0x46, 0x4f, 0x55, 0x4e,
	0x44, 0x10, 0x03, 0x42, 0x57, 0x0a, 0x20, 0x6e, 0x65, 0x74, 0x77, 0x6f, 0x72, 0x6b, 0x2e, 0x74,
	0x65, 0x6e, 0x38, 0x30, 0x2e, 0x70, 0x61, 0x72, 0x74, 0x6e, 0x65, 0x72, 0x2e, 0x70, 0x61, 0x72,
	0x74, 0x6e, 0x65, 0x72, 0x2e, 0x76, 0x31, 0x42, 0x0c, 0x50, 0x61, 0x72, 0x74, 0x6e, 0x65, 0x72,
	0x50, 0x72, 0x6f, 0x74, 0x6f, 0x5a, 0x1d, 0x74, 0x65, 0x6e, 0x38, 0x30, 0x2f, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x2f, 0x70, 0x61, 0x72, 0x74, 0x6e, 0x65, 0x72, 0x2f, 0x70, 0x61, 0x72, 0x74, 0x6e,
	0x65, 0x72, 0x76, 0x31, 0xa2, 0x02, 0x05, 0x54, 0x45, 0x4e, 0x38, 0x30, 0x62, 0x06, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_partner_partner_v1_partner_proto_rawDescOnce sync.Once
	file_partner_partner_v1_partner_proto_rawDescData = file_partner_partner_v1_partner_proto_rawDesc
)

func file_partner_partner_v1_partner_proto_rawDescGZIP() []byte {
	file_partner_partner_v1_partner_proto_rawDescOnce.Do(func() {
		file_partner_partner_v1_partner_proto_rawDescData = protoimpl.X.CompressGZIP(file_partner_partner_v1_partner_proto_rawDescData)
	})
	return file_partner_partner_v1_partner_proto_rawDescData
}

var file_partner_partner_v1_partner_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_partner_partner_v1_partner_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_partner_partner_v1_partner_proto_goTypes = []interface{}{
	(GetPartnerResponse_Status)(0), // 0: partner.partner.v1.GetPartnerResponse.Status
	(*PartnerAccount)(nil),         // 1: partner.partner.v1.PartnerAccount
	(*Partner)(nil),                // 2: partner.partner.v1.Partner
	(*GetPartnerRequest)(nil),      // 3: partner.partner.v1.GetPartnerRequest
	(*GetPartnerResponse)(nil),     // 4: partner.partner.v1.GetPartnerResponse
	(*timestamp.Timestamp)(nil),    // 5: google.protobuf.Timestamp
	(accounttypev1.AccountType)(0), // 6: common.enums.accounttype.v1.AccountType
	(currencyv1.Currency)(0),       // 7: common.enums.currency.v1.Currency
	(*v1.Address)(nil),             // 8: common.v1.Address
	(*v1.Error)(nil),               // 9: common.v1.Error
}
var file_partner_partner_v1_partner_proto_depIdxs = []int32{
	5,  // 0: partner.partner.v1.PartnerAccount.created:type_name -> google.protobuf.Timestamp
	5,  // 1: partner.partner.v1.PartnerAccount.updated:type_name -> google.protobuf.Timestamp
	6,  // 2: partner.partner.v1.PartnerAccount.account_type:type_name -> common.enums.accounttype.v1.AccountType
	7,  // 3: partner.partner.v1.PartnerAccount.currency:type_name -> common.enums.currency.v1.Currency
	5,  // 4: partner.partner.v1.Partner.created:type_name -> google.protobuf.Timestamp
	5,  // 5: partner.partner.v1.Partner.updated:type_name -> google.protobuf.Timestamp
	8,  // 6: partner.partner.v1.Partner.address:type_name -> common.v1.Address
	1,  // 7: partner.partner.v1.Partner.partner_accounts:type_name -> partner.partner.v1.PartnerAccount
	0,  // 8: partner.partner.v1.GetPartnerResponse.status:type_name -> partner.partner.v1.GetPartnerResponse.Status
	9,  // 9: partner.partner.v1.GetPartnerResponse.error:type_name -> common.v1.Error
	2,  // 10: partner.partner.v1.GetPartnerResponse.partner:type_name -> partner.partner.v1.Partner
	11, // [11:11] is the sub-list for method output_type
	11, // [11:11] is the sub-list for method input_type
	11, // [11:11] is the sub-list for extension type_name
	11, // [11:11] is the sub-list for extension extendee
	0,  // [0:11] is the sub-list for field type_name
}

func init() { file_partner_partner_v1_partner_proto_init() }
func file_partner_partner_v1_partner_proto_init() {
	if File_partner_partner_v1_partner_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_partner_partner_v1_partner_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
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
		file_partner_partner_v1_partner_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Partner); i {
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
		file_partner_partner_v1_partner_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetPartnerRequest); i {
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
		file_partner_partner_v1_partner_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetPartnerResponse); i {
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
			RawDescriptor: file_partner_partner_v1_partner_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_partner_partner_v1_partner_proto_goTypes,
		DependencyIndexes: file_partner_partner_v1_partner_proto_depIdxs,
		EnumInfos:         file_partner_partner_v1_partner_proto_enumTypes,
		MessageInfos:      file_partner_partner_v1_partner_proto_msgTypes,
	}.Build()
	File_partner_partner_v1_partner_proto = out.File
	file_partner_partner_v1_partner_proto_rawDesc = nil
	file_partner_partner_v1_partner_proto_goTypes = nil
	file_partner_partner_v1_partner_proto_depIdxs = nil
}
