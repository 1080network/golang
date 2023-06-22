// Copyright (c) 2022 Mica. All rights reserved. All software, including, without limitation, all source
// code and object code, is the intellectual property of Mica, Inc. and is protected by copyright, trademark and
// other intellectual property laws (collective "Mica Software"). You may not use, copy, reproduce, download, store,
// post, broadcast, transmit, modify, sell or make available to the public content from the Mica Software without a
// valid license or the prior written approval of Mica, Inc. Mica, its logos, slogans, taglines,
// products, feature names, and other trademarks are trademarks of Mica, Inc. and may not be used without
// permission.

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.29.1
// 	protoc        v3.21.12
// source: micashared/common/v1/organization.proto

package commonv1

import (
	_ "github.com/envoyproxy/protoc-gen-validate/validate"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	timestamppb "google.golang.org/protobuf/types/known/timestamppb"
	wrapperspb "google.golang.org/protobuf/types/known/wrapperspb"
	currencyv1 "github.com/1080network/golang/partner/proto/micashared/common/enums/currencyv1"
	organizationcategoryv1 "github.com/1080network/golang/partner/proto/micashared/common/enums/organizationcategoryv1"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type Organization struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The primary key of the organization in Mica.
	OrganizationKey string `protobuf:"bytes,1,opt,name=organization_key,json=organizationKey,proto3" json:"organization_key,omitempty"`
	Version         int64  `protobuf:"varint,2,opt,name=version,proto3" json:"version,omitempty"`
	// Date that the organization was created at Mica.
	Created *timestamppb.Timestamp `protobuf:"bytes,3,opt,name=created,proto3" json:"created,omitempty"`
	// Date that the organization was last updated at Mica.
	Updated *timestamppb.Timestamp `protobuf:"bytes,4,opt,name=updated,proto3" json:"updated,omitempty"`
	// The partner who owns this organization.
	PartnerKey string `protobuf:"bytes,5,opt,name=partner_key,json=partnerKey,proto3" json:"partner_key,omitempty"`
	// The name of this organization.
	Name string `protobuf:"bytes,7,opt,name=name,proto3" json:"name,omitempty"`
	// The list of categories of this organization (used by Service Provider's to as part of their processing rules).
	Categories []organizationcategoryv1.OrganizationCategory `protobuf:"varint,6,rep,packed,name=categories,proto3,enum=micashared.common.enums.organizationcategory.v1.OrganizationCategory" json:"categories,omitempty"`
	// The street address of the organization.
	Address *Address `protobuf:"bytes,9,opt,name=address,proto3" json:"address,omitempty"`
	// The organizations domestic ACH routing number (9 characters)
	DomesticAchRoutingNumber string `protobuf:"bytes,10,opt,name=domestic_ach_routing_number,json=domesticAchRoutingNumber,proto3" json:"domestic_ach_routing_number,omitempty"`
	// The organizations international ACH routing number (9 characters)
	InternationalAchRoutingNumber string `protobuf:"bytes,11,opt,name=international_ach_routing_number,json=internationalAchRoutingNumber,proto3" json:"international_ach_routing_number,omitempty"`
	// The organizations wire routing number (9 characters)
	WireRoutingNumber string `protobuf:"bytes,12,opt,name=wire_routing_number,json=wireRoutingNumber,proto3" json:"wire_routing_number,omitempty"`
	// The organizations swift routing number (9 characters)
	SwiftRoutingNumber string `protobuf:"bytes,13,opt,name=swift_routing_number,json=swiftRoutingNumber,proto3" json:"swift_routing_number,omitempty"`
	// The organizations bank account number (14 digits)
	BankAccountNumber string `protobuf:"bytes,14,opt,name=bank_account_number,json=bankAccountNumber,proto3" json:"bank_account_number,omitempty"`
}

func (x *Organization) Reset() {
	*x = Organization{}
	if protoimpl.UnsafeEnabled {
		mi := &file_micashared_common_v1_organization_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Organization) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Organization) ProtoMessage() {}

func (x *Organization) ProtoReflect() protoreflect.Message {
	mi := &file_micashared_common_v1_organization_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Organization.ProtoReflect.Descriptor instead.
func (*Organization) Descriptor() ([]byte, []int) {
	return file_micashared_common_v1_organization_proto_rawDescGZIP(), []int{0}
}

func (x *Organization) GetOrganizationKey() string {
	if x != nil {
		return x.OrganizationKey
	}
	return ""
}

func (x *Organization) GetVersion() int64 {
	if x != nil {
		return x.Version
	}
	return 0
}

func (x *Organization) GetCreated() *timestamppb.Timestamp {
	if x != nil {
		return x.Created
	}
	return nil
}

func (x *Organization) GetUpdated() *timestamppb.Timestamp {
	if x != nil {
		return x.Updated
	}
	return nil
}

func (x *Organization) GetPartnerKey() string {
	if x != nil {
		return x.PartnerKey
	}
	return ""
}

func (x *Organization) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *Organization) GetCategories() []organizationcategoryv1.OrganizationCategory {
	if x != nil {
		return x.Categories
	}
	return nil
}

func (x *Organization) GetAddress() *Address {
	if x != nil {
		return x.Address
	}
	return nil
}

func (x *Organization) GetDomesticAchRoutingNumber() string {
	if x != nil {
		return x.DomesticAchRoutingNumber
	}
	return ""
}

func (x *Organization) GetInternationalAchRoutingNumber() string {
	if x != nil {
		return x.InternationalAchRoutingNumber
	}
	return ""
}

func (x *Organization) GetWireRoutingNumber() string {
	if x != nil {
		return x.WireRoutingNumber
	}
	return ""
}

func (x *Organization) GetSwiftRoutingNumber() string {
	if x != nil {
		return x.SwiftRoutingNumber
	}
	return ""
}

func (x *Organization) GetBankAccountNumber() string {
	if x != nil {
		return x.BankAccountNumber
	}
	return ""
}

type LegacyConfiguration struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	InterchangePercentage *wrapperspb.StringValue `protobuf:"bytes,1,opt,name=interchange_percentage,json=interchangePercentage,proto3" json:"interchange_percentage,omitempty"`
	DisputeRatePercentage *wrapperspb.StringValue `protobuf:"bytes,2,opt,name=dispute_rate_percentage,json=disputeRatePercentage,proto3" json:"dispute_rate_percentage,omitempty"`
	Currency              currencyv1.Currency     `protobuf:"varint,3,opt,name=currency,proto3,enum=micashared.common.enums.currency.v1.Currency" json:"currency,omitempty"`
	DisputeFee            *wrapperspb.StringValue `protobuf:"bytes,4,opt,name=dispute_fee,json=disputeFee,proto3" json:"dispute_fee,omitempty"`
	DisputeManagementCost *wrapperspb.StringValue `protobuf:"bytes,5,opt,name=dispute_management_cost,json=disputeManagementCost,proto3" json:"dispute_management_cost,omitempty"`
	NetworkMembershipFee  *wrapperspb.StringValue `protobuf:"bytes,6,opt,name=network_membership_fee,json=networkMembershipFee,proto3" json:"network_membership_fee,omitempty"`
	NetworkTransactionFee *wrapperspb.StringValue `protobuf:"bytes,7,opt,name=network_transaction_fee,json=networkTransactionFee,proto3" json:"network_transaction_fee,omitempty"`
	NetworkReportingFee   *wrapperspb.StringValue `protobuf:"bytes,8,opt,name=network_reporting_fee,json=networkReportingFee,proto3" json:"network_reporting_fee,omitempty"`
}

func (x *LegacyConfiguration) Reset() {
	*x = LegacyConfiguration{}
	if protoimpl.UnsafeEnabled {
		mi := &file_micashared_common_v1_organization_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *LegacyConfiguration) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*LegacyConfiguration) ProtoMessage() {}

func (x *LegacyConfiguration) ProtoReflect() protoreflect.Message {
	mi := &file_micashared_common_v1_organization_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use LegacyConfiguration.ProtoReflect.Descriptor instead.
func (*LegacyConfiguration) Descriptor() ([]byte, []int) {
	return file_micashared_common_v1_organization_proto_rawDescGZIP(), []int{1}
}

func (x *LegacyConfiguration) GetInterchangePercentage() *wrapperspb.StringValue {
	if x != nil {
		return x.InterchangePercentage
	}
	return nil
}

func (x *LegacyConfiguration) GetDisputeRatePercentage() *wrapperspb.StringValue {
	if x != nil {
		return x.DisputeRatePercentage
	}
	return nil
}

func (x *LegacyConfiguration) GetCurrency() currencyv1.Currency {
	if x != nil {
		return x.Currency
	}
	return currencyv1.Currency(0)
}

func (x *LegacyConfiguration) GetDisputeFee() *wrapperspb.StringValue {
	if x != nil {
		return x.DisputeFee
	}
	return nil
}

func (x *LegacyConfiguration) GetDisputeManagementCost() *wrapperspb.StringValue {
	if x != nil {
		return x.DisputeManagementCost
	}
	return nil
}

func (x *LegacyConfiguration) GetNetworkMembershipFee() *wrapperspb.StringValue {
	if x != nil {
		return x.NetworkMembershipFee
	}
	return nil
}

func (x *LegacyConfiguration) GetNetworkTransactionFee() *wrapperspb.StringValue {
	if x != nil {
		return x.NetworkTransactionFee
	}
	return nil
}

func (x *LegacyConfiguration) GetNetworkReportingFee() *wrapperspb.StringValue {
	if x != nil {
		return x.NetworkReportingFee
	}
	return nil
}

var File_micashared_common_v1_organization_proto protoreflect.FileDescriptor

var file_micashared_common_v1_organization_proto_rawDesc = []byte{
	0x0a, 0x27, 0x6d, 0x69, 0x63, 0x61, 0x73, 0x68, 0x61, 0x72, 0x65, 0x64, 0x2f, 0x63, 0x6f, 0x6d,
	0x6d, 0x6f, 0x6e, 0x2f, 0x76, 0x31, 0x2f, 0x6f, 0x72, 0x67, 0x61, 0x6e, 0x69, 0x7a, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x14, 0x6d, 0x69, 0x63, 0x61, 0x73,
	0x68, 0x61, 0x72, 0x65, 0x64, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x76, 0x31, 0x1a,
	0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66,
	0x2f, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x1a, 0x1e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75,
	0x66, 0x2f, 0x77, 0x72, 0x61, 0x70, 0x70, 0x65, 0x72, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x1a, 0x32, 0x6d, 0x69, 0x63, 0x61, 0x73, 0x68, 0x61, 0x72, 0x65, 0x64, 0x2f, 0x63, 0x6f, 0x6d,
	0x6d, 0x6f, 0x6e, 0x2f, 0x65, 0x6e, 0x75, 0x6d, 0x73, 0x2f, 0x63, 0x75, 0x72, 0x72, 0x65, 0x6e,
	0x63, 0x79, 0x2f, 0x76, 0x31, 0x2f, 0x63, 0x75, 0x72, 0x72, 0x65, 0x6e, 0x63, 0x79, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x4b, 0x6d, 0x69, 0x63, 0x61, 0x73, 0x68, 0x61, 0x72, 0x65, 0x64,
	0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2f, 0x65, 0x6e, 0x75, 0x6d, 0x73, 0x2f, 0x6f, 0x72,
	0x67, 0x61, 0x6e, 0x69, 0x7a, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x63, 0x61, 0x74, 0x65, 0x67, 0x6f,
	0x72, 0x79, 0x2f, 0x76, 0x31, 0x2f, 0x6f, 0x72, 0x67, 0x61, 0x6e, 0x69, 0x7a, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x5f, 0x63, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x1a, 0x22, 0x6d, 0x69, 0x63, 0x61, 0x73, 0x68, 0x61, 0x72, 0x65, 0x64, 0x2f, 0x63, 0x6f,
	0x6d, 0x6d, 0x6f, 0x6e, 0x2f, 0x76, 0x31, 0x2f, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x17, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x65, 0x2f,
	0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x9a,
	0x06, 0x0a, 0x0c, 0x4f, 0x72, 0x67, 0x61, 0x6e, 0x69, 0x7a, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12,
	0x34, 0x0a, 0x10, 0x6f, 0x72, 0x67, 0x61, 0x6e, 0x69, 0x7a, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x5f,
	0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x42, 0x09, 0xfa, 0x42, 0x06, 0x72, 0x04,
	0x10, 0x1e, 0x18, 0x32, 0x52, 0x0f, 0x6f, 0x72, 0x67, 0x61, 0x6e, 0x69, 0x7a, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x4b, 0x65, 0x79, 0x12, 0x18, 0x0a, 0x07, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x07, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x12,
	0x34, 0x0a, 0x07, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62,
	0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x07, 0x63, 0x72,
	0x65, 0x61, 0x74, 0x65, 0x64, 0x12, 0x34, 0x0a, 0x07, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64,
	0x18, 0x04, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61,
	0x6d, 0x70, 0x52, 0x07, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x12, 0x2a, 0x0a, 0x0b, 0x70,
	0x61, 0x72, 0x74, 0x6e, 0x65, 0x72, 0x5f, 0x6b, 0x65, 0x79, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09,
	0x42, 0x09, 0xfa, 0x42, 0x06, 0x72, 0x04, 0x10, 0x1e, 0x18, 0x32, 0x52, 0x0a, 0x70, 0x61, 0x72,
	0x74, 0x6e, 0x65, 0x72, 0x4b, 0x65, 0x79, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18,
	0x07, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x65, 0x0a, 0x0a, 0x63,
	0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x69, 0x65, 0x73, 0x18, 0x06, 0x20, 0x03, 0x28, 0x0e, 0x32,
	0x45, 0x2e, 0x6d, 0x69, 0x63, 0x61, 0x73, 0x68, 0x61, 0x72, 0x65, 0x64, 0x2e, 0x63, 0x6f, 0x6d,
	0x6d, 0x6f, 0x6e, 0x2e, 0x65, 0x6e, 0x75, 0x6d, 0x73, 0x2e, 0x6f, 0x72, 0x67, 0x61, 0x6e, 0x69,
	0x7a, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x63, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x79, 0x2e, 0x76,
	0x31, 0x2e, 0x4f, 0x72, 0x67, 0x61, 0x6e, 0x69, 0x7a, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x43, 0x61,
	0x74, 0x65, 0x67, 0x6f, 0x72, 0x79, 0x52, 0x0a, 0x63, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x69,
	0x65, 0x73, 0x12, 0x37, 0x0a, 0x07, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x18, 0x09, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x1d, 0x2e, 0x6d, 0x69, 0x63, 0x61, 0x73, 0x68, 0x61, 0x72, 0x65, 0x64,
	0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x76, 0x31, 0x2e, 0x41, 0x64, 0x64, 0x72, 0x65,
	0x73, 0x73, 0x52, 0x07, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x12, 0x4e, 0x0a, 0x1b, 0x64,
	0x6f, 0x6d, 0x65, 0x73, 0x74, 0x69, 0x63, 0x5f, 0x61, 0x63, 0x68, 0x5f, 0x72, 0x6f, 0x75, 0x74,
	0x69, 0x6e, 0x67, 0x5f, 0x6e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x09,
	0x42, 0x0f, 0xfa, 0x42, 0x0c, 0x72, 0x0a, 0x32, 0x08, 0x5e, 0x7c, 0x5c, 0x77, 0x7b, 0x39, 0x7d,
	0x24, 0x52, 0x18, 0x64, 0x6f, 0x6d, 0x65, 0x73, 0x74, 0x69, 0x63, 0x41, 0x63, 0x68, 0x52, 0x6f,
	0x75, 0x74, 0x69, 0x6e, 0x67, 0x4e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x12, 0x58, 0x0a, 0x20, 0x69,
	0x6e, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x61, 0x6c, 0x5f, 0x61, 0x63, 0x68,
	0x5f, 0x72, 0x6f, 0x75, 0x74, 0x69, 0x6e, 0x67, 0x5f, 0x6e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x18,
	0x0b, 0x20, 0x01, 0x28, 0x09, 0x42, 0x0f, 0xfa, 0x42, 0x0c, 0x72, 0x0a, 0x32, 0x08, 0x5e, 0x7c,
	0x5c, 0x77, 0x7b, 0x39, 0x7d, 0x24, 0x52, 0x1d, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x61, 0x6c, 0x41, 0x63, 0x68, 0x52, 0x6f, 0x75, 0x74, 0x69, 0x6e, 0x67, 0x4e,
	0x75, 0x6d, 0x62, 0x65, 0x72, 0x12, 0x3f, 0x0a, 0x13, 0x77, 0x69, 0x72, 0x65, 0x5f, 0x72, 0x6f,
	0x75, 0x74, 0x69, 0x6e, 0x67, 0x5f, 0x6e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x18, 0x0c, 0x20, 0x01,
	0x28, 0x09, 0x42, 0x0f, 0xfa, 0x42, 0x0c, 0x72, 0x0a, 0x32, 0x08, 0x5e, 0x7c, 0x5c, 0x77, 0x7b,
	0x39, 0x7d, 0x24, 0x52, 0x11, 0x77, 0x69, 0x72, 0x65, 0x52, 0x6f, 0x75, 0x74, 0x69, 0x6e, 0x67,
	0x4e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x12, 0x41, 0x0a, 0x14, 0x73, 0x77, 0x69, 0x66, 0x74, 0x5f,
	0x72, 0x6f, 0x75, 0x74, 0x69, 0x6e, 0x67, 0x5f, 0x6e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x18, 0x0d,
	0x20, 0x01, 0x28, 0x09, 0x42, 0x0f, 0xfa, 0x42, 0x0c, 0x72, 0x0a, 0x32, 0x08, 0x5e, 0x7c, 0x5c,
	0x77, 0x7b, 0x39, 0x7d, 0x24, 0x52, 0x12, 0x73, 0x77, 0x69, 0x66, 0x74, 0x52, 0x6f, 0x75, 0x74,
	0x69, 0x6e, 0x67, 0x4e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x12, 0x40, 0x0a, 0x13, 0x62, 0x61, 0x6e,
	0x6b, 0x5f, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x5f, 0x6e, 0x75, 0x6d, 0x62, 0x65, 0x72,
	0x18, 0x0e, 0x20, 0x01, 0x28, 0x09, 0x42, 0x10, 0xfa, 0x42, 0x0d, 0x72, 0x0b, 0x32, 0x09, 0x5e,
	0x7c, 0x5c, 0x77, 0x7b, 0x31, 0x34, 0x7d, 0x24, 0x52, 0x11, 0x62, 0x61, 0x6e, 0x6b, 0x41, 0x63,
	0x63, 0x6f, 0x75, 0x6e, 0x74, 0x4e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x22, 0x9c, 0x05, 0x0a, 0x13,
	0x4c, 0x65, 0x67, 0x61, 0x63, 0x79, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x75, 0x72, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x12, 0x53, 0x0a, 0x16, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x63, 0x68, 0x61, 0x6e,
	0x67, 0x65, 0x5f, 0x70, 0x65, 0x72, 0x63, 0x65, 0x6e, 0x74, 0x61, 0x67, 0x65, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x1c, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x53, 0x74, 0x72, 0x69, 0x6e, 0x67, 0x56, 0x61, 0x6c, 0x75,
	0x65, 0x52, 0x15, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x63, 0x68, 0x61, 0x6e, 0x67, 0x65, 0x50, 0x65,
	0x72, 0x63, 0x65, 0x6e, 0x74, 0x61, 0x67, 0x65, 0x12, 0x54, 0x0a, 0x17, 0x64, 0x69, 0x73, 0x70,
	0x75, 0x74, 0x65, 0x5f, 0x72, 0x61, 0x74, 0x65, 0x5f, 0x70, 0x65, 0x72, 0x63, 0x65, 0x6e, 0x74,
	0x61, 0x67, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1c, 0x2e, 0x67, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x53, 0x74, 0x72, 0x69,
	0x6e, 0x67, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x52, 0x15, 0x64, 0x69, 0x73, 0x70, 0x75, 0x74, 0x65,
	0x52, 0x61, 0x74, 0x65, 0x50, 0x65, 0x72, 0x63, 0x65, 0x6e, 0x74, 0x61, 0x67, 0x65, 0x12, 0x49,
	0x0a, 0x08, 0x63, 0x75, 0x72, 0x72, 0x65, 0x6e, 0x63, 0x79, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0e,
	0x32, 0x2d, 0x2e, 0x6d, 0x69, 0x63, 0x61, 0x73, 0x68, 0x61, 0x72, 0x65, 0x64, 0x2e, 0x63, 0x6f,
	0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x65, 0x6e, 0x75, 0x6d, 0x73, 0x2e, 0x63, 0x75, 0x72, 0x72, 0x65,
	0x6e, 0x63, 0x79, 0x2e, 0x76, 0x31, 0x2e, 0x43, 0x75, 0x72, 0x72, 0x65, 0x6e, 0x63, 0x79, 0x52,
	0x08, 0x63, 0x75, 0x72, 0x72, 0x65, 0x6e, 0x63, 0x79, 0x12, 0x3d, 0x0a, 0x0b, 0x64, 0x69, 0x73,
	0x70, 0x75, 0x74, 0x65, 0x5f, 0x66, 0x65, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1c,
	0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66,
	0x2e, 0x53, 0x74, 0x72, 0x69, 0x6e, 0x67, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x52, 0x0a, 0x64, 0x69,
	0x73, 0x70, 0x75, 0x74, 0x65, 0x46, 0x65, 0x65, 0x12, 0x54, 0x0a, 0x17, 0x64, 0x69, 0x73, 0x70,
	0x75, 0x74, 0x65, 0x5f, 0x6d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x5f, 0x63,
	0x6f, 0x73, 0x74, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1c, 0x2e, 0x67, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x53, 0x74, 0x72, 0x69,
	0x6e, 0x67, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x52, 0x15, 0x64, 0x69, 0x73, 0x70, 0x75, 0x74, 0x65,
	0x4d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x43, 0x6f, 0x73, 0x74, 0x12, 0x52,
	0x0a, 0x16, 0x6e, 0x65, 0x74, 0x77, 0x6f, 0x72, 0x6b, 0x5f, 0x6d, 0x65, 0x6d, 0x62, 0x65, 0x72,
	0x73, 0x68, 0x69, 0x70, 0x5f, 0x66, 0x65, 0x65, 0x18, 0x06, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1c,
	0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66,
	0x2e, 0x53, 0x74, 0x72, 0x69, 0x6e, 0x67, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x52, 0x14, 0x6e, 0x65,
	0x74, 0x77, 0x6f, 0x72, 0x6b, 0x4d, 0x65, 0x6d, 0x62, 0x65, 0x72, 0x73, 0x68, 0x69, 0x70, 0x46,
	0x65, 0x65, 0x12, 0x54, 0x0a, 0x17, 0x6e, 0x65, 0x74, 0x77, 0x6f, 0x72, 0x6b, 0x5f, 0x74, 0x72,
	0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x66, 0x65, 0x65, 0x18, 0x07, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x1c, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x53, 0x74, 0x72, 0x69, 0x6e, 0x67, 0x56, 0x61, 0x6c, 0x75,
	0x65, 0x52, 0x15, 0x6e, 0x65, 0x74, 0x77, 0x6f, 0x72, 0x6b, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x61,
	0x63, 0x74, 0x69, 0x6f, 0x6e, 0x46, 0x65, 0x65, 0x12, 0x50, 0x0a, 0x15, 0x6e, 0x65, 0x74, 0x77,
	0x6f, 0x72, 0x6b, 0x5f, 0x72, 0x65, 0x70, 0x6f, 0x72, 0x74, 0x69, 0x6e, 0x67, 0x5f, 0x66, 0x65,
	0x65, 0x18, 0x08, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1c, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x53, 0x74, 0x72, 0x69, 0x6e, 0x67,
	0x56, 0x61, 0x6c, 0x75, 0x65, 0x52, 0x13, 0x6e, 0x65, 0x74, 0x77, 0x6f, 0x72, 0x6b, 0x52, 0x65,
	0x70, 0x6f, 0x72, 0x74, 0x69, 0x6e, 0x67, 0x46, 0x65, 0x65, 0x42, 0x52, 0x0a, 0x17, 0x69, 0x6f,
	0x2e, 0x6d, 0x69, 0x63, 0x61, 0x73, 0x68, 0x61, 0x72, 0x65, 0x64, 0x2e, 0x63, 0x6f, 0x6d, 0x6d,
	0x6f, 0x6e, 0x2e, 0x76, 0x31, 0x42, 0x11, 0x4f, 0x72, 0x67, 0x61, 0x6e, 0x69, 0x7a, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x5a, 0x1d, 0x6d, 0x69, 0x63, 0x61, 0x73, 0x68,
	0x61, 0x72, 0x65, 0x64, 0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2f, 0x76, 0x31, 0x3b, 0x63,
	0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x76, 0x31, 0xa2, 0x02, 0x04, 0x4d, 0x49, 0x43, 0x41, 0x62, 0x06,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_micashared_common_v1_organization_proto_rawDescOnce sync.Once
	file_micashared_common_v1_organization_proto_rawDescData = file_micashared_common_v1_organization_proto_rawDesc
)

func file_micashared_common_v1_organization_proto_rawDescGZIP() []byte {
	file_micashared_common_v1_organization_proto_rawDescOnce.Do(func() {
		file_micashared_common_v1_organization_proto_rawDescData = protoimpl.X.CompressGZIP(file_micashared_common_v1_organization_proto_rawDescData)
	})
	return file_micashared_common_v1_organization_proto_rawDescData
}

var file_micashared_common_v1_organization_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_micashared_common_v1_organization_proto_goTypes = []interface{}{
	(*Organization)(nil),                             // 0: micashared.common.v1.Organization
	(*LegacyConfiguration)(nil),                      // 1: micashared.common.v1.LegacyConfiguration
	(*timestamppb.Timestamp)(nil),                    // 2: google.protobuf.Timestamp
	(organizationcategoryv1.OrganizationCategory)(0), // 3: micashared.common.enums.organizationcategory.v1.OrganizationCategory
	(*Address)(nil),                                  // 4: micashared.common.v1.Address
	(*wrapperspb.StringValue)(nil),                   // 5: google.protobuf.StringValue
	(currencyv1.Currency)(0),                         // 6: micashared.common.enums.currency.v1.Currency
}
var file_micashared_common_v1_organization_proto_depIdxs = []int32{
	2,  // 0: micashared.common.v1.Organization.created:type_name -> google.protobuf.Timestamp
	2,  // 1: micashared.common.v1.Organization.updated:type_name -> google.protobuf.Timestamp
	3,  // 2: micashared.common.v1.Organization.categories:type_name -> micashared.common.enums.organizationcategory.v1.OrganizationCategory
	4,  // 3: micashared.common.v1.Organization.address:type_name -> micashared.common.v1.Address
	5,  // 4: micashared.common.v1.LegacyConfiguration.interchange_percentage:type_name -> google.protobuf.StringValue
	5,  // 5: micashared.common.v1.LegacyConfiguration.dispute_rate_percentage:type_name -> google.protobuf.StringValue
	6,  // 6: micashared.common.v1.LegacyConfiguration.currency:type_name -> micashared.common.enums.currency.v1.Currency
	5,  // 7: micashared.common.v1.LegacyConfiguration.dispute_fee:type_name -> google.protobuf.StringValue
	5,  // 8: micashared.common.v1.LegacyConfiguration.dispute_management_cost:type_name -> google.protobuf.StringValue
	5,  // 9: micashared.common.v1.LegacyConfiguration.network_membership_fee:type_name -> google.protobuf.StringValue
	5,  // 10: micashared.common.v1.LegacyConfiguration.network_transaction_fee:type_name -> google.protobuf.StringValue
	5,  // 11: micashared.common.v1.LegacyConfiguration.network_reporting_fee:type_name -> google.protobuf.StringValue
	12, // [12:12] is the sub-list for method output_type
	12, // [12:12] is the sub-list for method input_type
	12, // [12:12] is the sub-list for extension type_name
	12, // [12:12] is the sub-list for extension extendee
	0,  // [0:12] is the sub-list for field type_name
}

func init() { file_micashared_common_v1_organization_proto_init() }
func file_micashared_common_v1_organization_proto_init() {
	if File_micashared_common_v1_organization_proto != nil {
		return
	}
	file_micashared_common_v1_address_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_micashared_common_v1_organization_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Organization); i {
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
		file_micashared_common_v1_organization_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*LegacyConfiguration); i {
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
			RawDescriptor: file_micashared_common_v1_organization_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_micashared_common_v1_organization_proto_goTypes,
		DependencyIndexes: file_micashared_common_v1_organization_proto_depIdxs,
		MessageInfos:      file_micashared_common_v1_organization_proto_msgTypes,
	}.Build()
	File_micashared_common_v1_organization_proto = out.File
	file_micashared_common_v1_organization_proto_rawDesc = nil
	file_micashared_common_v1_organization_proto_goTypes = nil
	file_micashared_common_v1_organization_proto_depIdxs = nil
}
